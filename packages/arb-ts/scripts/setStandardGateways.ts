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

import { setStandardGateWays } from './lib'
import args from './getCLargs'

if (!(args.address || args.addresses)) {
  throw new Error(
    'Include token address(es) (--address 0xmyaddress) or   (--addresses 0xmyaddress,0xmyotheraddress,0xmythirdaddress)'
  )
}
const tokensAddresses: string[] = ((args.addresses || args.address) as string)
  .split(',')
  .map((address: string) => address.trim())

if (tokensAddresses.length === 0) {
  throw new Error('Include some tokens to set')
}
tokensAddresses.forEach((address: string) => {
  if (!(address.startsWith('0x') && address.length === 42)) {
    throw new Error(address + " doesn't look like a token address")
  }
})

console.log('Setting tokens to standard gateway:', tokensAddresses)

setStandardGateWays(tokensAddresses).then(() => {
  console.log('done')
})
