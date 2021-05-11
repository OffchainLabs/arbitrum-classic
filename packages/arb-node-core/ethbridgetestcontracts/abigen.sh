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
IGNORED_ARCH=$ARCH_PREFIX/Value.sol:Value,$ARCH_PREFIX/Marshaling.sol:Marshaling,$ARCH_PREFIX/Hashing.sol:Hashing,$ARCH_PREFIX/Machine.sol:Machine,$ARCH_PREFIX/OneStepProofCommon.sol:OneStepProofCommon
CHAL_PREFIX=$PREFIX/challenge
IGNORED_CHALLENGE=$CHAL_PREFIX/Challenge.sol:Challenge,$CHAL_PREFIX/ChallengeLib.sol:ChallengeLib,$CHAL_PREFIX/IChallengeFactory.sol:IChallengeFactory,$CHAL_PREFIX/IChallenge.sol:IChallenge
IGNORED_ROLLUP=$PREFIX/rollup/IInbox.sol:IInbox,$PREFIX/rollup/INodeFactory.sol:INodeFactory,$PREFIX/rollup/IRollup.sol:IRollup
IGNORED=$IGNORED_LIB,$IGNORED_CHALLENGE,$IGNORED_ROLLUP,$IGNORED_ARCH,$PREFIX/bridge/interfaces/IBridge.sol:IBridge,$PREFIX/bridge/interfaces/ISequencerInbox.sol:ISequencerInbox,$PREFIX/bridge/interfaces/IMessageProvider.sol:IMessageProvider
ROLLUP_LIB=$PREFIX/rollup/RollupLib.sol:RollupLib
ROLLUP=$PREFIX/rollup/Rollup.sol:Rollup
ROLLUP_CREATOR=$PREFIX/rollup/RollupCreator.sol:RollupCreator
OUTBOX=$PREFIX/rollup/Outbox.sol:Outbox
INBOX=$PREFIX/rollup/Inbox.sol:Inbox
MESSAGES=$PREFIX/bridge/Messages.sol:Messages
NODE=$PREFIX/rollup/Node.sol:Node
OUTBOX_ENTRY=$PREFIX/rollup/Outbox.sol:OutboxEntry
INODE=$PREFIX/rollup/INode.sol:INode
ROLLUP_LIBS=$INBOX,$OUTBOX,$ROLLUP_CREATOR,$ROLLUP,$ROLLUP_LIB,$MESSAGES,$NODE,$OUTBOX_ENTRY,$INODE

IGNORED_MORE=$IGNORED,$ROLLUP_LIBS

NM=$(realpath ./../../../node_modules)
OZ=$NM/@openzeppelin
OZUTILS=$OZ/contracts/utils
OZ_MATH=$OZ/contracts/math/SafeMath.sol:SafeMath
BASE=$(realpath ./../../arb-bridge-eth/contracts)

OZ_LIBS=$OZ/contracts/proxy/Clones.sol:Clones,$OZUTILS/Address.sol:Address,$OZ_MATH

solc --combined-json bin,abi,userdoc,devdoc,metadata --optimize --optimize-runs=1 --allow-paths $BASE,$NM @openzeppelin=$OZ ../../arb-bridge-eth/contracts/rollup/NodeFactory.sol --overwrite -o .
abigen --pkg=$PACKAGE --out=nodefactory.go --combined-json combined.json --exc=$IGNORED_MORE

solc --combined-json bin,abi,userdoc,devdoc,metadata --optimize --optimize-runs=1 --allow-paths $BASE,$NM @openzeppelin=$OZ ../../arb-bridge-eth/contracts/challenge/ChallengeFactory.sol --overwrite -o .
abigen --pkg=$PACKAGE --out=challengefactory.go --combined-json combined.json --exc=$IGNORED_MORE,$OZ_LIBS

solc --combined-json bin,abi,userdoc,devdoc,metadata --optimize --optimize-runs=1 --allow-paths $BASE,$NM @openzeppelin=$OZ ../../arb-bridge-eth/contracts/test_only/ChallengeTester.sol --overwrite -o .
abigen --pkg=$PACKAGE --out=challengeTester.go --combined-json combined.json --exc=$IGNORED_MORE,$OZ_LIBS,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof,$CHAL_PREFIX/ChallengeFactory.sol:ChallengeFactory

solc --combined-json bin,abi,userdoc,devdoc,metadata --optimize --optimize-runs=1 --allow-paths $BASE,$NM @openzeppelin=$OZ ../../arb-bridge-eth/contracts/rollup/RollupCreatorNoProxy.sol --overwrite -o .
abigen --pkg=$PACKAGE --out=rollupcreatornoproxy.go --combined-json combined.json --exc=$IGNORED_MORE,$OZ_LIBS,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof

rm combined.json

abigen --sol=$PREFIX/test_only/MachineTester.sol --pkg=$PACKAGE --out=machineTester.go --exc=$IGNORED_MORE
abigen --sol=$PREFIX/arch/OneStepProof.sol --pkg=$PACKAGE --out=onestepproof.go --exc=$IGNORED_MORE,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof
abigen --sol=$PREFIX/arch/OneStepProof2.sol --pkg=$PACKAGE --out=onestepproof2.go --exc=$IGNORED_MORE,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof
abigen --sol=$PREFIX/arch/OneStepProofHash.sol --pkg=$PACKAGE --out=onestepproofhash.go --exc=$IGNORED_MORE,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof
abigen --sol=$PREFIX/test_only/InboxHelperTester.sol --pkg=$PACKAGE --out=inboxhelpertester.go --exc=$IGNORED_MORE
