pragma solidity ^0.5.0;

contract Workout {
	mapping (uint => uint) kvs;

	function set(uint key, uint val) public {
 		kvs[key] = val;
	}
/*
	function get(uint key) public view returns (uint) {
		return kvs[key];
	}
*/
}
