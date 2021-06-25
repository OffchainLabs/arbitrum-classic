import { setStandardGateWays } from './lib'
const tokens = []
if (tokens.length === 0) {
  throw new Error('Include some tokens to set')
}

setStandardGateWays(tokens).then(() => {
  console.log('done')
})
