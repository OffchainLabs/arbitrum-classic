import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre
  const { deploy } = deployments
  const { deployer } = await getNamedAccounts()

  const osp1 = await deployments.get('OneStepProof')
  const osp2 = await deployments.get('OneStepProof2')
  const osp3 = await deployments.get('OneStepProofHash')

  await deploy('ChallengeFactory', {
    from: deployer,
    args: [[osp1.address, osp2.address, osp3.address]],
  })
}

module.exports = func
module.exports.tags = ['ChallengeFactory']
module.exports.dependencies = [
  'OneStepProof',
  'OneStepProof2',
  'OneStepProofHash',
]
