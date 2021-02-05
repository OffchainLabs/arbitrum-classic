#!/bin/bash
PACKAGE=ethbridgetestcontracts
PREFIX=../../arb-bridge-eth/contracts
MERKLELIB=$PREFIX/libraries/MerkleLib.sol:MerkleLib
BYTESLIB=$PREFIX/libraries/BytesLib.sol:BytesLib
CLONEFACTORY=$PREFIX/libraries/CloneFactory.sol:CloneFactory
DEBUGPRINT=$PREFIX/libraries/DebugPrint.sol:DebugPrint
ROLLUPTIME=$PREFIX/libraries/RollupTime.sol:RollupTime
CLONABLE=$PREFIX/libraries/Cloneable.sol:Cloneable
PRECOMPILES=$PREFIX/libraries/Precompiles.sol:Precompiles
ICLONABLE=$PREFIX/libraries/ICloneable.sol:ICloneable
SAFEMATH=$PREFIX/libraries/SafeMath.sol:SafeMath
IGNORED_LIB=$MERKLELIB,$BYTESLIB,$CLONEFACTORY,$DEBUGPRINT,$ROLLUPTIME,$CLONABLE,$ICLONABLE,$PRECOMPILES,$SAFEMATH
ARCH_PREFIX=$PREFIX/arch
IGNORED_ARCH=$ARCH_PREFIX/Value.sol:Value,$ARCH_PREFIX/Marshaling.sol:Marshaling,$ARCH_PREFIX/Hashing.sol:Hashing,$ARCH_PREFIX/Machine.sol:Machine,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof2,$ARCH_PREFIX/OneStepProofCommon.sol:OneStepProofCommon
CHAL_PREFIX=$PREFIX/challenge
IGNORED_CHALLENGE=$CHAL_PREFIX/Challenge.sol:Challenge,$CHAL_PREFIX/ChallengeLib.sol:ChallengeLib,$CHAL_PREFIX/IChallengeFactory.sol:IChallengeFactory,$CHAL_PREFIX/IChallenge.sol:IChallenge
IGNORED_ROLLUP=$PREFIX/rollup/IInbox.sol:IInbox,$PREFIX/rollup/INodeFactory.sol:INodeFactory,$PREFIX/rollup/IRollup.sol:IRollup
IGNORED=$IGNORED_LIB,$IGNORED_CHALLENGE,$IGNORED_ROLLUP,$IGNORED_ARCH
ROLLUP_LIB=$PREFIX/rollup/RollupLib.sol:RollupLib
ROLLUP=$PREFIX/rollup/Rollup.sol:Rollup
ROLLUP_CREATOR=$PREFIX/rollup/RollupCreator.sol:RollupCreator
OUTBOX=$PREFIX/rollup/Outbox.sol:Outbox
INBOX=$PREFIX/rollup/Inbox.sol:Inbox
MESSAGES=$PREFIX/bridge/Messages.sol:Messages
NODE=$PREFIX/rollup/Node.sol:Node
OUTBOX_ENTRY=$PREFIX/rollup/Outbox.sol:OutboxEntry
ROLLUP_LIBS=$INBOX,$OUTBOX,$ROLLUP_CREATOR,$ROLLUP,$ROLLUP_LIB,$MESSAGES,$NODE,$OUTBOX_ENTRY
IGNORED_MORE=$IGNORED,$ROLLUP_LIBS

abigen --sol=$PREFIX/test_only/ChallengeTester.sol --pkg=$PACKAGE --out=challengeTester.go --exc=$IGNORED_MORE
abigen --sol=$PREFIX/test_only/MachineTester.sol --pkg=$PACKAGE --out=machineTester.go --exc=$IGNORED_MORE
abigen --sol=$PREFIX/arch/OneStepProof.sol --pkg=$PACKAGE --out=onestepproof.go --exc=$IGNORED_MORE
abigen --sol=$PREFIX/arch/OneStepProof2.sol --pkg=$PACKAGE --out=onestepproof2.go --exc=$IGNORED_MORE
abigen --sol=$PREFIX/rollup/NodeFactory.sol --pkg=$PACKAGE --out=nodefactory.go --exc=$IGNORED_MORE
