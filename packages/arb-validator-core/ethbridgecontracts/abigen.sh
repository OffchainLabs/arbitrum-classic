#!/bin/bash
PREFIX=../../arb-bridge-eth/contracts
IGNORED_LIB=$PREFIX/libraries/MerkleLib.sol:MerkleLib,$PREFIX/libraries/BytesLib.sol:BytesLib,$PREFIX/libraries/CloneFactory.sol:CloneFactory,$PREFIX/libraries/BytesLib.sol:BytesLib,$PREFIX/libraries/DebugPrint.sol:DebugPrint,$PREFIX/libraries/RollupTime.sol:RollupTime
IGNORED_ARCH=$PREFIX/arch/Value.sol:Value,$PREFIX/arch/Protocol.sol:Protocol,$PREFIX/arch/Machine.sol:Machine
CHAL_PREFIX=$PREFIX/challenge
IGNORED_CHALLENGE=$CHAL_PREFIX/ChallengeUtils.sol:ChallengeUtils,$CHAL_PREFIX/IChallengeFactory.sol:IChallengeFactory,$CHAL_PREFIX/IBisectionChallenge.sol:IBisectionChallenge
IGNORED_INBOX=$PREFIX/inbox/IGlobalInbox.sol:IGlobalInbox,$PREFIX/inbox/Messages.sol:Messages
IGNORED_ROLLUP=$PREFIX/rollup/IStaking.sol:IStaking
IGNORED=$IGNORED_LIB,$IGNORED_ARCH,$IGNORED_CHALLENGE,$IGNORED_INBOX,$IGNORED_ROLLUP
IGNORED_WITH_CHALLENGES=$IGNORED,$CHAL_PREFIX/Challenge.sol:Challenge,$CHAL_PREFIX/BisectionChallenge.sol:BisectionChallenge
PACKAGE=ethbridgecontracts

abigen --sol=$PREFIX/rollup/ArbFactory.sol --pkg=$PACKAGE --out=arbfactory.go --exc=$IGNORED
abigen --sol=$PREFIX/rollup/ArbRollup.sol --pkg=$PACKAGE --out=arbrollup.go --exc=$IGNORED

abigen --sol=$PREFIX/challenge/ChallengeFactory.sol --pkg=$PACKAGE --out=challengefactory.go --exc=$IGNORED
abigen --sol=$PREFIX/challenge/InboxTopChallenge.sol --pkg=$PACKAGE --out=inboxtopchallenge.go --exc=$IGNORED
abigen --sol=$PREFIX/challenge/MessagesChallenge.sol --pkg=$PACKAGE --out=messageschallenge.go --exc=$IGNORED_WITH_CHALLENGES
abigen --sol=$PREFIX/challenge/ExecutionChallenge.sol --pkg=$PACKAGE --out=executionchallenge.go --exc=$IGNORED_WITH_CHALLENGES

abigen --sol=$PREFIX/inbox/GlobalInbox.sol --pkg=ethbridgecontracts --out=globalinbox.go --exc=$IGNORED
