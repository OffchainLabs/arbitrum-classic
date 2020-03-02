pragma solidity 0.5.15;

import './libraries/ITyped.sol';


contract ITime is ITyped {
    function withinTime() external view returns (bool);
    function pastTime() external view returns (bool);
    function timeLowerBound() external view returns(uint256);
    function timeUpperBound() external view returns(uint256);
    function getMessageInputTime() external view returns(uint256);
}
