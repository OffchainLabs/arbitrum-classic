import { providers, utils, Wallet, BigNumber, constants } from 'ethers'
import { Bridge } from '../src/lib/bridge'
import {
  DepositTokenEventResult,
  BridgeHelper,
} from '../src/lib/bridge_helpers'

import { expect } from 'chai'
import config from './config'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'
import { TestCustomTokenL1__factory } from '../src/lib/abi/factories/TestCustomTokenL1__factory'
import { TestArbCustomToken__factory } from '../src/lib/abi/factories/TestArbCustomToken__factory'

import { StandardArbERC20__factory } from '../src/lib/abi/factories/StandardArbERC20__factory'
import { StandardArbERC20 } from '../src/lib/abi/StandardArbERC20'
import { TestERC20 } from '../src/lib/abi/TestERC20'

import { EthERC20Bridge } from '../src/lib/abi/EthERC20Bridge'
import yargs from 'yargs/yargs'
import chalk from 'chalk'
import { TestCustomTokenL1 } from '../src/lib/abi/TestCustomTokenL1'
import { TestArbCustomToken } from '../src/lib/abi/TestArbCustomToken'

const { Zero, AddressZero } = constants
console.log(chalk.green(`Starting Token Bridge Integrations Tests!`))
console.log()

const prettyLog = (text: string) => {
  console.log(chalk.blue(`    *** ${text}`))
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
      `Network ${_network} not supported; supported networks: ${Object.keys(
        config
      ).join(',')}`
    )
    process.exit(1)
    return
  } else {
    prettyLog('Using network ' + _network)
    return _network.toString()
  }
})(argv.network as string | number | undefined) as 'kovan4' | 'devnet'

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
  existentTestERC20,
  existentCustomTokenL1,
  existentCustomTokenL2,
  defaultWait,
} = config[network]

if (
  [existentCustomTokenL1, existentCustomTokenL2].filter((x: any) => x)
    .length === 1
) {
  warn(
    'Include either both a pre-deployed / custom  token L1 and L2 or neither'
  )
  process.exit()
}

const ethProvider = new providers.JsonRpcProvider(ethRPC)
const arbProvider = new providers.JsonRpcProvider(arbRPC)

const preFundedWallet = new Wallet(preFundedSignerPK, ethProvider)

const testPk = utils.formatBytes32String(Math.random().toString())

const l1TestWallet = new Wallet(testPk, ethProvider)
const l2TestWallet = new Wallet(testPk, arbProvider)
const wait = (ms = 0) => {
  return new Promise(res => setTimeout(res, ms || defaultWait))
}

const depositAmount = '0.01'
let erc20Address = existentTestERC20
prettyLog('Using preFundedWallet: ' + preFundedWallet.address)
prettyLog('Randomly generated test wallet: ' + l1TestWallet.address)

const bridge = new Bridge(
  erc20BridgeAddress,
  arbTokenBridgeAddress,
  l1TestWallet,
  l2TestWallet
)

