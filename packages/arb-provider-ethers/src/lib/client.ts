/*
 * Copyright 2019, Offchain Labs, Inc.
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
'use strict';

import * as ArbValue from './value';

import * as ethers from 'ethers';

const fetch = require('node-fetch'); // eslint-disable-line @typescript-eslint/no-var-requires
const jaysonBrowserClient = require('jayson/lib/client/browser'); // eslint-disable-line @typescript-eslint/no-var-requires

export enum EVMCode {
    Revert = 0,
    Invalid = 1,
    Return = 2,
    Stop = 3,
    BadSequenceCode = 4,
}

interface Log {
    contractId: ethers.utils.BigNumber;
    data: Uint8Array;
    topics: ethers.utils.BigNumber[];
}

function logValToLog(val: ArbValue.Value): Log {
    const value = val as ArbValue.TupleValue;
    return {
        contractId: (value.get(0) as ArbValue.IntValue).bignum,
        data: ArbValue.sizedByteRangeToBytes(value.get(1) as ArbValue.TupleValue),
        topics: value.contents.slice(2).map(rawTopic => (rawTopic as ArbValue.IntValue).bignum),
    };
}

function stackValueToList(value: ArbValue.TupleValue): ArbValue.Value[] {
    const values = [];
    while (value.contents.length !== 0) {
        values.push(value.get(1));
        value = value.get(0) as ArbValue.TupleValue;
    }
    return values;
}

class OrigMessage {
    public data: Uint8Array;
    public calldataHash: string;
    public contractID: string;
    public sequenceNum: string;
    public timestamp: string;
    public blockHeight: ethers.utils.BigNumber;
    public txHash: string;
    public tokenType: string;
    public value: ethers.utils.BigNumber;
    public caller: string;

    constructor(value: ArbValue.TupleValue) {
        const wrappedData = value.get(0) as ArbValue.TupleValue;
        const calldata = wrappedData.get(0) as ArbValue.TupleValue;
        this.calldataHash = calldata.hash();
        this.data = ArbValue.sizedByteRangeToBytes(calldata.get(0) as ArbValue.TupleValue);
        this.contractID = ethers.utils.getAddress((calldata.get(1) as ArbValue.IntValue).bignum.toHexString());
        this.sequenceNum = (calldata.get(2) as ArbValue.IntValue).bignum.toHexString();
        this.timestamp = (wrappedData.get(1) as ArbValue.IntValue).bignum.toHexString();
        this.blockHeight = (wrappedData.get(2) as ArbValue.IntValue).bignum;
        this.txHash = (wrappedData.get(3) as ArbValue.IntValue).bignum.toHexString();
        this.tokenType = (value.get(3) as ArbValue.IntValue).bignum.toHexString();
        this.value = (value.get(2) as ArbValue.IntValue).bignum;
        this.caller = ethers.utils.getAddress((value.get(1) as ArbValue.IntValue).bignum.toHexString());
    }
}

export type EVMResult = EVMReturn | EVMRevert | EVMStop | EVMBadSequenceCode | EVMInvalid;

export class EVMReturn {
    public orig: OrigMessage;
    public data: Uint8Array;
    public logs: Log[];
    public returnType: EVMCode.Return;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.data = ArbValue.sizedByteRangeToBytes(value.get(2) as ArbValue.TupleValue);
        this.logs = stackValueToList(value.get(1) as ArbValue.TupleValue).map(logValToLog);
        this.returnType = EVMCode.Return;
    }
}

export class EVMRevert {
    public orig: OrigMessage;
    public data: Uint8Array;
    public returnType: EVMCode.Revert;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.data = ArbValue.sizedByteRangeToBytes(value.get(2) as ArbValue.TupleValue);
        this.returnType = EVMCode.Revert;
    }
}

export class EVMStop {
    public orig: OrigMessage;
    public logs: Log[];
    public returnType: EVMCode.Stop;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.logs = stackValueToList(value.get(1) as ArbValue.TupleValue).map(logValToLog);
        this.returnType = EVMCode.Stop;
    }
}

export class EVMBadSequenceCode {
    public orig: OrigMessage;
    public returnType: EVMCode.BadSequenceCode;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.returnType = EVMCode.BadSequenceCode;
    }
}

export class EVMInvalid {
    public orig: OrigMessage;
    public returnType: EVMCode.Invalid;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.returnType = EVMCode.Invalid;
    }
}

function processLog(value: ArbValue.TupleValue): EVMResult {
    const returnCode = value.get(3) as ArbValue.IntValue;
    switch (returnCode.bignum.toNumber()) {
        case EVMCode.Return:
            return new EVMReturn(value);
        case EVMCode.Revert:
            return new EVMRevert(value);
        case EVMCode.Stop:
            return new EVMStop(value);
        case EVMCode.BadSequenceCode:
            return new EVMBadSequenceCode(value);
        case EVMCode.Invalid:
            return new EVMInvalid(value);
        default:
            throw Error('processLogs Invalid EVM return code');
    }
}

interface GetVMInfoReply {
    vmID: string;
}

interface GetAssertionCountReply {
    assertionCount: number;
}

interface GetVMCreatedTxHashReply {
    vmCreatedTxHash: string;
}

interface GetMessageResultReply {
    found: boolean;
    rawVal: string;
    logPreHash: string;
    logPostHash: string;
    logValHashes: string[];
    validatorSigs: string[];
    partialHash: string;
    onChainTxHash: string;
}

interface SendMessageReply {
    txHash: string;
}

interface CallMessageReply {
    rawVal: string;
}

interface LogInfo {
    address: string;
    blockHash: string;
    blockNumber: string;
    data: string;
    logIndex: string;
    topics: string[];
    transactionIndex: string;
    transactionHash: string;
}

interface FindLogsReply {
    logs: LogInfo[];
}

function _arbClient(managerAddress: string): any {
    const callServer = (request: any, callback: any): void => {
        const options = {
            body: request, // request is a string
            headers: {
                'Content-Type': 'application/json',
            },
            method: 'POST',
        };

        fetch(managerAddress, options)
            .then((res: any) => {
                return res.text();
            })
            .then((text: string) => {
                callback(null, text);
            })
            .catch((err: Error) => {
                callback(err);
            });
    };

    return jaysonBrowserClient(callServer, {});
}

interface MessageResult {
    logPostHash: string;
    logPreHash: string;
    logValHashes: string[];
    onChainTxHash: string;
    partialHash: string;
    val: ArbValue.Value;
    validatorSigs: string[];
    vmId: string;
    evmVal: EVMResult;
}

export class ArbClient {
    public client: any;

    constructor(managerUrl: string) {
        this.client = _arbClient(managerUrl);
    }

    public async getMessageResult(txHash: string): Promise<MessageResult | null> {
        const messageResult = await new Promise<GetMessageResultReply>((resolve, reject): void => {
            this.client.request(
                'Validator.GetMessageResult',
                [
                    {
                        txHash,
                    },
                ],
                (err: Error, error: Error, result: GetMessageResultReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result);
                    }
                },
            );
        });
        if (messageResult.found) {
            const vmId = await this.getVmID();
            const val = ArbValue.unmarshal(messageResult.rawVal);
            const evmVal = processLog(val as ArbValue.TupleValue);
            let logValHashes = messageResult.logValHashes;
            if (!logValHashes) {
                logValHashes = [];
            }

            return {
                logPostHash: messageResult.logPostHash,
                logPreHash: messageResult.logPreHash,
                logValHashes,
                onChainTxHash: messageResult.onChainTxHash,
                partialHash: messageResult.partialHash,
                val,
                validatorSigs: messageResult.validatorSigs,
                vmId,
                evmVal,
            };
        } else {
            return null;
        }
    }

    public sendMessage(value: ArbValue.Value, sig: string, pubkey: string): Promise<string> {
        return this.sendRawMessage(ethers.utils.hexlify(ArbValue.marshal(value)), sig, pubkey);
    }

    public sendRawMessage(value: string, sig: string, pubkey: string): Promise<string> {
        return new Promise((resolve, reject): void => {
            this.client.request(
                'Validator.SendMessage',
                [
                    {
                        data: value,
                        pubkey,
                        signature: sig,
                    },
                ],
                (err: Error, error: Error, result: SendMessageReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result.txHash);
                    }
                },
            );
        });
    }

    public call(value: ArbValue.Value, sender: string): Promise<Uint8Array> {
        return new Promise((resolve, reject): void => {
            this.client.request(
                'Validator.CallMessage',
                [
                    {
                        data: ethers.utils.hexlify(ArbValue.marshal(value)),
                        sender,
                    },
                ],
                (err: Error, error: Error, result: CallMessageReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        const val = ArbValue.unmarshal(result.rawVal);
                        const evmVal = processLog(val as ArbValue.TupleValue);
                        switch (evmVal.returnType) {
                            case EVMCode.Return:
                                resolve(evmVal.data);
                                break;
                            case EVMCode.Stop:
                                resolve(new Uint8Array());
                                break;
                            default:
                                reject(new Error('Call was reverted'));
                                break;
                        }
                    }
                },
            );
        });
    }

    public findLogs(fromBlock: number, toBlock: number, address: string, topics: string[]): Promise<LogInfo[]> {
        return new Promise((resolve, reject): void => {
            return this.client.request(
                'Validator.FindLogs',
                [
                    {
                        address,
                        fromHeight: fromBlock,
                        toHeight: toBlock,
                        topics,
                    },
                ],
                (err: Error, error: Error, result: FindLogsReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result.logs);
                    }
                },
            );
        });
    }

    public getVmID(): Promise<string> {
        return new Promise((resolve, reject): void => {
            this.client.request('Validator.GetVMInfo', [], (err: Error, error: Error, result: GetVMInfoReply) => {
                if (err) {
                    reject(err);
                } else if (error) {
                    reject(error);
                } else {
                    resolve(result.vmID);
                }
            });
        });
    }

    public getAssertionCount(): Promise<number> {
        return new Promise((resolve, reject): void => {
            this.client.request(
                'Validator.GetAssertionCount',
                [],
                (err: Error, error: Error, result: GetAssertionCountReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result.assertionCount);
                    }
                },
            );
        });
    }

    public getVMCreatedTxHash(): Promise<string> {
        return new Promise((resolve, reject): void => {
            this.client.request(
                'Validator.GetVMCreatedTxHash',
                [],
                (err: Error, error: Error, result: GetVMCreatedTxHashReply) => {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result.vmCreatedTxHash);
                    }
                },
            );
        });
    }
}
