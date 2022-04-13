/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-avm-cpp/cmachine"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbtransaction"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks"
	"github.com/offchainlabs/arbitrum/packages/arb-util/transactauth"
)

const eip1820Tx = "0xf90a388085174876e800830c35008080b909e5608060405234801561001057600080fd5b506109c5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c010000000000000000000000000000000000000000000000000000000090048063a41e7d5111610078578063a41e7d51146101d4578063aabbb8ca1461020a578063b705676514610236578063f712f3e814610280576100a5565b806329965a1d146100aa5780633d584063146100e25780635df8122f1461012457806365ba36c114610152575b600080fd5b6100e0600480360360608110156100c057600080fd5b50600160a060020a038135811691602081013591604090910135166102b6565b005b610108600480360360208110156100f857600080fd5b5035600160a060020a0316610570565b60408051600160a060020a039092168252519081900360200190f35b6100e06004803603604081101561013a57600080fd5b50600160a060020a03813581169160200135166105bc565b6101c26004803603602081101561016857600080fd5b81019060208101813564010000000081111561018357600080fd5b82018360208201111561019557600080fd5b803590602001918460018302840111640100000000831117156101b757600080fd5b5090925090506106b3565b60408051918252519081900360200190f35b6100e0600480360360408110156101ea57600080fd5b508035600160a060020a03169060200135600160e060020a0319166106ee565b6101086004803603604081101561022057600080fd5b50600160a060020a038135169060200135610778565b61026c6004803603604081101561024c57600080fd5b508035600160a060020a03169060200135600160e060020a0319166107ef565b604080519115158252519081900360200190f35b61026c6004803603604081101561029657600080fd5b508035600160a060020a03169060200135600160e060020a0319166108aa565b6000600160a060020a038416156102cd57836102cf565b335b9050336102db82610570565b600160a060020a031614610339576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f7420746865206d616e616765720000000000000000000000000000000000604482015290519081900360640190fd5b6103428361092a565b15610397576040805160e560020a62461bcd02815260206004820152601a60248201527f4d757374206e6f7420626520616e204552433136352068617368000000000000604482015290519081900360640190fd5b600160a060020a038216158015906103b85750600160a060020a0382163314155b156104ff5760405160200180807f455243313832305f4143434550545f4d4147494300000000000000000000000081525060140190506040516020818303038152906040528051906020012082600160a060020a031663249cb3fa85846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182600160a060020a0316600160a060020a031681526020019250505060206040518083038186803b15801561047e57600080fd5b505afa158015610492573d6000803e3d6000fd5b505050506040513d60208110156104a857600080fd5b5051146104ff576040805160e560020a62461bcd02815260206004820181905260248201527f446f6573206e6f7420696d706c656d656e742074686520696e74657266616365604482015290519081900360640190fd5b600160a060020a03818116600081815260208181526040808320888452909152808220805473ffffffffffffffffffffffffffffffffffffffff19169487169485179055518692917f93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db15391a450505050565b600160a060020a03818116600090815260016020526040812054909116151561059a5750806105b7565b50600160a060020a03808216600090815260016020526040902054165b919050565b336105c683610570565b600160a060020a031614610624576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f7420746865206d616e616765720000000000000000000000000000000000604482015290519081900360640190fd5b81600160a060020a031681600160a060020a0316146106435780610646565b60005b600160a060020a03838116600081815260016020526040808220805473ffffffffffffffffffffffffffffffffffffffff19169585169590951790945592519184169290917f605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a43509190a35050565b600082826040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012090505b92915050565b6106f882826107ef565b610703576000610705565b815b600160a060020a03928316600081815260208181526040808320600160e060020a031996909616808452958252808320805473ffffffffffffffffffffffffffffffffffffffff19169590971694909417909555908152600284528181209281529190925220805460ff19166001179055565b600080600160a060020a038416156107905783610792565b335b905061079d8361092a565b156107c357826107ad82826108aa565b6107b85760006107ba565b815b925050506106e8565b600160a060020a0390811660009081526020818152604080832086845290915290205416905092915050565b6000808061081d857f01ffc9a70000000000000000000000000000000000000000000000000000000061094c565b909250905081158061082d575080155b1561083d576000925050506106e8565b61084f85600160e060020a031961094c565b909250905081158061086057508015155b15610870576000925050506106e8565b61087a858561094c565b909250905060018214801561088f5750806001145b1561089f576001925050506106e8565b506000949350505050565b600160a060020a0382166000908152600260209081526040808320600160e060020a03198516845290915281205460ff1615156108f2576108eb83836107ef565b90506106e8565b50600160a060020a03808316600081815260208181526040808320600160e060020a0319871684529091529020549091161492915050565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161590565b6040517f01ffc9a7000000000000000000000000000000000000000000000000000000008082526004820183905260009182919060208160248189617530fa90519096909550935050505056fea165627a7a72305820377f4a2d4301ede9949f163f319021a6e9c687c292a5e2b2c4734c126b524e6c00291ba01820182018201820182018201820182018201820182018201820182018201820a01820182018201820182018201820182018201820182018201820182018201820"

