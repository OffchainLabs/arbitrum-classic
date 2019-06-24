cmake_minimum_required (VERSION 3.5)

set (CMAKE_CXX_STANDARD 14)

project (avm LANGUAGES CXX)

set(LIB_HEADERS
  include/avm/bigint.hpp
  include/avm/cmachine.h
  include/avm/code.hpp
  include/avm/datastack.hpp
  include/avm/machine.hpp
  include/avm/opcodes.hpp
  include/avm/pool.hpp
  include/avm/tuple.hpp
  include/avm/util.hpp
  include/avm/value.hpp
)

set(LIB_SOURCES
  src/cmachine.cpp
  src/code.cpp
  src/machine.cpp
  src/pool.cpp
  src/value.cpp
  
  ${KECCAK_SOURCES}
)

add_library(avm STATIC ${LIB_HEADERS} ${LIB_SOURCES} )

target_compile_options(avm PRIVATE -Wall -Wextra -Wpedantic)

target_include_directories(avm PUBLIC
  $<BUILD_INTERFACE:${CMAKE_CURRENT_SOURCE_DIR}/include>
  $<INSTALL_INTERFACE:include>
  PRIVATE src
)

target_include_directories(avm PRIVATE
  ${CMAKE_CURRENT_SOURCE_DIR}/../external
)

target_link_libraries(avm variant optional Boost::boost)

source_group(avm FILES ${LIB_HEADERS} ${LIB_SOURCES})