import * as ArbValue from './value';
import * as ethers from 'ethers';
const jaysonBrowserClient = require('jayson/lib/client/browser');
const fetch = require('node-fetch');

export enum EVMCode {
    Revert = 0,
    Invalid = 1,
    Return = 2,
    Stop = 3,
    BadSequenceCode = 4,
}

function logValToLog(val: ArbValue.Value) {
    let value = val as ArbValue.TupleValue;
    return {
        contractId: (value.get(0) as ArbValue.IntValue).bignum,
        data: ArbValue.sizedByteRangeToBytes(value.get(1) as ArbValue.TupleValue),
        topics: value.contents.slice(2).map(val => (val as ArbValue.IntValue).bignum),
    };
}

function stackValueToList(value: ArbValue.TupleValue) {
    let values = [];
    while (value.contents.length != 0) {
        values.push(value.get(1));
        value = value.get(0) as ArbValue.TupleValue;
    }
    return values;
}

class OrigMessage {
    data: Uint8Array;
    calldataHash: string;
    contractID: string;
    sequenceNum: string;
    timestamp: string;
    blockHeight: string;
    txHash: string;
    tokenType: string;
    value: string;
    caller: string;

    constructor(value: ArbValue.TupleValue) {
        let wrappedData = value.get(0) as ArbValue.TupleValue;
        let calldata = wrappedData.get(0) as ArbValue.TupleValue;
        this.calldataHash = calldata.hash();
        this.data = ArbValue.sizedByteRangeToBytes(calldata.get(0) as ArbValue.TupleValue);
        this.contractID = ethers.utils.hexDataSlice((calldata.get(1) as ArbValue.IntValue).bignum.toHexString(), 12);
        this.sequenceNum = (calldata.get(2) as ArbValue.IntValue).bignum.toHexString();
        this.timestamp = (wrappedData.get(1) as ArbValue.IntValue).bignum.toHexString();
        this.blockHeight = (wrappedData.get(2) as ArbValue.IntValue).bignum.toHexString();
        this.txHash = (wrappedData.get(3) as ArbValue.IntValue).bignum.toHexString();
        this.tokenType = (value.get(3) as ArbValue.IntValue).bignum.toHexString();
        this.value = (value.get(2) as ArbValue.IntValue).bignum.toHexString();
        this.caller = ethers.utils.hexDataSlice((value.get(1) as ArbValue.IntValue).bignum.toHexString(), 12);
    }
}

export interface EVMResult {
    orig: OrigMessage;

    returnType(): EVMCode;
}

class EVMReturn {
    orig: OrigMessage;
    data: Uint8Array;
    logs: any;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.data = ArbValue.sizedByteRangeToBytes(value.get(2) as ArbValue.TupleValue);
        this.logs = stackValueToList(value.get(1) as ArbValue.TupleValue).map(logValToLog);
    }

    returnType(): EVMCode {
        return EVMCode.Return;
    }
}

class EVMRevert {
    orig: OrigMessage;
    data: Uint8Array;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.data = ArbValue.sizedByteRangeToBytes(value.get(2) as ArbValue.TupleValue);
    }

    returnType(): EVMCode {
        return EVMCode.Revert;
    }
}

class EVMStop {
    orig: OrigMessage;
    logs: any;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
        this.logs = stackValueToList(value.get(1) as ArbValue.TupleValue).map(logValToLog);
    }

    returnType(): EVMCode {
        return EVMCode.Stop;
    }
}

class EVMBadSequenceCode {
    orig: OrigMessage;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
    }

    returnType(): EVMCode {
        return EVMCode.BadSequenceCode;
    }
}

class EVMInvalid {
    orig: OrigMessage;

    constructor(value: ArbValue.TupleValue) {
        this.orig = new OrigMessage(value.get(0) as ArbValue.TupleValue);
    }

    returnType(): EVMCode {
        return EVMCode.Invalid;
    }
}

function processLog(value: ArbValue.TupleValue): EVMResult {
    let returnCode = value.get(3) as ArbValue.IntValue;
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
            throw 'processLogs Invalid EVM return code';
    }
}

interface GetVMInfoReply {
    vmID: string;
}

interface GetAssertionCountReply {
    assertionCount: number;
}

interface GetMessageResultReply {
    found: boolean;
    rawVal: string;
    logPreHash: string;
    logPostHash: string;
    logValHashes: Array<string>;
    validatorSigs: Array<string>;
    partialHash: string;
    onChainTxHash: string;
}

interface SendMessageReply {
    hash: string;
}

