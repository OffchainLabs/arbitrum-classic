import { instantiateBridge } from './instantiate_bridge'
import { ERC20__factory } from '../src/lib/abi/factories/ERC20__factory'
import dotenv from 'dotenv'
import args from './getCLargs'
import { constants, BigNumber, utils } from 'ethers'
import { L1TokenData } from '../src'
import axios from 'axios'
import prompt from 'prompts'
dotenv.config()

const privKey = process.env.PRIVKEY as string
if (!privKey) {
  throw new Error('Set PRIVKEY env variable')
}

if (!args.l1TokenAddress) {
  throw new Error(
    'Include l1TokenAddress arg (--l1TokenAddress 0xmytokenaddress)'
  )
}
const { l1TokenAddress: l1TokenAddress } = args

const main = async () => {
  const { bridge } = await instantiateBridge(privKey, privKey)
  const walletAddress = await bridge.l1Bridge.getWalletAddress()

  /* Looks like an L1 token: */
  let l1TokenData: L1TokenData | undefined
  try {
    l1TokenData = await bridge.l1Bridge.getL1TokenData(l1TokenAddress)
  } catch (err) {
    console.warn(`${l1TokenAddress} doesn't look like an L1 ERC20 token`)
    throw err
  }

  /** Check if disabled */
  const isDisabled = await bridge.l1Bridge.tokenIsDisabled(l1TokenAddress)
  if (isDisabled) {
    console.log(`Deploying ${l1TokenAddress} is currently disabled`)
    return
  }

  /* Handle warning tokens */
  const warningTokens = (
    await axios.get(
      'https://raw.githubusercontent.com/OffchainLabs/arb-token-lists/master/src/WarningList/warningTokens.json'
    )
  ).data
  const warningToken = warningTokens[l1TokenAddress.toLowerCase()]
  if (warningToken) {
    const description = (() => {
      switch (warningToken.type) {
        case 0:
          return 'a supply rebasing token'
        case 1:
          return 'an interest accruing token'
        default:
          return 'a non-standard ERC20 token'
      }
    })()
    console.log()
    console.log(
      `${l1TokenAddress} is ${description}; it will likely have unusual behavior when deployed as as standard token to Arbitrum. It is not recommended that you deploy it. (See https://developer.offchainlabs.com/docs/bridging_assets for more info.)`
    )
    console.log()

    const res = await prompt({
      type: 'confirm',
      name: 'value',
      message: 'Are you sure you would like to proceed?',
      initial: true,
    })
    if (!res.value) {
      console.log('Good decision; terminating 👋')
      return
    }
  }

  /* check that you have some eth */
  const walletBal = await bridge.l1Provider.getBalance(walletAddress)
  if (walletBal.eq(constants.Zero)) {
    throw new Error(`${walletAddress} has no Ether to pay for gas`)
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
    amount: BigNumber.from(0),
  }
  /* check for required gas */
  const gasNeeded = await bridge.estimateGasDeposit(depositParams)
  const { maxFeePerGas, gasPrice } = await bridge.l1Provider.getFeeData()
  const price = maxFeePerGas || gasPrice
  if (!price) {
    console.log(
      'Warning: could not get gas price estimate; will try depositing anyway'
    )
  } else {
    const fee = price.mul(gasNeeded)
    if (fee.gt(walletBal)) {
      console.log(
        `An estimated ${utils.formatEther(
          fee
        )} ether is needed for deposit; you only have ${utils.formatEther(
          walletBal
        )} ether. Will try depositing anyway:`
      )
    }
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
