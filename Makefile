### --------------------------------------------------------------------
### Makefile
### Runs docker-compose.yml
### --------------------------------------------------------------------

.PHONY: all clean src

all: src
	#cd compose/arbc-solidity && sudo docker build -t arbc-solidity .
	time sudo docker-compose build
	sudo docker-compose up

# Download all source repos from git if "compose" folder does not exist
src:
	\
if [ ! -d "compose" ]; then                                                 \
        mkdir compose;                                                      \
                                                                            \
        git clone https://github.com/OffchainLabs/demo-app.git              \
            compose/demo-app;                                               \
        git clone https://github.com/OffchainLabs/arb-web3-provider.git     \
            compose/demo-app/arb-web3-provider;                             \
                                                                            \
        git clone https://github.com/OffchainLabs/arbc-solidity.git         \
            compose/arbc-solidity;                                          \
                                                                            \
        git clone https://github.com/OffchainLabs/arb-ethbridge.git         \
            compose/arb-ethbridge;                                          \
                                                                            \
        git clone https://github.com/OffchainLabs/arb-validator.git         \
            compose/arb-validator;                                          \
        git clone https://github.com/OffchainLabs/arb-avm.git               \
            compose/arb-validator/arb-avm;                                  \
                                                                            \
        git clone https://github.com/OffchainLabs/truffle-deploy-demo.git   \
            compose/truffle-deploy-demo;                                    \
                                                                            \
        git clone https://github.com/OffchainLabs/arb-truffle-provider.git  \
            compose/truffle-deploy-demo/arb-truffle-provider;               \
fi
