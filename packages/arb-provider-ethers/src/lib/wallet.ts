import { ArbClient } from './client';
import * as ArbValue from './value';

import * as ethers from 'ethers';

export class ArbWallet extends ethers.Signer {
    public client: ArbClient;
    public contracts: any;
    public signer: any;
    public provider: any;
    public vmTracker: any;
    public seq: ethers.utils.BigNumber;
    public pubkey: any;

    constructor(client: ArbClient, contracts: any, signer: ethers.Signer, provider: any) {
        super();
        this.contracts = contracts;
        this.signer = signer;
        this.provider = provider;
        this.client = client;
        this.vmTracker = provider.vmTracker.connect(signer);
        this.seq = ethers.utils.bigNumberify(0);
        this.pubkey = null;
    }

    public async initialize() {
        if (this.seq.eq(ethers.utils.bigNumberify(0))) {
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

    public getAddress() {
        return this.signer.getAddress();
    }

    public async signMessage(message: ethers.utils.Arrayish | string): Promise<string> {
        return this.signer.signMessage;
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
            if (!transaction.value) {
                transaction.value = ethers.utils.bigNumberify(0); // eslint-disable-line require-atomic-updates
            }
            const args = [vmId, arbMsg.hash(), transaction.value, ethers.utils.hexZeroPad('0x00', 21)];
            const messageHash = ethers.utils.solidityKeccak256(['bytes32', 'bytes32', 'uint256', 'bytes21'], args);
            const fromAddress = await this.getAddress();
            const tx = {
                data: transaction.data,
                from: fromAddress,
                gasLimit: 1,
                gasPrice: 1,
                hash: messageHash,
                nonce: this.seq,
                to: dest,
                value: transaction.value,
            };
            if (ethers.utils.bigNumberify(await transaction.value).eq(0)) {
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
                    value: transaction.value,
                });

                await blockchainTx.wait();
            }
            return this.provider._wrapTransaction(tx, messageHash);
        } else {
            return this.signer.sendTransaction(transaction);
        }
    }
}
