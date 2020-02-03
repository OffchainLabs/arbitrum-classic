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

import "../archive/MixinAuthorizable.sol";


contract ERC20Proxy is
    MixinAuthorizable
{
    // Id of this proxy.
    bytes4 constant internal PROXY_ID = bytes4(keccak256("ERC20Token(address)"));

    // solhint-disable-next-line payable-fallback
    function ()
        external
    {
        assembly {
            // The first 4 bytes of calldata holds the function selector
            let selector := and(calldataload(0), 0xffffffff00000000000000000000000000000000000000000000000000000000)

            // `transferFrom` will be called with the following parameters:
            // assetData Encoded byte array.
            // from Address to transfer asset from.
            // to Address to transfer asset to.
            // amount Amount of asset to transfer.
            // bytes4(keccak256("transferFrom(bytes,address,address,uint256)")) = 0xa85e59e4
            if eq(selector, 0xa85e59e400000000000000000000000000000000000000000000000000000000) {

                // To lookup a value in a mapping, we load from the storage location keccak256(k, p),
                // where k is the key left padded to 32 bytes and p is the storage slot
                let start := mload(64)
                mstore(start, and(caller, 0xffffffffffffffffffffffffffffffffffffffff))
                mstore(add(start, 32), authorized_slot)

                // Revert if authorized[msg.sender] == false
                if iszero(sload(keccak256(start, 64))) {
                    // Revert with `Error("SENDER_NOT_AUTHORIZED")`
                    mstore(0, 0x08c379a000000000000000000000000000000000000000000000000000000000)
                    mstore(32, 0x0000002000000000000000000000000000000000000000000000000000000000)
                    mstore(64, 0x0000001553454e4445525f4e4f545f415554484f52495a454400000000000000)
                    mstore(96, 0)
                    revert(0, 100)
                }

                // `transferFrom`.
                // The function is marked `external`, so no abi decodeding is done for
                // us. Instead, we expect the `calldata` memory to contain the
                // following:
                //
                // | Area     | Offset | Length  | Contents                            |
                // |----------|--------|---------|-------------------------------------|
                // | Header   | 0      | 4       | function selector                   |
                // | Params   |        | 4 * 32  | function parameters:                |
                // |          | 4      |         |   1. offset to assetData (*)        |
                // |          | 36     |         |   2. from                           |
                // |          | 68     |         |   3. to                             |
                // |          | 100    |         |   4. amount                         |
                // | Data     |        |         | assetData:                          |
                // |          | 132    | 32      | assetData Length                    |
                // |          | 164    | **      | assetData Contents                  |
                //
                // (*): offset is computed from start of function parameters, so offset
                //      by an additional 4 bytes in the calldata.
                //
                // (**): see table below to compute length of assetData Contents
                //
                // WARNING: The ABIv2 specification allows additional padding between
                //          the Params and Data section. This will result in a larger
                //          offset to assetData.

                // Asset data itself is encoded as follows:
                //
                // | Area     | Offset | Length  | Contents                            |
                // |----------|--------|---------|-------------------------------------|
                // | Header   | 0      | 4       | function selector                   |
                // | Params   |        | 1 * 32  | function parameters:                |
                // |          | 4      | 12 + 20 |   1. token address                  |

                // We construct calldata for the `token.transferFrom` ABI.
                // The layout of this calldata is in the table below.
                //
                // | Area     | Offset | Length  | Contents                            |
                // |----------|--------|---------|-------------------------------------|
                // | Header   | 0      | 4       | function selector                   |
                // | Params   |        | 3 * 32  | function parameters:                |
                // |          | 4      |         |   1. from                           |
                // |          | 36     |         |   2. to                             |
                // |          | 68     |         |   3. amount                         |

                /////// Read token address from calldata ///////
                // * The token address is stored in `assetData`.
                //
                // * The "offset to assetData" is stored at offset 4 in the calldata (table 1).
                //   [assetDataOffsetFromParams = calldataload(4)]
                //
                // * Notes that the "offset to assetData" is relative to the "Params" area of calldata;
                //   add 4 bytes to account for the length of the "Header" area (table 1).
                //   [assetDataOffsetFromHeader = assetDataOffsetFromParams + 4]
                //
                // * The "token address" is offset 32+4=36 bytes into "assetData" (tables 1 & 2).
                //   [tokenOffset = assetDataOffsetFromHeader + 36 = calldataload(4) + 4 + 36]
                let token := calldataload(add(calldataload(4), 40))

                /////// Setup Header Area ///////
                // This area holds the 4-byte `transferFrom` selector.
                // Any trailing data in transferFromSelector will be
                // overwritten in the next `mstore` call.
                mstore(0, 0x23b872dd00000000000000000000000000000000000000000000000000000000)

                /////// Setup Params Area ///////
                // We copy the fields `from`, `to` and `amount` in bulk
                // from our own calldata to the new calldata.
                calldatacopy(4, 36, 96)

                /////// Call `token.transferFrom` using the calldata ///////
                let success := call(
                    gas,            // forward all gas
                    token,          // call address of token contract
                    0,              // don't send any ETH
                    0,              // pointer to start of input
                    100,            // length of input
                    0,              // write output over input
                    32              // output size should be 32 bytes
                )

                /////// Check return data. ///////
                // If there is no return data, we assume the token incorrectly
                // does not return a bool. In this case we expect it to revert
                // on failure, which was handled above.
                // If the token does return data, we require that it is a single
                // nonzero 32 bytes value.
                // So the transfer succeeded if the call succeeded and either
                // returned nothing, or returned a non-zero 32 byte value.
                success := and(success, or(
                    iszero(returndatasize),
                    and(
                        eq(returndatasize, 32),
                        gt(mload(0), 0)
                    )
                ))
                if success {
                    return(0, 0)
                }

                // Revert with `Error("TRANSFER_FAILED")`
                mstore(0, 0x08c379a000000000000000000000000000000000000000000000000000000000)
                mstore(32, 0x0000002000000000000000000000000000000000000000000000000000000000)
                mstore(64, 0x0000000f5452414e534645525f4641494c454400000000000000000000000000)
                mstore(96, 0)
                revert(0, 100)
            }

            // Revert if undefined function is called
            revert(0, 0)
        }
    }

    /// @dev Gets the proxy id associated with the proxy address.
    /// @return Proxy id.
    function getProxyId()
        external
        pure
        returns (bytes4)
    {
        return PROXY_ID;
    }
}
