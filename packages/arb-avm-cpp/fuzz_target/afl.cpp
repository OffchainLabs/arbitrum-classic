/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <unistd.h>
#include <cstdint>

extern "C" int LLVMFuzzerInitialize(int* argc, char*** argv);

extern "C" int LLVMFuzzerTestOneInput(uint8_t* buf, size_t len);

__AFL_FUZZ_INIT();

int main(int argc, char** argv) {
    LLVMFuzzerInitialize(&argc, &argv);

#ifdef __AFL_HAVE_MANUAL_CONTROL
    __AFL_INIT();
#endif

    unsigned char* buf = __AFL_FUZZ_TESTCASE_BUF;  // must be after __AFL_INIT
                                                   // and before __AFL_LOOP!

    while (__AFL_LOOP(10000)) {
        int len = __AFL_FUZZ_TESTCASE_LEN;  // don't use the macro directly in a
                                            // call!

        LLVMFuzzerTestOneInput(buf, len);
    }

    return 0;
}
