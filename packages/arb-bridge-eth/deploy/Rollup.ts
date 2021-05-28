import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre
  const { deploy } = deployments
  const { deployer } = await getNamedAccounts()

  await deploy('Rollup', {
    from: deployer,
    args: [],
  })

  await deploy('RollupUserFacet', {
    from: deployer,
    args: [],
  })

  await deploy('RollupAdminFacet', {
    from: deployer,
    args: [],
  })
}

module.exports = func
module.exports.tags = ['Rollup', 'live', 'test']
