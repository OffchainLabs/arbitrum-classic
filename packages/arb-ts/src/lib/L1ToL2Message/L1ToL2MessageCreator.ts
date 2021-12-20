import { Signer } from '@ethersproject/abstract-signer'
import { Provider } from '@ethersproject/abstract-provider'

import {
  L1ToL2MessageGasEstimator,
  L1toL2MessageGasValues,
} from './L1ToL2MessageGasEstimator'
import { L1ToL2Message } from './L1ToL2Message'
import { Inbox__factory } from '../abi/factories/Inbox__factory'
import { NetworksInfo, CustomNetworks } from '../utils/networks'
import { ContractReceipt, PayableOverrides } from '@ethersproject/contracts'
import { BigNumber } from 'ethers'

interface CreateRetryableTicketOpptions {
  excessFeeRefundAddress?: string
  callValueRefundAddress?: string
}
export class L1ToL2MessageCreator {
  sender?: string
  constructor(
    public readonly l1Signer: Signer,
    public readonly networksInfo: NetworksInfo
  ) {}
  static async init(
    l1Signer: Signer,
    l2ProviderOrChainID: Provider | number,
    customNetworks?: CustomNetworks
  ): Promise<L1ToL2MessageCreator> {
    const networksInfo = await (async () => {
      if (typeof l2ProviderOrChainID === 'number') {
        return NetworksInfo.initFromL2ChainID(
          l2ProviderOrChainID,
          customNetworks
        )
      } else {
        return NetworksInfo.initFromL2Provider(
          l2ProviderOrChainID,
          customNetworks
        )
      }
    })()
    return new L1ToL2MessageCreator(l1Signer, networksInfo)
  }

  public async createRetryableTicketFromGasParams(
    gasParams: L1toL2MessageGasValues,
    destAddr: string,
    callDataHex: string,
    options: CreateRetryableTicketOpptions = {
      excessFeeRefundAddress: undefined,
      callValueRefundAddress: undefined,
    },
    overrides: PayableOverrides = {}
  ): Promise<ContractReceipt> {
    const {
      maxGasPriceBid,
      maxSubmissionPriceBid,
      maxGasBid,
      totalDepositValue,
      l2CallValue,
    } = gasParams
    const sender = await this.getSender()
    const excessFeeRefundAddress = options.excessFeeRefundAddress || sender
    const callValueRefundAddress = options.callValueRefundAddress || sender

    const inboxAddress = this.networksInfo.ethBridge.inbox
    const inbox = Inbox__factory.connect(inboxAddress, this.l1Signer)

    const res = await inbox.createRetryableTicket(
      destAddr,
      l2CallValue,
      maxSubmissionPriceBid,
      excessFeeRefundAddress,
      callValueRefundAddress,
      maxGasBid,
      maxGasPriceBid,
      callDataHex,
      { value: totalDepositValue, ...overrides }
    )
    return res.wait()
  }

  public async createRetryableTicket(
    destAddr: string,
    callDataHex: string,
    l2CallValue: BigNumber,
    l2Provider: Provider,
    options: CreateRetryableTicketOpptions = {
      excessFeeRefundAddress: undefined,
      callValueRefundAddress: undefined,
    }
  ): Promise<L1ToL2Message[]> {
    const sender = await this.getSender()
    const gasEstimator = new L1ToL2MessageGasEstimator(l2Provider)
    const gasParams = await gasEstimator.estimateGasValuesL1ToL2Creation(
      sender,
      destAddr,
      callDataHex,
      l2CallValue
    )
    const rec = await this.createRetryableTicketFromGasParams(
      gasParams,
      destAddr,
      callDataHex,
      options
    )

    return L1ToL2Message.fromL1ReceiptAll(l2Provider, rec)
  }

  public async getSender(): Promise<string> {
    if (!this.sender) {
      const sender = await this.l1Signer.getAddress()
      this.sender = sender
      return sender
    }
    return this.sender
  }
}
