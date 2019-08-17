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
"use strict";

const ganache = require("ganache-core");
const path = require("path");
const fs = require("fs");
const callsite = require("callsite");
const spawnSync = require("child_process").spawnSync;

const filenameEVM = "compiled.json";
const filenameAO = "contract.ao";

function provider(outputFolder, buildLocation, options) {
  const outputLocationEVM = path.resolve(outputFolder, filenameEVM);
  const outputLocationAO = path.resolve(outputFolder, filenameAO);
  const stack = callsite();
  const rootPath = path.dirname(stack[1].getFileName());
  if (!buildLocation) {
    buildLocation = path.resolve(rootPath, "build/contracts");
  }
  options.network_id = 123456789;
  options.allowUnlimitedContractSize = true;
  const arbProvider = ganache.provider(options);

  const contractCode = {};

  let storageTrackFuncGen = function(address_string) {
    return function(err, code) {
      contractCode[address_string] = code;
    };
  };

  arbProvider.engine.on("block", function(block) {
    for (let [address_string, value] of Object.entries(storage)) {
      arbProvider.engine.manager.eth_getCode(
        address_string,
        "latest",
        storageTrackFuncGen(address_string)
      );
    }
  });

  const storage = {};
  const netID = arbProvider.options.network_id;
  arbProvider.engine.manager.waitForInitialization(function(err, state) {
    state.blockchain.vm.on("step", function(info) {
      let address_string = "0x" + info.address.toString("hex");
      if (!(address_string in storage)) {
        storage[address_string] = {};
      }
      if (info.opcode.name == "SSTORE") {
        let args = info.stack.slice(-2).map(arg => "0x" + arg.toString("hex"));
        storage[address_string][args[1]] = args[0];
      }
    });
  });
  process.on("exit", code => {
    const contracts = [];
    const files = fs.readdirSync(buildLocation, {});
    for (let filePath of files) {
      const contract = JSON.parse(
        fs.readFileSync(path.resolve(buildLocation, filePath))
      );
      const networkInfo = contract.networks[netID];
      if (networkInfo) {
        const address = networkInfo.address;
        contracts.push({
          name: contract.contractName,
          address: address,
          code: contractCode[address.toLowerCase()],
          storage: storage[address.toLowerCase()],
          abi: contract.abi
        });
      }
    }
    try {
      fs.writeFileSync(outputLocationEVM, JSON.stringify(contracts, null, 2));
    } catch (e) {
      console.log(
        "Error writing output to file: " +
          outputLocationEVM +
          "\n" +
          e.name +
          " " +
          e.message
      );
      throw e;
    }
    console.log("arbc-truffle " + filenameEVM + " " + filenameAO);
    try {
      var compile = spawnSync(
        "arbc-truffle",
        [outputLocationEVM, outputLocationAO],
        { encoding: "utf-8" }
      );
      console.log(compile.stdout);
      console.log(compile.stderr);
    } catch (e) {
      console.log("Error arbc-truffle: " + e.name + " " + e.message);
      throw e;
    }
  });
  return arbProvider;
}

module.exports = {
  provider
};
