pragma solidity 0.5.15;
pragma experimental ABIEncoderV2;

import 'ROOT/IAugurCreationDataGetter.sol';
import "ROOT/libraries/math/SafeMathUint256.sol";
import "ROOT/libraries/ContractExists.sol";
import "ROOT/libraries/token/IERC20.sol";
import "ROOT/external/IExchange.sol";
import "ROOT/trading/IFillOrder.sol";
import "ROOT/ICash.sol";
import "ROOT/trading/Order.sol";
import "ROOT/trading/IZeroXTrade.sol";
import "ROOT/trading/IAugurTrading.sol";
import 'ROOT/libraries/Initializable.sol';
import "ROOT/IAugur.sol";
import 'ROOT/libraries/token/IERC1155.sol';
import 'ROOT/libraries/LibBytes.sol';
import 'ROOT/CashSender.sol';
import 'ROOT/ISimpleDex.sol';


contract ZeroXTrade is Initializable, IZeroXTrade, IERC1155, CashSender {
    using SafeMathUint256 for uint256;
    using LibBytes for bytes;

    bool transferFromAllowed = false;

    uint256 constant public TRADE_INTERVAL_VALUE = 10 ** 19; // Trade value of 10 DAI
    uint256 constant public MIN_TRADE_INTERVAL = 10**14; // We ignore "dust" portions of the min interval and for huge scalars have a larger min value

    // ERC20Token(address)
    bytes4 constant private ERC20_PROXY_ID = 0xf47261b0;

    // ERC1155Assets(address,uint256[],uint256[],bytes)
    bytes4 constant private MULTI_ASSET_PROXY_ID = 0x94cfcdd7;

    // ERC1155Assets(address,uint256[],uint256[],bytes)
    bytes4 constant private ERC1155_PROXY_ID = 0xa7cb5fb7;

    // EIP191 header for EIP712 prefix
    string constant internal EIP191_HEADER = "\x19\x01";

    // EIP712 Domain Name value
    string constant internal EIP712_DOMAIN_NAME = "0x Protocol";

    // EIP712 Domain Version value
    string constant internal EIP712_DOMAIN_VERSION = "2";

    // EIP1271 Order With Hash Selector
    bytes4 constant public EIP1271_ORDER_WITH_HASH_SELECTOR = 0x3efe50c8;

    // Hash of the EIP712 Domain Separator Schema
    bytes32 constant internal EIP712_DOMAIN_SEPARATOR_SCHEMA_HASH = keccak256(
        abi.encodePacked(
        "EIP712Domain(",
        "string name,",
        "string version,",
        "address verifyingContract",
        ")"
    ));

    bytes32 constant internal EIP712_ORDER_SCHEMA_HASH = keccak256(
        abi.encodePacked(
        "Order(",
        "address makerAddress,",
        "address takerAddress,",
        "address feeRecipientAddress,",
        "address senderAddress,",
        "uint256 makerAssetAmount,",
        "uint256 takerAssetAmount,",
        "uint256 makerFee,",
        "uint256 takerFee,",
        "uint256 expirationTimeSeconds,",
        "uint256 salt,",
        "bytes makerAssetData,",
        "bytes takerAssetData",
        "bytes makerFeeAssetData,",
        "bytes takerFeeAssetData",
        ")"
    ));

    // Hash of the EIP712 Domain Separator data
    // solhint-disable-next-line var-name-mixedcase
    bytes32 public EIP712_DOMAIN_HASH;

    IAugurTrading public augurTrading;
    IFillOrder public fillOrder;
    ICash public cash;
    IShareToken public shareToken;
    IExchange public exchange;
    ISimpleDex public ethExchange;

    function initialize(IAugur _augur, IAugurTrading _augurTrading) public beforeInitialized {
        endInitialization();
        augurTrading = _augurTrading;
        cash = ICash(_augur.lookup("Cash"));
        require(cash != ICash(0));
        shareToken = IShareToken(_augur.lookup("ShareToken"));
        require(shareToken != IShareToken(0));
        exchange = IExchange(_augurTrading.lookup("ZeroXExchange"));
        require(exchange != IExchange(0));
        fillOrder = IFillOrder(_augurTrading.lookup("FillOrder"));
        require(fillOrder != IFillOrder(0));
        ethExchange = ISimpleDex(_augur.lookup("EthExchange"));
        require(ethExchange != ISimpleDex(0));

        initializeCashSender(_augur.lookup("DaiVat"), address(cash));

        EIP712_DOMAIN_HASH = keccak256(
            abi.encodePacked(
                EIP712_DOMAIN_SEPARATOR_SCHEMA_HASH,
                keccak256(bytes(EIP712_DOMAIN_NAME)),
                keccak256(bytes(EIP712_DOMAIN_VERSION)),
                uint256(address(this))
            )
        );
    }

    // ERC1155 Implementation
    /// @notice Transfers value amount of an _id from the _from address to the _to address specified.
    /// @dev MUST emit TransferSingle event on success.
    /// @param from    Source address
    /// @param to      Target address
    /// @param id      ID of the token type
    /// @param value   Transfer amount
    /// @param data    Additional data with no specified format, sent in call to `_to`
    function safeTransferFrom(address from, address to, uint256 id, uint256 value, bytes calldata data) external {
        require(transferFromAllowed);
        emit TransferSingle(msg.sender, from, to, id, value);
    }

    /// @notice Send multiple types of Tokens from a 3rd party in one transfer (with safety call).
    /// @dev MUST emit TransferBatch event on success.
    /// @param from    Source addresses
    /// @param to      Target addresses
    /// @param ids     IDs of each token type
    /// @param values  Transfer amounts per token type
    /// @param data    Additional data with no specified format, sent in call to `_to`
    function safeBatchTransferFrom(address from, address to, uint256[] calldata ids, uint256[] calldata values, bytes calldata data) external {
        require(transferFromAllowed);
        emit TransferBatch(msg.sender, from, to, ids, values);
    }

    /// @notice Get the balance of an account's Tokens.
    /// @param owner  The address of the token holder
    /// @param id     ID of the Token
    /// @return       The _owner's balance of the Token type requested
    function balanceOf(address owner, uint256 id) external view returns (uint256) {
        (address _market, uint256 _price, uint8 _outcome, uint8 _type) = unpackTokenId(id);
        // NOTE: An invalid order type will cause a failure here. That is malformed input so we don't mind reverting in such a case
        Order.Types _orderType = Order.Types(_type);
        if (_orderType == Order.Types.Ask) {
            return askBalance(owner, IMarket(_market), _outcome, _price);
        } else if (_orderType == Order.Types.Bid) {
            return bidBalance(owner, IMarket(_market), _outcome, _price);
        }
    }

    function totalSupply(uint256 id) external view returns (uint256) {
        return 0;
    }

    function bidBalance(address _owner, IMarket _market, uint8 _outcome, uint256 _price) public view returns (uint256) {
        uint256 _numberOfOutcomes = _market.getNumberOfOutcomes();
        // Figure out how many almost-complete-sets (just missing `outcome` share) the creator has
        uint256[] memory _shortOutcomes = new uint256[](_numberOfOutcomes - 1);
        uint256 _indexOutcome = 0;
        for (uint256 _i = 0; _i < _numberOfOutcomes - 1; _i++) {
            if (_i == _outcome) {
                _indexOutcome++;
            }
            _shortOutcomes[_i] = _indexOutcome;
            _indexOutcome++;
        }

        uint256 _attoSharesOwned = shareToken.lowestBalanceOfMarketOutcomes(_market, _shortOutcomes, _owner);

        uint256 _availableCash = cashAvailableForTransferFrom(_owner, address(fillOrder));
        uint256 _attoSharesPurchasable = _availableCash.div(_price);

        return _attoSharesOwned.add(_attoSharesPurchasable);
    }

    function askBalance(address _owner, IMarket _market, uint8 _outcome, uint256 _price) public view returns (uint256) {
        uint256 _attoSharesOwned = shareToken.balanceOfMarketOutcome(_market, _outcome, _owner);
        uint256 _availableCash = cashAvailableForTransferFrom(_owner, address(fillOrder));
        uint256 _attoSharesPurchasable = _availableCash.div(_market.getNumTicks().sub(_price));

        return _attoSharesOwned.add(_attoSharesPurchasable);
    }

    /// @notice Get the balance of multiple account/token pairs
    /// @param owners The addresses of the token holders
    /// @param ids    ID of the Tokens
    /// @return        The _owner's balance of the Token types requested
    function balanceOfBatch(address[] calldata owners, uint256[] calldata ids) external view returns (uint256[] memory balances_) {
        balances_ = new uint256[](owners.length);
        for (uint256 _i = 0; _i < owners.length; _i++) {
            balances_[_i] = this.balanceOf(owners[_i], ids[_i]);
        }
    }

    function setApprovalForAll(address operator, bool approved) external {
        revert("Not supported");
    }

    function isApprovedForAll(address owner, address operator) external view returns (bool) {
        return true;
    }

    // Trade functions

    /**
     * Perform Augur Trades using 0x signed orders
     *
     * @param  _requestedFillAmount  Share amount to fill
     * @param  _fingerprint          Fingerprint of the user to restrict affiliate fees
     * @param  _tradeGroupId         Random id to correlate these fills as one trade action
     * @param  _maxProtocolFeeDai    The maximum amount of DAI to spend on covering the 0x protocol fee
     * @param  _maxTrades            The maximum number of trades to actually take from the provided 0x orders
     * @param  _orders               Array of encoded Order struct data
     * @param  _signatures           Array of signature data
     * @return                       The amount the taker still wants
     */
    function trade(
        uint256 _requestedFillAmount,
        bytes32 _fingerprint,
        bytes32 _tradeGroupId,
        uint256 _maxProtocolFeeDai,
        uint256 _maxTrades,
        IExchange.Order[] memory _orders,
        bytes[] memory _signatures
    )
        public
        payable
        returns (uint256)
    {
        require(_orders.length > 0);
        uint256 _fillAmountRemaining = _requestedFillAmount;

        transferFromAllowed = true;

        uint256 _protocolFee = exchange.protocolFeeMultiplier().mul(tx.gasprice);
        coverProtocolFee(_protocolFee.mul(_orders.length), _maxProtocolFeeDai);

        // Do the actual asset exchanges
        for (uint256 i = 0; i < _orders.length && _fillAmountRemaining != 0; i++) {
            IExchange.Order memory _order = _orders[i];
            validateOrder(_order, _fillAmountRemaining);

            // Update 0x and pay protocol fee. This will also validate signatures and order state for us.
            IExchange.FillResults memory totalFillResults = exchange.fillOrder.value(_protocolFee)(
                _order,
                _fillAmountRemaining,
                _signatures[i]
            );

            if (totalFillResults.takerAssetFilledAmount == 0) {
                continue;
            }

            uint256 _amountTraded = doTrade(_order, totalFillResults.takerAssetFilledAmount, _fingerprint, _tradeGroupId, msg.sender);

            _fillAmountRemaining = _fillAmountRemaining.sub(_amountTraded);
            _maxTrades -= 1;
            if (_maxTrades == 0) {
                break;
            }
        }

        transferFromAllowed = false;

        if (address(this).balance > 0) {
            msg.sender.call.value(address(this).balance);
        }

        return _fillAmountRemaining;
    }

    function coverProtocolFee(uint256 _amountEthRequired, uint256 _maxProtocolFeeDai) internal {
        if (address(this).balance < _amountEthRequired) {
            uint256 _ethDeficit = _amountEthRequired - address(this).balance;
            uint256 _cost = ethExchange.getTokenPurchaseCost(_ethDeficit);
            require(_cost <= _maxProtocolFeeDai, "Cost of purchasing ETH to cover protocol Fee on the exchange was too high");
            cashTransferFrom(msg.sender, address(ethExchange), _cost);
            ethExchange.buyToken(address(this));
        }
    }

    function estimateProtocolFeeCostInCash(uint256 _numOrders, uint256 _gasPrice) public view returns (uint256) {
        uint256 _protocolFee = exchange.protocolFeeMultiplier().mul(_gasPrice);
        uint256 _amountEthRequired = _protocolFee.mul(_numOrders);
        return ethExchange.getTokenPurchaseCost(_amountEthRequired);
    }

    function validateOrder(IExchange.Order memory _order, uint256 _fillAmountRemaining) internal view {
        require(_order.takerAssetData.equals(encodeTakerAssetData()));
        require(_order.takerAssetAmount == _order.makerAssetAmount);
        (IERC1155 _zeroXTradeTokenMaker, uint256 _tokenIdMaker) = getZeroXTradeTokenData(_order.makerAssetData);
        (address _market, uint256 _price, uint8 _outcome, uint8 _type) = unpackTokenId(_tokenIdMaker);
        uint256 _numTicks = IMarket(_market).getNumTicks();
        uint256 _tradeInterval = TRADE_INTERVAL_VALUE.div(_numTicks);
        _tradeInterval = _tradeInterval.div(MIN_TRADE_INTERVAL).mul(MIN_TRADE_INTERVAL);
        _tradeInterval = MIN_TRADE_INTERVAL.max(_tradeInterval);
        require(_fillAmountRemaining.isMultipleOf(_tradeInterval), "Order must be a multiple of the market trade increment");
        require(_zeroXTradeTokenMaker == this);
    }

    function doTrade(IExchange.Order memory _order, uint256 _amount, bytes32 _fingerprint, bytes32 _tradeGroupId, address _taker) private returns (uint256 _amountFilled) {
        // parseOrderData will validate that the token being traded is the leigitmate one for the market
        AugurOrderData memory _augurOrderData = parseOrderData(_order);
        // If the signed order creator doesnt have enough funds we still want to continue and take their order out of the list
        // If the filler doesn't have funds this will just fail, which is fine
        if (!creatorHasFundsForTrade(_order, _amount)) {
            return 0;
        }
        // If the maker is also the taker we also just skip the trade but treat it as filled for amount remaining purposes
        if (_order.makerAddress == _taker) {
            return _amount;
        }
        (uint256 _amountRemaining, uint256 _fees) = fillOrder.fillZeroXOrder(IMarket(_augurOrderData.marketAddress), _augurOrderData.outcome, _augurOrderData.price, Order.Types(_augurOrderData.orderType), _order.makerAddress, _amount, _fingerprint, _tradeGroupId, _taker);
        _amountFilled = _amount.sub(_amountRemaining);
        logOrderFilled(_order, _augurOrderData, _taker, _tradeGroupId, _amountFilled, _fees);
        return _amountFilled;
    }

    function logOrderFilled(IExchange.Order memory _order, AugurOrderData memory _augurOrderData, address _taker, bytes32 _tradeGroupId, uint256 _amountFilled, uint256 _fees) private {
        bytes32 _orderHash = exchange.getOrderInfo(_order).orderHash;
        address[] memory _addressData = new address[](2);
        uint256[] memory _uint256Data = new uint256[](10);
        Order.Types _orderType = Order.Types(_augurOrderData.orderType);
        _addressData[0] = _order.makerAddress;
        _addressData[1] = _taker;
        _uint256Data[0] = _augurOrderData.price;
        _uint256Data[1] = 0;
        _uint256Data[2] = _augurOrderData.outcome;
        _uint256Data[5] = _fees;
        _uint256Data[6] = _amountFilled;
        _uint256Data[8] = 0;
        _uint256Data[9] = 0;
        augurTrading.logZeroXOrderFilled(IMarket(_augurOrderData.marketAddress).getUniverse(), IMarket(_augurOrderData.marketAddress), _orderHash, _tradeGroupId, uint8(_orderType), _addressData, _uint256Data);
    }

    function creatorHasFundsForTrade(IExchange.Order memory _order, uint256 _amount) public view returns (bool) {
        uint256 _tokenId = getTokenIdFromOrder(_order);
        return _amount <= this.balanceOf(_order.makerAddress, _tokenId);
    }

    function getTransferFromAllowed() public view returns (bool) {
        return transferFromAllowed;
    }

    /// @dev Encode MultiAsset proxy asset data into the format described in the AssetProxy contract specification.
    /// @param _market The address of the market to trade on
    /// @param _price The price used to trade
    /// @param _outcome The outcome to trade on
    /// @param _type Either BID == 0 or ASK == 1
    /// @return AssetProxy-compliant asset data describing the set of assets.
    function encodeAssetData(
        IMarket _market,
        uint256 _price,
        uint8 _outcome,
        uint8 _type
    )
        public
        view
        returns (bytes memory _assetData)
    {
        bytes[] memory _nestedAssetData = new bytes[](3);
        uint256[] memory _multiAssetValues = new uint256[](3);
        _nestedAssetData[0] = encodeTradeAssetData(_market, _price, _outcome, _type);
        _nestedAssetData[1] = encodeCashAssetData();
        _nestedAssetData[2] = encodeShareAssetData();
        _multiAssetValues[0] = 1;
        _multiAssetValues[1] = 0;
        _multiAssetValues[2] = 0;
        bytes memory _data = abi.encodeWithSelector(
            MULTI_ASSET_PROXY_ID,
            _multiAssetValues,
            _nestedAssetData
        );
        return _data;
    }

    /// @dev Encode ERC-1155 asset data into the format described in the AssetProxy contract specification.
    /// @param _market The address of the market to trade on
    /// @param _price The price used to trade
    /// @param _outcome The outcome to trade on
    /// @param _type Either BID == 0 or ASK == 1
    /// @return AssetProxy-compliant asset data describing the set of assets.
    function encodeTradeAssetData(
        IMarket _market,
        uint256 _price,
        uint8 _outcome,
        uint8 _type
    )
        private
        view
        returns (bytes memory _assetData)
    {
        uint256[] memory _tokenIds = new uint256[](1);
        uint256[] memory _tokenValues = new uint256[](1);

        uint256 _tokenId = getTokenId(address(_market), _price, _outcome, _type);
        _tokenIds[0] = _tokenId;
        _tokenValues[0] = 1;
        bytes memory _callbackData = new bytes(0);
        _assetData = abi.encodeWithSelector(
            ERC1155_PROXY_ID,
            address(this),
            _tokenIds,
            _tokenValues,
            _callbackData
        );

        return _assetData;
    }

    /// @dev Encode ERC-20 asset data into the format described in the AssetProxy contract specification.
    /// @return AssetProxy-compliant asset data describing the set of assets.
    function encodeCashAssetData()
        private
        view
        returns (bytes memory _assetData)
    {
        _assetData = abi.encodeWithSelector(
            ERC20_PROXY_ID,
            address(cash)
        );

        return _assetData;
    }

    /// @dev Encode ERC-1155 asset data into the format described in the AssetProxy contract specification.
    /// @return AssetProxy-compliant asset data describing the set of assets.
    function encodeShareAssetData()
        private
        view
        returns (bytes memory _assetData)
    {
        uint256[] memory _tokenIds = new uint256[](0);
        uint256[] memory _tokenValues = new uint256[](0);
        bytes memory _callbackData = new bytes(0);
        _assetData = abi.encodeWithSelector(
            ERC1155_PROXY_ID,
            address(shareToken),
            _tokenIds,
            _tokenValues,
            _callbackData
        );

        return _assetData;
    }

    /// @dev Encode ERC-1155 asset data into the format described in the AssetProxy contract specification.
    /// @return AssetProxy-compliant asset data describing the set of assets.
    function encodeTakerAssetData()
        private
        view
        returns (bytes memory _assetData)
    {
        uint256[] memory _tokenIds = new uint256[](0);
        uint256[] memory _tokenValues = new uint256[](0);
        bytes memory _callbackData = new bytes(0);
        _assetData = abi.encodeWithSelector(
            ERC1155_PROXY_ID,
            address(this),
            _tokenIds,
            _tokenValues,
            _callbackData
        );

        return _assetData;
    }

    function getTokenId(address _market, uint256 _price, uint8 _outcome, uint8 _type) public pure returns (uint256 _tokenId) {
        // NOTE: we're assuming no one needs a full uint256 for the price value here and cutting to uint80 so we can pack this in a uint256.
        bytes memory _tokenIdBytes = abi.encodePacked(_market, uint80(_price), _outcome, _type);
        assembly {
            _tokenId := mload(add(_tokenIdBytes, add(0x20, 0)))
        }
    }

    function unpackTokenId(uint256 _tokenId) public pure returns (address _market, uint256 _price, uint8 _outcome, uint8 _type) {
        assembly {
            _market := shr(96, and(_tokenId, 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF000000000000000000000000))
            _price := shr(16,  and(_tokenId, 0x0000000000000000000000000000000000000000FFFFFFFFFFFFFFFFFFFF0000))
            _outcome := shr(8, and(_tokenId, 0x000000000000000000000000000000000000000000000000000000000000FF00))
            _type :=           and(_tokenId, 0x00000000000000000000000000000000000000000000000000000000000000FF)
        }
    }

    /// @dev Decode MultiAsset asset data from the format described in the AssetProxy contract specification.
    /// @param _assetData AssetProxy-compliant asset data describing an ERC-1155 set of assets.
    /// @return The ERC-1155 AssetProxy identifier, the address of this ERC-1155
    /// contract hosting the assets, an array of the identifiers of the
    /// assets to be traded, an array of asset amounts to be traded, and
    /// callback data.  Each element of the arrays corresponds to the
    /// same-indexed element of the other array.  Return values specified as
    /// `memory` are returned as pointers to locations within the memory of
    /// the input parameter `assetData`.
    function decodeAssetData(bytes memory _assetData)
        public
        view
        returns (
            bytes4 _assetProxyId,
            address _tokenAddress,
            uint256[] memory _tokenIds,
            uint256[] memory _tokenValues,
            bytes memory _callbackData
        )
    {
         // Read the bytes4 from array memory
        assembly {
            _assetProxyId := mload(add(_assetData, 32))
            // Solidity does not require us to clean the trailing bytes. We do it anyway
            _assetProxyId := and(_assetProxyId, 0xFFFFFFFF00000000000000000000000000000000000000000000000000000000)
        }

        require(_assetProxyId == MULTI_ASSET_PROXY_ID, "WRONG_PROXY_ID");

        uint256[] memory _amounts;
        bytes[] memory _nestedAssetData;

        // Slice the selector off the asset data
        bytes memory _noSelectorAssetData = _assetData.slice(4, _assetData.length);

        (_amounts, _nestedAssetData) = abi.decode(_noSelectorAssetData, (uint256[], bytes[]));
        
        // Validate storage refs against the decoded values.
        {
            require(_amounts.length == 3);
            require(_amounts[0] == 1);
            require(_amounts[1] == 0);
            require(_amounts[2] == 0);
            require(_nestedAssetData[1].equals(encodeCashAssetData()));
            require(_nestedAssetData[2].equals(encodeShareAssetData()));
        }

        return decodeTradeAssetData(_nestedAssetData[0]);
    }

    /// @dev Decode ERC-1155 asset data from the format described in the AssetProxy contract specification.
    /// @param _assetData AssetProxy-compliant asset data describing an ERC-1155 set of assets.
    /// @return The ERC-1155 AssetProxy identifier, the address of this ERC-1155
    /// contract hosting the assets, an array of the identifiers of the
    /// assets to be traded, an array of asset amounts to be traded, and
    /// callback data.  Each element of the arrays corresponds to the
    /// same-indexed element of the other array.  Return values specified as
    /// `memory` are returned as pointers to locations within the memory of
    /// the input parameter `assetData`.
    function decodeTradeAssetData(bytes memory _assetData)
        public
        pure
        returns (
            bytes4 _assetProxyId,
            address _tokenAddress,
            uint256[] memory _tokenIds,
            uint256[] memory _tokenValues,
            bytes memory _callbackData
        )
    {
         // Read the bytes4 from array memory
        assembly {
            _assetProxyId := mload(add(_assetData, 32))
            // Solidity does not require us to clean the trailing bytes. We do it anyway
            _assetProxyId := and(_assetProxyId, 0xFFFFFFFF00000000000000000000000000000000000000000000000000000000)
        }

        require(_assetProxyId == ERC1155_PROXY_ID, "WRONG_PROXY_ID");

        assembly {
            let _length := mload(_assetData)
            // Skip the length (of bytes variable) and the selector to get to the first parameter.
            _assetData := add(_assetData, 36)
            // Read the value of the first parameter:
            _tokenAddress := mload(_assetData)
            _tokenIds := add(_assetData, mload(add(_assetData, 32)))
            _tokenValues := add(_assetData, mload(add(_assetData, 64)))
            _callbackData := add(_assetData, mload(add(_assetData, 96)))
        }

        return (
            _assetProxyId,
            _tokenAddress,
            _tokenIds,
            _tokenValues,
            _callbackData
        );
    }

    function parseOrderData(IExchange.Order memory _order) public view returns (AugurOrderData memory _data) {
        (bytes4 _assetProxyId, address _tokenAddress, uint256[] memory _tokenIds, uint256[] memory _tokenValues, bytes memory _callbackData) = decodeAssetData(_order.makerAssetData);
        (address _market, uint256 _price, uint8 _outcome, uint8 _type) = unpackTokenId(_tokenIds[0]);
        _data.marketAddress = _market;
        _data.price = _price;
        _data.orderType = _type;
        _data.outcome = _outcome;
    }

    function getZeroXTradeTokenData(bytes memory _assetData) public view returns (IERC1155 _token, uint256 _tokenId) {
        (bytes4 _assetProxyId, address _tokenAddress, uint256[] memory _tokenIds, uint256[] memory _tokenValues, bytes memory _callbackData) = decodeAssetData(_assetData);
        _tokenId = _tokenIds[0];
        _token = IERC1155(_tokenAddress);
    }

    function getTokenIdFromOrder(IExchange.Order memory _order) public view returns (uint256 _tokenId) {
        (bytes4 _assetProxyId, address _tokenAddress, uint256[] memory _tokenIds, uint256[] memory _tokenValues, bytes memory _callbackData) = decodeAssetData(_order.makerAssetData);
        _tokenId = _tokenIds[0];
    }

    function createZeroXOrder(uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) public view returns (IExchange.Order memory _zeroXOrder, bytes32 _orderHash) {
        return createZeroXOrderFor(msg.sender, _type, _attoshares, _price, _market, _outcome, _expirationTimeSeconds, _salt);
    }

    function createZeroXOrderFor(address _maker, uint8 _type, uint256 _attoshares, uint256 _price, address _market, uint8 _outcome, uint256 _expirationTimeSeconds, uint256 _salt) public view returns (IExchange.Order memory _zeroXOrder, bytes32 _orderHash) {
        bytes memory _assetData = encodeAssetData(IMarket(_market), _price, _outcome, _type);
        uint256 _numTicks = IMarket(_market).getNumTicks();
        uint256 _tradeInterval = TRADE_INTERVAL_VALUE / _numTicks;
        require(_attoshares.isMultipleOf(_tradeInterval), "Order must be a multiple of the market trade increment");
        _zeroXOrder.makerAddress = _maker;
        _zeroXOrder.makerAssetAmount = _attoshares;
        _zeroXOrder.takerAssetAmount = _attoshares;
        _zeroXOrder.expirationTimeSeconds = _expirationTimeSeconds;
        _zeroXOrder.salt = _salt;
        _zeroXOrder.makerAssetData = _assetData;
        _zeroXOrder.takerAssetData = encodeTakerAssetData();
        _orderHash = exchange.getOrderInfo(_zeroXOrder).orderHash;
    }

    function encodeEIP1271OrderWithHash(
        IExchange.Order memory _zeroXOrder,
        bytes32 _orderHash
    )
        public
        pure
        returns (bytes memory encoded)
    {
        return abi.encodeWithSelector(
            EIP1271_ORDER_WITH_HASH_SELECTOR,
            _zeroXOrder,
            _orderHash
        );
    }

    function () external payable {}
}
