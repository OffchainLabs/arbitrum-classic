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
