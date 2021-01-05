/*
 * Copyright 2019-2020, Offchain Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#ifndef valuetype_h
#define valuetype_h

// Proof values will only include values types up to TUPLE + 8 (11)
// All types declared with types creater than 11 are used for the internal
// marshalling format Used to pass values between the AVM and the validator
enum ValueTypes { NUM, CODEPT, HASH_PRE_IMAGE, TUPLE, BUFFER = 12, CODE_POINT_STUB = 13 };

#endif /* valuetype_h */
