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

package rollup

import (
	"context"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
)

// WARNING: The code in this file is badly behaved, on purpose. It is for testing only.
//     If you call this in production, you will be sorry.

type WrongAssertionType int

const (
	WrongPendingTopAssertion    = 0
	WrongMessagesSliceAssertion = 1
	WrongExecutionAssertion     = 2
)

type evilWrongAssertionListener struct {
	*ValidatorChainListener
	kind WrongAssertionType
}

func NewEvilWrongAssertionListener(
	rollupAddress common.Address,
	actor arbbridge.ArbRollup,
	kind WrongAssertionType,
) *evilWrongAssertionListener {
	return &evilWrongAssertionListener{NewValidatorChainListener(context.Background(), rollupAddress, actor), kind}
}

func (lis *evilWrongAssertionListener) AssertionPrepared(ctx context.Context, obs *ChainObserver, assertion *preparedAssertion) {
	badHash := common.Hash{}
	badHash[5] = 37
	switch lis.kind {
	case WrongPendingTopAssertion:
		assertion.claim.AfterPendingTop = badHash
		log.Println("Prepared EVIL pending top assertion")
	case WrongMessagesSliceAssertion:
		assertion.claim.ImportedMessagesSlice = badHash
		log.Println("Prepared EVIL imported messages assertion")
	case WrongExecutionAssertion:
		assertion.claim.AssertionStub.AfterHash = badHash
		log.Println("Prepared EVIL execution assertion")
	default:
		log.Fatal("unrecognized evil listener type")
	}
	lis.ValidatorChainListener.AssertionPrepared(ctx, obs, assertion)
}
