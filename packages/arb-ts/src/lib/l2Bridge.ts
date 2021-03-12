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
import { Signer, BigNumber, providers, ethers } from 'ethers'
import { ArbTokenBridgeFactory } from './abi/ArbTokenBridgeFactory'
import { ArbTokenBridge } from './abi/ArbTokenBridge'
import { ArbSys } from './abi/ArbSys'
import { ArbSys__factory } from './abi/factories/ArbSys__Factory'
import { StandardArbERC20 } from './abi/StandardArbERC20'
import { StandardArbERC20Factory } from './abi/StandardArbERC20Factory'
import { StandardArbERC777Factory } from './abi/StandardArbERC777Factory'
import { IArbToken } from './abi/IArbToken'
import { IArbTokenFactory } from './abi/IArbTokenFactory'

import { StandardArbERC777 } from './abi/StandardArbERC777'

const ARB_SYS_ADDRESS = '0x0000000000000000000000000000000000000064'

interface L2TokenData {
  ERC20?: { contract: StandardArbERC20; balance: BigNumber }
  ERC777?: { contract: StandardArbERC777; balance: BigNumber }
  CUSTOM?: { contract: IArbToken; balance: BigNumber }
}

export interface Tokens {
  [contractAddress: string]: L2TokenData | undefined
}

export class L2Bridge {
  l2Signer: Signer
  arbSys: ArbSys
  arbTokenBridge: ArbTokenBridge
  l2Tokens: Tokens
  l2Provider: providers.Provider
  walletAddressCache?: string

  constructor(arbTokenBridgeAddress: string, l2Signer: Signer) {
    this.l2Tokens = {}

    this.l2Signer = l2Signer

    const l2Provider = l2Signer.provider

    if (l2Provider === undefined) {
      throw new Error('Signer must be connected to an (Arbitrum) provider')
    }
    this.l2Provider = l2Provider

    this.arbSys = ArbSys__factory.connect(ARB_SYS_ADDRESS, l2Signer)

    this.arbTokenBridge = ArbTokenBridgeFactory.connect(
      arbTokenBridgeAddress,
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

    const tokenData = await this.getAndUpdateTokenData(erc20l1Address)
    const erc20TokenData = tokenData.ERC20

    if (!erc20TokenData) {
      throw new Error(
        `Can't withdraw; ArbERC20 for ${erc20l1Address} doesn't exist`
      )
    }
    return erc20TokenData.contract.withdraw(destination, amount)
  }

  public async withdrawERC777(
    erc20l1Address: string,
    amount: BigNumber,
    destinationAddress?: string
  ) {
    const destination = destinationAddress || (await this.getWalletAddress())

    const tokenData = await this.getAndUpdateTokenData(erc20l1Address)
    const erc777TokenData = tokenData.ERC777

    if (!erc777TokenData) {
      throw new Error(
        `Can't withdraw; ArbERC777 for ${erc20l1Address} doesn't exist`
      )
    }
    return erc777TokenData.contract.withdraw(destination, amount)
  }

  public async updateAllL2Tokens() {
    for (const l1Address in this.l2Tokens) {
      await this.getAndUpdateTokenData(l1Address)
    }
    return this.l2Tokens
  }

  public async getAndUpdateTokenData(erc20L1Address: string) {
    if (!this.l2Tokens[erc20L1Address]) {
      this.l2Tokens[erc20L1Address] = {
        ERC20: undefined,
        ERC777: undefined,
        CUSTOM: undefined,
      }
    }

    const tokenData = this.l2Tokens[erc20L1Address] as L2TokenData // truthiness is ensured above
    const walletAddress = await this.getWalletAddress()

    const customTokenAddress = await this.arbTokenBridge.customToken(
      erc20L1Address
    )
    if (customTokenAddress !== ethers.constants.AddressZero) {
      const customTokenContract = IArbTokenFactory.connect(
        customTokenAddress,
        this.l2Signer
      )
      tokenData.CUSTOM = {
        contract: customTokenContract,
        balance: BigNumber.from(0),
      }
      try {
        const balance = (await customTokenContract.balanceOf(
          walletAddress
        )) as BigNumber
        tokenData.CUSTOM.balance = balance
      } catch (err) {
        console.warn("Count not get custom token's balance", err)
      }
      return tokenData
    }

    const l2ERC20Address = await this.getERC20L2Address(erc20L1Address)
    const l2ERC777Address = await this.getERC777L2Address(erc20L1Address)

    if (!tokenData.ERC20) {
      if ((await this.l2Provider.getCode(l2ERC20Address)).length > 2) {
        const arbERC20TokenContract = await StandardArbERC20Factory.connect(
          l2ERC20Address,
          this.l2Signer
        )
        const balance = await arbERC20TokenContract.balanceOf(walletAddress)
        tokenData.ERC20 = {
          contract: arbERC20TokenContract,
          balance,
        }
      } else {
        console.info(
          `Corresponding ArbERC20 for ${erc20L1Address} not yet deployed`
        )
      }
    } else {
      const arbERC20TokenContract = await StandardArbERC20Factory.connect(
        l2ERC20Address,
        this.l2Signer
      )
      const balance = await arbERC20TokenContract.balanceOf(walletAddress)
      tokenData.ERC20.balance = balance
    }

    if (!tokenData.ERC777) {
      if ((await this.l2Provider.getCode(l2ERC777Address)).length > 2) {
        const arbERC77TokenContract = await StandardArbERC777Factory.connect(
          l2ERC777Address,
          this.l2Signer
        )
        const balance = await arbERC77TokenContract.balanceOf(walletAddress)
        tokenData.ERC777 = {
          contract: arbERC77TokenContract,
          balance,
        }
      }
      // else: : ERC777 not deployed
    } else {
      const arbERC777TokenContract = await StandardArbERC777Factory.connect(
        l2ERC777Address,
        this.l2Signer
      )
      const balance = await arbERC777TokenContract.balanceOf(walletAddress)
      tokenData.ERC777.balance = balance
    }

    return tokenData
  }

  public getERC20L2Address(erc20L1Address: string) {
    let address: string | undefined
    if ((address = this.l2Tokens[erc20L1Address]?.ERC20?.contract.address)) {
      return address
    }
    return this.arbTokenBridge.calculateBridgedERC20Address(erc20L1Address)
  }

  public getERC777L2Address(erc20L1Address: string) {
    let address: string | undefined
    if ((address = this.l2Tokens[erc20L1Address]?.ERC777?.contract.address)) {
      return address
    }
    return this.arbTokenBridge.calculateBridgedERC777Address(erc20L1Address)
  }

  public async getWalletAddress() {
    if (this.walletAddressCache) {
      return this.walletAddressCache
    }
    this.walletAddressCache = await this.l2Signer.getAddress()
    return this.walletAddressCache
  }
}
