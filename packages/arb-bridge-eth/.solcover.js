module.exports = {
  mocha: {
    reporter: 'mocha-junit-reporter',
    grep: '@skip-on-coverage',
    invert: true,
  },
  skipFiles: ['test_only'],
  providerOptions: {
    gasPrice: '0x2710',
    default_balance_ether: 100000000,
  },
}
