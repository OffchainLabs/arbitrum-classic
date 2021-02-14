// SPDX-License-Identifier: MIT

// Taken from https://github.com/optionality/clone-factory/blob/master/contracts/CloneFactory.sol

pragma solidity ^0.6.11;

/*
The MIT License (MIT)
Copyright (c) 2018 Murray Software, LLC.
Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:
The above copyright notice and this permission notice shall be included
in all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
*/
//solhint-disable max-line-length
//solhint-disable no-inline-assembly

import "./ICloneable.sol";
import "@openzeppelin/contracts/proxy/Clones.sol";

contract CloneFactory {
    using Clones for address;
    string private constant CLONE_MASTER = "CLONE_MASTER";

    function createClone(ICloneable target) internal returns (address result) {
        require(target.isMaster(), CLONE_MASTER);
        result = address(target).clone();
    }

    function create2Clone(ICloneable target, bytes32 salt) internal returns (address result) {
        require(target.isMaster(), CLONE_MASTER);
        result = address(target).cloneDeterministic(salt);
    }

    function calculateCreate2CloneAddress(ICloneable target, bytes32 salt) internal view returns (address calculatedAddress) {
        calculatedAddress = address(target).predictDeterministicAddress(salt, address(this));
    }
}
