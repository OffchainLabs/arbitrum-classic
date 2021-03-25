import { BigNumber } from '@ethersproject/bignumber'
import { concat, id, keccak256, zeroPad } from 'ethers/lib/utils'
import { ethers } from 'hardhat'
import deployments from '../deployment.json'

import { Bridge } from 'arb-ts/src'
import { BridgeHelper } from 'arb-ts/src/lib/bridge_helpers'
import { writeFileSync } from 'fs'
import { spawnSync } from 'child_process'

const main = async () => {
  const accounts = await ethers.getSigners()
  const inboxAddress =
    process.env.INBOX_ADDRESS || '0x0d0c1aDf6523D422ec7192506A7F6790253399fB'

  if (inboxAddress === '' || inboxAddress === undefined)
    throw new Error('Please set inbox address! INBOX_ADDRESS')

  console.log('deployer', accounts[0].address)

  const EthERC20Bridge = await ethers.getContractFactory('EthERC20Bridge')

  if (
    deployments.buddyDeployer === '' ||
    deployments.standardArbERC20 === '' ||
    deployments.standardArbERC777 === ''
  )
    throw new Error("Deployments.json doesn't include the necessary addresses")

  const maxSubmissionCost = 0
  const gasPrice = 0
  const maxGas = 100000000000
  const ethERC20Bridge = await EthERC20Bridge.deploy(
    inboxAddress,
    deployments.buddyDeployer,
    maxSubmissionCost,
    maxGas,
    gasPrice,
    deployments.standardArbERC777,
    deployments.standardArbERC20
  )

  await ethERC20Bridge.deployed()

  const arbTokenBridge = await ethERC20Bridge.l2Buddy()
  console.log('EthERC20Bridge deployed to:', ethERC20Bridge.address)
  console.log('L2 ArbBridge deployed to:', arbTokenBridge)

  const contracts = JSON.stringify({
    ...deployments,
    ethERC20Bridge: ethERC20Bridge.address,
    arbTokenBridge: arbTokenBridge,
  })
  const deployFilePath = './deployment.json'
  console.log(`Writing to JSON at ${deployFilePath}`)
  writeFileSync(deployFilePath, contracts)

  const deployReceipt = await ethers.provider.getTransactionReceipt(
    ethERC20Bridge.deployTransaction.hash
  )

  const seqNums = await BridgeHelper.getInboxSeqNumFromContractTransaction(
    deployReceipt,
    inboxAddress
  )

  if (!seqNums) throw new Error("Transaction didn't trigger inbox")
  if (seqNums.length !== 1)
    throw new Error('Transaction triggered inbox more than once')

  const inboxSequenceNumber = seqNums[0]

  // get proof data from L2 with child process
  const l2NetworkName: string = 'arbitrum'
  const { stdout, stderr } = await spawnSync(
    'yarn',
    ['hardhat', 'run', 'scripts/get_proof_l2.ts', '--network', l2NetworkName],
    {
      env: {
        ...process.env,
        INBOX_SEQ_NUM: inboxSequenceNumber.toHexString(),
      },
    }
  )

  if (stderr) {
    console.log(stderr)
    throw new Error('Error getting proof')
  }
  const proofData = JSON.parse(stdout.toString('utf-8'))
  // const proofData = JSON.parse(stdout)

  console.log('Got proof data')
  console.log(proofData)

  // trigger in L1
  const coreBridge = await BridgeHelper.getCoreBridgeFromInbox(inboxAddress, ethers.provider)

  const l1TxReceipt = await BridgeHelper.tryOutboxExecute(
    proofData,
    coreBridge,
    accounts[0]
  )

  console.log('Transaction executed in L1')
  console.log(l1TxReceipt)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
