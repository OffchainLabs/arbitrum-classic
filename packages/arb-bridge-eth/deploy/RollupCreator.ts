import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre
  const { deploy } = deployments
  const { deployer } = await getNamedAccounts()

  const challengeFactory = await deployments.get('ChallengeFactory')
  await deploy('RollupCreator', {
    from: deployer,
    args: [challengeFactory.address],
  })
}

module.exports = func
module.exports.tags = ['RollupCreator']
module.exports.dependencies = ['ChallengeFactory']
