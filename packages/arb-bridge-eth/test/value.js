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
});
