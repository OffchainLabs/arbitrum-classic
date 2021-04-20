---
id: Special_Features
title: Special Features
sidebar_label: qq
---

### Address Registry

Every Arbitrum chain includes a native, pre-compiled Address Table Registry contract. This contract allows users to register an address, mapping it to some index, which can thenceforth be used to retrieve the address, thereby saving bytes of calldata.

The interface is here:

```sol

/** @title Precompiled contract that exists in every Arbitrum chain at 0x0000000000000000000000000000000000000066.
* Allows registering / retrieving addresses at uint indices, saving calldata.
*/
interface ArbAddressTable {
    /**
    * @notice Register an address in the address table
    * @param addr address to register
    * @return index of the address (existing index, or newly created index if not already registered)
    */
    function register(address addr) external returns(uint);

    /**
    * @param addr address to lookup
    * @return index of an address in the address table (revert if address isn't in the table)
    */
    function lookup(address addr) external view returns(uint);

    /**
    * @notice Check whether an address exists in the address table
    * @param addr address to check for presence in table
    * @return true if address is in table
    */
    function addressExists(address addr) external view returns(bool);

    /**
    * @return size of address table (= first unused index)
     */
    function size() external view returns(uint);

    /**
    * @param index index to lookup address
    * @return address at a given index in address table (revert if index is beyond end of table)
    */
    function lookupIndex(uint index) external view returns(address);

    /**
    * @notice read a compressed address from a bytes buffer
    * @param buf bytes buffer containing an address
    * @param offset offset of target address
    * @return resulting address and updated offset into the buffer (revert if buffer is too short)
    */
    function decompress(bytes calldata buf, uint offset) external pure returns(address, uint);

    /**
    * @notice compress an address and return the result
    * @param addr address to comppress
    * @return compressed address bytes
    */
    function compress(address addr) external returns(bytes memory);
}

```

For example usage, see our [Arbiswap Demo](https://github.com/OffchainLabs/Arbiswap_V2_mono/blob/5b7c38ebbc97bf1784c23526b9b75879cd053cdf/packages/other_contracts/contracts/UniswapV2Router02.sol#L736).

### Parameter Byte Serialization

Generally speaking, L1 calldata will the primary contributor to gas cost for Arbitrum transactions. Thus, when you're looking to gas-optimize your contract for Arbitrum, if you can minimize the amount of calldata it uses, you probably should!

One way of doing so broadly applicable to most contracts is to replace a method's parameters with a serialized byte array and have the contract deserialize the data.

[arb-ts](qqq.md) offers convenience methods for this client side serialization (as well as interacting with the Address table).

For example usage, see our Arbiswap Demo:

- [Contract](https://github.com/OffchainLabs/Arbiswap_V2_mono/blob/5b7c38ebbc97bf1784c23526b9b75879cd053cdf/packages/other_contracts/contracts/UniswapV2Router02.sol#L121)
- [Client Side](https://github.com/OffchainLabs/Arbiswap_V2_mono/blob/5b7c38ebbc97bf1784c23526b9b75879cd053cdf/packages/uniswap-interface/src/hooks/useSwapCallback.ts#L59)

### BLS Signatures
