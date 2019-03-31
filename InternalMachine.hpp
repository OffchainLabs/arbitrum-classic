//
//  InternalMachine.hpp
//  AVMtest
//
//  Created by Timothy O'Bryan on 3/28/19.
//  Copyright © 2019 Timothy O'Bryan. All rights reserved.
//

#ifndef InternalMachine_hpp
#define InternalMachine_hpp

#include "Machine.h"

Assertion runMachine(vector<instr> &code, int state, uint64_t maxsteps, value &staticValue);

#endif /* InternalMachine_hpp */
