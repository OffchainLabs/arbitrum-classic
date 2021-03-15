import { ethers } from 'hardhat'
import { deploy1820Registry } from './utils'
import { writeFileSync } from 'fs'
import deployments from '../deployment.json'


const main = async () => {
  const accounts = await ethers.getSigners()
  const StandardArbERC20 = await ethers.getContractFactory(
    'StandardArbERC20'
  )
  const StandardArbERC777 = await ethers.getContractFactory(
    'StandardArbERC777'
  )
  const standardArbERC20 = await StandardArbERC20.deploy()
  const standardArbERC777 = await StandardArbERC777.deploy()

  const contracts = JSON.stringify({
    ...deployments,
    standardArbERC20: standardArbERC20.address,
    standardArbERC777: standardArbERC777.address,
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
