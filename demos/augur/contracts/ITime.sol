pragma solidity 0.5.15;

import './libraries/ITyped.sol';


contract ITime is ITyped {
    function getTimestamp() external view returns (uint256);
}
