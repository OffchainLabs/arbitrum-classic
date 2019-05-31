const ethers = require('ethers');
const ArbProvider = require('arb-ethers-provider');
var ProviderBridge = require('./ethers-web3-bridge');

module.exports = function(managerUrl, contracts, provider) {
	let wrappedProv = new ethers.providers.Web3Provider(provider)
	let arbProvider = new ArbProvider(managerUrl, contracts, wrappedProv)
	let wallet = arbProvider.getSigner(0);
	return new ProviderBridge(arbProvider, wallet);
}
