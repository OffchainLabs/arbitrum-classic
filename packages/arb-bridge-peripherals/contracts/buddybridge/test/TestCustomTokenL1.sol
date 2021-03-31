// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

import "../../tokenbridge/arbitrum/open-zeppelin/OZERC20.sol";
import "../../tokenbridge/ethereum/ICustomToken.sol";
import "../../tokenbridge/ethereum/EthERC20Bridge.sol";

contract TestCustomTokenL1 is OZERC20, ICustomToken {
    EthERC20Bridge public bridge;

    constructor(address _bridge) public {
        bridge = EthERC20Bridge(_bridge);
        _name = "TestCustomToken";
        _symbol = "CARB";
        _decimals = uint8(18);
    }

    function mint() external {
        _mint(msg.sender, 50000000);
    }

    function registerTokenOnL2(
        address l2CustomTokenAddress,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        address refundAddress
    ) public override {
        bridge.registerCustomL2Token(
            l2CustomTokenAddress,
            maxSubmissionCost,
            maxGas,
            gasPriceBid,
            refundAddress
        );
    }
}
