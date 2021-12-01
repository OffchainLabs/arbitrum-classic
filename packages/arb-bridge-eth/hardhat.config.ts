import { task } from 'hardhat/config'
import * as fs from 'fs'
import { config } from './hardhat.default-config'

try {
  require('arb-upgrades/ethBridgeTasks')
} catch (e) {
  // arb-upgrades dependency not available
}

task('create-chain', 'Creates a rollup chain')
  .addParam('sequencer', "The sequencer's address")
  .setAction(async (taskArguments, hre) => {
    const machineHash = fs.readFileSync('../MACHINEHASH').toString()
    console.log(
      `Creating chain for machine with hash ${machineHash} for sequencer ${taskArguments.sequencer}`
    )
    const { deployments, ethers } = hre
    const [deployer] = await ethers.getSigners()
    const rollupCreatorDep = await deployments.get('RollupCreator')
    const RollupCreator = await ethers.getContractFactory('RollupCreator')
    const rollupCreator = RollupCreator.attach(
      rollupCreatorDep.address
    ).connect(deployer)
    const tx = await rollupCreator.createRollup(
      machineHash,
      900,
      0,
      2000000000,
      ethers.utils.parseEther('.1'),
      ethers.constants.AddressZero,
      await deployer.getAddress(),
      taskArguments.sequencer,
      300,
      1500,
      '0x'
    )
    const receipt = await tx.wait()
    const ev = rollupCreator.interface.parseLog(
      receipt.logs[receipt.logs.length - 1]
    )
    console.log(ev)

    // const path = `rollup-${hre.network.name}.json`
    const path = `rollup-${hre.network.name}.json`
    const output = JSON.stringify({
      rollupAddress: ev.args[0],
      inboxAddress: ev.args[1],
    })

    fs.writeFileSync(path, output)
    console.log(
      'New rollup chain created and output written to:',
      `${process.cwd()}:${path}`
    )
  })

task('deposit', 'Deposit coins into ethbridge')
  .addPositionalParam('inboxAddress', "The rollup chain's address")
  .addPositionalParam('privkey', 'The private key of the depositer')
  .addPositionalParam('dest', "The destination account's address")
  .addPositionalParam('amount', 'The amount to deposit')
  .setAction(async ({ inboxAddress, privkey, dest, amount }, bre) => {
    const { ethers } = bre
    const wallet = new ethers.Wallet(privkey, ethers.provider)
    const GlobalInbox = await ethers.getContractFactory('Inbox')
    const inbox = GlobalInbox.attach(inboxAddress).connect(wallet)
    await inbox.depositEth(dest, { value: amount })
  })

module.exports = config
