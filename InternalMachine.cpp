//
//  InternalMachine.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/28/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//
#include <iostream>
#include <vector>
#include "uint256_t.h"
#include "datastack.hpp"
#include "code.hpp"
#include "Machine.h"

#include "InternalMachine.hpp"

using namespace std;

enum InstrType {
    BASIC_OP=0,
    IMMEDIATE_OP=1
};

/***********************************/
// test code
void push_num(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp, uint256_t num){
    instr *op;
    //push(1)
    pc++;
    tmp->set_num(num);
    op = new instr(pc,NOP,tmp);
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
}
void push_tuple(vector<instr> &code, unsigned long long &pc, int size, value *tpl, value *tmp){
    instr *op;
    
    if (size==5){
        tmp->set_num((uint256_t)11);
        tpl->set_tuple_elem(0, tmp);
        tmp->set_num((uint256_t)12);
        tpl->set_tuple_elem(1, tmp);
        tmp->set_num((uint256_t)13);
        tpl->set_tuple_elem(2, tmp);
        tmp->set_num((uint256_t)14);
        tpl->set_tuple_elem(3, tmp);
        tmp->set_num((uint256_t)15);
        tpl->set_tuple_elem(4, tmp);
    } else {
        tmp->set_num((uint256_t)21);
        tpl->set_tuple_elem(0, tmp);
        tmp->set_num((uint256_t)22);
        tpl->set_tuple_elem(1, tmp);
        tmp->set_num((uint256_t)23);
        tpl->set_tuple_elem(2, tmp);
    }
    //push Tuple
    pc++;
    op = new instr(pc,NOP,tpl);
    code.push_back(*op);
    delete op;
    // print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
}
void print_stack(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
    instr *op;
    pc++;
    op = new instr(pc,PRTSTK,NULL);
    code.push_back(*op);
}

void rset(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
    instr *op;
    //rset
    tmp->set_num((uint256_t)31);
    op = new instr(pc,RSET,tmp);
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
    pc++;
}

void test_pop(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
    instr *op;
    pc++;
    op = new instr(pc,POP,NULL);
    code.push_back(*op);
}

void test_tget( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
    instr *op;
    
    //test tget
    push_tuple( code, pc, 5, tpl, tmp);
    push_num( code, pc, tpl, tmp, (uint256_t)2);
    
    // tget()
    pc++;
    op = new instr(pc,TGET,NULL);
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
}

void test_add( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
    instr *op;
    //test add
    //push(10)
    push_num( code, pc, tpl, tmp, (uint256_t)10);
    push_num( code, pc, tpl, tmp, (uint256_t)20);
    op = new instr(pc,ADD,NULL);
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
}

void test_tset( vector<instr> &code, unsigned long long &pc, value *tpl, value *tpl2, value *tmp){
    instr *op;
    //test tset
    //push(10)
    //    push_num( code, pc, tpl, tmp, (uint256_t)10);
    push_tuple( code, pc, 5, tpl, tmp);
    push_tuple( code, pc, 3, tpl2, tmp);
    push_num( code, pc, tpl, tmp, (uint256_t)1);
    op = new instr(pc,TSET,NULL);
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
    pc++;
    op = new instr(pc,PRTSTK,NULL);
    code.push_back(*op);
}
void test_pcpush( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
    instr *op;
    //test pcpush
    //pcpush
    pc++;
    op = new instr(pc,PCPUSH,NULL);
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
    //rset
    pc++;
    op = new instr(pc,RSET,NULL);
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
}
void test_jump( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
    instr *op;
    //test jump
    //    rset( code, pc, tpl, tmp);
    op = new instr(pc,RPUSH,NULL); //rpush
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL); //print
    code.push_back(*op);
    pc++;
    op = new instr(pc,JUMP,NULL); //jmp
    //    op = new instr(pc,NOP,NULL); //jmp
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
}

void test_mul( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
    instr *op;
    //test mul
    push_num( code, pc, tpl, tmp, (uint256_t)10);
    push_num( code, pc, tpl, tmp, (uint256_t)20);
    
    op = new instr(pc,MUL,NULL);
    code.push_back(*op);
    //print top
    pc++;
    op = new instr(pc,PRTTOP,NULL);
    code.push_back(*op);
    
}

void setupCode( vector<instr> &code){
    unsigned long long pc=0;
    instr *op;
    
    value *tpl=new value((int)5);
    value *tpl3=new value((int)3);
    //print stack
    pc++;
    op = new instr(pc,PRTSTK,NULL);
    code.push_back(*op);
    value *tmp=new value;
    
    push_tuple(code, pc, 3, tpl3, tmp);
    test_tget(code, pc, tpl, tmp);
    print_stack(code, pc, tpl, tmp);
    //    test_tget( code, pc, tpl, tmp);
    //    print_stack(code, pc, tpl, tmp);
    //    test_add( code, pc, tpl, tmp);
    test_pcpush( code, pc, tpl, tmp);
    test_tset( code, pc, tpl, tpl3, tmp);
    print_stack(code, pc, tpl, tmp);
    test_pop( code, pc, tpl, tmp);
    
    test_jump( code, pc, tpl, tmp);
    //    test_mul( code, pc, tpl, tmp);
    //    test_pcpush( code, pc, tpl, tmp);
    
}
/***********************************/

