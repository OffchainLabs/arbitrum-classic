import { Provider } from '@ethersproject/abstract-provider'
import { SignerOrProvider } from '../utils/signerOrProvider'
import { Signer } from '@ethersproject/abstract-signer'

import { ArbAddressTable__factory } from '../abi/factories/ArbAddressTable__factory'
import { ARB_ADDRESS_TABLE_ADDRESS } from '../constants'
import { BigNumber, BigNumberish } from 'ethers'

export class ArbAddressTable {
  static getAddressTableContract(l2SignerOrProvider: SignerOrProvider) {
    return ArbAddressTable__factory.connect(
      ARB_ADDRESS_TABLE_ADDRESS,
      l2SignerOrProvider
    )
  }

  static async addressIsRegistered(
    addr: string,
    l2SignerOrProvider: SignerOrProvider
  ): Promise<boolean> {
    const addrTable = await ArbAddressTable.getAddressTableContract(
      l2SignerOrProvider
    )
    return addrTable.addressExists(addr)
  }

  static async lookupIndexOfAddress(
    addr: string,
    l2SignerOrProvider: SignerOrProvider
  ): Promise<BigNumber> {
    const addrTable = await ArbAddressTable.getAddressTableContract(
      l2SignerOrProvider
    )
    return addrTable.lookup(addr)
  }

  static async lookupAddressOfIndex(
    index: BigNumberish,
    l2SignerOrProvider: SignerOrProvider
  ): Promise<string> {
    const addrTable = await ArbAddressTable.getAddressTableContract(
      l2SignerOrProvider
    )
    return addrTable.lookupIndex(index)
  }

  static async registerAddress(addr: string, l2Signer: Signer) {
    const addrTable = await ArbAddressTable.getAddressTableContract(l2Signer)
    return addrTable.register(addr)
  }
}
