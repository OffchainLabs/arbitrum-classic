package accounttype

import (
	"encoding/json"
	"fmt"
)

type AccountType int

const (
	Compound AccountType = iota
	ExchangeAccount
	ExternalWallet
	FiatAccount
	GasStation
	InternalWallet
	NetworkConnection
	OneTimeAddress
	VaultAccount
)

var accountTypeStrMap = map[string]AccountType{
	"COMPOUND":           Compound,
	"EXCHANGE_ACCOUNT":   ExchangeAccount,
	"EXTERNAL_WALLET":    ExternalWallet,
	"FIAT_ACCOUNT":       FiatAccount,
	"GAS_STATION":        GasStation,
	"INTERNAL_WALLET":    InternalWallet,
	"NETWORK_CONNECTION": NetworkConnection,
	"ONE_TIME_ADDRESS":   OneTimeAddress,
	"VAULT_ACCOUNT":      VaultAccount,
}

var accountTypeMap = map[AccountType]string{}

func init() {
	for str, id := range accountTypeStrMap {
		accountTypeMap[id] = str
	}
}

func New(str string) (*AccountType, error) {
	if id, ok := accountTypeStrMap[str]; ok {
		return &id, nil
	}

	return nil, fmt.Errorf("'%s' is not an AccountType", str)
}

func (at *AccountType) Set(str string) error {
	if id, ok := accountTypeStrMap[str]; ok {
		*at = id
		return nil
	}

	return fmt.Errorf("'%s' is not an AccountType", str)
}

func (at *AccountType) UnmarshalJSON(b []byte) error {
	return at.Set(string(b))
}

func (at *AccountType) String() string {
	return accountTypeMap[*at]
}

func (at AccountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(at.String())
}
