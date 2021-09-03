---
id: ArbOS
title: ArbOS: The Arbitrum Operating System
sidebar_label: ArbOS
---

# ArbOS

ArbOS is a trusted "operating system” at Layer 2 that isolates untrusted contracts from each other, tracks and limits their resource usage, and manages the economic model that collects fees from users to fund the operation of a chain's validators. Much of the work that would otherwise have to be done expensively at Layer 1 is instead by ArbOS, trustlessly performing these functions at the speed and low cost of Layer 2.

Supporting these functions in Layer 2 trusted software, rather than building them in to the L1-enforced rules of the architecture as Ethereum does, offers significant advantages in cost because these operations can benefit from the lower cost of computation and storage at Layer 2, instead of having to manage those resources as part of the Layer 1 EthBridge contract. Having a trusted operating system at Layer 2 also has significant advantages in flexibility, because Layer 2 code is easier to evolve, or to customize for a particular chain, than a Layer-1 enforced VM architecture would be.

The use of a Layer 2 trusted operating system does require some support in the architecture, for example to allow the OS to limit and track resource usage by contracts

For a detailed specification describing the format of messages used for communication between clients, the EthBridge, and ArbOS, see the [ArbOS Message Formats Specification](ArbOS_Formats.md).
