import { providers, utils, Wallet } from 'ethers'
import { L1Bridge } from '../src/lib/L1Bridge'
import { withdrawEth } from '../src/lib/L2Bridge'
import { expect } from 'chai'

const arbChainAddress = process.argv[process.argv.length - 1]

if (!arbChainAddress.startsWith('0x')) {
  console.warn(
    `Error: include arb chain address, i.e,: yarn run test_bridge 0xaddress`
  )
  process.exit(1)
}
const { BigNumber } = utils
const ethProviderUrl = 'http://0.0.0.0:7545'
const arbProviderUrl = 'http://0.0.0.0:8547'

const walletAddress = '0xc7711f36b2C13E00821fFD9EC54B04A60AEfbd1b'
const walletKey =
  '0x979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76'
const ethProvider = new providers.JsonRpcProvider(ethProviderUrl)

const arbProvider = new providers.JsonRpcProvider(arbProviderUrl)

//  const ethSigner = ethProvider.getSigner(walletAddress)
const ethSigner = new Wallet(walletKey, ethProvider)
const arbSigner = new Wallet(walletKey, arbProvider)

const l1Bridge = new L1Bridge(ethSigner, arbChainAddress)

const testValue = utils.parseEther('0.1')

describe('setup bridge', () => {
  it('has funds', async () => {
    const ethInitialBalance = await ethProvider.getBalance(walletAddress)
    expect(ethInitialBalance.gt(new BigNumber(500))).to.be.true
  })
})

describe('l1 Bridge', () => {
  it('deposits ETH', async () => {
    const arbBalance = await arbProvider.getBalance(walletAddress)

    const tx = await l1Bridge.depositETH(walletAddress, testValue)
    await tx.wait()
    setTimeout(async () => {
      const newArbBalance = await arbProvider.getBalance(walletAddress)
      expect(newArbBalance.sub(testValue).eq(arbBalance)).to.be.true
    }, 3000)
  })
})

describe('l2', async () => {
  it('simple eth transfer', async () => {
    const randomAddress = Wallet.createRandom().address
    const value = utils.parseEther('0.01')
    const tx = await arbSigner.sendTransaction({
      to: randomAddress,
      value,
    })
    await tx.wait()
    setTimeout(async () => {
      const arbBalance = await arbProvider.getBalance(randomAddress)
      expect(arbBalance.eq(value)).to.be.true
    }, 3000)
  })
})

describe('l2 bridge', async () => {
  it('withdraws ETH', async () => {
    const setupTx = await l1Bridge.depositETH(walletAddress, testValue)
    await setupTx.wait()
    const arbBalance = await arbProvider.getBalance(walletAddress)
    expect(arbBalance.gte(testValue)).to.be.true
    const tx = await withdrawEth(arbSigner, testValue)
    console.info('withdrawEth response:', tx)
    console.info('waiting for receipt...')
    const rec = await tx.wait()
    console.info('withdraw receipt arrived!', rec)
    expect(true).to.be.true
  })
})
