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

import "../interfaces/PairedErc20.sol";
import "./FTRecords.sol";

contract PairedFTWallet is FTRecords {
    mapping(address => mapping(address => bool)) private pairedContractRecords;

    function registerContractPair(address _tokenContract, address _chain)
        internal
    {
        pairedContractRecords[_tokenContract][_chain] = true;
    }

    function isPairedContract(address _tokenContract, address _chain)
        public
        view
        returns (bool)
    {
        return pairedContractRecords[_tokenContract][_chain];
    }

    function withdrawPairedERC20(address _tokenContract) external {
        uint256 value = getERC20Balance(_tokenContract, msg.sender);
        require(
            removeToken(msg.sender, _tokenContract, value),
            "Wallet doesn't own sufficient balance of token"
        );

        PairedErc20(_tokenContract).mint(msg.sender, value);
    }

    function depositPairedERC20(
        address _tokenContract,
        address _destination,
        uint256 _value
    ) internal {
        require(
            isPairedContract(_tokenContract, _destination),
            "must be paired contract"
        );
        IERC20(_tokenContract).transferFrom(msg.sender, address(this), _value);
        PairedErc20(_tokenContract).burn(address(this), _value);
        addToken(_destination, _tokenContract, _value);
    }

    function transferPairedERC20(
        address _from,
        address _to,
        address _tokenContract,
        uint256 _value
    ) internal returns (bool) {
        uint256 balance = getERC20Balance(_tokenContract, _from);

        if (balance < _value) {
            removeToken(_from, _tokenContract, balance);
        } else {
            removeToken(_from, _tokenContract, _value);
        }

        addToken(_to, _tokenContract, _value);
        return true;
    }
}
