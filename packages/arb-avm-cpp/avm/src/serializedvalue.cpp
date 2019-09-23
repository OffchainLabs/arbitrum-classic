//
//  processstatus.cpp
//  avm
//
//  Created by Minh Truong on 9/22/19.
//

#include "serializedvalue.hpp"

// SerializedValue Serializer::operator()(const Tuple& val) const {
//
//    std::string return_value;
//
//    auto type_code = (char)TUPLE;
//    //auto hash_key = GetHashKey(val);
//
//    // make sure this works as intended
//    return_value += type_code;
//    //eturn_value += hash_key;
//
//    SerializedValue serialized_value{TUPLE, return_value};
//
//    return serialized_value;
//}
//
// SerializedValue Serializer::operator()(const uint256_t& val) const {
//
//    std::string return_value;
//
//    auto type_code = (char)NUM;
//    // makesure correct conversion
//    std::ostringstream ss;
//    ss << val;
//    auto value_str = ss.str();
//
//    // make sure this works as intended
//    return_value += type_code;
//    return_value += value_str;
//
//    SerializedValue serialized_value{NUM, return_value};
//
//    return serialized_value;
//}
//
// SerializedValue Serializer::operator()(const CodePoint& val) const {
//    std::string return_value;
//
//    auto type_code = (char)CODEPT;
//    // fine?
//    auto c = val.pc;
//    auto value_str = std::to_string(c);
//
//    // make sure this works as intended
//    return_value += type_code;
//    return_value += value_str;
//
//    SerializedValue serialized_value{CODEPT, return_value};
//
//    return serialized_value;
//}

// SerializedValue SerializeValue(const value& val) {
//    return nonstd::visit(Serializer{}, val);
//}
