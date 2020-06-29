/*
 * Copyright 2020, Offchain Labs, Inc.
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
/* eslint-env browser */
'use strict'

import * as ethers from 'ethers'

// TODO remove this dep
const jaysonBrowserClient = require('jayson/lib/client/browser') // eslint-disable-line @typescript-eslint/no-var-requires

function _aggregatorClient(managerAddress: string): any {
  const callServer = (request: any, callback: any): void => {
    const options = {
      body: request, // request is a string
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'POST',
    }

    fetch(managerAddress, options)
      .then((res: any) => {
        return res.text()
      })
      .then((text: string) => {
        callback(null, text)
      })
      .catch((err: Error) => {
        callback(err)
      })
  }

  return jaysonBrowserClient(callServer, {})
}

export class AggregatorClient {
  public client: any

  constructor(managerUrl: string) {
    this.client = _aggregatorClient(managerUrl)
  }

  public async sendTransaction(
    to: string,
    sequenceNum: ethers.utils.BigNumber,
    value: ethers.utils.BigNumber,
    data: string,
    pubkey: string,
    signature: string
  ): Promise<txaggregator.SendTransactionReply> {
    return new Promise<txaggregator.SendTransactionReply>(
      (resolve, reject): void => {
        const params: txaggregator.SendTransactionArgs = {
          to,
          sequenceNum: sequenceNum.toString(),
          value: value.toString(),
          data,
          pubkey,
          signature,
        }
        this.client.request(
          'TxAggregator.SendTransaction',
          [params],
          (
            err: Error,
            error: Error,
            result: txaggregator.SendTransactionReply
          ) => {
            if (err) {
              reject(err)
            } else if (error) {
              reject(error)
            } else {
              resolve(result)
            }
          }
        )
      }
    )
  }
}
