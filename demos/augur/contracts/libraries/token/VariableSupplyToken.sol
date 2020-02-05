pragma solidity 0.5.15;


import './ERC20.sol';


/**
 * @title Variable Supply Token
 * @notice A Standard Token wrapper which adds the ability to internally burn and mint tokens
 */
contract VariableSupplyToken is ERC20 {
    using SafeMathUint256 for uint256;

    function mint(address _target, uint256 _amount) internal returns (bool) {
        _mint(_target, _amount);
        onMint(_target, _amount);
        return true;
    }

    function burn(address _target, uint256 _amount) internal returns (bool) {
        _burn(_target, _amount);
        onBurn(_target, _amount);
        return true;
    }

    // Subclasses of this token may want to send additional logs through the centralized Augur log emitter contract
    function onMint(address, uint256) internal {
    }

    // Subclasses of this token may want to send additional logs through the centralized Augur log emitter contract
    function onBurn(address, uint256) internal {
    }
}
