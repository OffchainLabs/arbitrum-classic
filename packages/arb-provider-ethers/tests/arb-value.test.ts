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
/* eslint-env node, jest */
'use strict';

import * as ethers from 'ethers';

const utils = ethers.utils;
import * as arb from '../src/lib/value';
import testCases from './test_cases.json';

// Helper shortcuts
const bn = utils.bigNumberify;
const ZEROS_16B = '00000000000000000000000000000000';
const ZEROS_32B = ZEROS_16B + ZEROS_16B;
const ONES_16B = 'ffffffffffffffffffffffffffffffff';
const ONES_32B = ONES_16B + ONES_16B;
const EMPTY_TUPLE_HASH = '0x69c322e3248a5dfc29d73c5b0553b0185a35cd5bb6386747517ef7e53b15e287';

describe('Constructors', function() {
    const nullHash = '0x' + ZEROS_32B;

    test('BasicOp', function() {
        const bop = new arb.BasicOp(arb.OpCode.Halt);
        expect(bop.opcode).toBe(arb.OpCode.Halt);
    });

    test('ImmOp', function() {
        const iop = new arb.ImmOp(0x19, new arb.IntValue(utils.bigNumberify(9)));
        expect(iop.opcode).toBe(0x19);
        expect((iop.value as arb.IntValue).bignum.toNumber()).toBe(9);
    });

    test('IntValue', function() {
        const iv = new arb.IntValue(utils.bigNumberify(0));
        expect(iv.bignum.toNumber()).toBe(0);
        expect(iv.typeCode()).toBe(0);
    });

    test('CodePointValue', function() {
        const cpv = new arb.CodePointValue(0, new arb.BasicOp(0x60), nullHash);
        expect(cpv.insnNum).toBe(0);
        expect(cpv.op.opcode).toBe(0x60);
        expect(cpv.nextHash).toBe(nullHash);

        // Test BasicOp hash value
        const bopv = new arb.CodePointValue(99, new arb.BasicOp(0x60), EMPTY_TUPLE_HASH);
        const preCalc = '0xb9cffde57db229ede25012536613c9a9f9a7cde0e5f23381350737b6818852da';
        expect(bopv.hash()).toBe(preCalc);

        // Test ImmOp hash value
        const immop = new arb.ImmOp(0x60, new arb.IntValue(bn(0)));
        const immv = new arb.CodePointValue(100, immop, EMPTY_TUPLE_HASH);
        const preCalc2 = '0x9b6304f1c0d7299152b70c5097dcb370ed46668b1f09146586a333f507594619';
        expect(immv.hash()).toBe(preCalc2);
    });

    test('HashOnlyValue', function() {
        const hov = new arb.HashOnlyValue(nullHash, 0);
        expect(hov.hash()).toBe(nullHash);
        expect(hov.size).toBe(0);
        expect(hov.typeCode()).toBe(2);
    });
});

