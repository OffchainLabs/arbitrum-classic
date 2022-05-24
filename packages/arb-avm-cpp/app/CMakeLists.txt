# Copyright 2019, Offchain Labs, Inc.

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at

#    http://www.apache.org/licenses/LICENSE-2.0

# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

cmake_minimum_required (VERSION 3.9)

set (CMAKE_CXX_STANDARD 17)

project (avm_runner LANGUAGES CXX)
add_executable(avm_runner main.cpp)
target_link_libraries(avm_runner PRIVATE avm data_storage Boost::filesystem)
target_code_coverage(avm_runner AUTO ALL EXCLUDE external/* tests/*)

add_executable(arbcore_runner arbcorerunner.cpp)
target_link_libraries(arbcore_runner PRIVATE avm data_storage Boost::filesystem)
target_code_coverage(arbcore_runner AUTO ALL EXCLUDE external/* tests/*)
