import { Signer } from '@ethersproject/abstract-signer'
import hre from 'hardhat'
import * as fs from 'fs'
import {
  RollupAdminFacet__factory,
  RollupUserFacet__factory,
  Bridge__factory,
  Inbox__factory,
  SequencerInbox__factory,
  Outbox__factory,
  OldOutbox__factory,
  Rollup__factory,
  Node__factory,
} from '../build/types'

const ADDR_ONE = '0x0000000000000000000000000000000000000001'

export interface LogicAddresses {
  rollup: string
  oldOutbox: string
  outbox: string
  rollupAdmin: string
  rollupUser: string
  bridge: string
  inbox: string
  sequencerInbox: string
  node: string
}

async function deployContracts() {
  const accounts: Signer[] = await hre.ethers.getSigners()

  const RollupAdmin = new RollupAdminFacet__factory(accounts[0])
  const RollupUser = new RollupUserFacet__factory(accounts[0])
  const Bridge = new Bridge__factory(accounts[0])
  const Inbox = new Inbox__factory(accounts[0])
  const SequencerInbox = new SequencerInbox__factory(accounts[0])
  const Outbox = new Outbox__factory(accounts[0])
  const OldOutbox = new OldOutbox__factory(accounts[0])
  const Rollup = new Rollup__factory(accounts[0])
  const Node = new Node__factory(accounts[0])

  console.log('deploying Rollup')
  const rollup = await Rollup.deploy(1)
  await rollup.deployed()
  console.log(rollup.address)

  // rollup constructor makes this not initializable
  // const rollupInit = await rollup.initialize(...)
  // await rollupInit.wait()

  console.log('deploying OldOutbox')
  const oldOutbox = await OldOutbox.deploy()
  await oldOutbox.deployed()
  console.log(oldOutbox.address)

  const oldOutboxInit = await oldOutbox.initialize(ADDR_ONE, ADDR_ONE)
  await oldOutboxInit.wait()

  console.log('deploying Outbox')
  const outbox = await Outbox.deploy()
  await outbox.deployed()
  console.log(outbox.address)

  const outboxInit = await outbox.initialize(
    hre.ethers.constants.AddressZero,
    hre.ethers.constants.AddressZero
  )
  await outboxInit.wait()

  console.log('deploying rollup admin')
  const rollupAdmin = await RollupAdmin.deploy()
  await rollupAdmin.deployed()
  console.log(rollupAdmin.address)

  console.log('deploying rollup user')
  const rollupUser = await RollupUser.deploy()
  await rollupUser.deployed()
  console.log(rollupUser.address)

  console.log('init rollup user')
  const initRU = await rollupUser.initialize(hre.ethers.constants.AddressZero)
  await initRU.wait()

  console.log('deploying bridge')
  const bridge = await Bridge.deploy()
  await bridge.deployed()
  console.log(bridge.address)

  console.log('init bridge')
  const initBridge = await bridge.initialize()
  await initBridge.wait()

  console.log('deploying inbox')
  const inbox = await Inbox.deploy()
  await inbox.deployed()
  console.log(inbox.address)

  console.log('init inbox')
  const initInbox = await inbox.initialize(
    hre.ethers.constants.AddressZero,
    hre.ethers.constants.AddressZero
  )
  await initInbox.wait()

  console.log('deploying sequencer inbox')
  const sequencerInbox = await SequencerInbox.deploy()
  await sequencerInbox.deployed()
  console.log(sequencerInbox.address)

  console.log('init seq inbox')
  const initSeqInbox = await sequencerInbox.initialize(
    hre.ethers.constants.AddressZero,
    hre.ethers.constants.AddressZero,
    hre.ethers.constants.AddressZero
  )
  await initSeqInbox.wait()

  console.log('deploying node')
  const node = await Node.deploy()
  await node.deployed()
  console.log(node.address)

  const addresses: LogicAddresses = {
    rollup: rollup.address,
    oldOutbox: oldOutbox.address,
    outbox: outbox.address,
    rollupAdmin: rollupAdmin.address,
    rollupUser: rollupUser.address,
    bridge: bridge.address,
    inbox: inbox.address,
    sequencerInbox: sequencerInbox.address,
    node: node.address,
  }
  fs.writeFileSync('addresses.json', JSON.stringify(addresses))
}

const main = async () => {
  await deployContracts()
}

main()
  .then(() => console.log('done'))
  .catch(err => {
    console.error('error')
    console.error(err)
  })
