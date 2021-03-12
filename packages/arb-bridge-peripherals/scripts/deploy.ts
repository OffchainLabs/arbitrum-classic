import { ethers } from 'hardhat'
import { deploy1820Registry } from './utils'

const main = async () => {
  const accounts = await ethers.getSigners()

  const StandardArbERC20 = await ethers.getContractFactory(
    'StandardArbERC20'
  )
  const StandardArbERC777 = await ethers.getContractFactory(
    'StandardArbERC777'
  )
  const ArbTokenBridge = await ethers.getContractFactory(
    'ArbTokenBridge'
  )

  await deploy1820Registry(accounts[0])
  const standardArbERC20 = await StandardArbERC777.deploy()
  const standardArbERC777 = await StandardArbERC777.deploy()
  const arbTokenBridge = await ArbTokenBridge.deploy(0, 999999999999, standardArbERC777.address, standardArbERC20.address)

  console.warn('This should be deployed by the L1 EthErc20Bridge.')
  console.log('ArbTokenBridge deployed to:', arbTokenBridge.address)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
