import tokenBalancesData from '../../json_data/42161tokenBalances.json'

import { instantiateBridge } from './../instantiate_bridge'
import { ERC20__factory } from '../../src/lib/abi/factories/ERC20__factory'
const tokenAddress = '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8'
;async () => {
  const { bridge } = await instantiateBridge()
  const token = ERC20__factory.connect(tokenAddress, bridge.l2Bridge.l2Signer)
  const tokenBalances = tokenBalancesData.balances
  const totalTargetBalance = Object.keys(tokenBalances).reduce(
    (acc, currentAddress) => {
      return acc + tokenBalances[currentAddress]
    },
    0
  )

  const address = await bridge.l2Bridge.getWalletAddress()
  const myBalance = await token.getBalance(address)
  if (myBalance.lte(totalTargetBalance)) {
    throw new Error('Not enough token for distribution')
  } else {
    console.log('Enough token to distribute; distributing now')
  }
  for (const address in tokenBalances) {
    const res = await token.transfer(address, tokenBalances[address])
    const rec = await res.wait()
    console.log(`Successfully transferred to ${address}`)
  }
  console.log('done!')
}
