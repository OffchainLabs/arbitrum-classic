/*
 * Copyright 2022, Offchain Labs, Inc.
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

#include <data_storage/arbstorage.hpp>

#include <nlohmann/json.hpp>

#include <fstream>
#include <iostream>
#include <string>

constexpr size_t parallelKvsLayers = 2;  // 8**x is the total parallelism

// Warning: doesn't restore code segments or cache values
class SimpleValueLoader : public AbstractValueLoader {
   public:
    SimpleValueLoader(std::shared_ptr<DataStorage> data_storage_)
        : data_storage(data_storage_) {}

    Value loadValue(const uint256_t& hash) override {
        ReadTransaction tx(data_storage);
        std::set<uint64_t> segment_ids;
        ValueCache cache{1, 0};
        auto res = getValueImpl(tx, hash, segment_ids, cache, true);
        if (auto status = std::get_if<rocksdb::Status>(&res)) {
            throw std::runtime_error(std::string("Value loading failed: ") +
                                     status->ToString());
        }
        return std::get<CountedData<Value>>(res).data;
    }

    [[nodiscard]] std::unique_ptr<AbstractValueLoader> clone() const override {
        return std::make_unique<SimpleValueLoader>(data_storage);
    }

   protected:
    std::shared_ptr<DataStorage> data_storage;
};

Tuple resolveTuple(ValueLoader loader, Value val) {
    if (auto uv = get_if<UnloadedValue>(&val)) {
        val = loader.loadValue(uv->hash());
    }
    if (auto tup = get_if<Tuple>(&val)) {
        return *tup;
    }
    throw std::runtime_error("got unexpected type looking for tuple");
}

Tuple indexTup(ValueLoader loader, Tuple src, uint64_t index) {
    return resolveTuple(loader, src.get_element(index));
}

uint256_t assertInt(Value val) {
    if (auto x = get_if<uint256_t>(&val)) {
        return *x;
    }
    throw std::runtime_error("got unexpected type looking for int");
}

uint256_t indexInt(Tuple src, uint64_t index) {
    auto val = src.get_element(index);
    return assertInt(val);
}

Buffer indexBuffer(ValueLoader loader, Tuple src, uint64_t index) {
    auto val = src.get_element(index);
    if (auto uv = get_if<UnloadedValue>(&val)) {
        val = loader.loadValue(uv->hash());
    }
    if (auto buf = get_if<Buffer>(&val)) {
        return *buf;
    }
    throw std::runtime_error("got unexpected type looking for buffer");
}

template <typename F>
void kvsNodeForAll(ValueLoader loader,
                   Value node,
                   bool isStorageMap,
                   size_t parallelLayers,
                   F&& func) {
    if (auto x = get_if<uint256_t>(&node)) {
        if (*x == 0) {
            return;
        } else {
            throw std::runtime_error("kvs node is unexpected integer");
        }
    }
    auto tup = resolveTuple(loader, node);
    auto len = tup.tuple_size();
    if (len == 2) {
        auto key = tup.get_element(0);
        if (isStorageMap) {
            func(key, tup.get_element(1));
        } else {
            auto val = indexTup(loader, tup, 1);
            if (indexInt(val, 0)) {
                func(key, val.get_element(1));
            }
        }
    } else if (len == 8) {
        if (parallelLayers > 0) {
            std::vector<std::thread> threads;
            for (size_t i = 0; i < 8; i++) {
                threads.emplace_back([&, i]() {
                    kvsNodeForAll(loader, tup.get_element(i), isStorageMap,
                                  parallelLayers - 1, std::forward<F>(func));
                });
            }
            for (auto& thread : threads) {
                thread.join();
            }
        } else {
            for (size_t i = 0; i < 8; i++) {
                kvsNodeForAll(loader, tup.get_element(i), isStorageMap, 0,
                              std::forward<F>(func));
            }
        }
    } else {
        throw std::runtime_error("got unexpected kvs tuple length");
    }
}

template <typename F>
void kvsForAll(ValueLoader loader, Tuple kvs, F&& func) {
    kvsNodeForAll(loader, kvs.get_element(0), false, parallelKvsLayers,
                  std::forward<F>(func));
}

template <typename F>
void storageMapForAll(ValueLoader loader, Tuple map, F&& func) {
    kvsNodeForAll(loader, map.get_element(0), true, 0,
                  [&](Value key, Value value) {
                      func(assertInt(key), assertInt(value));
                  });
}

std::string hexLengthString(uint256_t x, size_t len) {
    auto hex = intx::to_string(x, 16);
    std::string prefix = "0x";
    for (auto i = 0; i < len * 2 - hex.size(); i++) {
        prefix += "0";
    }
    return prefix + hex;
}

std::string hashString(uint256_t hash) {
    return hexLengthString(hash, 32);
}

std::string addressString(uint256_t addr) {
    return hexLengthString(addr, 20);
}

std::vector<uint8_t> serializeBytes(ValueLoader loader, Tuple tup) {
    auto len = indexInt(tup, 0);
    auto offset = indexInt(tup, 1);
    auto buffer = indexBuffer(loader, tup, 2);
    auto bytes = buffer.toFlatVector();
    if (offset >= bytes.size()) {
        return std::vector<uint8_t>();
    }
    bytes.erase(bytes.begin(), bytes.begin() + size_t(offset));
    bytes.resize(size_t(len));
    return bytes;
}

nlohmann::json serializeRetryable(ValueLoader loader, Value retryable) {
    nlohmann::json json;
    auto tup = resolveTuple(loader, retryable);
    json["Id"] = hashString(indexInt(tup, 0));
    json["From"] = addressString(indexInt(tup, 1));
    json["To"] = addressString(indexInt(tup, 2));
    json["Callvalue"] = intx::to_string(indexInt(tup, 3));
    json["Beneficiary"] = addressString(indexInt(tup, 5));
    json["Calldata"] = serializeBytes(loader, indexTup(loader, tup, 6));
    auto rem = indexTup(loader, tup, 7);
    json["Timeout"] = uint64_t(indexInt(rem, 0));
    return json;
}

nlohmann::json serializeAccount(ValueLoader loader, Value account) {
    nlohmann::json json;
    auto tup = resolveTuple(loader, account);
    json["Addr"] = addressString(indexInt(tup, 0));
    json["Nonce"] = uint64_t(indexInt(tup, 2));
    json["Balance"] = intx::to_string(indexInt(tup, 3));

    auto contractInfo = indexTup(loader, tup, 4);
    nlohmann::json contractJson = nullptr;
    if (indexInt(contractInfo, 0)) {
        contractInfo = indexTup(loader, contractInfo, 1);
        contractJson = nlohmann::json();
        contractJson["Code"] =
            serializeBytes(loader, indexTup(loader, contractInfo, 1));
        nlohmann::json storage;
        storageMapForAll(loader, indexTup(loader, contractInfo, 4),
                         [&](uint256_t key, uint256_t value) {
                             storage[hashString(key)] = hashString(value);
                         });
        contractJson["ContractStorage"] = storage;
    }
    json["ContractInfo"] = contractJson;

    auto aggregatorInfo = indexTup(loader, tup, 5);
    nlohmann::json aggregatorInfoJson = nullptr;
    if (indexInt(aggregatorInfo, 0)) {
        aggregatorInfo = indexTup(loader, aggregatorInfo, 1);
        aggregatorInfoJson = nlohmann::json();
        aggregatorInfoJson["FeeCollector"] =
            addressString(indexInt(aggregatorInfo, 1));
        aggregatorInfoJson["BaseFeeL1Gas"] =
            intx::to_string(indexInt(aggregatorInfo, 2));
    }
    json["AggregatorInfo"] = aggregatorInfoJson;

    auto aggregatorToPay = indexTup(loader, tup, 6);
    nlohmann::json aggregatorToPayJson = nullptr;
    if (indexInt(aggregatorToPay, 0)) {
        aggregatorToPayJson = addressString(indexInt(aggregatorToPay, 1));
    }
    json["AggregatorToPay"] = aggregatorToPayJson;

    return json;
}

template <typename F>
void writeKvsToFile(ValueLoader loader,
                    Tuple kvs,
                    std::string name,
                    F&& serialize) {
    std::cerr << "Serializing " << name << "..." << std::endl;
    std::mutex mutex;
    std::ofstream retryables;
    retryables.open(name + ".json", std::ios::out | std::ios::trunc);
    uint64_t count = 0;
    kvsForAll(loader, kvs, [&](Value key, Value val) {
        auto serialized = serialize(loader, val);
        mutex.lock();
        retryables << serialized << std::endl;
        count++;
        mutex.unlock();
    });
    std::cout << "Finished serializing " << count << " " << name << std::endl;
    retryables.close();
}

int main(int argc, char* argv[]) {
    if (argc < 3) {
        std::cout << "Usage: \narbcore_runner dbpath arbospath\n";
        return 1;
    }
    auto dbpath = std::string(argv[1]);
    auto arbospath = std::string(argv[2]);
    ArbCoreConfig coreConfig{};
    coreConfig.lazy_load_core_machine = true;

    std::cout << "Loading db\n";
    ArbStorage storage{dbpath, coreConfig};
    std::cout << "Initializing arbstorage\n";
    auto result = storage.initialize(arbospath);
    if (result.finished) {
        // Nothing left to do
        return 0;
    }
    if (!result.status.ok()) {
        std::cerr << "Failed to initialize storage" << result.status.ToString()
                  << std::endl;
        return -1;
    }
    auto core = storage.getArbCore();
    auto msgCount = core->messageEntryInsertedCount();
    if (!msgCount.status.ok()) {
        std::cerr << "Failed to get message count" << msgCount.status.ToString()
                  << std::endl;
        return -1;
    }
    while (core->getLastMachineOutput().fully_processed_inbox.count <
           msgCount.data) {
        std::cerr << "Waiting for core machine to catch up..." << std::endl;
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }
    auto mach = core->getLastMachine();
    std::cout << "At L2 block " << mach->machine_state.output.l2_block_number
              << std::endl;
    std::cout << "Got register hash "
              << intx::to_string(hash_value(mach->machine_state.registerVal),
                                 16)
              << std::endl;

    auto l = ValueLoader{
        std::make_unique<SimpleValueLoader>(storage.getDataStorage())};

    auto root = resolveTuple(l, mach->machine_state.registerVal);
    auto accountStore = indexTup(l, indexTup(l, root, 6), 1);
    auto retryKvs = indexTup(l, indexTup(l, accountStore, 1), 0);
    auto accountsKvs = indexTup(l, accountStore, 0);

    writeKvsToFile(l, retryKvs, "retryables", serializeRetryable);
    writeKvsToFile(l, accountsKvs, "accounts", serializeAccount);

    return 0;
}
