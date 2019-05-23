const ethers = require('ethers');

// Valid opcode ranges (inclusive)
const OP_CODE_RANGES =
    [
        [0x00, 0x0b],
        [0x10, 0x1a],
        [0x20, 0x20],
        [0x30, 0x3e],
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

// Flips endianness by reversing the bytes in a hex string
const flipEndian = (s) => { s.match(/.{1,2}/g).reverse().join("") }

// Extracts first n bytes from s returning two separate strings as list
const extractBytes = (s, n) => {
    if (n < 0 || n*2 > s.length) {
        throw "Error extracting bytes: string is too short";
    }
    return [s.substring(0, n*2), s.substring(n*2, s.length)];
}

class Operation {
    // opcode: 1 byte number
    constructor(opcode) {
        this.opcode = opcode;
    }
}

class BasicOp extends Operation {}

class ImmOp extends Operation {
    constructor(opcode, val) {
        super(opcode);
        this.val = val;
    }
}

class Value {
    hash() { throw "unimplemented" }
    typeCode() { throw "unimplemented" }
}

class IntValue extends Value {
    // ival: 32 byte integer
    constructor(ival) {
        this.ival = ival;
        this.typeCode = () => TYPE_INT;
    }

    hash() {
        let bytes32 = utils.padZeros(utils.arrayify(this.ival), 32);
        return utils.soliditySha256(['bytes32'], [bytes32]);
    }
}

class CodePointValue extends Value {
    // insnNum: 8 byte integer
    // op: BasicOp or ImmOp
    // nextHash: 32 byte hash
    constructor(insnNum, op, nextHash) {
        this.insnNum = insnNum;
        this.op = op;
        this.nextHash = nextHash;
        this.typeCode = () => TYPE_CODE_POINT;
        this.haltCodePointHash = utils.sha256(utils.toUtf8Bytes("HaltCodePointHash"));
        this.errorCodePointHash = utils.sha256(utils.toUtf8Bytes("ErrorCodePointHash"));
    }

    hash() {
        if (this.insnNum == -1) {
            return this.haltCodePointHash();
        } else if (this.insnNum == -2) {
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
        this.hash = () => hash;
        this.size = size;
        this.typeCode = () => TYPE_HASH;
    }
}

class SizedTupleValue extends Value {
    // contents: array of Value(s)
    // size: num of Value(s) in contents
    constructor(contents, size) {
        if (size < 0 || size > MAX_TUPLE_SIZE) {
            throw ("Error SizedTupleValue: illegal size " + size);
        }
        this.contents = contents;
        this.itemCount = size;
        this.typeCode = () => TYPE_TUPLE_0;
        this.internalTypeCode = () => (TYPE_TUPLE_0 + size);
        let hash = function() {
            let hashesArr = this.contents.reduce((acc, value) => acc.concat(value.hash()), "");
            // TODO: missing the right pad of bytes. Why would we need this if all 32B?
            // TODO: check this hashing function is the same as the go one
            let hashes = '0x'.concat(hashes.replace(/0x/g, ''));
            return utils.soliditySha256(['uint8', 'bytes' + (this.itemCount * 32)],
                                        [this.internalTypeCode(), hashes]);
        }();
        this.hash = () => hash;
    }
}

function unmarshalValue(hexString) {
    let [head, tail] = extractBytes(hexString, 1);

    let ty = parseInt(head, 16);
    if (ty === TYPE_INT) {
        let [head, tail] = extractBytes(tail, 32);
        let i = utils.bigNumberify('0x' + head);
        return [IntValue(i), tail];
    } else if (ty === TYPE_CODE_POINT) {
        let [head, tail] = extractBytes(tail, 8);
        let pc = utils.bigNumberify(flipEndian(head)); // pc encoded as Little Endian
        let [op, tail] = unmarshalOp(tail);
        let [head, tail] = extractBytes(tail, 32);
        let nextHash = '0x' + head;
        return [CodePointValue(pc, op, nextHash), tail];
    } else if (ty === TYPE_HASH) {
        throw "Error unmarshalling: hash only value was not expected";
        let [head, tail] = extractBytes(tail, 8);
        let size = parseInt(flipEndian(head), 16); // size encoded as Little Endian
        let [head, tail] = extractBytes(tail, 32);
        let hash = '0x' + head;
        return [HashOnlyValue(hash, size), tail];
    } else if ((ty - TYPE_TUPLE_0) >= TYPE_TUPLE_0 && ty <= TYPE_TUPLE_MAX) {
        let size = ty - TYPE_TUPLE_0;
        let [contents, tail] = unmarshalTuple(tail, size);
        return [SizedTupleValue(contents, size), tail];
    } else {
        throw ("Error unmarshalling: invalid value type " + ty);
    }
}

function unmarshalOp(hexString) {
    let [head, tail] = extractBytes(hexString, 1);
    let immCount = parseInt(head, 16);
    if (immCount == 0) {
        let [opcode, tail] = unmarshalOpCode(tail);
        return [BasicOp(opcode), tail]
    } else if (immCount == 1) {
        let [opcode, tail] = unmarshalOpCode(tail);
        let [value, tail] = unmarshalValue(tail);
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
    let contents = new Array(size); // TODO: okay not to allocate MAX_TUPLE_SIZE right as go does?
    let value = undefined;
    let tail = hexString;
    for (let i = 0; i < size; i++) {
        [value, tail] = unmarshalValue(tail);
        contents[i] = value;
    }
    return [contents, tail];
}

