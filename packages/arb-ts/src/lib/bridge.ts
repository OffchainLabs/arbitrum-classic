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

export class Bridge extends L1Bridge {
  l2Bridge: L2Bridge
  walletAddressCache?: string

  constructor(
    inboxAddress: string,
    erc20BridgeAddress: string,
    arbERC20BridgeAddress: string,
    ethSigner: Signer,
    arbSigner: Signer
  ) {
    super(inboxAddress, erc20BridgeAddress, ethSigner)
    this.l2Bridge = new L2Bridge(arbERC20BridgeAddress, arbSigner)
  }

  public async withdrawETH(value: BigNumber, destinationAddress?: string) {
    return await this.l2Bridge.withdrawETH(value, destinationAddress)
  }
  public async withdrawERC20(
    erc20l1Address: string,
    amount: BigNumber,
    destinationAddress?: string
  ) {
    return await this.l2Bridge.withdrawERC20(
      erc20l1Address,
      amount,
      destinationAddress
    )
  }
  public async getERC20L1Address(erc20L2Address: string) {
    return await this.l2Bridge.getERC20L1Address(erc20L2Address)
  }
}
