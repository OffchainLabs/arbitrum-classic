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

package arbostest

import (
	"bytes"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/arbostestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-rpc-node/snapshot"
	"github.com/offchainlabs/arbitrum/packages/arb-util/arbos"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
)

func generateFib(t *testing.T, val *big.Int) []byte {
	t.Helper()
	fib, err := abi.JSON(strings.NewReader(arbostestcontracts.FibonacciABI))
	failIfError(t, err)
	return makeFuncData(t, fib.Methods["generateFib"], val)
}

func makeTxCountCall(account common.Address) message.L2Message {
	call := message.Call{
		BasicTx: message.BasicTx{
			MaxGas:      big.NewInt(10000000),
			GasPriceBid: big.NewInt(0),
			DestAddress: common.NewAddressFromEth(arbos.ARB_SYS_ADDRESS),
			Payment:     big.NewInt(0),
			Data:        snapshot.TransactionCountData(account),
		},
	}
	return message.NewSafeL2Message(call)
}

func checkTxCountResult(t *testing.T, res *evm.TxResult, correctCount *big.Int) {
	t.Helper()
	succeededTxCheck(t, res)
	txCount, err := snapshot.ParseTransactionCountResult(res)
	failIfError(t, err)
	if correctCount.Cmp(txCount) != 0 {
		t.Fatal("unexpected tx count")
	}
}

func TestTransactionCount(t *testing.T) {
	randDest := common.RandAddress()

	// Valid contract deployment
	tx1 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(0),
		DestAddress: common.Address{},
		Payment:     big.NewInt(0),
		Data:        hexutil.MustDecode(arbostestcontracts.FibonacciBin),
	}

	// Valid value tranfer to EOA
	tx2 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(1),
		DestAddress: randDest,
		Payment:     big.NewInt(300),
		Data:        []byte{},
	}

	// Invalid sequencer number
	tx3 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(3),
		DestAddress: randDest,
		Payment:     big.NewInt(10),
		Data:        []byte{},
	}

	// Insufficient balance
	tx4 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: randDest,
		Payment:     big.NewInt(30000),
		Data:        []byte{},
	}

	// Valid transaction to contract
	tx5 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(2),
		DestAddress: connAddress1,
		Payment:     big.NewInt(300),
		Data:        generateFib(t, big.NewInt(20)),
	}

	// Transaction to contract with incorrect sequence number
	tx6 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(4),
		DestAddress: connAddress1,
		Payment:     big.NewInt(300),
		Data:        generateFib(t, big.NewInt(20)),
	}

	// Transaction to contract with insufficient balance
	tx7 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: big.NewInt(3),
		DestAddress: connAddress1,
		Payment:     big.NewInt(100000),
		Data:        generateFib(t, big.NewInt(20)),
	}

	// Transaction to contract with insufficient balance
	tx8 := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(1000),
		SequenceNum: big.NewInt(3),
		DestAddress: connAddress1,
		Payment:     big.NewInt(0),
		Data:        generateFib(t, big.NewInt(20)),
	}

	chainTime := inbox.ChainTime{
		BlockNum:  common.NewTimeBlocksInt(0),
		Timestamp: big.NewInt(0),
	}

	messages := []inbox.InboxMessage{
		message.NewInboxMessage(initMsg(), chain, big.NewInt(0), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(1), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.Eth{Dest: sender, Value: big.NewInt(1000)}, sender, big.NewInt(2), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(3), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx1), sender, big.NewInt(4), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(5), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx2), sender, big.NewInt(6), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(7), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx3), sender, big.NewInt(8), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(9), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx4), sender, big.NewInt(10), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(11), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx5), sender, big.NewInt(12), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(13), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx6), sender, big.NewInt(14), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(15), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx7), sender, big.NewInt(16), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(17), big.NewInt(0), chainTime),
		message.NewInboxMessage(message.NewSafeL2Message(tx8), sender, big.NewInt(18), big.NewInt(0), chainTime),
		message.NewInboxMessage(makeTxCountCall(sender), common.Address{}, big.NewInt(19), big.NewInt(0), chainTime),
	}

	logs, _, _, _ := runAssertion(t, messages, len(messages)-2, 0)
	results := processTxResults(t, logs)

	checkTxCountResult(t, results[0], big.NewInt(0))

	// Deposit doesn't increase tx count
	checkTxCountResult(t, results[1], big.NewInt(0))

	// Contract deployment increases tx count
	succeededTxCheck(t, results[2])
	checkTxCountResult(t, results[3], big.NewInt(1))

	// Payment to EOA increases tx count
	succeededTxCheck(t, results[4])
	checkTxCountResult(t, results[5], big.NewInt(2))

	// Payment to EOA with incorrect sequence number shouldn't increase tx count
	txResultCheck(t, results[6], evm.BadSequenceCode)
	checkTxCountResult(t, results[7], big.NewInt(2))

	// Payment to EOA with insufficient funds shouldn't increase tx count
	txResultCheck(t, results[8], evm.InsufficientTxFundsCode)
	checkTxCountResult(t, results[9], big.NewInt(2))

	// Contract to contract increases tx count
	succeededTxCheck(t, results[10])
	checkTxCountResult(t, results[11], big.NewInt(3))

	// Tx call with incorrect sequence number doesn't affect the count
	txResultCheck(t, results[12], evm.BadSequenceCode)
	checkTxCountResult(t, results[13], big.NewInt(3))

	// Tx call with insufficient balance doesn't affect the count
	txResultCheck(t, results[14], evm.InsufficientTxFundsCode)
	checkTxCountResult(t, results[15], big.NewInt(3))

	// Tx call with insufficient gas funds doesn't affect the count
	txResultCheck(t, results[16], evm.InsufficientGasFundsCode)
	checkTxCountResult(t, results[17], big.NewInt(3))

	t.Log(crypto.CreateAddress(ethcommon.HexToAddress("0x3fab184622dc19b6109349b94811493bf2a45362"), 0).Hex())
}

