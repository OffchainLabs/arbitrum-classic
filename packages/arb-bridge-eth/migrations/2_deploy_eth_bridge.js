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

var BytesLib = artifacts.require("bytes/BytesLib.sol");

var DebugPrint = artifacts.require("./libraries/DebugPrint.sol");
var SigUtils = artifacts.require("./libraries/SigUtils.sol");
var MerkleLib = artifacts.require("./libraries/MerkleLib.sol");

var Machine = artifacts.require("./arch/Machine.sol");
var OneStepProof = artifacts.require("./arch/OneStepProof.sol");
var Protocol = artifacts.require("./arch/Protocol.sol");
var Value = artifacts.require("./arch/Value.sol");

var Challenge = artifacts.require("./challenge/Challenge.sol");
var ChallengeFactory = artifacts.require("./factories/ChallengeFactory.sol");

var VM = artifacts.require("./VM.sol");
var Disputable = artifacts.require("./Disputable.sol");
var Unanimous = artifacts.require("./Unanimous.sol");
var ArbChain = artifacts.require("./vm/ArbChain.sol");
var ArbChannel = artifacts.require("./vm/ArbChannel.sol");
var ChainFactory = artifacts.require("./factories/ChainFactory.sol");
var ChannelFactory = artifacts.require("./factories/ChannelFactory.sol");

var GlobalPendingInbox = artifacts.require("./GlobalPendingInbox.sol");

module.exports = async function(deployer, network, accounts) {
  deployer.deploy(DebugPrint);
  deployer.link(DebugPrint, []);

  deployer.deploy(MerkleLib);
  deployer.link(MerkleLib, [Challenge]);

  deployer.deploy(SigUtils);
  deployer.link(SigUtils, [GlobalPendingInbox, Unanimous]);

  deployer.deploy(BytesLib);
  deployer.link(BytesLib, []);

  deployer.deploy(Value);
  deployer.link(Value, [
    ArbChain,
    ArbChannel,
    Protocol,
    GlobalPendingInbox,
    OneStepProof
  ]);

  deployer.deploy(Protocol);
  deployer.link(Protocol, [Challenge, Disputable, Unanimous]);

  deployer.deploy(Machine);
  deployer.link(Machine, []);

  deployer.deploy(OneStepProof);
  deployer.link(OneStepProof, [Challenge]);

  deployer.deploy(VM);
  deployer.link(VM, [ArbChannel, Disputable, Unanimous]);

  deployer.deploy(Disputable);
  deployer.link(Disputable, [ArbChain, ArbChannel]);

  deployer.deploy(Unanimous);
  deployer.link(Unanimous, [ArbChannel]);

  await deployer.deploy(Challenge);
  await deployer.deploy(ArbChain);
  await deployer.deploy(ArbChannel);

  await deployer.deploy(GlobalPendingInbox);
  await deployer.deploy(ChallengeFactory, Challenge.address);
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
