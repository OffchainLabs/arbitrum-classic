import { ArbTsError } from '../errors'
import { JsonRpcProvider } from '@ethersproject/providers'

export const wait = (ms: number): Promise<void> =>
  new Promise(res => setTimeout(res, ms))

export interface ArbTransactionReceipt {
  to: string
  from: string
  contractAddress: string
  transactionIndex: string
  root?: string
  gasUsed: string
  logsBloom: string
  blockHash: string
  transactionHash: string
  logs: Array<any>
  blockNumber: string
  confirmations: string
  cumulativeGasUsed: string
  byzantium: boolean
  status?: string

  returnData: string
  returnCode: string
  // TODO expose this
  feeStats: {
    prices: {
      l1Transaction: string
      l1Calldata: string
      l2Storage: string
      l2Computation: string
    }
    unitsUsed: {
      l1Transaction: string
      l1Calldata: string
      l2Storage: string
      l2Computation: string
    }
    paid: {
      l1Transaction: string
      l1Calldata: string
      l2Storage: string
      l2Computation: string
    }
  }
}

export const getRawArbTransactionReceipt = async (
  l2Provider: JsonRpcProvider,
  txHash: string
): Promise<ArbTransactionReceipt> => {
  const rec = (await l2Provider.send('eth_getTransactionReceipt', [
    txHash,
  ])) as ArbTransactionReceipt
  if (!rec.returnCode)
    throw new ArbTsError(
      "Tx receipt doesn't have returnCode field. prob a l1 provider"
    )
  return rec
}
