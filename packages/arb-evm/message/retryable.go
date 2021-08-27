package message

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/inbox"
	"math/big"
)

type RetryableTx struct {
	Destination       common.Address
	Value             *big.Int
	Deposit           *big.Int
	MaxSubmissionCost *big.Int
	CreditBack        common.Address
	Beneficiary       common.Address
	MaxGas            *big.Int
	GasPriceBid       *big.Int
	Data              []byte
}

func NewRetryableTxFromData(data []byte) RetryableTx {
	destination, data := extractAddress(data)
	value, data := extractUInt256(data)
	deposit, data := extractUInt256(data)
	maxSubmissionCost, data := extractUInt256(data)
	creditBack, data := extractAddress(data)
	beneficiary, data := extractAddress(data)
	maxGas, data := extractUInt256(data)
	gasPriceBid, data := extractUInt256(data)
	dataLength, data := extractUInt256(data)
	return RetryableTx{
		Destination:       destination,
		Value:             value,
		Deposit:           deposit,
		MaxSubmissionCost: maxSubmissionCost,
		CreditBack:        creditBack,
		Beneficiary:       beneficiary,
		MaxGas:            maxGas,
		GasPriceBid:       gasPriceBid,
		Data:              data[:dataLength.Uint64()],
	}
}

func (t RetryableTx) AsData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, AddressData(t.Destination)...)
	ret = append(ret, math.U256Bytes(t.Value)...)
	ret = append(ret, math.U256Bytes(t.Deposit)...)
	ret = append(ret, math.U256Bytes(t.MaxSubmissionCost)...)
	ret = append(ret, AddressData(t.CreditBack)...)
	ret = append(ret, AddressData(t.Beneficiary)...)
	ret = append(ret, math.U256Bytes(t.MaxGas)...)
	ret = append(ret, math.U256Bytes(t.GasPriceBid)...)
	ret = append(ret, math.U256Bytes(big.NewInt(int64(len(t.Data))))...)
	ret = append(ret, t.Data...)
	return ret
}

func (t RetryableTx) Type() inbox.Type {
	return RetryableType
}

func (t RetryableTx) Equals(o RetryableTx) bool {
	return t.Destination == o.Destination &&
		t.Value.Cmp(o.Value) == 0 &&
		t.Deposit.Cmp(o.Deposit) == 0 &&
		t.MaxSubmissionCost.Cmp(o.MaxSubmissionCost) == 0 &&
		t.CreditBack == o.CreditBack &&
		t.Beneficiary == o.Beneficiary &&
		t.MaxGas.Cmp(o.MaxGas) == 0 &&
		t.GasPriceBid.Cmp(o.GasPriceBid) == 0 &&
		bytes.Equal(t.Data, o.Data)
}