type Config struct {
	client ethutils.EthClient
	auth   *bind.TransactOpts
	fb     *fireblocks.Fireblocks
}

var config *Config

func waitForTx(tx *types.Transaction, method string) error {
	fmt.Println("Waiting for receipt")
	_, err := transactauth.WaitForReceiptWithResults(context.Background(), config.client, config.auth.From, arbtransaction.NewArbTransaction(tx), method, transactauth.NewEthArbReceiptFetcher(config.client))
	if err != nil {
		return err
	}
	fmt.Println("Transaction completed successfully")
	return nil
}

type upgrade struct {
	Instructions []string `json:"instructions"`
}

func upgradeArbOS(upgradeFile string, targetMexe string, startMexe *string) error {
	targetMach, err := cmachine.New(targetMexe)
	if err != nil {
		return err
	}

	var startHash common.Hash
	if startMexe != nil {
		startMach, err := cmachine.New(*startMexe)
		if err != nil {
			return err
		}
		startHash = startMach.CodePointHash()
	}

	updateBytes, err := ioutil.ReadFile(upgradeFile)
	if err != nil {
		return err
	}
	upgrade := upgrade{}
	err = json.Unmarshal(updateBytes, &upgrade)
	if err != nil {
		return err
	}
	chunkSize := 50000
	chunks := []string{"0x"}
	for _, insn := range upgrade.Instructions {
		if len(chunks[len(chunks)-1])+len(insn) > chunkSize {
			chunks = append(chunks, "0x")
		}
		chunks[len(chunks)-1] += insn
	}

	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, config.client)
	if err != nil {
		return err
	}
	tx, err := arbOwner.StartCodeUpload(config.auth)
	if err != nil {
		return err
	}
	if err := waitForTx(tx, "StartCodeUpload"); err != nil {
		return err
	}

	fmt.Println("Submitting upgrade in", len(chunks), "chunks")
	for _, upgradeChunk := range chunks {
		tx, err = arbOwner.ContinueCodeUpload(config.auth, hexutil.MustDecode(upgradeChunk))
		if err != nil {
			return err
		}
		if err := waitForTx(tx, "ContinueCodeUpload"); err != nil {
			return err
		}
	}

	codeHash, err := arbOwner.GetUploadedCodeHash(&bind.CallOpts{})
	if err != nil {
		return err
	}
	if codeHash != targetMach.CodePointHash() {
		return errors.New("incorrect code segment uploaded")
	}

	_, err = arbOwner.FinishCodeUploadAsArbosUpgrade(config.auth, targetMach.CodePointHash(), startHash)
	if err != nil {
		return err
	}
	if err := waitForTx(tx, "FinishCodeUploadAsArbosUpgrade"); err != nil {
		return err
	}
	return nil
}

