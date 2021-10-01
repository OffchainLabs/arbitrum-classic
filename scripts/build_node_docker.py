#!/usr/bin/env python3

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
from support.run import run

ROOT_DIR = os.path.abspath(os.path.dirname(os.path.dirname(os.path.abspath(__file__))))
DOCKERFILE_CACHE = """FROM alpine:3.9
RUN mkdir /build /cpp-build /rocksdb
FROM scratch
COPY --from=0 /cpp-build /cpp-build
COPY --from=0 /build /build"""


# Bootstrap the build cache if it does not exist
def bootstrap_build_cache(name, sudo_flag):
    if (
        run(
            "docker images -q %s" % name,
            capture_stdout=True,
            quiet=True,
            sudo=sudo_flag,
        )
        == ""
    ):
        run("mkdir -p .tmp")
        run('echo "%s" > .tmp/Dockerfile' % DOCKERFILE_CACHE)
        run("docker build -t %s .tmp" % name, sudo=sudo_flag)
        run("rm -rf .tmp")


def build_node(sudo_flag=False):
    node_root = os.path.abspath(os.path.join(ROOT_DIR, "packages"))
    bootstrap_build_cache("arb-avm-cpp", sudo_flag)
    bootstrap_build_cache("arb-node", sudo_flag)

    run("git -C %s submodule update --init --recursive" % ROOT_DIR)

    return run(
        "docker build -t arb-node -f %s/arb-node.Dockerfile %s"
        % (node_root, node_root),
        sudo=sudo_flag,
    )


def is_built(sudo_flag=False):
    layer = run(
        "docker create arb-node", capture_stdout=True, quiet=True, sudo=sudo_flag
    ).strip()
    return layer != ""


if __name__ == "__main__":
    try:
        build_node()
    except KeyboardInterrupt:
        exit(1)
