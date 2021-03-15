import { ethers } from 'hardhat'
import deployments from '../deployment.json'

const main = async () => {
  // TODO: check buddy deployer address available
  // TODO: check 1820 registry
  const inboxAddress =
    process.env.INBOX_ADDRESS || '0xda0bB0f7aB435B0Fd3dD6Eac8c75D80A3daD6d1F'

  if (inboxAddress === '' || inboxAddress === undefined)
    throw new Error('Please set inbox address! INBOX_ADDRESS')

  const EthERC20Bridge = await ethers.getContractFactory('EthERC20Bridge')


  const gasPrice = 0
  const maxGas = 100000000000
  const ethERC20Bridge = await EthERC20Bridge.deploy(
    inboxAddress,
    deployments.buddyDeployer,
    maxGas,
    gasPrice,
    deployments.standardArbERC20,
    deployments.standardArbERC777
  )

  console.log('EthERC20Bridge deployed to:', ethERC20Bridge.address)

  // TODO: check if L2 counterpart worked
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
