import { ethers } from 'hardhat'

import { createInterface } from 'readline'

import { RollupCreator__factory } from 'arb-ts/src/lib/abi/factories/RollupCreator__factory'
import { Rollup__factory } from 'arb-ts/src/lib/abi/factories/Rollup__factory'
import { ProxyAdmin__factory } from 'arb-ts/src/lib/abi/factories/ProxyAdmin__factory'

interface DeployedContracts {
  bridgeCreator: string
  currentBridge: string
  challengeFactory: string
  nodeFactory: string
  osp: string
  osp2: string
  ospHash: string
  rollupCreator: string
  rollup: string
}

// TODO: load from deployments?
const prevAddresses: DeployedContracts = {
  bridgeCreator: '',
  currentBridge: '',
  challengeFactory: '',
  nodeFactory: '',
  osp: '',
  osp2: '',
  ospHash: '',
  rollupCreator: '',
  rollup: '',
}

const newAddresses: DeployedContracts = {
  bridgeCreator: '',
  currentBridge: '',
  challengeFactory: '',
  nodeFactory: '',
  osp: '',
  osp2: '',
  ospHash: '',
  rollupCreator: '',
  rollup: '',
}

async function promptUserToContinue() {
  const rl = createInterface({
    input: process.stdin, //or fileStream
    output: process.stdout,
  })
  rl.setPrompt('Continue? (Y/N) ')
  rl.prompt()
  const it = rl[Symbol.asyncIterator]()
  const line1 = await it.next()
  rl.close()

  switch (line1.value) {
    case 'y':
    case 'Y': {
      return
    }
    case 'n':
    case 'N': {
      console.log('Selected: No')
      process.exit(1)
      break
    }
    default: {
      console.log('Please input (Y)es or (N)o.')
      await promptUserToContinue()
    }
  }
}

async function main() {
  const accounts = await ethers.getSigners()

  const rollupCreator = RollupCreator__factory.connect(
    prevAddresses.rollupCreator,
    ethers.provider
  )

  const rollupCreatorOwner = await rollupCreator.owner()
  if (accounts[0].address.toLowerCase() !== rollupCreatorOwner.toLowerCase()) {
    throw new Error('Current account must be rollup creator owner')
  }

  const linkedBridgeCreator = await rollupCreator.bridgeCreator()
  const linkedRollupTemplate = await rollupCreator.rollupTemplate()
  const linkedChallengeFactory = await rollupCreator.challengeFactory()
  const linkedNodeFactory = await rollupCreator.nodeFactory()

  // check if current linked templates match the supplied previous ones
  if (
    linkedBridgeCreator.toLowerCase() !==
    prevAddresses.bridgeCreator.toLowerCase()
  ) {
    throw new Error('Wrong bridge creator')
  }
  if (
    linkedRollupTemplate.toLowerCase() !==
    prevAddresses.rollupCreator.toLowerCase()
  ) {
    throw new Error('Wrong rollup template')
  }
  if (
    linkedChallengeFactory.toLowerCase() !==
    prevAddresses.challengeFactory.toLowerCase()
  ) {
    throw new Error('Wrong challenge factory')
  }
  if (
    linkedNodeFactory.toLowerCase() !== prevAddresses.nodeFactory.toLowerCase()
  ) {
    throw new Error('Wrong node factory')
  }

  const rollupsCreated = await rollupCreator.queryFilter(
    rollupCreator.filters.RollupCreated(prevAddresses.rollup, null, null),
    0, // fromBlock
    'latest' // toBlock
  )
  const parsedLog = rollupCreator.interface.parseLog(rollupsCreated[0])
  const { rollupAddress, inboxAddress, adminProxy } = parsedLog.args

  console.log({ rollupAddress })
  console.log({ adminProxy })

  if (rollupsCreated.length !== 1) {
    throw new Error('Rollup was not created')
  }

  const rollupInstance = Rollup__factory.connect(
    prevAddresses.rollup,
    ethers.provider
  )
  const rollupOwner = await rollupInstance.owner()

  if (rollupOwner.toLowerCase() !== accounts[0].address.toLowerCase()) {
    throw new Error('Current account must be rollup owner')
  }

  const proxyAdmin = ProxyAdmin__factory.connect(
    adminProxy as string,
    ethers.provider
  )
  // TODO: we can check the direct storage slot as in 'packages/arb-bridge-peripherals/scripts/upgrade_bridge_logic.ts'
  const proxyAdminOwner = await proxyAdmin.owner()

  if (proxyAdminOwner.toLowerCase() !== accounts[0].address.toLowerCase()) {
    throw new Error('Current account must be admin proxy owner')
  }

  const latestConfirmedNode = await rollupInstance.latestConfirmed()
  const latestCreatedNode = await rollupInstance.latestNodeCreated()

  if (latestCreatedNode.gt(latestConfirmedNode)) {
    await promptUserToContinue()
  }
  // TODO: let user know if unresolved challenges

  // bridgeCreator is responsible for the following:
  // Bridge, SequencerInbox, Inbox, RollupEventBridge, Outbox

  // TODO: update templates in creators

  // TODO: deploy new instances, then update connected contracts
  // const updateTx = await rollupInstance.updateConnectedContracts(
  //   [
  // [delayedBridge, sequencerInbox, outbox, rollupEventBridge, challengeFactory, nodeFactory]
  //   ],
  //   [true, true, true, true, true, true]
  // )
  // const updateReceipt = await updateTx.wait()
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
