#!/bin/bash
PREFIX=../../arbos-contracts/contracts
PACKAGE=arboscontracts
NM=$(realpath ./../../../node_modules)
OZ=$NM/@openzeppelin
OZCONN=$OZ/contracts
IGNORED=$PREFIX/ArbSys.sol:ArbSys,$OZCONN/GSN/Context.sol:Context,$OZCONN/math/SafeMath.sol:SafeMath

abigen --sol=$PREFIX/ArbInfo.sol --pkg=$PACKAGE --out=arbinfo.go
abigen --sol=$PREFIX/ArbSys.sol --pkg=$PACKAGE --out=arbsys.go
abigen --sol=$PREFIX/ArbAddressTable.sol --pkg=$PACKAGE --out=arbaddresstable.go
abigen --sol=$PREFIX/ArbBLS.sol --pkg=$PACKAGE --out=arbbls.go
abigen --sol=$PREFIX/ArbFunctionTable.sol --pkg=$PACKAGE --out=arbfunctiontable.go
solc --combined-json bin,abi,userdoc,devdoc,metadata --allow-paths $NM @openzeppelin=$OZ ../../arbos-contracts/contracts/ArbERC20.sol --overwrite -o .
abigen --pkg=arboscontracts --out=arbERC20.go --combined-json combined.json --exc=$IGNORED
solc --combined-json bin,abi,userdoc,devdoc,metadata --allow-paths $NM @openzeppelin=$NM/@openzeppelin ../../arbos-contracts/contracts/ArbERC721.sol --overwrite -o .
abigen --pkg=arboscontracts --out=arbERC721.go --combined-json combined.json --exc=$IGNORED
rm combined.json
