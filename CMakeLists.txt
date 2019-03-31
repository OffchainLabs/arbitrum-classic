cmake_minimum_required (VERSION 2.6)

set (CMAKE_CXX_STANDARD 11)

project (AVMtest)

add_library(libAVMtest STATIC Machine.c InternalMachine.cpp code.cpp datastack.cpp value.cpp pool.cpp uint128_t.cpp uint256_t.cpp )

add_executable(AVMtest main.cpp)
target_link_libraries(AVMtest libAVMtest)
