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

set (CMAKE_CXX_STANDARD 17)

project (fuzz_target LANGUAGES CXX)

set(SHARED_SOURCE utils.cpp utils.hpp)

if (${AVM_FUZZER} MATCHES "afl")
    add_executable(fuzz_target afl.cpp libfuzzer.cpp ${SHARED_SOURCE})
elseif (${AVM_FUZZER} MATCHES "libfuzzer" OR ${AVM_FUZZER} MATCHES "honggfuzz")
    add_executable(fuzz_target libfuzzer.cpp ${SHARED_SOURCE})
else ()
    message(FATAL_ERROR "Unsupported AVM_FUZZER value \"${AVM_FUZZER}\" (must be afl, libfuzzer, or honggfuzz)")
endif ()

add_executable(manual_test manualtest.cpp ${SHARED_SOURCE})

target_link_libraries(fuzz_target PRIVATE avm)
target_link_libraries(manual_test PRIVATE avm)
