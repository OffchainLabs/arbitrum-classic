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

#include <avm_values/tuplestub.hpp>
#include <avm_values/value.hpp>

#include <array>
#include <deque>
#include <memory>
#include <vector>

struct RawTuple {
    HashPreImage cachedPreImage;
    std::vector<value> data;
    bool deferredHashing = true;

    RawTuple() : cachedPreImage({}, 0), deferredHashing(true) {}
};

using UniqueTuple = std::unique_ptr<RawTuple, void (*)(RawTuple*)>;

class TuplePool {
   private:
    std::array<std::vector<UniqueTuple>, 9> resources;
    bool shuttingDown = false;

    bool deleting = false;
    std::deque<UniqueTuple> delete_list;

    void deleteTuple(UniqueTuple tup);

    friend void tupleDeleter(RawTuple* p);

   public:
    static TuplePool& get_impl();

    ~TuplePool() { shuttingDown = true; }
    /**
     * Returns instance of Resource.
     *
     * New resource will be created if all the resources
     * were used at the time of the request.
     *
     * @return Resource instance.
     */
    std::shared_ptr<RawTuple> getResource(size_t s);
};

#endif /* pool_hpp */
