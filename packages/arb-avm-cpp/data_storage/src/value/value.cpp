/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

#include <data_storage/value/value.hpp>

#include "referencecount.hpp"
#include "utils.hpp"

#include <data_storage/datastorage.hpp>
#include <data_storage/storageresult.hpp>

#include <avm_values/tuple.hpp>

constexpr int TUP_TUPLE_LENGTH = 33;
constexpr int TUP_NUM_LENGTH = 33;
constexpr int TUP_CODEPT_LENGTH = 49;

namespace {
std::vector<unsigned char> getHashKey(const value& val) {
    auto hash_key = hash_value(val);
    std::vector<unsigned char> hash_key_vector;
    marshal_uint256_t(hash_key, hash_key_vector);

    return hash_key_vector;
}

struct SerializedTupStub {
    uint8_t size;
    uint256_t hash;
};

using ParsedTupVal =
    nonstd::variant<uint256_t, CodePointStub, SerializedTupStub>;

std::vector<ParsedTupVal> parseTuple(const std::vector<unsigned char>& data) {
    std::vector<ParsedTupVal> return_vector;

    auto iter = data.begin();
    uint8_t count = *iter - TUPLE;
    ++iter;

    for (uint8_t i = 0; i < count; i++) {
        auto value_type = static_cast<ValueTypes>(*iter);
        auto buf = reinterpret_cast<const char*>(&*iter);
        ++buf;

        switch (value_type) {
            case NUM: {
                return_vector.push_back(deserializeUint256t(buf));
                iter += TUP_NUM_LENGTH;
                break;
            }
            case CODE_POINT_STUB: {
                return_vector.push_back(deserializeCodePointStub(buf));
                iter += TUP_CODEPT_LENGTH;
                break;
            }
            case HASH_PRE_IMAGE: {
                throw std::runtime_error("HASH_ONLY item");
            }
            default: {
                uint8_t tup_size = value_type - TUPLE;
                if (tup_size > 8) {
                    throw std::runtime_error(
                        "tried to parse tuple with invalid typecode");
                }
                return_vector.push_back(
                    SerializedTupStub{tup_size, deserializeUint256t(buf)});
                iter += TUP_TUPLE_LENGTH;
                break;
            }
        }
    }
    return return_vector;
}

DbResult<value> getTuple(const Transaction& transaction,
                         const GetResults& results,
                         std::set<uint64_t>& segment_ids);

struct TupleGetter {
    const Transaction& transaction;
    std::set<uint64_t>& segment_ids;

    DbResult<value> operator()(const uint256_t& val) const {
        return {rocksdb::Status::OK(), 1, val};
    }

    DbResult<value> operator()(const CodePointStub& val) const {
        segment_ids.insert(val.pc.segment);
        return {rocksdb::Status::OK(), 1, val};
    }

    DbResult<value> operator()(const SerializedTupStub& val) const {
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(val.hash, hash_key);
        auto key = vecToSlice(hash_key);
        auto results = getRefCountedData(*transaction.transaction, key);

        if (!results.status.ok()) {
            return DbResult<value>{results.status, results.reference_count,
                                   Tuple()};
        }
        return getTuple(transaction, results, segment_ids);
    }
};

DbResult<value> getTuple(const Transaction& transaction,
                         const GetResults& results,
                         std::set<uint64_t>& segment_ids) {
    std::vector<value> values;
    for (auto& current_vector : parseTuple(results.stored_value)) {
        auto val = nonstd::visit(TupleGetter{transaction, segment_ids},
                                 current_vector);
        if (!val.status.ok()) {
            return DbResult<value>{val.status, val.reference_count, Tuple()};
        }
        values.push_back(std::move(val.data));
    }
    auto tuple = Tuple(std::move(values));
    return DbResult<value>{results.status, results.reference_count, tuple};
}

struct ValueSerializer {
    std::vector<unsigned char>& value_vector;
    std::map<uint64_t, uint64_t>& segment_counts;

