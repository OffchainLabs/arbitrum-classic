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

#include "bigint_utils.hpp"

#include <avm/machine.hpp>

#include <boost/algorithm/hex.hpp>

#include <sys/stat.h>
#include <fstream>
#include <iostream>
#include <string>
#include <thread>

std::vector<char> hexStringToBytes(const std::string& hexstr) {
    std::vector<char> bytes;
    bytes.reserve(hexstr.size() / 2);
    boost::algorithm::unhex(hexstr.begin(), hexstr.end(), bytes.begin());
    return bytes;
}

int main(int argc, char* argv[]) {
    using namespace std::chrono_literals;
    std::string filename;
    unsigned long long stepCount = 10000000000;
    if (argc < 2) {
        std::cout << "Usage: AVMTest <ao file>" << std::endl;
        std::cout << "   defaulting to use add.ao" << std::endl;
        filename = "add.ao";
    } else {
        filename = argv[1];
    }
    std::cout << filename << std::endl;

    Machine mach;
    mach.initializeMachine(filename);

    auto start = std::chrono::high_resolution_clock::now();

    auto msg1DataRaw = hexStringToBytes(
        "06000000000000000000000000000000000000000000000000000000000000000afe00"
        "7ec3319c8348582438bc9155320d99b577975540efd44c0878246f4007110342060000"
        "0000000000000000000000000000000000000000000000000000000000000000000000"
        "000000000000000000c7711f36b2c13e00821ffd9ec54b04a60aefbd1b070000000000"
        "0000000000000000d02bec7ee5ee73a271b144e829eed1c19218d81300ffffffffffff"
        "fffffffffffffffffffffffffffffffffffffffffffffffffffe000000000000000000"
        "000000000000000000000000000000000000000000000000050b030b03030303030303"
        "000aefbd1b000000000000000000000000000000000000000000000000000000000303"
        "0303030070a08231000000000000000000000000c7711f36b2c13e00821ffd9ec54b04"
        "a6000000000000000000000000000000000000000000000000000000000000000024");

    auto msg1DataRawPtr = const_cast<const char*>(msg1DataRaw.data());

    auto msg1 =
        nonstd::get<Tuple>(deserialize_value(msg1DataRawPtr, mach.getPool()));

    std::cout << "Msg " << msg1 << std::endl;

    Assertion assertion1 = mach.run(stepCount, {0, 0, 0, 0}, std::move(msg1),
                                    std::chrono::seconds{1000});

    auto finish = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double> elapsed = finish - start;
    std::cout << assertion1.stepCount << " "
              << " steps in " << elapsed.count() * 1000 << " milliseconds"
              << std::endl;
    std::cout << to_hex_str(mach.hash()) << "\n" << mach << std::endl;
    std::this_thread::sleep_for(1s);
    return 0;
}
