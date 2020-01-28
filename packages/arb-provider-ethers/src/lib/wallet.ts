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
/* eslint-env node */
'use strict';

import { ArbClient } from './client';
import { Contract } from './contract';
import { ArbProvider } from './provider';
import * as ArbValue from './value';
import { GlobalPendingInbox } from './abi/GlobalPendingInbox';
import { ArbSysFactory } from './abi/ArbSysFactory';

import * as ethers from 'ethers';

export class ArbWallet extends ethers.Signer {
    public client: ArbClient;
    public signer: ethers.Signer;
    public provider: ArbProvider;
    public inboxManagerCache?: GlobalPendingInbox;
    public seq: ethers.utils.BigNumber;
    public pubkey?: string;
    public channelMode: boolean;

    constructor(client: ArbClient, signer: ethers.Signer, provider: ArbProvider, channelMode: boolean) {
        super();
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

    public async withdrawEthFromChain(value: ethers.utils.BigNumberish): Promise<ethers.providers.TransactionResponse> {
        const valueNum = ethers.utils.bigNumberify(value);
        const arbsys = ArbSysFactory.connect('0x0000000000000000000000000000000000000064', this);
        return arbsys.withdrawEth(await this.getAddress(), valueNum);
    }

    public async withdrawEth(): Promise<ethers.providers.TransactionResponse> {
        const inboxManager = await this.globalInboxConn();
        return inboxManager.withdrawEth();
    }

    public async withdrawERC20(erc20: string): Promise<ethers.providers.TransactionResponse> {
        const inboxManager = await this.globalInboxConn();
        return inboxManager.withdrawERC20(erc20);
    }

    public async withdrawERC721(
        erc721: string,
        tokenId: ethers.utils.BigNumberish,
    ): Promise<ethers.providers.TransactionResponse> {
        const inboxManager = await this.globalInboxConn();
        return inboxManager.withdrawERC721(erc721, tokenId);
    }

    public async depositERC20(
        to: string,
        erc20: string,
        value: ethers.utils.BigNumberish,
    ): Promise<ethers.providers.TransactionResponse> {
        const sendValue = ethers.utils.bigNumberify(value);
        const chain = await this.provider.getVmID();
        const inboxManager = await this.globalInboxConn();
        const tx = await inboxManager.depositERC20Message(chain, to, erc20, sendValue);
        return this.provider._wrapTransaction(tx, tx.hash);
    }

    public async depositERC721(
        to: string,
        tokenAddress: string,
        value: ethers.utils.BigNumberish,
    ): Promise<ethers.providers.TransactionResponse> {
        const tokenValue = ethers.utils.bigNumberify(value);
        const chain = await this.provider.getVmID();
        const inboxManager = await this.globalInboxConn();
        const tx = await inboxManager.depositERC721Message(chain, to, tokenAddress, value);
        return this.provider._wrapTransaction(tx, tx.hash);
    }

    public async depositETH(
        to: string,
        value: ethers.utils.BigNumberish,
    ): Promise<ethers.providers.TransactionResponse> {
        const valueNum = ethers.utils.bigNumberify(value);
        const chain = await this.provider.getVmID();
        const inboxManager = await this.globalInboxConn();
        const tx = await inboxManager.depositEthMessage(chain, to, { value });
        return this.provider._wrapTransaction(tx, tx.hash);
    }

    // const vmId = await this.provider.getVmID();
    // const valueNum = ethers.utils.bigNumberify(value);
    // const args = [TxType.Transaction, vmId, to, from, this.seq, value, data];
    // const messageHash = ethers.utils.solidityKeccak256(
    //     ['uint8', 'address', 'address', 'address', 'uint256', 'uint256', 'bytes'],
    //     args,
    // );

    public async sendTransactionMessage(
        to: string,
        value: ethers.utils.BigNumberish,
        data: string,
    ): Promise<ethers.providers.TransactionResponse> {
        const vmId = await this.provider.getVmID();
        const from = await this.getAddress();
        const valueNum = ethers.utils.bigNumberify(value);
        const inboxManager = await this.globalInboxConn();
        const tx = await inboxManager.sendTransactionMessage(vmId, to, this.seq, valueNum, data);
        return this.provider._wrapTransaction(tx, tx.hash);
        // this.seq = this.seq.add(2);
        // if (this.channelMode && valueNum.eq(0)) {
        //     let hash = calculateTransactionHash(vmId, to, from, this.seq, valueNum, data);
        //     const messageHashBytes = ethers.utils.arrayify(hash);
        //     const sig = await this.signer.signMessage(messageHashBytes);
        //     if (!this.pubkey) {
        //         this.pubkey = ethers.utils.recoverPublicKey(
        //             ethers.utils.arrayify(ethers.utils.hashMessage(messageHashBytes)),
        //             sig,
        //         );
        //     }
        //     await this.client.sendMessage(to, this.seq, value, data, sig, this.pubkey);
        //     const tx = {
        //         data,
        //         from,
        //         gasLimit: ethers.utils.bigNumberify(1),
        //         gasPrice: ethers.utils.bigNumberify(1),
        //         hash,
        //         nonce: 0,
        //         to: vmId,
        //         value: ethers.utils.bigNumberify(0),
        //         chainId: 1578891852042,
        //     };
        //     return this.provider._wrapTransaction(tx, hash);
        // } else {

        // }
    }

    public async sendTransaction(
        transaction: ethers.providers.TransactionRequest,
    ): Promise<ethers.providers.TransactionResponse> {
        if (!transaction.to) {
            throw Error("Can't send transaction without destination");
        }
        const to = await transaction.to;
        let data = '0x';
        if (transaction.data) {
            data = ethers.utils.hexlify(await transaction.data);
        }

        let value = ethers.utils.bigNumberify(0);
        if (transaction.value) {
            value = ethers.utils.bigNumberify(await transaction.value); // eslint-disable-line require-atomic-updates
        }
        return this.sendTransactionMessage(to, value, data);
    }
}
