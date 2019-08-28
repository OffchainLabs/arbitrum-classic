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

project (avm_utils LANGUAGES CXX)

add_library(avm_utils STATIC bigint_utils.hpp bigint_utils.cpp )

set_target_properties(
    avm_utils
    PROPERTIES
    VERSION ${AVM_VERSION_STRING}
    SOVERSION ${AVM_VERSION_MAJOR}
)

target_compile_options(avm_utils PRIVATE -Wall -Wextra -Wpedantic)

target_include_directories(avm_utils PUBLIC
  $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}>
  PRIVATE src
)

target_link_libraries(avm_utils PUBLIC Boost::boost)

source_group(avm_utils FILES bigint_utils.hpp bigint_utils.cpp)
