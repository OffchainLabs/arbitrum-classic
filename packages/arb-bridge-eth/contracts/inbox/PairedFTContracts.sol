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

import "../interfaces/IPairedErc20.sol";

contract PairedFTContracts {
    mapping(address => bool) public isBuddyContract;
    mapping(address => address) private contractToChain;

    struct PairedFTWallet {
        address contractAddress;
        uint256 balance;
    }

    struct UserPairedFTWallet {
        mapping(address => uint256) ftIndex;
        PairedFTWallet[] ftList;
    }

    mapping(address => UserPairedFTWallet) private pairedFtWallets;

    function depositPairedERC20(
        address _tokenContract,
        address _destination,
        uint256 _value
    ) internal {
        PairedErc20(_tokenContract).burn(msg.sender, _value);
        addPairedToken(_destination, _tokenContract, _value);
    }

    function transferPairedERC20(
        address _from,
        address _to,
        address _tokenContract,
        uint256 _value
    ) internal returns (bool) {
        uint256 balance = getPairedERC20Balance(_tokenContract, _from);

        if (balance < _value) {
            removePairedToken(_from, _tokenContract, balance);
        } else {
            removePairedToken(_from, _tokenContract, _value);
        }

        addPairedToken(_to, _tokenContract, _value);
        return true;
    }

    function withdrawPairedERC20(address _tokenContract) external {
        uint256 value = getPairedERC20Balance(_tokenContract, msg.sender);
        require(
            removePairedToken(msg.sender, _tokenContract, value),
            "Wallet doesn't own sufficient balance of token"
        );

        PairedErc20(_tokenContract).mint(msg.sender, value);
    }

    function getPairedERC20Balance(address _tokenContract, address _owner)
        public
        view
        returns (uint256)
    {
        UserPairedFTWallet storage wallet = pairedFtWallets[_owner];
        uint256 index = wallet.ftIndex[_tokenContract];
        if (index == 0) {
            return 0;
        }
        return wallet.ftList[index - 1].balance;
    }

    function addPairedToken(
        address _user,
        address _tokenContract,
        uint256 _value
    ) private {
        if (_value == 0) {
            return;
        }
        UserPairedFTWallet storage wallet = pairedFtWallets[_user];
        uint256 index = wallet.ftIndex[_tokenContract];
        if (index == 0) {
            index = wallet.ftList.push(PairedFTWallet(_tokenContract, 0));
            wallet.ftIndex[_tokenContract] = index;
        }
        wallet.ftList[index - 1].balance += _value;
    }

    function removePairedToken(
        address _user,
        address _tokenContract,
        uint256 _value
    ) private returns (bool) {
        if (_value == 0) {
            return true;
        }
        UserPairedFTWallet storage wallet = pairedFtWallets[_user];
        uint256 walletIndex = wallet.ftIndex[_tokenContract];
        if (walletIndex == 0) {
            // Wallet has no coins from given ERC20 contract
            return false;
        }
        PairedFTWallet storage tokenWallet = wallet.ftList[walletIndex - 1];
        if (_value > tokenWallet.balance) {
            // Wallet does not own enough ERC20 tokens
            return false;
        }
        tokenWallet.balance -= _value;
        if (tokenWallet.balance == 0) {
            wallet.ftIndex[wallet.ftList[wallet.ftList.length - 1]
                .contractAddress] = walletIndex;
            wallet.ftList[walletIndex - 1] = wallet.ftList[wallet
                .ftList
                .length - 1];
            delete wallet.ftIndex[_tokenContract];
            wallet.ftList.pop();
        }
        return true;
    }
}
