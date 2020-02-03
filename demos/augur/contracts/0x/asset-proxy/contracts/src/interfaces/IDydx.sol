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


interface IDydx {

    /// @dev Represents the unique key that specifies an account
    struct AccountInfo {
        address owner;  // The address that owns the account
        uint256 number; // A nonce that allows a single address to control many accounts
    }

    enum ActionType {
        Deposit,   // supply tokens
        Withdraw,  // borrow tokens
        Transfer,  // transfer balance between accounts
        Buy,       // buy an amount of some token (externally)
        Sell,      // sell an amount of some token (externally)
        Trade,     // trade tokens against another account
        Liquidate, // liquidate an undercollateralized or expiring account
        Vaporize,  // use excess tokens to zero-out a completely negative account
        Call       // send arbitrary data to an address
    }

    /// @dev Arguments that are passed to Solo in an ordered list as part of a single operation.
    /// Each ActionArgs has an actionType which specifies which action struct that this data will be
    /// parsed into before being processed.
    struct ActionArgs {
        ActionType actionType;
        uint256 accountId;
        AssetAmount amount;
        uint256 primaryMarketId;
        uint256 secondaryMarketId;
        address otherAddress;
        uint256 otherAccountId;
        bytes data;
    }

    enum AssetDenomination {
        Wei, // the amount is denominated in wei
        Par  // the amount is denominated in par
    }

    enum AssetReference {
        Delta, // the amount is given as a delta from the current value
        Target // the amount is given as an exact number to end up at
    }

    struct AssetAmount {
        bool sign; // true if positive
        AssetDenomination denomination;
        AssetReference ref;
        uint256 value;
    }

    /// @dev The main entry-point to Solo that allows users and contracts to manage accounts.
    ///      Take one or more actions on one or more accounts. The msg.sender must be the owner or
    ///      operator of all accounts except for those being liquidated, vaporized, or traded with.
    ///      One call to operate() is considered a singular "operation". Account collateralization is
    ///      ensured only after the completion of the entire operation.
    /// @param  accounts  A list of all accounts that will be used in this operation. Cannot contain
    ///                   duplicates. In each action, the relevant account will be referred-to by its
    ///                   index in the list.
    /// @param  actions   An ordered list of all actions that will be taken in this operation. The
    ///                   actions will be processed in order.
    function operate(
        AccountInfo[] calldata accounts,
        ActionArgs[] calldata actions
    )
        external;
}
