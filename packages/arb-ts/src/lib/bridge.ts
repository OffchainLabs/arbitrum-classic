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
import { Signer, BigNumber } from 'ethers'
import { L1Bridge } from './l1Bridge'
import { L2Bridge } from './l2Bridge'

export class Bridge extends L2Bridge {
  l1Bridge: L1Bridge
  walletAddressCache?: string

  constructor(
    erc20BridgeAddress: string,
    arbERC20BridgeAddress: string,
    ethSigner: Signer,
    arbSigner: Signer
  ) {
    super(erc20BridgeAddress, ethSigner)
    // TODO can presumably get arbERC20BridgeAddress directly from the L! bridge
    this.l1Bridge = new L1Bridge(arbERC20BridgeAddress, arbSigner)
  }

  public async updateAllTokens() {
    const l1Tokens = await this.l1Bridge.updateAllL1Tokens()
    const l2Tokens = await this.updateAllL2Tokens()
    return { l1Tokens, l2Tokens }
  }

  public async approveToken(erc20L1Address: string) {
    return this.l1Bridge.approveToken(erc20L1Address)
  }

  public async depositETH(value: BigNumber, destinationAddress?: string) {
    return this.l1Bridge.depositETH(value, destinationAddress)
  }

  public async depositAsERC20(
    erc20L1Address: string,
    amount: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string
  ) {
    return this.l1Bridge.depositAsERC20(
      erc20L1Address,
      amount,
      maxGas,
      gasPriceBid,
      destinationAddress
    )
  }
  public async depositAsERC777(
    erc20L1Address: string,
    amount: BigNumber,
    maxGas: BigNumber,
    gasPriceBid: BigNumber,
    destinationAddress?: string
  ) {
    return this.l1Bridge.depositAsERC777(
      erc20L1Address,
      amount,
      maxGas,
      gasPriceBid,
      destinationAddress
    )
  }
}
