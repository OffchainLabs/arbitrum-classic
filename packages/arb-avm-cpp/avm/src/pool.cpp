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

#include "avm/pool.hpp"

#include "avm/tuple.hpp"
#include "avm/value.hpp"

#include <ostream>

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

void TuplePool::returnResource(std::shared_ptr<RawTuple>&& object) {
    if (!shuttingDown) {
        resources[object->data.size()].push_back(std::move(object));
    }
}
