---
id: EthBridge_version_support
title: EthBridge versions and where to find them
sidebar_label: EthBridge versions
---

Arbitrum's EthBridge protocol changes occasionally.  
The protocol versions are identified by numbers.

Every version of the protocol has its own Chain Factory, which is an Ethereum contract that you can call to launch Arbitrum Chains that use that version.
Each of these Chain Factories lives at its own Ethereum Address.
When you launch a new Arbitrum Chain, you can pick any protocol version that exists--although you'll normally want to use the latest version.
The tool you use to launch your Chain will connect to the Chain Factory for your chosen version, and launch a chain that uses that protocol version.

Client software needs to use a protocol version that is compatible with the version used by the Arbitrum Chain that the client is interacting with.
When you start your client software, it should check for version compatibility and report an error if the versions don't match.
(Maybe that's why you're reading this page.)

## Arbitrum protocol versions

Version: 1
* Chain Factory address: 
* Github link: 
