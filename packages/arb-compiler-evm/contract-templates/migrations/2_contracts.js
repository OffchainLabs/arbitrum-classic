const ArbERC20 = artifacts.require("ArbERC20");
const ArbERC721 = artifacts.require("ArbERC721");

module.exports = function(deployer) {
  deployer.deploy(ArbERC20);
  deployer.deploy(ArbERC721);
};
