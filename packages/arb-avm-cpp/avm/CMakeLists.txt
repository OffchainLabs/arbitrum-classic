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

project (avm LANGUAGES CXX)

set(LIB_HEADERS
  include/avm/bigint.hpp
  include/avm/codepoint.hpp
  include/avm/datastack.hpp
  include/avm/machine.hpp
  include/avm/opcodes.hpp
  include/avm/pool.hpp
  include/avm/tokenTracker.hpp
  include/avm/tuple.hpp
  include/avm/value.hpp
)

set(LIB_SOURCES
  src/bigint.cpp
  src/codepoint.cpp
  src/datastack.cpp
  src/machine.cpp
  src/pool.cpp
  src/tuple.cpp
  src/value.cpp
  src/tokenTracker.cpp

  src/util.hpp

  ${KECCAK_SOURCES}
)

add_library(avm STATIC ${LIB_HEADERS} ${LIB_SOURCES} )

set_target_properties(
    avm
    PROPERTIES
    VERSION ${AVM_VERSION_STRING}
    SOVERSION ${AVM_VERSION_MAJOR}
)

target_compile_options(avm PRIVATE -Wall -Wextra -Wpedantic)

target_include_directories(avm PUBLIC
  $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
  $<INSTALL_INTERFACE:include>
  PRIVATE src
)

target_include_directories(avm PRIVATE
  ${CMAKE_CURRENT_SOURCE_DIR}/../external
)

target_link_libraries(avm PUBLIC Boost::boost CONAN_PKG::variant-lite PRIVATE avm_utils)

target_code_coverage(avm AUTO ALL)

source_group(avm FILES ${LIB_HEADERS} ${LIB_SOURCES})
