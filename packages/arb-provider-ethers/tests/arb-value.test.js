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
"use strict";

const ethers = require("ethers");
const utils = ethers.utils;
const arb = require("../lib/arb-value");
const test_cases = require("./test_cases.json");

// Helper shortcuts
const bn = n => utils.bigNumberify(n);
const ZEROS_16B = "00000000000000000000000000000000";
const ZEROS_32B = ZEROS_16B + ZEROS_16B;
const ONES_16B = "ffffffffffffffffffffffffffffffff";
const ONES_32B = ONES_16B + ONES_16B;
const EMPTY_TUPLE_HASH =
  "0x69c322e3248a5dfc29d73c5b0553b0185a35cd5bb6386747517ef7e53b15e287";

describe("Helper Functions", function() {
  test("extractBytes", function() {
    expect(arb.extractBytes("", 0)).toEqual(["", ""]);
    expect(arb.extractBytes("hello, world", 1)).toEqual(["he", "llo, world"]);
    expect(arb.extractBytes("1234567890abcdef", 0)).toEqual([
      "",
      "1234567890abcdef"
    ]);
    expect(arb.extractBytes("1234567890abcdef", 1)).toEqual([
      "12",
      "34567890abcdef"
    ]);
    expect(arb.extractBytes("1234567890abcdef", 4)).toEqual([
      "12345678",
      "90abcdef"
    ]);
    expect(arb.extractBytes("1234567890abcdef", 8)).toEqual([
      "1234567890abcdef",
      ""
    ]);
    let e = "Error extracting bytes: string is too short";
    expect(() => arb.extractBytes("hello, world", -1)).toThrow(e);
    expect(() => arb.extractBytes("hello, world", 7)).toThrow(e);
  });

  test("intToBytes", function() {
    expect(arb.intToBytes(0, 0)).toBe("0");
    expect(arb.intToBytes(0, 1)).toBe("00");
    expect(arb.intToBytes(27, 1)).toBe("1b");
    expect(arb.intToBytes(365, 2)).toBe("016d");
    expect(arb.intToBytes(2048, 2)).toBe("0800");
    expect(arb.intToBytes(2048, 32)).toBe(
      ZEROS_32B.slice(0, 32 * 2 - 4) + "0800"
    );
  });

  test("uBigNumToBytes", function() {
    expect(arb.uBigNumToBytes(utils.bigNumberify(0))).toBe(ZEROS_32B);
    expect(arb.uBigNumToBytes(utils.bigNumberify(2048))).toBe(
      ZEROS_32B.slice(0, 32 * 2 - 4) + "0800"
    );
    expect(arb.uBigNumToBytes(utils.bigNumberify("0x" + ONES_16B))).toBe(
      ZEROS_16B + ONES_16B
    );
    expect(arb.uBigNumToBytes(utils.bigNumberify("0x" + ONES_32B))).toBe(
      ONES_32B
    );
  });
});

