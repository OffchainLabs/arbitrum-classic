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

package cmachine

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lcavm -lavm -ldata_storage -lavm_values -lstdc++ -lm -lrocksdb -ldl
#include "../cavm/cvaluecache.h"
#include <stdio.h>
#include <stdlib.h>
*/
import "C"
import (
	"github.com/pkg/errors"
	"runtime"
	"unsafe"
)

type ValueCache struct {
	c unsafe.Pointer
}

func NewValueCache() (*ValueCache, error) {
	cValueCache := C.createValueCache()

	if cValueCache == nil {
		return nil, errors.Errorf("error creating value cache")
	}
	ret := &ValueCache{cValueCache}
	runtime.SetFinalizer(ret, destroyValueCache)
	return ret, nil
}

func destroyValueCache(cValueCache *ValueCache) {
	C.destroyValueCache(cValueCache.c)
}

func (valueCache *ValueCache) Clear() {
	C.clearValueCache(valueCache.c)
}
