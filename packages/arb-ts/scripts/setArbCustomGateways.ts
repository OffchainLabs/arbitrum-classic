import { setArbCustomGateways } from './lib'
import args from './getCLargs'

if (!args.address) {
  throw new Error('Include token address (--address 0xmyaddress)')
}

const tokens = [args.address as string]
if (tokens.length === 0) {
  throw new Error('Include some tokens to set')
}

setArbCustomGateways(tokens).then(() => {
  console.log('done')
})
