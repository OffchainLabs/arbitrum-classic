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

import os
import subprocess


# Run commands in shell
def run(command, sudo=False, capture_stdout=False, quiet=False):
    command = ("sudo " if sudo else "") + command
    if not quiet:
        print("\n\033[1m$ %s" % command + "\033[0m")
    if not capture_stdout:
        return os.system(command)
    try:
        return subprocess.check_output(command, shell=True).decode("utf-8")
    except subprocess.CalledProcessError as err:
        print("Got subprocess error", err.output.decode("utf-8"))
        return ""
