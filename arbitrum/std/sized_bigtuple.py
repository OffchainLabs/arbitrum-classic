from . import bigtuple
from . import sized_common
from ..annotation import modifies_stack
from .. import value
from .struct import Struct

sized_bigtuple = Struct("sized_bigtuple", [
    ("data", bigtuple.typ),
    ("size", value.IntType())
])

def make():
    return sized_common.make(bigtuple.make)


# [] -> [tuple]
@modifies_stack(0, 1)
def new(vm):
    sized_common.new(vm, bigtuple.new)


@modifies_stack(1, 1)
def length(vm):
    sized_bigtuple.get("size")(vm)


# [tuple, index, value] -> [tuple]
@modifies_stack(3, 1)
def set_val(vm):
    sized_common.set_val(vm, sized_bigtuple, bigtuple.set_val, 1)


# [tuple, index] -> [value]
@modifies_stack(2, 1)
def get(vm):
    sized_common.get(vm, sized_bigtuple, bigtuple.get)


def get_static(val, index):
    return sized_common.get_static(val, index, bigtuple.get_static)


def set_static(val, index, value):
    return sized_common.set_static(val, index, value, bigtuple.set_static, 1)
