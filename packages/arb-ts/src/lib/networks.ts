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
      l1GatewayRouter: '0x753A61f15212327931FED5B089315302A6903E74',
      l1ERC20GatewayProxy: '0x3aa4c5C6D0Ec415E17ca2B75383208D459056F90',
      inbox: '0x578BAde599406A8fE3d24Fd7f7211c0911F5B29e',
    },
    l2: {
      chainId: 421611,
      l2GatewayRouter: '0xA35Ee00f10fF3195085dc1b0F8eF9C8843d67540',
      l2ERC20GatewayProxy: '0x12AcE03143cdD43659C38699C07d2E3FBb8EcC60',
    },
  },
}

export default networks
