package value

import (
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"io"
)

type HashOnlyValue struct {
	hash [32]byte
	size int64
}

func NewHashOnlyValue(hash [32]byte, size int64) HashOnlyValue {
	return HashOnlyValue{hash, size}
}

func NewHashOnlyValueFromValue(val Value) HashOnlyValue {
	return HashOnlyValue{val.Hash(), val.Size()}
}

func NewHashOnlyValueFromReader(rd io.Reader) (HashOnlyValue, error) {
	var size int64
	err := binary.Read(rd, binary.LittleEndian, &size)
	if err != nil {
		return HashOnlyValue{}, err
	}
	var hash [32]byte
	if _, err := io.ReadFull(rd, hash[:]); err != nil {
		return HashOnlyValue{}, err
	}
	return HashOnlyValue{hash, size}, nil
}

func (nv HashOnlyValue) Marshal(wr io.Writer) error {
	//if err := binary.Write(wr, binary.LittleEndian, &nv.size); err != nil {
	//	return err
	//}
	_, err := wr.Write(nv.hash[:])
	return err
}

func (nv HashOnlyValue) TypeCode() byte {
	return TypeCodeHashOnly
}

func (nv HashOnlyValue) InternalTypeCode() byte {
	return TypeCodeHashOnly
}

func (nv HashOnlyValue) Clone() Value {
	return HashOnlyValue{nv.hash, nv.size}
}

func (nv HashOnlyValue) CloneShallow() Value {
	return HashOnlyValue{nv.hash, nv.size}
}

func (nv HashOnlyValue) Size() int64 {
	return nv.size
}

func (nv HashOnlyValue) Equal(val Value) bool {
	return nv.Hash() == val.Hash()
}

func (nv HashOnlyValue) String() string {
	return fmt.Sprintf("HashOnlyValue(%v)", hexutil.Encode(nv.hash[:]))
	return "[HashOnlyValue]"
}

func (nv HashOnlyValue) Hash() [32]byte {
	return nv.hash
}
