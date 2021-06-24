import { setArbCustomGateways } from './lib'
const tokens = []
if (tokens.length === 0) {
  throw new Error('Include some tokens to set')
}

setArbCustomGateways(tokens).then(() => {
  console.log('done')
})
