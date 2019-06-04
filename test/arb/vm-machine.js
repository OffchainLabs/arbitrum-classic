const utils = require('ethereumjs-util');
const abi = require('ethereumjs-abi');

export class Stack {
	constructor(items) {
		this.items = items;
	}

	hash() {
		let stackHash = '0x0000000000000000000000000000000000000000000000000000000000000000';
		for (var i = 0; i < this.items.length; i++) {
			stackHash = '0x' + abi.soliditySHA3(['uint8', 'bytes32', 'bytes32'], [38, this.items[i].hash(), stackHash]).toString('hex');
		}
		return stackHash;
	}

	push(item) {
		this.items.push(item);
	}

	pop() {
		this.items.pop();
	}
}

export class Machine {
	constructor(instructionStack, dataStack, callStack, registerVal, staticVal) {
		this.instructionStack = instructionStack;
		this.dataStack = dataStack;
		this.callStack = callStack;
		this.registerVal = registerVal;
		this.staticVal = staticVal;
	}

	hash() {
		return '0x' + abi.soliditySHA3(
			['bytes32', 'bytes32', 'bytes32', 'bytes32', 'bytes32'], 
			[
				this.instructionStack.hash(), 
				this.dataStack.hash(),
				this.callStack.hash(),
				this.registerVal.hash(),
				this.staticVal.hash()
			]).toString('hex');
	}
}
