import { instantiateBridge } from './instantiate_bridge'
import {
  providers,
  utils,
  Wallet,
  BigNumber,
  constants,
  ethers,
  ContractReceipt,
} from 'ethers'
import { AeWETH__factory } from '../src/lib/abi/factories/AeWETH__factory'
;(async () => {
  const pk =
    '0x302e313135333138373134363135373731393200000000000000000000000000'

  const { bridge, l2Network, l1Network } = await instantiateBridge(pk)

  const walletAddress = await bridge.l1Bridge.getWalletAddress()
  const l2w = AeWETH__factory.connect(
    l2Network.tokenBridge.l2Weth,
    bridge.l2Bridge.l2Signer
  )
  const bal = await l2w.balanceOf(walletAddress)

  console.log('l2 weth balance', bal.toString())

  const withdrawRes = await bridge.withdrawERC20(
    l1Network.tokenBridge.l1Weth,
    bal
  )
  const withdrawRec = await withdrawRes.wait()

  /*do stuff here */
})()
