# run this from arbitrum project root
BRANCH_NOW=$(git branch --show-current)
echo "Currently on branch $BRANCH_NOW"
HASH_BEFORE=$1
HASH_AFTER=$2
OUTDIR=$3
echo "Comparing gas usage between $HASH_BEFORE and $HASH_AFTER. Store results to $OUTDIR."

# checkout before, run tests, create report
git checkout $HASH_BEFORE
cd ./packages/arb-bridge-eth
echo "Measuring gas usage for $HASH_BEFORE"
yarn hardhat clean
yarn hardhat compile
yarn test:gas:ci
cd ../..
mv gasReporterOutput.json $OUTDIR/gas-$HASH_BEFORE.json

# and after
git checkout $HASH_AFTER
cd ./packages/arb-bridge-eth
echo "Measuring gas usage for $HASH_AFTER"
yarn hardhat clean
yarn hardhat compile
yarn test:gas:ci
cd ../..
mv gasReporterOutput.json $OUTDIR/gas-$HASH_AFTER.json

# now compare by running output
git checkout $BRANCH_NOW
echo "Calculating comparison"
cd packages/tools
yarn dev:compare-gas --gasReport1=$OUTDIR/gas-$HASH_BEFORE.json --gasReport2=$OUTDIR/gas-$HASH_AFTER.json --outputFile=$OUTDIR/compare-results
cd ../..

