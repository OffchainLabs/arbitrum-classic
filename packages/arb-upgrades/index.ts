import { HardhatRuntimeEnvironment } from 'hardhat/types/runtime'
import { writeFileSync, readFileSync } from 'fs'
import childProcess from 'child_process'
import {
  QueuedUpdates,
  CurrentDeployments,
  ContractNames,
  CurrentDeployment,
  QueuedUpdate,
} from './types'

// import {
//   getStorageLayout,
//   assertUpgradeSafe,
//   getContractVersion,
//   assertStorageUpgradeSafe,
//   solcInputOutputDecoder,
//   validate,
//   RunValidation,
// } from '@openzeppelin/upgrades-core'

const adminSlot =
  '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103'
const implementationSlot =
  '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'

const currentCommit = childProcess
  .execSync('git rev-parse HEAD')
  .toString()
  .trim()

const ensureCleanGitTree = () => {
  const tree = childProcess.execSync('git diff-index HEAD --').toString().trim()
  const found = tree.split('\n').find(diff => diff.includes('contracts'))
  if (found) {
    throw new Error(
      `You have local changes to ${found}. Commit/stash/ get rid of them`
    )
  }
}

export const initUpgrades = (
  hre: HardhatRuntimeEnvironment,
  rootDir: string
) => {
  const getQueuedUpdates = async (): Promise<{
    path: string
    data: QueuedUpdates
  }> => {
    const network = await hre.ethers.provider.getNetwork()
    const path = `${rootDir}/deployments/${network.chainId}_queued-updates.json`
    try {
      const jsonBuff = readFileSync(path)
      return { path, data: JSON.parse(jsonBuff.toString()) as QueuedUpdates }
    } catch (err) {
      if (err.code === 'ENOENT') {
        console.log('New network; creating queued updates file')
        writeFileSync(path, JSON.stringify({}))
        return { path, data: {} }
      } else {
        throw err
      }
    }
  }

  const getDeployments = async (): Promise<{
    path: string
    data: CurrentDeployments
  }> => {
    const network = await hre.ethers.provider.getNetwork()
    const path = `${rootDir}/deployments/${network.chainId}_current_deployment.json`
    try {
      const jsonBuff = readFileSync(path)
      return {
        path,
        data: JSON.parse(jsonBuff.toString()) as CurrentDeployments,
      }
    } catch (err) {
      if (err.code === 'ENOENT') {
        console.log(
          'New network; need to set up _current_deployments.json file'
        )
        throw err
      } else {
        throw err
      }
    }
  }

  const getBuildInfoString = async (contractName: string) => {
    const contracts = (await hre.artifacts.getAllFullyQualifiedNames()).filter(
      curr => curr.includes(contractName)
    )
    if (contracts.length !== 1) throw new Error('Contract not found')
    const info = await hre.artifacts.getBuildInfo(contracts[0])
    return JSON.stringify(info)
  }

  const deployLogic = async (
    contractNames: ContractNames[] | ContractNames
  ) => {
    ensureCleanGitTree()

    if (!Array.isArray(contractNames)) {
      contractNames = [contractNames]
    }
    console.log('Deploying logic contracts for ', contractNames.join(','))
    const signers = await hre.ethers.getSigners()
    if (!signers.length) {
      throw new Error(
        'No signer - make sure a key is properly set (check hardhat config)'
      )
    }

    const signer = signers[0]
    console.log('Using signer', signer.address)

    const { path, data: queuedUpdatesData } = await getQueuedUpdates()

    for (const contractName of contractNames) {
      console.log('Deploying new logic for ', contractName)

      const contractFactory = (
        await hre.ethers.getContractFactory(contractName)
      ).connect(signer)

      const newLogic = await contractFactory.deploy()
      const deployedContract = await newLogic.deployed()
      const receipt = await deployedContract.deployTransaction.wait()

      const newLogicData: QueuedUpdate = {
        address: receipt.contractAddress,
        deployTxn: receipt.transactionHash,
        arbitrumCommitHash: currentCommit,
        buildInfo: await getBuildInfoString(contractName),
      }
      queuedUpdatesData[contractName] = newLogicData
      console.log(`Deployed ${contractName} Logic:`)
      console.log(newLogicData)
      console.log('')

      writeFileSync(path, JSON.stringify(queuedUpdatesData))
    }
  }

  const updateImplementations = async () => {
    ensureCleanGitTree()
    const res = await verifyCurrentImplementations()
    if (!res) {
      throw new Error(
        'Verification of current implementations failed; cancelling update'
      )
    }

    const { path: queuedUpdatesPath, data: queuedUpdatesData } =
      await getQueuedUpdates()
    const { path: deploymentsPath, data: deploymentsJsonData } =
      await getDeployments()

    const { proxyAdminAddress } = deploymentsJsonData

    const signers = await hre.ethers.getSigners()
    if (!signers.length) {
      throw new Error(
        'No signer - make sure a key is properly set (check hardhat config)'
      )
    }
    const signer = signers[0]

    const ProxyAdmin__factory = await hre.ethers.getContractFactory(
      'ProxyAdmin'
    )
    const proxyAdmin =
      ProxyAdmin__factory.attach(proxyAdminAddress).connect(signer)

    const proxyAdminOwner = await proxyAdmin.owner()
    if (proxyAdminOwner.toLowerCase() !== signer.address.toLowerCase()) {
      throw new Error(
        `Signer address ${signer.address} != ProxyAdmin owner ${proxyAdminOwner}`
      )
    }

    const contractsToUpdate = Object.keys(queuedUpdatesData) as ContractNames[]
    console.log(`Updating ${contractsToUpdate.length} contracts`)

    for (const contractName of contractsToUpdate) {
      const queuedUpdateData = queuedUpdatesData[contractName] as QueuedUpdate
      const deploymentData = deploymentsJsonData.contracts[
        contractName
      ] as CurrentDeployment
      if (!deploymentData) {
        console.warn(`Contract ${contractName} not recognized; skipping`)
        continue
      }
      console.log(`Updating ${contractName} to new implementation`)

      const upgradeTx = await proxyAdmin.upgrade(
        deploymentData.proxyAddress,
        queuedUpdateData.address
      )

      await upgradeTx.wait()

      const buildInfo = await getBuildInfoString(contractName)

      console.log(`Done updating ${contractName}`)
      const newDeploymentData: CurrentDeployment = {
        proxyAddress: deploymentData.proxyAddress,
        implAddress: queuedUpdateData.address,
        implDeploymentTxn: queuedUpdateData.deployTxn,
        implArbitrumCommitHash: queuedUpdateData.arbitrumCommitHash,
        implBuildInfo: buildInfo,
      }
      console.log('Setting new deployment data:', newDeploymentData)

      deploymentsJsonData.contracts[contractName] = newDeploymentData
      writeFileSync(deploymentsPath, JSON.stringify(deploymentsJsonData))

      delete queuedUpdatesData[contractName]
      writeFileSync(queuedUpdatesPath, JSON.stringify(queuedUpdatesData))
      console.log('')
    }
    return await verifyCurrentImplementations()
  }

  const verifyCurrentImplementations = async () => {
    console.log('Verifying implementations')

    const { data: deploymentsJsonData } = await getDeployments()
    let success = true
    for (const _contractName in deploymentsJsonData.contracts) {
      const contractName = _contractName as ContractNames
      const deploymentData = deploymentsJsonData.contracts[contractName]

      // check proxy admin
      let admin = await hre.ethers.provider.getStorageAt(
        deploymentData.proxyAddress,
        adminSlot
      )
      if (admin.length > 42) {
        admin = '0x' + admin.substr(admin.length - 40, 40)
      }
      if (
        admin.toLowerCase() !==
        deploymentsJsonData.proxyAdminAddress.toLowerCase()
      ) {
        console.log(
          'Verification failed: bad admin',
          admin,
          deploymentsJsonData.proxyAdminAddress
        )
        success = false
      }

      //  check implementation
      let implementation = await hre.ethers.provider.getStorageAt(
        deploymentData.proxyAddress,
        implementationSlot
      )
      if (implementation.length > 42) {
        implementation =
          '0x' + implementation.substr(implementation.length - 40, 40)
      }
      if (
        implementation.toLowerCase() !==
        deploymentData.implAddress.toLowerCase()
      ) {
        console.log('Verification failed; bad implementation', implementation)
        success = false
      }
    }
    console.log(success ? 'Verified successfully :)' : 'Failed verification :/')

    return success
  }

  // TODO:
  // const verifyUpgrade = async (contractName: string) => {
  //    const validationContext = {} as RunValidation

  //    const contractNameFull = (await artifacts.getAllFullyQualifiedNames()).find(
  //      curr => curr.includes(contractName)
  //    )

  //    if (!contractNameFull ) {
  //      throw new Error("Can't find artifacts for contracts")
  //    }

  //    for (const contract of contractNameFull) {
  //      const buildInfo = await artifacts.getBuildInfo(contract)

  //      const solcOutput = buildInfo.output
  //      const solcInput = buildInfo.input
  //      const decodeSrc = solcInputOutputDecoder(solcInput, solcOutput)
  //      Object.assign(validationContext, validate(solcOutput, decodeSrc))
  //    }

  //    const oldVersion = getContractVersion(validationContext, oldContract)
  //    const newVersion = getContractVersion(validationContext, newContract)

  //    // verifies for errors such as setting arguments in constructors
  //    assertUpgradeSafe([validationContext], oldVersion, { kind: 'transparent' })
  //    assertUpgradeSafe([validationContext], newVersion, { kind: 'transparent' })

  //    const oldStorage = getStorageLayout(validationContext, oldVersion)
  //    const newStorage = getStorageLayout(validationContext, newVersion)

  //    // verifies that storage layouts match
  //    assertStorageUpgradeSafe(oldStorage, newStorage)

  //    console.log('Upgrade validation complete, all is good.')
  //  }

  return {
    updateImplementations,
    verifyCurrentImplementations,
    deployLogic,
  }
}
