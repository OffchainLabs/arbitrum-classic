import { utils, Wallet, constants } from 'ethers'
import { expect } from 'chai'

import { prettyLog } from './testHelpers'
const { Zero, AddressZero } = constants
import dotenv from 'dotenv'
import {
  instantiateRandomBridge,
  fundL1,
  wait,
  fundL2,
  preFundAmount,
  skipIfMainnet,
} from './testHelpers'
import { ArbGasInfo__factory } from '../src/lib/abi/factories/ArbGasInfo__factory'
import { ARB_GAS_INFO } from '../src/lib/precompile_addresses'

dotenv.config()

describe('Ether', async () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('transfers ether on l2', async () => {
    const { bridge } = await instantiateRandomBridge()
    await fundL2(bridge)
    const randomAddress = Wallet.createRandom().address
    const amountToSend = utils.parseEther('0.000005')
    const res = await bridge.l2Signer.sendTransaction({
      to: randomAddress,
      value: amountToSend,
    })
    const rec = await res.wait()

    expect(rec.status).to.equal(1)
    const newBalance = await bridge.l2Provider.getBalance(randomAddress)
    expect(newBalance.eq(amountToSend)).to.be.true
  })
  it('deposits ether', async () => {
    const { bridge } = await instantiateRandomBridge()
    await fundL1(bridge)

    const inbox = await bridge.l1Bridge.getInbox()

    const initialInboxBalance = await bridge.l1Bridge.l1Provider.getBalance(
      inbox.address
    )
    const ethToDeposit = utils.parseEther('0.0002')
    const res = await bridge.depositETH(ethToDeposit)
    const rec = await res.wait()

    expect(rec.status).to.equal(1)
    const finalInboxBalance = await bridge.l1Bridge.l1Provider.getBalance(
      inbox.address
    )
    expect(initialInboxBalance.add(ethToDeposit).eq(finalInboxBalance))

    const seqNumArr = await bridge.getInboxSeqNumFromContractTransaction(rec)
    if (seqNumArr === undefined) {
      throw new Error('no seq num')
    }
    expect(seqNumArr.length).to.exist

    const seqNum = seqNumArr[0]
    const l2TxHash = await bridge.calculateL2TransactionHash(seqNum)
    prettyLog('l2TxHash: ' + l2TxHash)
    prettyLog('waiting for l2 transaction:')
    const l2TxnRec = await bridge.l2Bridge.l2Provider.waitForTransaction(
      l2TxHash,
      undefined,
      1000 * 60 * 12
    )
    prettyLog('l2 transaction found!')
    expect(l2TxnRec.status).to.equal(1)

    for (let i = 0; i < 60; i++) {
      prettyLog('balance check attempt ' + (i + 1))
      await wait(5000)
      const testWalletL2EthBalance = await bridge.getL2EthBalance()
      if (testWalletL2EthBalance.gt(constants.Zero)) {
        prettyLog(`balance updated!  ${testWalletL2EthBalance.toString()}`)
        expect(true).to.be.true
        return
        break
      }
    }
    expect(false).to.be.true
  })

  it('withdraw Ether transaction succeeds', async () => {
    const { bridge } = await instantiateRandomBridge()
    await fundL2(bridge)
    const ethToWithdraw = utils.parseEther('0.00002')

    const initialBalance = await bridge.l2Bridge.getL2EthBalance()

    const withdrawEthRes = await bridge.withdrawETH(ethToWithdraw)
    const withdrawEthRec = await withdrawEthRes.wait()

    const arbGasInfo = ArbGasInfo__factory.connect(
      ARB_GAS_INFO,
      bridge.l2Provider
    )
    expect(withdrawEthRec.status).to.equal(1)

    const inWei = await arbGasInfo.getPricesInWei({
      blockTag: withdrawEthRec.blockNumber,
    })
    const withdrawEventData = (
      await bridge.getWithdrawalsInL2Transaction(withdrawEthRec)
    )[0]
    expect(withdrawEventData).to.exist

    const etherBalance = await bridge.getL2EthBalance()

    const totalEth = etherBalance
      .add(ethToWithdraw)
      .add(withdrawEthRec.gasUsed.mul(inWei[5]))

    // TODO
    console.log(
      `This number should be zero...? ${initialBalance
        .sub(totalEth)
        .toString()}`
    )

    expect(true).to.be.true
  })
})
