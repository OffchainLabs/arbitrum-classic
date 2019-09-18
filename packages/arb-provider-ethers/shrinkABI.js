const fs = require('fs');

function shrinkABI(jsonPath) {
    const jsonFile = require(jsonPath);
    fs.writeFileSync(
        jsonPath,
        JSON.stringify({
            contractName: jsonFile.contractName,
            abi: jsonFile.abi,
        }),
    );
}

shrinkABI('./src/lib/ArbChain.json');
shrinkABI('./src/lib/ArbChannel.json');
shrinkABI('./src/lib/GlobalPendingInbox.json');
