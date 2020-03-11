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
	"errors"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/message"
	"math/big"
)

func depositEth(
	ethClient *goEthdata,
	vmAddress common.Address,
	destination common.Address,
	value *big.Int,
) {
	if balance, ok := ethClient.ethWallet[vmAddress]; ok {
		ethClient.ethWallet[vmAddress] = new(big.Int).Add(balance, value)
	} else {
		ethClient.ethWallet[vmAddress] = value
	}
}

func transferEth(ethClient *goEthdata, to common.Address, from common.Address, value *big.Int) error {
	fromWallet, ok := ethClient.ethWallet[from]
	if !ok {
		return errors.New("unknown from wallet")
	}
	if value.Cmp(fromWallet) > 0 {
		return errors.New("insufficient eth")
	}
	ethClient.ethWallet[from] = new(big.Int).Sub(fromWallet, value)
	toWallet, ok := ethClient.ethWallet[to]
	if ok {
		ethClient.ethWallet[to] = new(big.Int).Add(toWallet, value)
	} else {
		ethClient.ethWallet[to] = value
	}
	return nil
}

// ERC20 wallet
type userFTWallet struct {
	ftList map[common.Address]*big.Int
}

func depositERC20(
	ethClient *goEthdata,
	tokenContract common.Address,
	destination common.Address,
	value *big.Int,
) {
	addToken(ethClient, destination, tokenContract, value)
}

func transferERC20(
	ethClient *goEthdata,
	erc20 message.ERC20,
) bool {
	err := removeToken(ethClient, erc20.From, erc20.TokenAddress, erc20.Value)
	if err {
		return false
	}
	addToken(ethClient, erc20.To, erc20.TokenAddress, erc20.Value)
	return true
}

func addToken(
	ethClient *goEthdata,
	user common.Address,
	token common.Address,
	value *big.Int,
) {
	if value.Cmp(big.NewInt(0)) == 0 {
		return
	}
	userWallet := ethClient.ftWallets[user]
	if _, ok := userWallet.ftList[token]; !ok {
		userWallet.ftList[token] = big.NewInt(0)
	}
	userWallet.ftList[token] = new(big.Int).Add(userWallet.ftList[token], value)
}

func removeToken(
	ethClient *goEthdata,
	user common.Address,
	token common.Address,
	value *big.Int,
) bool {
	if value.Cmp(big.NewInt(0)) == 0 {
		return true
	}
	userWallet := ethClient.ftWallets[user]
	tokenWallet, ok := userWallet.ftList[token]
	if !ok {
		// Wallet has no coins from given ERC20 contract
		return false
	}

	if value.Cmp(tokenWallet) == 1 {
		// Wallet does not own enough ERC20 tokens
		return false
	}
	tokenWallet = new(big.Int).Sub(tokenWallet, value)
	if tokenWallet.Cmp(big.NewInt(0)) == 0 {
		delete(userWallet.ftList, token)
	}
	return true
}

type userNFTWallet struct {
	nftWalletList map[common.Address]map[*big.Int]bool //map of contract address to (map of token ids to owned bool)
}

func depositERC721(
	ethClient *goEthdata,
	tokenContract common.Address,
	destination common.Address,
	value *big.Int,
) {
	addNFTToken(ethClient, destination, tokenContract, value)
}

func transferNFTToken(
	ethClient *goEthdata,
	erc721 message.ERC721,
) bool {
	err := removeToken(ethClient, erc721.From, erc721.TokenAddress, erc721.Id)
	if err {
		return false
	}
	addToken(ethClient, erc721.To, erc721.TokenAddress, erc721.Id)
	return true
}

func addNFTToken(
	ethClient *goEthdata,
	user common.Address,
	erc721 common.Address,
	value *big.Int,
) error {
	userWallet := ethClient.nftWallets[user]
	tokens, ok := userWallet.nftWalletList[erc721]
	if !ok {
		userWallet.nftWalletList[erc721] = make(map[*big.Int]bool)
		tokens = userWallet.nftWalletList[erc721]
	}
	if _, ok := tokens[value]; !ok {
		return errors.New("can't add already owned token")
	}
	tokens[value] = true
	return nil
}

func removeNFTToken(
	ethClient *goEthdata,
	user common.Address,
	erc721 common.Address,
	value *big.Int,
) bool {
	userWallet := ethClient.nftWallets[user]
	tokens, ok := userWallet.nftWalletList[erc721]
	if !ok {
		// Wallet has no coins from given ERC721 contract
		return false
	}

	if tokens[value] {
		// Wallet does not own ERC721 tokens
		return false
	}
	delete(tokens, value)
	return true
}
