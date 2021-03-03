import { task } from 'hardhat/config'
import 'dotenv/config'
import * as fs from 'fs'

import 'hardhat-deploy'

import '@nomiclabs/hardhat-waffle'
import 'hardhat-typechain'
import 'solidity-coverage'
import 'hardhat-spdx-license-identifier'
import 'hardhat-gas-reporter'
import '@nomiclabs/hardhat-etherscan'

const verifyTask = require('./scripts/verifyTask') // eslint-disable-line @typescript-eslint/no-var-requires
const setupVerifyTask = verifyTask.default
setupVerifyTask()

task('accounts', 'Prints the list of accounts', async (taskArgs, bre) => {
  const accounts = await bre.ethers.getSigners()

  for (const account of accounts) {
    console.log(await account.getAddress())
  }
})

task('create-chain', 'Creates a rollup chain').setAction(
  async (taskArguments, hre) => {
    const machineHash = fs.readFileSync('../MACHINEHASH').toString()
    console.log(`Creating chain for machine with hash ${machineHash}`)
    const { deployments, ethers } = hre
    const [deployer] = await ethers.getSigners()
    const rollupCreatorDep = await deployments.get('RollupCreator')
    const RollupCreator = await ethers.getContractFactory('RollupCreator')
    const rollupCreator = RollupCreator.attach(
      rollupCreatorDep.address
    ).connect(deployer)
    const tx = await rollupCreator.createRollup(
      machineHash,
      900,
      0,
      2000000000,
      ethers.utils.parseEther('.1'),
      ethers.constants.AddressZero,
      await deployer.getAddress(),
      '0x'
    )
    const receipt = await tx.wait()
    console.log(receipt)
    console.log(receipt.logs)
    const ev = rollupCreator.interface.parseLog(
      receipt.logs[receipt.logs.length - 1]
    )
    console.log(ev)
  }
)

task('deposit', 'Deposit coins into ethbridge')
  .addPositionalParam('chain', "The rollup chain's address")
  .addPositionalParam('privkey', 'The private key of the depositer')
  .addPositionalParam('dest', "The destination account's address")
  .addPositionalParam('amount', 'The amount to deposit')
  .setAction(async ({ chain, privkey, dest, amount }, bre) => {
    const { deployments, ethers } = bre
    const inboxDep = await deployments.getOrNull('GlobalInbox')
    if (!inboxDep) {
      throw Error('GlobalInbox not deployed')
    }

    const wallet = new ethers.Wallet(privkey, ethers.provider)
    const GlobalInbox = await ethers.getContractFactory('GlobalInbox')
    const inbox = GlobalInbox.attach(inboxDep.address).connect(wallet)
    await inbox.depositEthMessage(chain, dest, { value: amount })
  })

const config = {
  defaultNetwork: 'hardhat',
  paths: {
    artifacts: 'build/contracts',
  },
  solc: {
    version: '0.5.17',
    optimizer: {
      enabled: true,
      runs: 200,
    },
  },
  typechain: {
    outDir: 'build/types',
    target: 'ethers-v5',
  },
  spdxLicenseIdentifier: {
    overwrite: false,
    runOnCompile: true,
  },
  gasReporter: {
    currency: 'USD',
    gasPrice: 20,
    enabled: process.env.REPORT_GAS ? true : false,
  },
  namedAccounts: {
    deployer: {
      default: 0,
    },
  },
  networks: {
    hardhat: {
      allowUnlimitedContractSize: true,
    },
    parity: {
      url: 'http://127.0.0.1:7545',
    },
    devnet: {
      url: 'https://devnet.arbitrum.io/rpc',
      accounts: process.env['DEVNET_PRIVKEY']
        ? [process.env['DEVNET_PRIVKEY']]
        : [],
    },
  },
  etherscan: {
    apiKey: process.env['ETHERSCAN_API_KEY'],
  },
  solidity: {
    version: '0.6.11',
    settings: {
      optimizer: {
        enabled: true,
        runs: 100,
      },
    },
  },
}

if (process.env['RINKEBY_URL'] && process.env['RINKEBY_MNEMONIC']) {
  ;(config.networks as any)['rinkeby'] = {
    url: process.env['RINKEBY_URL'] || '',
    accounts: [process.env['RINKEBY_MNEMONIC'] || ''],
    network_id: 4,
    confirmations: 1,
  }
}

if (process.env['ROPSTEN_URL'] && process.env['ROPSTEN_MNEMONIC']) {
  ;(config.networks as any)['ropsten'] = {
    url: process.env['ROPSTEN_URL'] || '',
    accounts: [process.env['ROPSTEN_MNEMONIC'] || ''],
    network_id: 3,
    confirmations: 1,
  }
}

if (process.env['KOVAN_URL'] && process.env['KOVAN_MNEMONIC']) {
  ;(config.networks as any)['ropsten'] = {
    url: process.env['KOVAN_URL'] || '',
    accounts: [process.env['KOVAN_MNEMONIC'] || ''],
    network_id: 42,
    confirmations: 4,
  }
}

module.exports = config
