// SPDX-License-Identifier: Apache-2.0

pragma solidity ^0.6.11;

import "../libraries/aeERC20.sol";
import "../ethereum/ICustomToken.sol";
import "../ethereum/gateway/L1CustomGateway.sol";
import "../ethereum/gateway/L1GatewayRouter.sol";
import "@openzeppelin/contracts/GSN/Context.sol";
import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract TestCustomTokenL1 is aeERC20, ICustomToken {
    address public bridge;
    address public router;
    bool private shouldRegisterGateway;

    constructor(address _bridge, address _router) public {
        bridge = _bridge;
        router = _router;
        aeERC20._initialize("TestCustomToken", "CARB", uint8(18));
    }

    function mint() external {
        _mint(msg.sender, 50000000);
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public virtual override(ERC20Upgradeable, ICustomToken) returns (bool) {
        return ERC20Upgradeable.transferFrom(sender, recipient, amount);
    }

    function balanceOf(address account)
        public
        view
        virtual
        override(ERC20Upgradeable, ICustomToken)
        returns (uint256)
    {
        return ERC20Upgradeable.balanceOf(account);
    }

    /// @dev we only set shouldRegisterGateway to true when in `registerTokenOnL2`
    function isArbitrumEnabled() external view override returns (uint8) {
        require(shouldRegisterGateway, "NOT_EXPECTED_CALL");
        return uint8(0xa4b1);
    }

    function registerTokenOnL2(
        address l2CustomTokenAddress,
        uint256 maxSubmissionCostForCustomBridge,
        uint256 maxSubmissionCostForRouter,
        uint256 maxGasForCustomBridge,
        uint256 maxGasForRouter,
        uint256 gasPriceBid,
        uint256 valueForGateway,
        uint256 valueForRouter,
        address creditBackAddress
    ) public payable override {
        // we temporarily set `shouldRegisterGateway` to true for the callback in registerTokenToL2 to succeed
        bool prev = shouldRegisterGateway;
        shouldRegisterGateway = true;

        // L1CustomGateway(bridge).registerTokenToL2{value: valueForGateway}(
        L1CustomGateway(bridge).registerTokenToL2(
            l2CustomTokenAddress,
            maxGasForCustomBridge,
            gasPriceBid,
            maxSubmissionCostForCustomBridge,
            creditBackAddress
        );

        // L1GatewayRouter(router).setGateway{value: valueForRouter}(
        L1GatewayRouter(router).setGateway(
            bridge,
            maxGasForRouter,
            gasPriceBid,
            maxSubmissionCostForRouter,
            creditBackAddress
        );

        shouldRegisterGateway = prev;
    }
}

contract MintableTestCustomTokenL1 is L1MintableToken, TestCustomTokenL1 {
    constructor(address _bridge) public TestCustomTokenL1(_bridge) {}

    function bridgeMint(address account, uint256 amount) public override(L1MintableToken) {
        _mint(account, amount);
    }

    function balanceOf(address account)
        public
        view
        override(TestCustomTokenL1, ICustomToken)
        returns (uint256 amount)
    {
        return super.balanceOf(account);
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public override(TestCustomTokenL1, ICustomToken) returns (bool) {
        return super.transferFrom(sender, recipient, amount);
    }
}