func makeSyscallTx(data []byte, seq *big.Int, addr common.Address) message.Message {
	tx := message.Transaction{
		MaxGas:      big.NewInt(10000000),
		GasPriceBid: big.NewInt(0),
		SequenceNum: seq,
		DestAddress: addr,
		Payment:     big.NewInt(0),
		Data:        data,
	}
	return message.NewSafeL2Message(tx)
}

func TestAddressTable(t *testing.T) {
	targetAddress := common.RandAddress()
	targetAddress2 := common.RandAddress()
	targetAddress3 := common.RandAddress()
	unregisteredAddress := common.RandAddress()

	encodedIndex2, err := message.CompressedAddressIndex{Int: big.NewInt(2)}.Encode()
	failIfError(t, err)
	encodedIndex3, err := message.CompressedAddressIndex{Int: big.NewInt(3)}.Encode()
	failIfError(t, err)
	encodedAddress3, err := message.CompressedAddressFull{Address: targetAddress3}.Encode()
	failIfError(t, err)

	addressTableCalls := [][]byte{
		// lookup nonexistent key
		snapshot.AddressTableLookupData(targetAddress),
		// register that key
		snapshot.AddressTableRegisterData(targetAddress),
		// now lookup the same key
		snapshot.AddressTableLookupData(targetAddress),
		// register a different ket
		snapshot.AddressTableRegisterData(targetAddress2),
		// call register on the first key again
		snapshot.AddressTableRegisterData(targetAddress),
		// Check to make sure that key exists
		snapshot.AddressTableAddressExistsData(targetAddress),
		// Check to make sure a different key doesn't exist
		snapshot.AddressTableAddressExistsData(unregisteredAddress),
		// Check to make sure the address table is the right size
		snapshot.AddressTableSizeData(),
		// Lookup the address with a registered index
		snapshot.AddressTableLookupIndexData(big.NewInt(2)),
		// Lookup the address with an index which is too high
		snapshot.AddressTableLookupIndexData(big.NewInt(3)),
		// Decompress a compressed address index for an address that exists
		snapshot.AddressTableDecompressData(encodedIndex2, big.NewInt(0)),
		// Decompress a compressed address index for an address that does't exist
		snapshot.AddressTableDecompressData(encodedIndex3, big.NewInt(0)),
		// Decompress a compressed full address
		snapshot.AddressTableDecompressData(encodedAddress3, big.NewInt(0)),
		// Compress an unregistered address
		snapshot.AddressTableCompressData(unregisteredAddress),
		// Compress a registerted address
		snapshot.AddressTableCompressData(targetAddress2),
	}

	senderSeq := int64(0)
	var messages []message.Message
	for _, msg := range addressTableCalls {
		messages = append(messages, makeSyscallTx(msg, big.NewInt(senderSeq), common.NewAddressFromEth(arbos.ARB_ADDRESS_TABLE_ADDRESS)))
		senderSeq++
	}

	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), len(addressTableCalls), 0)
	results := processTxResults(t, logs)

	revertedTxCheck(t, results[0])

	succeededTxCheck(t, results[1])
	index := returnedInt(t, results[1])
	if index.Cmp(big.NewInt(1)) != 0 {
		t.Error("address registered at unexpected index", index)
	}

	succeededTxCheck(t, results[2])
	lookedUpIndex := returnedInt(t, results[2])
	if lookedUpIndex.Cmp(index) != 0 {
		t.Error("looked up index doesn't match registered index")
	}

	succeededTxCheck(t, results[3])
	index2 := returnedInt(t, results[3])
	if index2.Cmp(big.NewInt(2)) != 0 {
		t.Error("second address registered at unexpected index", index2)
	}

	succeededTxCheck(t, results[4])
	reRegisteredIndex := returnedInt(t, results[4])
	if index.Cmp(reRegisteredIndex) != 0 {
		t.Error("second call to register address returned different index", index, reRegisteredIndex)
	}

	succeededTxCheck(t, results[5])
	doesExist := returnedInt(t, results[5])
	if doesExist.Cmp(big.NewInt(1)) != 0 {
		t.Error("address should exist")
	}

	succeededTxCheck(t, results[6])
	doesntExist := returnedInt(t, results[6])
	if doesntExist.Cmp(big.NewInt(0)) != 0 {
		t.Error("address shouldn't exist")
	}

	succeededTxCheck(t, results[7])
	tableSize := returnedInt(t, results[7])
	if tableSize.Cmp(big.NewInt(3)) != 0 {
		t.Error("wrong table size", tableSize)
	}

	succeededTxCheck(t, results[8])
	index2LookedUpAddress := returnedInt(t, results[8])
	if index2LookedUpAddress.Cmp(new(big.Int).SetBytes(targetAddress2.Bytes())) != 0 {
		t.Error("wrong address at index", index2LookedUpAddress)
	}

	revertedTxCheck(t, results[9])

	succeededTxCheck(t, results[10])
	index2DecompressedAddress, offset := returnedDecompressed(t, results[10])
	if index2DecompressedAddress != targetAddress2 {
		t.Error("wrong address decompressed", index2LookedUpAddress)
	}
	if offset.Cmp(big.NewInt(int64(len(encodedIndex2)))) != 0 {
		t.Error("unexpected offset after decompressing", offset)
	}

	revertedTxCheck(t, results[11])

	succeededTxCheck(t, results[12])
	address3DecompressedAddress, offset := returnedDecompressed(t, results[12])
	if address3DecompressedAddress != targetAddress3 {
		t.Error("wrong address decompressed", address3DecompressedAddress)
	}
	if offset.Cmp(big.NewInt(int64(len(encodedAddress3)))) != 0 {
		t.Error("unexpected offset after decompressing", offset)
	}

	succeededTxCheck(t, results[13])
	compressedUnregistereddAddress := returnedBytes(t, results[13])
	decoded, err := message.DecodeAddress(bytes.NewReader(compressedUnregistereddAddress))
	failIfError(t, err)
	addr, ok := decoded.(message.CompressedAddressFull)
	if !ok {
		t.Fatal("decoded to wrong type of address")
	}
	if addr.Address != unregisteredAddress {
		t.Fatal("got wrong address")
	}

	succeededTxCheck(t, results[14])
	compressedRegisteredAddress := returnedBytes(t, results[14])
	decoded, err = message.DecodeAddress(bytes.NewReader(compressedRegisteredAddress))
	failIfError(t, err)
	addressIndex, ok := decoded.(message.CompressedAddressIndex)
	if !ok {
		t.Fatal("decoded to wrong type of address")
	}
	if addressIndex.Cmp(big.NewInt(2)) != 0 {
		t.Fatal("got wrong address")
	}
}

