//
//  datastack.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/30/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#include "datastack.hpp"

void datastack::push(value &newdata){
    value* tmp=new value(newdata);
    basedatastack.push(tmp);
    
};
void datastack::push(uint256_t val, int type){
    value* tmp=new value(val,type);
    basedatastack.push(tmp);
    
};
void datastack::pop(){
    value* A=basedatastack.top();
    basedatastack.pop();
    delete A;
    
};
void datastack::popNoDel(){
    basedatastack.pop();
};
value *datastack::top(){
    if (basedatastack.size()==0){
        //error
        return NULL;
    }
    return basedatastack.top();
};

uint64_t datastack::stacksize(){
    return basedatastack.size();
};

void datastack::rpush(value &val){
    push(val);
};

void datastack::rset(value &val){
    A = top();
    val=*A;
    pop();
};

void datastack::pcpush(uint64_t i, uint64_t j){
    pcmap.insert(make_pair(i, j-1));
    push((uint256_t)i,CODEPT);
}

pair<int,uint64_t> datastack::jmp(){
    A = top();
    if (A->type != CODEPT){
        pop();
        return make_pair(-1, NULL);
    }
    uint64_t val=A->num.lower().lower();
    pop();
    
    if (pcmap.find(val)==pcmap.end()){
        return make_pair(0, val);
    }else{
        return make_pair(1, pcmap[val]);
    }
}

int datastack::tget(){
    if (stacksize()<2){
        //error
        return ERROR;
    }
    A=top();
    popNoDel();
    B=top();
    popNoDel();
    if ((A->type != NUM) || (B->type != TUPLE) || (A->num >= B->tplsize)){
        delete A;
        delete B;
        return ERROR;
    }
    push(*(B->get_tuple_elem((uint)A->num)));
    delete A;
    delete B;
    return EXTENSIVE;
}

int datastack::tset(){
    A=top(); // slot
    popNoDel();
    B=top(); // tuple
    popNoDel();
    C=top(); // val
    popNoDel();
    if ((A->type != NUM) ||
        (B->type != TUPLE) ||
        (A->num >= B->tplsize)){
        delete A;
        delete B;
        delete C;
        return ERROR;
    }
    
    B->set_tuple_elem((uint)A->num, C);
    push(*B);
    delete A;
    delete B;
    delete C;
    return EXTENSIVE;
}

int datastack::add(){
    A=top();
    popNoDel();
    B=top();
    popNoDel();
    if ((A->type != NUM) || (B->type != NUM)){
        delete A;
        delete B;
        return ERROR;
    }
    
    uint256_t sum=A->num+B->num;
    A->set_num(sum);
    push(*A);
    delete A;
    delete B;
    return EXTENSIVE;
}
int datastack::mul(){
    A=top();
    popNoDel();
    B=top();
    popNoDel();
    if ((A->type != NUM) || (B->type != NUM)){
        delete A;
        delete B;
        return -1;
    }
    uint256_t sum=A->num*B->num;
    A->set_num(sum);
    push(*A);
    delete A;
    delete B;
    return EXTENSIVE;
}
