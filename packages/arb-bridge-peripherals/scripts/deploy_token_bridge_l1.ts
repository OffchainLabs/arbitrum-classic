import { BigNumber } from '@ethersproject/bignumber'
import { concat, id, keccak256, zeroPad } from 'ethers/lib/utils'
import { ethers } from 'hardhat'
import deployments from '../deployment.json'
import {BridgeFactory} from "arb-ts/src/lib/abi/BridgeFactory"
import {InboxFactory} from "arb-ts/src/lib/abi/InboxFactory"
import {OutboxFactory} from "arb-ts/src/lib/abi/OutboxFactory"

const main = async () => {
  const accounts = await ethers.getSigners()
  // TODO: check buddy deployer address available
  // TODO: check 1820 registry
  const inboxAddress =
    process.env.INBOX_ADDRESS || '0x0d0c1aDf6523D422ec7192506A7F6790253399fB'

  if (inboxAddress === '' || inboxAddress === undefined)
    throw new Error('Please set inbox address! INBOX_ADDRESS')

  console.log('deployer', accounts[0].address)

  const SafeERC20Namer = await ethers.getContractFactory('SafeERC20Namer')
  const safeERC20Namer = await SafeERC20Namer.deploy()

  const EthERC20Bridge = await ethers.getContractFactory('EthERC20Bridge', {
    libraries: {
      SafeERC20Namer: safeERC20Namer.address,
    },
  })

  const gasPrice = 0
  const maxGas = 100000000000
  const ethERC20Bridge = await EthERC20Bridge.deploy(
    inboxAddress,
    deployments.buddyDeployer,
    maxGas,
    gasPrice,
    deployments.standardArbERC777,
    deployments.standardArbERC20
  )

  console.log('EthERC20Bridge deployed to:', ethERC20Bridge.address)
  console.log('L2 ArbBridge deployed to:', await ethERC20Bridge.l2Buddy())

  const deployReceipt = await ethers.provider.getTransactionReceipt(
    ethERC20Bridge.deployTransaction.hash
  )

  const inboxEventSignature = [
    id('InboxMessageDelivered(uint256,bytes)'),
    id('InboxMessageDeliveredFromOrigin(uint256)'),
  ]
  const inboxEvent = deployReceipt.logs.filter(
    log =>
      log.topics[0] === inboxEventSignature[0] ||
      log.topics[0] === inboxEventSignature[1]
  )

  if (inboxEvent.length !== 1) {
    console.log(inboxEvent)
    throw new Error('Triggered inbox multiple times?')
  }
  const inboxSequenceNumber = inboxEvent[0].topics[1]

  const l2DeployTxHash = keccak256(
    concat([
      zeroPad(deployments.l2ChainId, 32),
      zeroPad(inboxSequenceNumber, 32),
    ])
  )

  const l2Provider = new ethers.providers.JsonRpcProvider(
    'https://devnet-l2.arbitrum.io/rpc'
  )
  const l2TxReceipt = await l2Provider.getTransactionReceipt(l2DeployTxHash)

  const l2ToL1EventId = id(
    'L2ToL1Transaction(address,address,uint256,uint256,uint256,uint256,uint256,uint256,uint256,bytes)'
  )
  const logs = l2TxReceipt.logs.filter(log => log.topics[0] === l2ToL1EventId)

  if (logs.length !== 1)
    throw new Error('Not exactly 1 log emitted of L2 to L1 tx?')

  const abi = [
    `event L2ToL1Transaction(address caller, address indexed destination, uint indexed uniqueId,
      uint indexed batchNumber, uint indexInBatch, uint arbBlockNum,
      uint ethBlockNum, uint timestamp, uint callvalue, bytes data)`,
  ]
  const iface = new ethers.utils.Interface(abi)

  const {
    caller,
    destination,
    uniqueId,
    batchNumber,
    indexInBatch,
    arbBlockNum,
    ethBlockNum,
    timestamp,
    callvalue,
    data,
  } = iface.parseLog(logs[0]).args

  const wait = (ms: number) => new Promise(res => setTimeout(res, ms))

  const getProof = async (): Promise<{
    proof: Array<string>
    path: BigNumber
    l2Sender: string
    l1Dest: string
    l2Block: BigNumber
    l1Block: BigNumber
    timestamp: BigNumber
    amount: BigNumber
    calldataForL1: string
  }> => {
    const nodeInterfaceAddress = '0x00000000000000000000000000000000000000C8'
    const foo = new ethers.utils.Interface([
      `function lookupMessageBatchProof(uint256 batchNum, uint64 index)
          external
          view
          returns (
              bytes32[] memory proof,
              uint256 path,
              address l2Sender,
              address l1Dest,
              uint256 l2Block,
              uint256 l1Block,
              uint256 timestamp,
              uint256 amount,
              bytes memory calldataForL1
          )`,
    ])
    const nodeInterface = new ethers.Contract(
      nodeInterfaceAddress,
      foo
    ).connect(l2Provider)
    try {
      const res = await nodeInterface.callStatic.lookupMessageBatchProof(
        batchNumber,
        indexInBatch
      )
      return res
    } catch (e) {
      const expectedError = "batch doesn't exist"
      if (e.error.message === expectedError) {
        console.log(
          'Withdrawal detected, but batch not created yet. Going to wait a bit.'
        )
        await wait(1000)
      } else {
        console.log("Withdrawal proof didn't work. Not sure why")
        console.log(e)
        console.log('Going to try again after waiting')
        await wait(1000)
      }
      console.log('New attempt starting')
      return getProof()
    }
  }

  const {
    proof,
    path,
    l2Sender,
    l1Dest,
    l2Block,
    l1Block,
    timestamp: proofTimestamp,
    amount,
    calldataForL1,
  } = await getProof()
  
    const inbox = InboxFactory.connect(inboxAddress, ethers.provider)
    const bridge = BridgeFactory.connect(await inbox.bridge(), ethers.provider)
    const outbox = OutboxFactory.connect(await bridge.activeOutbox(), ethers.provider).connect(accounts[0])

    // TODO: wait until assertion is confirmed before execute
    console.log("executing outbox")
    const outboxExecute = await outbox.executeTransaction(
      batchNumber, proof, path, l2Sender, l1Dest, l2Block,
      l1Block, proofTimestamp, amount, calldataForL1
    )
    console.log("executed")
    console.log(outboxExecute)

}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
