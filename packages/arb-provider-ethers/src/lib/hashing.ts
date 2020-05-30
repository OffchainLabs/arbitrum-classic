/*
 * Copyright 2019-2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/* eslint-env node */
'use strict'

import * as ethers from 'ethers'

export enum TxType {
  Transaction = 0,
  DepositEth = 1,
  DepositERC20 = 2,
  DepositERC721 = 3,
  TransactionBatch = 5,
}

export function calculateTransactionHash(
  chain: string,
  to: string,
  from: string,
  sequenceNum: ethers.utils.BigNumber,
  value: ethers.utils.BigNumber,
  data: string
): string {
  return ethers.utils.solidityKeccak256(
    ['uint8', 'address', 'address', 'address', 'uint256', 'uint256', 'bytes32'],
    [
      TxType.Transaction,
      chain,
      to,
      from,
      sequenceNum,
      value,
      ethers.utils.solidityKeccak256(['bytes'], [data]),
    ]
  )
}

export function calculateBatchTransactionHash(
  chain: string,
  to: string,
  sequenceNum: ethers.utils.BigNumber,
  value: ethers.utils.BigNumber,
  data: string
): string {
  return ethers.utils.solidityKeccak256(
    ['address', 'address', 'uint256', 'uint256', 'bytes32'],
    [
      chain,
      to,
      sequenceNum,
      value,
      ethers.utils.solidityKeccak256(['bytes'], [data]),
    ]
  )
}
