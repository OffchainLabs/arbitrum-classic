import * as ethers from 'ethers'
import { EventFragment } from '@ethersproject/abi'
import { ArbOwner__factory } from 'arb-ts/dist/lib/abi/factories/ArbOwner__factory'
import * as yargs from 'yargs'
import * as fs from 'fs-extra'
import * as upgrade from '../../arb-os/arb_os/upgrade.json'

const provider = new ethers.providers.JsonRpcProvider('http://localhost:8547')
const wallet = provider.getSigner(0)

async function updateArbOS(): Promise<void> {
  const batchedUpgrades: string[] = ['0x']
  const maxSegmentLength = 200000
  for (const insn of upgrade.instructions) {
    if (batchedUpgrades[batchedUpgrades.length - 1].length > maxSegmentLength) {
      batchedUpgrades.push('0x')
    }
    batchedUpgrades[batchedUpgrades.length - 1] += insn
  }

  console.log(`Upgrade split into ${batchedUpgrades.length} segments`)
  const arbOwner = ArbOwner__factory.connect(
    '0x000000000000000000000000000000000000006B',
    wallet
  )
  await arbOwner.startCodeUpload()
  for (const batch of batchedUpgrades) {
    await arbOwner.continueCodeUpload(batch)
  }
  await arbOwner.finishCodeUploadAsArbosUpgrade({ gasLimit: 1000000000000 })
}

if (require.main === module) {
  yargs.command(
    'init [rollup] [ethurl]',
    'initialize validators for the given rollup chain',
    yargsBuilder => yargsBuilder.options({}),
    args => {
      updateArbOS()
    }
  ).argv
}
