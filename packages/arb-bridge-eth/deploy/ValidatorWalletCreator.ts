import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, ethers } = hre
  const { deploy } = deployments
  const [deployer] = await ethers.getSigners()

  await deploy('ValidatorWalletCreator', {
    from: await deployer.getAddress(),
    args: [],
  })
}

module.exports = func
module.exports.tags = ['ValidatorWalletCreator', 'live', 'test']
module.exports.dependencies = []
