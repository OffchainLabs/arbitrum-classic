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

#include <variant>

#include <avm/checkpoint/machinestatesaver.hpp>
#include <avm/machinestate/tokenTracker.hpp>
#include <avm/value/codepoint.hpp>
#include <avm/value/tuple.hpp>

MachineStateSaver::MachineStateSaver(CheckpointStorage& storage)
    : checkpoint_storage(storage) {}

SaveResults MachineStateSaver::saveValue(const value& val) {
    auto serialized_value = checkpoint::utils::serializeValue(val);
    auto type = static_cast<ValueTypes>(serialized_value[0]);

    if (type == TUPLE) {
        auto tuple = nonstd::get<Tuple>(val);
        return saveTuple(tuple);
    } else {
        auto hash_key = GetHashKey(val);
        return checkpoint_storage.saveValue(hash_key, serialized_value);
    }
}

SaveResults MachineStateSaver::saveTuple(const Tuple& val) {
    auto hash_key = GetHashKey(val);
    auto results = checkpoint_storage.getValue(hash_key);

    auto incr_ref_count = results.status.ok() && results.reference_count > 0;

    if (incr_ref_count) {
        return checkpoint_storage.incrementReference(hash_key);
    } else {
        std::vector<unsigned char> value_vector{
            static_cast<unsigned char>(TUPLE)};

        for (uint64_t i = 0; i < val.tuple_size(); i++) {
            auto current_val = val.get_element(i);
            auto serialized_val =
                checkpoint::utils::serializeValue(current_val);

            value_vector.insert(value_vector.end(), serialized_val.begin(),
                                serialized_val.end());

            auto type = static_cast<ValueTypes>(serialized_val[0]);
            if (type == TUPLE) {
                auto tup_val = nonstd::get<Tuple>(current_val);
                auto tuple_save_results = saveTuple(tup_val);
            }
        }
        return checkpoint_storage.saveValue(hash_key, value_vector);
    }
};

SaveResults MachineStateSaver::saveMachineState(
    ParsedState state_data,
    const std::vector<unsigned char>& checkpoint_name) {
    auto serialized_state = checkpoint::utils::serializeState(state_data);
    return checkpoint_storage.saveValue(checkpoint_name, serialized_state);
}