func TestArbSysBLS(t *testing.T) {
	x0a, x1a, y0a, y1a := common.RandBigInt(), common.RandBigInt(), common.RandBigInt(), common.RandBigInt()
	x0b, x1b, y0b, y1b := common.RandBigInt(), common.RandBigInt(), common.RandBigInt(), common.RandBigInt()

	arbSysCalls := [][]byte{
		// Lookup the key for the sender who hasn't registered
		snapshot.GetBLSPublicKeyData(sender),
		// Register a key
		snapshot.RegisterBLSKeyData(x0a, x1a, y0a, y1a),
		// Lookup the registered key
		snapshot.GetBLSPublicKeyData(sender),
		// Replace the currently registered key with a different one
		snapshot.RegisterBLSKeyData(x0b, x1b, y0b, y1b),
		// Make sure when we lookup, we get the new key
		snapshot.GetBLSPublicKeyData(sender),
	}

	senderSeq := int64(0)
	var messages []message.Message
	for _, msg := range arbSysCalls {
		messages = append(messages, makeSyscallTx(msg, big.NewInt(senderSeq), common.NewAddressFromEth(arbos.ARB_BLS_ADDRESS)))
		senderSeq++
	}

	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), len(arbSysCalls), 0)
	results := processTxResults(t, logs)

	revertedTxCheck(t, results[0])

	succeededTxCheck(t, results[1])
	if len(results[1].ReturnData) != 0 {
		t.Fatal("shouldn't have return data")
	}

	succeededTxCheck(t, results[2])
	x0aCheck, x1aCheck, y0aCheck, y1aCheck := returnedBLSKey(t, results[2])
	if x0aCheck.Cmp(x0a) != 0 || x1aCheck.Cmp(x1a) != 0 || y0aCheck.Cmp(y0a) != 0 || y1aCheck.Cmp(y1a) != 0 {
		t.Fatal("got wrong key")
	}

	succeededTxCheck(t, results[3])
	if len(results[3].ReturnData) != 0 {
		t.Fatal("shouldn't have return data")
	}

	succeededTxCheck(t, results[4])
	x0bCheck, x1bCheck, y0bCheck, y1bCheck := returnedBLSKey(t, results[4])
	if x0bCheck.Cmp(x0b) != 0 || x1bCheck.Cmp(x1b) != 0 || y0bCheck.Cmp(y0b) != 0 || y1bCheck.Cmp(y1b) != 0 {
		t.Fatal("got wrong key")
	}
}

