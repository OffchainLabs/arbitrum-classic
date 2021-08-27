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
  Outbox = 'Outbox',
  OutboxEntry = 'OutboxEntry',
}

export enum UpgradeableType {
  BeaconOwnedByEOA = 'BeaconOwnedByEOA',
  BeaconOwnedByRollup = 'BeaconOwnedByRollup',
  TransparentProxy = 'TransparentProxy',
  Proxy = 'Proxy',
  RollupUserFacet = 'RollupUserFacet',
  RollupAdminFacet = 'RollupAdminFacet',
}

export const proxyType = (contractName: ContractNames) => {
  switch (contractName) {
    case ContractNames.StandardArbERC20:
      return UpgradeableType.BeaconOwnedByEOA
    case ContractNames.Node:
    case ContractNames.Challenge:
    case ContractNames.OutboxEntry:
      return UpgradeableType.BeaconOwnedByRollup
    case ContractNames.RollupAdminFacet:
      return UpgradeableType.RollupAdminFacet
    case ContractNames.RollupUserFacet:
      return UpgradeableType.RollupUserFacet
    default:
      return UpgradeableType.TransparentProxy
  }
}

export const isBeacon = (contractName: ContractNames) =>
  isBeaconOwnedByEOA(contractName) || isBeaconOwnedByRollup(contractName)

export const isBeaconOwnedByEOA = (contractName: ContractNames) =>
  proxyType(contractName) === UpgradeableType.BeaconOwnedByEOA ||
  proxyType(contractName) === UpgradeableType.BeaconOwnedByRollup

export const isBeaconOwnedByRollup = (contractName: ContractNames) =>
  proxyType(contractName) === UpgradeableType.BeaconOwnedByRollup

export const isRollupUserFacet = (contractName: ContractNames) =>
  proxyType(contractName) === UpgradeableType.RollupUserFacet

export const isRollupAdminFacet = (contractName: ContractNames) =>
  proxyType(contractName) === UpgradeableType.RollupAdminFacet
export const getLayer = (contractName: ContractNames) => {
  switch (contractName) {
    case 'L2ERC20Gateway':
    case 'L2GatewayRouter':
    case 'L2WethGateway':
    case 'L2CustomGateway':
    case 'StandardArbERC20':
      return 2
    default:
      return 1
  }
}

export const hasPostInitHook = (contractName: ContractNames) => {
  switch (contractName) {
    case ContractNames.L1GatewayRouter:
    case ContractNames.L1ERC20Gateway:
    case ContractNames.L1CustomGateway:
    case ContractNames.L1WethGateway:
    case ContractNames.L2ERC20Gateway:
    case ContractNames.L2GatewayRouter:
    case ContractNames.L2CustomGateway:
    case ContractNames.L2WethGateway:
    case ContractNames.Rollup:
    case ContractNames.SequencerInbox:
      // case ContractNames.Inbox: // Uncomment if upgrading to Inbox with postInitHook
      return true
    default:
      return false
  }
}
