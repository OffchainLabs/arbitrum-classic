//
//  pool.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/23/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include <avm/pool.hpp>

#include <avm/tuple.hpp>
#include <avm/value.hpp>

#include <iostream>

/**
 * Returns instance of Resource.
 *
 * New resource will be created if all the resources
 * were used at the time of the request.
 *
 * @return Resource instance.
 */
std::shared_ptr<RawTuple> TuplePool::getResource(int s) {
    if (s == 0) {
        return nullptr;
    }
    std::shared_ptr<RawTuple> resource;
    if (resources[s].empty()) {
        resource = std::make_shared<RawTuple>();
    } else {
        resource = resources[s].back();
        resources[s].pop_back();
    }
    resource->data.clear();
    resource->data.reserve(s);
    return resource;
}
