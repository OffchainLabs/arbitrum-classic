from . import tup
from .struct import Struct
from ..annotation import modifies_stack
from ..value import TupleType, IntType, ValueType
from .. import value


def make_bigtuple_type(typ, default_val=None):
    base_typ = value.TupleType([
        value.TupleType(),
        value.TupleType(),
        value.TupleType(),
        value.TupleType(),
        value.TupleType(),
        value.TupleType(),
        value.TupleType(),
        typ
    ])
    bigtuple_type = Struct(f"bigtuple[{typ}]", [
        ("tup", base_typ)
    ])

    if default_val is None:
        default_val = value.Tuple([])
    base_val = value.Tuple([
        value.Tuple([]),
        value.Tuple([]),
        value.Tuple([]),
        value.Tuple([]),
        value.Tuple([]),
        value.Tuple([]),
        value.Tuple([]),
        default_val
    ])

    class BigTuple:
        @staticmethod
        def make():
            return base_val

        @staticmethod
        @modifies_stack([], [bigtuple_type.typ], default_val)
        def new(vm):
            vm.push(BigTuple.make())
            vm.cast(bigtuple_type)

        @staticmethod
        @modifies_stack([bigtuple_type.typ, IntType(), typ], [bigtuple_type.typ], default_val)
        def set_val(vm):
            bigtuple_type.get("tup")(vm)
            BigTuple._set_val_impl(vm)
            bigtuple_type.set_val("tup")(vm)

        # [tuple, index, value] -> [tuple]
        @staticmethod
        @modifies_stack(
            [TupleType(), IntType(), typ],
            [base_typ],
            default_val
        )
        def _set_val_impl(vm):
            # check if tuple is empty
            vm.dup0()
            vm.tnewn(0)
            vm.eq()

            vm.ifelse(lambda vm: [vm.pop(), vm.push(base_val)])
            # tuple must be size 8
            vm.cast(base_typ)

            vm.swap1()
            vm.dup0()
            vm.push(0)
            vm.eq()

            # check if index is 0
            vm.ifelse(
                # [index, tuple, value]
                lambda vm: [
                    vm.pop(),
                    vm.tsetn(7),
                    vm.cast(base_typ),
                ], lambda vm: [
                    vm.swap2(),
                    vm.auxpush(),
                    vm.swap1(),

                    vm.push(7),
                    vm.dup1(),
                    vm.mod(),

                    vm.swap1(),
                    vm.push(7),
                    vm.swap1(),
                    vm.div(),
                    # [index / 7, index % 7, tuple]

                    vm.dup2(),
                    vm.dup2(),
                    vm.tget(),
                    # [tuple[index % 7], index / 7, index % 7, tuple]
                    vm.cast(value.TupleType()),
                    vm.swap1(),
                    vm.auxpop(),
                    vm.swap2(),
                    BigTuple._set_val_impl(vm),
                    # [tup, index % 7, tuple]
                    vm.swap2(),
                    vm.swap1(),
                    vm.tset(),
                    vm.cast(base_typ)
                ]
            )

        # [tuple, index] -> [tuple]
        @staticmethod
        def read_modify_write(closure):
            # [tuple, index] -> [tuple]
            @modifies_stack(
                [TupleType(), IntType(), closure.typ],
                [base_typ],
                f"{default_val}_{closure.typ.name}"
            )
            def read_modify_write_impl(vm):
                # check if tuple is empty
                vm.dup0()
                vm.tnewn(0)
                vm.eq()

                vm.ifelse(lambda vm: [vm.pop(), vm.push(base_val)])
                # tuple must be size 8
                vm.cast(base_typ)

                vm.swap1()
                vm.dup0()
                vm.push(0)
                vm.eq()

                # check if index is 0
                vm.ifelse(
                    # [index, tuple, closure]
                    lambda vm: [
                        vm.pop(),
                        vm.swap1(),
                        vm.dup1(),
                        vm.tgetn(7),
                        vm.cast(typ),
                        vm.swap1(),
                        closure.call(vm),
                        vm.swap1(),
                        vm.tsetn(7),
                        vm.cast(base_typ)
                    ], lambda vm: [
                        vm.swap2(),
                        vm.auxpush(),
                        vm.swap1(),

                        vm.push(7),
                        vm.dup1(),
                        vm.mod(),

                        vm.swap1(),
                        vm.push(7),
                        vm.swap1(),
                        vm.div(),
                        # [index / 7, index % 7, tuple]

                        vm.dup2(),
                        vm.dup2(),
                        vm.tget(),
                        # [tuple[index % 7], index / 7, index % 7, tuple]
                        vm.cast(value.TupleType()),
                        vm.swap1(),
                        vm.auxpop(),
                        vm.swap2(),
                        read_modify_write_impl(vm),
                        # [tup, index % 7, tuple]
                        vm.swap2(),
                        vm.swap1(),
                        vm.tset(),
                        vm.cast(base_typ)
                    ]
                )

            @modifies_stack(
                [bigtuple_type.typ, IntType(), closure.typ],
                [bigtuple_type.typ],
                f"{default_val}_{closure.typ.name}"
            )
            def read_modify_write(vm):
                bigtuple_type.get("tup")(vm)
                read_modify_write_impl(vm)
                bigtuple_type.set_val("tup")(vm)

            return read_modify_write

        # [tuple, index] -> [value]
        @staticmethod
        @modifies_stack([bigtuple_type.typ, IntType()], [typ], default_val)
        def get(vm):
            # tuple, index
            vm.while_loop(lambda vm: [
                vm.dup0(),
                vm.tnewn(0),
                vm.eq(),
                vm.dup2(),
                vm.push(0),
                vm.eq(),
                vm.bitwise_or(),
                vm.iszero()
            ], lambda vm: [
                vm.cast(base_typ),
                # tuple index
                vm.push(7),
                vm.dup2(),
                vm.div(),
                # index/7 tuple index

                vm.swap2(),
                vm.push(7),
                vm.swap1(),
                vm.mod(),
                # index/7 tuple index%7
                vm.tget()
            ])
            vm.swap1()
            vm.push(0)
            vm.eq()
            # found_tup, tup
            vm.dup1()
            vm.tnewn(0)
            vm.eq()
            vm.iszero()
            vm.bitwise_and()
            vm.ifelse(lambda vm: [
                vm.cast(base_typ),
                vm.tgetn(7),
            ], lambda vm: [
                vm.pop(),
                vm.push(default_val),
                vm.cast(typ)
            ])

        # [source_tuple, start offset, end offset, dest tuple, dest offset]
        @staticmethod
        @modifies_stack(5, 1, default_val)
        def copy(vm):
            tup.make(5)(vm)
            BigTuple._copy_impl(vm)
            vm.tgetn(3)

        @staticmethod
        @modifies_stack(1, 1, default_val)
        def _copy_impl(vm):
            vm.dup0()
            vm.tgetn(2)
            vm.dup1()
            vm.tgetn(1)
            vm.lt()
            vm.ifelse(
                lambda vm: [
                    # get source[in offset]
                    vm.dup0(),
                    vm.tgetn(1),
                    vm.dup1(),
                    vm.tgetn(0),
                    BigTuple.get(vm),

                    # get destination[out offset]
                    vm.dup1(),
                    vm.tgetn(4),
                    vm.dup2(),
                    vm.tgetn(3),
                    BigTuple.set_val(vm),
                    vm.swap1(),
                    vm.tsetn(3),

                    # increment destination offset
                    vm.dup0(),
                    vm.tgetn(4),
                    vm.push(1),
                    vm.add(),
                    vm.swap1(),
                    vm.tsetn(4),

                    # increment source offset
                    vm.dup0(),
                    vm.tgetn(1),
                    vm.push(1),
                    vm.add(),
                    vm.swap1(),
                    vm.tsetn(1),

                    BigTuple._copy_impl(vm)
                ]
            )

        # [source_tuple, start offset, end offset]
        @staticmethod
        @modifies_stack(3, 1, default_val)
        def get_subset(vm):
            tup.make(3)(vm)
            BigTuple.new(vm)
            vm.push(0)
            vm.swap2()
            tup.tbreak(3)(vm)
            BigTuple.copy(vm)

        @staticmethod
        def get_static(tup_val, index):
            while tup_val != value.Tuple([]) and index != 0:
                tup_val = tup_val[index % 7]
                index //= 7

            if tup_val == value.Tuple([]):
                return 0

            return tup_val[7]

        @staticmethod
        def set_static(tup_val, index, val):
            if tup_val == value.Tuple([]):
                tup_val = base_val

            if index == 0:
                return tup_val.set_tup_val(7, val)

            return tup_val.set_tup_val(
                index % 7,
                BigTuple.set_static(tup_val[index % 7], index // 7, val)
            )

        @staticmethod
        def fromints(data):
            ret = value.Tuple([])
            for i, val in enumerate(data):
                ret = BigTuple.set_static(ret, i, val)
            return ret

    BigTuple.typ = bigtuple_type.typ
    return BigTuple


bigtuple = make_bigtuple_type(value.ValueType())
bigtuple_int = make_bigtuple_type(value.IntType(), 0)

