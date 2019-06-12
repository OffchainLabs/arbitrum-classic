//
//  CMachine.c
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include "cmachine.h"
#include "machine.hpp"

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <iostream>
#include <fstream>
#include <sys/types.h>
#include <sys/stat.h>

struct cmachine {
    void *obj;
};


//cassertion machine_run(cmachine_t *m, uint64_t maxSteps);


struct cmachine;
typedef struct cmachine cmachine_t;

typedef struct {
    uint64_t stepCount;
} cassertion;

//void *machine_create();
void machine_destroy(cmachine_t *m);

Machine read_file (std::string filename) {
    std::cout<<"In read_file. reading - "<<filename<<std::endl;
    std::ifstream myfile;
    
    struct stat filestatus;
    stat( filename.c_str(), &filestatus );
    
    char *buf = (char *)malloc(filestatus.st_size);
    
    myfile.open(filename, std::ios::in);
    if (myfile.is_open())
    {
        myfile.read((char *)buf, filestatus.st_size);
        myfile.close();
    }
    return Machine(buf);
}

//cmachine_t *machine_create(char *data)
void *machine_create(char *filename)
{
    std::cout<<"In machine_create "<<std::endl;
    Machine mach = read_file(filename);
    cmachine_t *m;
//    Machine *obj;
    std::cout<<"In machine_create machine created"<<std::endl;

    m      = (typeof(m))malloc(sizeof(*m));
    std::cout<<"In machine_create m malloced"<<std::endl;
//    obj    = new Machine(data);
    m->obj = &mach;
    std::cout<<"In machine_create m->obj set"<<std::endl;

    return (void *)&mach;
}

void machine_destroy(cmachine_t *m) {
    if (m == NULL)
        return;
    delete static_cast<Machine *>(m->obj);
    free(m);
}

cassertion machine_run(cmachine_t *m, uint64_t maxSteps) {
    Machine *obj;
    
    if (m == NULL)
        return cassertion{0};
    
    obj = static_cast<Machine *>(m->obj);
    Assertion assertion = obj->run(maxSteps);
    return cassertion{assertion.stepCount};
}
