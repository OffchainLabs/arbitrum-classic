// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.6.11;

import "arb-bridge-eth/contracts/libraries/Cloneable.sol";
import "./open-zeppelin/OZERC777.sol";
import "../libraries/DecimalConverter.sol";
import "./IArbToken.sol";
import "./ArbTokenBridge.sol";
import "../libraries/BytesParser.sol";

contract StandardArbERC777 is OZERC777, Cloneable, IArbToken {
    ArbTokenBridge public bridge;
    address public l1Address;
    uint8 public l1Decimals;

    modifier onlyBridge {
        require(msg.sender == address(bridge), "ONLY_BRIDGE");
        _;
    }

    function bridgeInit(address _l1Address, bytes memory _data) external override returns (bool) {
        require(address(l1Address) == address(0), "Already inited");
        bridge = ArbTokenBridge(msg.sender);
        l1Address = _l1Address;

        (bytes memory name, bytes memory symbol, bytes memory decimals) =
            abi.decode(_data, (bytes, bytes, bytes));
        // what if decode reverts? shouldn't as this is encoded by L1 contract

        _name = BytesParserWithDefault.toString(name, "");
        _symbol = BytesParserWithDefault.toString(symbol, "");

        uint8 _decimals = BytesParserWithDefault.toUint8(decimals, 18);
        // require(_decimals <= 18, "Decimals must be less than or equal to 18");
        l1Decimals = _decimals;
        OZERC777.initialize(DecimalConverter.decimalsToGranularity(_decimals));
        return true;
    }

    function bridgeMint(
        address account,
        uint256 amount,
        bytes memory data
    ) external override onlyBridge {
        _mint(account, amount, data, "");
    }

    function withdraw(address destination, uint256 amount) external {
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
            ""
        );
    }
}
