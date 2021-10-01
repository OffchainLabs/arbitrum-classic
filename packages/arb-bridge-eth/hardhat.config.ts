import { task } from 'hardhat/config'
import { HardhatRuntimeEnvironment } from 'hardhat/types/runtime'
import 'dotenv/config'
import * as fs from 'fs'

import 'hardhat-deploy'

import '@nomiclabs/hardhat-waffle'
import '@typechain/hardhat'
import 'solidity-coverage'
import 'hardhat-spdx-license-identifier'
import 'hardhat-gas-reporter'
import '@nomiclabs/hardhat-etherscan'
import { initUpgrades } from 'arb-upgrades'

const verifyTask = require('./scripts/verifyTask') // eslint-disable-line @typescript-eslint/no-var-requires
const setupVerifyTask = verifyTask.default
setupVerifyTask()

const handleFork = async (hre: HardhatRuntimeEnvironment) => {
  const network = hre.network
  if (network.name === 'fork') {
    await hre.network.provider.send('hardhat_setBalance', [
      (await hre.ethers.getSigners())[0].address,
      '0x16189AD417E380000',
    ])
  }
  return true
}

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

task('core-deploy-logic-one', 'deploy one logic')
  .addParam('contract', 'contract to deploy')
  .setAction(async (args, hre) => {
    await handleFork(hre)
    const { contract } = args
    const { deployLogic } = initUpgrades(hre, __dirname)
    await deployLogic(contract)
  })

task('core-deploy-logic-all', 'deploy all logic contracts').setAction(
  async (_, hre) => {
    await handleFork(hre)
    const { deployLogicAll } = initUpgrades(hre, __dirname)
    await deployLogicAll()
  }
)

task('core-trigger-upgrades', 'triggers upgrade').setAction(async (_, hre) => {
  await handleFork(hre)
  const { updateImplementations } = initUpgrades(hre, __dirname)
  await updateImplementations()
})

task('core-verify-deployments', 'verifies implementations').setAction(
  async (_, hre) => {
    await handleFork(hre)
    const { verifyCurrentImplementations } = initUpgrades(hre, __dirname)
    await verifyCurrentImplementations()
  }
)

task('core-transfer-beacon-owner', 'transfers beacon owner')
  .addParam('address', 'beacon contract')
  .addParam('newowner', 'beacon contract')

  .setAction(async (args, hre) => {
    await handleFork(hre)
    const { transferBeaconOwner } = initUpgrades(hre, __dirname)
    await transferBeaconOwner(args.address, args.newowner)
  })

task('core-transfer-admin', 'transfer proxy admin')
  .addParam('proxyaddress', 'proxy address')
  .addParam('newadmin', 'address of new admin')
  .setAction(async (args, hre) => {
    const { transferAdmin } = initUpgrades(hre, __dirname)
    await transferAdmin(args.proxyaddress, args.newadmin)
  })

task('etherscan-verify', 'verify current deployments in etherscan').setAction(
  async (_, hre) => {
    const { verifyDeployments } = await initUpgrades(hre, __dirname)
    await verifyDeployments()
  }
)

task(
  'remove-build-info',
  'remove giant build info string from current_deployments json'
).setAction(async (_, hre) => {
  const { removeBuildInfoFiles } = initUpgrades(hre, __dirname)
  await removeBuildInfoFiles()
})

task('deploy-outbox-logic', 'deploy and set a new outbox').setAction(
  async (_, hre) => {
    const OutboxFactory = await hre.ethers.getContractFactory('Outbox')
    console.log('Deploying outbox logic')
    const OutboxLogic = await OutboxFactory.deploy()
    await OutboxLogic.deployed()
    console.log('Outbox logic deployed at:', OutboxLogic.address)
  }
)

