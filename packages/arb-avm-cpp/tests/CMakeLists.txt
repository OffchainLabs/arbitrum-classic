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

cmake_minimum_required (VERSION 3.5)
set (CMAKE_CXX_STANDARD 14)

project (avm_tests)

set(TEST_CASES_PATH ${CMAKE_CURRENT_BINARY_DIR}/test_cases.json)

configure_file (
    "${CMAKE_CURRENT_SOURCE_DIR}/config.hpp.in"
    "${CMAKE_CURRENT_SOURCE_DIR}/config.hpp"
)

configure_file(test_cases.json ${CMAKE_CURRENT_BINARY_DIR}/test_cases.json COPYONLY)

add_executable(avm_tests main.cpp opcodes.cpp value.cpp balanceTracker.cpp)

target_link_libraries(
	avm_tests
PRIVATE
	avm
    avm_utils
	CONAN_PKG::Catch2
	CONAN_PKG::jsonformoderncpp
)

include(CTest)
include(lib/cmake/Catch2/Catch)
catch_discover_tests(avm_tests)

target_code_coverage(avm_tests AUTO ALL EXCLUDE external/* tests/*)
