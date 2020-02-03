/*

  Copyright 2019 ZeroEx Intl.

  Licensed under the Apache License, Version 2.0 (the "License");
  you may not use this file except in compliance with the License.
  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

  Unless required by applicable law or agreed to in writing, software
  distributed under the License is distributed on an "AS IS" BASIS,
  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
  See the License for the specific language governing permissions and
  limitations under the License.

*/

pragma solidity 0.5.15;

import "../../../../utils/contracts/src/LibEIP712.sol";


contract LibEIP712CoordinatorDomain {

    // EIP712 Domain Name value for the Coordinator
    string constant public EIP712_COORDINATOR_DOMAIN_NAME = "0x Protocol Coordinator";

    // EIP712 Domain Version value for the Coordinator
    string constant public EIP712_COORDINATOR_DOMAIN_VERSION = "3.0.0";

    // Hash of the EIP712 Domain Separator data for the Coordinator
    // solhint-disable-next-line var-name-mixedcase
    bytes32 public EIP712_COORDINATOR_DOMAIN_HASH;

    /// @param chainId Chain ID of the network this contract is deployed on.
    /// @param verifyingContractAddressIfExists Address of the verifying contract (null if the address of this contract)
    constructor (
        uint256 chainId,
        address verifyingContractAddressIfExists
    )
        public
    {
        address verifyingContractAddress = verifyingContractAddressIfExists == address(0)
            ? address(this)
            : verifyingContractAddressIfExists;
        EIP712_COORDINATOR_DOMAIN_HASH = LibEIP712.hashEIP712Domain(
            EIP712_COORDINATOR_DOMAIN_NAME,
            EIP712_COORDINATOR_DOMAIN_VERSION,
            chainId,
            verifyingContractAddress
        );
    }

    /// @dev Calculates EIP712 encoding for a hash struct in the EIP712 domain
    ///      of this contract.
    /// @param hashStruct The EIP712 hash struct.
    /// @return result EIP712 hash applied to this EIP712 Domain.
    function _hashEIP712CoordinatorMessage(bytes32 hashStruct)
        internal
        view
        returns (bytes32 result)
    {
        return LibEIP712.hashEIP712Message(EIP712_COORDINATOR_DOMAIN_HASH, hashStruct);
    }
}
