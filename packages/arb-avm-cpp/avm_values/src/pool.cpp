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

#include <avm_values/pool.hpp>
#include <avm_values/tuple.hpp>
#include <avm_values/value.hpp>

#include <ostream>

RawTuple::~RawTuple() {
    thread_local std::deque<std::shared_ptr<RawTuple>> deletion_queue;
    thread_local bool deleting_shutdown = false;

    for (auto& item : data) {
        if (auto tup = get_if<Tuple>(&item)) {
            if (tup->tpl != nullptr) {
                deletion_queue.push_back(std::move(tup->tpl));
            }
        }
    }
    if (!deleting_shutdown) {
        deleting_shutdown = true;
        while (!deletion_queue.empty()) {
            deletion_queue.pop_front();
        }
        deleting_shutdown = false;
    }
}

/**
 * Returns instance of Resource.
 *
 * New resource will be created if all the resources
 * were used at the time of the request.
 *
 * @return Resource instance.
 */
std::shared_ptr<RawTuple> TuplePool::getResource(size_t s) {
    if (s == 0) {
        return nullptr;
    }
    auto resource = std::make_shared<RawTuple>();
    resource->data.reserve(s);
    return resource;
}

TuplePool& TuplePool::get_impl() {
    thread_local TuplePool singleton;
    return singleton;
}
