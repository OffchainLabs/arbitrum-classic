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

import { BigNumber, ethers } from 'ethers'

import { L2Network } from '../utils/networks'

/**
 * Base for bridging assets from l1 to l2 and back
 */
export abstract class AssetBridger<DepositParams, WithdrawParams> {
  public constructor(public readonly l2Network: L2Network) {}

  /**
   * Estimate gas for transfering assets from L1 to L2
   * @param params
   */
  public abstract depositEstimateGas(params: DepositParams): Promise<BigNumber>

  /**
   * Transfer assets from L1 to L2
   * @param params
   */
  public abstract deposit(
    params: DepositParams
  ): Promise<ethers.ContractTransaction>

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
  ): Promise<ethers.ContractTransaction>
}
