import { task } from '@nomiclabs/buidler/config'

const EtherScanUrls = {
  1: 'https://api.etherscan.io/api',
  3: 'https://api-ropsten.etherscan.io/api',
  4: 'https://api-rinkeby.etherscan.io/api',
  42: 'https://api-kovan.etherscan.io/api',
}

const ContractLocations = {
  ArbFactory: 'contracts/vm/ArbFactory.sol',
  ArbRollup: 'contracts/vm/ArbRollup.sol',
  ChallengeFactory: 'contracts/challenge/ChallengeFactory.sol',
  ExecutionChallenge: 'contracts/challenge/ExecutionChallenge.sol',
  GlobalInbox: 'contracts/GlobalInbox.sol',
  InboxTopChallenge: 'contracts/challenge/InboxTopChallenge.sol',
  MessagesChallenge: 'contracts/challenge/MessagesChallenge.sol',
}

export default function setupVerifyTask(): void {
  task(
    'verify',
    'Verifies arbitrum deployment on etherscan',
    async (taskArgs, bre) => {
      const { deployments } = bre
      await deployments.getOrNull('ArbFactory')
      const contracts = await deployments.all()
      const network = await bre.ethers.provider.getNetwork()
      const chainId = network.chainId

      if (bre.config.etherscan && chainId in EtherScanUrls) {
        console.log('Verifying contracts on etherscan')
        bre.config.etherscan.url =
          EtherScanUrls[chainId as keyof typeof EtherScanUrls]
        for (const contractName in contracts) {
          const contractInfo = contracts[contractName as keyof typeof contracts]
          const location =
            ContractLocations[contractName as keyof typeof ContractLocations]
          console.log(`Verifying ${contractName}`)
          try {
            await bre.run('verify-contract', {
              contractName: `${location}:${contractName}`,
              address: contractInfo.address,
              constructorArguments: contractInfo.args,
            })
            console.log(`Verified ${contractName}`)
          } catch (e) {
            console.log('Failed to verify', e)
          }
        }
      }
    }
  )
}
