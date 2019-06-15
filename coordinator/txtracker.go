package coordinator

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/offchainlabs/arb-avm/evm"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-validator/valmessage"
	"log"
	"math/big"
	"strconv"
)

type ValidatorRequest interface {
}

type AssertionCountRequest struct {
	resultChan chan<- int
}

type TxRequest struct {
	txHash     [32]byte
	resultChan chan<- TxInfo
}

type FindLogsRequest struct {
	fromHeight *int64
	toHeight   *int64
	address    *big.Int
	topics     [][32]byte

	resultChan chan<- []LogInfo
}

type LogsInfo struct {
	msg  evm.EthMsg
	Logs []evm.Log
}

type TxInfo struct {
	Found          bool
	assertionIndex int
	RawVal         value.Value
}

type AssertionInfo struct {
	TxLogs []LogsInfo
}

type LogResponse struct {
	Log evm.Log
	Msg evm.EthMsg
}

func (a *AssertionInfo) FindLogs(address *big.Int, topics [][32]byte) []LogResponse {
	logs := make([]LogResponse, 0)
	for _, txLogs := range a.TxLogs {
		for _, evmLog := range txLogs.Logs {
			if address != nil && !value.NewIntValue(address).Equal(evmLog.ContractId) {
				continue
			}

			if len(topics) > len(evmLog.Topics) {
				continue
			}

			for i, topic := range topics {
				if topic != evmLog.Topics[i] {
					continue
				}
			}
			logs = append(logs, LogResponse{evmLog, txLogs.msg})
		}
	}
	return logs
}

func NewAssertionInfo() *AssertionInfo {
	logs := make([]LogsInfo, 0)
	return &AssertionInfo{logs}
}


type TxTracker struct {
	txRequestIndex int
	transactions   map[[32]byte]TxInfo
	assertionInfo  []*AssertionInfo
	accountNonces  map[common.Address]uint64
	vmId           [32]byte
}

func NewTxTracker(vmId [32]byte) *TxTracker {
	return &TxTracker{
		txRequestIndex: 0,
		transactions:   make(map[[32]byte]TxInfo),
		assertionInfo:  make([]*AssertionInfo, 0),
		accountNonces:  make(map[common.Address]uint64),
		vmId:           vmId,
	}
}

func (tr *TxTracker) processFinalizedAssertion(assertion valmessage.FinalizedAssertion) {
	log.Println("Coordinator produced finalized assertion")
	info := NewAssertionInfo()
	for _, res := range assertion.NewLogs() {
		evmVal, err := evm.ProcessLog(res)
		if err != nil {
			log.Printf("VM produced invalid evm result: %v\n", err)
		}

		msg := evmVal.GetEthMsg()
		msgHash := msg.MsgHash(tr.vmId)

		log.Println("Coordinator got response for", hexutil.Encode(msgHash[:]))
		txInfo := TxInfo{
			Found:          true,
			assertionIndex: 0,
			RawVal:         res,
		}
		txInfo.assertionIndex = len(tr.assertionInfo)
		switch evmVal := evmVal.(type) {
		case evm.Stop:
			info.TxLogs = append(info.TxLogs, LogsInfo{evmVal.Msg, evmVal.Logs})
		case evm.Return:
			info.TxLogs = append(info.TxLogs, LogsInfo{evmVal.Msg, evmVal.Logs})
		case evm.Revert:
		}
		tr.transactions[msgHash] = txInfo
	}
	tr.assertionInfo = append(tr.assertionInfo, info)
}

func (tr *TxTracker) processRequest(request ValidatorRequest) {
	switch request := request.(type) {
	case AssertionCountRequest:
		request.resultChan <- len(tr.assertionInfo) - 1
	case TxRequest:
		tx, ok := tr.transactions[request.txHash]
		if ok {
			request.resultChan <- tx
		} else {
			request.resultChan <- TxInfo{Found: false}
		}
	case FindLogsRequest:
		startHeight := int64(0)
		endHeight := int64(len(tr.assertionInfo))
		if request.fromHeight != nil && *request.fromHeight > int64(0) {
			startHeight = *request.fromHeight
		}
		if request.toHeight != nil {
			endHeight = *request.toHeight + 1
		}
		logs := make([]LogInfo, 0)
		if startHeight >= int64(len(tr.assertionInfo)) {
			request.resultChan <- logs
			break
		}
		assertions := tr.assertionInfo[startHeight:endHeight]

		for i, assertion := range assertions {
			assertionLogs := assertion.FindLogs(request.address, request.topics)
			for j, evmLog := range assertionLogs {
				addressBytes := evmLog.Log.ContractId.ToBytes()
				topicStrings := make([]string, 0, len(evmLog.Log.Topics))
				for _, topic := range evmLog.Log.Topics {
					topicStrings = append(topicStrings, hexutil.Encode(topic[:]))
				}
				txHash := evmLog.Msg.MsgHash(tr.vmId)
				logs = append(logs, LogInfo{
					Address:          hexutil.Encode(addressBytes[12:]),
					BlockHash:        hexutil.Encode(txHash[:]),
					BlockNumber:      "0x" + strconv.FormatInt(int64(i), 16),
					Data:             hexutil.Encode(evmLog.Log.Data[:]),
					LogIndex:         "0x" + strconv.FormatInt(int64(j), 16),
					Topics:           topicStrings,
					TransactionIndex: "0x00",
					TransactionHash:  hexutil.Encode(txHash[:]),
				})
			}
		}
		request.resultChan <- logs
	}
}

func (tr *TxTracker) HandleTxResults(completedCalls chan valmessage.FinalizedAssertion, requests chan ValidatorRequest) {
	for {
		select {
		case finalizedAssertion := <-completedCalls:
			tr.processFinalizedAssertion(finalizedAssertion)
		case request := <-requests:
			tr.processRequest(request)

		}
	}
}