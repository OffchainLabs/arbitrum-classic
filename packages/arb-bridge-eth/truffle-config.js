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
      network_id: "*",
      from: "0xe83f8ae25F873b1e17e05bda065ABEAc2FbD2E82"
    }
  },
  plugins: ["truffle-security"],
  compilers: {
    solc: {
      version: "0.5.10",
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
