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

from setuptools import setup

setup(name='arbc-solidity',
      version='0.1.0',
      description='Compiler from solidity to AVM bytecode',
      url='http://github.com/OffchainLabs/arbc-solidity',
      author='Harry Kalodner',
      author_email='harry@offchainlabs.com',
      license='Apache-2.0',
      packages=['arbitrum', 'arbitrum.std', 'arbitrum.evm'],
      zip_safe=False,
      scripts=['bin/arbc-truffle-compile', 'bin/arbc-compile'],
      test_suite='nose.collector',
      tests_require=['nose'],
      python_requires='>=3.6',
      install_requires=[
          "eth-utils",
          "eth-abi>=2.0.0b0",
          "pyevmasm",
          "web3>=5.0.0b1",
          "py-solc-x",
          "networkx"
      ]
)
