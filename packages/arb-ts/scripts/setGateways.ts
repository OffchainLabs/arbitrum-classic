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

import { instantiateBridge } from './instantiate_bridge'
import { validateToken, validateGateway } from './lib'
import args from './getCLargs'

async function main() {
  if (!args.tokens) {
    throw new Error(
      'Include token addresses (--tokens 0xmyaddress,0xmyaddress)'
    )
  }
  if (!args.gateways) {
    throw new Error(
      'Include gateway addresses (--gateways 0xmyaddress,0xmyaddress)'
    )
  }

  const tokensAddresses: string[] = args.tokens
    .split(',')
    .map((address: string) => address.trim())

  const gatewayAddresses: string[] = args.gateways
    .split(',')
    .map((address: string) => address.trim())

  if (tokensAddresses.length != gatewayAddresses.length) {
    throw new Error('Must have same number of gateways and tokens')
  }

  const { bridge } = await instantiateBridge()
  for (let i = 0; i < tokensAddresses.length; i++) {
    const token = tokensAddresses[i]
    const gateway = gatewayAddresses[i]
    await validateToken(bridge.l1Bridge.l1Provider, token)
    await validateGateway(bridge.l1Bridge.l1Provider, token, gateway)
  }
  console.log(`Registering ${tokensAddresses.length} tokens with gateways`)
  const tx = await bridge.setGateways(tokensAddresses, gatewayAddresses)
  await tx.wait()
  console.log('Done')
}

main()
