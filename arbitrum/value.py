import eth_utils
import eth_abi

INT_TYPE_CODE = 0
CODE_POINT_CODE = 1
HASH_ONLY_CODE = 2
TUPLE_TYPE_CODE = 3


class IntType:
    def __repr__(self):
        return "IntType()"

    def empty_val(self):
        return 0

    def typecode(self):
        return 0

    def accepts(self, val):
        return isinstance(val, (IntType, int))

    def common(self, other):
        if isinstance(other, (IntType, int)):
            return IntType()

        return ValueType()


class TupleType:
    def __init__(self, types=None):
        if types and isinstance(types, int):
            types = [ValueType()]*types
        self.types = types

    def __repr__(self):
        if self.types is None:
            return "TupleType()"

        return f"TupleType({', '.join([repr(typ) for typ in self.types])})"

    def empty_val(self):
        return Tuple([])

    def typecode(self):
        return 2

    def size(self):
        if self.types is None:
            return IntType()

        return len(self.types)

    def get_tup(self, index):
        if self.types is None:
            raise Exception("Can't get from uncounted tup")

        if isinstance(index, int):
            if index >= len(self.types):
                raise Exception("TypeCheck: Can't get from tuple that is too small")
            return self.types[index]

        return ValueType()

    def set_tup_val(self, index, value):
        if self.types is None:
            raise Exception("Can't set to uncounted tup")

        if isinstance(index, int):
            if index < len(self.types):
                new_types = list(self.types)
                new_types[index] = value
                return TupleType(new_types)
            else:
                raise Exception("TypeCheck: Can't set to tuple that is too small")

        return TupleType(len(self.types))

    def has_member_at_index(self, index):
        return True

    def accepts(self, val):
        if not isinstance(val, (TupleType, Tuple)):
            return False

        if self.types is None:
            return True

        if isinstance(val, TupleType):
            if val.types is None:
                return False
            return all(typeA.accepts(typeB) for (typeA, typeB) in zip(self.types, val.types))
        
        if isinstance(val, Tuple):
            return all(typeA.accepts(typeB) for (typeA, typeB) in zip(self.types, val.val))

        return False

    def common(self, other):
        if isinstance(other, TupleType):
            if self.types is None or other.types is None:
                return TupleType()

            if self.types == other.types:
                return TupleType(list(self.types))

            return TupleType()
        return ValueType()


class CodePointType:
    def __repr__(self):
        return "CodePointType()"

    def empty_val(self):
        return AVMCodePoint(-2, None, b'')

    def typecode(self):
        return 1

    def accepts(self, val):
        return isinstance(val, CodePointType)

    def common(self, other):
        if isinstance(other, CodePointType):
            return CodePointType()
        return ValueType()


class NamedType:

    def __init__(self, name):
        self.name = name

    def __repr__(self):
        return f"Struct({self.name})"

    def common(self, other):
        if isinstance(other, NamedType):
            if self.name == other.name:
                return self
        return ValueType()

    def accepts(self, val):
        if isinstance(val, NamedType):
            return self.name == val.name
        return False


class ValueType:
    def __repr__(self):
        return "ValueType()"

    def empty_val(self):
        return Tuple([])

    def accepts(self, val):
        return True

    def common(self, other):
        return ValueType()


