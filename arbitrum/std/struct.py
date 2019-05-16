from .array import Array
from .. import value
from ..annotation import modifies_stack


class Struct:

    # Makes a structmap object describing a struct
    # with the specified field names
    def __init__(self, name, fieldNames):
        self.typ = value.NamedType(name)
        self.field_index = {}
        types = []
        for i, val in enumerate(fieldNames):
            if isinstance(val, tuple):
                types.append(val[1])
                self.field_index[val[0]] = i
            else:
                types.append(value.ValueType())
                self.field_index[val] = i

        self.array = Array(types)

    def __len__(self):
        return len(self.field_index)

    def update_type(self, name, typ):
        self.array.update_type(self.field_index[name], typ)

    def make(self):
        return self.array.make()

    @property
    def new(self):
        @modifies_stack(0, [self.typ])
        def new(vm):
            vm.push(self.make())
            vm.cast(self.typ)

        return new

    @property
    def build(self):
        @modifies_stack(self.array.types, [self.typ], self.typ.name)
        def build(vm):
            self.array.build(vm)
            vm.cast(self.typ)

        return build

    def get(self, names):
        if not isinstance(names, list):
            names = [names]
        types = [
            self.array.types[self.field_index[name]]
            for name in names
        ]

        def binder(struct_type, field_names):
            @modifies_stack(
                [struct_type],
                types,
                f"{self.typ.name}_{'_'.join(field_names)}"
            )
            def get(vm):
                vm.cast(self.array.typ)
                for name in field_names[1:][::-1]:
                    vm.dup0()
                    self.array.get(self.field_index[name])(vm)
                    vm.swap1()
                self.array.get(self.field_index[field_names[0]])(vm)
            return get
        return binder(self.typ, names)

    def set_val(self, names):
        if not isinstance(names, list):
            names = [names]
        types = [
            self.array.types[self.field_index[name]]
            for name in names
        ]

        if len(self) == 1:
            @modifies_stack(
                [self.array.types[self.field_index[names[0]]]],
                [self.typ],
                f"{self.typ.name}_{names[0]}"
            )
            def set_val(vm):
                vm.cast(self.array.typ)
                self.array.set_val(self.field_index[names[0]])(vm)
                vm.cast(self.typ)
            return set_val

        def binder(struct_type, field_names):
            @modifies_stack(
                [struct_type] + types,
                [struct_type],
                f"{self.typ.name}_{'_'.join(field_names)}"
            )
            def set_val(vm):
                vm.cast(self.array.typ)
                for name in field_names:
                    self.array.set_val(self.field_index[name])(vm)
                vm.cast(struct_type)
            return set_val
        return binder(self.typ, names)
