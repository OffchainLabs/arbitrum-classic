pragma solidity 0.5.15;

import '../reporting/IMarket.sol';


library TokenId {

    function getTokenId(IMarket _market, uint256 _outcome) internal pure returns (uint256 _tokenId) {
        bytes memory _tokenIdBytes = abi.encodePacked(_market, uint8(_outcome));
        assembly {
            _tokenId := mload(add(_tokenIdBytes, add(0x20, 0)))
        }
    }

    function getTokenIds(IMarket _market, uint256[] memory _outcomes) internal pure returns (uint256[] memory _tokenIds) {
        _tokenIds = new uint256[](_outcomes.length);
        for (uint256 _i = 0; _i < _outcomes.length; _i++) {
            _tokenIds[_i] = getTokenId(_market, _outcomes[_i]);
        }
    }

    function unpackTokenId(uint256 _tokenId) internal pure returns (address _market, uint256 _outcome) {
        assembly {
            _market := shr(96,  and(_tokenId, 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF000000000000000000000000))
            _outcome := shr(88, and(_tokenId, 0x0000000000000000000000000000000000000000FF0000000000000000000000))
        }
    }
}