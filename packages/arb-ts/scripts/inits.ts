import { instantiateBridge } from './instantiate_bridge'
import { AeWETH__factory } from '../src/lib/abi/factories/AeWETH__factory'
import { TransparentUpgradeableProxy__factory } from '../src/lib/abi/factories/TransparentUpgradeableProxy__factory'
import { utils, Contract } from 'ethers'

const updgradeWeth = async () => {
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

const initUSDC = async () => {
  const { bridge, l2Network } = await instantiateBridge()
  const { l2Signer } = bridge.l2Bridge

  const contractInterface = new utils.Interface([
    ` function initialize(
        address _gatewayAddress,
        address _l1Address,
        address owner,
        string memory name,
        string memory symbol,
        uint8 decimals
    ) external
  `,
  ])
  const usdcContract = new Contract(
    '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8',
    contractInterface
  ).connect(l2Signer)

  const res = await usdcContract.initialize(
    l2Network.tokenBridge.l2CustomGateway,
    '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
    '0xBB1a241DCBd6A3894cB61F659034874Dc9CF65D4',
    'USD Coin (Arb1)',
    'USDC',
    6
  )
  const rec = await res.wait()
  console.log('initialized USDC', rec)

  console.log('setting post init hook:')

  const customGatewayInterface = new utils.Interface([
    ` function postUpgradeInit() external`,
  ])

  const customGatewayContract = new Contract(
    l2Network.tokenBridge.l2CustomGateway,
    customGatewayInterface
  ).connect(l2Signer)

  const postInitRes = await customGatewayContract.postUpgradeInit()
  const postInitRec = await postInitRes.wait()

  console.log('ran postUpgradeInit', postInitRec)
}

const initBoth = async () => {
  console.log('initing weth')

  await initWETH()
  console.log('done initing weth')

  console.log('initing usdc')
  await initUSDC()
  console.log('done initing usdc')
}

initBoth()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
