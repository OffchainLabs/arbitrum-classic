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

const ethers = require('ethers');
const utils = ethers.utils;

// Valid opcode ranges (inclusive)
const OP_CODE_RANGES =
    [
        [0x00, 0x0b],
        [0x10, 0x1a],
        [0x20, 0x20],
        [0x30, 0x3f],
        [0x40, 0x44],
        [0x50, 0x53],
        [0x60, 0x61],
        [0x70, 0x72],
    ]
const VALID_OP_CODES = OP_CODE_RANGES.reduce(
    ((acc, range) => acc.concat(
        Array(range[1] - range[0] + 1).fill().map((_, i) => range[0] + i)
    )), []
);

// Max tuple size
const MAX_TUPLE_SIZE = 8;

// Arbitrum value type identifiers
const TYPE_INT        = 0;
const TYPE_CODE_POINT = 1;
const TYPE_HASH       = 2;
const TYPE_TUPLE_0    = 3;
const TYPE_TUPLE_MAX  = 3 + MAX_TUPLE_SIZE;

// Flips endianness by reversing the bytes in a hex string. Must not have "0x".
const flipEndian = (s) => s.match(/.{1,2}/g).reverse().join("");

// Extracts first n bytes from s returning two separate strings as list
const extractBytes = (s, n) => {
    if (n < 0 || n*2 > s.length) {
        throw "Error extracting bytes: string is too short";
    }
    return [s.substring(0, n*2), s.substring(n*2, s.length)];
}

// Convert unsigned int i to hexstring of n bytes. Does not include "0x".
const intToBytes = (i, n) => i.toString(16).padStart(n*2, '0');

// Convert BigNumber to hexstring of 32 bytes. Does not include "0x".
const bigNumToBytes = (bn) => bn.toHexString().slice(2).padStart(32*2, '0');

// Operation identifiers
const BASIC_OP_IMM_COUNT = 0;
const IMM_OP_IMM_COUNT   = 1;

class Operation {
    // opcode: 1 byte number
    constructor(opcode) {
        this.opcode = opcode;
    }
}

class BasicOp extends Operation {
    constructor(opcode) {
        super(opcode);
        this.immCount = BASIC_OP_IMM_COUNT;
    }
}

class ImmOp extends Operation {
    constructor(opcode, val) {
        super(opcode);
        this.immCount = IMM_OP_IMM_COUNT;
        this.val = val;
    }
}

class Value {
    hash() { throw "unimplemented" }
    typeCode() { throw "unimplemented" }
}

class IntValue extends Value {
    // bignum: 32 byte integer (BigNumber)
    constructor(bignum) {
        super();
        this.bignum = ethers.utils.bigNumberify(bignum);
        this.typeCode = () => TYPE_INT;
    }

    hash() {
        return utils.solidityKeccak256(['uint256'], [this.bignum]);
    }

    toString() {
        return this.bignum.toString()
    }
}

class CodePointValue extends Value {
    // insnNum: 8 byte integer
    // op: BasicOp or ImmOp
    // nextHash: 32 byte hash
    constructor(insnNum, op, nextHash) {
        super();
        this.insnNum = insnNum;
        this.op = op;
        this.nextHash = nextHash;
        this.typeCode = () => TYPE_CODE_POINT;
        let zero = utils.formatBytes32String("");
        let one = zero.slice(0, -1) + '1';
        this.haltCodePointHash = zero;
        this.errorCodePointHash = one;
    }

    hash() {
        if (this.insnNum === -1) {
            return this.haltCodePointHash();
        } else if (this.insnNum === -2) {
            return this.errorCodePointHash;
        }

        if (this.op instanceof BasicOp) {
            // 34 bytes total (2 + 32)
            let packed = new uint8Array(
                [this.typeCode(), this.op.opcode].concat(this.nextHash)
            );
            return utils.keccak256(packed);
        } else if (this.op instanceof ImmOp) {
            // 66 bytes total (2 + 32 + 32)
            let packed = new uint8Array(
                [this.typeCode(), this.op.opcode].concat(this.op.val.hash())
                    .concat(this.nextHash)
            );
            return utils.keccak256(packed);
        } else {
            throw "Error: CodePointValue must be instanceof BasicOp or ImmOp";
        }
    }
}

class HashOnlyValue extends Value {
    // hash: 32 byte hash
    // size: 8 byte integer
    constructor(hash, size) {
        super();
        this.hash = () => hash;
        this.size = size;
        this.typeCode = () => TYPE_HASH;
    }
}

