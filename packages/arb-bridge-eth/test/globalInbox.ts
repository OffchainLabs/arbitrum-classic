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

import { ethers, waffle } from '@nomiclabs/buidler'
import * as bre from '@nomiclabs/buidler'
import { utils, Signer } from 'ethers'
import * as chai from 'chai'
import * as chaiAsPromised from 'chai-as-promised'
import { deployMockContract } from '@ethereum-waffle/mock-contract'
import { GlobalInbox } from '../build/types/GlobalInbox'
import { ValueTester } from '../build/types/ValueTester'
import { ArbValue } from 'arb-provider-ethers'

chai.use(require('chai-as-promised'))

const { assert, expect } = chai

async function getMessageData(
  sender: utils.BigNumberish,
  receiver: utils.BigNumberish,
  value: utils.BigNumberish,
  value_tester: ValueTester
): Promise<Uint8Array> {
  let msgType = 1

  const msg = new ArbValue.TupleValue([
    new ArbValue.IntValue(1),
    new ArbValue.IntValue(ethers.utils.bigNumberify(sender).toString()),
    new ArbValue.TupleValue([
      new ArbValue.IntValue(ethers.utils.bigNumberify(receiver).toString()),
      new ArbValue.IntValue(value),
    ]),
  ])

  const msg_data = ArbValue.marshal(msg)
  let res = await value_tester.deserializeMessageData(msg_data, 0)

  assert.isTrue(res['0'], 'did not deserialize first part corrctly')

  let offset = res['1'].toNumber()
  assert.equal(res['2'].toNumber(), 1, 'Incorrect message type, must be ethMsg')
  assert.equal(res['3'], sender, 'Incorrect sender')

  let res2 = await value_tester.getEthMsgData(msg_data, offset)
  assert.isTrue(res2['0'], "value didn't deserialize correctly")
  assert.equal(res2['2'], receiver, 'Incorrect receiver')

  assert.equal(res2['3'].toNumber(), value, 'Incorrect value sent')

  return msg_data
}

function calcTxHash(
  chain: string,
  to: string,
  sequenceNum: utils.BigNumberish,
  value: utils.BigNumberish,
  messageData: string
) {
  return ethers.utils.solidityKeccak256(
    ['address', 'address', 'uint256', 'uint256', 'bytes32'],
    [
      chain,
      to,
      sequenceNum,
      value,
      ethers.utils.solidityKeccak256(['bytes'], [messageData]),
    ]
  )
}

async function generateTxData(
  accounts: Signer[],
  chain: string,
  messageCount: utils.BigNumberish
) {
  let txDataTemplate = {
    to: '0xffffffffffffffffffffffffffffffffffffffff',
    sequenceNum: 2000,
    value: 54254535454544,
    messageData: '0x00',
  }

  let transactionsData = []
  for (let i = 0; i < messageCount; i++) {
    transactionsData.push(txDataTemplate)
  }

  let data = '0x'

  for (var i = 0; i < transactionsData.length; i++) {
    let txData = transactionsData[i]

    let txHash = calcTxHash(
      chain,
      txData['to'],
      txData['sequenceNum'],
      txData['value'],
      txData['messageData']
    )
    let signedTxHash = await accounts[0].signMessage(txHash)
    let packedTxData = ethers.utils.solidityPack(
      ['uint16', 'address', 'uint256', 'uint256', 'bytes', 'bytes'],
      [
        (txData['messageData'].length - 2) / 2,
        txData['to'],
        txData['sequenceNum'],
        txData['value'],
        signedTxHash,
        txData['messageData'],
      ]
    )
    data += packedTxData.slice(2)
  }
  return data
}

let chain_address = ethers.utils.getAddress(
  '0xffffffffffffffffffffffffffffffffffffffff'
)

let accounts: Signer[]
let global_inbox: GlobalInbox
let value_tester: ValueTester

let nodeHash =
  '0x10c9d77c3846591fdfc3f966935819eb7dd71ebb7e71d5d081b880868ca33e4d'
let nodeHash2 =
  '0x20c9d77c3846591fdfc3f966935819eb7dd71ebb7e71d5d081b880868ca33e4d'
let messageIndex = 0
let originalOwner: string
let address2: string
let address3: string
let address4: string

