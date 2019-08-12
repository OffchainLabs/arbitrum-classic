/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package checkpoint

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/dgraph-io/badger"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

const PrefixValue byte = iota

func (cp *Checkpointer) writeValue(wr io.Writer, val value.Value) ([]value.Value, error) {
	typecode := val.TypeCode()
	if _, err := wr.Write([]byte{typecode}); err != nil {
		return nil, err
	}
	switch typecode {
	case value.TypeCodeInt:
		theBytes := val.(value.IntValue).ToBytes()
		if _, err := wr.Write(theBytes[:]); err != nil {
			return nil, err
		}
		return nil, nil
	case value.TypeCodeCodePoint:
		codepoint := val.(value.CodePointValue)
		insnNum := codepoint.InsnNum
		if err := binary.Write(wr, binary.LittleEndian, &insnNum); err != nil {
			return nil, err
		}
		op := codepoint.Op
		val, err := writeOp(wr, op)
		if err != nil {
			return nil, err
		}
		if _, err := wr.Write(codepoint.NextHash[:]); err != nil {
			return nil, err
		}
		if val == nil {
			return nil, nil
		}
		return []value.Value{val}, nil
	case value.TypeCodeTuple:
		tup := val.(value.TupleValue)
		size := tup.Len()
		contents := tup.Contents()
		if _, err := wr.Write([]byte{byte(size)}); err != nil {
			return nil, err
		}
		for i := int64(0); i < size; i++ {
			h := contents[i].Hash()
			if _, err := wr.Write(h[:]); err != nil {
				return nil, err
			}
		}
		return contents, nil
	case value.TypeCodeHashOnly:
		h := val.Hash()
		if _, err := wr.Write(h[:]); err != nil {
			return nil, err
		}
		return nil, nil
	default:
		panic("Unrecognized typecode")
	}
}

func (cp *Checkpointer) addRefToValueInTxn(txn *badger.Txn, val value.Value) error {
	h := val.Hash()
	hkey := append([]byte{PrefixValue}, h[:]...)
	item, err := txn.Get(hkey)
	switch err {
	case nil:
		// value found; increment its refcount
		var refCount uint64
		var valCopy []byte
		if err := item.Value(func(barr []byte) error {
			valCopy = append([]byte{}, barr...)
			rd := bytes.NewReader(barr[:8])
			return binary.Read(rd, binary.LittleEndian, &refCount)
		}); err != nil {
			return err
		}

		refCount++
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, &refCount); err != nil {
			return err
		}
		return txn.Set(hkey, append(buf.Bytes(), valCopy[8:]...))
	case badger.ErrKeyNotFound:
		// value not found; create it with refcount=1, and add refs to its children
		var buf bytes.Buffer
		refCount := uint64(1)
		if err := binary.Write(&buf, binary.LittleEndian, &refCount); err != nil {
			return err
		}
		more, err := cp.writeValue(&buf, val)
		if err != nil {
			return err
		}
		if err := txn.Set(hkey, buf.Bytes()); err != nil {
			return err
		}
		for _, v := range more {
			if err := cp.addRefToValueInTxn(txn, v); err != nil {
				return err
			}
		}
		return nil
	default:
		return err
	}
}

func (cp *Checkpointer) AddRefToValue(val value.Value) error {
	return cp.db.Update(func(txn *badger.Txn) error {
		return cp.addRefToValueInTxn(txn, val)
	})
}

func (cp *Checkpointer) RemoveRefToValue(hash [32]byte) {
	go func() {
		_ = cp.synchronousRemoveRefToValue(hash) // accept that error will leave unneeded values laying around
	}()
}

