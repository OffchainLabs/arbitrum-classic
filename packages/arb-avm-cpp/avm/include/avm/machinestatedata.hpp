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
#include <nonstd/variant.hpp>

enum class Status { Extensive, Halted, Error };

std::vector<unsigned char> Serialize(Status status);

typedef std::array<uint256_t, 2> TimeBounds;

enum BlockType { Not, Halt, Error, Breakpoint, Inbox, Send };

struct NotBlocked {
    BlockType type = Not;
};

struct HaltBlocked {
    BlockType type = Halt;
};

struct ErrorBlocked {
    BlockType type = Error;
};

struct BreakpointBlocked {
    BlockType type = Breakpoint;
};

struct InboxBlocked {
    BlockType type = Inbox;
    uint256_t inbox;
};

struct SendBlocked {
    BlockType type = Send;
    uint256_t currency;
    TokenType tokenType;
};

using BlockReason = nonstd::variant<NotBlocked,
                                    HaltBlocked,
                                    ErrorBlocked,
                                    BreakpointBlocked,
                                    InboxBlocked,
                                    SendBlocked>;

std::vector<unsigned char> SerializeBlockReason(const BlockReason& val);

BlockReason deserializeBlockReason(std::vector<unsigned char> data);

#endif /* machinestatedata_hpp */
