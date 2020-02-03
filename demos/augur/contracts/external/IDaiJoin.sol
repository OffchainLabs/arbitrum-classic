pragma solidity 0.5.15;


contract IDaiJoin {
    uint256 public live = 1;  // Access Flag
    function join(address urn, uint wad) public;
    function exit(address usr, uint wad) public;
}