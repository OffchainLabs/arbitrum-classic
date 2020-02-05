pragma solidity 0.5.15;

import './IUniverse.sol';
import '../libraries/token/IERC20.sol';


contract IOICash is IERC20 {
    function initialize(IAugur _augur, IUniverse _universe) external;
}
