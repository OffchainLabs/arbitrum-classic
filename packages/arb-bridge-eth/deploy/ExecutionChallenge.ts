import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  let contract = await deployments.getOrNull('ExecutionChallenge')
  if (!contract) {
    const deployResult = await deploy('ExecutionChallenge', { from: deployer })
    contract = await deployments.get('ExecutionChallenge')
    if (deployResult.newlyDeployed && deployResult.receipt) {
      log(
        `ExecutionChallenge deployed at ${contract.address} for ${deployResult.receipt.gasUsed}`
      )
    }
  }
}

module.exports = func
module.exports.tags = ['ExecutionChallenge']
