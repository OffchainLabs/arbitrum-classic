pragma solidity 0.5.15;

import './IAffiliateValidator.sol';


contract IAffiliates {
    function setFingerprint(bytes32 _fingerprint) external;
    function setReferrer(address _referrer) external;
    function getAccountFingerprint(address _account) external returns (bytes32);
    function getReferrer(address _account) external returns (address);
    function getAndValidateReferrer(address _account, IAffiliateValidator affiliateValidator) external returns (address);
    function affiliateValidators(address _affiliateValidator) external returns (bool);
}