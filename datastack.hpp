//
//  datastack.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef datastack_hpp
#define datastack_hpp

#include <stdio.h>
#include <stack>
#include <unordered_map>
#include "pool.hpp"
#include "value.hpp"

class datastack{
public:
    
    std::vector<value> basedatastack;
    std::unordered_map<uint64_t, uint64_t> pcmap;
    unsigned int size;
    
    datastack() {
        basedatastack.reserve(1000);
    }
    
    void push(value && newdata);
    value pop();
    value &peek();
    void popSet(value &val);
    uint64_t stacksize();
    
    void pcpush(uint64_t i, uint64_t j);
    std::pair<int,uint64_t> jmp();
    int tget();
    int tset();
};

#endif /* datastack_hpp */
