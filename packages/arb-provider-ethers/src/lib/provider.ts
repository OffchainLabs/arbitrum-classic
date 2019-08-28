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

import { ArbClient, EVMCode, EVMResult } from './client';
import * as ArbValue from './value';
import { ArbWallet } from './wallet';
import { Contract } from './contract';

import * as ethers from 'ethers';

const promisePoller = require('promise-poller').default;

import vmTrackerJson from './VMTracker.json';

function sleep(ms: number): Promise<void> {
    return new Promise((resolve: any): void => {
        setTimeout(resolve, ms);
    });
}

// EthBridge event names
const EB_EVENT_VMC = 'VMCreated';
const EB_EVENT_CDA = 'ConfirmedDisputableAssertion';

interface MessageResult {
    evmVal: EVMResult;
    txHash: string;
}

interface Message {
    value: string;
    sig: string;
    pubkey: string;
}

export class ArbProvider extends ethers.providers.BaseProvider {
    public chainId: number;
    public provider: ethers.providers.JsonRpcProvider;
    public client: ArbClient;
    public vmTracker: ethers.Contract;
    public contracts: Map<string, Contract>;

    private validatorAddressesCache?: string[];
    private vmIdCache?: string;

    constructor(managerUrl: string, contracts: Contract[], provider: ethers.providers.JsonRpcProvider) {
        super(123456789);
        this.chainId = 123456789;
        this.provider = provider;
        this.client = new ArbClient(managerUrl);
        const contractAddress = '0x5EBF59dBff8dCDa41610738634b396DfCB24A7c7';
        this.vmTracker = new ethers.Contract(contractAddress, vmTrackerJson.abi, provider);
        this.contracts = new Map<string, Contract>();
        for (const contract of contracts) {
            this.contracts.set(contract.address.toLowerCase(), contract);
        }
    }

    public async getSigner(index: number): Promise<ArbWallet> {
        const wallet = new ArbWallet(this.client, this.contracts, this.provider.getSigner(index), this);
        await wallet.initialize();
        return wallet;
    }

    public async getValidatorAddresses(): Promise<string[]> {
        if (!this.validatorAddressesCache) {
            const eventTxHash = await this.client.getVMCreatedTxHash();
            const receipt = await this.provider.waitForTransaction(eventTxHash);
            if (!receipt.logs) {
                throw new Error('VMCreated Tx has no logs');
            }
            const events = receipt.logs.map(l => this.vmTracker.interface.parseLog(l));
            const vmCreatedEvent = events.find(event => event.name === EB_EVENT_VMC);
            if (!vmCreatedEvent) {
                throw new Error('VMCreated Event not found');
            }

            // Get vmId
            const vmId = await this.getVmID();
            if (vmCreatedEvent.values.vmId !== vmId) {
                throw new Error(
                    'VMCreated Event TxHash is from the wrong VM ID:' +
                        vmCreatedEvent.values.vmId +
                        '\nExpected:' +
                        vmId,
                );
            }

            const addresses = vmCreatedEvent.values.validators
                .map((addr: string) => addr.toLowerCase().slice(2))
                .sort();

            // Cache the set of lowercase validator addresses (without "0x")
            this.validatorAddressesCache = addresses;
            return addresses;
        }
        return this.validatorAddressesCache;
    }

    public async sendMessages(messages: Message[]): Promise<string> {
        let txHash: Promise<string> = new Promise<string>((): string => '');
        for (const message of messages) {
            txHash = this.client.sendRawMessage(message.value, message.sig, message.pubkey);
            await sleep(1);
        }
        return txHash;
    }

    public async getVmID(): Promise<string> {
        if (!this.vmIdCache) {
            const vmId = await this.client.getVmID();
            // Guard against race condition
            if (!this.vmIdCache) {
                this.vmIdCache = vmId;
            }
        }
        return this.vmIdCache;
    }

