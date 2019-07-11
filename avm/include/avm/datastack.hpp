//
//  datastack.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef datastack_hpp
#define datastack_hpp

#include "tuple.hpp"
#include "value.hpp"

#include <iostream>
#include <vector>

class datastack {
    static constexpr int lazyCount = 100;

    void addHash() const;

   public:
    std::vector<value> values;
    mutable std::vector<uint256_t> hashes;

    datastack() {
        values.reserve(1000);
        hashes.reserve(1000);
    }

    void push(value&& newdata) {
        values.push_back(std::move(newdata));
        if (values.size() > hashes.size() + lazyCount) {
            addHash();
        }
    };
    
    const value &operator[](size_t index) const {
        return values[values.size() - 1 - index];
    }
    
    value &operator[](size_t index) {
        return values[values.size() - 1 - index];
    }

    value pop() {
        auto stackEmpty = values.empty();
        if (stackEmpty) {
            throw std::runtime_error("Stack is empty");
        }

        auto val = std::move(values.back());
        popClear();
        return val;
    }
    
    void prepForMod(int count) {
        while(hashes.size() > values.size() - count) {
            hashes.pop_back();
        }
    }
    
    void popClear() {
        values.pop_back();
        if (hashes.size() > values.size()) {
            hashes.pop_back();
        }
    }
    
    uint256_t SolidityProofValue(std::vector<bool>& stackInfo, std::vector<value>& vals);

    value& peek() {
        if (values.size() == 0) {
            throw std::runtime_error("Stack is empty");
        }

        return values.back();
    }

    uint64_t stacksize() { return values.size(); }

    uint256_t hash() const;
};

std::ostream& operator<<(std::ostream& os, const datastack& val);

#endif /* datastack_hpp */
