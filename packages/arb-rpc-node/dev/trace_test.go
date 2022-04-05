package dev

import (
	"bytes"
	"context"
	"encoding/json"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/test"
)

func TestTrace(t *testing.T) {
	skipBelowVersion(t, 25)
	ctx := context.Background()
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	senderKey := test.MustGenerateKey(t)
	ownerKey := test.MustGenerateKey(t)
	owner := crypto.PubkeyToAddress(ownerKey.PublicKey)

	backend, _, srv, cancelDevNode := NewSimpleTestDevNode(t, config, common.NewAddressFromEth(owner))
	defer cancelDevNode()

	senderAuth, err := bind.NewKeyedTransactorWithChainID(senderKey, backend.chainID)
	test.FailIfError(t, err)

	ethServer := web3.NewServer(srv, web3.DefaultConfig, nil)
	tracer := web3.NewTracer(ethServer, configuration.DefaultCoreSettingsMaxExecution())

	client := web3.NewEthClient(srv, true)

	simpleAddr, _, _, err := arbostestcontracts.DeploySimple(senderAuth, client)
	test.FailIfError(t, err)
	simpleAddr2, _, _, err := arbostestcontracts.DeploySimple(senderAuth, client)
	test.FailIfError(t, err)

	simpleABI, err := arbostestcontracts.SimpleMetaData.GetAbi()
	test.FailIfError(t, err)

	traceMethod := simpleABI.Methods["trace"]
	traceInpData, err := traceMethod.Inputs.Pack(big.NewInt(4234))
	test.FailIfError(t, err)
	gas := hexutil.Uint64(100000000)
	blockNum := rpc.LatestBlockNumber
	data := hexutil.Bytes(append(traceMethod.ID, traceInpData...))
	callTraceData, err := tracer.Call(ctx, web3.CallTxArgs{
		From: &senderAuth.From,
		To:   &simpleAddr,
		Data: &data,
		Gas:  &gas,
	}, []string{"trace", "deletedContracts"}, rpc.BlockNumberOrHash{BlockNumber: &blockNum})
	test.FailIfError(t, err)

	signer := types.NewEIP155Signer(backend.chainID)
	userTx1 := types.NewTx(&types.LegacyTx{
		Nonce:    2,
		GasPrice: big.NewInt(10),
		Gas:      uint64(gas),
		To:       &simpleAddr,
		Value:    big.NewInt(0),
		Data:     data,
	})
	userTx1, err = types.SignTx(userTx1, signer, senderKey)
	test.FailIfError(t, err)

	userTx2 := types.NewTx(&types.LegacyTx{
		Nonce:    3,
		GasPrice: big.NewInt(10),
		Gas:      uint64(gas),
		To:       &simpleAddr2,
		Value:    big.NewInt(0),
		Data:     data,
	})
	userTx2, err = types.SignTx(userTx2, signer, senderKey)
	test.FailIfError(t, err)

	arbMsg, err := message.NewTransactionBatchFromMessages([]message.AbstractL2Message{
		message.NewCompressedECDSAFromEth(userTx1),
		message.NewCompressedECDSAFromEth(userTx2),
	})
	test.FailIfError(t, err)
	_, err = backend.AddInboxMessage(message.NewSafeL2Message(arbMsg), common.Address{})
	test.FailIfError(t, err)

	tx1TraceData, err := tracer.ReplayTransaction(ctx, userTx1.Hash().Bytes(), []string{"trace", "deletedContracts"})
	test.FailIfError(t, err)

	tx2TraceData, err := tracer.ReplayTransaction(ctx, userTx2.Hash().Bytes(), []string{"trace", "deletedContracts"})
	test.FailIfError(t, err)

	txReq, _, _, err := backend.db.GetRequest(common.NewHashFromEth(userTx1.Hash()))
	test.FailIfError(t, err)
	l2BlockNum := rpc.BlockNumber(txReq.IncomingRequest.L2BlockNumber.Int64())

	blockTraceData, err := tracer.ReplayBlockTransactions(ctx, rpc.BlockNumberOrHash{BlockNumber: &l2BlockNum}, []string{"trace", "deletedContracts"})
	test.FailIfError(t, err)

	for i := range blockTraceData {
		var txHash []byte
		if i == 0 {
			txHash = userTx1.Hash().Bytes()
		} else {
			txHash = userTx2.Hash().Bytes()
		}
		for j := range blockTraceData[i].Trace {
			if !bytes.Equal(*blockTraceData[i].Trace[j].TransactionHash, txHash) {
				t.Error("wrong tx hash")
			}
			blockTraceData[i].Trace[j].TransactionHash = nil
		}
	}

	assertTraceEqual(t, tx1TraceData, blockTraceData[0])
	assertTraceEqual(t, tx2TraceData, blockTraceData[1])

	// Gas is slightly different between call and transaction so we clear for comparison
	clearGasData(tx1TraceData)
	clearGasData(callTraceData)
	assertTraceEqual(t, callTraceData, tx1TraceData)
}

func clearGasData(trace *web3.TraceResult) {
	for j := range trace.Trace {
		if trace.Trace[j].Result != nil {
			trace.Trace[j].Result.GasUsed = 0
		}
		trace.Trace[j].Action.Gas = 0
	}
}

