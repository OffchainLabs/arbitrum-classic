const ArbProvider = require("arb-provider-truffle");
const mnemonic =
  "jar deny prosper gasp flush glass core corn alarm treat leg smart";
module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // for more about customizing your Truffle configuration!
  networks: {
    development: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "*" // Match any network id
    },
    arbitrum: {
      provider: function() {
        if (typeof this.provider.prov == "undefined") {
          this.provider.prov = ArbProvider.provider(
            __dirname,
            "build/contracts",
            {
              mnemonic: mnemonic
            }
          );
        }
        return this.provider.prov;
      },
      network_id: "*"
    }
  },
  compilers: {
    solc: {
      version: "0.4.25", // Fetch exact version from solc-bin (default: truffle's version)
      docker: true, // Use "0.5.3" you've installed locally with docker (default: false)
      settings: {
        // See the solidity docs for advice about optimization and evmVersion
        optimizer: {
          enabled: true,
          runs: 200
        }
      }
    }
  }
};
