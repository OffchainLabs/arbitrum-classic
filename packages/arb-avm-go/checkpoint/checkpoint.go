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

// How to use this:
//   When you start a new VM, call
//		cp, err := checkpoint.NewCheckpointer(machine, true)
//   To checkpoint a VM, call
//	    err := cp.SaveMachine("your checkpoint name", machine)
//   If you restart and want to restore a checkpointed VM, call
//      cp, err := checkpoint.NewCheckpointer(nil, false)
//      machine, err := cp.RestoreMachine("your checkpoint name")
//   Before exiting, it's polite to clean up by calling
//      err := cp.Close()
//   [error handling code omitted]

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/offchainlabs/arbitrum/packages/arb-avm-go/vm"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

type Checkpointer struct {
	db          *badger.DB
	closeSignal chan struct{}
}

const (
	defaultCheckpointPath = "/tmp/arb-validator-checkpoint"
	vcpVersionNumsKey     = "VersionedCheckpointer:versionNums"
)

type Error struct {
	str string
}

func (e Error) Error() string {
	return e.str
}

func NewCheckpointer(machine *vm.Machine, destroyOldCheckpoints bool) (*Checkpointer, error) {
	if destroyOldCheckpoints {
		if err := os.RemoveAll(defaultCheckpointPath); err != nil {
			return nil, err
		}
	}

	opts := badger.DefaultOptions(defaultCheckpointPath)
	opts.ValueDir = defaultCheckpointPath
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	ret := &Checkpointer{db, make(chan struct{})}
	if machine != nil {
		// TODO: save the code asynchronously; have machine checkpoints wait for completion
		//  open question: how to handle errors in saving the code; probably best to just retry
		if err := ret.SaveCode(machine); err != nil {
			return nil, err
		}
	}

	// start Badger garbage collector
	go func() {
		ticker := time.NewTicker(5 * time.Minute)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
			again:
				err := ret.db.RunValueLogGC(0.7)
				if err == nil {
					goto again
				}
			case <-ret.closeSignal:
				return
			}
		}
	}()

	return ret, nil
}

func (cp *Checkpointer) Close() error {
	if err := cp.db.Close(); err != nil {
		return err
	}
	close(cp.closeSignal)
	return nil
}

type VersionedCheckpointer struct {
	cp         *Checkpointer
	minVersion int64
	maxVersion int64
}

func NewVersionedCheckpointer(cp *Checkpointer) (*VersionedCheckpointer, error) {
	minVersion := int64(0)
	maxVersion := int64(-1)
	restoring := false

	if err := cp.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(vcpVersionNumsKey))
		if err != nil {
			if err == badger.ErrKeyNotFound {
				return nil
			}
			return err
		}
		if err := item.Value(func(val []byte) error {
			rd := bytes.NewReader(val)
			if err := binary.Read(rd, binary.LittleEndian, &minVersion); err != nil {
				return err
			}
			if err := binary.Read(rd, binary.LittleEndian, &maxVersion); err != nil {
				return err
			}
			return nil
		}); err != nil {
			return err
		}

		restoring = true
		return nil
	}); err != nil {
		return nil, err
	}
	ret := &VersionedCheckpointer{cp, minVersion, maxVersion}
	if !restoring {
		if err := ret.saveState(); err != nil {
			return nil, err
		}
	}
	return ret, nil
}

func (vcp *VersionedCheckpointer) Close() error {
	return vcp.cp.Close()
}

func (vcp *VersionedCheckpointer) saveStateInTxn(txn *badger.Txn) error {
	var buf bytes.Buffer
	if err := binary.Write(&buf, binary.LittleEndian, &vcp.minVersion); err != nil {
		return err
	}
	if err := binary.Write(&buf, binary.LittleEndian, &vcp.maxVersion); err != nil {
		return err
	}
	return txn.Set([]byte(vcpVersionNumsKey), buf.Bytes())
}

