import { BuidlerRuntimeEnvironment } from '@nomiclabs/buidler/types'
import { Contract } from 'ethers'
import { OneStepProofFactory } from '../build/types/OneStepProofFactory'

type ContractName =
  | 'ArbFactory'
  | 'ChallengeFactory'
  | 'ArbRollup'
  | 'GlobalInbox'
  | 'InboxTopChallenge'
  | 'ExecutionChallenge'
  | 'OneStepProof'

const logDeploy = (contractName: string, contract: Contract) => {
  console.log(
    `Submitting ${contractName} at ${contract.address} in tx ${contract.deployTransaction.hash}`
  )
}

export default async function deploy_contracts(
  bre: BuidlerRuntimeEnvironment
): Promise<Record<ContractName, Contract>> {
  const ethers = bre.ethers

  const UtilLibrary = await ethers.getContractFactory('MerkleUtil')
  const utilLibrary = await UtilLibrary.deploy()
  const ExecutionChallenge = await ethers.getContractFactory(
    'ExecutionChallenge'
  )
  const OneStepProof = await ethers.getContractFactory('OneStepProof')
  const OneStepProof2 = await ethers.getContractFactory('OneStepProof2')
  const InboxTopChallenge = await ethers.getContractFactory('InboxTopChallenge')
  const ArbRollup = await ethers.getContractFactory('ArbRollup')
  const GlobalInbox = await ethers.getContractFactory('GlobalInbox')
  const ChallengeFactory = await ethers.getContractFactory('ChallengeFactory')
  const ArbFactory = await ethers.getContractFactory('ArbFactory')

  const oneStepProof = await OneStepProof.deploy()
  logDeploy('OneStepProof', oneStepProof)
  const oneStepProof2 = await OneStepProof2.deploy()
  logDeploy('OneStepProof2', oneStepProof2)
  const inboxTopChallenge = await InboxTopChallenge.deploy()
  logDeploy('InboxTopChallenge', inboxTopChallenge)
  const executionChallenge = await ExecutionChallenge.deploy()
  logDeploy('ExecutionChallenge', executionChallenge)
  const arbRollup = await ArbRollup.deploy()
  logDeploy('ArbRollup', arbRollup)
  const globalInbox = await GlobalInbox.deploy()
  logDeploy('GlobalInbox', globalInbox)

  const challengeFactory = await ChallengeFactory.deploy(
    inboxTopChallenge.address,
    executionChallenge.address,
    oneStepProof.address,
    oneStepProof2.address
  )
  logDeploy('ChallengeFactory', challengeFactory)

  const arbFactory = await ArbFactory.deploy(
    arbRollup.address,
    globalInbox.address,
    challengeFactory.address
  )
  logDeploy('ArbFactory', arbFactory)

  await Promise.all([
    oneStepProof.deployed().then(() => {
      console.log('OneStepProof deployed')
    }),
    oneStepProof2.deployed().then(() => {
      console.log('OneStepProof2 deployed')
    }),
    inboxTopChallenge.deployed().then(() => {
      console.log('InboxTopChallenge deployed')
    }),
    executionChallenge.deployed().then(() => {
      console.log('ExecutionChallenge deployed')
    }),
    arbRollup.deployed().then(() => {
      console.log('ArbRollup deployed')
    }),
    globalInbox.deployed().then(() => {
      console.log('GlobalInbox deployed')
    }),
    challengeFactory.deployed().then(() => {
      console.log('ChallengeFactory deployed')
    }),
    arbFactory.deployed().then(() => {
      console.log('ArbFactory deployed')
    }),
  ])

  const contracts: Record<ContractName, Contract> = {
    ArbFactory: arbFactory,
    ChallengeFactory: challengeFactory,
    ArbRollup: arbRollup,
    GlobalInbox: globalInbox,
    InboxTopChallenge: inboxTopChallenge,
    ExecutionChallenge: executionChallenge,
    OneStepProof: oneStepProof,
  }
  return contracts
}
