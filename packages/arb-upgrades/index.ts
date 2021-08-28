import { HardhatRuntimeEnvironment } from 'hardhat/types/runtime'
import { writeFileSync, readFileSync, unlinkSync, existsSync } from 'fs'
import childProcess from 'child_process'
// @ts-ignore (module doesn't have types declared)
import prompt from 'prompt-promise'
import {
  QueuedUpdates,
  CurrentDeployments,
  ContractNames,
  CurrentDeployment,
  QueuedUpdate,
  isBeacon,
  isRollupUserFacet,
  isRollupAdminFacet,
  getLayer,
  hasPostInitHook,
  isBeaconOwnedByEOA,
  isBeaconOwnedByRollup,
} from './types'

const ADMIN_SLOT =
  '0xb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103'
const IMPLEMENTATION_SLOT =
  '0x360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc'
const BEACON_SLOT =
  '0xa3f0ad74e5423aebfd80d3ef4346578335a9a72aeaee59ff6cb3582b35133d50'

const getAdminFromProxyStorage = async (
  hre: HardhatRuntimeEnvironment,
  proxyAddress: string
) => {
  let admin = await hre.ethers.provider.getStorageAt(proxyAddress, ADMIN_SLOT)
  if (admin.length > 42) {
    admin = '0x' + admin.substr(admin.length - 40, 40)
  }
  return admin
}