func (cp *Checkpointer) synchronousRemoveRefToValue(hash [32]byte) error {
	var more [][32]byte
	err := cp.db.Update(func(txn *badger.Txn) error {
		key := append([]byte{PrefixValue}, hash[:]...)
		item, err := txn.Get(key)
		if err != nil {
			return err
		}

		var refCount uint64
		var valCopy []byte
		if err := item.Value(func(val []byte) error {
			valCopy = append([]byte{}, val...)
			rd := bytes.NewReader(val[:8])
			return binary.Read(rd, binary.LittleEndian, &refCount)
		}); err != nil {
			return err
		}

		refCount--
		if refCount == 0 {
			if err := txn.Delete(key); err != nil {
				return err
			}
			if valCopy[8] == value.TypeCodeTuple {
				size := int(valCopy[9])
				rd := bytes.NewReader(valCopy[10:])
				more = make([][32]byte, size)
				for i := 0; i < size; i++ {
					h := [32]byte{}
					if _, err := io.ReadFull(rd, h[:]); err != nil {
						return err
					}
					more[i] = h
				}
			}
			return nil
		}
		var buf bytes.Buffer
		if err := binary.Write(&buf, binary.LittleEndian, &refCount); err != nil {
			return err
		}
		return txn.Set(key, append(buf.Bytes(), valCopy[8:]...))
	})
	if err != nil {
		return err
	}
	for _, h := range more {
		cp.RemoveRefToValue(h)
	}
	return nil
}

func (cp *Checkpointer) RestoreValueFromHash(hash [32]byte) (value.Value, error) {
	txn := cp.db.NewTransaction(false) // open a read-only transaction
	defer txn.Discard()

	return cp.restoreValueFromHashInTxn(txn, hash)
}

func (cp *Checkpointer) restoreValueFromHashInTxn(txn *badger.Txn, hash [32]byte) (value.Value, error) {
	hkey := append([]byte{PrefixValue}, hash[:]...)
	item, err := txn.Get(hkey)
	if err != nil {
		return nil, err
	}
	var bytesRead []byte
	if err := item.Value(func(bytesVal []byte) error {
		bytesRead = append([]byte{}, bytesVal...)
		return nil
	}); err != nil {
		return nil, err
	}

	rd := bytes.NewReader(bytesRead)
	var unusedRefCount uint64
	if err := binary.Read(rd, binary.LittleEndian, &unusedRefCount); err != nil {
		return nil, err
	}
	var typeCode byte
	if err := binary.Read(rd, binary.LittleEndian, &typeCode); err != nil {
		return nil, err
	}
	switch typeCode {
	case value.TypeCodeInt:
		return value.NewIntValueFromReader(rd)
	case value.TypeCodeCodePoint:
		var insnNum int64
		if err2 := binary.Read(rd, binary.LittleEndian, &insnNum); err2 != nil {
			return nil, err2
		}
		op, err2 := cp.restoreOp(txn, rd)
		if err2 != nil {
			return nil, err2
		}
		var nextHash [32]byte
		if _, err2 := io.ReadFull(rd, nextHash[:]); err2 != nil {
			return nil, err2
		}
		return value.CodePointValue{InsnNum: insnNum, Op: op, NextHash: nextHash}, nil
	case value.TypeCodeTuple:
		var sizeAsByte byte
		if err2 := binary.Read(rd, binary.LittleEndian, &sizeAsByte); err2 != nil {
			return nil, err2
		}
		size := int(sizeAsByte)
		contents := make([]value.Value, size)
		for i := 0; i < size; i++ {
			var subHash [32]byte
			if _, err2 := io.ReadFull(rd, subHash[:]); err2 != nil {
				return nil, err
			}
			contents[i], err = cp.restoreValueFromHashInTxn(txn, subHash)
			if err != nil {
				return nil, err
			}
		}
		return value.NewTupleFromSlice(contents)
	case value.TypeCodeHashOnly:
		var h [32]byte
		if _, err := io.ReadFull(rd, h[:]); err != nil {
			return nil, err
		}
		return value.NewHashOnlyValue(h, 32), nil
	default:
		return nil, Error{"Unexpected typecode in checkpoint"}
	}
}
