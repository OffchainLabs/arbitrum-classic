/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/* eslint-env node */
'use strict'

import { BigNumber } from 'ethers'
import { ArbTsError } from '../errors'
import { L1ContractTransaction } from '../message/L1ToL2Message'
import { L2ContractTransaction } from '../message/L2ToL1Message'

import networks, { L1Network, L2Network } from '../utils/networks'
import {
  SignerOrProvider,
  SignerProviderUtils,
} from '../utils/signerOrProvider'

/**
 * Base for bridging assets from l1 to l2 and back
 */
export abstract class AssetBridger<DepositParams, WithdrawParams> {
  public readonly l1Network: L1Network

  public constructor(public readonly l2Network: L2Network) {
    this.l1Network = networks.l1Networks[l2Network.partnerChainID]
    if (!this.l1Network) {
      throw new ArbTsError(
        `Unknown l1 network chain id: ${l2Network.partnerChainID}`
      )
    }
  }

  /**
   * Check the signer/provider matches the l1Network, throws if not
   * @param sop
   */
  protected checkL1Network(sop: SignerOrProvider): void {
    SignerProviderUtils.checkNetworkMatches(
      sop,
      parseInt(this.l1Network.chainID)
    )
  }

  /**
   * Check the signer/provider matches the l2Network, throws if not
   * @param sop
   */
  protected checkL2Network(sop: SignerOrProvider): void {
    SignerProviderUtils.checkNetworkMatches(
      sop,
      parseInt(this.l2Network.chainID)
    )
  }

  /**
   * Estimate gas for transfering assets from L1 to L2
   * @param params
   */
  public abstract depositEstimateGas(params: DepositParams): Promise<BigNumber>

  /**
   * Transfer assets from L1 to L2
   * @param params
   */
  public abstract deposit(params: DepositParams): Promise<L1ContractTransaction>

  /**
   * Estimate gas for transfering assets from L2 to L1
   * @param params
   */
  public abstract withdrawEstimateGas(
    params: WithdrawParams
  ): Promise<BigNumber>

  /**
   * Transfer assets from L2 to L1
   * @param params
   */
  public abstract withdraw(
    params: WithdrawParams
  ): Promise<L2ContractTransaction>
}