const POST_UPGRADE_INIT_SIG = '0x95fcea78'

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
  const compileTask = hre.run('compile')

  const getQueuedUpdates = async (): Promise<{
    path: string
    data: QueuedUpdates
  }> => {
    const network = await hre.ethers.provider.getNetwork()
    const path = `${rootDir}/_deployments/${network.chainId}_queued-updates.json`
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
    const path = `${rootDir}/_deployments/${network.chainId}_current_deployment.json`
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
      }
      throw err
    }
  }

  const createOrLoadTmpDeploymentsFile = async (): Promise<{
    path: string
    data: CurrentDeployments
  }> => {
    const { data: currentDeployments } = await getDeployments()
    const val = await loadTmpDeployments()
    if (val) return val

    console.log('Creating a new tmp deployments file:')
    const path = await tmpDeploymentsPath()
    writeFileSync(path, JSON.stringify(currentDeployments))
    return {
      path,
      data: currentDeployments,
    }
  }
  const tmpDeploymentsPath = async () => {
    const network = await hre.ethers.provider.getNetwork()
    return `${rootDir}/_deployments/${network.chainId}_tmp_deployment.json`
  }

  const loadTmpDeployments = async (): Promise<
    | {
        path: string
        data: CurrentDeployments
      }
    | undefined
  > => {
    const path = await tmpDeploymentsPath()
    if (existsSync(path)) {
      console.log(
        `tmp deployments file found; do you want to resume deployments with it? ('Yes' to continue)`
      )
      const res = await prompt('')
      if (res !== 'Yes') {
        console.log('exiting')
        process.exit(0)
      }
      const jsonBuff = readFileSync(path)
      return {
        path,
        data: JSON.parse(jsonBuff.toString()) as CurrentDeployments,
      }
    }
  }
  const getBuildInfoString = async (contractName: string) => {
    const names = await hre.artifacts.getAllFullyQualifiedNames()
    const contracts = names.filter(curr => curr.endsWith(`:${contractName}`))
    if (contracts.length !== 1) throw new Error('Contract not found')
    const info = await hre.artifacts.getBuildInfo(contracts[0])
    return JSON.stringify(info)
  }

  const deployLogic = async (
    contractNames: ContractNames[] | ContractNames
  ) => {
    await compileTask
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
      if (queuedUpdatesData[contractName]) {
        console.log(
          `Update already queued up for ${contractName}; would you redeploy it? ('Yes' to redeploy, otherwise we'll skip and used the queued update)`
        )
        const res = await prompt('')
        if (res.trim().toLowerCase() !== 'yes') {
          console.log('Skipping redeploy and using the queued update')
          continue
        } else {
          console.log('Redeploying ', contractName)
        }
      }

      const layerOfContract = getLayer(contractName)
      const currentLayer =
        (
          await hre.ethers.provider.getCode(
            '0x0000000000000000000000000000000000000064'
          )
        ).length > 2
          ? 2
          : 1
      if (layerOfContract !== currentLayer) {
        throw new Error(
          `Warning: trying to deploy ${contractName} onto the wrong layer!`
        )
      }
      console.log('Deploying new logic for ', contractName)

      const contractFactory = (
        await hre.ethers.getContractFactory(contractName)
      ).connect(signer)

      // handle Rollup's constructor:
      const newLogic =
        contractName === ContractNames.Rollup
          ? await contractFactory.deploy(42161)
          : await contractFactory.deploy()
      const deployedContract = await newLogic.deployed()
      const receipt = await deployedContract.deployTransaction.wait()

      const newLogicData: QueuedUpdate = {
        address: receipt.contractAddress,
        deployTxn: receipt.transactionHash,
        arbitrumCommitHash: currentCommit,
        buildInfo: '' /* await getBuildInfoString(contractName) */,
      }
      queuedUpdatesData[contractName] = newLogicData
      console.log(`Deployed ${contractName} Logic:`)
      console.log(receipt)
      console.log('')

      writeFileSync(path, JSON.stringify(queuedUpdatesData))
    }
  }

  const updateImplementations = async () => {
    await compileTask
    ensureCleanGitTree()
    const res = await verifyCurrentImplementations()
    if (!res) {
      throw new Error(
        'Verification of current implementations failed; cancelling update'
      )
    }

    const { path: queuedUpdatesPath, data: queuedUpdatesData } =
      await getQueuedUpdates()
    const { path: deploymentsPath } = await getDeployments()
    const { path: tmpDeploymentsPath, data: tmpDeploymentsJsonData } =
      await createOrLoadTmpDeploymentsFile()

    const { proxyAdminAddress } = tmpDeploymentsJsonData

    const ProxyAdmin__factory = await hre.ethers.getContractFactory(
      'ProxyAdmin'
    )
    let proxyAdmin = ProxyAdmin__factory.attach(proxyAdminAddress).connect(
      hre.ethers.provider
    )

    const proxyAdminOwner = await proxyAdmin.owner()

    const getSigner = async (networkName: string) => {
      if (networkName === 'fork') {
        await hre.network.provider.request({
          method: 'hardhat_impersonateAccount',
          params: [proxyAdminOwner],
        })

        await hre.network.provider.send('hardhat_setBalance', [
          proxyAdminOwner,
          '0x16189AD417E380000',
        ])

        return hre.ethers.getSigner(proxyAdminOwner)
      } else {
        const signers = await hre.ethers.getSigners()
        if (!signers.length) {
          throw new Error(
            'No signer - make sure a key is properly set (check hardhat config)'
          )
        }
        return signers[0]
      }
    }
    const signer = await getSigner(hre.network.name)

    proxyAdmin = proxyAdmin.connect(signer)

    if (proxyAdminOwner.toLowerCase() !== signer.address.toLowerCase()) {
      throw new Error(
        `Signer address ${signer.address} != ProxyAdmin owner ${proxyAdminOwner}`
      )
    }

    const contractsToUpdate = Object.keys(queuedUpdatesData) as ContractNames[]
    if (contractsToUpdate.length === 0) {
      throw new Error(
        'No logic implementations to upgrade to for current network / package'
      )
    }
    console.log(`Updating ${contractsToUpdate.length} contracts`)
    // TODO: explicitly check for storage layout clashes

    contractsToUpdate.sort((a, b) =>
      a === ContractNames.SequencerInbox ? -1 : 1
    )

    for (const contractName of contractsToUpdate) {
      const queuedUpdateData = queuedUpdatesData[contractName] as QueuedUpdate
      const deploymentData = tmpDeploymentsJsonData.contracts[
        contractName
      ] as CurrentDeployment
      if (!deploymentData) {
        console.warn(`Contract ${contractName} not recognized; skipping`)
        continue
      }
      console.log(`Updating ${contractName} to new implementation`)

      let upgradeTx: any
      if (isBeaconOwnedByEOA(contractName)) {
        // handle UpgradeableBeacon proxy owned by EOA
        const UpgradeableBeacon = (
          await hre.ethers.getContractFactory('UpgradeableBeacon')
        )
          .attach(deploymentData.proxyAddress)
          .connect(signer)
        upgradeTx = await UpgradeableBeacon.upgradeTo(queuedUpdateData.address)
      } else if (isBeaconOwnedByRollup(contractName)) {
        // handle UpgradeableBeacon proxy owned by Rollup
        const rollupAddress =
          tmpDeploymentsJsonData.contracts.Rollup.proxyAddress
        const RollupAdmin = (
          await hre.ethers.getContractFactory(ContractNames.RollupAdminFacet)
        )
          .attach(rollupAddress)
          .connect(signer)
        upgradeTx = await RollupAdmin.upgradeBeacon(
          deploymentData.proxyAddress,
          queuedUpdateData.address
        )
      } else if (
        isRollupAdminFacet(contractName) ||
        isRollupUserFacet(contractName)
      ) {
        // Handle diamond proxy pattern
        const userFacetAddress = isRollupUserFacet(contractName)
          ? queuedUpdateData.address
          : tmpDeploymentsJsonData.contracts.RollupUserFacet.implAddress
        const adminFacetAddress = isRollupAdminFacet(contractName)
          ? queuedUpdateData.address
          : tmpDeploymentsJsonData.contracts.RollupAdminFacet.implAddress
        const RollupAdmin = (
          await hre.ethers.getContractFactory(ContractNames.RollupAdminFacet)
        )
          .attach(tmpDeploymentsJsonData.contracts.Rollup.proxyAddress)
          .connect(signer)
        upgradeTx = await RollupAdmin.setFacets(
          adminFacetAddress,
          userFacetAddress
        )
      } else {
        // handle TransparentUpgradeableProxy

        if (hasPostInitHook(contractName)) {
          upgradeTx = await proxyAdmin.upgradeAndCall(
            deploymentData.proxyAddress,
            queuedUpdateData.address,
            POST_UPGRADE_INIT_SIG
          )
        } else {
          upgradeTx = await proxyAdmin.upgrade(
            deploymentData.proxyAddress,
            queuedUpdateData.address
          )
        }
      }
      const rec = await upgradeTx.wait()
      console.log('Upgrade receipt:', rec)

      // const buildInfo = await getBuildInfoString(contractName)

      console.log(`Done updating ${contractName}`)
      const newDeploymentData: CurrentDeployment = {
        proxyAddress: deploymentData.proxyAddress,
        implAddress: queuedUpdateData.address,
        implDeploymentTxn: queuedUpdateData.deployTxn,
        implArbitrumCommitHash: queuedUpdateData.arbitrumCommitHash,
        implBuildInfo: '',
      }
      console.log('Setting new tmp: deployment data')

      tmpDeploymentsJsonData.contracts[contractName] = newDeploymentData
      writeFileSync(tmpDeploymentsPath, JSON.stringify(tmpDeploymentsJsonData))

      delete queuedUpdatesData[contractName]
      writeFileSync(queuedUpdatesPath, JSON.stringify(queuedUpdatesData))
      console.log('')
    }
    console.log('Finished all deployments: setting data to current deployments')
    writeFileSync(deploymentsPath, JSON.stringify(tmpDeploymentsJsonData))
    // removing tmp file
    unlinkSync(tmpDeploymentsPath)

    return await verifyCurrentImplementations()
  }

  const verifyCurrentImplementations = async () => {
    await compileTask
    console.log('Verifying deployments:')

    const { data: deploymentsJsonData } = await getDeployments()
    const tmpDeploymentsJsonData = await loadTmpDeployments()

    let success = true
    const ProxyAdmin__factory = await hre.ethers.getContractFactory(
      'ProxyAdmin'
    )
    const proxyAdmin = ProxyAdmin__factory.attach(
      deploymentsJsonData.proxyAdminAddress
    )
    const proxyAdminOwner = await proxyAdmin.owner()
    console.log('proxyAdmin owner:', proxyAdminOwner)

    for (const _contractName in deploymentsJsonData.contracts) {
      const contractName = _contractName as ContractNames
      const _currentDeploymentData = deploymentsJsonData.contracts[contractName]
      const _tmpDeploymentData =
        tmpDeploymentsJsonData &&
        tmpDeploymentsJsonData.data.contracts[contractName]

      const deploymentData = _tmpDeploymentData
        ? _tmpDeploymentData
        : _currentDeploymentData

      if (isBeacon(contractName)) {
        const UpgradeableBeacon = (
          await hre.ethers.getContractFactory('UpgradeableBeacon')
        ).attach(deploymentData.proxyAddress)

        const implementation = await UpgradeableBeacon.implementation()
        const beaconOwner = await UpgradeableBeacon.owner()
        if (
          implementation.toLowerCase() !==
          deploymentData.implAddress.toLowerCase()
        ) {
          console.log(
            contractName + ' Verification failed: bad implementation',
            implementation,
            deploymentData.implAddress
          )
          success = false
        }

        const expectedBeaconOwner = isBeaconOwnedByRollup(contractName)
          ? deploymentsJsonData.contracts.Rollup.proxyAddress
          : proxyAdminOwner

        if (beaconOwner.toLowerCase() !== expectedBeaconOwner.toLowerCase()) {
          console.log(
            `${contractName} Verification failed: bad admin`,
            beaconOwner,
            proxyAdminOwner
          )
          success = false
        }
        continue
      }

      if (isRollupAdminFacet(contractName) || isRollupUserFacet(contractName)) {
        const Rollup = (
          await hre.ethers.getContractFactory(ContractNames.Rollup)
        ).attach(deploymentsJsonData.contracts.Rollup.proxyAddress)
        const facet = isRollupUserFacet(contractName)
          ? await Rollup.getUserFacet()
          : await Rollup.getAdminFacet()
        if (facet.toLowerCase() !== deploymentData.implAddress.toLowerCase()) {
          console.log(
            `${contractName} Verification failed; bad implementation`,
            facet
          )
          success = false
        }
        continue
      }

      // check proxy admin
      const admin = await getAdminFromProxyStorage(
        hre,
        deploymentData.proxyAddress
      )
      if (
        admin.toLowerCase() !==
        deploymentsJsonData.proxyAdminAddress.toLowerCase()
      ) {
        console.log(
          `${contractName} Verification failed: bad admin`,
          admin,
          deploymentsJsonData.proxyAdminAddress
        )
        success = false
      }
      //  check implementation
      let implementation = await hre.ethers.provider.getStorageAt(
        deploymentData.proxyAddress,
        IMPLEMENTATION_SLOT
      )
      if (implementation.length > 42) {
        implementation =
          '0x' + implementation.substr(implementation.length - 40, 40)
      }
      if (
        implementation.toLowerCase() !==
        deploymentData.implAddress.toLowerCase()
      ) {
        console.log(
          `${contractName} Verification failed; bad implementation`,
          implementation,
          deploymentData.implAddress
        )
        success = false
      }
    }
    console.log(success ? 'Verified successfully :)' : 'Failed verification :/')

    return success
  }

  const deployLogicAll = async () => {
    await compileTask
    const { path: deploymentsPath, data: deploymentsJsonData } =
      await getDeployments()
    const contractsNames = Object.keys(
      deploymentsJsonData.contracts
    ) as ContractNames[]
    await deployLogic(contractsNames)
  }

  const transferAdmin = async (proxyAddress: string, newAdmin: string) => {
    await compileTask

    const proxyAdminAddr = await getAdminFromProxyStorage(hre, proxyAddress)
    const ProxyAdmin__factory = await hre.ethers.getContractFactory(
      'ProxyAdmin'
    )
    const proxyAdmin = ProxyAdmin__factory.attach(proxyAdminAddr).connect(
      hre.ethers.provider
    )
    const proxyAdminOwner = await proxyAdmin.owner()

    if (newAdmin.toLowerCase() === proxyAdminOwner.toLowerCase()) {
      throw new Error('User trying to update admin to current admin address')
    }

    const { path: deploymentsPath, data } = await getDeployments()
    const signers = await hre.ethers.getSigners()
    if (!signers.length) {
      throw new Error(
        'No signer - make sure a key is properly set (check hardhat config)'
      )
    }
    const signer = signers[0]

    if (signer.address.toLowerCase() !== proxyAdminOwner.toLowerCase()) {
      throw new Error('User signer is not the owner of proxy admin')
    }

    const res = await proxyAdmin
      .connect(signer)
      .changeProxyAdmin(proxyAddress, newAdmin)
    const rec = await res.wait()
  }

  const transferBeaconOwner = async (
    upgradableBeaconAddress: string,
    newOwner: string
  ) => {
    await compileTask
    const signers = await hre.ethers.getSigners()
    if (!signers.length) {
      throw new Error(
        'No signer - make sure a key is properly set (check hardhat config)'
      )
    }
    const signer = signers[0]
    const UpgradeableBeacon = (
      await hre.ethers.getContractFactory('UpgradeableBeacon')
    )
      .attach(upgradableBeaconAddress)
      .connect(signer)

    const beaconOwner = await UpgradeableBeacon.owner()
    if (beaconOwner.toLowerCase() !== signer.address.toLowerCase()) {
      throw new Error(
        `Not connecetd as owner ${beaconOwner}, instead running as ${signer.address}`
      )
    }

    console.log(
      `You are about to transfer owner ship of ${upgradableBeaconAddress} to ${newOwner}. You sure? ('Yes' to proceeed)`
    )

    const confirm = await prompt('')
    if (confirm !== 'Yes') {
      console.log('Cancelling')
      return
    }
    const res = await UpgradeableBeacon.transferOwnership(newOwner)
    const rec = await res.wait()
    console.log('ownership transfer complete')
  }

  const removeBuildInfoFiles = async () => {
    console.log(
      `You sure you want to remove build info files for the current network's current_deployments file? You might want to make sure they're backed up first.  ('Yes' to continue)`
    )
    const res = await prompt('')
    if (res !== 'Yes') {
      console.log('exiting')
      process.exit(0)
    }
    const { data, path } = await getDeployments()
    for (const _contractName of Object.keys(data.contracts)) {
      const contractName = _contractName as ContractNames
      data.contracts[contractName].implBuildInfo = ''
    }
    writeFileSync(path, JSON.stringify(data))
  }

  return {
    updateImplementations,
    verifyCurrentImplementations,
    deployLogic,
    deployLogicAll,
    transferAdmin,
    transferBeaconOwner,
    removeBuildInfoFiles,
  }
}
