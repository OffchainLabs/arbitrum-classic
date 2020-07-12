import { ArbProvider } from '../src/lib/provider'
import { ArbClient } from '../src/lib/client'
import { ArbWallet } from '../src/lib/wallet'
import * as ethers from 'ethers'

const ethProviderUrl = 'http://0.0.0.0:7545'
const arbProviderUrl = 'http://0.0.0.0:1235'
const ethereumProvider = new ethers.providers.JsonRpcProvider(ethProviderUrl)
import fetch from 'node-fetch'

const pubAddresss = '0xc7711f36b2C13E00821fFD9EC54B04A60AEfbd1b'
// @ts-ignore: override fetch w/ node fetch for test suite
if (!global.fetch) {
  // @ts-ignore
  global.fetch = fetch
}

describe('arbProvider tests', function () {
  const arbProvider = new ArbProvider(
    arbProviderUrl,
    ethereumProvider,
    '',
    true
  )
  test('instantiates', async function () {
    expect(arbProvider.client instanceof ArbClient).toBe(true)
  })
})

describe('arbWallet tests', function () {
  const arbProvider = new ArbProvider(
    arbProviderUrl,
    ethereumProvider,
    undefined,
    true
  )
  const ethereumWallet = new ethers.Wallet(
    '0x979f020f6f6f71577c09db93ba944c89945f10fade64cfc7eb26137d5816fb76',
    ethereumProvider
  )
  const arbWallet = new ArbWallet(ethereumWallet, arbProvider)

  test('arbWallet instantiates properly', async function () {
    const address = await arbWallet.getAddress()
    expect(address).toBe(pubAddresss)
  })

  let tx: ethers.providers.TransactionReceipt
  let txnHash: string

  test('deposit ETH', async function () {
    jest.setTimeout(30000)
    const initialEthBalance = await ethereumWallet.getBalance()
    const initialArbBalance = await arbProvider.getBalance(pubAddresss)

    const response = await arbWallet.depositETH(pubAddresss, 1000)
    tx = await response.wait()
    txnHash = tx.transactionHash || ''

    expect(txnHash).toBeTruthy()
    const newEthBalance = await ethereumWallet.getBalance()
    const newArbBalance = await arbProvider.getBalance(pubAddresss)

    expect(newEthBalance.add(1000).lt(initialEthBalance)).toBe(true)
    expect(newArbBalance.sub(1000).eq(initialArbBalance)).toBe(true)
  })

  test('gets message result data from hash', async function () {
    const result = await arbProvider.getMessageResult(txnHash)
    expect(result).toBeTruthy()
    // todo: better parsing/checking message
    expect(typeof (result && result.nodeInfo.nodeHash)).toEqual('string')
  })
  test('gets txn receipt from txn hash', async function () {
    const result = await arbProvider.perform('getTransactionReceipt', {
      transactionHash: txnHash,
    })
    result.contractAddress = null
    // todo: is this failing sometimes?
    expect(result).toEqual(tx)
  })

  test('sendTransactionMessage and assertion count', async function () {
    jest.setTimeout(30000)
    let preAssertionCount = await arbProvider.client.getAssertionCount()

    const rec = '0x755449b9901f91deC52DB39AF8c655206C63eD8e'
    const initialArbSenderBalance = await arbProvider.getBalance(pubAddresss)
    const initialRecipientBalance = await arbProvider.getBalance(rec)
    const oldSeq = await arbWallet.generateSeq()

    const response = await arbWallet.sendTransactionMessage(rec, 150, '0xabc')
    tx = await response.wait()
    txnHash = tx.transactionHash || ''
    expect(txnHash).toBeTruthy()

    const newSenderBalance = await arbProvider.getBalance(pubAddresss)
    const newRecipientBalance = await arbProvider.getBalance(rec)

    expect(newSenderBalance.add(150).eq(initialArbSenderBalance)).toBe(true)
    expect(newRecipientBalance.sub(150).eq(initialRecipientBalance)).toBe(true)
    expect((oldSeq || 0) + 1).toBe(arbWallet.seqCache)

    let postAssetionCount = await arbProvider.client.getAssertionCount()
    expect(preAssertionCount + 1).toBe(postAssetionCount)
  })

  test('widthdraw ETH', async function () {
    jest.setTimeout(30000)
    const initialArbBalance = await arbProvider.getBalance(pubAddresss)

    const response = await arbWallet.withdrawEthFromChain(100)
    tx = await response.wait()
    txnHash = tx.transactionHash || ''

    expect(txnHash).toBeTruthy()
    const newArbBalance = await arbProvider.getBalance(pubAddresss)
    expect(newArbBalance.add(100).eq(initialArbBalance)).toBe(true)
  })

  test('gets node location', async function () {
    const nodeLocation = await arbProvider.client.getLatestNodeLocation()
    expect(nodeLocation).toBeTruthy()
    const pendingNodeLocation = await arbProvider.client.getLatestPendingNodeLocation()
    expect(pendingNodeLocation).toBeTruthy()
  })
})
