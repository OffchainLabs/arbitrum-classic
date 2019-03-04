### --------------------------------------------------------------------
### Dockerfile
### Deploy demo app
### --------------------------------------------------------------------

# base
FROM ubuntu:18.04 as base
RUN apt-get update && apt-get install -y \
    git golang go-dep npm python3 python3-pip wget && \

# Files
ENV DEMO=demo-app
ENV COMPILER=arbc-solidity
ENV CLIENT=arbc-solidity
# Paths
ENV GOPATH=$HOME

# demo-app
FROM node:alpine as demo
COPY $DEMO $DEMO
WORKDIR $DEMO
RUN npm install

# arbitrum-verifier-solidity
FROM node:alpine as verifier
WORKDIR $VERIFIER
COPY $VERIFIER/package.json $VERIFIER/
RUN npm install
RUN truffle compile # TODO truffle compile

# arbitrum-go
FROM golang:alpine as client
COPY $CLIENT src/arbitrum
WORKDIR $HOME/src
RUN dep init arbitrum
RUN go install arbitrum/cmd/managerServer

# arbc-solidity
FROM base as compiler
COPY $COMPILER $HOME/$COMPILER
WORKDIR $HOME/$COMPILER
RUN pip3 install -r requirements.txt
RUN python3 $HOME/$COMPILER/compile.py $HOME/$DEMO/fibonacci.sol \
    $HOME/$COMPILER/fibonacci.ao

# main
FROM base as main
COPY $VERIFIER $HOME/$VERIFIER
COPY --from=verifier $HOME/$VERIFIER/node_modules: $HOME/$VERIFIER
COPY --from=client $HOME/$CLIENT $HOME/$CLIENT
COPY --from=demo $HOME/$DEMO $HOME/$DEMO
COPY --from=compiler $HOME/$COMPILER $HOME/$COMPILER
# Run Verifier, Manager, and Demo
CMD ganache-cli -p 7545 -m "$TRUFFLE_MNEMONIC" & \
    cd $VERIFIER && truffle migrate -- reset && \
    go run $CLIENT/cmd/managerServer/managerServer.go \
        $VERIFIER/addresses.json $COMPILER/fibonacci.ao & \
    sleep 25 && \
    npm start --prefix $DEMO
EXPOSE 8080 7545
