import * as ethers from 'ethers'
import { abi, Program, ArbConversion } from 'arb-provider-ethers'
import * as yargs from 'yargs'
import * as fs from 'fs-extra'

import * as addresses from '../../arb-bridge-eth/bridge_eth_addresses.json'

const argv = yargs.options({
  force: {
    description: 'clear any existing state',
  },
  validatorcount: {
    description: 'number of validators to deploy',
    default: 1,
  },
  blocktime: {
    description: 'expected length of time between blocks',
    default: 2,
  },
}).argv

const arbConversion = new ArbConversion()

interface RollupCreatedParams {
  vmAddress: string
}

const provider = new ethers.providers.JsonRpcProvider('http://localhost:7545')

const wallet = provider.getSigner(0)
const validatorsPath = '../../validator-states'

const optionDefinitions = [{ name: 'force', alias: 'f', type: Boolean }]

async function setupRollup(arbOSData: string) {
  const arbOSHash = Program.programMachineHash(arbOSData)

  const factoryAddress = addresses['ArbFactory']

  const factory = abi.ArbFactoryFactory.connect(factoryAddress, wallet)

  console.log(`Initializing rollup chain for machine with hash ${arbOSHash}`)

  const tx = await factory.createRollup(
    arbOSHash,
    arbConversion.blocksToTicks(30),
    80000000,
    10000000000,
    ethers.utils.parseEther('.01'),
    ethers.utils.hexZeroPad('0x', 20)
  )
  const result = await tx.wait()

  const e = result.events?.find((e: ethers.Event) =>
    e.topics.includes(
      (factory.interface.events.RollupCreated as ethers.utils.EventDescription)
        .topic
    )
  )

  const {
    vmAddress,
  }: RollupCreatedParams = (e?.args as any) as RollupCreatedParams

  return vmAddress
}

async function initializeWallets(count: number): Promise<ethers.Wallet[]> {
  const wallets: ethers.Wallet[] = []
  const waits = []
  for (let i = 0; i < count; i++) {
    const newWallet = ethers.Wallet.createRandom()
    const tx = {
      to: newWallet.address,
      value: ethers.utils.parseEther('5.0'),
    }
    const send = await wallet.sendTransaction(tx)
    wallets.push(newWallet)
    waits.push(send.wait())
  }
  await Promise.all(waits)
  return wallets
}

async function setupValidatorStates(count: number, blocktime: number) {
  if (count < 1) {
    throw Error('must create at least 1 validator')
  }
  if (fs.existsSync(validatorsPath)) {
    if (argv.force) {
      fs.removeSync(validatorsPath)
    } else {
      throw Error(
        'validator-states folder already exists. First manually delete it or run with --force'
      )
    }
  }
  const wallets = await initializeWallets(count)

  const arbOSData = fs.readFileSync('../../arbos.mexe', 'utf8')
  const rollup = await setupRollup(arbOSData)
  console.log('Created rollup', rollup)

  const config = {
    rollup_address: rollup,
    eth_url: 'http://localhost:7545',
    password: 'pass',
    blocktime: blocktime,
  }

  fs.mkdirSync(validatorsPath)
  let i = 0
  for (const wallet of wallets) {
    const valPath = validatorsPath + '/validator' + i
    const walletPath = valPath + '/wallets'
    fs.mkdirSync(valPath)
    fs.mkdirSync(walletPath)
    const encryptedWallet = await wallet.encrypt('pass')
    fs.writeFileSync(walletPath + '/' + wallet.address, encryptedWallet)
    fs.writeFileSync(valPath + '/config.json', JSON.stringify(config))
    fs.writeFileSync(valPath + '/contract.mexe', arbOSData)
    i++
  }
}

if (require.main === module) {
  setupValidatorStates(argv.validatorcount, argv.blocktime)
}
