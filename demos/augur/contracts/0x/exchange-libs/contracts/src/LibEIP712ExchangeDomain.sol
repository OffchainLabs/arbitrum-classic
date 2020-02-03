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

import "../../../utils/contracts/src/LibEIP712.sol";


contract LibEIP712ExchangeDomain {

    // EIP712 Exchange Domain Name value
    string constant internal _EIP712_EXCHANGE_DOMAIN_NAME = "0x Protocol";

    // EIP712 Exchange Domain Version value
    string constant internal _EIP712_EXCHANGE_DOMAIN_VERSION = "3.0.0";

    // solhint-disable var-name-mixedcase
    /// @dev Hash of the EIP712 Domain Separator data
    /// @return 0 Domain hash.
    bytes32 public EIP712_EXCHANGE_DOMAIN_HASH;
    // solhint-enable var-name-mixedcase

    /// @param chainId Chain ID of the network this contract is deployed on.
    /// @param verifyingContractAddressIfExists Address of the verifying contract (null if the address of this contract)
    constructor (
        uint256 chainId,
        address verifyingContractAddressIfExists
    )
        public
    {
        address verifyingContractAddress = verifyingContractAddressIfExists == address(0) ? address(this) : verifyingContractAddressIfExists;
        EIP712_EXCHANGE_DOMAIN_HASH = LibEIP712.hashEIP712Domain(
            _EIP712_EXCHANGE_DOMAIN_NAME,
            _EIP712_EXCHANGE_DOMAIN_VERSION,
            chainId,
            verifyingContractAddress
        );
    }
}
