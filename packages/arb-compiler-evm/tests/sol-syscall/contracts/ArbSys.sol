pragma solidity >=0.4.21 <0.6.0;

interface ArbSys {
	function sendERC20(address dest, address tokenAddress, uint256 amount) external;
	function sendERC721(address dest, address tokenAddress, uint256 id) external;
}
