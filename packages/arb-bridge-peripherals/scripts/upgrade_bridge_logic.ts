import { BigNumber } from '@ethersproject/bignumber'
import { ethers } from 'hardhat'
// import { TransparentUpgradeableProxy__factory } from "arb-ts/src/lib/abi/factories/TransparentUpgradeableProxy__factory"
import { ProxyAdmin__factory } from 'arb-ts/src/lib/abi/factories/ProxyAdmin__factory'

const main = async () => {
  const accounts = await ethers.getSigners()
  // The standard arb erc20 uses a beacon proxy, so this update method won't work!
  const proxyAddress = '0x64b92d4f02cE1b4BDE2D16B6eAEe521E27f28e07'
  const oldLogicAddress = '0xeA86B7cA0F476012A03a03C3B9641692b8c5D5b3'

  const getAdmin = () => {
    // beacon stores this somewhere different
    const adminStorageSlot =
      '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103'
    return ethers.provider.getStorageAt(proxyAddress, adminStorageSlot)
  }

  const admin = await getAdmin()

  const proxyAdmin = await ProxyAdmin__factory.connect(
    admin,
    accounts[0]
  ).deployed()

  const currAdmin = await proxyAdmin.getProxyAdmin(proxyAddress)
  if (currAdmin.toLowerCase() !== proxyAdmin.address.toLowerCase())
    throw new Error("ProxyAdmin doens't control TransparentProxy")

  const proxyAdminOwner = await proxyAdmin.owner()

  if (proxyAdminOwner.toLowerCase() !== accounts[0].address.toLowerCase()) {
    console.log('Current admin', currAdmin)
    console.log('Current account', accounts[0].address)
    throw new Error('Current account does not control ProxyAdmin')
  }

  const currLogicAddress = await proxyAdmin.getProxyImplementation(proxyAddress)
  if (currLogicAddress !== oldLogicAddress)
    throw new Error(
      'You are expecting the wrong logic address for the current implementation'
    )

  // deploy new logic
  console.log('Deploying logic contract')
  const Factory = await ethers.getContractFactory('ArbTokenBridge')
  const logicContract = await Factory.deploy()
  await logicContract.deployed()
  console.log('Deployed logic contract')
  // Watch out as initialize is already called!

  console.log('Upgrading logic contract')
  const upgradeTx = await proxyAdmin.upgrade(
    proxyAddress,
    logicContract.address
  )
  const upgradeTxReceipt = await upgradeTx.wait()
  console.log('Upgraded logic address')

  // verify proxy points to new logic
  console.log('Verifying new logic address')
  const newLogicAddress = await proxyAdmin.getProxyImplementation(proxyAddress)

  if (newLogicAddress.toLowerCase() !== logicContract.address.toLowerCase()) {
    console.log('logicContract deployTx')
    console.log(logicContract.deployTransaction)

    console.log('upgradeTxReceipt')
    console.log(upgradeTxReceipt)

    throw new Error('New logic address is not correctly set.')
  }

  console.log('Upgrade success!')
  console.log(`Old logic address ${oldLogicAddress}`)
  console.log(`New logic address ${logicContract.address}`)
  console.log(`In transaction ${upgradeTx.hash}`)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
