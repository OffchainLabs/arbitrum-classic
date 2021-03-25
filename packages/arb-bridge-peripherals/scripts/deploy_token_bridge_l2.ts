import { ethers } from 'hardhat'
import { deploy1820Registry } from './utils'
import { writeFileSync } from 'fs'
import deployments from '../deployment.json'

const main = async () => {
  const accounts = await ethers.getSigners()
  const StandardArbERC20 = await ethers.getContractFactory('StandardArbERC20')
  const StandardArbERC777 = await ethers.getContractFactory('StandardArbERC777')
  const TransparentUpgradeableProxy = await ethers.getContractFactory(
    'TransparentUpgradeableProxy'
  )
  const ProxyAdmin = await ethers.getContractFactory('ProxyAdmin')

  const standardArbERC20 = await StandardArbERC20.deploy()
  const standardArbERC777 = await StandardArbERC777.deploy()
  const proxyAdmin = await ProxyAdmin.deploy()
  const arbERC20Proxy = await TransparentUpgradeableProxy.deploy(
    standardArbERC20.address,
    proxyAdmin.address,
    '0x'
  )
  const arbERC777Proxy = await TransparentUpgradeableProxy.deploy(
    standardArbERC777.address,
    proxyAdmin.address,
    '0x'
  )

  const contracts = JSON.stringify({
    ...deployments,
    standardArbERC20: arbERC20Proxy.address,
    standardArbERC777: arbERC777Proxy.address,
    standardArbERC20Raw: standardArbERC20.address,
    standardArbERC777Raw: standardArbERC777.address,
    proxyAdmin: proxyAdmin.address,
    l2ChainId: ethers.BigNumber.from(
      ethers.provider.network.chainId
    ).toHexString(),
  })

  const path = './deployment.json'
  console.log(`Writing to JSON at ${path}`)

  // TODO: should append/check if previous entries
  writeFileSync(path, contracts)

  console.log('Almost done')
  await deploy1820Registry(accounts[0])
  console.log('Done')
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
