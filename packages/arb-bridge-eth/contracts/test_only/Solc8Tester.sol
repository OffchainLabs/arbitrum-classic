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

pragma solidity ^0.8.7;

import "../arch/IOneStepProof.sol";
import "../bridge/interfaces/IBridge.sol";
import "../bridge/interfaces/IInbox.sol";
import "../bridge/interfaces/IMessageProvider.sol";
import "../bridge/interfaces/IOutbox.sol";
import "../bridge/interfaces/ISequencerInbox.sol";
import "../challenge/IChallenge.sol";
import "../challenge/IChallengeFactory.sol";
import "../interfaces/IERC20.sol";
import "../interfaces/IERC721.sol";
import "../libraries/ICloneable.sol";
import "../rollup/INode.sol";
import "../rollup/INodeFactory.sol";
import "../rollup/IRollupCore.sol";
import "../rollup/facets/IRollupFacets.sol";
import "../validator/IGasRefunder.sol";

import "arb-bridge-peripherals/contracts/tokenbridge/arbitrum/IArbToken.sol";
import "arb-bridge-peripherals/contracts/tokenbridge/ethereum/ICustomToken.sol";
import "arb-bridge-peripherals/contracts/tokenbridge/libraries/IWETH9.sol";
import "arb-bridge-peripherals/contracts/tokenbridge/libraries/gateway/ICustomGateway.sol";
import "arb-bridge-peripherals/contracts/tokenbridge/libraries/gateway/ITokenGateway.sol";
