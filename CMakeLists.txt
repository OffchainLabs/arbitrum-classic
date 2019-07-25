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

include(${CMAKE_BINARY_DIR}/conanbuildinfo.cmake)
conan_basic_setup()
# ${CONAN_LIBS}
find_package(Boost 1.65 REQUIRED)
find_package(variant-lite 1.2.2 REQUIRED)
find_package(Catch2 2.9.1 REQUIRED)
find_package(nlohmann_json 3.6.1 REQUIRED)


set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} ${CONAN_CMAKE_MODULE_PATH} ${PROJECT_SOURCE_DIR}/cmake)

include(CodeCoverage)
include(Formatting)

file(GLOB_RECURSE ALL_CODE_FILES
    ${PROJECT_SOURCE_DIR}/avm/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/app/*.[ch]pp
    ${PROJECT_SOURCE_DIR}/tests/*.[h]pp
)

clang_format(TARGET_NAME ${ALL_CODE_FILES})



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

set(CMAKE_MODULE_PATH ${CMAKE_MODULE_PATH} "${CMAKE_CURRENT_SOURCE_DIR}/external/Catch2/contrib/")

file(GLOB KECCAK_SOURCES
  external/keccak/*.c
)

add_subdirectory(avm)
add_subdirectory(cavm)
add_subdirectory(app)

add_subdirectory(tests)

