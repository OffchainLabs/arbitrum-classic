import { LogicAddresses } from './deployClassicLogic'
import hre from 'hardhat'
import fs from 'fs'

if (!process.env['ETHERSCAN_API_KEY'])
  throw new Error('Please set ETHERSCAN_API_KEY')

async function verifyContracts(addresses: LogicAddresses) {
  await hre.run('verify:verify', {
    address: addresses.rollup,
    constructorArguments: [1],
    contract: 'contracts/rollup/Rollup.sol:Rollup',
  })

  await hre.run('verify:verify', {
    address: addresses.oldOutbox,
    constructorArguments: [],
  })

  await hre.run('verify:verify', {
    address: addresses.outbox,
    constructorArguments: [],
  })
  await hre.run('verify:verify', {
    address: addresses.rollupAdmin,
    constructorArguments: [],
  })

  await hre.run('verify:verify', {
    address: addresses.rollupUser,
    constructorArguments: [],
  })

  await hre.run('verify:verify', {
    address: addresses.bridge,
    constructorArguments: [],
  })

  await hre.run('verify:verify', {
    address: addresses.inbox,
    constructorArguments: [],
  })

  await hre.run('verify:verify', {
    address: addresses.sequencerInbox,
    constructorArguments: [],
  })

  await hre.run('verify:verify', {
    address: addresses.node,
    constructorArguments: [],
  })
}

const main = async () => {
  const addresses = JSON.parse(
    fs.readFileSync('addresses.json').toString()
  ) as LogicAddresses
  await verifyContracts(addresses)
}

main()
  .then(() => console.log('done'))
  .catch(err => {
    console.error('error')
    console.error(err)
  })
