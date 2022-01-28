import { Provider } from '@ethersproject/abstract-provider'
import { SignerOrProvider } from '../utils/signerOrProvider'
import { Signer } from '@ethersproject/abstract-signer'

import { ArbAddressTable__factory } from '../abi/factories/ArbAddressTable__factory'
import { ARB_ADDRESS_TABLE_ADDRESS } from '../constants'
import { BigNumber, BigNumberish } from 'ethers'

export class ArbAddressTable {
  getAddressTableContract(l2SignerOrProvider: SignerOrProvider) {
    return ArbAddressTable__factory.connect(
      ARB_ADDRESS_TABLE_ADDRESS,
      l2SignerOrProvider
    )
  }

  async addressIsRegistered(
    addr: string,
    l2SignerOrProvider: SignerOrProvider
  ): Promise<boolean> {
    const addrTable = await this.getAddressTableContract(l2SignerOrProvider)
    return addrTable.addressExists(addr)
  }

  async lookupIndexOfAddress(
    addr: string,
    l2SignerOrProvider: SignerOrProvider
  ): Promise<BigNumber> {
    const addrTable = await this.getAddressTableContract(l2SignerOrProvider)
    return addrTable.lookup(addr)
  }

  async lookupAddressOfIndex(
    index: BigNumberish,
    l2SignerOrProvider: SignerOrProvider
  ): Promise<string> {
    const addrTable = await this.getAddressTableContract(l2SignerOrProvider)
    return addrTable.lookupIndex(index)
  }

  async registerAddress(addr: string, l2Signer: Signer) {
    const addrTable = await this.getAddressTableContract(l2Signer)
    return addrTable.register(addr)
  }
}
