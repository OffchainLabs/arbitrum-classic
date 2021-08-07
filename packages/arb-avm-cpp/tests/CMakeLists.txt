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
set (CMAKE_CXX_STANDARD 17)

project (avm_tests)

set(TEST_CASES_PATH ${CMAKE_CURRENT_SOURCE_DIR}/test_cases.json)
set(TEST_CONTRACT_PATH ${CMAKE_CURRENT_SOURCE_DIR}/contract.mexe)
set(ARB_OS_DIR ${CMAKE_CURRENT_SOURCE_DIR}/../../arb-os)
set(ARB_OS_PATH ${ARB_OS_DIR}/arb_os/arbos.mexe)
set(ARB_OS_TEST_CASES_PATH ${ARB_OS_DIR}/replayTests)
set(MACHINE_TEST_CASES_PATH ${CMAKE_CURRENT_SOURCE_DIR}/machine-cases)

configure_file (
    "${CMAKE_CURRENT_SOURCE_DIR}/config.hpp.in"
    "${CMAKE_CURRENT_SOURCE_DIR}/config.hpp"
)

add_executable(
		avm_tests

		aggregator.cpp
		arbos.cpp
		arbcore.cpp
		buffer.cpp
		checkpoint.cpp
		checkpointedmachine.cpp
		code.cpp
		datastack.cpp
		ecops.cpp
		ecops.hpp
		helper.cpp
        machine.cpp
		main.cpp
        messagestore.cpp
		opcodes.cpp
		sideloadcache.cpp
		value.cpp
        valuecache.cpp)

target_compile_options(avm_tests PRIVATE -Wall -Wextra -Wpedantic)
# target_compile_options(avm_tests PRIVATE -Wall -Wextra -Wpedantic -Weverything -Wno-c++98-compat -Wno-padded)

set_source_files_properties( buffer.cpp PROPERTIES COMPILE_FLAGS "-Wconversion -Wparentheses" )

target_link_libraries(
	avm_tests
PRIVATE
    cavm
	avm
	Catch2::Catch2
	nlohmann_json::nlohmann_json
    Boost::filesystem
    Boost::system
)

include(CTest)
include(Catch)
catch_discover_tests(avm_tests)

target_code_coverage(avm_tests AUTO ALL EXCLUDE external/* tests/*)