task('deploy-outbox-proxy', 'deploy outbox proxy')
  .addParam('outboxlogic', 'outbox logic')
  .setAction(async (args, hre) => {
    const { getDeployments } = initUpgrades(hre, __dirname)
    const { data } = await getDeployments()
    const proxyAdminAddress = data.proxyAdminAddress

    console.log('Deploying Outbox TransparentUpgradeableProxy')
    const TransparentUpgradeableProxyFactory =
      await hre.ethers.getContractFactory('TransparentUpgradeableProxy')
    const OutboxProxyDeployed = await TransparentUpgradeableProxyFactory.deploy(
      args.outboxlogic,
      proxyAdminAddress,
      '0x'
    )
    await OutboxProxyDeployed.deployed()
    console.log('Outbox proxy deployed at', OutboxProxyDeployed.address)
  })

task('init-outbox', 'deploy and set a new outbox')
  .addParam('outboxproxy', '')

  .setAction(async (args, hre) => {
    const { getDeployments } = initUpgrades(hre, __dirname)
    const { data } = await getDeployments()
    const rollupAddress = data.contracts.Rollup.proxyAddress
    const bridgeAddress = data.contracts.Bridge.proxyAddress

    const Outbox = (await hre.ethers.getContractFactory('Outbox')).attach(
      args.outboxproxy
    )
    const initializeRes = await Outbox.initialize(rollupAddress, bridgeAddress)
    const initializeRec = await initializeRes.wait()
    console.log('Outbox initialized', initializeRec)
  })

task('set-outbox', 'deploy and set a new outbox')
  .addParam('outboxproxy', '')

  .setAction(async (args, hre) => {
    const { getDeployments } = initUpgrades(hre, __dirname)
    const { data } = await getDeployments()
    const rollupAddress = data.contracts.Rollup.proxyAddress
    console.log('Sanity checking ')
    const Rollup = (await hre.ethers.getContractFactory('Rollup')).attach(
      rollupAddress
    )
    await Rollup.getUserFacet()
    console.log('Rollup sanity checked ')
    const RollupAdmin = (
      await hre.ethers.getContractFactory('RollupAdminFacet')
    ).attach(rollupAddress)
    const setRollupRes = await RollupAdmin.setOutbox(args.outboxproxy)
    const setRollupRec = await setRollupRes.wait()
    console.log('Outbox set', setRollupRec)
    console.log('all set 👍')
  })

const config = {
  defaultNetwork: 'hardhat',
  paths: {
    artifacts: 'build/contracts',
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
      throwOnTransactionFailures: true,
      allowUnlimitedContractSize: true,
      accounts: {
        accountsBalance: '1000000000000000000000000000',
      },
      blockGasLimit: 200000000,
      // mining: {
      //   auto: false,
      //   interval: 1000,
      // },
      forking: {
        url: 'https://mainnet.infura.io/v3/' + process.env['INFURA_KEY'],
        enabled: process.env['SHOULD_FORK'] === '1',
      },
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
    mainnet: {
      url: 'https://mainnet.infura.io/v3/' + process.env['INFURA_KEY'],
      accounts: process.env['MAINNET_PRIVKEY']
        ? [process.env['MAINNET_PRIVKEY']]
        : [],
    },
    fork: {
      url: 'http://127.0.0.1:8545/',
    },
    arbitrum1: {
      url: 'https://arb1.arbitrum.io/rpc',
      accounts: process.env['MAINNET_PRIVKEY']
        ? [process.env['MAINNET_PRIVKEY']]
        : [],
    },
    rinkeby: {
      url: 'https://rinkeby.infura.io/v3/' + process.env['INFURA_KEY'],
      accounts: process.env['DEVNET_PRIVKEY']
        ? [process.env['DEVNET_PRIVKEY']]
        : [],
    },
    arbRinkeby: {
      gasPrice: 0,
      url: 'https://rinkeby.arbitrum.io/rpc',
      accounts: process.env['DEVNET_PRIVKEY']
        ? [process.env['DEVNET_PRIVKEY']]
        : [],
    },
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
    bail: true,
  },
  etherscan: {
    apiKey: process.env['ETHERSCAN_API_KEY'],
  },
  solidity: {
    compilers: [
      {
        version: '0.6.11',
        settings: {
          optimizer: {
            enabled: true,
            runs: 100,
          },
        },
      },
      {
        version: '0.8.7',
        settings: {
          optimizer: {
            enabled: true,
            runs: 100,
          },
        },
      },
    ],
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
