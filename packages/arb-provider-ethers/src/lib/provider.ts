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
import * as ethers from 'ethers';
import * as ArbValue from './value';
import { ArbClient, EVMCode } from './client';
import { ArbWallet } from './wallet';

const promisePoller = require('promise-poller').default;
const vmTrackerJson = require('./VMTracker.json');

// EthBridge event names
const EB_EVENT_VMC = 'VMCreated';
const EB_EVENT_CUA = 'ConfirmedUnanimousAssertion';
const EB_EVENT_FUA = 'FinalUnanimousAssertion';
const EB_EVENT_CDA = 'ConfirmedDisputableAssertion';

export class ArbProvider extends ethers.providers.BaseProvider {
    chainId: number;
    provider: ethers.providers.JsonRpcProvider;
    client: ArbClient;
    vmTracker: ethers.Contract;
    contracts: Map<string, any>;

    _validatorAddresses: any;
    _vmId: any;

    constructor(managerUrl: string, contracts: any, provider: ethers.providers.JsonRpcProvider) {
        super(123456789);
        this.chainId = 123456789;
        this.provider = provider;
        this.client = new ArbClient(managerUrl);
        let contractAddress = '0x5EBF59dBff8dCDa41610738634b396DfCB24A7c7';
        this.vmTracker = new ethers.Contract(contractAddress, vmTrackerJson.abi, provider);
        this.contracts = new Map<string, any>();
        for (var contract of contracts) {
            this.contracts.set(contract.address.toLowerCase(), contract);
        }
    }

    async getSigner(index: number) {
        let wallet = new ArbWallet(this.client, this.contracts, this.provider.getSigner(index), this);
        await wallet.initialize();
        return wallet;
    }

