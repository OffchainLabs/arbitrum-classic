import { utils } from 'ethers'
import dotenv from 'dotenv'
dotenv.config()

const mainnetConfig = {
  preFundedSignerPK:
    process.env['DEVNET_PRIVKEY'] ||
    '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
  existentTestERC20: '0x9f8F72aA9304c8B593d555F12eF6589cC3A579A2',
  executeOutGoingMessages: false,
  outBoxUpdateTimeout: 100000,
  existentTestCustomToken: '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48',
}

export default {
  '4': {
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    existentTestERC20: '0x2616Fd3e4e89dB180F570b200b13195597bEb337',
    executeOutGoingMessages: false,
    outBoxUpdateTimeout: 100000,
    existentTestCustomToken: '0x7A58e7f78893bcC15E1DAf6bfD08E527567C0552',
  },
  '1337': mainnetConfig,
  '1': mainnetConfig,
} as any
