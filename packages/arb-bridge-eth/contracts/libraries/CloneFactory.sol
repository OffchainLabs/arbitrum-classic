// Taken from https://github.com/AugurProject/augur/blob/master/packages/augur-core/source/contracts/libraries/CloneFactory.sol

pragma solidity ^0.5.10;

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

contract Spawn {
  constructor(address logicContract) public {
    // place eip-1167 runtime code in memory.
    bytes memory runtimeCode = abi.encodePacked(
      bytes10(0x363d3d373d3d3d363d73),
      logicContract,
      bytes15(0x5af43d82803e903d91602b57fd5bf3)
    );

    // return eip-1167 code to write it to spawned contract runtime.
    assembly {
      return(add(0x20, runtimeCode), 45) // eip-1167 runtime code, length
    }
  }
}


contract CloneFactory {
    function createClone(address target) internal returns (address spawnedContract) {
        bytes memory initCode = abi.encodePacked(
            type(Spawn).creationCode,
            abi.encode(target)
        );
        assembly {
            let encoded_data := add(0x20, initCode) // load initialization code.
            let encoded_size := mload(initCode)     // load the init code's length.
            spawnedContract := create(              // call `CREATE` w/ 3 arguments.
                callvalue,                          // forward any supplied endowment.
                encoded_data,                       // pass in initialization code.
                encoded_size                        // pass in init code's length.
            )

          // pass along failure message from failed contract deployment and revert.
            if iszero(spawnedContract) {
                returndatacopy(0, 0, returndatasize)
                revert(0, returndatasize)
            }
        }
    }

    function create2Clone(address target, uint256 salt) internal returns (address spawnedContract) {
        bytes memory initCode = abi.encodePacked(
            type(Spawn).creationCode,
            abi.encode(target)
        );
        assembly {
            let encoded_data := add(0x20, initCode) // load initialization code.
            let encoded_size := mload(initCode)     // load the init code's length.
            spawnedContract := create2(             // call `CREATE2` w/ 4 arguments.
                callvalue,                          // forward any supplied endowment.
                encoded_data,                       // pass in initialization code.
                encoded_size,                       // pass in init code's length.
                salt                                // pass in the salt value.
            )

          // pass along failure message from failed contract deployment and revert.
            if iszero(spawnedContract) {
                returndatacopy(0, 0, returndatasize)
                revert(0, returndatasize)
            }
        }
    }

    function cloneCodeHash(address target) internal pure returns (bytes32 result) {
        bytes memory initCode = abi.encodePacked(
            type(Spawn).creationCode,
            abi.encode(target)
        );
        return keccak256(initCode);
    }
}
