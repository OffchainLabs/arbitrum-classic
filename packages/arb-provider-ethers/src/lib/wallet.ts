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

import { ArbClient } from './client';
import { Contract } from './contract';
import { ArbProvider } from './provider';
import * as ArbValue from './value';

import * as ethers from 'ethers';

export class ArbWallet extends ethers.Signer {
    public client: ArbClient;
    public contracts: Map<string, Contract>;
    public signer: ethers.Signer;
    public provider: ArbProvider;
    public vmTracker: ethers.Contract;
    public seq: ethers.utils.BigNumber;
    public pubkey?: string;

    constructor(client: ArbClient, contracts: Map<string, Contract>, signer: ethers.Signer, provider: ArbProvider) {
        super();
        this.contracts = contracts;
        this.signer = signer;
        this.provider = provider;
        this.client = client;
        this.vmTracker = provider.vmTracker.connect(signer);
        this.seq = ethers.utils.bigNumberify(0);
        this.pubkey = undefined;
    }

    public async initialize(): Promise<void> {
        if (!this.seq.eq(ethers.utils.bigNumberify(0))) {
            return;
        }

        return this.provider.provider.getBlockNumber().then((height: number) => {
            let seq = ethers.utils.bigNumberify(height);
            for (let i = 0; i < 128; i++) {
                seq = seq.mul(2);
            }
            const timeStamp = Math.floor(Date.now());
            seq = seq.add(timeStamp);
            seq = seq.mul(2);
            this.seq = seq;
        });
    }

    public getAddress(): Promise<string> {
        return this.signer.getAddress();
    }

    public async signMessage(message: ethers.utils.Arrayish | string): Promise<string> {
        return this.signer.signMessage(message);
    }

    public async sendTransaction(
        transaction: ethers.providers.TransactionRequest,
    ): Promise<ethers.providers.TransactionResponse> {
        if (!transaction.to) {
            throw Error("Can't send transaction without destination");
        }
        const dest = await transaction.to;
        const contract = this.contracts.get(dest.toLowerCase());
        if (contract) {
            this.seq = this.seq.add(2);
            const vmId = await this.provider.getVmID();
            let encodedData = new ArbValue.TupleValue([new ArbValue.TupleValue([]), new ArbValue.IntValue(0)]);
            if (transaction.data) {
                encodedData = ArbValue.hexToSizedByteRange(await transaction.data);
            }
            const arbMsg = new ArbValue.TupleValue([
                encodedData,
                new ArbValue.IntValue(dest),
                new ArbValue.IntValue(this.seq),
            ]);
            let value = ethers.utils.bigNumberify(0);
            if (transaction.value) {
                value = ethers.utils.bigNumberify(await transaction.value); // eslint-disable-line require-atomic-updates
            }
            const args = [vmId, arbMsg.hash(), value, ethers.utils.hexZeroPad('0x00', 21)];
            const messageHash = ethers.utils.solidityKeccak256(['bytes32', 'bytes32', 'uint256', 'bytes21'], args);
            const fromAddress = await this.getAddress();
            if (value.eq(0)) {
                const messageHashBytes = ethers.utils.arrayify(messageHash);
                const sig = await this.signer.signMessage(messageHashBytes);
                if (!this.pubkey) {
                    this.pubkey = ethers.utils.recoverPublicKey(
                        ethers.utils.arrayify(ethers.utils.hashMessage(messageHashBytes)),
                        sig,
                    );
                }
                await this.client.sendMessage(arbMsg, sig, this.pubkey);
            } else {
                const blockchainTx = await this.vmTracker.sendEthMessage(vmId, ArbValue.marshal(arbMsg), {
                    value,
                });

                await blockchainTx.wait();
            }
            let txData = '';
            if (transaction.data) {
                txData = ethers.utils.hexlify(await transaction.data);
            }

            const tx = {
                data: txData,
                from: fromAddress,
                gasLimit: ethers.utils.bigNumberify(1),
                gasPrice: ethers.utils.bigNumberify(1),
                hash: messageHash,
                nonce: 0,
                to: dest,
                value: value,
                chainId: 123456789,
            };
            return this.provider._wrapTransaction(tx, messageHash);
        } else {
            return this.signer.sendTransaction(transaction);
        }
    }
}