    public async getMessageResult(txHash: string): Promise<MessageResult | null> {
        const result = await this.client.getMessageResult(txHash);
        if (!result) {
            return null;
        }
        const {
            val,
            logPreHash,
            logPostHash,
            logValHashes,
            validatorSigs,
            partialHash,
            onChainTxHash,
            evmVal,
        } = result;

        const vmId = await this.getVmID();
        const txHashCheck = ethers.utils.solidityKeccak256(
            ['bytes32', 'bytes32', 'uint256', 'bytes21'],
            [
                vmId,
                evmVal.orig.calldataHash,
                evmVal.orig.value,
                ethers.utils.hexZeroPad(ethers.utils.hexDataSlice(evmVal.orig.tokenType, 21), 21),
            ],
        );

        // Check txHashCheck matches txHash
        if (txHash !== txHashCheck) {
            throw Error('txHash did not match its pre-image ' + txHash + ' ' + txHashCheck);
        }

        // Step 1: prove that val is in logPostHash
        if (!this.processLogsProof(val, logPreHash, logPostHash, logValHashes)) {
            throw Error('Failed to prove val is in logPostHash');
        }

        // Step 2: prove that logPostHash is in assertion and assertion is valid
        if (validatorSigs && validatorSigs.length > 0) {
            this.processUnanimousAssertion(partialHash, logPostHash, validatorSigs);
        } else {
            this.processConfirmedDisputableAssertion(logPostHash, onChainTxHash);
        }

        return {
            evmVal,
            txHash: txHashCheck,
        };
    }

    // This should return a Promise (and may throw errors)
    // method is the method name (e.g. getBalance) and params is an
    // object with normalized values passed in, depending on the method
    public async perform(method: string, params: any): Promise<any> {
        // console.log('perform', method, params);
        switch (method) {
            case 'getCode': {
                const contract = this.contracts.get(params.address.toLowerCase());
                if (contract) {
                    return contract.code;
                }
                break;
            }
            case 'getBlockNumber': {
                return this.client.getAssertionCount();
            }
            case 'getTransactionReceipt': {
                const result = await this.getMessageResult(params.transactionHash);
                if (result) {
                    let status = 0;
                    if (result.evmVal.returnType === EVMCode.Return || result.evmVal.returnType === EVMCode.Stop) {
                        status = 1;
                    }
                    return {
                        blockHash: result.txHash,
                        blockNumber: result.evmVal.orig.blockHeight.toNumber(),
                        confirmations: 1000,
                        cumulativeGasUsed: ethers.utils.bigNumberify(1),
                        from: result.evmVal.orig.caller,
                        gasUsed: ethers.utils.bigNumberify(1),
                        logs: [],
                        status,
                        to: result.evmVal.orig.contractID,
                        transactionHash: result.txHash,
                        transactionIndex: 0,
                        byzantium: true,
                    } as ethers.providers.TransactionReceipt;
                } else {
                    return null;
                }
            }
            case 'getTransaction': {
                const getMessageRequest = async (): Promise<ethers.providers.TransactionResponse | null> => {
                    const result = await this.getMessageResult(params.transactionHash);
                    if (result) {
                        const tx = {
                            blockHash: result.txHash,
                            blockNumber: result.evmVal.orig.blockHeight.toNumber(),
                            confirmations: 1000,
                            data: ethers.utils.hexlify(result.evmVal.orig.data),
                            from: result.evmVal.orig.caller,
                            gasLimit: ethers.utils.bigNumberify(1),
                            gasPrice: ethers.utils.bigNumberify(1),
                            hash: result.txHash,
                            nonce: 0,
                            to: result.evmVal.orig.contractID,
                            value: result.evmVal.orig.value,
                            chainId: 123456789,
                        } as ethers.providers.TransactionResponse;
                        return this.provider._wrapTransaction(tx);
                    } else {
                        return null;
                    }
                };
                return promisePoller({
                    interval: 100,
                    shouldContinue: (reason?: any, value?: any): boolean => {
                        if (reason) {
                            return true;
                        } else if (value) {
                            return false;
                        } else {
                            return true;
                        }
                    },
                    taskFn: getMessageRequest,
                });
            }
            case 'getLogs': {
                return this.client.findLogs(
                    params.filter.fromBlock,
                    params.filter.toBlock,
                    params.filter.address,
                    params.filter.topics,
                );
            }
        }
        const forwardResponse = this.provider.perform(method, params);
        // console.log('Forwarding query to provider', method, forwardResponse);
        return forwardResponse;
    }

