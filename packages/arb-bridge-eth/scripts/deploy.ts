import deploy_contracts from './deploylib'

async function main() {
  // Buidler always runs the compile task when running scripts through it.
  // If this runs in a standalone fashion you may want to call compile manually
  // to make sure everything is compiled
  // await bre.run('compile');
  const { arb_factory } = await deploy_contracts()
  const fs = require('fs')
  let addresses = {
    ArbFactory: arb_factory.address,
  }
  fs.writeFileSync('bridge_eth_addresses.json', JSON.stringify(addresses))
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
