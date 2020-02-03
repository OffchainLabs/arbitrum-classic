pragma solidity 0.5.15;

contract IRepSymbol {
    function getRepSymbol(address _augur, address _universe) external view returns (string memory);
}