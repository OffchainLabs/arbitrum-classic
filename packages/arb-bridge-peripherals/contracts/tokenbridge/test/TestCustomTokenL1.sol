// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

import "../libraries/aeERC20.sol";
import "../ethereum/ICustomToken.sol";
import "../ethereum/gateway/L1CustomGateway.sol";
import "@openzeppelin/contracts/GSN/Context.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract TestCustomTokenL1 is aeERC20, ICustomToken {
    address public bridge;

    constructor(address _bridge) public {
        bridge = _bridge;
        aeERC20._initialize("TestCustomToken", "CARB", uint8(18));
    }

    function mint() external {
        _mint(msg.sender, 50000000);
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public override(ERC20Upgradeable, ICustomToken) returns (bool) {
        return ERC20Upgradeable.transferFrom(sender, recipient, amount);
    }

    function balanceOf(address account)
        public
        view
        override(ERC20Upgradeable, ICustomToken)
        returns (uint256)
    {
        return ERC20Upgradeable.balanceOf(account);
    }

    function registerTokenOnL2(
        address l2CustomTokenAddress,
        uint256 maxSubmissionCost,
        uint256 maxGas,
        uint256 gasPriceBid,
        address creditBackAddress
    ) public override {
        L1CustomGateway(bridge).registerTokenToL2(
            l2CustomTokenAddress,
            maxGas,
            gasPriceBid,
            maxSubmissionCost,
            creditBackAddress
        );
    }
}
