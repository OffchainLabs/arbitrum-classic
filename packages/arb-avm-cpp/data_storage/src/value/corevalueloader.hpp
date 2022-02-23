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

#ifndef corevalueloader_hpp
#define corevalueloader_hpp

#include <avm_values/code.hpp>
#include <avm_values/value.hpp>
#include <avm_values/valueloader.hpp>
#include <data_storage/datastorage.hpp>
#include <data_storage/value/valuecache.hpp>

class CoreValueLoader : public AbstractValueLoader {
   public:
    CoreValueLoader(std::shared_ptr<DataStorage>,
                    std::shared_ptr<CoreCode>,
                    ValueCache);

    Value loadValue(const uint256_t& hash) override;

    std::unique_ptr<AbstractValueLoader> clone() const override;

   protected:
    std::shared_ptr<DataStorage> data_storage;
    std::shared_ptr<CoreCode> core_code;
    ValueCache cache;
};

#endif /* corevalueloader_hpp */