func version() error {
	con, err := arboscontracts.NewArbSys(arbos.ARB_SYS_ADDRESS, config.client)
	if err != nil {
		return err
	}
	version, err := con.ArbOSVersion(&bind.CallOpts{})
	if err != nil {
		return err
	}
	fmt.Println("ArbOS Version:", version)
	return nil
}

func feeInfo(blockNum *big.Int) error {
	con, err := arboscontracts.NewArbGasInfo(arbos.ARB_GAS_INFO_ADDRESS, config.client)
	if err != nil {
		return err
	}
	opts := &bind.CallOpts{
		BlockNumber: blockNum,
	}
	perL2TxWei,
		perL1CalldataByteWei,
		perStorageWei,
		perArgGasBaseWei,
		perArbGasCongestionWei,
		perArbGasTotalWei,
		err := con.GetPricesInWei(opts)
	if err != nil {
		return err
	}
	fmt.Println("perL2TxWei:", perL2TxWei)
	fmt.Println("perL1CalldataByteWei:", perL1CalldataByteWei)
	fmt.Println("perStorageWei:", perStorageWei)
	fmt.Println("perArgGasBaseWei:", perArgGasBaseWei)
	fmt.Println("perArbGasCongestionWei:", perArbGasCongestionWei)
	fmt.Println("perArbGasTotalWei:", perArbGasTotalWei)

	perL2Tx, perL1CalldataByte, perStorage, err := con.GetPricesInArbGas(opts)
	if err != nil {
		return err
	}
	fmt.Println("perL2Tx:", perL2Tx)
	fmt.Println("perL1CalldataByte:", perL1CalldataByte)
	fmt.Println("perStorage:", perStorage)

	speedLimitPerSecond, gasPoolMax, maxTxGasLimit, err := con.GetGasAccountingParams(opts)
	if err != nil {
		return err
	}
	fmt.Println("speedLimitPerSecond:", speedLimitPerSecond)
	fmt.Println("gasPoolMax:", gasPoolMax)
	fmt.Println("maxTxGasLimit:", maxTxGasLimit)
	return nil
}

func switchFees(enabled bool) error {
	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, config.client)
	if err != nil {
		return err
	}
	tx, err := arbOwner.SetChainParameter(config.auth, arbos.FeesEnabledParamId, big.NewInt(1))
	if err != nil {
		return err
	}
	return waitForTx(tx, "SetChainParameters")
}

func setDefaultAggregator(agg ethcommon.Address) error {
	arbAggregator, err := arboscontracts.NewArbAggregator(arbos.ARB_AGGREGATOR_ADDRESS, config.client)
	if err != nil {
		return err
	}
	tx, err := arbAggregator.SetDefaultAggregator(config.auth, agg)
	if err != nil {
		return err
	}
	return waitForTx(tx, "SetDefaultAggregator")
}

func setFairGasPriceSender(sender ethcommon.Address) error {
	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, config.client)
	if err != nil {
		return err
	}
	tx, err := arbOwner.SetFairGasPriceSender(config.auth, sender, true)
	if err != nil {
		return err
	}
	return waitForTx(tx, "SetFairGasPriceSender")
}

func deploy1820() error {
	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, config.client)
	if err != nil {
		return err
	}

	tx1820 := new(types.Transaction)
	if err := rlp.DecodeBytes(hexutil.MustDecode(eip1820Tx), tx1820); err != nil {
		return err
	}
	sender1820, err := types.Sender(types.HomesteadSigner{}, tx1820)
	if err != nil {
		return err
	}

	_, err = arbOwner.DeployContract(config.auth, tx1820.Data(), sender1820, new(big.Int).SetUint64(tx1820.Nonce()))
	if err != nil {
		return err
	}
	return nil
}

func estimateTransferGas() error {
	dest := common.RandAddress().ToEthAddress()
	msg := ethereum.CallMsg{
		From: config.auth.From,
		To:   &dest,
	}
	gas, err := config.client.EstimateGas(context.Background(), msg)
	if err != nil {
		return err
	}
	fmt.Println("Gas estimate:", gas)
	return nil
}

