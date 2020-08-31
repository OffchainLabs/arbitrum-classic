#!/bin/bash
PREFIX=../../arb-bridge-eth/contracts
TESTER_PREFIX=$PREFIX/test_only
IGNORED_LIB=$PREFIX/libraries/Keccak.sol:Keccak,$PREFIX/libraries/MerkleLib.sol:MerkleLib,$PREFIX/libraries/BytesLib.sol:BytesLib,$PREFIX/libraries/CloneFactory.sol:CloneFactory,$PREFIX/libraries/BytesLib.sol:BytesLib,$PREFIX/libraries/DebugPrint.sol:DebugPrint,$PREFIX/libraries/RollupTime.sol:RollupTime
IGNORED_ARCH=$PREFIX/arch/Value.sol:Value,$PREFIX/arch/Marshaling.sol:Marshaling,$PREFIX/arch/Hashing.sol:Hashing,$PREFIX/arch/Protocol.sol:Protocol,$PREFIX/arch/Machine.sol:Machine
CHAL_PREFIX=$PREFIX/challenge
IGNORED_CHALLENGE=$CHAL_PREFIX/ChallengeUtils.sol:ChallengeUtils,$CHAL_PREFIX/IChallengeFactory.sol:IChallengeFactory,$CHAL_PREFIX/IBisectionChallenge.sol:IBisectionChallenge
IGNORED_INBOX=$PREFIX/inbox/IGlobalInbox.sol:IGlobalInbox,$PREFIX/inbox/Messages.sol:Messages
IGNORED_ROLLUP=$PREFIX/rollup/IStaking.sol:IStaking,$PREFIX/rollup/IArbRollup.sol:IArbRollup
IGNORED=$IGNORED_LIB,$IGNORED_ARCH,$IGNORED_CHALLENGE,$IGNORED_INBOX,$IGNORED_ROLLUP
PACKAGE=ethbridgetestcontracts
abigen --sol=$TESTER_PREFIX/ChallengeTester.sol --pkg=$PACKAGE --out=challengetester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/MachineTester.sol --pkg=$PACKAGE --out=machinetester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/MessageTester.sol --pkg=$PACKAGE --out=messagetester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/KeccakTester.sol --pkg=$PACKAGE --out=keccaktester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/ValueTester.sol --pkg=$PACKAGE --out=valuetester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/RollupTester.sol --pkg=$PACKAGE --out=rolluptester.go --exc=$IGNORED
#abigen --sol=$TESTER_PREFIX/BuddyERC20.sol --pkg=$PACKAGE --out=buddyERC20.go --exc=$IGNORED
