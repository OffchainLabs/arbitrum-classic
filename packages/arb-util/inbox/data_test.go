/*
* Copyright 2020, Offchain Labs, Inc.
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

package inbox

import (
	"testing"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

func TestListToStackValue(t *testing.T) {
	vals := make([]value.Value, 0)
	for i := int64(0); i < 10; i++ {
		vals = append(vals, value.NewInt64Value(i))
	}
	stackVal := ListToStackValue(vals)

	vals2, err := StackValueToList(stackVal)
	if err != nil {
		t.Fatal(err)
	}

	if len(vals) != len(vals2) {
		t.Fatal("wrong val count")
	}

	for i, val := range vals {
		if !value.Eq(val, vals2[i]) {
			t.Fatal("val not equal")
		}
	}
}

func TestStackValueToListFailures(t *testing.T) {
	intVal := value.NewInt64Value(0)
	if _, err := StackValueToList(intVal); err == nil {
		t.Error("should fail when passed non-tuple")
	}

	// Static slice correct size, so error can be ignored
	tup, _ := value.NewTupleFromSlice([]value.Value{intVal, intVal, intVal})
	if _, err := StackValueToList(tup); err == nil {
		t.Error("should fail when passed tuple not of size 2")
	}
}
