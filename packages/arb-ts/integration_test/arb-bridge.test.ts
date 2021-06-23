import {
  providers,
  utils,
  Wallet,
  BigNumber,
  constants,
  ethers,
  ContractReceipt,
} from 'ethers'
import { Bridge } from '../src/lib/bridge'
import { Network } from '../src/lib/networks'

import {
  OutboundTransferInitiatedResult,
  BridgeHelper,
  L2ToL1EventResult,
} from '../src/lib/bridge_helpers'
import { expect } from 'chai'
import config from './config'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'
import { Rollup__factory } from '../src/lib/abi/factories/Rollup__factory'
import { StandardArbERC20__factory } from '../src/lib/abi/factories/StandardArbERC20__factory'
import { Outbox__factory } from '../src/lib/abi/factories/Outbox__factory'
import { L2GatewayRouter__factory } from '../src/lib/abi/factories/L2GatewayRouter__factory'
import { AeWETH__factory } from '../src/lib/abi/factories/AeWETH__factory'

import { Inbox__factory } from '../src/lib/abi/factories/Inbox__factory'
import { StandardArbERC20 } from '../src/lib/abi/StandardArbERC20'
import { TestERC20 } from '../src/lib/abi/TestERC20'
import { instantiateBridge } from '../scripts/instantiate_bridge'
import yargs from 'yargs/yargs'
import chalk from 'chalk'
import { testRetryableTicket, prettyLog, warn } from './testHelpers'
const { Zero, AddressZero } = constants
console.log(chalk.green(`Starting Token Bridge Integrations Tests!`))
console.log()

const argv = yargs(process.argv.slice(2)).argv

