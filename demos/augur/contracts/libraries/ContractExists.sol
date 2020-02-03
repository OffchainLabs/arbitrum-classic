pragma solidity 0.5.15;


// Utility to check if the address actually contains a contract based on size.
// Note: This will fail if called from the contract's constructor
library ContractExists {
    function exists(address _address) internal view returns (bool) {
        uint256 size;
        assembly { size := extcodesize(_address) }
        return size > 0;
    }
}