before('setup', () => {
  it("'pre-funded wallet' is indeed pre-funded", async () => {
    const balance = await preFundedWallet.getBalance()
    const hasBalance = balance.gt(utils.parseEther(depositAmount))
    expect(hasBalance).to.be.true
    if (!hasBalance) {
      warn(
        preFundedWallet.address +
          ' not pre-funded; set a funded wallet via env-var DEVNET_PRIVKEY. exiting.'
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
    const testWalletBalance = await l1TestWallet.getBalance()
    expect(testWalletBalance.eq(parseEther(depositAmount))).to.be.true
  })
})

describe('Ether', () => {
  let testWalletL1EthBalance: BigNumber
  let testWalletL2EthBalance: BigNumber

  before('has expected initial values', async () => {
    testWalletL1EthBalance = await bridge.getAndUpdateL1EthBalance()
    testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
    expect(testWalletL1EthBalance.eq(parseEther(depositAmount))).to.be.true
    expect(testWalletL2EthBalance.eq(Zero)).to.be.true
  })

  const ethToL2DepositAmount = parseEther('0.0001')
  const ethFromL2WithdrawAmount = parseEther('0.00001')

  it('deposit ether transaction succeeds', async () => {
    const inbox = await bridge.l1Bridge.getInbox()
    const initialInboxBalance = await ethProvider.getBalance(inbox.address)
    const res = await bridge.depositETH(ethToL2DepositAmount)
    const rec = await res.wait()

    expect(rec.status).to.equal(1)
    const finalInboxBalance = await ethProvider.getBalance(inbox.address)
    expect(initialInboxBalance.add(ethToL2DepositAmount).eq(finalInboxBalance))
  })
  it('L2 address has expected balance after deposit eth', async () => {
    await wait()
    testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
    expect(testWalletL2EthBalance.eq(ethToL2DepositAmount)).to.be.true
  })
  it('withdraw Ether transaction succeeds and emits event', async () => {
    const withdrawEthRes = await bridge.withdrawETH(ethFromL2WithdrawAmount)
    const withdrawEthRec = await withdrawEthRes.wait()

    expect(withdrawEthRec.status).to.equal(1)
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawEthRec)
    )[0]
    expect(withdrawEventData).to.exist
  })

  it('balance deducted after withdraw eth', async () => {
    wait()
    const bal = await bridge.getAndUpdateL2EthBalance()
    // lazy check, will update once gasPrice is activated
    expect(bal.lt(ethToL2DepositAmount)).to.be.true
  })
})
const tokenDepositAmount = BigNumber.from(50)
const tokenWithdrawAmount = BigNumber.from(2)

const tokenDepositAmountE18 = utils.parseUnits('50', 18)
const tokenWithdrawAmountE18 = utils.parseUnits('2', 18)

describe('ERC20', () => {
  before('create/ensure l1 erc20 w initial supply', async () => {
    wait(10000)
    const testTokenFactory = await new TestERC20__factory(preFundedWallet)
    const testToken = await (async () => {
      if (erc20Address) {
        prettyLog('Using token already deployed at l1 address: ' + erc20Address)
        return testTokenFactory.attach(erc20Address)
      } else {
        prettyLog(
          `No L1 standard ERC20 token address provided; we'll do it live!`
        )
        const res = await testTokenFactory.deploy()
        const rec = await res.deployTransaction.wait()
        expect(rec.status).to.equal(1)
        prettyLog('New token deployed at ' + res.address)
        return res
      }
    })()
    const mintRes = await testToken.mint()
    const mintRec = await mintRes.wait()
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
    const tokenContract = TestERC20__factory.connect(erc20Address, ethProvider)
    const initialBridgeTokenBalance = await tokenContract.balanceOf(
      bridge.ethERC20Bridge.address
    )
    const depositRes = await bridge.deposit(
      erc20Address,
      tokenDepositAmount,
      BigNumber.from(10000000000000),
      BigNumber.from(0),
      undefined,
      { gasLimit: 210000, gasPrice: l1gasPrice }
    )

    const depositRec = await depositRes.wait()
    await wait()

    expect(depositRec.status).to.equal(1)
    const finalBridgeTokenBalance = await tokenContract.balanceOf(
      bridge.ethERC20Bridge.address
    )
    expect(
      initialBridgeTokenBalance
        .add(tokenDepositAmount)
        .eq(finalBridgeTokenBalance)
    )

    const tokenDepositData = (
      await bridge.getDepositTokenEventData(depositRec)
    )[0] as DepositTokenEventResult
    const seqNum = tokenDepositData.seqNum

    const l2RetryableHash = await bridge.calculateL2RetryableTransactionHash(
      seqNum
    )

    const retryableReceipt = await arbProvider.waitForTransaction(
      l2RetryableHash
    )

    expect(retryableReceipt.status).to.equal(1)
  })

  it('L2 wallet has expected balance after erc20 deposit', async () => {
    const l2Data = await bridge.getAndUpdateL2TokenData(erc20Address)

    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance

    expect(testWalletL2Balance && testWalletL2Balance.eq(tokenDepositAmount)).to
      .be.true
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
    expect(balance.eq(tokenDepositAmount)).to.be.true
  })

  it('L1 and L2 implementations of calculateL2ERC20Address match', async () => {
    // this uses the ArbTokenBridge implementation
    const erc20L2AddressAsPerL2 = await bridge.getERC20L2Address(erc20Address)
    const erc20L2AddressAsPerL1 = await bridge.ethERC20Bridge.calculateL2TokenAddress(
      erc20Address
    )
    prettyLog('Token L2 Address: ' + erc20L2AddressAsPerL1)
    expect(erc20L2AddressAsPerL2).to.equal(erc20L2AddressAsPerL1)
  })

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
  })

  it('balance properly deducted after erc20 withdraw', async () => {
    await wait()
    const l2Data = await bridge.getAndUpdateL2TokenData(erc20Address)
    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance
    expect(
      testWalletL2Balance &&
        testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenDepositAmount)
    ).to.be.true
  })
})

