import { Bridge } from '../src/lib/bridge'
import { expect } from 'chai'
import { BigNumber, ContractReceipt, Wallet } from 'ethers'
import chalk from 'chalk'
import { instantiateBridge } from '../scripts/instantiate_bridge'
import { utils } from 'ethers'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'
import yargs from 'yargs/yargs'
import config from './config'

const argv = yargs(process.argv.slice(2)).argv
let networkID = argv.networkID as string

networkID = networkID || '4'
if (!config[networkID]) {
  throw new Error('network not supported')
}

const {
  existentTestERC20: _existentTestERC20,
  existentTestCustomToken: _existentTestCustomToken,
} = config[networkID]

export const existentTestERC20 = _existentTestERC20 as string
export const existentTestCustomToken = _existentTestCustomToken as string

export const preFundAmount = utils.parseEther('0.001')

export const retryableTicketReceipt = async (
  bridge: Bridge, 
  retryableTicket: string
) => {
  prettyLog('Waiting for retryable ticket')

  const retryableTicketReceipt = await bridge.l2Bridge.l2Provider.waitForTransaction(
    retryableTicket,
    undefined,
    1000 * 60 * 15
  )

  prettyLog('retryableTicketReceipt found:')

  expect(retryableTicketReceipt.status).to.equal(1) 

  return retryableTicketReceipt
}

export const getRetryableTicket = async( 
  bridge: Bridge, 
  rec: ContractReceipt,
  seqNum: BigNumber
) => {

  prettyLog('Getting the receipt for the retryable ticket')

  const retryableTicket = await bridge.calculateL2TransactionHash(seqNum)
  
  prettyLog(
    `retryableTicket: ${retryableTicket}`
  )
  
  return retryableTicket
}

export const autoRedeemReceipt = async (
  bridge: Bridge, 
  seqNum: BigNumber 
) => {
  
  prettyLog(`Waiting for auto redeem transaction (this shouldn't take long`)
  const autoRedeem = await bridge.calculateRetryableAutoRedeemTxnHash(seqNum) 
  const redeemTransaction = await bridge.calculateL2RetryableTransactionHash(
    seqNum
  ) 

  prettyLog(`autoredeem: ${autoRedeem}, redeem: ${redeemTransaction}`)

  const autoRedeemReceipt = await bridge.l2Bridge.l2Provider.waitForTransaction(
    autoRedeem,
    undefined,
    1000 * 60
  )
  prettyLog('autoRedeem receipt found!')

  expect(autoRedeemReceipt.status).to.equal(1)
  prettyLog('Getting redemption')
  const redemptionReceipt = await bridge.l2Bridge.l2Provider.getTransactionReceipt(
    redeemTransaction
  )

  expect(redemptionReceipt && redemptionReceipt.status).equals(1) 

  return redemptionReceipt
}

export const testRetryableTicketNoAutoRedeem = async (
  bridge: Bridge,
  rec: ContractReceipt
) => {
  prettyLog(`testing retryable for ${rec.transactionHash}`) 

  const seqNums = await bridge.getInboxSeqNumFromContractTransaction(rec)
  const seqNum = seqNums && seqNums[0]  
  if (!seqNum) {
    throw new Error('Seq num not found')
  } 

  const retryableTicket = await getRetryableTicket(bridge, rec, seqNum)  
  const _retryableTicketReceipt = await retryableTicketReceipt(bridge, retryableTicket)

  return retryableTicket
  
} 

export const testRetryableTicket = async (
  bridge: Bridge,
  rec: ContractReceipt
) => {
  prettyLog(`testing retryable for ${rec.transactionHash}`) 

  const seqNums = await bridge.getInboxSeqNumFromContractTransaction(rec)
  const seqNum = seqNums && seqNums[0] 
  if (!seqNum) {
    throw new Error('Seq num not found')
  }
  const retryableTicket = await getRetryableTicket(bridge, rec, seqNum)   
  const _retryableTicketReceipt = await retryableTicketReceipt(bridge, retryableTicket)
  //auto redeem  
  const redeemReceipt = await autoRedeemReceipt(bridge, seqNum)  

  return retryableTicket

}  

export const prettyLog = (text: string) => {
  console.log(chalk.blue(`    *** ${text}`))
  console.log()
}

export const warn = (text: string) => {
  console.log(chalk.red(`WARNING: ${text}`))
  console.log()
} 

export const generateRandomWallet = () => {
  const testPk = utils.formatBytes32String(Math.random().toString())
  prettyLog(
    `Generated wallet, pk: ${testPk} address: ${new Wallet(testPk).address} `
  ) 
  return testPk
}

export const instantiateBridgeWithRandomWallet = () => { 
  const testPk = generateRandomWallet()  
  return instantiateBridge(testPk)
}

const _preFundedWallet = new Wallet(process.env.DEVNET_PRIVKEY as string)
const _preFundedL2Wallet = new Wallet(process.env.DEVNET_PRIVKEY as string)

