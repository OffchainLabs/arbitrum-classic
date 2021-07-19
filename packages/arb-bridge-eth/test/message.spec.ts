/*
 * Copyright 2020, Offchain Labs, Inc.
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

/* eslint-env node, mocha */

// import { ethers } from 'hardhat'
// import { assert, expect } from 'chai'
// import { MessageTester } from '../build/types/MessageTester'
// import {  Message } from 'arb-provider-ethers'

// let messageTester: MessageTester

// const token = '0xc7711f36b2C13E00821fFD9EC54B04A60AEfbd1b'
// const dest = '0xFeCd3992654bFC565c3aFc6C4d7b14dCe603EbF5'
// const sender = '0x7073c616a8A3F277Ea4511fCe9EBB2656a1b87B8'
// const value = ethers.BigNumber.from(563356543)

// describe('Messages', async () => {
//   before(async () => {
//     const MessageTester = await ethers.getContractFactory('MessageTester')
//     messageTester = (await MessageTester.deploy()) as MessageTester
//     await messageTester.deployed()
//   })

//   it('can hash incoming messages as values', async () => {
//     const msg = new Message.EthMessage(dest, value)
//     const inMsg = new Message.IncomingMessage(msg, 1000, 5345346, sender, 65465)

//     const calcMsgHash = await messageTester.messageValueHash(
//       inMsg.msg.kind,
//       inMsg.blockNumber,
//       inMsg.timestamp,
//       inMsg.sender,
//       inMsg.inboxSeqNum,
//       inMsg.msg.asData()
//     )
//     expect(calcMsgHash).to.equal(inMsg.asValue().hash())
//   })

//   it('can hash incoming messages as commitments', async () => {
//     const msg = new Message.EthMessage(dest, value)
//     const inMsg = new Message.IncomingMessage(msg, 1000, 5345346, sender, 65465)

//     const calcMsgHash = await messageTester.messageHash(
//       inMsg.msg.kind,
//       inMsg.sender,
//       inMsg.blockNumber,
//       inMsg.timestamp,
//       inMsg.inboxSeqNum,
//       ethers.utils.keccak256(inMsg.msg.asData())
//     )
//     expect(calcMsgHash).to.equal(inMsg.commitmentHash())
//   })
// })