func assertTraceEqual(t *testing.T, trace1 interface{}, trace2 interface{}) {
	t.Helper()
	jsonData1, err := json.MarshalIndent(trace1, "", " ")
	test.FailIfError(t, err)
	jsonData2, err := json.MarshalIndent(trace2, "", " ")
	test.FailIfError(t, err)
	if !bytes.Equal(jsonData1, jsonData2) {
		t.Errorf("traces not equal")
	}
}

func TestTraceConstructor(t *testing.T) {
	skipBelowVersion(t, 25)
	ctx := context.Background()
	config := protocol.ChainParams{
		GracePeriod:               common.NewTimeBlocksInt(3),
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	senderKey := test.MustGenerateKey(t)
	ownerKey := test.MustGenerateKey(t)
	owner := crypto.PubkeyToAddress(ownerKey.PublicKey)

	backend, _, srv, cancelDevNode := NewTestDevNode(t, *arbosfile, config, common.NewAddressFromEth(owner), nil, false)
	defer cancelDevNode()

	signer := types.NewEIP155Signer(backend.chainID)

	senderAuth, err := bind.NewKeyedTransactorWithChainID(senderKey, backend.chainID)
	test.FailIfError(t, err)

	ethServer := web3.NewServer(srv, web3.DefaultConfig, nil)
	tracer := web3.NewTracer(ethServer, configuration.DefaultCoreSettingsMaxExecution())

	client := web3.NewEthClient(srv, true)

	_, successTx, _, err := arbostestcontracts.DeploySimple(senderAuth, client)
	test.FailIfError(t, err)

	senderAuth.GasLimit = 5000
	_, failedTx, _, err := arbostestcontracts.DeploySimple(senderAuth, client)
	test.FailIfError(t, err)

	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(senderAuth.From),
				Payment:     new(big.Int).Exp(big.NewInt(10), big.NewInt(20), nil),
				Data:        nil,
			},
		}),
	}
	_, err = backend.AddInboxMessage(deposit, message.L1RemapAccount(common.NewAddressFromEth(senderAuth.From)))
	test.FailIfError(t, err)

	retryableTx := message.RetryableTx{
		Destination:       common.Address{},
		Value:             big.NewInt(20),
		Deposit:           big.NewInt(100),
		MaxSubmissionCost: big.NewInt(30),
		CreditBack:        common.RandAddress(),
		Beneficiary:       common.RandAddress(),
		MaxGas:            big.NewInt(100000000),
		GasPriceBid:       big.NewInt(10),
		Data:              successTx.Data(),
	}

	retryableRequestId, err := backend.AddInboxMessage(retryableTx, message.L1RemapAccount(common.NewAddressFromEth(senderAuth.From)))
	test.FailIfError(t, err)

	redeemId := hashing.SoliditySHA3(hashing.Bytes32(retryableRequestId), hashing.Uint256(big.NewInt(1)))

	txTrace, err := tracer.Transaction(ctx, redeemId.Bytes())
	test.FailIfError(t, err)
	for _, frame := range txTrace {
		if frame.Type != "call" {
			t.Error("unexpected frame type")
		}
	}

	successCreateDeposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.Address{},
				Payment:     big.NewInt(0),
				Data:        successTx.Data(),
			},
		}),
	}

	successDepositRequestId, err := backend.AddInboxMessage(successCreateDeposit, message.L1RemapAccount(common.NewAddressFromEth(senderAuth.From)))
	test.FailIfError(t, err)

	failedCreateDeposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.Address{},
				Payment:     big.NewInt(0),
				Data:        successTx.Data(),
			},
		}),
	}

	failedDepositRequestId, err := backend.AddInboxMessage(failedCreateDeposit, message.L1RemapAccount(common.NewAddressFromEth(senderAuth.From)))
	test.FailIfError(t, err)

	checkCreateRequest := func(txHash []byte, data []byte, sender ethcommon.Address, nonce uint64, success bool) {
		t.Helper()
		txTrace, err := tracer.Transaction(ctx, txHash)
		test.FailIfError(t, err)

		if len(txTrace) == 0 {
			t.Fatal("expected at least one frame")
		}
		if txTrace[0].Type != "create" {
			t.Error("expected top level frame to be create")
		}

		deployedCode, err := client.CodeAt(ctx, crypto.CreateAddress(sender, nonce), nil)
		test.FailIfError(t, err)

		if success {
			if txTrace[0].Result == nil || txTrace[0].Result.Code == nil || !bytes.Equal(deployedCode, *txTrace[0].Result.Code) {
				t.Error("expected correct code in create")
			}
		} else {
			if txTrace[0].Result != nil {
				t.Error("expected no result")
			}
		}

		if !bytes.Equal(data, txTrace[0].Action.Init) {
			t.Log("Correct init", hexutil.Encode(data))
			t.Log("Real init", hexutil.Encode(txTrace[0].Action.Init))
			t.Error("expected correct init in create")
		}
	}

	checkCreateTx := func(tx *types.Transaction, success bool) {
		t.Helper()
		sender, err := types.Sender(signer, tx)
		test.FailIfError(t, err)
		checkCreateRequest(tx.Hash().Bytes(), tx.Data(), sender, tx.Nonce(), success)
	}

	checkCreateTx(successTx, true)
	checkCreateTx(failedTx, false)
	checkCreateRequest(successDepositRequestId.Bytes(), successTx.Data(), senderAuth.From, failedTx.Nonce()+1, true)
	checkCreateRequest(failedDepositRequestId.Bytes(), successTx.Data(), senderAuth.From, failedTx.Nonce()+2, false)
}
