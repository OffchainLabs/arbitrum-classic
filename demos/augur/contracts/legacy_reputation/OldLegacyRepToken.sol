pragma solidity 0.5.15;


import 'ROOT/reporting/IReputationToken.sol';
import 'ROOT/legacy_reputation/DelegationTarget.sol';
import 'ROOT/libraries/ITyped.sol';
import 'ROOT/libraries/Initializable.sol';
import 'ROOT/libraries/token/VariableSupplyToken.sol';
import 'ROOT/libraries/token/IERC20.sol';
import 'ROOT/reporting/IUniverse.sol';
import 'ROOT/reporting/IMarket.sol';
import 'ROOT/reporting/Reporting.sol';
import 'ROOT/reporting/IDisputeCrowdsourcer.sol';
import 'ROOT/libraries/math/SafeMathUint256.sol';


contract OldLegacyReputationToken is DelegationTarget, ITyped, Initializable, VariableSupplyToken, IReputationToken {
    using SafeMathUint256 for uint256;

    string constant public name = "Reputation";
    string constant public symbol = "REP";
    uint8 constant public decimals = 18;
    IUniverse private universe;
    uint256 private totalMigrated;
    mapping(address => uint256) migratedToSibling;
    uint256 private parentTotalTheoreticalSupply;
    uint256 private totalTheoreticalSupply;

    // Auto migration related state
    bool private isMigratingFromLegacy;
    uint256 private targetSupply;

    /**
     * @dev modifier to allow actions only when the contract IS paused
     */
    modifier whenMigratingFromLegacy() {
        require(isMigratingFromLegacy);
        _;
    }

    /**
     * @dev modifier to allow actions only when the contract IS paused
     */
    modifier whenNotMigratingFromLegacy() {
        require(!isMigratingFromLegacy);
        _;
    }

    function initialize(IUniverse _universe) public beforeInitialized {
        endInitialization();
        require(_universe != IUniverse(0));
        universe = _universe;
        updateParentTotalTheoreticalSupply();
        IERC20 _legacyRepToken = getLegacyRepToken();
        // Initialize migration related state. If this is Genesis universe REP the balances from the Legacy contract must be migrated before we enable usage
        isMigratingFromLegacy = _universe.getParentUniverse() == IUniverse(0);
        targetSupply = _legacyRepToken.totalSupply();
    }

    function migrateOutByPayout(uint256[] memory _payoutNumerators, uint256 _attotokens) public whenNotMigratingFromLegacy returns (bool) {
        require(_attotokens > 0);
        IUniverse _destinationUniverse = universe.createChildUniverse(_payoutNumerators);
        IReputationToken _destination = _destinationUniverse.getReputationToken();
        burn(msg.sender, _attotokens);
        _destination.migrateIn(msg.sender, _attotokens);
        return true;
    }

    function migrateOut(IReputationToken _destination, uint256 _attotokens) public whenNotMigratingFromLegacy returns (bool) {
        require(_attotokens > 0);
        assertReputationTokenIsLegitChild(_destination);
        burn(msg.sender, _attotokens);
        _destination.migrateIn(msg.sender, _attotokens);
        return true;
    }

    function migrateIn(address _reporter, uint256 _attotokens) public whenNotMigratingFromLegacy returns (bool) {
        IUniverse _parentUniverse = universe.getParentUniverse();
        require(msg.sender == address(_parentUniverse.getReputationToken()));
        mint(_reporter, _attotokens);
        totalMigrated += _attotokens;
        // Award a bonus if migration is done before the fork period is over, even if it has finalized
        if (getTimestamp() < _parentUniverse.getForkEndTime()) {
            uint256 _bonus = _attotokens.div(20);
            mint(_reporter, _bonus);
            totalTheoreticalSupply += _bonus;
        }
        // Update the fork tenative winner and finalize if we can
        if (!_parentUniverse.getForkingMarket().isFinalized()) {
            _parentUniverse.updateTentativeWinningChildUniverse(universe.getParentPayoutDistributionHash());
        }
        return true;
    }

    function mintForReportingParticipant(uint256 _amountMigrated) public whenNotMigratingFromLegacy returns (bool) {
        IUniverse _parentUniverse = universe.getParentUniverse();
        IReportingParticipant _reportingParticipant = IReportingParticipant(msg.sender);
        require(_parentUniverse.isContainerForReportingParticipant(_reportingParticipant));
        uint256 _bonus = _amountMigrated.div(2);
        mint(address(_reportingParticipant), _bonus);
        totalTheoreticalSupply += _bonus;
        return true;
    }

    function transfer(address _to, uint _value) public whenNotMigratingFromLegacy returns (bool) {
        return super.transfer(_to, _value);
    }

    function transferFrom(address _from, address _to, uint _value) public whenNotMigratingFromLegacy returns (bool) {
        require(_value <= allowances[_from][msg.sender], "Not enough funds allowed");
        allowances[_from][msg.sender] = allowances[_from][msg.sender].sub(_value);
        balances[_from] = balances[_from].sub(_value);
        balances[_to] = balances[_to].add(_value);
        emit Transfer(_from, _to, _value);
        return true;
    }

    function trustedUniverseTransfer(address _source, address _destination, uint256 _attotokens) public whenNotMigratingFromLegacy returns (bool) {
        require(IUniverse(msg.sender) == universe);
        _transfer(_source, _destination, _attotokens);
        return true;
    }

    function trustedMarketTransfer(address _source, address _destination, uint256 _attotokens) public whenNotMigratingFromLegacy returns (bool) {
        require(universe.isContainerForMarket(IMarket(msg.sender)));
        _transfer(_source, _destination, _attotokens);
        return true;
    }

    function trustedReportingParticipantTransfer(address _source, address _destination, uint256 _attotokens) public whenNotMigratingFromLegacy returns (bool) {
        require(universe.isContainerForReportingParticipant(IReportingParticipant(msg.sender)));
        _transfer(_source, _destination, _attotokens);
        return true;
    }

    function trustedDisputeWindowTransfer(address _source, address _destination, uint256 _attotokens) public whenNotMigratingFromLegacy returns (bool) {
        require(universe.isContainerForDisputeWindow(IDisputeWindow(msg.sender)));
        _transfer(_source, _destination, _attotokens);
        return true;
    }

    function assertReputationTokenIsLegitChild(IReputationToken _shadyReputationToken) private view {
        IUniverse _shadyUniverse = _shadyReputationToken.getUniverse();
        require(universe.isParentOf(_shadyUniverse));
        IUniverse _legitUniverse = _shadyUniverse;
        require(_legitUniverse.getReputationToken() == _shadyReputationToken);
    }

    function getTypeName() public view returns (bytes32) {
        return "ReputationToken";
    }

    function getUniverse() public view returns (IUniverse) {
        return universe;
    }

    function getTotalMigrated() public view returns (uint256) {
        return totalMigrated;
    }

    function getLegacyRepToken() public pure returns (IERC20) {
        return IERC20(address(0xE94327D07Fc17907b4DB788E5aDf2ed424adDff6));
    }

    function updateSiblingMigrationTotal(IReputationToken _token) public whenNotMigratingFromLegacy returns (bool) {
        require(_token != this);
        IUniverse _shadyUniverse = _token.getUniverse();
        require(_token == universe.getParentUniverse().getChildUniverse(_shadyUniverse.getParentPayoutDistributionHash()).getReputationToken());
        totalTheoreticalSupply += migratedToSibling[address(_token)];
        migratedToSibling[address(_token)] = _token.getTotalMigrated();
        totalTheoreticalSupply -= migratedToSibling[address(_token)];
        return true;
    }

    function updateParentTotalTheoreticalSupply() public whenNotMigratingFromLegacy returns (bool) {
        IUniverse _parentUniverse = universe.getParentUniverse();
        totalTheoreticalSupply -= parentTotalTheoreticalSupply;
        if (_parentUniverse == IUniverse(0)) {
            parentTotalTheoreticalSupply = Reporting.getInitialREPSupply();
        } else {
            parentTotalTheoreticalSupply = _parentUniverse.getReputationToken().getTotalTheoreticalSupply();
        }
        totalTheoreticalSupply += parentTotalTheoreticalSupply;
        return true;
    }

    function getTotalTheoreticalSupply() public view returns (uint256) {
        return totalTheoreticalSupply;
    }

        /**
     * @dev Copies the balance of a batch of addresses from the legacy contract
     * @param _holders Array of addresses to migrate balance
     * @return True if operation was completed
     */
    function migrateBalancesFromLegacyRep(address[] memory _holders) public whenMigratingFromLegacy returns (bool) {
        IERC20 _legacyRepToken = getLegacyRepToken();
        for (uint256 i = 0; i < _holders.length; i++) {
            migrateBalanceFromLegacyRep(_holders[i], _legacyRepToken);
        }
        return true;
    }

    /**
     * @dev Copies the balance of a single addresses from the legacy contract
     * @param _holder Address to migrate balance
     * @return True if balance was copied, false if was already copied or address had no balance
     */
    function migrateBalanceFromLegacyRep(address _holder, IERC20 _legacyRepToken) private whenMigratingFromLegacy returns (bool) {
        if (balances[_holder] > 0) {
            return false; // Already copied, move on
        }

        uint256 amount = _legacyRepToken.balanceOf(_holder);
        if (amount == 0) {
            return false; // Has no balance in legacy contract, move on
        }

        mint(_holder, amount);

        if (targetSupply == totalSupply) {
            isMigratingFromLegacy = false;
        }
        return true;
    }

    /**
     * @dev Copies the allowances of a batch of addresses from the legacy contract. This is an optional step which may only be done before the migration is complete but is not required to complete it.
     * @param _owners Array of owner addresses to migrate allowances
     * @param _spenders Array of spender addresses to migrate allowances
     * @return True if operation was completed
     */
    function migrateAllowancesFromLegacyRep(address[] memory _owners, address[] memory _spenders) public whenMigratingFromLegacy returns (bool) {
        IERC20 _legacyRepToken = getLegacyRepToken();
        for (uint256 i = 0; i < _owners.length; i++) {
            address _owner = _owners[i];
            address _spender = _spenders[i];
            uint256 _allowance = _legacyRepToken.allowance(_owner, _spender);
            _approve(_owner, _spender, _allowance);
        }
        return true;
    }

    function getIsMigratingFromLegacy() public view returns (bool) {
        return isMigratingFromLegacy;
    }

    function getTargetSupply() public view returns (uint256) {
        return targetSupply;
    }

    function getTimestamp() public view returns (uint256) {
        return block.timestamp;
    }
}