func TestArbSysFunctionTable(t *testing.T) {
	functionTable1 := message.FunctionTable{
		message.NewRandomFunctionTableEntry(),
		message.NewRandomFunctionTableEntry(),
	}
	functionTable2 := message.FunctionTable{
		message.NewRandomFunctionTableEntry(),
	}

	functionTableEncoded1, err := functionTable1.Encode()
	failIfError(t, err)
	functionTableEncoded2, err := functionTable2.Encode()
	failIfError(t, err)

	arbSysCalls := [][]byte{
		// Get size of non existent table
		snapshot.FunctionTableSizeData(sender),
		// Get row of non existent table
		snapshot.FunctionTableGetData(sender, big.NewInt(0)),
		// Upload valid table
		snapshot.UploadFunctionTableData(functionTableEncoded1),
		// Get size of uploaded table
		snapshot.FunctionTableSizeData(sender),
		// Get row from uploaded table
		snapshot.FunctionTableGetData(sender, big.NewInt(1)),
		// Upload a new function table
		snapshot.UploadFunctionTableData(functionTableEncoded2),
		// Get new table size
		snapshot.FunctionTableSizeData(sender),
		// Lookup from new table
		snapshot.FunctionTableGetData(sender, big.NewInt(0)),
	}

	senderSeq := int64(0)
	var messages []message.Message
	for _, msg := range arbSysCalls {
		messages = append(messages, makeSyscallTx(msg, big.NewInt(senderSeq), common.NewAddressFromEth(arbos.ARB_FUNCTION_TABLE_ADDRESS)))
		senderSeq++
	}

	logs, _, _, _ := runAssertion(t, makeSimpleInbox(messages), len(arbSysCalls), 0)
	results := processTxResults(t, logs)

	revertedTxCheck(t, results[0])

	revertedTxCheck(t, results[1])

	succeededTxCheck(t, results[2])

	succeededTxCheck(t, results[3])
	table1Size := returnedInt(t, results[3])
	if table1Size.Cmp(big.NewInt(int64(len(functionTable1)))) != 0 {
		t.Error("wrong uploaded function table size")
	}

	succeededTxCheck(t, results[4])
	ft1Row := returnedFunctionTableEntry(t, results[4])
	if !ft1Row.Equals(functionTable1[1]) {
		t.Error("got incorrect function entry from lookup")
	}

	succeededTxCheck(t, results[5])

	succeededTxCheck(t, results[6])
	table2Size := returnedInt(t, results[6])
	if table2Size.Cmp(big.NewInt(int64(len(functionTable2)))) != 0 {
		t.Error("wrong uploaded function table size")
	}

	succeededTxCheck(t, results[7])
	ft2Row := returnedFunctionTableEntry(t, results[7])
	if !ft2Row.Equals(functionTable2[0]) {
		t.Error("got incorrect function entry from lookup")
	}
}

