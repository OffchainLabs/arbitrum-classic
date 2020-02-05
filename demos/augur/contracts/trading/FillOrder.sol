pragma solidity 0.5.15;
pragma experimental ABIEncoderV2;


import './IFillOrder.sol';
import '../libraries/ReentrancyGuard.sol';
import '../libraries/math/SafeMathUint256.sol';
import '../reporting/IMarket.sol';
import '../ICash.sol';
import './IOrders.sol';
import '../reporting/IShareToken.sol';
import './IProfitLoss.sol';
import './Order.sol';
import '../libraries/Initializable.sol';
import './IAugurTrading.sol';
import '../libraries/TokenId.sol';
import '../external/IDaiVat.sol';


library Trade {
    using SafeMathUint256 for uint256;

    enum Direction {
        Long,
        Short
    }

    struct StoredContracts {
        IAugur augur;
        IAugurTrading augurTrading;
        IOrders orders;
        ICash denominationToken;
        IProfitLoss profitLoss;
        IShareToken shareToken;
    }

    struct Contracts {
        IOrders orders;
        IMarket market;
        ICash denominationToken;
        IShareToken shareToken;
        IAugur augur;
        IUniverse universe;
        IProfitLoss profitLoss;
        IAugurTrading augurTrading;
    }

    struct FilledOrder {
        bytes32 orderId;
        uint256 outcome;
        uint256 sharePriceRange;
        uint256 sharePriceLong;
        uint256 sharePriceShort;
    }

    struct Participant {
        address participantAddress;
        Direction direction;
        uint256 startingSharesToSell;
        uint256 startingSharesToBuy;
        uint256 sharesToSell;
        uint256 sharesToBuy;
    }

    struct Data {
        Contracts contracts;
        FilledOrder order;
        Participant creator;
        Participant filler;
        uint256 longOutcome;
        uint256[] shortOutcomes;
        address longFundsAccount;
        address shortFundsAccount;
        bytes32 fingerprint;
    }

    struct OrderData {
        IMarket market;
        uint256 outcome;
        uint256 price;
        Order.Types orderType;
        uint256 sharesEscrowed;
        uint256 amount;
        address creator;
        bytes32 orderId;
    }

    //
    // Constructor
    //

    function create(StoredContracts memory _storedContracts, bytes32 _orderId, address _fillerAddress, uint256 _fillerSize, bytes32 _fingerprint) internal view returns (Data memory) {
        OrderData memory _orderData = createOrderDataWithOrderId(_storedContracts, _orderId);

        return createWithData(_storedContracts, _orderData, _fillerAddress, _fillerSize, _fingerprint);
    }

    function createWithData(StoredContracts memory _storedContracts, OrderData memory _orderData, address _fillerAddress, uint256 _fillerSize, bytes32 _fingerprint) internal view returns (Data memory) {
        Contracts memory _contracts = getContracts(_storedContracts, _orderData.market, _orderData.outcome);
        FilledOrder memory _order = getOrder(_contracts, _orderData.outcome, _orderData.price, _orderData.orderId);
        Participant memory _creator = getMaker(_orderData.sharesEscrowed, _orderData.amount, _orderData.creator, _orderData.orderType);
        uint256[] memory _shortOutcomes = getShortOutcomes(_contracts.market, _orderData.outcome);
        Participant memory _filler = getFiller(_contracts, _orderData.outcome, _shortOutcomes, _orderData.orderType, _fillerAddress, _fillerSize);

        // Signed orders which have no order id get their funds from the signed order "creator" whereas on chain orders have funds escrowed in Augur Trading.
        address _creatorFundsSource = _orderData.orderId == bytes32(0) ? _creator.participantAddress : address(_contracts.augurTrading);

        return Data({
            contracts: _contracts,
            order: _order,
            creator: _creator,
            filler: _filler,
            longOutcome: _orderData.outcome,
            shortOutcomes: _shortOutcomes,
            longFundsAccount: _creator.direction == Direction.Long ? _creatorFundsSource : _filler.participantAddress,
            shortFundsAccount: _creator.direction == Direction.Short ? _creatorFundsSource : _filler.participantAddress,
            fingerprint: _fingerprint
        });
    }

    function createOrderDataWithOrderId(StoredContracts memory _storedContracts, bytes32 _orderId) internal view returns (OrderData memory) {
        IOrders _orders = _storedContracts.orders;

        return OrderData({
            market: _orders.getMarket(_orderId),
            outcome: _orders.getOutcome(_orderId),
            price: _orders.getPrice(_orderId),
            orderType: _orders.getOrderType(_orderId),
            sharesEscrowed: _orders.getOrderSharesEscrowed(_orderId),
            amount: _orders.getAmount(_orderId),
            creator: _orders.getOrderCreator(_orderId),
            orderId: _orderId
        });
    }

    function createOrderData(IShareToken _shareToken, IMarket _market, uint256 _outcome, uint256 _price, Order.Types _orderType, uint256 _amount, address _creator) internal view returns (OrderData memory) {
        uint256 _sharesAvailable = getSharesAvailable(_shareToken, _market, _orderType, _outcome, _amount, _creator);

        return OrderData({
            market: _market,
            outcome: _outcome,
            price: _price,
            orderType: _orderType,
            sharesEscrowed: _sharesAvailable,
            amount: _amount,
            creator: _creator,
            orderId: bytes32(0)
        });
    }

    function getSharesAvailable(IShareToken _shareToken, IMarket _market, Order.Types _orderType, uint256 _outcome, uint256 _amount, address _creator) private view returns (uint256) {
        // Figure out how many almost-complete-sets (just missing `outcome` share) the creator has
        uint256 _numberOfOutcomes = _market.getNumberOfOutcomes();
        if (_orderType == Order.Types.Bid) {
            return _shareToken.lowestBalanceOfMarketOutcomes(_market, getShortOutcomes(_market, _outcome), _creator).min(_amount);
        }
        return _shareToken.balanceOfMarketOutcome(_market, _outcome, _creator).min(_amount);
    }

    function getShortOutcomes(IMarket _market, uint256 _outcome) private view returns (uint256[] memory) {
        uint256 _numberOfOutcomes = _market.getNumberOfOutcomes();
        uint256[] memory _shortOutcomes = new uint256[](_numberOfOutcomes - 1);
        uint256 _indexOutcome = 0;
        for (uint256 _i = 0; _i < _numberOfOutcomes - 1; _i++) {
            if (_i == _outcome) {
                _indexOutcome++;
            }
            _shortOutcomes[_i] = _indexOutcome;
            _indexOutcome++;
        }
        return _shortOutcomes;
    }

    //
    // "public" functions
    //

    function tradeMakerSharesForFillerShares(Data memory _data) internal returns (uint256 _marketCreatorFees, uint256 _reporterFees) {
        uint256 _numberOfCompleteSets = _data.creator.sharesToSell.min(_data.filler.sharesToSell);
        if (_numberOfCompleteSets == 0) {
            return (0, 0);
        }

        // transfer shares and sell complete sets distributing payouts based on the price

        // Sell both account shares
        (_marketCreatorFees, _reporterFees) = _data.contracts.shareToken.sellCompleteSetsForTrade(_data.contracts.market, _data.longOutcome, _numberOfCompleteSets, _data.shortFundsAccount, _data.longFundsAccount, getShortShareSellerDestination(_data), getLongShareSellerDestination(_data), _data.order.sharePriceLong, _data.filler.participantAddress, _data.fingerprint);

        // update available shares for creator and filler
        _data.creator.sharesToSell -= _numberOfCompleteSets;
        _data.filler.sharesToSell -= _numberOfCompleteSets;
        return (_marketCreatorFees, _reporterFees);
    }

    function tradeMakerSharesForFillerTokens(Data memory _data) internal returns (bool) {
        uint256 _numberOfSharesToTrade = _data.creator.sharesToSell.min(_data.filler.sharesToBuy);
        if (_numberOfSharesToTrade == 0) {
            return true;
        }

        // transfer shares from creator (escrowed in market) to filler
        if (_data.creator.direction == Direction.Short) {
            _data.contracts.shareToken.unsafeTransferFrom(_data.shortFundsAccount, _data.filler.participantAddress, TokenId.getTokenId(_data.contracts.market, _data.longOutcome), _numberOfSharesToTrade);
        } else {
            uint256[] memory _values = new uint256[](_data.shortOutcomes.length);
            for (uint256 _i = 0; _i < _data.shortOutcomes.length; _i++) {
                _values[_i] = _numberOfSharesToTrade;
            }
            _data.contracts.shareToken.unsafeBatchTransferFrom(_data.longFundsAccount, _data.filler.participantAddress, TokenId.getTokenIds(_data.contracts.market, _data.shortOutcomes), _values);
        }

        uint256 _tokensToCover = getTokensToCover(_data, _data.filler.direction, _numberOfSharesToTrade);
        _data.contracts.denominationToken.transferFrom(_data.filler.participantAddress, _data.creator.participantAddress, _tokensToCover);

        // update available assets for creator and filler
        _data.creator.sharesToSell -= _numberOfSharesToTrade;
        _data.filler.sharesToBuy -= _numberOfSharesToTrade;
        return true;
    }

    function tradeMakerTokensForFillerShares(Data memory _data) internal returns (bool) {
        uint256 _numberOfSharesToTrade = _data.filler.sharesToSell.min(_data.creator.sharesToBuy);
        if (_numberOfSharesToTrade == 0) {
            return true;
        }

        // transfer shares from filler to creator
        if (_data.filler.direction == Direction.Short) {
            _data.contracts.shareToken.unsafeTransferFrom(_data.filler.participantAddress, _data.creator.participantAddress, TokenId.getTokenId(_data.contracts.market, _data.longOutcome), _numberOfSharesToTrade);
        } else {
            uint256[] memory _values = new uint256[](_data.shortOutcomes.length);
            for (uint256 _i = 0; _i < _data.shortOutcomes.length; _i++) {
                _values[_i] = _numberOfSharesToTrade;
            }
            _data.contracts.shareToken.unsafeBatchTransferFrom(_data.filler.participantAddress, _data.creator.participantAddress, TokenId.getTokenIds(_data.contracts.market, _data.shortOutcomes), _values);
        }

        // transfer tokens from creator (taken from the signer for signed orders, escrowed in Augur Trading for on chain orders) to filler
        uint256 _tokensToCover = getTokensToCover(_data, _data.creator.direction, _numberOfSharesToTrade);
        if (_data.order.orderId == bytes32(0)) {
            // No order Id indicates this is a signed order
            _data.contracts.denominationToken.transferFrom(_data.creator.participantAddress, _data.filler.participantAddress, _tokensToCover);
        } else {
            _data.contracts.denominationToken.transferFrom(address(_data.contracts.augurTrading), _data.filler.participantAddress, _tokensToCover);
        }

        // update available assets for creator and filler
        _data.creator.sharesToBuy -= _numberOfSharesToTrade;
        _data.filler.sharesToSell -= _numberOfSharesToTrade;
        return true;
    }

    function tradeMakerTokensForFillerTokens(Data memory _data) internal returns (uint256) {
        uint256 _numberOfCompleteSets = _data.creator.sharesToBuy.min(_data.filler.sharesToBuy);
        if (_numberOfCompleteSets == 0) {
            return 0;
        }

        // If someone is filling their own order with CASH both ways we just return the CASH
        if (_data.creator.participantAddress == _data.filler.participantAddress) {
            uint256 _creatorTokensToCover = getTokensToCover(_data, _data.creator.direction, _numberOfCompleteSets);
            uint256 _fillerTokensToCover = getTokensToCover(_data, _data.filler.direction, _numberOfCompleteSets);
            _data.contracts.denominationToken.transferFrom(address(_data.contracts.augurTrading), _data.creator.participantAddress, _creatorTokensToCover);

            _data.creator.sharesToBuy -= _numberOfCompleteSets;
            _data.filler.sharesToBuy -= _numberOfCompleteSets;
            return _creatorTokensToCover.add(_fillerTokensToCover);
        }

        // buy complete sets and distribute shares to participants
        uint256 _longCost = _numberOfCompleteSets.mul(_data.order.sharePriceLong);
        uint256 _shortCost = _numberOfCompleteSets.mul(_data.order.sharePriceShort);
        IUniverse _universe = _data.contracts.market.getUniverse();

        // Bring in cash from both parties
        if (_data.longFundsAccount == address(_data.contracts.market)) {
            _data.contracts.denominationToken.transferFrom(address(_data.contracts.augurTrading), address(this), _longCost);
        } else {
            _data.contracts.denominationToken.transferFrom(_data.longFundsAccount, address(this), _longCost);
        }

        if (_data.shortFundsAccount == address(_data.contracts.market)) {
            _data.contracts.denominationToken.transferFrom(address(_data.contracts.augurTrading), address(this), _shortCost);
        } else {
            _data.contracts.denominationToken.transferFrom(_data.shortFundsAccount, address(this), _shortCost);
        }

        // Buy and distribute complete sets
        address _longRecipient = getLongShareBuyerDestination(_data);
        address _shortRecipient = getShortShareBuyerDestination(_data);
        _data.contracts.shareToken.buyCompleteSetsForTrade(_data.contracts.market, _numberOfCompleteSets, _data.order.outcome, _longRecipient, _shortRecipient);

        _data.creator.sharesToBuy -= _numberOfCompleteSets;
        _data.filler.sharesToBuy -= _numberOfCompleteSets;
        return 0;
    }

    //
    // Helpers
    //

    function getLongShareBuyerDestination(Data memory _data) internal pure returns (address) {
        return (_data.creator.direction == Direction.Long) ? _data.creator.participantAddress : _data.filler.participantAddress;
    }

    function getShortShareBuyerDestination(Data memory _data) internal pure returns (address) {
        return (_data.creator.direction == Direction.Short) ? _data.creator.participantAddress : _data.filler.participantAddress;
    }

    function getLongShareSellerSource(Data memory _data) internal pure returns (address) {
        return (_data.creator.direction == Direction.Short) ? address(_data.contracts.market) : _data.filler.participantAddress;
    }

    function getShortShareSellerSource(Data memory _data) internal pure returns (address) {
        return (_data.creator.direction == Direction.Long) ? address(_data.contracts.market) : _data.filler.participantAddress;
    }

    function getLongShareSellerDestination(Data memory _data) internal pure returns (address) {
        return (_data.creator.direction == Direction.Short) ? _data.creator.participantAddress : _data.filler.participantAddress;
    }

    function getShortShareSellerDestination(Data memory _data) internal pure returns (address) {
        return (_data.creator.direction == Direction.Long) ? _data.creator.participantAddress : _data.filler.participantAddress;
    }

    function getMakerSharesDepleted(Data memory _data) internal pure returns (uint256) {
        return _data.creator.startingSharesToSell.sub(_data.creator.sharesToSell);
    }

    function getFillerSharesDepleted(Data memory _data) internal pure returns (uint256) {
        return _data.filler.startingSharesToSell.sub(_data.filler.sharesToSell);
    }

    function getMakerTokensDepleted(Data memory _data) internal pure returns (uint256) {
        return getTokensDepleted(_data, _data.creator.direction, _data.creator.startingSharesToBuy, _data.creator.sharesToBuy);
    }

    function getFillerTokensDepleted(Data memory _data) internal pure returns (uint256) {
        return getTokensDepleted(_data, _data.filler.direction, _data.filler.startingSharesToBuy, _data.filler.sharesToBuy);
    }

    function getTokensDepleted(Data memory _data, Direction _direction, uint256 _startingSharesToBuy, uint256 _endingSharesToBuy) internal pure returns (uint256) {
        return _startingSharesToBuy
            .sub(_endingSharesToBuy)
            .mul((_direction == Direction.Long) ? _data.order.sharePriceLong : _data.order.sharePriceShort);
    }

    function getTokensToCover(Data memory _data, Direction _direction, uint256 _numShares) internal pure returns (uint256) {
        return getTokensToCover(_direction, _data.order.sharePriceLong, _data.order.sharePriceShort, _numShares);
    }

    //
    // Construction helpers
    //

    function getContracts(StoredContracts memory _storedContracts, IMarket _market, uint256 _outcome) private view returns (Contracts memory) {
        return Contracts({
            orders: _storedContracts.orders,
            market: _market,
            denominationToken: _storedContracts.denominationToken,
            augur: _storedContracts.augur,
            augurTrading: _storedContracts.augurTrading,
            universe: _market.getUniverse(),
            profitLoss: _storedContracts.profitLoss,
            shareToken: _storedContracts.shareToken
        });
    }

    function getOrder(Contracts memory _contracts, uint256 _outcome, uint256 _price, bytes32 _orderId) private view returns (FilledOrder memory) {
        uint256 _sharePriceRange;
        uint256 _sharePriceLong;
        uint256 _sharePriceShort;
        (_sharePriceRange, _sharePriceLong, _sharePriceShort) = getSharePriceDetails(_contracts.market, _price);
        return FilledOrder({
            orderId: _orderId,
            outcome: _outcome,
            sharePriceRange: _sharePriceRange,
            sharePriceLong: _sharePriceLong,
            sharePriceShort: _sharePriceShort
        });
    }

    function getMaker(uint256 _sharesEscrowed, uint256 _amount, address _creator, Order.Types _orderOrderType) private pure returns (Participant memory) {
        Direction _direction = (_orderOrderType == Order.Types.Bid) ? Direction.Long : Direction.Short;
        uint256 _sharesToBuy = _amount.sub(_sharesEscrowed);
        return Participant({
            participantAddress: _creator,
            direction: _direction,
            startingSharesToSell: _sharesEscrowed,
            startingSharesToBuy: _sharesToBuy,
            sharesToSell: _sharesEscrowed,
            sharesToBuy: _sharesToBuy
        });
    }

    function getFiller(Contracts memory _contracts, uint256 _longOutcome, uint256[] memory _shortOutcomes, Order.Types _orderOrderType, address _address, uint256 _size) private view returns (Participant memory) {
        Direction _direction = (_orderOrderType == Order.Types.Bid) ? Direction.Short : Direction.Long;
        uint256 _sharesToSell = 0;
        _sharesToSell = getFillerSharesToSell(_contracts, _longOutcome, _shortOutcomes, _address, _direction, _size);
        uint256 _sharesToBuy = _size.sub(_sharesToSell);
        return Participant({
            participantAddress: _address,
            direction: _direction,
            startingSharesToSell: _sharesToSell,
            startingSharesToBuy: _sharesToBuy,
            sharesToSell: _sharesToSell,
            sharesToBuy: _sharesToBuy
        });
    }

    function getTokensToCover(Direction _direction, uint256 _sharePriceLong, uint256 _sharePriceShort, uint256 _numShares) internal pure returns (uint256) {
        return _numShares.mul((_direction == Direction.Long) ? _sharePriceLong : _sharePriceShort);
    }

    function getSharePriceDetails(IMarket _market, uint256 _price) private view returns (uint256 _sharePriceRange, uint256 _sharePriceLong, uint256 _sharePriceShort) {
        uint256 _numTicks = _market.getNumTicks();
        _sharePriceShort = uint256(_numTicks.sub(_price));
        return (_numTicks, _price, _sharePriceShort);
    }

    function getFillerSharesToSell(Contracts memory _contracts, uint256 _longOutcome, uint256[] memory _shortOutcomes, address _filler, Direction _fillerDirection, uint256 _fillerSize) private view returns (uint256) {
        if (_fillerDirection == Direction.Short) {
            return _contracts.shareToken.balanceOfMarketOutcome(_contracts.market, _longOutcome, _filler).min(_fillerSize);
        }
        return _contracts.shareToken.lowestBalanceOfMarketOutcomes(_contracts.market, _shortOutcomes, _filler).min(_fillerSize);
    }
}


