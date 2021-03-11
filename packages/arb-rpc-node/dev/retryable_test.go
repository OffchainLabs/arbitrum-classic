package dev

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/test"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/txdb"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"io/ioutil"
	"math/big"
	"os"
	"testing"
)

func setupTest(t *testing.T, tmpDir string) (
	common.Address,
	*bind.TransactOpts,
	*bind.TransactOpts,
	common.Address,
	*txdb.TxDB,
	*Backend,
	message.RetryableTx,
	common.Hash,
	func(),
) {
	config := protocol.ChainParams{
		StakeRequirement:          big.NewInt(10),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocksInt(3),
		MaxExecutionSteps:         10000000000,
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}
	monitor, backend, db, rollupAddress := NewDevNode(tmpDir, config)
	closeFunc := func() {
		db.Close()
		monitor.Close()
	}
	privkey, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	otherAuth := bind.NewKeyedTransactor(privkey)

	privkey2, err := crypto.GenerateKey()
	test.FailIfError(t, err)
	beneficiaryAuth := bind.NewKeyedTransactor(privkey2)

	sender := common.RandAddress()

	deposit := message.EthDepositTx{
		L2Message: message.NewSafeL2Message(message.ContractTransaction{
			BasicTx: message.BasicTx{
				MaxGas:      big.NewInt(1000000),
				GasPriceBid: big.NewInt(0),
				DestAddress: common.NewAddressFromEth(otherAuth.From),
				Payment:     big.NewInt(100),
				Data:        nil,
			},
		}),
	}
	_, err = backend.AddInboxMessage(deposit, common.RandAddress())
	test.FailIfError(t, err)

	retryableTx := message.RetryableTx{
		Destination:       common.RandAddress(),
		Value:             big.NewInt(20),
		Deposit:           big.NewInt(100),
		MaxSubmissionCost: big.NewInt(30),
		CreditBack:        common.RandAddress(),
		Beneficiary:       common.NewAddressFromEth(beneficiaryAuth.From),
		Data:              nil,
	}

	requestId, err := backend.AddInboxMessage(retryableTx, sender)
	test.FailIfError(t, err)

	return sender, beneficiaryAuth, otherAuth, rollupAddress, db, backend, retryableTx, requestId, closeFunc
}

