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

import { writeFileSync } from 'fs'
import axios from 'axios'
import Ajv from 'ajv'
import addFormats from 'ajv-formats'
import { AddressZero } from '@ethersproject/constants'
import { TokenInfo, TokenList, schema } from '@uniswap/token-lists'

import { StandardArbERC20__factory } from '../src/lib/abi/factories/StandardArbERC20__factory'
import { ERC20__factory } from '../src/lib/abi/factories/ERC20__factory'

import { instantiateBridge } from '../scripts/instantiate_bridge'

const ajv = new Ajv()
addFormats(ajv)
const validate = ajv.compile(schema)

const gen = async () => {
  const tokens: TokenInfo[] = []

  const { bridge, l2Network } = await instantiateBridge()
  const path = `./token_lists/lists/token-list-${l2Network.chainID}.json`
  const previousJSON = (
    await axios.get('https://bridge.arbitrum.io/token-list-42161.json')
  ).data

  // alt: use your  local copy:
  // const previousJSON:TokenList = JSON.parse(readFileSync(path).toString())

  const valid = validate(previousJSON)
  if (!valid) {
    console.log('Prev schema invalid')
    console.log(validate.errors)
    return
  } else {
    console.log('Previous list conforms to schema')
  }

  const previousTokens = new Set(
    previousJSON.tokens.map((token: TokenInfo) => token.address.toLowerCase())
  )

  const gatewaySetData = await bridge.getL2GatewaySetEventData()
  const excludeList = [
    '0x0CE51000d5244F1EAac0B313a792D5a5f96931BF',
    '0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f',
    '0xEDA6eFE5556e134Ef52f2F858aa1e81c84CDA84b',
    '0xe54942077Df7b8EEf8D4e6bCe2f7B58B0082b0cd',
  ]
  // flatten in case they were set more than once
  const tokenAddresses = [
    ...new Set(gatewaySetData.map(data => data.l1Token)),
  ].filter((address: string) => !excludeList.includes(address))
  const logoUris = (
    await axios.get('https://zapper.fi/api/token-list')
  ).data.tokens.reduce((acc: any, currentToken: any) => {
    return {
      ...acc,
      [currentToken.address.toLocaleLowerCase()]: currentToken.logoURI,
    }
  }, {})

  console.log(`Checking ${tokenAddresses.length} addresses`)

  let minorVersionBump = false
  let majorVersionBump = false

  for (const l1Address of tokenAddresses) {
    let l1GatewayAddress: string
    try {
      l1GatewayAddress = await bridge.l1Bridge.getGatewayAddress(l1Address)
    } catch (err) {
      console.log(`Could not get gateway for ${l1Address}; moving on!`)
      continue
    }

    if (l1GatewayAddress === AddressZero) {
      throw new Error(`Token ${l1Address} not registered in L1 router`)
    }
    if (l1Address === '0x0CE51000d5244F1EAac0B313a792D5a5f96931BF') {
      continue
    }

    const l1TokenContract = ERC20__factory.connect(
      l1Address,
      bridge.l1Bridge.l1Provider
    )

    const l2Address = await bridge.getERC20L2Address(l1Address)
    const l2AddressLowerCase = l2Address.toLowerCase()
    const code = await bridge.l2Bridge.l2Provider.getCode(l2Address)
    if (code.length <= 2)
      throw new Error(
        `${l1Address} registered at (or defaults to) ${l1GatewayAddress} but not yet deployed on L2`
      )
    const arbToken = await StandardArbERC20__factory.connect(
      l2Address,
      bridge.l2Bridge.l2Provider
    )

    const l1Name =
      l1Address === '0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2'
        ? 'Maker'
        : await l1TokenContract.name()
    const l1Symbol =
      l1Address === '0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2'
        ? 'MKR'
        : await l1TokenContract.symbol()
    const l1Decimals = await l1TokenContract.decimals()

    const l2Symbol =
      l1Address === '0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2'
        ? 'MKR'
        : await l1TokenContract.symbol()
    const l2Decimals = await arbToken.decimals()
    const l2Name =
      l1Address === '0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2'
        ? 'Maker'
        : await arbToken.name()
    if (l2Symbol !== l1Symbol) {
      console.warn(
        `******* Warning! Symbols don't match for deployment of ${l1Address}: L1 symbol ${l1Symbol}. L2 symbol ${l2Symbol}`
      )
      continue
    }
    if (l2Decimals !== l1Decimals) {
      console.warn(
        `******* Warning! decimals don't match for deployment of ${l1Address}: L1 decimals ${l1Decimals}. L2 decimals: ${l2Decimals}`
      )
      continue
    }

    if (l2Name !== l1Name) {
      console.info(
        `******* NOTE: names don't match for deployment of ${l1Address}: L1 name ${l1Name}. L2 name: ${l2Name}`
      )
    }

    let arbTokenInfo: TokenInfo = {
      chainId: +l2Network.chainID,
      address: l2Address,
      name: l2Name,
      symbol: l2Symbol,
      decimals: l2Decimals,
      extensions: {
        l1Address,
        l1GatewayAddress,
      },
    }
    const logoURI = logoUris[l1Address.toLowerCase()]
    if (logoURI) {
      arbTokenInfo = { ...{ logoURI }, ...arbTokenInfo }
    }

    console.info('Adding', l2Symbol)
    tokens.push(arbTokenInfo)

    if (!previousTokens.has(l2AddressLowerCase)) {
      console.log(`${l2Symbol} ${l2AddressLowerCase} is a new token`)
      minorVersionBump = true
    }

    previousTokens.delete(l2AddressLowerCase)
  }

  tokens.sort((a, b) => (a.symbol < b.symbol ? -1 : 1))

  if (previousTokens.size) {
    console.log(
      `${previousTokens.size} tokens removed:`,
      [...previousTokens].join(',')
    )
    majorVersionBump = true
  }

  let minorVersion = previousJSON.version.minor
  if (minorVersionBump) {
    minorVersion++
  }

  let majorVersion = previousJSON.version.major

  if (majorVersionBump) {
    majorVersion++
  }

  const tokenList: TokenList = {
    name: l2Network.name,
    timestamp: new Date().toISOString(),
    version: {
      major: majorVersion,
      minor: minorVersion,
      patch: previousJSON.version.patch,
    },
    tokens,
  }

  const validateResult = ajv.compile(schema)
  const validResult = validateResult(previousJSON)
  if (!validResult) {
    console.log('Schematic errors in new JSON, cancelling write:')
    console.log(validateResult.errors)
    return
  } else {
    console.log('new JSON is valid')
  }

  console.log(`Generating JSON with ${tokens.length} tokens`)

  const listData = JSON.stringify(tokenList)
  writeFileSync(path, listData)
}

gen()
