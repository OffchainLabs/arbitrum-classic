pragma solidity 0.5.15;

import 'ROOT/libraries/CloneFactory.sol';
import 'ROOT/reporting/IDisputeCrowdsourcer.sol';
import 'ROOT/reporting/IMarket.sol';
import 'ROOT/IAugur.sol';
import 'ROOT/factories/IDisputeCrowdsourcerFactory.sol';


/**
 * @title Dispute Crowdsourcer Factory
 * @notice A Factory contract to create Dispute Crowdsourcer delegate contracts
 * @dev Cannot be used directly. Only called by Market contracts
 */
contract DisputeCrowdsourcerFactory is CloneFactory, IDisputeCrowdsourcerFactory {
    function createDisputeCrowdsourcer(IAugur _augur, uint256 _size, bytes32 _payoutDistributionHash, uint256[] memory _payoutNumerators, uint256 _crowdsourcerGeneration) public returns (IDisputeCrowdsourcer) {
        IMarket _market = IMarket(msg.sender);
        require(_augur.isKnownMarket(_market), "DisputeCrowdsourcerFactory: Market specified is unrecognized by Augur");
        IDisputeCrowdsourcer _disputeCrowdsourcer = IDisputeCrowdsourcer(createClone(_augur.lookup("DisputeCrowdsourcer")));
        _disputeCrowdsourcer.initialize(_augur, _market, _size, _payoutDistributionHash, _payoutNumerators, _crowdsourcerGeneration);
        return _disputeCrowdsourcer;
    }
}
