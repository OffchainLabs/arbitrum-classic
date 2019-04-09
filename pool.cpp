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

/**
 * Returns instance of Resource.
 *
 * New resource will be created if all the resources
 * were used at the time of the request.
 *
 * @return Resource instance.
 */
std::shared_ptr<std::vector<value>> TuplePool::getResource(int s)
{
    if (resources[s].empty())
    {
        auto newTup = std::make_shared<std::vector<value>>();
        newTup->reserve(s);
        for (int i = 0; i < s; i++) {
            newTup->push_back(Tuple(0, this));
        }
        return newTup;
    }
    else
    {
        std::shared_ptr<std::vector<value>> resource = resources[s].back();
        resources[s].pop_back();
        resource->clear();
        for (int i = 0; i < s; i++) {
            resource->push_back(Tuple(0, this));
        }
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

void TuplePool::returnResource(std::shared_ptr<std::vector<value>> && object) {
    if (object.use_count() == 1){
        resources[object->size()].push_back(std::move(object));
    }
}
