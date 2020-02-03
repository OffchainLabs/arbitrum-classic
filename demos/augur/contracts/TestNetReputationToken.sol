pragma solidity 0.5.15;

import 'ROOT/reporting/ReputationToken.sol';
import 'ROOT/IAugur.sol';
import 'ROOT/reporting/IUniverse.sol';


contract TestNetReputationToken is ReputationToken {
    uint256 private constant DEFAULT_FAUCET_AMOUNT = 47 ether;

    constructor(IAugur _augur, IUniverse _universe, IUniverse _parentUniverse) ReputationToken(_augur, _universe, _parentUniverse) public {
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
