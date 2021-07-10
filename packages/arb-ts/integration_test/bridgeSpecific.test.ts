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

import { expect } from 'chai'
import config from './config'
import { TestERC20__factory } from '../src/lib/abi/factories/TestERC20__factory'
import yargs from 'yargs/yargs'
import chalk from 'chalk'
import {
  fundL1,
  fundL2,
  testRetryableTicket,
  testRetryableTicketNoAutoRedeem,
  depositFunc,
  generateRandomWallet,
  prettyLog,
  warn,
  instantiateBridgeWithRandomWallet,
  fundL2Token,
  tokenFundAmount,
  skipIfMainnet,
  existentTestERC20,
  retryableTicketReceipt,
} from './testHelpers'
import { Inbox } from '../src/lib/abi/Inbox'
const { Zero, AddressZero } = constants

describe('Bridge Specific', () => {
  beforeEach('skipIfMainnet', function () {
    skipIfMainnet(this)
  })

  it('should not be able to redeem after the retryable has been cancelled', async () => {
    const { bridge } = await instantiateBridgeWithRandomWallet()

    await fundL1(bridge)
    await fundL2(bridge)
    //console.log("here before deposit token test")
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
    console.log('here before deposit Res')
    const depositRes = await bridge.deposit(
      existentTestERC20,
      tokenDepositAmount,
      { maxGas: BigNumber.from(0) }
    )

    const depositRec = await depositRes.wait()
    console.log('here after deposit Res')

    const finalBridgeTokenBalance = await testToken.balanceOf(
      expectedL1GatewayAddress
    )

    expect(
      initialBridgeTokenBalance
        .add(tokenDepositAmount)
        .eq(finalBridgeTokenBalance)
    ).to.be.true

    //console.log("test retryable - not autoredeem")
    const cancelTx = await bridge.cancelRetryableTicket(depositRec)
    console.log('cancelled ticket')
    const retryableTicket = await testRetryableTicketNoAutoRedeem(
      bridge,
      depositRec
    )
    console.log('here')
    //const bridgeTxn  = await depositFunc(bridge, {maxGas: BigNumber.from(0)}, false) //no autoredeem
    //console.log("here after deposit token test")

    //CANCEL ticket (need to fund the l2 side)

    //const cancelTx = await bridge.cancelRetryableTicket(bridgeTxn.depositRec)

    //revert on attempt to redeem
    const redeemTxs = await bridge.redeemRetryableTicket(retryableTicket)
    //const redeemTxs = await bridge.redeemRetryableTicket(bridgeTxn.retryableTicket)
    const reedemTx = await redeemTxs.wait()
    expect(reedemTx.status).to.equal(0)
  })

  it('should not have anything in the retry buffer after cancelling', async () => {
    const { bridge } = await instantiateBridgeWithRandomWallet()
    await fundL1(bridge)
    const depositTx = await depositFunc(
      bridge,
      { maxGas: BigNumber.from(0) },
      false
    ) //no autoredeem

    //check retry buffer for ticket
    const timeoutBefore = await bridge.l2Bridge.arbRetryableTx.getTimeout(
      depositTx.retryableTicket
    )
    expect(timeoutBefore).to.not.equal(0) //retry buffer should not empty

    //CANCEL ticket
    await bridge.cancelRetryableTicket(depositTx.depositRec)

    //check retry buffer for ticket
    const timeout = await bridge.l2Bridge.arbRetryableTx.getTimeout(
      depositTx.retryableTicket
    )
    expect(timeout).to.equal(0) //returns 0 retry buffer is empty
  })

  it('beneficiary should have call value in funds after cancellation', async () => {
    const { bridge } = await instantiateBridgeWithRandomWallet()
    await fundL1(bridge)
    const bridgeTxn = await depositFunc(
      bridge,
      { maxGas: BigNumber.from(0) },
      false
    ) //no autoredeem

    //CANCEL
    await bridge.cancelRetryableTicket(bridgeTxn.depositRec)

    const beneficiaryAddr = await bridge.l2Bridge.arbRetryableTx.getBeneficiary(
      bridgeTxn.retryableTicket
    )
    const beneficiaryBalance = bridge.l2Provider.getBalance(beneficiaryAddr)

    expect(beneficiaryBalance).to.equal(0.001) //check against l1 fund amount
  })

  /*it.skip('cannot cancel twice', async () => {
          const { bridge } = await instantiateBridgeWithRandomWallet() 
          await fundL1(bridge) 
          await fundL2(bridge)
          const depositTx = await depositFunc(bridge, {maxGas: BigNumber.from(0)}, false) //no autoredeem    
   
          //CANCEL
          const cancelTx = await bridge.cancelRetryableTicket(depositTx.retryableTicket) //should work
          //cannot cancel again
          try {
            const cancelTxAgain = await bridge.cancelRetryableTicket() //id
          } catch (err) {
            console.log("cannot cancel tx twice!") 
          }   
          
        })  */

  /*it('appropriate amount of funds are returned to the credit back address', async () => { 
          const { bridge } = await instantiateBridgeWithRandomWallet() 
          //create custom retryable to do this 
          await bridge.l1Bridge.getInbox()
          await fundL1(bridge)    
          const retryableTicketReceipt = await depositTokenTest(bridge) 
      
          ##################### [SEE DEPOSIT FOR A WORKING EXAMPLE OF THIS]
           const tokenContract = TestERC20__factory.connect(erc20Address, ethProvider) //l1 token contract
          const erc20L2Address = await bridge.getERC20L2Address(erc20Address)
          //l2 token address 
          const arbERC20 = StandardArbERC20__factory.connect(
            erc20L2Address,
            arbProvider
          )  
          const maxSubmissionCost = 100      
          const data = ethers.utils.defaultAbiCoder.encode(
            ['uint256', 'bytes'],
            [maxSubmissionCost, '0x']
          ) 
      
          //params for retryable ticket 
          const l2CallValue = BigNumber.from(.1)
          const inbox = await bridge.l1Bridge.getInbox() 
          const maxGas = 1000000000  
       
          const gasPrice = bridge.l2Bridge.getTxnSubmissionPrice(data)  //getSubmissionPrie 
          const destAddr = erc20L2Address
          const creditBackAddr = '0x'  
          const beneficiaryAddr = '0x'
            //maxSubmissionCost + callvalue 
            //deposit here 
          const txid = inbox['createRetryableTicket(address,uint256,uint256,address,address,uint256,uint256,bytes)'](
            destAddr, 
            l2CallValue, 
            maxSubmissionCost, 
            creditBackAddr, 
            beneficiaryAddr, 
            maxGas, 
            gasPrice, 
            data, 
            {value: } )
      
            const txId = 'placeholder'
            const beneficiaryBridgeAddr = await bridge.l2Bridge.arbRetryableTx.getBeneficiary(txId) 
            expect(beneficiaryAddr).to.equal(beneficiaryBridgeAddr) //make sure addreses are equal 
            
            await bridge.cancelRetryableTicket(retryableTx) //CANCEL
            expect(await tokenContract.balanceOf(beneficiaryAddr)).to.equal(l2CallValue)
            ##################### 
          bridge.
        })*/
})
