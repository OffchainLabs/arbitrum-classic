cmake_minimum_required (VERSION 3.5)
set (CMAKE_CXX_STANDARD 14)

project (avm_tests)

add_executable(tests main.cpp opcodes.cpp)
target_link_libraries(tests PRIVATE avm Catch2::Catch2)

include(CTest)
include(Catch)
catch_discover_tests(tests)