describe('GlobalInbox', async () => {
  before(async () => {
    accounts = await ethers.getSigners()

    const GlobalInbox = await ethers.getContractFactory('GlobalInbox')
    global_inbox = (await GlobalInbox.deploy()) as GlobalInbox
    await global_inbox.deployed()

    const ValueTester = await ethers.getContractFactory('ValueTester')
    value_tester = (await ValueTester.deploy()) as ValueTester
    await value_tester.deployed()

    originalOwner = await accounts[0].getAddress()
    address2 = await accounts[1].getAddress()
    address3 = await accounts[2].getAddress()
    address4 = await accounts[3].getAddress()
  })

  it('should deposit eth', async () => {
    await expect(
      global_inbox.getEthBalance(chain_address),
      'Eth balance should start at 0'
    ).to.eventually.equal(0)

    await global_inbox.depositEthMessage(
      chain_address,
      await accounts[0].getAddress(),
      {
        value: 1000,
      }
    )

    await expect(
      global_inbox.getEthBalance(chain_address),
      "Eth balance wasn't deposited successfully"
    ).to.eventually.equal(1000)
  })

  // These tests use a waffle mock which depends on buidlerevm
  if (bre.network.name == 'buidlerevm') {
    it('should deposit an ERC20', async () => {
      const [mock_creator] = await waffle.provider.getWallets()
      const IERC20 = await ethers.getContractFactory('IERC20')
      const mockERC20 = await deployMockContract(
        mock_creator,
        IERC20.interface.abi
      )

      await mockERC20.mock.transferFrom.returns(1)

      await expect(
        global_inbox.getERC20Balance(mockERC20.address, chain_address),
        'ERC20 balance should start at 0'
      ).to.eventually.equal(0)

      await global_inbox.depositERC20Message(
        chain_address,
        chain_address,
        mockERC20.address,
        50
      )

      await expect(
        global_inbox.getERC20Balance(mockERC20.address, chain_address),
        "ERC20 Balance wasn't deposited successfully"
      ).to.eventually.equal(50)
    })

    it('should deposit an ERC721', async () => {
      const [mock_creator] = await waffle.provider.getWallets()
      const IERC721 = await ethers.getContractFactory('IERC721')
      const mockERC721 = await deployMockContract(
        mock_creator,
        IERC721.interface.abi
      )
      await expect(
        global_inbox.hasERC721(mockERC721.address, chain_address, 1234),
        'ERC721 Balance should start 0'
      ).to.eventually.be.false

      await mockERC721.mock.transferFrom.returns()
      await global_inbox.depositERC721Message(
        chain_address,
        chain_address,
        mockERC721.address,
        1234
      )

      await expect(
        global_inbox.hasERC721(mockERC721.address, chain_address, 1234),
        "ERC721 Balance wasn't deposited successfully"
      ).to.eventually.be.true
    })
  }

  it('should make initial call', async () => {
    await global_inbox.sendTransactionMessage(
      chain_address,
      chain_address,
      2000,
      54254535454544,
      '0x'
    )
  })

  it('should make second call', async () => {
    await global_inbox.sendTransactionMessage(
      chain_address,
      chain_address,
      2000,
      54254535454544,
      '0x'
    )
  })

  it('should make bigger call', async () => {
    await global_inbox.sendTransactionMessage(
      chain_address,
      chain_address,
      2000,
      54254535454544,
      // 64 bytes of data
      '0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff'
    )
  })

  it('should make a batch call', async () => {
    let messageCount = 500

    // console.log(data);

    let data = await generateTxData(accounts, chain_address, messageCount)

    let tx_promise = global_inbox.deliverTransactionBatch(chain_address, data)

    await expect(tx_promise)
      .to.emit(global_inbox, 'TransactionMessageBatchDelivered')
      .withArgs(chain_address)

    let tx = await tx_promise
    let [chainInput, txDataInput] = ethers.utils.defaultAbiCoder.decode(
      ['address', 'bytes'],
      ethers.utils.hexDataSlice(tx.data, 4)
    )

    assert.equal(
      chainInput.toLowerCase(),
      chain_address.toLowerCase(),
      'incorrect chain from input'
    )

    assert.equal(
      txDataInput.toLowerCase(),
      data.toLowerCase(),
      'incorrect tx data from input'
    )
  })

  it('tradeable-exits: initial', async () => {
    await expect(
      global_inbox.getPaymentOwner(originalOwner, nodeHash, messageIndex),
      'current owner must be original owner'
    ).to.eventually.equal(originalOwner)

    await expect(
      global_inbox
        .connect(accounts[0])
        .transferPayment(originalOwner, address2, nodeHash, messageIndex)
    ).to.emit(global_inbox, 'PaymentTransfer')

    await expect(
      global_inbox.getPaymentOwner(originalOwner, nodeHash, messageIndex),
      'current owner must be new owner (address2)'
    ).to.eventually.equal(address2)
  })

  it('tradeable-exits: subsequent transfers', async () => {
    await expect(
      global_inbox.getPaymentOwner(originalOwner, nodeHash, messageIndex),
      'current owner must be address2'
    ).to.eventually.equal(address2)

    await expect(
      global_inbox
        .connect(accounts[0])
        .transferPayment(originalOwner, address2, nodeHash, messageIndex)
    ).to.be.revertedWith('Must be payment owner')

    await expect(
      global_inbox
        .connect(accounts[1])
        .transferPayment(originalOwner, address3, nodeHash, messageIndex)
    ).to.emit(global_inbox, 'PaymentTransfer')

    await expect(
      global_inbox.getPaymentOwner(originalOwner, nodeHash, messageIndex),
      'current owner must be new owner (address3)'
    ).to.eventually.equal(address3)

    await expect(
      global_inbox
        .connect(accounts[1])
        .transferPayment(originalOwner, address2, nodeHash, messageIndex)
    ).to.be.revertedWith('Must be payment owner.')
  })

  it('tradeable-exits: commiting transfers', async () => {
    let curr_owner = await global_inbox.getPaymentOwner(
      originalOwner,
      nodeHash,
      messageIndex
    )
    await expect(
      global_inbox
        .connect(accounts[0])
        .depositEthMessage(address4, originalOwner, {
          value: 100,
        })
    ).to.emit(global_inbox, 'EthDepositMessageDelivered')

    await expect(global_inbox.getEthBalance(address4)).to.eventually.equal(100)
    await expect(global_inbox.getEthBalance(curr_owner)).to.eventually.equal(0)
    await expect(global_inbox.getEthBalance(originalOwner)).to.eventually.equal(
      0
    )

    let msg_data = await getMessageData(
      originalOwner,
      curr_owner,
      50,
      value_tester
    )

    await global_inbox
      .connect(accounts[3])
      .sendMessages(msg_data, [1], [nodeHash])

    await expect(global_inbox.getEthBalance(address4)).to.eventually.equal(50)
    await expect(global_inbox.getEthBalance(curr_owner)).to.eventually.equal(50)
    await expect(global_inbox.getEthBalance(originalOwner)).to.eventually.equal(
      0
    )

    let msg_data2 = await getMessageData(
      address4,
      originalOwner,
      50,
      value_tester
    )

    await global_inbox
      .connect(accounts[3])
      .sendMessages(msg_data2, [1], [nodeHash2])

    await expect(global_inbox.getEthBalance(address4)).to.eventually.equal(0)
    await expect(global_inbox.getEthBalance(curr_owner)).to.eventually.equal(50)
    await expect(global_inbox.getEthBalance(originalOwner)).to.eventually.equal(
      50
    )
  })

  it('tradeable-exits: commiting transfers, different mnsg indexes', async () => {
    let chain_address = await accounts[4].getAddress()
    let dest_address = await accounts[5].getAddress()
    await expect(
      global_inbox
        .connect(accounts[1])
        .depositEthMessage(chain_address, dest_address, {
          value: 100,
        })
    ).to.emit(global_inbox, 'EthDepositMessageDelivered')

    let msg_data = await getMessageData(
      chain_address,
      dest_address,
      10,
      value_tester
    )

    await global_inbox
      .connect(accounts[4])
      .sendMessages(msg_data, [0], [nodeHash2])

    await expect(global_inbox.getEthBalance(chain_address)).to.eventually.equal(
      100
    )
    await expect(global_inbox.getEthBalance(dest_address)).to.eventually.equal(
      0
    )

    await global_inbox
      .connect(accounts[4])
      .sendMessages(msg_data, [1], [nodeHash2])

    await expect(global_inbox.getEthBalance(chain_address)).to.eventually.equal(
      90
    )
    await expect(global_inbox.getEthBalance(dest_address)).to.eventually.equal(
      10
    )
  })
})
