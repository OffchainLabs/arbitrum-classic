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

import { Provider, Filter, EventFilter } from '@ethersproject/abstract-provider'
import { Contract, Event } from '@ethersproject/contracts'

/**
 * Fetches and parses blockchain logs
 */
export class EventFetcher {
  public constructor(public readonly provider: Provider) {}

  /**
   * Fetch logs and parse logs
   * @param addr The address of the contract emitting the events
   * @param contractFactory A contract factory for generating a contract of type TContract at the addr
   * @param topicGenerator Generator function for creating
   * @param filter Block filter parameters
   * @returns
   */
  public async getEvents<TContract extends Contract, TEvent extends Event>(
    addr: string,
    contractFactory: {
      connect(address: string, provider: Provider): TContract
    },
    topicGenerator: (t: TContract) => EventFilter,
    filter?: Omit<Filter, 'topics' | 'address'>
  ): Promise<TEvent['args'][]> {
    const contract = contractFactory.connect(addr, this.provider)
    const eventFilter = topicGenerator(contract)
    const fullFilter = {
      ...eventFilter,
      fromBlock: filter?.fromBlock || 0,
      toBlock: filter?.toBlock || 'latest',
    }
    const logs = await this.provider.getLogs(fullFilter)
    return logs.map(
      l => contract.interface.parseLog(l).args
    ) as TEvent['args'][]
  }
}
