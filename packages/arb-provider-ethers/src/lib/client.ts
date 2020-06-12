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
import { EVMCode, processLog } from './message'

import * as ethers from 'ethers'

import { evm } from './abi/evm.evm.d'
import { validatorserver } from './abi/validatorserver.server.d'

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

export interface NodeInfo {
  nodeHash: string
  nodeHeight: number
  l1TxHash?: string
  l1Confirmations?: number
}

export interface AVMProof {
  logPostHash: string
  logPreHash: string
  logValHashes: string[]
}

interface RawMessageResult {
  val: ArbValue.Value
  proof?: AVMProof
  nodeInfo: NodeInfo
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

function extractAVMProof(proof?: evm.AVMLogProof): AVMProof | undefined {
  if (proof === undefined) {
    return undefined
  }

  let logValHashes = proof.logValHashes
  if (!logValHashes) {
    logValHashes = []
  }

  if (proof.logPostHash === undefined || proof.logPreHash === undefined) {
    return undefined
  }
  return {
    logPostHash: proof.logPostHash,
    logPreHash: proof.logPreHash,
    logValHashes,
  }
}

function extractNodeInfo(nodeInfo?: evm.NodeLocation): NodeInfo | undefined {
  if (
    nodeInfo === undefined ||
    nodeInfo.nodeHash === undefined ||
    nodeInfo.nodeHeight === undefined
  ) {
    return undefined
  }
  return {
    nodeHash: nodeInfo.nodeHash,
    nodeHeight: nodeInfo.nodeHeight,
    l1TxHash: nodeInfo.l1TxHash,
  }
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

  public async getMessageResult(
    txHash: string
  ): Promise<RawMessageResult | null> {
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
      if (tx.rawVal === undefined) {
        return null
      }
      const nodeInfo = extractNodeInfo(tx.location)
      if (nodeInfo === undefined) {
        return null
      }

      const val = ArbValue.unmarshal(tx.rawVal)
      return {
        val,
        nodeInfo: nodeInfo,
        proof: extractAVMProof(tx.proof),
      }
    } else {
      return null
    }
  }

  private _call(
    callFunc: string,
    contractAddress: string,
    sender: string | undefined,
    data: string
  ): Promise<Uint8Array> {
    return new Promise((resolve, reject): void => {
      const params: validatorserver.CallMessageArgs = {
        contractAddress,
        data,
        sender,
      }
      this.client.request(
        callFunc,
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

  public call(
    contractAddress: string,
    sender: string | undefined,
    data: string
  ): Promise<Uint8Array> {
    return this._call('Validator.CallMessage', contractAddress, sender, data)
  }

  public pendingCall(
    contractAddress: string,
    sender: string | undefined,
    data: string
  ): Promise<Uint8Array> {
    return this._call('Validator.PendingCall', contractAddress, sender, data)
  }

  public findLogs(filter: ethers.providers.Filter): Promise<evm.FullLogBuf[]> {
    return new Promise((resolve, reject): void => {
      const addresses: string[] = []
      if (filter.address !== undefined) {
        addresses.push(filter.address)
      }

      const params: validatorserver.FindLogsArgs = {
        addresses,
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

  private async getNodeLocation(
    funcType: string
  ): Promise<NodeInfo | undefined> {
    const params: validatorserver.GetLatestNodeLocationArgs = {}
    return new Promise((resolve, reject): void => {
      this.client.request(
        funcType,
        [params],
        (
          err: Error,
          error: Error,
          result: validatorserver.GetLatestNodeLocationReply
        ) => {
          if (err) {
            reject(err)
          } else if (error) {
            reject(error)
          } else {
            resolve(extractNodeInfo(result.location))
          }
        }
      )
    })
  }

  public async getLatestNodeLocation(): Promise<NodeInfo> {
    const nodeInfo = await this.getNodeLocation(
      'Validator.GetLatestNodeLocation'
    )
    if (!nodeInfo) {
      throw Error('node location not set')
    }
    return nodeInfo
  }

  public async getLatestPendingNodeLocation(): Promise<NodeInfo> {
    const nodeInfo = await this.getNodeLocation(
      'Validator.GetLatestPendingNodeLocation'
    )
    if (nodeInfo === undefined) {
      return this.getLatestNodeLocation()
    }
    return nodeInfo
  }
}
