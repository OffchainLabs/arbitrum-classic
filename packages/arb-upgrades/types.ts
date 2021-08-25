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
}

export enum UpgradeableType {
  Beacon = 'Beacon',
  TransparentProxy = 'TransparentProxy',
  Proxy = 'Proxy',
}

export const proxyType = (contractName: ContractNames) => {
  switch (contractName) {
    case ContractNames.StandardArbERC20:
      return UpgradeableType.Beacon
    default:
      return UpgradeableType.TransparentProxy
  }
}

export const isBeacon = (contractName: ContractNames) =>
  proxyType(contractName) === UpgradeableType.Beacon
