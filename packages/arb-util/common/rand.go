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

package common

import (
	"math/big"
	"math/rand"
)

func fillRand(slice []byte) {
	_, err := rand.Read(slice)
	if err != nil {
		panic(err)
	}
}

func RandAddress() Address {
	var v Address
	fillRand(v[:])
	return v
}

func RandHash() Hash {
	var v Hash
	fillRand(v[:])
	return v
}

func RandBytes(size int) []byte {
	v := make([]byte, size)
	fillRand(v[:])
	return v
}

func RandBigInt() *big.Int {
	raw := RandBytes(32)
	return new(big.Int).SetBytes(raw)
}

func RandBigIntBelowBound(max *big.Int) *big.Int {
	x := RandBigInt()
	for ; x.Cmp(max) >= 0; x = RandBigInt() {
	}
	return x
}
