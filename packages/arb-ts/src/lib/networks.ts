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

import dotenv from 'dotenv'
dotenv.config()

export interface Network {
  chainID: string
  name: string
  isArbitrum: boolean
  explorerUrl: string
  partnerChainID: string
  tokenBridge: TokenBridge
  ethBridge: EthBridge | undefined
  gif?: string
  confirmPeriodBlocks?: number
  blockTime?: number //seconds
  rpcURL: string
}

export interface TokenBridge {
  l1GatewayRouter: string
  l2GatewayRouter: string
  l1ERC20Gateway: string
  l2ERC20Gateway: string
  l1CustomGateway: string
  l2CustomGateway: string
  l1WethGateway: string
  l2WethGateway: string
  l2Weth: string
  l1Weth: string
  l1ProxyAdmin: string
  l2ProxyAdmin: string
  l1DaiGateway: string
  l2DaiGateway: string
}

export interface EthBridge {
  inbox: string
  sequencerInbox: string
  outbox: string
  rollup: string
}

export interface Networks {
  [id: string]: Network
}

export const MAINNET_WHITELIST_ADDRESS =
  '0xD485e5c28AA4985b23f6DF13dA03caa766dcd459'

const mainnetBridge: TokenBridge = {
  l1GatewayRouter: '0x72Ce9c846789fdB6fC1f34aC4AD25Dd9ef7031ef',
  l2GatewayRouter: '0x5288c571Fd7aD117beA99bF60FE0846C4E84F933',
  l1ERC20Gateway: '0xa3A7B6F88361F48403514059F1F16C8E78d60EeC',
  l2ERC20Gateway: '0x09e9222E96E7B4AE2a407B98d48e330053351EEe',
  l1CustomGateway: '0xcEe284F754E854890e311e3280b767F80797180d',
  l2CustomGateway: '0x096760F208390250649E3e8763348E783AEF5562',
  l1WethGateway: '0xd92023E9d9911199a6711321D1277285e6d4e2db',
  l2WethGateway: '0x6c411aD3E74De3E7Bd422b94A27770f5B86C623B',
  l2Weth: '0x82aF49447D8a07e3bd95BD0d56f35241523fBab1',
  l1Weth: '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2',
  l1ProxyAdmin: '0x9aD46fac0Cf7f790E5be05A0F15223935A0c0aDa',
  l2ProxyAdmin: '0xd570aCE65C43af47101fC6250FD6fC63D1c22a86',
  l1DaiGateway: '0xD3B5b60020504bc3489D6949d545893982BA3011',
  l2DaiGateway: '0x467194771dAe2967Aef3ECbEDD3Bf9a310C76C65',
}

const rinkebyBridge: TokenBridge = {
  l1GatewayRouter: '0x70C143928eCfFaf9F5b406f7f4fC28Dc43d68380',
  l2GatewayRouter: '0x9413AD42910c1eA60c737dB5f58d1C504498a3cD',
  l1ERC20Gateway: '0x91169Dbb45e6804743F94609De50D511C437572E',
  l2ERC20Gateway: '0x195C107F3F75c4C93Eba7d9a1312F19305d6375f',
  l1CustomGateway: '0x917dc9a69F65dC3082D518192cd3725E1Fa96cA2',
  l2CustomGateway: '0x9b014455AcC2Fe90c52803849d0002aeEC184a06',
  l1WethGateway: '0x81d1a19cf7071732D4313c75dE8DD5b8CF697eFD',
  l2WethGateway: '0xf94bc045c4E926CC0b34e8D1c41Cd7a043304ac9',
  l2Weth: '0xB47e6A5f8b33b3F17603C83a0535A9dcD7E32681',
  l1Weth: '0xc778417E063141139Fce010982780140Aa0cD5Ab',
  l1ProxyAdmin: '0x0DbAF24efA2bc9Dd1a6c0530DD252BCcF883B89A',
  l2ProxyAdmin: '0x58816566EB91815Cc07f3Ad5230eE0820fe1A19a',
  l1DaiGateway: '0x10E6593CDda8c58a1d0f14C5164B376352a55f2F',
  l2DaiGateway: '0x467194771dAe2967Aef3ECbEDD3Bf9a310C76C65',
}

const rinkebyETHBridge: EthBridge = {
  inbox: '0x578BAde599406A8fE3d24Fd7f7211c0911F5B29e',
  sequencerInbox: '0xe1ae39e91c5505f7f0ffc9e2bbf1f6e1122dcfa8',
  outbox: '0x2360A33905dc1c72b12d975d975F42BaBdcef9F3',
  rollup: '0xFe2c86CF40F89Fe2F726cFBBACEBae631300b50c',
}

const mainnetETHBridge: EthBridge = {
  inbox: '0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f',
  sequencerInbox: '0x4c6f947Ae67F572afa4ae0730947DE7C874F95Ef',
  outbox: '0x760723CD2e632826c38Fef8CD438A4CC7E7E1A40',
  rollup: '0xC12BA48c781F6e392B49Db2E25Cd0c28cD77531A',
}

export const networks: Networks = {
  '1': {
    chainID: '1',
    name: 'Mainnet',
    explorerUrl: 'https://etherscan.io',
    isArbitrum: false,
    partnerChainID: '42161',
    tokenBridge: mainnetBridge,
    ethBridge: mainnetETHBridge,
    blockTime: 15,
    rpcURL: process.env['MAINNET_RPC'] as string,
  },
  '1337': {
    chainID: '1337',
    name: 'Hardhat_Mainnet_Fork',
    explorerUrl: 'https://etherscan.io',
    isArbitrum: false,
    partnerChainID: '42161', // TODO: use sequencer fork ID
    tokenBridge: mainnetBridge,
    ethBridge: mainnetETHBridge,
    blockTime: 15,
    rpcURL: process.env['HARDHAT_RPC'] || 'http://127.0.0.1:8545/',
  },
  '42161': {
    chainID: '42161',
    name: 'Arbitrum One',
    explorerUrl: 'https://mainnet-arb-explorer.netlify.app',
    partnerChainID: '1',
    isArbitrum: true,
    tokenBridge: mainnetBridge,
    ethBridge: undefined,
    confirmPeriodBlocks: 45818,
    rpcURL: process.env['ARB_ONE_RPC'] || 'https://arb1.arbitrum.io/rpc',
  },
  '4': {
    chainID: '4',
    name: 'Rinkeby',
    explorerUrl: 'https://rinkeby.etherscan.io',
    partnerChainID: '421611',
    isArbitrum: false,
    tokenBridge: rinkebyBridge,
    ethBridge: rinkebyETHBridge,
    confirmPeriodBlocks: 6545, // TODO
    blockTime: 15,
    rpcURL: process.env['RINKEBY_RPC'] as string,
  },
  '421611': {
    chainID: '421611',
    name: 'ArbRinkeby',
    explorerUrl: 'https://rinkeby-explorer.arbitrum.io',
    partnerChainID: '4',
    isArbitrum: true,
    tokenBridge: rinkebyBridge,
    ethBridge: undefined,
    confirmPeriodBlocks: 6545, // TODO
    rpcURL: process.env['RINKARBY_RPC'] || 'https://rinkeby.arbitrum.io/rpc',
  },
}

export default networks
