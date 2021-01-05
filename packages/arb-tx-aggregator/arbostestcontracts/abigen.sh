#!/bin/bash
abigen --sol=simple.sol --pkg=arbostestcontracts --out=simple.go
abigen --sol=failedsend.sol --pkg=arbostestcontracts --out=failedsend.go
abigen --sol=storage.sol --pkg=arbostestcontracts --out=storage.go
abigen --sol=create2.sol --pkg=arbostestcontracts --out=create2.go
abigen --sol=receiver.sol --pkg=arbostestcontracts --out=receiver.go
abigen --sol=opcodes.sol --pkg=arbostestcontracts --out=opcodes.go
abigen --sol=transfer.sol --pkg=arbostestcontracts --out=transfer.go
abigen --sol=gasused.sol --pkg=arbostestcontracts --out=gasused.go
abigen --sol=Fibonacci.sol --pkg=arbostestcontracts --out=fibonacci.go
abigen --sol=token.sol --pkg=arbostestcontracts --out=token.go