    // value: *Value
    // logPreHash: hexString
    // logPostHash: hexString
    // logValHashes: []hexString
    // Returns true if the hash of value is in logPostHash and false otherwise
    processLogsProof(value: ArbValue.Value, logPreHash: string, logPostHash: string, logValHashes: Array<string>) {
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
    // Returns true if assertionHash is signed by all validators
    async processUnanimousAssertion(partialHash: string, logPostHash: string, validatorSigs: Array<string>) {
        const vmId = await this.getVmID();
        const validatorAddresses = await this.getValidatorAddresses();
        if (validatorAddresses.length !== validatorSigs.length) {
            console.error('Expected:', validatorAddresses.length, 'signatures.\n', 'Received:', validatorSigs.length);
            return false;
        }

        let assertionHash = ethers.utils.solidityKeccak256(
            ['bytes32', 'bytes32', 'bytes32'],
            [vmId, partialHash, logPostHash],
        );

        let addresses = validatorSigs
            .map(sig =>
                ethers.utils
                    .verifyMessage(ethers.utils.arrayify(assertionHash), sig)
                    .toLowerCase()
                    .slice(2),
            )
            .sort();

        for (let i = 0; i < validatorAddresses; i++) {
            if (validatorAddresses[i] !== addresses[i]) {
                console.error('Invalid signature');
                return false;
            }
        }
        return true;
    }

    // logPostHash: hexString
    // onChainTxHash: hexString
    // Returns true if assertionHash is logged by the onChainTxHash
    async processConfirmedDisputableAssertion(logPostHash: string, onChainTxHash: string) {
        let receipt = await this.provider.waitForTransaction(onChainTxHash);
        if (!receipt.logs) {
            console.error('DisputableAssertion tx had no logs');
            return false;
        }
        let events = receipt.logs.map(l => this.vmTracker.interface.parseLog(l));
        // DisputableAssertion Event
        let cda = events.find(event => event.name === EB_EVENT_CDA);
        if (cda) {
            const vmId = await this.getVmID();
            // Check correct VM
            if (cda.values.vmId !== vmId) {
                console.error(
                    'DisputableAssertion Event is from a different VM:',
                    cda.values.vmId,
                    '\nExpected VM ID:',
                    vmId,
                );
                return false;
            }

            // Check correct logs hash
            if (cda.values.logsAccHash !== logPostHash) {
                console.error(
                    'DisputableAssertion Event on-chain logPostHash is:',
                    cda.values.logsAccHash,
                    '\nExpected:',
                    logPostHash,
                );
                return false;
            }

            // DisputableAssertion is correct
            // TODO: must wait for finality (past the re-org period)
            return true;
        } else {
            console.error('DisputableAssertion', onChainTxHash, 'not found on chain');
            return false;
        }
    }

    async getValidatorAddresses() {
        if (!this._validatorAddresses) {
            let eventTxHash = await this.client.getVMCreatedTxHash();
            let receipt = await this.provider.waitForTransaction(eventTxHash);
            if (!receipt.logs) {
                throw new Error('VMCreated Tx has no logs');
            }
            let events = receipt.logs.map(l => this.vmTracker.interface.parseLog(l));
            let vmCreatedEvent = events.find(event => event.name === EB_EVENT_VMC);
            if (!vmCreatedEvent) {
                throw new Error('VMCreated Event not found');
            }

            // Get vmId
            const vmId = await this.getVmID();
            if (vmCreatedEvent.values.vmId !== vmId) {
                console.error(
                    'VMCreated Event TxHash is from the wrong VM ID:',
                    vmCreatedEvent.values.vmId,
                    '\nExpected:',
                    vmId,
                );
                throw new Error('VMCreated Event vmId does not match');
            }

            // Cache the set of lowercase validator addresses (without "0x")
            this._validatorAddresses = vmCreatedEvent.values.validators
                .map((addr: string) => addr.toLowerCase().slice(2))
                .sort();
            console.log('Validator Addresses are:', this._validatorAddresses);
        }
        return this._validatorAddresses;
    }

    async getVmID() {
        if (!this._vmId) {
            const vmId = await this.client.getVmID();
            // Guard against race condition
            if (!this._vmId) {
                this._vmId = vmId;
                console.log('VM ID is:', vmId);
            }
        }
        return this._vmId;
    }

    async getMessageResult(txHash: string) {
        let result = await this.client.getMessageResult(txHash);
        if (!result) {
            return null;
        }
        let { data, evmVal } = result;
        let { val, logPreHash, logPostHash, logValHashes, validatorSigs, partialHash, onChainTxHash } = data;

        const vmId = await this.getVmID();
        let txHashCheck = ethers.utils.solidityKeccak256(
            ['bytes32', 'bytes32', 'uint256', 'bytes21'],
            [vmId, evmVal.orig.calldataHash, evmVal.orig.value, ethers.utils.hexDataSlice(evmVal.orig.tokenType, 11)],
        );

        // Check txHashCheck matches txHash
        if (txHash !== txHashCheck) {
            console.error('txHash did not match its pre-image', txHash, txHashCheck);
            return null;
        }

        // Step 1: prove that val is in logPostHash
        if (!this.processLogsProof(val, logPreHash, logPostHash, logValHashes)) {
            console.error('Failed to prove val is in logPostHash');
            return null;
        }

        // Step 2: prove that logPostHash is in assertion and assertion is valid
        if (validatorSigs && validatorSigs.length > 0) {
            if (!this.processUnanimousAssertion(partialHash, logPostHash, validatorSigs)) {
                return null;
            }
        } else {
            // TODO: enable disputable assertion checks
            if (!this.processConfirmedDisputableAssertion(logPostHash, onChainTxHash)) {
                return null;
            }
        }

        return {
            evmVal: evmVal,
            txHash: txHashCheck,
        };
    }

    // This should return a Promise (and may throw errors)
    // method is the method name (e.g. getBalance) and params is an
    // object with normalized values passed in, depending on the method
    perform(method: string, params: any): Promise<any> {
        // console.log("perform", method, params)
        var self = this;
        switch (method) {
            case 'getCode':
                let contract = self.contracts.get(params.address.toLowerCase());
                if (contract) {
                    return new Promise((resolve, reject) => {
                        resolve(contract.code);
                    });
                }
                break;
            case 'getBlockNumber':
                return this.client.getAssertionCount();
            case 'getTransactionReceipt':
                return this.getMessageResult(params.transactionHash).then(result => {
                    if (result) {
                        let status = 0;
                        if (
                            result.evmVal.returnType() == EVMCode.Return ||
                            result.evmVal.returnType() == EVMCode.Stop
                        ) {
                            status = 1;
                        }
                        return {
                            to: result.evmVal.orig.contractID,
                            from: result.evmVal.orig.caller,
                            transactionIndex: 0,
                            gasUsed: 1,
                            blockHash: result.txHash,
                            transactionHash: result.txHash,
                            logs: [],
                            blockNumber: result.evmVal.orig.blockHeight,
                            confirmations: 1000,
                            cumulativeGasUsed: 1,
                            status: status,
                        };
                    } else {
                        return null;
                    }
                });
            case 'getTransaction':
                var getMessageRequest = () => {
                    return self.getMessageResult(params.transactionHash).then(result => {
                        if (result) {
                            return {
                                hash: result.txHash,
                                blockHash: result.txHash,
                                blockNumber: result.evmVal.orig.blockHeight,
                                transactionIndex: 0,
                                confirmations: 1000,
                                from: result.evmVal.orig.caller,
                                gasPrice: 1,
                                gasLimit: 1,
                                to: result.evmVal.orig.contractID,
                                cumulativeGasUsed: 1,
                                value: result.evmVal.orig.value,
                                nonce: 0,
                                data: result.evmVal.orig.data,
                                status:
                                    result.evmVal.returnType() == EVMCode.Return ||
                                    result.evmVal.returnType() == EVMCode.Stop,
                            };
                        } else {
                            return null;
                        }
                    });
                };
                return promisePoller({
                    taskFn: getMessageRequest,
                    interval: 100,
                    shouldContinue: (reason: any, value: any) => {
                        if (reason) {
                            return true;
                        } else if (value) {
                            return false;
                        } else {
                            return true;
                        }
                    },
                });
            case 'getLogs':
                return this.client.findLogs(
                    params.filter.fromBlock,
                    params.filter.toBlock,
                    params.filter.address,
                    params.filter.topics,
                );
        }
        let forwardResponse = self.provider.perform(method, params);
        console.log('Forwarding query to provider', method, forwardResponse);
        return forwardResponse;
    }

    async call(
        transaction: ethers.providers.TransactionRequest,
        blockTag?: ethers.providers.BlockTag | Promise<ethers.providers.BlockTag>,
    ) {
        if (!transaction.to) {
            throw Error('Cannot create call without a destination');
        }
        let dest = await transaction.to;
        let contractData = this.contracts.get(dest.toLowerCase());
        if (contractData) {
            var maxSeq = ethers.utils.bigNumberify(2);
            for (var i = 0; i < 255; i++) {
                maxSeq = maxSeq.mul(2);
            }
            maxSeq = maxSeq.sub(2);
            let txData = new ArbValue.TupleValue([new ArbValue.TupleValue([]), new ArbValue.IntValue(0)]);
            if (transaction.data) {
                txData = ArbValue.hexToSizedByteRange(await transaction.data);
            }
            let arbMsg = new ArbValue.TupleValue([txData, new ArbValue.IntValue(dest), new ArbValue.IntValue(maxSeq)]);
            let sender = await this.provider.getSigner(0).getAddress();
            return this.client.call(arbMsg, sender);
        } else {
            return this.provider.call(transaction);
        }
    }
}
