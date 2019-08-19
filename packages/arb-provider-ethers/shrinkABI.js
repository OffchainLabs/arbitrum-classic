const fs = require('fs');

const jsonPath = './src/lib/VMTracker.json';
const vmTrackerJson = require(jsonPath);
fs.writeFileSync(
    jsonPath,
    JSON.stringify({
        contractName: vmTrackerJson.contractName,
        abi: vmTrackerJson.abi,
    }),
);
