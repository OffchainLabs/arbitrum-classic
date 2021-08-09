import { artifacts } from 'hardhat'

import {
  getStorageLayout,
  assertUpgradeSafe,
  getContractVersion,
  assertStorageUpgradeSafe,
  solcInputOutputDecoder,
  validate,
  RunValidation,
} from '@openzeppelin/upgrades-core'

const oldContract = 'L1ERC20Gateway'
const newContract = 'L2ERC20Gateway'

const main = async () => {
  const validationContext = {} as RunValidation

  const contracts = (await artifacts.getAllFullyQualifiedNames()).filter(
    curr => curr.includes(oldContract) || curr.includes(newContract)
  )

  if (contracts.length < 1) {
    throw new Error("Can't find artifacts for contracts")
  }

  for (const contract of contracts) {
    const buildInfo = await artifacts.getBuildInfo(contract)
    if (buildInfo === undefined) {
      throw new Error(`Build info not found for contract ${contract}`)
    }
    const solcOutput = buildInfo.output
    const solcInput = buildInfo.input
    const decodeSrc = solcInputOutputDecoder(solcInput, solcOutput)
    Object.assign(validationContext, validate(solcOutput, decodeSrc))
  }

  console.log(`Looking at ${oldContract} and ${newContract}`)

  const oldVersion = getContractVersion(validationContext, oldContract)
  const newVersion = getContractVersion(validationContext, newContract)

  // verifies for errors such as setting arguments in constructors
  assertUpgradeSafe([validationContext], oldVersion, { kind: 'transparent' })
  assertUpgradeSafe([validationContext], newVersion, { kind: 'transparent' })

  const oldStorage = getStorageLayout(validationContext, oldVersion)
  const newStorage = getStorageLayout(validationContext, newVersion)

  // verifies that storage layouts match
  assertStorageUpgradeSafe(oldStorage, newStorage)

  console.log('Upgrade validation complete, all is good.')
}

main()
  .then(() => process.exit(0))
  .catch(error => {
    console.error(error)
    process.exit(1)
  })
