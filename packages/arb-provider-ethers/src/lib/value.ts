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

function assertNever(x: never): never {
    throw new Error('Unexpected object: ' + x);
}

// Error and Halt opcodes
export enum OpCode {
    Error = 0x74,
    Halt = 0x75,
}

// Valid opcode ranges (inclusive)
const OP_CODE_RANGES: Array<[number, number]> = [
    [0x00, 0x0a],
    [0x10, 0x1b],
    [0x20, 0x22],
    [0x30, 0x3d],
    [0x40, 0x44],
    [0x50, 0x52],
    [0x60, 0x61],
    [0x70, 0x75],
];

const VALID_OP_CODES = OP_CODE_RANGES.reduce(
    (acc: number[], range: [number, number]) =>
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
function extractBytes(s: Uint8Array, offset: number, n: number): [Uint8Array, number] {
    if (n < 0 || n > s.length) {
        throw Error('Error extracting bytes: Uint8Array is too short');
    }
    return [s.slice(offset, offset + n), offset + n];
}

// Convert unsigned int i to byte array of n bytes.
function intToBytes(i: ethers.utils.BigNumberish, n: number): Uint8Array {
    return ethers.utils.padZeros(ethers.utils.arrayify(ethers.utils.bigNumberify(i)), n);
}

// Convert unsigned BigNumber to hexstring of 32 bytes. Does not include "0x".
function uBigNumToBytes(bn: ethers.utils.BigNumber): Uint8Array {
    return ethers.utils.padZeros(ethers.utils.arrayify(bn), 32);
}

export enum OperationType {
    Basic = 0,
    Immediate = 1,
}

type Operation = BasicOp | ImmOp;

export class BasicOp {
    public opcode: number;
    public kind: OperationType.Basic;

    constructor(opcode: number) {
        this.opcode = opcode;
        this.kind = OperationType.Basic;
    }
}

export class ImmOp {
    public opcode: number;
    public value: Value;
    public kind: OperationType.Immediate;

    constructor(opcode: number, value: Value) {
        this.opcode = opcode;
        this.value = value;
        this.kind = OperationType.Immediate;
    }
}

export type Value = IntValue | TupleValue | CodePointValue | HashOnlyValue;

export class IntValue {
    public bignum: ethers.utils.BigNumber;

    constructor(bignum: ethers.utils.BigNumberish) {
        this.bignum = ethers.utils.bigNumberify(bignum);
    }

    public typeCode(): ValueType {
        return ValueType.Int;
    }

    public hash(): string {
        return ethers.utils.solidityKeccak256(['uint256'], [this.bignum]);
    }

    public toString(): string {
        return this.bignum.toString();
    }
}

export class CodePointValue {
    public insnNum: ethers.utils.BigNumber;
    public op: Operation;
    public nextHash: string;
    // insnNum: 8 byte integer
    // op: BasicOp or ImmOp
    // nextHash: 32 byte hash
    constructor(insnNum: ethers.utils.BigNumberish, op: Operation, nextHash: string) {
        this.insnNum = ethers.utils.bigNumberify(insnNum);
        this.op = op;
        this.nextHash = nextHash;
    }

    public typeCode(): ValueType {
        return ValueType.CodePoint;
    }

    public hash(): string {
        switch (this.op.kind) {
            case OperationType.Basic: {
                return ethers.utils.solidityKeccak256(
                    ['uint8', 'uint8', 'bytes32'],
                    [this.typeCode(), this.op.opcode, this.nextHash],
                );
            }
            case OperationType.Immediate: {
                return ethers.utils.solidityKeccak256(
                    ['uint8', 'uint8', 'bytes32', 'bytes32'],
                    [this.typeCode(), this.op.opcode, this.op.value.hash(), this.nextHash],
                );
            }
            default:
                assertNever(this.op);
                return '';
        }
    }

    public toString(): string {
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
    public hashVal: string;
    public size: ethers.utils.BigNumber;

    // hash: 32 byte hash
    // size: 8 byte integer
    constructor(hash: string, size: ethers.utils.BigNumberish) {
        this.hashVal = hash;
        this.size = ethers.utils.bigNumberify(size);
    }

    public typeCode(): ValueType {
        return ValueType.HashOnly;
    }

    public hash(): string {
        return this.hashVal;
    }

    public toString(): string {
        return 'HashOnlyValue(' + this.hash() + ')';
    }
}

export class TupleValue {
    public contents: Value[];
    public cachedHash: string;
    // contents: array of Value(s)
    // size: num of Value(s) in contents
    constructor(contents: Value[]) {
        if (contents.length > MAX_TUPLE_SIZE) {
            throw Error('Error TupleValue: illegal size ' + contents.length);
        }
        this.contents = contents;
        const hashes = this.contents.map((value): string => value.hash());
        const types = ['uint8'].concat(Array(contents.length).fill('bytes32'));
        const values: any[] = [this.typeCode()];
        this.cachedHash = ethers.utils.solidityKeccak256(types, values.concat(hashes));
    }

    public typeCode(): ValueType {
        return ValueType.Tuple + this.contents.length;
    }

    public hash(): string {
        return this.cachedHash;
    }

    // index: uint8
    public get(index: number): Value {
        if (index < 0 || index >= this.contents.length) {
            throw Error('Error TupleValue get: index out of bounds ' + index);
        }
        return this.contents[index];
    }

    // Non-mutating
    // index: uint8
    // value: *Value
    public set(index: number, value: Value): TupleValue {
        if (index < 0 || index >= this.contents.length) {
            throw Error('Error TupleValue set: index out of bounds ' + index);
        }
        const contents = [...this.contents];
        contents[index] = value;
        return new TupleValue(contents);
    }

    public toString(): string {
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
    } else if (index === 0) {
        return tuple.get(LAST_INDEX);
    } else {
        const path = index % LAST_INDEX_BIG_NUM;
        const subTup = tuple.get(path) as TupleValue;
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

    if (index === 0) {
        return tuple.set(LAST_INDEX, value);
    } else {
        const path = index % LAST_INDEX_BIG_NUM;
        const subTup = tuple.get(path) as TupleValue;
        const newSubTup = setBigTuple(subTup, Math.floor(index / LAST_INDEX_BIG_NUM), value);
        return tuple.set(path, newSubTup);
    }
}

function bytesToIntValues(bytearray: Uint8Array): ethers.utils.BigNumber[] {
    const bignums: ethers.utils.BigNumber[] = [];
    const sizeBytes = bytearray.length;
    const chunkCount = Math.ceil(sizeBytes / 32);
    for (let i = 0; i < chunkCount; i++) {
        const byteSlice = bytearray.slice(i * 32, (i + 1) * 32);
        const nextNumBytes = new Uint8Array(32);
        nextNumBytes.set(byteSlice);
        const bignum = ethers.utils.bigNumberify(nextNumBytes);
        bignums.push(bignum);
    }
    return bignums;
}

// twoTupleValue: (byterange: SizedTupleValue, size: IntValue)
export function sizedByteRangeToBytes(twoTupleValue: TupleValue): Uint8Array {
    const byterange = twoTupleValue.get(0) as TupleValue;
    const sizeInt = twoTupleValue.get(1) as IntValue;
    const sizeBytes = sizeInt.bignum.toNumber();
    const chunkCount = Math.ceil(sizeBytes / 32);
    const result = new Uint8Array(chunkCount * 32);
    for (let i = 0; i < chunkCount; i++) {
        const value = getBigTuple(byterange, i) as IntValue;
        result.set(ethers.utils.padZeros(ethers.utils.arrayify(value.bignum), 32), i * 32);
    }
    return result.slice(0, sizeBytes);
}

// hexString: must be a byte string (hexString.length % 2 === 0)
export function hexToSizedByteRange(hex: ethers.utils.Arrayish): TupleValue {
    const bytearray = ethers.utils.arrayify(hex);
    const sizeBytes = bytearray.length;
    const bignums = bytesToIntValues(bytearray);
    // Empty tuple
    let t = new TupleValue([]);
    for (let i = 0; i < bignums.length; i++) {
        t = setBigTuple(t, i, new IntValue(bignums[i]));
    }
    return new TupleValue([t, new IntValue(sizeBytes)]);
}

// twoTupleValue: (byterange: SizedTupleValue, size: IntValue)
export function bytestackToBytes(twoTupleValue: TupleValue): Uint8Array {
    const sizeInt = twoTupleValue.get(0) as IntValue;
    let stack = twoTupleValue.get(1) as TupleValue;

    const sizeBytes = sizeInt.bignum.toNumber();
    const chunkCount = Math.ceil(sizeBytes / 32);
    const result = new Uint8Array(chunkCount * 32);

    let i = 0;
    while (stack.contents.length == 2) {
        const value = stack.get(1) as IntValue;
        stack = stack.get(0) as TupleValue;
        const chunk = ethers.utils.padZeros(ethers.utils.arrayify(value.bignum), 32);
        const offset = (chunkCount - 1 - i) * 32;
        result.set(chunk, offset);
        i++;
    }
    return result.slice(0, sizeBytes);
}

// hexString: must be a byte string (hexString.length % 2 === 0)
export function hexToBytestack(hex: ethers.utils.Arrayish): TupleValue {
    const bytearray = ethers.utils.arrayify(hex);
    const sizeBytes = bytearray.length;
    const bignums = bytesToIntValues(bytearray);
    // Empty tuple
    let t = new TupleValue([]);
    for (let i = 0; i < bignums.length; i++) {
        t = new TupleValue([t, new IntValue(bignums[i])]);
    }
    return new TupleValue([new IntValue(sizeBytes), t]);
}

function _marshalValue(acc: Uint8Array, v: Value): Uint8Array {
    const ty = v.typeCode();
    let accTy = ethers.utils.concat([acc, intToBytes(ty, 1)]);
    if (ty === ValueType.Int) {
        const val = v as IntValue;
        // 1B type; 32B hex int
        if (val.bignum.lt(0)) {
            throw Error('Error marshaling IntValue: negative values not supported');
        }
        return ethers.utils.concat([accTy, uBigNumToBytes(val.bignum)]);
    } else if (ty === ValueType.CodePoint) {
        const val = v as CodePointValue;
        // 1B type; 8B insnNum; 1B immCount; 1B opcode; Val?; 32B hash
        const packed = ethers.utils.concat([
            accTy,
            intToBytes(val.insnNum, 8),
            intToBytes(val.op.kind, 1),
            intToBytes(val.op.opcode, 1),
        ]);
        switch (val.op.kind) {
            case OperationType.Basic: {
                return ethers.utils.concat([packed, val.nextHash]);
            }
            case OperationType.Immediate: {
                const op = val.op as ImmOp;
                return ethers.utils.concat([_marshalValue(packed, op.value), val.nextHash]);
            }
            default: {
                assertNever(val.op);
                return new Uint8Array();
            }
        }
    } else if (ty === ValueType.HashOnly) {
        const val = v as HashOnlyValue;
        // 1B type; 8B size; 32B hash
        return ethers.utils.concat([accTy, intToBytes(val.size, 8), val.hash()]);
    } else if (ty >= ValueType.Tuple && ty <= ValueType.TupleMax) {
        const val = v as TupleValue;
        // 1B type; (ty-TYPE_TUPLE_0) number of Values
        for (const subVal of val.contents) {
            accTy = _marshalValue(accTy, subVal);
        }
        return accTy;
    } else {
        throw Error('Error marshaling value no such TYPE: ' + ty);
    }
}

export function marshal(someValue: Value): Uint8Array {
    return _marshalValue(new Uint8Array(), someValue);
}

function unmarshalContract(array: Uint8Array): [Operation[], Value] {
    let offset = 0;
    let versionBytes;
    [versionBytes, offset] = extractBytes(array, offset, 4);
    let extensionVersion = ethers.utils.bigNumberify(1);
    while (!extensionVersion.eq(0)) {
        let extensionVersionBytes;
        [extensionVersionBytes, offset] = extractBytes(array, offset, 4);
        extensionVersion = ethers.utils.bigNumberify(extensionVersionBytes);
        if (!extensionVersion.eq(0)) {
            let extensionLengthBytes;
            [extensionLengthBytes, offset] = extractBytes(array, offset, 4);
            const extensionLength = ethers.utils.bigNumberify(extensionLengthBytes);
            offset += extensionLength.toNumber();
        }
    }

    let codeCountBytes;
    [codeCountBytes, offset] = extractBytes(array, offset, 8);
    const codeCount = ethers.utils.bigNumberify(codeCountBytes).toNumber();
    const ops: Operation[] = [];
    for (let i = 0; i < codeCount; ++i) {
        let op;
        [op, offset] = unmarshalOp(array, offset);
        ops.push(op);
    }
    const [staticVal] = _unmarshalValue(array, offset);
    return [ops, staticVal];
}

function opsToCodePoints(ops: Operation[]): CodePointValue[] {
    const cps: CodePointValue[] = [];
    for (const op of ops) {
        cps.push(new CodePointValue(0, op, ethers.utils.hexZeroPad('0x00', 32)));
    }
    for (let i = cps.length - 2; i >= 0; i--) {
        cps[i].nextHash = cps[i + 1].hash();
    }
    return cps;
}

export function contractMachineHash(array: Uint8Array): string {
    const [ops, staticVal] = unmarshalContract(array);
    const codePoints = opsToCodePoints(ops);

    return machineHash(
        codePoints[0],
        new TupleValue([]),
        new TupleValue([]),
        new TupleValue([]),
        staticVal,
        new CodePointValue('18446744073709551615', new BasicOp(0), ethers.utils.hexZeroPad('0x00', 32)),
    );
}

export function machineHash(
    pc: Value,
    stack: Value,
    auxstack: Value,
    registerVal: Value,
    staticVal: Value,
    errPc: Value,
): string {
    return ethers.utils.solidityKeccak256(
        ['bytes32', 'bytes32', 'bytes32', 'bytes32', 'bytes32', 'bytes32'],
        [pc.hash(), stack.hash(), auxstack.hash(), registerVal.hash(), staticVal.hash(), errPc.hash()],
    );
}

function unmarshalOpCode(array: Uint8Array, offset: number): [number, number] {
    let head;
    [head, offset] = extractBytes(array, offset, 1);
    const opcode = ethers.utils.bigNumberify(head).toNumber();
    if (!VALID_OP_CODES.includes(opcode)) {
        throw Error('Error unmarshalOpCode no such opcode: 0x' + opcode.toString(16));
    }
    return [opcode, offset];
}

function unmarshalOp(array: Uint8Array, offset: number): [Operation, number] {
    let head = new Uint8Array();
    [head, offset] = extractBytes(array, offset, 1);
    const kind: OperationType = ethers.utils.bigNumberify(head).toNumber();
    if (kind === OperationType.Basic) {
        let opcode;
        [opcode, offset] = unmarshalOpCode(array, offset);
        return [new BasicOp(opcode), offset];
    } else if (kind === OperationType.Immediate) {
        let opcode: number;
        let value: Value;
        [opcode, offset] = unmarshalOpCode(array, offset);
        [value, offset] = _unmarshalValue(array, offset);
        return [new ImmOp(opcode, value), offset];
    } else {
        throw Error('Error unmarshalOp no such immCount: ' + kind);
    }
}

function unmarshalTuple(size: number, array: Uint8Array, offset: number): [Value[], number] {
    const contents = new Array(size);
    const tail = array;
    for (let i = 0; i < size; i++) {
        let value;
        [value, offset] = _unmarshalValue(array, offset);
        contents[i] = value;
    }
    return [contents, offset];
}

function _unmarshalValue(array: Uint8Array, offset: number): [Value, number] {
    let head;
    [head, offset] = extractBytes(array, offset, 1);
    const ty = ethers.utils.bigNumberify(head).toNumber();
    if (ty === ValueType.Int) {
        [head, offset] = extractBytes(array, offset, 32);
        const i = ethers.utils.bigNumberify(head);
        return [new IntValue(i), offset];
    } else if (ty === ValueType.CodePoint) {
        [head, offset] = extractBytes(array, offset, 8);
        const pc = ethers.utils.bigNumberify(head);
        let op;
        [op, offset] = unmarshalOp(array, offset);
        [head, offset] = extractBytes(array, offset, 32);
        const nextHash = ethers.utils.hexlify(head);
        return [new CodePointValue(pc, op, nextHash), offset];
    } else if (ty === ValueType.HashOnly) {
        [head, offset] = extractBytes(array, offset, 8);
        const size = ethers.utils.bigNumberify(head);
        [head, offset] = extractBytes(array, offset, 32);
        const hash = '0x' + head;
        // return [new HashOnlyValue(hash, size), tail];
        throw Error('Error unmarshaling: HashOnlyValue was not expected');
    } else if (ty >= ValueType.Tuple && ty <= ValueType.TupleMax) {
        const size = ty - ValueType.Tuple;
        let contents;
        [contents, offset] = unmarshalTuple(size, array, offset);
        return [new TupleValue(contents), offset];
    } else {
        throw Error('Error unmarshaling value no such TYPE: ' + ty.toString(16));
    }
}

export function unmarshal(array: ethers.utils.Arrayish): Value {
    return _unmarshalValue(ethers.utils.arrayify(array), 0)[0];
}
