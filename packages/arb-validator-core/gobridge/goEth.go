/*
 * Copyright 2019, Offchain Labs, Inc.
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

package gobridge

import (
	"errors"
	"log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/arbbridge"
	"math/big"
	rand2 "math/rand"
	"sync"
	"time"
)

type EthState int

const (
	Uninitialized EthState = iota
	Waiting
	PendingDisputable
	PendingUnanimous
)

// for future channel development - should be moved to a channel module
//type channelData struct {
//	state                  EthState
//	gracePeriod            uint64
//	maxSteps               uint32
//	escrowRequired         *big.Int
//	owner                  common.Address
//	ValidatorCount         int64
//	ActiveValidators       int64
//	validatorKeys          []common.Address
//	MachineHash            [32]byte
//	PendingHash            [32]byte
//	Inbox                  [32]byte
//	Asserter               common.Address
//	Deadline               uint64
//	SequenceNum            uint64
//	ActiveChallengeManager common.Address
//	ValidatorBalances      map[common.Address]*big.Int
//}

type staker struct {
	location           common.Hash
	creationTimeBlocks *common.TimeBlocks
	inChallenge        bool
	balance            *big.Int
}

type contractMsgs struct {
	msgs map[common.Address][]arbbridge.Event
}

// goEthData one per 'URL'
type goEthdata struct {
	ethAddress               common.Address
	goEthMutex               sync.Mutex
	rollups                  map[common.Address]*arbRollup // contract instance to rollupData
	challenges               map[common.Address]*challenge
	nextAddress              common.Address // unique 'contractAddress'
	blockMsgs                map[*common.BlockId]contractMsgs
	arbFactoryContract       *arbFactory
	globalInbox              *globalInbox
	challengeFactoryContract *challengeFactory
	rootBlock                *common.BlockId
	NextBlock                *common.BlockId
	LastMinedBlock           *common.BlockId
	blockNumbers             map[uint64]*common.BlockId      // block height to blockId
	blockHashes              map[common.Hash]*common.BlockId // block hash to blockId
	parentHashes             map[common.BlockId]common.Hash  // blokcId to block hash
	ethWallet                map[common.Address]*big.Int
	ftWallets                map[common.Address]*userFTWallet
	nftWallets               map[common.Address]*userNFTWallet
}

var GoEth map[string]*goEthdata

var once map[string]*sync.Once
var onceMutex sync.Mutex

func init() {
	GoEth = make(map[string]*goEthdata)
}

func StartGoEth(ethURL string) {
	getGoEth(ethURL)
}

func getGoEth(ethURL string) *goEthdata {
	// once for each ethURL, set up data
	onceMutex.Lock()
	defer onceMutex.Unlock()
	tmpOnce, ok := once[ethURL]
	if !ok {
		once = make(map[string]*sync.Once)
		once[ethURL] = new(sync.Once)
		tmpOnce = once[ethURL]
	}
	tmpOnce.Do(func() {
		mEthData := new(goEthdata)
		GoEth[ethURL] = mEthData
		mEthData.ethAddress = common.BigIntToAddress(big.NewInt(1))
		mEthData.nextAddress = common.BigIntToAddress(big.NewInt(2))
		//mEthData.challenges = make(map[common.Address]*challengeData)
		//mEthData.balances = make(map[common.Address]*big.Int)
		//mEthData.challengeWatcherEvents = make(map[*challenge]map[*common.BlockId][]arbbridge.Event)
		mEthData.blockHashes = make(map[common.Hash]*common.BlockId)
		mEthData.blockNumbers = make(map[uint64]*common.BlockId)
		mEthData.parentHashes = make(map[common.BlockId]common.Hash)
		//init header number to 0 at startup of each eth URL
		rootBlockHash := hashing.SoliditySHA3(hashing.Uint256(big.NewInt(rand2.Int63())))
		mEthData.rootBlock = new(common.BlockId)
		mEthData.rootBlock.Height = common.NewTimeBlocks(big.NewInt(0))
		mEthData.rootBlock.HeaderHash = rootBlockHash
		mEthData.blockHashes[rootBlockHash] = mEthData.rootBlock
		mEthData.blockMsgs = make(map[*common.BlockId]contractMsgs)
		mEthData.blockMsgs[mEthData.rootBlock] = contractMsgs{make(map[common.Address][]arbbridge.Event)}
		mEthData.LastMinedBlock = mEthData.rootBlock
		mEthData.NextBlock = new(common.BlockId)
		mEthData.NextBlock.Height = common.NewTimeBlocks(big.NewInt(1))
		mEthData.blockMsgs[mEthData.NextBlock] = contractMsgs{make(map[common.Address][]arbbridge.Event)}
		mEthData.blockNumbers[mEthData.LastMinedBlock.Height.AsInt().Uint64()] = mEthData.LastMinedBlock
		mEthData.parentHashes[*mEthData.LastMinedBlock] = hashing.SoliditySHA3(hashing.Uint256(big.NewInt(0)))
		mEthData.ethWallet = make(map[common.Address]*big.Int)
		mEthData.ftWallets = make(map[common.Address]*userFTWallet)
		mEthData.nftWallets = make(map[common.Address]*userNFTWallet)
		go func() {
			for x := range time.Tick(2 * time.Second) {
				mine(mEthData, x)
			}
		}()
		deployGlobalInbox(mEthData)
		deployRollupFactory(mEthData)
		deployChallengeFactory(mEthData)
	})
	return GoEth[ethURL]
}

func (m *goEthdata) getNextAddress() common.Address {
	addr := m.nextAddress
	addrInt := new(big.Int)
	addrInt.SetBytes(addr[:])
	m.nextAddress = common.BigIntToAddress(addrInt.Add(addrInt, big.NewInt(1)))
	return addr
}

func (m *goEthdata) getCurrentBlock() *common.BlockId {
	return m.NextBlock
}

func (m *goEthdata) getLastBlock() *common.BlockId {
	return m.LastMinedBlock
}

func (m *goEthdata) getBlockFromHeight(height *common.TimeBlocks) (*common.BlockId, error) {
	b, ok := m.blockNumbers[height.AsInt().Uint64()]
	if !ok {
		return nil, errors.New("not found")
	}

	return b, nil
}

func (m *goEthdata) pubMsg(addr common.Address, msg arbbridge.Event) {
	m.blockMsgs[msg.GetChainInfo().BlockId].msgs[addr] = append(m.blockMsgs[msg.GetChainInfo().BlockId].msgs[addr], msg)
}

func mine(m *goEthdata, t time.Time) {
	m.goEthMutex.Lock()
	defer m.goEthMutex.Unlock()
	blockHash := hashing.SoliditySHA3(hashing.Uint256(big.NewInt(rand2.Int63())))
	m.NextBlock.HeaderHash = blockHash
	lastBlock := m.LastMinedBlock
	m.blockNumbers[m.NextBlock.Height.AsInt().Uint64()] = m.NextBlock
	m.blockHashes[m.NextBlock.HeaderHash] = m.NextBlock
	m.parentHashes[*m.NextBlock] = lastBlock.HeaderHash
	m.LastMinedBlock = m.NextBlock
	log.Println("mined block number", m.NextBlock)
	newBlock := new(common.BlockId)
	newBlock.Height = common.NewTimeBlocks(new(big.Int).Add(m.LastMinedBlock.Height.AsInt(), big.NewInt(1)))
	m.NextBlock = newBlock
	m.blockMsgs[m.NextBlock] = contractMsgs{make(map[common.Address][]arbbridge.Event)}
	blockEvent := arbbridge.NewTimeEvent{
		ChainInfo: arbbridge.ChainInfo{
			BlockId: m.LastMinedBlock,
		},
	}

	m.pubMsg(m.ethAddress, blockEvent)
}
