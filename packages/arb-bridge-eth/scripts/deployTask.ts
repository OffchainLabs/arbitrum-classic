import { task } from '@nomiclabs/buidler/config'
import deploy_contracts from './deployLib'
import fs from 'fs'

const EtherScanUrls = {
  1: 'https://api.etherscan.io/api',
  3: 'https://api-ropsten.etherscan.io/api',
  4: 'https://api-rinkeby.etherscan.io/api',
  42: 'https://api-kovan.etherscan.io/api',
}

export default function setupDeployTask() {
  task('deploy', 'Prints the list of accounts', async (taskArgs, bre) => {
    await bre.run('compile')
    const contracts = await deploy_contracts(bre)
    const network = await bre.ethers.provider.getNetwork()
    const chainId = network.chainId
    fs.writeFileSync(`deployment_${chainId}.json`, JSON.stringify(contracts))
  })

  task(
    'verify',
    'Verifies arbitrum deployment on etherscan',
    async (taskArgs, bre) => {
      const network = await bre.ethers.provider.getNetwork()
      const chainId = network.chainId
      let rawdata = fs.readFileSync(`deployment_${chainId}.json`)
      let contracts = JSON.parse(rawdata.toString())

      if (bre.config.etherscan && chainId in EtherScanUrls) {
        console.log('Verifying contracts on etherscan')
        bre.config.etherscan.url =
          EtherScanUrls[chainId as keyof typeof EtherScanUrls]
        for (const contractName in contracts) {
          const contractInfo = contracts[contractName as keyof typeof contracts]
          console.log(`Verifying ${contractName}`)
          await bre.run('verify-contract', {
            contractName: `${contractInfo.path}:${contractName}`,
            address: contractInfo.address,
            libraries: contractInfo.libraries,
            constructorArguments: contractInfo.constructorArguments,
          })
          console.log(`Verified ${contractName}`)
        }
      } else {
        console.log('Not verifying contracts since network is unsupported')
        for (const contractName in contracts) {
          const contractInfo = contracts[contractName as keyof typeof contracts]
          console.log('verify-contract', {
            contractName: `${contractInfo.path}:${contractName}`,
            address: contractInfo.address,
            libraries: contractInfo.libraries,
            constructorArguments: contractInfo.constructorArguments,
          })
        }
      }
    }
  )
}
