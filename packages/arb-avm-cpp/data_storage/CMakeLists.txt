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

project (data_storage LANGUAGES CXX)

set(LIB_HEADERS
  include/data_storage/checkpointstorage.hpp
  include/data_storage/checkpointdeleter.hpp
  include/data_storage/checkpointutils.hpp
  include/data_storage/checkpointresult.hpp
  include/data_storage/transaction.hpp
)

set(LIB_SOURCES
  src/checkpointstorage.cpp
  src/checkpointdeleter.cpp
  src/checkpointutils.cpp
  src/transaction.cpp
)

add_library(data_storage STATIC ${LIB_HEADERS} ${LIB_SOURCES})

set_target_properties(
    data_storage
    PROPERTIES
    VERSION ${AVM_VERSION_STRING}
    SOVERSION ${AVM_VERSION_MAJOR}
)

target_compile_options(data_storage PRIVATE -Wall -Wextra -Wpedantic)

target_include_directories(data_storage PUBLIC
  $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
  $<INSTALL_INTERFACE:include>
  PRIVATE src
)

target_link_libraries(data_storage PUBLIC avm_values CONAN_PKG::variant-lite Threads::Threads PRIVATE avm_utils)

target_code_coverage(data_storage AUTO ALL)

if(${CMAKE_VERSION} VERSION_LESS "3.8.0")
    source_group(data_storage FILES ${LIB_HEADERS} ${LIB_SOURCES})
else()
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/include/data_storage" FILES ${LIB_HEADERS})
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/src" FILES ${LIB_SOURCES})
endif()

