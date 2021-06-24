import { instantiateBridge } from './../instantiate_bridge'
import { BigNumber } from 'ethers'

import { IWETH9L1__factory as IWETH9L2__factory } from '../../src/lib/abi/factories/IWETH9L1__factory'
import wethBalancesMainnetData from '../../json_data/42161wethBalances.json'
import wethBalancesRinkArbyData from '../../json_data/421611wethBalances.json'
;async () => {
  const { bridge, l2Network } = await instantiateBridge()

  const WETH9 = IWETH9L2__factory.connect(
    l2Network.tokenBridge.l2Weth,
    bridge.l2Bridge.l2Signer
  )

  const res = await WETH9.deposit({
    value: 200000000000000000,
  })

  const rec = await res.wait()
  const data =
    l2Network.chainID === '42161'
      ? wethBalancesMainnetData
      : wethBalancesRinkArbyData
  const wethBalances = data.balances
  const totalTargetBalance = Object.keys(wethBalances).reduce(
    (acc, currentAddress) => {
      return acc + wethBalances[currentAddress]
    },
    0
  )
  const address = await bridge.l2Bridge.getWalletAddress()
  const myBalance = BigNumber.from(await WETH9.balanceOf(address))
  if (myBalance.lte(totalTargetBalance)) {
    throw new Error('Not enough WETH for distribution')
  } else {
    console.log('Enough WETH to distribute; distributing now')
  }

  for (const address in wethBalances) {
    const res = await WETH9.transfer(address, wethBalances[address])
    const rec = await res.wait()
    console.log(`Successfully transferred to ${address}`)
  }
  console.log('done!')
}
