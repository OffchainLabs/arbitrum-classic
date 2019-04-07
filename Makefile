### --------------------------------------------------------------------
### Makefile
### Runs docker-compose.yml
### --------------------------------------------------------------------

.PHONY: all clean src
BIN=compose

all: src
	# Prebuild images that don't run
	cd $(BIN)/truffle-deploy-demo && sudo docker build -t truffle-deploy-demo .
	cd $(BIN)/arbc-solidity && sudo docker build -t arbc-solidity .
	# Build and run demo, ethbridge, and validator
	cd $(BIN) && time sudo docker-compose build
	cd $(BIN) && sudo docker-compose up

clean:
	rm -rf $(BIN)

# Download all source repos from git if $(BIN) does not exist
src:
	if [ ! -d "$(BIN)" ]; then                                                 \
        mkdir $(BIN);                                                      \
                                                                            \
        git clone https://github.com/OffchainLabs/demo-app.git              \
            $(BIN)/demo-app;                                               \
        git clone https://github.com/OffchainLabs/arb-web3-provider.git     \
            $(BIN)/demo-app/arb-web3-provider;                             \
                                                                            \
        git clone https://github.com/OffchainLabs/arbc-solidity.git         \
            $(BIN)/arbc-solidity;                                          \
                                                                            \
        git clone https://github.com/OffchainLabs/arb-ethbridge.git         \
            $(BIN)/arb-ethbridge;                                          \
                                                                            \
        git clone https://github.com/OffchainLabs/arb-validator.git         \
            $(BIN)/arb-validator;                                          \
        git clone https://github.com/OffchainLabs/arb-avm.git               \
            $(BIN)/arb-validator/arb-avm;                                  \
                                                                            \
        git clone https://github.com/OffchainLabs/truffle-deploy-demo.git   \
            $(BIN)/truffle-deploy-demo;                                    \
                                                                            \
        git clone https://github.com/OffchainLabs/arb-truffle-provider.git  \
            $(BIN)/truffle-deploy-demo/arb-truffle-provider;               \
    fi
	cp docker-compose.yml $(BIN) || true
	rm -f $(BIN)/demo-app/Makefile
