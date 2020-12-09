pragma solidity >=0.4.21 <0.7.0;

interface ArbSys {
    // Send given amount of ERC-20 tokens to dest with token contract sender.
    // This is safe to freely call since the sender is authenticated and thus
    // you can only send fake tokens, not steal real ones
    function withdrawERC20(address dest, uint256 amount) external;

    // Send given ERC-721 token to dest with token contract sender.
    // This is safe by the above arguement
    function withdrawERC721(address dest, uint256 id) external;

    // Send given amount of Eth to dest with from sender.
    function withdrawEth(address dest) external payable;

    // Return the number of transactions issued by the given external account
    // or the account sequence number of the given contract
    function getTransactionCount(address account) external view returns (uint256);

    // Return the value of the storage slot for the given account at the given index
    // This function is only callable from address 0 to prevent contracts from being
    // able to call it
    function getStorageAt(address account, uint256 index) external view returns (uint256);

    event EthWithdrawal(address indexed destAddr, uint256 amount);
    event ERC20Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 amount);
    event ERC721Withdrawal(address indexed destAddr, address indexed tokenAddr, uint256 indexed id);
}
