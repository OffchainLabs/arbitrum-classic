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
  include/data_storage/aggregator.hpp
  include/data_storage/arbcore.hpp
  include/data_storage/arbstorage.hpp
  include/data_storage/blockstore.hpp
  include/data_storage/checkpoint.hpp
  include/data_storage/datastorage.hpp
  include/data_storage/inboxmessage.hpp
  include/data_storage/keyvaluestore.hpp
  include/data_storage/messageentry.hpp
  include/data_storage/messagestore.hpp
  include/data_storage/storageresult.hpp
  include/data_storage/storageresultfwd.hpp
  include/data_storage/value/code.hpp
  include/data_storage/value/machine.hpp
  include/data_storage/value/value.hpp
  include/data_storage/value/valuecache.hpp
)

set(LIB_PRIVATE_HEADERS
  src/value/referencecount.hpp
  src/value/utils.hpp
)

set(LIB_SOURCES
  src/aggregator.cpp
  src/arbcore.cpp
  src/arbstorage.cpp
  src/blockstore.cpp
  src/checkpoint.cpp
  src/datastorage.cpp
  src/inboxmessage.cpp
  src/keyvaluestore.cpp
  src/messageentry.cpp
  src/messagestore.cpp
  src/value/code.cpp
  src/value/machine.cpp
  src/value/referencecount.cpp
  src/value/value.cpp
  src/value/valuecache.cpp
)

add_library(data_storage STATIC ${LIB_HEADERS} ${LIB_PRIVATE_HEADERS} ${LIB_SOURCES})

set_target_properties(
    data_storage
    PROPERTIES
    VERSION ${AVM_VERSION_STRING}
    SOVERSION ${AVM_VERSION_MAJOR}
)

target_compile_options(data_storage PRIVATE -Wall -Wextra -Wpedantic)
# target_compile_options(data_storage PRIVATE -Wall -Wextra -Wpedantic -Weverything -Wno-c++98-compat -Wno-padded)

target_include_directories(data_storage PUBLIC
  $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
  $<INSTALL_INTERFACE:include>
  PRIVATE src
)

target_include_directories(data_storage PRIVATE
  ${CMAKE_CURRENT_SOURCE_DIR}/../external
)

target_link_libraries(data_storage PUBLIC avm avm_values cavm Threads::Threads)

target_code_coverage(data_storage AUTO ALL)

if(${CMAKE_VERSION} VERSION_LESS "3.8.0")
    source_group(data_storage FILES ${LIB_HEADERS} ${LIB_PRIVATE_HEADERS} ${LIB_SOURCES})
else()
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/include/data_storage" FILES ${LIB_HEADERS})
    source_group(TREE "${CMAKE_CURRENT_SOURCE_DIR}/src" FILES ${LIB_SOURCES} ${LIB_PRIVATE_HEADERS})
endif()

