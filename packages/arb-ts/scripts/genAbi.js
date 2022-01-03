const { runTypeChain, glob } = require('typechain')
const { execSync } = require('child_process')

const getPackagePath = packageName => {
  const path = require.resolve(`${packageName}/package.json`)
  return path.substr(0, path.indexOf('package.json'))
}

async function main() {
  const cwd = process.cwd()

  const arbosPath = getPackagePath('arbos-precompiles')
  const ethBridgePath = getPackagePath('arb-bridge-eth')
  const peripheralsPath = getPackagePath('arb-bridge-peripherals')

  console.log('Compiling paths.')
  console.log('building arbos')
  const stdout1 = execSync(`cd ${arbosPath} && yarn build`)
  console.log('building ethbridge')
  const stdout2 = execSync(`cd ${ethBridgePath} && yarn build`)
  console.log('building peripherals')
  const stdout3 = execSync(`cd ${peripheralsPath} && yarn build`)

  console.log('Done compiling')

  const allFiles = glob(cwd, [
    `${arbosPath}/artifacts/!(build-info)/**/builtin/**/+([a-zA-Z0-9_]).json`,
    `${ethBridgePath}/build/contracts/!(build-info)/**/+([a-zA-Z0-9_]).json`,
    `${peripheralsPath}/build/contracts/!(build-info)/**/+([a-zA-Z0-9_]).json`,
  ])

  const files = allFiles.filter(curr => !curr.toLowerCase().includes('test'))

  const result = await runTypeChain({
    cwd,
    filesToProcess: files,
    allFiles: files,
    outDir: './src/lib/abi/',
    target: 'ethers-v5',
  })

  console.log('Typechain generated')
}

main()
  .then(() => console.log('Done.'))
  .catch(console.error)
