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
IGNORED_LIB=$MERKLELIB,$BYTESLIB,$CLONEFACTORY,$DEBUGPRINT,$ROLLUPTIME,$CLONABLE,$ICLONABLE,$PRECOMPILES,$PREFIX/libraries/SafeMath.sol:SafeMath
ARCH_PREFIX=$PREFIX/arch
IGNORED_ARCH=$ARCH_PREFIX/Value.sol:Value,$ARCH_PREFIX/Marshaling.sol:Marshaling,$ARCH_PREFIX/Hashing.sol:Hashing,$ARCH_PREFIX/Machine.sol:Machine,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof,$ARCH_PREFIX/IOneStepProof.sol:IOneStepProof2,$ARCH_PREFIX/OneStepProofCommon.sol:OneStepProofCommon
CHAL_PREFIX=$PREFIX/challenge
IGNORED_CHALLENGE=$CHAL_PREFIX/ChallengeLib.sol:ChallengeLib,$CHAL_PREFIX/IChallengeFactory.sol:IChallengeFactory,$CHAL_PREFIX/IChallenge.sol:IChallenge
#IGNORED_INBOX=$PREFIX/inbox/IGlobalInbox.sol:IGlobalInbox,$PREFIX/inbox/Messages.sol:Messages
IGNORED_ROLLUP=$PREFIX/rollup/IInbox.sol:IInbox,$PREFIX/rollup/INodeFactory.sol:INodeFactory,$PREFIX/rollup/IRollup.sol:IRollup
#IGNORED=$IGNORED_LIB,$IGNORED_ARCH,$IGNORED_CHALLENGE,$IGNORED_INBOX,$IGNORED_ROLLUP
#IGNORED_WITH_CHALLENGES=$IGNORED,$CHAL_PREFIX/Challenge.sol:Challenge,$CHAL_PREFIX/BisectionChallenge.sol:BisectionChallenge
MESSAGES=$PREFIX/bridge/Messages.sol:Messages
BRIDGE_LIBS=$PREFIX/bridge/interfaces/IBridge.sol:IBridge,$PREFIX/bridge/interfaces/ISequencerInbox.sol:ISequencerInbox,$PREFIX/bridge/interfaces/IOutbox.sol:IOutbox,$PREFIX/bridge/interfaces/IMessageProvider.sol:IMessageProvider,$MESSAGES

ROLLUP_LIB=$PREFIX/rollup/RollupLib.sol:RollupLib
ROLLUP=$PREFIX/rollup/Rollup.sol:Rollup
OUTBOX=$PREFIX/rollup/Outbox.sol:Outbox
INBOX=$PREFIX/rollup/Inbox.sol:Inbox

INODE=$PREFIX/rollup/INode.sol:INode
OUTBOX_ENTRY=$PREFIX/rollup/Outbox.sol:OutboxEntry
ROLLUP_LIBS=$INBOX,$OUTBOX,$ROLLUP,$ROLLUP_LIB,$INODE,$OUTBOX_ENTRY,$PREFIX/rollup/RollupCore.sol:RollupCore,$PREFIX/rollup/RollupEventBridge.sol:RollupEventBridge

IGNORED_INTERFACES=$PREFIX/interfaces/IERC20.sol:IERC20

CURRPATH=$(pwd)
FILEROOT=${CURRPATH%/*/*/*}
NM=$FILEROOT/node_modules
OZ=$NM/@openzeppelin
BASE=$FILEROOT/packages/arb-bridge-eth/contracts

OZUTILS=$OZ/contracts/utils
OZ_TOKENS=$OZ/contracts/token/ERC20/IERC20.sol:IERC20
OZ_MATH=$OZ/contracts/math/SafeMath.sol:SafeMath
OZ_PROXY=$OZ/contracts/proxy/Proxy.sol:Proxy,$OZ/contracts/proxy/TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy,$OZ/contracts/proxy/UpgradeableProxy.sol:UpgradeableProxy
OZ_LIBS=$OZUTILS/Address.sol:Address,$OZUTILS/Pausable.sol:Pausable,$OZ/contracts/utils/Context.sol:Context,$OZ/contracts/access/Ownable.sol:Ownable,$OZ_PROXY
#OZCONN=$OZ/contracts
IGNORED=$IGNORED_LIB,$IGNORED_CHALLENGE,$IGNORED_ROLLUP,$IGNORED_ARCH,$BRIDGE_LIBS,$OZ_MATH,$OZ_TOKENS,$OZ_LIBS,$IGNORED_INTERFACES
IGNORED_MORE=$IGNORED,$ROLLUP_LIBS,$OZ/contracts/proxy/ProxyAdmin.sol:ProxyAdmin

solc --combined-json bin,abi,userdoc,devdoc,metadata --optimize --optimize-runs=1 --allow-paths $BASE,$NM @openzeppelin=$OZ ../../arb-bridge-eth/contracts/validator/ValidatorUtils.sol --overwrite -o .
abigen --pkg=$PACKAGE --out=validatorutils.go --combined-json combined.json --exc=$IGNORED

solc --combined-json bin,abi,userdoc,devdoc,metadata --optimize --optimize-runs=1 --allow-paths $BASE,$NM @openzeppelin=$OZ ../../arb-bridge-eth/contracts/rollup/RollupCreator.sol --overwrite -o .
abigen --pkg=$PACKAGE --out=rollupcreator.go --combined-json combined.json --exc=$IGNORED_MORE

rm combined.json

abigen --sol=$PREFIX/challenge/Challenge.sol --pkg=$PACKAGE --out=challenge.go --exc=$IGNORED_MORE
abigen --sol=$PREFIX/validator/Validator.sol --pkg=$PACKAGE --out=validator.go --exc=$IGNORED_MORE
