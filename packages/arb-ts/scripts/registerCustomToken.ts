import args from './getCLargs'
import { instantiateBridge } from './instantiate_bridge'
import {
  L1CustomGateway__factory,
  IArbToken__factory,
  NodeInterface__factory,
} from '../src/lib/abi'
import { utils } from 'ethers'

if (!args.l1Address) {
  throw new Error('Include token address (--l1Address 0xmyaddress)')
}
if (!args.l2Address) {
  throw new Error('Include token address (--l2Address 0xmyaddress)')
}

const main = async () => {
  const { l1Address, l2Address } = args as { [key: string]: string }
  const { bridge, l1Network, l2Network } = await instantiateBridge()

  const l1CustomGateway = await L1CustomGateway__factory.connect(
    l1Network.tokenBridge.l1CustomGateway,
    bridge.l1Signer
  )

  const owner = await l1CustomGateway.owner()

  if (
    owner.toLowerCase() !== (await bridge.l1Signer.getAddress()).toLowerCase()
  ) {
    throw new Error(
      `Current L1 signer ${await bridge.l1Signer.getAddress()} is not owner ${owner}`
    )
  }

  const token = await IArbToken__factory.connect(l2Address, bridge.l2Provider)

  if (l1Address.toLowerCase() !== (await token.l1Address()).toLowerCase()) {
    throw new Error('L2 token set to different L1 token')
  }

  const maxSubmissionCost = (await bridge.l2Bridge.getTxnSubmissionPrice(80))[0]

  const nodeInterface = NodeInterface__factory.connect(
    '0x00000000000000000000000000000000000000C8',
    bridge.l2Provider
  )

  const gasPrice = (await bridge.l2Provider.getGasPrice()).mul(2)

  const maxGas = (
    await nodeInterface.estimateRetryableTicket(
      l1Network.tokenBridge.l1CustomGateway,
      utils.parseEther('0.5'),
      l2Network.tokenBridge.l2CustomGateway,
      0,
      maxSubmissionCost,
      await bridge.l2Signer.getAddress(),
      await bridge.l2Signer.getAddress(),
      10000000,
      gasPrice,
      '0x'
    )
  )[0].mul(2)
  console.log('Max gas', maxGas)

  console.log('sending L1 tx')
  const l1Tx = await l1CustomGateway.forceRegisterTokenToL2(
    [l1Address],
    [l2Address],
    maxGas,
    gasPrice,
    maxSubmissionCost
  )
  console.log('waiting for tx to be mined')
  const l1Receipt = await l1Tx.wait(3)
  console.log('got L1 tx mined iwth hash ', l1Receipt.transactionHash)

  const l2Retryable = await bridge.getL2TxHashByRetryableTicket(l1Receipt)
  console.log('waiting for l2 tx with hash ', l2Retryable)

  const l2RetryableReceipt = await bridge.l2Provider.waitForTransaction(
    l2Retryable
  )
  console.log('Got it!')
}

main()
  .then(() => console.log('Done registering.'))
  .catch(err => console.error(err))
