/*
 * Copyright 2019, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef pool_hpp
#define pool_hpp

#include <avm/value.hpp>

#include <array>
#include <memory>
#include <vector>

struct RawTuple {
    std::vector<value> data;
    uint256_t cachedHash = 0;
};

class TuplePool {
   private:
    std::array<std::vector<std::shared_ptr<RawTuple>>, 9> resources;
    bool shuttingDown = false;

   public:
    ~TuplePool() { shuttingDown = true; }
    /**
     * Returns instance of Resource.
     *
     * New resource will be created if all the resources
     * were used at the time of the request.
     *
     * @return Resource instance.
     */
    std::shared_ptr<RawTuple> getResource(int s);

    /**
     * Return resource back to the pool.
     *
     * The resource must be initialized back to
     * the default settings before someone else
     * attempts to use it.
     *
     * @param object Resource instance.
     */
    void returnResource(std::shared_ptr<RawTuple>&& object);
};

#endif /* pool_hpp */
