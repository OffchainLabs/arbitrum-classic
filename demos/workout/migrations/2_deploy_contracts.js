var Workout = artifacts.require("Workout");

module.exports = function(deployer) {
  deployer.deploy(Workout);
};