describe('TupleValue', function() {
    test('Empty', function() {
        const emptyTuple = new arb.TupleValue([]);
        expect(emptyTuple.contents).toEqual([]);
        expect(emptyTuple.typeCode()).toBe(3);
        expect(emptyTuple.hash()).toBe(EMPTY_TUPLE_HASH);
    });

    test('Two Tuple', function() {
        const twoTuple = new arb.TupleValue([
            new arb.IntValue(utils.bigNumberify(0)),
            new arb.IntValue(utils.bigNumberify(100)),
        ]);
        expect((twoTuple.get(0) as arb.IntValue).bignum.isZero()).toBe(true);
        expect((twoTuple.get(1) as arb.IntValue).bignum.eq(utils.bigNumberify(100))).toBe(true);
        expect(twoTuple.typeCode()).toBe(3 + 2);
    });

    test('Largest Tuple', function() {
        const mtsv = new arb.TupleValue(Array(arb.MAX_TUPLE_SIZE).fill(new arb.TupleValue([])));
        expect(mtsv.contents.length).toBe(arb.MAX_TUPLE_SIZE);
        for (let i = 0; i < arb.MAX_TUPLE_SIZE; i++) {
            expect((mtsv.get(i) as arb.TupleValue).contents.length).toBe(0);
        }
        expect(mtsv.typeCode()).toBe(3 + arb.MAX_TUPLE_SIZE);
        // Pre calculated hash
        const p = '0xbfa3a3ab6c2f6c71c78354b8186a0e206a99b73b9bb6d8e6d45b733466decbf8';
        expect(mtsv.hash()).toBe(p);
    });

    test('Greater than MAX_TUPLE_SIZE', function() {
        expect(() => new arb.TupleValue(Array(arb.MAX_TUPLE_SIZE + 1))).toThrow(
            'Error TupleValue: illegal size ' + (arb.MAX_TUPLE_SIZE + 1),
        );
    });

    test('get and set', function() {
        const emptyTuple = new arb.TupleValue(Array(arb.MAX_TUPLE_SIZE).fill(new arb.TupleValue([])));
        let t = emptyTuple;
        // set
        for (let i = 0; i < arb.MAX_TUPLE_SIZE; i++) {
            t = t.set(i, new arb.IntValue(utils.bigNumberify(i)));
            expect((t.contents[i] as arb.IntValue).bignum.toNumber()).toBe(i);
        }

        expect(() => t.set(arb.MAX_TUPLE_SIZE, new arb.TupleValue([]))).toThrow(
            'Error TupleValue set: index out of bounds ' + arb.MAX_TUPLE_SIZE,
        );
        expect(() => t.set(-1, new arb.TupleValue([]))).toThrow('Error TupleValue set: index out of bounds ' + -1);

        // get
        for (let i = 0; i < arb.MAX_TUPLE_SIZE; i++) {
            expect((t.get(i) as arb.IntValue).bignum.toNumber()).toBe(i);
        }

        expect(() => t.get(arb.MAX_TUPLE_SIZE)).toThrow(
            'Error TupleValue get: index out of bounds ' + arb.MAX_TUPLE_SIZE,
        );
        expect(() => t.get(-1)).toThrow('Error TupleValue get: index out of bounds ' + -1);
    });
});

describe('BigTuple', function() {
    test('getBigTuple and setBigTuple', function() {
        const emptyBigTup = new arb.TupleValue([]);
        expect((arb.getBigTuple(emptyBigTup, 93) as arb.IntValue).bignum.toNumber()).toBe(0);
        expect((arb.getBigTuple(emptyBigTup, 1234567890) as arb.IntValue).bignum.eq(0)).toBe(true);

        let t = emptyBigTup;
        for (let i = 0; i < 100; i++) {
            t = arb.setBigTuple(t, i, new arb.IntValue(i));
            expect((arb.getBigTuple(t, i) as arb.IntValue).bignum.toNumber()).toBe(i);
        }
    });
});

// Marshaled sizes as hexstrings
const M_INT_VALUE_SIZE = 1 + 32;
const M_CODE_POINT_SIZE = 1 + 8 + 1 + 1 + 0 + 32; // Without val
const M_HASH_ONLY_SIZE = 1 + 8 + 32;
const M_TUPLE_SIZE = 1 + 0; // Without other vals

