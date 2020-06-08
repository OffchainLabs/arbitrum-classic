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

var OneStepProof = artifacts.require('./arch/OneStepProof.sol')

var MessagesChallenge = artifacts.require('./challenge/MessagesChallenge.sol')
var InboxTopChallenge = artifacts.require('./challenge/InboxTopChallenge.sol')
var ExecutionChallenge = artifacts.require('./challenge/ExecutionChallenge.sol')
var ChallengeFactory = artifacts.require('./factories/ChallengeFactory.sol')

var ArbRollup = artifacts.require('./vm/ArbRollup.sol')
var ArbFactory = artifacts.require('./vm/ArbFactory.sol')

var Value = artifacts.require('./arch/Value.sol')

var GlobalInbox = artifacts.require('./GlobalInbox.sol')

module.exports = async function (deployer, network, accounts) {
  deployer.deploy(OneStepProof)
  deployer.link(OneStepProof, [ExecutionChallenge])

  deployer.deploy(Value)
  deployer.link(Value, [GlobalInbox])

  let contractPromises1 = []
  contractPromises1.push(deployer.deploy(MessagesChallenge))
  contractPromises1.push(deployer.deploy(InboxTopChallenge))
  contractPromises1.push(deployer.deploy(ExecutionChallenge))
  contractPromises1.push(deployer.deploy(ArbRollup))
  contractPromises1.push(deployer.deploy(GlobalInbox))

  await Promise.all(contractPromises1)

  await deployer.deploy(
    ChallengeFactory,
    MessagesChallenge.address,
    InboxTopChallenge.address,
    ExecutionChallenge.address
  )

  await deployer.deploy(
    ArbFactory,
    ArbRollup.address,
    GlobalInbox.address,
    ChallengeFactory.address
  )

  const fs = require('fs')
  let addresses = {
    ArbFactory: ArbFactory.address,
  }
  fs.writeFileSync('bridge_eth_addresses.json', JSON.stringify(addresses))
}