func (vcp *VersionedCheckpointer) saveState() error {
	return vcp.cp.db.Update(func(txn *badger.Txn) error {
		return vcp.saveStateInTxn(txn)
	})
}

func vcpStateDataKey(versionNum int64) []byte {
	return []byte("VersionedCheckpointer:stateData:" + string(versionNum))
}

func vcpMachineVersionKey(versionNum int64) string {
	return "versioned:" + string(versionNum)
}

func (vcp *VersionedCheckpointer) SaveVersion(machine *vm.Machine, stateData []byte) (versionNum int64, returnErr error) {
	returnErr = vcp.cp.db.Update(func(txn *badger.Txn) error {
		versionNum = 1 + vcp.maxVersion
		nameSuffix := vcpMachineVersionKey(versionNum)
		if err := vcp.cp.saveMachineInTxn(txn, []byte(nameSuffix), machine); err != nil {
			return err
		}
		if stateData != nil {
			if err := txn.Set(vcpStateDataKey(versionNum), stateData); err != nil {
				return err
			}
		}
		vcp.maxVersion = versionNum
		if err := vcp.saveStateInTxn(txn); err != nil {
			vcp.maxVersion-- // revert maxVersion, because transaction will abort
			return err
		}
		return nil
	})
	return
}

func (vcp *VersionedCheckpointer) RestoreVersion(versionNum int64) (machine *vm.Machine, stateData []byte, retError error) {
	machine = nil
	stateData = nil
	retError = vcp.cp.db.View(func(txn *badger.Txn) error {
		if !vcp.IsKnownVersion(versionNum) {
			return Error{"Can't restore; invalid version number"}
		}

		item, err := txn.Get(vcpStateDataKey(versionNum))
		if err == nil {
			if err := item.Value(func(val []byte) error {
				stateData = append([]byte{}, val...)
				return nil
			}); err != nil {
				return err
			}
		} else {
			if err != badger.ErrKeyNotFound {
				return err
			}
		}

		machine, err = vcp.cp.restoreMachineInTxn(txn, []byte(vcpMachineVersionKey(versionNum)))
		if err != nil {
			return err
		}
		return nil
	})
	return
}

func (vcp *VersionedCheckpointer) KnownVersions() (minVersionNum, maxVersionNum int64) {
	minVersionNum = vcp.minVersion
	maxVersionNum = vcp.maxVersion
	return
}

func (vcp *VersionedCheckpointer) IsKnownVersion(num int64) bool {
	return (num >= vcp.minVersion) && (num <= vcp.maxVersion)
}

