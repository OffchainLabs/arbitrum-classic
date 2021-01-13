#!/bin/bash
PACKAGE=ethbridgecontracts
PREFIX=../../arb-bridge-eth/contracts
MERKLELIB=$PREFIX/libraries/MerkleLib.sol:MerkleLib
BYTESLIB=$PREFIX/libraries/BytesLib.sol:BytesLib
CLONEFACTORY=$PREFIX/libraries/CloneFactory.sol:CloneFactory
DEBUGPRINT=$PREFIX/libraries/DebugPrint.sol:DebugPrint
ROLLUPTIME=$PREFIX/libraries/RollupTime.sol:RollupTime
CLONABLE=$PREFIX/libraries/Cloneable.sol:Cloneable
PRECOMPILES=$PREFIX/libraries/Precompiles.sol:Precompiles
ICLONABLE=$PREFIX/libraries/ICloneable.sol:ICloneable
IGNORED_LIB=$MERKLELIB,$BYTESLIB,$CLONEFACTORY,$DEBUGPRINT,$ROLLUPTIME,$CLONABLE,$ICLONABLE,$PRECOMPILES
ARCH_PREFIX=$PREFIX/arch
IGNORED_ARCH=$ARCH_PREFIX/Value.sol:Value,$ARCH_PREFIX/Marshaling.sol:Marshaling,$ARCH_PREFIX/Hashing.sol:Hashing,$ARCH_PREFIX/Machine.sol:Machine,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof2,$ARCH_PREFIX/OneStepProofCommon.sol:OneStepProofCommon
CHAL_PREFIX=$PREFIX/challenge
IGNORED_CHALLENGE=$CHAL_PREFIX/ChallengeLib.sol:ChallengeLib,$CHAL_PREFIX/IChallengeFactory.sol:IChallengeFactory,$CHAL_PREFIX/IChallenge.sol:IChallenge
#IGNORED_INBOX=$PREFIX/inbox/IGlobalInbox.sol:IGlobalInbox,$PREFIX/inbox/Messages.sol:Messages
IGNORED_ROLLUP=$PREFIX/rollup/IInbox.sol:IInbox,$PREFIX/rollup/INodeFactory.sol:INodeFactory,$PREFIX/rollup/IRollup.sol:IRollup
#IGNORED=$IGNORED_LIB,$IGNORED_ARCH,$IGNORED_CHALLENGE,$IGNORED_INBOX,$IGNORED_ROLLUP
#IGNORED_WITH_CHALLENGES=$IGNORED,$CHAL_PREFIX/Challenge.sol:Challenge,$CHAL_PREFIX/BisectionChallenge.sol:BisectionChallenge
IGNORED=$IGNORED_LIB,$IGNORED_CHALLENGE,$IGNORED_ROLLUP,$IGNORED_ARCH
ROLLUP_LIB=$PREFIX/rollup/RollupLib.sol:RollupLib
ROLLUP=$PREFIX/rollup/Rollup.sol:Rollup
ROLLUP_CREATOR=$PREFIX/rollup/RollupCreator.sol:RollupCreator
OUTBOX=$PREFIX/rollup/Outbox.sol:Outbox
INBOX=$PREFIX/rollup/Inbox.sol:Inbox
MESSAGES=$PREFIX/rollup/Messages.sol:Messages
INODE=$PREFIX/rollup/INode.sol:INode
OUTBOX_ENTRY=$PREFIX/rollup/Outbox.sol:OutboxEntry
ROLLUP_LIBS=$INBOX,$OUTBOX,$ROLLUP_CREATOR,$ROLLUP,$ROLLUP_LIB,$MESSAGES,$INODE,$OUTBOX_ENTRY
IGNORED_MORE=$IGNORED,$ROLLUP_LIBS
abigen --sol=$PREFIX/rollup/RollupCreator.sol --pkg=$PACKAGE --out=rollupcreator.go --exc=$IGNORED
abigen --sol=$PREFIX/validator/ValidatorUtils.sol --pkg=$PACKAGE --out=validatorutils.go --exc=$IGNORED_MORE
abigen --sol=$PREFIX/challenge/Challenge.sol --pkg=$PACKAGE --out=challenge.go --exc=$IGNORED_MORE
abigen --sol=$PREFIX/validator/Validator.sol --pkg=$PACKAGE --out=validator.go --exc=$IGNORED_MORE
