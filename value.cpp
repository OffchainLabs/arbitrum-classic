//
//  value.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/28/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include "value.hpp"
#include "pool.hpp"
#include "util.hpp"

#include <iostream>

uint256_t deserialize_int(char *&bufptr) {
    uint256_t ret = from_big_endian(bufptr, bufptr + 32);
    bufptr += 32;
    return ret;
}

Tuple deserialize_tuple(char *&bufptr, int size, TuplePool &pool) {
    Tuple tup(size, &pool);
    for (int i = 0; i < size; i++) {
        tup.set_element(i, deserialize_value(bufptr, pool));
    }
    return tup;
}

value deserialize_value(char *&srccode, TuplePool &pool) {
    uint8_t valType;
    memcpy(&valType, srccode, sizeof(valType));
    srccode += sizeof(valType);
    switch (valType) {
        case NUM:
            return deserialize_int(srccode);
        case CODEPT:
            throw std::runtime_error("Tried to deserialize unhandled codepoint");
        default:
            if (valType >= TUPLE && valType <= TUPLE + 8) {
                return deserialize_tuple(srccode, valType - TUPLE, pool);
            } else {
                throw std::runtime_error("Tried to deserialize unhandled type");
            }
    }
}

Tuple::Tuple(int size_, TuplePool *pool) :
tuplePool(pool),
size(size_ + 1),
tpl(pool->getResource(size_)) {}

Tuple::Tuple(const Tuple &tup) :
tuplePool(tup.tuplePool),
size(tup.size),
tpl(tup.tpl) {}

Tuple::~Tuple(){
    tuplePool->returnResource(std::move(tpl));
}

int Tuple::tuple_size() const {
    return tpl->size();
}

void Tuple::set_element(int pos, value && newval) {
    if (pos >= tuple_size()){
        throw bad_tuple_index{};
    }

    if (tpl.use_count() > 1) {
        //make new copy tuple
        std::shared_ptr<std::vector<value>> tmp = tuplePool->getResource(tpl->size());
        std::copy(tpl->begin(), tpl->end(), tmp->begin());
        tpl=tmp;
    }
    (*tpl)[pos] = std::move(newval);
}

value Tuple::get_element(int pos) const {
    if (pos >= tuple_size()){
        throw bad_tuple_index{};
    }
    return (*tpl)[pos];
}

std::vector<unsigned char> value_hash_raw(const value &value) {
    return mpark::visit([](const auto &val) {
        return value_hash_raw(val);
    }, value);
}

std::vector<unsigned char> value_hash_raw(const uint256_t &value) {
    std::vector<unsigned char> intData;
    intData.resize(32);
    to_big_endian(value, intData.begin());
    
    std::vector<unsigned char> hashData;
    hashData.resize(32);
    evm::Keccak_256(intData.data(), 32, hashData.data());
    return hashData;
}

std::vector<unsigned char> value_hash_raw(const Tuple &tup) {
    
    std::vector<unsigned char> tupData;
    tupData.resize(1 + tup.tuple_size() * 32);
    auto oit = tupData.begin();
    tupData[0] = TUPLE + tup.tuple_size();
    ++oit;
    for (int i = 0; i < tup.tuple_size(); i++) {
        auto valHash = value_hash_raw(tup.get_element(i));
        std::copy(valHash.begin(), valHash.end(), oit);
        oit += 32;
    }
    
    std::vector<unsigned char> hashData;
    hashData.resize(32);
    evm::Keccak_256(tupData.data(), 32, hashData.data());
    return hashData;
}

std::vector<unsigned char> value_hash_raw(const CodePoint &cp) {
    throw std::runtime_error("CodePoint hash not supported");
}

uint256_t value_hash(const value &value) {
    return mpark::visit([](const auto &val) {
        auto raw = value_hash_raw(val);
        return from_big_endian(raw.begin(), raw.end());
    }, value);
}

struct ValuePrinter {
    std::ostream& os;
    
    std::ostream &operator()(const Tuple &val) {
        os << "tuple=" << " [";
        for (int i = 0; i < val.tuple_size(); i++){
            std::cout<< val.get_element(i) << ((i < val.tuple_size()-1) ? "," : "");
        }
        std::cout<<"]";
        return os;
    }
    
    std::ostream &operator()(const uint256_t &val) {
        os << "num=" << val;
        return os;
    }
    
    std::ostream &operator()(const CodePoint &val) {
        os << "codept=" << val;
        return os;
    }
};

std::ostream& operator<<(std::ostream& os, const value& val) {
    return mpark::visit(ValuePrinter{os}, val);
}
