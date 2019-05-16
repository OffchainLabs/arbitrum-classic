from .struct import Struct


class Locals:
    def __init__(self, vm, fields):
        self.vm = vm
        self.struc = Struct(
            f"Locals[{', '.join((str(x) for x in fields))}",
            fields
        )

    def new(self):
        self.struc.new(self.vm)
        self.vm.auxpush()

    def make(self):
        self.struc.build(self.vm)
        self.vm.auxpush()

    def get(self, fields):
        self.vm.auxpop()
        self.vm.dup0()
        self.vm.auxpush()
        self.struc.get(fields)(self.vm)

    def set_val(self, fields):
        self.vm.auxpop()
        self.struc.set_val(fields)(self.vm)
        self.vm.auxpush()

    def discard(self, varToSave=None):
        self.vm.auxpop()
        if varToSave == None:
            self.vm.pop()
        else:
            self.struc.get(varToSave)(self.vm)
