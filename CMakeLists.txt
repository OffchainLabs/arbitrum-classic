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

project (AVMtest)

set(CMAKE_MODULE_PATH ${PROJECT_SOURCE_DIR}/cmake)

include(CodeCoverage)
include(Formatting)

file(GLOB_RECURSE ALL_CODE_FILES
    ${PROJECT_SOURCE_DIR}/avm/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/app/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/tests/*.[h]pp
)

clang_format(TARGET_NAME ${ALL_CODE_FILES})

# Code Coverage Configuration
# add_library(coverage_config INTERFACE)

# option(CODE_COVERAGE "Enable coverage reporting" OFF)
# if(CODE_COVERAGE)
#     if("${CMAKE_C_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang" OR "${CMAKE_CXX_COMPILER_ID}" MATCHES "(Apple)?[Cc]lang")
#         message("Building with llvm Code Coverage Tools")

#         # Warning/Error messages
#         if(NOT LLVM_COV_PATH)
#             message(FATAL_ERROR "llvm-cov not found! Aborting.")
#         endif()
#         set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -fprofile-instr-generate -fcoverage-mapping")
#         set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fprofile-instr-generate -fcoverage-mapping")
#     endif()

#   # # Add required flags (GCC & LLVM/Clang)
#   # target_compile_options(coverage_config INTERFACE
#   #   -O0        # no optimization
#   #   -g         # generate debug info
#   #   --coverage # sets all required flags
#   # )
#   # if(CMAKE_VERSION VERSION_GREATER_EQUAL 3.13)
#   #   target_link_options(coverage_config INTERFACE --coverage)
#   # else()
#   #   target_link_libraries(coverage_config INTERFACE --coverage)
#   # endif()
# endif()

find_package(Boost 1.65 REQUIRED)

enable_testing()
add_code_coverage_all_targets()

#from https://cliutils.gitlab.io/modern-cmake/chapters/projects/submodule.html
find_package(Git QUIET)
if(GIT_FOUND AND EXISTS "${PROJECT_SOURCE_DIR}/.git")
# Update submodules as needed
    option(GIT_SUBMODULE "Check submodules during build" ON)
    if(GIT_SUBMODULE)
        message(STATUS "Submodule update")
        execute_process(COMMAND ${GIT_EXECUTABLE} submodule update --init --recursive
                        WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
                        RESULT_VARIABLE GIT_SUBMOD_RESULT)
        if(NOT GIT_SUBMOD_RESULT EQUAL "0")
            message(FATAL_ERROR "git submodule update --init failed with ${GIT_SUBMOD_RESULT}, please checkout submodules")
        endif()
    endif()
endif()

add_subdirectory(external/Catch2)

set(JSON_BuildTests OFF CACHE INTERNAL "")
add_subdirectory(external/json)

set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/external/Catch2/contrib/")

file(GLOB KECCAK_SOURCES
  external/keccak/*.c
)

add_library(variant INTERFACE)
target_include_directories(variant INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/external/variant/include/)

add_library(optional INTERFACE)
target_include_directories(optional INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/external/optional-lite/include/)

add_subdirectory(avm)
add_subdirectory(cavm)
add_subdirectory(app)

add_subdirectory(tests)

