import { config } from "dotenv"; 
config({ path: ".env" });
const fs = require('fs')

const {resolve} = require("path");
const network:string = process.env['DEPLOY_ON'] || '';
const valDetails = require(`../rollups/${network}/validator0/config.json`);
const walletPath = resolve(`rollups/${network}/validator0/wallets`);
const arbNodePath = resolve(`packages/arb-rpc-node/cmd/arb-node`)


const command = `cd ${arbNodePath} && go run arb-node.go \
                --l1.url=${valDetails.eth_url} \
                --rollup.address=${valDetails.rollup_address} \
                --bridge-utils-address=${valDetails.bridge_utils_address} \
                --node.type=sequencer \
                --wallet.local.pathname=${walletPath} \
                --wallet.local.password=${valDetails.password}`

fs.writeFile('./scripts/tmpRPC.sh', command, err => {
    if (err) {
        console.error(err)
        return
    }});   