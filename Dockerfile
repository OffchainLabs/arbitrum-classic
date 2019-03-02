### --------------------------------------------------------------------
### Dockerfile
### Deploy demo app
### --------------------------------------------------------------------

FROM ubuntu:18.04

# Ubuntu Deps
RUN apt-get update
RUN apt-get install -y git python3 python3-pip npm golang go-dep wget

# Testnet Deps
RUN npm install -g ganache-cli truffle

# Install destinations
ENV HOME=/root
ENV DEMO=$HOME/demo-app
ENV COMPILER=$HOME/arbc-solidity
ENV VERIFIER=$HOME/arbitrum-verifier-solidity
ENV CLIENT=$HOME/src/arbitrum

# Copy source code
COPY demo-app $DEMO
COPY arbc-solidity $COMPILER
COPY arbitrum-verifier-solidity $VERIFIER
COPY arbitrum-go $CLIENT

# Verifier Deps
WORKDIR $VERIFIER
RUN npm install
RUN cd $VERIFIER && truffle compile

# Client Deps
ENV GOPATH=$HOME
WORKDIR $CLIENT/..
RUN dep init arbitrum
RUN go install arbitrum/cmd/managerServer

# Demo App Deps
WORKDIR $DEMO
RUN npm install

# Run Compiler
WORKDIR $COMPILER
RUN pip3 install -r requirements.txt
RUN python3 $COMPILER/compile.py $DEMO/fibonacci.sol \
    $DEMO/fibonacci.ao

# Mnemonic
ENV TRUFFLE_MNEMONIC=\
"jar deny prosper gasp flush glass core corn alarm treat leg smart"

# Run Testnet
WORKDIR $VERIFIER
RUN ganache-cli -p 7545 -m "$TRUFFLE_MNEMONIC" & \
    truffle migrate -- reset

# Run Manager
RUN go run $CLIENT/cmd/managerServer/managerServer.go \
    $VERIFIER/addresses.json $DEMO/fibonacci.ao &

# Run Website
CMD npm start --prefix $DEMO
EXPOSE 8080
