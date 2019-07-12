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
    
    typedef struct {
        unsigned char *data;
        int length;
    } ByteSlice;
    
    typedef void CMachine;
    
    CMachine* machineCreate(const char* filename, const char* inboxfile);
    void machineDestroy(CMachine *m);
    
    // Ret must have 32 bytes of storage allocated for returned hash
    void machineHash(CMachine *m, char *ret);
    CMachine* machineClone(CMachine *m);
    
    // Ret must have 32 bytes of storage allocated for returned hash
    void machineInboxHash(CMachine *m, char *ret);
    
    bool machineHasPendingMessages(CMachine *m);
    void machineSendOnchainMessage(CMachine *m, char *data, int size);
    void machineDeliverOnchainMessage(CMachine *m);
    void machineSendOffchainMessages(CMachine *m, char *data, int size);
    
    uint64_t machineExecuteAssertion(CMachine* m, uint64_t maxSteps, uint64_t timeboundStart, uint64_t timeboundEnd);
    
    ByteSlice machineMarshallForProof(CMachine *m);

#ifdef __cplusplus
}
#endif

#endif /* Machine_h */
