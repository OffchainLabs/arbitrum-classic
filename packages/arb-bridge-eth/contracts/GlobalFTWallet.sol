// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

pragma solidity ^0.5.3;

import "./interfaces/IERC20.sol";


contract GlobalFTWallet {

    struct FTWallet {
        address contractAddress;
        uint256 balance;
    }

    struct UserFTWallet {
        mapping(address => uint256) ftIndex;
        FTWallet[] ftList;
    }

    mapping(address => UserFTWallet) ftWallets;

    function ownedERC20s(address _owner) external view returns (address[] memory) {
        UserFTWallet storage wallet = ftWallets[_owner];
        address[] memory addresses = new address[](wallet.ftList.length);
        uint256 addressCount = addresses.length;
        for (uint256 i = 0; i < addressCount; i++) {
            addresses[i] = wallet.ftList[i].contractAddress;
        }
        return addresses;
    }

    function withdrawERC20(address _tokenContract) external {
        uint256 value = getERC20Balance(_tokenContract, msg.sender);
        require(
            removeToken(msg.sender, _tokenContract, value),
            "Wallet doesn't own sufficient balance of token"
        );
        IERC20(_tokenContract).transfer(msg.sender, value);
    }

    function getERC20Balance(address _tokenContract, address _owner) public view returns (uint256) {
        UserFTWallet storage wallet = ftWallets[_owner];
        uint256 index = wallet.ftIndex[_tokenContract];
        if (index == 0) {
            return 0;
        }
        return wallet.ftList[index - 1].balance;
    }

    function depositERC20(address _tokenContract, address _destination, uint256 _value) internal {
        IERC20(_tokenContract).transferFrom(msg.sender, address(this), _value);
        addToken(_destination, _tokenContract, _value);
    }

    function transferERC20(
        address _from,
        address _to,
        address _tokenContract,
        uint256 _value
    )
        internal
        returns (bool)
    {
        if (!removeToken(_from, _tokenContract, _value)) {
            return false;
        }
        addToken(_to, _tokenContract, _value);
        return true;
    }

    function addToken(address _user, address _tokenContract, uint256 _value) private {
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

    function removeToken(address _user, address _tokenContract, uint256 _value) private returns (bool) {
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
            wallet.ftIndex[wallet.ftList[wallet.ftList.length - 1].contractAddress] = walletIndex;
            wallet.ftList[walletIndex - 1] = wallet.ftList[wallet.ftList.length - 1];
            delete wallet.ftIndex[_tokenContract];
            wallet.ftList.pop();
        }
        return true;
    }
}
