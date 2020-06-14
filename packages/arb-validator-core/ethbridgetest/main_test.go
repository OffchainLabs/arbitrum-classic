package ethbridgetest

import (
	"context"
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/protocoltester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/valuetester"
	"log"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/messagetester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetest/sigutilstester"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
)

var privHex = "27e926925fb5903ee038c894d9880f74d3dd6518e23ab5e5651de93327c7dffa"

var auth *bind.TransactOpts
var tester *messagetester.MessageTester
var sigTester *sigutilstester.SigUtilsTester
var valueTester *valuetester.ValueTester
var protocolTester *protocoltester.ProtocolTester
var client *ethclient.Client

var addr1 common.Address
var addr2 common.Address
var addr3 common.Address

var errHash = errors.New("ethbridge calculated wrong hash")
var errMsgHash = errors.New("ethbridge calculated wrong message hash")

func TestMain(m *testing.M) {
	addr1[0] = 76
	addr1[19] = 93
	addr2[0] = 43
	addr2[19] = 12
	addr3[0] = 73
	addr3[19] = 85

	var err error

	auth, err = test.SetupAuth(privHex)
	if err != nil {
		log.Fatal(err)
	}
	client, err = ethclient.Dial(test.GetEthUrl())
	if err != nil {
		log.Fatal(err)
	}
	_, tx, deployedTester, err := messagetester.DeployMessageTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		tx,
		"DeployMessageTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	_, sigTx, deployedSigTester, err := sigutilstester.DeploySigUtilsTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		sigTx,
		"DeploySigUtilsTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	_, valTx, deployedValueTester, err := valuetester.DeployValueTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		valTx,
		"DeployValueTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	_, protocolTx, deployedProtocolTester, err := protocoltester.DeployProtocolTester(
		auth,
		client,
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = ethbridge.WaitForReceiptWithResults(
		context.Background(),
		client,
		auth.From,
		protocolTx,
		"DeployValueTester",
	)
	if err != nil {
		log.Fatal(err)
	}

	sigTester = deployedSigTester
	tester = deployedTester
	valueTester = deployedValueTester
	protocolTester = deployedProtocolTester
	code := m.Run()
	os.Exit(code)
}
