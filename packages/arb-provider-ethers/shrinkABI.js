const fs = require('fs');

const vmTrackerPath = './src/lib/VMTracker.json';
const vmTrackerJson = require(vmTrackerPath);
fs.writeFileSync(
    vmTrackerPath,
    JSON.stringify({
        contractName: vmTrackerJson.contractName,
        abi: vmTrackerJson.abi,
    }),
);

const inboxPath = './src/lib/GlobalPendingInbox.json';
const inboxJson = require(inboxPath);
fs.writeFileSync(
    inboxPath,
    JSON.stringify({
        contractName: inboxJson.contractName,
        abi: inboxJson.abi,
    }),
);
