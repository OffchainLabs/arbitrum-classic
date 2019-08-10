# arb-provider-go

This Go module allows your client-side Go code to interact with Arbitrum VMs. It is designed to be compatible with the ethclient functionality of go-ethereum.

If your Go-based client code runs with go-ethereum, you can switch to go-arbitrum by finding the go-ethereum `ethclient.Dial` call that your code is using to connect to Ethereum, and replacing it with a `goarbitrum.Dial` call.

Specifically, the call to connect to Arbitrum is `goarbitrum.Dial(url, myAddress, privateKey, hexPubkey)`, where `url` is the URL of an Arbitrum validator you want to connect to (or pass an empty string and it will guess that you want the local URL that arb-deploy uses), `myAddress` is the Ethereum address you are using, `privateKey` is the private key corresponding to that address, and `hexPubkey` is the corresponding public key hex-encoded as by `hexutil.Encode`.

This package implements the interface necessary to support the code that is produced by the standard `abigen` tool. But note that some of the less common functions in that interface are not implemented. Trying to call one of the not implemented calls will generate an error that conveys that you have called a functions that is not yet implemented.

Arbitrum technologies are patent pending. This repository is offered under the Apache 2.0 license. See LICENSE for details.
