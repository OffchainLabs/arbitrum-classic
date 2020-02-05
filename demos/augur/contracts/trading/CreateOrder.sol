// Copyright (C) 2015 Forecast Foundation OU, full GPL notice in LICENSE

pragma solidity 0.5.15;


import '../libraries/ReentrancyGuard.sol';
import './Order.sol';
import './ICreateOrder.sol';
import '../libraries/Initializable.sol';
import '../libraries/token/IERC20.sol';
import '../IAugur.sol';
import './IProfitLoss.sol';
import './IAugurTrading.sol';
import '../libraries/math/SafeMathUint256.sol';
import '../CashSender.sol';
import '../libraries/TokenId.sol';


/**
 * @title Create Order
 * @notice Exposes functions to place an order on the book for other parties to take
 */
contract CreateOrder is Initializable, ReentrancyGuard, CashSender {
    using SafeMathUint256 for uint256;
    using Order for Order.Data;

    IAugur public augur;
    IAugurTrading public augurTrading;
    address public trade;
    IProfitLoss public profitLoss;
    IOrders public orders;


    function initialize(IAugur _augur, IAugurTrading _augurTrading) public beforeInitialized {
        endInitialization();
        augur = _augur;

        augurTrading = _augurTrading;
        trade = _augurTrading.lookup("Trade");
        require(trade != address(0));
        profitLoss = IProfitLoss(_augurTrading.lookup("ProfitLoss"));
        require(profitLoss != IProfitLoss(0));
        orders = IOrders(_augurTrading.lookup("Orders"));

        initializeCashSender(_augur.lookup("DaiVat"), _augur.lookup("Cash"));
        require(orders != IOrders(0));
    }

    /**
     * @notice Create an order
     * @param _type The type of order. Either BID==0, or ASK==1
     * @param _attoshares The number of attoShares desired
     * @param _price The price in attoCash. Must be within the market range (1 to numTicks-1)
     * @param _market The associated market
     * @param _outcome The associated outcome of the market
     * @param _betterOrderId The id of an order which is better than this one. Used to reduce gas costs when sorting
     * @param _worseOrderId The id of an order which is worse than this one. Used to reduce gas costs when sorting
     * @param _tradeGroupId A Bytes32 value used when attempting to associate multiple orderbook actions with a single TX
     * @return The Bytes32 orderid of the created order
     */
    function publicCreateOrder(Order.Types _type, uint256 _attoshares, uint256 _price, IMarket _market, uint256 _outcome, bytes32 _betterOrderId, bytes32 _worseOrderId, bytes32 _tradeGroupId) external returns (bytes32) {
        bytes32 _result = this.createOrder(msg.sender, _type, _attoshares, _price, _market, _outcome, _betterOrderId, _worseOrderId, _tradeGroupId);
        return _result;
    }

    function createOrder(address _creator, Order.Types _type, uint256 _attoshares, uint256 _price, IMarket _market, uint256 _outcome, bytes32 _betterOrderId, bytes32 _worseOrderId, bytes32 _tradeGroupId) external nonReentrant returns (bytes32) {
        require(augur.isKnownMarket(_market));
        require(msg.sender == trade || msg.sender == address(this));
        Order.Data memory _orderData = Order.create(augur, augurTrading, _creator, _outcome, _type, _attoshares, _price, _market, _betterOrderId, _worseOrderId);
        escrowFunds(_orderData);
        profitLoss.recordFrozenFundChange(_market.getUniverse(), _market, _creator, _outcome, int256(_orderData.moneyEscrowed));
        /* solium-disable indentation */
        {
            IOrders _orders = orders;
            require(_orders.getAmount(Order.getOrderId(_orderData, _orders)) == 0, "Createorder.createOrder: Order duplication in same block");
            return Order.saveOrder(_orderData, _tradeGroupId, _orders);
        }
        /* solium-enable indentation */
    }

    /**
     * @notice Create multiple orders
     * @param _outcomes Array of associated outcomes for each order
     * @param _types Array of the type of each order. Either BID==0, or ASK==1
     * @param _attoshareAmounts Array of the number of attoShares desired for each order
     * @param _prices Array of the price in attoCash for each order. Must be within the market range (1 to numTicks-1)
     * @param _market The associated market
     * @param _tradeGroupId A Bytes32 value used when attempting to associate multiple orderbook actions with a single TX
     * @return Array of Bytes32 ids of the created orders
     */
    function publicCreateOrders(uint256[] memory _outcomes, Order.Types[] memory _types, uint256[] memory _attoshareAmounts, uint256[] memory _prices, IMarket _market, bytes32 _tradeGroupId) public nonReentrant returns (bytes32[] memory _orders) {
        require(augur.isKnownMarket(_market));
        require(_outcomes.length == _types.length);
        require(_outcomes.length == _attoshareAmounts.length);
        require(_outcomes.length == _prices.length);
        _orders = new bytes32[]( _types.length);

        IUniverse _universe = _market.getUniverse();
        for (uint256 i = 0; i <  _types.length; i++) {
            Order.Data memory _orderData = Order.create(augur, augurTrading, msg.sender, _outcomes[i], _types[i], _attoshareAmounts[i], _prices[i], _market, bytes32(0), bytes32(0));
            escrowFunds(_orderData);
            profitLoss.recordFrozenFundChange(_universe, _market, msg.sender, _outcomes[i], int256(_orderData.moneyEscrowed));
            /* solium-disable indentation */
            {
                IOrders _ordersContract = orders;
                require(_ordersContract.getAmount(Order.getOrderId(_orderData, _ordersContract)) == 0, "Createorder.publicCreateOrders: Order duplication in same block");
                _orders[i] = Order.saveOrder(_orderData, _tradeGroupId, _ordersContract);
            }
            /* solium-enable indentation */
        }

        return _orders;
    }

    //
    // Private functions
    //

    function escrowFunds(Order.Data memory _orderData) internal returns (bool) {
        if (_orderData.orderType == Order.Types.Ask) {
            return escrowFundsForAsk(_orderData);
        } else if (_orderData.orderType == Order.Types.Bid) {
            return escrowFundsForBid(_orderData);
        }
    }

    function escrowFundsForBid(Order.Data memory _orderData) private returns (bool) {
        require(_orderData.moneyEscrowed == 0, "Order.escrowFundsForBid: New order had money escrowed. This should not be possible");
        require(_orderData.sharesEscrowed == 0, "Order.escrowFundsForBid: New order had shares escrowed. This should not be possible");
        uint256 _attosharesToCover = _orderData.amount;
        uint256 _numberOfShortOutcomes = _orderData.market.getNumberOfOutcomes() - 1;

        uint256[] memory _shortOutcomes = new uint256[](_numberOfShortOutcomes);
        uint256 _indexOutcome = 0;
        for (uint256 _i = 0; _i < _numberOfShortOutcomes; _i++) {
            if (_i == _orderData.outcome) {
                _indexOutcome++;
            }
            _shortOutcomes[_i] = _indexOutcome;
            _indexOutcome++;
        }

        // Figure out how many almost-complete-sets (just missing `outcome` share) the creator has
        uint256 _attosharesHeld = _orderData.shareToken.lowestBalanceOfMarketOutcomes(_orderData.market, _shortOutcomes, _orderData.creator);

        // Take shares into escrow if they have any almost-complete-sets
        if (_attosharesHeld > 0) {
            _orderData.sharesEscrowed = SafeMathUint256.min(_attosharesHeld, _attosharesToCover);
            _attosharesToCover -= _orderData.sharesEscrowed;
            uint256[] memory _values = new uint256[](_numberOfShortOutcomes);
            for (uint256 _i = 0; _i < _numberOfShortOutcomes; _i++) {
                _values[_i] = _orderData.sharesEscrowed;
            }
            _orderData.shareToken.unsafeBatchTransferFrom(_orderData.creator, address(_orderData.augurTrading), TokenId.getTokenIds(_orderData.market, _shortOutcomes), _values);
        }

        // If not able to cover entire order with shares alone, then cover remaining with tokens
        if (_attosharesToCover > 0) {
            _orderData.moneyEscrowed = _attosharesToCover.mul(_orderData.price);
            cashTransferFrom(_orderData.creator, address(_orderData.augurTrading), _orderData.moneyEscrowed);
        }

        return true;
    }

    function escrowFundsForAsk(Order.Data memory _orderData) private returns (bool) {
        require(_orderData.moneyEscrowed == 0, "Order.escrowFundsForAsk: New order had money escrowed. This should not be possible");
        require(_orderData.sharesEscrowed == 0, "Order.escrowFundsForAsk: New order had shares escrowed. This should not be possible");
        uint256 _attosharesToCover = _orderData.amount;

        // Figure out how many shares of the outcome the creator has
        uint256 _attosharesHeld = _orderData.shareToken.balanceOfMarketOutcome(_orderData.market, _orderData.outcome, _orderData.creator);

        // Take shares in escrow if user has shares
        if (_attosharesHeld > 0) {
            _orderData.sharesEscrowed = SafeMathUint256.min(_attosharesHeld, _attosharesToCover);
            _attosharesToCover -= _orderData.sharesEscrowed;
            _orderData.shareToken.unsafeTransferFrom(_orderData.creator, address(_orderData.augurTrading), TokenId.getTokenId(_orderData.market, _orderData.outcome), _orderData.sharesEscrowed);
        }

        // If not able to cover entire order with shares alone, then cover remaining with tokens
        if (_attosharesToCover > 0) {
            _orderData.moneyEscrowed = _orderData.market.getNumTicks().sub(_orderData.price).mul(_attosharesToCover);
            cashTransferFrom(_orderData.creator, address(_orderData.augurTrading), _orderData.moneyEscrowed);
        }

        return true;
    }
}
