import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre
  const { deploy } = deployments
  const { deployer } = await getNamedAccounts()

  await deploy('OneStepProof2', {
    from: deployer,
    args: [],
  })
}

module.exports = func
module.exports.tags = ['OneStepProof2', 'live', 'test']