describe("Constructors", function() {
  const nullHash = "0x" + ZEROS_32B;

  test("BasicOp", function() {
    let bop = new arb.BasicOp(arb.OP_CODE_HALT);
    expect(bop.opcode).toBe(arb.OP_CODE_HALT);
    expect(bop.immCount).toBe(0);
  });

  test("ImmOp", function() {
    let iop = new arb.ImmOp(0x19, new arb.IntValue(utils.bigNumberify(9)));
    expect(iop.opcode).toBe(0x19);
    expect(iop.immCount).toBe(1);
    expect(iop.val.bignum.toNumber()).toBe(9);
  });

  test("Value (abstract", function() {
    let v = new arb.Value();
    expect(() => v.hash()).toThrow("unimplemented");
    expect(() => v.typeCode()).toThrow("unimplemented");
  });

  test("IntValue", function() {
    let iv = new arb.IntValue(utils.bigNumberify(0));
    expect(iv.bignum.toNumber()).toBe(0);
    expect(iv.typeCode()).toBe(0);
  });

  test("CodePointValue", function() {
    let cpv = new arb.CodePointValue(0, new arb.BasicOp(0x60), nullHash);
    expect(cpv.insnNum).toBe(0);
    expect(cpv.op.opcode).toBe(0x60);
    expect(cpv.nextHash).toBe(nullHash);

    // Test BasicOp hash value
    let bopv = new arb.CodePointValue(
      99,
      new arb.BasicOp(0x60),
      EMPTY_TUPLE_HASH
    );
    let preCalc =
      "0xb9cffde57db229ede25012536613c9a9f9a7cde0e5f23381350737b6818852da";
    expect(bopv.hash()).toBe(preCalc);

    // Test ImmOp hash value
    let immop = new arb.ImmOp(0x60, new arb.IntValue(bn(0)));
    let immv = new arb.CodePointValue(100, immop, EMPTY_TUPLE_HASH);
    let preCalc2 =
      "0x9b6304f1c0d7299152b70c5097dcb370ed46668b1f09146586a333f507594619";
    expect(immv.hash()).toBe(preCalc2);

    // Test invalid Operand type
    let invalid = new arb.CodePointValue(0, new arb.Operation(0x0), nullHash);
    expect(() => invalid.hash()).toThrow(
      "CodePointValue must be instanceof BasicOp or ImmOp"
    );
    expect(() => invalid.toString()).toThrow(
      "CodePointValue must be instanceof BasicOp or ImmOp"
    );
  });

  test("HashOnlyValue", function() {
    let hov = new arb.HashOnlyValue(nullHash, 0);
    expect(hov.hash()).toBe(nullHash);
    expect(hov.size).toBe(0);
    expect(hov.typeCode()).toBe(2);
  });
});

describe("TupleValue", function() {
  test("Empty", function() {
    let emptyTuple = new arb.TupleValue([]);
    expect(emptyTuple.contents).toEqual([]);
    expect(emptyTuple.typeCode()).toBe(3);
    expect(emptyTuple.hash()).toBe(EMPTY_TUPLE_HASH);
  });

  test("Two Tuple", function() {
    let twoTuple = new arb.TupleValue([
      new arb.IntValue(utils.bigNumberify(0)),
      new arb.IntValue(utils.bigNumberify(100))
    ]);
    expect(twoTuple.get(0).bignum.isZero()).toBe(true);
    expect(twoTuple.get(1).bignum.eq(utils.bigNumberify(100))).toBe(true);
    expect(twoTuple.typeCode()).toBe(3 + 2);
  });

  test("Largest Tuple", function() {
    let mtsv = new arb.TupleValue(
      Array(arb.MAX_TUPLE_SIZE).fill(new arb.TupleValue([]))
    );
    expect(mtsv.contents.length).toBe(arb.MAX_TUPLE_SIZE);
    for (let i = 0; i < arb.MAX_TUPLE_SIZE; i++) {
      expect(mtsv.get(i).contents.length).toBe(0);
    }
    expect(mtsv.typeCode()).toBe(3 + arb.MAX_TUPLE_SIZE);
    // Pre calculated hash
    let p =
      "0xbfa3a3ab6c2f6c71c78354b8186a0e206a99b73b9bb6d8e6d45b733466decbf8";
    expect(mtsv.hash()).toBe(p);
  });

  test("Greater than arb.MAX_TUPLE_SIZE", function() {
    expect(() => new arb.TupleValue(Array(arb.MAX_TUPLE_SIZE + 1))).toThrow(
      "Error TupleValue: illegal size " + (arb.MAX_TUPLE_SIZE + 1)
    );
  });

  test("get and set", function() {
    let emptyTuple = new arb.TupleValue(
      Array(arb.MAX_TUPLE_SIZE).fill(new arb.TupleValue([]))
    );
    let t = emptyTuple;
    // set
    for (let i = 0; i < arb.MAX_TUPLE_SIZE; i++) {
      t = t.set(i, new arb.IntValue(utils.bigNumberify(i)));
      expect(t.contents[i].bignum.toNumber()).toBe(i);
    }

    expect(() => t.set(arb.MAX_TUPLE_SIZE, new arb.TupleValue([]))).toThrow(
      "Error TupleValue set: index out of bounds " + arb.MAX_TUPLE_SIZE
    );
    expect(() => t.set(-1, new arb.TupleValue([]))).toThrow(
      "Error TupleValue set: index out of bounds " + -1
    );

    // get
    for (let i = 0; i < arb.MAX_TUPLE_SIZE; i++) {
      expect(t.get(i).bignum.toNumber()).toBe(i);
    }

    expect(() => t.get(arb.MAX_TUPLE_SIZE)).toThrow(
      "Error TupleValue get: index out of bounds " + arb.MAX_TUPLE_SIZE
    );
    expect(() => t.get(-1)).toThrow(
      "Error TupleValue get: index out of bounds " + -1
    );
  });
});

