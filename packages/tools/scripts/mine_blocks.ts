import { findEnv } from '../find-env';
import { config } from "dotenv"; 
import {ethers} from "ethers";
config({ path: findEnv() });

const network:string = process.env['DEPLOY_ON'] || '';
const network_url:string = process.env[network.toUpperCase() + "_NETWORK"] || '';

const provider = new ethers.providers.JsonRpcProvider(network_url);

async function mineBlocks(numOfBlocks: number) {
    while (numOfBlocks > 0) {
      await provider.send("evm_mine", []);
      numOfBlocks--;
    }
}

if (require.main === module) {
  const numBlocks = Number(process.argv[2]) || 10;
  console.log(`Mining ${numBlocks} blocks`)
  mineBlocks(numBlocks)
}