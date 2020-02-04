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
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/arbbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/structures"
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
	pendingHash       [32]byte // Lock pending and confirm asserts together
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
	state           EthState
	vmState         common.Hash
	gracePeriod     common.TimeTicks
	maxSteps        uint64
	escrowRequired  *big.Int
	owner           common.Address
	events          map[*structures.BlockId][]arbbridge.Event
	creation        *structures.BlockId
	stakers         map[common.Address]*staker
	leaves          map[common.Hash]bool
	lastConfirmed   common.Hash
	contractAddress common.Address
	nextConfirmed   common.Hash
}

type challengeData struct {
	sync.Mutex
	deadline common.TimeTicks
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

type void struct{}

var Void void

// goEthData one per 'URL'
type goEthdata struct {
	blockMutex        sync.Mutex
	Vm                map[common.Address]*VmData
	channels          map[common.Address]*channelData
	rollups           map[common.Address]*rollupData    // contract address to rollup
	challenges        map[common.Address]*challengeData // contract address to rollup
	challengeWatchers map[*challengeData]map[*structures.BlockId][]arbbridge.Event
	arbFactory        common.Address // eth address to factory address
	nextAddress       common.Address // unique 'address'
	nextMsgs          map[*structures.BlockId][]arbbridge.MaybeEvent
	BlockNumber       uint64
	pending           map[common.Address]*PendingInbox
	NextBlock         *structures.BlockId
	LastMinedBlock    *structures.BlockId
	//LatestHeight *big.Int
	blockNumbers map[uint64]*structures.BlockId      // block height to blockId
	blockHashes  map[common.Hash]*structures.BlockId // block hash to blockId
	parentHashes map[structures.BlockId]common.Hash  // blokcId to block hash
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

func getGoEth(ethURL string) *goEthdata {
	// once for each ethURL, set up data
	onceMutex.Lock()
	tmpOnce, ok := once[ethURL]
	if !ok {
		once = make(map[string]*sync.Once)
		once[ethURL] = new(sync.Once)
		tmpOnce = once[ethURL]
	}
	tmpOnce.Do(func() {
		blockHash := common.NewHashFromEth(ethcommon.BigToHash(big.NewInt(rand2.Int63())))
		mEthData := new(goEthdata)
		GoEth[ethURL] = mEthData
		mEthData.Vm = make(map[common.Address]*VmData)
		mEthData.channels = make(map[common.Address]*channelData)
		mEthData.rollups = make(map[common.Address]*rollupData)
		mEthData.challenges = make(map[common.Address]*challengeData)
		mEthData.challengeWatchers = make(map[*challengeData]map[*structures.BlockId][]arbbridge.Event)
		mEthData.arbFactory = mEthData.getNextAddress()
		mEthData.pending = make(map[common.Address]*PendingInbox)
		mEthData.blockHashes = make(map[common.Hash]*structures.BlockId)
		mEthData.blockNumbers = make(map[uint64]*structures.BlockId)
		mEthData.parentHashes = make(map[structures.BlockId]common.Hash)
		//init header number to 0 at startup
		mEthData.LastMinedBlock = new(structures.BlockId)
		mEthData.LastMinedBlock.Height = common.NewTimeBlocks(big.NewInt(0))
		mEthData.LastMinedBlock.HeaderHash = blockHash
		mEthData.NextBlock = new(structures.BlockId)
		mEthData.NextBlock.Height = common.NewTimeBlocks(big.NewInt(1))
		mEthData.nextMsgs = make(map[*structures.BlockId][]arbbridge.MaybeEvent)
		mEthData.blockHashes[blockHash] = mEthData.LastMinedBlock
		mEthData.blockNumbers[mEthData.LastMinedBlock.Height.AsInt().Uint64()] = mEthData.LastMinedBlock
		mEthData.parentHashes[*mEthData.LastMinedBlock] = common.NewHashFromEth(ethcommon.BigToHash(big.NewInt(0)))
		mEthData.nextAddress = common.NewAddressFromEth(ethcommon.BigToAddress(big.NewInt(1)))
		mEthData.outchans = make(map[chan arbbridge.MaybeEvent]void)
		mEthData.chanMgr = make(chan chan arbbridge.MaybeEvent)
		mEthData.pubchan = make(chan arbbridge.MaybeEvent)
		//mEthData.ChannelWallet = make(map[common.Address]protocol.TokenTracker)
		//mEthData
		onceMutex.Unlock()
		go func() {
			for {
				select {
				case ch := <-mEthData.chanMgr: // register outchan
					mEthData.outchans[ch] = Void
				case msg := <-mEthData.pubchan: // publish to outchans
					for ch, _ := range mEthData.outchans {
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
	m.nextAddress = common.NewAddressFromEth(ethcommon.BigToAddress(addrInt.Add(addrInt, big.NewInt(1))))
	return addr
}

func (m *goEthdata) getCurrentBlock() *structures.BlockId {
	m.blockMutex.Lock()
	defer m.blockMutex.Unlock()

	return m.NextBlock
}

func (m *goEthdata) getLastBlock() *structures.BlockId {
	m.blockMutex.Lock()
	defer m.blockMutex.Unlock()

	return m.LastMinedBlock
}

func (m *goEthdata) getBlockFromHeight(height *common.TimeBlocks) (*structures.BlockId, error) {
	m.blockMutex.Lock()
	defer m.blockMutex.Unlock()

	b, ok := m.blockNumbers[height.AsInt().Uint64()]
	if !ok {
		return nil, errors.New("not found")
	}

	return b, nil
}

func (m *goEthdata) registerOutChan(oc chan arbbridge.MaybeEvent) {
	fmt.Println("registering outchan")
	m.chanMgr <- oc
}

func (m *goEthdata) pubMsg(msg arbbridge.MaybeEvent) {
	//fmt.Println("publishing event", msg)
	//if msg.Event.GetChainInfo().BlockId == m.NextBlock {
	m.nextMsgs[msg.Event.GetChainInfo().BlockId] = append(m.nextMsgs[msg.Event.GetChainInfo().BlockId], msg)
	//} else {
	//	for _, ru := range m.rollups {
	//		ru.events[msg.Event.GetChainInfo().BlockId] = append(ru.events[msg.Event.GetChainInfo().BlockId], msg.Event)
	//	}
	//	for _, watcher := range m.challengeWatchers {
	//		watcher[msg.Event.GetChainInfo().BlockId] = append(watcher[msg.Event.GetChainInfo().BlockId], msg.Event)
	//	}
	//	m.pubchan <- msg
	//}
}

func mine(m *goEthdata, t time.Time) {
	m.blockMutex.Lock()
	blockHash := common.NewHashFromEth(ethcommon.BigToHash(big.NewInt(rand2.Int63())))
	m.NextBlock.HeaderHash = blockHash
	lastBlock := m.LastMinedBlock
	m.blockNumbers[m.NextBlock.Height.AsInt().Uint64()] = m.NextBlock
	m.blockHashes[m.NextBlock.HeaderHash] = m.NextBlock
	m.parentHashes[*m.NextBlock] = lastBlock.HeaderHash
	m.LastMinedBlock = m.NextBlock
	for _, event := range m.nextMsgs[m.NextBlock] {
		for _, ru := range m.rollups {
			ru.events[m.NextBlock] = append(ru.events[m.NextBlock], event.Event)
		}
		for _, watcher := range m.challengeWatchers {
			watcher[m.NextBlock] = append(watcher[m.NextBlock], event.Event)
		}
		m.pubchan <- event
	}
	fmt.Println("mined block number", m.NextBlock)
	newBlock := new(structures.BlockId)
	newBlock.Height = common.NewTimeBlocks(new(big.Int).Add(m.LastMinedBlock.Height.AsInt(), big.NewInt(1)))
	m.NextBlock = newBlock
	blockEvent := arbbridge.NewTimeEvent{
		arbbridge.ChainInfo{
			BlockId: m.LastMinedBlock,
		},
	}

	m.blockMutex.Unlock()
	fmt.Println("publishing NewTimeEvent block - ", blockEvent.BlockId)
	m.pubMsg(arbbridge.MaybeEvent{
		Event: blockEvent,
	})
}

func addToken(source common.Address, data value.Value, destination common.Address, amount *big.Int) {
	if amount.Cmp(big.NewInt(0)) == 0 {
		return
	}

}

// TODO have to figure out all this data
func emitFinalizedUnanimousAssertion(vm common.Address, unanHash [32]byte) {
	//raw := types.Log{
	//	Address:     vm,
	//	Topics:      nil,
	//	Data:        nil,
	//	BlockNumber: MockEthData.BlockNumber,
	//	TxHash:      common.Hash{},
	//	TxIndex:     0,
	//	BlockHash:   common.Hash{},
	//	Index:       0,
	//	Removed:     false,
	//}
	//for _, unanAss := range MockEthData.FinUnanAssList {
	//	unanAss <- &channellauncher.ArbChannelFinalizedUnanimousAssertion{unanHash, raw}
	//}
}
func pullPendingMessages(address common.Address) [32]byte {
	//bytes32 messages = pending[msg.sender];
	//pending[msg.sender] = ArbValue.hashEmptyTuple();
	//messages := MockEthData.pending[address]
	//MockEthData.pending[address] = value.NewEmptyTuple().Hash()
	//return messages
	return *new([32]byte)
}

func sendUnpaidMessage(src common.Address, dest common.Address, tokType [21]byte, amount *big.Int, data []byte) error {

	//srcBt := MockEthData.balanceTrackers[src]
	//err := srcBt.Spend(tokType, amount)
	//if err != nil {
	//	return err
	//}
	//destBt := MockEthData.balanceTrackers[dest]
	//destBt.Add(tokType, amount)

	// create message to send
	// send message

	return nil
	//        _deliverMessage(
	//            _destination,
	//            _tokenType,
	//            _value,
	//            _sender,
	//            _data
	//        );
	//        if (pending[_destination] != 0) {
	//            bytes32 dataHash = ArbValue.deserializeValueHash(_data);
	//            bytes32 txHash = keccak256(
	//                abi.encodePacked(
	//                    _destination,
	//                    dataHash,
	//                    _value,
	//                    _tokenType
	//                )
	//            );
	//            ArbValue.Value[] memory dataValues = new ArbValue.Value[](4);
	//            dataValues[0] = ArbValue.newHashOnlyValue(dataHash);
	//            dataValues[1] = ArbValue.newIntValue(block.timestamp);
	//            dataValues[2] = ArbValue.newIntValue(block.number);
	//            dataValues[3] = ArbValue.newIntValue(uint(txHash));
	//
	//            ArbValue.Value[] memory values = new ArbValue.Value[](4);
	//            values[0] = ArbValue.newTupleValue(dataValues);
	//            values[1] = ArbValue.newIntValue(uint256(_sender));
	//            values[2] = ArbValue.newIntValue(_value);
	//            values[3] = ArbValue.newIntValue(uint256(bytes32(_tokenType)));
	//            bytes32 messageHash =  ArbValue.newTupleValue(values).hash().hash;
	//
	//            pending[_destination] = ArbValue.hashTupleValue([
	//                ArbValue.newIntValue(0),
	//                ArbValue.newHashOnlyValue(pending[_destination]),
	//                ArbValue.newHashOnlyValue(messageHash)
	//            ]);
	//        }
	//
	//        emit IGlobalPendingInbox.MessageDelivered(
	//            _destination,
	//            _sender,
	//            _tokenType,
	//            _value,
	//            _data
	//        );
	//}

}

//function calculateBeforeValues(
//bytes21[] memory _tokenTypes,
//uint16[] memory _messageTokenNums,
//uint256[] memory _messageAmounts
//)
//public
//pure
//returns(uint256[] memory)
//{
//uint messageCount = _messageTokenNums.length;
//uint256[] memory beforeBalances = new uint256[](_tokenTypes.length);
//
//for (uint i = 0; i < messageCount; i++) {
//uint16 tokenNum = _messageTokenNums[i];
//if (_tokenTypes[tokenNum][20] == 0x00) {
//beforeBalances[tokenNum] += _messageAmounts[i];
//} else {
//require(beforeBalances[tokenNum] == 0, "Can't include NFT token twice");
//require(_messageAmounts[i] != 0, "NFT token must have non-zero id");
//beforeBalances[tokenNum] = _messageAmounts[i];
//}
//}
//return beforeBalances;
//}
