import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  let messages_challenge = await deployments.get('MessagesChallenge')
  let inbox_top_challenge = await deployments.get('InboxTopChallenge')
  let execution_challenge = await deployments.get('ExecutionChallenge')

  let contract = await deployments.getOrNull('ChallengeFactory')
  if (!contract) {
    const deployResult = await deploy('ChallengeFactory', {
      from: deployer,
      args: [
        messages_challenge.address,
        inbox_top_challenge.address,
        execution_challenge.address,
      ],
    })
    contract = await deployments.get('ChallengeFactory')
    if (deployResult.newlyDeployed) {
      log(
        `ChallengeFactory deployed at ${contract.address} for ${deployResult.receipt.gasUsed}`
      )
    }
  }
}

module.exports = func
module.exports.tags = ['ChallengeFactory']
module.exports.dependencies = [
  'MessagesChallenge',
  'InboxTopChallenge',
  'ExecutionChallenge',
]
