import {
  BuidlerRuntimeEnvironment,
  DeployFunction,
} from '@nomiclabs/buidler/types'

const func: DeployFunction = async (bre: BuidlerRuntimeEnvironment) => {
  const { deployments, getNamedAccounts } = bre
  const { deploy, log } = deployments
  const { deployer } = await getNamedAccounts()

  let contract = await deployments.getOrNull('OneStepProof')
  if (!contract) {
    const utilLibrary = await deploy("MerkleUtil", { from: deployer });

    const deployResult = await deploy('OneStepProof', {
      from: deployer,
      libraries: {
        ['MerkleUtil']: utilLibrary.address,
      }
    })
    contract = await deployments.get('OneStepProof')
    if (deployResult.newlyDeployed && deployResult.receipt) {
      log(
        `OneStepProof deployed at ${contract.address} for ${deployResult.receipt.gasUsed}`
      )
    }
  }
}

module.exports = func
module.exports.tags = ['OneStepProof']
