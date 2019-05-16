
from .import tup
from ..annotation import modifies_stack
from .locals import Locals

@modifies_stack(1, 1)   # seed -> generator
def new(vm):
	vm.hash()

@modifies_stack(1, 2)   # gen -> value gen
def getint(vm): 
	vm.push(1)
	vm.dup1()
	tup.make(2)(vm)
	vm.hash()
	# newgen oldgen
	vm.swap1()
	vm.push(0)
	vm.swap1()
	tup.make(2)(vm)
	vm.hash()

@modifies_stack(2, 2)   # gen n -> value gen
def getmodn(vm):   # get a random int, 0<=result<n
	local_vars = Locals(vm, ['cutoff', 'dummy'])
	# gen n
	vm.dup1()
	vm.push((1<<256)-1)
	vm.div()
	# ff//n gen n
	vm.dup2()
	vm.mul()
	vm.dup0()
	# cutoff dummy gen n
	local_vars.make()

	# gen n
	getint(vm)
	# val gen n
	vm.while_loop(lambda vm: [
		vm.dup0(),
		local_vars.get('cutoff'),
		# cutoff val val gen n
		vm.gt(),
		vm.iszero(),
	], lambda vm: [
		# val gen n
		vm.pop(),
		getint(vm),
		])
	# can now discard cutoff from auxstack
	local_vars.discard()

	# val gen n
	vm.swap1()
	vm.swap2()
	# n val gen
	vm.swap1()
	vm.mod()
	# value gen
