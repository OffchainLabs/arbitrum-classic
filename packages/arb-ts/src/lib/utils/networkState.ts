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

import { Provider, Filter } from '@ethersproject/abstract-provider'
import { BigNumber } from 'ethers'

import { Whitelist__factory, Rollup__factory, Rollup } from '../abi'
import { NodeConfirmedEvent, NodeCreatedEvent } from '../abi/Rollup'
import { EventFetcher } from '../utils/eventFetcher'
import { ADDRESS_ALIAS_OFFSET } from '../constants'

/**
 * General information about the current network state
 */
export class NetworkState {
  public constructor(public readonly l1Provider: Provider) {}

  /**
   * Find the L2 alias of an L1 address
   * @param l1Address
   * @returns
   */
  public static applyL1ToL2Alias(l1Address: string): BigNumber {
    return BigNumber.from(l1Address).add(ADDRESS_ALIAS_OFFSET)
  }

  /**
   * Find the L1 alias of an L2 address
   * @param l2Address
   * @returns
   */
  public static undoL1ToL2Alias(l2Address: string): BigNumber {
    return BigNumber.from(l2Address).sub(ADDRESS_ALIAS_OFFSET)
  }

  /**
   * Check if an address is whitelisted
   * @param address The address to check
   * @param whiteListAddress The whitelist contract address
   * @returns
   */
  public isWhiteListed(
    address: string,
    whiteListAddress: string
  ): Promise<boolean> {
    const whiteList = Whitelist__factory.connect(
      whiteListAddress,
      this.l1Provider
    )
    return whiteList.isAllowed(address)
  }

  /**
   * Get the NodeCreated events
   * @param rollupAddress
   * @param nodeNum
   * @param parentNodeHash
   * @param filter
   * @returns
   */
  public async getNodeCreatedEvents(
    rollupAddress: string,
    nodeNum?: BigNumber,
    parentNodeHash?: string,
    filter?: Omit<Filter, 'address' | 'topics'>
  ): Promise<NodeCreatedEvent['args'][]> {
    const eventFetcher = new EventFetcher(this.l1Provider)
    return await eventFetcher.getEvents<Rollup, NodeCreatedEvent>(
      rollupAddress,
      Rollup__factory,
      r => r.filters.NodeCreated(nodeNum, parentNodeHash),
      filter
    )
  }

  /**
   * Get the NodeConfirmed events
   * @param rollupAddress
   * @param nodeNum
   * @param parentNodeHash
   * @param filter
   * @returns
   */
  public getNodeConfirmedEvents(
    rollupAddress: string,
    nodeNum?: BigNumber,
    filter?: Omit<Filter, 'address' | 'topics'>
  ): Promise<NodeConfirmedEvent['args'][]> {
    const eventFetcher = new EventFetcher(this.l1Provider)
    return eventFetcher.getEvents<Rollup, NodeConfirmedEvent>(
      rollupAddress,
      Rollup__factory,
      r => r.filters.NodeConfirmed(nodeNum),
      filter
    )
  }

  public async contractExists(contractAddress: string): Promise<boolean> {
    const contractCode = await this.l1Provider.getCode(contractAddress)
    return !(contractCode.length > 2)
  }
}
