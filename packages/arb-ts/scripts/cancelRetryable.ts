import { Bridge } from '../src/lib/bridge'
import { instantiateBridge } from './instantiate_bridge'
import { ContractReceipt } from 'ethers'

const l1Txn: string | ContractReceipt = ''

if (!l1Txn) {
  throw new Error('Need to set l1 txn hash')
}

;(async () => {
  const { bridge } = await instantiateBridge()
  const res = await bridge.cancelRetryableTicket(l1Txn)
  const rec = await res.wait()
  console.log('done:', rec)
  console.log(rec.status === 1 ? 'success!' : 'failed...')
})()
