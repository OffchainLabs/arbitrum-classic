import * as yargs from 'yargs'
import * as fs from 'fs-extra'

const root = '../../'
const rollupsPath = root + 'rollups/'

export interface Config {
  rollup_address: string
  inbox_address: string
  validator_utils_address: string
  validator_wallet_factory_address: string
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

  const arbOSData = fs.readFileSync('../arb-os/arb_os/arbos.mexe', 'utf8')

  const rollupPath = rollupsPath + folder + '/'
  if (fs.existsSync(rollupPath)) {
    throw Error(`${rollupPath} folder already exists`)
  }

  fs.mkdirSync(rollupPath)
  for (let i = 0; i < count; i++) {
    const valPath = rollupPath + 'validator' + i + '/'
    fs.mkdirSync(valPath)
    fs.writeFileSync(valPath + 'config.json', JSON.stringify(config))
    fs.writeFileSync(valPath + 'arbos.mexe', arbOSData)
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
        .positional('inbox', {
          describe: 'address of the rollup chain inbox',
          type: 'string',
          demandOption: true,
        })
        .positional('validatorutils', {
          describe: 'address of the validator utils contract',
          type: 'string',
          demandOption: true,
        })
        .positional('validatorwallet', {
          describe: 'address of the validator wallet creator contract',
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
        inbox_address: args.inbox,
        validator_utils_address: args.validatorutils,
        validator_wallet_factory_address: args.validatorwallet,
        eth_url: args.ethurl,
        blocktime: args.blocktime,
      }

      setupValidatorStates(
        args.validatorcount + 1,
        config.rollup_address,
        config
      )
    }
  ).argv
}
