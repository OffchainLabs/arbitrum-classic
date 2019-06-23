cmake_minimum_required (VERSION 3.5)

set (CMAKE_CXX_STANDARD 14)

project (avm_runner LANGUAGES CXX)
add_executable(avm_runner main.cpp)
target_link_libraries(avm_runner avm)
