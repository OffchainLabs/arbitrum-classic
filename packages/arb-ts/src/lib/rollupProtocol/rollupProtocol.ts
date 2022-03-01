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

import { Provider, BlockTag } from '@ethersproject/abstract-provider'
import { BigNumber } from 'ethers'

import { Rollup__factory } from '../abi'
import { NodeConfirmedEvent, NodeCreatedEvent } from '../abi/Rollup'
import { EventFetcher } from '../utils/eventFetcher'

/**
 * General information about protocol level actions, WIP
 */
export class RollupProtocol {
  public constructor(public readonly l1Provider: Provider) {}

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
    filter: { fromBlock: BlockTag; toBlock: BlockTag },
    nodeNum?: BigNumber,
    parentNodeHash?: string
  ): Promise<NodeCreatedEvent['args'][]> {
    const eventFetcher = new EventFetcher(this.l1Provider)
    return (
      await eventFetcher.getEvents(
        rollupAddress,
        Rollup__factory,
        r => r.filters.NodeCreated(nodeNum, parentNodeHash),
        filter
      )
    ).map(a => a.event)
  }

  /**
   * Get the NodeConfirmed events
   * @param rollupAddress
   * @param nodeNum
   * @param parentNodeHash
   * @param filter
   * @returns
   */
  public async getNodeConfirmedEvents(
    rollupAddress: string,
    filter: { fromBlock: BlockTag; toBlock: BlockTag },
    nodeNum?: BigNumber
  ): Promise<NodeConfirmedEvent['args'][]> {
    const eventFetcher = new EventFetcher(this.l1Provider)
    return (
      await eventFetcher.getEvents(
        rollupAddress,
        Rollup__factory,
        r => r.filters.NodeConfirmed(nodeNum),
        filter
      )
    ).map(a => a.event)
  }
}
