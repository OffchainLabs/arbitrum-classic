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

    std::ifstream myfile;

    struct stat filestatus;
    stat(filename.c_str(), &filestatus);

    char* buf = (char*)malloc(filestatus.st_size);

    myfile.open(filename, std::ios::in);
    if (myfile.is_open()) {
        myfile.read((char*)buf, filestatus.st_size);
        myfile.close();
    }
    std::cout << "In read_files. Done reading " << filename << std::endl;

    Machine mach;
    mach.deserialize(buf);

    auto start = std::chrono::high_resolution_clock::now();

    auto msg1DataRaw = hexStringToBytes(
        "0706050b03030303030303002d35a8a200000000000000000000000000000000000000"
        "0000000000000000000000000000000000000000000000000000000000000000000000"
        "0000000000000400000000000000000000000000895521964d724c8362a36608aaf09a"
        "3d7d0a0445000000000000000000000000000000000000000000000000000000000000"
        "00000a0000000000000000000000000000000000000000000000000000000000000000"
        "0000000000000000000000000000000000000000000000000000000000000000000000"
        "0000000000000000000000000000000000000000000000000000000000000000");
    auto msg2DataRaw = hexStringToBytes(
        "0706050b030b0303030303030300000000010000000000000000000000000000000000"
        "00000000000000000000000303030303003477ee2e0000000000000000000000000000"
        "0000000000000000000000000000000000000000000000000000000000000000000000"
        "00000000000000000000002400000000000000000000000000895521964d724c8362a3"
        "6608aaf09a3d7d0a044500000000000000000000000000000000000000000000000000"
        "000000000000000c000000000000000000000000000000000000000000000000000000"
        "0000000000000000000000000000000000000000000000000000000000000000000000"
        "0000000000000000000000000000000000000000000000000000000000000000000000"
        "0000");
    auto msg3DataRaw = hexStringToBytes(
        "0706050b030b0303030303030300000000020000000000000000000000000000000000"
        "00000000000000000000000303030303003477ee2e0000000000000000000000000000"
        "0000000000000000000000000000000000000000000000000000000000000000000000"
        "00000000000000000000002400000000000000000000000000895521964d724c8362a3"
        "6608aaf09a3d7d0a044500000000000000000000000000000000000000000000000000"
        "000000000000000e000000000000000000000000000000000000000000000000000000"
        "0000000000000000000000000000000000000000000000000000000000000000000000"
        "0000000000000000000000000000000000000000000000000000000000000000000000"
        "0000");

    auto msg1DataRawPtr = msg1DataRaw.data();
    auto msg2DataRawPtr = msg2DataRaw.data();
    auto msg3DataRawPtr = msg3DataRaw.data();

    auto msg1Data = deserialize_value(msg1DataRawPtr, mach.getPool());
    auto msg2Data = deserialize_value(msg2DataRawPtr, mach.getPool());
    auto msg3Data = deserialize_value(msg3DataRawPtr, mach.getPool());

    mach.sendOnchainMessage(Message{msg1Data, 0, 0, {}});
    mach.deliverOnchainMessages();
    Assertion assertion1 = mach.run(stepCount, 0, 0);
    mach.sendOnchainMessage(Message{msg2Data, 0, 0, {}});
    mach.deliverOnchainMessages();
    Assertion assertion2 = mach.run(stepCount, 0, 0);
    mach.sendOnchainMessage(Message{msg3Data, 0, 0, {}});
    mach.deliverOnchainMessages();
    Assertion assertion3 = mach.run(stepCount, 0, 0);

    auto finish = std::chrono::high_resolution_clock::now();
    std::chrono::duration<double> elapsed = finish - start;
    std::cout << assertion1.stepCount << " " << assertion2.stepCount << " "
              << assertion3.stepCount << " steps in " << elapsed.count() * 1000
              << " milliseconds" << std::endl;
    std::cout << to_hex_str(mach.hash()) << "\n" << mach << std::endl;
    std::this_thread::sleep_for(1s);
    return 0;
}
