import { runTypeChain, glob } from 'typechain'
import { getPackagePath } from './util'
import { execSync } from 'child_process'

async function main() {
  const cwd = process.cwd()

  const arbosPath = getPackagePath('arbos-precompiles')
  const ethBridgePath = getPackagePath('arb-bridge-eth')
  const peripheralsPath = getPackagePath('arb-bridge-eth')

  console.log('Compiling paths.')
  console.log('building arbos')
  const stdout1 = execSync(`cd ${arbosPath} && yarn build`)
  console.log('building ethbridge')
  const stdout2 = execSync(`cd ${ethBridgePath} && yarn build`)
  console.log('building peripherals')
  const stdout3 = execSync(`cd ${peripheralsPath} && yarn build`)

  console.log('Done compiling')

  const allFiles = glob(cwd, [
    `${getPackagePath(
      'arbos-precompiles'
    )}/artifacts/!(build-info)/**/builtin/**/+([a-zA-Z0-9_]).json`,
    `${getPackagePath(
      `arb-bridge-eth`
    )}/build/contracts/!(build-info)/**/+([a-zA-Z0-9_]).json`,
    `${getPackagePath(
      `arb-bridge-peripherals`
    )}/build/contracts/!(build-info)/**/+([a-zA-Z0-9_]).json`,
  ])

  const result = await runTypeChain({
    cwd,
    filesToProcess: allFiles.filter(curr => !curr.includes('test')),
    allFiles,
    outDir: './src/lib/abi/',
    target: 'ethers-v5',
  })

  console.log('Typechain generated')
}

main()
  .then(() => console.log('Done.'))
  .catch(console.error)
