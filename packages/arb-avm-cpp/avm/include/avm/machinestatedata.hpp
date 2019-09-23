//
//  machinestatedata.hpp
//  avm
//
//  Created by Minh Truong on 9/23/19.
//

#ifndef machinestatedata_hpp
#define machinestatedata_hpp

#include <stdio.h>
#include <avm/tokenTracker.hpp>
#include <avm/value.hpp>

enum class Status { Extensive, Halted, Error };

typedef std::array<uint256_t, 2> TimeBounds;

struct NotBlocked {};

struct HaltBlocked {};

struct ErrorBlocked {};

struct BreakpointBlocked {};

struct InboxBlocked {
    uint256_t inbox;
};

struct SendBlocked {
    uint256_t currency;
    TokenType tokenType;
};

using BlockReason = nonstd::variant<NotBlocked,
                                    HaltBlocked,
                                    ErrorBlocked,
                                    BreakpointBlocked,
                                    InboxBlocked,
                                    SendBlocked>;

#endif /* machinestatedata_hpp */
