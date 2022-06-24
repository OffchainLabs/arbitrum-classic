import { JsonRpcProvider } from '@ethersproject/providers'
import { Contract } from 'ethers'
import ArbOwnerAbi from './ArbOwner.json'
import * as yargs from 'yargs'
import * as upgrade from '../../arb-os/arb_os/upgrade.json'

const provider = new JsonRpcProvider('http://localhost:8547')
const wallet = provider.getSigner(0)

async function updateArbOS(
  newCodeHash: string,
  oldCodeHash: string
): Promise<void> {
  const batchedUpgrades: string[] = ['0x']
  const maxSegmentLength = 200000
  for (const insn of upgrade.instructions) {
    if (batchedUpgrades[batchedUpgrades.length - 1].length > maxSegmentLength) {
      batchedUpgrades.push('0x')
    }
    batchedUpgrades[batchedUpgrades.length - 1] += insn
  }

  console.log(`Upgrade split into ${batchedUpgrades.length} segments`)
  const arbOwner = new Contract(
    '0x000000000000000000000000000000000000006B',
    ArbOwnerAbi
  ).connect(wallet)
  await arbOwner.startCodeUpload()
  for (const batch of batchedUpgrades) {
    await arbOwner.continueCodeUpload(batch)
  }
  await arbOwner.finishCodeUploadAsArbosUpgrade(newCodeHash, oldCodeHash, {
    gasLimit: 1000000000000,
  })
}

if (require.main === module) {
  yargs.command(
    'init [rollup] [ethurl]',
    'initialize validators for the given rollup chain',
    yargsBuilder =>
      yargsBuilder.options({
        newCodeHash: {
          type: 'string',
          required: true,
        },
        oldCodeHash: {
          type: 'string',
          required: true,
        },
      }),
    args => {
      updateArbOS(args.newCodeHash, args.oldCodeHash)
    }
  ).argv
}
