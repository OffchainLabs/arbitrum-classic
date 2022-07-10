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

project (cavm LANGUAGES CXX)


set(LIB_HEADERS
  caggregator.h
  carbcore.h
  carbstorage.h
  cexecutioncursor.h
  cmachine.h
  ctypes.h
  utils.hpp
)

set(LIB_SOURCES
  caggregator.cpp
  carbcore.cpp
  carbstorage.cpp
  cexecutioncursor.cpp
  cmachine.cpp
  dumpstate.cpp
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

get_property(_isMultiConfig GLOBAL PROPERTY GENERATOR_IS_MULTI_CONFIG)
if (NOT _isMultiConfig)
    set(INFILE ${PROJECT_SOURCE_DIR}/../cmachine/flags.go.in)
    set(MIDDLEFILE ${PROJECT_BINARY_DIR}/../cmachine/flags.go.tmp)
    set(OUTFILE ${PROJECT_SOURCE_DIR}/../cmachine/flags.go)

    set(C_LIBRARY_PATH ${PROJECT_BINARY_DIR}/../lib)
    set(CFLAGS "-I.")
    string(CONCAT LDFLAGS
        " -L${C_LIBRARY_PATH}"
        " -L${ETHASH_LIB_FOLDER}"
        " $<TARGET_LINKER_FILE:cavm>"
        " $<TARGET_LINKER_FILE:data_storage>"
        " $<TARGET_LINKER_FILE:avm>"
        " $<TARGET_LINKER_FILE:avm_values>"
        " $<TARGET_LINKER_FILE:ff>"
        " $<TARGET_LINKER_FILE:secp256k1>"
        " $<TARGET_LINKER_FILE:OpenSSL::Crypto>"
        " $<TARGET_LINKER_FILE:RocksDB::RocksDB>"
        " $<TARGET_LINKER_FILE:ethash::keccak>"
        " $<TARGET_LINKER_FILE:Boost::filesystem>"
        " ${GPROF_LIBRARY}"
        " ${MALLOC_LIBRARY}"
        " ${GMP_LIBRARY}"
        " ${GMPXX_LIBRARY}"
    )
    configure_file( ${INFILE} ${MIDDLEFILE} )
    file(GENERATE OUTPUT ${OUTFILE} INPUT ${MIDDLEFILE})
endif()
