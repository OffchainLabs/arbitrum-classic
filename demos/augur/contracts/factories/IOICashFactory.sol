pragma solidity 0.5.15;

import 'ROOT/IAugur.sol';
import 'ROOT/reporting/IOICash.sol';


contract IOICashFactory {
    function createOICash(IAugur _augur) public returns (IOICash);
}
