import { ethers } from 'hardhat'

import MainnetDeployments from '../deployment-42161.json'

const infuraKey = process.env['INFURA_KEY']
if (!infuraKey) throw new Error('No INFURA_KEY')

const l1Prov = new ethers.providers.JsonRpcProvider(
  'https://mainnet.infura.io/v3/' + infuraKey
)
const l2Prov = new ethers.providers.JsonRpcProvider(
  'https://arb1.arbitrum.io/rpc'
)

const l1privKey = process.env['L1_PRIVKEY']
if (!l1privKey) throw new Error('No L1_PRIVKEY')

const l2privKey = process.env['L2_PRIVKEY']
if (!l2privKey) throw new Error('No L2_PRIVKEY')

const L1Signer = ethers.Wallet.fromMnemonic(l1privKey)
const L2Signer = ethers.Wallet.fromMnemonic(l2privKey)

const l1Signer = L1Signer.connect(l1Prov)
const l2Signer = L1Signer.connect(l2Prov)

const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

const l1GatewayRouterAddr = MainnetDeployments.l1GatewayRouter
const l2GatewayRouterAddr = MainnetDeployments.l2GatewayRouter

const l1token = '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48'
const l2token = '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8'

// L1 to L2 call parameters
// const gasPriceBid = '0'
// const maxGas = '800000000'
// const maxSubmissionCost = '259829212830'
// TODO: actually calculate the needed value
// const deposit = ethers.utils.parseEther('0.01')

const main = async () => {
  const l1Router = (
    await ethers.getContractAt('L1GatewayRouter', l1GatewayRouterAddr)
  ).connect(l1Signer)
  const l2Router = (
    await ethers.getContractAt('L2GatewayRouter', l2GatewayRouterAddr)
  ).connect(l2Signer)

  //   check if user owns router
  const expectedOwner = await l1Router.owner()

  if (expectedOwner.toLowerCase() !== l1Signer.address.toLowerCase()) {
    throw new Error('Not router owner')
  }

  // get factories and deploy logic

  const L1CustomGateway = (
    await ethers.getContractFactory('L1CustomGateway')
  ).connect(l1Signer)
  const L2CustomGateway = (
    await ethers.getContractFactory('L2CustomGateway')
  ).connect(l2Signer)

  const l1CustomGateway = await L1CustomGateway.attach(
    '0xcEe284F754E854890e311e3280b767F80797180d'
  )
  await l1CustomGateway.deployed()
  console.log('got L1 Custom gateway')

  const l2CustomGateway = await L2CustomGateway.attach(
    '0x096760F208390250649E3e8763348E783AEF5562'
  )
  await l2CustomGateway.deployed()
  console.log('got L2 Custom gateway')

  const maxGas = 4000000
  const gasPriceBid = (await l2Signer.getGasPrice()).mul(2)
  const maxSubmissionCost = 400000000000

  // const l2WethFromL2 = await l2Router.calculateL2TokenAddress(l1token)
  // const l2WethFromL1 = await l1Router.calculateL2TokenAddress(l1token)
  // const expectedGateway = await l1Router.getGateway(l1token)

  // console.log({
  //   l1counterpart: await l1CustomGateway.counterpartGateway(),
  //   l2counterpart: await l2CustomGateway.counterpartGateway()
  // })

  // console.log({
  //   l2WethFromL1,
  //   l2WethFromL2,
  //   l2token,
  //   expectedGateway
  // })

  // init
  console.log('init L1')
  const l1SetTx = await l1CustomGateway.forceRegisterTokenToL2(
    [l1token],
    [l2token],
    maxGas,
    gasPriceBid,
    maxSubmissionCost,
    {
      value: gasPriceBid.mul(maxGas).add(maxSubmissionCost),
      // nonce: 16,
      gasPrice: 30000000000,
    }
  )
  const l1Receipt = await l1SetTx.wait()

  console.log(l1Receipt)

  // const bridge = await Bridge.init(
  //     l1Signer,
  //     l2Signer
  // )

  // const l1Receipt = await l1Signer.provider.getTransactionReceipt("0xa47653f4150ffcb5c7475afeed691ab6e93c2311093794f07896cd42cbd44f86")
  // const seqNum = await bridge.getInboxSeqNumFromContractTransaction(l1Receipt)

  // if(!seqNum) {
  //     throw new Error("no seq")
  // }

  // const l2Ticket = await bridge.calculateL2RetryableTransactionHash(seqNum[0])
  // const l2UserTx = await bridge.calculateL2TransactionHash(seqNum[0])
  // const l2Redeem = await bridge.calculateRetryableAutoRedeemTxnHash(seqNum[0])

  // console.log({
  //     l2Ticket,
  //     l2UserTx,
  //     l2Redeem
  //   })

  //   console.log("waiting for tx")
  //   const tx = await l2Signer.provider.waitForTransaction(l2UserTx)
  //   console.log({tx})

  // set gateway in router

  // const setGatewayTx = await l1Router.setGateways(
  //   [l1WethAddr],
  //   [l1WethGatewayProxy.address],
  //   maxGas,
  //   gasPriceBid,
  //   maxSubmissionCost,
  //   {
  //     value: deposit,
  //   }
  // )

  // console.log({ setGatewayTx })

  // await setGatewayTx.wait()
}

main()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
