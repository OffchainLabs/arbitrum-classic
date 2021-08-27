package operationtype

import (
	"encoding/json"
	"fmt"
)

type OperationType int

const (
	Burn OperationType = iota
	ContractCall
	Mint
	Raw
	RedeemFromCompound
	SupplyToCompound
	Transfer
)

var operationTypeStrMap = map[string]OperationType{
	"BURN":                 Burn,
	"CONTRACT_CALL":        ContractCall,
	"MINT":                 Mint,
	"RAW":                  Raw,
	"REEDEM_FROM_COMPOUND": RedeemFromCompound,
	"SUPPLY_TO_COMPOUND":   SupplyToCompound,
	"TRANSFER":             Transfer,
}

var operationTypeMap = map[OperationType]string{}

func init() {
	for str, id := range operationTypeStrMap {
		operationTypeMap[id] = str
	}
}

func New(str string) (*OperationType, error) {
	if id, ok := operationTypeStrMap[str]; ok {
		return &id, nil
	}

	return nil, fmt.Errorf("'%s' is not an OperationType", str)
}

func (at *OperationType) Set(str string) error {
	if id, ok := operationTypeStrMap[str]; ok {
		*at = id
		return nil
	}

	return fmt.Errorf("'%s' is not an OperationType", str)
}

func (at *OperationType) UnmarshalJSON(b []byte) error {
	return at.Set(string(b))
}

func (at *OperationType) String() string {
	return operationTypeMap[*at]
}

func (at *OperationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(at.String())
}
