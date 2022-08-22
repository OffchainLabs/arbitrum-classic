import { HardhatRuntimeEnvironment } from 'hardhat/types'
import { DeployFunction } from 'hardhat-deploy/types'

const func: DeployFunction = async (hre: HardhatRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = hre
  const { deploy } = deployments
  const { deployer } = await getNamedAccounts()

  const NitroMigrator = await deployments.get('NitroMigrator')
  const ProxyAdmin = await deployments.get('ProxyAdmin')

  await deploy('TransparentUpgradeableProxy', {
    from: deployer,
    args: [NitroMigrator.address, ProxyAdmin.address, '0x8129fc1c'],
  })
}

module.exports = func
module.exports.tags = ['NitroMigratorProxy', 'live']
module.exports.dependencies = ['NitroMigrator', 'NitroMigratorProxyAdmin']
