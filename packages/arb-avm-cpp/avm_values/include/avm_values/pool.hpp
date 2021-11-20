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
    std::vector<Value> data;
    bool deferredHashing = true;

    RawTuple() : cachedPreImage({}, 0), deferredHashing(true) {}
    RawTuple(const RawTuple&) = delete;
    RawTuple& operator=(const RawTuple&) = delete;
    ~RawTuple();
};

void tupleDeleter(RawTuple* p);

using UniqueTuple = std::unique_ptr<RawTuple, void (*)(RawTuple*)>;

class TuplePool {
   public:
    static TuplePool& get_impl();

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