describe("BigTuple", function() {
  test("getBigTuple and setBigTuple", function() {
    let emptyBigTup = new arb.TupleValue([]);
    expect(
      arb.getBigTuple(emptyBigTup, utils.bigNumberify(93)).bignum.toNumber()
    ).toBe(0);
    expect(
      arb.getBigTuple(emptyBigTup, utils.bigNumberify(1234567890)).bignum.eq(0)
    ).toBe(true);

    let t = emptyBigTup;
    for (let i = bn(0); i.lt(100); i = i.add(1)) {
      t = arb.setBigTuple(t, i, new arb.IntValue(i));
      expect(arb.getBigTuple(t, i).bignum.toNumber()).toBe(i.toNumber());
    }
  });
});

// Marshaled sizes as hexstrings
const M_INT_VALUE_SIZE = (1 + 32) * 2;
const M_CODE_POINT_SIZE = (1 + 8 + 1 + 1 + 0 + 32) * 2; // Without val
const M_HASH_ONLY_SIZE = (1 + 8 + 32) * 2;
const M_TUPLE_SIZE = (1 + 0) * 2; // Without other vals

describe("Marshaling", function() {
  test("marshal and unmarshal IntValue", function() {
    for (let i of [0, 1, 100, "0x9271342394932492394"]) {
      let iv = new arb.IntValue(bn(i));
      let marshaledBytes = arb.marshal(iv);
      expect(marshaledBytes.slice(2).length).toBe(M_INT_VALUE_SIZE);
      let unmarshaledValue = arb.unmarshal(marshaledBytes);
      expect(unmarshaledValue.bignum.eq(bn(i))).toBe(true);
    }

    // Test that negative IntValues throw on marshal
    expect(() => arb.marshal(new arb.IntValue(bn(-1)))).toThrow(
      "Error marshaling IntValue: negative values not supported"
    );

    // Test without "0x"
    let iv = new arb.IntValue(bn(99));
    let marshaledBytes = arb.marshal(iv).slice(2);
    expect(marshaledBytes.length).toBe(M_INT_VALUE_SIZE);
    let unmarshaledValue = arb.unmarshal(marshaledBytes);
    expect(unmarshaledValue.bignum.eq(bn(99))).toBe(true);
  });

  test("marshal and unmarshal CodePointValue", function() {
    let [pc, op, nextHash] = [
      bn(0),
      new arb.BasicOp(arb.OP_CODE_HALT),
      "0x" + ZEROS_32B
    ];
    let basic_tcv = new arb.CodePointValue(pc, op, nextHash);
    let marshaledBytes = arb.marshal(basic_tcv);
    expect(marshaledBytes.slice(2).length).toBe((1 + 8 + 1 + 1 + 0 + 32) * 2);
    let revValue = arb.unmarshal(marshaledBytes);
    expect(revValue.insnNum).toEqual(pc);
    expect(revValue.op.opcode).toBe(op.opcode);
    expect(revValue.nextHash).toEqual(nextHash);
    expect(revValue.toString()).toEqual(basic_tcv.toString());

    let iv = new arb.IntValue(bn(60));
    expect(arb.marshal(iv).slice(2).length).toBe(M_INT_VALUE_SIZE);
    let imm_tcv = new arb.CodePointValue(pc, new arb.ImmOp(0x19, iv), nextHash);
    let mb = arb.marshal(imm_tcv);
    expect(mb.slice(2).length).toBe(M_CODE_POINT_SIZE + M_INT_VALUE_SIZE);
    let revImmValue = arb.unmarshal(mb);
    expect(revImmValue.insnNum).toEqual(pc);
    expect(revImmValue.op.opcode).toBe(0x19);
    expect(revImmValue.op.val.bignum.toNumber()).toBe(60);
    expect(revImmValue.nextHash).toEqual(nextHash);
    expect(revImmValue.toString()).toEqual(imm_tcv.toString());
  });

  test("marshal and unmarshal HashOnlyValue", function() {
    // HashOnlyValue should not be used
    let hv = new arb.HashOnlyValue("0x" + ZEROS_32B, 0);
    let marshaledBytes = arb.marshal(hv);
    expect(marshaledBytes.slice(2).length).toBe((1 + 8 + 32) * 2);
    expect(() => arb.unmarshal(marshaledBytes)).toThrow(
      "Error unmarshaling: HashOnlyValue was not expected"
    );
    expect(hv.toString()).toEqual("HashOnlyValue(" + hv.hash() + ")");
  });

  test("marshal and unmarshal TupleValue", function() {
    // Empty Tuple
    let etv = new arb.TupleValue([]);
    let etvm = arb.marshal(etv);
    expect(etvm.slice(2).length).toBe(M_TUPLE_SIZE);
    let etv_rev = arb.unmarshal(etvm);
    expect(etv_rev.toString()).toEqual(etv.toString());

    // Full Tuple of Empty Tuple"s
    let ftv = new arb.TupleValue(Array(8).fill(new arb.TupleValue([])));
    let ftvm = arb.marshal(ftv);
    expect(ftvm.slice(2).length).toBe(M_TUPLE_SIZE + M_TUPLE_SIZE * 8);
    let ftv_rev = arb.unmarshal(ftvm);
    expect(ftv_rev.toString()).toEqual(ftv.toString());

    // Full Tuple of IntValue"s
    let fitv = new arb.TupleValue(Array(8).fill(new arb.IntValue(bn(0))));
    let fitvm = arb.marshal(fitv);
    expect(fitvm.slice(2).length).toBe(M_TUPLE_SIZE + M_INT_VALUE_SIZE * 8);
    let fitv_rev = arb.unmarshal(fitvm);
    expect(fitv_rev.toString()).toEqual(fitv.toString());
    expect(fitv_rev.get(0).bignum.toNumber()).toBe(0);
    expect(fitv_rev.get(7).bignum.toNumber()).toBe(0);
    expect(() => fitv_rev.get(8)).toThrow(
      "Error TupleValue get: index out of bounds 8"
    );
  });

  test("illegal inputs", function() {
    // Illegal Value
    let erroneous = new arb.Value();
    erroneous.typeCode = () => -1;
    expect(() => arb.marshal(erroneous)).toThrow(
      "Error marshaling value no such TYPE: -1"
    );
    expect(() => arb.unmarshal("0x99")).toThrow(
      "Error unmarshaling value no such TYPE: 99"
    );

    // Illegal Operation immCount
    let cpv = new arb.CodePointValue(
      bn(0),
      new arb.Operation(0),
      "0x" + ZEROS_32B
    );
    cpv.op.immCount = 99;
    expect(() => arb.marshal(cpv)).toThrow(
      "Error marshaling CodePointValue illegal immCount: 99"
    );
    let [ty_code_point, pc, erroneous_op_ty] = [
      "01",
      Array(8)
        .fill("00")
        .join(""),
      "FF"
    ];
    expect(() => arb.unmarshal(ty_code_point + pc + erroneous_op_ty)).toThrow(
      "Error unmarshalOp no such immCount: 0xff"
    );

    // Illegal OpCode
    let ILLEGAL_OP_CODE = "FF",
      immop;
    [ty_code_point, pc, immop] = [
      "01",
      Array(8)
        .fill("00")
        .join(""),
      "00"
    ];
    expect(() =>
      arb.unmarshal(ty_code_point + pc + immop + ILLEGAL_OP_CODE)
    ).toThrow("Error unmarshalOpCode no such opcode: 0xff");
  });
});

