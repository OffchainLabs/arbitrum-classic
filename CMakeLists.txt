cmake_minimum_required (VERSION 3.5)

set (CMAKE_CXX_STANDARD 14)

project (AVMtest)

enable_testing()

find_package(Boost 1.67 REQUIRED)

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
add_subdirectory(app)

add_subdirectory(tests)