    void operator()(const Tuple& val) const {
        value_vector.push_back(TUPLE + val.tuple_size());
        auto hash_key = hash_value(val);
        marshal_uint256_t(hash_key, value_vector);
    }

    void operator()(const uint256_t& val) const {
        value_vector.push_back(NUM);
        marshal_uint256_t(val, value_vector);
    }

    void operator()(const CodePointStub& val) const {
        value_vector.push_back(CODE_POINT_STUB);
        val.marshal(value_vector);
        ++segment_counts[val.pc.segment];
    }

    void operator()(const HashPreImage& val) const {
        value_vector.push_back(HASH_PRE_IMAGE);
        val.marshal(value_vector);
    }
};

SaveResults saveTuple(Transaction& transaction,
                      const Tuple& orig_val,
                      std::map<uint64_t, uint64_t>& segment_counts) {
    std::vector<Tuple> tups{orig_val};
    bool first = true;
    SaveResults save_ret;
    while (!tups.empty()) {
        auto tup = std::move(tups.back());
        tups.pop_back();
        auto hash_key = getHashKey(tup);
        auto key = vecToSlice(hash_key);
        auto results = getRefCountedData(*transaction.transaction, key);
        SaveResults ret;
        if (results.status.ok() && results.reference_count > 0) {
            ret = incrementReference(*transaction.transaction, key);
        } else {
            std::vector<unsigned char> value_vector;
            value_vector.push_back(TUPLE + tup.tuple_size());
            for (uint64_t i = 0; i < tup.tuple_size(); i++) {
                auto current_val = tup.get_element(i);
                nonstd::visit(ValueSerializer{value_vector, segment_counts},
                              current_val);
                if (nonstd::holds_alternative<Tuple>(current_val)) {
                    tups.push_back(current_val.get<Tuple>());
                }
            }
            ret =
                saveRefCountedData(*transaction.transaction, key, value_vector);
        }
        if (first) {
            save_ret = ret;
            first = false;
        }
    }
    return save_ret;
}

struct ValueSaver {
    Transaction& transaction;
    std::map<uint64_t, uint64_t>& segment_counts;

    template <typename T>
    SaveResults saveImpl(const T& val, bool allow_replacement) const {
        std::vector<unsigned char> serialized_value;
        ValueSerializer{serialized_value, segment_counts}(val);
        auto hash_key = getHashKey(val);
        auto key = vecToSlice(hash_key);
        return saveRefCountedData(*transaction.transaction, key,
                                  serialized_value, 1, allow_replacement);
    }

    SaveResults operator()(const Tuple& val) const {
        return saveTuple(transaction, val, segment_counts);
    }

    SaveResults operator()(const CodePointStub& val) const {
        // The same code point can exist in different segments with different
        // serializations mapping to the same hash. If this occurs, the
        // different versions are interchangeable
        return saveImpl(val, true);
    }

    template <typename T>
    SaveResults operator()(const T& val) const {
        return saveImpl(val, false);
    }
};

DeleteResults deleteTuple(Transaction& transaction,
                          const uint256_t& tuple_hash,
                          std::map<uint64_t, uint64_t>& segment_counts) {
    std::vector<uint256_t> tuples_to_delete{tuple_hash};
    DeleteResults ret;
    while (!tuples_to_delete.empty()) {
        auto tup_hash = std::move(tuples_to_delete.back());
        tuples_to_delete.pop_back();
        std::vector<unsigned char> hash_key;
        marshal_uint256_t(tup_hash, hash_key);
        auto key = vecToSlice(hash_key);

        auto results = getRefCountedData(*transaction.transaction, key);
        if (!results.status.IsNotFound()) {
            // value was already deleted, this shouldn't happen, but we can
            // continue
            continue;
        }

        if (!results.status.ok()) {
            // Some unexpected error occured
            return DeleteResults{0, results.status};
        }

        if (results.reference_count == 1) {
            // This was the last reference to this tuple so decrement the
            // reference count of the references values
            for (const auto& nested : parseTuple(results.stored_value)) {
                if (nonstd::holds_alternative<SerializedTupStub>(nested)) {
                    tuples_to_delete.push_back(
                        nested.get<SerializedTupStub>().hash);
                } else if (nonstd::holds_alternative<CodePointStub>(nested)) {
                    segment_counts[nested.get<CodePointStub>().pc.segment]++;
                }
            }
        }

        ret = deleteRefCountedData(*transaction.transaction, key);
    }
    return ret;
}
}  // namespace

