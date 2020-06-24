import { Artifact, BuidlerRuntimeEnvironment } from '@nomiclabs/buidler/types'
import { Contract } from 'ethers'
const { readArtifact } = require('@nomiclabs/buidler/plugins')

interface ContractInfo {
  path: string
  address: string
  libraries: string
  constructorArguments: string[]
}

type ContractName =
  | 'ArbFactory'
  | 'ChallengeFactory'
  | 'ArbRollup'
  | 'GlobalInbox'
  | 'InboxTopChallenge'
  | 'MessagesChallenge'
  | 'ExecutionChallenge'

const logDeploy = (contractName: string, contract: Contract) => {
  console.log(
    `Submitting ${contractName} at ${contract.address} in tx ${contract.deployTransaction.hash}`
  )
}

export default async function deploy_contracts(bre: BuidlerRuntimeEnvironment) {
  const ethers = bre.ethers
  const config = bre.config

  const OneStepProof = await ethers.getContractFactory('OneStepProof')
  const MessagesChallenge = await ethers.getContractFactory('MessagesChallenge')
  const InboxTopChallenge = await ethers.getContractFactory('InboxTopChallenge')
  const ArbRollup = await ethers.getContractFactory('ArbRollup')
  const GlobalInbox = await ethers.getContractFactory('GlobalInbox')
  const ChallengeFactory = await ethers.getContractFactory('ChallengeFactory')
  const ArbFactory = await ethers.getContractFactory('ArbFactory')

  const one_step_proof = await OneStepProof.deploy()
  logDeploy('OneStepProof', one_step_proof)

  const cArtifact = await readArtifact(
    config.paths.artifacts,
    'ExecutionChallenge'
  )
  const linkedBytecode = linkBytecode(cArtifact, {
    OneStepProof: one_step_proof.address,
  })
  const ExecutionChallenge = await ethers.getContractFactory(
    cArtifact.abi,
    linkedBytecode
  )

  const message_challenge = await MessagesChallenge.deploy()
  logDeploy('MessagesChallenge', message_challenge)
  const inbox_top_challenge = await InboxTopChallenge.deploy()
  logDeploy('InboxTopChallenge', inbox_top_challenge)
  const execution_challenge = await ExecutionChallenge.deploy()
  logDeploy('ExecutionChallenge', execution_challenge)
  const arb_rollup = await ArbRollup.deploy()
  logDeploy('ArbRollup', arb_rollup)
  const global_inbox = await GlobalInbox.deploy()
  logDeploy('GlobalInbox', global_inbox)

  const challenge_factory = await ChallengeFactory.deploy(
    message_challenge.address,
    inbox_top_challenge.address,
    execution_challenge.address
  )
  logDeploy('ChallengeFactory', challenge_factory)

  const arb_factory = await ArbFactory.deploy(
    arb_rollup.address,
    global_inbox.address,
    challenge_factory.address
  )
  logDeploy('ArbFactory', arb_factory)

  await Promise.all([
    one_step_proof.deployed().then(() => {
      console.log('OneStepProof deployed')
    }),
    message_challenge.deployed().then(() => {
      console.log('MessagesChallenge deployed')
    }),
    inbox_top_challenge.deployed().then(() => {
      console.log('InboxTopChallenge deployed')
    }),
    execution_challenge.deployed().then(() => {
      console.log('ExecutionChallenge deployed')
    }),
    arb_rollup.deployed().then(() => {
      console.log('ArbRollup deployed')
    }),
    global_inbox.deployed().then(() => {
      console.log('GlobalInbox deployed')
    }),
    challenge_factory.deployed().then(() => {
      console.log('ChallengeFactory deployed')
    }),
    arb_factory.deployed().then(() => {
      console.log('ArbFactory deployed')
    }),
  ])

  const contracts: Record<ContractName, ContractInfo> = {
    ArbFactory: {
      path: 'contracts/vm/ArbFactory.sol',
      address: arb_factory.address,
      libraries: JSON.stringify({}),
      constructorArguments: [
        arb_rollup.address,
        global_inbox.address,
        challenge_factory.address,
      ],
    },
    ChallengeFactory: {
      path: 'contracts/challenge/ChallengeFactory.sol',
      address: challenge_factory.address,
      libraries: JSON.stringify({}),
      constructorArguments: [
        message_challenge.address,
        inbox_top_challenge.address,
        execution_challenge.address,
      ],
    },
    ArbRollup: {
      path: 'contracts/vm/ArbRollup.sol',
      address: arb_rollup.address,
      libraries: JSON.stringify({}),
      constructorArguments: [],
    },
    GlobalInbox: {
      path: 'contracts/GlobalInbox.sol',
      address: global_inbox.address,
      libraries: JSON.stringify({}),
      constructorArguments: [],
    },
    InboxTopChallenge: {
      path: 'contracts/challenge/InboxTopChallenge.sol',
      address: inbox_top_challenge.address,
      libraries: JSON.stringify({}),
      constructorArguments: [],
    },
    MessagesChallenge: {
      path: 'contracts/challenge/MessagesChallenge.sol',
      address: message_challenge.address,
      libraries: JSON.stringify({}),
      constructorArguments: [],
    },
    ExecutionChallenge: {
      path: 'contracts/challenge/ExecutionChallenge.sol',
      address: execution_challenge.address,
      libraries: JSON.stringify({ OneStepProof: one_step_proof.address }),
      constructorArguments: [],
    },
  }
  return contracts
}

function linkBytecode(artifact: Artifact, libraries: any) {
  let bytecode = artifact.bytecode

  for (const [fileName, fileReferences] of Object.entries(
    artifact.linkReferences
  )) {
    for (const [libName, fixups] of Object.entries(fileReferences)) {
      const addr = libraries[libName]
      if (addr === undefined) {
        continue
      }

      for (const fixup of fixups) {
        bytecode =
          bytecode.substr(0, 2 + fixup.start * 2) +
          addr.substr(2) +
          bytecode.substr(2 + (fixup.start + fixup.length) * 2)
      }
    }
  }

  return bytecode
}
