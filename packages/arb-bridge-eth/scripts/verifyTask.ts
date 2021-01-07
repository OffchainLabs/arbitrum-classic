import { task } from 'hardhat/config'

export default function setupVerifyTask(): void {
  task(
    'verify-contracts',
    'Verifies arbitrum deployment on etherscan',
    async (_taskArgs, bre) => {
      const { deployments } = bre
      await deployments.getOrNull('ArbFactory')
      const contracts = await deployments.all()

      console.log('Verifying contracts on etherscan')
      for (const contractName in contracts) {
        const contractInfo = contracts[contractName as keyof typeof contracts]
        console.log(`Verifying ${contractName}`)
        try {
          await bre.run('verify', {
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
