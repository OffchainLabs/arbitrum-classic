pragma solidity 0.5.15;


interface IRepExchange {
    function initialize(address _augurAddress, address _token) external;
    function pokePrice() external returns (uint256);
}