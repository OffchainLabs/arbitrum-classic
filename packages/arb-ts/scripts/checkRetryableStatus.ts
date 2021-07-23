import { checkRetryableStatus } from './lib'

import args from './getCLargs'

if (!args.txid) {
  throw new Error('Include txid (--txid 0xmytxid)')
}
const txId = args.txid as string

if (!txId) {
  throw new Error('Need to set an L1 hash')
}

checkRetryableStatus(txId)
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
