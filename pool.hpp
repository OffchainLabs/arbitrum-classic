//
//  pool.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/23/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef pool_hpp
#define pool_hpp

#include <stdio.h>
#include <list>
#include <array>
#include "uint256_t.h"

using namespace std;


class value;
//class vTuple;
class vTuple{
public:
    int ref;
    value* vals;
};

/* Note, that this class is a singleton. */
class ObjectPool
{
private:
    std::array<std::list<vTuple*>,9> resources;

    static ObjectPool* instance;
    ObjectPool() {}
public:
    /**
     * Static method for accessing class instance.
     * Part of Singleton design pattern.
     *
     * @return ObjectPool instance.
     */
    static ObjectPool* getInstance();
    /**
     * Returns instance of Resource.
     *
     * New resource will be created if all the resources
     * were used at the time of the request.
     *
     * @return Resource instance.
     */
    vTuple* getResource(int s);
    /**
     * Return resource back to the pool.
     *
     * The resource must be initialized back to
     * the default settings before someone else
     * attempts to use it.
     *
     * @param object Resource instance.
     */
    void returnResource(int size, vTuple* object);
};

#endif /* pool_hpp */
