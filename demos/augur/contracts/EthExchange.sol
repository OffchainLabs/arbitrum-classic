pragma solidity 0.5.15;

import "./BaseSimpleDex.sol";


contract EthExchange is BaseSimpleDex {

    function initialize(address _augurAddress) public beforeInitialized {
        initializeInternal(_augurAddress, address(0));
    }

    function transferToken(address _to, uint256 _value) private {
        address payable _payable = address(uint160(_to));
        _payable.transfer(_value);
    }

    function getTokenBalance() public returns (uint256) {
        return address(this).balance;
    }

    function autoSellToken(address _recipient, uint256 _tokenAmount) external payable returns (uint256 _cashAmount) {
        sellToken(_recipient);
    }

    function publicMintAuto(address _to, uint256 _cashAmount) external payable returns (uint256 _liquidity) {
        augur.trustedCashTransfer(msg.sender, address(this), _cashAmount);
        publicMint(_to);
    }

    function onUpdate(uint256 _blocksElapsed, uint256 _priceCumulativeIncrease) internal {}

    function () external payable {
    }
}