describe('CustomToken', () => {
  let l1CustomToken: TestCustomTokenL1
  let l2CustomToken: TestArbCustomToken

  before('sets up a new custom token, L1 and L2', async () => {
    if (!existentCustomTokenL1 && !existentCustomTokenL2) {
      prettyLog("No custom token addresses given; we'll do it live!")
      const customTokenFactory = await new TestCustomTokenL1__factory(
        preFundedWallet
      )
      l1CustomToken = await customTokenFactory.deploy(
        bridge.ethERC20Bridge.address
      )
      let rec = await l1CustomToken.deployTransaction.wait()
      expect(rec.status).to.equal(1)
      prettyLog('Deployed a new customL1 Token at ' + l1CustomToken.address)

      const customL2TokenFactory = await new TestArbCustomToken__factory(
        l2TestWallet
      )
      l2CustomToken = await customL2TokenFactory.deploy(
        bridge.arbTokenBridge.address,
        l1CustomToken.address
      )
      rec = await l2CustomToken.deployTransaction.wait()
      expect(rec.status).to.equal(1)
      prettyLog('Deployed a new custom L2 token at ' + l2CustomToken.address)

      const registerRes = await l1CustomToken.registerTokenOnL2(
        l2CustomToken.address,
        Zero,
        BigNumber.from(10000000000000),
        Zero,
        l1TestWallet.address
      )
      const registerRec = await registerRes.wait()
      expect(registerRec.status).to.equal(1)

      const eventData = (
        await BridgeHelper.getActivateCustomTokenEventResult(
          registerRec,
          bridge.ethERC20Bridge.address
        )
      )[0]

      expect(eventData).to.exist

      const { seqNum } = eventData

      const l2RetryableHash = await bridge.calculateL2RetryableTransactionHash(
        seqNum
      )

      const retryableReceipt = await arbProvider.waitForTransaction(
        l2RetryableHash
      )

      expect(retryableReceipt.status).to.equal(1)

      wait()
      const l2AddressHopefully = await bridge.arbTokenBridge.customL2Token(
        l1CustomToken.address
      )
      expect(l2AddressHopefully).to.equal(l2CustomToken.address)
    } else {
      prettyLog(
        "Connecting to pre-deployed custom tokens and ensuring they're property registered:"
      )
      l1CustomToken = TestCustomTokenL1__factory.connect(
        existentCustomTokenL1,
        preFundedWallet
      )
      l2CustomToken = TestArbCustomToken__factory.connect(
        existentCustomTokenL2,
        l2TestWallet
      )

      const l2CustomTokenAddressInEthBridge = await bridge.ethERC20Bridge.customL2Token(
        existentCustomTokenL1
      )

      expect(l2CustomTokenAddressInEthBridge).to.equal(existentCustomTokenL2)
      prettyLog(
        `Connected to pre-deployed, pre-registered custom token addresses. L1:${l1CustomToken.address} L2: ${l2CustomToken.address}`
      )
    }
  })

  it('setup: mint some custom token and send to test address', async () => {
    const mintRes = await l1CustomToken.mint()
    const mintRec = await mintRes.wait()
    const bal = await l1CustomToken.balanceOf(preFundedWallet.address)
    expect(bal.gt(BigNumber.from(40000000))).to.be.true

    erc20Address = l1CustomToken.address
    const res = await l1CustomToken.transfer(
      l1TestWallet.address,
      BigNumber.from(200)
    )
    const rec = await res.wait()
    const data = await bridge.getAndUpdateL1TokenData(erc20Address)
    const testWalletBal = data.ERC20 && data.ERC20.balance
    expect(testWalletBal && testWalletBal.eq(BigNumber.from(200))).to.be.true
  })

  it('approve token for bridge contract', async () => {
    const approveRes = await bridge.approveToken(l1CustomToken.address)
    const approveRec = await approveRes.wait()
    expect(approveRec.status).to.equal(1)

    const data = await bridge.getAndUpdateL1TokenData(l1CustomToken.address)
    const allowed = data.ERC20 && data.ERC20.allowed
    expect(allowed).to.be.true
  })
  it('deposits custom token', async () => {
    const tokenContract = TestCustomTokenL1__factory.connect(
      erc20Address,
      ethProvider
    )
    const initialBridgeTokenBalance = await tokenContract.balanceOf(
      bridge.ethERC20Bridge.address
    )
    const depositRes = await bridge.deposit(
      l1CustomToken.address,
      tokenDepositAmount,
      BigNumber.from(10000000000000),
      BigNumber.from(0),
      '',
      { gasLimit: 210000, gasPrice: l1gasPrice }
    )

    const depositRec = await depositRes.wait()
    expect(depositRec.status).to.equal(1)
    const finalBridgeTokenBalance = await tokenContract.balanceOf(
      bridge.ethERC20Bridge.address
    )
    expect(
      initialBridgeTokenBalance
        .add(tokenDepositAmount)
        .eq(finalBridgeTokenBalance)
    )
    await wait(10000)

    const tokenDepositData = (
      await bridge.getDepositTokenEventData(depositRec)
    )[0] as DepositTokenEventResult

    const seqNum = tokenDepositData.seqNum

    const l2RetryableHash = await bridge.calculateL2RetryableTransactionHash(
      seqNum
    )

    const retryableReceipt = await arbProvider.waitForTransaction(
      l2RetryableHash
    )

    expect(retryableReceipt.status).to.equal(1)
  })

  it('wallet has expected balance after custom token deposit', async () => {
    const data = await bridge.getAndUpdateL2TokenData(l1CustomToken.address)
    const customTokenData = data && data.CUSTOM
    expect(customTokenData && customTokenData.balance.eq(tokenDepositAmount))
  })

  it('withdraw custom token succeeds and emits event data', async () => {
    const withdrawRes = await l2CustomToken.withdraw(
      l1TestWallet.address,
      tokenWithdrawAmount
    )
    const withdrawRec = await withdrawRes.wait()
    expect(withdrawRec.status).to.equal(1)
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawRec)
    )[0]
    expect(withdrawEventData).to.exist
  })

  it('balance properly deducted after custom withdraw', async () => {
    await wait()
    const l2Data = await bridge.getAndUpdateL2TokenData(l1CustomToken.address)
    const testWalletL2Balance = l2Data && l2Data.CUSTOM && l2Data.CUSTOM.balance
    expect(
      testWalletL2Balance &&
        testWalletL2Balance.add(tokenWithdrawAmount).eq(tokenDepositAmount)
    ).to.be.true
  })
})

