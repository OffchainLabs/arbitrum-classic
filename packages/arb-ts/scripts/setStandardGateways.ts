import { setStandardGateWays } from './lib'
import yargs from 'yargs/yargs'
import args from './getCLargs'

if (!args.address) {
  throw new Error('Include token address (--address 0xmyaddress)')
}

const tokens: string[] = [args.address as string]
if (tokens.length === 0) {
  throw new Error('Include some tokens to set')
}

setStandardGateWays(tokens).then(() => {
  console.log('done')
})
