//
//  datastack.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include "datastack.hpp"

#include "opcodes.hpp"

void datastack::push(value && newdata){
    basedatastack.push_back(std::move(newdata));
    
};

value datastack::pop() {
    if (basedatastack.size()==0){
        throw std::runtime_error("Stack is empty");
    }
    
    auto val = std::move(basedatastack.back());
    basedatastack.pop_back();
    return val;
};

value &datastack::peek() {
    if (basedatastack.size()==0){
        throw std::runtime_error("Stack is empty");
    }
    
    return basedatastack.back();
};

value &peek();

void datastack::popSet(value &val) {
    if (basedatastack.size()==0){
        throw std::runtime_error("Stack is empty");
    }
    
    val = std::move(basedatastack.back());
    basedatastack.pop_back();
};

uint64_t datastack::stacksize(){
    return basedatastack.size();
};

//void datastack::pcpush(uint64_t i, uint64_t j){
////    pcmap.insert(std::make_pair(i, j-1));
//    push(CodePoint{i});
//}

//std::pair<int,uint64_t> datastack::jmp(){
//    auto A = pop();
//    auto cp = mpark::get_if<CodePoint>(&A);
//    if (!cp) {
//        return std::make_pair(-1, NULL);
//    }
//    
//    uint64_t val = cp->pc;
//    if (pcmap.find(val)==pcmap.end()){
//        return std::make_pair(0, val);
//    }else{
//        return std::make_pair(1, pcmap[val]);
//    }
//}

//int datastack::tset(){
//    auto A = pop(); // slot
//    auto B = pop(); // tuple
//    auto C = pop(); // val
//    auto aIndex = mpark::get_if<uint256_t>(&A);
//    auto bTup = mpark::get_if<Tuple>(&B);
//    if (!aIndex || bTup ||
//        *aIndex >= bTup->tuple_size()){
//        return ERROR;
//    }
//
//    bTup->set_element(static_cast<uint32_t>(*aIndex), std::move(C));
//    push(std::move(*bTup));
//    return EXTENSIVE;
//}

