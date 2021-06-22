import { checkRetryableStatus } from './lib'
const l1Hash = ''

if (!l1Hash) {
  throw new Error('Need to set an L1 hash')
}

checkRetryableStatus(l1Hash)
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
