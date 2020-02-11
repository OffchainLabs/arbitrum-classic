---
id: Executable_Creation
title: Executable Creation
sidebar_label: Executable Creation
---

Arbitrum supports compiling Ethereum Virtual Machine code into Arbitrum Virtual Machine (AVM) executable files. This allows standard solidity smart to be deployed on Arbitrum Chains.

`arbc-compiler` is the main tool for transforming EVM code into AVM code. It works on the level of EVM code rather than solidity, so additional tooling is useful for starting from solidity smart contracts.

To make life easy, we supply a plugin for the popular Truffle development framework which allows you to take a standard truffle project and convert its smart contracts into an AVM executable with minimal changes.

In order to integrate Arbitrum, you need to add a new network to your `truffle-config.js`

1.  First add the `arb-provider-truffle` to your project:

    ```bash
    yarn add arb-provider-truffle
    ```

2.  Edit the `truffle-config.js`:

    -   Import `arb-provider-truffle` and set the mnemonic at the top of the file. This mnemonic is used to set the caller of your contract's constructor when migrating.:

    ```js
    const ArbProvider = require("arb-provider-truffle");
    const mnemonic =
        "jar deny prosper gasp flush glass core corn alarm treat leg smart";
    ```

    -   Add the `arbitrum` network to `module.exports`:

    ```js
    module.exports = {
        networks: {
            arbitrum: {
                provider: function() {
                    if (!this.provider.prov) {
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
        }
    };
    ```

Now that the truffle project is setup correctly, the next step is to compile
the Truffle project into an AVM executable:

```bash
truffle migrate --reset --network arbitrum
```

After this command finished, the current directory will contain a `contract.ao` file which contains the code of your AVM executable.
