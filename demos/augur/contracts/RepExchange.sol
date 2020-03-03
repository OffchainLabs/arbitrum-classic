pragma solidity 0.5.15;

import "./BaseSimpleDex.sol";
import "./reporting/IV2ReputationToken.sol";
import './ITime.sol';


contract RepExchange is BaseSimpleDex {

    uint256 public lastUpdateTimestamp;
    uint256 public price;
    ITime public time;

    uint256 constant public period = 3 days; // TODO: revisit if this is an appropriate period

    uint256 constant public genesisInitialRepPriceinAttoCash = 9 * 10**18;

    function initialize(address _augurAddress, address _token) public beforeInitialized {
        initializeInternal(_augurAddress, _token);
        lastUpdateTimestamp =time.timeLowerBound();
        IV2ReputationToken _repToken = IV2ReputationToken(_token);
        IUniverse _parentUniverse = _repToken.parentUniverse();
        if (_parentUniverse == IUniverse(0)) {
            price = genesisInitialRepPriceinAttoCash;
        } else {
            price = RepExchange(address(_parentUniverse.repExchange())).price();
        }
    }

    function transferToken(address _to, uint256 _value) private {
        IV2ReputationToken(token).transfer(_to, _value);
    }

    function getTokenBalance() public returns (uint256) {
        return IV2ReputationToken(token).balanceOf(address(this));
    }

    function autoSellToken(address _recipient, uint256 _tokenAmount) external payable returns (uint256 _cashAmount) {
        IV2ReputationToken(token).trustedREPExchangeTransfer(msg.sender, address(this), _tokenAmount);
        sellToken(_recipient);
    }

    function publicMintAuto(address _to, uint256 _tokenAmount, uint256 _cashAmount) external returns (uint256 _liquidity) {
        augur.trustedCashTransfer(msg.sender, address(this), _cashAmount);
        IV2ReputationToken(token).trustedREPExchangeTransfer(msg.sender, address(this), _tokenAmount);
        publicMint(_to);
    }

    function publicBurnAuto(address _to, uint256 _tokenAmount, uint256 _cashAmount) external returns (uint256 _liquidity) {
        augur.trustedCashTransfer(msg.sender, address(this), _cashAmount);
        IV2ReputationToken(token).trustedREPExchangeTransfer(msg.sender, address(this), _tokenAmount);
        publicMint(_to);
    }

    function pokePrice() public returns (uint256) {
        update(getTokenBalance(), getCashBalance());
        return price;
    }

    function onUpdate(uint256 _blocksElapsed, uint256 _priceCumulativeIncrease) internal {
        uint256 _blockTimestamp = time.timeLowerBound();

        uint256 _price = _priceCumulativeIncrease.mul(10**18) / _blocksElapsed / 10**18;
        require(_price > 0, "Price should not be 0");

        uint256 _secondsElapsed = _blockTimestamp.sub(lastUpdateTimestamp);
        uint256 _priceAverage = _price;

        if (_secondsElapsed < period) {
            _priceAverage = price.mul(period.sub(_secondsElapsed)).add(_price.mul(_secondsElapsed)) / period;
        }

        lastUpdateTimestamp = _blockTimestamp;
        price = _priceAverage;
    }
}