import { getPackagePath } from './util'
// this breaks on browsers, we can instead do a dynamic require to surface the error cleanly?
import { execSync } from 'child_process'

const cleanCompileContracts = async () => {
  const arbosPath = getPackagePath('arbos-precompiles')
  const ethBridgePath = getPackagePath('arb-bridge-eth')
  const peripheralsPath = getPackagePath('arb-bridge-eth')

  console.log('Clean building arbos')
  const stdout1 = execSync(`cd ${arbosPath} && yarn clean:build`)

  console.log('Clean building ethbridge')
  const stdout2 = execSync(`cd ${ethBridgePath} && yarn clean:build`)

  console.log('Clean building peripherals')
  const stdout3 = execSync(`cd ${peripheralsPath} && yarn clean:build`)

  console.log('All clean and all built.')
}

cleanCompileContracts()
  .then(() => process.exit(process.exitCode))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
