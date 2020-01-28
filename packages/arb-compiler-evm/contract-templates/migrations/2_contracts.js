const ArbERC20 = artifacts.require("ArbERC20");
const ArbERC721 = artifacts.require("ArbERC721");
const ArbInfo = artifacts.require("ArbInfo");

module.exports = function(deployer) {
  deployer.deploy(ArbERC20);
  deployer.deploy(ArbERC721);
  deployer.deploy(ArbInfo);
};
