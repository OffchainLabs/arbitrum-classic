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

import "./libs/LibSafeDowncast.sol";
import "./immutable/MixinStorage.sol";
import "./immutable/MixinConstants.sol";
import "./interfaces/IStorageInit.sol";
import "./interfaces/IStakingProxy.sol";


/// #dev The 0x Staking contract.
contract StakingProxy is
    IStakingProxy,
    MixinStorage,
    MixinConstants
{
    using LibSafeDowncast for uint256;

    /// @dev Constructor.
    /// @param _stakingContract Staking contract to delegate calls to.
    constructor(address _stakingContract)
        public
        MixinStorage()
    {
        // Deployer address must be authorized in order to call `init`
        _addAuthorizedAddress(msg.sender);

        // Attach the staking contract and initialize state
        _attachStakingContract(_stakingContract);

        // Remove the sender as an authorized address
        _removeAuthorizedAddressAtIndex(msg.sender, 0);
    }

    /// @dev Delegates calls to the staking contract, if it is set.
    function ()
        external
        payable
    {
        // Sanity check that we have a staking contract to call
        address stakingContract_ = stakingContract;
        if (stakingContract_ == NIL_ADDRESS) {
            revert();
        }

        // Call the staking contract with the provided calldata.
        (bool success, bytes memory returnData) = stakingContract_.delegatecall(msg.data);

        // Revert on failure or return on success.
        assembly {
            switch success
            case 0 {
                revert(add(0x20, returnData), mload(returnData))
            }
            default {
                return(add(0x20, returnData), mload(returnData))
            }
        }
    }

    /// @dev Attach a staking contract; future calls will be delegated to the staking contract.
    /// Note that this is callable only by an authorized address.
    /// @param _stakingContract Address of staking contract.
    function attachStakingContract(address _stakingContract)
        external
        onlyAuthorized
    {
        _attachStakingContract(_stakingContract);
    }

    /// @dev Detach the current staking contract.
    /// Note that this is callable only by an authorized address.
    function detachStakingContract()
        external
        onlyAuthorized
    {
        stakingContract = NIL_ADDRESS;
        emit StakingContractDetachedFromProxy();
    }

    /// @dev Batch executes a series of calls to the staking contract.
    /// @param data An array of data that encodes a sequence of functions to
    ///             call in the staking contracts.
    function batchExecute(bytes[] calldata data)
        external
        returns (bytes[] memory batchReturnData)
    {
        // Initialize commonly used variables.
        bool success;
        bytes memory returnData;
        uint256 dataLength = data.length;
        batchReturnData = new bytes[](dataLength);
        address staking = stakingContract;

        // Ensure that a staking contract has been attached to the proxy.
        if (staking == NIL_ADDRESS) {
            revert();
        }

        // Execute all of the calls encoded in the provided calldata.
        for (uint256 i = 0; i != dataLength; i++) {
            // Call the staking contract with the provided calldata.
            (success, returnData) = staking.delegatecall(data[i]);

            // Revert on failure.
            if (!success) {
                assembly {
                    revert(add(0x20, returnData), mload(returnData))
                }
            }

            // Add the returndata to the batch returndata.
            batchReturnData[i] = returnData;
        }

        return batchReturnData;
    }

    /// @dev Asserts that an epoch is between 5 and 30 days long.
    //       Asserts that 0 < cobb douglas alpha value <= 1.
    //       Asserts that a stake weight is <= 100%.
    //       Asserts that pools allow >= 1 maker.
    //       Asserts that all addresses are initialized.
    function assertValidStorageParams()
        public
        view
    {
        // Epoch length must be between 5 and 30 days long
        uint256 _epochDurationInSeconds = epochDurationInSeconds;
        if (_epochDurationInSeconds < 5 days || _epochDurationInSeconds > 30 days) {
            revert();
        }

        // Alpha must be 0 < x <= 1
        uint32 _cobbDouglasAlphaDenominator = cobbDouglasAlphaDenominator;
        if (cobbDouglasAlphaNumerator > _cobbDouglasAlphaDenominator || _cobbDouglasAlphaDenominator == 0) {
            revert();
        }

        // Weight of delegated stake must be <= 100%
        if (rewardDelegatedStakeWeight > PPM_DENOMINATOR) {
            revert();
        }

        // Minimum stake must be > 1
        if (minimumPoolStake < 2) {
            revert();
        }
    }

    /// @dev Attach a staking contract; future calls will be delegated to the staking contract.
    /// @param _stakingContract Address of staking contract.
    function _attachStakingContract(address _stakingContract)
        internal
    {
        // Attach the staking contract
        stakingContract = _stakingContract;
        emit StakingContractAttachedToProxy(_stakingContract);

        // Call `init()` on the staking contract to initialize storage.
        (bool didInitSucceed, bytes memory initReturnData) = stakingContract.delegatecall(
            abi.encodeWithSelector(IStorageInit(0).init.selector)
        );

        if (!didInitSucceed) {
            assembly {
                revert(add(initReturnData, 0x20), mload(initReturnData))
            }
        }

        // Assert initialized storage values are valid
        assertValidStorageParams();
    }
}