class TypeStack:

    def __init__(self, stack=None, auxstack=None):
        if stack is None:
            stack = []
        if auxstack is None:
            auxstack = []
        self.stack = stack
        self.auxstack = auxstack

    def __len__(self):
        return len(self.stack)

    def __repr__(self):
        return f"TypeStack({self.stack}, {self.auxstack})"

    def __getitem__(self, index):
        return self.stack[index]

    def clone(self):
        return TypeStack(list(self.stack), list(self.auxstack))

    def merge(self, other):
        ret = []
        if len(self) != len(other):
            print(self[:])
            print(other[:])
            raise Exception(f"Can't merge stack's of different length {len(self)} and {len(other)}")
        for a_type, b_type in zip(self.stack, other.stack):
            ret.append(arbtype(a_type).common(arbtype(b_type)))
        self.stack = ret

    def pop(self, pop_type=None):
        if pop_type is None:
            pop_type = ValueType()
        typ = self.stack[0]
        try:
            if not pop_type.accepts(typ):
                raise Exception(f"TypeStack wanted {pop_type} but got {self.stack[0]}")
        except Exception as err:
            raise Exception(f"TypeStack: included non-type {typ}. Got err {err}")
        self.stack = self.stack[1:]
        return typ

    def pop_aux(self, pop_type=None):
        if pop_type is None:
            pop_type = ValueType()
        typ = self.auxstack[0]
        try:
            if not pop_type.accepts(typ):
                raise Exception(f"TypeStack wanted {pop_type} but got {self.stack[0]}")
        except Exception as err:
            raise Exception(f"TypeStack: included non-type {typ}. Got err {err}")
        self.auxstack = self.auxstack[1:]
        return typ

    def push(self, push_type):
        self.stack.insert(0, push_type)

    def push_aux(self, push_type):
        self.auxstack.insert(0, push_type)


class Tuple:
    def __init__(self, val=None):
        if val is None:
            val = []
        if not isinstance(val, list):
            raise Exception(f"Tuple must be created from list not {type(val)}")
        elif len(val) > 8:
            raise Exception("Tuple must be created from list of size <= 8")
        self.val = tuple(val)

    def __repr__(self):
        return f"Tuple([{', '.join([repr(v) for v in self.val])}])"

    def __len__(self):
        return len(self.val)

    def size(self):
        return len(self)

    def __getitem__(self, index):
        return self.val[index]

    def __eq__(self, other):
        return isinstance(other, Tuple) and self.val == other.val

    def __hash__(self):
        return self.val.__hash__()

    def __ne__(self, other):
        if not isinstance(other, Tuple):
            return False
        return self.val != other.val

    def __iter__(self):
        return self.val.__iter__()

    def has_member_at_index(self, index):
        return index < len(self.val)

    def get_tup(self, index):
        return self.val[index]

    def set_tup_val(self, index, value):
        if index >= len(self):
            raise Exception(f"Can't set value {value} to index {index} of tuple {self.val}")
        new_tup = list(self.val)
        new_tup[index] = value
        return Tuple(new_tup)


class AVMCodePoint:
    def __init__(self, pc, op, next_hash, path=None):
        if path is None:
            path = []
        self.pc = pc
        self.op = op
        self.next_hash = next_hash
        self.path = path

    def __repr__(self):
        return f"AVMCodePoint({self.pc}, {self.op})"


def value_hash(val):
    if isinstance(val, int):
        return eth_utils.keccak(eth_abi.encode_single(
            '(uint256)',
            [val]
        ))
        # return eth_utils.keccak(eth_abi.encode_single(
        #     '(uint8,uint256)',
        #     [INT_TYPE_CODE, val]
        # ))
    if isinstance(val, Tuple):
        return eth_utils.keccak(eth_abi.encode_single(
            '(uint8' + ',bytes32'*len(val) + ')',
            [TUPLE_TYPE_CODE] + [value_hash(v) for v in val.val]
        ))
    if isinstance(val, AVMCodePoint):
        if hasattr(val.op, "op_code"):
            return eth_utils.keccak(eth_abi.encode_single(
                '(uint8,uint8,bytes32)',
                [2, val.op.op_code, val.next_hash]
            ))
        if hasattr(val.op, "val"):
            return eth_utils.keccak(eth_abi.encode_single(
                '(uint8,uint8,bytes32,bytes32)',
                [2, val.op.op.op_code, value_hash(val.op.val), val.next_hash]
            ))
        raise Exception(f"Bad op type {val.op}")

    raise Exception(f"Can't hash {val}")


def arbtype(val):
    if isinstance(val, int):
        return IntType()
    if isinstance(val, Tuple):
        return TupleType([arbtype(v) for v in val])
    return val
