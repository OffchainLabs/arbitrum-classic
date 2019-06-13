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

const ganache = require("ganache-core");
const path = require("path");
const fs = require("fs");
const callsite = require('callsite');
const spawnSync = require("child_process").spawnSync;

function provider(outputFolder, buildLocation, options) {
  var outputLocation = path.resolve(outputFolder, 'compiled.json')
  var outputAOLocation = path.resolve(outputFolder, 'contract.ao')
  var stack = callsite()
  var rootPath = path.dirname(stack[1].getFileName())
  if (!buildLocation) {
    buildLocation = path.resolve(rootPath, 'build/contracts');
  }
  options["network_id"] = 123456789
  let arbProvider = ganache.provider(options)

  var contractCode = {}
  arbProvider.engine.on("block", function(block) {
    for (const [ address_string, value ] of Object.entries(storage)) {
        arbProvider.engine.manager.eth_getCode(address_string, "latest", function(err, code) {
          contractCode[address_string] = code
        })
      }
  })

  var storage = {}
  var netID = arbProvider.options.network_id;
  arbProvider.engine.manager.waitForInitialization(function(err, state) {
    state.blockchain.vm.on("step", function(info) {
      var address_string = '0x' + info.address.toString('hex')
      if (!(address_string in storage)) {
        storage[address_string] = {}
      }
      if (info.opcode.name == "SSTORE") {
        var args = info.stack
          .slice(-info.opcode.in)
          .map((arg) => '0x' + arg.toString('hex'))
        
        storage[address_string][args[1]] = args[0]
      }
    });
  })
  process.on('exit', (code) => {
    var contracts = []
    var files = fs.readdirSync(buildLocation, {})
    for (var filePath of files) {
      var contract = JSON.parse(fs.readFileSync(path.resolve(buildLocation, filePath)))
      var networkInfo = contract.networks[netID]
      if (networkInfo) {
        var address = networkInfo.address
        contracts.push({
          name: contract.contractName,
          address: address,
          code: contractCode[address.toLowerCase()],
          storage: storage[address.toLowerCase()], 
          abi: contract.abi
        })
      }
    }
    try {
      fs.writeFileSync(outputLocation, JSON.stringify(contracts, null, 2))
    } catch (e) {
      console.log("Error writing output to file: " + outputLocation +
        "\n" + e.name + " " + e.message);
      throw e;
    }
    console.log("arbc-truffle-compile compiled.json contract.ao");
    try {
      var compile = spawnSync("arbc-truffle-compile",
        [outputLocation, outputAOLocation], {encoding: 'utf-8'});
      console.log(compile.stdout);
      console.log(compile.stderr);
    } catch (e) {
        console.log("Error arbc-truffle-compile: " + e.name + " " + e.message);
        throw e;
    }
  })
  return arbProvider;
}
module.exports = {
  provider: provider
}