func returnedInt(t *testing.T, res *evm.TxResult) *big.Int {
	if len(res.ReturnData) != 32 {
		t.Fatal("unexpected return data length")
	}
	return new(big.Int).SetBytes(res.ReturnData)
}

func returnedBLSKey(t *testing.T, res *evm.TxResult) (x0, x1, y0, y1 *big.Int) {
	if len(res.ReturnData) != 32*4 {
		t.Fatal("unexpected return data length")
	}
	x0 = new(big.Int).SetBytes(res.ReturnData[:32])
	x1 = new(big.Int).SetBytes(res.ReturnData[32:64])
	y0 = new(big.Int).SetBytes(res.ReturnData[64:96])
	y1 = new(big.Int).SetBytes(res.ReturnData[96:128])
	return
}

func returnedFunctionTableEntry(t *testing.T, res *evm.TxResult) message.FunctionTableEntry {
	if len(res.ReturnData) != 96 {
		t.Fatal("unexpected return data length")
	}
	entry, err := snapshot.ParseFunctionTableGetDataResult(res.ReturnData)
	failIfError(t, err)
	return entry
}

func returnedBytes(t *testing.T, res *evm.TxResult) []byte {
	if len(res.ReturnData) < 32 {
		t.Fatal("return data too short")
	}
	offset := new(big.Int).SetBytes(res.ReturnData[:32])
	if offset.Cmp(big.NewInt(32)) != 0 {
		t.Fatal("expected offset of 32", offset)
	}
	data := res.ReturnData[32:]
	if len(data) < 32 {
		t.Fatal("return data too short")
	}
	length := new(big.Int).SetBytes(data[:32])
	data = data[32:]
	if uint64(len(data)) != (length.Uint64()+31)/32*32 {
		t.Fatal("unexpected data length")
	}
	return data[:length.Uint64()]
}

func returnedDecompressed(t *testing.T, res *evm.TxResult) (common.Address, *big.Int) {
	if len(res.ReturnData) != 64 {
		t.Fatal("unexpected return data length", len(res.ReturnData))
	}
	if !bytes.Equal(res.ReturnData[:12], make([]byte, 12)) {
		t.Fatal("first 12 bytes should be blank")
	}
	var addr common.Address
	copy(addr[:], res.ReturnData[12:32])
	return addr, new(big.Int).SetBytes(res.ReturnData[32:])
}
