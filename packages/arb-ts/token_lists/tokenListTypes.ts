/* from https://github.com/Uniswap/token-lists/blob/master/src/types.ts */

export interface TokenInfo {
  readonly chainId: number
  readonly address: string
  readonly name: string
  readonly decimals: number
  readonly symbol: string
  readonly logoURI?: string
  readonly tags?: string[]
  readonly extensions?: {
    readonly [key: string]: string | number | boolean | null
  }
}

export interface Version {
  readonly major: number
  readonly minor: number
  readonly patch: number
}

export interface Tags {
  readonly [tagId: string]: {
    readonly name: string
    readonly description: string
  }
}

export interface TokenList {
  readonly name: string
  readonly timestamp: string
  readonly version: Version
  readonly tokens: TokenInfo[]
  readonly keywords?: string[]
  readonly tags?: Tags
  readonly logoURI?: string
}
