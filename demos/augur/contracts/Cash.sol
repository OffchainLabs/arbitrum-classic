pragma solidity 0.5.15;

import './ICashFaucet.sol';
import './IAugur.sol';
import './ICash.sol';
import './libraries/ITyped.sol';
import './external/IDaiVat.sol';
import './external/IDaiJoin.sol';


/**
 * @title Cash
 * @dev Test contract for CASH (Dai)
 */
contract Cash is ITyped, ICash, ICashFaucet {
    using SafeMathUint256 for uint256;
    uint256 public constant ETERNAL_APPROVAL_VALUE = 2 ** 256 - 1;

    uint256 public totalSupply;
    
    event Mint(address indexed target, uint256 value);
    event Burn(address indexed target, uint256 value);

    mapping (address => uint) public wards;
    modifier auth {
        require(wards[msg.sender] == 1);
        _;
    }

    string constant public name = "Cash";
    string constant public symbol = "CASH";

    uint256 constant public RAY = 10 ** 27;

    mapping(address => uint) internal balances;
    mapping(address => mapping(address => uint256)) internal allowed;

    uint8 constant public decimals = 18;

    IDaiVat public daiVat;
    IDaiJoin public daiJoin;

    function initialize(IAugur _augur) public returns (bool) {
        daiJoin = IDaiJoin(_augur.lookup("DaiJoin"));
        daiVat = IDaiVat(_augur.lookup("DaiVat"));
        require(daiJoin != IDaiJoin(0));
        require(daiVat != IDaiVat(0));
        wards[address(this)] = 1;
        wards[address(daiJoin)] = 1;
        return true;
    }

    function transfer(address _to, uint256 _amount) public returns (bool) {
        require(_to != address(0), "Cannot send to 0x0");
        internalTransfer(msg.sender, _to, _amount);
        return true;
    }

    function transferFrom(address _from, address _to, uint256 _amount) public returns (bool) {
        uint256 _allowance = allowed[_from][msg.sender];
        require(_amount <= _allowance, "Not enough funds allowed");
        if (_allowance != ETERNAL_APPROVAL_VALUE) {
            allowed[_from][msg.sender] = _allowance.sub(_amount);
        }

        internalTransfer(_from, _to, _amount);
        return true;
    }

    function internalTransfer(address _from, address _to, uint256 _amount) internal returns (bool) {
        require(_to != address(0), "Cannot send to 0x0");
        require(balances[_from] >= _amount, "SEND Not enough funds");

        balances[_from] = balances[_from].sub(_amount);
        balances[_to] = balances[_to].add(_amount);
        emit Transfer(_from, _to, _amount);
        return true;
    }

    function balanceOf(address _owner) public view returns (uint256) {
        return balances[_owner];
    }

    function approve(address _spender, uint256 _amount) public returns (bool) {
        approveInternal(msg.sender, _spender, _amount);
        return true;
    }

    function increaseApproval(address _spender, uint _addedValue) public returns (bool) {
        approveInternal(msg.sender, _spender, allowed[msg.sender][_spender].add(_addedValue));
        return true;
    }

    function decreaseApproval(address _spender, uint _subtractedValue) public returns (bool) {
        uint oldValue = allowed[msg.sender][_spender];
        if (_subtractedValue > oldValue) {
            approveInternal(msg.sender, _spender, 0);
        } else {
            approveInternal(msg.sender, _spender, oldValue.sub(_subtractedValue));
        }
        return true;
    }

    function approveInternal(address _owner, address _spender, uint256 _allowance) internal returns (bool) {
        allowed[_owner][_spender] = _allowance;
        emit Approval(_owner, _spender, _allowance);
        return true;
    }

    function allowance(address _owner, address _spender) public view returns (uint256) {
        return allowed[_owner][_spender];
    }

    function faucet(uint256 _amount) public returns (bool) {
        daiVat.faucet(address(daiJoin), _amount * RAY);
        mint(msg.sender, _amount);
        return true;
    }

    function sub(uint x, uint y) internal pure returns (uint z) {
        require((z = x - y) <= x, "math-sub-underflow");
    }

    function joinMint(address usr, uint wad) public auth returns (bool) {
        return mint(usr, wad);
    }

    function joinBurn(address usr, uint wad) public returns (bool) {
        if (usr != msg.sender && allowed[usr][msg.sender] != uint(-1)) {
            allowed[usr][msg.sender] = sub(allowed[usr][msg.sender], wad);
        }
        return burn(usr, wad);
    }

    function mint(address _target, uint256 _amount) internal returns (bool) {
        balances[_target] = balances[_target].add(_amount);
        totalSupply = totalSupply.add(_amount);
        emit Mint(_target, _amount);
        emit Transfer(address(0), _target, _amount);
        return true;
    }

    function burn(address _target, uint256 _amount) internal returns (bool) {
        require(balanceOf(_target) >= _amount, "BURN Not enough funds");

        balances[_target] = balances[_target].sub(_amount);
        totalSupply = totalSupply.sub(_amount);

        emit Burn(_target, _amount);
        return true;
    }

    function getTypeName() public view returns (bytes32) {
        return "Cash";
    }
}