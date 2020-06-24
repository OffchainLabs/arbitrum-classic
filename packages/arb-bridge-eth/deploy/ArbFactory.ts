import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  let arb_rollup = await deployments.get('ArbRollup')
  let global_inbox = await deployments.get('GlobalInbox')
  let challenge_factory = await deployments.get('ChallengeFactory')

  let contract = await deployments.getOrNull('ArbFactory')
  if (!contract) {
    const deployResult = await deploy('ArbFactory', {
      from: deployer,
      args: [
        arb_rollup.address,
        global_inbox.address,
        challenge_factory.address,
      ],
    })
    contract = await deployments.get('ArbFactory')
    if (deployResult.newlyDeployed) {
      log(
        `ArbFactory deployed at ${contract.address} for ${deployResult.receipt.gasUsed}`
      )
    }
  }
}

module.exports = func
module.exports.tags = ['ArbFactory']
module.exports.dependencies = ['ArbRollup', 'GlobalInbox', 'ChallengeFactory']
