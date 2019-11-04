# Copyright (c) 2018-present, Facebook, Inc. and its affiliates.
# All rights reserved.
#
# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree.

# Build the RocksDB library.
#
# Variables used by this module, they can change the default behaviour and need
# to be set before calling find_package:
#
# Variables defined by this module:
#
#  ROCKSDB_FOUND               RocksDB library/headers were found
#  ROCKSDB_LIBRARIES           The RocksDB library.
#  ROCKSDB_INCLUDE_DIRS        The location of RocksDB headers.

message( ${CMAKE_CURRENT_SOURCE_DIR} )
message( ${CMAKE_SOURCE_DIR} )
message( ${CMAKE_CURRENT_BINARY_DIR} )

include(ExternalProject)

ExternalProject_Add(rocksdb
    GIT_REPOSITORY "https://github.com/facebook/rocksdb.git"
    GIT_TAG "master"
    PREFIX "${CMAKE_CURRENT_BINARY_DIR}/rocksdb"
    CMAKE_ARGS 
        -DCMAKE_CXX_STANDARD=14
        -DWITH_TESTS=OFF
    CONFIGURE_COMMAND ""
    BUILD_IN_SOURCE ON
    BUILD_COMMAND make -j static_lib
    UPDATE_COMMAND "" 
)

ExternalProject_Get_Property(rocksdb BINARY_DIR)
set(ROCKSDB_LIBRARIES
    ${BINARY_DIR}/librocksdb.a)

message( ${BINARY_DIR} )
link_directories(${BINARY_DIR})

set(ROCKSDB_FOUND TRUE)

set(ROCKSDB_INCLUDE_DIRS
    ${BINARY_DIR}/include)
message(STATUS "Found RocksDB library: ${ROCKSDB_LIBRARIES}")
message(STATUS "Found RocksDB includes: ${ROCKSDB_INCLUDE_DIRS}")

add_library(rocks INTERFACE)
add_dependencies(rocks rocksdb)
target_include_directories(rocks INTERFACE ${ROCKSDB_INCLUDE_DIRS})
target_link_libraries(rocks INTERFACE ${ROCKSDB_LIBRARIES})

mark_as_advanced(
    ROCKSDB_LIBRARIES
    ROCKSDB_INCLUDE_DIRS
)