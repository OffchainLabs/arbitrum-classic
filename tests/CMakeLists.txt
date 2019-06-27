cmake_minimum_required (VERSION 3.5)
set (CMAKE_CXX_STANDARD 14)

project (avm_tests)

set(TEST_CASES_PATH ${CMAKE_CURRENT_BINARY_DIR}/test_cases.json)

configure_file (
    "${CMAKE_CURRENT_SOURCE_DIR}/config.hpp.in"
    "${CMAKE_CURRENT_SOURCE_DIR}/config.hpp"
)

configure_file(test_cases.json ${CMAKE_CURRENT_BINARY_DIR}/test_cases.json COPYONLY)

add_executable(tests main.cpp opcodes.cpp value.cpp balanceTracker.cpp)
target_link_libraries(
	tests
PRIVATE
	avm
	Catch2::Catch2
	nlohmann_json::nlohmann_json
)

include(CTest)
include(Catch)
catch_discover_tests(tests)
