import { ethers } from 'hardhat'
import deployments from '../deployment.json'
import { writeFileSync } from 'fs'

const main = async () => {
  // TODO: check buddy deployer address available
  // TODO: check 1820 registry
  const inboxAddress =
    process.env.INBOX_ADDRESS || '0x0d0c1aDf6523D422ec7192506A7F6790253399fB'

  if (inboxAddress === '' || inboxAddress === undefined)
    throw new Error('Please set inbox address! INBOX_ADDRESS')

  console.log('deployer', (await ethers.getSigners())[0].address)

  const SafeERC20Namer = await ethers.getContractFactory('SafeERC20Namer')
  const safeERC20Namer = await SafeERC20Namer.deploy()

  const EthERC20Bridge = await ethers.getContractFactory('EthERC20Bridge', {
    libraries: {
      SafeERC20Namer: safeERC20Namer.address,
    },
  })
  const gasPrice = 0
  const maxGas = 100000000000
  const ethERC20Bridge = await EthERC20Bridge.deploy(
    inboxAddress,
    deployments.buddyDeployer,
    maxGas,
    gasPrice,
    deployments.standardArbERC777,
    deployments.standardArbERC20
  )
  const arbTokenBridge = await ethERC20Bridge.l2Buddy()
  console.log('EthERC20Bridge deployed to:', ethERC20Bridge.address)
  console.log('L2 ArbBridge deployed to:', arbTokenBridge)

  const contracts = JSON.stringify({
    ...deployments,
    ethERC20Bridge: ethERC20Bridge.address,
    arbTokenBridge: arbTokenBridge,
  })
  const path = './deployment.json'
  console.log(`Writing to JSON at ${path}`)
  writeFileSync(path, contracts)

  // TODO: check if L2 counterpart worked
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
