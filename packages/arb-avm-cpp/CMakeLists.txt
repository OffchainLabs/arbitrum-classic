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

find_package(Git)
if(GIT_FOUND AND EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/../../.git")
# Update submodules as needed
    message(STATUS "Submodule update")
    execute_process(COMMAND ${GIT_EXECUTABLE} submodule update --init --recursive
                    WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
                    RESULT_VARIABLE GIT_SUBMOD_RESULT)
    if(NOT GIT_SUBMOD_RESULT EQUAL "0")
        message(FATAL_ERROR "git submodule update --init failed with ${GIT_SUBMOD_RESULT}, please checkout submodules")
    endif()
endif()

if(NOT EXISTS "${CMAKE_CURRENT_SOURCE_DIR}/cmake/cable/bootstrap.cmake")
    message(FATAL_ERROR "The submodules were not downloaded! GIT_SUBMODULE was turned off or failed. Please update submodules and try again.")
endif()

include(cmake/cable/bootstrap.cmake)
include(CableBuildType)
include(CableToolchains)

cable_configure_toolchain(DEFAULT cxx17)

include("cmake/HunterGate.cmake")
HunterGate(
    URL "https://github.com/cpp-pm/hunter/archive/v0.23.260.tar.gz"
    SHA1 "13775235910a3fa85644568d1c5be8271de72e1c"
)

cable_set_build_type(DEFAULT RelWithDebInfo CONFIGURATION_TYPES Debug Release RelWithDebInfo)

project (arb-avm-cpp)

include(CTest)

set(AVM_VERSION_MAJOR 0)
set(AVM_VERSION_MINOR 8)
set(AVM_VERSION_PATCH 0)

set(AVM_VERSION_STRING ${AVM_VERSION_MAJOR}.${AVM_VERSION_MINOR}.${AVM_VERSION_PATCH})

set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} ${PROJECT_SOURCE_DIR}/cmake)

option(ENABLE_JEMALLOC "Enable jemalloc" OFF)

set(AVM_FUZZER "" CACHE STRING "Build fuzzing target for the requested fuzzer")

if (${AVM_FUZZER} MATCHES "afl")
    add_compile_options(-fsanitize=address)
    add_link_options(-fsanitize=address)
elseif (${AVM_FUZZER} MATCHES "libfuzzer")
    add_compile_options(-fsanitize=address,fuzzer-no-link)
    add_link_options(-fsanitize=address,fuzzer)
elseif (${AVM_FUZZER} MATCHES "honggfuzz")
    add_compile_options(-fsanitize=address)
    add_link_options(-fsanitize=address)
endif ()

find_package(Threads REQUIRED)
find_package(GMP REQUIRED)
find_package(GMPXX REQUIRED)

if (ENABLE_JEMALLOC)
    find_package(PkgConfig REQUIRED)
    pkg_check_modules (JEMALLOC REQUIRED jemalloc)
    message(STATUS "Building with jemalloc")
    include_directories(${JEMALLOC_INCLUDE_DIRS})
    link_directories(${JEMALLOC_LIBRARY_DIRS})
endif (ENABLE_JEMALLOC)

# On macOS, search Homebrew for keg-only versions of OpenSSL
if (CMAKE_HOST_SYSTEM_NAME MATCHES "Darwin")
    execute_process(
        COMMAND brew --prefix OpenSSL
        RESULT_VARIABLE BREW_OPENSSL
        OUTPUT_VARIABLE BREW_OPENSSL_PREFIX
        OUTPUT_STRIP_TRAILING_WHITESPACE
    )
    if (BREW_OPENSSL EQUAL 0 AND EXISTS "${BREW_OPENSSL_PREFIX}")
        message(STATUS "Found OpenSSL keg installed by Homebrew at ${BREW_OPENSSL_PREFIX}")
        set(OPENSSL_ROOT_DIR "${BREW_OPENSSL_PREFIX}/")
        set(OPENSSL_INCLUDE_DIR "${BREW_OPENSSL_PREFIX}/include")
        set(OPENSSL_LIBRARIES "${BREW_OPENSSL_PREFIX}/lib")
        set(OPENSSL_CRYPTO_LIBRARY "${BREW_OPENSSL_PREFIX}/lib/libcrypto.dylib")
        set(OPENSSL_SSL_LIBRARY (ADVANCED) "${BREW_OPENSSL_PREFIX}/lib/libssl.dylib")
    endif()
