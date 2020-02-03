
pragma solidity 0.5.15;


/**
 * @title BytesToString
 * @dev Bytes operations to convert to string and remove null characters
 */
library BytesToString {

    function bytes32ToString(bytes32 _bytes32) internal pure returns (string memory) {
        uint256 _stringLen = 0;
        for (uint256 _j = 0; _j < 32; _j++) {
            byte _char = _bytes32[_j];
            if (_char != byte(0)) {
                _stringLen += 1;
            }
        }
        bytes memory _bytesArray = new bytes(_stringLen);
        uint256 _bytesArrayIndex = 0;
        for (uint256 _j = 0; _j < 32; _j++) {
            byte _char = _bytes32[_j];
            if (_char != byte(0)) {
                _bytesArray[_bytesArrayIndex] = _char;
            }
            _bytesArrayIndex += 1;
        }
        return string(_bytesArray);
    }
}