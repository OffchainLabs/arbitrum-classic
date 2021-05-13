/*
 * Copyright 2020-2021, Offchain Labs, Inc.
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
	"crypto/ecdsa"
	"flag"
	"fmt"
	"io/ioutil"
	golog "log"
	"math/big"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/arboscontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/aggregator"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/dev"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/web3"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"

	accounts2 "github.com/ethereum/go-ethereum/accounts"
	gethlog "github.com/ethereum/go-ethereum/log"
	"github.com/miguelmota/go-ethereum-hdwallet"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/rpc"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

var logger zerolog.Logger
var pprofMux *http.ServeMux

var canceled = false

const eip1820Tx = "0xf90a388085174876e800830c35008080b909e5608060405234801561001057600080fd5b506109c5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a5576000357c010000000000000000000000000000000000000000000000000000000090048063a41e7d5111610078578063a41e7d51146101d4578063aabbb8ca1461020a578063b705676514610236578063f712f3e814610280576100a5565b806329965a1d146100aa5780633d584063146100e25780635df8122f1461012457806365ba36c114610152575b600080fd5b6100e0600480360360608110156100c057600080fd5b50600160a060020a038135811691602081013591604090910135166102b6565b005b610108600480360360208110156100f857600080fd5b5035600160a060020a0316610570565b60408051600160a060020a039092168252519081900360200190f35b6100e06004803603604081101561013a57600080fd5b50600160a060020a03813581169160200135166105bc565b6101c26004803603602081101561016857600080fd5b81019060208101813564010000000081111561018357600080fd5b82018360208201111561019557600080fd5b803590602001918460018302840111640100000000831117156101b757600080fd5b5090925090506106b3565b60408051918252519081900360200190f35b6100e0600480360360408110156101ea57600080fd5b508035600160a060020a03169060200135600160e060020a0319166106ee565b6101086004803603604081101561022057600080fd5b50600160a060020a038135169060200135610778565b61026c6004803603604081101561024c57600080fd5b508035600160a060020a03169060200135600160e060020a0319166107ef565b604080519115158252519081900360200190f35b61026c6004803603604081101561029657600080fd5b508035600160a060020a03169060200135600160e060020a0319166108aa565b6000600160a060020a038416156102cd57836102cf565b335b9050336102db82610570565b600160a060020a031614610339576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f7420746865206d616e616765720000000000000000000000000000000000604482015290519081900360640190fd5b6103428361092a565b15610397576040805160e560020a62461bcd02815260206004820152601a60248201527f4d757374206e6f7420626520616e204552433136352068617368000000000000604482015290519081900360640190fd5b600160a060020a038216158015906103b85750600160a060020a0382163314155b156104ff5760405160200180807f455243313832305f4143434550545f4d4147494300000000000000000000000081525060140190506040516020818303038152906040528051906020012082600160a060020a031663249cb3fa85846040518363ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018083815260200182600160a060020a0316600160a060020a031681526020019250505060206040518083038186803b15801561047e57600080fd5b505afa158015610492573d6000803e3d6000fd5b505050506040513d60208110156104a857600080fd5b5051146104ff576040805160e560020a62461bcd02815260206004820181905260248201527f446f6573206e6f7420696d706c656d656e742074686520696e74657266616365604482015290519081900360640190fd5b600160a060020a03818116600081815260208181526040808320888452909152808220805473ffffffffffffffffffffffffffffffffffffffff19169487169485179055518692917f93baa6efbd2244243bfee6ce4cfdd1d04fc4c0e9a786abd3a41313bd352db15391a450505050565b600160a060020a03818116600090815260016020526040812054909116151561059a5750806105b7565b50600160a060020a03808216600090815260016020526040902054165b919050565b336105c683610570565b600160a060020a031614610624576040805160e560020a62461bcd02815260206004820152600f60248201527f4e6f7420746865206d616e616765720000000000000000000000000000000000604482015290519081900360640190fd5b81600160a060020a031681600160a060020a0316146106435780610646565b60005b600160a060020a03838116600081815260016020526040808220805473ffffffffffffffffffffffffffffffffffffffff19169585169590951790945592519184169290917f605c2dbf762e5f7d60a546d42e7205dcb1b011ebc62a61736a57c9089d3a43509190a35050565b600082826040516020018083838082843780830192505050925050506040516020818303038152906040528051906020012090505b92915050565b6106f882826107ef565b610703576000610705565b815b600160a060020a03928316600081815260208181526040808320600160e060020a031996909616808452958252808320805473ffffffffffffffffffffffffffffffffffffffff19169590971694909417909555908152600284528181209281529190925220805460ff19166001179055565b600080600160a060020a038416156107905783610792565b335b905061079d8361092a565b156107c357826107ad82826108aa565b6107b85760006107ba565b815b925050506106e8565b600160a060020a0390811660009081526020818152604080832086845290915290205416905092915050565b6000808061081d857f01ffc9a70000000000000000000000000000000000000000000000000000000061094c565b909250905081158061082d575080155b1561083d576000925050506106e8565b61084f85600160e060020a031961094c565b909250905081158061086057508015155b15610870576000925050506106e8565b61087a858561094c565b909250905060018214801561088f5750806001145b1561089f576001925050506106e8565b506000949350505050565b600160a060020a0382166000908152600260209081526040808320600160e060020a03198516845290915281205460ff1615156108f2576108eb83836107ef565b90506106e8565b50600160a060020a03808316600081815260208181526040808320600160e060020a0319871684529091529020549091161492915050565b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff161590565b6040517f01ffc9a7000000000000000000000000000000000000000000000000000000008082526004820183905260009182919060208160248189617530fa90519096909550935050505056fea165627a7a72305820377f4a2d4301ede9949f163f319021a6e9c687c292a5e2b2c4734c126b524e6c00291ba01820182018201820182018201820182018201820182018201820182018201820a01820182018201820182018201820182018201820182018201820182018201820"

func init() {
	pprofMux = http.DefaultServeMux
	http.DefaultServeMux = http.NewServeMux()

	// Enable line numbers in logging
	golog.SetFlags(golog.LstdFlags | golog.Lshortfile)

	// Print stack trace when `.Error().Stack().Err(err).` is added to zerolog call
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	gethlog.Root().SetHandler(gethlog.LvlFilterHandler(gethlog.LvlInfo, gethlog.StreamHandler(os.Stderr, gethlog.TerminalFormat(true))))

	// Print line number that log was created on
	logger = log.With().Caller().Stack().Str("component", "arb-dev-node").Logger()
}

func main() {
	if err := startup(); err != nil {
		logger.Error().Err(err).Msg("Error running dev node")
	}
}

func startup() error {
	fs := flag.NewFlagSet("", flag.ContinueOnError)

	enablePProf := fs.Bool("pprof", false, "enable profiling server")
	saveMessages := fs.String("save", "", "save messages")
	walletcount := fs.Int("walletcount", 10, "number of wallets to fund")
	walletbalance := fs.Int64("walletbalance", 100, "amount of funds in each wallet (Eth)")
	arbosPath := fs.String("arbos", "", "ArbOS version")
	enableFees := fs.Bool("with-fees", false, "Run arbos with fees on")
	mnemonic := fs.String(
		"mnemonic",
		"jar deny prosper gasp flush glass core corn alarm treat leg smart",
		"mnemonic to generate accounts from",
	)
	gethLogLevel, arbLogLevel := cmdhelp.AddLogFlags(fs)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return errors.Wrap(err, "error parsing arguments")
	}

	if err := cmdhelp.ParseLogFlags(gethLogLevel, arbLogLevel); err != nil {
		return err
	}

	if *enablePProf {
		go func() {
			err := http.ListenAndServe("localhost:8081", pprofMux)
			log.Error().Err(err).Msg("profiling server failed")
		}()
	}

	tmpDir, err := ioutil.TempDir(".", "arbitrum")
	if err != nil {
		return errors.Wrap(err, "error generating temporary directory")
	}

	wallet, err := hdwallet.NewFromMnemonic(*mnemonic)
	if err != nil {
		return err
	}

	depositSize, ok := new(big.Int).SetString("1000000000000000000", 10)
	if !ok {
		return errors.New("invalid value for deposit amount")
	}
	depositSize = depositSize.Mul(depositSize, big.NewInt(*walletbalance))

	accounts := make([]accounts2.Account, 0)
	for i := 0; i < *walletcount; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%v", i))
		account, err := wallet.Derive(path, false)
		if err != nil {
			return err
		}
		accounts = append(accounts, account)
	}

	if *arbosPath == "" {
		arbosPathStr, err := arbos.Path()
		if err != nil {
			return err
		}
		arbosPath = &arbosPathStr
	}

	errChan := make(chan error, 10)
	defer close(errChan)

	ctx := context.Background()

	config := protocol.ChainParams{
		StakeRequirement:          big.NewInt(10),
		StakeToken:                common.Address{},
		GracePeriod:               common.NewTimeBlocksInt(3),
		MaxExecutionSteps:         10000000000,
		ArbGasSpeedLimitPerSecond: 2000000000000,
	}

	var configOptions []message.ChainConfigOption
	aggInit := message.DefaultAggConfig{Aggregator: common.NewAddressFromEth(accounts[1].Address)}
	if *enableFees {
		configOptions = append(configOptions, aggInit)

		netFeeRecipient := common.RandAddress()
		congestionFeeRecipient := common.RandAddress()
		feeConfigInit := message.FeeConfig{
			SpeedLimitPerSecond:    new(big.Int).SetUint64(config.ArbGasSpeedLimitPerSecond),
			L1GasPerL2Tx:           big.NewInt(3700),
			ArbGasPerL2Tx:          big.NewInt(0),
			L1GasPerL2Calldata:     big.NewInt(1),
			ArbGasPerL2Calldata:    big.NewInt(0),
			L1GasPerStorage:        big.NewInt(2000),
			ArbGasPerStorage:       big.NewInt(0),
			ArbGasDivisor:          big.NewInt(10000),
			NetFeeRecipient:        netFeeRecipient,
			CongestionFeeRecipient: congestionFeeRecipient,
		}

		configOptions = append(configOptions, feeConfigInit)
	}

	backend, db, rollupAddress, cancelDevNode, devNodeErrChan, err := dev.NewDevNode(
		ctx,
		tmpDir,
		*arbosPath,
		config,
		common.NewAddressFromEth(accounts[0].Address),
		configOptions,
	)
	if err != nil {
		return err
	}

	go func() {
		errChan <- <-devNodeErrChan
	}()

	cancel := func() {
		if !canceled {
			cancelDevNode()
			if err := os.RemoveAll(tmpDir); err != nil {
				panic(err)
			}
			canceled = true
		}
	}
	defer cancel()

	for _, account := range accounts {
		deposit := message.EthDepositTx{
			L2Message: message.NewSafeL2Message(message.ContractTransaction{
				BasicTx: message.BasicTx{
					MaxGas:      big.NewInt(1000000),
					GasPriceBid: big.NewInt(0),
					DestAddress: common.NewAddressFromEth(account.Address),
					Payment:     depositSize,
					Data:        nil,
				},
			}),
		}
		if _, err := backend.AddInboxMessage(deposit, common.RandAddress()); err != nil {
			return err
		}
	}

	privateKeys := make([]*ecdsa.PrivateKey, 0)
	for _, account := range accounts {
		privKey, err := wallet.PrivateKey(account)
		if err != nil {
			return err
		}
		privateKeys = append(privateKeys, privKey)
	}

	chainId := message.ChainAddressToID(rollupAddress)
	ownerAuth, err := bind.NewKeyedTransactorWithChainID(privateKeys[0], chainId)
	if err != nil {
		return err
	}
	signer := types.NewEIP155Signer(chainId)

	srv := aggregator.NewServer(backend, rollupAddress, db)

	client := web3.NewEthClient(srv, true)
	arbOwner, err := arboscontracts.NewArbOwner(arbos.ARB_OWNER_ADDRESS, client)
	if err != nil {
		return err
	}

	tx1820 := new(types.Transaction)
	if err := rlp.DecodeBytes(hexutil.MustDecode(eip1820Tx), tx1820); err != nil {
		return err
	}
	sender1820, err := types.Sender(signer, tx1820)
	if err != nil {
		return err
	}

	_, err = arbOwner.DeployContract(ownerAuth, tx1820.Data(), sender1820, new(big.Int).SetUint64(tx1820.Nonce()))
	if err != nil {
		return err
	}

	if *enableFees {
		_, err = arbOwner.SetFairGasPriceSender(ownerAuth, aggInit.Aggregator.ToEthAddress())
		if err != nil {
			return err
		}

		_, err = arbOwner.SetFeesEnabled(ownerAuth, true)
		if err != nil {
			return err
		}
		if _, err := backend.AddInboxMessage(message.NewSafeL2Message(message.HeartbeatMessage{}), common.RandAddress()); err != nil {
			return err
		}
	}

	fmt.Println("Arbitrum Dev Chain")
	fmt.Println("")
	fmt.Println("Available Accounts")
	fmt.Println("==================")
	for i, account := range accounts {
		fmt.Printf("(%v) %v (100 ETH)\n", i, account.Address.Hex())
	}

	fmt.Println("\nPrivate Keys")
	fmt.Println("==================")
	for i, account := range accounts {
		privKey, err := wallet.PrivateKeyHex(account)
		if err != nil {
			return err
		}
		fmt.Printf("(%v) 0x%v\n", i, privKey)
	}
	fmt.Println("")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		if *saveMessages != "" {
			data, err := backend.ExportData()
			if err != nil {
				errChan <- errors.Wrap(err, "error exporting data from backend")
				return
			}

			if err := ioutil.WriteFile(*saveMessages, data, 777); err != nil {
				errChan <- errors.Wrap(err, "error saving exported data")
				return
			}
		}
		errChan <- nil
	}()

	plugins := make(map[string]interface{})
	plugins["evm"] = dev.NewEVM(backend)

	web3Server, err := web3.GenerateWeb3Server(srv, privateKeys, true, plugins)
	if err != nil {
		return err
	}

	go func() {
		errChan <- rpc.LaunchPublicServer(ctx, web3Server, "8547", "8548")
	}()

	err = <-errChan
	return err
}
