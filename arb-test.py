import arbitrum as arb


def test(vm):
    for i in range(10000):
        vm.push(i)
    for _ in range(9999):
        vm.add()
    vm.push(5)
    vm.push(arb.value.Tuple([1, 2, 3, 4]))
    vm.tsetn(1)
    vm.tgetn(3)
    vm.halt()


code = arb.compile_block(test)
vm = arb.compile_program(arb.ast.BlockStatement([]), code)
# print(vm.code)
with open("test.ao", "wb") as f:
    arb.marshall.marshall_vm(vm, f)

# vm2 = arb.VM()
# test(vm2)
# print(vm2.stack[:])