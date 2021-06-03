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
    existentTestERC20: '0xA27B9a8185a840c46F1D5034bCC59450C221540F',
    existentCustomTokenL1: '0xCcec300d98325b491a5286983EC14295531C0B7F',
    existentCustomTokenL2: '0x5f612047a664D8Af77C68B91d6237C879C99f4E6',
    // existentTestERC20: '',
    // existentCustomTokenL1: '',
    // existentCustomTokenL2: '',
    defaultWait: 0,
    executeOutGoingMessages: false,
    outBoxUpdateTimeout: 10000,
  },
  kovan5: {
    ethRPC: 'https://kovan.infura.io/v3/' + process.env['INFURA_KEY'],
    arbRPC: 'https://kovan5.arbitrum.io/rpc',
    // arbRPC: 'http://localhost:8547/' /* for port-forwarding */,
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0x1d750369c91b129524B68f308512b0FE2C903d71',
    arbTokenBridgeAddress: '0x2EEBB8EE9c377caBC476654ca4aba016ECA1B9fc',
    l1gasPrice: utils.parseUnits('10', 'gwei'),
    existentTestERC20: '0xFa3933797baF9572c4Cf8fA9a28f09ad3C8d30BB',
    existentCustomTokenL1: '0x2091C1d0d20270e4cA15270b737421Ba90Cf0470',
    existentCustomTokenL2: '0x88692aE036594c95F723204c3DB070E2fB4d8eE4',

    defaultWait: 0,
    executeOutGoingMessages: false,
    outBoxUpdateTimeout: 100000,
  },
  arbRinkeby: {
    ethRPC: 'https://rinkeby.infura.io/v3/' + process.env['INFURA_KEY'],
    arbRPC: 'https://rinkeby.arbitrum.io/rpc',
    // arbRPC: 'http://localhost:8547/' /* for port-forwarding */,
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    erc20BridgeAddress: '0xdEee8125D79812E45491f2b760D420b43407F8Bd',
    arbTokenBridgeAddress: '0xB80954bdD2A41193B74A7a49edAD766C2385466C',
    l1gasPrice: utils.parseUnits('4', 'gwei'),
    existentTestERC20: '0xB83bCBB164a23817301f09dBd426E907E9d28ad4',
    defaultWait: 0,
    executeOutGoingMessages: false,
    outBoxUpdateTimeout: 100000,
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
    executeOutGoingMessages: false,
    outBoxUpdateTimeout: 10000,
  },
}
