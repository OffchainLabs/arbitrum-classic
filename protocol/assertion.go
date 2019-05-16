package protocol

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/value"
	"io"
)

type Assertion struct {
	AfterHash [32]byte
	NumSteps  uint32
	OutMsgs   []Message
	Logs      []value.Value
}

type MultiReader interface {
	io.Reader
	io.ByteReader
}

func NewAssertion(afterHash [32]byte, numSteps uint32, outMsgs []Message, logs []value.Value) *Assertion {
	return &Assertion{afterHash, numSteps, outMsgs, logs}
}

func NewAssertionFromReader(rd io.Reader) (*Assertion, error) {
	length := uint64(0)
	err := binary.Read(rd, binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 0, length)
	_, err = io.ReadFull(rd, buf)
	if err != nil {
		return nil, err
	}
	assertion := &AssertionBuf{}
	err = proto.Unmarshal(buf, assertion)
	if err != nil {
		return nil, err
	}
	return NewAssertionFromBuf(assertion)
}

func (a *Assertion) Marshal(wr io.Writer) error {
	assertionData, err := proto.Marshal(NewAssertionBuf(a))
	if err != nil {
		return err
	}
	length := uint64(len(assertionData))
	err = binary.Write(wr, binary.LittleEndian, &length)
	if err != nil {
		return err
	}
	_, err = wr.Write(assertionData)
	return err
}

func (a *Assertion) Equals(b *Assertion) bool {
	if a.AfterHash != b.AfterHash || (a.NumSteps != b.NumSteps) || (len(a.OutMsgs) != len(b.OutMsgs)) {
		return false
	}
	for i, ao := range a.OutMsgs {
		if !ao.Equals(b.OutMsgs[i]) {
			return false
		}
	}
	return true
}

func (a *Assertion) LogsHash() [32]byte {
	var logHash [32]byte
	for _, logVal := range a.Logs {
		next := solsha3.SoliditySHA3(solsha3.Bytes32(logHash), solsha3.Bytes32(logVal.Hash()))
		copy(logHash[:], next)
	}
	return logHash
}

func (a *Assertion) Stub() *AssertionStub {
	tracker := NewBalanceTrackerFromMessages(a.OutMsgs)
	var lastHash [32]byte
	for _, msg := range a.OutMsgs {
		next := solsha3.SoliditySHA3(solsha3.Bytes32(lastHash), solsha3.Bytes32(msg.Hash()))
		copy(lastHash[:], next)
	}

	return &AssertionStub{
		a.AfterHash,
		a.NumSteps,
		[32]byte{},
		lastHash,
		[32]byte{},
		a.LogsHash(),
		tracker.TokenAmounts,
	}
}
