pragma solidity ^0.5.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/ERC20Detailed.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721Full.sol";
import "@openzeppelin/contracts/drafts/Counters.sol";


contract TestToken is ERC20, ERC20Detailed {
    constructor(uint256 initialSupply) ERC20Detailed("Gold", "GLD", 18) public {}

    function mint(address dest, uint amount) public {
        _mint(dest, amount);
    }

}


contract TestItem is ERC721Full {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;

    constructor() ERC721Full("GameItem", "ITM") public {

    }

    function tokensOfOwner(address owner) public view returns (uint256[] memory) {
        return _tokensOfOwner(owner);
    }

    function mintItem(address player, uint itemId) public {
        _mint(player, itemId);
    }
}
