import { ethers } from 'hardhat'
import { deploy1820Registry } from './utils'

const main = async () => {
  const accounts = await ethers.getSigners()

  const SymmetricBridge = await ethers.getContractFactory(
    'ArbSymmetricTokenBridge'
  )
  const symmetricBridge = await SymmetricBridge.deploy()

  console.log('ArbSymmetricTokenBridge deployed to:', symmetricBridge.address)

  await deploy1820Registry(accounts[0])
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
