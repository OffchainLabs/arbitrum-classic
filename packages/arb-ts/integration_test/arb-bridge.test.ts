import { providers, utils, Wallet, BigNumber, constants } from 'ethers'
import { Bridge } from '../src/lib/bridge'
import { DepositTokenEventResult } from '../src/lib/bridge_helpers'

import { expect } from 'chai'
import config from './config'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'
import { TestERC777__factory } from '../src/lib/abi/factories/TestERC777__factory'
import { StandardArbERC20__factory } from '../src/lib/abi/factories/StandardArbERC20__factory'
import { StandardArbERC777__factory } from '../src/lib/abi/factories/StandardArbERC777__factory'

import { EthERC20Bridge } from '../src/lib/abi/EthERC20Bridge'
import yargs from 'yargs/yargs'
import chalk from 'chalk'

console.log()
console.log(chalk.green(`Starting Token Bridge Integrations Tests!`))
console.log()

const prettyLog = (text: string) => {
  console.log(chalk.blue(`*** ${text}`))
  console.log()
}

const warn = (text: string) => {
  console.log(chalk.red(`WARNING: ${text}`))
  console.log()
}
const argv = yargs(process.argv.slice(2)).argv

const network = ((_network?: string | number) => {
  if (!_network) {
    prettyLog('Using default network: kovan4')
    return 'kovan4'
  } else if (!Object.keys(config).includes(_network.toString())) {
    warn(
      `Network ${_network} not suppored; supported networks: ${Object.keys(
        config
      ).join(',')}`
    )
    process.exit(1)
    return
  } else {
    prettyLog('Using network ' + _network)
    return _network
  }
})(argv.network as string | number | undefined)

if (network === 'kovan4' && !process.env.INFURA_KEY) {
  warn('To use kovan, set env var INFURA_KEY')
  process.exit(1)
}

const { parseEther } = utils
const {
  ethRPC,
  arbRPC,
  preFundedSignerPK,
  erc20BridgeAddress,
  arbTokenBridgeAddress,
  l1gasPrice,
  existantTestERC20,
} = config.kovan4

const ethProvider = new providers.JsonRpcProvider(ethRPC)
const arbProvider = new providers.JsonRpcProvider(arbRPC)

const preFundedWallet = new Wallet(preFundedSignerPK, ethProvider)

const testPk = utils.formatBytes32String(Math.random().toString())

const l1TestWallet = new Wallet(testPk, ethProvider)
const l2TestWallet = new Wallet(testPk, arbProvider)
const wait = (ms = 10000) => new Promise(res => setTimeout(res, ms))

const depositAmount = '0.01'
let erc20Address = existantTestERC20
if (erc20Address) {
  prettyLog('Using token already deployed at l1 address: ' + erc20Address)
} else {
  prettyLog(`No L1 token address provided; we'll do it live!`)
}

prettyLog('Using preFundedWallet: ' + preFundedWallet.address)
prettyLog('Randomly generated test wallet: ' + l1TestWallet.address)

const bridge = new Bridge(
  erc20BridgeAddress,
  arbTokenBridgeAddress,
  l1TestWallet,
  l2TestWallet
)

describe('setup', () => {
  it('prefunded wallet is prefunded', async () => {
    const balance = await preFundedWallet.getBalance()
    const hasBalance = balance.gt(utils.parseEther(depositAmount))
    expect(hasBalance).to.be.true
    if (!hasBalance) {
      warn(
        preFundedWallet.address +
          ' not prefunded; set a funded wallet via env-var DEVNET_PRIVKEY. exiting.'
      )
      process.exit()
    }
  })
  it('fund l1 test wallet with eth', async () => {
    const res = await preFundedWallet.sendTransaction({
      to: l1TestWallet.address,
      value: utils.parseEther(depositAmount),
    })
    const rec = await res.wait()
    const testWAlletBalance = await l1TestWallet.getBalance()
    expect(testWAlletBalance.eq(parseEther(depositAmount))).to.be.true
  })
})

