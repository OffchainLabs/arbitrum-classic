//
//  utilities.hpp
//  avm
//
//  Created by Minh Truong on 9/20/19.
//

#ifndef exceptions_hpp
#define exceptions_hpp

#include <nonstd/variant.hpp>

class bad_tuple_index : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "bad_tuple_index";
    }
};

class bad_pop_type : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "bad_variant_access";
    }
};

class int_out_of_bounds : public std::exception {
   public:
    virtual const char* what() const noexcept override {
        return "int_out_of_bounds";
    }
};

#endif /* exceptions_hpp */