describe('Marshaling', function() {
    test('marshal and unmarshal IntValue', function() {
        for (const i of [0, 1, 100, '0x9271342394932492394']) {
            const iv = new arb.IntValue(bn(i));
            const marshaledBytes = arb.marshal(iv);
            expect(marshaledBytes.length).toBe(M_INT_VALUE_SIZE);
            const unmarshaledValue = arb.unmarshal(marshaledBytes);
            expect((unmarshaledValue as arb.IntValue).bignum.eq(bn(i))).toBe(true);
        }

        // Test that negative IntValues throw on marshal
        expect(() => arb.marshal(new arb.IntValue(bn(-1)))).toThrow(
            'Error marshaling IntValue: negative values not supported',
        );

        // Test without "0x"
        const iv = new arb.IntValue(bn(99));
        const marshaledBytes = arb.marshal(iv);
        expect(marshaledBytes.length).toBe(M_INT_VALUE_SIZE);
        const unmarshaledValue = arb.unmarshal(marshaledBytes);
        expect((unmarshaledValue as arb.IntValue).bignum.eq(bn(99))).toBe(true);
    });

    test('marshal and unmarshal CodePointValue', function() {
        const pc = 0;
        const op = new arb.BasicOp(arb.OpCode.Halt);
        const nextHash = '0x' + ZEROS_32B;
        const basicTCV = new arb.CodePointValue(pc, op, nextHash);
        const marshaledBytes = arb.marshal(basicTCV);
        expect(marshaledBytes.length).toBe(M_CODE_POINT_SIZE);
        const revValue = arb.unmarshal(marshaledBytes) as arb.CodePointValue;
        expect(revValue.insnNum).toEqual(pc);
        expect(revValue.op.opcode).toBe(op.opcode);
        expect(revValue.nextHash).toEqual(nextHash);
        expect(revValue.toString()).toEqual(basicTCV.toString());

        const iv = new arb.IntValue(bn(60));
        expect(arb.marshal(iv).length).toBe(M_INT_VALUE_SIZE);
        const immTCV = new arb.CodePointValue(pc, new arb.ImmOp(0x19, iv), nextHash);
        const mb = arb.marshal(immTCV);
        expect(mb.length).toBe(M_CODE_POINT_SIZE + M_INT_VALUE_SIZE);
        const revImmValue = arb.unmarshal(mb) as arb.CodePointValue;
        expect(revImmValue.insnNum).toEqual(pc);
        expect(revImmValue.op.opcode).toBe(0x19);
        expect(((revImmValue.op as arb.ImmOp).value as arb.IntValue).bignum.toNumber()).toBe(60);
        expect(revImmValue.nextHash).toEqual(nextHash);
        expect(revImmValue.toString()).toEqual(immTCV.toString());
    });

    test('marshal and unmarshal HashOnlyValue', function() {
        // HashOnlyValue should not be used
        const hv = new arb.HashOnlyValue('0x' + ZEROS_32B, 0);
        const marshaledBytes = arb.marshal(hv);
        expect(marshaledBytes.length).toBe(M_HASH_ONLY_SIZE);
        expect(() => arb.unmarshal(marshaledBytes)).toThrow('Error unmarshaling: HashOnlyValue was not expected');
        expect(hv.toString()).toEqual('HashOnlyValue(' + hv.hash() + ')');
    });

    test('marshal and unmarshal TupleValue', function() {
        // Empty Tuple
        const etv = new arb.TupleValue([]);
        const etvm = arb.marshal(etv);
        expect(etvm.length).toBe(M_TUPLE_SIZE);
        const etvRev = arb.unmarshal(etvm);
        expect(etvRev.toString()).toEqual(etv.toString());

        // Full Tuple of Empty Tuple"s
        const ftv = new arb.TupleValue(Array(8).fill(new arb.TupleValue([])));
        const ftvm = arb.marshal(ftv);
        expect(ftvm.length).toBe(M_TUPLE_SIZE + M_TUPLE_SIZE * 8);
        const ftvRev = arb.unmarshal(ftvm);
        expect(ftvRev.toString()).toEqual(ftv.toString());

        // Full Tuple of IntValue"s
        const fitv = new arb.TupleValue(Array(8).fill(new arb.IntValue(bn(0))));
        const fitvm = arb.marshal(fitv);
        expect(fitvm.length).toBe(M_TUPLE_SIZE + M_INT_VALUE_SIZE * 8);
        const fitvRev = arb.unmarshal(fitvm) as arb.TupleValue;
        expect(fitvRev.toString()).toEqual(fitv.toString());
        expect((fitvRev.get(0) as arb.IntValue).bignum.toNumber()).toBe(0);
        expect((fitvRev.get(7) as arb.IntValue).bignum.toNumber()).toBe(0);
        expect(() => fitvRev.get(8)).toThrow('Error TupleValue get: index out of bounds 8');
    });

    test('illegal inputs', function() {
        // Illegal Value
        expect(() => arb.unmarshal('0x99')).toThrow('Error unmarshaling value no such TYPE: 99');

        const [tyCodePoint, pc, erroneousOpTy] = [
            '0x01',
            Array(8)
                .fill('00')
                .join(''),
            'FF',
        ];
        expect(() => arb.unmarshal(tyCodePoint + pc + erroneousOpTy)).toThrow(
            'Error unmarshalOp no such immCount: 255',
        );

        // Illegal OpCode
        const ILLEGAL_OP_CODE = 'FF';
        const [tyCodePoint2, pc2, immop2] = [
            '0x01',
            Array(8)
                .fill('00')
                .join(''),
            '00',
        ];
        expect(() => arb.unmarshal(tyCodePoint2 + pc2 + immop2 + ILLEGAL_OP_CODE)).toThrow(
            'Error unmarshalOpCode no such opcode: 0xff',
        );

        expect(() => arb.unmarshal('0x01')).toThrow('Error extracting bytes: Uint8Array is too short');
    });
});

