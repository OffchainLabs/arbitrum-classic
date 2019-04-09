//
//  main.cpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/19/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

/*#include <iostream>

struct value{
    
};

int main(int argc, const char * argv[]) {
    // insert code here...
    std::cout << "Hello, World!\n";
    return 0;
}*/

#include <iostream>
#include <stack>
#include <list>
#include <unordered_map>
#include <vector>
#include <iostream>
#include <fstream>
#include <sys/types.h>
#include <sys/stat.h>
#include <unistd.h>

#include "pool.hpp"
#include "value.hpp"
#include "datastack.hpp"
#include "code.hpp"
#include "machine.hpp"

//struct stk{
//    value *stkdata;
//    stk* rest;
//};





//
//void push_num(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp, uint256_t num){
//    instr *op;
//    //push(1)
//    pc++;
//    tmp->set_num(num);
//    op = new instr(pc,NOP,tmp);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
//void push_tuple(vector<instr> &code, unsigned long long &pc, int size, value *tpl, value *tmp){
//    instr *op;
//    
//    if (size==5){
//        tmp->set_num((uint256_t)11);
//        tpl->set_tuple_elem(0, tmp);
//        tmp->set_num((uint256_t)12);
//        tpl->set_tuple_elem(1, tmp);
//        tmp->set_num((uint256_t)13);
//        tpl->set_tuple_elem(2, tmp);
//        tmp->set_num((uint256_t)14);
//        tpl->set_tuple_elem(3, tmp);
//        tmp->set_num((uint256_t)15);
//        tpl->set_tuple_elem(4, tmp);
//    } else {
//        tmp->set_num((uint256_t)21);
//        tpl->set_tuple_elem(0, tmp);
//        tmp->set_num((uint256_t)22);
//        tpl->set_tuple_elem(1, tmp);
//        tmp->set_num((uint256_t)23);
//        tpl->set_tuple_elem(2, tmp);
//    }
//    //push Tuple
//    pc++;
//    op = new instr(pc,NOP,tpl);
//    code.push_back(*op);
//    delete op;
//    // print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
//void print_stack(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//    pc++;
//    op = new instr(pc,PRTSTK,NULL);
//    code.push_back(*op);
//}
//
//void rset(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//    //rset
//    tmp->set_num((uint256_t)31);
//    op = new instr(pc,RSET,tmp);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//    pc++;
//}
//
//void test_pop(vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//    pc++;
//    op = new instr(pc,POP,NULL);
//    code.push_back(*op);
//}
//
//void test_tget( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//
//    //test tget
//    push_tuple( code, pc, 5, tpl, tmp);
//    push_num( code, pc, tpl, tmp, (uint256_t)2);
//
//    // tget()
//    pc++;
//    op = new instr(pc,TGET,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
//
//void test_add( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//    //test add
//    //push(10)
//    push_num( code, pc, tpl, tmp, (uint256_t)10);
//    push_num( code, pc, tpl, tmp, (uint256_t)20);
//    op = new instr(pc,ADD,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
//
//void test_tset( vector<instr> &code, unsigned long long &pc, value *tpl, value *tpl2, value *tmp){
//    instr *op;
//    //test tset
//    //push(10)
////    push_num( code, pc, tpl, tmp, (uint256_t)10);
//    push_tuple( code, pc, 5, tpl, tmp);
//    push_tuple( code, pc, 3, tpl2, tmp);
//    push_num( code, pc, tpl, tmp, (uint256_t)1);
//    op = new instr(pc,TSET,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//    pc++;
//    op = new instr(pc,PRTSTK,NULL);
//    code.push_back(*op);
//}
//void test_pcpush( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//    //test pcpush
//    //pcpush
//    pc++;
//    op = new instr(pc,PCPUSH,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//    //rset
//    pc++;
//    op = new instr(pc,RSET,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
//void test_jump( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//    //test jump
////    rset( code, pc, tpl, tmp);
//    op = new instr(pc,RPUSH,NULL); //rpush
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL); //print
//    code.push_back(*op);
//    pc++;
//    op = new instr(pc,JUMP,NULL); //jmp
////    op = new instr(pc,NOP,NULL); //jmp
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//}
//
//void test_mul( vector<instr> &code, unsigned long long &pc, value *tpl, value *tmp){
//    instr *op;
//    //test mul
//    push_num( code, pc, tpl, tmp, (uint256_t)10);
//    push_num( code, pc, tpl, tmp, (uint256_t)20);
//
//    op = new instr(pc,MUL,NULL);
//    code.push_back(*op);
//    //print top
//    pc++;
//    op = new instr(pc,PRTTOP,NULL);
//    code.push_back(*op);
//
//}
//
//void setupCode( vector<instr> &code){
//    unsigned long long pc=0;
//    instr *op;
//    
//    value *tpl=new value((int)5);
//    value *tpl3=new value((int)3);
//    //print stack
//    pc++;
//    op = new instr(pc,PRTSTK,NULL);
//    code.push_back(*op);
//    value *tmp=new value;
//
//    push_tuple(code, pc, 3, tpl3, tmp);
//    test_tget(code, pc, tpl, tmp);
//    print_stack(code, pc, tpl, tmp);
////    test_tget( code, pc, tpl, tmp);
////    print_stack(code, pc, tpl, tmp);
////    test_add( code, pc, tpl, tmp);
//    test_pcpush( code, pc, tpl, tmp);
//    test_tset( code, pc, tpl, tpl3, tmp);
//    print_stack(code, pc, tpl, tmp);
//    test_pop( code, pc, tpl, tmp);
//
//    test_jump( code, pc, tpl, tmp);
////    test_mul( code, pc, tpl, tmp);
////    test_pcpush( code, pc, tpl, tmp);
//
//}

