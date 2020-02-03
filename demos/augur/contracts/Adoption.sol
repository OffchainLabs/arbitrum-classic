pragma solidity ^0.5.0;

contract Adoption {
	address[16] public adopters = [
		address(0),
		address(0),
		address(0),
		address(0),
		address(0),
		address(0),
		address(0),
		address(0),	
		address(0),
		address(0),
		address(0),
		address(0),
		address(0),
		address(0),
		address(0),
		address(0)
	];

	// Adopting a pet
	function adopt(uint petId) public returns (uint) {
  		require(petId >= 0 && petId <= 15);

  		adopters[petId] = msg.sender;

  		return petId;
	}

	// Retrieving the adopters
	function getAdopters() public view returns (address[16] memory) {
  		return adopters;
	}
}