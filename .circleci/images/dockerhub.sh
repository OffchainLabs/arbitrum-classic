#!/bin/bash
### --------------------------------------------------------------------
### install-validator
### --------------------------------------------------------------------

# Exit on error
set -e

# cpp-base depended on by everything else so build first
docker build -t offchainlabs/cpp-base:0.6.2 cpp-base
docker push offchainlabs/cpp-base:0.6.2

docker build -t offchainlabs/backend-base:0.6.2 backend-base
docker push offchainlabs/backend-base:0.6.2

docker build -t offchainlabs/cpp-asan-base:0.6.2 cpp-asan-base
docker push offchainlabs/cpp-asan-base:0.6.2

docker build -t offchainlabs/cpp-tsan-base:0.6.2 cpp-tsan-base
docker push offchainlabs/cpp-tsan-base:0.6.2

docker build -t offchainlabs/dist-base:0.6.2 dist-base
docker push offchainlabs/dist-base:0.6.2

docker build -t offchainlabs/frontend-base:0.6.2 frontend-base
docker push offchainlabs/frontend-base:0.6.2

docker build -t offchainlabs/integration-base:0.6.2 integration-base
docker push offchainlabs/integration-base:0.6.2

