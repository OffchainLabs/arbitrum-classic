pragma solidity 0.5.15;

import "../ArbSys.sol";

// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-1167.md

/* Template Code for the create clone method:
  function createClone(address target) internal returns (address result) {
    bytes20 targetBytes = bytes20(target)${bytes == 20 ? "" : "<<" + ((20 - bytes) * 8)};
    assembly {
      let clone := mload(0x40)
      mstore(clone, 0x${code.substring(0, 2*(cloner.labels.address + 1)).padEnd(64, '0')})
      mstore(add(clone, 0x${(cloner.labels.address + 1).toString(16)}), targetBytes)
      mstore(add(clone, 0x${(cloner.labels.address + bytes + 1).toString(16)}), 0x${code.substring(2*(cloner.labels.address + bytes + 1), 2*(cloner.labels.address+bytes+1) + 30).padEnd(64, '0')})
      result := create(0, clone, 0x${(code.length / 2).toString(16)})
    }
  }
*/

contract CloneFactory {
    address public templateAddress;

    function initializeFactory(address _templateAddress) public {
        require(templateAddress == address(0));
        require(_templateAddress != address(0));
        templateAddress = _templateAddress;
    }

    function createNewContract() public returns (address) {
        require(templateAddress != address(0));

        address payable newContractAddress = address(uint160(ArbSys(100).cloneContract(templateAddress)));
        return newContractAddress;
    }


    // function createClone(address target) internal returns (address result) {
    //     // convert address to bytes20 for assembly use
    //     bytes20 targetBytes = bytes20(target);
    //     assembly {
    //         // allocate clone memory
    //         let clone := mload(0x40)
    //         // store initial portion of the delegation contract code in bytes form
    //         mstore(clone, 0x3d602d80600a3d3981f3363d3d373d3d3d363d73000000000000000000000000)
    //         // store the provided address
    //         mstore(add(clone, 0x14), targetBytes)
    //         // store the remaining delegation contract code
    //         mstore(add(clone, 0x28), 0x5af43d82803e903d91602b57fd5bf30000000000000000000000000000000000)
    //         // create the actual delegate contract reference and return its address
    //         result := create(0, clone, 0x37)
    //     }
    // }
}
