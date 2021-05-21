#!/bin/bash
PREFIX=../../arb-os/contracts/arbos/builtin
PREFIX2=../../arb-bridge-peripherals/contracts/rpc-utils
PACKAGE=arboscontracts

abigen --sol=$PREFIX/ArbInfo.sol --pkg=$PACKAGE --out=arbinfo.go
abigen --sol=$PREFIX/ArbSys.sol --pkg=$PACKAGE --out=arbsys.go
abigen --sol=$PREFIX/ArbAddressTable.sol --pkg=$PACKAGE --out=arbaddresstable.go
abigen --sol=$PREFIX/ArbAggregator.sol --pkg=$PACKAGE --out=arbaggregator.go
abigen --sol=$PREFIX/ArbBLS.sol --pkg=$PACKAGE --out=arbbls.go
abigen --sol=$PREFIX/ArbFunctionTable.sol --pkg=$PACKAGE --out=arbfunctiontable.go
abigen --sol=$PREFIX/ArbOwner.sol --pkg=$PACKAGE --out=arbowner.go
abigen --sol=$PREFIX/ArbGasInfo.sol --pkg=$PACKAGE --out=arbgasinfo.go
abigen --sol=$PREFIX/ArbRetryableTx.sol --pkg=$PACKAGE --out=arbretryable.go

abigen --sol=$PREFIX2/NodeInterface.sol --pkg=$PACKAGE --out=nodeinterface.go
abigen --sol=$PREFIX2/RetryableTicketCreator.sol --pkg=$PACKAGE --out=retryableticketcreator.go
