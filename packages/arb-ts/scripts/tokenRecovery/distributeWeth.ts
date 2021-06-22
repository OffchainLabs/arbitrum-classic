import wethBalances from '../../wethBalances.json'
import { instantiateBridge } from './../instantiate_bridge'
import { BigNumber } from 'ethers'

import { IWETH9L2__factory } from '../../src/lib/abi/factories/IWETH9L2__factory'
const wethAddress = '0x82aF49447D8a07e3bd95BD0d56f35241523fBab1'

;async () => {
  const { bridge } = await instantiateBridge()
  const WETH9 = IWETH9L2__factory.connect(wethAddress, bridge.l2Bridge.l2Signer)
  console.log('Wrapping some weth:')

  const res = await WETH9.deposit({
    value: 200000000000000000,
  })

  const rec = await res.wait()
  console.log('Done wrapping eth', rec)

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
