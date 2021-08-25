import * as config from 'arb-bridge-eth/hardhat.config'
import { task } from 'hardhat/config'
import { initUpgrades } from 'arb-upgrades'
import 'hardhat-deploy-ethers'

task('deploy-logic-one', 'deploy one logic')
  .addParam('contract', 'contract to deploy')
  .setAction(async (args, hre) => {
    await hre.network.provider.send('hardhat_setBalance', [
      (await hre.ethers.getSigners())[0].address,
      '0x16189AD417E380000',
    ])

    const { contract } = args
    const { deployLogic } = initUpgrades(hre, __dirname)
    await deployLogic(contract)
  })

task('deploy-logic-all', 'deploy all logic contracts').setAction(
  async (_, hre) => {
    await hre.network.provider.send('hardhat_setBalance', [
      (await hre.ethers.getSigners())[0].address,
      '0x16189AD417E380000',
    ])

    const { deployLogicAll } = initUpgrades(hre, __dirname)
    await deployLogicAll()
  }
)

task('trigger-upgrades', 'triggers upgrade').setAction(async (_, hre) => {
  await hre.network.provider.send('hardhat_setBalance', [
    (await hre.ethers.getSigners())[0].address,
    '0x16189AD417E380000',
  ])

  const { updateImplementations } = initUpgrades(hre, __dirname)
  await updateImplementations()
})

task('verify-deployments', 'verifies implementations').setAction(
  async (_, hre) => {
    await hre.network.provider.send('hardhat_setBalance', [
      (await hre.ethers.getSigners())[0].address,
      '0x16189AD417E380000',
    ])

    const { verifyCurrentImplementations } = initUpgrades(hre, __dirname)
    await verifyCurrentImplementations()
  }
)

module.exports = config
