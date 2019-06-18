//
//  Machine.h
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef Machine_h
#define Machine_h

#include "stdint.h"

#ifdef __cplusplus
extern "C" {
#endif

void *machine_create(const char *filename, const char *inboxfile);
//void machine_destroy(cmachine_t *m);
//cassertion machine_run(cmachine_t *m, uint64_t maxSteps);
uint64_t machine_run(void *m, uint64_t maxSteps);


//struct cmachine;
//typedef struct cmachine cmachine_t;
//
//typedef struct {
//    uint64_t stepCount;
//} cassertion;
//
//cmachine_t *machine_create();
//void machine_destroy(cmachine_t *m);
//cassertion machine_run(cmachine_t *m, uint64_t maxSteps);

#ifdef __cplusplus
}
#endif

#endif /* Machine_h */
