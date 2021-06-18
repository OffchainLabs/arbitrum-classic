import { checkRetryableStatus } from './lib'
const l1Hash =
  '0x7703437080b3f1b85b09ca00c5a64e51e3b715b525b20835a93fd050d73f27f4'

if (!l1Hash) {
  throw new Error('Need to set an L1 hash')
}

checkRetryableStatus(l1Hash).then(() => {
  console.log('done')
})
