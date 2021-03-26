---
id: Validators
title: Inside Arbitrum: Validators
sidebar_label: Validators
---
Some Arbitrum nodes will choose to act as *validators*. This means that they watch the progress of the rollup protocol and participate in that protocol to advance the state of the chain securely.

Not all nodes will choose to do this. Because the rollup protocol doesn’t decide what the chain will do but merely confirms the correct behavior that is fully determined by the inbox messages, a node can ignore the rollup protocol and simply compute for itself the correct behavior. For more on what such nodes might do, see the [Full Nodes](Full_Nodes.md) section.

Being a validator is permissionless--anyone can do it. Offchain Labs provides open source validator software, including a pre-built Docker image. 

Every validator can choose their own approach, but we expect validators to follow three common strategies. 

* The *active validator* strategy tries to advance the state of the chain by proposing new rollup blocks. An active validator is always staked, because creating a rollup block requires being staked. A chain really only needs one honest active validator; any more is an inefficient use of resources. For the flagship Arbitrum chain, Offchain Labs will run an active validator.
* The *defensive validator* strategy watches the rollup protocol operate. If only correct rollup blocks are proposed, this strategy does nothing. But if an incorrect block is proposed, this strategy intervenes by posting a correct block or staking on a correct block that another party has posted. This strategy avoids staking when things are going well, but if someone is dishonest it stakes in order to defend the correct outcome. 
* The *watchtower validator* strategy never stakes. It simply watches the rollup protocol and if an incorrect block is proposed, it raises the alarm (by whatever means it chooses) so that others can intervene. This strategy assumes that other parties who are willing to stake will be willing to intervene in order to take some of the dishonest proposer’s stake, and that that can happen before the dishonest block’s deadline expires. (In practice this will allow several days for a response.)

Under normal conditions, validators using the defensive and watchtower strategies won’t do anything except observe. A malicious actor who is considering whether to try cheating won’t be able to tell how many defensive and watchtower validators are operating incognito. Perhaps some defensive validators will announce themselves, but others probably won’t, so a would-be attacker will always have to worry that defenders are waiting to emerge.

Who will be validators? Anyone can do it, but most people will choose not to. In practice we expect people to validate a chain for several reasons.

- Some validators will be paid, by the party that created the chain or someone else. On the flagship Arbitrum chain, Offchain Labs will hire some validators.
- Parties who have significant assets at stake on a chain, such as dapp developers, exchanges, power-users, and liquidity providers, may choose to validate in order to protect their investment.
- Anyone who chooses to validate can do so. Some users will probably choose to validate in order to protect their own interests or just to be good citizens. But ordinary users don’t need to validate, and we expect that the vast majority of users won’t.
