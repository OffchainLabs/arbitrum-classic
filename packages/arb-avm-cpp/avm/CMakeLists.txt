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

  include/avm/machine.hpp
  include/avm/opcodes.hpp
  include/avm/exceptions.hpp
  include/avm/value/tuple.hpp
  include/avm/value/value.hpp
  include/avm/value/bigint.hpp
  include/avm/value/codepoint.hpp
  include/avm/value/pool.hpp
  include/avm/machinestate/machinestate.hpp
  include/avm/machinestate/tokenTracker.hpp
  include/avm/machinestate/datastack.hpp
  include/avm/machinestate/blockreason.hpp
  include/avm/machinestate/machineoperation.hpp
  include/avm/machinestate/messagestack.hpp
  include/avm/checkpoint/checkpointutils.hpp
  include/avm/checkpoint/checkpointdeleter.hpp
  include/avm/checkpoint/machinestatesaver.hpp
  include/avm/checkpoint/machinestatefetcher.hpp
  include/avm/checkpoint/checkpointstorage.hpp
)

set(LIB_SOURCES
  src/util.hpp
  src/machine.cpp
  src/value/bigint.cpp
  src/value/codepoint.cpp
  src/value/pool.cpp
  src/value/tuple.cpp
  src/value/value.cpp
  src/machinestate/tokenTracker.cpp
  src/machinestate/machinestate.cpp
  src/machinestate/machineoperation.cpp
  src/machinestate/datastack.cpp
  src/machinestate/blockreason.cpp
  src/checkpoint/machinestatesaver.cpp
  src/checkpoint/checkpointutils.cpp
  src/checkpoint/machinestatefetcher.cpp
  src/checkpoint/checkpointdeleter.cpp
  src/checkpoint/checkpointstorage.cpp
)

add_library(avm STATIC ${LIB_HEADERS} ${LIB_SOURCES} ${KECCAK_SOURCES} )

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

target_link_libraries(avm PUBLIC Boost::boost boost_filesystem CONAN_PKG::variant-lite rocksdb PRIVATE avm_utils)

target_code_coverage(avm AUTO ALL)

if(${CMAKE_VERSION} VERSION_LESS "3.8.0")
    source_group(avm FILES ${LIB_HEADERS} ${LIB_SOURCES} ${KECCAK_SOURCES})
else()
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/include/avm" FILES ${LIB_HEADERS})
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/src" FILES ${LIB_SOURCES})
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/../external" FILES ${KECCAK_SOURCES})
endif()

