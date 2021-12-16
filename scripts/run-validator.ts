import { config } from "dotenv"; 
config({ path: ".env" });
const fs = require('fs')

const {resolve} = require("path");
const network:string = process.env['DEPLOY_ON'] || '';
const valDetails = require(`../rollups/${network}/validator1/config.json`);
const walletPath = resolve(`rollups/${network}/validator1/wallets`);
const arbNodePath = resolve(`packages/arb-node-core/cmd/arb-validator`)


const command = `cd ${arbNodePath} && go run arb-validator.go \
                --l1.url=${valDetails.eth_url} \
                --rollup.address=${valDetails.rollup_address} \
                --bridge-utils-address=${valDetails.bridge_utils_address} \
                --validator.strategy=MakeNodes \
                --wallet.local.pathname=${walletPath} \
                --wallet.local.password=${valDetails.password} \ 
                --validator.utils-address=${valDetails.validator_utils_address} \
                --validator.wallet-factory=${valDetails.validator_wallet_factory_address}`

fs.writeFile('./scripts/tmpVAL.sh', command, err => {
    if (err) {
        console.error(err)
        return
    }});   