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
/* eslint-env node */
'use strict';
import * as ethers from 'ethers';

// Error and Halt opcodes
export enum OpCode {
    Error = 0x74,
    Halt = 0x75,
}

// Valid opcode ranges (inclusive)
const OP_CODE_RANGES: Array<[number, number]> = [
    [0x01, 0x0a],
    [0x10, 0x1b],
    [0x20, 0x21],
    [0x30, 0x3d],
    [0x40, 0x44],
    [0x50, 0x52],
    [0x60, 0x61],
    [0x70, 0x75],
];

const VALID_OP_CODES = OP_CODE_RANGES.reduce(
    (acc: Array<number>, range: [number, number]) =>
        acc.concat(
            Array(range[1] - range[0] + 1)
                .fill(0)
                .map((_, i) => range[0] + i),
        ),
    [],
);

// Max tuple size
export const MAX_TUPLE_SIZE = 8;

export enum ValueType {
    Int = 0,
    CodePoint = 1,
    HashOnly = 2,
    Tuple = 3,
    TupleMax = 3 + MAX_TUPLE_SIZE,
}

// Extracts first n bytes from s returning two separate strings as list
function extractBytes(s: Uint8Array, n: number): [Uint8Array, Uint8Array] {
    if (n < 0 || n > s.length) {
        throw 'Error extracting bytes: Uint8Array is too short';
    }
    return [s.slice(0, n), s.slice(n)];
}

// Convert unsigned int i to byte array of n bytes.
function intToBytes(i: number, n: number): Uint8Array {
    return ethers.utils.padZeros(ethers.utils.arrayify(ethers.utils.bigNumberify(i)), n);
}

// Convert unsigned BigNumber to hexstring of 32 bytes. Does not include "0x".
var uBigNumToBytes = function(bn: ethers.utils.BigNumber): Uint8Array {
    return ethers.utils.padZeros(ethers.utils.arrayify(bn), 32);
};

// Operation identifiers
const BASIC_OP_IMM_COUNT = 0;
const IMM_OP_IMM_COUNT = 1;

export enum OperationType {
    Basic = 0,
    Immediate = 1,
}

type Operation = BasicOp | ImmOp;

export class BasicOp {
    opcode: number;
    kind: OperationType.Basic;

    constructor(opcode: number) {
        this.opcode = opcode;
        this.kind = OperationType.Basic;
    }
}

export class ImmOp {
    opcode: number;
    value: Value;
    kind: OperationType.Immediate;

    constructor(opcode: number, value: Value) {
        this.opcode = opcode;
        this.value = value;
        this.kind = OperationType.Immediate;
    }
}

export type Value = IntValue | TupleValue | CodePointValue | HashOnlyValue;

export class IntValue {
    bignum: ethers.utils.BigNumber;

    constructor(bignum: ethers.utils.BigNumberish) {
        this.bignum = ethers.utils.bigNumberify(bignum);
    }

    typeCode() {
        return ValueType.Int;
    }

    hash(): string {
        return ethers.utils.solidityKeccak256(['uint256'], [this.bignum]);
    }

    toString() {
        return this.bignum.toString();
    }
}

export class CodePointValue {
    insnNum: number;
    op: Operation;
    nextHash: string;
    // insnNum: 8 byte integer
    // op: BasicOp or ImmOp
    // nextHash: 32 byte hash
    constructor(insnNum: number, op: Operation, nextHash: string) {
        this.insnNum = insnNum;
        this.op = op;
        this.nextHash = nextHash;
    }

    typeCode() {
        return ValueType.CodePoint;
    }

    hash(): string {
        switch (this.op.kind) {
            case OperationType.Basic: {
                let packed =
                    '0x' +
                    this.typeCode()
                        .toString()
                        .padStart(2, '0') +
                    this.op.opcode.toString().padStart(2, '0') +
                    this.nextHash.slice(2);
                return ethers.utils.keccak256(packed);
            }
            case OperationType.Immediate: {
                let packed =
                    '0x' +
                    this.typeCode()
                        .toString()
                        .padStart(2, '0') +
                    this.op.opcode.toString().padStart(2, '0') +
                    this.op.value.hash().slice(2) +
                    this.nextHash.slice(2);
                return ethers.utils.keccak256(packed);
            }
            default:
                assertNever(this.op);
                return '';
        }
    }

    toString(): string {
        switch (this.op.kind) {
            case OperationType.Basic: {
                return 'Basic(OpCode(0x' + this.op.opcode.toString() + '))';
            }
            case OperationType.Immediate: {
                return 'Immediate(OpCode(0x' + this.op.opcode.toString() + '), ' + this.op.value.toString() + ')';
            }
            default:
                assertNever(this.op);
                return '';
        }
    }
}

export class HashOnlyValue {
    hashVal: string;
    size: number;

    // hash: 32 byte hash
    // size: 8 byte integer
    constructor(hash: string, size: number) {
        this.hashVal = hash;
        this.size = size;
    }