/**
 * @title Fill Order
 * @notice Exposes functions to fill an order on the book
 */
contract FillOrder is Initializable, ReentrancyGuard, IFillOrder {
    using SafeMathUint256 for uint256;
    using Trade for Trade.Data;

    IAugur public augur;
    IAugurTrading public augurTrading;
    address public zeroXTrade;
    address public trade;

    Trade.StoredContracts private storedContracts;

    mapping (address => uint256[]) public marketOutcomeVolumes;

    uint256 private constant MAX_APPROVAL_AMOUNT = 2 ** 256 - 1;

    function initialize(IAugur _augur, IAugurTrading _augurTrading) public beforeInitialized {
        endInitialization();
        augur = _augur;
        augurTrading = _augurTrading;
        ICash _cash = ICash(augur.lookup("Cash"));
        storedContracts = Trade.StoredContracts({
            augur: _augur,
            augurTrading: _augurTrading,
            orders: IOrders(_augurTrading.lookup("Orders")),
            denominationToken: _cash,
            profitLoss: IProfitLoss(_augurTrading.lookup("ProfitLoss")),
            shareToken: IShareToken(_augur.lookup("ShareToken"))
        });
        require(storedContracts.orders != IOrders(0));
        require(storedContracts.profitLoss != IProfitLoss(0));
        require(storedContracts.shareToken != IShareToken(0));
        trade = _augurTrading.lookup("Trade");
        require(trade != address(0));
        zeroXTrade = _augurTrading.lookup("ZeroXTrade");
        require(zeroXTrade != address(0));
        _cash.approve(address(_augur), MAX_APPROVAL_AMOUNT);
        IDaiVat _vat = IDaiVat(augur.lookup("DaiVat"));
        _vat.hope(address(augur));
    }

    /**
     * @notice Fill an order
     * @param _orderId The id of the order to fill
     * @param _amountFillerWants The number of attoShares desired
     * @param _tradeGroupId A Bytes32 value used when attempting to associate multiple orderbook actions with a single TX
     * @param _fingerprint Fingerprint of the filler used to naively restrict affiliate fee dispursement
     * @return The amount remaining the filler wants
     */
    function publicFillOrder(bytes32 _orderId, uint256 _amountFillerWants, bytes32 _tradeGroupId, bytes32 _fingerprint) external returns (uint256) {
        address _filler = msg.sender;
        Trade.Data memory _tradeData = Trade.create(storedContracts, _orderId, _filler, _amountFillerWants, _fingerprint);
        (uint256 _amountRemaining, uint256 _fees) = fillOrderInternal(_filler, _tradeData, _amountFillerWants, _tradeGroupId);
        return _amountRemaining;
    }

    function fillOrder(address _filler, bytes32 _orderId, uint256 _amountFillerWants, bytes32 _tradeGroupId, bytes32 _fingerprint) external returns (uint256) {
        require(msg.sender == trade);
        Trade.Data memory _tradeData = Trade.create(storedContracts, _orderId, _filler, _amountFillerWants, _fingerprint);
        (uint256 _amountRemaining, uint256 _fees) = fillOrderInternal(_filler, _tradeData, _amountFillerWants, _tradeGroupId);
        return _amountRemaining;
    }

    function fillZeroXOrder(IMarket _market, uint256 _outcome, uint256 _price, Order.Types _orderType, address _creator, uint256 _amount, bytes32 _fingerprint, bytes32 _tradeGroupId, address _filler) external returns (uint256 _amountRemaining, uint256 _fees) {
        require(msg.sender == zeroXTrade);
        require(augur.isKnownMarket(_market));
        Trade.OrderData memory _orderData = Trade.createOrderData(storedContracts.shareToken, _market, _outcome, _price, _orderType, _amount, _creator);
        Trade.Data memory _tradeData = Trade.createWithData(storedContracts, _orderData, _filler, _amount, _fingerprint);
        return fillOrderInternal(_filler, _tradeData, _amount, _tradeGroupId);
    }


    function fillOrderInternal(address _filler, Trade.Data memory _tradeData, uint256 _amountFillerWants, bytes32 _tradeGroupId) internal nonReentrant returns (uint256 _amountRemainingFillerWants, uint256 _totalFees) {
        uint256 _marketCreatorFees;
        uint256 _reporterFees;

        (_marketCreatorFees, _reporterFees) = _tradeData.tradeMakerSharesForFillerShares();
        _tradeData.tradeMakerTokensForFillerShares();
        _tradeData.tradeMakerSharesForFillerTokens();
        uint256 _tokensRefunded = _tradeData.tradeMakerTokensForFillerTokens();

        sellCompleteSets(_tradeData);

        uint256 _amountRemainingFillerWants = _tradeData.filler.sharesToSell.add(_tradeData.filler.sharesToBuy);
        uint256 _amountFilled = _amountFillerWants.sub(_amountRemainingFillerWants);
        if (_tradeData.order.orderId != bytes32(0)) {
            _tradeData.contracts.orders.recordFillOrder(_tradeData.order.orderId, _tradeData.getMakerSharesDepleted(), _tradeData.getMakerTokensDepleted(), _amountFilled);
        }
        _totalFees = _marketCreatorFees.add(_reporterFees);
        if (_tradeData.order.orderId != bytes32(0)) {
            logOrderFilled(_tradeData, _tradeData.order.sharePriceLong, _totalFees, _amountFilled, _tradeGroupId);
        }
        logAndUpdateVolume(_tradeData);
        if (_totalFees > 0) {
            uint256 _longFees = _totalFees.mul(_tradeData.order.sharePriceLong).div(_tradeData.contracts.market.getNumTicks());
            uint256 _shortFees = _totalFees.sub(_longFees);
            _tradeData.contracts.profitLoss.adjustTraderProfitForFees(_tradeData.contracts.market, _tradeData.getLongShareBuyerDestination(), _tradeData.order.outcome, _longFees);
            _tradeData.contracts.profitLoss.adjustTraderProfitForFees(_tradeData.contracts.market, _tradeData.getShortShareBuyerDestination(), _tradeData.order.outcome, _shortFees);
        }
        updateProfitLoss(_tradeData, _amountFilled);
        if (_tradeData.creator.participantAddress == _tradeData.filler.participantAddress) {
            _tradeData.contracts.profitLoss.recordFrozenFundChange(_tradeData.contracts.universe, _tradeData.contracts.market, _tradeData.creator.participantAddress, _tradeData.order.outcome, -int256(_tokensRefunded));
        }

        return (_amountRemainingFillerWants, _totalFees);
    }

    function sellCompleteSets(Trade.Data memory _tradeData) internal returns (bool) {
        address _filler = _tradeData.filler.participantAddress;
        address _creator = _tradeData.creator.participantAddress;
        IMarket _market = _tradeData.contracts.market;
        uint256 _numOutcomes = _market.getNumberOfOutcomes();

        uint256[] memory _outcomes = new uint256[](_numOutcomes);
        for (uint256 _i = 0; _i < _numOutcomes; _i++) {
            _outcomes[_i] = _i;
        }
        uint256 _fillerCompleteSets = _tradeData.contracts.shareToken.lowestBalanceOfMarketOutcomes(_market, _outcomes, _filler);
        uint256 _creatorCompleteSets = _tradeData.contracts.shareToken.lowestBalanceOfMarketOutcomes(_market, _outcomes, _creator);

        if (_fillerCompleteSets > 0) {
            _tradeData.contracts.shareToken.sellCompleteSets(_market, _filler, _filler, _fillerCompleteSets, _tradeData.fingerprint);
        }

        if (_creatorCompleteSets > 0) {
            _tradeData.contracts.shareToken.sellCompleteSets(_market, _creator, _creator, _creatorCompleteSets, _tradeData.fingerprint);
        }

        return true;
    }

    function logOrderFilled(Trade.Data memory _tradeData, uint256 _price, uint256 _fees, uint256 _amountFilled, bytes32 _tradeGroupId) private returns (bool) {
        _tradeData.contracts.augurTrading.logOrderFilled(_tradeData.contracts.universe, _tradeData.creator.participantAddress, _tradeData.filler.participantAddress, _price, _fees, _amountFilled, _tradeData.order.orderId, _tradeGroupId);
        return true;
    }

    function logAndUpdateVolume(Trade.Data memory _tradeData) private {
        IMarket _market = _tradeData.contracts.market;
        uint256 _makerSharesDepleted = _tradeData.getMakerSharesDepleted();
        uint256 _fillerSharesDepleted = _tradeData.getFillerSharesDepleted();
        uint256 _makerTokensDepleted = _tradeData.getMakerTokensDepleted();
        uint256 _fillerTokensDepleted = _tradeData.getFillerTokensDepleted();
        uint256 _completeSetTokens = _makerSharesDepleted.min(_fillerSharesDepleted).mul(_market.getNumTicks());
        if (marketOutcomeVolumes[address(_market)].length == 0) {
            marketOutcomeVolumes[address(_market)].length = _tradeData.shortOutcomes.length + 1;
        }
        marketOutcomeVolumes[address(_market)][_tradeData.order.outcome] = marketOutcomeVolumes[address(_market)][_tradeData.order.outcome].add(_makerTokensDepleted).add(_fillerTokensDepleted).add(_completeSetTokens);

        uint256[] memory tmpMarketOutcomeVolumes = marketOutcomeVolumes[address(_market)];
        uint256 _volume;
        for (uint256 i = 0; i < tmpMarketOutcomeVolumes.length; i++) {
            _volume += tmpMarketOutcomeVolumes[i];
        }

        _tradeData.contracts.augurTrading.logMarketVolumeChanged(_tradeData.contracts.universe, address(_market), _volume, tmpMarketOutcomeVolumes);
    }

    function updateProfitLoss(Trade.Data memory _tradeData, uint256 _amountFilled) private returns (bool) {
        uint256 makerTokensDepleted = _tradeData.order.orderId != bytes32(0) ? 0 : _tradeData.getMakerTokensDepleted();
        uint256 _numLongTokens = _tradeData.creator.direction == Trade.Direction.Long ? makerTokensDepleted : _tradeData.getFillerTokensDepleted();
        uint256 _numShortTokens = _tradeData.creator.direction == Trade.Direction.Short ? makerTokensDepleted : _tradeData.getFillerTokensDepleted();
        uint256 _numLongShares = _tradeData.creator.direction == Trade.Direction.Long ? _tradeData.getMakerSharesDepleted() : _tradeData.getFillerSharesDepleted();
        uint256 _numShortShares = _tradeData.creator.direction == Trade.Direction.Short ? _tradeData.getMakerSharesDepleted() : _tradeData.getFillerSharesDepleted();
        _tradeData.contracts.profitLoss.recordTrade(_tradeData.contracts.universe, _tradeData.contracts.market, _tradeData.getLongShareBuyerDestination(), _tradeData.getShortShareBuyerDestination(), _tradeData.order.outcome, int256(_amountFilled), int256(_tradeData.order.sharePriceLong), _numLongTokens, _numShortTokens, _numLongShares, _numShortShares);
        return true;
    }

    function getMarketOutcomeValues(IMarket _market) public view returns (uint256[] memory) {
        return marketOutcomeVolumes[address(_market)];
    }

    function getMarketVolume(IMarket _market) public view returns (uint256) {
        uint256[] memory tmpMarketOutcomeVolumes = marketOutcomeVolumes[address(_market)];
        uint256 _volume;
        for (uint256 i = 0; i < tmpMarketOutcomeVolumes.length; i++) {
            _volume += tmpMarketOutcomeVolumes[i];
        }
        return _volume;
    }
}
