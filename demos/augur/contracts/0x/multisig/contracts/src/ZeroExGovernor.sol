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

import "./MultiSigWalletWithTimeLock.sol";
import "../../../utils/contracts/src/LibBytes.sol";
import "../../../utils/contracts/src/LibSafeMath.sol";


contract ZeroExGovernor is
    MultiSigWalletWithTimeLock
{
    using LibBytes for bytes;
    using LibSafeMath for uint256;

    struct TimeLock {
        bool hasCustomTimeLock;
        uint128 secondsTimeLocked;
    }

    event FunctionCallTimeLockRegistration(
        bytes4 functionSelector,
        address destination,
        bool hasCustomTimeLock,
        uint128 newSecondsTimeLocked
    );

    // Function selector => destination => seconds timelocked
    mapping (bytes4 => mapping (address => TimeLock)) public functionCallTimeLocks;

    /// @dev Contract constructor sets initial owners, required number of confirmations, and default time lock
    ///      It will also register unique timelocks for each passed in function selector / destination combo.
    /// @param _functionSelectors Array of function selectors for registered functions.
    /// @param _destinations Array of destinations for registered function calls.
    /// @param _functionCallTimeLockSeconds Array of seconds that each registered function call will be timelocked.
    /// @param _owners List of initial owners.
    /// @param _required Number of required confirmations.
    /// @param _defaultSecondsTimeLocked Default duration in seconds needed after a transaction is confirmed to become executable.
    constructor (
        bytes4[] memory _functionSelectors,
        address[] memory _destinations,
        uint128[] memory _functionCallTimeLockSeconds,
        address[] memory _owners,
        uint256 _required,
        uint256 _defaultSecondsTimeLocked
    )
        public
        MultiSigWalletWithTimeLock(
            _owners,
            _required,
            _defaultSecondsTimeLocked
        )
    {
        uint256 length = _functionSelectors.length;
        require(
            length == _destinations.length && length == _functionCallTimeLockSeconds.length,
            "EQUAL_LENGTHS_REQUIRED"
        );

        // Register function timelocks
        for (uint256 i = 0; i != length; i++) {
            _registerFunctionCall(
                true,  // all functions registered in constructor are assumed to have a custom timelock
                _functionSelectors[i],
                _destinations[i],
                _functionCallTimeLockSeconds[i]
            );
        }
    }

    /// @dev Registers a custom timelock to a specific function selector / destination combo
    /// @param hasCustomTimeLock True if timelock is custom.
    /// @param functionSelector 4 byte selector of registered function.
    /// @param destination Address of destination where function will be called.
    /// @param newSecondsTimeLocked Duration in seconds needed after a transaction is confirmed to become executable.
    function registerFunctionCall(
        bool hasCustomTimeLock,
        bytes4 functionSelector,
        address destination,
        uint128 newSecondsTimeLocked
    )
        external
        onlyWallet
    {
        _registerFunctionCall(
            hasCustomTimeLock,
            functionSelector,
            destination,
            newSecondsTimeLocked
        );
    }

    /// @dev Allows anyone to execute a confirmed transaction.
    ///      Transactions *must* encode the values with the signature "bytes[] data, address[] destinations, uint256[] values"
    ///      The `destination` and `value` fields of the transaction in storage are ignored.
    ///      All function calls must be successful or the entire call will revert.
    /// @param transactionId Transaction ID.
    function executeTransaction(uint256 transactionId)
        public
        notExecuted(transactionId)
        fullyConfirmed(transactionId)
    {
        Transaction storage transaction = transactions[transactionId];
        transaction.executed = true;

        // Decode batch transaction data from transaction.data
        // `destination` and `value` fields of transaction are ignored
        // Note that `destination` must be non-0, or the transaction cannot be submitted
        // solhint-disable
        (
            bytes[] memory data,
            address[] memory destinations,
            uint256[] memory values
        ) = abi.decode(
            transaction.data,
            (bytes[], address[], uint256[])
        );
        // solhint-enable

        // Ensure lengths of array properties are equal
        uint256 length = data.length;
        require(
            length == destinations.length && length == values.length,
            "EQUAL_LENGTHS_REQUIRED"
        );

        uint256 transactionConfirmationTime = confirmationTimes[transactionId];
        for (uint i = 0; i != length; i++) {
            // Ensure that each function call is past its timelock
            _assertValidFunctionCall(
                transactionConfirmationTime,
                data[i],
                destinations[i]
            );
            // Call each function
            // solhint-disable-next-line avoid-call-value
            (bool didSucceed,) = destinations[i].call.value(values[i])(data[i]);
            // Ensure that function call was successful
            require(
                didSucceed,
                "FAILED_EXECUTION"
            );
        }
        emit Execution(transactionId);
    }

    /// @dev Registers a custom timelock to a specific function selector / destination combo
    /// @param hasCustomTimeLock True if timelock is custom.
    /// @param functionSelector 4 byte selector of registered function.
    /// @param destination Address of destination where function will be called.
    /// @param newSecondsTimeLocked Duration in seconds needed after a transaction is confirmed to become executable.
    function _registerFunctionCall(
        bool hasCustomTimeLock,
        bytes4 functionSelector,
        address destination,
        uint128 newSecondsTimeLocked
    )
        internal
    {
        // Clear the previous secondsTimeLocked if custom timelock not used
        uint128 _secondsTimeLocked = hasCustomTimeLock ? newSecondsTimeLocked : 0;
        TimeLock memory timeLock = TimeLock({
            hasCustomTimeLock: hasCustomTimeLock,
            secondsTimeLocked: _secondsTimeLocked
        });
        functionCallTimeLocks[functionSelector][destination] = timeLock;
        emit FunctionCallTimeLockRegistration(
            functionSelector,
            destination,
            hasCustomTimeLock,
            _secondsTimeLocked
        );
    }

    /// @dev Ensures that the function call has past its timelock.
    /// @param transactionConfirmationTime Timestamp at which transaction was fully confirmed.
    /// @param data Function calldata.
    /// @param destination Address to call function on.
    function _assertValidFunctionCall(
        uint256 transactionConfirmationTime,
        bytes memory data,
        address destination
    )
        internal
        view
    {
        bytes4 functionSelector = data.readBytes4(0);
        TimeLock memory timeLock = functionCallTimeLocks[functionSelector][destination];
        // solhint-disable not-rely-on-time
        if (timeLock.hasCustomTimeLock) {
            require(
                block.timestamp >= transactionConfirmationTime.safeAdd(timeLock.secondsTimeLocked),
                "CUSTOM_TIME_LOCK_INCOMPLETE"
            );
        } else {
            require(
                block.timestamp >= transactionConfirmationTime.safeAdd(secondsTimeLocked),
                "DEFAULT_TIME_LOCK_INCOMPLETE"
            );
        }
        // solhint-enable not-rely-on-time
    }
}
