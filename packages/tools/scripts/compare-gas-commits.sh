# run this from arbitrum project root
BRANCH_NOW=$(git branch --show-current)
echo "Currently on branch $BRANCH_NOW"
HASH_BEFORE=$1
HASH_AFTER=$2
echo "Comparing gas usage between $HASH_BEFORE and $HASH_AFTER"

# checkout before, run tests, create report
git checkout $HASH_BEFORE
cd ./packages/arb-bridge-eth
echo "Measuring gas usage for $HASH_BEFORE"
yarn hardhat clean
yarn hardhat compile
yarn test:gas:ci
cd ../..
mv gasReporterOutput.json ./packages/tools/gas-$HASH_BEFORE.json

# and after
git checkout $HASH_AFTER
cd ./packages/arb-bridge-eth
echo "Measuring gas usage for $HASH_AFTER"
yarn hardhat clean
yarn hardhat compile
yarn test:gas:ci
cd ../..
mv gasReporterOutput.json ./packages/tools/gas-$HASH_AFTER.json

# now compare by running output
git checkout $BRANCH_NOW
echo "Calculating comparison"
cd packages/tools
yarn dev:compare-gas --gasReport1=./gas-$HASH_BEFORE.json --gasReport2=./gas-$HASH_AFTER.json --outputFile=./comparison-output.csv
cd ../..