class TupleValue extends Value {
    // contents: array of Value(s)
    // size: num of Value(s) in contents
    constructor(contents) {
        if (contents.length > MAX_TUPLE_SIZE) {
            throw ("Error TupleValue: illegal size " + contents.length);
        }
        super();
        this.contents = contents;
        this.typeCode = () => TYPE_TUPLE_0 + this.contents.length;
        // Calculate the hash
        this.calcHash = () => {
            let hashes = this.contents.map((value, _) => value.hash());
            let types = ['uint8'].concat(Array(contents.length).fill('bytes32'));
            return utils.solidityKeccak256(types, [this.typeCode()].concat(hashes));
        }
        let hash = this.calcHash();
        this.hash = () => hash;
    }

    // index: uint8
    get(index) {
        if (index < 0 || index >= this.contents.length) {
            throw ("Error TupleValue get: index out of bounds " + index);
        }
        return this.contents[index];
    }

    // Non-mutating
    // index: uint8
    // value: *Value
    set(index, value) {
        if (index < 0 || index >= this.contents.length) {
            throw ("Error TupleValue set: index out of bounds " + index);
        }
        let contents = [...this.contents];
        contents[index] = value;
        return new TupleValue(contents);
    }

    toString() {
        let ret = "Tuple(["
        ret += this.contents.map(val => val.toString()).join(", ")
        ret += "])"
        return ret
    }
}

// Useful for BigTuple operations
const LAST_INDEX = MAX_TUPLE_SIZE-1;
const LAST_INDEX_BIG_NUM = utils.bigNumberify(LAST_INDEX);

// tuple: TupleValue
// index: BigNumber
// returns: *Value
function getBigTuple(tuple, index) {
    if (tuple.contents.length === 0) {
        return new IntValue(utils.bigNumberify(0));
    } else if (index.isZero()) {
        return tuple.get(LAST_INDEX);
    } else {
        let subTup = tuple.get(index.mod(LAST_INDEX_BIG_NUM).toNumber());
        return getBigTuple(subTup, index.div(LAST_INDEX_BIG_NUM));
    }
}

// tuple: TupleValue
// index: BigNumber
// value: *Value
// Non-Mutating returns TupleValue
function setBigTuple(tupleValue, index, value) {
    let tuple = tupleValue;
    if (tuple.contents.length === 0) {
        tuple = new TupleValue(Array(MAX_TUPLE_SIZE).fill(new TupleValue([])));
    }

    if (index.isZero()) {
        return tuple.set(LAST_INDEX, value);
    } else {
        let path = index.mod(LAST_INDEX_BIG_NUM).toNumber();
        let subTup = tuple.get(path);
        let newSubTup = setBigTuple(subTup, index.div(LAST_INDEX_BIG_NUM), value);
        return tuple.set(path, newSubTup);
    }
}

// twoTupleValue: (byterange: SizedTupleValue, size: IntValue)
function sizedByteRangeToHex(twoTupleValue) {
    let byterange = twoTupleValue.get(0);
    let size = twoTupleValue.get(1).bignum;
    let accumulator = '';
    for (let i = utils.bigNumberify(0); i.lt(size); i = i.add(1)) {
        let value = getBigTuple(byterange, i);
        accumulator += value.bignum.toHexString().slice(2);
    }
    return '0x' + accumulator.slice(0, 2*size);
}

// hexString: must be a byte string (hexString.length % 2 === 0)
function hexToSizedByteRange(hexString) {
    let h = hexString;
    // Remove prefix
    if (h.slice(0, 2) === '0x') {
        h = hexString.slice(2);
    }

    // Emtpy tuple
    let t = new TupleValue([]);

    // Array of 32B BigNums
    // TODO: Check endianness of entering into the tuple
    let numBytes = h.length / 2;
    let size = utils.bigNumberify(Math.ceil(numBytes / 32));
    for (let i = utils.bigNumberify(0); i.lt(size); i = i.add(1)) {
        // TODO for hexnum see: https://github.com/OffchainLabs/arb-avm/blob/6b67cd8cdde54735fbe76c7c5667a33e2c742756/evm/data.go#L149
        let hexnum = h.slice(i*32*2, i*32*2 + 32*2).padEnd(2*32, '0');
        let bignum = utils.bigNumberify('0x' + hexnum);
        t = setBigTuple(t, i, new IntValue(bignum));
    }
    return new TupleValue([t, new IntValue(h.length / 2)])
}

function marshal(someValue) {
    return _marshalValue('0x', someValue);
}

