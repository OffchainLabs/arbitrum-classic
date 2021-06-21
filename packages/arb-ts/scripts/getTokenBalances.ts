import { BridgeHelper } from '../src/lib/bridge_helpers'
import { ERC20__factory } from '../src/lib/abi/factories/ERC20__factory'
import { instantiateBridge } from './instantiate_bridge'
import { BigNumber } from 'ethers'
import { writeFileSync } from 'fs'

const tokenAddress = '0xFF970A61A04b1cA14834A43f5dE4533eBDDB5CC8'

export interface TransferEvent {
  from: string
  to: string
}
export interface balancesMap {
  [address: string]: string
}
;(async () => {
  const { bridge } = await instantiateBridge()
  const token = ERC20__factory.connect(tokenAddress, bridge.l2Bridge.l2Provider)
  const candidateAddresses: Set<string> = new Set([])

  const transfers = (
    await BridgeHelper.getEventLogs('Transfer', token, [])
  ).map(
    (log: any) =>
      (token.interface.parseLog(log).args as unknown) as TransferEvent
  )

  for (const transferLog of transfers) {
    candidateAddresses.add(transferLog.from)
    candidateAddresses.add(transferLog.to)
  }

  const balancesMap: balancesMap = {}
  let totalBalance: BigNumber = BigNumber.from(0)
  for (const address of candidateAddresses) {
    const bal = await token.balanceOf(address)
    if (bal.isZero()) {
      console.log(`${address} has balance of 0`)
      continue
    }

    balancesMap[address] = bal.toString()
    totalBalance = totalBalance.add(bal)
  }
  const supply = await token.totalSupply()
  if (supply.eq(totalBalance)) {
    console.log('Full token supply properly accounted for, generating JSON')

    const listData = JSON.stringify(balancesMap)
    writeFileSync(`./tokenBalances.json`, listData)
  } else {
    throw new Error(
      `Sanity check failed: total balance counted: ${totalBalance.toString()}; total supply: ${
        supply.toString
      }`
    )
  }
})()
