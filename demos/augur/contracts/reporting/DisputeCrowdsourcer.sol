pragma solidity 0.5.15;

import './IDisputeCrowdsourcer.sol';
import '../libraries/token/VariableSupplyToken.sol';
import './BaseReportingParticipant.sol';
import '../libraries/Initializable.sol';
import './IUniverse.sol';
import '../IAugur.sol';


/**
 * @title Dispute Crowdsourcer
 * @notice A bond used during the disputing phase of a Market lifecycle.
 */
contract DisputeCrowdsourcer is VariableSupplyToken, BaseReportingParticipant, IDisputeCrowdsourcer, Initializable {
    IUniverse public universe;
    uint256 public crowdsourcerGeneration;

    string constant public name = "Dispute Crowdsourcer Token";
    string constant public symbol = "DISP";

    function initialize(IAugur _augur, IMarket _market, uint256 _size, bytes32 _payoutDistributionHash, uint256[] memory _payoutNumerators, uint256 _crowdsourcerGeneration) public beforeInitialized {
        endInitialization();
        augur = _augur;
        market = _market;
        universe = market.getUniverse();
        reputationToken = market.getReputationToken();
        size = _size;
        payoutNumerators = _payoutNumerators;
        payoutDistributionHash = _payoutDistributionHash;
        crowdsourcerGeneration = _crowdsourcerGeneration;
    }

    /**
     * @notice Redeems any token balance of this bond for the provided redeemer in exchange for owed REP
     * @param _redeemer The account to redeem for
     * @return bool True
     */
    function redeem(address _redeemer) public returns (bool) {
        bool _isDisavowed = isDisavowed();
        if (!_isDisavowed && !market.isFinalized()) {
            market.finalize();
        }
        uint256 _reputationSupply = reputationToken.balanceOf(address(this));
        uint256 _supply = totalSupply;
        uint256 _amount = balances[_redeemer];
        uint256 _reputationShare = _reputationSupply.mul(_amount).div(_supply);
        burn(_redeemer, _amount);
        require(reputationToken.transfer(_redeemer, _reputationShare));
        augur.logDisputeCrowdsourcerRedeemed(universe, _redeemer, address(market), _amount, _reputationShare, payoutNumerators);
        return true;
    }

    function contribute(address _participant, uint256 _amount, bool _overload) public returns (uint256) {
        require(IMarket(msg.sender) == market);
        if (_overload) {
            universe.updateForkValues();
            _amount = _amount.min(universe.getDisputeThresholdForDisputePacing().sub(totalSupply));
        } else {
            _amount = _amount.min(size.sub(totalSupply));
        }
        if (_amount == 0) {
            return 0;
        }
        reputationToken.trustedReportingParticipantTransfer(_participant, address(this), _amount);
        mint(_participant, _amount);
        assert(reputationToken.balanceOf(address(this)) >= totalSupply);
        return _amount;
    }

    /**
     * @notice Used in the event of the market forking. First forks this bond into the appropriate child universe and then redeems there for the msg sender.
     * @return bool True
     */
    function forkAndRedeem() public returns (bool) {
        fork();
        redeem(msg.sender);
        return true;
    }

    /**
     * @return The amount of REP remaining needed to fill this bond.
     */
    function getRemainingToFill() public view returns (uint256) {
        return size.sub(totalSupply);
    }

    function setSize(uint256 _size) public {
        require(IMarket(msg.sender) == market);
        size = _size;
    }

    /**
     * @return The amount of REP currently staked in this bond
     */
    function getStake() public view returns (uint256) {
        return totalSupply;
    }

    function onTokenTransfer(address _from, address _to, uint256 _value) internal {
        augur.logDisputeCrowdsourcerTokensTransferred(universe, _from, _to, _value, balances[_from], balances[_to]);
    }

    function onMint(address _target, uint256 _amount) internal {
        augur.logDisputeCrowdsourcerTokensMinted(universe, _target, _amount, totalSupply, balances[_target]);
    }

    function onBurn(address _target, uint256 _amount) internal {
        augur.logDisputeCrowdsourcerTokensBurned(universe, _target, _amount, totalSupply, balances[_target]);
    }

    /**
     * @return The REP token associated with this bond.
     */
    function getReputationToken() public view returns (IReputationToken) {
        return reputationToken;
    }

    function correctSize() public returns (bool) {
        require(IMarket(msg.sender) == market);
        size = totalSupply;
        return true;
    }

    function getCrowdsourcerGeneration() public view returns (uint256) {
        return crowdsourcerGeneration;
    }
}
