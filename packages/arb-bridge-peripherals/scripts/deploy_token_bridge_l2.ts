import { ethers } from 'hardhat'
import { deploy1820Registry } from './utils'
import { HardhatRuntimeEnvironment } from 'hardhat/types'
import hre from 'hardhat'

const main = async () => {
  const accounts = await ethers.getSigners()
  const { deployments } = hre

  await deploy1820Registry(accounts[0])
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
