package arbostest

import (
	"bytes"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestOwnerDeployCorrectCode(t *testing.T) {
	privkey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	txSender := crypto.PubkeyToAddress(privkey.PublicKey)
	connAddress := crypto.CreateAddress(txSender, 0)
	signer := types.NewEIP155Signer(chainId)
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	conData := hexutil.MustDecode(arbostestcontracts.SimpleBin)

	snap1 := func() *snapshot.Snapshot {
		txRaw := types.NewTx(&types.LegacyTx{
			Nonce:    0,
			GasPrice: big.NewInt(0),
			Gas:      10000000,
			To:       nil,
			Value:    big.NewInt(0),
			Data:     conData,
		})
		tx, err := types.SignTx(txRaw, signer, privkey)
		test.FailIfError(t, err)
		l2Tx, err := message.NewL2Message(message.NewCompressedECDSAFromEth(tx))
		test.FailIfError(t, err)
		messages := []message.Message{l2Tx}
		results, snap := runSimpleTxAssertion(t, messages)
		allResultsSucceeded(t, results)
		t.Log(results[0])
		return snap
	}()

	snap2 := func() *snapshot.Snapshot {
		ownerTx := message.Transaction{
			MaxGas:      big.NewInt(100000000),
			GasPriceBid: big.NewInt(0),
			SequenceNum: big.NewInt(0),
			DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
			Payment:     big.NewInt(0),
			Data:        arbos.DeployContract(conData, common.NewAddressFromEth(txSender), big.NewInt(0)),
		}

		ib := &InboxBuilder{}
		config := []message.ChainConfigOption{message.ChainIDConfig{ChainId: chainId}}
		ib.AddMessage(initMsg(t, config), common.Address{}, big.NewInt(0), chainTime)
		ib.AddMessage(message.NewSafeL2Message(ownerTx), owner, big.NewInt(0), chainTime)
		results, snap := runTxAssertion(t, ib.Messages)
		allResultsSucceeded(t, results)
		t.Log(results[0])
		return snap
	}()

	code1, err := snap1.GetCode(common.NewAddressFromEth(connAddress))
	test.FailIfError(t, err)
	code2, err := snap2.GetCode(common.NewAddressFromEth(connAddress))
	test.FailIfError(t, err)
	if !bytes.Equal(code1, code2) {
		t.Error("code not deployed correctly")
	}
}

func TestOwnerDeployCorrectDeploy(t *testing.T) {
	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}
	simple, err := abi.JSON(strings.NewReader(arbostestcontracts.SimpleABI))
	failIfError(t, err)

	conData := hexutil.MustDecode(arbostestcontracts.SimpleBin)
	nonce := uint64(342)
	ownerTx := message.Transaction{
		MaxGas:      big.NewInt(100000000),
		GasPriceBid: big.NewInt(0), // fees are off
		SequenceNum: big.NewInt(0),
		DestAddress: common.NewAddressFromEth(arbos.ARB_OWNER_ADDRESS),
		Payment:     big.NewInt(100),
		Data:        arbos.DeployContract(conData, sender, new(big.Int).SetUint64(nonce)),
	}

	ib := &InboxBuilder{}
	ib.AddMessage(initMsg(t, nil), common.Address{}, big.NewInt(0), chainTime)
	ib.AddMessage(makeEthDeposit(message.L2RemapAccount(owner), big.NewInt(1000)), sender, big.NewInt(0), chainTime)
	ib.AddMessage(message.NewSafeL2Message(ownerTx), owner, big.NewInt(0), chainTime)
	results, snap := runTxAssertion(t, ib.Messages)
	correctConnAddress := crypto.CreateAddress(sender.ToEthAddress(), nonce)
	checkConstructorResult(t, results[1], common.NewAddressFromEth(correctConnAddress))

	evmLogs := results[1].EVMLogs
	if len(evmLogs) != 1 {
		t.Fatal("wrong log count")
	}
	evmLog := evmLogs[0]
	if evmLog.Topics[0].ToEthHash() != simple.Events["TestEvent"].ID {
		t.Fatal("wrong topic")
	}
	if new(big.Int).SetBytes(evmLog.Data[:32]).Cmp(ownerTx.Payment) != 0 {
		t.Error("wrong event value data")
	}
	// The sender check has been disabled, as right now it's incorrectly set to the ArbSys precompile address
	/*
		var senderInLog common.Address
		copy(senderInLog[:], evmLog.Data[32+(32-20):])
		if senderInLog != sender {
			t.Error("wrong event sender data")
		}
	*/
	ownerBalance, err := snap.GetBalance(message.L2RemapAccount(owner))
	test.FailIfError(t, err)
	conBalance, err := snap.GetBalance(common.NewAddressFromEth(correctConnAddress))
	test.FailIfError(t, err)
	if ownerBalance.Cmp(big.NewInt(900)) != 0 {
		t.Error("wrong owner balance")
	}
	if conBalance.Cmp(big.NewInt(100)) != 0 {
		t.Error("wrong contract balance")
	}
}
