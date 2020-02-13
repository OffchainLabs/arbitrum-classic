var Augur = artifacts.require("./Augur.sol");
var AugurTrading = artifacts.require("./AugurTrading.sol");

module.exports = async function(deployer) {
  await deployer.deploy(Augur);
  await deployer.deploy(AugurTrading, Augur.address);
};
