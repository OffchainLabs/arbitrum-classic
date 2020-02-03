pragma solidity 0.5.15;

import 'ROOT/ITime.sol';


/**
 * @title Time
 * @notice Contract to encapsulate time on chain for easy mocking out and testing
 */
contract Time is ITime {
    function getTimestamp() external view returns (uint256) {
        return block.timestamp;
    }

    function getTypeName() public view returns (bytes32) {
        return "Time";
    }
}
