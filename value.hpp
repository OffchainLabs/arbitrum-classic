//
//  value.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/25/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef value_h
#define value_h

#include "uint256_t.h"
#include "pool.hpp"

enum types{ NONE, NUM, CODEPT, TUPLE };

class value{
private:
    ObjectPool* tuplePool = ObjectPool::getInstance();

public:
    int type;
    int size;
    int tplsize;
    int refcount;
    //    uint256_t reg;
    uint256_t num;
    vTuple *tpl;
    
    value();
    value(const value &obj);
    value(uint256_t val, int type);
    value(int s);
    ~value();
    value &operator = (const value &val );
    void reset();
    void set_num(uint256_t val);
    void set_codept(uint256_t val);
    value *dup();

    int set_tuple_elem(int pos, value *newval);
    
    value *get_tuple_elem(int pos);
    void print();
    void printstack();
};

#endif /* value_h */
