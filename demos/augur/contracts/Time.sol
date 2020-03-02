pragma solidity 0.5.15;

import './ITime.sol';
import "./ArbSys.sol";


/**
 * @title Time
 * @notice Contract to encapsulate time on chain for easy mocking out and testing
 */
contract Time is ITime {

    function withinTime(uint256 time) external view returns(bool) {
    	return time >= ArbSys(100).timeUpperBound();
  	}

  	function pastTime(uint256 time) external view returns(bool) {
    	return time < block.number;
  	}

  	function timeLowerBound() external view returns(uint256) {
  		return block.number;
  	}

  	function timeUpperBound() external view returns(uint256) {
  		return ArbSys(100).timeUpperBound();
  	}

  	function getMessageInputTime() external view returns(uint256){
  		return ArbSys(100).currentMessageTime();
  	}

    function getTypeName() public view returns (bytes32) {
        return "Time";
    }
}
