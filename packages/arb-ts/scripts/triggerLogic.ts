import { Rollup__factory } from '../src/lib/abi/factories/Rollup__factory'
import { RollupEventBridge__factory } from '../src/lib/abi/factories/RollupEventBridge__factory'
import { Bridge__factory } from '../src/lib/abi/factories/Bridge__factory'

import { ethers } from 'ethers'

const main = async () => {
  const provider = new ethers.providers.JsonRpcProvider('http://127.0.0.1:8545')
  const signer = new ethers.Wallet(process.env['DEVNET_PRIVKEY']).connect(
    provider
  )

  // const rollup = "0xc12ba48c781f6e392b49db2e25cd0c28cd77531a"
  const delayedBridgeLogic = '0xcb0da32914a683286ed2d4890e8157ffecc9bd06'
  const rollupLogic = '0xae71755b42d1ef5fb365aeb4a74cb73992dd9fbe'
  const eventBridgeLogic = '0x5c7355e46d5486583a1cc211701e25004231d9dd'

  const rollup = Rollup__factory.connect(rollupLogic, signer)
  const delayedBridge = Bridge__factory.connect(delayedBridgeLogic, signer)

  const delayInit = await delayedBridge.initialize()
  await delayInit.wait()

  const setInbox = await delayedBridge.transferOwnership(rollupLogic)
  await setInbox.wait()

  const rollupEventBridge = RollupEventBridge__factory.connect(
    eventBridgeLogic,
    signer
  )
  const initEventBridge = await rollupEventBridge.initialize(
    delayedBridgeLogic,
    rollupLogic
  )
  await initEventBridge.wait()

  const nodeFactory = '0x4442a659aaa27eeddf0a5d6ef7659290e92f678c'
  const sequencerInbox = '0x0000000000000000000000000000000000000000'
  const outbox = '0x0000000000000000000000000000000000000000'
  const challengeFactory = '0x0000000000000000000000000000000000000000'

  const tx = await rollup.initialize(
    '0x0000000000000000000000000000000000000000000000000000000000000000',
    [1, 0, 0, 0],
    '0x0000000000000000000000000000000000000000',
    '0x0000000000000000000000000000000000000000',
    '0x',
    [
      delayedBridgeLogic,
      sequencerInbox,
      outbox,
      eventBridgeLogic,
      challengeFactory,
      nodeFactory,
    ],
    [
      '0x0000000000000000000000000000000000000000',
      '0x0000000000000000000000000000000000000000',
    ],
    [0, 0]
  )
  await tx.wait()

  try {
    const initAgain = await rollup.initialize(
      '0x0000000000000000000000000000000000000000000000000000000000000000',
      [1, 0, 0, 0],
      '0x0000000000000000000000000000000000000000',
      '0x0000000000000000000000000000000000000000',
      '0x',
      [
        delayedBridgeLogic,
        sequencerInbox,
        outbox,
        eventBridgeLogic,
        challengeFactory,
        nodeFactory,
      ],
      [
        '0x0000000000000000000000000000000000000000',
        '0x0000000000000000000000000000000000000000',
      ],
      [0, 0]
    )
    await initAgain.wait()
    console.log('this should have reverted :angry:')
  } catch (e) {
    console.log('reverted as intended')
  }
  // console.log({tx})
  // const receipt = await tx.wait()
  // console.log({receipt})
}

main()
  .then(() => console.log('done'))
  .catch(err => console.error(err))
