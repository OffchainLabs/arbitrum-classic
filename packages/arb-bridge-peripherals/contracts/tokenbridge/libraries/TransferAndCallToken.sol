// SPDX-License-Identifier: MIT
pragma solidity >0.6.0 <0.8.0;

import "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import "./ITransferAndCall.sol";

// Implementation from https://github.com/smartcontractkit/LinkToken/blob/master/contracts/v0.6/TransferAndCallToken.sol
/**
 * @notice based on Implementation from https://github.com/smartcontractkit/LinkToken/blob/master/contracts/v0.6/ERC677Token.sol
 * The implementation doesn't return a bool on onTokenTransfer. This is similar to the proposed 677 standard, but still incompatible - thus we don't refer to it as such.
 */
abstract contract TransferAndCallToken is ERC20Upgradeable, ITransferAndCall {
    /**
     * @dev transfer token to a contract address with additional data if the recipient is a contact.
     * @param _to The address to transfer to.
     * @param _value The amount to be transferred.
     * @param _data The extra data to be passed to the receiving contract.
     */
    function transferAndCall(
        address _to,
        uint256 _value,
        bytes memory _data
    ) public virtual override returns (bool success) {
        super.transfer(_to, _value);
        emit Transfer(msg.sender, _to, _value, _data);
        if (isContract(_to)) {
            contractFallback(_to, _value, _data);
        }
        return true;
    }

    // PRIVATE

    function contractFallback(
        address _to,
        uint256 _value,
        bytes memory _data
    ) private {
        ITransferAndCallReceiver receiver = ITransferAndCallReceiver(_to);
        receiver.onTokenTransfer(msg.sender, _value, _data);
    }

    function isContract(address _addr) private view returns (bool hasCode) {
        uint256 length;
        assembly {
            length := extcodesize(_addr)
        }
        return length > 0;
    }
}
