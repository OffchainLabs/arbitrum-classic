package dev

import (
	"bytes"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
)

func TestL2ToL1Tx(t *testing.T) {
	tmpDir, err := ioutil.TempDir(".", "arbitrum")
	if err != nil {
		logger.Fatal().Err(err).Msg("error generating temporary directory")
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()

	config := protocol.ChainParams{
		StakeRequirement:          big.NewInt(10),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocksInt(3),
		MaxExecutionSteps:         10000000000,
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	monitor, backend, db, rollupAddress := NewDevNode(tmpDir, config)
	defer monitor.Close()

	srv := aggregator.NewServer(backend, rollupAddress, db)
	client := web3.NewEthClient(srv, true)
	arbSys, err := arboscontracts.NewArbSys(arbos.ARB_SYS_ADDRESS, client)
	if err != nil {
		t.Fatal(err)
	}
	privkey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	auth := bind.NewKeyedTransactor(privkey)

	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(auth.From),
				Payment:     big.NewInt(100),
				Data:        nil,
			},
		}),
	}
	if err := backend.AddInboxMessage(deposit, common.RandAddress()); err != nil {
		logger.Fatal().Stack().Err(err).Send()
	}

	for i := 0; i < 12; i++ {
		dest := common.RandAddress().ToEthAddress()
		t.Log("Send tx to L1", dest.Hex())
		_, err = arbSys.SendTxToL1(auth, dest, nil)
		if err != nil {
			t.Fatal(err)
		}
		if i%8 == 0 {
			backend.l1Emulator.IncreaseTime(10)
		}
	}

	batches := make([]*evm.MerkleRootResult, 0)
	for i := 0; i < 3; i++ {
		batch, err := db.GetMessageBatch(big.NewInt(int64(i)))
		if err != nil {
			t.Fatal(err)
		}
		if batch.BatchNumber.Cmp(big.NewInt(int64(i))) != 0 {
			t.Fatal("wrong batch num")
		}
		batches = append(batches, batch)
	}

	clnt, pks := test.SimulatedBackend()
	ethAuth := bind.NewKeyedTransactor(pks[0])

	bridgeAddress, _, bridge, err := ethbridgecontracts.DeployBridge(ethAuth, clnt)
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()

	outboxAddress, _, outbox, err := ethbridgecontracts.DeployOutbox(ethAuth, clnt, ethAuth.From, bridgeAddress)
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()

	inboxAddress, _, inbox, err := ethbridgecontracts.DeployInbox(ethAuth, clnt, bridgeAddress)
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()

	_, err = bridge.SetOutbox(ethAuth, outboxAddress, true)
	if err != nil {
		t.Fatal(err)
	}
	_, err = bridge.SetInbox(ethAuth, inboxAddress, true)
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()

	_, err = inbox.DepositEth(&bind.TransactOpts{
		From:     ethAuth.From,
		Nonce:    ethAuth.Nonce,
		Signer:   ethAuth.Signer,
		Value:    big.NewInt(100000000),
		GasPrice: nil,
		GasLimit: 0,
		Context:  nil,
	}, common.RandAddress().ToEthAddress())
	if err != nil {
		t.Fatal(err)
	}

	sendCount, err := backend.arbcore.GetSendCount()
	if err != nil {
		t.Fatal(err)
	}
	sends, err := backend.arbcore.GetSends(big.NewInt(0), sendCount)
	if err != nil {
		t.Fatal(err)
	}

	var sendsData []byte
	var sendLengths []*big.Int
	for _, send := range sends {
		sendsData = append(sendsData, send...)
		sendLengths = append(sendLengths, big.NewInt(int64(len(send))))
	}
	_, err = outbox.ProcessOutgoingMessages(ethAuth, sendsData, sendLengths)
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()

	nodeInterface, err := arboscontracts.NewNodeInterface(arbos.ARB_NODE_INTERFACE_ADDRESS, client)
	if err != nil {
		t.Fatal(err)
	}

	for i, batch := range batches {
		entries := batch.Tree.Entries()
		for j, entry := range entries {
			send, err := evm.NewVirtualSendResultFromData(entry)
			if err != nil {
				t.Fatal(err)
			}
			res, ok := send.(*evm.L2ToL1TxResult)
			if !ok {
				t.Fatal("expected l2 to l1 result")
			}
			batchNum := big.NewInt(int64(i))
			index := uint64(j)
			msgData, err := nodeInterface.LookupMessageBatchProof(&bind.CallOpts{}, batchNum, index)
			if err != nil {
				t.Fatal(err)
			}
			if msgData.L2Sender != res.L2Sender.ToEthAddress() {
				t.Fatal("wrong l2 sender")
			}
			if msgData.L1Dest != res.L1Dest.ToEthAddress() {
				t.Fatal("wrong l1 dest")
			}
			if msgData.L2Block.Cmp(res.L2Block) != 0 {
				t.Fatal("wrong l2 block")
			}
			if msgData.L1Block.Cmp(res.L1Block) != 0 {
				t.Fatal("wrong l1 block")
			}
			if msgData.Timestamp.Cmp(res.Timestamp) != 0 {
				t.Fatal("wrong timestamp")
			}
			if msgData.Amount.Cmp(res.Value) != 0 {
				t.Fatal("wrong amount")
			}
			if !bytes.Equal(msgData.CalldataForL1, res.Calldata) {
				t.Fatal("wrong calldata")
			}
			t.Log("Execute", msgData.L1Dest.Hex(), msgData.Amount, hexutil.Encode(msgData.CalldataForL1))

			_, err = outbox.ExecuteTransaction(
				ethAuth,
				batchNum,
				msgData.Proof,
				msgData.Path,
				msgData.L2Sender,
				msgData.L1Dest,
				msgData.L2Block,
				msgData.L1Block,
				msgData.Timestamp,
				msgData.Amount,
				msgData.CalldataForL1,
			)
			if err != nil {
				t.Fatal(err)
			}
			clnt.Commit()
		}
	}
}

func TestBridge(t *testing.T) {
	clnt, pks := test.SimulatedBackend()
	auth := bind.NewKeyedTransactor(pks[0])

	_, _, bridge, err := ethbridgecontracts.DeployBridge(auth, clnt)
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()

	_, err = bridge.SetOutbox(auth, auth.From, true)
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()

	_, err = bridge.ExecuteCall(auth, common.RandAddress().ToEthAddress(), big.NewInt(0), nil)
	if err != nil {
		t.Fatal(err)
	}
	clnt.Commit()
}