func spam() error {
	dest := common.RandAddress().ToEthAddress()
	ctx := context.Background()
	for {
		nonce, err := config.client.PendingNonceAt(ctx, config.auth.From)
		if err != nil {
			return err
		}
		gasPrice, err := config.client.SuggestGasPrice(ctx)
		if err != nil {
			return err
		}
		tx := types.NewTx(&types.LegacyTx{
			Nonce:    nonce,
			GasPrice: gasPrice,
			Gas:      1000000,
			To:       &dest,
			Value:    big.NewInt(1),
			Data:     nil,
		})
		tx, err = config.auth.Signer(config.auth.From, tx)
		if err != nil {
			return err
		}
		if err := config.client.SendTransaction(ctx, tx); err != nil {
			return err
		}
		time.Sleep(time.Minute)
	}
}

func handleCommand(fields []string) error {
	switch fields[0] {
	case "enable-fees":
		if len(fields) != 2 {
			return errors.New("Expected a true or false argument")
		}
		enabled, err := strconv.ParseBool(fields[1])
		if err != nil {
			return err
		}
		return switchFees(enabled)
	case "estimate-transfer-gas":
		return estimateTransferGas()
	case "set-default-agg":
		if len(fields) != 2 {
			return errors.New("Expected address argument")
		}
		agg := ethcommon.HexToAddress(fields[1])
		return setDefaultAggregator(agg)
	case "set-fair-gas-sender":
		if len(fields) != 2 {
			return errors.New("Expected address argument")
		}
		agg := ethcommon.HexToAddress(fields[1])
		return setFairGasPriceSender(agg)
	case "deploy-1820":
		return deploy1820()
	case "fee-info":
		var blockNum *big.Int
		if len(fields) == 2 {
			var ok bool
			blockNum, ok = new(big.Int).SetString(fields[1], 10)
			if !ok {
				return errors.New("expected arg to be int")
			}
		}
		return feeInfo(blockNum)
	case "upgrade":
		if len(fields) != 3 && len(fields) != 4 {
			return errors.New("Expected upgrade file and target mexe arguments")
		}
		var source *string
		if len(fields) == 4 {
			source = &fields[3]
		}
		return upgradeArbOS(fields[1], fields[2], source)
	case "version":
		return version()
	case "spam":
		return spam()
	default:
		fmt.Println("Unknown command")
	}
	return nil
}

func executor(t string) {
	if t == "exit" {
		os.Exit(0)
	}
	fields := strings.Fields(t)
	err := handleCommand(fields)
	if err != nil {
		fmt.Println("Error running command", err)
	}
}

func completer(t prompt.Document) []prompt.Suggest {
	return []prompt.Suggest{
		{Text: "enable-fees"},
		{Text: "exit"},
	}
}

func run(ctx context.Context) error {
	if len(os.Args) < 3 {
		fmt.Println("Expected: arb-cli rpcurl privkey")
	}
	arbUrl := os.Args[1]
	privKeystr := os.Args[2]

	client, err := ethutils.NewRPCEthClient(arbUrl)
	if err != nil {
		return err
	}
	chainId, err := client.ChainID(ctx)
	if err != nil {
		return err
	}
	privKey, err := crypto.HexToECDSA(privKeystr)
	if err != nil {
		return err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainId)
	if err != nil {
		return err
	}
	fmt.Println("Sending from address", auth.From)

	config = &Config{
		client: client,
		auth:   auth,
		fb:     nil,
	}

	if len(os.Args) >= 3 {
		if err := handleCommand(os.Args[3:]); err != nil {
			return fmt.Errorf("error running command: %w", err)
		}
		return nil
	}

	p := prompt.New(
		executor,
		completer,
	)
	p.Run()
	return nil
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Println("Error running app", err)
	}
}