func (vcp *VersionedCheckpointer) discardVersion(num int64) error {
	var refs [][32]byte
	if err := vcp.cp.db.Update(func(txn *badger.Txn) error {
		if err := txn.Delete(vcpStateDataKey(num)); err != nil {
			return err
		}
		mkey := []byte("machine:" + vcpMachineVersionKey(num))
		item, err := txn.Get(mkey)
		if err != nil {
			return nil
		}
		if err := item.Value(func(barr []byte) error {
			rd := bytes.NewReader(barr)
			refs = make([][32]byte, 5)
			for i := 0; i < 5; i++ {
				if _, err := rd.Read(refs[i][:]); err != nil {
					return err
				}
			}
			return nil
		}); err != nil {
			return err
		}

		if err := txn.Delete(mkey); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	for _, h := range refs {
		if err := vcp.cp.synchronousRemoveRefToValue(h); err != nil {
			return err
		}
	}
	return nil
}

func (vcp *VersionedCheckpointer) DiscardVersions(newMinVersionNum int64) error {
	if newMinVersionNum > 1+vcp.maxVersion {
		return Error{"Can't discard versions that don't yet exist"}
	}
	oldMinVersion := vcp.minVersion
	vcp.minVersion = newMinVersionNum

	// asynchronously discard old versions
	go func() {
		for i := oldMinVersion; i < newMinVersionNum; i++ {
			_ = vcp.discardVersion(i) // ignore errors; error will leave a useless old version sitting around
		}
	}()
	return nil
}

type EventChainCheckpointer struct {
	cp          *Checkpointer
	fullKey     []byte
	machineHash [32]byte
	timeBounds  [2]uint64
	balances    *protocol.BalanceTracker
	nextSeqNo   uint64
	discarded   bool
}

const _eventChainCheckpointerPrefix = "EventChain:"

func NewEventChainCheckpointer(
	cp *Checkpointer,
	keySuffix []byte,
	machine *vm.Machine,
	timeBounds [2]uint64,
	balances *protocol.BalanceTracker) (*EventChainCheckpointer, error) {
	var buf bytes.Buffer
	fullKey := append([]byte(_eventChainCheckpointerPrefix), keySuffix...)

	machineHash := machine.Hash()
	if _, err := buf.Write(machineHash[:]); err != nil {
		return nil, err
	}
	for i := 0; i < 2; i++ {
		if err := binary.Write(&buf, binary.LittleEndian, &timeBounds[i]); err != nil {
			return nil, err
		}
	}
	if err := balances.Marshal(&buf); err != nil {
		return nil, err
	}

	var seqNumBuf bytes.Buffer
	seqNum := uint64(0)
	if err := binary.Write(&seqNumBuf, binary.LittleEndian, &seqNum); err != nil {
		return nil, err
	}

	if err := cp.db.Update(func(txn *badger.Txn) error {
		machineKey := append(fullKey, []byte(":machine:")...)
		seqNumKey := append(fullKey, []byte(":nextseqnum:")...)
		if err := cp.saveMachineInTxn(txn, machineKey, machine); err != nil {
			return err
		}
		if err := txn.Set(seqNumKey, seqNumBuf.Bytes()); err != nil {
			return err
		}
		if err := txn.Set(fullKey, buf.Bytes()); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &EventChainCheckpointer{
		cp,
		fullKey,
		machineHash,
		[2]uint64{timeBounds[0], timeBounds[1]},
		balances.Clone(),
		uint64(0),
		false,
	}, nil
}

func (ecc *EventChainCheckpointer) Discard() error {
	if !ecc.discarded {
		ecc.discarded = true
		var inboxHash [32]byte
		needToRemoveInboxRef := false
		if err := ecc.cp.db.Update(func(txn *badger.Txn) error {
			item, err := txn.Get(ecc.fullKey)
			if err != nil {
				return err
			}
			if err := item.Value(func(val []byte) error {
				if len(val) < 64 {
					return errors.New("EventChainCheckpointer::Discard: checkpointed item is too small")
				}
				copy(inboxHash[:], val[32:64])
				return nil
			}); err != nil {
				return err
			}
			needToRemoveInboxRef = true
			return txn.Delete(ecc.fullKey)
		}); err != nil {
			return nil
		}
		if needToRemoveInboxRef {
			ecc.cp.RemoveRefToValue(inboxHash) // error will be ignored
		}

		// asynchronously delete the info checkpointed for this event chain
		go func() {
			// ignore all errors in here--no way to recover, and worst possible outcome is that orphaned data is left in database
			for i := uint64(0); i < ecc.nextSeqNo; i++ {
				needRemove := false
				var inboxHash [32]byte
				_ = ecc.cp.db.Update(func(txn *badger.Txn) error {
					keyIntent := ecc.eccKeyForSeqNum(i, "intentToSign")
					keySigs := ecc.eccKeyForSeqNum(i, "recordSignatures")
					item, err := txn.Get(keyIntent)
					if err == nil {
						_ = item.Value(func(itemBytes []byte) error {
							copy(inboxHash[:], itemBytes[32:64])
							return nil
						})
						needRemove = true

					}
					_ = txn.Delete(keyIntent)
					_ = txn.Delete(keySigs)
					return nil
				})
				if needRemove {
					ecc.cp.RemoveRefToValue(inboxHash)
				}
			}
		}()
	}
	return nil
}

func (ecc *EventChainCheckpointer) eccKeyForSeqNum(seqNum uint64, kind string) []byte {
	return append(ecc.fullKey, []byte(string(seqNum)+kind)...)
}

func (ecc *EventChainCheckpointer) RecordIntentToSign(seqNum uint64, machine *vm.Machine, inbox value.Value) error {
	if ecc.discarded {
		return errors.New("can't record intent-to-sign on discarded EventChainCheckpointer")
	}
	if seqNum < ecc.nextSeqNo {
		return errors.New("can't reuse sequence number in EventChainCheckpointer")
	}
	machineHash := machine.Hash()
	inboxHash := inbox.Hash()

	var buf bytes.Buffer
	if _, err := buf.Write(machineHash[:]); err != nil {
		return err
	}
	if _, err := buf.Write(inboxHash[:]); err != nil {
		return err
	}

	ecc.nextSeqNo = seqNum + 1
	var seqNumBuf bytes.Buffer
	if err := binary.Write(&seqNumBuf, binary.LittleEndian, &seqNum); err != nil {
		return err
	}
	seqNumKey := append(ecc.fullKey, []byte(":nextseqnum:")...)
	key := ecc.eccKeyForSeqNum(seqNum, "intentToSign")
	return ecc.cp.db.Update(func(txn *badger.Txn) error {
		if err := ecc.cp.addRefToValueInTxn(txn, inbox); err != nil {
			return err
		}
		if err := ecc.cp.saveMachineInTxn(txn, key, machine); err != nil {
			return err
		}
		if err := txn.Set(seqNumKey, seqNumBuf.Bytes()); err != nil {
			return err
		}
		return txn.Set(key, buf.Bytes())
	})
}

func (ecc *EventChainCheckpointer) RecordSignatures(seqNum uint64, marshaledSigs []byte) error {
	if ecc.discarded {
		return errors.New("can't record intent-to-sign on discarded EventChainCheckpointer")
	}
	if seqNum >= ecc.nextSeqNo {
		return errors.New("EventChainCheckpointer::RecordSignatures: invalid sequence number")
	}
	key := ecc.eccKeyForSeqNum(seqNum, "recordSignatures")
	return ecc.cp.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, marshaledSigs)
	})
}

func RestoreEventChainCheckpointer(cp *Checkpointer, keySuffix []byte) (*EventChainCheckpointer, error) {
	fullKey := append([]byte(_eventChainCheckpointerPrefix), keySuffix...)
	var recordedBytes []byte
	if err := cp.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(fullKey)
		if err != nil {
			return err
		}
		if err := item.Value(func(byteArr []byte) error {
			recordedBytes = append([]byte{}, byteArr...)
			return nil
		}); err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	rd := bytes.NewReader(recordedBytes)

	var machineHash [32]byte
	if _, err := rd.Read(machineHash[:]); err != nil {
		return nil, err
	}

	var timeBounds [2]uint64
	for i := 0; i < 2; i++ {
		if err := binary.Read(rd, binary.LittleEndian, &timeBounds[i]); err != nil {
			return nil, err
		}
	}

	balanceTracker, err := protocol.NewBalanceTrackerFromReader(rd)
	if err != nil {
		return nil, err
	}

	seqNumKey := append(fullKey, []byte(":nextseqnum:")...)
	nextSeqNum := uint64(0)
	if err := cp.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(seqNumKey)
		if err != nil {
			return err
		}
		return item.Value(func(byteArr []byte) error {
			rd := bytes.NewReader(byteArr)
			return binary.Read(rd, binary.LittleEndian, &nextSeqNum)
		})
	}); err != nil {
		return nil, err
	}
	return &EventChainCheckpointer{
		cp,
		fullKey,
		machineHash,
		timeBounds,
		balanceTracker,
		nextSeqNum,
		false,
	}, nil
}

