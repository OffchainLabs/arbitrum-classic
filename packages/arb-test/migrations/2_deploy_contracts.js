const Fibonacci = artifacts.require("Fibonacci");

module.exports = async function(deployer) {
  await deployer.deploy(Fibonacci);
};
