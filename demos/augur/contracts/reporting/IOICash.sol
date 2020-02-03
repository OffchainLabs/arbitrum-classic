pragma solidity 0.5.15;

import 'ROOT/reporting/IUniverse.sol';
import 'ROOT/libraries/token/IERC20.sol';


contract IOICash is IERC20 {
    function initialize(IAugur _augur, IUniverse _universe) external;
}
