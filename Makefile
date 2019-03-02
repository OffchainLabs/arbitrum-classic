### --------------------------------------------------------------------
### Makefile
### Runs Dockerfile. Puts repos in the $(BIN) directory.
### --------------------------------------------------------------------

# Repository names
DEMO=demo-app
COMPILER=arbc-solidity
VERIFIER=arbitrum-verifier-solidity
CLIENT=arbitrum-go

.PHONY: all clean src
BIN=deploy

all: src
	sudo docker build -t $(DEMO) $(BIN)
	sudo docker run  -p 80:8080 -it $(DEMO)

clean:
	rm -rf $(BIN)

# Download all source repos from git if $(BIN) does not exist
src:
	if [ ! -d "$(BIN)" ]; then                                  \
	    mkdir $(BIN);                                           \
	    mkdir $(BIN)/$(DEMO);                                   \
	    cp -r * $(BIN)/$(DEMO) || true;                         \
	    git clone git@github.com:OffchainLabs/$(COMPILER).git   \
	        $(BIN)/$(COMPILER);                                 \
	    git clone git@github.com:OffchainLabs/$(VERIFIER).git   \
	        $(BIN)/$(VERIFIER);                                 \
	    git clone git@github.com:OffchainLabs/$(CLIENT).git     \
	        $(BIN)/$(CLIENT);                                   \
	fi
	cp Dockerfile $(BIN) || true
