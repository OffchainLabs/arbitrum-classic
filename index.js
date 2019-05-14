var ArbProvider = require('./lib/index.js');
var ProviderBridge = require('ethers-web3-bridge');

module.exports = function(managerUrl, contracts, provider) {
	let arbProvider = new ArbProvider(managerUrl, contracts, provider)
	let wallet = arbProvider.getSigner(0);
	return new ProviderBridge(arbProvider, wallet);
}