describe('Ether', () => {
  let testWalletL1EthBalance: BigNumber
  let testWalletL2EthBalance: BigNumber

  it('has expected initial values', async () => {
    testWalletL1EthBalance = await bridge.getAndUpdateL1EthBalance()
    testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
    expect(testWalletL1EthBalance.eq(parseEther(depositAmount))).to.be.true
    expect(testWalletL2EthBalance.eq(constants.Zero)).to.be.true
  })

  const ethToL2DepositAmount = parseEther('0.0001')
  const ethFromL2WithdrawAmmount = parseEther('0.00001')

  it('deposit ether transaction succeeds', async () => {
    const res = await bridge.depositETH(ethToL2DepositAmount)
    const rec = await res.wait()

    expect(rec.status).to.equal(1)
  })
  it('L2 address has expected balance after deposit eth', async () => {
    await wait()
    testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
    expect(testWalletL2EthBalance.eq(ethToL2DepositAmount)).to.be.true
  })
  it('withdraw Ether transaction succeeds and emits event', async () => {
    const withdrawEthRes = await bridge.withdrawETH(ethFromL2WithdrawAmmount)
    const withdrawEthRec = await withdrawEthRes.wait()

    expect(withdrawEthRec.status).to.equal(1)
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawEthRec)
    )[0]
    expect(withdrawEventData).to.exist

    const { indexInBatch, batchNumber } = withdrawEventData
    await wait()
    const proof = await bridge.tryGetProof(batchNumber, indexInBatch)
    expect(proof).to.exist
  })

  it('balance deducted after withdraw eth', async () => {
    wait()
    const bal = await bridge.getAndUpdateL2EthBalance()
    // lazy check, will update once gasPrice is activated
    expect(bal.lt(ethToL2DepositAmount)).to.be.true
  })
})
const tokenDepositAmmount = BigNumber.from(50)

const tokenDepositAmountE18 = utils.parseUnits('50', 18)
const tokenWithdrawAmountE18 = utils.parseUnits('2', 18)

describe('ERC20', () => {
  it('create/ensure l1 erc20 w initial supply', async () => {
    const testTokenFactory = await new TestERC20__factory(preFundedWallet)
    const testToken = erc20Address
      ? await testTokenFactory.attach(erc20Address)
      : await testTokenFactory.deploy()

    const bal = await testToken.balanceOf(preFundedWallet.address)
    expect(bal.gt(BigNumber.from(40000000))).to.be.true

    erc20Address = testToken.address
    const res = await testToken.transfer(
      l1TestWallet.address,
      BigNumber.from(200)
    )
    const rec = await res.wait()
    const data = await bridge.getAndUpdateL1TokenData(erc20Address)
    const testWalletBal = data.ERC20 && data.ERC20.balance
    expect(testWalletBal && testWalletBal.eq(BigNumber.from(200))).to.be.true
  })

  it('approve token for bridge contract', async () => {
    const approveRes = await bridge.approveToken(erc20Address)
    const approveRec = await approveRes.wait()
    expect(approveRec.status).to.equal(1)

    const data = await bridge.getAndUpdateL1TokenData(erc20Address)
    const allowed = data.ERC20 && data.ERC20.allowed
    expect(allowed).to.be.true
  })

  it('initial erc20 deposit txns — L1 and L2 — both succeed', async () => {
    const despositRes = await bridge.depositAsERC20(
      erc20Address,
      tokenDepositAmmount,
      BigNumber.from(10000000000000),
      BigNumber.from(0),
      undefined,
      { gasLimit: 210000, gasPrice: l1gasPrice }
    )

    const depositRec = await despositRes.wait()
    await wait()

    const tokenDepositData = (
      await bridge.getDepositTokenEventData(depositRec, 'ERC20')
    )[0] as DepositTokenEventResult
    const seqNum = tokenDepositData.seqNum

    const l2RetriableHash = await bridge.calculateL2RetryableTransactionHash(
      seqNum
    )

    const retriableReceipt = await arbProvider.waitForTransaction(
      l2RetriableHash
    )

    expect(depositRec.status === 1).to.be.true
    expect(retriableReceipt.status === 1).to.be.true
  })

  it('L2 wallet has expected balance after erc20 deposit', async () => {
    const l2Data = await bridge.getAndUpdateL2TokenData(erc20Address)

    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance

    expect(testWalletL2Balance && testWalletL2Balance.eq(tokenDepositAmmount))
      .to.be.true
  })

  it('erc20 contract is properly deployed in L2', async () => {
    const erc20L2Address = await bridge.getERC20L2Address(erc20Address)
    const arbERC20 = StandardArbERC20__factory.connect(
      erc20L2Address,
      arbProvider
    )
    const l2Code = await arbProvider.getCode(erc20L2Address)
    expect(l2Code.length > 2).to.be.true

    const balance = await arbERC20.balanceOf(l1TestWallet.address)
    expect(balance.eq(tokenDepositAmmount)).to.be.true
  })

  const tokenWithdrawAmount = BigNumber.from(2)

  it('withdraw erc20 succeeds and emits event data', async () => {
    const withdrawRes = await bridge.withdrawERC20(
      erc20Address,
      tokenWithdrawAmount
    )
    const withdrawRec = await withdrawRes.wait()
    expect(withdrawRec.status).to.equal(1)
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawRec)
    )[0]

    expect(withdrawEventData).to.exist
    const { indexInBatch, batchNumber } = withdrawEventData
    await wait()
    const proof = await bridge.tryGetProof(batchNumber, indexInBatch)
    expect(proof).to.exist
  })

  it('balance properly deducted after erc20 withdraw', async () => {
    await wait()
    const l2Data = await bridge.getAndUpdateL2TokenData(erc20Address)
    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance
    expect(
      testWalletL2Balance &&
        testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenDepositAmmount)
    ).to.be.true
  })
  // it('withdraw events queries works as expected',async ()=>{
  //   const l1EventData = await bridge.getL2ToL1EventData(l1TestWallet.address)
  //   const withdrawTokenData = await bridge.getTokenWithdrawEventData(
  //     l1TestWallet.address
  //   )
  //   console.log(l1EventData.length, withdrawTokenData.length)
  //   expect(l1EventData.length).to.equal(withdrawTokenData.length)
  // })
})

