//
//  helper.hpp
//  arb-avm-cpp
//
//  Created by Harry Kalodner on 5/17/20.
//

#ifndef avm_tests_helper_hpp
#define avm_tests_helper_hpp

#include <boost/filesystem.hpp>

#include <string>

extern std::string dbpath;

struct DBDeleter {
    DBDeleter() { boost::filesystem::remove_all(dbpath); }
    ~DBDeleter() { boost::filesystem::remove_all(dbpath); }
};

#endif /* avm_tests_helper_hpp */
