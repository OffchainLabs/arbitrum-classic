---
id: Above_or_Below_the_Line
title: Inside Arbitrum: Above or Below the Line?
sidebar_label: Above or Below the Line?
---
We often say that the key dividing line in the Arbitrum architecture is the AVM interface which divides Layer 1 from Layer 2. It can be useful to think about whether a particular activity is below the line or above the line.

![img](https://lh5.googleusercontent.com/1qwGMCrLQjJMv9zhWIUYkQXoDR2IksU5IzcSUPNJ5pWkY81pCvr7WkTf4-sb41cVohcnL-i6y8M1LU8v-4RXT_fdOsaMuLXnjwerSuKTQdHE-Hrvf4qBhRQ2r7qjxuAi3mk3hgkh)

Below the line functions are concerned with ensuring that the AVM, and therefore the chain, executes correctly. Above the line functions assume that the AVM will execute correctly, and focus on interacting with the software running at Layer 2.

As an example, Arbitrum validators operate below the line, because they participate in the rollup protocol, which is managed below-the-line by the EthBridge, to ensure that correct execution of the AVM is confirmed.

On the other hand, Arbitrum full nodes operate above the line, because they run a copy of the AVM locally, and assume that below-the-line mechanisms will ensure that the same result that they compute locally will eventually be confirmed by below-the-line mechanisms that they don’t monitor.

Most users, most of the time, will be thinking in above the line terms. They will be interacting with an Arbitrum chain as just another chain, without worrying about the below-the-line details that ensure that the chain won’t go wrong. 