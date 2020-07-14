#!/bin/bash
PREFIX=../../arbos-contracts/contracts
PACKAGE=arboscontracts

abigen --sol=$PREFIX/ArbInfo.sol --pkg=$PACKAGE --out=arbinfo.go
abigen --sol=$PREFIX/ArbSys.sol --pkg=$PACKAGE --out=arbsys.go
