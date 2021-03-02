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
import { ArbTokenBridgeFactory } from './abi/ArbTokenBridgeFactory'
import { ArbTokenBridge } from './abi/ArbTokenBridge'
import { ArbSys } from './abi/ArbSys'
import { ArbSysFactory } from './abi/ArbSysFactory'

const ARB_SYS_ADDRESS = '0x0000000000000000000000000000000000000064'

export class L2Bridge {
  l2Signer: Signer
  arbSys: ArbSys
  arbERC20Bridge: ArbTokenBridge
  walletAddressCache?: string

  constructor(arbERC20BridgeAddress: string, l2Signer: Signer) {
    this.l2Signer = l2Signer

    this.arbSys = ArbSysFactory.connect(ARB_SYS_ADDRESS, l2Signer)

    this.arbERC20Bridge = ArbTokenBridgeFactory.connect(
      arbERC20BridgeAddress,
      l2Signer
    )
  }

  public async withdrawETH(value: BigNumber, destinationAddress?: string) {
    const address = destinationAddress || (await this.getWalletAddress())
    return this.arbSys.withdrawEth(address, {
      value,
    })
  }

  public async withdrawERC20(
    erc20l1Address: string,
    amount: BigNumber,
    destinationAddress?: string
  ) {
    const destination = destinationAddress || (await this.getWalletAddress())
    return await this.arbERC20Bridge.withdraw(
      erc20l1Address,
      destination,
      amount
    )
  }

  public getERC20L1Address(erc20L2Address: string) {
    return this.arbERC20Bridge.customToken(erc20L2Address)
  }

  public async getWalletAddress() {
    if (this.walletAddressCache) {
      return this.walletAddressCache
    }
    this.walletAddressCache = await this.l2Signer.getAddress()
    return this.walletAddressCache
  }
}
