import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  let contract = await deployments.getOrNull('ArbRollup')
  if (!contract) {
    const deployResult = await deploy('ArbRollup', { from: deployer })
    contract = await deployments.get('ArbRollup')
    if (deployResult.newlyDeployed && deployResult.receipt) {
      log(
        `ArbRollup deployed at ${contract.address} for ${deployResult.receipt.gasUsed}`
      )
    }
  }
}

module.exports = func
module.exports.tags = ['ArbRollup']
