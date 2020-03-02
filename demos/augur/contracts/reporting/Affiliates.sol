pragma solidity 0.5.15;

import './IAffiliateValidator.sol';
import './AffiliateValidator.sol';
import "../ArbSys.sol";

/**
 * @title Affiliates
 * @notice A contract used to record an account's referrer and their browser fingerprint for use in the affiliate system
 */
contract Affiliates {
    // Maps an account to their fingerprint. Used to naievly filter out attempts at self reference
    mapping (address => bytes32) public fingerprints;

    // Maps an account to the referral account. Used to apply affiliate fees on settlement.
    mapping (address => address) public referrals;

    // Mapping of valid Affiliate Validators
    mapping (address => bool) public affiliateValidators;

    address public validatorTemplate;

    function initializeValidatorFactory(address template) public {
        require(validatorTemplate == address(0));
        require(template != address(0));
        validatorTemplate = template;
    }

    /**
     * @notice Create a new Affiliate Validator contract to be used in markets
     * @return AffiliateValidator
     */
    function createAffiliateValidator() public returns (AffiliateValidator) {
        require(validatorTemplate != address(0));

        address payable validatorAddress = address(uint160(ArbSys(100).cloneContract(validatorTemplate)));
        AffiliateValidator _affiliateValidator = AffiliateValidator(validatorAddress);

        _affiliateValidator.transferOwnership(msg.sender);
        affiliateValidators[address(_affiliateValidator)] = true;
        
        return _affiliateValidator;
    }

    /**
     * @notice Sets the browser fingerprint for an account
     * @param _fingerprint The account browser fingerprint
     */
    function setFingerprint(bytes32 _fingerprint) external {
        fingerprints[msg.sender] = _fingerprint;
    }

    /**
     * @notice Set the referring account for the sender.
     * @param _referrer The referrer who should recieve affiliate fees when possible for this account
     */
    function setReferrer(address _referrer) external {
        require(msg.sender != _referrer);

        if (referrals[msg.sender] != address(0)) {
            return;
        }

        referrals[msg.sender] = _referrer;
    }

    /**
     * @notice Get the fingerprint for an account
     * @param _account The account whose fingerprint to look up
     * @return bytes32
     */
    function getAccountFingerprint(address _account) external view returns (bytes32) {
        return fingerprints[_account];
    }

    /**
     * @notice Get the referrer for an account
     * @param _account The account whose referrer to look up
     * @return address
     */
    function getReferrer(address _account) external view returns (address) {
        return referrals[_account];
    }

    function getAndValidateReferrer(address _account, IAffiliateValidator _affiliateValidator) external view returns (address) {
        address _referrer = referrals[_account];
        if (_referrer == address(0) || _account == _referrer) {
            return address(0);
        }
        if (_affiliateValidator == IAffiliateValidator(0)) {
            return _referrer;
        }
        bool _success = _affiliateValidator.validateReference(_account, _referrer);
        if (!_success) {
            return address(0);
        }
        return _referrer;
    }
}