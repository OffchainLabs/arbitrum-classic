package utils

import (
	"flag"
	"fmt"
	"math"
	"math/big"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"golang.org/x/crypto/ssh/terminal"
)

type WalletFlags struct {
	passphrase *string
	gasPrice   *float64
}

func AddFlags(fs *flag.FlagSet) WalletFlags {
	passphrase := fs.String("password", "", "password=pass")
	gasPrice := fs.Float64("gasprice", 4.5, "gasprice=FloatInGwei")

	return WalletFlags{
		passphrase: passphrase,
		gasPrice:   gasPrice,
	}
}

func GetKeystore(validatorFolder string, args WalletFlags, flags *flag.FlagSet) (*bind.TransactOpts, error) {
	ks := keystore.NewKeyStore(filepath.Join(validatorFolder, "wallets"), keystore.StandardScryptN, keystore.StandardScryptP)

	found := false
	flags.Visit(func(f *flag.Flag) {
		if f.Name == "password" {
			found = true
		}
	})

	var passphrase string
	if !found {
		if len(ks.Accounts()) == 0 {
			fmt.Print("Enter new account password: ")
		} else {
			fmt.Print("Enter account password: ")
		}

		bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil, err
		}
		passphrase = string(bytePassword)

		passphrase = strings.TrimSpace(passphrase)
	} else {
		passphrase = *args.passphrase
	}

	var account accounts.Account
	if len(ks.Accounts()) == 0 {
		var err error
		account, err = ks.NewAccount(passphrase)
		if err != nil {
			return nil, err
		}
	} else {
		account = ks.Accounts()[0]
	}
	err := ks.Unlock(account, passphrase)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyStoreTransactor(ks, account)
	if err != nil {
		return nil, err
	}

	gasPriceAsFloat := 1e9 * (*args.gasPrice)
	if gasPriceAsFloat < math.MaxInt64 {
		auth.GasPrice = big.NewInt(int64(gasPriceAsFloat))
	}
	return auth, nil
}

const WalletArgsString = "[--password=pass] [--gasprice==FloatInGwei]"
