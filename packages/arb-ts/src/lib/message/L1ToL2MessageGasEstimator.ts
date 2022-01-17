import { Provider } from '@ethersproject/abstract-provider'
import { NodeInterface__factory, ArbRetryableTx__factory } from '../abi'

import { ARB_RETRYABLE_TX_ADDRESS, NODE_INTERFACE_ADDRESS } from '../constants'
import { BigNumber } from '@ethersproject/bignumber'
import { constants } from 'ethers'
import { utils } from 'ethers'

const DEFAULT_SUBMISSION_PRICE_PERCENT_INCREASE = BigNumber.from(340)

/**
 * An optional big number percentage increase
 */
export type PercentIncrease = {
  /**
   * If provided, will override the estimated base
   */
  base?: BigNumber

  /**
   * How much to increase the base by. If not provided system defaults may be used.
   */
  percentIncrease?: BigNumber
}

export interface GasOverrides {
  maxGas?: PercentIncrease & {
    /**
     * Set a minimum max gas
     */
    min?: BigNumber
  }
  maxSubmissionPrice?: PercentIncrease
  maxGasPrice?: PercentIncrease
  sendL2CallValueFromL1?: boolean
}

const defaultL1ToL2MessageEstimateOptions = {
  maxSubmissionFeePercentIncrease: DEFAULT_SUBMISSION_PRICE_PERCENT_INCREASE,
  maxGasPercentIncrease: constants.Zero,
  maxGasPricePercentIncrease: constants.Zero,
  sendL2CallValueFromL1: true,
}

export interface L1toL2MessageGasValues {
  maxGasPriceBid: BigNumber
  maxSubmissionPriceBid: BigNumber
  maxGasBid: BigNumber

  totalDepositValue: BigNumber
  l2CallValue: BigNumber
}

export class L1ToL2MessageGasEstimator {
  constructor(public readonly l2Provider: Provider) {}

  private percentIncrease(num: BigNumber, increase: BigNumber): BigNumber {
    return num.add(num.mul(increase).div(100))
  }

  private applySubmissionPriceDefaults(maxSubmissionPrice?: PercentIncrease) {
    return {
      base: maxSubmissionPrice?.base,
      percentIncrease:
        maxSubmissionPrice?.percentIncrease ||
        defaultL1ToL2MessageEstimateOptions.maxSubmissionFeePercentIncrease,
    }
  }

  public async getSubmissionPrice(
    callDataSize: BigNumber | number,
    options?: {
      base?: BigNumber
      percentIncrease?: BigNumber
    }
  ): Promise<{
    submissionPrice: BigNumber
    nextUpdateTimestamp: BigNumber
  }> {
    const defaultedOptions = this.applySubmissionPriceDefaults(options)

    const arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      this.l2Provider
    )
    const [currentSubmissionPrice, nextUpdateTimestamp] =
      await arbRetryableTx.getSubmissionPrice(callDataSize)
    // Apply percent increase
    const submissionPrice = this.percentIncrease(
      defaultedOptions.base || currentSubmissionPrice,
      defaultedOptions.percentIncrease
    )
    return {
      submissionPrice,
      nextUpdateTimestamp,
    }
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
    gasPriceBid: BigNumber,
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
        gasPriceBid,
        calldata
      )
    )[0]
  }

  private applyDefaults(options?: GasOverrides) {
    return {
      maxGas: {
        base: options?.maxGasPrice?.base,
        percentIncrease:
          options?.maxGas?.percentIncrease ||
          defaultL1ToL2MessageEstimateOptions.maxGasPercentIncrease,
        min: options?.maxGas?.min || constants.Zero,
      },
      maxGasPrice: {
        base: options?.maxGasPrice?.base,
        percentIncrease:
          options?.maxGasPrice?.percentIncrease ||
          defaultL1ToL2MessageEstimateOptions.maxGasPricePercentIncrease,
      },
      sendL2CallValueFromL1:
        typeof options?.sendL2CallValueFromL1 === 'boolean'
          ? options?.sendL2CallValueFromL1
          : defaultL1ToL2MessageEstimateOptions.sendL2CallValueFromL1,
    }
  }

  public async estimateGasValuesL1ToL2Creation(
    sender: string,
    destAddr: string,
    callDataHex: string,
    l2CallValue: BigNumber,
    options?: GasOverrides
  ): Promise<{
    maxGasBid: BigNumber
    maxSubmissionPriceBid: BigNumber
    maxGasPriceBid: BigNumber
    totalDepositValue: BigNumber
  }> {
    const defaultedOptions = this.applyDefaults(options)

    const maxGasPriceBid = this.percentIncrease(
      defaultedOptions.maxGasPrice.base ||
        (await this.l2Provider.getGasPrice()),
      defaultedOptions.maxGasPrice.percentIncrease
    )

    const maxSubmissionPriceBid = (
      await this.getSubmissionPrice(
        utils.hexDataLength(callDataHex),
        options?.maxSubmissionPrice
      )
    ).submissionPrice

    const calculatedMaxGas = this.percentIncrease(
      defaultedOptions.maxGas.base ||
        (await this.estimateGasRetryableTicket(
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
        )),
      defaultedOptions.maxGas.percentIncrease
    )
    // always ensure the max gas is greater than the min
    const maxGas = calculatedMaxGas.gt(defaultedOptions.maxGas.min)
      ? calculatedMaxGas
      : defaultedOptions.maxGas.min

    let totalDepositValue = maxSubmissionPriceBid.add(
      maxGasPriceBid.mul(maxGas)
    )

    if (defaultedOptions.sendL2CallValueFromL1) {
      totalDepositValue = totalDepositValue.add(l2CallValue)
    }
    return {
      maxGasBid: maxGas,
      maxSubmissionPriceBid,
      maxGasPriceBid,
      totalDepositValue,
    }
  }
}