DbResult<value> getValueImpl(const Transaction& transaction,
                             uint256_t value_hash,
                             std::set<uint64_t>& segment_ids) {
    std::vector<unsigned char> hash_key;
    marshal_uint256_t(value_hash, hash_key);
    auto key = vecToSlice(hash_key);
    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        return DbResult<value>{results.status, results.reference_count,
                               Tuple()};
    }

    auto buf = reinterpret_cast<const char*>(results.stored_value.data());
    auto value_type = static_cast<ValueTypes>(*buf);
    ++buf;

    switch (value_type) {
        case NUM: {
            auto val = deserializeUint256t(buf);
            return DbResult<value>{results.status, results.reference_count,
                                   val};
        }
        case CODE_POINT_STUB: {
            auto code_point = deserializeCodePointStub(buf);
            segment_ids.insert(code_point.pc.segment);
            return DbResult<value>{results.status, results.reference_count,
                                   code_point};
        }
        case HASH_PRE_IMAGE: {
            throw std::runtime_error("HASH_ONLY item");
        }
        default: {
            if (value_type - TUPLE > 8) {
                throw std::runtime_error("can't get value with invalid type");
            }
            return getTuple(transaction, results, segment_ids);
        }
    }
}

DbResult<value> getValue(const Transaction& transaction, uint256_t value_hash) {
    std::set<uint64_t> segment_ids;
    return getValueImpl(transaction, value_hash, segment_ids);
}

SaveResults saveValueImpl(Transaction& transaction,
                          const value& val,
                          std::map<uint64_t, uint64_t>& segment_counts) {
    return nonstd::visit(ValueSaver{transaction, segment_counts}, val);
}

SaveResults saveValue(Transaction& transaction, const value& val) {
    std::map<uint64_t, uint64_t> segment_counts;
    return saveValueImpl(transaction, val, segment_counts);
}

DeleteResults deleteValueImpl(Transaction& transaction,
                              const uint256_t& value_hash,
                              std::map<uint64_t, uint64_t>& segment_counts) {
    std::vector<unsigned char> hash_key;
    marshal_uint256_t(value_hash, hash_key);
    auto key = vecToSlice(hash_key);

    auto results = getRefCountedData(*transaction.transaction, key);

    if (!results.status.ok()) {
        return DeleteResults{0, results.status};
    }

    auto ret = deleteRefCountedData(*transaction.transaction, key);

    auto value_type = static_cast<ValueTypes>(*results.stored_value.data());

    if (results.reference_count == 1 &&
        (value_type >= TUPLE && value_type <= TUPLE + 8)) {
        // The value we deleted was a tuple
        std::vector<uint256_t> tuples_to_delete;
        for (const auto& nested : parseTuple(results.stored_value)) {
            if (nonstd::holds_alternative<SerializedTupStub>(nested)) {
                deleteTuple(transaction, nested.get<SerializedTupStub>().hash,
                            segment_counts);
            } else if (nonstd::holds_alternative<CodePointStub>(nested)) {
                segment_counts[nested.get<CodePointStub>().pc.segment]++;
            }
        }
    }

    return ret;
}

DeleteResults deleteValue(Transaction& transaction, uint256_t value_hash) {
    std::map<uint64_t, uint64_t> segment_counts;
    return deleteValueImpl(transaction, value_hash, segment_counts);
}
