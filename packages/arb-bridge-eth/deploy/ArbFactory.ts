import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  const arbRollup = await deployments.get('ArbRollup')
  const globalInbox = await deployments.get('GlobalInbox')
  const challengeFactory = await deployments.get('ChallengeFactory')

  let contract = await deployments.getOrNull('ArbFactory')
  if (!contract) {
    const deployResult = await deploy('ArbFactory', {
      from: deployer,
      args: [arbRollup.address, globalInbox.address, challengeFactory.address],
    })
    contract = await deployments.get('ArbFactory')
    if (deployResult.newlyDeployed && deployResult.receipt) {
      log(
        `ArbFactory deployed at ${contract.address} for ${deployResult.receipt.gasUsed}`
      )
    }
  }
}

module.exports = func
module.exports.tags = ['ArbFactory']
module.exports.dependencies = ['ArbRollup', 'GlobalInbox', 'ChallengeFactory']