interface CallMessageReply {
    ReturnVal: string;
    Success: boolean;
}

interface LogInfo {
    address: string;
    blockHash: string;
    blockNumber: string;
    data: string;
    logIndex: string;
    topics: Array<string>;
    transactionIndex: string;
    transactionHash: string;
}

interface FindLogsReply {
    logs: Array<LogInfo>;
}

function _arbClient(managerAddress: string) {
    var callServer = function(request: any, callback: any) {
        var options = {
            method: 'POST',
            body: request, // request is a string
            headers: {
                'Content-Type': 'application/json',
            },
        };

        fetch(managerAddress, options)
            .then(function(res: any) {
                return res.text();
            })
            .then(function(text: any) {
                callback(null, text);
            })
            .catch(function(err: any) {
                callback(err);
            });
    };

    return jaysonBrowserClient(callServer, {});
}

export class ArbClient {
    client: any;

    constructor(managerUrl: string) {
        this.client = _arbClient(managerUrl);
    }

    async getMessageResult(txHash: string) {
        let self = this;
        let result = await new Promise<GetMessageResultReply>((resolve, reject) => {
            self.client.request(
                'Validator.GetMessageResult',
                [
                    {
                        txHash: txHash,
                    },
                ],
                function(err: Error, error: Error, result: GetMessageResultReply) {
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
        if (result.found) {
            let vmId = await self.getVmID();
            let val = ArbValue.unmarshal(result.rawVal);
            let evmVal = processLog(val as ArbValue.TupleValue);

            let data = {
                vmId: vmId,
                val: val,
                logPreHash: result.logPreHash,
                logPostHash: result.logPostHash,
                logValHashes: result.logValHashes,
                validatorSigs: result.validatorSigs,
                partialHash: result.partialHash,
                onChainTxHash: result.onChainTxHash,
            };

            return {
                data: data,
                evmVal: evmVal,
            };
        } else {
            return null;
        }
    }

    sendMessage(value: ArbValue.Value, sig: string, pubkey: string): Promise<string> {
        let self = this;
        return new Promise(function(resolve, reject) {
            self.client.request(
                'Validator.SendMessage',
                [
                    {
                        data: ArbValue.marshal(value),
                        signature: sig,
                        pubkey: pubkey,
                    },
                ],
                function(err: Error, error: Error, result: SendMessageReply) {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        resolve(result.hash);
                    }
                },
            );
        });
    }

    call(value: ArbValue.Value, sender: string): Promise<string> {
        let self = this;
        return new Promise(function(resolve, reject) {
            self.client.request(
                'Validator.CallMessage',
                [
                    {
                        data: ArbValue.marshal(value),
                        sender: sender,
                    },
                ],
                function(err: Error, error: Error, result: CallMessageReply) {
                    if (err) {
                        reject(err);
                    } else if (error) {
                        reject(error);
                    } else {
                        if (result.Success) {
                            resolve(result.ReturnVal);
                        } else {
                            reject(new Error('Call was reverted'));
                        }
                    }
                },
            );
        });
    }

    findLogs(fromBlock: number, toBlock: number, address: string, topics: Array<string>): Promise<Array<LogInfo>> {
        let self = this;
        return new Promise(function(resolve, reject) {
            return self.client.request(
                'Validator.FindLogs',
                [
                    {
                        fromHeight: fromBlock,
                        toHeight: toBlock,
                        address: address,
                        topics: topics,
                    },
                ],
                function(err: Error, error: Error, result: FindLogsReply) {
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

    getVmID(): Promise<string> {
        let self = this;
        return new Promise(function(resolve, reject) {
            self.client.request('Validator.GetVMInfo', [], function(err: Error, error: Error, result: GetVMInfoReply) {
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

    getAssertionCount(): Promise<number> {
        let self = this;
        return new Promise(function(resolve, reject) {
            self.client.request('Validator.GetAssertionCount', [], function(
                err: Error,
                error: Error,
                result: GetAssertionCountReply,
            ) {
                if (err) {
                    reject(err);
                } else if (error) {
                    reject(error);
                } else {
                    resolve(result.assertionCount);
                }
            });
        });
    }

    getVMCreatedTxHash(): Promise<string> {
        let self = this;
        return new Promise(function(resolve, reject) {
            self.client.request('Validator.GetVMCreatedTxHash', [], function(err: Error, error: Error, result: any) {
                if (err) {
                    reject(err);
                } else if (error) {
                    reject(error);
                } else {
                    resolve(result.vmCreatedTxHash);
                }
            });
        });
    }
}
