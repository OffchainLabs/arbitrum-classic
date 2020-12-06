/*
 * Copyright 2020, Offchain Labs, Inc.
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

package ethbridgetest

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/ethutils"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/test"
	"math/big"
	"math/rand"
	"testing"
)

func TestKeccak(t *testing.T) {
	ctx := context.Background()
	backend, pks := test.SimulatedBackend()
	client := &ethutils.SimulatedEthClient{SimulatedBackend: backend}
	auth := bind.NewKeyedTransactor(pks[0])
	authClient, err := ethbridge.NewEthAuthClient(ctx, client, auth)
	if err != nil {
		t.Fatal(err)
	}

	precompilesTesterAddr, _, err := authClient.MakeContract(ctx, func(auth *bind.TransactOpts) (ethcommon.Address, *types.Transaction, interface{}, error) {
		return ethbridgetestcontracts.DeployPrecompilesTester(auth, client)
	})
	if err != nil {
		t.Fatal(err)
	}

	client.Commit()

	precompilesTester, err := ethbridgetestcontracts.NewPrecompilesTester(precompilesTesterAddr, client)
	if err != nil {
		t.Fatal(err)
	}

	var data [25]uint64
	var bigData [25]*big.Int
	for i := 0; i < 25; i++ {
		data[i] = rand.Uint64()
	}

	for i := 0; i < 25; i++ {
		bigData[i] = new(big.Int).SetUint64(data[5*(i%5)+i/5])
	}

	ret, err := precompilesTester.KeccakF(nil, bigData)
	if err != nil {
		t.Fatal(err)
	}

	keccakF1600(&data)

	t.Log(ret)
	t.Log(data)

	var permuted [25]uint64
	for i := range ret {
		permuted[i] = data[5*(i%5)+i/5]
	}

	for i := range ret {
		t.Log(hexutil.EncodeUint64(ret[i].Uint64()), permuted)
	}

	for i := range ret {
		if ret[i].Cmp(new(big.Int).SetUint64(permuted[i])) != 0 {
			t.Fatal("result hash didn't match")
		}
	}
}

// The following impl was taken from golang.org/x/crypto/sha3/keccakf.go
// under the BSD license

// rc stores the round constants for use in the ι step.
var rc = [24]uint64{
	0x0000000000000001,
	0x0000000000008082,
	0x800000000000808A,
	0x8000000080008000,
	0x000000000000808B,
	0x0000000080000001,
	0x8000000080008081,
	0x8000000000008009,
	0x000000000000008A,
	0x0000000000000088,
	0x0000000080008009,
	0x000000008000000A,
	0x000000008000808B,
	0x800000000000008B,
	0x8000000000008089,
	0x8000000000008003,
	0x8000000000008002,
	0x8000000000000080,
	0x000000000000800A,
	0x800000008000000A,
	0x8000000080008081,
	0x8000000000008080,
	0x0000000080000001,
	0x8000000080008008,
}

// keccakF1600 applies the Keccak permutation to a 1600b-wide
// state represented as a slice of 25 uint64s.
func keccakF1600(a *[25]uint64) {
	// Implementation translated from Keccak-inplace.c
	// in the keccak reference code.
	var t, bc0, bc1, bc2, bc3, bc4, d0, d1, d2, d3, d4 uint64

	for i := 0; i < 24; i += 4 {
		// Combines the 5 steps in each round into 2 steps.
		// Unrolls 4 rounds per loop and spreads some steps across rounds.

		// Round 1
		bc0 = a[0] ^ a[5] ^ a[10] ^ a[15] ^ a[20]
		bc1 = a[1] ^ a[6] ^ a[11] ^ a[16] ^ a[21]
		bc2 = a[2] ^ a[7] ^ a[12] ^ a[17] ^ a[22]
		bc3 = a[3] ^ a[8] ^ a[13] ^ a[18] ^ a[23]
		bc4 = a[4] ^ a[9] ^ a[14] ^ a[19] ^ a[24]
		d0 = bc4 ^ (bc1<<1 | bc1>>63)
		d1 = bc0 ^ (bc2<<1 | bc2>>63)
		d2 = bc1 ^ (bc3<<1 | bc3>>63)
		d3 = bc2 ^ (bc4<<1 | bc4>>63)
		d4 = bc3 ^ (bc0<<1 | bc0>>63)

		bc0 = a[0] ^ d0
		t = a[6] ^ d1
		bc1 = t<<44 | t>>(64-44)
		t = a[12] ^ d2
		bc2 = t<<43 | t>>(64-43)
		t = a[18] ^ d3
		bc3 = t<<21 | t>>(64-21)
		t = a[24] ^ d4
		bc4 = t<<14 | t>>(64-14)
		a[0] = bc0 ^ (bc2 &^ bc1) ^ rc[i]
		a[6] = bc1 ^ (bc3 &^ bc2)
		a[12] = bc2 ^ (bc4 &^ bc3)
		a[18] = bc3 ^ (bc0 &^ bc4)
		a[24] = bc4 ^ (bc1 &^ bc0)

		t = a[10] ^ d0
		bc2 = t<<3 | t>>(64-3)
		t = a[16] ^ d1
		bc3 = t<<45 | t>>(64-45)
		t = a[22] ^ d2
		bc4 = t<<61 | t>>(64-61)
		t = a[3] ^ d3
		bc0 = t<<28 | t>>(64-28)
		t = a[9] ^ d4
		bc1 = t<<20 | t>>(64-20)
		a[10] = bc0 ^ (bc2 &^ bc1)
		a[16] = bc1 ^ (bc3 &^ bc2)
		a[22] = bc2 ^ (bc4 &^ bc3)
		a[3] = bc3 ^ (bc0 &^ bc4)
		a[9] = bc4 ^ (bc1 &^ bc0)

		t = a[20] ^ d0
		bc4 = t<<18 | t>>(64-18)
		t = a[1] ^ d1
		bc0 = t<<1 | t>>(64-1)
		t = a[7] ^ d2
		bc1 = t<<6 | t>>(64-6)
		t = a[13] ^ d3
		bc2 = t<<25 | t>>(64-25)
		t = a[19] ^ d4
		bc3 = t<<8 | t>>(64-8)
		a[20] = bc0 ^ (bc2 &^ bc1)
		a[1] = bc1 ^ (bc3 &^ bc2)
		a[7] = bc2 ^ (bc4 &^ bc3)
		a[13] = bc3 ^ (bc0 &^ bc4)
		a[19] = bc4 ^ (bc1 &^ bc0)

		t = a[5] ^ d0
		bc1 = t<<36 | t>>(64-36)
		t = a[11] ^ d1
		bc2 = t<<10 | t>>(64-10)
		t = a[17] ^ d2
		bc3 = t<<15 | t>>(64-15)
		t = a[23] ^ d3
		bc4 = t<<56 | t>>(64-56)
		t = a[4] ^ d4
		bc0 = t<<27 | t>>(64-27)
		a[5] = bc0 ^ (bc2 &^ bc1)
		a[11] = bc1 ^ (bc3 &^ bc2)
		a[17] = bc2 ^ (bc4 &^ bc3)
		a[23] = bc3 ^ (bc0 &^ bc4)
		a[4] = bc4 ^ (bc1 &^ bc0)

		t = a[15] ^ d0
		bc3 = t<<41 | t>>(64-41)
		t = a[21] ^ d1
		bc4 = t<<2 | t>>(64-2)
		t = a[2] ^ d2
		bc0 = t<<62 | t>>(64-62)
		t = a[8] ^ d3
		bc1 = t<<55 | t>>(64-55)
		t = a[14] ^ d4
		bc2 = t<<39 | t>>(64-39)
		a[15] = bc0 ^ (bc2 &^ bc1)
		a[21] = bc1 ^ (bc3 &^ bc2)
		a[2] = bc2 ^ (bc4 &^ bc3)
		a[8] = bc3 ^ (bc0 &^ bc4)
		a[14] = bc4 ^ (bc1 &^ bc0)

		// Round 2
		bc0 = a[0] ^ a[5] ^ a[10] ^ a[15] ^ a[20]
		bc1 = a[1] ^ a[6] ^ a[11] ^ a[16] ^ a[21]
		bc2 = a[2] ^ a[7] ^ a[12] ^ a[17] ^ a[22]
		bc3 = a[3] ^ a[8] ^ a[13] ^ a[18] ^ a[23]
		bc4 = a[4] ^ a[9] ^ a[14] ^ a[19] ^ a[24]
		d0 = bc4 ^ (bc1<<1 | bc1>>63)
		d1 = bc0 ^ (bc2<<1 | bc2>>63)
		d2 = bc1 ^ (bc3<<1 | bc3>>63)
		d3 = bc2 ^ (bc4<<1 | bc4>>63)
		d4 = bc3 ^ (bc0<<1 | bc0>>63)

		bc0 = a[0] ^ d0
		t = a[16] ^ d1
		bc1 = t<<44 | t>>(64-44)
		t = a[7] ^ d2
		bc2 = t<<43 | t>>(64-43)
		t = a[23] ^ d3
		bc3 = t<<21 | t>>(64-21)
		t = a[14] ^ d4
		bc4 = t<<14 | t>>(64-14)
		a[0] = bc0 ^ (bc2 &^ bc1) ^ rc[i+1]
		a[16] = bc1 ^ (bc3 &^ bc2)
		a[7] = bc2 ^ (bc4 &^ bc3)
		a[23] = bc3 ^ (bc0 &^ bc4)
		a[14] = bc4 ^ (bc1 &^ bc0)

		t = a[20] ^ d0
		bc2 = t<<3 | t>>(64-3)
		t = a[11] ^ d1
		bc3 = t<<45 | t>>(64-45)
		t = a[2] ^ d2
		bc4 = t<<61 | t>>(64-61)
		t = a[18] ^ d3
		bc0 = t<<28 | t>>(64-28)
		t = a[9] ^ d4
		bc1 = t<<20 | t>>(64-20)
		a[20] = bc0 ^ (bc2 &^ bc1)
		a[11] = bc1 ^ (bc3 &^ bc2)
		a[2] = bc2 ^ (bc4 &^ bc3)
		a[18] = bc3 ^ (bc0 &^ bc4)
		a[9] = bc4 ^ (bc1 &^ bc0)

		t = a[15] ^ d0
		bc4 = t<<18 | t>>(64-18)
		t = a[6] ^ d1
		bc0 = t<<1 | t>>(64-1)
		t = a[22] ^ d2
		bc1 = t<<6 | t>>(64-6)
		t = a[13] ^ d3
		bc2 = t<<25 | t>>(64-25)
		t = a[4] ^ d4
		bc3 = t<<8 | t>>(64-8)
		a[15] = bc0 ^ (bc2 &^ bc1)
		a[6] = bc1 ^ (bc3 &^ bc2)
		a[22] = bc2 ^ (bc4 &^ bc3)
		a[13] = bc3 ^ (bc0 &^ bc4)
		a[4] = bc4 ^ (bc1 &^ bc0)

		t = a[10] ^ d0
		bc1 = t<<36 | t>>(64-36)
		t = a[1] ^ d1
		bc2 = t<<10 | t>>(64-10)
		t = a[17] ^ d2
		bc3 = t<<15 | t>>(64-15)
		t = a[8] ^ d3
		bc4 = t<<56 | t>>(64-56)
		t = a[24] ^ d4
		bc0 = t<<27 | t>>(64-27)
		a[10] = bc0 ^ (bc2 &^ bc1)
		a[1] = bc1 ^ (bc3 &^ bc2)
		a[17] = bc2 ^ (bc4 &^ bc3)
		a[8] = bc3 ^ (bc0 &^ bc4)
		a[24] = bc4 ^ (bc1 &^ bc0)

		t = a[5] ^ d0
		bc3 = t<<41 | t>>(64-41)
		t = a[21] ^ d1
		bc4 = t<<2 | t>>(64-2)
		t = a[12] ^ d2
		bc0 = t<<62 | t>>(64-62)
		t = a[3] ^ d3
		bc1 = t<<55 | t>>(64-55)
		t = a[19] ^ d4
		bc2 = t<<39 | t>>(64-39)
		a[5] = bc0 ^ (bc2 &^ bc1)
		a[21] = bc1 ^ (bc3 &^ bc2)
		a[12] = bc2 ^ (bc4 &^ bc3)
		a[3] = bc3 ^ (bc0 &^ bc4)
		a[19] = bc4 ^ (bc1 &^ bc0)

		// Round 3
		bc0 = a[0] ^ a[5] ^ a[10] ^ a[15] ^ a[20]
		bc1 = a[1] ^ a[6] ^ a[11] ^ a[16] ^ a[21]
		bc2 = a[2] ^ a[7] ^ a[12] ^ a[17] ^ a[22]
		bc3 = a[3] ^ a[8] ^ a[13] ^ a[18] ^ a[23]
		bc4 = a[4] ^ a[9] ^ a[14] ^ a[19] ^ a[24]
		d0 = bc4 ^ (bc1<<1 | bc1>>63)
		d1 = bc0 ^ (bc2<<1 | bc2>>63)
		d2 = bc1 ^ (bc3<<1 | bc3>>63)
		d3 = bc2 ^ (bc4<<1 | bc4>>63)
		d4 = bc3 ^ (bc0<<1 | bc0>>63)

		bc0 = a[0] ^ d0
		t = a[11] ^ d1
		bc1 = t<<44 | t>>(64-44)
		t = a[22] ^ d2
		bc2 = t<<43 | t>>(64-43)
		t = a[8] ^ d3
		bc3 = t<<21 | t>>(64-21)
		t = a[19] ^ d4
		bc4 = t<<14 | t>>(64-14)
		a[0] = bc0 ^ (bc2 &^ bc1) ^ rc[i+2]
		a[11] = bc1 ^ (bc3 &^ bc2)
		a[22] = bc2 ^ (bc4 &^ bc3)
		a[8] = bc3 ^ (bc0 &^ bc4)
		a[19] = bc4 ^ (bc1 &^ bc0)

		t = a[15] ^ d0
		bc2 = t<<3 | t>>(64-3)
		t = a[1] ^ d1
		bc3 = t<<45 | t>>(64-45)
		t = a[12] ^ d2
		bc4 = t<<61 | t>>(64-61)
		t = a[23] ^ d3
		bc0 = t<<28 | t>>(64-28)
		t = a[9] ^ d4
		bc1 = t<<20 | t>>(64-20)
		a[15] = bc0 ^ (bc2 &^ bc1)
		a[1] = bc1 ^ (bc3 &^ bc2)
		a[12] = bc2 ^ (bc4 &^ bc3)
		a[23] = bc3 ^ (bc0 &^ bc4)
		a[9] = bc4 ^ (bc1 &^ bc0)

		t = a[5] ^ d0
		bc4 = t<<18 | t>>(64-18)
		t = a[16] ^ d1
		bc0 = t<<1 | t>>(64-1)
		t = a[2] ^ d2
		bc1 = t<<6 | t>>(64-6)
		t = a[13] ^ d3
		bc2 = t<<25 | t>>(64-25)
		t = a[24] ^ d4
		bc3 = t<<8 | t>>(64-8)
		a[5] = bc0 ^ (bc2 &^ bc1)
		a[16] = bc1 ^ (bc3 &^ bc2)
		a[2] = bc2 ^ (bc4 &^ bc3)
		a[13] = bc3 ^ (bc0 &^ bc4)
		a[24] = bc4 ^ (bc1 &^ bc0)

		t = a[20] ^ d0
		bc1 = t<<36 | t>>(64-36)
		t = a[6] ^ d1
		bc2 = t<<10 | t>>(64-10)
		t = a[17] ^ d2
		bc3 = t<<15 | t>>(64-15)
		t = a[3] ^ d3
		bc4 = t<<56 | t>>(64-56)
		t = a[14] ^ d4
		bc0 = t<<27 | t>>(64-27)
		a[20] = bc0 ^ (bc2 &^ bc1)
		a[6] = bc1 ^ (bc3 &^ bc2)
		a[17] = bc2 ^ (bc4 &^ bc3)
		a[3] = bc3 ^ (bc0 &^ bc4)
		a[14] = bc4 ^ (bc1 &^ bc0)

		t = a[10] ^ d0
		bc3 = t<<41 | t>>(64-41)
		t = a[21] ^ d1
		bc4 = t<<2 | t>>(64-2)
		t = a[7] ^ d2
		bc0 = t<<62 | t>>(64-62)
		t = a[18] ^ d3
		bc1 = t<<55 | t>>(64-55)
		t = a[4] ^ d4
		bc2 = t<<39 | t>>(64-39)
		a[10] = bc0 ^ (bc2 &^ bc1)
		a[21] = bc1 ^ (bc3 &^ bc2)
		a[7] = bc2 ^ (bc4 &^ bc3)
		a[18] = bc3 ^ (bc0 &^ bc4)
		a[4] = bc4 ^ (bc1 &^ bc0)

		// Round 4
		bc0 = a[0] ^ a[5] ^ a[10] ^ a[15] ^ a[20]
		bc1 = a[1] ^ a[6] ^ a[11] ^ a[16] ^ a[21]
		bc2 = a[2] ^ a[7] ^ a[12] ^ a[17] ^ a[22]
		bc3 = a[3] ^ a[8] ^ a[13] ^ a[18] ^ a[23]
		bc4 = a[4] ^ a[9] ^ a[14] ^ a[19] ^ a[24]
		d0 = bc4 ^ (bc1<<1 | bc1>>63)
		d1 = bc0 ^ (bc2<<1 | bc2>>63)
		d2 = bc1 ^ (bc3<<1 | bc3>>63)
		d3 = bc2 ^ (bc4<<1 | bc4>>63)
		d4 = bc3 ^ (bc0<<1 | bc0>>63)

		bc0 = a[0] ^ d0
		t = a[1] ^ d1
		bc1 = t<<44 | t>>(64-44)
		t = a[2] ^ d2
		bc2 = t<<43 | t>>(64-43)
		t = a[3] ^ d3
		bc3 = t<<21 | t>>(64-21)
		t = a[4] ^ d4
		bc4 = t<<14 | t>>(64-14)
		a[0] = bc0 ^ (bc2 &^ bc1) ^ rc[i+3]
		a[1] = bc1 ^ (bc3 &^ bc2)
		a[2] = bc2 ^ (bc4 &^ bc3)
		a[3] = bc3 ^ (bc0 &^ bc4)
		a[4] = bc4 ^ (bc1 &^ bc0)

		t = a[5] ^ d0
		bc2 = t<<3 | t>>(64-3)
		t = a[6] ^ d1
		bc3 = t<<45 | t>>(64-45)
		t = a[7] ^ d2
		bc4 = t<<61 | t>>(64-61)
		t = a[8] ^ d3
		bc0 = t<<28 | t>>(64-28)
		t = a[9] ^ d4
		bc1 = t<<20 | t>>(64-20)
		a[5] = bc0 ^ (bc2 &^ bc1)
		a[6] = bc1 ^ (bc3 &^ bc2)
		a[7] = bc2 ^ (bc4 &^ bc3)
		a[8] = bc3 ^ (bc0 &^ bc4)
		a[9] = bc4 ^ (bc1 &^ bc0)

		t = a[10] ^ d0
		bc4 = t<<18 | t>>(64-18)
		t = a[11] ^ d1
		bc0 = t<<1 | t>>(64-1)
		t = a[12] ^ d2
		bc1 = t<<6 | t>>(64-6)
		t = a[13] ^ d3
		bc2 = t<<25 | t>>(64-25)
		t = a[14] ^ d4
		bc3 = t<<8 | t>>(64-8)
		a[10] = bc0 ^ (bc2 &^ bc1)
		a[11] = bc1 ^ (bc3 &^ bc2)
		a[12] = bc2 ^ (bc4 &^ bc3)
		a[13] = bc3 ^ (bc0 &^ bc4)
		a[14] = bc4 ^ (bc1 &^ bc0)

		t = a[15] ^ d0
		bc1 = t<<36 | t>>(64-36)
		t = a[16] ^ d1
		bc2 = t<<10 | t>>(64-10)
		t = a[17] ^ d2
		bc3 = t<<15 | t>>(64-15)
		t = a[18] ^ d3
		bc4 = t<<56 | t>>(64-56)
		t = a[19] ^ d4
		bc0 = t<<27 | t>>(64-27)
		a[15] = bc0 ^ (bc2 &^ bc1)
		a[16] = bc1 ^ (bc3 &^ bc2)
		a[17] = bc2 ^ (bc4 &^ bc3)
		a[18] = bc3 ^ (bc0 &^ bc4)
		a[19] = bc4 ^ (bc1 &^ bc0)

		t = a[20] ^ d0
		bc3 = t<<41 | t>>(64-41)
		t = a[21] ^ d1
		bc4 = t<<2 | t>>(64-2)
		t = a[22] ^ d2
		bc0 = t<<62 | t>>(64-62)
		t = a[23] ^ d3
		bc1 = t<<55 | t>>(64-55)
		t = a[24] ^ d4
		bc2 = t<<39 | t>>(64-39)
		a[20] = bc0 ^ (bc2 &^ bc1)
		a[21] = bc1 ^ (bc3 &^ bc2)
		a[22] = bc2 ^ (bc4 &^ bc3)
		a[23] = bc3 ^ (bc0 &^ bc4)
		a[24] = bc4 ^ (bc1 &^ bc0)
	}
}
