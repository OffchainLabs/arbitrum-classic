import { BigNumberish } from 'ethers'

export const addressToSymbol = (erc20L1Address: string) => {
  return erc20L1Address.substr(erc20L1Address.length - 3).toUpperCase() + '?'
}

export class TransactionOverrides {
  nonce?: BigNumberish | Promise<BigNumberish>
  gasLimit?: BigNumberish | Promise<BigNumberish>
  gasPrice?: BigNumberish | Promise<BigNumberish>
  value?: BigNumberish | Promise<BigNumberish>
  chainId?: number | Promise<number>
}
