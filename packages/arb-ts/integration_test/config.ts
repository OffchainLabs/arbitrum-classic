import { utils } from 'ethers'
export default {
  kovan4: {
    ethRPC: 'https://kovan.infura.io/v3/' + process.env['INFURA_KEY'],
    arbRPC: 'https://kovan4.arbitrum.io/rpc',
    // arbRPC: 'http://localhost:8547/' /* for port-forwarding */,
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x2948ac43e4AfF448f6af0F7a11F18Bb6062dd271',
    arbTokenBridgeAddress: '0x64b92d4f02cE1b4BDE2D16B6eAEe521E27f28e07',
    l1gasPrice: utils.parseUnits('10', 'gwei'),
    existantTestERC20: '0xA27B9a8185a840c46F1D5034bCC59450C221540F',
    existantCustomTokenL1: '0xCcec300d98325b491a5286983EC14295531C0B7F',
    existantCustomTokenL2: '0x5f612047a664D8Af77C68B91d6237C879C99f4E6',
    // existantTestERC20: '',
    // existantCustomTokenL1: '',
    // existantCustomTokenL2: '',
    defaultWait: 0,
  },
  devnet: {
    ethRPC: 'https://devnet.arbitrum.io/rpc',
    arbRPC: 'https://devnet-l2.arbitrum.io/rpc',
    preFundedSignerPK:
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x9DDede4e09DCF6B2C04C62b16B8abEaCD4B3C7aE',
    arbTokenBridgeAddress: '0xefaA73f05e5441b57C9dB7498e8bA5dd77Cfd8a2',
    l1gasPrice: 0,
    existantTestERC20: '',
    existantCustomTokenL1: '',
    existantCustomTokenL2: '',
    defaultWait: 10000,
  },
}
