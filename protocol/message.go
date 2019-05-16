package protocol

import (
	"fmt"
	"github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/value"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type Identity [32]byte
type TokenType [21]byte

func (t TokenType) IsToken() bool {
	return t[20] == 0
}

func tokenTypeEncoded(input [21]byte) []byte {
	return common.RightPadBytes(input[:], 21)
}

func tokenTypeArrayEncoded(input [][21]byte) []byte {
	var values []byte
	for _, val := range input {
		values = append(values, tokenTypeEncoded(val)...)
	}
	return values
}

type Message struct {
	Data        value.Value
	TokenType   [21]byte
	Currency    *big.Int
	Destination [32]byte
}

func NewMessage(data value.Value, tokenType [21]byte, currency *big.Int, destination [32]byte) Message {
	return Message{data, tokenType, currency, destination}
}

func NewSimpleMessage(data value.Value, tokenType [21]byte, currency *big.Int, sender common.Address) Message {
	senderArr := [32]byte{}
	copy(senderArr[:], sender.Bytes())
	return Message{data, tokenType, currency, senderArr}
}

func NewMessageFromReader(rd io.Reader) (Message, error) {
	data, err := value.UnmarshalValue(rd)
	if err != nil {
		return Message{}, err
	}

	tokenType := [21]byte{}
	_, err = rd.Read(tokenType[:])
	if err != nil {
		return Message{}, fmt.Errorf("Error unmarshalling OutgoingMessage: %v", err)
	}

	currency, err := value.NewIntValueFromReader(rd)
	if err != nil {
		return Message{}, fmt.Errorf("Error unmarshalling OutgoingMessage: %v", err)
	}

	dest := [32]byte{}
	_, err = rd.Read(tokenType[:])
	if err != nil {
		return Message{}, fmt.Errorf("Error unmarshalling OutgoingMessage: %v", err)
	}

	return NewMessage(data, tokenType, currency.BigInt(), dest), nil
}

func (msg Message) Marshal(w io.Writer) error {
	if err := value.MarshalValue(msg.Data, w); err != nil {
		return err
	}

	_, err := w.Write(msg.TokenType[:])
	if err != nil {
		return err
	}

	err = value.NewIntValue(msg.Currency).Marshal(w)
	if err != nil {
		return err
	}
	_, err = w.Write(msg.Destination[:])
	if err != nil {
		return err
	}
	return nil
}

func (msg Message) Hash() [32]byte {
	var ret [32]byte
	hashVal := solsha3.SoliditySHA3(
		solsha3.Bytes32(msg.Data.Hash),
		tokenTypeEncoded(msg.TokenType),
		solsha3.Uint256(msg.Currency),
		solsha3.Bytes32(msg.Destination),
	)
	copy(ret[:], hashVal)
	return ret
}

func (msg Message) AsValue() value.Value {
	destination := big.NewInt(0)
	destination.SetBytes(msg.Destination[:])
	tokTypeBytes := [32]byte{}
	copy(tokTypeBytes[:], msg.TokenType[:])
	tokTypeInt := big.NewInt(0)
	tokTypeInt.SetBytes(tokTypeBytes[:])
	newTup, _ := value.NewTupleFromSlice([]value.Value{
		msg.Data,
		value.NewIntValue(destination),
		value.NewIntValue(msg.Currency),
		value.NewIntValue(tokTypeInt),
	})
	return newTup
}

func (msg Message) Equals(b Message) bool {
	if msg.TokenType != b.TokenType {
		return false
	}
	if !value.Eq(msg.Data, b.Data) {
		return false
	}
	if msg.Currency.Cmp(b.Currency) != 0 {
		return false
	}
	if msg.Destination != b.Destination {
		return false
	}
	return true
}
