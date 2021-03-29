import { utils } from 'ethers'
export default {
  kovan4: {
    ethRPC: 'https://kovan.infura.io/v3/' + process.env['INFURA_KEY'],
    arbRPC: 'https://kovan4.arbitrum.io/rpc',
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x5F530aBc9e173107C07c6A0b4Dc1Ebb3d7F79a5f',
    arbTokenBridgeAddress: '0x0Bf4eB4CFfA03541603Cb142Ea3120b529200C69',
    l1gasPrice: utils.parseUnits('4', 'gwei'),
    existantTestERC20: '0x39cA0c02a8c8e6b45cbb71E3a08dd8fB5204A12C',
    defaultWait: 10000,
  },
  devnet: {
    ethRPC: 'https://devnet.arbitrum.io/rpc',
    arbRPC: 'https://devnet-l2.arbitrum.io/rpc',
    preFundedSignerPK:
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x9DDede4e09DCF6B2C04C62b16B8abEaCD4B3C7aE',
    arbTokenBridgeAddress: '0xefaA73f05e5441b57C9dB7498e8bA5dd77Cfd8a2',
    l1gasPrice: 0,
    defaultWait: 0,
  },
}
