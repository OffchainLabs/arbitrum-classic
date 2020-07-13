import * as yargs from 'yargs'
import * as fs from 'fs-extra'

const root = '../../'
const rollupsPath = root + 'rollups/'

export interface Config {
  rollup_address: string
  eth_url: string
  password?: string
  blocktime: number
}

export async function setupValidatorStates(
  count: number,
  folder: string,
  config: Config
): Promise<void> {
  if (count < 1) {
    throw Error('must create at least 1 validator')
  }
  if (!fs.existsSync(rollupsPath)) {
    fs.mkdirSync(rollupsPath)
  }

  const arbOSData = fs.readFileSync('../../arbos.mexe', 'utf8')

  const rollupPath = rollupsPath + folder + '/'
  if (fs.existsSync(rollupPath)) {
    throw Error(`${rollupPath} folder already exists`)
  }

  fs.mkdirSync(rollupPath)
  for (let i = 0; i < count; i++) {
    const valPath = rollupPath + 'validator' + i + '/'
    fs.mkdirSync(valPath)
    fs.writeFileSync(valPath + 'config.json', JSON.stringify(config))
    fs.writeFileSync(valPath + 'contract.mexe', arbOSData)
  }
}

if (require.main === module) {
  const argv = yargs.command(
    'init [rollup] [ethurl]',
    'initialize validators for the given rollup chain',
    yargsBuilder =>
      yargsBuilder
        .positional('rollup', {
          describe: 'address of the rollup chain',
          type: 'string',
          demandOption: true,
        })
        .positional('ethurl', {
          describe: 'url for ethereum node',
          type: 'string',
          demandOption: true,
        })
        .options({
          validatorcount: {
            description: 'number of validators to deploy',
            default: 1,
          },
          blocktime: {
            description: 'expected length of time between blocks',
            default: 2,
          },
        }),
    args => {
      if (!args.rollup || !args.ethurl) {
        console.error('Must supply rollup address and eth url')
        return
      }
      const config: Config = {
        rollup_address: args.rollup,
        eth_url: args.ethurl,
        blocktime: args.blocktime,
      }

      setupValidatorStates(args.validatorcount, config.rollup_address, config)
    }
  ).argv
}
