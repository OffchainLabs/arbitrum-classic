// SPDX-License-Identifier: Apache-2.0

//SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

library DecimalConverter {
  function from777to20(uint8 decimals, uint amount) internal pure returns (uint256) {
    require(decimals <= 18, 'DEC');
    return amount / (10 ** uint256(18 - decimals));
  }

  function from20to777(uint8 decimals, uint amount) internal pure returns (uint256) {
    require(decimals <= 18, 'DEC');
    return amount * (10 ** uint256(18 - decimals));
  }
}
