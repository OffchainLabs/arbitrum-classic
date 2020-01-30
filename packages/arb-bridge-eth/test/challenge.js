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

const ChallengeTester = artifacts.require("ChallengeTester");
const ChallengeFactory = artifacts.require("ChallengeFactory");
const Challenge = artifacts.require("Challenge");

contract("ChallengeTester", accounts => {
  it("should initialize", async () => {
    const challenge_factory = await ChallengeFactory.deployed();
    const challenge_tester = await ChallengeTester.new(
      challenge_factory.address
    );
    const pendingTopChallengeTemplate = await challenge_factory.pendingTopChallengeTemplate();
    const codePart1 = "3d602d80600a3d3981f3363d3d373d3d3d363d73";
    const codePart2 = "5af43d82803e903d91602b57fd5bf3";
    const code = codePart1 + pendingTopChallengeTemplate.slice(2) + codePart2;
    const codehash = web3.utils.soliditySha3({ type: "bytes", value: code });
    const nonce1 = web3.utils.soliditySha3(
      { type: "address", value: accounts[0] },
      { type: "address", value: accounts[1] },
      { type: "address", value: challenge_tester.address }
    );
    const calcAddress = web3.utils.soliditySha3(
      { type: "bytes1", value: "0xff" },
      { type: "address", value: challenge_factory.address },
      { type: "uint256", value: nonce1 },
      { type: "bytes32", value: codehash }
    );

    let createTx = await challenge_tester.startChallenge(
      accounts[0],
      accounts[1],
      0,
      "0x",
      0
    );
    const challengeTemplate = createTx.logs[0].args.challengeTemplate;
    const nonce2 = createTx.logs[0].args.nonce;
    const codeHash2 = createTx.logs[0].args.codeHash;
    const cloneAddress = createTx.logs[0].args.cloneAddress;
    // const code2 = createTx.logs[0].args.code2;
    let challengeAddress = createTx.receipt.rawLogs[0].address;
    console.log("challengeAddress", challengeAddress);
    console.log("cloneAddress", cloneAddress.toLowerCase());
    // console.log("code2", code2);
    // let challenge = await Challenge.at(challengeAddress);
    // await challenge.timeoutChallenge();
  });
});
