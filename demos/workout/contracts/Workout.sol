pragma solidity ^0.5.0;

contract Workout {
	mapping (uint => uint) kvs;
	uint spinner = 2;

	uint constant bigPrime = (1<<254)-245;

	function set(uint key, uint val) public {
 		kvs[key] = val;
	}

	function get(uint key) public view returns (uint) {
		return kvs[key];
	}

	function compute(uint numSteps) public {
		for (uint i=0; i<numSteps; i++) {
			spinner = (2 * spinner) % bigPrime;
		}
	}
}
