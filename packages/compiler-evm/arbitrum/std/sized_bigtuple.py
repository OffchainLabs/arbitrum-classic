# Copyright 2019, Offchain Labs, Inc.
# 
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

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
