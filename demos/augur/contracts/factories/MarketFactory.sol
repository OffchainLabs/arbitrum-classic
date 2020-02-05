pragma solidity 0.5.15;


import '../libraries/CloneFactory.sol';
import '../reporting/IMarket.sol';
import '../reporting/IReputationToken.sol';
import '../reporting/IAffiliateValidator.sol';
import '../ICash.sol';
import '../factories/IMarketFactory.sol';
import '../libraries/math/SafeMathUint256.sol';
import '../IAugur.sol';


/**
 * @title Market Factory
 * @notice A Factory contract to create Market delegate contracts
 * @dev This factory is a trusted Augur contract which means it has REP transfer privileges. The additional validation logic that occurs here and not in the Market contract itself is to reduce the Market contract's size as each require error message adds an additional ~100 bytes
 */
contract MarketFactory is CloneFactory, IMarketFactory {
    using SafeMathUint256 for uint256;

    uint256 private constant MAX_FEE_PER_CASH_IN_ATTOCASH = 15 * 10**16; // 15%
    address private constant NULL_ADDRESS = address(0);
    uint256 private constant MIN_OUTCOMES = 2; // Does not Include Invalid
    uint256 private constant MAX_OUTCOMES = 7; // Does not Include Invalid

    function createMarket(IAugur _augur, uint256 _endTime, uint256 _feePerCashInAttoCash, IAffiliateValidator _affiliateValidator, uint256 _affiliateFeeDivisor, address _designatedReporterAddress, address _sender, uint256 _numOutcomes, uint256 _numTicks) public returns (IMarket _market) {
        _market = IMarket(createClone(_augur.lookup("Market")));
        IUniverse _universe = IUniverse(msg.sender);
        require(_augur.isKnownUniverse(_universe), "MarketFactory: Universe specified is unrecognized by Augur");
        IReputationToken _reputationToken = _universe.getReputationToken();
        require(_reputationToken.transfer(address(_market), _reputationToken.balanceOf(address(this))));
        if (_sender != _augur.lookup("WarpSync")) {
            require(_augur.trustedCashTransfer(_sender, address(_market), _universe.getOrCacheValidityBond()));
        }

        // Market param validation
        require(_numTicks.isMultipleOf(2), "MarketFactory.createMarket: numTicks must be multiple of 2");
        require(MIN_OUTCOMES <= _numOutcomes && _numOutcomes <= MAX_OUTCOMES, "MarketFactory.createMarket: numOutcomes out of range");
        require(_designatedReporterAddress != NULL_ADDRESS, "MarketFactory.createMarket: designated rpeorter address is 0x0");
        require(_numTicks >= _numOutcomes, "MarketFactory.createMarket: numTicks lower than numOutcomes");
        require(_feePerCashInAttoCash <= MAX_FEE_PER_CASH_IN_ATTOCASH, "MarketFactory.createMarket: market creator fee too high");
        require(_sender != NULL_ADDRESS, "MarketFactory.createMarket: market creator address is 0x0");
        require(_augur.getTimestamp() < _endTime, "MarketFactory.createMarket: endTime before current time");
        require(_endTime < _augur.getMaximumMarketEndDate(), "MarketFactory.createMarket: endTime too far ahead");
        require(!_universe.isForking(), "MarketFactory.createMarket: Universe is forking");

        _market.initialize(_augur, _universe, _endTime, _feePerCashInAttoCash, _affiliateValidator, _affiliateFeeDivisor, _designatedReporterAddress, _sender, _numOutcomes, _numTicks);
        return _market;
    }
}
