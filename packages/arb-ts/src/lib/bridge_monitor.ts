import { Bridge } from './bridge'
import { Signer, constants } from 'ethers'
import { ArbOwner__factory } from './abi/factories/ArbOwner__factory'
import { ARB_OWNER } from './precompile_addresses'
import { OutgoingMessageState } from './bridge_helpers'
import { ERC20__factory } from './abi/factories/ERC20__factory'
// TODO: set block number everywhere....
export class BridgeMonitor extends Bridge {
  static async init(
    ethSigner: Signer,
    arbSigner: Signer,
    l1GatewayRouterAddress?: string,
    l2GatewayRouterAddress?: string
  ) {
    const { l1BridgeObj, l2BridgeObj } = await Bridge._createBridges(
      ethSigner,
      arbSigner,
      l1GatewayRouterAddress,
      l2GatewayRouterAddress
    )
    return new BridgeMonitor(l1BridgeObj, l2BridgeObj)
  }

  public confirmEthSupply = async () => {
    const ethL2 = await this.l2EthSupply()

    const ethL1 = await this.l1EthSupply()
    const ethIn = await this.totalIncomingEth()

    const ethOut = await this.totalOutgoingEth()

    const accountedForEth = ethL2.add(ethOut).add(ethIn)

    return ethL1.eq(accountedForEth)
  }

  public l1EthSupply = async (l1BlockNumber?: number) => {
    const bridge = await this.getBridgeContract()
    return this.l1Provider.getBalance(bridge.address, l1BlockNumber)
  }

  public l2EthSupply = async (l2BlockNumber?: number) => {
    const arbOwner = ArbOwner__factory.connect(ARB_OWNER, this.l2Provider)
    // TODO: does this include pending incoming eth?
    return arbOwner.getTotalOfEthBalances()
  }

  public totalIncomingEth = async () => {
    const allRetryables = await this.getRetryablesL1()
    const l1TxnRecs = await Promise.all(
      allRetryables.map(retryLog =>
        this.l1Provider.getTransactionReceipt(retryLog.txHash)
      )
    )

    const l1TxnCallvalueArr = (
      await Promise.all(
        l1TxnRecs.map(rec =>
          this.l1Provider.getTransaction(rec.transactionHash)
        )
      )
    ).map(res => res.value)

    const l2RetryRedeemRecs = await Promise.all(
      l1TxnRecs.map(rec => this.getL2TxnFromL1Txn(rec))
    )

    // sum up the call value from unredeemed retryables
    //TODO can  the redeem receipts ever get emitted with a status code of zero?
    // TODO: this doesn't cover the callvalue for unredeemed retryables

    return l2RetryRedeemRecs.reduce((acc, redeemRec, i) => {
      if (!redeemRec) {
        return acc.add(l1TxnCallvalueArr[i])
      }
      return acc
    }, constants.Zero)
  }

  public totalOutgoingEth = async () => {
    const allEthWithdrawals = (await this.getEthWithdrawals()).filter(event =>
      event.callvalue.gt(constants.AddressZero)
    )

    const queries = allEthWithdrawals.map(l2L1Event => {
      return this.getOutGoingMessageState(
        l2L1Event.batchNumber,
        l2L1Event.indexInBatch
      )
    })

    const messageStates = await Promise.all(queries)

    // get sum all unclaimed ether value
    return messageStates.reduce((acc, messageState, i) => {
      if (messageState !== OutgoingMessageState.EXECUTED) {
        return acc.add(allEthWithdrawals[i].callvalue)
      }
      return acc
    }, constants.Zero)
  }

  public l1TokenSupply = async (tokenAddress: string) => {
    const token = await ERC20__factory.connect(tokenAddress, this.l1Provider)
    const gatewayAddress = await this.l1Bridge.getGatewayAddress(tokenAddress)
    return token.balanceOf(gatewayAddress)
  }

  public l2TokenSupply = async (tokenAddress: string) => {
    const l2Address = await this.l1Bridge.getERC20L2Address(tokenAddress)
    const token = await ERC20__factory.connect(l2Address, this.l2Provider)
    return token.totalSupply()
  }

  //   totalIncomingTokenBalance(tokenAddress: string) {}

  //   totalOutgoingTokenBalance(tokenAddress: string) {}
}
