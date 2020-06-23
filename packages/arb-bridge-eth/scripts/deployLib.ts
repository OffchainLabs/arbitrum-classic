import { Artifact, BuidlerRuntimeEnvironment } from '@nomiclabs/buidler/types'
import { Contract } from 'ethers'
const { readArtifact } = require('@nomiclabs/buidler/plugins')

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
  const execution_challenge_promise = one_step_proof
    .deployed()
    .then(async (one_step_proof: Contract) => {
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
      return await ExecutionChallenge.deploy()
    })

  const contractPromises1 = []
  contractPromises1.push(MessagesChallenge.deploy())
  contractPromises1.push(InboxTopChallenge.deploy())
  contractPromises1.push(execution_challenge_promise)
  contractPromises1.push(ArbRollup.deploy())
  contractPromises1.push(GlobalInbox.deploy())

  const [
    message_challenge,
    inbox_top_challenge,
    execution_challenge,
    arb_rollup,
    global_inbox,
  ] = await Promise.all(contractPromises1)

  const challenge_factory = await ChallengeFactory.deploy(
    message_challenge.address,
    inbox_top_challenge.address,
    execution_challenge.address
  )

  const arb_factory = await ArbFactory.deploy(
    arb_rollup.address,
    global_inbox.address,
    challenge_factory.address
  )

  await one_step_proof.deployed()
  await message_challenge.deployed()
  await inbox_top_challenge.deployed()
  await execution_challenge.deployed()
  await arb_rollup.deployed()
  await global_inbox.deployed()
  await challenge_factory.deployed()
  await arb_factory.deployed()

  console.log('ArbFactory deployed at', arb_factory.address)

  return {
    arb_factory,
    challenge_factory,
    arb_rollup,
    global_inbox,
    inbox_top_challenge,
    message_challenge,
    execution_challenge,
  }
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
