pragma solidity >=0.4.21 <0.7.0;

interface ArbosTest {
    function installAccount(
        address addr,
        bool isEOA,
        uint256 balance,
        uint256 nonce,
        bytes calldata code,
        bytes calldata initStorage
    ) external;

    function getMarshalledStorage(address addr) external view; // returns raw returndata

    function getAccountInfo(address addr) external view; // returns raw returndata

    function burnArbGas(uint256 gasAmount) external view;

    function setNonce(address addr, uint256 nonce) external;

    function setBalance(address addr, uint256 balance) external;

    function setCode(address addr, bytes calldata code) external;

    function setState(address addr, bytes calldata state) external;

    function store(
        address addr,
        uint32 key,
        uint32 value
    ) external;
}

contract EthCallTester {
    uint256 x;

    uint256 constant abostest_address = 105;

    constructor() public {
        x = 0x100;
    }

    function failStore() external {
        ArbosTest(abostest_address).store(address(this), 0x0, 0x3000000);
    }

    function getX() external returns (uint256) {
        return x;
    }

    function sLoad(uint256 key) external returns (uint256) {
        uint256 val;
        assembly {
            val := sload(key)
        }
        return val;
    }

    function getBalance() external returns (uint256) {
        return address(this).balance;
    }
}
