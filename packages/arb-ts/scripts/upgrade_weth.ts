import { instantiateBridge } from './instantiate_bridge'
import { AeWETH__factory } from '../src/lib/abi/factories/AeWETH__factory'
import { TransparentUpgradeableProxy__factory } from '../src/lib/abi/factories/TransparentUpgradeableProxy__factory'

const main = async () => {
  const { bridge, l2Network } = await instantiateBridge()
  const { l2Signer } = bridge.l2Bridge

  const aeWeth = new AeWETH__factory(l2Signer)
  const res = await aeWeth.deploy()
  const rec = await res.deployTransaction.wait()
  const logicAddress = res.address

  console.log('aeWeth logic deployed to ', logicAddress)

  const connectedProxy = TransparentUpgradeableProxy__factory.connect(
    l2Network.tokenBridge.l2Weth,
    l2Signer
  )
  const upgradeRes = await connectedProxy.upgradeTo(logicAddress)
  const upgradeRec = await upgradeRes.wait()

  console.log('successfully upgraded WETH logic', upgradeRec)
}

const initWETH = async () => {
  const { bridge, l2Network } = await instantiateBridge()
  const { l2Signer } = bridge.l2Bridge

  const aeWeth = AeWETH__factory.connect(l2Network.tokenBridge.l2Weth, l2Signer)
  const res = await aeWeth.initialize(
    'Wrapped Ether',
    'WETH',
    18,
    l2Network.tokenBridge.l2WethGateway,
    l2Network.tokenBridge.l1Weth
  )

  const rec = await res.wait()
  console.warn('initialized weth', rec)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
