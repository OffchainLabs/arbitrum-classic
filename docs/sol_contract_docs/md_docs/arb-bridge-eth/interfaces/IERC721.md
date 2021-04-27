---
title: IERC721.sol Spec
---

ERC-721 Non-Fungible Token Standard

See https://eips.ethereum.org/EIPS/eip-721
Note: the ERC-165 identifier for this interface is 0x80ac58cd.

### `balanceOf(address _owner) → uint256` (external)

Count all NFTs assigned to an owner

NFTs assigned to the zero address are considered invalid, and this
function throws for queries about the zero address.

- `_owner`: An address for whom to query the balance

**Returns**: The: number of NFTs owned by `_owner`, possibly zero

### `ownerOf(uint256 _tokenId) → address` (external)

Find the owner of an NFT

NFTs assigned to zero address are considered invalid, and queries
about them do throw.

- `_tokenId`: The identifier for an NFT

**Returns**: The: address of the owner of the NFT

### `safeTransferFrom(address _from, address _to, uint256 _tokenId, bytes data)` (external)

Transfers the ownership of an NFT from one address to another address

Throws unless `msg.sender` is the current owner, an authorized
operator, or the approved address for this NFT. Throws if `_from` is
not the current owner. Throws if `_to` is the zero address. Throws if
`_tokenId` is not a valid NFT. When transfer is complete, this function
checks if `_to` is a smart contract (code size > 0). If so, it calls
`onERC721Received` on `_to` and throws if the return value is not
`bytes4(keccak256("onERC721Received(address,address,uint256,bytes)"))`.

- `_from`: The current owner of the NFT

- `_to`: The new owner

- `_tokenId`: The NFT to transfer

- `data`: Additional data with no specified format, sent in call to `_to`

### `safeTransferFrom(address _from, address _to, uint256 _tokenId)` (external)

Transfers the ownership of an NFT from one address to another address

This works identically to the other function with an extra data parameter,
except this function just sets data to "".

- `_from`: The current owner of the NFT

- `_to`: The new owner

- `_tokenId`: The NFT to transfer

### `transferFrom(address _from, address _to, uint256 _tokenId)` (external)

Transfer ownership of an NFT -- THE CALLER IS RESPONSIBLE
TO CONFIRM THAT `_to` IS CAPABLE OF RECEIVING NFTS OR ELSE
THEY MAY BE PERMANENTLY LOST

Throws unless `msg.sender` is the current owner, an authorized
operator, or the approved address for this NFT. Throws if `_from` is
not the current owner. Throws if `_to` is the zero address. Throws if
`_tokenId` is not a valid NFT.

- `_from`: The current owner of the NFT

- `_to`: The new owner

- `_tokenId`: The NFT to transfer

### `approve(address _approved, uint256 _tokenId)` (external)

Change or reaffirm the approved address for an NFT

The zero address indicates there is no approved address.
Throws unless `msg.sender` is the current NFT owner, or an authorized
operator of the current owner.

- `_approved`: The new approved NFT controller

- `_tokenId`: The NFT to approve

### `setApprovalForAll(address _operator, bool _approved)` (external)

Enable or disable approval for a third party ("operator") to manage
all of `msg.sender`'s assets

Emits the ApprovalForAll event. The contract MUST allow
multiple operators per owner.

- `_operator`: Address to add to the set of authorized operators

- `_approved`: True if the operator is approved, false to revoke approval

### `getApproved(uint256 _tokenId) → address` (external)

Get the approved address for a single NFT

Throws if `_tokenId` is not a valid NFT.

- `_tokenId`: The NFT to find the approved address for

**Returns**: The: approved address for this NFT, or the zero address if there is none

### `isApprovedForAll(address _owner, address _operator) → bool` (external)

Query if an address is an authorized operator for another address

- `_owner`: The address that owns the NFTs

- `_operator`: The address that acts on behalf of the owner

**Returns**: True: if `_operator` is an approved operator for `_owner`, false otherwise

### `Transfer(address _from, address _to, uint256 _tokenId)`

This emits when ownership of any NFT changes by any mechanism.
This event emits when NFTs are created (`from` == 0) and destroyed
(`to` == 0). Exception: during contract creation, any number of NFTs
may be created and assigned without emitting Transfer. At the time of
any transfer, the approved address for that NFT (if any) is reset to none.

### `Approval(address _owner, address _approved, uint256 _tokenId)`

This emits when the approved address for an NFT is changed or
reaffirmed. The zero address indicates there is no approved address.
When a Transfer event emits, this also indicates that the approved
address for that NFT (if any) is reset to none.

### `ApprovalForAll(address _owner, address _operator, bool _approved)`

This emits when an operator is enabled or disabled for an owner.
The operator can manage all NFTs of the owner.
