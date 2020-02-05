
pragma solidity 0.5.15;
pragma experimental ABIEncoderV2;

import "../external/IExchange.sol";


contract IZeroXTrade {

    struct AugurOrderData {
        address marketAddress;                  // Market Address
        uint256 price;                          // Price
        uint8 outcome;                          // Outcome
        uint8 orderType;                        // Order Type
    }

    function parseOrderData(IExchange.Order memory _order) public view returns (AugurOrderData memory _data);
    function unpackTokenId(uint256 _tokenId) public pure returns (address _market, uint256 _price, uint8 _outcome, uint8 _type);
}