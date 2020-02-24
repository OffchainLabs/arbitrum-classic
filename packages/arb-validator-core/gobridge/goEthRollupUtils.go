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

package gobridge

import (
	"fmt"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"math/big"
)

func calculatePath(from common.Hash, proof []common.Hash) common.Hash {
	node := from
	if len(proof) > 0 {
		for _, val := range proof {
			//node = keccak256(abi.encodePacked(node, proof[i]));
			node = hashing.SoliditySHA3(hashing.Bytes32(node), hashing.Bytes32(val))
		}
	}
	return node
}

//function transactionHash(
//address chain,
//address to,
//address from,
//uint256 seqNumber,
//uint256 value,
//bytes memory data,
//uint256 blockNumber
//)
//internal
//pure
//returns(bytes32)
//{
//return keccak256(
//abi.encodePacked(
//TRANSACTION_MSG,
//chain,
//to,
//from,
//seqNumber,
//value,
//data,
//blockNumber
//)
//);
//}

func transactionHash(
	chain common.Address,
	to common.Address,
	from common.Address,
	seqNumber *big.Int,
	value *big.Int,
	data []byte,
	blockNumber *common.TimeBlocks,
) common.Hash {
	fmt.Println("0", 0)
	fmt.Println("chain", chain)
	fmt.Println("to", to)
	fmt.Println("from", from)
	fmt.Println("seqNumber", seqNumber)
	fmt.Println("value", value)
	fmt.Println("data", data)
	fmt.Println("blockNumber.AsInt()", blockNumber.AsInt())
	return hashing.SoliditySHA3(
		hashing.Uint8(0), // TRANSACTION_MSG
		hashing.Address(chain),
		hashing.Address(to),
		hashing.Address(from),
		hashing.Uint256(seqNumber),
		hashing.Uint256(value),
		data,
		hashing.Uint256(blockNumber.AsInt()),
	)
}

//func transactionMessageHash(
//	chain common.Address,
//	to common.Address,
//	from common.Address,
//	seqNumber *big.Int,
//	value *big.Int,
//	data []byte,
//	blockNumber *common.TimeBlocks,
//) common.Hash {
//	txHash := hashing.SoliditySHA3(
//		hashing.Uint8(0), // TRANSACTION_MSG
//		hashing.Address(chain),
//		hashing.Address(to),
//		hashing.Address(from),
//		hashing.Uint256(seqNumber),
//		hashing.Uint256(value),
//		data,
//	)
//bytes32 dataHash = Value.bytesToBytestackHash(data);
// val1
//Value.Data[] memory msgValues = new Value.Data[](4);
//msgValues[0] = Value.newInt(uint256(to));
//msgValues[1] = Value.newInt(seqNumber);
//msgValues[2] = Value.newInt(value);
//msgValues[3] = Value.newHashOnly(dataHash);

