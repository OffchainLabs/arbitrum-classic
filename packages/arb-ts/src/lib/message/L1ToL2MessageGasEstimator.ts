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
import { ArbTsError } from '../errors'

const DEFAULT_SUBMISSION_PRICE_PERCENT_INCREASE = BigNumber.from(340)
const DEFAULT_MAX_GAS_PERCENT_INCREASE = BigNumber.from(50)

export type PercentIncrease = {
  base?: BigNumber
  percentIncrease?: BigNumber
}

// CHRIS: better name for this
export interface GasOverrides {
  maxGas?: PercentIncrease
  maxSubmissionPrice?: PercentIncrease
  maxGasPrice?: PercentIncrease
  sendL2CallValueFromL1?: boolean
}

// CHRIS: remove this
// interface L1ToL2MessageEstimateOptions {
//   maxSubmissionFeePercentIncrease?: BigNumber
//   maxGasPercentIncrease?: BigNumber

//   // CHRIS: naming here?
//   gasPriceBid?: BigNumber
//   maxGasPricePercentIncrease?: BigNumber

//   sendL2CallValueFromL1?: boolean
// }

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
    const [
      currentSubmissionPrice,
      nextUpdateTimestamp,
    ] = await arbRetryableTx.getSubmissionPrice(callDataSize)
    // Apply percent increase
    const submissionPrice = percentIncrease(
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
  ) {
    const defaultedOptions = this.applyDefaults(options)

    const maxGasPriceBid = percentIncrease(
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

    const maxGasBid = percentIncrease(
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

    let totalDepositValue = maxSubmissionPriceBid.add(
      maxGasPriceBid.mul(maxGasBid)
    )
    if (defaultedOptions.sendL2CallValueFromL1) {
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
