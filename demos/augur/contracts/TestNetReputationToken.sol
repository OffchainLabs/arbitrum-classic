pragma solidity 0.5.15;

import './reporting/ReputationToken.sol';
import './IAugur.sol';
import './reporting/IUniverse.sol';


contract TestNetReputationToken is ReputationToken {
    uint256 private constant DEFAULT_FAUCET_AMOUNT = 47 ether;

    function initialize(IAugur _augur, IUniverse _universe, IUniverse _parentUniverse) public {
        initializeRepToken(_augur, _universe, _parentUniverse);
    }

    function faucet(uint256 _amount) public returns (bool) {
        if (_amount == 0) {
            _amount = DEFAULT_FAUCET_AMOUNT;
        }
        require(_amount < 2 ** 128, "TestNetReputationToken.faucet: amount is greater than 2**128");
        mint(msg.sender, _amount);
        return true;
    }
}
