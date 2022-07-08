// SPDX-License-Identifier: MIT
// Licensed under OpenZeppelin's license: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/8b778fa20d6d76340c5fac1ed66c80273f05b95a/LICENSE

import "../libraries/ProxyUtil.sol";

pragma solidity ^0.6.11;

contract NonDelegatingProxy {
	address private proxyTarget;

	function postUpgradeInit(address newProxyTarget) external {
		require(msg.sender == ProxyUtil.getProxyAdmin(), "NOT_PROXY_ADMIN");
		proxyTarget = newProxyTarget;
	}

	fallback() external {
		address target = proxyTarget;
		// Taken from https://github.com/OpenZeppelin/openzeppelin-contracts/blob/8b778fa20d6d76340c5fac1ed66c80273f05b95a/contracts/proxy/Proxy.sol
		// Modified to use staticcall instead of delegatecall
		assembly {
			// Copy msg.data. We take full control of memory in this inline assembly
            // block because it will not return to Solidity code. We overwrite the
            // Solidity scratch pad at memory position 0.
            calldatacopy(0, 0, calldatasize())

            // Call the implementation.
            // out and outsize are 0 because we don't know the size yet.
            let result := staticcall(gas(), target, 0, calldatasize(), 0, 0)

            // Copy the returned data.
            returndatacopy(0, 0, returndatasize())

            switch result
            // staticcall returns 0 on error.
            case 0 {
                revert(0, returndatasize())
            }
            default {
                return(0, returndatasize())
            }
		}
	}
}
