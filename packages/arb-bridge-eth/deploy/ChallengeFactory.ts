import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  const oneStepProof = await deployments.get('OneStepProof')
  const inboxTopChallenge = await deployments.get('InboxTopChallenge')
  const executionChallenge = await deployments.get('ExecutionChallenge')

  let contract = await deployments.getOrNull('ChallengeFactory')
  if (!contract) {
    const deployResult = await deploy('ChallengeFactory', {
      from: deployer,
      args: [
        inboxTopChallenge.address,
        executionChallenge.address,
        oneStepProof.address,
      ],
    })
    contract = await deployments.get('ChallengeFactory')
    if (deployResult.newlyDeployed && deployResult.receipt) {
      log(
        `ChallengeFactory deployed at ${contract.address} for ${deployResult.receipt.gasUsed}`
      )
    }
  }
}

module.exports = func
module.exports.tags = ['ChallengeFactory']
module.exports.dependencies = [
  'InboxTopChallenge',
  'ExecutionChallenge',
  'OneStepProof',
]
