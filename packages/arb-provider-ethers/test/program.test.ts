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
/* eslint-env node, jest */
'use strict'

import * as program from '../src/lib/program'
import * as fs from 'fs'

describe('Contract Parsing', function () {
  it('should have no backward references', async () => {
    const data = fs.readFileSync('../../arb-os/arb_os/arbos.mexe', 'utf8')
    const prog = JSON.parse(data) as program.Program

    let i = 0
    for (const op of prog.code) {
      if (op.immediate && op.immediate.CodePoint) {
        if (op.immediate.CodePoint.Internal <= i) {
          throw Error(
            `invalid internal reference ${i} ${op.immediate.CodePoint.Internal}`
          )
        }
      }
      i += 1
    }
  })

  it.skip("should load ArbOS and calulculate it's hash", function () {
    const data = fs.readFileSync('../../arb-os/arb_os/arbos.mexe', 'utf8')
    const machineHash = program.programMachineHash(data)
    console.log(machineHash)
  }).timeout(5000)
})
