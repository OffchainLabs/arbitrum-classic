//
//  pool.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/23/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//
#include <iostream>
#include <list>
#include <array>
#include "pool.hpp"
#include "value.hpp"

/* Note, that this class is a singleton. */
/**
 * Static method for accessing class instance.
 * Part of Singleton design pattern.
 *
 * @return ObjectPool instance.
 */
ObjectPool* ObjectPool::getInstance()
{
    if (instance == 0)
    {
        instance = new ObjectPool;
    }
    return instance;
}
/**
 * Returns instance of Resource.
 *
 * New resource will be created if all the resources
 * were used at the time of the request.
 *
 * @return Resource instance.
 */
vTuple* ObjectPool::getResource(int s)
{
    if (resources[s].empty())
    {
        vTuple *ret=new vTuple;
        ret->ref=1;
        ret->vals = new value[s];
        return ret;
    }
    else
    {
        vTuple* resource = resources[s].front();
        resources[s].pop_front();
        return resource;
    }
}
/**
 * Return resource back to the pool.
 *
 * The resource must be initialized back to
 * the default settings before someone else
 * attempts to use it.
 *
 * @param object Resource instance.
 */

void ObjectPool::returnResource(int size, vTuple* object)
{
    if (object->ref > 1){
        object->ref--;
    } else {
        memset(object->vals,0,size*sizeof(value));
        resources[size].push_back(object);
    }
}
ObjectPool* ObjectPool::instance = 0;


