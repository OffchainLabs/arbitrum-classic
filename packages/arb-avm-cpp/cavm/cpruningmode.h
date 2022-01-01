/*
 * Copyright 2021, Offchain Labs, Inc.
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

#ifndef ARB_AVM_CPP_CPRUNINGMODE_H
#define ARB_AVM_CPP_CPRUNINGMODE_H

typedef enum {
    PRUNING_MODE_OFF = 0,
    PRUNING_MODE_ON = 1,
    PRUNING_MODE_DEFAULT = 2
} CPruningMode;

#endif  // ARB_AVM_CPP_CPRUNINGMODE_H
