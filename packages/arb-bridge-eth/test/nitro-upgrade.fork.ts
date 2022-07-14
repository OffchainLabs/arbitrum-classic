import { ethers, network } from 'hardhat'
import { CurrentDeployments } from 'arb-upgrades/types'
import { readFileSync } from 'fs'
import { constants, Wallet } from 'ethers'
import { Provider } from '@ethersproject/providers'
import { NitroMigrationManager } from './nitroMigrationManager'

describe('Nitro upgrade', () => {
  const getDeployments = async (provider: Provider) => {
    const chainId = (await provider.getNetwork()).chainId
    const deploymentData = readFileSync(
      `./_deployments/${chainId}_current_deployment.json`
    )
    const deployments = JSON.parse(
      deploymentData.toString()
    ) as CurrentDeployments

    return deployments.contracts.Outbox
      ? deployments
      : {
          ...deployments,
          contracts: {
            ...deployments.contracts,
            Outbox: deployments.contracts.OldOutbox,
          },
        }
  }

  const getProxyAdminSigner = async (proxyAdminAddr: string) => {
    const proxyAdmin = await ethers.getContractAt('ProxyAdmin', proxyAdminAddr)
    const owner = await proxyAdmin.owner()
    // airdrop
    await network.provider.send('hardhat_setBalance', [
      owner,
      '0x16189AD417E380000',
    ])

    // use owner
    await network.provider.request({
      method: 'hardhat_impersonateAccount',
      params: [owner],
    })

    return await ethers.provider.getSigner(owner)
  }

  const getNewFundedSigner = async () => {
    const wallet = Wallet.createRandom().connect(ethers.provider)
    await network.provider.send('hardhat_setBalance', [
      wallet.address,
      '0x16189AD417E380000',
    ])
    return wallet
  }

  const deploymentsToClassicConfig = (deployments: CurrentDeployments) => ({
    proxyAdminAddr: deployments.proxyAdminAddress,
    inboxAddr: deployments.contracts.Inbox.proxyAddress,
    rollupAddr: deployments.contracts.Rollup.proxyAddress,
    sequencerInboxAddr: deployments.contracts.SequencerInbox.proxyAddress,
    bridgeAddr: deployments.contracts.Bridge.proxyAddress,
    outboxV1: deployments.contracts.OldOutbox.proxyAddress,
    outboxV2: deployments.contracts.Outbox.proxyAddress,
    rollupEventBridgeAddr: deployments.contracts.RollupEventBridge.proxyAddress,
  })

  const getNitroConfig = async (rollupAddr: string) => {
    const provider = ethers.provider
    const rollupFac = await ethers.getContractFactory('Rollup')
    const prevRollup = await rollupFac.attach(rollupAddr)
    const wasmModuleRoot =
      '0x9900000000000000000000000000000000000000000000000000000000000010'
    const loserStakeEscrow = constants.AddressZero
    return {
      confirmPeriodBlocks: await prevRollup.confirmPeriodBlocks(),
      extraChallengeTimeBlocks: await prevRollup.extraChallengeTimeBlocks(),
      stakeToken: await prevRollup.stakeToken(),
      baseStake: await prevRollup.baseStake(),
      wasmModuleRoot: wasmModuleRoot,
      chainId: (await provider.getNetwork()).chainId,
      loserStakeEscrow: loserStakeEscrow,
      sequencerInboxMaxTimeVariation: {
        delayBlocks: (60 * 60 * 24) / 15,
        futureBlocks: 12,
        delaySeconds: 60 * 60 * 24,
        futureSeconds: 60 * 60,
      },
    }
  }

  it.only('run succeeds', async () => {
    const provider = ethers.provider
    const deployments = await getDeployments(provider)
    const proxyAdminSigner = await getProxyAdminSigner(
      deployments.proxyAdminAddress
    )
    const nitroDeployer = await getNewFundedSigner()
    const classicConfig = deploymentsToClassicConfig(deployments)
    const nitroConfig = await getNitroConfig(
      deployments.contracts.Rollup.proxyAddress
    )

    const migrationManager = await NitroMigrationManager.deploy(
      nitroDeployer,
      true,
      true
    )

    await migrationManager.run(
      nitroDeployer,
      proxyAdminSigner,
      nitroConfig,
      classicConfig,
      true,
      true
    )
  })
})