    public async call(
        transaction: ethers.providers.TransactionRequest,
        blockTag?: ethers.providers.BlockTag | Promise<ethers.providers.BlockTag>,
    ): Promise<string> {
        if (!transaction.to) {
            throw Error('Cannot create call without a destination');
        }
        const dest = await transaction.to;
        const contractData = this.contracts.get(dest.toLowerCase());
        if (contractData) {
            let maxSeq = ethers.utils.bigNumberify(2);
            for (let i = 0; i < 255; i++) {
                maxSeq = maxSeq.mul(2);
            }
            maxSeq = maxSeq.sub(2);
            let txData = new ArbValue.TupleValue([new ArbValue.TupleValue([]), new ArbValue.IntValue(0)]);
            if (transaction.data) {
                txData = ArbValue.hexToSizedByteRange(await transaction.data);
            }
            const arbMsg = new ArbValue.TupleValue([
                txData,
                new ArbValue.IntValue(dest),
                new ArbValue.IntValue(maxSeq),
            ]);
            const sender = await this.provider.getSigner(0).getAddress();
            const resultData = await this.client.call(arbMsg, sender);
            return ethers.utils.hexlify(resultData);
        } else {
            return this.provider.call(transaction);
        }
    }

    // value: *Value
    // logPreHash: hexString
    // logPostHash: hexString
    // logValHashes: []hexString
    // Returns true if the hash of value is in logPostHash and false otherwise
    private processLogsProof(
        value: ArbValue.Value,
        logPreHash: string,
        logPostHash: string,
        logValHashes: string[],
    ): boolean {
        const startHash = ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [logPreHash, value.hash()]);
        const checkHash = logValHashes.reduce(
            (acc, hash) => ethers.utils.solidityKeccak256(['bytes32', 'bytes32'], [acc, hash]),
            startHash,
        );

        return logPostHash === checkHash;
    }

    // partialHash: hexString
    // logPostHash: hexString
    // validatorSigs: []hexString
    // Throws error if assertionHash is not signed by all validators
    private async processUnanimousAssertion(
        partialHash: string,
        logPostHash: string,
        validatorSigs: string[],
    ): Promise<void> {
        const vmId = await this.getVmID();
        const validatorAddresses = await this.getValidatorAddresses();
        if (validatorAddresses.length !== validatorSigs.length) {
            throw Error('Expected: ' + validatorAddresses.length + ' signatures.\nReceived: ' + validatorSigs.length);
        }

        const assertionHash = ethers.utils.solidityKeccak256(
            ['bytes32', 'bytes32', 'bytes32'],
            [vmId, partialHash, logPostHash],
        );

        const addresses = validatorSigs
            .map(sig =>
                ethers.utils
                    .verifyMessage(ethers.utils.arrayify(assertionHash), sig)
                    .toLowerCase()
                    .slice(2),
            )
            .sort();

        for (let i = 0; i < validatorAddresses.length; i++) {
            if (validatorAddresses[i] !== addresses[i]) {
                throw Error('Invalid signature');
            }
        }
    }

    // logPostHash: hexString
    // onChainTxHash: hexString
    // Returns true if assertionHash is logged by the onChainTxHash
    private async processConfirmedDisputableAssertion(logPostHash: string, onChainTxHash: string): Promise<void> {
        const receipt = await this.provider.waitForTransaction(onChainTxHash);
        if (!receipt.logs) {
            throw Error('DisputableAssertion tx had no logs');
        }
        const events = receipt.logs.map(l => this.vmTracker.interface.parseLog(l));
        // DisputableAssertion Event
        const cda = events.find(event => event.name === EB_EVENT_CDA);
        if (!cda) {
            throw Error('DisputableAssertion ' + onChainTxHash + ' not found on chain');
        }
        const vmId = await this.getVmID();
        // Check correct VM
        if (cda.values.vmId !== vmId) {
            throw Error(
                'DisputableAssertion Event is from a different VM: ' + cda.values.vmId + '\nExpected VM ID: ' + vmId,
            );
        }

        // Check correct logs hash
        if (cda.values.logsAccHash !== logPostHash) {
            throw Error(
                'DisputableAssertion Event on-chain logPostHash is: ' +
                    cda.values.logsAccHash +
                    '\nExpected: ' +
                    logPostHash,
            );
        }

        // DisputableAssertion is correct
        // TODO: must wait for finality (past the re-org period)
    }
}
