/*
 * Copyright 2019, Offchain Labs, Inc.
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

import {ArbVM} from './vmTrackerWrapper'

export class ArbManager {
  constructor(vmTracker, arbMachineLib, arbValueLib, address) {
    this.vmTracker = vmTracker;
    this.arbMachineLib = arbMachineLib;
    this.arbValueLib = arbValueLib;
    this.address = address;
  }

  async createVm(vmId, config, vmState, signatures, cMan) {
    let self = this;

    var sigVs = [];
    var sigRs = [];
    var sigSs = [];

    signatures.forEach(function(signature) {
      sigVs.push(signature.v);
      sigRs.push('0x' + signature.r.toString('hex'));
      sigSs.push('0x' + signature.s.toString('hex'));
    });

    let createHash = await self.vmTracker.createVMHash(
      config.gracePeriod,
      config.escrowRequired,
      config.maxExecutionSteps,
      vmState,
      config.challengeVerifier,
      config.assertKeys
    );
    return self.vmTracker.createVm(
        config["gracePeriod"], 
        config["escrowRequired"],
        config["maxExecutionSteps"],
        vmId,
        vmState,
        config["challengeVerifier"],
        createHash,
        sigVs,
        sigRs,
        sigSs,
        {from: self.address}
    ).then(function(result) {
      config.challengeVerifier = cMan;
      return new ArbVM(self.vmTracker, self.arbMachineLib, self.arbValueLib, vmId, config, self.address);
    });
  }

  createDefaultVm(vmId, managers, challengeVerifier, vmState) {
    return this.createVm(vmId, {
        "gracePeriod": 10,
        "escrowRequired": 50000,
        "assertKeys": managers,
        "maxExecutionSteps":100000,
        "challengeVerifier":challengeVerifier
      }, vmState
    );
  }

  getVm(vmId, challengeManager, assertKeys) {
    return ArbVM.vmWithId(challengeManager, this.vmTracker, this.arbMachineLib, this.arbValueLib, vmId, this.address, assertKeys);
  }

  async sendMessage(destination, value, data) {
    let result = await this.vmTracker.sendMessage(destination, value, data, {from: this.address});
  }
}
