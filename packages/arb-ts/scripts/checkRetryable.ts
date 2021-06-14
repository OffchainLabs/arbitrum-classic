import { Bridge, BridgeHelper } from '../src'
import { providers, Wallet } from 'ethers'
import { ArbRetryableTx__factory } from 'arb-ts/src/lib/abi/factories/ArbRetryableTx__factory'
// import { L2GatewayRouter__factory } from 'arb-ts/src/lib/abi/factories/L2GatewayRouter__factory'

const l1Hash =
  '0x9945b0773f0da582f729507b018c25c0e2f34baa46663926e497d07a34eb9b37'

const rinkInbox = '0x578BAde599406A8fE3d24Fd7f7211c0911F5B29e'
const mainnetInbox = '0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f'
const l2Provider = new providers.JsonRpcProvider(
  'https://rinkeby.arbitrum.io/rpc'
  // https://arb1.arbitrum.io/rpc
)

const l2Singer = new Wallet(process.env['DEVNET_PRIVKEY'] as string, l2Provider)

const l1Provider = new providers.JsonRpcProvider(
  // 'https://mainnet.infura.io/v3/ + process.env['INFURA_KEY']
  'https://rinkeby.infura.io/v3/' + process.env['INFURA_KEY']
)
const check = async () => {
  const rec = await l1Provider.getTransactionReceipt(l1Hash)
  const _seqNums = await BridgeHelper.getInboxSeqNumFromContractTransaction(
    rec,
    rinkInbox
  )
  if (!_seqNums) throw new Error('no seq nums')
  const seqNum = _seqNums[0]

  const autoHash = await BridgeHelper.calculateRetryableAutoRedeemTxnHash(
    seqNum,
    l2Provider
  )
  const autoRec = await l2Provider.getTransactionReceipt(autoHash)

  const retryHash = await BridgeHelper.calculateL2RetryableTransactionHash(
    seqNum,
    l2Provider
  )
  const retryRec = await l2Provider.getTransactionReceipt(retryHash)

  const reqId = await BridgeHelper.calculateL2TransactionHash(
    seqNum,
    l2Provider
  )

  const reqIdRec = await l2Provider.waitForTransaction(reqId)

  console.log('*** autoRedeemHash', autoHash)
  console.log('*** autoRedeem status', autoRec ? autoRec.status : autoRec)
  if (autoRec && autoRec.status !== 1) {
    console.log('**** autoredeem receipt', autoRec)
  }

  console.log('*** retryableHash', retryHash)
  console.log('*** retryable status', retryRec ? retryRec.status : retryRec)
  if (retryRec && retryRec.status !== 1) {
    console.log('**** retryable receipt', autoRec)
  }

  console.log('*** reqIdHash', reqId)
  console.log('*** reqId status', reqIdRec ? reqIdRec.status : reqIdRec)
  if (reqIdRec && reqIdRec.status !== 1) {
    console.log('**** reqId receipt', autoRec)
  }
}

check()
