// SPDX-License-Identifier: Apache-2.0

pragma solidity >=0.6.0 <0.8.0;

import "@openzeppelin/contracts/proxy/BeaconProxy.sol";
import "@openzeppelin/contracts/proxy/UpgradeableBeacon.sol";
import "@openzeppelin/contracts/utils/Create2.sol";

interface ProxySetter {
    function beacon() external view returns (address);
}

contract ClonableBeaconProxy is BeaconProxy {
    constructor() public BeaconProxy(ProxySetter(msg.sender).beacon(), "") {}
}

contract BeaconProxyFactory is ProxySetter {
    bytes32 public constant cloneableProxyHash = keccak256(type(ClonableBeaconProxy).creationCode);

    /**
     * @notice utility function used in ClonableBeaconProxy.
     * @dev this method makes it possible to use ClonableBeaconProxy.creationCode without encoding constructor parameters
     * @return the beacon to be used by the proxy contract.
     */
    address public override beacon;

    function initialize(address _beacon) external {
        require(_beacon != address(0), "INVALID_BEACON");
        require(beacon == address(0), "ALREADY_INIT");
        beacon = _beacon;
    }

    function getSalt(address user, bytes32 userSalt) public pure returns (bytes32) {
        return keccak256(abi.encode(user, userSalt));
    }

    function createProxy(bytes32 userSalt) external returns (address) {
        // consumer should check if createdContract == address(0), ie fail deploy
        bytes32 salt = getSalt(msg.sender, userSalt);
        address createdContract = address(new ClonableBeaconProxy{ salt: salt }());
        return createdContract;
    }

    function calculateExpectedAddress(address user, bytes32 userSalt)
        public
        view
        returns (address)
    {
        bytes32 salt = getSalt(user, userSalt);
        return Create2.computeAddress(salt, cloneableProxyHash, address(this));
    }

    function calculateExpectedAddress(bytes32 salt) public view returns (address) {
        return Create2.computeAddress(salt, cloneableProxyHash, address(this));
    }
}
