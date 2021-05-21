import { ethers } from 'hardhat'
import { deploy1820Registry } from './utils'
import { writeFileSync } from 'fs'
import deployments from '../deployment.json'

const main = async () => {
  const accounts = await ethers.getSigners()
  const StandardArbERC20 = await ethers.getContractFactory('StandardArbERC20')

  const standardArbERC20Logic = await StandardArbERC20.deploy()
  await standardArbERC20Logic.deployed()
  console.log(`erc20 logic at ${standardArbERC20Logic.address}`)

  // const ProxyAdmin = await ethers.getContractFactory('ProxyAdmin')
  // const proxyAdmin = await ProxyAdmin.deploy()
  // await proxyAdmin.deployed()
  // console.log("Admin proxy deployed to", proxyAdmin.address)

  const UpgradeableBeacon = await ethers.getContractFactory('UpgradeableBeacon')

  const standardArbERC20Proxy = await UpgradeableBeacon.deploy(
    standardArbERC20Logic.address
  )
  await standardArbERC20Proxy.deployed()
  console.log(`erc20 proxy at ${standardArbERC20Proxy.address}`)

  const contracts = JSON.stringify({
    ...deployments,
    // standardArbERC20: standardArbERC20Logic.address,
    standardArbERC20: standardArbERC20Proxy.address,
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