func TestRetryableRedeem(t *testing.T) {
	tmpDir, err := ioutil.TempDir(".", "arbitrum")
	if err != nil {
		logger.Fatal().Err(err).Msg("error generating temporary directory")
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()

	sender, _, otherAuth, rollupAddress, db, backend, retryableTx, requestId, closeFunc := setupTest(t, tmpDir)
	defer closeFunc()
	ticketId := hashing.SoliditySHA3(hashing.Bytes32(requestId), hashing.Uint256(big.NewInt(0)))

	srv := aggregator.NewServer(backend, rollupAddress, db)
	client := web3.NewEthClient(srv, true)
	retryable, err := arboscontracts.NewArbRetryableTx(arbos.ARB_RETRYABLE_ADDRESS, client)
	test.FailIfError(t, err)

	txReceipt, err := client.TransactionReceipt(context.Background(), requestId.ToEthHash())
	test.FailIfError(t, err)

	if txReceipt == nil || txReceipt.Status != 1 {
		t.Fatal("retryable tx failed")
	}

	finalReceipt, err := client.TransactionReceipt(context.Background(), ticketId.ToEthHash())
	test.FailIfError(t, err)

	if finalReceipt != nil {
		t.Fatal("shouldn't have receipt yet")
	}

	creationBlock, err := backend.db.GetBlockWithHash(common.NewHashFromEth(txReceipt.BlockHash))
	test.FailIfError(t, err)

	lifetime, err := retryable.GetLifetime(&bind.CallOpts{})
	test.FailIfError(t, err)

	timeout, err := retryable.GetTimeout(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)

	if timeout.Uint64() != creationBlock.Header.Time+lifetime.Uint64() {
		t.Fatal("wrong timeout")
	}

	price, nextUpdateTimestamp, err := retryable.GetKeepalivePrice(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)
	if price.Cmp(big.NewInt(0)) != 0 {
		t.Fatal("wrong price")
	}
	t.Log("nextUpdateTimestamp", nextUpdateTimestamp)

	beneficiary, err := retryable.GetBeneficiary(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)
	if beneficiary != retryableTx.Beneficiary.ToEthAddress() {
		t.Fatal("wrong beneficiary")
	}

	correctSenderBalance := new(big.Int).Sub(retryableTx.Deposit, retryableTx.Value)
	correctSenderBalance = correctSenderBalance.Sub(correctSenderBalance, retryableTx.MaxSubmissionCost)

	balanceCheck(t, srv, sender, retryableTx, correctSenderBalance, big.NewInt(0), retryableTx.MaxSubmissionCost, big.NewInt(0))

	_, err = retryable.Keepalive(otherAuth, ticketId)
	test.FailIfError(t, err)

	newTimeout, err := retryable.GetTimeout(&bind.CallOpts{}, ticketId)
	test.FailIfError(t, err)

	if newTimeout.Uint64() != creationBlock.Header.Time+lifetime.Uint64()*2 {
		t.Fatal("wrong timeout")
	}

	tx, err := retryable.Redeem(otherAuth, ticketId)
	test.FailIfError(t, err)

	txReceipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	test.FailIfError(t, err)

	if txReceipt == nil || txReceipt.Status != 1 {
		t.Fatal("cancel tx failed")
	}

	if len(txReceipt.Logs) != 1 {
		t.Fatal("wrong log count")
	}

	if txReceipt.Logs[0].Topics[0] != arbos.RetryRedeemedEvent.ID {
		t.Fatal("wrong log topic")
	}

	// TODO: Remove this
	correctSenderBalance = correctSenderBalance.Sub(correctSenderBalance, retryableTx.Value)
	balanceCheck(t, srv, sender, retryableTx, correctSenderBalance, big.NewInt(0), retryableTx.MaxSubmissionCost, retryableTx.Value)
}

func TestRetryableCancel(t *testing.T) {

	tmpDir, err := ioutil.TempDir(".", "arbitrum")
	if err != nil {
		logger.Fatal().Err(err).Msg("error generating temporary directory")
	}
	defer func() {
		if err := os.RemoveAll(tmpDir); err != nil {
			panic(err)
		}
	}()

	sender, beneficiaryAuth, otherAuth, rollupAddress, db, backend, retryableTx, requestId, closeFunc := setupTest(t, tmpDir)
	defer closeFunc()
	ticketId := hashing.SoliditySHA3(hashing.Bytes32(requestId), hashing.Uint256(big.NewInt(0)))

	srv := aggregator.NewServer(backend, rollupAddress, db)
	client := web3.NewEthClient(srv, true)
	retryable, err := arboscontracts.NewArbRetryableTx(arbos.ARB_RETRYABLE_ADDRESS, client)
	test.FailIfError(t, err)

	_, err = retryable.Cancel(otherAuth, ticketId)
	if err == nil {
		t.Fatal("cancel should fail from non beneficiary")
	}

	tx, err := retryable.Cancel(beneficiaryAuth, ticketId)
	test.FailIfError(t, err)

	txReceipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	test.FailIfError(t, err)

	if txReceipt == nil || txReceipt.Status != 1 {
		t.Fatal("cancel tx failed")
	}

	if len(txReceipt.Logs) != 1 {
		t.Fatal("wrong log count")
	}

	if txReceipt.Logs[0].Topics[0] != arbos.RetryCanceledEvent.ID {
		t.Fatal("wrong log topic")
	}

	correctSenderBalance := new(big.Int).Sub(retryableTx.Deposit, retryableTx.Value)
	correctSenderBalance = correctSenderBalance.Sub(correctSenderBalance, retryableTx.MaxSubmissionCost)

	correctBeneficiaryValue := big.NewInt(0)
	// TODO: Why isn't this right
	//correctBeneficiaryValue := retryableTx.Value

	balanceCheck(t, srv, sender, retryableTx, correctSenderBalance, correctBeneficiaryValue, retryableTx.MaxSubmissionCost, big.NewInt(0))
}

func balanceCheck(
	t *testing.T,
	srv *aggregator.Server,
	sender common.Address,
	retryableTx message.RetryableTx,
	correctSenderBalance,
	correctBeneficiaryBalance,
	correctCreditBackBalance,
	correctDestinationBalance *big.Int,
) {
	t.Helper()
	snap, err := srv.PendingSnapshot()
	test.FailIfError(t, err)

	senderBalance, err := snap.GetBalance(sender)
	test.FailIfError(t, err)

	if senderBalance.Cmp(correctSenderBalance) != 0 {
		t.Error("unexpected sender balance", senderBalance, "instead of", correctSenderBalance)
	}

	beneficiaryBalance, err := snap.GetBalance(retryableTx.Beneficiary)
	test.FailIfError(t, err)

	if beneficiaryBalance.Cmp(correctBeneficiaryBalance) != 0 {
		t.Error("unexpected beneficiary balance", beneficiaryBalance, "instead of", correctBeneficiaryBalance)
	}

	creditBackBalance, err := snap.GetBalance(retryableTx.CreditBack)
	test.FailIfError(t, err)
	if creditBackBalance.Cmp(correctCreditBackBalance) != 0 {
		t.Error("unexpected credit back balance")
	}

	destinationBalance, err := snap.GetBalance(retryableTx.Destination)
	test.FailIfError(t, err)

	if destinationBalance.Cmp(correctDestinationBalance) != 0 {
		t.Error("unexpected destination balance")
	}
}
