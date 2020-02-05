pragma solidity 0.5.15;


import './IOrders.sol';
import '../libraries/math/SafeMathUint256.sol';
import '../libraries/math/SafeMathInt256.sol';
import './Order.sol';
import '../reporting/IMarket.sol';
import '../libraries/Initializable.sol';
import '../libraries/token/IERC20.sol';
import './IAugurTrading.sol';
import './IProfitLoss.sol';


/**
 * @title Orders
 * @notice Storage of all data associated with orders
 */
contract Orders is IOrders, Initializable {
    using Order for Order.Data;
    using SafeMathUint256 for uint256;

    struct MarketOrders {
        uint256 totalEscrowed;
        mapping(uint256 => uint256) prices;
    }

    mapping(bytes32 => Order.Data) private orders;
    mapping(address => MarketOrders) private marketOrderData;
    mapping(bytes32 => bytes32) private bestOrder;
    mapping(bytes32 => bytes32) private worstOrder;

    IAugurTrading public augurTrading;
    ICash public cash;
    address public trade;
    address public fillOrder;
    address public cancelOrder;
    address public createOrder;
    IProfitLoss public profitLoss;

    function initialize(IAugur _augur, IAugurTrading _augurTrading) public beforeInitialized {
        endInitialization();
        cash = ICash(_augur.lookup("Cash"));

        augurTrading = _augurTrading;
        createOrder = _augurTrading.lookup("CreateOrder");
        fillOrder = _augurTrading.lookup("FillOrder");
        cancelOrder = _augurTrading.lookup("CancelOrder");
        trade = _augurTrading.lookup("Trade");
        profitLoss = IProfitLoss(_augurTrading.lookup("ProfitLoss"));
        require(createOrder != address(0));
        require(fillOrder != address(0));
        require(cancelOrder != address(0));
        require(trade != address(0));
        require(profitLoss != IProfitLoss(0));
    }

    /**
     * @param _orderId The id of the order
     * @return The market associated with the order id
     */
    function getMarket(bytes32 _orderId) public view returns (IMarket) {
        return orders[_orderId].market;
    }

    /**
     * @param _orderId The id of the order
     * @return The order type (BID==0,ASK==1) associated with the order
     */
    function getOrderType(bytes32 _orderId) public view returns (Order.Types) {
        return orders[_orderId].orderType;
    }

    /**
     * @param _orderId The id of the order
     * @return The outcome associated with the order
     */
    function getOutcome(bytes32 _orderId) public view returns (uint256) {
        return orders[_orderId].outcome;
    }

    /**
     * @param _orderId The id of the order
     * @return The remaining amount of the order
     */
    function getAmount(bytes32 _orderId) public view returns (uint256) {
        return orders[_orderId].amount;
    }

    /**
     * @param _orderId The id of the order
     * @return The price of the order
     */
    function getPrice(bytes32 _orderId) public view returns (uint256) {
        return orders[_orderId].price;
    }

    /**
     * @param _orderId The id of the order
     * @return The creator of the order
     */
    function getOrderCreator(bytes32 _orderId) public view returns (address) {
        return orders[_orderId].creator;
    }

    /**
     * @param _orderId The id of the order
     * @return The remaining shares escrowed in the order
     */
    function getOrderSharesEscrowed(bytes32 _orderId) public view returns (uint256) {
        return orders[_orderId].sharesEscrowed;
    }

    /**
     * @param _orderId The id of the order
     * @return The remaining Cash tokens escrowed in the order
     */
    function getOrderMoneyEscrowed(bytes32 _orderId) public view returns (uint256) {
        return orders[_orderId].moneyEscrowed;
    }

    function getOrderDataForCancel(bytes32 _orderId) public view returns (uint256 _moneyEscrowed, uint256 _sharesEscrowed, Order.Types _type, IMarket _market, uint256 _outcome, address _creator) {
        Order.Data storage _order = orders[_orderId];
        _moneyEscrowed = _order.moneyEscrowed;
        _sharesEscrowed = _order.sharesEscrowed;
        _type = _order.orderType;
        _market = _order.market;
        _outcome = _order.outcome;
        _creator = _order.creator;
    }

    function getOrderDataForLogs(bytes32 _orderId) public view returns (Order.Types _type, address[] memory _addressData, uint256[] memory _uint256Data) {
        Order.Data storage _order = orders[_orderId];
        _addressData = new address[](2);
        _uint256Data = new uint256[](10);
        _addressData[0] = _order.creator;
        _uint256Data[0] = _order.price;
        _uint256Data[1] = _order.amount;
        _uint256Data[2] = _order.outcome;
        _uint256Data[8] = _order.sharesEscrowed;
        _uint256Data[9] = _order.moneyEscrowed;
        return (_order.orderType, _addressData, _uint256Data);
    }

    /**
     * @param _market The address of the market
     * @return The amount of Cash escrowed for all orders for the specified market
     */
    function getTotalEscrowed(IMarket _market) public view returns (uint256) {
        return marketOrderData[address(_market)].totalEscrowed;
    }

    /**
     * @param _market The address of the market
     * @param _outcome The outcome number
     * @return The price for the last completed trade for the specified market and outcome
     */
    function getLastOutcomePrice(IMarket _market, uint256 _outcome) public view returns (uint256) {
        return marketOrderData[address(_market)].prices[_outcome];
    }

    /**
     * @param _orderId The id of the order
     * @return The id (if there is one) of the next order better than the provided one
     */
    function getBetterOrderId(bytes32 _orderId) public view returns (bytes32) {
        return orders[_orderId].betterOrderId;
    }

    /**
     * @param _orderId The id of the order
     * @return The id (if there is one) of the next order worse than the provided one
     */
    function getWorseOrderId(bytes32 _orderId) public view returns (bytes32) {
        return orders[_orderId].worseOrderId;
    }

    /**
     * @param _type The type of order. Either BID==0, or ASK==1
     * @param _market The market of the order
     * @param _outcome The outcome of the order
     * @return The id (if there is one) of the best order that satisfies the given parameters
     */
    function getBestOrderId(Order.Types _type, IMarket _market, uint256 _outcome) public view returns (bytes32) {
        return bestOrder[getBestOrderWorstOrderHash(_market, _outcome, _type)];
    }

    /**
     * @param _type The type of order. Either BID==0, or ASK==1
     * @param _market The market of the order
     * @param _outcome The outcome of the order
     * @return The id (if there is one) of the worst order that satisfies the given parameters
     */
    function getWorstOrderId(Order.Types _type, IMarket _market, uint256 _outcome) public view returns (bytes32) {
        return worstOrder[getBestOrderWorstOrderHash(_market, _outcome, _type)];
    }

    /**
     * @param _type The type of order. Either BID==0, or ASK==1
     * @param _market The market of the order
     * @param _amount The amount of the order
     * @param _price The price of the order
     * @param _sender The creator of the order
     * @param _blockNumber The blockNumber which the order was created in
     * @param _outcome The outcome of the order
     * @param _moneyEscrowed The amount of Cash tokens escrowed in the order
     * @param _sharesEscrowed The outcome Share tokens escrowed in the order
     * @return The order id that satisfies the given parameters
     */
    function getOrderId(Order.Types _type, IMarket _market, uint256 _amount, uint256 _price, address _sender, uint256 _blockNumber, uint256 _outcome, uint256 _moneyEscrowed, uint256 _sharesEscrowed) public pure returns (bytes32) {
        return Order.calculateOrderId(_type, _market, _amount, _price, _sender, _blockNumber, _outcome, _moneyEscrowed, _sharesEscrowed);
    }

    function isBetterPrice(Order.Types _type, uint256 _price, bytes32 _orderId) public view returns (bool) {
        if (_type == Order.Types.Bid) {
            return (_price > orders[_orderId].price);
        } else if (_type == Order.Types.Ask) {
            return (_price < orders[_orderId].price);
        }
    }

    function isWorsePrice(Order.Types _type, uint256 _price, bytes32 _orderId) public view returns (bool) {
        if (_type == Order.Types.Bid) {
            return (_price <= orders[_orderId].price);
        } else {
            return (_price >= orders[_orderId].price);
        }
    }

    function assertIsNotBetterPrice(Order.Types _type, uint256 _price, bytes32 _betterOrderId) public view returns (bool) {
        require(!isBetterPrice(_type, _price, _betterOrderId), "Orders.assertIsNotBetterPrice: Is better price");
        return true;
    }

    function assertIsNotWorsePrice(Order.Types _type, uint256 _price, bytes32 _worseOrderId) public returns (bool) {
        require(!isWorsePrice(_type, _price, _worseOrderId), "Orders.assertIsNotWorsePrice: Is worse price");
        return true;
    }

    function insertOrderIntoList(Order.Data storage _order, bytes32 _betterOrderId, bytes32 _worseOrderId) private returns (bool) {
        bytes32 _bestOrderWorstOrderHash = getBestOrderWorstOrderHash(_order.market, _order.outcome, _order.orderType);
        bytes32 _bestOrderId = bestOrder[_bestOrderWorstOrderHash];
        bytes32 _worstOrderId = worstOrder[_bestOrderWorstOrderHash];
        (_betterOrderId, _worseOrderId) = findBoundingOrders(_order.orderType, _order.price, _bestOrderId, _worstOrderId, _betterOrderId, _worseOrderId);
        if (_order.orderType == Order.Types.Bid) {
            _bestOrderId = updateBestBidOrder(_order.id, _order.price, _order.outcome, _bestOrderWorstOrderHash, _bestOrderId);
            _worstOrderId = updateWorstBidOrder(_order.id, _order.price, _order.outcome, _bestOrderWorstOrderHash, _worstOrderId);
        } else {
            _bestOrderId = updateBestAskOrder(_order.id, _order.price, _order.outcome, _bestOrderWorstOrderHash, _bestOrderId);
            _worstOrderId = updateWorstAskOrder(_order.id, _order.price, _order.outcome, _bestOrderWorstOrderHash, _worstOrderId);
        }
        if (_bestOrderId == _order.id) {
            _betterOrderId = bytes32(0);
        }
        if (_worstOrderId == _order.id) {
            _worseOrderId = bytes32(0);
        }
        if (_betterOrderId != bytes32(0)) {
            orders[_betterOrderId].worseOrderId = _order.id;
            _order.betterOrderId = _betterOrderId;
        }
        if (_worseOrderId != bytes32(0)) {
            orders[_worseOrderId].betterOrderId = _order.id;
            _order.worseOrderId = _worseOrderId;
        }
        return true;
    }

    // _amount = _uints[0]
    // _price = _uints[1]
    // _outcome = _uints[2]
    // _moneyEscrowed = _uints[3]
    // _sharesEscrowed = _uints[4]
    // _betterOrderId = _bytes32s[0]
    // _worseOrderId = _bytes32s[1]
    // _tradeGroupId = _bytes32s[2]
    // _orderId = _bytes32s[3]
    function saveOrder(uint256[] calldata _uints, bytes32[] calldata _bytes32s, Order.Types _type, IMarket _market, address _sender) external returns (bytes32 _orderId) {
        require(msg.sender == createOrder || msg.sender == address(this));
        require(_uints[2] < _market.getNumberOfOutcomes(), "Orders.saveOrder: Outcome not in market range");
        _orderId = _bytes32s[3];
        Order.Data storage _order = orders[_orderId];
        _order.market = _market;
        _order.id = _orderId;
        _order.orderType = _type;
        _order.outcome = _uints[2];
        _order.price = _uints[1];
        _order.amount = _uints[0];
        _order.creator = _sender;
        _order.moneyEscrowed = _uints[3];
        marketOrderData[address(_market)].totalEscrowed += _uints[3];
        _order.sharesEscrowed = _uints[4];
        insertOrderIntoList(_order, _bytes32s[0], _bytes32s[1]);
        augurTrading.logOrderCreated(_order.market.getUniverse(), _orderId, _bytes32s[2]);
        return _orderId;
    }

    function removeOrder(bytes32 _orderId) external returns (bool) {
        require(msg.sender == cancelOrder || msg.sender == address(this));
        removeOrderFromList(_orderId);
        Order.Data storage _order = orders[_orderId];
        marketOrderData[address(_order.market)].totalEscrowed -= _order.moneyEscrowed;
        delete orders[_orderId];
        return true;
    }

    function recordFillOrder(bytes32 _orderId, uint256 _sharesFilled, uint256 _tokensFilled, uint256 _fill) external returns (bool) {
        require(msg.sender == fillOrder || msg.sender == address(this));
        Order.Data storage _order = orders[_orderId];
        require(_order.outcome < _order.market.getNumberOfOutcomes(), "Orders.recordFillOrder: Outcome is not in market range");
        require(_orderId != bytes32(0), "Orders.recordFillOrder: orderId is 0x0");
        require(_sharesFilled <= _order.sharesEscrowed, "Orders.recordFillOrder: shares filled higher than order amount");
        require(_tokensFilled <= _order.moneyEscrowed, "Orders.recordFillOrder: tokens filled higher than order amount");
        require(_order.price <= _order.market.getNumTicks(), "Orders.recordFillOrder: Price outside of market range");
        require(_fill <= _order.amount, "Orders.recordFillOrder: Fill higher than order amount");
        _order.amount -= _fill;
        _order.moneyEscrowed -= _tokensFilled;
        marketOrderData[address(_order.market)].totalEscrowed -= _tokensFilled;
        _order.sharesEscrowed -= _sharesFilled;
        if (_order.amount == 0) {
            require(_order.moneyEscrowed == 0, "Orders.recordFillOrder: Money left in filled order");
            require(_order.sharesEscrowed == 0, "Orders.recordFillOrder: Shares left in filled order");
            removeOrderFromList(_orderId);
            _order.price = 0;
            _order.creator = address(0);
            _order.betterOrderId = bytes32(0);
            _order.worseOrderId = bytes32(0);
        }
        return true;
    }

    function setPrice(IMarket _market, uint256 _outcome, uint256 _price) external returns (bool) {
        require(msg.sender == trade);
        marketOrderData[address(_market)].prices[_outcome] = _price;
        return true;
    }

    function removeOrderFromList(bytes32 _orderId) private returns (bool) {
        Order.Data storage _order = orders[_orderId];
        bytes32 _betterOrderId = _order.betterOrderId;
        bytes32 _worseOrderId = _order.worseOrderId;
        bytes32 _bestOrderWorstOrderHash = getBestOrderWorstOrderHash(_order.market, _order.outcome, _order.orderType);
        if (bestOrder[_bestOrderWorstOrderHash] == _orderId) {
            bestOrder[_bestOrderWorstOrderHash] = _worseOrderId;
        }
        if (worstOrder[_bestOrderWorstOrderHash] == _orderId) {
            worstOrder[_bestOrderWorstOrderHash] = _betterOrderId;
        }
        if (_betterOrderId != bytes32(0)) {
            orders[_betterOrderId].worseOrderId = _worseOrderId;
        }
        if (_worseOrderId != bytes32(0)) {
            orders[_worseOrderId].betterOrderId = _betterOrderId;
        }
        _order.betterOrderId = bytes32(0);
        _order.worseOrderId = bytes32(0);
        return true;
    }

    /**
     * @dev If best bid is not set or price higher than best bid price, this order is the new best bid.
     */
    function updateBestBidOrder(bytes32 _orderId, uint256 _price, uint256 _outcome, bytes32 _bestOrderWorstOrderHash, bytes32 _bestBidOrderId) private returns (bytes32) {
        if (_bestBidOrderId == bytes32(0) || _price > orders[_bestBidOrderId].price) {
            bestOrder[_bestOrderWorstOrderHash] = _orderId;
            return _orderId;
        } else {
            return _bestBidOrderId;
        }
    }

    /**
     * @dev If worst bid is not set or price lower than worst bid price, this order is the new worst bid.
     */
    function updateWorstBidOrder(bytes32 _orderId, uint256 _price, uint256 _outcome, bytes32 _bestOrderWorstOrderHash, bytes32 _worstBidOrderId) private returns (bytes32) {
        if (_worstBidOrderId == bytes32(0) || _price <= orders[_worstBidOrderId].price) {
            worstOrder[_bestOrderWorstOrderHash] = _orderId;
            return _orderId;
        } else {
            return _worstBidOrderId;
        }
    }

    /**
     * @dev If best ask is not set or price lower than best ask price, this order is the new best ask.
     */
    function updateBestAskOrder(bytes32 _orderId, uint256 _price, uint256 _outcome, bytes32 _bestOrderWorstOrderHash, bytes32 _bestAskOrderId) private returns (bytes32) {
        if (_bestAskOrderId == bytes32(0) || _price < orders[_bestAskOrderId].price) {
            bestOrder[_bestOrderWorstOrderHash] = _orderId;
            return _orderId;
        } else {
            return _bestAskOrderId;
        }
    }

    /**
     * @dev If worst ask is not set or price higher than worst ask price, this order is the new worst ask.
     */
    function updateWorstAskOrder(bytes32 _orderId, uint256 _price, uint256 _outcome, bytes32 _bestOrderWorstOrderHash, bytes32 _worstAskOrderId) private returns (bytes32) {
        if (_worstAskOrderId == bytes32(0) || _price >= orders[_worstAskOrderId].price) {
            worstOrder[_bestOrderWorstOrderHash] = _orderId;
            return _orderId;
        } else {
            return _worstAskOrderId;
        }
    }

    function getBestOrderWorstOrderHash(IMarket _market, uint256 _outcome, Order.Types _type) private pure returns (bytes32) {
        return sha256(abi.encodePacked(_market, _outcome, _type));
    }

    function ascendOrderList(Order.Types _type, uint256 _price, bytes32 _lowestOrderId) public view returns (bytes32 _betterOrderId, bytes32 _worseOrderId) {
        _worseOrderId = _lowestOrderId;
        bool _isWorstPrice;
        if (_type == Order.Types.Bid) {
            _isWorstPrice = _price <= getPrice(_worseOrderId);
        } else if (_type == Order.Types.Ask) {
            _isWorstPrice = _price >= getPrice(_worseOrderId);
        }
        if (_isWorstPrice) {
            return (_worseOrderId, getWorseOrderId(_worseOrderId));
        }
        bool _isBetterPrice = isBetterPrice(_type, _price, _worseOrderId);
        while (_isBetterPrice && getBetterOrderId(_worseOrderId) != 0 && _price != getPrice(getBetterOrderId(_worseOrderId))) {
            _betterOrderId = getBetterOrderId(_worseOrderId);
            _isBetterPrice = isBetterPrice(_type, _price, _betterOrderId);
            if (_isBetterPrice) {
                _worseOrderId = getBetterOrderId(_worseOrderId);
            }
        }
        _betterOrderId = getBetterOrderId(_worseOrderId);
        return (_betterOrderId, _worseOrderId);
    }

    function descendOrderList(Order.Types _type, uint256 _price, bytes32 _highestOrderId) public view returns (bytes32 _betterOrderId, bytes32 _worseOrderId) {
        _betterOrderId = _highestOrderId;
        bool _isBestPrice;
        if (_type == Order.Types.Bid) {
            _isBestPrice = _price > getPrice(_betterOrderId);
        } else if (_type == Order.Types.Ask) {
            _isBestPrice = _price < getPrice(_betterOrderId);
        }
        if (_isBestPrice) {
            return (0, _betterOrderId);
        }
        bool _isWorsePrice = isWorsePrice(_type, _price, _betterOrderId);
        while (_isWorsePrice && getWorseOrderId(_betterOrderId) != 0) {
            _worseOrderId = getWorseOrderId(_betterOrderId);
            _isWorsePrice = isWorsePrice(_type, _price, _worseOrderId);
            if (_isWorsePrice || _price == getPrice(getWorseOrderId(_betterOrderId))) {
                _betterOrderId = getWorseOrderId(_betterOrderId);
            }
        }
        _worseOrderId = getWorseOrderId(_betterOrderId);
        return (_betterOrderId, _worseOrderId);
    }

    function findBoundingOrders(Order.Types _type, uint256 _price, bytes32 _bestOrderId, bytes32 _worstOrderId, bytes32 _betterOrderId, bytes32 _worseOrderId) public returns (bytes32 betterOrderId, bytes32 worseOrderId) {
        if (_bestOrderId == _worstOrderId) {
            if (_bestOrderId == bytes32(0)) {
                return (bytes32(0), bytes32(0));
            } else if (isBetterPrice(_type, _price, _bestOrderId)) {
                return (bytes32(0), _bestOrderId);
            } else {
                return (_bestOrderId, bytes32(0));
            }
        }
        if (_betterOrderId != bytes32(0)) {
            if (getPrice(_betterOrderId) == 0) {
                _betterOrderId = bytes32(0);
            } else {
                assertIsNotBetterPrice(_type, _price, _betterOrderId);
            }
        }
        if (_worseOrderId != bytes32(0)) {
            if (getPrice(_worseOrderId) == 0) {
                _worseOrderId = bytes32(0);
            } else {
                assertIsNotWorsePrice(_type, _price, _worseOrderId);
            }
        }
        if (_betterOrderId == bytes32(0) && _worseOrderId == bytes32(0)) {
            return (descendOrderList(_type, _price, _bestOrderId));
        } else if (_betterOrderId == bytes32(0)) {
            return (ascendOrderList(_type, _price, _worseOrderId));
        } else if (_worseOrderId == bytes32(0)) {
            return (descendOrderList(_type, _price, _betterOrderId));
        }
        if (getWorseOrderId(_betterOrderId) != _worseOrderId) {
            return (descendOrderList(_type, _price, _betterOrderId));
        } else if (getBetterOrderId(_worseOrderId) != _betterOrderId) {
            // Coverage: This condition is likely unreachable or at least seems to be. Rather than remove it I'm keeping it for now just to be paranoid
            return (ascendOrderList(_type, _price, _worseOrderId));
        }
        return (_betterOrderId, _worseOrderId);
    }
}
