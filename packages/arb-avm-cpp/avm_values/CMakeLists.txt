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

project (avm_values LANGUAGES CXX)

set(LIB_HEADERS
  include/avm_values/exceptions.hpp
  include/avm_values/tuple.hpp
  include/avm_values/tuplestub.hpp
  include/avm_values/unloadedvalue.hpp
  include/avm_values/value.hpp
  include/avm_values/valuetype.hpp
  include/avm_values/bigint.hpp
  include/avm_values/buffer.hpp
  include/avm_values/code.hpp
  include/avm_values/codepoint.hpp
  include/avm_values/codepointstub.hpp
  include/avm_values/pool.hpp
  include/avm_values/opcodes.hpp
  include/avm_values/vmValueParser.hpp
)

set(LIB_SOURCES
  src/bigint.cpp
  src/buffer.cpp
  src/codepoint.cpp
  src/codepointstub.cpp
  src/exceptions.cpp
  src/pool.cpp
  src/tuple.cpp
  src/tuplestub.cpp
  src/value.cpp
  src/vmValueParser.cpp
  src/unloadedvalue.cpp
)

add_library(avm_values STATIC ${LIB_HEADERS} ${LIB_SOURCES})

set_target_properties(
    avm_values
    PROPERTIES
    VERSION ${AVM_VERSION_STRING}
    SOVERSION ${AVM_VERSION_MAJOR}
)

target_compile_options(avm_values PRIVATE -Wall -Wextra -Wpedantic)
# target_compile_options(avm_values PRIVATE -Wall -Wextra -Wpedantic -Weverything -Wno-c++98-compat -Wno-padded)

set_source_files_properties( src/buffer.cpp PROPERTIES COMPILE_FLAGS "-Wconversion -Wparentheses" )



target_include_directories(avm_values PUBLIC
  $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
  $<INSTALL_INTERFACE:include>
  PRIVATE src
)

target_include_directories(avm_values PRIVATE
  ${CMAKE_CURRENT_SOURCE_DIR}/../external
)

target_link_libraries(avm_values PUBLIC Threads::Threads nlohmann_json::nlohmann_json ethash::keccak intx::intx ${arb_node_LINKER_LIBS} PRIVATE Boost::boost)

target_code_coverage(avm_values AUTO ALL)

if(${CMAKE_VERSION} VERSION_LESS "3.8.0")
    source_group(avm_values FILES ${LIB_HEADERS} ${LIB_SOURCES})
else()
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/include/avm_values" FILES ${LIB_HEADERS})
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/src" FILES ${LIB_SOURCES})
endif()

