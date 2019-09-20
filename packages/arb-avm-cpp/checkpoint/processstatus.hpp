//
//  processstatus.hpp
//  avm
//
//  Created by Minh Truong on 9/19/19.
//

#ifndef processstatus_hpp
#define processstatus_hpp

#include <stdio.h>
#include "rocksdb/db.h"

struct ProcessStatus {
    rocksdb::Status status;
    std::string string_value;
};

#endif /* processstatus_hpp */
