pragma solidity ^0.6.11;

contract ArbSysMock {
    function sendTxToL1(address destination, bytes calldata calldataForL1) external payable returns(uint){
        return 0;
    }
}