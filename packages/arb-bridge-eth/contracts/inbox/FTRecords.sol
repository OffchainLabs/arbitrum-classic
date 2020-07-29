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

contract FTRecords {
    struct FTWallet {
        address contractAddress;
        uint256 balance;
    }

    struct UserFTWallet {
        mapping(address => uint256) ftIndex;
        FTWallet[] ftList;
    }

    mapping(address => UserFTWallet) internal ftWallets;

    function getERC20Balance(address _tokenContract, address _owner)
        public
        view
        returns (uint256)
    {
        UserFTWallet storage wallet = ftWallets[_owner];
        uint256 index = wallet.ftIndex[_tokenContract];
        if (index == 0) {
            return 0;
        }
        return wallet.ftList[index - 1].balance;
    }

    function addToken(
        address _user,
        address _tokenContract,
        uint256 _value
    ) internal {
        if (_value == 0) {
            return;
        }
        UserFTWallet storage wallet = ftWallets[_user];
        uint256 index = wallet.ftIndex[_tokenContract];
        if (index == 0) {
            index = wallet.ftList.push(FTWallet(_tokenContract, 0));
            wallet.ftIndex[_tokenContract] = index;
        }
        wallet.ftList[index - 1].balance += _value;
    }

    function removeToken(
        address _user,
        address _tokenContract,
        uint256 _value
    ) internal returns (bool) {
        if (_value == 0) {
            return true;
        }
        UserFTWallet storage wallet = ftWallets[_user];
        uint256 walletIndex = wallet.ftIndex[_tokenContract];
        if (walletIndex == 0) {
            // Wallet has no coins from given ERC20 contract
            return false;
        }
        FTWallet storage tokenWallet = wallet.ftList[walletIndex - 1];
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
