import { TransactionReceipt, Provider } from '@ethersproject/providers'
import { BigNumber } from '@ethersproject/bignumber'

export const getTxnReceipt = async (
  txn: string | TransactionReceipt,
  provider?: Provider
): Promise<TransactionReceipt> => {
  if (typeof txn === 'string') {
    if (!provider) throw new Error('Must include provider')
    const txnReceipt = await provider.getTransactionReceipt(txn)
    if (!txnReceipt) throw new Error('Txn receipt not found')
    return txnReceipt
  } else {
    return txn
  }
}

export const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

export const percentIncrease = (
  num: BigNumber,
  increase: BigNumber
): BigNumber => {
  return num.add(num.mul(increase).div(100))
}

export function isError(error: Error): error is NodeJS.ErrnoException {
  return error instanceof Error
}
