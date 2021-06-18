import { instantiateBridge } from '../scripts/instantiate_bridge'
import { providers, utils, Wallet, BigNumber, constants, ethers } from 'ethers'
import { StandardArbERC20__factory } from '../src/lib/abi/factories/StandardArbERC20__factory'
import { ERC20__factory } from '../src/lib/abi/factories/ERC20__factory'

import { TokenInfo, TokenList } from './tokenListTypes'
import { writeFileSync } from 'fs'
import axios from 'axios'

const gen = async () => {
  const tokens: TokenInfo[] = []

  const { bridge, l1Network, l2Network } = await instantiateBridge()
  const gatewaySetData = await bridge.getL2GatewaySetEventData()
  // flatten in case they were set more than once
  const tokenAddresses = [...new Set(gatewaySetData.map(data => data.l1Token))]
  const logoUris = (
    await axios.get('https://zapper.fi/api/token-list')
  ).data.tokens.reduce((acc: any, currentToken: any) => {
    return {
      ...acc,
      [currentToken.address.toLocaleLowerCase()]: currentToken.logoURI,
    }
  }, {})

  console.log(`Checking ${tokenAddresses.length} addresses`)

  for (const l1Address of tokenAddresses) {
    const l1GatewayAddress = await bridge.l1Bridge.getGatewayAddress(l1Address)
    if (l1GatewayAddress === constants.AddressZero) {
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
      console.warn(
        `******* Warning! names don't match for deployment of ${l1Address}: L1 name ${l1Name}. L2 name: ${l2Name}`
      )
      continue
    }

    const arbTokenInfo: TokenInfo = {
      chainId: +l2Network.chainID,
      address: l2Address,
      name: l2Name,
      symbol: l2Symbol,
      decimals: l2Decimals,
      logoURI: logoUris[l1Address.toLocaleLowerCase()] || '',
      extensions: {
        l1Address,
        l1GatewayAddress,
      },
    }
    tokens.push(arbTokenInfo)
  }
  tokens.sort((a, b) => (a.symbol < b.symbol ? -1 : 1))
  const tokenList: TokenList = {
    name: l2Network.name,
    timestamp: new Date().toISOString(),
    version: {
      major: 0,
      minor: 1,
      patch: 0,
    },
    tokens,
  }
  console.log(`Generating JSON with ${tokens.length} tokens`)

  const listData = JSON.stringify(tokenList)
  writeFileSync(
    `./token_lists/lists/token-list-${l2Network.chainID}.json`,
    listData
  )
}

gen()
