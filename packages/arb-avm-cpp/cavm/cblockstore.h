//
//  cblockstore.h
//  avm
//
//  Created by Harry Kalodner on 5/27/20.
//

#ifndef cblockstore_h
#define cblockstore_h

#include "ctypes.h"

#include <stdio.h>

#ifdef __cplusplus
extern "C" {
#endif

void deleteBlockStore(CBlockStore* m);

int putBlock(CBlockStore* storage_ptr,
             const void* height,
             const void* hash,
             const void* data,
             int data_length);
int deleteBlock(CBlockStore* storage_ptr, const void* height, const void* hash);
ByteSliceResult getBlock(const CBlockStore* storage_ptr,
                         const void* height,
                         const void* hash);
HashList blockHashesAtHeight(const CBlockStore* storage_ptr,
                             const void* height);
int isBlockStoreEmpty(const CBlockStore* storage_ptr);
void* maxBlockStoreHeight(const CBlockStore* storage_ptr);
void* minBlockStoreHeight(const CBlockStore* storage_ptr);

#ifdef __cplusplus
}
#endif

#endif /* cblockstore_h */
