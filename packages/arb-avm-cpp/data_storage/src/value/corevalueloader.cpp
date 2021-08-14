/*
 * Copyright 2021, Offchain Labs, Inc.
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

#include "corevalueloader.hpp"

#include <data_storage/value/value.hpp>

CoreValueLoader::CoreValueLoader(std::shared_ptr<DataStorage> data_storage_,
                                 ValueCache cache_)
    : data_storage(data_storage_), cache(cache_) {}

value CoreValueLoader::loadValue(const uint256_t& hash) {
    ReadTransaction tx(data_storage);
    auto res = getValue(tx, hash, cache, true);
    if (auto status = std::get_if<rocksdb::Status>(&res)) {
        throw std::runtime_error(std::string("Value loading failed: ") +
                                 status->ToString());
    }
    return std::get<CountedData<value>>(res).data;
}

std::unique_ptr<AbstractValueLoader> CoreValueLoader::clone() const {
    return std::make_unique<CoreValueLoader>(data_storage, cache);
}
