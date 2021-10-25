import { JsonRpcProvider } from '@ethersproject/providers'
import { NODE_INTERFACE_ADDRESS } from '../src/lib/precompile_addresses'
import { NodeInterface__factory } from '../src/lib/abi/factories/NodeInterface__factory'
;(async () => {
  const rpcURL = process.env['DEV_RPC']
  const arbProvider = new JsonRpcProvider(rpcURL)
  const nodeInterface = NodeInterface__factory.connect(
    NODE_INTERFACE_ADDRESS,
    arbProvider
  )

  for (let i = 0; i < 10; i++) {
    console.log('Starting round', i)
    const promises = []
    for (let j = 0; j < 200; j++) {
      promises.push(
        nodeInterface.estimateRetryableTicket(
          '0xb6ed7644c69416d67b522e20bc294a9a9b405b31',
          0,
          '0xb6ed7644c69416d67b522e20bc294a9a9b405b31',
          0,
          0,
          '0xb6ed7644c69416d67b522e20bc294a9a9b405b31',
          '0xb6ed7644c69416d67b522e20bc294a9a9b405b31',
          0,
          0,
          '0x'
        )
      )
    }
    try {
      await Promise.all(promises)
    } catch (e) {
      // do nothing
    }
  }
})()
