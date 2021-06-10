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
  inbox: string
}

export interface Networks {
  [id: string]: Network
}

export const MAINNET_WHITELIST_ADDRESS =
  '0xD485e5c28AA4985b23f6DF13dA03caa766dcd459'

const mainnetBridge: TokenBridge = {
  l1GatewayRouter: '0x6074515Bc580C76c0C68AE2E13408B680d670F8C',
  l2GatewayRouter: '0x29f86A78551Fac44217A8763A45540027c3F7cA5',
  l1ERC20Gateway: '0xEd66239C7400f9C29D9127C5C95c18c03DDF3106',
  l2ERC20Gateway: '0x952A6d5D6757BA28dFe11Fc82d85b5880d300E58',
  inbox: '0x4Dbd4fc535Ac27206064B68FfCf827b0A60BAB3f',
}

const RinkebyBridge: TokenBridge = {
  l1GatewayRouter: '0x70C143928eCfFaf9F5b406f7f4fC28Dc43d68380',
  l2GatewayRouter: '0x9413AD42910c1eA60c737dB5f58d1C504498a3cD',
  l1ERC20Gateway: '0x91169Dbb45e6804743F94609De50D511C437572E',
  l2ERC20Gateway: '0x195C107F3F75c4C93Eba7d9a1312F19305d6375f',
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
