import { Contract } from 'ethers'
import { ethers } from 'hardhat'

type ContractName =
  | 'RollupCreator'
  | 'NodeFactory'
  | 'ChallengeFactory'
  | 'OneStepProof'
  | 'OneStepProof2'

const logDeploy = (contractName: string, contract: Contract) => {
  console.log(
    `Submitting ${contractName} at ${contract.address} in tx ${contract.deployTransaction.hash}`
  )
}

export default async function deploy_contracts(): Promise<
  Record<ContractName, Contract>
> {
  const OneStepProof = await ethers.getContractFactory('OneStepProof')
  const OneStepProof2 = await ethers.getContractFactory('OneStepProof2')
  const ChallengeFactory = await ethers.getContractFactory('ChallengeFactory')
  const NodeFactory = await ethers.getContractFactory('NodeFactory')
  const RollupCreator = await ethers.getContractFactory('RollupCreator')

  const oneStepProof = await OneStepProof.deploy()
  logDeploy('OneStepProof', oneStepProof)
  const oneStepProof2 = await OneStepProof2.deploy()
  logDeploy('OneStepProof2', oneStepProof2)

  const challengeFactory = await ChallengeFactory.deploy(
    oneStepProof.address,
    oneStepProof2.address
  )
  logDeploy('ChallengeFactory', challengeFactory)

  const nodeFactory = await NodeFactory.deploy()
  logDeploy('NodeFactory', nodeFactory)

  const rollupCreator = await RollupCreator.deploy(
    challengeFactory.address,
    nodeFactory.address
  )
  logDeploy('RollupCreator', rollupCreator)

  await Promise.all([
    oneStepProof.deployed().then(() => {
      console.log('OneStepProof deployed')
    }),
    oneStepProof2.deployed().then(() => {
      console.log('OneStepProof2 deployed')
    }),
    challengeFactory.deployed().then(() => {
      console.log('ChallengeFactory deployed')
    }),
    nodeFactory.deployed().then(() => {
      console.log('NodeFactory deployed')
    }),
    rollupCreator.deployed().then(() => {
      console.log('RollupCreator deployed')
    }),
  ])

  const contracts: Record<ContractName, Contract> = {
    RollupCreator: rollupCreator,
    NodeFactory: nodeFactory,
    ChallengeFactory: challengeFactory,
    OneStepProof: oneStepProof,
    OneStepProof2: oneStepProof2,
  }
  return contracts
}
