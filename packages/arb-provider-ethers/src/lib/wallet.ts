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

export enum TxType {
    Transaction = 0,
    DepositEth = 1,
    DepositERC20 = 2,
    DepositERC721 = 3,
}

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
        to: string,
        erc20: string,
        value: ethers.utils.BigNumberish,
    ): Promise<ethers.providers.TransactionResponse> {
        const sendValue = ethers.utils.bigNumberify(value);
        const chain = await this.provider.getVmID();
        const valueNum = ethers.utils.bigNumberify(value);

        const inboxManager = await this.globalInboxConn();
        console.log('vm add: ' + chain);
        const blockchainTx = await inboxManager.depositERC20Message(chain, to, erc20, sendValue);
        await blockchainTx.wait();
        console.log('vm add22: ' + chain);

        return this.wrapTransaction(TxType.DepositERC20, chain, erc20, to, valueNum);
    }

    public async depositERC721(
        to: string,
        tokenAddress: string,
        value: ethers.utils.BigNumberish,
    ): Promise<ethers.providers.TransactionResponse> {
        const tokenValue = ethers.utils.bigNumberify(value);
        const chain = await this.provider.getVmID();
        const valueNum = ethers.utils.bigNumberify(value);

        const inboxManager = await this.globalInboxConn();
        const blockchainTx = await inboxManager.depositERC721Message(chain, tokenAddress, to, tokenValue);
        await blockchainTx.wait();

        return this.wrapTransaction(TxType.DepositERC721, chain, tokenAddress, to, valueNum);
    }

    private async wrapTransaction(
        messageType: TxType,
        chain: string,
        tokenAddress: string,
        to: string,
        value: ethers.utils.BigNumber,
    ): Promise<ethers.providers.TransactionResponse> {
        const fromAddress = await this.getAddress();
        const args = [messageType, chain, fromAddress, to, tokenAddress, value];
        const messageHash = ethers.utils.solidityKeccak256(
            ['uint8', 'address', 'address', 'address', 'address', 'uint256'],
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
            to: to,
            chainId: 1578891852042,
        };
        return this.provider._wrapTransaction(tx, messageHash);
    }

    public async depositETH(
        to: string,
        value: ethers.utils.BigNumberish,
    ): Promise<ethers.providers.TransactionResponse> {
        const chain = await this.provider.getVmID();
        const valueNum = ethers.utils.bigNumberify(value);
        const fromAddress = await this.getAddress();

        const args = [TxType.DepositEth, chain, to, fromAddress, value];
        const messageHash = ethers.utils.solidityKeccak256(['uint8', 'address', 'address', 'address', 'uint256'], args);

        const inboxManager = await this.globalInboxConn();
        const blockchainTx = await inboxManager.sendEthMessage(chain, to, { value });

        await blockchainTx.wait();

        const tx = {
            data: '',
            from: fromAddress,
            gasLimit: ethers.utils.bigNumberify(1),
            gasPrice: ethers.utils.bigNumberify(1),
            hash: messageHash,
            nonce: 0,
            to: chain,
            value: valueNum,
            chainId: 1578891852042,
        };
        return this.provider._wrapTransaction(tx, messageHash);
    }

    public async sendTransactionMessage(
        to: string,
        value: ethers.utils.BigNumberish,
        data: ArbValue.Value,
    ): Promise<ethers.providers.TransactionResponse> {
        const from = await this.getAddress();
        this.seq = this.seq.add(2);
        const vmId = await this.provider.getVmID();
        const valueNum = ethers.utils.bigNumberify(value);
        const args = [TxType.Transaction, vmId, to, from, this.seq, value, data.hash()];
        const messageHash = ethers.utils.solidityKeccak256(
            ['uint8', 'address', 'address', 'address', 'uint256', 'uint256', 'bytes32'],
            args,
        );
        const fromAddress = await this.getAddress();
        if (this.channelMode && valueNum.eq(0)) {
            const messageHashBytes = ethers.utils.arrayify(messageHash);
            const sig = await this.signer.signMessage(messageHashBytes);
            if (!this.pubkey) {
                this.pubkey = ethers.utils.recoverPublicKey(
                    ethers.utils.arrayify(ethers.utils.hashMessage(messageHashBytes)),
                    sig,
                );
            }
            await this.client.sendMessage(to, this.seq, value, data, sig, this.pubkey);
        } else {
            const inboxManager = await this.globalInboxConn();
            const blockchainTx = await inboxManager.sendTransactionMessage(
                vmId,
                to,
                this.seq,
                valueNum,
                ArbValue.marshal(data),
            );

            await blockchainTx.wait();
        }

        const tx = {
            data: ethers.utils.hexlify(ArbValue.marshal(data)),
            from: fromAddress,
            gasLimit: ethers.utils.bigNumberify(1),
            gasPrice: ethers.utils.bigNumberify(1),
            hash: messageHash,
            nonce: 0,
            to: vmId,
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
        const to = await transaction.to;
        let encodedData = new ArbValue.TupleValue([new ArbValue.TupleValue([]), new ArbValue.IntValue(0)]);
        if (transaction.data) {
            encodedData = ArbValue.hexToSizedByteRange(await transaction.data);
        }
        let value = ethers.utils.bigNumberify(0);
        if (transaction.value) {
            value = ethers.utils.bigNumberify(await transaction.value); // eslint-disable-line require-atomic-updates
        }
        return this.sendTransactionMessage(to, value, encodedData);
    }
}