func (ecc *EventChainCheckpointer) RestoreChainStartMachine() (*vm.Machine, error) {
	return ecc.cp.RestoreMachine(ecc.fullKey)
}

func (ecc *EventChainCheckpointer) RestoreFromSeqNum(seqNum uint64) (*vm.Machine /*inbox*/, value.Value /*marshaledSigs*/, []byte, error) {
	if ecc.discarded {
		return nil, nil, nil, errors.New("can't restore from discarded EventChainCheckpointer")
	}
	if seqNum >= ecc.nextSeqNo {
		return nil, nil, nil, errors.New("invalid sequence number in EventChainCheckpointer::RestoreFromSeqNum")
	}

	intentKey := ecc.eccKeyForSeqNum(seqNum, "intentToSign")
	var machineHash [32]byte
	var inboxHash [32]byte
	if err := ecc.cp.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(intentKey)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			if len(val) < 32 {
				return errors.New("EventChainCheckpointer: intentToSign record is too small")
			}
			copy(machineHash[:], val[:32])
			copy(inboxHash[:], val[32:])
			return nil
		})
	}); err != nil {
		return nil, nil, nil, err
	}
	machine, err := ecc.cp.RestoreMachine(intentKey)
	if err != nil {
		return nil, nil, nil, err
	}
	inbox, err := ecc.cp.RestoreValueFromHash(inboxHash)
	if err != nil {
		return nil, nil, nil, err
	}

	var marshaledSigs []byte
	sigsKey := ecc.eccKeyForSeqNum(seqNum, "recordSignatures")
	if err := ecc.cp.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(sigsKey)
		if err == badger.ErrKeyNotFound {
			return nil
		} else if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			marshaledSigs = append([]byte{}, val...)
			return nil
		})
	}); err != nil {
		return nil, nil, nil, err
	}
	return machine, inbox, marshaledSigs, nil
}

