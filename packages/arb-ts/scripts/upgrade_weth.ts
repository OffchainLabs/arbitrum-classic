/*
 * Copyright 2021, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/* eslint-env node */
'use strict'

import { AeWETH__factory } from '../src/lib/abi/factories/AeWETH__factory'
import { TransparentUpgradeableProxy__factory } from '../src/lib/abi/factories/TransparentUpgradeableProxy__factory'

import { instantiateBridge } from './instantiate_bridge'

const main = async () => {
  const { bridge, l2Network } = await instantiateBridge()
  const { l2Signer } = bridge.l2Bridge

  const aeWeth = new AeWETH__factory(l2Signer)
  const res = await aeWeth.deploy()
  await res.deployTransaction.wait()
  const logicAddress = res.address

  console.log('aeWeth logic deployed to ', logicAddress)

  const connectedProxy = TransparentUpgradeableProxy__factory.connect(
    l2Network.tokenBridge.l2Weth,
    l2Signer
  )
  const upgradeRes = await connectedProxy.upgradeTo(logicAddress)
  const upgradeRec = await upgradeRes.wait()

  console.log('successfully upgraded WETH logic', upgradeRec)
}

// const initWETH = async () => {
//   const { bridge, l2Network } = await instantiateBridge()
//   const { l2Signer } = bridge.l2Bridge

//   const aeWeth = AeWETH__factory.connect(l2Network.tokenBridge.l2Weth, l2Signer)
//   const res = await aeWeth.initialize(
//     'Wrapped Ether',
//     'WETH',
//     18,
//     l2Network.tokenBridge.l2WethGateway,
//     l2Network.tokenBridge.l1Weth
//   )

//   const rec = await res.wait()
//   console.warn('initialized weth', rec)
// }

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
