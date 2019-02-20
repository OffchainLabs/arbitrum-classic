pragma solidity ^0.5.1;

contract Fibonacci {

uint[] fibseries;

  // n = how many in the series to return
  function generateFib(uint n) public payable {

    // set 1st and 2nd entries
    fibseries.push(1);
    fibseries.push(1);

    // generate subsequent entries
    for (uint i=2; i < n ; i++) {
      fibseries.push(fibseries[i-1] + fibseries[i-2]);
    }

  }

  function getFib(uint n) public view returns (uint) {
    return fibseries[n];
  }

  // function getFib(uint n) public view returns (uint, uint) {
  //   return (fibseries[n], fibseries[n + 1]);
  // }

}