describe.skip('ERC777', () => {
  it('initial ERC777 deposit works', async () => {
    const despositRes = await bridge.depositAsERC777(
      erc20Address,
      tokenDepositAmmount,
      BigNumber.from(10000000000000),
      BigNumber.from(0),
      undefined,
      { gasLimit: 210000, gasPrice: l1gasPrice }
    )

    const depositRec = await despositRes.wait()

    const tokenDepositData = (
      await bridge.getDepositTokenEventData(depositRec, 'ERC777')
    )[0] as DepositTokenEventResult
    const seqNum = tokenDepositData.seqNum

    const l2RetriableHash = await bridge.calculateL2RetryableTransactionHash(
      seqNum
    )

    const retriableReceipt = await arbProvider.waitForTransaction(
      l2RetriableHash
    )
    expect(depositRec.status === 1).to.be.true
    expect(retriableReceipt.status === 1).to.be.true

    const l2Data = await bridge.getAndUpdateL2TokenData(erc20Address)

    const testWalletL2Balance = l2Data && l2Data.ERC777 && l2Data.ERC777.balance
    expect(testWalletL2Balance && testWalletL2Balance.eq(tokenDepositAmountE18))
      .to.be.true
  })

  it('erc777 is properly deployed in L2', async () => {
    const erc777L2Address = await bridge.getERC777L2Address(erc20Address)

    const arbERC777 = StandardArbERC777__factory.connect(
      erc777L2Address,
      arbProvider
    )
    const l2Code = await arbProvider.getCode(erc777L2Address)
    expect(l2Code.length > 2).to.be.true
    const balance = await arbERC777.balanceOf(l1TestWallet.address)
    expect(balance.eq(tokenDepositAmountE18)).to.be.true
  })

  it('withdraw erc77', async () => {
    const withdrawRes = await bridge.withdrawERC777(
      erc20Address,
      tokenWithdrawAmountE18
    )
    const withdrawRec = await withdrawRes.wait()
    expect(withdrawRec.status).to.equal(1)
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawRec)
    )[0]

    expect(withdrawEventData).to.exist
    await wait()

    const l2Data = await bridge.getAndUpdateL2TokenData(erc20Address)
    const testWalletL2Balance = l2Data && l2Data.ERC777 && l2Data.ERC777.balance
    expect(
      testWalletL2Balance &&
        testWalletL2Balance
          .add(tokenWithdrawAmountE18)
          .eq(tokenDepositAmountE18)
    ).to.be.true
  })
})
