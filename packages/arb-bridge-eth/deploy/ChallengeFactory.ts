import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  const messagesChallenge = await deployments.get('MessagesChallenge')
  const inboxTopChallenge = await deployments.get('InboxTopChallenge')
  const executionChallenge = await deployments.get('ExecutionChallenge')

  let contract = await deployments.getOrNull('ChallengeFactory')
  if (!contract) {
    const deployResult = await deploy('ChallengeFactory', {
      from: deployer,
      args: [
        messagesChallenge.address,
        inboxTopChallenge.address,
        executionChallenge.address,
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
