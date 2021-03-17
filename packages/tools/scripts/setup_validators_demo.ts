import * as ethers from 'ethers'
import { EventFragment } from '@ethersproject/abi'
import { ArbConversion, L1Bridge } from 'arb-ts'
import { RollupCreator__factory } from 'arb-ts/dist/lib/abi/factories/RollupCreator__factory'
import { Inbox__factory } from 'arb-ts/dist/lib/abi/factories/Inbox__factory'
import * as yargs from 'yargs'
import * as fs from 'fs-extra'
import { setupValidatorStates } from './setup_validators'

import * as addresses from '../../arb-bridge-eth/bridge_eth_addresses.json'

const arbConversion = new ArbConversion()

interface RollupCreatedParams {
  rollupAddress: string
}

const provider = new ethers.providers.JsonRpcProvider('http://localhost:7545')

const wallet = provider.getSigner(0)
const root = '../../'
const rollupsPath = root + 'rollups/'

export interface RollupCreatedEvent {
  rollupAddress: string
  inboxAddress: string
}

async function setupRollup(): Promise<RollupCreatedEvent> {
  const machineHash = fs.readFileSync('../MACHINEHASH').toString()
  console.log(`Creating chain for machine with hash ${machineHash}`)

  const factoryAddress = addresses['contracts']['RollupCreator'].address
  const rollupCreator = RollupCreator__factory.connect(factoryAddress, wallet)

  const tx = await rollupCreator.createRollup(
    machineHash,
    900,
    0,
    2000000000,
    ethers.utils.parseEther('.1'),
    ethers.constants.AddressZero,
    await wallet.getAddress(),
    '0x'
  )
  const receipt = await tx.wait()
  const ev = rollupCreator.interface.parseLog(
    receipt.logs[receipt.logs.length - 1]
  )

  if (ev.name != 'RollupCreated') {
    throw 'expected RollupCreated event'
  }

  const parsedEv = (ev as any) as {
    args: RollupCreatedEvent
  }
  return parsedEv.args
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

async function initializeClientWallets(inboxAddress: string): Promise<void> {
  const addresses = [
    '0xc7711f36b2C13E00821fFD9EC54B04A60AEfbd1b',
    '0x38299D74a169e68df4Da85Fb12c6Fd22246aDD9F',
    '0xAf40F7D235A9786a420bb89B188910958fD7EF93',
    '0xFcC598b3E3575CA937AF7F0E804a8BAb5E92a3f6',
    '0x755449b9901f91deC52DB39AF8c655206C63eD8e',
  ]

  const inbox = Inbox__factory.connect(inboxAddress, wallet)
  const amount = ethers.utils.parseEther('100')

  for (const address of addresses) {
    await inbox.depositEth(address, { value: amount })
  }
}

async function setupValidators(
  count: number,
  blocktime: number,
  force: boolean
): Promise<void> {
  const { rollupAddress, inboxAddress } = await setupRollup()
  console.log('Created rollup', rollupAddress)

  const validatorsPath = rollupsPath + 'local/'

  if (count < 2) {
    throw Error('must create at least 1 validator')
  }

  if (!fs.existsSync(rollupsPath)) {
    fs.mkdirSync(rollupsPath)
  }

  if (fs.existsSync(validatorsPath)) {
    if (force) {
      fs.removeSync(validatorsPath)
    } else {
      throw Error(
        `${validatorsPath} already exists. First manually delete it or run with --force`
      )
    }
  }

  const config = {
    rollup_address: rollupAddress,
    inbox_address: inboxAddress,
    validator_utils_address: addresses['contracts']['ValidatorUtils'].address,
    eth_url: 'http://localhost:7545',
    password: 'pass',
    blocktime: blocktime,
  }

  await setupValidatorStates(count, 'local', config)

  const wallets = await initializeWallets(count)
  let i = 0
  for (const wallet of wallets) {
    const valPath = validatorsPath + 'validator' + i + '/'
    const walletPath = valPath + 'wallets/'
    fs.mkdirSync(walletPath)
    const encryptedWallet = await wallet.encrypt('pass')
    fs.writeFileSync(walletPath + wallet.address, encryptedWallet)
    i++
  }

  await initializeClientWallets(inboxAddress)
}

if (require.main === module) {
  yargs.command(
    'init [rollup] [ethurl]',
    'initialize validators for the given rollup chain',
    yargsBuilder =>
      yargsBuilder.options({
        force: {
          description: 'clear any existing state',
          type: 'boolean',
          default: false,
        },
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
      setupValidators(args.validatorcount + 1, args.blocktime, args.force)
    }
  ).argv
}
