import { ethers } from 'hardhat'
import { deploy1820Registry } from './utils'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import hre from 'hardhat'

const main = async () => {
  const accounts = await ethers.getSigners()
  const { deployments } = hre

  const inbox = await deployments.get('Inbox')
  const buddyDeployer = await deployments.get('BuddyDeployer')

  const EthERC20Bridge = await ethers.getContractFactory('EthERC20Bridge')
  const ethERC20Bridge = await EthERC20Bridge.deploy(
    inbox.address,
    buddyDeployer.address,
    10000000,
    10000000
  )

  console.log('EthERC20Bridge deployed to:', ethERC20Bridge.address)

  await deploy1820Registry(accounts[0])
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
