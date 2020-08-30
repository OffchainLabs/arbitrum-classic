#!/bin/bash
PREFIX=../../arb-bridge-eth/contracts
MERKLELIB=$PREFIX/libraries/MerkleLib.sol:MerkleLib
BYTESLIB=$PREFIX/libraries/BytesLib.sol:BytesLib
CLONEFACTORY=$PREFIX/libraries/CloneFactory.sol:CloneFactory
DEBUGPRINT=$PREFIX/libraries/DebugPrint.sol:DebugPrint
ROLLUPTIME=$PREFIX/libraries/RollupTime.sol:RollupTime
CLONABLE=$PREFIX/libraries/Cloneable.sol:Cloneable
ICLONABLE=$PREFIX/libraries/ICloneable.sol:ICloneable
IGNORED_LIB=$MERKLELIB,$BYTESLIB,$CLONEFACTORY,$DEBUGPRINT,$ROLLUPTIME,$CLONABLE,$ICLONABLE
ARCH_PREFIX=$PREFIX/arch
IGNORED_ARCH=$ARCH_PREFIX/Value.sol:Value,$ARCH_PREFIX/Marshaling.sol:Marshaling,$ARCH_PREFIX/Hashing.sol:Hashing,$ARCH_PREFIX/Protocol.sol:Protocol,$ARCH_PREFIX/Machine.sol:Machine,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof
CHAL_PREFIX=$PREFIX/challenge
IGNORED_CHALLENGE=$CHAL_PREFIX/ChallengeUtils.sol:ChallengeUtils,$CHAL_PREFIX/IChallengeFactory.sol:IChallengeFactory,$CHAL_PREFIX/IBisectionChallenge.sol:IBisectionChallenge,$CHAL_PREFIX/IExecutionChallenge.sol:IExecutionChallenge
IGNORED_INBOX=$PREFIX/inbox/IGlobalInbox.sol:IGlobalInbox,$PREFIX/inbox/Messages.sol:Messages
IGNORED_ROLLUP=$PREFIX/rollup/IStaking.sol:IStaking,$PREFIX/rollup/IArbRollup.sol:IArbRollup
IGNORED=$IGNORED_LIB,$IGNORED_ARCH,$IGNORED_CHALLENGE,$IGNORED_INBOX,$IGNORED_ROLLUP
IGNORED_WITH_CHALLENGES=$IGNORED,$CHAL_PREFIX/Challenge.sol:Challenge,$CHAL_PREFIX/BisectionChallenge.sol:BisectionChallenge
PACKAGE=ethbridgecontracts

abigen --sol=$PREFIX/rollup/ArbFactory.sol --pkg=$PACKAGE --out=arbfactory.go --exc=$IGNORED
abigen --sol=$PREFIX/rollup/ArbRollup.sol --pkg=$PACKAGE --out=arbrollup.go --exc=$IGNORED

abigen --sol=$PREFIX/challenge/ChallengeFactory.sol --pkg=$PACKAGE --out=challengefactory.go --exc=$IGNORED
abigen --sol=$PREFIX/challenge/InboxTopChallenge.sol --pkg=$PACKAGE --out=inboxtopchallenge.go --exc=$IGNORED
abigen --sol=$PREFIX/challenge/ExecutionChallenge.sol --pkg=$PACKAGE --out=executionchallenge.go --exc=$IGNORED_WITH_CHALLENGES

abigen --sol=$PREFIX/arch/OneStepProof.sol --pkg=$PACKAGE --out=onestepproof.go --exc=$IGNORED

abigen --sol=$PREFIX/inbox/GlobalInbox.sol --pkg=ethbridgecontracts --out=globalinbox.go --exc=$IGNORED,$PREFIX/interfaces/IERC20.sol:IERC20
