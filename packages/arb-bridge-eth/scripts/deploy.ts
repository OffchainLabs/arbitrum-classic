import { BuidlerRuntimeEnvironment } from '@nomiclabs/buidler/types'
import { Contract } from 'ethers'

type ContractName =
  | 'RollupCreator'
  | 'ChallengeFactory'
  | 'Challenge'
  | 'OneStepProof'
  | 'OneStepProof2'

const logDeploy = (contractName: string, contract: Contract) => {
  console.log(
    `Submitting ${contractName} at ${contract.address} in tx ${contract.deployTransaction.hash}`
  )
}

export default async function deploy_contracts(
  bre: BuidlerRuntimeEnvironment
): Promise<Record<ContractName, Contract>> {
  const ethers = bre.ethers

  const Challenge = await ethers.getContractFactory('Challenge')
  const OneStepProof = await ethers.getContractFactory('OneStepProof')
  const OneStepProof2 = await ethers.getContractFactory('OneStepProof2')
  const ChallengeFactory = await ethers.getContractFactory('ChallengeFactory')
  const RollupCreator = await ethers.getContractFactory('RollupCreator')

  const oneStepProof = await OneStepProof.deploy()
  logDeploy('OneStepProof', oneStepProof)
  const oneStepProof2 = await OneStepProof2.deploy()
  logDeploy('OneStepProof2', oneStepProof2)
  const challenge = await Challenge.deploy()
  logDeploy('Challenge', challenge)

  const challengeFactory = await ChallengeFactory.deploy(
    challenge.address,
    oneStepProof.address,
    oneStepProof2.address
  )
  logDeploy('ChallengeFactory', challengeFactory)

  const rollupCreator = await RollupCreator.deploy(challengeFactory.address)
  logDeploy('RollupCreator', rollupCreator)

  await Promise.all([
    oneStepProof.deployed().then(() => {
      console.log('OneStepProof deployed')
    }),
    oneStepProof2.deployed().then(() => {
      console.log('OneStepProof2 deployed')
    }),
    challenge.deployed().then(() => {
      console.log('Challenge deployed')
    }),
    challengeFactory.deployed().then(() => {
      console.log('ChallengeFactory deployed')
    }),
    rollupCreator.deployed().then(() => {
      console.log('RollupCreator deployed')
    }),
  ])

  const contracts: Record<ContractName, Contract> = {
    RollupCreator: rollupCreator,
    ChallengeFactory: challengeFactory,
    Challenge: challenge,
    OneStepProof: oneStepProof,
    OneStepProof2: oneStepProof2,
  }
  return contracts
}
