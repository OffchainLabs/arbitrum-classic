export interface Network {
  chainID: string
  name: string
  isArbitrum: boolean
  explorerUrl: string
  partnerChainID: string
  tokenBridge: TokenBridge
  gif?: string
  confirmPeriodBlocks?: number
  blockTime?: number //seconds
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
  inbox: string
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
  l1CustomGateway: '',
  l2CustomGateway: '',
  l1WethGateway: '',
  l2WethGateway: '',
  inbox: '0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f',
}

const RinkebyBridge: TokenBridge = {
  l1GatewayRouter: '0x70C143928eCfFaf9F5b406f7f4fC28Dc43d68380',
  l2GatewayRouter: '0x9413AD42910c1eA60c737dB5f58d1C504498a3cD',
  l1ERC20Gateway: '0x91169Dbb45e6804743F94609De50D511C437572E',
  l2ERC20Gateway: '0x195C107F3F75c4C93Eba7d9a1312F19305d6375f',
  l1CustomGateway: '0x917dc9a69F65dC3082D518192cd3725E1Fa96cA2',
  l2CustomGateway: '0x9b014455AcC2Fe90c52803849d0002aeEC184a06',
  l1WethGateway: '0xf94bc045c4E926CC0b34e8D1c41Cd7a043304ac9',
  l2WethGateway: '0xF90EB31045d5b924900afF29344dEb42EAe0b087',
  inbox: '0x578BAde599406A8fE3d24Fd7f7211c0911F5B29e',
}

const networks: Networks = {
  '1': {
    chainID: '1',
    name: 'Mainnet',
    explorerUrl: 'https://etherscan.io',
    isArbitrum: false,
    partnerChainID: '42161',
    tokenBridge: mainnetBridge,
    blockTime: 15,
  },
  '42161': {
    chainID: '42161',
    name: 'Arb1',
    explorerUrl: 'https://mainnet-arb-explorer.netlify.app',
    partnerChainID: '1',
    isArbitrum: true,
    tokenBridge: mainnetBridge,
    confirmPeriodBlocks: 45818,
  },
  '4': {
    chainID: '4',
    name: 'Rinkeby',
    explorerUrl: 'https://rinkeby.etherscan.io',
    partnerChainID: '421611',
    isArbitrum: false,
    tokenBridge: RinkebyBridge,
    confirmPeriodBlocks: 6545, // TODO
    blockTime: 15,
  },
  '421611': {
    chainID: '421611',
    name: 'ArbRinkeby',
    explorerUrl: 'https://rinkeby-explorer.arbitrum.io',
    partnerChainID: '4',
    isArbitrum: true,
    tokenBridge: RinkebyBridge,
    confirmPeriodBlocks: 6545, // TODO
  },
}

export default networks
