import { task } from '@nomiclabs/buidler/config'
import fs from 'fs'

export default function setupVerifyTask(): void {
  task(
    'verify-contracts',
    'Verifies arbitrum deployment on etherscan',
    async (taskArgs, bre) => {
      const { deployments } = bre
      await deployments.getOrNull('ArbFactory')
      const contracts = await deployments.all()
      const network = await bre.ethers.provider.getNetwork()
      const chainId = network.chainId

      console.log('Verifying contracts on etherscan')
      for (const contractName in contracts) {
        const contractInfo = contracts[contractName as keyof typeof contracts]
        console.log(`Verifying ${contractName}`)
        try {
          await bre.run('verify-contract', {
            contractName,
            address: contractInfo.address,
            constructorArguments: contractInfo.args,
          })
          console.log(`Verified ${contractName}`)
        } catch (e) {
          console.log('Failed to verify', e)
        }
      }
    }
  )
}
