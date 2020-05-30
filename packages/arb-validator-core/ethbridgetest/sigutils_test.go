/*
 * Copyright 2019, Offchain Labs, Inc.
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
	"bytes"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
)

func TestSigParse(t *testing.T) {
	var messageHash [32]byte
	messageHash[0] = 65
	messageHash[2] = 23
	messageHash[4] = 68
	messageHash[6] = 87
	messageHash[31] = 12

	ethMessageHash := hashing.SoliditySHA3WithPrefix(messageHash[:])

	privateKey, err := crypto.HexToECDSA(privHex)
	if err != nil {
		t.Fatal(err)
	}

	sigBytes, err := crypto.Sign(ethMessageHash.Bytes(), privateKey)
	if err != nil {
		t.Fatal(err)
	}
	parsedSig, err := sigTester.ParseSignature(
		nil,
		sigBytes,
		big.NewInt(0),
	)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(sigBytes[0:32], parsedSig.R[:]) {
		t.Error("Parsed signature has incorrect R component")
	}

	if !bytes.Equal(sigBytes[32:64], parsedSig.S[:]) {
		t.Error("Parsed signature has incorrect S component")
	}

	if sigBytes[64]+27 != parsedSig.V {
		t.Error("Parsed signature has incorrect V component")
	}
}

func TestSigRecover(t *testing.T) {
	var messageHash [32]byte
	messageHash[0] = 65
	messageHash[2] = 23
	messageHash[4] = 68
	messageHash[6] = 87
	messageHash[31] = 12

	ethMessageHash := hashing.SoliditySHA3WithPrefix(messageHash[:])

	privateKey, err := crypto.HexToECDSA(privHex)
	if err != nil {
		t.Fatal(err)
	}

	sigBytes, err := crypto.Sign(ethMessageHash.Bytes(), privateKey)
	if err != nil {
		t.Fatal(err)
	}
	signer, err := sigTester.RecoverAddressFromData(
		nil,
		messageHash,
		sigBytes,
		big.NewInt(0),
	)
	if err != nil {
		t.Fatal(err)
	}
	if signer != auth.From {
		t.Fatal("Message signer not calculated correctly: got", hexutil.Encode(signer[:]), "instead of", hexutil.Encode(auth.From[:]))
	}
}
