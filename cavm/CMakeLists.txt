cmake_minimum_required (VERSION 3.5)

set (CMAKE_CXX_STANDARD 14)

project (cavm LANGUAGES CXX)

add_library(cavm STATIC cmachine.h cmachine.cpp )

target_compile_options(cavm PRIVATE -Wall -Wextra -Wpedantic)

target_link_libraries(cavm avm)

source_group(cavm FILES cmachine.h cmachine.cpp)
