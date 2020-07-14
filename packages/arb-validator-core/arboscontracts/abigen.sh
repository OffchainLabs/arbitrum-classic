#!/bin/bash
PREFIX=../../arb-compiler-evm/contract-templates/contracts
PACKAGE=arboscontracts

abigen --sol=$PREFIX/ArbInfo.sol --pkg=$PACKAGE --out=arbinfo.go
abigen --sol=$PREFIX/ArbSys.sol --pkg=$PACKAGE --out=arbsys.go
