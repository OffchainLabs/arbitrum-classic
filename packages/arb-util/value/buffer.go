package value

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Buffer struct {
	data []byte
}

func NewBufferFromReader(rd io.Reader) (*Buffer, error) {
	var length uint64
	if err := binary.Read(rd, binary.BigEndian, &length); err != nil {
		return nil, err
	}
	data := make([]byte, length)
	_, err := io.ReadFull(rd, data)
	if err != nil {
		return nil, err
	}
	return &Buffer{data: data}, nil
}

func (b *Buffer) TypeCode() uint8 {
	return TypeCodeBuffer
}

func (b *Buffer) Equal(other Value) bool {
	o, ok := other.(*Buffer)
	if !ok {
		return false
	}
	return bytes.Equal(b.data, o.data)
}

func (b *Buffer) Size() int64 {
	return 1
}

func (b *Buffer) String() string {
	return fmt.Sprintf("Buffer(0x%x)", b.data)
}

func (b *Buffer) Data() []byte {
	return b.data
}
