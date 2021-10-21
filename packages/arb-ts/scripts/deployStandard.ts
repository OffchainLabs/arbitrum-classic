import { instantiateBridge } from './instantiate_bridge'
import { ERC20__factory } from '../src/lib/abi/factories/ERC20__factory'
import dotenv from 'dotenv'
import args from './getCLargs'
import { constants, BigNumber, utils } from 'ethers'
import { L1TokenData } from '../src'

dotenv.config()

const privKey = process.env.PRIVKEY as string
if (!privKey) {
  throw new Error('Set PRIVKEY env variable')
}

if (!args.l1_token_address) {
  throw new Error(
    'Include l1_token_address arg (--l1_token_address 0xmytokenaddress)'
  )
}
if (!args.amount) {
  throw new Error('Include amount arg (--amount 123)')
}

const { l1_token_address: l1TokenAddress, amount } = args
const amountBigNum = BigNumber.from(amount)

const main = async () => {
  const { bridge } = await instantiateBridge(privKey, privKey)
  const walletAddress = await bridge.l1Bridge.getWalletAddress()

  const l1Token = ERC20__factory.connect(l1TokenAddress, bridge.l1Provider)
  // TODO: check blacklist

  /* Looks like an L1 token: */
  let l1TokenData: L1TokenData | undefined
  try {
    l1TokenData = await bridge.l1Bridge.getL1TokenData(l1TokenAddress)
  } catch (err) {
    console.warn(`${l1TokenAddress} doesn't look like an L1 ERC20 token`)
    throw err
  }
  /* check that you have some eth */
  const walletBal = await bridge.l1Provider.getBalance(walletAddress)
  if (walletBal.eq(constants.Zero)) {
    throw new Error(`${walletAddress} has no Ether to pay for gas`)
  }
  /* check token bal */
  if (l1TokenData.balance.lt(amountBigNum)) {
    throw new Error(
      `Insufficient token balance for deposit; you tried depositing ${amount} but you only have ${l1TokenData.balance.toString()}`
    )
  }

  /* check token not yet deployed */
  const l2TokenAddress = await bridge.l1Bridge.getERC20L2Address(l1TokenAddress)
  if (l2TokenAddress === constants.AddressZero) {
    throw new Error(`${l1TokenAddress} can't be bridged`)
  }

  if ((await bridge.l2Provider.getCode(l2TokenAddress)).length > 2) {
    throw new Error(
      `${l1TokenAddress} already deployed on L2 at ${l2TokenAddress}`
    )
  }

  /* set allowance */
  if (!l1TokenData.allowed) {
    console.log('Setting allowance on gateway')

    const res = await bridge.approveToken(l1TokenAddress)
    const rec = await res.wait(2)
    console.log(
      `Allowance successfully set — L1 tx hash: ${rec.transactionHash}`
    )
  }
  const depositParams = {
    erc20L1Address: l1TokenAddress,
    amount: amountBigNum,
  }
  /* check for required gas */
  const gasNeeded = await bridge.estimateGasDeposit(depositParams)

  if (gasNeeded.gt(walletBal)) {
    throw new Error(
      `An estimated ${utils.formatEther(
        gasNeeded
      )} ether is needed for deposit; you only have ${utils.formatEther(
        walletBal
      )} ether`
    )
  }
  console.log('Depositing / deploying standard token contract:')

  const res = await bridge.deposit(depositParams)
  const rec = await res.wait(2)
  console.log(`L1 deposit txn confirmed — L1 txn hash: ${rec.transactionHash}`)
  const seqNums = await bridge.getInboxSeqNumFromContractTransaction(rec)
  if (!seqNums) {
    throw new Error(
      `Sequence number not found for ${rec.transactionHash} (???)`
    )
  }
  const seqNum = seqNums[0]

  const userTxnHash = await bridge.calculateL2RetryableTransactionHash(seqNum)
  console.log(
    `Waiting for L2 txn; this takes ~10 minutes; waiting for L2 txn hash: ${userTxnHash}`
  )

  await bridge.waitForRetryableReceipt(seqNum, 2)
  console.log(`Done; your token is deployed on L2 at ${l2TokenAddress}`)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
