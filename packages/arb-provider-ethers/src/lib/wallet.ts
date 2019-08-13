import * as ArbValue from './value';
import { ArbClient } from './client';
import * as ethers from 'ethers';

export class ArbWallet extends ethers.Signer {
    client: ArbClient;
    contracts: any;
    signer: any;
    provider: any;
    vmTracker: any;
    seq: ethers.utils.BigNumber;
    pubkey: any;

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

    async initialize() {
        if (this.seq.eq(ethers.utils.bigNumberify(0))) {
            return;
        }

        return this.provider.provider.getBlockNumber().then((height: number) => {
            var seq = ethers.utils.bigNumberify(height);
            for (var i = 0; i < 128; i++) {
                seq = seq.mul(2);
            }
            var timeStamp = Math.floor(Date.now());
            seq = seq.add(timeStamp);
            seq = seq.mul(2);
            this.seq = seq;
        });
    }

    getAddress() {
        return this.signer.getAddress();
    }

    async signMessage(message: ethers.utils.Arrayish | string): Promise<string> {
        return this.signer.signMessage;
    }

    async sendTransaction(
        transaction: ethers.providers.TransactionRequest,
    ): Promise<ethers.providers.TransactionResponse> {
        let self = this;
        if (!transaction.to) {
            throw Error("Can't send transaction without destination");
        }
        let dest = await transaction.to;
        if (self.contracts[dest.toLowerCase()]) {
            self.seq = self.seq.add(2);
            let vmId = await self.provider.getVmID();
            let encodedData = new ArbValue.TupleValue([new ArbValue.TupleValue([]), new ArbValue.IntValue(0)]);
            if (transaction.data) {
                let test = await transaction.data;

                encodedData = ArbValue.hexToSizedByteRange(test);
            }
            let arbMsg = new ArbValue.TupleValue([
                encodedData,
                new ArbValue.IntValue(dest),
                new ArbValue.IntValue(self.seq),
            ]);
            if (!transaction.value) {
                transaction.value = ethers.utils.bigNumberify(0); // eslint-disable-line require-atomic-updates
            }
            let args = [vmId, arbMsg.hash(), transaction.value, ethers.utils.hexZeroPad('0x00', 21)];
            let messageHash = ethers.utils.solidityKeccak256(['bytes32', 'bytes32', 'uint256', 'bytes21'], args);
            let fromAddress = await self.getAddress();
            let tx = {
                hash: messageHash,
                from: fromAddress,
                gasPrice: 1,
                gasLimit: 1,
                to: dest,
                value: transaction.value,
                nonce: self.seq,
                data: transaction.data,
            };
            if (ethers.utils.bigNumberify(await transaction.value).eq(0)) {
                let messageHashBytes = ethers.utils.arrayify(messageHash);
                let sig = await self.signer.signMessage(messageHashBytes);
                if (!self.pubkey) {
                    self.pubkey = ethers.utils.recoverPublicKey(
                        ethers.utils.arrayify(ethers.utils.hashMessage(messageHashBytes)),
                        sig,
                    );
                }
                await self.client.sendMessage(arbMsg, sig, self.pubkey);
            } else {
                let tx = await self.vmTracker.sendEthMessage(vmId, ArbValue.marshal(arbMsg), {
                    value: transaction.value,
                });

                await tx.wait();
            }
            return self.provider._wrapTransaction(tx, messageHash);
        } else {
            return self.signer.sendTransaction(transaction);
        }
    }
}