    typeCode(): number {
        return ValueType.HashOnly;
    }

    hash(): string {
        return this.hashVal;
    }

    toString() {
        return 'HashOnlyValue(' + this.hash() + ')';
    }
}

export class TupleValue {
    contents: Array<Value>;
    cachedHash: string;
    // contents: array of Value(s)
    // size: num of Value(s) in contents
    constructor(contents: Array<Value>) {
        if (contents.length > MAX_TUPLE_SIZE) {
            throw 'Error TupleValue: illegal size ' + contents.length;
        }
        this.contents = contents;
        let hashes = this.contents.map((value, _) => value.hash());
        let types = ['uint8'].concat(Array(contents.length).fill('bytes32'));
        let values: Array<any> = [this.typeCode()];
        this.cachedHash = ethers.utils.solidityKeccak256(types, values.concat(hashes));
    }

    typeCode(): number {
        return ValueType.Tuple + this.contents.length;
    }

    hash(): string {
        return this.cachedHash;
    }

    // index: uint8
    get(index: number) {
        if (index < 0 || index >= this.contents.length) {
            throw 'Error TupleValue get: index out of bounds ' + index;
        }
        return this.contents[index];
    }

    // Non-mutating
    // index: uint8
    // value: *Value
    set(index: number, value: Value) {
        if (index < 0 || index >= this.contents.length) {
            throw 'Error TupleValue set: index out of bounds ' + index;
        }
        let contents = [...this.contents];
        contents[index] = value;
        return new TupleValue(contents);
    }

    toString() {
        let ret = 'Tuple([';
        ret += this.contents.map(val => val.toString()).join(', ');
        ret += '])';
        return ret;
    }
}

// Useful for BigTuple operations
const LAST_INDEX = MAX_TUPLE_SIZE - 1;
const LAST_INDEX_BIG_NUM = LAST_INDEX;

// tuple: TupleValue
// index: BigNumber
// returns: *Value
export function getBigTuple(tuple: TupleValue, index: number): Value {
    if (tuple.contents.length === 0) {
        return new IntValue(ethers.utils.bigNumberify(0));
    } else if (index == 0) {
        return tuple.get(LAST_INDEX);
    } else {
        let path = index % LAST_INDEX_BIG_NUM;
        let subTup = tuple.get(path) as TupleValue;
        return getBigTuple(subTup, Math.floor(index / LAST_INDEX_BIG_NUM));
    }
}

// tuple: TupleValue
// index: BigNumber
// value: *Value
// Non-Mutating returns TupleValue
export function setBigTuple(tupleValue: TupleValue, index: number, value: Value): TupleValue {
    let tuple = tupleValue;
    if (tuple.contents.length === 0) {
        tuple = new TupleValue(Array(MAX_TUPLE_SIZE).fill(new TupleValue([])));
    }

    if (index == 0) {
        return tuple.set(LAST_INDEX, value);
    } else {
        let path = index % LAST_INDEX_BIG_NUM;
        let subTup = tuple.get(path) as TupleValue;
        let newSubTup = setBigTuple(subTup, Math.floor(index / LAST_INDEX_BIG_NUM), value);
        return tuple.set(path, newSubTup);
    }
}

// twoTupleValue: (byterange: SizedTupleValue, size: IntValue)
export function sizedByteRangeToBytes(twoTupleValue: TupleValue): Uint8Array {
    let byterange = twoTupleValue.get(0) as TupleValue;
    let sizeInt = twoTupleValue.get(1) as IntValue;
    let sizeBytes = sizeInt.bignum.toNumber();
    let chunkCount = Math.ceil(sizeBytes / 32);
    let result = new Uint8Array(chunkCount * 32);
    for (let i = 0; i < chunkCount; i++) {
        let value = getBigTuple(byterange, i) as IntValue;
        result.set(ethers.utils.arrayify(value.bignum), i * 32);
    }
    return result.slice(0, sizeBytes);
}

// hexString: must be a byte string (hexString.length % 2 === 0)
export function hexToSizedByteRange(hex: ethers.utils.Arrayish): TupleValue {
    let bytearray = ethers.utils.arrayify(hex);

    // Emtpy tuple
    let t = new TupleValue([]);

    // Array of 32B BigNums
    let sizeBytes = bytearray.length;
    let chunkCount = Math.ceil(sizeBytes / 32);
    for (let i = 0; i < chunkCount; i++) {
        let byteSlice = bytearray.slice(i * 32, (i + 1) * 32);
        let nextNumBytes = new Uint8Array(32);
        nextNumBytes.set(byteSlice);
        let bignum = ethers.utils.bigNumberify(nextNumBytes);
        t = setBigTuple(t, i, new IntValue(bignum));
    }
    return new TupleValue([t, new IntValue(sizeBytes)]);
}

export function marshal(someValue: Value): Uint8Array {
    return _marshalValue(new Uint8Array(), someValue);
}

function assertNever(x: never): never {
    throw new Error('Unexpected object: ' + x);
}

