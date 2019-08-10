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

const ethers = require("ethers");
const ArbProvider = require("arb-provider-ethers");
var ProviderBridge = require("./ethers-web3-bridge");

module.exports = async function(managerUrl, contracts, provider) {
  let wrappedProv = new ethers.providers.Web3Provider(provider);
  let arbProvider = new ArbProvider(managerUrl, contracts, wrappedProv);
  let wallet = await arbProvider.getSigner(0);
  return new ProviderBridge(arbProvider, wallet);
};
