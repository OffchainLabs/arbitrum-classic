import { instantiateBridge } from './instantiate_bridge'
import { BigNumber } from '@ethersproject/bignumber'
;(async () => {
  const { bridge } = await instantiateBridge()
  for (let i = 0; i < 10; i++) {
    console.log('Starting round', i)
    const promises = []
    for (let j = 0; j < 200; j++) {
      promises.push(
        bridge.getDepositTxParams({
          erc20L1Address: '0xb6ed7644c69416d67b522e20bc294a9a9b405b31',
          amount: BigNumber.from(0),
          retryableGasArgs: {},
          destinationAddress: '0xb6ed7644c69416d67b522e20bc294a9a9b405b31',
        })
      )
    }
    await Promise.all(promises)
  }
})()
