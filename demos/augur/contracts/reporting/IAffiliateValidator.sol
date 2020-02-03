pragma solidity 0.5.15;


contract IAffiliateValidator {
    function validateReference(address _account, address _referrer) external view returns (bool);
}