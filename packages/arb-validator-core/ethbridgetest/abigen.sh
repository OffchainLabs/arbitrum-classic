#!/bin/bash
PREFIX=../../arb-bridge-eth/contracts
TESTER_PREFIX=$PREFIX/test_only
IGNORED=$PREFIX/libraries/BytesLib.sol:BytesLib,$PREFIX/arch/Value.sol:Value,BytesLib,$PREFIX/arch/Protocol.sol:Protocol,$PREFIX/arch/Machine.sol:Machine,$PREFIX/libraries/BytesLib.sol:BytesLib,$PREFIX/libraries/DebugPrint.sol:DebugPrint
PACKAGE=ethbridgetest
abigen --sol=$TESTER_PREFIX/ChallengeTester.sol --pkg=$PACKAGE --out=challengetester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/MachineTester.sol --pkg=$PACKAGE --out=machinetester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/MessageTester.sol --pkg=$PACKAGE --out=messagetester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/OneStepProofTester.sol --pkg=$PACKAGE --out=onestepprooftester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/ValueTester.sol --pkg=$PACKAGE --out=valuetester.go --exc=$IGNORED
abigen --sol=$TESTER_PREFIX/RollupTester.sol --pkg=$PACKAGE --out=rolluptester.go --exc=$IGNORED
