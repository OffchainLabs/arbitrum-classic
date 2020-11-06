import { isAddress as _isAddress } from '@ethersproject/address'
import { concat, hexZeroPad } from '@ethersproject/bytes'
import { BigNumber } from '@ethersproject/bignumber'

import { getAddressIndex } from 'arb-provider-ethers'
import { JsonRpcProvider } from 'ethers/providers'

export const initSerializeAndLookUpIndices = (arbProviderUrl: string) => {
  const arbProvider = new JsonRpcProvider(arbProviderUrl)
  return async (params: any) => {
    return await serializeParams(params, async (address: string) => {
      return await getAddressIndex(address, arbProvider)
    })
  }
}

const isAddress = (input: any) => typeof input === 'string' && _isAddress(input)

type PrimativeType = string | number | boolean | BigNumber

const formatPrimative = (value: PrimativeType) => {
  if (isAddress(value)) {
    return value as string
  } else if (typeof value === 'boolean') {
    return new Uint8Array([value ? 1 : 0])
  } else if (
    typeof value === 'number' ||
    +value ||
    BigNumber.isBigNumber(value)
  ) {
    return hexZeroPad(BigNumber.from(value).toHexString(), 32)
  } else {
    throw new Error('unsupported type')
  }
}

export const serializeParams = async (
  params: (PrimativeType | PrimativeType[])[],
  addressToIndex: (address: string) => Promise<number> = () =>
    new Promise(exec => exec(-1))
) => {
  const formattedParams: (string | Uint8Array | Uint32Array)[] = []

  for (let i = 0; i < params.length; i++) {
    const param = params[i]

    if (Array.isArray(param)) {
      let paramArray: PrimativeType[] | number[] = param as PrimativeType[]
      formattedParams.push(new Uint8Array([paramArray.length]))

      if (isAddress(paramArray[0])) {
        const indices = await Promise.all(
          paramArray.map(
            async address => await addressToIndex(address as string)
          )
        )
        if (indices.every(i => i > -1)) {
          paramArray = indices as number[]
          formattedParams.push(new Uint8Array([1]))
          paramArray.forEach(value => {
            formattedParams.push(
              hexZeroPad(BigNumber.from(value).toHexString(), 4)
            )
          })
        } else {
          formattedParams.push(new Uint8Array([0]))
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
      //  not an array: handle address indexing
      if (isAddress(param)) {
        const index = await addressToIndex(param as string)
        if (index > -1) {
          formattedParams.push(new Uint8Array([1]))
          formattedParams.push(
            hexZeroPad(BigNumber.from(index).toHexString(), 4)
          )
        } else {
          formattedParams.push(new Uint8Array([0]))
          formattedParams.push(formatPrimative(param))
        }
      } else {
        formattedParams.push(formatPrimative(param))
      }
    }
  }
  const res = concat(formattedParams)
  console.info('*** Serialized byte array of length ***', res.length)
  return res
}
