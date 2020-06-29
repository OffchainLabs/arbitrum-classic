import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  let contract = await deployments.getOrNull('GlobalInbox')
  if (!contract) {
    const deployResult = await deploy('GlobalInbox', { from: deployer })
    contract = await deployments.get('GlobalInbox')
    if (deployResult.newlyDeployed) {
      log(
        `GlobalInbox deployed at ${contract.address} for ${deployResult.receipt.gasUsed}`
      )
    }
  }
}

module.exports = func
module.exports.tags = ['GlobalInbox']