//
//Value.Data[] memory msgType = new Value.Data[](3);
//val2
//msgType[0] = Value.newInt(TRANSACTION_MSG);
//msgType[1] = Value.newInt(uint256(from));
//msgType[2] = Value.newTuple(msgValues);
//
//return Value.hashTuple([
//Value.newInt(blockNumber),
//Value.newInt(uint256(txHash)),
//Value.newTuple(msgType)
//]);
//}
//}
//
//function ethHash(
//address to,
//address from,
//uint256 value,
//uint256 blockNumber,
//uint256 messageNum
//)
//internal
//pure
//returns(bytes32)
//{
//return keccak256(
//abi.encodePacked(
//ETH_DEPOSIT,
//to,
//from,
//value,
//blockNumber,
//messageNum
//)
//);
//}
//
//function ethMessageHash(
//address to,
//address from,
//uint256 value,
//uint256 blockNumber,
//uint256 messageNum
//)
//internal
//pure
//returns(bytes32)
//{
//Value.Data[] memory msgValues = new Value.Data[](2);
//msgValues[0] = Value.newInt(uint256(to));
//msgValues[1] = Value.newInt(value);
//
//Value.Data[] memory msgType = new Value.Data[](3);
//msgType[0] = Value.newInt(ETH_DEPOSIT);
//msgType[1] = Value.newInt(uint256(from));
//msgType[2] = Value.newTuple(msgValues);
//
//Value.Data[] memory ethMsg = new Value.Data[](3);
//ethMsg[0] = Value.newInt(blockNumber);
//ethMsg[1] = Value.newInt(messageNum);
//ethMsg[2] = Value.newTuple(msgType);
//
//return Value.newTuple(ethMsg).hash().hash;
//}
//
//function erc20Hash(
//address to,
//address from,
//address erc20,
//uint256 value,
//uint256 blockNumber,
//uint256 messageNum
//)
//internal
//pure
//returns(bytes32)
//{
//return tokenHash(
//ERC20_DEPOSIT,
//to,
//from,
//erc20,
//value,
//blockNumber,
//messageNum
//);
//}
//
//function erc20MessageHash(
//address to,
//address from,
//address erc20,
//uint256 value,
//uint256 blockNumber,
//uint256 messageNum
//)
//internal
//pure
//returns(bytes32)
//{
//return tokenMessageHash(
//ERC20_DEPOSIT,
//to,
//from,
//erc20,
//value,
//blockNumber,
//messageNum
//);
//}
//
//function erc721Hash(
//address to,
//address from,
//address erc721,
//uint256 id,
//uint256 blockNumber,
//uint256 messageNum
//)
//internal
//pure
//returns(bytes32)
//{
//return tokenHash(
//ERC721_DEPOSIT,
//to,
//from,
//erc721,
//id,
//blockNumber,
//messageNum
//);
//}
//
//function erc721MessageHash(
//address to,
//address from,
//address erc721,
//uint256 id,
//uint256 blockNumber,
//uint256 messageNum
//)
//internal
//pure
//returns(bytes32)
//{
//return tokenMessageHash(
//ERC721_DEPOSIT,
//to,
//from,
//erc721,
//id,
//blockNumber,
//messageNum
//);
//}
//
//function contractTransactionHash(
//address to,
//address from,
//uint256 value,
//bytes memory data,
//uint256 blockNumber,
//uint256 messageNum
//)
//internal
//pure
//returns(bytes32)
//{
//return keccak256(
//abi.encodePacked(
//CONTRACT_TRANSACTION_MSG,
//to,
//from,
//value,
//data,
//blockNumber,
//messageNum
//)
//);
//}
//
//function contractTransactionMessageHash(
//address to,
//address from,
//uint256 value,
//bytes memory data,
//uint256 blockNumber,
//uint256 messageNum
//)
//internal
//pure
//returns(bytes32)
//{
//bytes32 dataHash = Value.bytesToBytestackHash(data);
//Value.Data[] memory msgValues = new Value.Data[](3);
//msgValues[0] = Value.newInt(uint256(to));
//msgValues[2] = Value.newInt(value);
//msgValues[3] = Value.newHashOnly(dataHash);
//
//Value.Data[] memory msgType = new Value.Data[](3);
//msgType[0] = Value.newInt(CONTRACT_TRANSACTION_MSG);
//msgType[1] = Value.newInt(uint256(from));
//msgType[2] = Value.newTuple(msgValues);
//
//return Value.hashTuple([
//Value.newInt(blockNumber),
//Value.newInt(messageNum),
//Value.newTuple(msgType)
//]);
//}
//
//function tokenHash(
//uint8 messageType,
//address to,
//address from,
//address tokenContract,
//uint256 value,
//uint256 blockNumber,
//uint256 messageNum
//)
//private
//pure
//returns(bytes32)
//{
//return keccak256(
//abi.encodePacked(
//messageType,
//to,
//from,
//tokenContract,
//value,
//blockNumber,
//messageNum
//)
//);
//}
//
//function tokenMessageHash(
//uint8 messageType,
//address to,
//address from,
//address tokenContract,
//uint256 value,
//uint256 blockNumber,
//uint256 messageNum
//)
//private
//pure
//returns(bytes32)
//{
//Value.Data[] memory msgValues = new Value.Data[](3);
//msgValues[0] = Value.newInt(uint256(tokenContract));
//msgValues[1] = Value.newInt(uint256(to));
//msgValues[2] = Value.newInt(value);
//
//Value.Data[] memory msgType = new Value.Data[](3);
//msgType[0] = Value.newInt(messageType);
//msgType[1] = Value.newInt(uint256(from));
//msgType[2] = Value.newTuple(msgValues);
//
//Value.Data[] memory ercTokenMsg = new Value.Data[](3);
//ercTokenMsg[0] = Value.newInt(blockNumber);
//ercTokenMsg[1] = Value.newInt(messageNum);
//ercTokenMsg[2] = Value.newTuple(msgType);
//
//return  Value.newTuple(ercTokenMsg).hash().hash;
//}
//}
