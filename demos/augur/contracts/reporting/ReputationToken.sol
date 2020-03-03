pragma solidity 0.5.15;

import './IV2ReputationToken.sol';
import '../libraries/ITyped.sol';
import '../libraries/token/VariableSupplyToken.sol';
import '../libraries/token/IERC20.sol';
import './IUniverse.sol';
import './IMarket.sol';
import './Reporting.sol';
import './IDisputeWindow.sol';
import './IDisputeCrowdsourcer.sol';
import '../libraries/math/SafeMathUint256.sol';
import '../utility/IRepSymbol.sol';
import '../libraries/Initializable.sol';


/**
 * @title Reputation Token
 * @notice The Reputation Token for a particular universe
 */
contract ReputationToken is Initializable, VariableSupplyToken, IV2ReputationToken {
    using SafeMathUint256 for uint256;

    string constant public name = "Reputation";
    IUniverse internal universe;
    IUniverse public parentUniverse;
    uint256 internal totalMigrated;
    IERC20 public legacyRepToken;
    IAugur public augur;
    address public warpSync;

    function initializeRepToken(IAugur _augur, IUniverse _universe, IUniverse _parentUniverse) public beforeInitialized {
        endInitialization();
        augur = _augur;
        universe = _universe;
        parentUniverse = _parentUniverse;
        warpSync = _augur.lookup("WarpSync");
        legacyRepToken = IERC20(_augur.lookup("LegacyReputationToken"));
        require(warpSync != address(0));
        require(legacyRepToken != IERC20(0));
    }


    function symbol() public view returns (string memory) {
        return IRepSymbol(augur.lookup("RepSymbol")).getRepSymbol(address(augur), address(universe));
    }

    /**
     * @notice Migrate to a Child Universe by indicating the Market payout associated with it
     * @param _payoutNumerators The array of payouts for the market associated with the desired universe
     * @param _attotokens The amount of tokens to migrate
     * @return Bool True
     */
    function migrateOutByPayout(uint256[] memory _payoutNumerators, uint256 _attotokens) public returns (bool) {
        require(_attotokens > 0);
        IUniverse _destinationUniverse = universe.createChildUniverse(_payoutNumerators);
        IReputationToken _destination = _destinationUniverse.getReputationToken();
        burn(msg.sender, _attotokens);
        _destination.migrateIn(msg.sender, _attotokens);
        return true;
    }

    function migrateIn(address _reporter, uint256 _attotokens) public returns (bool) {
        IUniverse _parentUniverse = parentUniverse;
        require(ReputationToken(msg.sender) == _parentUniverse.getReputationToken());
        require(augur.getTimestamp() < _parentUniverse.getForkEndTime());
        mint(_reporter, _attotokens);
        totalMigrated += _attotokens;
        // Update the fork tentative winner and finalize if we can
        if (!_parentUniverse.getForkingMarket().isFinalized()) {
            _parentUniverse.updateTentativeWinningChildUniverse(universe.getParentPayoutDistributionHash());
        }
        return true;
    }

    function mintForReportingParticipant(uint256 _amountMigrated) public returns (bool) {
        IReportingParticipant _reportingParticipant = IReportingParticipant(msg.sender);
        require(parentUniverse.isContainerForReportingParticipant(_reportingParticipant));
        // simulate a 40% ROI which would have occured during a normal dispute had this participant's outcome won the dispute
        uint256 _bonus = _amountMigrated.mul(2) / 5;
        mint(address(_reportingParticipant), _bonus);
        return true;
    }

    function mintForWarpSync(uint256 _amountToMint, address _target) public returns (bool) {
        require(warpSync == msg.sender);
        mint(_target, _amountToMint);
        universe.updateForkValues();
        return true;
    }

    function burnForMarket(uint256 _amountToBurn) public returns (bool) {
        require(universe.isContainerForMarket(IMarket(msg.sender)));
        burn(msg.sender, _amountToBurn);
        return true;
    }

    function trustedUniverseTransfer(address _source, address _destination, uint256 _attotokens) public returns (bool) {
        require(IUniverse(msg.sender) == universe);
        _transfer(_source, _destination, _attotokens);
        return true;
    }

    function trustedMarketTransfer(address _source, address _destination, uint256 _attotokens) public returns (bool) {
        require(universe.isContainerForMarket(IMarket(msg.sender)));
        _transfer(_source, _destination, _attotokens);
        return true;
    }

    function trustedReportingParticipantTransfer(address _source, address _destination, uint256 _attotokens) public returns (bool) {
        require(universe.isContainerForReportingParticipant(IReportingParticipant(msg.sender)));
        _transfer(_source, _destination, _attotokens);
        return true;
    }

    function trustedDisputeWindowTransfer(address _source, address _destination, uint256 _attotokens) public returns (bool) {
        require(universe.isContainerForDisputeWindow(IDisputeWindow(msg.sender)));
        _transfer(_source, _destination, _attotokens);
        return true;
    }

    function trustedREPExchangeTransfer(address _source, address _destination, uint256 _attotokens) public returns (bool) {
        require(address(universe.repExchange()) == msg.sender);
        _transfer(_source, _destination, _attotokens);
        return true;
    }

    function assertReputationTokenIsLegitChild(IReputationToken _shadyReputationToken) private view {
        IUniverse _universe = _shadyReputationToken.getUniverse();
        require(universe.isParentOf(_universe));
        require(_universe.getReputationToken() == _shadyReputationToken);
    }

    /**
     * @return The universe associated with this Reputation Token
     */
    function getUniverse() public view returns (IUniverse) {
        return universe;
    }

    /**
     * @return The total amount of parent REP migrated into this version of REP
     */
    function getTotalMigrated() public view returns (uint256) {
        return totalMigrated;
    }

    /**
     * @return The V1 Rep token
     */
    function getLegacyRepToken() public view returns (IERC20) {
        return legacyRepToken;
    }

    /**
     * @return The maximum possible total supply for this version of REP.
     */
    function getTotalTheoreticalSupply() public view returns (uint256) {
        uint256 _totalSupply = totalSupply;
        if (parentUniverse == IUniverse(0)) {
            return _totalSupply.add(legacyRepToken.totalSupply()).sub(legacyRepToken.balanceOf(address(1))).sub(legacyRepToken.balanceOf(address(0)));
        } else if (augur.getTimestamp() >= parentUniverse.getForkEndTime()) {
            return _totalSupply;
        } else {
            return _totalSupply + parentUniverse.getReputationToken().getTotalTheoreticalSupply();
        }
    }

    function onTokenTransfer(address _from, address _to, uint256 _value) internal {
        augur.logReputationTokensTransferred(universe, _from, _to, _value, balances[_from], balances[_to]);
    }

    function onMint(address _target, uint256 _amount) internal {
        augur.logReputationTokensMinted(universe, _target, _amount, totalSupply, balances[_target]);
    }

    function onBurn(address _target, uint256 _amount) internal {
        augur.logReputationTokensBurned(universe, _target, _amount, totalSupply, balances[_target]);
    }

    /**
     * @notice Migrate V1 REP to V2
     * @dev This can only be done for the Genesis Universe in V2. If a fork occurs and the window ends V1 REP is stuck in V1 forever
     * @return Bool True
     */
    function migrateFromLegacyReputationToken() public returns (bool) {
        require(parentUniverse == IUniverse(0));
        uint256 _legacyBalance = legacyRepToken.balanceOf(msg.sender);
        require(legacyRepToken.transferFrom(msg.sender, address(1), _legacyBalance));
        mint(msg.sender, _legacyBalance);
        return true;
    }
}
