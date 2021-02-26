// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.6.11;

import "arb-bridge-eth/contracts/libraries/Cloneable.sol";
import "./open-zeppelin/OZERC777.sol";
import "../libraries/DecimalConverter.sol";
import "./IArbToken.sol";
import "./ArbTokenBridge.sol";

contract StandardArbERC777 is ERC777, Cloneable, IArbToken {
    ArbTokenBridge public bridge;
    address public l1Address;
    uint8 public l1Decimals;

    modifier onlyBridge {
        require(msg.sender == address(bridge), "ONLY_BRIDGE");
        _;
    }

    function initialize(address _bridge, address _l1Address, uint8 _decimals) external override {
        require(address(bridge) != address(0), "ALREADY_INIT");
        bridge = ArbTokenBridge(_bridge);
        l1Address = _l1Address;

        require(_decimals <= 18);
        l1Decimals = _decimals;
        _granularity = 10 ** uint256(18 - _decimals);
    }

    function updateInfo(string memory newName, string memory newSymbol) public override onlyBridge {
        if (bytes(newName).length != 0) {
            _name = newName;
        }
        if (bytes(newSymbol).length != 0) {
            _symbol = newSymbol;
        }
    }

    function bridgeMint(address account, uint256 amount) external override onlyBridge {
        _mint(account, DecimalConverter.from20to777(l1Decimals, amount), '', '');
    }

    function withdraw(address destination, uint256 amount) external {
        _burn(msg.sender, amount, '', '');
        bridge.withdraw(l1Address, destination, DecimalConverter.from777to20(l1Decimals, amount));
    }

    function migrate(address target, uint256 amount) external {
        _burn(msg.sender, amount, '', '');
        bridge.migrate(l1Address, target, msg.sender, amount);
    }
}