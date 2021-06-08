type Network = {
  [string: number]: {
    l1: {
      chainId: number
      l1GatewayRouter: string
      l1ERC20GatewayProxy: string
      inbox: string
    }
    l2: {
      chainId: number
      l2GatewayRouter: string
      l2ERC20GatewayProxy: string
    }
  }
}

const networks: Network = {
  // rinkeby
  421611: {
    l1: {
      chainId: 4,
      l1GatewayRouter: '0x70C143928eCfFaf9F5b406f7f4fC28Dc43d68380',
      l1ERC20GatewayProxy: '0x91169Dbb45e6804743F94609De50D511C437572E',
      inbox: '0x578BAde599406A8fE3d24Fd7f7211c0911F5B29e',
    },
    l2: {
      chainId: 421611,
      l2GatewayRouter: '0x9413AD42910c1eA60c737dB5f58d1C504498a3cD',
      l2ERC20GatewayProxy: '0x195C107F3F75c4C93Eba7d9a1312F19305d6375f',
    },
  },
}

export default networks
