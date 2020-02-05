pragma solidity 0.5.15;

import './ITime.sol';
import './IAugur.sol';


contract TimeControlled is ITime {
    uint256 public timestamp = 1;

    IAugur public augur;

    constructor() public {
        timestamp = block.timestamp;
    }

    function initialize(IAugur _augur) public returns (bool) {
        augur = _augur;
        return true;
    }

    function getTimestamp() external view returns (uint256) {
        return timestamp;
    }

    function incrementTimestamp(uint256 _amount) external returns (bool) {
        timestamp += _amount;
        augur.logTimestampSet(timestamp);
        return true;
    }

    function setTimestamp(uint256 _timestamp) external returns (bool) {
        timestamp = _timestamp;
        augur.logTimestampSet(timestamp);
        return true;
    }

    function getTypeName() public view returns (bytes32) {
        return "TimeControlled";
    }
}
