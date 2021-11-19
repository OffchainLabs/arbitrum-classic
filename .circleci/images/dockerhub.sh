#!/bin/bash
### --------------------------------------------------------------------
### install-validator
### --------------------------------------------------------------------

# Exit on error
set -e

docker build -t python-base python-base
docker tag python-base:latest offchainlabs/python-base:0.1.0
docker push offchainlabs/python-base:0.1.0

docker build -t yarn-base yarn-base
docker tag yarn-base:latest offchainlabs/yarn-base:0.1.0
docker push offchainlabs/yarn-base:0.1.0

docker build -t go-base go-base
docker tag go-base:latest offchainlabs/go-base:0.1.0
docker push offchainlabs/go-base:0.1.0

docker build -t cpp-base cpp-base
docker tag cpp-base:latest offchainlabs/cpp-base:0.1.0
docker push offchainlabs/cpp-base:0.1.0

docker build -t cpp-tsan-base cpp-tsan-base
docker tag cpp-tsan-base:latest offchainlabs/cpp-tsan-base:0.1.0
docker push offchainlabs/cpp-tsan-base:0.1.0

docker build -t cpp-asan-base cpp-asan-base
docker tag cpp-asan-base:latest offchainlabs/cpp-asan-base:0.1.0
docker push offchainlabs/cpp-asan-base:0.1.0

docker build -t dist-base dist-base
docker tag dist-base:latest offchainlabs/dist-base:0.1.0
docker push offchainlabs/dist-base:0.1.0

docker build -t ethbridge-base ethbridge-base
docker tag ethbridge-base:latest offchainlabs/ethbridge-base:0.1.0
docker push offchainlabs/ethbridge-base:0.1.0

docker build -t backend-base backend-base
docker tag backend-base:latest offchainlabs/backend-base:0.1.0
docker push offchainlabs/backend-base:0.1.0
