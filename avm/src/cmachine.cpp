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

        struct stat filestatus;
        stat(inboxfile.c_str(), &filestatus);

        inbox = (char*)malloc(filestatus.st_size);

        myfile.open(inboxfile, std::ios::in);
        if (myfile.is_open()) {
            myfile.read((char*)inbox, filestatus.st_size);
            myfile.close();
        }
    }
    return new Machine(buf, inbox);
}

// cmachine_t *machine_create(char *data)
void* machine_create(const char* filename, const char* inboxfile) {
    Machine* mach = read_files(filename, inboxfile);
    return static_cast<void*>(mach);
}

void machine_destroy(void* m) {
    if (m == NULL)
        return;
    delete static_cast<Machine*>(m);
    free(m);
}

// cassertion machine_run(cmachine_t *m, uint64_t maxSteps) {
uint64_t machine_run(void* m, uint64_t maxSteps) {
    if (m == NULL)
        return 0;
    Machine* mach = static_cast<Machine*>(m);
    Assertion assertion = mach->run(maxSteps, 0, 0);
    printf("%llu steps ran\n", assertion.stepCount);
    return assertion.stepCount;
}

void inbox_add_message(void *m, char *inbox){
    Machine *mach = static_cast<Machine*>(m);
    mach->addInboxMessage(inbox);
}

void machineSettime_bounds(void* m, uint64_t timeboundStart, uint64_t timeboundEnd){
    Machine *mach = static_cast<Machine*>(m);
    mach->setTimebounds(timeboundStart, timeboundEnd);
}
