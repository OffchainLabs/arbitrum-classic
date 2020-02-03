pragma solidity 0.5.15;


contract IGnosisSafe {
    address public masterCopy;
    function getThreshold() public view returns (uint256);
    function getOwners() public view returns (address[] memory);
    function setup(
        address[] calldata _owners,
        uint256 _threshold,
        address to,
        bytes calldata data,
        address fallbackHandler,
        address paymentToken,
        uint256 payment,
        address payable paymentReceiver
        ) external;
}