describe("Integration", function() {
  test("sizedByteRangeToHex and hexToSizedByteRange", function() {
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
        new arb.IntValue(
          bn(
            "0x5ce0c8f1e004fe36aa260ecd02c68ca0c6dea5a4acdfe0b8b10d7b526360046b"
          )
        )
      ]),
      new arb.TupleValue([
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.IntValue(
          bn(
            "0x781371cb80a394c637cebf3e3d48a268a44ad21cd68239afb3c3a37196d582c1"
          )
        )
      ]),
      new arb.TupleValue([
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.TupleValue([]),
        new arb.IntValue(
          bn(
            "0x32edc9a100000000000000000000000000000000000000000000000000000000"
          )
        )
      ]),
      new arb.TupleValue([]),
      new arb.TupleValue([]),
      new arb.TupleValue([]),
      new arb.IntValue(
        bn("0x90e130e5da79003b67479a3ed2caf5585e93ae6771de6cdec6d7641bd2e60180")
      )
    ]);
    const marshaledBytes = arb.marshal(myValue);
    const expectedMessageBytes =
      "0x0b030b03030303030303005ce0c8f1e004fe36aa260ecd02c68ca0c6dea5a4acdfe0b8b10d7b526360046b0b0303030303030300781371cb80a394c637cebf3e3d48a268a44ad21cd68239afb3c3a37196d582c10b030303030303030032edc9a1000000000000000000000000000000000000000000000000000000000303030090e130e5da79003b67479a3ed2caf5585e93ae6771de6cdec6d7641bd2e60180";
    expect(marshaledBytes).toBe(expectedMessageBytes);
    const val = arb.unmarshal(expectedMessageBytes);
    expect(val.toString()).toEqual(
      "Tuple([Tuple([]), Tuple([Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), 42009942682379378059947058450083587892049528549641310042571988458584210932843]), Tuple([Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), 54311897307976383387700091809425625413681576808311458615378223571158439396033]), Tuple([Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), Tuple([]), 23035776775082914819351730844503175383119388692051102633761137171302117801984]), Tuple([]), Tuple([]), Tuple([]), 65530928266225785593959077233184075030766656784123302077071474652886342173056])"
    );
    const sizedByteRange = new arb.TupleValue([val, new arb.IntValue(bn(100))]);
    const hex = arb.sizedByteRangeToHex(sizedByteRange);
    expect(hex).toBe(
      "0x90e130e5da79003b67479a3ed2caf5585e93ae6771de6cdec6d7641bd2e601805ce0c8f1e004fe36aa260ecd02c68ca0c6dea5a4acdfe0b8b10d7b526360046b781371cb80a394c637cebf3e3d48a268a44ad21cd68239afb3c3a37196d582c132edc9a1"
    );
    const sizedByteRangeReverse = arb.hexToSizedByteRange(hex);
    const sizeReverse = sizedByteRangeReverse.get(1);
    expect(sizeReverse.bignum.toNumber()).toBe(100);
    const valReverse = sizedByteRangeReverse.get(0);
    const messageReverse = arb.marshal(valReverse);
    expect(messageReverse).toBe(expectedMessageBytes);

    // hexToSizedByteRange without "0x" starting
    const sizedByteRangeReverse2 = arb.hexToSizedByteRange(hex.slice(2));
    const sizeReverse2 = sizedByteRangeReverse2.get(1);
    expect(sizeReverse2.bignum.toNumber()).toBe(100);
    const valReverse2 = sizedByteRangeReverse2.get(0);
    const messageReverse2 = arb.marshal(valReverse2);
    expect(messageReverse2).toBe(expectedMessageBytes);
  });
});

describe("test_cases.json", function() {
  for (let i = 0; i < test_cases.length; i++) {
    it(test_cases[i].name, function() {
      let expectedHash = test_cases[i].hash;
      let value = arb.unmarshal(test_cases[i].value);
      let hash = value.hash().slice(2);
      if (hash !== expectedHash) {
        console.log(value.toString());
      }
      expect(hash).toEqual(expectedHash);
    });
  }
});
