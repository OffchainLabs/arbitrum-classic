pragma solidity 0.5.15;


contract IDaiVat {
    struct Ilk {
        uint256 Art;   // Total Normalised Debt     [wad]
        uint256 rate;  // Accumulated Rates         [ray]
        uint256 spot;  // Price with Safety Margin  [ray]
        uint256 line;  // Debt Ceiling              [rad]
        uint256 dust;  // Urn Debt Floor            [rad]
    }

    mapping (bytes32 => Ilk)                       public ilks;
    uint256 public live = 1;  // Access Flag

    mapping (address => uint256) public dai;  // [rad]
    mapping(address => mapping (address => uint)) public can;
    function hope(address usr) public;
    function move(address src, address dst, uint256 rad) public;
    function suck(address u, address v, uint rad) public;
    function frob(bytes32 i, address u, address v, address w, int dink, int dart) external;
    function faucet(address _target, uint256 _amount) public; // NOTE: this only exists in our mock VAT used for local testing
}