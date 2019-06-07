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

from .struct import Struct


class Locals:
    def __init__(self, vm, fields):
        self.vm = vm
        self.struc = Struct(
            "Locals[{}]".format(', '.join((str(x) for x in fields))),
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
