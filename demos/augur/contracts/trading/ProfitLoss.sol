pragma solidity 0.5.15;

import 'ROOT/IAugur.sol';
import 'ROOT/trading/IOrders.sol';
import 'ROOT/libraries/Initializable.sol';
import 'ROOT/libraries/math/SafeMathUint256.sol';
import 'ROOT/libraries/math/SafeMathInt256.sol';
import 'ROOT/trading/IAugurTrading.sol';


/**
 * @title Profit Loss
 * @notice Storage of Profit Loss data.
 */
contract ProfitLoss is Initializable {
    using SafeMathUint256 for uint256;
    using SafeMathInt256 for int256;

    IAugurTrading public augurTrading;
    IOrders public orders;
    address public shareToken;
    address public createOrder;
    address public cancelOrder;
    address public fillOrder;

    struct OutcomeData {
        int256 netPosition;
        int256 avgPrice; // Cannot actually be negative. Typed for code convenience
        int256 realizedProfit;
        int256 frozenFunds;
        int256 realizedCost; // Also cannot be negative.
    }

    // User => Market => Outcome => Data
    mapping (address => mapping(address => mapping(uint256 => OutcomeData))) private profitLossData;

    function initialize(IAugur _augur, IAugurTrading _augurTrading) public beforeInitialized {
        endInitialization();
        shareToken = _augur.lookup("ShareToken");
        require(shareToken != address(0));

        augurTrading = _augurTrading;
        createOrder = _augurTrading.lookup("CreateOrder");
        cancelOrder = _augurTrading.lookup("CancelOrder");
        fillOrder = _augurTrading.lookup("FillOrder");
        orders = IOrders(_augurTrading.lookup("Orders"));
        require(createOrder != address(0));
        require(fillOrder != address(0));
        require(cancelOrder != address(0));
        require(orders != IOrders(0));
    }

    function recordFrozenFundChange(IUniverse _universe, IMarket _market, address _account, uint256 _outcome, int256 _frozenFundDelta) external returns (bool) {
        require(msg.sender == createOrder || msg.sender == cancelOrder || msg.sender == address(orders) || msg.sender == fillOrder);
        OutcomeData storage _outcomeData = profitLossData[_account][address(_market)][_outcome];
        _outcomeData.frozenFunds += _frozenFundDelta;
        augurTrading.logProfitLossChanged(_market, _account, _outcome, _outcomeData.netPosition, uint256(_outcomeData.avgPrice), _outcomeData.realizedProfit, _outcomeData.frozenFunds,  _outcomeData.realizedCost);
        return true;
    }

    function adjustTraderProfitForFees(IMarket _market, address _trader, uint256 _outcome, uint256 _fees) external returns (bool) {
        require(msg.sender == fillOrder);
        profitLossData[_trader][address(_market)][_outcome].realizedProfit -= int256(_fees);
        return true;
    }

    function recordTrade(IUniverse _universe, IMarket _market, address _longAddress, address _shortAddress, uint256 _outcome, int256 _amount, int256 _price, uint256 _numLongTokens, uint256 _numShortTokens, uint256 _numLongShares, uint256 _numShortShares) external returns (bool) {
        require(msg.sender == fillOrder);
        int256 _numTicks = int256(_market.getNumTicks());
        int256  _longFrozenTokenDelta = int256(_numLongTokens).sub(int256(_numLongShares).mul(_numTicks.sub(_price)));
        int256  _shortFrozenTokenDelta = int256(_numShortTokens).sub(int256(_numShortShares).mul(_price));
        adjustForTrader(_universe, _market, _shortAddress, _outcome, -_amount, _price, _shortFrozenTokenDelta);
        adjustForTrader(_universe, _market, _longAddress, _outcome, _amount, _price, _longFrozenTokenDelta);
        return true;
    }

    function adjustForTrader(IUniverse _universe, IMarket _market, address _address, uint256 _outcome, int256 _amount, int256 _price, int256 _frozenTokenDelta) internal returns (bool) {
        OutcomeData storage _outcomeData = profitLossData[_address][address(_market)][_outcome];
        OutcomeData memory _tmpOutcomeData = profitLossData[_address][address(_market)][_outcome];

        bool _sold = _tmpOutcomeData.netPosition < 0 &&  _amount > 0 || _tmpOutcomeData.netPosition > 0 &&  _amount < 0;
        if (_tmpOutcomeData.netPosition != 0 && _sold) {
            int256 _amountSold = _tmpOutcomeData.netPosition.abs().min(_amount.abs());
            int256 _profit = (_tmpOutcomeData.netPosition < 0 ? _tmpOutcomeData.avgPrice.sub(_price) : _price.sub(_tmpOutcomeData.avgPrice)).mul(_amountSold);
            _tmpOutcomeData.realizedProfit += _profit;
            _tmpOutcomeData.realizedCost += (_tmpOutcomeData.netPosition < 0 ? int256(_market.getNumTicks()).sub(_tmpOutcomeData.avgPrice) : _tmpOutcomeData.avgPrice).mul(_amountSold);
            _tmpOutcomeData.frozenFunds += _profit + _frozenTokenDelta;

            _outcomeData.realizedProfit = _tmpOutcomeData.realizedProfit;
            _outcomeData.realizedCost = _tmpOutcomeData.realizedCost;
            _outcomeData.frozenFunds = _tmpOutcomeData.frozenFunds;
        } else {
            _tmpOutcomeData.frozenFunds += _frozenTokenDelta;
            _outcomeData.frozenFunds = _tmpOutcomeData.frozenFunds;
        }

        int256 _newNetPosition = _tmpOutcomeData.netPosition.add(_amount);
        bool _reversed = _tmpOutcomeData.netPosition < 0 && _newNetPosition > 0 || _tmpOutcomeData.netPosition > 0 && _newNetPosition < 0;
        if (_newNetPosition == 0) {
            _tmpOutcomeData.avgPrice = 0;
            _outcomeData.avgPrice = _tmpOutcomeData.avgPrice;
        } else if (_reversed) {
            _tmpOutcomeData.avgPrice = _price;
            _outcomeData.avgPrice = _tmpOutcomeData.avgPrice;
        } else if (!_sold) {
            _tmpOutcomeData.avgPrice = _tmpOutcomeData.netPosition.abs().mul(_tmpOutcomeData.avgPrice).add(_amount.abs().mul(_price)).div(_newNetPosition.abs());
            _outcomeData.avgPrice = _tmpOutcomeData.avgPrice;
        }

        _outcomeData.netPosition = _newNetPosition;
        augurTrading.logProfitLossChanged(_market, _address, _outcome, _outcomeData.netPosition, uint256(_tmpOutcomeData.avgPrice), _tmpOutcomeData.realizedProfit, _tmpOutcomeData.frozenFunds,  _tmpOutcomeData.realizedCost);
        return true;
    }

    function recordClaim(IMarket _market, address _account, uint256[] memory _outcomeFees) public returns (bool) {
        require(msg.sender == address(augurTrading));
        uint256 _numOutcomes = _market.getNumberOfOutcomes();
        IUniverse _universe = _market.getUniverse();
        for (uint256 _outcome = 0; _outcome < _numOutcomes; _outcome++) {
            OutcomeData storage _outcomeData = profitLossData[_account][address(_market)][_outcome];
            if (_outcomeData.netPosition == 0) {
                continue;
            }
            int256 _salePrice = int256(_market.getWinningPayoutNumerator(_outcome));
            int256 _amount = _outcomeData.netPosition.abs();
            _outcomeData.realizedProfit += (_outcomeData.netPosition < 0 ? _outcomeData.avgPrice.sub(_salePrice) : _salePrice.sub(_outcomeData.avgPrice)).mul(_amount);
            _outcomeData.realizedProfit -= int256(_outcomeFees[_outcome]);
            _outcomeData.realizedCost += (_outcomeData.netPosition < 0 ? int256(_market.getNumTicks()).sub(_outcomeData.avgPrice) : _outcomeData.avgPrice).mul(_amount);
            _outcomeData.avgPrice = 0;
            _outcomeData.frozenFunds = 0;
            _outcomeData.netPosition = 0;
            augurTrading.logProfitLossChanged(_market, _account, _outcome, 0, 0, _outcomeData.realizedProfit, 0, _outcomeData.realizedCost);
        }
        return true;
    }

    function getNetPosition(address _market, address _account, uint256 _outcome) external view returns (int256) {
        return profitLossData[_account][_market][_outcome].netPosition;
    }

    function getAvgPrice(address _market, address _account, uint256 _outcome) external view returns (int256) {
        return profitLossData[_account][_market][_outcome].avgPrice;
    }

    function getRealizedProfit(address _market, address _account, uint256 _outcome) external view returns (int256) {
        return profitLossData[_account][_market][_outcome].realizedProfit;
    }

    function getFrozenFunds(address _market, address _account, uint256 _outcome) external view returns (int256) {
        return profitLossData[_account][_market][_outcome].frozenFunds;
    }

    function getRealizedCost(address _market, address _account, uint256 _outcome) external view returns (int256) {
        return profitLossData[_account][_market][_outcome].realizedCost;
    }
}
