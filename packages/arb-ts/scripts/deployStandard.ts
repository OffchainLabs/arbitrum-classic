import { instantiateBridge } from './instantiate_bridge'
import dotenv from 'dotenv'
import args from './getCLargs'
import { constants, BigNumber, utils } from 'ethers'
import { MultiCaller } from '../src'
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
  const { l1Signer, tokenBridger, l2Signer } = await instantiateBridge(
    privKey,
    privKey
  )
  const l1Provider = l1Signer.provider!
  const l2Provider = l2Signer.provider!
  const gatewayAddress = await tokenBridger.getL1GatewayAddress(
    l1TokenAddress,
    l1Provider
  )
  const l1SignerAddr = await l1Signer.getAddress()

  /* Looks like an L1 token: */
  const multicaller = await MultiCaller.fromProvider(l1Provider)
  let l1TokenData: { allowance: BigNumber | undefined }
  try {
    l1TokenData = (
      await multicaller.getTokenData([l1TokenAddress], {
        allowance: {
          owner: l1SignerAddr,
          spender: gatewayAddress,
        },
      })
    )[0]
  } catch (err) {
    console.warn(`${l1TokenAddress} doesn't look like an L1 ERC20 token`)
    throw err
  }

  /** Check if disabled */
  const isDisabled = await tokenBridger.l1TokenIsDisabled(
    l1TokenAddress,
    l1Provider
  )
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
  const walletBal = await l1Provider.getBalance(l1SignerAddr)
  if (walletBal.eq(constants.Zero)) {
    throw new Error(`${l1SignerAddr} has no Ether to pay for gas`)
  }

  /* check token not yet deployed */
  const l2TokenAddress = await tokenBridger.getL2ERC20Address(
    l1TokenAddress,
    l1Provider
  )
  if (l2TokenAddress === constants.AddressZero) {
    throw new Error(`${l1TokenAddress} can't be bridged`)
  }

  if ((await l2Provider.getCode(l2TokenAddress)).length > 2) {
    throw new Error(
      `${l1TokenAddress} already deployed on L2 at ${l2TokenAddress}`
    )
  }

  /* set allowance */
  const amount = BigNumber.from(0)
  const approveAmount = BigNumber.from(1000)
  if (l1TokenData.allowance.lt(approveAmount)) {
    console.log('Setting allowance on gateway')

    const res = await tokenBridger.approveToken({
      erc20L1Address: l1TokenAddress,
      l1Signer: l1Signer,
      amount,
    })
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
  const gasNeeded = await tokenBridger.depositEstimateGas({
    amount,
    erc20L1Address: l1TokenAddress,
    l1Signer,
    l2Provider: l2Provider,
  })
  const price = await l2Provider.getGasPrice()
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

  const res = await tokenBridger.deposit({
    amount,
    erc20L1Address: l1TokenAddress,
    l1Signer,
    l2Provider,
  })
  const rec = await res.wait(2)
  console.log(`L1 deposit txn confirmed — L1 txn hash: ${rec.transactionHash}`)
  const message = await rec.getL1ToL2Message(l2Provider)

  await message.wait(undefined, 2)
  console.log(`Done; your token is deployed on L2 at ${l2TokenAddress}`)
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