function _marshalValue(acc, v) {
    let ty = v.typeCode();
    let accTy = acc + intToBytes(ty, 1);
    if (ty === TYPE_INT) {
        // 1B type; 32B hex int
        return accTy + bigNumToBytes(v.bignum);
    } else if (ty === TYPE_CODE_POINT) {
        // 1B type; 8B LittleEndian insnNum; 1B immCount; 1B opcode; Val?; 32B hash
        let packed = (accTy + flipEndian(intToBytes(v.insnNum, 8)) +
                      intToBytes(v.op.immCount, 1) + intToBytes(v.op.opcode, 1));
        if (v.op.immCount === BASIC_OP_IMM_COUNT) {
            return packed + v.nextHash.slice(2);
        } else if (v.op.immCount === IMM_OP_IMM_COUNT) {
            return _marshalValue(packed, v.op.val) + v.nextHash.slice(2);
        }
    } else if (ty === TYPE_HASH) {
        // 1B type; 8B LittleEndian size; 32B hash
        return accTy + flipEndian(intToBytes(v.size, 8)) + v.hash.slice(2);
    } else if (ty >= TYPE_TUPLE_0 && ty <= TYPE_TUPLE_MAX) {
        // 1B type; (ty-TYPE_TUPLE_0) number of Values
        for (let i = 0; i < v.contents.length; i++) {
            accTy = _marshalValue(accTy, v.contents[i]);
        }
        return accTy;
    }
}

function unmarshal(hexString) {
    let h = hexString;
    // Remove prefix if exists
    if (h.slice(0, 2) === '0x') {
        h = h.slice(2);
    }
    return _unmarshalValue(h)[0];
}

function _unmarshalValue(hexString) {
    let [head, tail] = extractBytes(hexString, 1);

    let ty = parseInt(head, 16);
    if (ty === TYPE_INT) {
        [head, tail] = extractBytes(tail, 32);
        let i = utils.bigNumberify('0x' + head);
        return [new IntValue(i), tail];
    } else if (ty === TYPE_CODE_POINT) {
        [head, tail] = extractBytes(tail, 8);
        let pc = utils.bigNumberify(flipEndian(head)); // pc encoded as Little Endian
        let [op, tail] = unmarshalOp(tail);
        [head, tail] = extractBytes(tail, 32);
        let nextHash = '0x' + head;
        return [new CodePointValue(pc, op, nextHash), tail];
    } else if (ty === TYPE_HASH) {
        throw "Error unmarshalling: hash only value was not expected";
        [head, tail] = extractBytes(tail, 8);
        let size = parseInt(flipEndian(head), 16); // size encoded as Little Endian
        [head, tail] = extractBytes(tail, 32);
        let hash = '0x' + head;
        return [new HashOnlyValue(hash, size), tail];
    } else if (ty >= TYPE_TUPLE_0 && ty <= TYPE_TUPLE_MAX) {
        let size = ty - TYPE_TUPLE_0;
        [contents, tail] = unmarshalTuple(tail, size);
        return [new TupleValue(contents), tail];
    } else {
        throw ("Error unmarshalling: invalid value type " + ty);
    }
}

function unmarshalOp(hexString) {
    let [head, tail] = extractBytes(hexString, 1);
    let immCount = parseInt(head, 16);
    if (immCount === BASIC_OP_IMM_COUNT) {
        [opcode, tail] = unmarshalOpCode(tail);
        return [BasicOp(opcode), tail]
    } else if (immCount === IMM_OP_IMM_COUNT) {
        [opcode, tail] = unmarshalOpCode(tail);
        [value, tail] = _unmarshalValue(tail);
        return [ImmOp(opcode, value), tail]
    } else {
        throw "Error unmarshalling operand: immediate count must be 0 or 1";
    }
}

function unmarshalOpCode(hexString) {
    let [head, tail] = extractBytes(hexString, 1);
    let opcode = parseInt(head, 16);
    if (!VALID_OP_CODES.includes(opcode)) {
        throw ("Error unmarshalling: Invalid opcode: " + opcode)
    }
    return opcode;
}

function unmarshalTuple(hexString, size) {
    let contents = new Array(size);
    let value = undefined;
    let tail = hexString;
    for (let i = 0; i < size; i++) {
        [value, tail] = _unmarshalValue(tail);
        contents[i] = value;
    }
    return [contents, tail];
}

module.exports = {
    Value: Value,
    TupleValue: TupleValue,
    IntValue: IntValue,
    CodePointValue: CodePointValue,

    marshal: marshal,
    unmarshal: unmarshal,

    hexToSizedByteRange: hexToSizedByteRange,
    sizedByteRangeToHex: sizedByteRangeToHex
}
