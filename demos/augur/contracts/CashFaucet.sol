pragma solidity 0.5.15;

import './IAugur.sol';
import './ICashFaucet.sol';
import './libraries/math/SafeMathUint256.sol';
import './libraries/token/IERC20.sol';
import './external/IDaiJoin.sol';
import './external/IDaiVat.sol';
import './external/IDaiFaucet.sol';
import './CashFaucetProxy.sol';


/**
 * @title Cash Faucet
 * @dev Faucet contract for Tesnet CASH (Dai)
 */
contract CashFaucet is ICashFaucet {
    IDaiVat public vat;
    IDaiJoin public daiJoin;
    IERC20 public col;
    IDaiJoin public colJoin;
    bytes32 public colIlk;
    IDaiFaucet public mcdFaucet;
    IERC20 public dai;

    constructor(IAugur _augur) public {
        vat = IDaiVat(_augur.lookup("DaiVat"));
        mcdFaucet = IDaiFaucet(_augur.lookup("MCDFaucet"));
        col = IERC20(_augur.lookup("MCDCol"));
        colJoin = IDaiJoin(_augur.lookup("MCDColJoin"));
        daiJoin = IDaiJoin(_augur.lookup("DaiJoin"));
        dai = IERC20(_augur.lookup("Cash"));

        colIlk = bytes32("REP-A");

        col.approve(address(colJoin), 2**256 - 1);
        dai.approve(address(daiJoin), 2**256 - 1);
        col.approve(address(vat), 2**256 - 1);
        dai.approve(address(vat), 2**256 - 1);
        vat.hope(address(daiJoin));
    }

    function faucet(uint256) public returns (bool) {
        // generate collateral by creating a proxy that will faucet for us. We do this because the MCD faucet only allows one use per address
        new CashFaucetProxy(mcdFaucet, col);

        // get balance of collateral
        uint256 balance = col.balanceOf(address(this));

        // Deposit collateral
        colJoin.join(address(this), balance);

        // Open a CDP and issue max DAI
        (uint256 art, uint256 rate, uint256 spot, uint256 line, uint256 dust) = vat.ilks(colIlk);
        uint256 daiReceived = spot * balance / 10**27 - 10**18;
        vat.frob(colIlk, address(this), address(this),address(this), int256(balance), int256(daiReceived));

        // Mint DAI for the sender
        daiJoin.exit(msg.sender, daiReceived);

        return true;
    }
}