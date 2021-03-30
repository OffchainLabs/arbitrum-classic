// /*
//  * Copyright 2020, Offchain Labs, Inc.
//  *
//  * Licensed under the Apache License, Version 2.0 (the "License");
//  * you may not use this file except in compliance with the License.
//  * You may obtain a copy of the License at
//  *
//  *    http://www.apache.org/licenses/LICENSE-2.0
//  *
//  * Unless required by applicable law or agreed to in writing, software
//  * distributed under the License is distributed on an "AS IS" BASIS,
//  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  * See the License for the specific language governing permissions and
//  * limitations under the License.
//  */

// /* eslint-env node, mocha */

// import { ethers, waffle } from '@nomiclabs/buidler'
// import * as bre from '@nomiclabs/buidler'
// import { utils, Signer } from 'ethers'
// import * as chai from 'chai'
// import chaiAsPromised from 'chai-as-promised'
// import { deployMockContract } from '@ethereum-waffle/mock-contract'
// import { GlobalInbox } from '../build/types/GlobalInbox'
// import { EthBuddyErc20 } from '../build/types/EthBuddyErc20'
// import { ArbValue, Message } from 'arb-provider-ethers'
// import { initializeAccounts } from './utils'

// chai.use(chaiAsPromised)

// const { expect } = chai

// function getEthMessageData(
//   sender: string,
//   receiver: string,
//   value: utils.BigNumberish
// ): Uint8Array {
//   const ethMsg = new Message.EthMessage(receiver, value)
//   const msg = new Message.OutgoingMessage(ethMsg, sender)
//   return ArbValue.marshal(msg.asValue())
// }

// function getERC20MessageData(
//   sender: string,
//   receiver: string,
//   token: string,
//   value: utils.BigNumberish
// ): Uint8Array {
//   const erc20Msg = new Message.ERC20Message(token, receiver, value)
//   const msg = new Message.OutgoingMessage(erc20Msg, sender)
//   return ArbValue.marshal(msg.asValue())
// }

// function getBuddyMessageData(sender: string, valid: boolean): Uint8Array {
//   const buddyMsg = new Message.BuddyRegisteredMessage(valid)
//   const msg = new Message.OutgoingMessage(buddyMsg, sender)
//   return ArbValue.marshal(msg.asValue())
// }

// const chainAddress = ethers.utils.getAddress(
//   '0xffffffffffffffffffffffffffffffffffffffff'
// )

// let accounts: Signer[]
// let globalInbox: GlobalInbox

// const messageIndex = 0
// let originalOwner: string
// let address2: string
// let address3: string
// let address4: string

// describe('GlobalInbox', () => {
//   before(async () => {
//     accounts = await initializeAccounts()

//     const GlobalInbox = await ethers.getContractFactory('GlobalInbox')
//     globalInbox = (await GlobalInbox.deploy()) as GlobalInbox
//     await globalInbox.deployed()

//     originalOwner = await accounts[0].getAddress()
//     address2 = await accounts[1].getAddress()
//     address3 = await accounts[2].getAddress()
//     address4 = await accounts[3].getAddress()
//   })

//   it('should deposit eth', async () => {
//     await expect(
//       globalInbox.getEthBalance(chainAddress),
//       'Eth balance should start at 0'
//     ).to.eventually.equal(0)

//     await globalInbox.depositEthMessage(
//       chainAddress,
//       await accounts[0].getAddress(),
//       {
//         value: 1000,
//       }
//     )

//     await expect(
//       globalInbox.getEthBalance(chainAddress),
//       "Eth balance wasn't deposited successfully"
//     ).to.eventually.equal(1000)
//   })

//   it('should deposit an ERC20', async () => {
//     const EthBuddyErc20Contract = await ethers.getContractFactory(
//       'EthBuddyERC20'
//     )
//     const erc20 = (await EthBuddyErc20Contract.deploy(
//       globalInbox.address
//     )) as EthBuddyErc20
//     await erc20.deployed()

//     await expect(
//       globalInbox.depositERC20Message(
//         chainAddress,
//         erc20.address,
//         chainAddress,
//         50
//       ),
//       'Must approve before depositing'
//     ).to.be.revertedWith('ERC20: transfer amount exceeds allowance')

//     await erc20.approve(globalInbox.address, 1000000)

//     await expect(
//       globalInbox.getERC20Balance(erc20.address, chainAddress),
//       'ERC20 balance should start at 0'
//     ).to.eventually.equal(0)

