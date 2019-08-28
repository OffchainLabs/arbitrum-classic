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

project (arb-avm-cpp)

set(AVM_VERSION_MAJOR 0)
set(AVM_VERSION_MINOR 2)
set(AVM_VERSION_PATCH 0)

set(AVM_VERSION_STRING ${AVM_VERSION_MAJOR}.${AVM_VERSION_MINOR}.${AVM_VERSION_PATCH})

include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)
conan_basic_setup(TARGETS)

find_package(Boost 1.65 REQUIRED)

set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} ${PROJECT_SOURCE_DIR}/cmake)

message(STATUS "CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH}")

include(CodeCoverage)
include(Formatting)

file(GLOB_RECURSE ALL_CODE_FILES
    ${PROJECT_SOURCE_DIR}/avm/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/app/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/tests/*.[h]pp
    ${PROJECT_SOURCE_DIR}/utils/*.[ch]pp
)

clang_format(TARGET_NAME ${ALL_CODE_FILES})

enable_testing()
add_code_coverage_all_targets()

file(GLOB KECCAK_SOURCES
  external/keccak/*.c
)

add_subdirectory(utils)
add_subdirectory(avm)
add_subdirectory(cavm)
add_subdirectory(app)

add_subdirectory(tests)

