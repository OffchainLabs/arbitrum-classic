### --------------------------------------------------------------------
### Makefile
### Runs docker-compose.yml
### --------------------------------------------------------------------

# Repository names
DEMO=demo-app
COMPILER=arbc-solidity
ETHBRIDGE=arb-ethbridge
VALIDATOR=arb-validator
AWP=arb-web3-provider
TDD=truffle-deploy-demo
ATP=arb-truffle-provider

.PHONY: all clean src
BIN=compose

all: src
	# Prebuild images that don't run
	cd $(BIN)/$(TDD) && sudo docker build -t $(TDD) .
	cd $(BIN)/$(COMPILER) && sudo docker build -t $(COMPILER) .
	# Prebuild addresses.json image to avoid depends_on
	cd $(BIN)/$(ETHBRIDGE) && sudo docker build -t $(ETHBRIDGE) .
	# Build and run demo, ethbridge, and validator
	cd $(BIN) && time sudo docker-compose build
	cd $(BIN) && sudo docker-compose up

clean:
	rm -rf $(BIN)

# Download all source repos from git if $(BIN) does not exist
src:
	if [ ! -d "$(BIN)" ]; then                                      \
	    mkdir $(BIN);                                               \
	    git clone https://github.com/OffchainLabs/$(DEMO).git       \
	        $(BIN)/$(DEMO);                                         \
	    git clone https://github.com/OffchainLabs/$(AWP).git        \
	        $(BIN)/$(DEMO)/$(AWP);                                  \
	    git clone https://github.com/OffchainLabs/$(COMPILER).git   \
	        $(BIN)/$(COMPILER);                                     \
	    git clone https://github.com/OffchainLabs/$(ETHBRIDGE).git  \
	        $(BIN)/$(ETHBRIDGE);                                    \
	    git clone https://github.com/OffchainLabs/$(VALIDATOR).git  \
	        $(BIN)/$(VALIDATOR);                                    \
	    git clone https://github.com/OffchainLabs/arb-avm.git       \
	        $(BIN)/$(VALIDATOR)/arb-avm;                            \
	    git clone https://github.com/OffchainLabs/$(TDD).git        \
	        $(BIN)/$(TDD);                                          \
	    git clone https://github.com/OffchainLabs/$(ATP).git        \
	        $(BIN)/$(TDD)/$(ATP);                                   \
	fi
	cp docker-compose.yml $(BIN) || true
	rm -f $(BIN)/$(DEMO)/Makefile
