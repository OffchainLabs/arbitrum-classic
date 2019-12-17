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

project (avm_values LANGUAGES CXX)

set(LIB_HEADERS
  include/tuple.hpp
  include/value.hpp
  include/bigint.hpp
  include/codepoint.hpp
  include/pool.hpp
  include/opcodes.hpp
)

set(LIB_SOURCES
  src/util.hpp
  src/bigint.cpp
  src/codepoint.cpp
  src/pool.cpp
  src/tuple.cpp
  src/value.cpp
)

add_library(avm_values STATIC ${LIB_HEADERS} ${LIB_SOURCES} ${KECCAK_SOURCES})

set_target_properties(
    avm_values
    PROPERTIES
    VERSION ${AVM_VERSION_STRING}
    SOVERSION ${AVM_VERSION_MAJOR}
)

target_compile_options(avm_values PRIVATE -Wall -Wextra -Wpedantic)

target_include_directories(avm_values PUBLIC
  $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
  $<INSTALL_INTERFACE:include>
  PRIVATE src
)

target_include_directories(avm_values PRIVATE
  ${CMAKE_CURRENT_SOURCE_DIR}/../external
)

target_link_libraries(avm_values PUBLIC CONAN_PKG::variant-lite Threads::Threads PRIVATE avm_utils)

target_code_coverage(avm_values AUTO ALL)

if(${CMAKE_VERSION} VERSION_LESS "3.8.0")
    source_group(avm_values FILES ${LIB_HEADERS} ${LIB_SOURCES} ${KECCAK_SOURCES})
else()
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/include" FILES ${LIB_HEADERS})
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/src" FILES ${LIB_SOURCES})
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/../external" FILES ${KECCAK_SOURCES})
endif()

