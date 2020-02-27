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

	//"fmt"
	//ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
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

type VmData struct {
	machineHash       [32]byte
	pendingHash       [32]byte // Lock inbox and confirm asserts together
	inbox             [32]byte
	asserter          common.Address
	escrowRequired    big.Int
	deadline          uint64
	sequenceNum       uint64
	gracePeriod       uint32
	maxExecutionSteps uint32
	State             EthState
	inChallenge       bool
}

type channelData struct {
	state                  EthState
	gracePeriod            uint64
	maxSteps               uint32
	escrowRequired         *big.Int
	owner                  common.Address
	ValidatorCount         int64
	ActiveValidators       int64
	validatorKeys          []common.Address
	MachineHash            [32]byte
	PendingHash            [32]byte
	Inbox                  [32]byte
	Asserter               common.Address
	Deadline               uint64
	SequenceNum            uint64
	ActiveChallengeManager common.Address
	ValidatorBalances      map[common.Address]*big.Int
}

type rollupData struct {
	initVMHash              common.Hash
	VMstate                 machine.Status
	state                   EthState
	gracePeriod             common.TimeTicks
	maxSteps                uint64
	maxTimeBoundsWidth      uint64
	arbGasSpeedLimitPerTick uint64
	escrowRequired          *big.Int
	owner                   common.Address
	events                  map[*common.BlockId][]arbbridge.Event
	creation                *common.BlockId
	stakers                 map[common.Address]*staker
	leaves                  map[common.Hash]bool
	lastConfirmed           common.Hash
	contractAddress         common.Address
	nextConfirmed           common.Hash
}

type challengeData struct {
	sync.Mutex
	deadline             common.TimeTicks
	challengerDataHash   common.Hash
	state                int
	challengePeriodTicks common.TimeTicks
	asserter             common.Address
	challenger           common.Address
	challengeType        *big.Int
}

type pendingTopChallengeData struct {
	challengeData
}

type messagesChallengeData struct {
	challengeData
}

type execChallengeData struct {
	challengeData
}

type staker struct {
	location           common.Hash
	creationTimeBlocks *common.TimeBlocks
	inChallenge        bool
	balance            *big.Int
}

type nodeGraph struct {
	stakeRequirement *big.Int
	stakers          map[common.Address]*staker
	leaves           map[common.Hash]bool
	lastConfirmed    common.Hash
}

type goMaybeEvent struct {
	challenge *challengeData
	event     arbbridge.MaybeEvent
}

type inbox struct {
	value common.Hash
	count *big.Int
}

func (ib *inbox) addMessageToInbox(msg common.Hash) {
	ib.value = hashing.SoliditySHA3(
		hashing.Bytes32(ib.value),
		hashing.Bytes32(msg),
	)
	ib.count = new(big.Int).Add(ib.count, big.NewInt(1))
}

type void struct{}

var Void void

// goEthData one per 'URL'
type goEthdata struct {
	blockMutex             sync.Mutex
	msgMutex               sync.Mutex
	Vm                     map[common.Address]*VmData
	channels               map[common.Address]*channelData
	rollups                map[common.Address]*rollupData // contract address to rollup
	challenges             map[common.Address]*challengeData
	balances               map[common.Address]*big.Int
	challengeWatchersMutex sync.Mutex
	challengeWatcherEvents map[*challengeData]map[*common.BlockId][]arbbridge.Event
	arbFactory             common.Address // eth address to factory address
	nextAddress            common.Address // unique 'address'
	nextMsgs               map[*common.BlockId][]goMaybeEvent
	BlockNumber            uint64
	inbox                  map[common.Address]*inbox
	NextBlock              *common.BlockId
	LastMinedBlock         *common.BlockId
	//LatestHeight *big.Int
	blockNumbers map[uint64]*common.BlockId      // block height to blockId
	blockHashes  map[common.Hash]*common.BlockId // block hash to blockId
	parentHashes map[common.BlockId]common.Hash  // blokcId to block hash
	ethWallet    map[common.Address]*big.Int
	//headerNumber map[*big.Int]common.Hash
	//headerhash   map[common.Hash]*big.Int

	// need to hold list of out chans to publish to
	outchans map[chan arbbridge.MaybeEvent]void
	chanMgr  chan chan arbbridge.MaybeEvent
	pubchan  chan arbbridge.MaybeEvent
	//ChannelWallet map[common.Address]protocol.TokenTracker
	//create channel manager for adding, removing and publishing to list of outchans
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
		blockHash := hashing.SoliditySHA3(hashing.Uint256(big.NewInt(rand2.Int63())))
		mEthData := new(goEthdata)
		GoEth[ethURL] = mEthData
		mEthData.Vm = make(map[common.Address]*VmData)
		mEthData.channels = make(map[common.Address]*channelData)
		mEthData.rollups = make(map[common.Address]*rollupData)
		mEthData.challenges = make(map[common.Address]*challengeData)
		mEthData.balances = make(map[common.Address]*big.Int)
		mEthData.challengeWatcherEvents = make(map[*challengeData]map[*common.BlockId][]arbbridge.Event)
		mEthData.arbFactory = mEthData.getNextAddress()
		mEthData.inbox = make(map[common.Address]*inbox)
		mEthData.blockHashes = make(map[common.Hash]*common.BlockId)
		mEthData.blockNumbers = make(map[uint64]*common.BlockId)
		mEthData.parentHashes = make(map[common.BlockId]common.Hash)
		mEthData.LastMinedBlock = new(common.BlockId)
		//init header number to 0 at startup of each eth URL
		mEthData.LastMinedBlock.Height = common.NewTimeBlocks(big.NewInt(0))
		mEthData.LastMinedBlock.HeaderHash = blockHash
		mEthData.NextBlock = new(common.BlockId)
		mEthData.NextBlock.Height = common.NewTimeBlocks(big.NewInt(1))
		mEthData.nextMsgs = make(map[*common.BlockId][]goMaybeEvent)
		mEthData.blockHashes[blockHash] = mEthData.LastMinedBlock
		mEthData.blockNumbers[mEthData.LastMinedBlock.Height.AsInt().Uint64()] = mEthData.LastMinedBlock
		mEthData.parentHashes[*mEthData.LastMinedBlock] = hashing.SoliditySHA3(hashing.Uint256(big.NewInt(0)))
		mEthData.nextAddress = common.BigIntToAddress(big.NewInt(1))
		mEthData.outchans = make(map[chan arbbridge.MaybeEvent]void)
		mEthData.chanMgr = make(chan chan arbbridge.MaybeEvent)
		mEthData.pubchan = make(chan arbbridge.MaybeEvent)
		mEthData.ethWallet = make(map[common.Address]*big.Int)
		//mEthData.ChannelWallet = make(map[common.Address]protocol.TokenTracker)
		//mEthData
		//onceMutex.Unlock()
		go func() {
			for {
				select {
				case ch := <-mEthData.chanMgr: // register outchan
					mEthData.outchans[ch] = Void
				case msg := <-mEthData.pubchan: // publish to outchans
					for ch := range mEthData.outchans {
						ch <- msg
					}
				}
			}
		}()
		go func() {
			for x := range time.Tick(2 * time.Second) {
				mine(mEthData, x)
			}
		}()
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
	m.blockMutex.Lock()
	defer m.blockMutex.Unlock()

	return m.NextBlock
}

