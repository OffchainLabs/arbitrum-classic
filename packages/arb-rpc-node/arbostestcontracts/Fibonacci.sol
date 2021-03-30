pragma solidity >=0.4.21 <0.7.0;

contract Fibonacci {
    uint256[] fibseries;

    event TestEvent(uint256 number);

    // n = how many in the series to return
    function generateFib(uint256 n) public payable {
        // set 1st and 2nd entries
        fibseries.push(1);
        fibseries.push(1);

        // generate subsequent entries
        for (uint256 i = 2; i < n; i++) {
            fibseries.push(fibseries[i - 1] + fibseries[i - 2]);
        }

        emit TestEvent(n);
    }

    function getFib(uint256 n) public view returns (uint256) {
        return fibseries[n];
    }

    // function getFib(uint n) public view returns (uint, uint) {
    //   return (fibseries[n], fibseries[n + 1]);
    // }
}
