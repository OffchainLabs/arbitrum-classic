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
import { GlobalPendingInbox } from './GlobalPendingInbox';

import * as ethers from 'ethers';

export class ArbWallet extends ethers.Signer {
    public client: ArbClient;
    public contracts: Map<string, Contract>;
    public signer: ethers.Signer;
    public provider: ArbProvider;
    public inboxManagerCache?: GlobalPendingInbox;
    public seq: ethers.utils.BigNumber;
    public pubkey?: string;
    public channelMode: boolean;

    constructor(
        client: ArbClient,
        contracts: Map<string, Contract>,
        signer: ethers.Signer,
        provider: ArbProvider,
        channelMode: boolean,
    ) {
        super();
        this.contracts = contracts;
        this.signer = signer;
        this.provider = provider;
        this.client = client;
        this.seq = ethers.utils.bigNumberify(0);
        this.pubkey = undefined;
        this.channelMode = channelMode;
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

    public async globalInboxConn(): Promise<GlobalPendingInbox> {
        if (!this.inboxManagerCache) {
            const inboxManager = await this.provider.globalInboxConn();
            const linkedInboxManager = inboxManager.connect(this.signer);
            this.inboxManagerCache = linkedInboxManager;
            return linkedInboxManager;
        }
        return this.inboxManagerCache;
    }

    public getAddress(): Promise<string> {
        return this.signer.getAddress();
    }

    public async signMessage(message: ethers.utils.Arrayish | string): Promise<string> {
        return this.signer.signMessage(message);
    }

    public async depositERC20(
        tokenAddress: string,
        destAddress: string,
        value: number,
    ): Promise<ethers.providers.TransactionResponse> {
        const sendValue = ethers.utils.bigNumberify(value);
        const vmAddress = await this.provider.getVmID();

        const inboxManager = await this.globalInboxConn();
        console.log('vm add: ' + vmAddress);
        const blockchainTx = await inboxManager.depositERC20Message(vmAddress, tokenAddress, destAddress, sendValue);
        await blockchainTx.wait();
        console.log('vm add22: ' + vmAddress);

        return this.wrapTransaction(vmAddress, tokenAddress, destAddress, value);
    }

    public async depositERC721(
        tokenAddress: string,
        destAddress: string,
        value: number,
    ): Promise<ethers.providers.TransactionResponse> {
        const tokenValue = ethers.utils.bigNumberify(value);
        const vmAddress = await this.provider.getVmID();

        const inboxManager = await this.globalInboxConn();
        const blockchainTx = await inboxManager.depositERC721Message(vmAddress, tokenAddress, destAddress, tokenValue);
        await blockchainTx.wait();

        return this.wrapTransaction(vmAddress, tokenAddress, destAddress, value);
    }

    public async wrapTransaction(
        vmAddress: string,
        tokenAddress: string,
        destAddress: string,
        value: number,
    ): Promise<ethers.providers.TransactionResponse> {
        const fromAddress = await this.getAddress();
        const args = [vmAddress, fromAddress, destAddress, tokenAddress, value];
        const messageHash = ethers.utils.solidityKeccak256(
            ['address', 'address', 'address', 'address', 'uint256'],
            args,
        );

        const tx = {
            data: '',
            value: ethers.utils.bigNumberify(0),
            from: fromAddress,
            gasLimit: ethers.utils.bigNumberify(1),
            gasPrice: ethers.utils.bigNumberify(1),
            hash: messageHash,
            nonce: 0,
            to: destAddress,
            chainId: 1578891852042,
        };
        return this.provider._wrapTransaction(tx, messageHash);
    }

    public async depositETH(
        destAddress: string,
        transaction: ethers.providers.TransactionRequest,
    ): Promise<ethers.providers.TransactionResponse> {
        if (!transaction.to) {
            throw Error("Can't send transaction without destination");
        }

        const vmAddress = await this.provider.getVmID();

        let value = ethers.utils.bigNumberify(0);
        if (transaction.value) {
            value = ethers.utils.bigNumberify(await transaction.value); // eslint-disable-line require-atomic-updates
        }

        let txData = '';
        if (transaction.data) {
            txData = ethers.utils.hexlify(await transaction.data);
        }

        const args = [vmAddress, destAddress, value, ethers.utils.hexZeroPad('0x00', 21)];
        const messageHash = ethers.utils.solidityKeccak256(['address', 'address', 'uint256', 'bytes21'], args);
        const fromAddress = await this.getAddress();

        const inboxManager = await this.globalInboxConn();
        const blockchainTx = await inboxManager.sendEthMessage(vmAddress, destAddress, {
            value,
        });

        await blockchainTx.wait();

        const tx = {
            data: txData,
            from: fromAddress,
            gasLimit: ethers.utils.bigNumberify(1),
            gasPrice: ethers.utils.bigNumberify(1),
            hash: messageHash,
            nonce: 0,
            to: vmAddress,
            value: value,
            chainId: 1578891852042,
        };
        return this.provider._wrapTransaction(tx, messageHash);
    }

    public async sendTransactionMessage(
        value: number,
        transaction: ethers.providers.TransactionRequest,
    ): Promise<ethers.providers.TransactionResponse> {
        const vmAddress = await this.provider.getVmID();
        this.seq = this.seq.add(2);

        let encodedData = new ArbValue.TupleValue([new ArbValue.TupleValue([]), new ArbValue.IntValue(0)]);
        if (transaction.data) {
            encodedData = ArbValue.hexToSizedByteRange(await transaction.data);
        }
        const arbMsg = new ArbValue.TupleValue([
            encodedData,
            new ArbValue.IntValue(vmAddress),
            new ArbValue.IntValue(this.seq),
        ]);

        const inboxManager = await this.globalInboxConn();
        const blockchainTx = await inboxManager.sendTransactionMessage(
            vmAddress,
            this.seq,
            value,
            ArbValue.marshal(arbMsg),
        );

        await blockchainTx.wait();

        const fromAddress = await this.getAddress();
        let txData = '';
        if (transaction.data) {
            txData = ethers.utils.hexlify(await transaction.data);
        }

        const args = [vmAddress, this.seq, value, arbMsg.hash()];
        const messageHash = ethers.utils.solidityKeccak256(['address', 'uint256', 'uint256', 'bytes21'], args);

        const tx = {
            data: txData,
            from: fromAddress,
            gasLimit: ethers.utils.bigNumberify(1),
            gasPrice: ethers.utils.bigNumberify(1),
            hash: messageHash,
            nonce: 0,
            to: vmAddress,
            value: ethers.utils.bigNumberify(0),
            chainId: 1578891852042,
        };
        return this.provider._wrapTransaction(tx, messageHash);
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
            const messageHash = ethers.utils.solidityKeccak256(['address', 'bytes32', 'uint256', 'bytes21'], args);
            const fromAddress = await this.getAddress();
            if (this.channelMode && value.eq(0)) {
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
                const inboxManager = await this.globalInboxConn();
                const blockchainTx = await inboxManager.sendEthMessage(vmId, ArbValue.marshal(arbMsg), {
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
