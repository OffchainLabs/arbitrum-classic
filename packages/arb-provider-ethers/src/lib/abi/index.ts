import {
  BigNumberish,
  EventDescription,
  FunctionDescription,
} from 'ethers/utils'

export class TransactionOverrides {
  nonce?: BigNumberish | Promise<BigNumberish>
  gasLimit?: BigNumberish | Promise<BigNumberish>
  gasPrice?: BigNumberish | Promise<BigNumberish>
  value?: BigNumberish | Promise<BigNumberish>
  chainId?: number | Promise<number>
}

export interface TypedEventDescription<
  T extends Pick<EventDescription, 'encodeTopics'>
> extends EventDescription {
  encodeTopics: T['encodeTopics']
}

export interface TypedFunctionDescription<
  T extends Pick<FunctionDescription, 'encode'>
> extends FunctionDescription {
  encode: T['encode']
}

export type { ArbErc20 } from './ArbErc20'
export type { ArbErc721 } from './ArbErc721.d'
export type { ArbFactory } from './ArbFactory.d'
export type { ArbInfo } from './ArbInfo.d'
export type { ArbRollup } from './ArbRollup.d'
export type { ArbSys } from './ArbSys.d'
export type { GlobalInbox } from './GlobalInbox.d'

export { ArbFactoryFactory } from './ArbFactoryFactory'
export { ArbRollupFactory } from './ArbRollupFactory'
export { GlobalInboxFactory } from './GlobalInboxFactory'
export { ArbErc20Factory } from './ArbErc20Factory'
export { ArbErc721Factory } from './ArbErc721Factory'
