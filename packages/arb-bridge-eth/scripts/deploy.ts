import { Contract } from 'ethers'
import { ethers } from 'hardhat'

type ContractName =
  | 'RollupCreatorNoProxy'
  | 'BridgeCreator'
  | 'NodeFactory'
  | 'ChallengeFactory'
  | 'OneStepProof'
  | 'OneStepProof2'
  | 'OneStepProofHash'

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
  const OneStepProofHash = await ethers.getContractFactory('OneStepProofHash')
  const ChallengeFactory = await ethers.getContractFactory('ChallengeFactory')
  const NodeFactory = await ethers.getContractFactory('NodeFactory')
  const Rollup = await ethers.getContractFactory('Rollup')
  const BridgeCreator = await ethers.getContractFactory('BridgeCreator')
  const RollupCreatorNoProxy = await ethers.getContractFactory(
    'RollupCreatorNoProxy'
  )

  const oneStepProof = await OneStepProof.deploy()
  logDeploy('OneStepProof', oneStepProof)
  const oneStepProof2 = await OneStepProof2.deploy()
  logDeploy('OneStepProof2', oneStepProof2)
  const oneStepProof3 = await OneStepProofHash.deploy()
  logDeploy('OneStepProofHash', oneStepProof3)

  const challengeFactory = await ChallengeFactory.deploy([
    oneStepProof.address,
    oneStepProof2.address,
    oneStepProof3.address,
  ])
  logDeploy('ChallengeFactory', challengeFactory)

  const nodeFactory = await NodeFactory.deploy()
  logDeploy('NodeFactory', nodeFactory)

  const rollupTemplate = await Rollup.deploy()
  logDeploy('Rollup', rollupTemplate)

  const bridgeCreator = await BridgeCreator.deploy()
  logDeploy('BridgeCreator', bridgeCreator)

  const rollupCreatorNoProxy = await RollupCreatorNoProxy.deploy()
  logDeploy('RollupCreatorNoProxy', rollupCreatorNoProxy)

  await Promise.all([
    oneStepProof.deployed().then(() => {
      console.log('OneStepProof deployed')
    }),
    oneStepProof2.deployed().then(() => {
      console.log('OneStepProof2 deployed')
    }),
    oneStepProof3.deployed().then(() => {
      console.log('OneStepProofHash deployed')
    }),
    challengeFactory.deployed().then(() => {
      console.log('ChallengeFactory deployed')
    }),
    nodeFactory.deployed().then(() => {
      console.log('NodeFactory deployed')
    }),
    rollupTemplate.deployed().then(() => {
      console.log('Rollup deployed')
    }),
    bridgeCreator.deployed().then(() => {
      console.log('BridgeCreator deployed')
    }),
    rollupCreatorNoProxy.deployed().then(() => {
      console.log('RollupCreatorNoProxy deployed')
    }),
  ])

  await rollupCreatorNoProxy.setTemplates(
    bridgeCreator.address,
    rollupTemplate.address,
    challengeFactory.address,
    nodeFactory.address
  )

  const contracts: Record<ContractName, Contract> = {
    RollupCreatorNoProxy: rollupCreatorNoProxy,
    BridgeCreator: bridgeCreator,
    NodeFactory: nodeFactory,
    ChallengeFactory: challengeFactory,
    OneStepProof: oneStepProof,
    OneStepProof2: oneStepProof2,
    OneStepProofHash: oneStepProof3,
  }
  return contracts
}
