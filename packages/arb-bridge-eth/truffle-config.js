module.exports = {
  networks: {
    development: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*",
      websockets: true
    },
    parity: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*"
    },
    rinkeby: {
      host: "localhost", // Connect to geth on the specified
      port: 8545,
      from: "0xf4a7f2c6bbe40a67e74f1b44bed16c6302eb07f6", // default address to use for any transaction Truffle makes during migrations
      network_id: 4
    }
  },
  plugins: ["truffle-security"],
  compilers: {
    solc: {
      version: "0.5.15",
      // docker: true,
      settings: {
        optimizer: {
          enabled: true,
          runs: 200
        }
      }
    }
  }
};
