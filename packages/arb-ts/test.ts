import { Bridge } from './src/b'
import { ethers } from 'ethers'
const l1Prov = new ethers.providers.JsonRpcProvider(
  'https://kovan.infura.io/v3/c13a0d6955b14bf181c924bf4c7797fc'
)
const l2Prov = new ethers.providers.JsonRpcProvider(
  'https://kovan5.arbitrum.io/rpc'
)
const signer = new ethers.Wallet(
  '0xffb2b26161e081f0cdf9db67200ee0ce25499d5ee683180a9781e6cceb791c39'
)
const l2Signer = signer.connect(l2Prov)
const bridge = new Bridge(
  '0x1d750369c91b129524B68f308512b0FE2C903d71',
  '0x2EEBB8EE9c377caBC476654ca4aba016ECA1B9fc',
  signer.connect(l1Prov),
  l2Signer
)
const wait = (ms: number) => new Promise(res => setTimeout(res, ms))
const main = async () => {
  console.log('startin')
  const res = await bridge.l1Bridge.ethERC20Bridge.queryFilter(
    bridge.l1Bridge.ethERC20Bridge.filters.DepositToken(
      null,
      null,
      null,
      null,
      null
    ),
    // "0x017C70B5",
    // "0x017C70B6",
    '0x017C7A3B',
    'latest'
  )

  if (res.length === 0) {
    console.log('No deposits')
  }
  console.log(`investigating ${res.length} txs`)
  const txs = res.map(curr => curr.transactionHash)
  for (const tx of txs) {
    console.log(`Looking into L1 tx ${tx}`)
    const l1Tx = await bridge.getL1Transaction(tx)
    const _seq = await bridge.getInboxSeqNumFromContractTransaction(l1Tx)
    if (!_seq) throw new Error('aa')
    const seq = _seq[0]
    const l2TxHash = await bridge.calculateRetryableAutoReedemTxnHash(seq)
    try {
      const l2Receipt = await bridge.getL2Transaction(l2TxHash)
      console.log('Already redeemed!')
    } catch (e) {
      console.log('no receipt')
      try {
        const tx = await l2Signer.sendTransaction({
          to: '0x000000000000000000000000000000000000006E',
          // data: "e64017b4306f5b28fc0878eeb57a4c441ec83bbe13c4a4c34f13abe1081b57f4"
          data: `0xeda1122c${l2TxHash.substr(2)}`,
        })
        console.log(`wait for receipt now ${tx.hash}`)
        const receipt = await tx.wait()
        console.log('Redeem success :)')
      } catch (e) {
        const l2TxFoo = await bridge.calculateL2TransactionHash(seq)
        console.log(
          `failed on redeem autoredeem ${l2TxHash} also check ${l2TxFoo}`
        )
      }
    }
  }
}
const exec = async (runs = 1) => {
  await main()
  console.log(`ran ${runs} times, waiting`)
  await wait(30000)
  console.log('running again')
  await exec(runs + 1)
}
exec()
  .then(() => console.log('done'))
  .catch(err => console.log(err))
