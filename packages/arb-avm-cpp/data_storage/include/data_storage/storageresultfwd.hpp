//
//  storageresultfwd.hpp
//  arb-avm-cpp
//
//  Created by Harry Kalodner on 5/27/20.
//

#ifndef storageresultfwd_h
#define storageresultfwd_h

#include <vector>

struct GetResults;
struct SaveResults;
struct DeleteResults;

template <typename T>
struct DbResult;

template <typename T>
struct ValueResult;

using DataResults = ValueResult<std::vector<unsigned char>>;

#endif /* storageresultfwd_h */
