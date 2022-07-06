// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2022, Offchain Labs, Inc.
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

pragma solidity >=0.6.9 <0.9.0;

import "arb-bridge-eth/contracts/arch/IOneStepProof.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IBridge.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IMessageProvider.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IOutbox.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/ISequencerInbox.sol";
import "arb-bridge-eth/contracts/challenge/IChallenge.sol";
import "arb-bridge-eth/contracts/challenge/IChallengeFactory.sol";
import "arb-bridge-eth/contracts/interfaces/IERC20.sol";
import "arb-bridge-eth/contracts/interfaces/IERC721.sol";
import "arb-bridge-eth/contracts/libraries/ICloneable.sol";
import "arb-bridge-eth/contracts/rollup/INode.sol";
import "arb-bridge-eth/contracts/rollup/INodeFactory.sol";
import "arb-bridge-eth/contracts/rollup/IRollupCore.sol";
import "arb-bridge-eth/contracts/rollup/facets/IRollupFacets.sol";
import "arb-bridge-eth/contracts/validator/IGasRefunder.sol";

import "../arbitrum/IArbToken.sol";
import "../ethereum/ICustomToken.sol";
import "../libraries/IWETH9.sol";
import "../libraries/gateway/ICustomGateway.sol";
import "../libraries/gateway/ITokenGateway.sol";
