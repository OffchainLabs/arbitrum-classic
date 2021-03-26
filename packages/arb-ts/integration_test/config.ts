import { utils } from 'ethers'
export default {
  kovan4: {
    ethRPC: 'https://kovan.infura.io/v3/' + process.env['INFURA_KEY'],
    arbRPC: 'https://kovan4.arbitrum.io/rpc',
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x70d05779a4c6c3541051Ed3e358323a6487B1d27',
    arbTokenBridgeAddress: '0x7DdCB8365BC78b7520EA92d088A1b548EbC73c15',
    l1gasPrice: utils.parseUnits('4', 'gwei'),
    existantTestERC20: '',
    // existantTestERC20: '0xB422417c4b7D981C32e99b4250f6d20Cb09EC118',
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
