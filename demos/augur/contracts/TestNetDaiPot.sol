pragma solidity 0.5.15;

import 'ROOT/external/IDaiVat.sol';
import 'ROOT/external/IDaiPot.sol';
import 'ROOT/ITime.sol';


contract TestNetDaiPot is IDaiPot {
    uint256 public Pie;  // total Savings Dai

    IDaiVat public vat;  // CDP engine
    uint256 public rho;  // Time of last drip

    ITime public time;

    uint constant ONE = 10 ** 27;

    constructor(address vat_, ITime _time) public {
        vat = IDaiVat(vat_);
        dsr = ONE;
        chi = ONE;
        time = _time;
        rho = time.getTimestamp();
    }

    function rpow(uint x, uint n, uint base) internal pure returns (uint z) {
        assembly {
            switch x case 0 {switch n case 0 {z := base} default {z := 0}}
            default {
                switch mod(n, 2) case 0 { z := base } default { z := x }
                let half := div(base, 2)  // for rounding.
                for { n := div(n, 2) } n { n := div(n,2) } {
                    let xx := mul(x, x)
                    if iszero(eq(div(xx, x), x)) { revert(0,0) }
                    let xxRound := add(xx, half)
                    if lt(xxRound, xx) { revert(0,0) }
                    x := div(xxRound, base)
                    if mod(n,2) {
                        let zx := mul(z, x)
                        if and(iszero(iszero(x)), iszero(eq(div(zx, x), z))) { revert(0,0) }
                        let zxRound := add(zx, half)
                        if lt(zxRound, zx) { revert(0,0) }
                        z := div(zxRound, base)
                    }
                }
            }
        }
    }

    function rmul(uint x, uint y) internal pure returns (uint z) {
        z = Mul(x, y) / ONE;
    }

    function Add(uint x, uint y) internal pure returns (uint z) {
        require((z = x + y) >= x);
    }

    function Sub(uint x, uint y) internal pure returns (uint z) {
        require((z = x - y) <= x);
    }

    function Mul(uint x, uint y) internal pure returns (uint z) {
        require(y == 0 || (z = x * y) / y == x);
    }

    function drip() public returns (uint256) {
        uint256 _now = time.getTimestamp();
        require(_now >= rho);
        uint chi_ = Sub(rmul(rpow(dsr, _now - rho, ONE), chi), chi);
        chi = Add(chi, chi_);
        rho  = _now;
        vat.suck(address(0), address(this), Mul(Pie, chi_));
        return chi;
    }

    function setDSR(uint256 _dsr) public returns (bool) {
        dsr = _dsr;
        return true;
    }

    // --- Savings Dai Management ---
    function join(uint wad) public {
        pie[msg.sender] = Add(pie[msg.sender], wad);
        Pie = Add(Pie, wad);
        vat.move(msg.sender, address(this), Mul(chi, wad));
    }

    function exit(uint wad) public {
        pie[msg.sender] = Sub(pie[msg.sender], wad);
        Pie = Sub(Pie, wad);
        vat.move(address(this), msg.sender, Mul(chi, wad));
    }
}