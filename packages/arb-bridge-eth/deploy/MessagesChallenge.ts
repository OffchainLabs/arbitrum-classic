import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  let contract = await deployments.getOrNull('MessagesChallenge')
  if (!contract) {
    const deployResult = await deploy('MessagesChallenge', { from: deployer })
    contract = await deployments.get('MessagesChallenge')
    if (deployResult.newlyDeployed) {
      log(
        `MessagesChallenge deployed at ${contract.address} for ${deployResult.receipt.gasUsed}`
      )
    }
  }
}

module.exports = func
module.exports.tags = ['MessagesChallenge']
