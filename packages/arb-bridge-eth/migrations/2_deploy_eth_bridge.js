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
var OneStepProof = artifacts.require("./OneStepProof.sol");
var ArbMachine = artifacts.require("./ArbMachine.sol");
var BytesLib = artifacts.require("bytes/BytesLib.sol");
var MerkleLib = artifacts.require("./MerkleLib.sol");
var SigUtils = artifacts.require("./SigUtils.sol");

var ArbChain = artifacts.require("./vm/ArbChain.sol");
var ArbChannel = artifacts.require("./vm/ArbChannel.sol");
var ArbChallenge = artifacts.require("./challenge/ArbChallenge.sol");

var ChallengeFactory = artifacts.require("./ChallengeFactory.sol");
var ChainFactory = artifacts.require("./ChainFactory.sol");
var ChannelFactory = artifacts.require("./ChannelFactory.sol");
var GlobalPendingInbox = artifacts.require("./GlobalPendingInbox.sol");

module.exports = async function(deployer, network, accounts) {
  deployer.deploy(DebugPrint);
  deployer.link(DebugPrint, []);

  deployer.deploy(MerkleLib);
  deployer.link(MerkleLib, [Bisection]);

  deployer.deploy(SigUtils);
  deployer.link(SigUtils, [GlobalPendingInbox, Unanimous]);

  deployer.deploy(BytesLib);
  deployer.link(BytesLib, []);

  deployer.deploy(ArbValue);
  deployer.link(ArbValue, [
    ArbChain,
    ArbChannel,
    ArbProtocol,
    GlobalPendingInbox,
    OneStepProof
  ]);

  deployer.deploy(ArbProtocol);
  deployer.link(ArbProtocol, [Bisection, Disputable, OneStepProof, Unanimous]);

  deployer.deploy(ArbMachine);
  deployer.link(ArbMachine, []);

  deployer.deploy(OneStepProof);
  deployer.link(OneStepProof, [ArbChallenge]);

  deployer.deploy(Bisection);
  deployer.link(Bisection, [ArbChallenge]);

  deployer.deploy(VM);
  deployer.link(VM, [ArbChannel, Disputable, Unanimous]);

  deployer.deploy(Disputable);
  deployer.link(Disputable, [ArbChain, ArbChannel]);

  deployer.deploy(Unanimous);
  deployer.link(Unanimous, [ArbChannel]);

  await deployer.deploy(ArbChallenge);
  await deployer.deploy(ArbChain);
  await deployer.deploy(ArbChannel);

  await deployer.deploy(GlobalPendingInbox);
  await deployer.deploy(ChallengeFactory, ArbChallenge.address);
  await deployer.deploy(
    ChainFactory,
    ArbChain.address,
    GlobalPendingInbox.address,
    ChallengeFactory.address
  );
  await deployer.deploy(
    ChannelFactory,
    ArbChannel.address,
    GlobalPendingInbox.address,
    ChallengeFactory.address
  );

  const fs = require("fs");
  let addresses = {
    ChainFactory: ChainFactory.address,
    ChannelFactory: ChannelFactory.address,
    GlobalPendingInbox: GlobalPendingInbox.address,
    OneStepProof: OneStepProof.address
  };
  fs.writeFileSync("bridge_eth_addresses.json", JSON.stringify(addresses));
};
