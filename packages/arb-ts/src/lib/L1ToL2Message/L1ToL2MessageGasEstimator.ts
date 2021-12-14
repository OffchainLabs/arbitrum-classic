import { Provider } from '@ethersproject/abstract-provider'
import { ArbRetryableTx__factory } from '../abi/factories/ArbRetryableTx__factory'
import { NodeInterface__factory } from '../abi/factories/NodeInterface__factory'

import {
  ARB_RETRYABLE_TX_ADDRESS,
  NODE_INTERFACE_ADDRESS,
} from '../precompile_addresses'
import { BigNumber } from '@ethersproject/bignumber'
import { percentIncrease } from '../utils/lib'
import { constants } from 'ethers'
import { utils } from 'ethers'

const DEFAULT_SUBMISSION_PRICE_PERCENT_INCREASE = BigNumber.from(340)
const DEFAULT_MAX_GAS_PERCENT_INCREASE = BigNumber.from(50)

interface L1ToL2MessageEstimateOptions {
  maxSubmissionFeePercentIncrease?: BigNumber
  maxGasPercentIncrease?: BigNumber
  maxGasPricePercentIncrease?: BigNumber
  sendL2CallValueFromL1?: boolean
}

const defaultL1ToL2MessageEstimateOptions = {
  maxSubmissionFeePercentIncrease: DEFAULT_SUBMISSION_PRICE_PERCENT_INCREASE,
  maxGasPercentIncrease: constants.Zero,
  maxGasPricePercentIncrease: constants.Zero,
  sendL2CallValueFromL1: true,
}

interface L1toL2MessageGasValues {
  maxGasPriceBid: BigNumber
  maxSubmissionPriceBid: BigNumber
  maxGasBid: BigNumber
  totalDepositValue: BigNumber
}

export class L1ToL2MessageGasEstimator {
  constructor(public readonly l2Provider: Provider) {}

  public async getSubmissionPrice(
    callDataSize: BigNumber | number,
    options: {
      percentIncrease: BigNumber
    } = {
      percentIncrease: BigNumber.from(
        DEFAULT_SUBMISSION_PRICE_PERCENT_INCREASE // include percent increase by default
      ),
    }
  ): Promise<{
    submissionPrice: BigNumber
    nextUpdateTimestamp: BigNumber
  }> {
    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Provider
    )
    const [currentSubmissionPrice, nextUpdateTimestamp] =
      await arbRetryableTx.getSubmissionPrice(callDataSize)
    // Apply percent increase
    const submissionPrice = percentIncrease(
      currentSubmissionPrice,
      options.percentIncrease
    )
    return {
      submissionPrice,
      nextUpdateTimestamp,
    }
  }

  public getL2GasPrice(): Promise<BigNumber> {
    return this.l2Provider.getGasPrice()
  }

  public async estimateGasRetryableTicket(
    sender: string,
    senderDeposit: BigNumber,
    destAddr: string,
    l2CallValue: BigNumber,
    maxSubmissionCost: BigNumber,
    excessFeeRefundAddress: string,
    callValueRefundAddress: string,
    maxGas: BigNumber,
    gasPriceBig: BigNumber,
    calldata: string
  ): Promise<BigNumber> {
    const nodeInterface = NodeInterface__factory.connect(
      NODE_INTERFACE_ADDRESS,
      this.l2Provider
    )
    return (
      await nodeInterface.estimateRetryableTicket(
        sender,
        senderDeposit,
        destAddr,
        l2CallValue,
        maxSubmissionCost,
        excessFeeRefundAddress,
        callValueRefundAddress,
        maxGas,
        gasPriceBig,
        calldata
      )
    )[0]
  }

  public applyDefafultEstimateGasValues(
    options: L1ToL2MessageEstimateOptions
  ): {
    maxSubmissionFeePercentIncrease: BigNumber
    maxGasPercentIncrease: BigNumber
    maxGasPricePercentIncrease: BigNumber
    sendL2CallValueFromL1: boolean
  } {
    return {
      maxSubmissionFeePercentIncrease:
        options.maxSubmissionFeePercentIncrease ||
        defaultL1ToL2MessageEstimateOptions.maxGasPricePercentIncrease,
      maxGasPercentIncrease:
        options.maxGasPercentIncrease ||
        defaultL1ToL2MessageEstimateOptions.maxGasPercentIncrease,
      maxGasPricePercentIncrease:
        options.maxGasPricePercentIncrease ||
        defaultL1ToL2MessageEstimateOptions.maxGasPricePercentIncrease,
      sendL2CallValueFromL1:
        typeof options.sendL2CallValueFromL1 === 'boolean'
          ? options.sendL2CallValueFromL1
          : defaultL1ToL2MessageEstimateOptions.sendL2CallValueFromL1,
    }
  }

  public async estimateGasValuesL1ToL2Creation(
    callDataHex: string,
    sender: string,
    destAddr: string,
    l2CallValue: BigNumber,
    options: L1ToL2MessageEstimateOptions
  ): Promise<L1toL2MessageGasValues> {
    const {
      maxSubmissionFeePercentIncrease,
      maxGasPercentIncrease,
      maxGasPricePercentIncrease,
      sendL2CallValueFromL1,
    } = this.applyDefafultEstimateGasValues(options)

    const maxGasPriceBid = percentIncrease(
      await this.getL2GasPrice(),
      maxGasPricePercentIncrease as BigNumber
    )
    const { submissionPrice } = await this.getSubmissionPrice(
      utils.hexDataLength(callDataHex)
    )
    const maxSubmissionPriceBid = percentIncrease(
      submissionPrice,
      maxSubmissionFeePercentIncrease as BigNumber
    )

    const retryableGas = await this.estimateGasRetryableTicket(
      sender,
      utils
        .parseEther('1')
        .add(
          l2CallValue
        ) /** we add a 1 ether "deposit" buffer to pay for execution in the gas estimation  */,
      destAddr,
      l2CallValue,
      maxSubmissionPriceBid,
      sender,
      sender,
      constants.Zero,
      maxGasPriceBid,
      callDataHex
    )

    const maxGasBid = percentIncrease(
      retryableGas,
      maxGasPercentIncrease as BigNumber
    )

    let totalDepositValue = maxSubmissionPriceBid.add(
      maxGasPriceBid.mul(maxGasBid)
    )
    if (sendL2CallValueFromL1) {
      totalDepositValue = totalDepositValue.add(l2CallValue)
    }
    return {
      maxGasBid,
      maxSubmissionPriceBid,
      maxGasPriceBid,
      totalDepositValue,
    }
  }
}
