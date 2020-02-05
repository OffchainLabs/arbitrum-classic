// Copyright (C) 2015 Forecast Foundation OU, full GPL notice in LICENSE

pragma solidity 0.5.15;


import '../IAugur.sol';
import '../libraries/ReentrancyGuard.sol';
import './Order.sol';
import '../reporting/IMarket.sol';
import './ICreateOrder.sol';
import './IOrders.sol';
import './IFillOrder.sol';
import '../libraries/Initializable.sol';
import '../libraries/token/IERC20.sol';
import './IAugurTrading.sol';


/**
 * @title Trade
 * @notice Provides functions to perform actions on an orderbook which will fill or create orders as appropriate
 */
contract Trade is Initializable, ReentrancyGuard {

    struct Data {
        Order.TradeDirections direction;
        IMarket market;
        uint256 outcome;
        uint256 amount;
        uint256 price;
        bytes32 betterOrderId;
        bytes32 worseOrderId;
        bytes32 tradeGroupId;
        uint256 loopLimit;
        bytes32 fingerprint;
        address sender;
    }

    IAugur public augur;
    ICreateOrder public createOrder;
    IFillOrder public fillOrder;
    IOrders public orders;
    ICash public cash;

    address private constant NULL_ADDRESS = address(0);

    function initialize(IAugur _augur, IAugurTrading _augurTrading) public beforeInitialized {
        endInitialization();
        augur = _augur;
        cash = ICash(augur.lookup("Cash"));
        require(cash != ICash(0));

        createOrder = ICreateOrder(_augurTrading.lookup("CreateOrder"));
        fillOrder = IFillOrder(_augurTrading.lookup("FillOrder"));
        orders = IOrders(_augurTrading.lookup("Orders"));
        require(createOrder != ICreateOrder(0));
        require(fillOrder != IFillOrder(0));
        require(orders != IOrders(0));
    }

    function create(Order.TradeDirections _direction, IMarket _market, uint256 _outcome, uint256 _amount, uint256 _price, bytes32 _betterOrderId, bytes32 _worseOrderId, bytes32 _tradeGroupId, uint256 _loopLimit, bytes32 _fingerprint, address _sender) internal pure returns (Data memory) {
        require(_amount > 0, "Trade.create: Trade amount cannot be 0");

        return Data({
            direction: _direction,
            market: _market,
            outcome: _outcome,
            amount: _amount,
            price: _price,
            betterOrderId: _betterOrderId,
            worseOrderId: _worseOrderId,
            tradeGroupId: _tradeGroupId,
            loopLimit: _loopLimit,
            fingerprint: _fingerprint,
            sender: _sender
        });
    }

    /**
     * @notice Perform a trade which will fill orders and if a desired amount remains will create an order
     * @param _direction The trade direction of order. Either LONG==0, or SHORT==1
     * @param _market The associated market
     * @param _outcome The associated outcome of the market
     * @param _amount The number of attoShares desired
     * @param _price The price in attoCash. Must be within the market range (1 to numTicks-1)
     * @param _betterOrderId The id of an order which is better than this one. Used to reduce gas costs when sorting
     * @param _worseOrderId The id of an order which is worse than this one. Used to reduce gas costs when sorting
     * @param _tradeGroupId A Bytes32 value used when attempting to associate multiple orderbook actions with a single TX
     * @param _loopLimit The number of orders to take from the book before completing the tx. Used to limit gas cost
     * @param _fingerprint Fingerprint of the filler used to naively restrict affiliate fee dispursement
     * @return The Bytes32 orderid of the created order
     */
    function publicTrade(Order.TradeDirections _direction, IMarket _market, uint256 _outcome, uint256 _amount, uint256 _price, bytes32 _betterOrderId, bytes32 _worseOrderId, bytes32 _tradeGroupId, uint256 _loopLimit, bytes32 _fingerprint) external returns (bytes32) {
        return internalTrade(_direction, _market, _outcome, _amount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit, _fingerprint, msg.sender);
    }

    /**
     * @notice Perform a trade which will only fill orders
     * @param _direction The trade direction of order. Either LONG==0, or SHORT==1
     * @param _market The associated market
     * @param _outcome The associated outcome of the market
     * @param _amount The number of attoShares desired
     * @param _price The price in attoCash. Must be within the market range (1 to numTicks-1)
     * @param _tradeGroupId A Bytes32 value used when attempting to associate multiple orderbook actions with a single TX
     * @param _loopLimit The number of orders to take from the book before completing the tx. Used to limit gas cost
     * @param _fingerprint Fingerprint of the filler used to naively restrict affiliate fee dispursement
     * @return The desired amount remaining that was not filled
     */
    function publicFillBestOrder(Order.TradeDirections _direction, IMarket _market, uint256 _outcome, uint256 _amount, uint256 _price, bytes32 _tradeGroupId, uint256 _loopLimit, bytes32 _fingerprint) external returns (uint256) {
        return internalFillBestOrder(_direction, _market, _outcome, _amount, _price, _tradeGroupId, _loopLimit, _fingerprint, msg.sender);
    }

    function internalTrade(Order.TradeDirections _direction, IMarket _market, uint256 _outcome, uint256 _amount, uint256 _price, bytes32 _betterOrderId, bytes32 _worseOrderId, bytes32 _tradeGroupId, uint256 _loopLimit, bytes32 _fingerprint, address _sender) internal returns (bytes32) {
        require(augur.isKnownMarket(_market));
        Data memory _tradeData = create(_direction, _market, _outcome, _amount, _price, _betterOrderId, _worseOrderId, _tradeGroupId, _loopLimit, _fingerprint, _sender);
        bytes32 _result = trade(_tradeData);
        return _result;
    }

    function internalFillBestOrder(Order.TradeDirections _direction, IMarket _market, uint256 _outcome, uint256 _amount, uint256 _price, bytes32 _tradeGroupId, uint256 _loopLimit, bytes32 _fingerprint, address _sender) internal returns (uint256) {
        require(augur.isKnownMarket(_market));
        Data memory _tradeData = create(_direction, _market, _outcome, _amount, _price, bytes32(0), bytes32(0), _tradeGroupId, _loopLimit, _fingerprint, _sender);
        uint256 _result = fillBestOrder(_tradeData);
        return _result;
    }

    function trade(Data memory _tradeData) internal returns (bytes32) {
        uint256 _bestAmount = fillBestOrder(_tradeData);
        if (_bestAmount == 0) {
            return bytes32(uint256(1));
        }
        return createOrder.createOrder(_tradeData.sender, Order.getOrderTradingTypeFromMakerDirection(_tradeData.direction), _bestAmount, _tradeData.price, _tradeData.market, _tradeData.outcome, _tradeData.betterOrderId, _tradeData.worseOrderId, _tradeData.tradeGroupId);
    }

    function fillBestOrder(Data memory _tradeData) internal nonReentrant returns (uint256 _bestAmount) {
        IOrders _orders = orders;
        // we need to fill a BID if we want to SELL and we need to fill an ASK if we want to BUY
        Order.Types _type = Order.getOrderTradingTypeFromFillerDirection(_tradeData.direction);
        bytes32 _orderId = _orders.getBestOrderId(_type, _tradeData.market, _tradeData.outcome);
        _bestAmount = _tradeData.amount;
        uint256 _orderPrice = _orders.getPrice(_orderId);
        uint256 _lastTradePrice = 0;
        // If the price is acceptable relative to the trade type
        while (_orderId != 0 && _bestAmount > 0 && _tradeData.loopLimit > 0 && isMatch(_orderId, _type, _orderPrice, _tradeData.price)) {
            bytes32 _nextOrderId = _orders.getWorseOrderId(_orderId);
            _lastTradePrice = _orderPrice;
            _bestAmount = fillOrder.fillOrder(_tradeData.sender, _orderId, _bestAmount, _tradeData.tradeGroupId, _tradeData.fingerprint);
            _orderId = _nextOrderId;
            _orderPrice = _orders.getPrice(_orderId);
            _tradeData.loopLimit -= 1;
        }

        if (_lastTradePrice != 0) {
            _orders.setPrice(_tradeData.market, _tradeData.outcome, _lastTradePrice);
        }

        if (isMatch(_orderId, _type, _orderPrice, _tradeData.price)) {
            return 0;
        }
        return _bestAmount;
    }

    function isMatch(bytes32 _orderId, Order.Types _type, uint256 _orderPrice, uint256 _price) private pure returns (bool) {
        if (_orderId == 0) {
            return false;
        }
        return _type == Order.Types.Bid ? _orderPrice >= _price : _orderPrice <= _price;
    }
}
