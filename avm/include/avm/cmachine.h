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
        void *data;
        int length;
    } ByteSlice;
    
    typedef struct {
        unsigned char *outMessageData;
        int outMessageLength;
        unsigned char *logData;
        int logLength;
        uint64_t numSteps;
    } RawAssertion;
    
    typedef void CMachine;
    
    CMachine* machineCreate(const char* filename);
    void machineDestroy(CMachine *m);
    
    // Ret must have 32 bytes of storage allocated for returned hash
    void machineHash(CMachine *m, void *ret);
    CMachine* machineClone(CMachine *m);
    
    // Ret must have 32 bytes of storage allocated for returned hash
    void machineInboxHash(CMachine *m, void *ret);
    
    int machineHasPendingMessages(CMachine *m);
    void machineSendOnchainMessage(CMachine *m, void *data);
    void machineDeliverOnchainMessages(CMachine *m);
    void machineSendOffchainMessages(CMachine *m, void *data, int size);
    
    RawAssertion machineExecuteAssertion(CMachine* m, uint64_t maxSteps, uint64_t timeboundStart, uint64_t timeboundEnd);
    
    ByteSlice machineMarshallForProof(CMachine *m);

#ifdef __cplusplus
}
#endif

#endif /* Machine_h */
