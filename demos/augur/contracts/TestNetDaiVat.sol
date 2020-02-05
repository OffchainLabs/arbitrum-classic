pragma solidity 0.5.15;

import './external/IDaiVat.sol';


contract TestNetDaiVat is IDaiVat {
    function cage() external {
        live = 0;
    }

    function add(uint x, uint y) internal pure returns (uint z) {
        require((z = x + y) >= x);
    }

    function sub(uint x, uint y) internal pure returns (uint z) {
        require((z = x - y) <= x);
    }

    function hope(address usr) public { can[msg.sender][usr] = 1; }
    function nope(address usr) public { can[msg.sender][usr] = 0; }
    function wish(address bit, address usr) internal view returns (bool) {
        return bit == usr || can[bit][usr] == 1;
    }

    function suck(address, address v, uint rad) public {
        dai[v] = add(dai[v], rad);
    }

    function move(address src, address dst, uint256 rad) public {
        require(wish(src, msg.sender));
        dai[src] = sub(dai[src], rad);
        dai[dst] = add(dai[dst], rad);
    }

    function frob(bytes32 i, address u, address v, address w, int dink, int dart) external {
        // Just here for interface fulfilment
        return;
    }

    function faucet(address _target, uint256 _amount) public {
        dai[_target] = add(dai[_target], _amount);
    }
}