const network = ((_network?: string | number) => {
  if (!_network) {
    prettyLog('Using default network: rinkArby')
    return 'rinkArby'
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
})(argv.network as string | number | undefined) as 'rinkArby' | 'devnet'

if (network === 'rinkArby' && !process.env.INFURA_KEY) {
  warn('To use rinkArby, set env var INFURA_KEY')
  process.exit(1)
}

const { parseEther } = utils
const {
  ethRPC,
  arbRPC,
  preFundedSignerPK,
  l1GatewayRouterAddress,
  l2GatewayRouterAddress,
  l1gasPrice,
  existentTestERC20,
  defaultWait,
  executeOutGoingMessages,
  outBoxUpdateTimeout,
} = config[network]

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

const outGoingMessages: L2ToL1EventResult[] = []

let bridge: Bridge
before('setup', () => {
  it("'pre-funded wallet' is indeed pre-funded", async () => {
    bridge = await Bridge.init(l1TestWallet, l2TestWallet)
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

const tokenDepositAmount = BigNumber.from(50)
const tokenWithdrawAmount = BigNumber.from(2)

describe.skip('ERC20', () => {
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
    const expectedL1GatewayAddress = await bridge.l1Bridge.getGatewayAddress(
      tokenContract.address
    )
    const initialBridgeTokenBalance = await tokenContract.balanceOf(
      expectedL1GatewayAddress
    )
    prettyLog('depositing')
    const depositRes = await bridge.deposit(
      erc20Address,
      tokenDepositAmount,
      {},
      undefined
    )
    prettyLog('depositing done')

    const depositRec = await depositRes.wait()
    prettyLog('depositing rec hash ' + depositRec.transactionHash)

    await wait()

    expect(depositRec.status).to.equal(1)
    const finalBridgeTokenBalance = await tokenContract.balanceOf(
      expectedL1GatewayAddress
    )
    prettyLog('final bridge bal passed')
    expect(
      initialBridgeTokenBalance
        .add(tokenDepositAmount)
        .eq(finalBridgeTokenBalance)
    )

    const tokenDepositData = (
      await bridge.getDepositTokenEventData(depositRec)
    )[0] as OutboundTransferInitiatedResult
    expect(tokenDepositData).to.exist
    prettyLog('token data emitted')

    const seqNums = await bridge.getInboxSeqNumFromContractTransaction(
      depositRec
    )
    if (seqNums === undefined) {
      throw new Error('no seq num')
    }
    expect(seqNums.length).to.equal(1)

    const seqNum = seqNums[0]
    prettyLog('seqnum' + seqNum)

    const l2RetryableHash = await bridge.calculateL2RetryableTransactionHash(
      seqNum
    )
    prettyLog('l2RetryableHash ' + l2RetryableHash)

    const l2RedeemHash = await bridge.calculateRetryableAutoRedeemTxnHash(
      seqNum
    )
    prettyLog('l2RedeemHash ' + l2RedeemHash)

    const redeemReceipt = await arbProvider.waitForTransaction(
      l2RedeemHash,
      undefined,
      1000 * 60 * 10
    )
    expect(redeemReceipt.status).to.equal(1)
    prettyLog('auto-redeem suceeeded ' + l2RedeemHash)

    const retryableReceipt = await arbProvider.waitForTransaction(
      l2RetryableHash
    )

    expect(retryableReceipt.status).to.equal(1)

    prettyLog('retryable succeeded ' + l2RetryableHash)
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
    // this uses the l2ERC20Gateway implementation
    const erc20L2AddressAsPerL1 = await bridge.getERC20L2Address(erc20Address)
    const l2gr = L2GatewayRouter__factory.connect(
      l2GatewayRouterAddress,
      arbProvider
    )
    const erc20L2AddressAsPerL2 = await l2gr.calculateL2TokenAddress(
      erc20Address
    )

    prettyLog(
      `Token L2 Address: ${erc20L2AddressAsPerL1}, ${erc20L2AddressAsPerL2}`
    )
    expect(erc20L2AddressAsPerL2).to.equal(erc20L2AddressAsPerL1)
  })
})
describe.skip('Ether', () => {
  let testWalletL1EthBalance: BigNumber
  let testWalletL2EthBalance: BigNumber
  let initialTestWalletEth2Balance: BigNumber
  let preWithdrawalL2Balance: BigNumber

  const ethToL2DepositAmount = parseEther('0.0005')
  const ethFromL2WithdrawAmount = parseEther('0.00001')

  it('deposit ether transaction succeeds', async () => {
    initialTestWalletEth2Balance = await bridge.getAndUpdateL2EthBalance()
    const inbox = await bridge.l1Bridge.getInbox()
    const initialInboxBalance = await ethProvider.getBalance(inbox.address)
    const res = await bridge.depositETH(ethToL2DepositAmount)
    const rec = await res.wait()

    expect(rec.status).to.equal(1)
    const finalInboxBalance = await ethProvider.getBalance(inbox.address)
    expect(initialInboxBalance.add(ethToL2DepositAmount).eq(finalInboxBalance))

    const seqNumArr = await bridge.getInboxSeqNumFromContractTransaction(rec)
    if (seqNumArr === undefined) {
      throw new Error('no seq num')
    }
    expect(seqNumArr.length).to.exist

    const seqNum = seqNumArr[0]
    const l2TxHash = await bridge.calculateL2TransactionHash(seqNum)
    prettyLog('l2TxHash: ' + l2TxHash)
    // Note:these will pass once the node is updated to deliver this tx format
    prettyLog('waiting for l2 transaction:')
    const l2TxnRec = await arbProvider.waitForTransaction(
      l2TxHash,
      undefined,
      1000 * 60 * 12
    )
    prettyLog('l2 transaction found!')
    expect(l2TxnRec.status).to.equal(1)
  })
  it('L2 address has expected balance after deposit eth', async () => {
    for (let i = 0; i < 60; i++) {
      prettyLog('balance check attempt ' + (i + 1))
      await wait(5000)
      testWalletL2EthBalance = await bridge.getAndUpdateL2EthBalance()
      if (!initialTestWalletEth2Balance.eq(testWalletL2EthBalance)) {
        prettyLog(
          `balance updated! ${initialTestWalletEth2Balance.toString()} ${testWalletL2EthBalance.toString()}`
        )
        break
      }
    }

    expect(testWalletL2EthBalance.gte(initialTestWalletEth2Balance)).to.be.true
  })
  it('withdraw Ether transaction succeeds and emits event', async () => {
    preWithdrawalL2Balance = await bridge.getAndUpdateL2EthBalance()
    const withdrawEthRes = await bridge.withdrawETH(ethFromL2WithdrawAmount, '')
    const withdrawEthRec = await withdrawEthRes.wait()

    expect(withdrawEthRec.status).to.equal(1)
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawEthRec)
    )[0]
    expect(withdrawEventData).to.exist
    outGoingMessages.push(withdrawEventData)
  })

  it('balance deducted after withdraw eth', async () => {
    wait()
    const bal = await bridge.getAndUpdateL2EthBalance()
    // lazy check, will update once gasPrice is activated
    expect(bal.lt(preWithdrawalL2Balance)).to.be.true
  })
})

describe.skip('ERC20', () => {
  it('2nd pass, prefunded: initial erc20 deposit txns — L1 and L2 — both succeed', async () => {
    const tokenContract = TestERC20__factory.connect(erc20Address, ethProvider)
    const defaultL1GatewayAddress = (await bridge.defaultL1Gateway()).address
    const initialBridgeTokenBalance = await tokenContract.balanceOf(
      defaultL1GatewayAddress
    )
    const depositRes = await bridge.deposit(
      erc20Address,
      tokenDepositAmount,
      {},
      undefined,
      { gasLimit: 210000, gasPrice: l1gasPrice }
    )

    const depositRec = await depositRes.wait()

    await wait()

    expect(depositRec.status).to.equal(1)
    const finalBridgeTokenBalance = await tokenContract.balanceOf(
      defaultL1GatewayAddress
    )
    expect(
      initialBridgeTokenBalance
        .add(tokenDepositAmount)
        .eq(finalBridgeTokenBalance)
    )

    const tokenDepositData = (
      await bridge.getDepositTokenEventData(depositRec)
    )[0] as OutboundTransferInitiatedResult
    expect(tokenDepositData).to.exist

    const seqNums = await bridge.getInboxSeqNumFromContractTransaction(
      depositRec
    )
    if (seqNums === undefined) {
      throw new Error('no seq num')
    }
    expect(seqNums.length).to.equal(1)

    const seqNum = seqNums[0]

    const l2RetryableHash = await bridge.calculateL2RetryableTransactionHash(
      seqNum
    )

    const l2RedeemHash = await bridge.calculateRetryableAutoRedeemTxnHash(
      seqNum
    )
    const redeemReceipt = await arbProvider.waitForTransaction(l2RedeemHash)
    expect(redeemReceipt.status).to.equal(1)
    prettyLog('auto-redeem succeeded ' + l2RedeemHash)

    const retryableReceipt = await arbProvider.waitForTransaction(
      l2RetryableHash
    )

    expect(retryableReceipt.status).to.equal(1)

    prettyLog('retryable succeeded ' + l2RetryableHash)
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
    outGoingMessages.push(withdrawEventData)
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

describe('WETH', async () => {
  const depositWETHAmount = '0.0005'
  const withdrawWETHAmount = '0.00005'
  const l1WethAddress = '0xc778417E063141139Fce010982780140Aa0cD5Ab'

  it('setup: wrap some ether', async () => {
    const l1WETH = AeWETH__factory.connect(
      l1WethAddress,
      bridge.l1Bridge.l1Signer
    )
    const res = await l1WETH.deposit({
      value: utils.parseEther('0.001'),
    })
    const rec = await res.wait()
    prettyLog('wrapped some ether')
    expect(rec.status).to.equal(1)
  })

  it('setup: approve WETH gateway', async () => {
    const approveRes = await bridge.approveToken(l1WethAddress)
    const approveRec = await approveRes.wait()
    expect(approveRec.status).to.equal(1)

    const data = await bridge.getAndUpdateL1TokenData(l1WethAddress)
    const allowed = data.ERC20 && data.ERC20.allowed
    expect(allowed).to.be.true
  })

  it('deposits WETH', async () => {
    const res = await bridge.deposit(
      l1WethAddress,
      utils.parseEther(depositWETHAmount),
      {},
      undefined,
      { gasLimit: 230000 }
    )
    const rec = await res.wait()
    prettyLog(
      `Deposit WETH: L1 tx succeeded ${rec.transactionHash}; now handling l2 retryable`
    )
    expect(rec.status).to.equal(1)
    await testRetryableTicket(bridge, rec)

    it('L2 wallet has expected balance after erc20 deposit', async () => {
      const l2Data = await bridge.getAndUpdateL2TokenData(l1WethAddress)

      const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance

      expect(testWalletL2Balance && testWalletL2Balance.eq(tokenDepositAmount))
        .to.be.true
    })
  })

  it('withdraw WETH succeeds and emits event data', async () => {
    const withdrawRes = await bridge.withdrawERC20(
      l1WethAddress,
      utils.parseEther(withdrawWETHAmount)
    )
    const withdrawRec = await withdrawRes.wait()
    expect(withdrawRec.status).to.equal(1)
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawRec)
    )[0]

    expect(withdrawEventData).to.exist
    outGoingMessages.push(withdrawEventData)
  })

  it('balance properly deducted after weth withdraw', async () => {
    await wait()
    const l2Data = await bridge.getAndUpdateL2TokenData(l1WethAddress)
    const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance
    expect(
      testWalletL2Balance &&
        testWalletL2Balance.add(withdrawWETHAmount).eq(depositWETHAmount)
    ).to.be.true
  })
})

describe.skip('trigger outgoing messages', async () => {
  if (!executeOutGoingMessages) {
    return
  }
  const outboxAddress = await bridge.getOutboxAddress()
  const outbox = Outbox__factory.connect(outboxAddress, ethProvider)

  const targetBatchNumber = outGoingMessages.reduce(
    (highestBatchNumber, message) => {
      return message.batchNumber.gte(highestBatchNumber)
        ? message.batchNumber
        : highestBatchNumber
    },
    BigNumber.from(0)
  )
  prettyLog(
    `I will wait a total of ${
      outBoxUpdateTimeout / 1000
    } seconds for the outbox entry to arrive:`
  )

  for (let i = 0; i < 10; i++) {
    const len = await outbox.outboxesLength()
    if (len.lte(targetBatchNumber)) {
      wait(outBoxUpdateTimeout / 10)
      prettyLog(`not yet...`)
    } else {
      prettyLog(`outbox entry created! executing:`)
      for (const outgoingMessage of outGoingMessages) {
        const receipt = await bridge.triggerL2ToL1Transaction(
          outgoingMessage.batchNumber,
          outgoingMessage.indexInBatch
        )
        expect(receipt.status).to.equal(1)
      }
    }
  }
})