func (cp *Checkpointer) EntryExists(key []byte) (bool, error) {
	txn := cp.db.NewTransaction(false)
	defer txn.Discard()

	_, err := txn.Get(key)
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (cp *Checkpointer) SaveMachine(keySuffix []byte, machine *vm.Machine) error {
	return cp.db.Update(func(txn *badger.Txn) error {
		return cp.saveMachineInTxn(txn, keySuffix, machine)
	})
}

func (cp *Checkpointer) saveMachineInTxn(txn *badger.Txn, keySuffix []byte, machine *vm.Machine) error {
	key := append([]byte("machine:"), keySuffix...)
	var buf bytes.Buffer
	var vals [6]value.Value
	vals[0] = machine.Stack().FullyExpandedValue()
	vals[1] = machine.AuxStack().FullyExpandedValue()
	vals[2] = machine.Register().Get()
	vals[3] = machine.Static().Get()
	vals[4] = machine.GetPC()
	vals[5] = machine.GetErrHandler()
	for i := 0; i < len(vals); i++ {
		if err := cp.addRefToValueInTxn(txn, vals[i]); err != nil {
			return err
		}
		h := vals[i].Hash()
		if _, err := buf.Write(h[:]); err != nil {
			return err
		}
	}
	sizeLimit := machine.GetSizeLimit()
	if err := binary.Write(&buf, binary.LittleEndian, &sizeLimit); err != nil {
		return err
	}
	return txn.Set(key, buf.Bytes())
}

func (cp *Checkpointer) RestoreMachine(keySuffix []byte) (*vm.Machine, error) {
	txn := cp.db.NewTransaction(false)
	defer txn.Discard()

	return cp.restoreMachineInTxn(txn, keySuffix)
}

func (cp *Checkpointer) restoreMachineInTxn(txn *badger.Txn, keySuffix []byte) (*vm.Machine, error) {
	key := append([]byte("machine:"), keySuffix...)

	item, err := txn.Get(key)
	if err != nil {
		return nil, err
	}
	var machineBytes []byte
	if err := item.Value(func(bytesVal []byte) error {
		machineBytes = append([]byte{}, bytesVal...)
		return nil
	}); err != nil {
		return nil, err
	}

	rd := bytes.NewReader(machineBytes)
	var vals [6]value.Value
	var h [32]byte
	for i := 0; i < len(vals); i++ {
		if _, err := io.ReadFull(rd, h[:]); err != nil {
			return nil, err
		}
		vals[i], err = cp.restoreValueFromHashInTxn(txn, h)
		if err != nil {
			return nil, err
		}
	}
	var sizeLimit int64
	if err := binary.Read(rd, binary.LittleEndian, &sizeLimit); err != nil {
		return nil, err
	}

	// codeOps, err := cp.restoreCodeInTxn(txn)
	// if err != nil {
	//	return nil, err
	//}
	//
	// errHandler, ok := vals[5].(value.CodePointValue)
	// if !ok {
	//	return nil, errors.New("6th value must be a codepoint")
	//}

	return nil, errors.New("ERROR: Unimplemented")
	// return vm.RestoreMachine(codeOps, vals[0], vals[1], vals[2], vals[3], vals[4], errHandler, sizeLimit), nil
}

func writeOp(wr io.Writer, op value.Operation) (value.Value, error) {
	var val value.Value
	if op.TypeCode() == 1 {
		iop := op.(value.ImmediateOperation)
		val = iop.Val
		h := val.Hash()
		contents := append([]byte{1, byte(iop.Op)}, h[:]...)
		if _, err := wr.Write(contents); err != nil {
			return nil, err
		}
	} else {
		bop := op.(value.BasicOperation)
		if _, err := wr.Write([]byte{0, byte(bop.Op)}); err != nil {
			return nil, err
		}
	}
	return val, nil
}

func (cp *Checkpointer) restoreOp(txn *badger.Txn, rd io.Reader) (value.Operation, error) {
	var buf [2]byte
	if _, err := io.ReadFull(rd, buf[:]); err != nil {
		return nil, err
	}
	if buf[0] == 1 {
		var h [32]byte
		if _, err := io.ReadFull(rd, h[:]); err != nil {
			return nil, err
		}
		immedVal, err := cp.restoreValueFromHashInTxn(txn, h)
		if err != nil {
			return nil, err
		}
		return value.ImmediateOperation{Op: value.Opcode(buf[1]), Val: immedVal}, nil
	}
	return value.BasicOperation{Op: value.Opcode(buf[1])}, nil
}

func (cp *Checkpointer) SaveCode(machine *vm.Machine) error {
	var buf bytes.Buffer
	ops := machine.GetAllOperations()
	key := []byte("code")

	numOps := uint64(len(ops))
	if err := binary.Write(&buf, binary.LittleEndian, &numOps); err != nil {
		return err
	}

	for i := uint64(0); i < numOps; i++ {
		op := ops[i]
		val, err := writeOp(&buf, op)
		if err != nil {
			return err
		}
		if val != nil {
			_ = cp.AddRefToValue(val) // accept that will be orphaned if txn below fails
		}
	}
	return cp.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, buf.Bytes())
	})
}

func (cp *Checkpointer) restoreCodeInTxn(txn *badger.Txn) ([]value.Operation, error) {
	key := []byte("code")

	item, err := txn.Get(key)
	if err != nil {
		return nil, err
	}
	var codeBytes []byte
	if err := item.Value(func(bytesVal []byte) error {
		codeBytes = append([]byte{}, bytesVal...)
		return nil
	}); err != nil {
		return nil, err
	}

	rd := bytes.NewReader(codeBytes)

	var numOps uint64
	if err := binary.Read(rd, binary.LittleEndian, &numOps); err != nil {
		return nil, err
	}
	ops := make([]value.Operation, numOps)
	for i := uint64(0); i < numOps; i++ {
		ops[i], err = cp.restoreOp(txn, rd)
		if err != nil {
			return nil, err
		}
	}
	return ops, nil
}
