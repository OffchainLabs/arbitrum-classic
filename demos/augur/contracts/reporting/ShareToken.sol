pragma solidity 0.5.15;

import './IShareToken.sol';
import '../libraries/token/ERC1155.sol';
import '../libraries/ReentrancyGuard.sol';
import '../libraries/ITyped.sol';
import '../libraries/Initializable.sol';
import './IMarket.sol';
import '../IAugur.sol';
import '../CashSender.sol';
import '../libraries/TokenId.sol';


/**
 * @title Share Token
 * @notice ERC1155 contract to hold all Augur share token balances
 */
contract ShareToken is ITyped, Initializable, ERC1155, IShareToken, ReentrancyGuard, CashSender {

    string constant public name = "Shares";
    string constant public symbol = "SHARE";

    struct MarketData {
        uint256 numOutcomes;
        uint256 numTicks;
    }

    mapping(address => MarketData) markets;

    IAugur public augur;
    ICash public cash;

    function initialize(IAugur _augur) external beforeInitialized {
        endInitialization();
        augur = _augur;
        cash = ICash(_augur.lookup("Cash"));

        initializeCashSender(_augur.lookup("DaiVat"), address(cash));
        // require(cash != ICash(0));
    }

    /**
        @dev Transfers `value` amount of an `id` from the `from` address to the `to` address specified.
        Caller must be approved to manage the tokens being transferred out of the `from` account.
        Regardless of if the desintation is a contract or not this will not call `onERC1155Received` on `to`
        @param _from Source address
        @param _to Target address
        @param _id ID of the token type
        @param _value Transfer amount
    */
    function unsafeTransferFrom(address _from, address _to, uint256 _id, uint256 _value) public {
        _transferFrom(_from, _to, _id, _value, bytes(""), false);
    }

    /**
        @dev Transfers `values` amount(s) of `ids` from the `from` address to the
        `to` address specified. Caller must be approved to manage the tokens being
        transferred out of the `from` account. Regardless of if the desintation is
        a contract or not this will not call `onERC1155Received` on `to`
        @param _from Source address
        @param _to Target address
        @param _ids IDs of each token type
        @param _values Transfer amounts per token type
    */
    function unsafeBatchTransferFrom(address _from, address _to, uint256[] memory _ids, uint256[] memory _values) public {
        _batchTransferFrom(_from, _to, _ids, _values, bytes(""), false);
    }

    function initializeMarket(IMarket _market, uint256 _numOutcomes, uint256 _numTicks) public {
        // require (augur.isKnownUniverse(IUniverse(msg.sender)));
        markets[address(_market)].numOutcomes = _numOutcomes;
        markets[address(_market)].numTicks = _numTicks;
    }

    /**
     * @notice Buy some amount of complete sets for a market
     * @param _market The market to purchase complete sets in
     * @param _amount The number of complete sets to purchase
     * @return Bool True
     */
    function publicBuyCompleteSets(IMarket _market, uint256 _amount) external returns (bool) {
        buyCompleteSetsInternal(_market, msg.sender, _amount);
        augur.logCompleteSetsPurchased(_market.getUniverse(), _market, msg.sender, _amount);
    }

    /**
     * @notice Buy some amount of complete sets for a market
     * @param _market The market to purchase complete sets in
     * @param _account The account receiving the complete sets
     * @param _amount The number of complete sets to purchase
     * @return Bool True
     */
    function buyCompleteSets(IMarket _market, address _account, uint256 _amount) external returns (bool) {
        buyCompleteSetsInternal(_market, _account, _amount);
    }

    function buyCompleteSetsInternal(IMarket _market, address _account, uint256 _amount) internal returns (bool) {
        uint256 _numOutcomes = markets[address(_market)].numOutcomes;
        uint256 _numTicks = markets[address(_market)].numTicks;

        // require(_numOutcomes != 0, "Invalid Market provided");

        IUniverse _universe = _market.getUniverse();

        uint256 _cost = _amount.mul(_numTicks);
        _universe.deposit(msg.sender, _cost, address(_market));

        uint256[] memory _tokenIds = new uint256[](_numOutcomes);
        uint256[] memory _values = new uint256[](_numOutcomes);

        for (uint256 _i = 0; _i < _numOutcomes; _i++) {
            _tokenIds[_i] = TokenId.getTokenId(_market, _i);
            _values[_i] = _amount;
        }

        _mintBatch(_account, _tokenIds, _values, bytes(""), false);

        if (!_market.isFinalized()) {
            _universe.incrementOpenInterest(_cost);
        }

        augur.logMarketOIChanged(_universe, _market);

        assertBalances(_market);
        return true;
    }

    /**
     * @notice Buy some amount of complete sets for a market and distribute the shares according to the positions of two accounts
     * @param _market The market to purchase complete sets in
     * @param _amount The number of complete sets to purchase
     * @param _longOutcome The outcome for the trade being fulfilled
     * @param _longRecipient The account which should recieve the _longOutcome shares
     * @param _shortRecipient The account which should recieve shares of every outcome other than _longOutcome
     * @return Bool True
     */
    function buyCompleteSetsForTrade(IMarket _market, uint256 _amount, uint256 _longOutcome, address _longRecipient, address _shortRecipient) external returns (bool) {
        uint256 _numOutcomes = markets[address(_market)].numOutcomes;

        // require(_numOutcomes != 0, "Invalid Market provided");
        // require(_longOutcome < _numOutcomes);

        IUniverse _universe = _market.getUniverse();

        {
            uint256 _numTicks = markets[address(_market)].numTicks;
            uint256 _cost = _amount.mul(_numTicks);
            _universe.deposit(msg.sender, _cost, address(_market));

            if (!_market.isFinalized()) {
                _universe.incrementOpenInterest(_cost);
            }
        }

        uint256[] memory _tokenIds = new uint256[](_numOutcomes - 1);
        uint256[] memory _values = new uint256[](_numOutcomes - 1);
        uint256 _outcome = 0;

        for (uint256 _i = 0; _i < _numOutcomes - 1; _i++) {
            if (_outcome == _longOutcome) {
                _outcome++;
            }
            _tokenIds[_i] = TokenId.getTokenId(_market, _outcome);
            _values[_i] = _amount;
            _outcome++;
        }

        _mintBatch(_shortRecipient, _tokenIds, _values, bytes(""), false);
        _mint(_longRecipient, TokenId.getTokenId(_market, _longOutcome), _amount, bytes(""), false);

        augur.logMarketOIChanged(_universe, _market);

        assertBalances(_market);
        return true;
    }

    /**
     * @notice Sell some amount of complete sets for a market
     * @param _market The market to sell complete sets in
     * @param _amount The number of complete sets to sell
     * @return (uint256 _creatorFee, uint256 _reportingFee) The fees taken for the market creator and reporting respectively
     */
    function publicSellCompleteSets(IMarket _market, uint256 _amount) external returns (uint256 _creatorFee, uint256 _reportingFee) {
        (uint256 _payout, uint256 _creatorFee, uint256 _reportingFee) = burnCompleteSets(_market, msg.sender, _amount, msg.sender, bytes32(0));

        cashTransfer(msg.sender, _payout);

        IUniverse _universe = _market.getUniverse();
        augur.logCompleteSetsSold(_universe, _market, msg.sender, _amount, _creatorFee.add(_reportingFee));

        assertBalances(_market);
        return (_creatorFee, _reportingFee);
    }

    /**
     * @notice Sell some amount of complete sets for a market
     * @param _market The market to sell complete sets in
     * @param _holder The holder of the complete sets
     * @param _recipient The recipient of funds from the sale
     * @param _amount The number of complete sets to sell
     * @param _fingerprint Fingerprint of the filler used to naively restrict affiliate fee dispursement
     * @return (uint256 _creatorFee, uint256 _reportingFee) The fees taken for the market creator and reporting respectively
     */
    function sellCompleteSets(IMarket _market, address _holder, address _recipient, uint256 _amount, bytes32 _fingerprint) external returns (uint256 _creatorFee, uint256 _reportingFee) {
        // require(_holder == msg.sender || isApprovedForAll(_holder, msg.sender) == true, "ERC1155: need operator approval to sell complete sets");
        
        (uint256 _payout, uint256 _creatorFee, uint256 _reportingFee) = burnCompleteSets(_market, _holder, _amount, _holder, _fingerprint);

        cashTransfer(_recipient, _payout);

        assertBalances(_market);
        return (_creatorFee, _reportingFee);
    }

    /**
     * @notice Sell some amount of complete sets for a market
     * @param _market The market to sell complete sets in
     * @param _amount The number of complete sets to sell
     * @param _shortParticipant The account which should provide the short party portion of shares
     * @param _longParticipant The account which should provide the long party portion of shares
     * @param _longRecipient The account which should receive the remaining payout for providing the matching shares to the short recipients shares
     * @param _shortRecipient The account which should recieve the (price * shares provided) payout for selling their side of the sale
     * @param _price The price of the trade being done. This determines how much each recipient recieves from the sale proceeds
     * @param _fingerprint Fingerprint of the filler used to naively restrict affiliate fee dispursement
     * @return (uint256 _creatorFee, uint256 _reportingFee) The fees taken for the market creator and reporting respectively
     */
    function sellCompleteSetsForTrade(IMarket _market, uint256 _outcome, uint256 _amount, address _shortParticipant, address _longParticipant, address _shortRecipient, address _longRecipient, uint256 _price, address _sourceAccount, bytes32 _fingerprint) external returns (uint256 _creatorFee, uint256 _reportingFee) {
        // require(isApprovedForAll(_shortParticipant, msg.sender) == true, "ERC1155: need operator approval to burn short account shares");
        // require(isApprovedForAll(_longParticipant, msg.sender) == true, "ERC1155: need operator approval to burn long account shares");

        _internalTransferFrom(_shortParticipant, _longParticipant, getTokenId(_market, _outcome), _amount, bytes(""), false);

        // NOTE: burnCompleteSets will validate the market provided is legitimate
        (uint256 _payout, uint256 _creatorFee, uint256 _reportingFee) = burnCompleteSets(_market, _longParticipant, _amount, _sourceAccount, _fingerprint);

        {
            uint256 _longPayout = _payout.mul(_price) / _market.getNumTicks();
            cashTransfer(_longRecipient, _longPayout);
            cashTransfer(_shortRecipient, _payout.sub(_longPayout));
        }

        assertBalances(_market);
        return (_creatorFee, _reportingFee);
    }

    function burnCompleteSets(IMarket _market, address _account, uint256 _amount, address _sourceAccount, bytes32 _fingerprint) private returns (uint256 _payout, uint256 _creatorFee, uint256 _reportingFee) {
        uint256 _numOutcomes = markets[address(_market)].numOutcomes;
        uint256 _numTicks = markets[address(_market)].numTicks;

        // require(_numOutcomes != 0, "Invalid Market provided");

        // solium-disable indentation
        {
            uint256[] memory _tokenIds = new uint256[](_numOutcomes);
            uint256[] memory _values = new uint256[](_numOutcomes);

            for (uint256 i = 0; i < _numOutcomes; i++) {
                _tokenIds[i] = TokenId.getTokenId(_market, i);
                _values[i] = _amount;
            }

            _burnBatch(_account, _tokenIds, _values, bytes(""), false);
        }
        // solium-enable indentation

        _payout = _amount.mul(_numTicks);
        IUniverse _universe = _market.getUniverse();

        if (!_market.isFinalized()) {
            _universe.decrementOpenInterest(_payout);
        }

        _creatorFee = _market.deriveMarketCreatorFeeAmount(_payout);
        uint256 _reportingFeeDivisor = _universe.getOrCacheReportingFeeDivisor();
        _reportingFee = _payout.div(_reportingFeeDivisor);
        _payout = _payout.sub(_creatorFee).sub(_reportingFee);

        if (_creatorFee != 0) {
            _market.recordMarketCreatorFees(_creatorFee, _sourceAccount, _fingerprint);
        }

        _universe.withdraw(address(this), _payout.add(_reportingFee), address(_market));

        if (_reportingFee != 0) {
            cashTransfer(address(_universe.getOrCreateNextDisputeWindow(false)), _reportingFee);
        }

        augur.logMarketOIChanged(_universe, _market);
    }

    /**
     * @notice Claims winnings for a market and for a particular shareholder
     * @param _market The market to claim winnings for
     * @param _shareHolder The account to claim winnings for
     * @param _fingerprint Fingerprint of the filler used to naively restrict affiliate fee dispursement
     * @return Bool True
     */
    function claimTradingProceeds(IMarket _market, address _shareHolder, bytes32 _fingerprint) external nonReentrant returns (uint256[] memory _outcomeFees) {
        return claimTradingProceedsInternal(_market, _shareHolder, _fingerprint);
    }

    function claimTradingProceedsInternal(IMarket _market, address _shareHolder, bytes32 _fingerprint) internal returns (uint256[] memory _outcomeFees) {
        // require(augur.isKnownMarket(_market));
        if (!_market.isFinalized()) {
            _market.finalize();
        }
        _outcomeFees = new uint256[](8);
        for (uint256 _outcome = 0; _outcome < _market.getNumberOfOutcomes(); ++_outcome) {
            uint256 _numberOfShares = balanceOfMarketOutcome(_market, _outcome, _shareHolder);

            if (_numberOfShares > 0) {
                uint256 _proceeds;
                uint256 _shareHolderShare;
                uint256 _creatorShare;
                uint256 _reporterShare;
                uint256 _tokenId = TokenId.getTokenId(_market, _outcome);
                (_proceeds, _shareHolderShare, _creatorShare, _reporterShare) = divideUpWinnings(_market, _outcome, _numberOfShares);

                // always destroy shares as it gives a minor gas refund and is good for the network
                _burn(_shareHolder, _tokenId, _numberOfShares, bytes(""), false);
                logTradingProceedsClaimed(_market, _outcome, _shareHolder, _numberOfShares, _shareHolderShare, _creatorShare.add(_reporterShare));

                if (_proceeds > 0) {
                    _market.getUniverse().withdraw(address(this), _shareHolderShare.add(_reporterShare), address(_market));
                    distributeProceeds(_market, _shareHolder, _shareHolderShare, _creatorShare, _reporterShare, _fingerprint);
                }
                _outcomeFees[_outcome] = _creatorShare.add(_reporterShare);
            }
        }

        assertBalances(_market);
        return _outcomeFees;
    }

    function distributeProceeds(IMarket _market, address _shareHolder, uint256 _shareHolderShare, uint256 _creatorShare, uint256 _reporterShare, bytes32 _fingerprint) private {
        if (_shareHolderShare > 0) {
            cashTransfer(_shareHolder, _shareHolderShare);
        }
        if (_creatorShare > 0) {
            _market.recordMarketCreatorFees(_creatorShare, _shareHolder, _fingerprint);
        }
        if (_reporterShare > 0) {
            cashTransfer(address(_market.getUniverse().getOrCreateNextDisputeWindow(false)), _reporterShare);
        }
    }

    function logTradingProceedsClaimed(IMarket _market, uint256 _outcome, address _sender, uint256 _numShares, uint256 _numPayoutTokens, uint256 _fees) private {
        augur.logTradingProceedsClaimed(_market.getUniverse(), _sender, address(_market), _outcome, _numShares, _numPayoutTokens, _fees);
    }

    function divideUpWinnings(IMarket _market, uint256 _outcome, uint256 _numberOfShares) public returns (uint256 _proceeds, uint256 _shareHolderShare, uint256 _creatorShare, uint256 _reporterShare) {
        _proceeds = calculateProceeds(_market, _outcome, _numberOfShares);
        _creatorShare = calculateCreatorFee(_market, _proceeds);
        _reporterShare = calculateReportingFee(_market, _proceeds);
        _shareHolderShare = _proceeds.sub(_creatorShare).sub(_reporterShare);
        return (_proceeds, _shareHolderShare, _creatorShare, _reporterShare);
    }

    function calculateProceeds(IMarket _market, uint256 _outcome, uint256 _numberOfShares) public view returns (uint256) {
        uint256 _payoutNumerator = _market.getWinningPayoutNumerator(_outcome);
        return _numberOfShares.mul(_payoutNumerator);
    }

    function calculateReportingFee(IMarket _market, uint256 _amount) public returns (uint256) {
        uint256 _reportingFeeDivisor = _market.getUniverse().getOrCacheReportingFeeDivisor();
        return _amount.div(_reportingFeeDivisor);
    }

    function calculateCreatorFee(IMarket _market, uint256 _amount) public view returns (uint256) {
        return _market.deriveMarketCreatorFeeAmount(_amount);
    }

    function getTypeName() public view returns(bytes32) {
        return "ShareToken";
    }

    /**
     * @return The market associated with this Share Token ID
     */
    function getMarket(uint256 _tokenId) external view returns(IMarket) {
        (address _market, uint256 _outcome) = TokenId.unpackTokenId(_tokenId);
        return IMarket(_market);
    }

    /**
     * @return The outcome associated with this Share Token ID
     */
    function getOutcome(uint256 _tokenId) external view returns(uint256) {
        (address _market, uint256 _outcome) = TokenId.unpackTokenId(_tokenId);
        return _outcome;
    }

    function totalSupplyForMarketOutcome(IMarket _market, uint256 _outcome) public view returns (uint256) {
        uint256 _tokenId = TokenId.getTokenId(_market, _outcome);
        return totalSupply(_tokenId);
    }

    function balanceOfMarketOutcome(IMarket _market, uint256 _outcome, address _account) public view returns (uint256) {
        uint256 _tokenId = TokenId.getTokenId(_market, _outcome);
        return balanceOf(_account, _tokenId);
    }

    function lowestBalanceOfMarketOutcomes(IMarket _market, uint256[] memory _outcomes, address _account) public view returns (uint256) {
        uint256 _lowest = SafeMathUint256.getUint256Max();
        for (uint256 _i = 0; _i < _outcomes.length; ++_i) {
            uint256 _tokenId = TokenId.getTokenId(_market, _outcomes[_i]);
            _lowest = balanceOf(_account, _tokenId).min(_lowest);
        }
        return _lowest;
    }

    function assertBalances(IMarket _market) public view {
        uint256 _expectedBalance = 0;
        uint256 _numTicks = _market.getNumTicks();
        uint256 _numOutcomes = _market.getNumberOfOutcomes();
        // Market Open Interest. If we're finalized we need actually calculate the value
        if (_market.isFinalized()) {
            for (uint8 i = 0; i < _numOutcomes; i++) {
                _expectedBalance = _expectedBalance.add(totalSupplyForMarketOutcome(_market, i).mul(_market.getWinningPayoutNumerator(i)));
            }
        } else {
            _expectedBalance = totalSupplyForMarketOutcome(_market, 0).mul(_numTicks);
        }

        assert(_market.getUniverse().marketBalance(address(_market)) >= _expectedBalance);
    }

    function getTokenId(IMarket _market, uint256 _outcome) public pure returns (uint256 _tokenId) {
        return TokenId.getTokenId(_market, _outcome);
    }

    function getTokenIds(IMarket _market, uint256[] memory _outcomes) public pure returns (uint256[] memory _tokenIds) {
        return TokenId.getTokenIds(_market, _outcomes);
    }

    function unpackTokenId(uint256 _tokenId) public pure returns (address _market, uint256 _outcome) {
        return TokenId.unpackTokenId(_tokenId);
    }

    function onTokenTransfer(uint256 _tokenId, address _from, address _to, uint256 _value) internal {
        (address _marketAddress, uint256 _outcome) = TokenId.unpackTokenId(_tokenId);
        augur.logShareTokensBalanceChanged(_from, IMarket(_marketAddress), _outcome, balanceOf(_from, _tokenId));
        augur.logShareTokensBalanceChanged(_to, IMarket(_marketAddress), _outcome, balanceOf(_to, _tokenId));
    }

    function onMint(uint256 _tokenId, address _target, uint256 _amount) internal {
        (address _marketAddress, uint256 _outcome) = TokenId.unpackTokenId(_tokenId);
        augur.logShareTokensBalanceChanged(_target, IMarket(_marketAddress), _outcome, balanceOf(_target, _tokenId));
    }

    function onBurn(uint256 _tokenId, address _target, uint256 _amount) internal {
        (address _marketAddress, uint256 _outcome) = TokenId.unpackTokenId(_tokenId);
        augur.logShareTokensBalanceChanged(_target, IMarket(_marketAddress), _outcome, balanceOf(_target, _tokenId));
    }
}
