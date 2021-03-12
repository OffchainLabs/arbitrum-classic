/*
 * Copyright 2020, Offchain Labs, Inc.
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

#ifndef storageresultfwd_h
#define storageresultfwd_h

#include <rocksdb/status.h>

#include <variant>
#include <vector>

struct GetResults;
struct SaveResults;
struct DeleteResults;

template <typename T>
struct CountedData;

template <typename T>
using DbResult = std::variant<rocksdb::Status, CountedData<T>>;

template <typename T>
struct ValueResult;

using DataResults = ValueResult<std::vector<unsigned char>>;

#endif /* storageresultfwd_h */
