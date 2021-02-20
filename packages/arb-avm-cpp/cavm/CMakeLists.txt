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

project (cavm LANGUAGES CXX)


set(LIB_HEADERS
  caggregator.h
  carbcore.h
  carbstorage.h
  cblockstore.h
  cexecutioncursor.h
  cmachine.h
  ctypes.h
  cvaluecache.h
  utils.hpp
)

set(LIB_SOURCES
  caggregator.cpp
  carbcore.cpp
  carbstorage.cpp
  cblockstore.cpp
  cexecutioncursor.cpp
  cmachine.cpp
  cvaluecache.cpp
)

add_library(cavm STATIC ${LIB_HEADERS} ${LIB_SOURCES})

set_target_properties(
    cavm
    PROPERTIES
    VERSION ${AVM_VERSION_STRING}
    SOVERSION ${AVM_VERSION_MAJOR}
)

target_include_directories(cavm PUBLIC
    $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}>
)

target_compile_options(cavm PRIVATE -Wall -Wextra -Wpedantic)
# target_compile_options(cavm PRIVATE -Wall -Wextra -Wpedantic -Weverything -Wno-c++98-compat -Wno-padded)

target_link_libraries(cavm PUBLIC data_storage avm)

source_group(cavm FILES ${LIB_HEADERS} ${LIB_SOURCES})