endif()

find_package(Boost 1.65 COMPONENTS filesystem system REQUIRED)
if(NOT TARGET Boost::boost)
    add_library(Boost::boost IMPORTED INTERFACE)
    set_property(TARGET Boost::boost PROPERTY
        INTERFACE_INCLUDE_DIRECTORIES ${Boost_INCLUDE_DIR})
endif()
if(NOT TARGET Boost::filesystem)
    add_library(Boost::filesystem IMPORTED INTERFACE)
    set_property(TARGET Boost::filesystem PROPERTY
        INTERFACE_INCLUDE_DIRECTORIES ${Boost_INCLUDE_DIR})
    set_property(TARGET Boost::filesystem PROPERTY
        INTERFACE_LINK_LIBRARIES ${Boost_LIBRARIES})
endif()
if(NOT TARGET Boost::system)
    add_library(Boost::system IMPORTED INTERFACE)
    set_property(TARGET Boost::system PROPERTY
        INTERFACE_INCLUDE_DIRECTORIES ${Boost_INCLUDE_DIR})
    set_property(TARGET Boost::system PROPERTY
        INTERFACE_LINK_LIBRARIES ${Boost_LIBRARIES})
endif()
find_package(RocksDB REQUIRED)
include(CodeCoverage)
include(Formatting)

file(GLOB_RECURSE ALL_CODE_FILES
	${PROJECT_SOURCE_DIR}/avm_values/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/avm/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/app/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/tests/*.[h]pp
    ${PROJECT_SOURCE_DIR}/utils/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/data_storage/*.[ch]pp
)

clang_format(TARGET_NAME ${ALL_CODE_FILES})

enable_testing()
add_code_coverage_all_targets()

set(CMAKE_ARCHIVE_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_LIBRARY_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/lib)
set(CMAKE_RUNTIME_OUTPUT_DIRECTORY ${CMAKE_BINARY_DIR}/bin)

hunter_add_package(ethash)
find_package(ethash CONFIG REQUIRED)

hunter_add_package(intx)
find_package(intx CONFIG REQUIRED)

get_target_property(ETHASH_LIB ethash::keccak "LOCATION_Release")
get_filename_component(ETHASH_LIB_FOLDER "${ETHASH_LIB}" DIRECTORY)

add_subdirectory(external)
set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} ${CMAKE_CURRENT_SOURCE_DIR}/external/Catch2/contrib)


if(IS_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}/avm_values")
    add_subdirectory(avm_values)
endif()
if(IS_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}/data_storage" AND NOT AVM_FUZZER)
    add_subdirectory(data_storage)
endif()
if(IS_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}/avm")
    add_subdirectory(avm)
endif()
if(IS_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}/cavm" AND NOT AVM_FUZZER)
    add_subdirectory(cavm)

    get_property(_isMultiConfig GLOBAL PROPERTY GENERATOR_IS_MULTI_CONFIG)
    if (NOT _isMultiConfig)
        set(INFILE ${PROJECT_SOURCE_DIR}/cmachine/flags.go.in)
        set(MIDDLEFILE ${PROJECT_BINARY_DIR}/cmachine/flags.go.tmp)
        set(OUTFILE ${PROJECT_SOURCE_DIR}/cmachine/flags.go)

        set(C_LIBRARY_PATH ${PROJECT_BINARY_DIR}/lib)
        set(CFLAGS "-I.")
        string(CONCAT LDFLAGS
            " -L${C_LIBRARY_PATH}"
            " -L${ETHASH_LIB_FOLDER}"
            " $<TARGET_LINKER_FILE:cavm>"
            " $<TARGET_LINKER_FILE:data_storage>"
            " $<TARGET_LINKER_FILE:avm>"
            " $<TARGET_LINKER_FILE:avm_values>"
            " $<TARGET_LINKER_FILE:ff>"
            " ${JEMALLOC_LINK_LIBRARIES}"
        )
        configure_file( ${INFILE} ${MIDDLEFILE} )
        file(GENERATE OUTPUT ${OUTFILE} INPUT ${MIDDLEFILE})
    endif()

endif()
if(AVM_FUZZER)
    add_subdirectory(fuzz_target)
elseif(IS_DIRECTORY "${CMAKE_CURRENT_SOURCE_DIR}/app")
    add_subdirectory(app)
endif()

if (BUILD_TESTING AND NOT AVM_FUZZER)
  add_subdirectory(tests)
endif()

