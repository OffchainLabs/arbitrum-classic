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

import "./ConfirmRoots.sol";
import "./ExitLiquidityProvider.sol";
import "./L1Buddy.sol";
import "../arbitrum/ArbERC20Bridge.sol";
import "arb-bridge-eth/contracts/bridge/interfaces/IInbox.sol";
import "arb-bridge-eth/contracts/libraries/MerkleLib.sol";

contract EthERC20Bridge is L1Buddy {
    uint256 internal constant SendType_sendTxToL1 = 0;

    ConfirmRoots confirmRoots;
    // exitNum => exitDataHash => LP
    mapping(bytes32 => address) redirectedExits;

    constructor(IInbox _inbox, address _confirmRoots) public L1Buddy(_inbox) {
        confirmRoots = ConfirmRoots(_confirmRoots);
    }

    function connectToChain(uint256 maxGas, uint256 gasPriceBid) external payable {
        // Pay for gas
        if (msg.value > 0) {
            inbox.depositEth{ value: msg.value }(address(this));
        }
        inbox.deployL2ContractPair(
            maxGas, // max gas
            gasPriceBid, // gas price
            0, // payment
            type(ArbERC20Bridge).creationCode
        );
    }

    function fastWithdrawalFromL2(
        address liquidityProvider,
        uint256 nodeNum,
        bytes32[] calldata proof,
        uint256 merklePath,
        uint256 l2Block,
        uint256 l2Timestamp,
        address erc20,
        uint256 amount,
        uint256 exitNum
    ) external {
        bytes32 withdrawData = keccak256(abi.encodePacked(exitNum, msg.sender, erc20, amount));
        require(redirectedExits[withdrawData] == address(0), "ALREADY_EXITED");
        redirectedExits[withdrawData] = liquidityProvider;
        bytes memory data =
            abi.encodeWithSignature(
                "withdrawFromL2(uint256,address,address,uint256)",
                exitNum,
                msg.sender,
                erc20,
                amount
            );
        bytes32 userTx =
            keccak256(
                abi.encodePacked(
                    SendType_sendTxToL1,
                    uint256(uint160(bytes20(address(this)))),
                    uint256(uint160(bytes20(address(this)))),
                    l2Block,
                    l2Timestamp,
                    uint256(0),
                    data
                )
            );
        bytes32 calcRoot =
            MerkleLib.calculateRoot(proof, merklePath, keccak256(abi.encodePacked(userTx)));
        require(confirmRoots.confirmRoots(calcRoot, nodeNum), "INVALID_ROOT");

        ExitLiquidityProvider(liquidityProvider).provideLiquidity(
            nodeNum,
            msg.sender,
            erc20,
            amount
        );
    }

    function withdrawFromL2(
        uint256 exitNum,
        address erc20,
        address destination,
        uint256 amount
    ) external onlyIfConnected onlyL2 {
        // This method is only callable by this contract's buddy contract on L2
        require(l2Sender() == address(this), "L2_SENDER");
        bytes32 withdrawData = keccak256(abi.encodePacked(exitNum, destination, erc20, amount));
        address exitAddress = redirectedExits[withdrawData];
        // Unsafe external calls must occur below checks and effects
        if (exitAddress != address(0)) {
            require(IERC20(erc20).transfer(exitAddress, amount));
        } else {
            require(IERC20(erc20).transfer(destination, amount));
        }
    }

    function updateTokenInfo(
        address erc20,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable onlyIfConnected {
        string memory name;
        string memory symbol;
        uint8 decimals;
        try ERC20(erc20).name() returns (string memory _name) {
            name = _name;
        } catch {}
        try ERC20(erc20).symbol() returns (string memory _symbol) {
            symbol = _symbol;
        } catch {}
        try ERC20(erc20).decimals() returns (uint8 _decimals) {
            decimals = _decimals;
        } catch {}
        inbox.sendL1FundedContractTransaction{ value: msg.value }(
            maxGas,
            gasPriceBid,
            address(this),
            abi.encodeWithSignature(
                "updateTokenInfo(address,string,string,uint8)",
                erc20,
                name,
                symbol,
                decimals
            )
        );
    }

    function deposit(
        address erc20,
        address destination,
        uint256 amount,
        uint256 maxGas,
        uint256 gasPriceBid
    ) external payable onlyIfConnected {
        require(IERC20(erc20).transferFrom(msg.sender, address(this), amount));
        // This transfers along any ETH sent for to pay for gas in L2
        inbox.sendL1FundedContractTransaction{ value: msg.value }(
            maxGas,
            gasPriceBid,
            address(this),
            abi.encodeWithSignature(
                "mintFromL1(address,address,uint256)",
                erc20,
                destination,
                amount
            )
        );
    }
}
