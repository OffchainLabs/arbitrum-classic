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

#include <boost/smart_ptr/make_local_shared.hpp>

#include <iostream>

/**
 * Returns instance of Resource.
 *
 * New resource will be created if all the resources
 * were used at the time of the request.
 *
 * @return Resource instance.
 */
boost::local_shared_ptr<std::vector<value>> TuplePool::getResource(int s) {
    if (resources[s].empty()) {
        auto newTup = boost::make_local_shared<std::vector<value>>();
        newTup->reserve(s);
        for (int i = 0; i < s; i++) {
            newTup->push_back(Tuple(0, this));
        }
        return newTup;
    } else {
        boost::local_shared_ptr<std::vector<value>> resource =
            resources[s].back();
        resources[s].pop_back();
        resource->clear();
        for (int i = 0; i < s; i++) {
            resource->push_back(Tuple(0, this));
        }
        return resource;
    }
}
