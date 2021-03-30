import { isAddress as _isAddress } from '@ethersproject/address'
import { concat, hexZeroPad } from '@ethersproject/bytes'
import { BigNumber } from '@ethersproject/bignumber'

import { getAddressIndex } from './l2bridge'
import { JsonRpcProvider } from 'ethers/providers'

type PrimativeType = string | number | boolean | BigNumber
type PrimativeOrPrimativeArray = PrimativeType | PrimativeType[]
type BytesNumber = 1 | 4 | 8 | 16 | 32

// to use:
// const mySerializeParamsFunction = argSerializerConstructor("rpcurl")
export const argSerializerConstructor = (
  arbProvider: JsonRpcProvider
): ((params: PrimativeOrPrimativeArray[]) => Promise<Uint8Array>) => {
  return async (params: PrimativeOrPrimativeArray[]) => {
    return await serializeParams(params, async (address: string) => {
      return await getAddressIndex(address, arbProvider)
    })
  }
}

const isAddress = (input: PrimativeType) =>
  typeof input === 'string' && _isAddress(input)

const toUint = (val: PrimativeType, bytes: BytesNumber) =>
  hexZeroPad(BigNumber.from(val).toHexString(), bytes)

//  outputs string suitable for formatting
const formatPrimative = (value: PrimativeType) => {
  if (isAddress(value)) {
    return value as string
  } else if (typeof value === 'boolean') {
    return toUint(value ? 1 : 0, 1)
  } else if (
    typeof value === 'number' ||
    +value ||
    BigNumber.isBigNumber(value)
  ) {
    return toUint(value, 32)
  } else {
    throw new Error('unsupported type')
  }
}

export const serializeParams = async (
  params: PrimativeOrPrimativeArray[],
  addressToIndex: (address: string) => Promise<number> = () =>
    new Promise(exec => exec(-1))
): Promise<Uint8Array> => {
  const formattedParams: string[] = []

  for (const param of params) {
    // handle arrays
    if (Array.isArray(param)) {
      let paramArray: PrimativeType[] = param as PrimativeType[]
      formattedParams.push(toUint(paramArray.length, 1))

      if (isAddress(paramArray[0])) {
        const indices = await Promise.all(
          paramArray.map(
            async address => await addressToIndex(address as string)
          )
        )
        // If all addresses are registered, serialize as indices
        if (indices.every(i => i > -1)) {
          paramArray = indices as number[]
          formattedParams.push(toUint(1, 1))
          paramArray.forEach(value => {
            formattedParams.push(toUint(value, 4))
          })
          // otherwise serialize as address
        } else {
          formattedParams.push(toUint(0, 1))
          paramArray.forEach(value => {
            formattedParams.push(formatPrimative(value))
          })
        }
      } else {
        paramArray.forEach(value => {
          formattedParams.push(formatPrimative(value))
        })
      }
    } else {
      //  handle non-arrays
      if (isAddress(param)) {
        const index = await addressToIndex(param as string)
        if (index > -1) {
          formattedParams.push(toUint(1, 1), toUint(index, 4))
        } else {
          formattedParams.push(toUint(0, 1), formatPrimative(param))
        }
      } else {
        formattedParams.push(formatPrimative(param))
      }
    }
  }
  return concat(formattedParams)
}
