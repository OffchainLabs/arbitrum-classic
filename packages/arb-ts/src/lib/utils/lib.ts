import { BigNumber, utils } from 'ethers'
import { ADDRESS_ALIAS_OFFSET } from '../dataEntities/constants'
import { ArbTsError } from '../dataEntities/errors'

export const wait = (ms: number): Promise<void> =>
  new Promise(res => setTimeout(res, ms))

export const throwIfNotAddress = (address: string) => {
  if (!utils.isAddress(address))
    throw new ArbTsError(`The supplied '${address}' is not a valid address`)
}

/**
 * Find the L2 alias of an L1 address
 * @param l1Address
 * @returns
 */
export const applyL1ToL2Alias = (l1Address: string): BigNumber => {
  throwIfNotAddress(l1Address)
  return BigNumber.from(l1Address).add(ADDRESS_ALIAS_OFFSET)
}

/**
 * Find the L1 alias of an L2 address
 * @param l2Address
 * @returns
 */
export const undoL1ToL2Alias = (l2Address: string): BigNumber => {
  throwIfNotAddress(l2Address)
  return BigNumber.from(l2Address).sub(ADDRESS_ALIAS_OFFSET)
}
