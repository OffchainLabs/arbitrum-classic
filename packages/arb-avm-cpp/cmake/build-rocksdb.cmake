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

set(CMAKE_THREAD_PREFER_PTHREAD TRUE)
set(THREADS_PREFER_PTHREAD_FLAG TRUE)
find_package(Threads REQUIRED)
# find_package(zlib REQUIRED)
# find_package(bzip2 REQUIRED)

if("${CMAKE_GENERATOR}" STREQUAL "Unix Makefiles")
  set(BUILD_CMD $(MAKE) rocksdb)
else()
  set(BUILD_CMD ${CMAKE_COMMAND} --build <BINARY_DIR> --target rocksdb)
endif()

ExternalProject_Add(rocks
    GIT_REPOSITORY "https://github.com/facebook/rocksdb.git"
    GIT_TAG "v6.4.6"
    SOURCE_DIR "${CMAKE_CURRENT_SOURCE_DIR}/rocksdb"
    BINARY_DIR "${CMAKE_CURRENT_BINARY_DIR}/rocksdb"
    CMAKE_ARGS
        -DCMAKE_CXX_STANDARD=14
        -DWITH_TESTS=OFF
        -DWITH_TOOLS=OFF
        -DCMAKE_POSITION_INDEPENDENT_CODE=True
        -DWITH_ZLIB=OFF
        -DWITH_BZ2=OFF
        -DWITH_LZ4=OFF
        -DWITH_ZSTD=OFF
        -DWITH_SNAPPY=OFF
        -DWITH_JEMALLOC=OFF
        -DWITH_GFLAGS=OFF
        -DFAIL_ON_WARNINGS=OFF
    BUILD_COMMAND ${BUILD_CMD}
    INSTALL_COMMAND cmake -E echo "Skipping install step."
)

ExternalProject_Get_Property(rocks BINARY_DIR)
set(ROCKSDB_LIBRARIES ${BINARY_DIR}/${CMAKE_CFG_INTDIR}/librocksdb.a)

link_directories(${BINARY_DIR})

set(ROCKSDB_FOUND TRUE)

set(ROCKSDB_INCLUDE_DIRS
    ${CMAKE_CURRENT_SOURCE_DIR}/rocksdb/include)

add_library(RocksDB::RocksDB INTERFACE)
add_dependencies(RocksDB::RocksDB rocks)
target_include_directories(RocksDB::RocksDB INTERFACE ${ROCKSDB_INCLUDE_DIRS})
# ZLIB::ZLIB BZip2::BZip2
target_link_libraries(RocksDB::RocksDB INTERFACE ${ROCKSDB_LIBRARIES} Threads::Threads)

mark_as_advanced(
    ROCKSDB_LIBRARIES
    ROCKSDB_INCLUDE_DIRS
)
