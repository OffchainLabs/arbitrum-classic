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

pragma solidity ^0.5.3;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";


contract GlobalWallet {

    using SafeMath for uint256;

    address internal constant ETH_ADDRESS = address(0);

    struct NFTWallet {
        address contractAddress;
        mapping(uint256 => uint256) tokenIndex;
        uint256[] tokenList;
    }

    struct TokenWallet {
        address contractAddress;
        uint256 balance;
    }

    struct Wallet {
        mapping(address => uint256) tokenIndex;
        TokenWallet[] tokenList;

        mapping(address => uint256) nftWalletIndex;
        NFTWallet[] nftWalletList;
    }

    mapping(address => Wallet) wallets;

    function getTokenBalances(address _owner) external view returns (address[] memory, uint256[] memory) {
        Wallet storage wallet = wallets[_owner];
        address[] memory addresses = new address[](wallet.tokenList.length);
        uint256[] memory values = new uint256[](addresses.length);
        uint addressCount = addresses.length;
        for (uint i = 0; i < addressCount; i++) {
            addresses[i] = wallet.tokenList[i].contractAddress;
            values[i] = wallet.tokenList[i].balance;
        }
        return (addresses, values);
    }

    function getNFTTokens(address _owner) external view returns (address[] memory, uint256[] memory) {
        Wallet storage wallet = wallets[_owner];
        uint totalLength = 0;
        uint i;
        for (i = 0; i < wallet.nftWalletList.length; i++) {
            totalLength += wallet.nftWalletList[i].tokenList.length;
        }
        address[] memory addresses = new address[](totalLength);
        uint256[] memory tokens = new uint256[](totalLength);
        uint count = 0;
        uint nftCount = wallet.nftWalletList.length;
        for (i = 0; i < nftCount; i++) {
            NFTWallet storage nftWallet = wallet.nftWalletList[i];
            uint tokenCount = nftWallet.tokenList.length;
            for (uint j = 0; j < tokenCount; j++) {
                addresses[count] = nftWallet.contractAddress;
                tokens[count] = nftWallet.tokenList[j];
                count++;
            }
        }
        return (addresses, tokens);
    }

    function withdrawEth(uint256 _value) external {
        require(
            removeToken(msg.sender, ETH_ADDRESS, _value),
            "Wallet doesn't own sufficient balance of token"
        );
        msg.sender.transfer(_value);
    }

    function withdrawEth(address _destination, uint256 _value) external {
        require(
            removeToken(msg.sender, ETH_ADDRESS, _value),
            "Wallet doesn't own sufficient balance of token"
        );

        _destination.transfer(_value);
    }

    function depositERC20(address _tokenContract, address _destination, uint256 _value) external {
        IERC20(_tokenContract).transferFrom(msg.sender, address(this), _value);
        addToken(_destination, _tokenContract, _value);
    }

    function depositERC20(address _tokenContract, address _sender, address _destination, uint256 _value) external {
        IERC20(_tokenContract).transferFrom(_sender, address(this), _value);
        addToken(_destination, _tokenContract, _value);
    }

    function withdrawERC20(address _tokenContract, uint256 _value) external {
        require(
            removeToken(msg.sender, _tokenContract, _value),
            "Wallet doesn't own sufficient balance of token"
        );
        IERC20(_tokenContract).transfer(msg.sender, _value);
    }

    function withdrawERC20(address _tokenContract, address _destination, uint256 _value) external {
        require(
            removeToken(msg.sender, _tokenContract, _value),
            "Wallet doesn't own sufficient balance of token"
        );
        IERC20(_tokenContract).transfer(_destination, _value);
    }

    function onERC721Received(
        address,
        address _from,
        uint256 _tokenId,
        bytes calldata
    )
        external
        returns(bytes4)
    {
        addNFTToken(_from, msg.sender, _tokenId);
        return bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"));
    }

    function depositERC721(address _tokenContract, address _destination, uint256 _tokenId) external {
        IERC721(_tokenContract).transferFrom(msg.sender, address(this), _tokenId);
        addNFTToken(_destination, _tokenContract, _tokenId);
    }

    function depositERC721(address _tokenContract, address _sender, address _destination, uint256 _tokenId) external {
        IERC721(_tokenContract).transferFrom(_sender, address(this), _tokenId);
        addNFTToken(_destination, _tokenContract, _tokenId);
    }

    function withdrawERC721(address _tokenContract, uint256 _tokenId) external {
        require(
            removeNFTToken(msg.sender, _tokenContract, _tokenId),
            "Wallet doesn't own token"
        );
        IERC721(_tokenContract).safeTransferFrom(address(this), msg.sender, _tokenId);
    }

    function withdrawERC721(address _tokenContract, address _destination, uint256 _tokenId) external {
        require(
            removeNFTToken(msg.sender, _tokenContract, _tokenId),
            "Wallet doesn't own token"
        );
        IERC721(_tokenContract).safeTransferFrom(address(this), _destination, _tokenId);
    }

    function depositEth(address _destination) public payable {
        addToken(_destination, ETH_ADDRESS, msg.value);
    }

    function getTokenBalance(address _tokenContract, address _owner) public view returns (uint256) {
        Wallet storage wallet = wallets[_owner];
        uint index = wallet.tokenIndex[_tokenContract];
        if (index == 0) {
            return 0;
        }
        return wallet.tokenList[index - 1].balance;
    }

    function hasNFT(address _tokenContract, address _owner, uint256 _tokenId) public view returns (bool) {
        Wallet storage wallet = wallets[_owner];
        uint index = wallet.nftWalletIndex[_tokenContract];
        if (index == 0) {
            return false;
        }
        return wallet.nftWalletList[index - 1].tokenIndex[_tokenId] != 0;
    }

    function transferToken(
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

    function transferNFT(
        address _from,
        address _to,
        address _tokenContract,
        uint256 _tokenId
    )
        internal
        returns (bool)
    {
        if(!removeNFTToken(_from, _tokenContract, _tokenId)) {
            return false;
        }
        addNFTToken(_to, _tokenContract, _tokenId);
        return true;
    }

    function addNFTToken(address _user, address _tokenContract, uint256 _tokenId) private {
        Wallet storage wallet = wallets[_user];
        uint index = wallet.nftWalletIndex[_tokenContract];
        if (index == 0) {
            index = wallet.nftWalletList.push(NFTWallet(_tokenContract, new uint256[](0)));
            wallet.nftWalletIndex[_tokenContract] = index;
        }
        NFTWallet storage nftWallet = wallet.nftWalletList[index - 1];
        require(nftWallet.tokenIndex[_tokenId] == 0, "can't add already owned token");
        nftWallet.tokenList.push(_tokenId);
        nftWallet.tokenIndex[_tokenId] = nftWallet.tokenList.length;
    }

    function addToken(address _user, address _tokenContract, uint256 _value) private {
        if (_value == 0) {
            return;
        }
        Wallet storage wallet = wallets[_user];
        uint index = wallet.tokenIndex[_tokenContract];
        if (index == 0) {
            index = wallet.tokenList.push(TokenWallet(_tokenContract, 0));
            wallet.tokenIndex[_tokenContract] = index;
        }
        TokenWallet storage tokenWallet = wallet.tokenList[index - 1];
        tokenWallet.balance = tokenWallet.balance.add(_value);
    }

    function removeNFTToken(address _user, address _tokenContract, uint256 _tokenId) private returns (bool) {
        Wallet storage wallet = wallets[_user];
        uint walletIndex = wallet.nftWalletIndex[_tokenContract];
        if (walletIndex == 0) {
            // Wallet has no coins from given NFT contract
            return false;
        }
        NFTWallet storage nftWallet = wallet.nftWalletList[walletIndex - 1];
        uint tokenIndex = nftWallet.tokenIndex[_tokenId];
        if (tokenIndex == 0) {
            // Wallet does not own specific NFT
            return false;
        }
        nftWallet.tokenIndex[nftWallet.tokenList[nftWallet.tokenList.length - 1]] = tokenIndex;
        nftWallet.tokenList[tokenIndex - 1] = nftWallet.tokenList[nftWallet.tokenList.length - 1];
        delete nftWallet.tokenIndex[_tokenId];
        nftWallet.tokenList.pop();
        if (nftWallet.tokenList.length == 0) {
            wallet.nftWalletIndex[wallet.nftWalletList[wallet.nftWalletList.length - 1].contractAddress] = walletIndex;
            wallet.nftWalletList[walletIndex - 1] = wallet.nftWalletList[wallet.nftWalletList.length - 1];
            delete wallet.nftWalletIndex[_tokenContract];
            wallet.nftWalletList.pop();
        }
        return true;
    }

    function removeToken(address _user, address _tokenContract, uint256 _value) private returns (bool) {
        if (_value == 0) {
            return true;
        }
        Wallet storage wallet = wallets[_user];
        uint walletIndex = wallet.tokenIndex[_tokenContract];
        if (walletIndex == 0) {
            // Wallet has no coins from given ERC20 contract
            return false;
        }
        TokenWallet storage tokenWallet = wallet.tokenList[walletIndex - 1];
        if (_value > tokenWallet.balance) {
            // Wallet does not own enough ERC20 tokens
            return false;
        }
        tokenWallet.balance = tokenWallet.balance.sub(_value);
        if (tokenWallet.balance == 0) {
            wallet.tokenIndex[wallet.tokenList[wallet.tokenList.length - 1].contractAddress] = walletIndex;
            wallet.tokenList[walletIndex - 1] = wallet.tokenList[wallet.tokenList.length - 1];
            delete wallet.tokenIndex[_tokenContract];
            wallet.tokenList.pop();
        }
        return true;
    }
}
