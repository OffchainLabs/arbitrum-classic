import { TransactionReceipt, Provider } from '@ethersproject/providers'

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
