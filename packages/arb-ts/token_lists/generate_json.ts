import { instantiateBridge } from '../scripts/instantiate_bridge'
import uniswapDefaultList from './uniswapDefaultList'
import arbDefaultLists from './arbDefaultLists'

import yargs from 'yargs/yargs'
import { providers, utils, Wallet, BigNumber, constants, ethers } from 'ethers'
import { networks } from '../src'
import { StandardArbERC20__factory } from '../src/lib/abi/factories/StandardArbERC20__factory'
import { ERC20__factory } from '../src/lib/abi/factories/ERC20__factory'

import { TokenInfo, TokenList } from './tokenListTypes'
import { writeFileSync } from 'fs'
;(async () => {
  const { bridge, l1Network, l2Network } = await instantiateBridge()

  const tokens: TokenInfo[] = []
  const l1NetworkID = +l1Network.chainID
  const arbDefaultList = arbDefaultLists[l1NetworkID] || ([] as TokenInfo[])
  for (const l1Token of uniswapDefaultList.concat(arbDefaultList)) {
    const l1TokenContract = ERC20__factory.connect(
      l1Token.address,
      bridge.l1Bridge.l1Provider
    )
    const l1Address = l1Token.address
    let l1GatewayAddress: string

    try {
      l1GatewayAddress = await bridge.l1Bridge.getGatewayAddress(l1Address)
      if (l1GatewayAddress === constants.AddressZero) {
        /* token not registered */
        continue
      }
    } catch (err) {
      continue
    }

    const l2Address = await bridge.getERC20L2Address(l1Address)
    const code = await bridge.l2Bridge.l2Provider.getCode(l2Address)
    if (code.length <= 2) {
      console.log(
        `${l1Token.name} (${l1Token.address}) registered at (or defaults to) ${l1GatewayAddress} but not yet deployed on L2`
      )
      continue
    }
    // skip check for MKR, known unorthodox
    if (l1Token.address !== '0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2') {
      try {
        const l1Name = await l1TokenContract.name()
        const l1Symbol = await l1TokenContract.symbol()
        const l1Decimals = await l1TokenContract.decimals()

        const arbToken = await StandardArbERC20__factory.connect(
          l2Address,
          bridge.l2Bridge.l2Provider
        )
        const l2Symbol = await arbToken.symbol()
        const l2Decimals = await arbToken.decimals()
        const l2Name = await arbToken.name()
        if (l2Symbol !== l1Symbol) {
          console.warn(
            `******* Warning! Symbols don't match for standard deployment of ${l1Token.address}: L1 symbol ${l1Symbol}. L2 symbol ${l2Symbol}`
          )
          continue
        }
        if (l2Decimals !== l1Decimals) {
          console.warn(
            `******* Warning! decimals don't match for standard deployment of ${l1Token.address}: L1 decimals ${l1Decimals}. L2 decimals: ${l2Decimals}`
          )
          continue
        }

        if (l2Name !== l1Name) {
          console.warn(
            `******* Warning! names don't match for standard deployment of ${l1Token.address}: L1 name ${l1Name}. L2 name: ${l2Name}`
          )
          continue
        }
      } catch (err) {
        console.warn(
          `**** Warning! Error verifying standard deployment of ${l1Token.address}`,
          err
        )
        continue
      }
    }
    if (l1GatewayAddress === l1Network.tokenBridge.l1ERC20Gateway) {
      /* is registered as standard token
      check that it looks right */

      const arbTokenInfo: TokenInfo = {
        chainId: +l2Network.chainID,
        address: l2Address,
        name: l1Token.name,
        symbol: l1Token.symbol,
        decimals: l1Token.decimals,
        logoURI: l1Token.logoURI,
        extensions: {
          l1Address,
          l1GatewayAddress,
        },
      }
      tokens.push(arbTokenInfo)
    } else if (l1GatewayAddress === l1Network.tokenBridge.l1CustomGateway) {
      /* is registered as custom token */

      const arbTokenInfo: TokenInfo = {
        chainId: +l2Network.chainID,
        address: l2Address,
        name: l1Token.name,
        symbol: l1Token.symbol,
        decimals: l1Token.decimals,
        logoURI: l1Token.logoURI,
        extensions: {
          l1Address,
          l1GatewayAddress,
        },
      }
      tokens.push(arbTokenInfo)
    } else {
      console.warn(
        `Warning: ${l1Token.name} (${l1Token.address}) registered at unrecognized gateway: ${l1GatewayAddress}`
      )
    }
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
})()
