// SPDX-License-Identifier: Apache-2.0

/*
 * Copyright 2021, Offchain Labs, Inc.
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

pragma solidity ^0.6.11;

import "../rollup/Rollup.sol";
import "../rollup/Node.sol";

contract Validator {
    function successorNodes(Rollup rollup, uint256 nodeNum)
        external
        view
        returns (uint256[] memory)
    {
        uint256[] memory nodes = new uint256[](100000);
        uint256 index = 0;
        for (uint256 i = nodeNum + 1; i <= rollup.latestNodeCreated(); i++) {
            Node node = Node(rollup.nodes(i));
            if (node.prev() == nodeNum) {
                nodes[index] = i;
                index++;
            }
        }
        // Shrink array down to real size
        assembly {
            mstore(nodes, index)
        }
        return nodes;
    }

    function stakedNodes(Rollup rollup, address staker) external view returns (uint256[] memory) {
        uint256[] memory nodes = new uint256[](100000);
        uint256 index = 0;
        for (uint256 i = rollup.latestConfirmed(); i <= rollup.latestNodeCreated(); i++) {
            Node node = Node(rollup.nodes(i));
            if (node.stakers(staker)) {
                nodes[index] = i;
                index++;
            }
        }
        // Shrink array down to real size
        assembly {
            mstore(nodes, index)
        }
        return nodes;
    }
}
