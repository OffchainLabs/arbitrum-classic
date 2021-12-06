import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, ethers } = hre
  const { deploy } = deployments
  const [deployer] = await ethers.getSigners()

  const bridgeCreator = await deployments.get('BridgeCreator')
  const rollup = await deployments.get('Rollup')
  const challengeFactory = await deployments.get('ChallengeFactory')
  const nodeFactory = await deployments.get('NodeFactory')
  const RollupAdminLogic = await deployments.get('RollupAdminLogic')
  const RollupUserLogic = await deployments.get('RollupUserLogic')

  const dep = await deploy('RollupCreator', {
    from: await deployer.getAddress(),
    args: [],
  })

  const RollupCreator = await ethers.getContractFactory('RollupCreator')
  const rollupCreator = RollupCreator.attach(dep.address).connect(deployer)
  await rollupCreator.setTemplates(
    bridgeCreator.address,
    rollup.address,
    challengeFactory.address,
    nodeFactory.address,
    RollupAdminLogic.address,
    RollupUserLogic.address
  )
}

module.exports = func
module.exports.tags = ['RollupCreator', 'live', 'test']
module.exports.dependencies = [
  'Rollup',
  'ChallengeFactory',
  'NodeFactory',
  'BridgeCreator',
]
