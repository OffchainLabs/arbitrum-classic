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
	time sudo docker-compose build -f demo-app/docker-compose.yml
	sudo docker-compose run -f demo-app/docker-compose.yml

clean:
	rm -rf $(BIN)

# Download all source repos from git if $(BIN) does not exist
src:
	if [ ! -d "$(BIN)" ]; then                                  \
	    mkdir $(BIN);                                           \
	    mkdir $(BIN)/$(DEMO);                                   \
	    cp -r * $(BIN)/$(DEMO) || true;                         \
	    git clone git@github.com:OffchainLabs/$(AWP).git        \
	        $(BIN)/$(DEMO)/$(AWP);                              \
	    git clone git@github.com:OffchainLabs/$(COMPILER).git   \
	        $(BIN)/$(COMPILER);                                 \
	    git clone git@github.com:OffchainLabs/$(ETHBRIDGE).git  \
	        $(BIN)/$(ETHBRIDGE);                                \
	    git clone git@github.com:OffchainLabs/$(VALIDATOR).git  \
	        $(BIN)/$(VALIDATOR);                                \
	    git clone git@github.com:OffchainLabs/arb-avm.git       \
	        $(BIN)/$(VALIDATOR)/arb-avm;                        \
	    git clone git@github.com:OffchainLabs/$(TDD).git        \
	        $(BIN)/$(TDD);                                      \
	    git clone git@github.com:OffchainLabs/$(ATP).git        \
	        $(BIN)/$(TDD)/$(ATP);                               \
	fi