func (m *goEthdata) getLastBlock() *common.BlockId {
	m.blockMutex.Lock()
	defer m.blockMutex.Unlock()

	return m.LastMinedBlock
}

func (m *goEthdata) getBlockFromHeight(height *common.TimeBlocks) (*common.BlockId, error) {
	m.blockMutex.Lock()
	defer m.blockMutex.Unlock()

	b, ok := m.blockNumbers[height.AsInt().Uint64()]
	if !ok {
		return nil, errors.New("not found")
	}

	return b, nil
}

func (m *goEthdata) registerOutChan(oc chan arbbridge.MaybeEvent) {
	m.chanMgr <- oc
}

func (m *goEthdata) pubMsg(challenge *challengeData, msg arbbridge.MaybeEvent) {
	m.msgMutex.Lock()
	defer m.msgMutex.Unlock()
	m.nextMsgs[msg.Event.GetChainInfo().BlockId] = append(m.nextMsgs[msg.Event.GetChainInfo().BlockId], goMaybeEvent{
		challenge: challenge,
		event:     msg,
	})
}

func mine(m *goEthdata, t time.Time) {
	m.blockMutex.Lock()
	blockHash := hashing.SoliditySHA3(hashing.Uint256(big.NewInt(rand2.Int63())))
	m.NextBlock.HeaderHash = blockHash
	lastBlock := m.LastMinedBlock
	m.blockNumbers[m.NextBlock.Height.AsInt().Uint64()] = m.NextBlock
	m.blockHashes[m.NextBlock.HeaderHash] = m.NextBlock
	m.parentHashes[*m.NextBlock] = lastBlock.HeaderHash
	m.LastMinedBlock = m.NextBlock
	m.msgMutex.Lock()
	for _, event := range m.nextMsgs[m.NextBlock] {
		for _, ru := range m.rollups {
			ru.events[m.NextBlock] = append(ru.events[m.NextBlock], event.event.Event)
		}
		m.challengeWatchersMutex.Lock()
		for challenge, watcher := range m.challengeWatcherEvents {
			// if this is a challenge specific event and it is not for this challenge ignore it
			if event.challenge != nil && event.challenge != challenge {
				continue
			}
			watcher[m.NextBlock] = append(watcher[m.NextBlock], event.event.Event)
		}
		m.challengeWatchersMutex.Unlock()
		m.pubchan <- event.event
	}
	m.msgMutex.Unlock()
	log.Println("mined block number", m.NextBlock)
	newBlock := new(common.BlockId)
	newBlock.Height = common.NewTimeBlocks(new(big.Int).Add(m.LastMinedBlock.Height.AsInt(), big.NewInt(1)))
	m.NextBlock = newBlock
	blockEvent := arbbridge.NewTimeEvent{
		arbbridge.ChainInfo{
			BlockId: m.LastMinedBlock,
		},
	}

	m.blockMutex.Unlock()
	m.pubMsg(nil, arbbridge.MaybeEvent{
		Event: blockEvent,
	})
}

func addToken(source common.Address, data value.Value, destination common.Address, amount *big.Int) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		return
	}

}

func (m *goEthdata) deliverMessage(address common.Address, msgHash common.Hash) {
	hash := hashing.SoliditySHA3(
		hashing.Bytes32(m.inbox[address].value),
		hashing.Bytes32(msgHash),
	)
	m.inbox[address].value = hash
	m.inbox[address].count = new(big.Int).Add(m.inbox[address].count, big.NewInt(1))
}
