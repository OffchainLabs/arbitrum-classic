import { utils } from 'ethers'
export default {
  kovan4: {
    ethRPC: 'https://kovan.infura.io/v3/' + process.env['INFURA_KEY'],
    arbRPC: 'https://kovan5.arbitrum.io/rpc',
    // arbRPC: 'http://localhost:8547/' /* for port-forwarding */,
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x2948ac43e4AfF448f6af0F7a11F18Bb6062dd271',
    arbTokenBridgeAddress: '0x64b92d4f02cE1b4BDE2D16B6eAEe521E27f28e07',
    l1gasPrice: utils.parseUnits('10', 'gwei'),
    existentTestERC20: '0xA27B9a8185a840c46F1D5034bCC59450C221540F',
    existentCustomTokenL1: '0xCcec300d98325b491a5286983EC14295531C0B7F',
    existentCustomTokenL2: '0x5f612047a664D8Af77C68B91d6237C879C99f4E6',
    // existentTestERC20: '',
    // existentCustomTokenL1: '',
    // existentCustomTokenL2: '',
    defaultWait: 0,
  },
  kovan5: {
    ethRPC: 'https://kovan.infura.io/v3/' + process.env['INFURA_KEY'],
    arbRPC: 'https://kovan5.arbitrum.io/rpc',
    // arbRPC: 'http://localhost:8547/' /* for port-forwarding */,
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0xB70e15f33275d626fb958b63FcD4AA30d8D55072',
    arbTokenBridgeAddress: '0x4C4A16436cB80274725861ED798C7f4e3D6C4d55',
    l1gasPrice: utils.parseUnits('10', 'gwei'),
    existentTestERC20: '0xFa3933797baF9572c4Cf8fA9a28f09ad3C8d30BB',
    existentCustomTokenL1: '',
    existentCustomTokenL2: '',
    // existentTestERC20: '',
    // existentCustomTokenL1: '',
    // existentCustomTokenL2: '',
    defaultWait: 0,
  },

  devnet: {
    ethRPC: 'https://devnet.arbitrum.io/rpc',
    arbRPC: 'https://devnet-l2.arbitrum.io/rpc',
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x13bd06778868edC76bdeBB20657cD72Ce4540748',
    arbTokenBridgeAddress: '0xf4A1465abAb18B66B79bAc2EA23B1cac9f3e4d8E',
    l1gasPrice: 0,
    existentTestERC20: '',
    existentCustomTokenL1: '',
    existentCustomTokenL2: '',
    defaultWait: 10000,
  },
}
