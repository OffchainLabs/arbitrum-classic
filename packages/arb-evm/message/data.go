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

package message

import (
	"fmt"
	ethmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"math/big"
)

func extractUInt256(data []byte) (*big.Int, []byte) {
	val := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	return val, data
}

func extractAddress(data []byte) (common.Address, []byte) {
	data = data[12:] // Skip first 12 bytes of 32 byte address data
	var addr common.Address
	copy(addr[:], data[:])
	data = data[20:]
	return addr, data
}

func addressData(addr common.Address) []byte {
	ret := make([]byte, 0, 32)
	ret = append(ret, make([]byte, 12)...)
	ret = append(ret, addr[:]...)
	return ret
}

func marshaledBytesHash(data []byte) common.Hash {
	var ret common.Hash
	copy(ret[:], ethmath.U256Bytes(big.NewInt(int64(len(data)))))
	chunks := make([]common.Hash, 0)
	for len(data) > 0 {
		var nextVal common.Hash
		copy(nextVal[:], data[:])
		chunks = append(chunks, nextVal)
		if len(data) <= 32 {
			break
		}
		data = data[32:]
	}

	for i := range chunks {
		ret = hashing.SoliditySHA3(
			hashing.Bytes32(ret),
			hashing.Bytes32(chunks[len(chunks)-1-i]),
		)
	}
	return ret
}

func encodeAmount(amount *big.Int) ([]byte, error) {
	mod := big.NewInt(10)
	zero := big.NewInt(0)
	exp := byte(0)
	for amount.Cmp(zero) > 0 && new(big.Int).Mod(amount, mod).Cmp(zero) == 0 {
		amount = amount.Div(amount, mod)
		exp++
	}
	amountData, err := rlp.EncodeToBytes(amount)
	if err != nil {
		return nil, err
	}

	if amount.Cmp(zero) == 0 {
		return amountData, nil
	}
	return append(amountData, exp), nil
}

func decodeAmount(r io.Reader) (*big.Int, error) {
	a := new(big.Int)
	if err := rlp.Decode(r, a); err != nil {
		return nil, err
	}
	if a.Cmp(big.NewInt(0)) == 0 {
		return a, nil
	}
	bData := make([]byte, 1)
	n, _ := r.Read(bData)
	if n != 1 {
		return nil, errors.New("not enough data for exponent in amount")
	}
	amount := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(bData[0])), nil)
	return amount.Mul(amount, a), nil
}

type CompressedAddress interface {
	fmt.Stringer
	isCompressedAddress()
	Encode() ([]byte, error)
}

type CompressedAddressIndex struct {
	*big.Int
}

func (c CompressedAddressIndex) isCompressedAddress() {}

func (c CompressedAddressIndex) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(c.Int)
}

func (c CompressedAddressIndex) String() string {
	return fmt.Sprintf("Index[%v]", c)
}

type CompressedAddressFull struct {
	common.Address
}

func (c CompressedAddressFull) isCompressedAddress() {}

func (c CompressedAddressFull) Encode() ([]byte, error) {
	return rlp.EncodeToBytes(c.Bytes())
}

func (c CompressedAddressFull) String() string {
	return fmt.Sprintf("Address[%v]", c.Hex())
}

func DecodeAddress(r io.Reader) (CompressedAddress, error) {
	addressBytes := make([]byte, 0)
	if err := rlp.Decode(r, &addressBytes); err != nil {
		return nil, errors.Wrap(err, "couldn't parse address")
	}

	if len(addressBytes) == 0 {
		return nil, nil
	}

	if len(addressBytes) < 20 {
		return CompressedAddressIndex{new(big.Int).SetBytes(addressBytes)}, nil
	}

	if len(addressBytes) == 20 {
		var address common.Address
		copy(address[:], addressBytes)
		return CompressedAddressFull{address}, nil
	}

	return nil, errors.Errorf("unexpected address length %v", len(addressBytes))
}

func encodeUnsignedTx(tx CompressedTx) ([]byte, error) {
	nonceData, err := rlp.EncodeToBytes(tx.SequenceNum)
	if err != nil {
		return nil, err
	}
	gasPriceData, err := rlp.EncodeToBytes(tx.GasPrice)
	if err != nil {
		return nil, err
	}
	gasLimitData, err := rlp.EncodeToBytes(tx.GasLimit)
	if err != nil {
		return nil, err
	}
	paymentData, err := encodeAmount(tx.Payment)
	if err != nil {
		return nil, err
	}

	var data []byte
	data = append(data, nonceData...)
	data = append(data, gasPriceData...)
	data = append(data, gasLimitData...)
	if tx.To == nil {
		data = append(data, 0x80)
	} else {
		destData, err := tx.To.Encode()
		if err != nil {
			return nil, err
		}
		data = append(data, destData...)
	}
	data = append(data, paymentData...)
	data = append(data, tx.Calldata...)
	return data, nil
}

func decodeCompressedTx(r io.Reader) (CompressedTx, error) {
	nonce := new(big.Int)
	if err := rlp.Decode(r, nonce); err != nil {
		return CompressedTx{}, err
	}

	gasPrice := new(big.Int)
	if err := rlp.Decode(r, gasPrice); err != nil {
		return CompressedTx{}, err
	}

	gasLimit := new(big.Int)
	if err := rlp.Decode(r, gasLimit); err != nil {
		return CompressedTx{}, err
	}

	address, err := DecodeAddress(r)
	if err != nil {
		return CompressedTx{}, err
	}

	payment, err := decodeAmount(r)
	if err != nil {
		return CompressedTx{}, err
	}

	calldata, err := ioutil.ReadAll(r)
	if err != nil {
		return CompressedTx{}, err
	}
	return CompressedTx{
		SequenceNum: nonce,
		GasPrice:    gasPrice,
		GasLimit:    gasLimit,
		To:          address,
		Payment:     payment,
		Calldata:    calldata,
	}, nil

}

func encodeECDSASig(v byte, r, s *big.Int) []byte {
	data := make([]byte, 0, 65)
	data = append(data, ethmath.U256Bytes(new(big.Int).Set(r))...)
	data = append(data, ethmath.U256Bytes(new(big.Int).Set(s))...)
	data = append(data, v)
	return data
}

func decodeECDSASig(rd io.Reader) (v byte, r, s *big.Int, err error) {
	rData := make([]byte, 32)
	if count, _ := rd.Read(rData); count != len(rData) {
		return 0, nil, nil, errors.New("couldn't read r")
	}
	sData := make([]byte, 32)
	if count, _ := rd.Read(sData); count != len(sData) {
		return 0, nil, nil, errors.New("couldn't read s")
	}
	vData := make([]byte, 1)
	if count, _ := rd.Read(vData); count != len(vData) {
		return 0, nil, nil, errors.New("couldn't read v")
	}
	return vData[0], new(big.Int).SetBytes(rData), new(big.Int).SetBytes(sData), nil
}
