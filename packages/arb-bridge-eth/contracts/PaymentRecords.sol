// SPDX-License-Identifier: Apache-2.0

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

pragma solidity ^0.5.11;

contract PaymentRecords {
	mapping(bytes32 => mapping(uint256 => mapping(address => address))) internal paymentMap;

    event PaymentTransfer(
        bytes32 nodeHash,
        uint256 messageIndex,
        address originalOwner,
        address prevOwner,
        address newOwner);


    function transferPayment(
        address originalOwner,
        address newOwner,
        bytes32 nodeHash,
        uint256 messageIndex) external
    {
        address currentOwner = getPaymentOwner(originalOwner, nodeHash, messageIndex);
        require(msg.sender == currentOwner, "Must be payment owner.");

        paymentMap[nodeHash][messageIndex][originalOwner] = newOwner;

        emit PaymentTransfer(nodeHash, messageIndex, originalOwner, currentOwner, newOwner);
    }

    function getPaymentOwner(
        address originalOwner,
        bytes32 nodeHash,
        uint256 messageIndex) public view returns(address)
    {
    	address currentOwner = paymentMap[nodeHash][messageIndex][originalOwner];

    	if(currentOwner == address(0)){
    		return originalOwner;
		}else{
			return currentOwner;
		}
    }
}
