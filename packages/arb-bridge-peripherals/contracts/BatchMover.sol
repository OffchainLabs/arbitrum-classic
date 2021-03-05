// SPDX-License-Identifier: Apache-2.0

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

pragma solidity ^0.6.11;

import "./MMR.sol";
import "./tokenbridge/arbitrum/StandardArbERC20.sol";
import "./buddybridge/ethereum/L1Buddy.sol";
import "./buddybridge/util/BuddyUtil.sol";

import "arbos-contracts/arbos/builtin/ArbSys.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

contract ArbBatchTokenMover {
    using MMR for MMR.Tree;

    MMR.Tree withdrawalTree;
    uint256 exitBlock;
    StandardArbERC20 erc20;

    function withdrawInBatch(uint256 amount) external {
        require(erc20.transferFrom(msg.sender, address(this), amount), "TRANSFER_FAILED");
        withdrawalTree.append(abi.encode(msg.sender, amount));
    }

    function exitToL1() external {
        require(block.number >= exitBlock, "TOO_SOON");
        ArbSys(100).sendTxToL1(
            address(this),
            abi.encodeWithSignature(
                "distributeBatch(bytes32,address)",
                withdrawalTree.getRoot(),
                erc20.l1Address()
            )
        );

        erc20.withdraw(address(this), erc20.balanceOf(address(this)));
        selfdestruct(msg.sender);
    }
}

contract EthBatchTokenReceiver is L1Buddy {
    bytes32 root;
    IERC20 erc20;
    mapping(uint256 => bool) redeemed;
    address l2Buddy;

    constructor(
        address _inbox,
        address _l2Deployer,
        uint256 _maxGas,
        uint256 _gasPrice
    ) public payable L1Buddy(_inbox, _l2Deployer) {
        L1Buddy.initiateBuddyDeploy(
            _maxGas,
            _gasPrice,
            type(ArbBatchTokenMover).creationCode
        );
    }

    function handleDeploySuccess() internal override {
        // TODO: should we check if connection state is pending?
        l2Buddy = BuddyUtil.calculateL2Address(
            address(L1Buddy.l2Deployer),
            address(this),
            keccak256(type(ArbBatchTokenMover).creationCode)
        );
        // this deletes the codehash from state!
        L1Buddy.handleDeploySuccess();
    }
    function handleDeployFail() internal override {}

    modifier onlyIfConnected {
        require(L1Buddy.l2Connection == L1Buddy.L2Connection.Complete, "Not connected");
        _;
    }

    modifier onlyL2Buddy {
        require(l2Buddy != address(0), "l2 buddy not set");
        IOutbox outbox = IOutbox(L1Buddy.inbox.bridge().activeOutbox());
        require(l2Buddy == outbox.l2ToL1Sender(), "Not from l2 buddy");
        _;
    }

    function distributeBatch(bytes32 _root) external onlyIfConnected onlyL2Buddy {
        root = _root;
    }

    function redeemWithdrawal(
        address dest,
        uint256 amount,
        uint256 width,
        uint256 index,
        bytes32[] memory peaks,
        bytes32[] memory siblings
    ) public {
        require(root != 0, "NOT_INITIALIZED");
        require(!redeemed[index], "ALREADY_REDEEMED");
        redeemed[index] = true;
        require(
            MMR.inclusionProof(root, width, index, abi.encode(dest, amount), peaks, siblings) ==
                true,
            "BAD_PROOF"
        );
        require(erc20.transfer(dest, amount), "BAD_TRANSFER");
    }
}
