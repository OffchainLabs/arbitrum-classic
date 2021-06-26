import { utils } from 'ethers'
import dotenv from 'dotenv'
dotenv.config()

export default {
  '4': {
    preFundedSignerPK:
      process.env['DEVNET_PRIVKEY'] ||
      '0x8803565d1ab75cf6a04656e2a638c65a2984f810ce2f5f8270601aca4e25e067',
    existentTestERC20: '0x2616Fd3e4e89dB180F570b200b13195597bEb337',
    defaultWait: 0,
    executeOutGoingMessages: false,
    outBoxUpdateTimeout: 100000,
  },
} as any