console.warn('using prefunded wallet ', _preFundedWallet.address)

export const fundL1 = async (bridge: Bridge) => {
  const testWalletAddress = await bridge.l1Bridge.getWalletAddress()
  const preFundedWallet = _preFundedWallet.connect(bridge.l1Provider) 
  // const balanceWallet = (await preFundedWallet.getBalance()) 
  // console.log(balanceWallet)
  // if (balanceWallet == BigNumber.from(0)) {  
  //   throw new Error("balance in L1 wallet is insufficient")
  // }
  const res = await preFundedWallet.sendTransaction({
    to: testWalletAddress,
    value: preFundAmount,
  }) 
  const rec = await res.wait()
  prettyLog('Funded L1 account')
}
export const fundL2 = async (bridge: Bridge) => {
  const testWalletAddress = await bridge.l2Bridge.getWalletAddress()
  const preFundedL2Wallet = _preFundedL2Wallet.connect(bridge.l2Provider) 
  console.log("inside fund l2") 
  //const balanceWallet = await preFundedL2Wallet.getBalance()
  // if (balanceWallet < BigNumber.from(.001)) { 
  //   console.log("pre-funded wallet balance " + balanceWallet)  
  //   throw new Error ("insufficient balance on l2 to complete tx")
  // }
   
  const res = await preFundedL2Wallet.sendTransaction({
    to: testWalletAddress,
    value: preFundAmount, 
  })
  const rec = await res.wait()
  prettyLog('Funded L2 account')
}

export const tokenFundAmount = BigNumber.from(2)
export const fundL2Token = async (bridge: Bridge, tokenAddress: string) => {
  try {
    const testWalletAddress = await bridge.l2Bridge.getWalletAddress()
    const preFundedL2Wallet = _preFundedL2Wallet.connect(bridge.l2Provider)
    const l2Address = await bridge.getERC20L2Address(tokenAddress)
    const testToken = TestERC20__factory.connect(l2Address, preFundedL2Wallet)

    const res = await testToken.transfer(testWalletAddress, tokenFundAmount)
    const rec = await res.wait()

    const result = rec.status === 1
    result && prettyLog('Funded L2 account w/ tokens')

    return result
  } catch (err) {
    console.warn('err', err)

    return false
  }
}

export const wait = (ms = 0) => {
  return new Promise(res => setTimeout(res, ms))
} 

export const depositFunc = async (bridge: Bridge, retryableTxParams={}, autoRedeem: Boolean = true) => {
  const tokenDepositAmount = BigNumber.from(1)

  const testToken = TestERC20__factory.connect(
    existentTestERC20,
    bridge.l1Signer
  )
  const mintRes = await testToken.mint()
  const mintRec = await mintRes.wait()

  const approveRes = await bridge.approveToken(existentTestERC20)
  const approveRec = await approveRes.wait()

  const data = await bridge.getAndUpdateL1TokenData(existentTestERC20)
  const allowed = data.ERC20 && data.ERC20.allowed
  expect(allowed).to.be.true

  const expectedL1GatewayAddress = await bridge.l1Bridge.getGatewayAddress(
    testToken.address
  )
  const initialBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )
  console.log("here before deposit Res")
  const depositRes = await bridge.deposit(existentTestERC20, tokenDepositAmount, retryableTxParams)
  
 
  const depositRec = await depositRes.wait()
  console.log("here after deposit Res")

  const finalBridgeTokenBalance = await testToken.balanceOf(
    expectedL1GatewayAddress
  )

  expect(
    initialBridgeTokenBalance
      .add(tokenDepositAmount)
      .eq(finalBridgeTokenBalance)
  ).to.be.true  
  
  var retryableTicket; 
  if (autoRedeem) {
    retryableTicket = await testRetryableTicket(bridge, depositRec)   
  } else {
    retryableTicket = await testRetryableTicketNoAutoRedeem(bridge, depositRec)
  } 
 
  const l2Data = await bridge.getAndUpdateL2TokenData(existentTestERC20)

  const testWalletL2Balance = l2Data && l2Data.ERC20 && l2Data.ERC20.balance

  expect(testWalletL2Balance && testWalletL2Balance.eq(tokenDepositAmount)).to
    .be.true 
  // //error handling retryable
  // if (retryableTicket == null) {
  //   throw new Error("cannot get null retryable ticket receipt")
  // } else if (typeof retryableTicket == 'undefined') {
  //   throw new Error("cannot get undefined retryable ticket receipt")
  // } 

  console.log("was able to deposit")
  return { depositRec, retryableTicket }
}




export const skipIfMainnet = (() => {
  let chainId = ''
  return async (testContext: Mocha.Context) => {
    if (!chainId) {
      const { l1Network } = await instantiateBridgeWithRandomWallet()
      chainId = l1Network.chainID
    }
    if (chainId === '1') {
      console.log("You're writing to the chain on mainnet lol stop")
      testContext.skip()
    }
  }
})()
