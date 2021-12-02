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

project (fuzz_target LANGUAGES CXX)

set(SHARED_SOURCE utils.cpp utils.hpp)

if(UNIX AND NOT APPLE)
    set(LINUX TRUE)
endif()

if(LINUX)
    SET(LIBNAME soname)
    SET(LIB_EXTENSION so)
    # linux stuff here
elseif(APPLE)
    SET(LIBNAME install_name)
    SET(LIB_EXTENSION dylib)
else()
    message(FATAL_ERROR "not support on non-unix systems")
endif()

set(GO_SRC_DIR ${CMAKE_CURRENT_SOURCE_DIR}/proofchecker)
set(GO_SRCS ${GO_SRC_DIR}/proofchecker.go)
set(GO_LIBNAME libproofchecker.${LIB_EXTENSION})
set(GOPATH ${CMAKE_CURRENT_BINARY_DIR})



# Custom command for 'go build -buildmode=c-shared ...'
# to create a library from Go codes.
add_custom_command(OUTPUT ${CMAKE_CURRENT_BINARY_DIR}/${GO_LIBNAME}
  DEPENDS ${GO_SRCS}
  WORKING_DIRECTORY ${GO_SRC_DIR}
  COMMAND go build
  -buildmode=c-shared
  -ldflags '-extldflags -Wl,-${LIBNAME},${CMAKE_CURRENT_BINARY_DIR}/${GO_LIBNAME}'
  -o "${CMAKE_CURRENT_BINARY_DIR}/${GO_LIBNAME}"
  ${CMAKE_GO_FLAGS} ./...
  COMMENT "Building Go library")

add_custom_target(lib_proofchecker_go DEPENDS ${GO_LIBNAME} libproofchecker.h)
add_library(lib_proofchecker SHARED IMPORTED GLOBAL)
add_dependencies(lib_proofchecker lib_proofchecker_go)
set_target_properties(lib_proofchecker
  PROPERTIES
  IMPORTED_LOCATION ${CMAKE_CURRENT_BINARY_DIR}/${GO_LIBNAME}
  INTERFACE_INCLUDE_DIRECTORIES ${CMAKE_CURRENT_BINARY_DIR}
)

if (${AVM_FUZZER} MATCHES "afl")
    add_executable(fuzz_target afl.cpp libfuzzer.cpp ${SHARED_SOURCE})
elseif (${AVM_FUZZER} MATCHES "libfuzzer" OR ${AVM_FUZZER} MATCHES "honggfuzz")
    add_executable(fuzz_target libfuzzer.cpp ${SHARED_SOURCE})
elseif (NOT ${AVM_FUZZER} MATCHES "manual_test")
    message(FATAL_ERROR "Unsupported AVM_FUZZER value \"${AVM_FUZZER}\" (must be afl, libfuzzer, or honggfuzz)")
endif ()

add_executable(manual_test manualtest.cpp ${SHARED_SOURCE})

if (NOT ${AVM_FUZZER} MATCHES "manual_test")
    target_link_libraries(fuzz_target PRIVATE avm lib_proofchecker)

endif ()
target_link_libraries(manual_test PRIVATE avm lib_proofchecker)