uint256_t buf2uint256(char *bufptr){
    uint256_t tmpval;
    char tmpchar;

    memcpy(&tmpchar, bufptr, 1);
    bufptr+=sizeof(tmpchar);
    tmpval=tmpchar;
    for (int i=0; i<31; i++){
        memcpy(&tmpchar, bufptr, 1);
        bufptr+=sizeof(tmpchar);
        tmpval=tmpval<<8;
        tmpval += tmpchar;
    }
    return tmpval;
}

//void loadAndRunCode(void *input, unsigned long long stepcount){
Assertion run_machine(Machine *Machine, unsigned long long stepCount)
{
    Assertion ret;

    instr *op;
    vector<instr> code;
    uint256_t val;
    int state=EXTENSIVE;
    uint256_t staticVal;

//    setupCode(code);
    for(unsigned long long i=0; i<Machine->pcCount; i++){
        value *valptr=NULL;
        uint256_t immedval;
        if (Machine->code[i].InstrType==IMMEDIATE_OP){
            immedval = buf2uint256(Machine->code[i].val);
            valptr=new value(immedval, NUM);
        }
        op = new instr(i, Machine->code[i].Instr, valptr);
        code.push_back(*op);

    }
    value staticValue(buf2uint256(Machine->staticValue), NUM);
    
    ret = runMachine(code, state, stepCount, staticValue);
    
    return ret;
}

Assertion runMachine(vector<instr> &code, int state, unsigned long long maxsteps, value &staticValue){
    datastack* stk=new datastack();
    value registerVal;

    int stepcount=0;

    cout<<"starting machine code size="<<code.size()<<endl;
    for (uint64_t i=0; i<code.size(); i++){
        if (state==ERROR){
            //set error return
            cout<<"error state"<<endl;
            break;
        }
        if (state==HALTED){
            //set error return
            cout<<"halted state"<<endl;
            cout<<"full stack - size="<<stk->stacksize()<<endl;
            while (stk->stacksize()>0){
                value *A=stk->top();
                A->print();
                cout<<endl;
                stk->pop();
            }

            break;
        }
        stepcount++;
        if (stepcount >= maxsteps){
            cout<<"max steps reached"<<endl;
            break;
        }
        if (code[i].getimmediate()!=NULL){
            stk->push(*(code[i].getimmediate()));
        }
        switch (code[i].opcode) {
            case HALT:
                state=HALTED;
                break;
            case ADD:
                stk->add();
                break;
            case MUL:
                stk->mul();
                break;
            case POP:
                stk->pop();
                break;
            case RPUSH:
                stk->rpush(registerVal);
                break;
            case RSET:
                stk->rset(registerVal);
                cout<<"register set ";
                registerVal.print();
                cout<<endl;
                break;
            case JUMP:{
                pair<int,uint64_t> p=stk->jmp();
                if (p.first==1){
                    i=p.second; //found jump back
                }else if (p.first==0){
                    //not found must be forward
                    while ((i<code.size()) && (i!=p.second)){
                        i++;
                    }
                }else{
                    state=ERROR;
                }
                cout<<"jumping to "<<i<<endl;
                break;
            }
            case PCPUSH:
                cout<<"**** PCPUSH i="<<i<<endl;
                stk->pcpush(i, i);
                break;
            case NOP:
                //nop
                break;
            case TGET:
                stk->tget();
                break;
            case TSET:
                stk->tset();
                break;
            case PRTTOP:
                cout<<"stack size="<<stk->stacksize();
                if (stk->stacksize()>0){
                    value *A=stk->top();
                    cout<<" top is ";
                    A->print();
                }
                cout<<endl;
                break;
            case PRTSTK:{
                datastack* tmpstk=new datastack();
                cout<<endl;
                cout<<"full stack - size="<<stk->stacksize()<<endl;
                while (stk->stacksize()>0){
                    value *A=stk->top();
                    A->print();
                    cout<<endl;
                    tmpstk->push(*A);
                    stk->pop();
                }
                while (tmpstk->stacksize()>0){
                    value *A=tmpstk->top();
                    stk->push(*A);
                    tmpstk->pop();
                }
                cout<<"register val=";
                registerVal.print();
                cout<<endl;
                cout<<endl;
                break;
            }
                
            default:
                break;
        }
    }
    if (state==ERROR){
        //set error return
        cout<<"error state"<<endl;
    }
    if (state==HALTED){
        //set error return
        cout<<"halted state"<<endl;
    }
    cout<<"full stack - size="<<stk->stacksize()<<endl;
    while (stk->stacksize()>0){
        value *A=stk->top();
        A->print();
        cout<<endl;
        stk->pop();
    }
    cout<<"Total steps executed="<<stepcount<<endl;

    delete stk;
    return stepcount;
}
