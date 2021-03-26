import { utils } from 'ethers'
export default {
  kovan4: {
    ethRPC: 'https://kovan.infura.io/v3/' + process.env['INFURA_KEY'],
    arbRPC: 'https://kovan4.arbitrum.io/rpc',
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x47791c40b7022728493178087E46f04D5bA421FC',
    arbTokenBridgeAddress: '0xfD6608908a10e9Ab7a8325B18A465365260E2842',
    l1gasPrice: utils.parseUnits('4', 'gwei'),
    existantTestERC20: '0x9709549E22EB664E5B54E2f7d2ef5221dA25141b',
  },
  devnet: {
    ethRPC: 'https://devnet.arbitrum.io/rpc',
    arbRPC: 'https://devnet-l2.arbitrum.io/rpc',
    preFundedSignerPK:
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x9DDede4e09DCF6B2C04C62b16B8abEaCD4B3C7aE',
    arbTokenBridgeAddress: '0xefaA73f05e5441b57C9dB7498e8bA5dd77Cfd8a2',
    l1gasPrice: 0,
  },
}
