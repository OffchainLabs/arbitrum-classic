//
//  datastack.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef datastack_hpp
#define datastack_hpp

#include "pool.hpp"
#include "tuple.hpp"
#include "value.hpp"

#include <stdio.h>
#include <stack>
#include <unordered_map>

class datastack{
public:
    
    std::vector<value> basedatastack;
    unsigned int size;
    
    datastack() {
        basedatastack.reserve(1000);
    }
    
    void push(value && newdata) {
        basedatastack.push_back(std::move(newdata));
    };
    
    value pop() {
        auto stackEmpty = basedatastack.empty();
        if (stackEmpty){
            throw std::runtime_error("Stack is empty");
        }
        
        auto val = std::move(basedatastack.back());
        basedatastack.pop_back();
        return val;
    }
    
    value &peek() {
        if (basedatastack.size()==0){
            throw std::runtime_error("Stack is empty");
        }
        
        return basedatastack.back();
    }
    
    void popSet(value &val) {
        if (basedatastack.size()==0){
            throw std::runtime_error("Stack is empty");
        }
        
        val = std::move(basedatastack.back());
        basedatastack.pop_back();
    }
    
    uint64_t stacksize() {
        return basedatastack.size();
    }
};

#endif /* datastack_hpp */
