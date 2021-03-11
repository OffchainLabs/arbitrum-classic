package message

import (
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
	Data              []byte
}

func NewRetryableTxFromData(data []byte) RetryableTx {
	destination, data := extractAddress(data)
	value, data := extractUInt256(data)
	deposit, data := extractUInt256(data)
	maxSubmissionCost, data := extractUInt256(data)
	creditBack, data := extractAddress(data)
	beneficiary, data := extractAddress(data)
	dataLength, data := extractUInt256(data)
	return RetryableTx{
		Destination:       destination,
		Value:             value,
		Deposit:           deposit,
		MaxSubmissionCost: maxSubmissionCost,
		CreditBack:        creditBack,
		Beneficiary:       beneficiary,
		Data:              data[:dataLength.Uint64()],
	}
}

func (t RetryableTx) AsData() []byte {
	ret := make([]byte, 0)
	ret = append(ret, addressData(t.Destination)...)
	ret = append(ret, math.U256Bytes(t.Value)...)
	ret = append(ret, math.U256Bytes(t.Deposit)...)
	ret = append(ret, math.U256Bytes(t.MaxSubmissionCost)...)
	ret = append(ret, addressData(t.CreditBack)...)
	ret = append(ret, addressData(t.Beneficiary)...)
	ret = append(ret, math.U256Bytes(big.NewInt(int64(len(t.Data))))...)
	ret = append(ret, t.Data...)
	return ret
}

func (t RetryableTx) Type() inbox.Type {
	return RetryableType
}
