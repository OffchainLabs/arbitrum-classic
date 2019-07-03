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

    void* machine_create(const char* filename, const char* inboxfile);
    void machine_destroy(void *m);
    void* machine_clone(void *m);
    uint64_t machine_run(void* m, uint64_t maxSteps);
    uint64_t machine_run_until_stop(void* m);
    
    void machine_add_to_inbox(void *m, char *inbox);
    void machine_set_time_bounds(void* m, uint64_t timeboundStart, uint64_t timeboundEnd);

#ifdef __cplusplus
}
#endif

#endif /* Machine_h */
