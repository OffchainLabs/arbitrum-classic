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

/* eslint-env node, mocha */
import { ethers, run, deployments } from 'hardhat'

import { NodeFactory } from '../build/types/NodeFactory'

let nodeFactory: NodeFactory
describe('NodeFactory', () => {
  it('should deploy contracts', async function () {
    await run('deploy', { tags: 'test' })

    const NodeFactoryDeployment = await deployments.get('NodeFactory')
    const nodeFactoryDeployment = await ethers.getContractAt(
      'NodeFactory',
      NodeFactoryDeployment.address
    )

    nodeFactory = nodeFactoryDeployment as NodeFactory
  })

  it('should create node', async function () {
    await nodeFactory.createNode(
      ethers.constants.HashZero,
      ethers.constants.HashZero,
      ethers.constants.HashZero,
      0,
      0
    )
  })
})
