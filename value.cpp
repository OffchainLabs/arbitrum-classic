//
//  value.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/28/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//
#include <iostream>

#include "value.hpp"

    value::value(){
        type=NONE;
        size=1;
        tplsize=0;
        refcount=1;
    }
    
    value::value(const value &obj){
        type=obj.type;
        size=obj.size;
        tplsize=obj.tplsize;
        refcount=1;
        num=obj.num;
        if (type==TUPLE){
            tpl=obj.tpl;
            tpl->ref++;
        }
    }
    
//    value::value(uint256_t val){
//        type=NUM;
//        size=1;
//        tplsize=0;
//        num=val;
//        refcount=1;
//    }

    value::value(uint256_t val, int newtype){
        type=newtype;
        size=1;
        tplsize=0;
        num=val;
        refcount=1;
    }

    value::value(int s){
        type=TUPLE;
        size=s+1;
        tplsize=s;
        tpl=tuplePool->getResource(s);
        refcount=1;
        
    }
    
    value::~value(){
        refcount--;
        //        cout<<"value obj "<<this<<" delete type="<<type<<" refcount="<<refcount<<endl;
        if (type==TUPLE && refcount==0){
            tuplePool->returnResource(tplsize, tpl);
        }
    }
    
    value &value::operator = (const value &val ){
        type = val.type;
        size = val.size;
        tplsize = val.tplsize;
        num = val.num;
        tpl = val.tpl;
        if (type == TUPLE){
            tpl->ref++;
        }

        return *this;
    }
    //    value& operator=(value other);
    //    value(const value &v2) { }
    void value::reset(){
        // TODO     release resource
        type=NONE;
        size=1;
        tplsize=0;
        
    }
    
    void value::set_num(uint256_t val) {
        type=NUM;
        size=1;
        tplsize=0;
        num=val;
    }

    void value::set_codept(uint256_t val) {
        type=CODEPT;
        size=1;
        tplsize=0;
        num=val;
    }
    //    void set_reg(uint256_t val) {
    //        reg=val;
    //    }
    
    /*    void set_tuple(int count, value *vals) {
     type=TUPLE;
     size=count;
     tplsize=count;
     value *newtuple=new value[count];
     tpl = newtuple;
     cout<<tpl<<endl;
     //        memcpy(tpl, vals, sizeof(value)*count);
     for (int i=0; i<count; i++){
     newtuple[i] = vals[i];
     size+=newtuple[i].size;
     }
     }
     */
    //   copy_tuple(value *){
    
    //   }
    value *value::dup(){
        value *tmp=new value;
        tmp->type = type;
        if (type==NUM){
            tmp->num=num;
        }
        if (type==CODEPT){
            tmp->num=num;
        }
        if (type==TUPLE){
            tmp->tpl=tpl;
            tmp->tplsize=tplsize;
        }
        
        return tmp;
    }
    int value::set_tuple_elem(int pos, value *newval) {
        if (type != TUPLE) return -1;
        if (pos >= tplsize) return -1;
//        value *oldval=&tpl->vals[pos];
        if (tpl->ref>1){
            //make new copy tuple
            vTuple* tmp=tuplePool->getResource(tplsize);
            memcpy(tmp->vals, tpl->vals, sizeof(value)*tplsize);
            tpl->ref--;
            tpl=tmp;
            for(int i=0; i<tplsize; i++){
                if (tpl->vals[i].type==TUPLE){
                    tpl->ref++;
                }
            }
//            value *oldval=&tpl->vals[pos];
        }
        if (tpl->vals[pos].type == TUPLE){
            //slot is currently tuple return old
            tuplePool->returnResource(tpl->vals[pos].tplsize, tpl->vals[pos].tpl);
        }
        tpl->vals[pos].type=newval->type;
        tpl->vals[pos].size=newval->size;
        tpl->vals[pos].tplsize=newval->tplsize;
        tpl->vals[pos].num=newval->num;
        
        tpl->vals[pos].refcount=newval->refcount;
        if (newval->type == TUPLE){
            tpl->vals[pos].tpl=newval->tpl;
            tpl->vals[pos].tpl->ref++;
            //copy Tuple
        } else {
            tpl->vals[pos].num=newval->num;
        }
        return 0;
    }
    
    value *value::get_tuple_elem(int pos) {
        if (type != TUPLE) return NULL;
        if (pos >= tplsize) return NULL;
        
        return &tpl->vals[pos];
    }
    void value::print(){
        if (type==NUM){
            cout<<"num="<<num;
        }
        if (type==CODEPT){cout<<"codept="<<num;}
        if (type==TUPLE){
            value *val=tpl->vals;
            cout<<"tuple="<<tpl<<" [";
            for (int i=0; i<tplsize; i++){
                val[i].print();
                cout<<((i<tplsize-1)?",":"");
            }
            cout<<"]";
        }
    }

void value::printstack(){
    if (type==NUM){cout<<"num="<<num;}
    if (type==CODEPT){cout<<"codept="<<num;}
    if (type==TUPLE){
        value *val=tpl->vals;
        cout<<"tuple="<<tpl<<" [";
        for (int i=0; i<tplsize; i++){
            val[i].print();
            cout<<((i<tplsize-1)?",":"");
        }
        cout<<"]";
    }
}

