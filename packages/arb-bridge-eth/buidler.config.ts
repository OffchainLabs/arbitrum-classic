import { task, usePlugin } from '@nomiclabs/buidler/config'
import 'dotenv/config'

usePlugin('buidler-deploy')
if (!process.env.DOCKER) {
  usePlugin('@nomiclabs/buidler-waffle')
  usePlugin('buidler-typechain')
  // usePlugin('solidity-coverage')
  // usePlugin('buidler-spdx-license-identifier')
  // usePlugin('buidler-gas-reporter')
  // usePlugin('@nomiclabs/buidler-etherscan')

  // const verifyTask = require('./scripts/verifyTask') // eslint-disable-line @typescript-eslint/no-var-requires
  // const setupVerifyTask = verifyTask.default
  // setupVerifyTask()
}

task('accounts', 'Prints the list of accounts', async (taskArgs, bre) => {
  const accounts = await bre.ethers.getSigners()

  for (const account of accounts) {
    console.log(await account.getAddress())
  }
})

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

module.exports = {
  defaultNetwork: 'buidlerevm',
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
    target: 'ethers-v4',
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
    buidlerevm: {},
    parity: {
      url: 'http://127.0.0.1:7545',
    },
    rinkeby: {
      url: process.env['RINKEBY_URL'] || '',
      accounts: [process.env['RINKEBY_MNEMONIC'] || ''],
      network_id: 4,
      confirmations: 1,
    },
    ropsten: {
      url: process.env['ROPSTEN_URL'] || '',
      accounts: [process.env['ROPSTEN_MNEMONIC'] || ''],
      network_id: 3,
      confirmations: 1,
    },
    kovan: {
      url: process.env['KOVAN_URL'] || '',
      accounts: [process.env['KOVAN_MNEMONIC'] || ''],
      network_id: 42,
      confirmations: 4,
    },
  },
  etherscan: {
    apiKey: process.env['ETHERSCAN_API_KEY'],
  },
}