describe('Integration', function() {
    test('sizedByteRangeToBytes and hexToSizedByteRange', function() {
        // Create test value
        const myValue = new arb.TupleValue([
            new arb.TupleValue([]),
            new arb.TupleValue([
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.IntValue(bn('0x5ce0c8f1e004fe36aa260ecd02c68ca0c6dea5a4acdfe0b8b10d7b526360046b')),
            ]),
            new arb.TupleValue([
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.IntValue(bn('0x781371cb80a394c637cebf3e3d48a268a44ad21cd68239afb3c3a37196d582c1')),
            ]),
            new arb.TupleValue([
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.TupleValue([]),
                new arb.IntValue(bn('0x32edc9a100000000000000000000000000000000000000000000000000000000')),
            ]),
            new arb.TupleValue([]),
            new arb.TupleValue([]),
            new arb.TupleValue([]),
            new arb.IntValue(bn('0x90e130e5da79003b67479a3ed2caf5585e93ae6771de6cdec6d7641bd2e60180')),
        ]);
        const marshaledBytes = arb.marshal(myValue);
        const expectedMessageBytes =
            '0x0b030b03030303030303005ce0c8f1e004fe36aa260ecd02c68ca0c6dea5a4acdfe0b8b10d7b526360046b0b0303030303030300781371cb80a394c637cebf3e3d48a268a44ad21cd68239afb3c3a37196d582c10b030303030303030032edc9a1000000000000000000000000000000000000000000000000000000000303030090e130e5da79003b67479a3ed2caf5585e93ae6771de6cdec6d7641bd2e60180';
        expect(ethers.utils.hexlify(marshaledBytes)).toBe(expectedMessageBytes);
        const val = arb.unmarshal(expectedMessageBytes);
        expect(val.toString()).toEqual(
            'Tuple([Tuple([]), Tuple([Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), 42009942682379378059947058450083587892049528549641310042571988458584210932843]), Tuple([Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), 54311897307976383387700091809425625413681576808311458615378223571158439396033]), Tuple([Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), 23035776775082914819351730844503175383119388692051102633761137171302117801984]), Tuple([]), Tuple([]), Tuple([]), 65530928266225785593959077233184075030766656784123302077071474652886342173056])',
        );
        const sizedByteRange = new arb.TupleValue([val, new arb.IntValue(bn(100))]);
        const hex = arb.sizedByteRangeToBytes(sizedByteRange);
        expect(ethers.utils.hexlify(hex)).toBe(
            '0x90e130e5da79003b67479a3ed2caf5585e93ae6771de6cdec6d7641bd2e601805ce0c8f1e004fe36aa260ecd02c68ca0c6dea5a4acdfe0b8b10d7b526360046b781371cb80a394c637cebf3e3d48a268a44ad21cd68239afb3c3a37196d582c132edc9a1',
        );
        const sizedByteRangeReverse = arb.hexToSizedByteRange(hex);
        const sizeReverse = sizedByteRangeReverse.get(1);
        expect((sizeReverse as arb.IntValue).bignum.toNumber()).toBe(100);
        const valReverse = sizedByteRangeReverse.get(0);
        const messageReverse = arb.marshal(valReverse);
        expect(ethers.utils.hexlify(messageReverse)).toBe(expectedMessageBytes);
    });
});

describe('test_cases.json', function() {
    for (let i = 0; i < testCases.length; i++) {
        it(testCases[i].name, function() {
            const expectedHash = testCases[i].hash;
            const value = arb.unmarshal('0x' + testCases[i].value);
            const hash = value.hash().slice(2);
            if (hash !== expectedHash) {
                console.log(value.toString());
            }
            expect(hash).toEqual(expectedHash);
        });
    }
});
