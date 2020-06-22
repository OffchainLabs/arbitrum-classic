/*
 * Copyright 2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

  pragma solidity ^0.5.3;

  import "../vm/RollupUtils.sol";
  import "../vm/NodeGraphUtils.sol";

  contract RollupTester {

  	function calculateLeafFromPath(
  		bytes32 from,
        bytes32[] memory proof) public pure returns(bytes32) 
  	{
  		return RollupUtils.calculateLeafFromPath(from, proof);
  	}

  	function childNodeHash(
  		bytes32 prevNodeHash,
        uint256 deadlineTicks,
        bytes32 nodeDataHash,
        uint256 childType,
        bytes32 vmProtoStateHash) public pure returns(bytes32)
  	{
  		return RollupUtils.childNodeHash(
  			prevNodeHash, 
  			deadlineTicks, 
  			nodeDataHash, 
  			childType, 
  			vmProtoStateHash);
  	}

  	function computeProtoHashBefore(
  		bytes32 machineHash,
        bytes32 inboxTop,
        uint256 inboxCount) public pure returns (bytes32)
  	{
  		return RollupUtils.protoStateHash(machineHash, inboxTop, inboxCount);
  	}

  	function computePrevLeaf(
  		bytes32[9] memory _fields,
        uint256 _beforeInboxCount,
        uint256 _prevDeadlineTicks,
        uint32 _prevChildType,
        uint64 _numSteps,
        uint128[4] memory _timeBounds,
        uint256 _importedMessageCount,
        bool _didInboxInsn,
        uint64 _numArbGas) public pure returns (bytes32, bytes32)
  	{
  		NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils.AssertionData(
            _fields[0],
            _fields[1],
            _beforeInboxCount,

            _fields[2],
            _prevDeadlineTicks,
            _fields[3],
            _prevChildType,

            _numSteps,
            _timeBounds,
            _importedMessageCount,

            _fields[4],

            _fields[5],

            _fields[6],
            _didInboxInsn,
            _numArbGas,
            _fields[7],
            _fields[8]
        );

        return NodeGraphUtils.computePrevLeaf(assertData);
  	}

	function generateInvalidInboxTopLeaf(
        uint256[4] memory invalidInboxData,
  		bytes32[9] memory _fields,
        uint256 _beforeInboxCount,
        uint256 _prevDeadlineTicks,
        uint32 _prevChildType,
        uint64 _numSteps,
        uint128[4] memory _timeBounds,
        uint256 _importedMessageCount,
        bool _didInboxInsn,
        uint64 _numArbGas) public pure returns(bytes32)
	{
		NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils.AssertionData(
            _fields[0],
            _fields[1],
            _beforeInboxCount,

            _fields[2],
            _prevDeadlineTicks,
            _fields[3],
            _prevChildType,

            _numSteps,
            _timeBounds,
            _importedMessageCount,

            _fields[4],

            _fields[5],

            _fields[6],
            _didInboxInsn,
            _numArbGas,
            _fields[7],
            _fields[8]
        );

		return _generateInvalidInboxTopLeaf(
            assertData,
            invalidInboxData
        );
	}

	function _generateInvalidInboxTopLeaf(
		NodeGraphUtils.AssertionData memory assertData,
		uint256[4] memory invalidInboxData
  	) 
		private pure returns(bytes32)
	{
		(bytes32 prevLeaf, bytes32 vmProtoHashBefore) = NodeGraphUtils.computePrevLeaf(assertData);

		return NodeGraphUtils.generateInvalidInboxTopLeaf(
            assertData,
            prevLeaf,
            invalidInboxData[3],
            bytes32(invalidInboxData[0]),
            invalidInboxData[1],
            vmProtoHashBefore,
            invalidInboxData[2]
        );
	}

	function generateInvalidMessagesLeaf(
		uint256 gracePeriodTicks,
		uint256 deadlineTicks,
  		bytes32[9] memory _fields,
        uint256 _beforeInboxCount,
        uint256 _prevDeadlineTicks,
        uint32 _prevChildType,
        uint64 _numSteps,
        uint128[4] memory _timeBounds,
        uint256 _importedMessageCount,
        bool _didInboxInsn,
        uint64 _numArbGas) public pure returns(bytes32)
	{
		NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils.AssertionData(
            _fields[0],
            _fields[1],
            _beforeInboxCount,

            _fields[2],
            _prevDeadlineTicks,
            _fields[3],
            _prevChildType,

            _numSteps,
            _timeBounds,
            _importedMessageCount,

            _fields[4],

            _fields[5],

            _fields[6],
            _didInboxInsn,
            _numArbGas,
            _fields[7],
            _fields[8]
        );

		return _generateInvalidMessagesLeaf(
            assertData,
            gracePeriodTicks,
            deadlineTicks
        );
	}

	function _generateInvalidMessagesLeaf(
		NodeGraphUtils.AssertionData memory assertData,
		uint256 gracePeriodTicks,
		uint256 deadlineTicks
	) 
		private pure returns(bytes32)
	{
		(bytes32 prevLeaf, bytes32 vmProtoHashBefore) = NodeGraphUtils.computePrevLeaf(assertData);

		return NodeGraphUtils.generateInvalidMessagesLeaf(
            assertData,
            prevLeaf,
            deadlineTicks,
            vmProtoHashBefore,
            gracePeriodTicks
        );
	}

	function generateInvalidExecutionLeaf(
		uint256 gracePeriodTicks,
		uint256 checkTimeTicks,
		uint256 deadlineTicks,
  		bytes32[9] memory _fields,
        uint256 _beforeInboxCount,
        uint256 _prevDeadlineTicks,
        uint32 _prevChildType,
        uint64 _numSteps,
        uint128[4] memory _timeBounds,
        uint256 _importedMessageCount,
        bool _didInboxInsn,
        uint64 _numArbGas) public pure returns(bytes32)
	{
		NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils.AssertionData(
            _fields[0],
            _fields[1],
            _beforeInboxCount,

            _fields[2],
            _prevDeadlineTicks,
            _fields[3],
            _prevChildType,

            _numSteps,
            _timeBounds,
            _importedMessageCount,

            _fields[4],

            _fields[5],

            _fields[6],
            _didInboxInsn,
            _numArbGas,
            _fields[7],
            _fields[8]
        );

		return _generateInvalidExecutionLeaf(
            assertData,
            gracePeriodTicks,
            checkTimeTicks,
            deadlineTicks
        );
	}

	function _generateInvalidExecutionLeaf(
		NodeGraphUtils.AssertionData memory assertData,
		uint256 gracePeriodTicks,
		uint256 checkTimeTicks,
		uint256 deadlineTicks
	)
	 	private pure returns(bytes32)
	{
		(bytes32 prevLeaf, bytes32 vmProtoHashBefore) = NodeGraphUtils.computePrevLeaf(assertData);

		return NodeGraphUtils.generateInvalidExecutionLeaf(
            assertData,
            prevLeaf,
            deadlineTicks,
            vmProtoHashBefore,
            gracePeriodTicks,
            checkTimeTicks
        );
	}

	function generateValidLeaf(
		uint256 deadlineTicks,
  		bytes32[9] memory _fields,
        uint256 _beforeInboxCount,
        uint256 _prevDeadlineTicks,
        uint32 _prevChildType,
        uint64 _numSteps,
        uint128[4] memory _timeBounds,
        uint256 _importedMessageCount,
        bool _didInboxInsn,
        uint64 _numArbGas) public pure returns(bytes32)
	{
		NodeGraphUtils.AssertionData memory assertData = NodeGraphUtils.AssertionData(
            _fields[0],
            _fields[1],
            _beforeInboxCount,

            _fields[2],
            _prevDeadlineTicks,
            _fields[3],
            _prevChildType,

            _numSteps,
            _timeBounds,
            _importedMessageCount,

            _fields[4],

            _fields[5],

            _fields[6],
            _didInboxInsn,
            _numArbGas,
            _fields[7],
            _fields[8]
        );

		return _generateValidLeaf(
            assertData,
            deadlineTicks
        );
	}

	function _generateValidLeaf(
		NodeGraphUtils.AssertionData memory assertData,
		uint256 deadlineTicks) private pure returns(bytes32)
	{
		(bytes32 prevLeaf, ) = NodeGraphUtils.computePrevLeaf(assertData);

		return NodeGraphUtils.generateValidLeaf(
            assertData,
            prevLeaf,
            deadlineTicks
        );
	}

}