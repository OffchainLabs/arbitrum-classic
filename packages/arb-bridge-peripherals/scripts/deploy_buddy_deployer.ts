import { ethers } from 'hardhat'
import { writeFileSync } from 'fs'

const main = async () => {
  const BuddyDeployer = await ethers.getContractFactory('BuddyDeployer')
  const buddyDeployer = await BuddyDeployer.deploy()
  console.log('BuddyDeployer deployed to:', buddyDeployer.address)

  const contracts = JSON.stringify({
    buddyDeployer: buddyDeployer.address,
  })
  const path = './deployment.json'
  console.log(`Writing to JSON at ${path}`)

  // TODO: should append/check if previous entries
  writeFileSync(path, contracts)
  
  console.log('Done')
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
