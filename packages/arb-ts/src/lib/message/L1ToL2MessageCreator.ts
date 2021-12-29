import { Signer } from '@ethersproject/abstract-signer'
import { Provider } from '@ethersproject/abstract-provider'

import {
  L1ToL2MessageGasEstimator,
  L1toL2MessageGasValues,
} from './L1ToL2MessageGasEstimator'
import { L1ToL2Message, L1TransactionReceipt } from './L1ToL2Message'
import { Inbox__factory } from '../abi/factories/Inbox__factory'
import { l2Networks } from '../utils/networks'
import { ContractReceipt, PayableOverrides } from '@ethersproject/contracts'
import { BigNumber } from 'ethers'

interface CreateRetryableTicketOpptions {
  excessFeeRefundAddress?: string
  callValueRefundAddress?: string
}
export class L1ToL2MessageCreator {
  sender?: string
  constructor(public readonly l1Signer: Signer) {}

  public async createRetryableTicketFromGasParams(
    gasParams: L1toL2MessageGasValues,
    destAddr: string,
    callDataHex: string,
    l2ChainID: string,
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

    const inboxAddress = l2Networks[l2ChainID].ethBridge.inbox
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
  ): Promise<L1TransactionReceipt> {
    const sender = await this.getSender()
    const gasEstimator = new L1ToL2MessageGasEstimator(l2Provider)
    const gasParams = await gasEstimator.estimateGasValuesL1ToL2Creation(
      sender,
      destAddr,
      callDataHex,
      l2CallValue
    )
    const l2ChainID = (await l2Provider.getNetwork()).chainId.toString()
    const rec = await this.createRetryableTicketFromGasParams(
      gasParams,
      destAddr,
      callDataHex,
      l2ChainID,
      options
    )

    return new L1TransactionReceipt(rec)
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