describe('CustomToken: no-L2-yet-fallback case', () => {
  let l1CustomToken: TestCustomTokenL1
  before(
    'deploys a new custom token, mints, approves, and registered the L2 side as some rando-address',
    async () => {
      prettyLog('deploying a new custom token')

      const customTokenFactory = await new TestCustomTokenL1__factory(
        l1TestWallet
      )
      l1CustomToken = await customTokenFactory.deploy(
        bridge.ethERC20Bridge.address
      )
      const rec = await l1CustomToken.deployTransaction.wait()
      expect(rec.status).to.equal(1)
      prettyLog('Deployed a new customL1 Token at ' + l1CustomToken.address)

      const approveRes = await bridge.approveToken(l1CustomToken.address)
      const approveRec = await approveRes.wait()
      expect(approveRec.status).to.equal(1)

      const data = await bridge.getAndUpdateL1TokenData(l1CustomToken.address)
      const allowed = data.ERC20 && data.ERC20.allowed
      expect(allowed).to.be.true

      const mintRes = await l1CustomToken.mint()
      const mintRec = await mintRes.wait()

      const randoAddress = Wallet.createRandom().address

      const registerRes = await l1CustomToken.registerTokenOnL2(
        randoAddress,
        Zero,
        BigNumber.from(10000000000000),
        Zero,
        l1TestWallet.address
      )
      const registerRec = await registerRes.wait()
      expect(registerRec.status).to.equal(1)

      const eventData = (
        await BridgeHelper.getActivateCustomTokenEventResult(
          registerRec,
          bridge.ethERC20Bridge.address
        )
      )[0]

      expect(eventData).to.exist

      const { seqNum } = eventData

      const l2RetryableHash = await bridge.calculateL2RetryableTransactionHash(
        seqNum
      )

      const retryableReceipt = await arbProvider.waitForTransaction(
        l2RetryableHash
      )

      expect(retryableReceipt.status).to.equal(1)

      wait()
    }
  )

  it('did not register at rando address', async () => {
    const customTokenAddr = await bridge.arbTokenBridge.customL2Token(
      l1CustomToken.address
    )
    const customTokenCode = await bridge.l2Provider.getCode(customTokenAddr)
    // no custom token deployed
    expect(customTokenCode).to.equal('0x')

    const erc20L2Address = await bridge.getERC20L2Address(l1CustomToken.address)

    const prevErc20Code = await bridge.l2Provider.getCode(erc20L2Address)
    expect(prevErc20Code).to.equal('0x')
  })

  it('deposit into TMT works', async () => {
    const depositRes = await bridge.deposit(
      l1CustomToken.address,
      tokenDepositAmount,
      BigNumber.from(10000000000000),
      BigNumber.from(0),
      undefined,
      { gasLimit: 210000, gasPrice: l1gasPrice }
    )
    const depositRec = await depositRes.wait()
    await wait()

    const tokenDepositData = (
      await bridge.getDepositTokenEventData(depositRec)
    )[0] as DepositTokenEventResult
    const seqNum = tokenDepositData.seqNum

    const l2RetryableHash = await bridge.calculateL2RetryableTransactionHash(
      seqNum
    )

    const retryableReceipt = await arbProvider.waitForTransaction(
      l2RetryableHash
    )

    expect(depositRec.status).to.equal(1)

    expect(retryableReceipt.status).to.equal(1)

    const erc20L2Address = await bridge.arbTokenBridge.calculateL2ERC20TokenAddress(
      l1CustomToken.address
    )

    const post = await bridge.l2Provider.getCode(erc20L2Address)
    expect(post).not.to.equal('0x')

    const arbERC20 = StandardArbERC20__factory.connect(
      erc20L2Address,
      arbProvider
    )
    const symbol = await arbERC20.symbol()
    expect(symbol).to.equal('TMT')

    const newCustomTokenBal = await arbERC20.balanceOf(l1TestWallet.address)

    expect(newCustomTokenBal.eq(tokenDepositAmount)).to.be.true
    // user should be able to withdraw
    // or migrate to custom token once deployed
  })
})

describe.skip('scrap paper', async () => {
  it('', () => {
    console.log('just for scraps')
  })
})
