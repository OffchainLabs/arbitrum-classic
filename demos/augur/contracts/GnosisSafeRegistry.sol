pragma solidity 0.5.15;

import './IAugur.sol';
import './trading/IAugurTrading.sol';
import './external/IGnosisSafe.sol';
import './external/IProxyFactory.sol';
import './libraries/Initializable.sol';
import './external/IProxy.sol';
import './libraries/token/IERC1155.sol';
import './reporting/IAffiliates.sol';

/**
 * @title Gnosis Safe Registry
 * @notice A contract that contains a mapping of user addresses to Gnosis Safes which have been set up for easier Augur use
 */
contract GnosisSafeRegistry is Initializable {
    // mapping of user to created safes. The first safe wins but we store additional safes in case a user somehow makes multiple. The current safe may be de-registered by itself and the current safe will simply become the next one in line
    mapping (address => IGnosisSafe[]) public accountSafes;
    mapping (address => uint256) public accountSafeIndexes;

    IAugur public augur;
    IAugurTrading public augurTrading;
    bytes32 public proxyCodeHash;
    bytes32 public deploymentData;
    address public gnosisSafeMasterCopy;
    IProxyFactory public proxyFactory;

    address public cash;
    address public affiliates;
    address public shareToken;
    address public createOrder;
    address public fillOrder;
    address public zeroXTrade;

    uint256 private constant MAX_APPROVAL_AMOUNT = 2 ** 256 - 1;

    function initialize(IAugur _augur, IAugurTrading _augurTrading) public beforeInitialized returns (bool) {
        endInitialization();
        augur = _augur;
        gnosisSafeMasterCopy = _augur.lookup("GnosisSafe");
        require(gnosisSafeMasterCopy != address(0));
        proxyFactory = IProxyFactory(_augur.lookup("ProxyFactory"));
        proxyCodeHash = keccak256(proxyFactory.proxyRuntimeCode());
        deploymentData = keccak256(abi.encodePacked(proxyFactory.proxyCreationCode(), uint256(gnosisSafeMasterCopy)));
        cash = _augur.lookup("Cash");
        affiliates = augur.lookup("Affiliates");
        shareToken = augur.lookup("ShareToken");

        augurTrading = _augurTrading;
        createOrder = _augurTrading.lookup("CreateOrder");
        fillOrder = _augurTrading.lookup("FillOrder");
        zeroXTrade = _augurTrading.lookup("ZeroXTrade");
        return true;
    }

    // The misdirection here is because this is called through a delegatecall execution initially. We just direct that into making an actual call to the register method
    function setupForAugur(address _augur, address _createOrder, address _fillOrder, address _zeroXTrade, IERC20 _cash, IERC1155 _shareToken, IAffiliates _affiliates, bytes32 _fingerprint, address _referralAddress) public {
        _cash.approve(_augur, MAX_APPROVAL_AMOUNT);

        _cash.approve(_createOrder, MAX_APPROVAL_AMOUNT);
        _shareToken.setApprovalForAll(_createOrder, true);

        _cash.approve(_fillOrder, MAX_APPROVAL_AMOUNT);
        _shareToken.setApprovalForAll(_fillOrder, true);

        _cash.approve(_zeroXTrade, MAX_APPROVAL_AMOUNT);

        _affiliates.setFingerprint(_fingerprint);

        if (_referralAddress != address(0)) {
            _affiliates.setReferrer(_referralAddress);
        }
    }

    function proxyCreated(address _proxy, address _mastercopy, bytes calldata _initializer, uint256 _saltNonce) external {
        require(msg.sender == address(proxyFactory));
        IGnosisSafe _safe = IGnosisSafe(_proxy);
        bytes32 _codeHash;
        assembly {
            _codeHash := extcodehash(_safe)
        }
        require(_codeHash == proxyCodeHash, "Safe instance does not match expected code hash of the Proxy contract");
        require(_mastercopy == gnosisSafeMasterCopy, "Proxy master contract is not the Gnosis Safe");
        require(_safe.getThreshold() == 1, "Safe may only have a threshold of 1");
        address[] memory _owners = _safe.getOwners();
        require(_owners.length == 1, "Safe may only have 1 user");

        validateSafeCreation(_owners, _proxy, _mastercopy, _initializer, _saltNonce);

        address _owner = _owners[0];
        accountSafes[_owner].push(_safe);
        augurTrading.logGnosisSafeRegistered(address(_safe), _owner);
    }

    function validateSafeCreation(address[] memory _owners, address _proxy, address _mastercopy, bytes memory _initializer, uint256 _saltNonce) internal {
        IAffiliates _affilifates = IAffiliates(affiliates);
        bytes32 _fingerprint = _affilifates.getAccountFingerprint(_proxy);
        address _referralAddress = _affilifates.getReferrer(_proxy);
        bytes memory _expectedSetupForAugurData = abi.encodeWithSelector(this.setupForAugur.selector,
            address(augur),
            createOrder,
            fillOrder,
            zeroXTrade,
            cash,
            shareToken,
            affiliates,
            _fingerprint,
            _referralAddress
        );
        address _to;
        bytes memory _data;
        assembly {
            // Skip selector and length:
            let _initializerData := add(_initializer, 36)
            // Load the to and data params
            _to := mload(add(_initializerData, 64))
            _data := add(_initializerData, mload(add(_initializerData, 96)))
        }
        require(_to == address(this));
        // Requires the expected and actual data arguments are equal
        require(_data.length == _expectedSetupForAugurData.length && keccak256(_data) == keccak256(_expectedSetupForAugurData));
        // Now validate the address of the safe is what should have been produced via create2
        uint256 _saltNonceWithCallback = uint256(keccak256(abi.encodePacked(_saltNonce, address(this))));
        bytes32 _salt = keccak256(abi.encodePacked(keccak256(_initializer), _saltNonceWithCallback));
        address _expectedAddress = generateCreate2(address(proxyFactory), _salt, deploymentData);
        require(_proxy == _expectedAddress);
    }

    function generateCreate2(address _address, bytes32 _salt, bytes32 _hashedInitCode) public pure returns (address) {
        bytes1 _const = 0xff;
        return address(uint160(uint256(keccak256(abi.encodePacked(
            _const,
            _address,
            _salt,
            _hashedInitCode
        )))));
    }

    /**
     * @notice De-register the current Gnosis Safe for the msg.sender
     */
    function deRegister() external {
        IGnosisSafe _safe = IGnosisSafe(msg.sender);
        address[] memory _owners = _safe.getOwners();
        address _owner = _owners[0];
        require(_safe == getSafe(_owner), "Can only deRegister the current account safe");
        accountSafeIndexes[_owner] += 1;
        augurTrading.logGnosisSafeDeRegistered(address(_safe), _owner);
    }

    /**
     * @notice Get the registered Gnosis Safe for the given account
     * @param _account The account to look up
     * @return IGnosisSafe for the account or 0x if none is registered
     */
    function getSafe(address _account) public view returns (IGnosisSafe) {
        uint256 accountSafeIndex = accountSafeIndexes[_account];
        if (accountSafes[_account].length < accountSafeIndex + 1) {
            return IGnosisSafe(0);
        }
        return accountSafes[_account][accountSafeIndex];
    }
}
