const { runTypeChain, glob } = require('typechain')
const { execSync } = require('child_process')
const { unlinkSync } = require('fs')

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

  const npmExec = process.env['npm_execpath']
  if (!npmExec || npmExec === '')
    throw new Error(
      'No support for npm_execpath env variable in package manager'
    )

  // https://yarnpkg.com/advanced/rulebook#packages-should-never-write-inside-their-own-folder-outside-of-postinstall
  // instead of writing in postinstall in each of those packages, we target a local folder in arb-ts' postinstall
  const artifactsTarget = `${cwd}/contract-artifacts`
  const envConfig = { ...process.env, ARTIFACT_PATH: artifactsTarget }

  console.log('building arbos')
  execSync(`${npmExec} run hardhat:prod compile`, {
    cwd: arbosPath,
    env: envConfig,
  })

  console.log('building ethbridge')
  execSync(`${npmExec} run hardhat:prod compile`, {
    cwd: ethBridgePath,
    env: envConfig,
  })

  console.log('building peripherals')
  execSync(`${npmExec} run hardhat:prod compile`, {
    cwd: peripheralsPath,
    env: envConfig,
  })

  console.log('Done compiling')

  const allFiles = glob(cwd, [
    `${arbosPath}/artifacts/!(build-info)/**/builtin/**/+([a-zA-Z0-9_]).json`,
    `${ethBridgePath}/build/contracts/!(build-info)/**/+([a-zA-Z0-9_]).json`,
    `${peripheralsPath}/build/contracts/!(build-info)/**/+([a-zA-Z0-9_]).json`,
  ])

  // TODO: remove test files
  await runTypeChain({
    cwd,
    filesToProcess: allFiles,
    allFiles: allFiles,
    outDir: './src/lib/abi/',
    target: 'ethers-v5',
  })

  // we delete the index file since it doesn't play well with tree shaking
  unlinkSync(`${cwd}/src/lib/abi/index.ts`)

  console.log('Typechain generated')
}

main()
  .then(() => console.log('Done.'))
  .catch(console.error)
