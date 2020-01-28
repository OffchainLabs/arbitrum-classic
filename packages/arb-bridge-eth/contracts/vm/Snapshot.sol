/*
 * Copyright 2019, Offchain Labs, Inc.
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

contract Snapshot {
	mapping (address => mapping (uint256 => bytes32)) snapshots;

	enum SnapshotType { LatestConfirmed, TwoStakers, DeadlineStakers, NodeExists }

	event SavedLatestConfirmedSnapshot(
		address client, 
		bytes32 latestConfirmed, 
		bytes32 snapshot
	);
	event SavedTwoStakersSnapshot(
		address client, 
		address addr1, 
		bytes32 location1, 
		address addr2, 
		bytes32 location2, 
		bytes32 snapshot
	);
	event SavedDeadlineStakersSnapshot(
		address client, 
		uint256 deadlineTicks, 
		bytes32[] stakerLocations, 
		bytes32 snapshot
	);
	event SavedNodeExistsSnapshot(
		address client, 
		bytes32 nodeHash, 
		bytes32 snapshot
	);

	function getMySnapshot(uint256 idx) public view returns(bytes32) {
		return snapshots[msg.sender][idx];
	}

	function calcLatestConfirmedSnapshot(bytes32 lcHash) internal pure returns(bytes32) {
		return keccak256(
				abi.encodePacked(
					SnapshotType.LatestConfirmed,
					lcHash
				)
			);
	}

	function saveLatestConfirmedSnapshot(uint256 idx, bytes32 lcHash) internal {
		bytes32 snap = calcLatestConfirmedSnapshot(lcHash);
		snapshots[msg.sender][idx] = snap;
		emit SavedLatestConfirmedSnapshot(msg.sender, lcHash, snap);
	}

	function calcTwoStakersSnapshot(address addr1, bytes32 loc1, address addr2, bytes32 loc2) 
		internal 
		pure 
		returns(bytes32) 
	{
		return keccak256(
			abi.encodePacked(
				SnapshotType.TwoStakers,
				addr1,
				loc1,
				addr2,
				loc2
			)
		);
	}

	function saveTwoStakersSnapshot(uint256 idx, address addr1, bytes32 loc1, address addr2, bytes32 loc2) 
		internal 
	{
		bytes32 snap = calcTwoStakersSnapshot(addr1, loc1, addr2, loc2);
		snapshots[msg.sender][idx] = snap;
		emit SavedTwoStakersSnapshot(msg.sender, addr1, loc1, addr2, loc2, snap);
	}

	function calcDeadlineStakersSnapshot(uint256 deadlineTicks, bytes32[] memory locations) 
		internal 
		pure 
		returns(bytes32) 
	{
		bytes32 acc = bytes32(0);
		for(uint i=0; i<locations.length; i++) {
			acc = keccak256(
				abi.encodePacked(
					locations[i],
					acc
				)
			);
		}
		return keccak256(
			abi.encodePacked(
				SnapshotType.DeadlineStakers,
				deadlineTicks,
				acc
			)
		);
	}

	function saveDeadlineStakersSnapshot(
		uint256 idx, 
		uint256 deadlineTicks, 
		bytes32[] memory locations
	)
		internal
	{
		bytes32 snap = calcDeadlineStakersSnapshot(deadlineTicks, locations);
		snapshots[msg.sender][idx] = snap;
		emit SavedDeadlineStakersSnapshot(msg.sender, deadlineTicks, locations, snap);
	}

	function calcNodeExistsSnapshot(bytes32 location) internal pure returns(bytes32) {
		return keccak256(
			abi.encodePacked(
				SnapshotType.NodeExists,
				location
			)
		);
	}

	function saveNodeExistsSnapshot(uint256 idx, bytes32 location) internal {
		bytes32 snap = calcNodeExistsSnapshot(location);
		snapshots[msg.sender][idx] = snap;
		emit SavedNodeExistsSnapshot(msg.sender, location, snap);
	}
}

