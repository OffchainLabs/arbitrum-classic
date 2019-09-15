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

var DebugPrint = artifacts.require("./DebugPrint.sol");

var ArbProtocol = artifacts.require("./ArbProtocol.sol");
var VM = artifacts.require("./VM.sol");
var ArbValue = artifacts.require("./ArbValue.sol");
var Disputable = artifacts.require("./Disputable.sol");
var Unanimous = artifacts.require("./Unanimous.sol");
var Bisection = artifacts.require("./Bisection.sol");
var VMCreator = artifacts.require("./VMCreator.sol");
var OneStepProof = artifacts.require("./OneStepProof.sol");
var ArbMachine = artifacts.require("./ArbMachine.sol");
var BytesLib = artifacts.require("bytes/BytesLib.sol");
var MerkleLib = artifacts.require("./MerkleLib.sol");
var SigUtils = artifacts.require("./SigUtils.sol");
var ChallengeManager = artifacts.require("./ChallengeManager.sol");
var GlobalPendingInbox = artifacts.require("./GlobalPendingInbox.sol");

module.exports = async function(deployer, network, accounts) {
  deployer.deploy(DebugPrint);
  deployer.link(DebugPrint, [ArbMachine, Unanimous]);

  deployer.deploy(MerkleLib);
  deployer.link(MerkleLib, [Bisection]);

  deployer.deploy(SigUtils);
  deployer.link(SigUtils, [VMCreator, Unanimous, GlobalPendingInbox]);

  deployer.deploy(BytesLib);
  deployer.link(BytesLib, [ArbValue]);

  deployer.deploy(ArbValue);
  deployer.link(ArbValue, [
    GlobalPendingInbox,
    OneStepProof,
    ArbMachine,
    ArbProtocol,
    VM,
    VMCreator
  ]);

  deployer.deploy(ArbProtocol);
  deployer.link(ArbProtocol, [
    GlobalPendingInbox,
    VMCreator,
    ChallengeManager,
    OneStepProof,
    VM,
    Disputable,
    Bisection
  ]);

  deployer.deploy(ArbMachine);
  deployer.link(ArbMachine, [OneStepProof]);

  deployer.deploy(OneStepProof);
  deployer.link(OneStepProof, ChallengeManager);

  deployer.deploy(Bisection);
  deployer.link(Bisection, ChallengeManager);

  deployer.deploy(VM);
  deployer.link(VM, [Disputable, Unanimous, VMCreator]);

  deployer.deploy(Disputable);
  deployer.link(Disputable, VMCreator);

  deployer.deploy(Unanimous);
  deployer.link(Unanimous, VMCreator);

  await deployer.deploy(GlobalPendingInbox);
  await deployer.deploy(ChallengeManager);
  await deployer.deploy(
    VMCreator,
    GlobalPendingInbox.address,
    ChallengeManager.address
  );

  const fs = require("fs");
  let addresses = {
    vmCreator: VMCreator.address,
    globalPendingInbox: GlobalPendingInbox.address
  };
  fs.writeFileSync("bridge_eth_addresses.json", JSON.stringify(addresses));
};
