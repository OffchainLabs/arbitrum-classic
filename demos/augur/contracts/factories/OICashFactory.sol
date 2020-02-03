pragma solidity 0.5.15;


import 'ROOT/libraries/CloneFactory.sol';
import 'ROOT/reporting/IUniverse.sol';
import 'ROOT/reporting/IOICash.sol';
import 'ROOT/IAugur.sol';


/**
 * @title OI Cash Factory
 * @notice A Factory contract to create OI Cash Token contracts
 * @dev Should not be used directly. Only intended to be used by Universe contracts
 */
contract OICashFactory is CloneFactory {
    function createOICash(IAugur _augur) public returns (IOICash) {
        IUniverse _universe = IUniverse(msg.sender);
        IOICash _openInterestCash = IOICash(createClone(_augur.lookup("OICash")));
        _openInterestCash.initialize(_augur, _universe);
        return _openInterestCash;
    }
}