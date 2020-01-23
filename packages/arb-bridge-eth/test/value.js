/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

const ValueTester = artifacts.require("ValueTester");
const ArbValue = require("arb-provider-ethers").ArbValue;
const test_cases = require("./test_cases.json");

contract("ArbRollup", accounts => {
  it("should initialize", async () => {
    let value_tester = await ValueTester.new();
    let val = new ArbValue.IntValue(100);
    let res = await value_tester.deserializeHashed(ArbValue.marshal(val), 0);
    assert.isTrue(res["0"], "value didn't deserialize correctly");
    assert.equal(val.hash(), res["2"], "value hashes incorrectly");
  });

  for (let i = 0; i < test_cases.length; i++) {
    it(test_cases[i].name, async () => {
      let value_tester = await ValueTester.new();
      const expectedHash = test_cases[i].hash;
      let res = await value_tester.deserializeHashed(
        "0x" + test_cases[i].value,
        0
      );
      assert.isTrue(res["0"], "value didn't deserialize correctly");
      assert.equal("0x" + expectedHash, res["2"], "value hashes incorrectly");
    });
  }

  it("should parse erc20 message", async () => {
    let value_tester = await ValueTester.new();
    const val = new ArbValue.TupleValue([
      new ArbValue.IntValue(2),
      new ArbValue.IntValue(
        "1454660323771124265538360532739934987166685588469"
      ),
      new ArbValue.TupleValue([
        new ArbValue.IntValue(
          "641988807973089174688456409219473706566398216120"
        ),
        new ArbValue.IntValue(
          "1454660323771124265538360532739934987166685588469"
        ),
        new ArbValue.IntValue(1543)
      ])
    ]);

    const val_data = ArbValue.marshal(val);
    let res = await value_tester.deserializeMessageData(val_data, 0);
    let offset = res["1"].toNumber();
    assert.isTrue(res["0"], "value didn't deserialize correctly");
    assert.equal(res["2"].toNumber(), 2, "Incorrect message type");
    assert.equal(
      res["3"],
      "0xFeCd3992654bFC565c3aFc6C4d7b14dCe603EbF5",
      "Incorrect sender"
    );

    let res2 = await value_tester.getERCTokenMsgData(val_data, offset);
    assert.isTrue(res2["0"], "value didn't deserialize correctly");
    assert.equal(
      res2["2"],
      "0x7073c616a8A3F277Ea4511fCe9EBB2656a1b87B8",
      "Incorrect token contract"
    );
    assert.equal(
      res2["3"],
      "0xFeCd3992654bFC565c3aFc6C4d7b14dCe603EbF5",
      "Incorrect dest"
    );
    assert.equal(res2["4"].toNumber(), 1543, "Incorrect value");
  });
});