//     await globalInbox.depositERC20Message(
//       chainAddress,
//       erc20.address,
//       chainAddress,
//       50
//     )

//     await expect(
//       globalInbox.getERC20Balance(erc20.address, chainAddress),
//       "ERC20 Balance wasn't deposited successfully"
//     ).to.eventually.equal(50)
//   })

//   it('should support paired ERC20s', async () => {
//     const chainAddress = await accounts[6].getAddress()
//     const EthBuddyErc20Contract = await ethers.getContractFactory(
//       'EthBuddyERC20'
//     )
//     const erc20 = (await EthBuddyErc20Contract.deploy(
//       globalInbox.address
//     )) as EthBuddyErc20
//     await erc20.deployed()

//     await expect(
//       globalInbox.isPairedContract(erc20.address, chainAddress),
//       "shouldn't be paired"
//     ).to.eventually.to.equal(0)

//     await erc20.connectToChain(chainAddress)

//     await expect(
//       globalInbox.isPairedContract(erc20.address, chainAddress),
//       'should be initializing'
//     ).to.eventually.to.equal(1)

//     await globalInbox
//       .connect(accounts[6])
//       .sendMessages(await getBuddyMessageData(erc20.address, false), 0, 1)

//     await expect(
//       globalInbox.isPairedContract(erc20.address, chainAddress),
//       'should be reset to unpaired'
//     ).to.eventually.to.equal(0)

//     await erc20.connectToChain(chainAddress)

//     await globalInbox
//       .connect(accounts[6])
//       .sendMessages(await getBuddyMessageData(erc20.address, true), 0, 1)

//     await expect(
//       globalInbox.isPairedContract(erc20.address, chainAddress),
//       'should be paired'
//     ).to.eventually.equal(2)

//     await expect(
//       erc20.balanceOf(address3),
//       'ERC20 balance of dest should start at 0'
//     ).to.eventually.equal(0)

//     await expect(
//       globalInbox.getERC20Balance(erc20.address, chainAddress),
//       'Global inbox ERC20 balance should start at 0'
//     ).to.eventually.equal(0)

//     await globalInbox.depositERC20Message(
//       chainAddress,
//       erc20.address,
//       address4,
//       50
//     )

//     await expect(
//       erc20.balanceOf(globalInbox.address),
//       'Depositing paired token should burn'
//     ).to.eventually.equal(0)

//     await expect(
//       globalInbox.getERC20Balance(erc20.address, chainAddress),
//       'Global inbox ERC20 balance should still be 0'
//     ).to.eventually.equal(0)

//     const erc20MsgData = getERC20MessageData(
//       address2,
//       address3,
//       erc20.address,
//       100000
//     )
//     await globalInbox.connect(accounts[6]).sendMessages(erc20MsgData, 1, 2)

//     await expect(
//       globalInbox.getERC20Balance(erc20.address, address3),
//       'ERC20 balance should increase'
//     ).to.eventually.equal(100000)

//     await globalInbox.connect(accounts[2]).withdrawERC20(erc20.address)

//     await expect(
//       erc20.balanceOf(address3),
//       'Withdrawing should mint tokens'
//     ).to.eventually.equal(100000)
//   })

//   // These tests use a waffle mock which depends on buidlerevm
//   if (bre.network.name == 'buidlerevm') {
//     it('should reject a failed ERC20 deposit', async () => {
//       const [mockCreator] = await waffle.provider.getWallets()
//       const IERC20 = await ethers.getContractFactory('IERC20')
//       const mockERC20 = await deployMockContract(
//         mockCreator,
//         IERC20.interface.abi
//       )

//       await mockERC20.mock.transferFrom.returns(0)
//       await expect(
//         globalInbox.depositERC20Message(
//           chainAddress,
//           mockERC20.address,
//           chainAddress,
//           50
//         )
//       ).to.eventually.be.rejectedWith('FAILED_TRANSFER')
//     })

//     it('should deposit an ERC721', async () => {
//       const [mockCreator] = await waffle.provider.getWallets()
//       const IERC721 = await ethers.getContractFactory('IERC721')
//       const mockERC721 = await deployMockContract(
//         mockCreator,
//         IERC721.interface.abi
//       )
//       await expect(
//         globalInbox.hasERC721(mockERC721.address, chainAddress, 1234),
//         'ERC721 Balance should start 0'
//       ).to.eventually.be.false

