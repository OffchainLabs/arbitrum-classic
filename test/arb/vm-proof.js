const abi = require('ethereumjs-abi')

export default function(baseMachine, instruction, stackVals) {
	var proof = 
	baseMachine.instructionStack.hash() + 
	baseMachine.dataStack.hash().substring(2) + 
	baseMachine.callStack.hash().substring(2) + 
	baseMachine.registerVal.hash().substring(2) +
	baseMachine.staticVal.hash().substring(2) +
	abi.solidityPack(['uint8'], [instruction]).toString('hex');
	stackVals.forEach(function(val) {
		proof += val.serialize().toString('hex');
	});
	return proof;
}