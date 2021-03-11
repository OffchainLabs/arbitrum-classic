#!/bin/bash
PREFIX=../../arb-os/contracts/arbos/builtin
PACKAGE=arboscontracts

abigen --sol=$PREFIX/ArbInfo.sol --pkg=$PACKAGE --out=arbinfo.go
abigen --sol=$PREFIX/ArbSys.sol --pkg=$PACKAGE --out=arbsys.go
abigen --sol=$PREFIX/ArbAddressTable.sol --pkg=$PACKAGE --out=arbaddresstable.go
abigen --sol=$PREFIX/ArbBLS.sol --pkg=$PACKAGE --out=arbbls.go
abigen --sol=$PREFIX/ArbFunctionTable.sol --pkg=$PACKAGE --out=arbfunctiontable.go
abigen --sol=$PREFIX/ArbOwner.sol --pkg=$PACKAGE --out=arbowner.go
abigen --sol=$PREFIX/ArbRetryableTx.sol --pkg=$PACKAGE --out=arbretryable.go
abigen --sol=NodeInterface.sol --pkg=$PACKAGE --out=nodeinterface.go
