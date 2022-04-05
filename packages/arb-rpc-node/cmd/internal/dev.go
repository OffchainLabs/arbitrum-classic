package internal

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func InitializeWallet(mnemonic string, walletCount int) (*hdwallet.Wallet, []accounts.Account, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, nil, err
	}

	accounts := make([]accounts.Account, 0)
	for i := 0; i < walletCount; i++ {
		path := hdwallet.MustParseDerivationPath(fmt.Sprintf("m/44'/60'/0'/0/%v", i))
		account, err := wallet.Derive(path, false)
		if err != nil {
			return nil, nil, err
		}
		accounts = append(accounts, account)
	}
	return wallet, accounts, err
}

func PrintAccountInfo(wallet *hdwallet.Wallet, accounts []accounts.Account) error {
	fmt.Println("Arbitrum Dev Chain")
	fmt.Println("")
	fmt.Println("Available Accounts")
	fmt.Println("==================")
	for i, account := range accounts {
		fmt.Printf("(%v) %v (100 ETH)\n", i, account.Address.Hex())
	}

	fmt.Println("\nPrivate Keys")
	fmt.Println("==================")
	for i, account := range accounts {
		privKey, err := wallet.PrivateKeyHex(account)
		if err != nil {
			return err
		}
		fmt.Printf("(%v) 0x%v\n", i, privKey)
	}
	fmt.Println("")
	return nil
}
