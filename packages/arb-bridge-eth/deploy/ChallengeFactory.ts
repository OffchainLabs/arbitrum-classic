import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre
  const { deploy } = deployments
  const { deployer } = await getNamedAccounts()

  const challenge = await deployments.get('Challenge')
  const osp1 = await deployments.get('OneStepProof')
  const osp2 = await deployments.get('OneStepProof2')

  await deploy('ChallengeFactory', {
    from: deployer,
    args: [challenge.address, osp1.address, osp2.address],
  })
}

module.exports = func
module.exports.tags = ['ChallengeFactory']
module.exports.dependencies = ['Challenge', 'OneStepProof', 'OneStepProof2']
