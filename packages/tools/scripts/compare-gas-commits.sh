# run this from arbitrum project root
HASH_NOW=$(git rev-parse HEAD)
HASH_BEFORE=
HASH_AFTER=

# checkout before, run tests, create report
git checkout HASH_BEFORE
cd ./packages/arb-eth-bridge
yarn test:gas:ci
cd ../..
mv gasReporterOutput.json ./packages/tools/gas-$HASH_BEFORE.json

# and after
git checkout HASH_AFTER
cd ./packages/arb-eth-bridge
yarn test:gas:ci
cd ../..
mv gasReporterOutput.json ./packages/tools/gas-$HASH_AFTER.json

# now compare by running output
echo "./gas-$HASH_BEFORE.json ./gas-$HASH_AFTER.json ./comparison-output.csv"
cd packages/tools
yarn dev:compare-gas -- "./gas-$HASH_BEFORE.json ./gas-$HASH_AFTER.json ./comparison-output.csv"


