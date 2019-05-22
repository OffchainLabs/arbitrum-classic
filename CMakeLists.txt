cmake_minimum_required (VERSION 2.6)

set (CMAKE_CXX_STANDARD 14)

project (AVMtest)

set(Boost_ADDITIONAL_VERSIONS "1.67" "1.67.0")
find_package(Boost 1.60.0 REQUIRED)

file(GLOB KECCAK_SOURCES
  external/keccak/*.c
)

add_library(variant INTERFACE)
target_include_directories(variant INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/external/variant/include/)

add_library(optional INTERFACE)
target_include_directories(optional INTERFACE ${CMAKE_CURRENT_SOURCE_DIR}/external/optional-lite/include/)


set(LIB_HEADERS
  bigint.hpp
  cmachine.h
  machine.hpp
  code.hpp
  datastack.hpp
  value.hpp
  pool.hpp
  opcodes.hpp
)

set(LIB_SOURCES
  cmachine.cpp
  machine.cpp
  code.cpp
  datastack.cpp
  value.cpp
  pool.cpp
  ${KECCAK_SOURCES}
)


add_library(libAVMtest STATIC ${LIB_HEADERS} ${LIB_SOURCES} )

target_include_directories(libAVMtest PUBLIC
  ${Boost_INCLUDE_DIR}
  ${CMAKE_CURRENT_SOURCE_DIR}/external
)

target_link_libraries(libAVMtest variant optional)

add_executable(AVMtest main.cpp)
target_link_libraries(AVMtest libAVMtest)
