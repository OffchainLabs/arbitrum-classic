/*
 * Copyright 2020, Offchain Labs, Inc.
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

import * as ArbValue from './value'
import * as ethers from 'ethers'

export interface CodePointRef {
  Internal: number
}

export interface Value {
  Int?: string
  Tuple?: Value[]
  CodePoint?: CodePointRef
}

export interface Opcode {
  AVMOpcode: number
}

export interface Operation {
  opcode: Opcode
  immediate: Value | undefined
}

export interface Program {
  code: Operation[]
  static_val: Value
}

function loadValue(
  val: Value,
  codePoints: ArbValue.CodePointValue[],
  totalOps: number
): ArbValue.Value {
  if (val.Int) {
    return new ArbValue.IntValue('0x' + val.Int)
  }

  if (val.Tuple) {
    return new ArbValue.TupleValue(
      val.Tuple.map(subVal => loadValue(subVal, codePoints, totalOps))
    )
  }

  if (val.CodePoint) {
    const cpIndex = totalOps - val.CodePoint.Internal
    if (cpIndex < 0 || cpIndex >= codePoints.length) {
      throw Error(
        `invalid internal reference ${val.CodePoint.Internal} ${totalOps} ${cpIndex} ${codePoints.length}`
      )
    }
    return codePoints[cpIndex]
  }

  throw Error('Invalid value')
}

function loadOperation(
  op: Operation,
  codePoints: ArbValue.CodePointValue[],
  totalOps: number
): ArbValue.Operation {
  if (op.immediate) {
    return new ArbValue.ImmOp(
      op.opcode.AVMOpcode,
      loadValue(op.immediate, codePoints, totalOps)
    )
  } else {
    return new ArbValue.BasicOp(op.opcode.AVMOpcode)
  }
}

function loadProgram(
  progString: string
): [ArbValue.CodePointValue[], ArbValue.Value] {
  const prog = JSON.parse(progString) as Program

  let nextHash = ethers.utils.hexZeroPad('0x', 32)

  const codePoints: ArbValue.CodePointValue[] = [
    new ArbValue.CodePointValue(new ArbValue.BasicOp(0), nextHash),
  ]
  nextHash = codePoints[0].hash()

  for (let i = 0; i < prog.code.length; i++) {
    const rawOp = prog.code[prog.code.length - 1 - i]
    const op = loadOperation(rawOp, codePoints, prog.code.length)
    const cp = new ArbValue.CodePointValue(op, nextHash)
    nextHash = cp.hash()
    codePoints.push(cp)
  }

  const staticVal = loadValue(prog.static_val, codePoints, prog.code.length)

  return [codePoints.reverse(), staticVal]
}

export function programMachineHash(progString: string): string {
  const [codePoints, staticVal] = loadProgram(progString)
  return machineHash(
    codePoints[0],
    new ArbValue.TupleValue([]),
    new ArbValue.TupleValue([]),
    new ArbValue.TupleValue([]),
    staticVal,
    new ArbValue.CodePointValue(
      new ArbValue.BasicOp(0),
      ethers.utils.hexZeroPad('0x00', 32)
    ),
    new ArbValue.IntValue(ethers.constants.MaxUint256),
    new ArbValue.TupleValue([])
  )
}

export function machineHash(
  pc: ArbValue.Value,
  stack: ArbValue.Value,
  auxstack: ArbValue.Value,
  registerVal: ArbValue.Value,
  staticVal: ArbValue.Value,
  errPc: ArbValue.Value,
  arbGasRemaining: ArbValue.IntValue,
  pendingMessage: ArbValue.Value
): string {
  return ethers.utils.solidityKeccak256(
    [
      'bytes32',
      'bytes32',
      'bytes32',
      'bytes32',
      'bytes32',
      'uint256',
      'bytes32',
      'bytes32',
    ],
    [
      pc.hash(),
      stack.hash(),
      auxstack.hash(),
      registerVal.hash(),
      staticVal.hash(),
      arbGasRemaining.bignum,
      errPc.hash(),
      pendingMessage.hash(),
    ]
  )
}
