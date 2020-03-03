pragma solidity 0.5.15;


import '../libraries/CloneFactory.sol';
import '../reporting/IUniverse.sol';
import '../reporting/IOICash.sol';
import '../IAugur.sol';


/**
 * @title OI Cash Factory
 * @notice A Factory contract to create OI Cash Token contracts
 * @dev Should not be used directly. Only intended to be used by Universe contracts
 */
contract OICashFactory is CloneFactory {
    function createOICash(IAugur _augur) public returns (IOICash) {
        IUniverse _universe = IUniverse(msg.sender);

        address newContractAddress = createNewContract();
        IOICash _openInterestCash = IOICash(newContractAddress);
        _openInterestCash.initialize(_augur, _universe);

        return _openInterestCash;
    }
}