#!/bin/bash
if [[ $PWD == */packages/arb-bridge-eth ]];
    then $npm_execpath run hardhat compile;
    else $npm_execpath run hardhat:prod compile;
fi
