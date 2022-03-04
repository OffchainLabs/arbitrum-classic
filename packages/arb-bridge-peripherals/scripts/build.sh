if [[ $PWD == */packages/arb-bridge-peripherals ]];
    then $npm_execpath run hardhat compile;
    else $npm_execpath run hardhat:prod compile;
fi
