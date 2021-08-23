export interface QueuedUpdate {
  address: string
  deployTxn: string
  arbitrumCommitHash: string
  buildInfo: string
}

export interface CurrentDeployment {
  proxyAddress: string
  implAddress: string
  implDeploymentTxn: string
  implArbitrumCommitHash: string
  implBuildInfo: string
  isBeacon?: boolean
}

export type QueuedUpdates = {
  [key in ContractNames]?: QueuedUpdate
}

export interface CurrentDeployments {
  proxyAdminAddress: string
  contracts: {
    [key in ContractNames]: CurrentDeployment
  }
}

export enum ContractNames {
  L1GatewayRouter = 'L1GatewayRouter',
  L1ERC20Gateway = 'L1ERC20Gateway',
  L1CustomGateway = 'L1CustomGateway',
  L1WethGateway = 'L1WethGateway',

  L2ERC20Gateway = 'L2ERC20Gateway',
  L2GatewayRouter = 'L2GatewayRouter',
  L2CustomGateway = 'L2CustomGateway',
  L2WethGateway = 'L2WethGateway',
  StandardArbERC20 = 'StandardArbERC20',
  Rollup = 'Rollup',
  RollupAdminFacet = 'RollupAdminFacet',
  RollupUserFacet = 'RollupUserFacet',
  RollupEventBridge = 'RollupEventBridge',
  Node = 'Node',
  Challenge = 'Challenge',
  OneStepProof = 'OneStepProof',
  OneStepProof2 = 'OneStepProof2',
  OneStepProofHash = 'OneStepProofHash',
  Inbox = 'Inbox',
  Bridge = 'Bridge',
  SequencerInbox = 'SequencerInbox',
}

export enum UpgradeableType {
  Beacon = 'Beacon',
  TransparentProxy = 'TransparentProxy',
  Proxy = 'Proxy',
  RollupUserFacet = 'RollupUserFacet',
  RollupAdminFacet = 'RollupAdminFacet',
}

export const proxyType = (contractName: ContractNames) => {
  switch (contractName) {
    case ContractNames.StandardArbERC20:
    case ContractNames.Node:
    case ContractNames.Challenge:
      return UpgradeableType.Beacon
    case ContractNames.RollupAdminFacet:
      return UpgradeableType.RollupAdminFacet
    case ContractNames.RollupUserFacet:
      return UpgradeableType.RollupUserFacet
    default:
      return UpgradeableType.TransparentProxy
  }
}

export const isBeacon = (contractName: ContractNames) =>
  proxyType(contractName) === UpgradeableType.Beacon

export const isRollupUserFacet = (contractName: ContractNames) =>
  proxyType(contractName) === UpgradeableType.RollupUserFacet

export const isRollupAdminFacet = (contractName: ContractNames) =>
  proxyType(contractName) === UpgradeableType.RollupAdminFacet
