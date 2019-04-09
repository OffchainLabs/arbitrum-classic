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

struct cmachine {
    void *obj;
};

cmachine_t *machine_create();
void machine_destroy(cmachine_t *m);

cmachine_t *machine_create(char *data)
{
    cmachine_t *m;
    Machine *obj;
    
    m      = (typeof(m))malloc(sizeof(*m));
    obj    = new Machine(data);
    m->obj = obj;
    
    return m;
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
