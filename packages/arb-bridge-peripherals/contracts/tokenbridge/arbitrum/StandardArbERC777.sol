// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.6.11;

import "arb-bridge-eth/contracts/libraries/Cloneable.sol";
import "./open-zeppelin/OZERC777.sol";
import "../libraries/DecimalConverter.sol";
import "./IArbToken.sol";
import "./ArbTokenBridge.sol";

contract StandardArbERC777 is OZERC777, Cloneable, IArbToken {
    ArbTokenBridge public bridge;
    address public l1Address;
    uint8 public l1Decimals;

    modifier onlyBridge {
        require(msg.sender == address(bridge), "ONLY_BRIDGE");
        _;
    }

    function initialize(
        address _bridge,
        address _l1Address,
        uint8 _decimals
    ) external override {
        require(address(bridge) == address(0), "ALREADY_INIT");
        bridge = ArbTokenBridge(_bridge);
        l1Address = _l1Address;

        // require(_decimals <= 18, "Decimals must be less than or equal to 18");
        OZERC777.initialize(DecimalConverter.decimalsToGranularity(_decimals));
        l1Decimals = _decimals;
    }

    function updateInfo(string memory newName, string memory newSymbol, uint8 newDecimals) public override onlyBridge {
        _name = newName;
        _symbol = newSymbol;
        require(
            OZERC777._granularity == DecimalConverter.decimalsToGranularity(newDecimals),
            "777 granularity can't change"
        );
    }

    function bridgeMint(address account, uint256 amount, bytes memory data) external override onlyBridge {
        _mint(account, amount, data, "");
    }

    function withdraw(address destination, uint256 amount) external override {
        _burn(msg.sender, amount, "", "");
        bridge.withdraw(l1Address, destination, amount);
    }

    function migrate(address target, uint256 amount) external {
        _burn(msg.sender, amount, "", "");
        // migrating from 777 to 20, so no data
        bridge.migrate(
            l1Address,
            target,
            msg.sender,
            DecimalConverter.from777to20(amount, l1Decimals),
            ''
        );
    }
}
