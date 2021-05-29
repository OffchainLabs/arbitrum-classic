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

if (!process.env['DEVNET_PRIVKEY']) console.warn('No devnet privkey set')

task('accounts', 'Prints the list of accounts', async (taskArgs, bre) => {
  const accounts = await bre.ethers.getSigners()

  for (const account of accounts) {
    console.log(await account.getAddress())
  }
})

task('create-chain', 'Creates a rollup chain')
  .addParam('sequencer', "The sequencer's address")
  .setAction(async (taskArguments, hre) => {
    const machineHash = fs.readFileSync('../MACHINEHASH').toString()
    console.log(
      `Creating chain for machine with hash ${machineHash} for sequencer ${taskArguments.sequencer}`
    )
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
      taskArguments.sequencer,
      300,
      1500,
      '0x'
    )
    const receipt = await tx.wait()
    const ev = rollupCreator.interface.parseLog(
      receipt.logs[receipt.logs.length - 1]
    )
    console.log(ev)

    // const path = `rollup-${hre.network.name}.json`
    const path = `rollup-${hre.network.name}.json`
    const output = JSON.stringify({
      rollupAddress: ev.args[0],
      inboxAddress: ev.args[1],
    })

    fs.writeFileSync(path, output)
    console.log(
      'New rollup chain created and output written to:',
      `${process.cwd()}:${path}`
    )
  })

task('deposit', 'Deposit coins into ethbridge')
  .addPositionalParam('inboxAddress', "The rollup chain's address")
  .addPositionalParam('privkey', 'The private key of the depositer')
  .addPositionalParam('dest', "The destination account's address")
  .addPositionalParam('amount', 'The amount to deposit')
  .setAction(async ({ inboxAddress, privkey, dest, amount }, bre) => {
    const { ethers } = bre
    const wallet = new ethers.Wallet(privkey, ethers.provider)
    const GlobalInbox = await ethers.getContractFactory('Inbox')
    const inbox = GlobalInbox.attach(inboxAddress).connect(wallet)
    await inbox.depositEth(dest, { value: amount })
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
      chainId: 1337,
      allowUnlimitedContractSize: true,
      accounts: {
        accountsBalance: '10000000000000000000000000',
      },
      blockGasLimit: 20000000,
      // mining: {
      //   auto: false,
      //   interval: 1000,
      // },
    },
    local_development: {
      url: 'http://127.0.0.1:7545',
    },
    kovan: {
      url: 'https://kovan.infura.io/v3/' + process.env['INFURA_KEY'],
      accounts: process.env['DEVNET_PRIVKEY']
        ? [process.env['DEVNET_PRIVKEY']]
        : [],
    },
    // mainnet: {
    //   url: process.env['MAINNET_URL'],
    //   accounts: process.env['MAINNET_PRIVKEY']
    //     ? [process.env['MAINNET_PRIVKEY']]
    //     : [],
    // },
    arbkovan4: {
      gasPrice: 0,
      url: 'https://kovan4.arbitrum.io/rpc',
      accounts: process.env['DEVNET_PRIVKEY']
        ? [process.env['DEVNET_PRIVKEY']]
        : [],
    },
    kovan5: {
      gasPrice: 0,
      url: 'https://kovan5.arbitrum.io/rpc',
      accounts: process.env['DEVNET_PRIVKEY']
        ? [process.env['DEVNET_PRIVKEY']]
        : [],
    },
    devnet: {
      url: 'https://devnet.arbitrum.io/rpc',
      accounts: process.env['DEVNET_PRIVKEY']
        ? [process.env['DEVNET_PRIVKEY']]
        : [],
    },
    devnetL2: {
      url: 'https://devnet-l2.arbitrum.io/rpc',
      accounts: process.env['DEVNET_PRIVKEY']
        ? [process.env['DEVNET_PRIVKEY']]
        : [],
    },
    arbitrum: {
      url: 'http://127.0.0.1:8547',
      // url: 'https://kovan3.arbitrum.io/rpc',
      gas: 999999999999999,
      accounts: {
        mnemonic:
          'jar deny prosper gasp flush glass core corn alarm treat leg smart',
        path: "m/44'/60'/0'/0",
        initialIndex: 0,
        count: 10,
      },
      timeout: 100000,
    },
  },
  mocha: {
    timeout: 0,
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
  ;(config.networks as any)['kovan'] = {
    url: process.env['KOVAN_URL'] || '',
    accounts: [process.env['KOVAN_MNEMONIC'] || ''],
    network_id: 42,
    confirmations: 4,
  }
}

module.exports = config
