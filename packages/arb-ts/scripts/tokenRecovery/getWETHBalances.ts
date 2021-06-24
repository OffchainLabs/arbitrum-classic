import { BridgeHelper } from '../../src/lib/bridge_helpers'
import { IWETH9L1__factory as IWETH9L2__factory } from '../../src/lib/abi/factories/IWETH9L1__factory'
import { instantiateBridge } from './../instantiate_bridge'
import { BigNumber } from 'ethers'
import { writeFileSync } from 'fs'

export interface WethDepositEvent {
  dst: string
}

export interface WethTransferEvent {
  src: string
  dst: string
}
export interface balancesMap {
  [address: string]: string
}
;(async () => {
  const { bridge, l2Network } = await instantiateBridge()
  const WETH_ADDRESS = l2Network.tokenBridge.l2Weth
  const WETH9 = IWETH9L2__factory.connect(
    WETH_ADDRESS,
    bridge.l2Bridge.l2Provider
  )

  const blockNumber = await bridge.l2Bridge.l2Provider.getBlockNumber()

  const deposits = (await BridgeHelper.getEventLogs('Deposit', WETH9, [])).map(
    (log: any) =>
      (WETH9.interface.parseLog(log).args as unknown) as WethDepositEvent
  )
  const transfers = (
    await BridgeHelper.getEventLogs('Transfer', WETH9, [])
  ).map(
    (log: any) =>
      (WETH9.interface.parseLog(log).args as unknown) as WethTransferEvent
  )

  console.log(
    `Found ${deposits.length} deposits and ${transfers.length} transfers`
  )
  const candidateAddresses: Set<string> = new Set([])
  for (const depositLog of deposits) {
    candidateAddresses.add(depositLog.dst)
  }
  for (const transferLog of transfers) {
    candidateAddresses.add(transferLog.dst)
    candidateAddresses.add(transferLog.src)
  }

  const balancesMap: balancesMap = {}
  let totalBalance: BigNumber = BigNumber.from(0)
  for (const address of candidateAddresses) {
    const bal = await WETH9.balanceOf(address, { blockTag: blockNumber })
    if (bal.isZero()) {
      console.log(`${address} has balance of 0`)
      continue
    }

    balancesMap[address] = bal.toString()
    totalBalance = totalBalance.add(bal)
  }
  const supply = await WETH9.totalSupply({ blockTag: blockNumber })
  if (supply.eq(totalBalance)) {
    console.log(
      `Full WETH supply ${supply} properly accounted for, generating JSON`
    )

    const listData = JSON.stringify({
      blockNumber,
      balances: balancesMap,
      address: WETH_ADDRESS,
    })
    writeFileSync(`./json_data/${l2Network.chainID}wethBalances.json`, listData)
  } else {
    throw new Error(
      `Sanity check failed: total balance counted: ${totalBalance.toString()}; total supply: ${
        supply.toString
      }`
    )
  }
})()
