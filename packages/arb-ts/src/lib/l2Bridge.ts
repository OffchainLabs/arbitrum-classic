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
import { ArbTokenBridge__factory } from './abi/factories/ArbTokenBridge__factory'
import { ArbTokenBridge } from './abi/ArbTokenBridge'
import { ArbSys } from './abi/ArbSys'
import { ArbSys__factory } from './abi/factories/ArbSys__factory'
import { StandardArbERC20 } from './abi/StandardArbERC20'
import { ICustomToken } from './abi/ICustomToken'
import { ICustomToken__factory } from './abi/factories/ICustomToken__factory'
import { StandardArbERC20__factory } from './abi/factories/StandardArbERC20__factory'
import { IArbToken } from './abi/IArbToken'
import { IArbToken__factory } from './abi/factories/IArbToken__factory'
import { ArbRetryableTx__factory } from './abi/factories/ArbRetryableTx__factory'
import { ArbRetryableTx } from './abi/ArbRetryableTx'
import { PayableOverrides } from '@ethersproject/contracts'

export const ARB_SYS_ADDRESS = '0x0000000000000000000000000000000000000064'
const ARB_RETRYABLE_TX_ADDRESS = '0x000000000000000000000000000000000000006E'

export interface L2TokenData {
  ERC20?: { contract: StandardArbERC20; balance: BigNumber }
  CUSTOM?: { contract: ICustomToken; balance: BigNumber } // Here we don't use the particlar custom token's interface; for the sake of this sdk that's fine
}

export interface Tokens {
  [contractAddress: string]: L2TokenData | undefined
}
/**
 * L2 side only of {@link Bridge}
 */
export class L2Bridge {
  l2Signer: Signer
  arbSys: ArbSys
  arbTokenBridge: ArbTokenBridge
  l2Tokens: Tokens
  l2Provider: providers.Provider
  l2EthBalance: BigNumber
  arbRetryableTx: ArbRetryableTx
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

    this.arbTokenBridge = ArbTokenBridge__factory.connect(
      arbTokenBridgeAddress,
      l2Signer
    )

    this.arbRetryableTx = ArbRetryableTx__factory.connect(
      ARB_RETRYABLE_TX_ADDRESS,
      l2Signer
    )

    this.l2EthBalance = BigNumber.from(0)
  }

  public async withdrawETH(
    value: BigNumber,
    destinationAddress?: string,
    overrides?: PayableOverrides
  ) {
    const address = destinationAddress || (await this.getWalletAddress())
    return this.arbSys.functions.withdrawEth(address, {
      value,
      ...overrides,
    })
  }

  public getLatestBlock() {
    return this.l2Provider.getBlock('latest')
  }
  public async withdrawERC20(
    erc20l1Address: string,
    amount: BigNumber,
    destinationAddress?: string,
    overrides: PayableOverrides = {}
  ) {
    const destination = destinationAddress || (await this.getWalletAddress())

    const tokenData = await this.getAndUpdateL2TokenData(erc20l1Address)
    if (!tokenData) {
      throw new Error("Can't withdraw; token not deployed")
    }
    const erc20TokenData = tokenData.ERC20

    if (!erc20TokenData) {
      throw new Error(
        `Can't withdraw; ArbERC20 for ${erc20l1Address} doesn't exist`
      )
    }
    return erc20TokenData.contract.functions.withdraw(
      destination,
      amount,
      overrides
    )
  }

  public async updateAllL2Tokens() {
    for (const l1Address in this.l2Tokens) {
      await this.getAndUpdateL2TokenData(l1Address)
    }
    return this.l2Tokens
  }

  public async getAndUpdateL2TokenData(erc20L1Address: string) {
    const tokenData = this.l2Tokens[erc20L1Address] || {
      ERC20: undefined,
      ERC777: undefined,
      CUSTOM: undefined,
    }
    this.l2Tokens[erc20L1Address] = tokenData
    const walletAddress = await this.getWalletAddress()

    // handle custom L2 token:
    const [
      customTokenAddress,
    ] = await this.arbTokenBridge.functions.customL2Token(erc20L1Address)
    if (customTokenAddress !== ethers.constants.AddressZero) {
      const customTokenContract = ICustomToken__factory.connect(
        customTokenAddress,
        this.l2Signer
      )
      tokenData.CUSTOM = {
        contract: customTokenContract,
        balance: BigNumber.from(0),
      }
      try {
        const [balance] = await customTokenContract.functions.balanceOf(
          walletAddress
        )
        tokenData.CUSTOM.balance = balance
      } catch (err) {
        console.warn("Could not get custom token's balance", err)
      }
    }

    const l2ERC20Address = await this.getERC20L2Address(erc20L1Address)

    // check if standard arb erc20:
    if (!tokenData.ERC20) {
      if ((await this.l2Provider.getCode(l2ERC20Address)).length > 2) {
        const arbERC20TokenContract = await StandardArbERC20__factory.connect(
          l2ERC20Address,
          this.l2Signer
        )
        const [balance] = await arbERC20TokenContract.functions.balanceOf(
          walletAddress
        )
        tokenData.ERC20 = {
          contract: arbERC20TokenContract,
          balance,
        }
      } else {
        console.info(
          `Corresponding ArbERC20 for ${erc20L1Address} not yet deployed (would be at ${l2ERC20Address})`
        )
      }
    } else {
      const arbERC20TokenContract = await StandardArbERC20__factory.connect(
        l2ERC20Address,
        this.l2Signer
      )
      const [balance] = await arbERC20TokenContract.functions.balanceOf(
        walletAddress
      )
      tokenData.ERC20.balance = balance
    }

    if (tokenData.ERC20 || tokenData.CUSTOM) {
      return tokenData
    } else {
      console.warn(`No L2 token for ${erc20L1Address} found`)
      return
    }
  }

  public getERC20L2Address(erc20L1Address: string) {
    let address: string | undefined
    if ((address = this.l2Tokens[erc20L1Address]?.ERC20?.contract.address)) {
      return address
    }
    return this.arbTokenBridge.functions
      .calculateL2TokenAddress(erc20L1Address)
      .then(([res]) => res)
  }

  public getERC20L1Address(erc20L2Address: string) {
    try {
      const arbERC20 = StandardArbERC20__factory.connect(
        erc20L2Address,
        this.l2Signer
      )
      return arbERC20.functions.l1Address().then(([res]) => res)
    } catch (e) {
      console.warn('Could not get L1 Address')

      return undefined
    }
  }

  public getTxnSubmissionPrice(dataSize: BigNumber) {
    return this.arbRetryableTx.functions.getSubmissionPrice(dataSize)
  }

  public async getWalletAddress() {
    if (this.walletAddressCache) {
      return this.walletAddressCache
    }
    this.walletAddressCache = await this.l2Signer.getAddress()
    return this.walletAddressCache
  }

  public async getAndUpdateL2EthBalance(): Promise<BigNumber> {
    const bal = await this.l2Signer.getBalance()
    this.l2EthBalance = bal
    return bal
  }
}
