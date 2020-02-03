pragma solidity 0.5.15;

import 'ROOT/trading/Orders.sol';
import 'ROOT/libraries/ContractExists.sol';
import 'ROOT/libraries/token/IERC20.sol';


contract TestOrders is Orders {
    using ContractExists for address;

    address private constant FOUNDATION_REP_ADDRESS = address(0x1985365e9f78359a9B6AD760e32412f4a445E862);

    constructor() public {
        // This is to confirm we are not on foundation network
        require(!FOUNDATION_REP_ADDRESS.exists(), "TestOrders: Deploying test contract to production");
    }

    function testSaveOrder(uint256[] memory _uints, bytes32[] memory _bytes32s, Order.Types _type, IMarket _market, address _sender) public returns (bytes32 _orderId) {
        require(_uints.length == 5, "TestOrders: incorrect length for _uints array");
        require(_bytes32s.length == 4, "TestOrders: incorrect length for _bytes32s array");
        _bytes32s[3] = getOrderId(_type, _market, _uints[0], _uints[1], _sender, block.number, _uints[2], _uints[3], _uints[4]);
        return this.saveOrder(_uints, _bytes32s, _type, _market, _sender);
    }

    function testRemoveOrder(bytes32 _orderId) public returns (bool) {
        return this.removeOrder(_orderId);
    }

    function testRecordFillOrder(bytes32 _orderId, uint256 _sharesFilled, uint256 _tokensFilled, uint256 _fill) public returns (bool) {
        return this.recordFillOrder(_orderId, _sharesFilled, _tokensFilled, _fill);
    }
}
