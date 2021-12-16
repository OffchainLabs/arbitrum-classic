import { findEnv } from '../find-env';
import { config } from "dotenv"; 
import {ethers} from "ethers";
config({ path: findEnv() });

const {resolve} = require("path");

const network:string = process.env['DEPLOY_ON'] || '';
const networkUrl:string = process.env[network.toUpperCase() + "_NETWORK"] || '';
const pvtKey = process.env[network.toUpperCase() + "_PRIVATE_KEY"] || '';

const inboxAddress = require(`../../../rollups/${network}/validator0/config.json`)['inbox_address'];

const yargs = require('yargs');

// argument parser
const argv = yargs.options({
    wallet: {
      type: 'string',
      alias: 'w',
      demandOption: true,
      description: 'wallet address of the account to be funded on L2'
    },
    amount: {
        type: 'string',
        alias: 'a',
        demandOption: true,
        description: 'amount to be funded'
      }
}).argv;

async function fundWallet(
    destAddress:string, 
    amount:string, 
    inboxAddress: string, 
    faucetKey: 
    string, 
    network: string){
    
        amount = ethers.utils.parseEther(amount).toString();
        const inboxAbi = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"messageNum","type":"uint256"},{"indexed":false,"internalType":"bytes","name":"data","type":"bytes"}],"name":"InboxMessageDelivered","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"messageNum","type":"uint256"}],"name":"InboxMessageDeliveredFromOrigin","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bool","name":"enabled","type":"bool"}],"name":"PauseToggled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"bool","name":"enabled","type":"bool"}],"name":"RewriteToggled","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"newSource","type":"address"}],"name":"WhitelistSourceUpdated","type":"event"},{"inputs":[],"name":"bridge","outputs":[{"internalType":"contract IBridge","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"destAddr","type":"address"},{"internalType":"uint256","name":"l2CallValue","type":"uint256"},{"internalType":"uint256","name":"maxSubmissionCost","type":"uint256"},{"internalType":"address","name":"excessFeeRefundAddress","type":"address"},{"internalType":"address","name":"callValueRefundAddress","type":"address"},{"internalType":"uint256","name":"maxGas","type":"uint256"},{"internalType":"uint256","name":"gasPriceBid","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"createRetryableTicket","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"address","name":"destAddr","type":"address"},{"internalType":"uint256","name":"l2CallValue","type":"uint256"},{"internalType":"uint256","name":"maxSubmissionCost","type":"uint256"},{"internalType":"address","name":"excessFeeRefundAddress","type":"address"},{"internalType":"address","name":"callValueRefundAddress","type":"address"},{"internalType":"uint256","name":"maxGas","type":"uint256"},{"internalType":"uint256","name":"gasPriceBid","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"createRetryableTicketNoRefundAliasRewrite","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"maxSubmissionCost","type":"uint256"}],"name":"depositEth","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"contract IBridge","name":"_bridge","type":"address"},{"internalType":"address","name":"_whitelist","type":"address"}],"name":"initialize","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"isCreateRetryablePaused","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"isMaster","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"pauseCreateRetryables","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"maxGas","type":"uint256"},{"internalType":"uint256","name":"gasPriceBid","type":"uint256"},{"internalType":"address","name":"destAddr","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"sendContractTransaction","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"maxGas","type":"uint256"},{"internalType":"uint256","name":"gasPriceBid","type":"uint256"},{"internalType":"address","name":"destAddr","type":"address"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"sendL1FundedContractTransaction","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"uint256","name":"maxGas","type":"uint256"},{"internalType":"uint256","name":"gasPriceBid","type":"uint256"},{"internalType":"uint256","name":"nonce","type":"uint256"},{"internalType":"address","name":"destAddr","type":"address"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"sendL1FundedUnsignedTransaction","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"payable","type":"function"},{"inputs":[{"internalType":"bytes","name":"messageData","type":"bytes"}],"name":"sendL2Message","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"bytes","name":"messageData","type":"bytes"}],"name":"sendL2MessageFromOrigin","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"maxGas","type":"uint256"},{"internalType":"uint256","name":"gasPriceBid","type":"uint256"},{"internalType":"uint256","name":"nonce","type":"uint256"},{"internalType":"address","name":"destAddr","type":"address"},{"internalType":"uint256","name":"amount","type":"uint256"},{"internalType":"bytes","name":"data","type":"bytes"}],"name":"sendUnsignedTransaction","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"shouldRewriteSender","outputs":[{"internalType":"bool","name":"","type":"bool"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"startRewriteAddress","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"stopRewriteAddress","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"unpauseCreateRetryables","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newSource","type":"address"}],"name":"updateWhitelistSource","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"whitelist","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
        const provider =  new ethers.providers.JsonRpcProvider(network);
        const contract = new ethers.Contract(inboxAddress, inboxAbi);
        const wallet = new ethers.Wallet(faucetKey, provider);
        console.log("Depositing money...")
        await contract.connect(wallet).depositEth(destAddress, { value: amount })
        console.info("Deposited... will take some time to reflect on L2 chain")    
}

if (require.main === module){ 
    console.log('-----------------------------------------')
    console.info("Network:", network)
    console.info("Destination wallet address:", argv.wallet);
    console.info("Amount to deposit:", argv.amount);
    console.log('-----------------------------------------')

    fundWallet(
        argv.wallet, 
        argv.amount, 
        inboxAddress, 
        pvtKey, 
        networkUrl);
}