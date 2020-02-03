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
pragma experimental ABIEncoderV2;

import "./IZrxVault.sol";
import "./IStructs.sol";


interface IStorage {

    function stakingContract()
        external
        view
        returns (address);

    function lastPoolId()
        external
        view
        returns (bytes32);

    function numMakersByPoolId(bytes32 poolId)
        external
        view
        returns (uint256);

    function currentEpoch()
        external
        view
        returns (uint256);

    function currentEpochStartTimeInSeconds()
        external
        view
        returns (uint256);

    function protocolFeesThisEpochByPool(bytes32 poolId)
        external
        view
        returns (uint256);

    function validExchanges(address exchangeAddress)
        external
        view
        returns (bool);

    function epochDurationInSeconds()
        external
        view
        returns (uint256);

    function rewardDelegatedStakeWeight()
        external
        view
        returns(uint32);

    function minimumPoolStake()
        external
        view
        returns (uint256);

    function cobbDouglasAlphaNumerator()
        external
        view
        returns (uint32);

    function cobbDouglasAlphaDenominator()
        external
        view
        returns (uint32);
}
