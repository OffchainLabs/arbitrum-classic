//
//  cnodestore.h
//  avm
//
//  Created by Harry Kalodner on 5/27/20.
//

#ifndef cnodestore_h
#define cnodestore_h

#include "ctypes.h"

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void deleteNodeStore(CNodeStore* m);
int putNode(CNodeStore* node_store,
            uint64_t height,
            const void* hash,
            const void* data,
            int data_length);
ByteSliceResult getNode(CNodeStore* node_store,
                        uint64_t height,
                        const void* hash);
Uint64Result getNodeHeight(CNodeStore* node_store, const void* hash);
HashResult getNodeHash(CNodeStore* node_store, uint64_t height);
uint64_t longestNodeChainCount(CNodeStore* node_store);

#ifdef __cplusplus
}
#endif

#endif /* cnodestore_h */
