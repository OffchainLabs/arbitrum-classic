//
//  CMachine.c
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include <avm/cmachine.h>

#include <avm/machine.hpp>

#include <sys/stat.h>
#include <fstream>
#include <iostream>

typedef struct {
    uint64_t stepCount;
} cassertion;

Machine* read_files(std::string filename, std::string inboxfile) {
    std::cout << "In read_file. reading - " << filename << std::endl;
    std::ifstream myfile;

    struct stat filestatus;
    struct stat inboxfilestatus;
    stat(filename.c_str(), &filestatus);

    char* buf = (char*)malloc(filestatus.st_size);

    myfile.open(filename, std::ios::in);
    if (myfile.is_open()) {
        myfile.read((char*)buf, filestatus.st_size);
        myfile.close();
    }
    char* inbox = NULL;
    std::cout << "In read_files. Done reading " << filename << std::endl;
    if (!inboxfile.empty()) {
        std::cout << "In read_files. reading - " << inboxfile << std::endl;
        std::ifstream myfile;

        stat(inboxfile.c_str(), &inboxfilestatus);

        inbox = (char*)malloc(inboxfilestatus.st_size);

        myfile.open(inboxfile, std::ios::in);
        if (myfile.is_open()) {
            myfile.read((char*)inbox, inboxfilestatus.st_size);
            myfile.close();
        } else {
            std::cout << "In read_files. " << inboxfile << " not found" << std::endl;
        }
    }
    return new Machine(buf, inbox, inboxfilestatus.st_size);
}

// cmachine_t *machine_create(char *data)
CMachine* machineCreate(const char* filename, const char* inboxfile) {
    Machine* mach = read_files(filename, inboxfile);
    return static_cast<void*>(mach);
}

void machineDestroy(CMachine* m) {
    std::cout << "In machine_destroy"<<std::endl;
    if (m == NULL)
        return;
    delete static_cast<Machine*>(m);
}

void machineHash(CMachine* m, char *ret) {
    uint256_t retHash = static_cast<Machine*>(m)->hash();
    std::array<unsigned char, 32> val;
    to_big_endian(retHash, val.begin());
    std::copy(val.begin(), val.end(), ret);
}

void* machineClone(CMachine* m) {
    Machine *mach = new Machine(*(static_cast<Machine*>(m)));
    return static_cast<void*>(mach);
}

void machineInboxHash(CMachine *m, char *ret) {
    uint256_t retHash = static_cast<Machine*>(m)->inboxHash();
    std::array<unsigned char, 32> val;
    to_big_endian(retHash, val.begin());
    std::copy(val.begin(), val.end(), ret);
}

bool machineHasPendingMessages(CMachine *m) {
    Machine *mach = static_cast<Machine*>(m);
    return mach->hasPendingMessages();
}

void machineSendOnchainMessage(CMachine *m, char *data) {
    Machine *mach = static_cast<Machine*>(m);
    auto val = deserialize_value(data, mach->getPool());
    Message msg;
    auto success = msg.deserialize(val);
    if (!success) {
        throw std::runtime_error("Machine recieved invalid message");
    }
    mach->sendOnchainMessage(msg);
}

void machineDeliverOnchainMessages(CMachine *m) {
    Machine *mach = static_cast<Machine*>(m);
    mach->deliverOnchainMessages();
}

ByteSlice machineMarshallForProof(CMachine *m) {
    Machine *mach = static_cast<Machine*>(m);
    std::vector<unsigned char> buffer;
    auto proof = mach->marshalForProof();
    auto proofData = (unsigned char *)malloc(proof.size());
    std::copy(proof.begin(), proof.end(), proofData);
    return {proofData, static_cast<int>(proof.size())};
}

void machineSendOffchainMessages(CMachine *m, char *data, int size) {
    Machine *mach = static_cast<Machine*>(m);
    std::vector<Message> messages;
    auto end = data + size;
    while(data < end) {
        auto val = deserialize_value(data, mach->getPool());
        messages.emplace_back();
        auto success = messages.back().deserialize(val);
        if (!success) {
            throw std::runtime_error("Machine recieved invalid message");
        }
    }
    mach->sendOffchainMessages(messages);
}

uint64_t machineExecuteAssertion(CMachine* m, uint64_t maxSteps, uint64_t timeboundStart, uint64_t timeboundEnd) {
    Machine* mach = static_cast<Machine*>(m);
    Assertion assertion = mach->run(maxSteps, timeboundStart, timeboundEnd);
    printf("%llu steps ran\n", assertion.stepCount);
    return assertion.stepCount;
}
