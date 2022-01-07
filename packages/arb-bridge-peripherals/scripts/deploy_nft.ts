import { ethers } from 'hardhat'
import { networks } from 'arb-ts/src/lib/networks'

const infuraKey = process.env['INFURA_KEY']
if (!infuraKey) throw new Error('No INFURA_KEY')

const l1Prov = new ethers.providers.JsonRpcProvider(
  'https://rinkeby.infura.io/v3/' + infuraKey
)
const l2Prov = new ethers.providers.JsonRpcProvider(
  'https://rinkeby.arbitrum.io/rpc'
)

const l1PrivKey = process.env['DEVNET_PRIVKEY']
if (!l1PrivKey) throw new Error('No DEVNET_PRIVKEY')
const l2PrivKey = process.env['DEVNET_PRIVKEY']
if (!l2PrivKey) throw new Error('No DEVNET_PRIVKEY')

// const l1Wallet = ethers.Wallet.fromMnemonic(l1PrivKey).connect(l1Prov)
// const l2Wallet = ethers.Wallet.fromMnemonic(l2PrivKey).connect(l2Prov)
const l1Wallet = new ethers.Wallet(l1PrivKey).connect(l1Prov)
const l2Wallet = new ethers.Wallet(l2PrivKey).connect(l2Prov)

const main = async () => {
  const chainId = (await l1Wallet.provider.getNetwork()).chainId
  const ethBridge = networks[chainId].ethBridge
  if (!ethBridge) throw new Error('No inbox')
  const inboxAddr = ethBridge.inbox

  const L1NftBridge = await (
    await ethers.getContractFactory('L1NftGateway')
  ).connect(l1Wallet)
  const L2NftBridge = await (
    await ethers.getContractFactory('L2NftGateway')
  ).connect(l2Wallet)
  const BeaconProxyFactory = await (
    await ethers.getContractFactory('BeaconProxyFactory')
  ).connect(l2Wallet)
  const StandardArbERC721 = await (
    await ethers.getContractFactory('StandardArbERC721')
  ).connect(l2Wallet)
  const UpgradeableBeacon = await (
    await ethers.getContractFactory('UpgradeableBeacon')
  ).connect(l2Wallet)

  const l1NftBridge = await L1NftBridge.deploy()
  await l1NftBridge.deployed()

  const l2NftBridge = await L2NftBridge.deploy()
  await l2NftBridge.deployed()

  const beaconProxyFactory = await BeaconProxyFactory.deploy()
  await beaconProxyFactory.deployed()

  const standardArbERC721Logic = await StandardArbERC721.deploy()
  await standardArbERC721Logic.deployed()

  const beacon = await UpgradeableBeacon.deploy(standardArbERC721Logic.address)
  await beacon.deployed()

  const init0 = await beaconProxyFactory.initialize(beacon.address)
  await init0.wait()

  const init1 = await l1NftBridge.initialize(l2NftBridge.address, inboxAddr)
  await init1.wait()

  const init2 = await l2NftBridge.initialize(
    l1NftBridge.address,
    beaconProxyFactory.address
  )
  await init2.wait()

  console.log('Deployed', {
    l1NftBridge: l1NftBridge.address,
    l2NftBridge: l2NftBridge.address,
  })

  // const Token = await (await ethers.getContractFactory('TestERC721')).connect(l1Wallet)
  // const name = 'mock'
  // const symbol = 'mck'
  // const token = await Token.deploy(name, symbol)
  // send escrowed tokens to bridge
  // const tokenId = 3
  // const tokenUri = '0xasdasdasd'
  // console.log("deployed test 721", token.address)

  // await token.deployed()
  // const mint = await token.mint(l1Wallet.address, tokenId, tokenUri)
  // await mint.wait()

  // const approv = await token.approve(l1NftBridge.address, tokenId)
  // await approv.wait()
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
