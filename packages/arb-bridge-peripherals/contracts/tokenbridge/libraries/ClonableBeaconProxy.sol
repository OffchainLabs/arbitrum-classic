// SPDX-License-Identifier: Apache-2.0

pragma solidity >=0.6.0 <0.8.0;

import "@openzeppelin/contracts/proxy/BeaconProxy.sol";

contract ClonableBeaconProxy is BeaconProxy {
    constructor(address beacon) public BeaconProxy(beacon, "") {}
}
