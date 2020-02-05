pragma solidity 0.5.15;

import "./IAugur.sol";
import "./ISimpleDex.sol";
import "./CashSender.sol";
import "./libraries/token/IERC20.sol";
import "./libraries/token/VariableSupplyToken.sol";
import "./libraries/Initializable.sol";
import "./libraries/ReentrancyGuard.sol";
import "./libraries/math/SafeMathUint256.sol";


contract BaseSimpleDex is Initializable, ReentrancyGuard, VariableSupplyToken, CashSender {
    using SafeMathUint256  for uint256;

    string constant public name = "Simple Decentralized Exchange";
    string constant public symbol = "SDE";

    IAugur public augur;
    address public token;

    uint256 public tokenReserve;
    uint256 public cashReserve;
    uint256 public blockNumberLast;
    uint256 public tokenPriceCumulativeLast;

    event Mint(address indexed sender, uint256 tokenAmount, uint256 cashAmount);
    event Burn(address indexed sender, uint256 tokenAmount, uint256 cashAmount, address indexed to);
    event Swap(address indexed sender, address indexed tokenIn, uint256 amountIn, uint256 amountOut, address indexed to);
    event Sync(uint256 tokenReserve, uint256 cashReserve);

    function initializeInternal(address _augurAddress, address _token) internal {
        endInitialization();
        IAugur _augur = IAugur(_augurAddress);
        augur = _augur;
        token = _token;

        initializeCashSender(_augur.lookup("DaiVat"), _augur.lookup("Cash"));
        require(cash != IERC20(0));
    }

    function transferToken(address _to, uint256 _value) private;

    function getTokenBalance() public returns (uint256);

    function transferCash(address _to, uint _value) private {
        cashTransfer(_to, _value);
    }

    function getCashBalance() public view returns (uint256) {
        return cashBalance(address(this));
    }

    function update(uint256 _tokenBalance, uint256 _cashBalance) internal {
        uint256 _blockNumber = block.number;
        uint256 _blocksElapsed = _blockNumber - blockNumberLast;
        uint256 _cashReserve = cashReserve;
        uint256 _tokenReserve = tokenReserve;
        if (_tokenReserve != 0 && _cashReserve != 0 && _blocksElapsed > 0) {
            // cannot reasonably overflow unless the supply of Cash, REP, or ETH became several OOM larger
            uint256 _priceCumulativeIncrease = _cashReserve.mul(10**18).mul(_blocksElapsed).div(_tokenReserve);
            onUpdate(_blocksElapsed, _priceCumulativeIncrease);
            tokenPriceCumulativeLast += _priceCumulativeIncrease;
        }
        tokenReserve = _tokenBalance;
        cashReserve = _cashBalance;
        blockNumberLast = _blockNumber;
        emit Sync(_tokenBalance, _cashBalance);
    }

    function onUpdate(uint256 _blocksElapsed, uint256 _priceCumulativeIncrease) internal;

    function publicMint(address _to) public nonReentrant returns (uint256 _liquidity) {
        uint256 _tokenBalance = getTokenBalance();
        uint256 _cashBalance = getCashBalance();
        uint256 _tokenReserve = tokenReserve;
        uint256 _cashReserve = cashReserve;
        uint256 _totalSupply = totalSupply;
        uint256 _tokenAmount = _tokenBalance.sub(_tokenReserve);
        uint256 _cashAmount = _cashBalance.sub(_cashReserve);

        _liquidity = _totalSupply == 0 ?
            _tokenAmount.mul(_cashAmount).sqrt() :
            (_tokenAmount.mul(_totalSupply) / _tokenReserve).min(_cashAmount.mul(_totalSupply) / _cashReserve);
        require(_liquidity > 0, "Insufficient liquidity");
        mint(_to, _liquidity);

        update(_tokenBalance, _cashBalance);
        emit Mint(msg.sender, _tokenAmount, _cashAmount);
    }

    function publicBurnAuto(address _to, uint256 _amount) external returns (uint256 _tokenAmount, uint256 _cashAmount) {
        transferFrom(msg.sender, address(this), _amount);
        publicBurn(_to);
    }

    function publicBurn(address _to) public nonReentrant returns (uint256 _tokenAmount, uint256 _cashAmount) {
        uint256 _liquidity = balances[address(this)];
        uint256 _totalSupply = totalSupply;

        _tokenAmount = _liquidity.mul(tokenReserve) / _totalSupply;
        _cashAmount = _liquidity.mul(cashReserve) / _totalSupply;
        require(_tokenAmount > 0 && _cashAmount > 0, "Insufficient liquidity");
        burn(address(this), _liquidity);
        transferToken(_to, _tokenAmount);
        transferCash(_to, _cashAmount);


        update(getTokenBalance(), getCashBalance());
        emit Burn(msg.sender, _tokenAmount, _cashAmount, _to);
    }

    function getTokenSaleProceeds(uint256 _tokenAmount) public view returns (uint256) {
        return getSaleProceeds(_tokenAmount, tokenReserve, cashReserve);
    }

    function getCashSaleProceeds(uint256 _cashAmount) public view returns (uint256) {
        return getSaleProceeds(_cashAmount, cashReserve, tokenReserve);
    }

    function getSaleProceeds(uint256 _inputAmount, uint256 _inputReserve, uint256 _outputReserve) public pure returns (uint256) {
        require(_inputReserve > 0 && _outputReserve > 0, "inputReserve & outputReserve must be > 0");
        uint256 _amountInputWithFee = _inputAmount.mul(997);
        uint256 _numerator = _amountInputWithFee.mul(_outputReserve);
        uint256 _denominator = _inputReserve.mul(1000).add(_amountInputWithFee);
        return _numerator / _denominator;
    }

    function getTokenPurchaseCost(uint256 _tokenAmount) public view returns (uint256) {
        return getPurchaseCost(_tokenAmount, cashReserve, tokenReserve);
    }

    function getCashPurchaseCost(uint256 _cashAmount) public view returns (uint256) {
        return getPurchaseCost(_cashAmount, tokenReserve, cashReserve);
    }

    function getPurchaseCost(uint256 _outputAmount, uint256 _inputReserve, uint256 _outputReserve) public pure returns (uint256) {
        require(_inputReserve > 0 && _outputReserve > 0, "inputReserve & outputReserve must be > 0");
        uint256 _numerator = _outputAmount.mul(_inputReserve).mul(1000);
        uint256 _denominator = _outputReserve.sub(_outputAmount).mul(997);
        return _numerator / _denominator;
    }

    function autoSellToken(address _recipient, uint256 _tokenAmount) external payable returns (uint256 _cashAmount);

    function sellToken(address _recipient) public nonReentrant returns (uint _cashAmount) {
        uint256 _tokenBalance = getTokenBalance();
        uint256 _tokenAmount = _tokenBalance.sub(tokenReserve);

        _cashAmount = getTokenSaleProceeds(_tokenAmount);
        require(_cashAmount > 0, "cashAmount must be > 0");
        transferCash(_recipient, _cashAmount);

        update(_tokenBalance, getCashBalance());
        emit Swap(msg.sender, token, _tokenAmount, _cashAmount, _recipient);
    }

    function autoBuyToken(address _recipient, uint256 _cashAmount) external returns (uint256 _tokenAmount) {
        augur.trustedCashTransfer(msg.sender, address(this), _cashAmount);
        buyToken(_recipient);
    }

    function buyToken(address _recipient) public nonReentrant returns (uint256 _tokenAmount) {
        uint256 _cashBalance = getCashBalance();
        uint256 _cashAmount = _cashBalance.sub(cashReserve);

        _tokenAmount = getCashSaleProceeds(_cashAmount);
        require(_tokenAmount > 0, "tokenAmount must be > 0");
        transferToken(_recipient, _tokenAmount);

        update(getTokenBalance(), _cashBalance);
        emit Swap(msg.sender, address(cash), _cashAmount, _tokenAmount, _recipient);
    }

    function skim(address to) external nonReentrant {
        transferToken(to, getTokenBalance().sub(tokenReserve));
        transferCash(to, getCashBalance().sub(cashReserve));
    }

    function sync() external nonReentrant {
        update(getTokenBalance(), getCashBalance());
    }

    function onTokenTransfer(address _from, address _to, uint256 _value) internal {
    }

    function onMint(address _target, uint256 _amount) internal {
    }

    function onBurn(address _target, uint256 _amount) internal {
    }
}