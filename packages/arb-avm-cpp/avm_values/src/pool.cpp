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

#include <ostream>

void tupleDeleter(RawTuple* p) {
    auto& deleter = TuplePool::get_impl();
    if (!deleter.shuttingDown) {
        return deleter.deleteTuple({p, tupleDeleter});
    }

    static std::vector<std::unique_ptr<RawTuple>> deletion_queue;
    static bool deleting_shutdown = false;

    deletion_queue.push_back(std::unique_ptr<RawTuple>{p});
    if (!deleting_shutdown) {
        deleting_shutdown = true;
        while (!deletion_queue.empty()) {
            deletion_queue.pop_back();
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
    std::shared_ptr<RawTuple> resource;
    if (resources[s].empty()) {
        resource = {new RawTuple{}, tupleDeleter};
    } else {
        resource = {std::move(resources[s].back())};
        resources[s].pop_back();
    }
    resource->data.clear();
    resource->data.reserve(s);
    resource->deferredHashing = true;
    return resource;
}

void TuplePool::deleteTuple(UniqueTuple tup) {
    delete_list.push_front(std::move(tup));
    if (!deleting) {
        deleting = true;
        while (!delete_list.empty()) {
            auto& item = delete_list.back();
            item->data.clear();
            if (!shuttingDown) {
                resources[item->data.capacity()].push_back(std::move(item));
            } else {
                // Clear out the custom deleter
                std::unique_ptr<RawTuple>{item.release()};
            }
            delete_list.pop_back();
        }
        deleting = false;
    }
}

TuplePool& TuplePool::get_impl() {
    thread_local TuplePool singleton;
    return singleton;
}