//       await mockERC721.mock.transferFrom.returns()
//       await globalInbox.depositERC721Message(
//         chainAddress,
//         mockERC721.address,
//         chainAddress,
//         1234
//       )

//       await expect(
//         globalInbox.hasERC721(mockERC721.address, chainAddress, 1234),
//         "ERC721 Balance wasn't deposited successfully"
//       ).to.eventually.be.true
//     })
//   }

//   it('should make initial call', async () => {
//     const data = '0x' + 'ff'.repeat(100)
//     await globalInbox.sendL2Message(chainAddress, data)
//   })

//   it('should make second call', async () => {
//     const data = '0x' + 'ff'.repeat(100)
//     await globalInbox.sendL2Message(chainAddress, data)
//   })

//   it('should make bigger call', async () => {
//     const data = '0x' + 'ff'.repeat(1000)
//     await globalInbox.sendL2Message(chainAddress, data)
//   })

//   it('should make an empty L2 message', async () => {
//     const data = '0x'
//     const tx = await globalInbox.sendL2Message(chainAddress, data)
//     const receipt = await tx.wait()
//     console.log('Empty batch gas:', receipt.gasUsed?.toNumber())
//   })

//   it('tradeable-exits: initial', async () => {
//     await expect(
//       globalInbox.getPaymentOwner(originalOwner, messageIndex),
//       'current owner must be original owner'
//     ).to.eventually.equal(originalOwner)

//     await expect(
//       globalInbox
//         .connect(accounts[0])
//         .transferPayment(originalOwner, address2, messageIndex)
//     ).to.emit(globalInbox, 'PaymentTransfer')

//     await expect(
//       globalInbox.getPaymentOwner(originalOwner, messageIndex),
//       'current owner must be new owner (address2)'
//     ).to.eventually.equal(address2)
//   })

//   it('tradeable-exits: subsequent transfers', async () => {
//     await expect(
//       globalInbox.getPaymentOwner(originalOwner, messageIndex),
//       'current owner must be address2'
//     ).to.eventually.equal(address2)

//     await expect(
//       globalInbox
//         .connect(accounts[0])
//         .transferPayment(originalOwner, address2, messageIndex)
//     ).to.be.revertedWith('Must be payment owner')

//     await expect(
//       globalInbox
//         .connect(accounts[1])
//         .transferPayment(originalOwner, address3, messageIndex)
//     ).to.emit(globalInbox, 'PaymentTransfer')

//     await expect(
//       globalInbox.getPaymentOwner(originalOwner, messageIndex),
//       'current owner must be new owner (address3)'
//     ).to.eventually.equal(address3)

//     await expect(
//       globalInbox
//         .connect(accounts[1])
//         .transferPayment(originalOwner, address2, messageIndex)
//     ).to.be.revertedWith('Must be payment owner.')
//   })

//   it('tradeable-exits: commiting transfers', async () => {
//     const currOwner = await globalInbox.getPaymentOwner(
//       originalOwner,
//       messageIndex
//     )
//     await expect(
//       globalInbox
//         .connect(accounts[0])
//         .depositEthMessage(address4, originalOwner, {
//           value: 100,
//         })
//     ).to.emit(globalInbox, 'MessageDelivered')

//     await expect(globalInbox.getEthBalance(address4)).to.eventually.equal(100)
//     await expect(globalInbox.getEthBalance(currOwner)).to.eventually.equal(0)
//     await expect(globalInbox.getEthBalance(originalOwner)).to.eventually.equal(
//       0
//     )

//     const msgData = await getEthMessageData(originalOwner, currOwner, 50)

//     await globalInbox.connect(accounts[3]).sendMessages(msgData, 0, 1)

//     await expect(globalInbox.getEthBalance(address4)).to.eventually.equal(50)
//     await expect(globalInbox.getEthBalance(currOwner)).to.eventually.equal(50)
//     await expect(globalInbox.getEthBalance(originalOwner)).to.eventually.equal(
//       0
//     )

//     const msgData2 = await getEthMessageData(address4, originalOwner, 50)

//     await globalInbox.connect(accounts[3]).sendMessages(msgData2, 1, 2)

//     await expect(globalInbox.getEthBalance(address4)).to.eventually.equal(0)
//     await expect(globalInbox.getEthBalance(currOwner)).to.eventually.equal(50)
//     await expect(globalInbox.getEthBalance(originalOwner)).to.eventually.equal(
//       50
//     )
//   })
// })