function _marshalValue(acc: Uint8Array, v: Value): Uint8Array {
    let ty = v.typeCode();
    let accTy = ethers.utils.concat([acc, intToBytes(ty, 1)]);
    if (ty === ValueType.Int) {
        let val = v as IntValue;
        // 1B type; 32B hex int
        if (val.bignum.lt(0)) {
            throw 'Error marshaling IntValue: negative values not supported';
        }
        return ethers.utils.concat([accTy, uBigNumToBytes(val.bignum)]);
    } else if (ty === ValueType.CodePoint) {
        let val = v as CodePointValue;
        // 1B type; 8B insnNum; 1B immCount; 1B opcode; Val?; 32B hash
        let packed = ethers.utils.concat([
            accTy,
            intToBytes(val.insnNum, 8),
            intToBytes(val.op.kind, 1),
            intToBytes(val.op.opcode, 1),
        ]);
        switch (val.op.kind) {
            case OperationType.Basic:
                return ethers.utils.concat([packed, val.nextHash]);
            case OperationType.Immediate:
                let op = val.op as ImmOp;
                return ethers.utils.concat([_marshalValue(packed, op.value), val.nextHash]);
            default:
                assertNever(val.op);
                return new Uint8Array();
        }
    } else if (ty === ValueType.HashOnly) {
        let val = v as HashOnlyValue;
        // 1B type; 8B size; 32B hash
        return ethers.utils.concat([accTy, intToBytes(val.size, 8), val.hash()]);
    } else if (ty >= ValueType.Tuple && ty <= ValueType.TupleMax) {
        let val = v as TupleValue;
        // 1B type; (ty-TYPE_TUPLE_0) number of Values
        for (let i = 0; i < val.contents.length; i++) {
            accTy = _marshalValue(accTy, val.contents[i]);
        }
        return accTy;
    } else {
        throw 'Error marshaling value no such TYPE: ' + ty;
    }
}

export function unmarshal(array: ethers.utils.Arrayish) {
    return _unmarshalValue(ethers.utils.arrayify(array))[0];
}

function _unmarshalValue(array: Uint8Array): [Value, Uint8Array] {
    var head, tail, contents, op;
    [head, tail] = extractBytes(array, 1);

    let ty = ethers.utils.bigNumberify(head).toNumber();
    if (ty === ValueType.Int) {
        [head, tail] = extractBytes(tail, 32);
        let i = ethers.utils.bigNumberify(head);
        return [new IntValue(i), tail];
    } else if (ty === ValueType.CodePoint) {
        [head, tail] = extractBytes(tail, 8);
        let pc = ethers.utils.bigNumberify(head).toNumber();
        [op, tail] = unmarshalOp(tail);
        [head, tail] = extractBytes(tail, 32);
        let nextHash = ethers.utils.hexlify(head);
        return [new CodePointValue(pc, op, nextHash), tail];
    } else if (ty === ValueType.HashOnly) {
        [head, tail] = extractBytes(tail, 8);
        let size = ethers.utils.bigNumberify(head);
        [head, tail] = extractBytes(tail, 32);
        let hash = '0x' + head;
        // return [new HashOnlyValue(hash, size), tail];
        throw 'Error unmarshaling: HashOnlyValue was not expected';
    } else if (ty >= ValueType.Tuple && ty <= ValueType.TupleMax) {
        let size = ty - ValueType.Tuple;
        [contents, tail] = unmarshalTuple(tail, size);
        return [new TupleValue(contents), tail];
    } else {
        throw 'Error unmarshaling value no such TYPE: ' + ty.toString(16);
    }
}

function unmarshalOp(array: Uint8Array): [Operation, Uint8Array] {
    var head, tail, opcode, value;
    [head, tail] = extractBytes(array, 1);
    let kind: OperationType = ethers.utils.bigNumberify(head).toNumber();
    if (kind === OperationType.Basic) {
        [opcode, tail] = unmarshalOpCode(tail);
        return [new BasicOp(opcode), tail];
    } else if (kind === OperationType.Immediate) {
        [opcode, tail] = unmarshalOpCode(tail);
        [value, tail] = _unmarshalValue(tail);
        return [new ImmOp(opcode, value), tail];
    } else {
        throw 'Error unmarshalOp no such immCount: ' + kind;
    }
}

function unmarshalOpCode(array: Uint8Array): [number, Uint8Array] {
    let [head, tail] = extractBytes(array, 1);
    let opcode = ethers.utils.bigNumberify(head).toNumber();
    if (!VALID_OP_CODES.includes(opcode)) {
        throw 'Error unmarshalOpCode no such opcode: 0x' + opcode.toString(16);
    }
    return [opcode, tail];
}

function unmarshalTuple(array: Uint8Array, size: number): [Array<Value>, Uint8Array] {
    let contents = new Array(size);
    var value;
    let tail = array;
    for (let i = 0; i < size; i++) {
        [value, tail] = _unmarshalValue(tail);
        contents[i] = value;
    }
    return [contents, tail];
}
