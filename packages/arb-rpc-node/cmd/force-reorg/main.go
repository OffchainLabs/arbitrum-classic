/*
 * Copyright 2022, Offchain Labs, Inc.
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
	"crypto/rand"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-node-core/cmdhelp"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbtransaction"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-util/transactauth"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	flag "github.com/spf13/pflag"
)

var logger zerolog.Logger

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr}).With().Str("component", "force-reorg").Logger()

	keystorePath := flag.String("keystore", "", "path to keystore with sequencer key")
	rpcUrl := flag.String("rpc", "", "L1 RPC URL")
	chainId := flag.Uint64("chainid", 1337, "chain ID")
	sequencerInboxAddr := flag.String("sequencerinbox", "", "sequencer inbox address")

	flag.Parse()

	if !common.IsHexAddress(*sequencerInboxAddr) {
		panic("--sequencerinbox must be specified and a valid address")
	}

	ctx := context.Background()
	client, err := ethutils.NewRPCEthClient(*rpcUrl)
	if err != nil {
		panic(err)
	}
	keystore, account, _, err := cmdhelp.OpenKeystore("sequencer", *keystorePath, nil, false)
	if err != nil {
		panic(err)
	}
	auth, err := bind.NewKeyStoreTransactorWithChainID(keystore, *account, new(big.Int).SetUint64(*chainId))
	if err != nil {
		panic(err)
	}
	logger.Info().Str("address", auth.From.String()).Msg("Loaded wallet")

	sequencerInbox, err := ethbridgecontracts.NewSequencerInbox(common.HexToAddress(*sequencerInboxAddr), client)
	if err != nil {
		panic(err)
	}
	latestHeader, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		panic(err)
	}
	callOpts := &bind.CallOpts{}
	msgCount, err := sequencerInbox.MessageCount(callOpts)
	if err != nil {
		panic(err)
	}
	batchCount, err := sequencerInbox.GetInboxAccsLength(callOpts)
	if err != nil {
		panic(err)
	}
	prevAcc, err := sequencerInbox.InboxAccs(callOpts, new(big.Int).Sub(batchCount, big.NewInt(1)))
	if err != nil {
		panic(err)
	}
	delayedMessagesRead, err := sequencerInbox.TotalDelayedMessagesRead(callOpts)
	if err != nil {
		panic(err)
	}
	// Make a random invalid message
	msg := make([]byte, 33)
	msg[0] = 0xFF
	_, err = rand.Read(msg[1:])
	if err != nil {
		panic(err)
	}
	metadata := []*big.Int{big.NewInt(1), latestHeader.Number, new(big.Int).SetUint64(latestHeader.Time), delayedMessagesRead, new(big.Int)}
	newAcc := crypto.Keccak256Hash(
		prevAcc[:],
		math.U256Bytes(msgCount),
		crypto.Keccak256(
			auth.From.Bytes(),
			math.U256Bytes(metadata[1]),
			math.U256Bytes(metadata[2]),
		),
		crypto.Keccak256(msg),
	)

	logger.Info().Str("acc", newAcc.String()).Msg("Prepared batch")
	msgLength := []*big.Int{big.NewInt(int64(len(msg)))}
	tx, err := sequencerInbox.AddSequencerL2BatchFromOrigin(auth, msg, msgLength, metadata, newAcc)
	if err != nil {
		panic(err)
	}
	logger.Info().Str("tx", tx.Hash().String()).Msg("Sent transaction")

	_, err = transactauth.WaitForReceiptWithResults(ctx, client, auth.From, arbtransaction.NewArbTransaction(tx), "AddSequencerL2BatchFromOrigin", transactauth.NewEthArbReceiptFetcher(client))
	if err != nil {
		panic(err)
	}

	logger.Info().Msg("Done!")
}
