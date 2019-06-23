pragma solidity >=0.4.21 <0.6.0;

interface ArbSys {
	function sendERC20(bytes32 dest, address tokenAddress, uint256 amount) external;
	function sendERC721(bytes32 dest, address tokenAddress, uint256 id) external;
}