Machine read_file (std::string filename) {
    std::ifstream myfile;
    
    struct stat filestatus;
    stat( filename.c_str(), &filestatus );

    char *buf = (char *)malloc(filestatus.st_size);
    
    myfile.open(filename, std::ios::in);
    if (myfile.is_open())
    {
        myfile.read((char *)buf, filestatus.st_size);
        myfile.close();
    }
    return Machine(buf);
}

//void oldread_file (string filename, vector<instr> &code, char *staticValue) {
//    instr *op;
//    ifstream myfile;
//
//    struct stat filestatus;
//    stat( filename.c_str(), &filestatus );
//
//    myfile.open(filename, ios::in);
//    if (myfile.is_open())
//    {
//       // uint8_t pc_count;
//        long long pc_count;
//        char type;
//        char op_code;
//        char val_type;
//        uint256_t val;
//
//        myfile.read((char *)&pc_count, 8);
////        cout<<"pc_count="<<pc_count<<endl;
//        for (unsigned long long i=0; i<pc_count; i++){
//            value *valptr=NULL;
//            myfile.read((char*)&type, 1);
//            myfile.read((char*)&op_code, 1);
//            myfile.read((char*)&val_type, 1);
//            if (type==0x01){
//                unsigned char tmpval;
//                myfile.read((char *)&tmpval, 1);
//                val=tmpval;
//                for (int i=0; i<31; i++){
//                    myfile.read((char*)&tmpval, 1);
//                    val=val<<8;
//                    val += tmpval;
//                }
//                valptr = new value(val, NUM);
//            }
//            op = new instr(i,op_code,valptr);
//            code.push_back(*op);
////            cout<<hex<<type<<" "<<hex<<op_code<<" "<<val<<endl;
//        }
//        myfile.close();
//    }
////    return Init_machine(buf, staticValue);
//}

//int main() {
int main(int argc, char *argv[])
{
//    int state=EXTENSIVE;
    std::string filename;
    unsigned long long stepCount=1000000;
    if(argc!=2)
    {
        std::cout<<"Usage: AVMTest <ao file>"<<std::endl;
        std::cout<<"   defaulting to use add.ao"<<std::endl;
        filename = "add.ao";
    } else {
        filename = argv[1];
    }

//    oldread_file(filename, code, staticValue);
    Machine mach = read_file(filename);

//    setupCode( code );
    Assertion result = mach.run(stepCount);
    
//    runMachine(code, state, 200);
    
    exit(0);
    
}
