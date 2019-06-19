//
//  pool.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/23/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef pool_hpp
#define pool_hpp

#include "value.hpp"

#include <boost/smart_ptr/local_shared_ptr.hpp>

#include <stdio.h>
#include <list>
#include <array>
#include <vector>

class TuplePool
{
private:
    std::array<std::vector<boost::local_shared_ptr<std::vector<value>>>,9> resources;
public:
    /**
     * Returns instance of Resource.
     *
     * New resource will be created if all the resources
     * were used at the time of the request.
     *
     * @return Resource instance.
     */
    boost::local_shared_ptr<std::vector<value>> getResource(int s);

    /**
     * Return resource back to the pool.
     *
     * The resource must be initialized back to
     * the default settings before someone else
     * attempts to use it.
     *
     * @param object Resource instance.
     */
    void returnResource(boost::local_shared_ptr<std::vector<value>> && object) {
        resources[object->size()].push_back(std::move(object));
    }
};

#endif /* pool_hpp */
