import { ethers } from 'hardhat'
const main = async () => {
  const BuddyDeployer = await ethers.getContractFactory('BuddyDeployer')
  const buddyDeployer = await BuddyDeployer.deploy()
  console.log('BuddyDeployer deployed to:', buddyDeployer.address)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
