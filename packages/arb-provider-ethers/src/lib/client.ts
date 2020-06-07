/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

import * as ArbValue from './value'
import { EVMCode, EVMResult, processLog } from './message'

import * as ethers from 'ethers'

// TODO remove this dep
const jaysonBrowserClient = require('jayson/lib/client/browser') // eslint-disable-line @typescript-eslint/no-var-requires

function _arbClient(managerAddress: string): any {
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

interface MessageResult {
  logPostHash: string
  logPreHash: string
  logValHashes: string[]
  onChainTxHash: string
  val: ArbValue.Value
  vmId: string
  evmVal: EVMResult
}

interface OutputMessage {
  outputMsg: ArbValue.Value
}

function convertBlockTag(tag?: ethers.providers.BlockTag): string | undefined {
  if (tag === undefined || typeof tag == 'string') {
    return tag
  }

  return ethers.utils.bigNumberify(tag).toHexString()
}

function convertTopics(
  topicGroups?: Array<string | Array<string>>
): Array<validatorserver.TopicGroup> | undefined {
  if (topicGroups == undefined) {
    return topicGroups
  }
  return topicGroups.map(
    (topics): validatorserver.TopicGroup => {
      if (typeof topics == 'string') {
        return { topics: [topics] }
      } else {
        return { topics }
      }
    }
  )
}

export class ArbClient {
  public client: any

  constructor(managerUrl: string) {
    this.client = _arbClient(managerUrl)
  }

  public async getOutputMessage(
    AssertionNodeHash: string,
    msgIndex: string
  ): Promise<OutputMessage | null> {
    const params: validatorserver.GetOutputMessageArgs = {
      AssertionNodeHash,
      MsgIndex: msgIndex,
    }
    const msgResult = await new Promise<validatorserver.GetOutputMessageReply>(
      (resolve, reject): void => {
        this.client.request(
          'Validator.GetOutputMessage',
          [params],
          (
            err: Error,
            error: Error,
            result: validatorserver.GetOutputMessageReply
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
    if (msgResult.found && msgResult.rawVal !== undefined) {
      const val = ArbValue.unmarshal(msgResult.rawVal)

      return {
        outputMsg: val,
      }
    } else {
      return null
    }
  }

  public async getMessageResult(txHash: string): Promise<MessageResult | null> {
    const params: validatorserver.GetMessageResultArgs = {
      txHash,
    }
    const messageResult = await new Promise<
      validatorserver.GetMessageResultReply
    >((resolve, reject): void => {
      this.client.request(
        'Validator.GetMessageResult',
        [params],
        (
          err: Error,
          error: Error,
          result: validatorserver.GetMessageResultReply
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
    })
    if (messageResult.tx && messageResult.tx.found) {
      const tx = messageResult.tx
      const vmId = await this.getVmID()
      if (tx.rawVal === undefined) {
        return null
      }
      const val = ArbValue.unmarshal(tx.rawVal)
      const evmVal = processLog(val as ArbValue.TupleValue)
      if (tx.logValHashes === undefined) {
        return null
      }
      let logValHashes = tx.logValHashes
      if (!logValHashes) {
        logValHashes = []
      }

      if (
        tx.logPostHash === undefined ||
        tx.logPreHash === undefined ||
        tx.onChainTxHash === undefined
      ) {
        return null
      }

      return {
        logPostHash: tx.logPostHash,
        logPreHash: tx.logPreHash,
        logValHashes,
        onChainTxHash: tx.onChainTxHash,
        val,
        vmId,
        evmVal,
      }
    } else {
      return null
    }
  }

  public call(
    contractAddress: string,
    sender: string,
    data: string
  ): Promise<Uint8Array> {
    return new Promise((resolve, reject): void => {
      const params: validatorserver.CallMessageArgs = {
        contractAddress,
        data,
        sender,
      }
      this.client.request(
        'Validator.CallMessage',
        [params],
        (
          err: Error,
          error: Error,
          result: validatorserver.CallMessageReply
        ) => {
          if (err) {
            reject(err)
          } else if (error) {
            reject(error)
          } else {
            if (result.rawVal === undefined) {
              reject('call result empty')
              return
            }
            const val = ArbValue.unmarshal(result.rawVal)
            const evmVal = processLog(val as ArbValue.TupleValue)
            switch (evmVal.returnType) {
              case EVMCode.Return:
                resolve(evmVal.data)
                break
              case EVMCode.Stop:
                resolve(new Uint8Array())
                break
              default:
                reject(new Error('Call was reverted'))
                break
            }
          }
        }
      )
    })
  }

  public findLogs(filter: ethers.providers.Filter): Promise<evm.FullLogBuf[]> {
    return new Promise((resolve, reject): void => {
      const params: validatorserver.FindLogsArgs = {
        address: filter.address,
        fromHeight: convertBlockTag(filter.fromBlock),
        toHeight: convertBlockTag(filter.toBlock),
        topicGroups: convertTopics(filter.topics),
      }
      return this.client.request(
        'Validator.FindLogs',
        [params],
        (err: Error, error: Error, result: validatorserver.FindLogsReply) => {
          if (err) {
            reject(err)
          } else if (error) {
            reject(error)
          } else {
            resolve(result.logs)
          }
        }
      )
    })
  }

  public getVmID(): Promise<string> {
    const params: validatorserver.GetVMInfoArgs = {}
    return new Promise((resolve, reject): void => {
      this.client.request(
        'Validator.GetVMInfo',
        [params],
        (err: Error, error: Error, result: validatorserver.GetVMInfoReply) => {
          if (err) {
            reject(err)
          } else if (error) {
            reject(error)
          } else {
            resolve(result.vmID)
          }
        }
      )
    })
  }

  public getAssertionCount(): Promise<number> {
    const params: validatorserver.GetAssertionCountArgs = {}
    return new Promise((resolve, reject): void => {
      this.client.request(
        'Validator.GetAssertionCount',
        [params],
        (
          err: Error,
          error: Error,
          result: validatorserver.GetAssertionCountReply
        ) => {
          if (err) {
            reject(err)
          } else if (error) {
            reject(error)
          } else {
            resolve(result.assertionCount)
          }
        }
      )
    })
  }
}
