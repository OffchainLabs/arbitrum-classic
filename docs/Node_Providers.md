---
id: Node_Providers
title: Alternatives
sidebar_label: Node Providers
---

To interct with Arbitrum One and/or Arbirum public testnet, you can also rely on popular node providers that offer access to the nodes of these two networks. All you need to do is to choose the provider and the required node, and the node service provider will carry out out all the settings and installation. Here, we list these that help you build on our mainnet anf testnet.

## Alchemy

Alchemy provides two networks on Arbitrum: Arbitrum One (Mainnet) and Rinkeby testnet. Depending one which network you want to interact with, you can add one of the following endpoints to your RPC URL:
    
    Mainnet: https://arb-mainnet.g.alchemy.com/v2/your-api-key
    Rinkeby: https://arb-rinkeby.g.alchemy.com/v2/your-api-key

*Not sure how to create an Alchemy API key?*
To use Alchemy's products, you need an API key to authenticate your requests. For dteilaed steps of how to create one, see the [Alchemy's Tutorials](https://docs.alchemy.com/alchemy/introduction/getting-started#1.create-an-alchemy-key).


## Infura
Infura is another node service provider that provides access to Arbitrum One (Mainnet) and Rinkeby testnet. To start interacting with either of these networks, you need to set up an Infura account and create a project. Next, you will be given a ***Project ID*** (For detailed steps, see the [Infura's Tutorials](https://blog.infura.io/getting-started-with-infura-28e41844cc89/)). Finally, you can now select which endpoint you want to connect to. This could be Arbirum One, or it could be Rinkeby testnet. Depending on which one you choose, your endpoint URL will look like below. You can use these endpoint URLs to start interacting with Arbitrum.

    https://<network>.infura.io/v3/YOUR-PROJECT-ID

