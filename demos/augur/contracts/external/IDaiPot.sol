pragma solidity 0.5.15;


contract IDaiPot {
    mapping (address => uint256) public pie;  // user Savings Dai
    uint256 public dsr;  // The Dai Savings Rate
    uint256 public chi;  // The Rate Accumulator

    function drip() public returns (uint256);
    function join(uint wad) public;
    function exit(uint wad) public;
}