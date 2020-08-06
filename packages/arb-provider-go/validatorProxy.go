package goarbitrum

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/rpc/json"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/offchainlabs/arbitrum/packages/arb-evm/evm"
	"github.com/offchainlabs/arbitrum/packages/arb-evm/message"
	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
)

var Namespace = "Aggregator"

type ValidatorProxy interface {
	GetBlockCount(ctx context.Context) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) (common.Hash, error)
	BlockInfo(ctx context.Context, height uint64) (machine.BlockInfo, error)
	GetRequestResult(ctx context.Context, txHash common.Hash) (value.Value, error)
	GetChainAddress(ctx context.Context) (ethcommon.Address, error)
	FindLogs(ctx context.Context, fromHeight, toHeight *uint64, addresses []ethcommon.Address, topics [][]ethcommon.Hash) ([]evm.FullLog, error)
	Call(ctx context.Context, msg message.ContractTransaction, sender ethcommon.Address) (value.Value, error)
	PendingCall(ctx context.Context, msg message.ContractTransaction, sender ethcommon.Address) (value.Value, error)
}

type ValidatorProxyImpl struct {
	url string
}

func NewValidatorProxyImpl(url string) ValidatorProxy {
	if url == "" {
		url = "http://localhost:1235"
	}
	return &ValidatorProxyImpl{url}
}

func _encodeInt(i *uint64) string {
	if i == nil {
		return ""
	}

	return "0x" + strconv.FormatUint(*i, 16)
}

func _encodeByteArraySlice(slice []ethcommon.Hash) []string {
	ret := make([]string, len(slice))
	for i, arr := range slice {
		ret[i] = hexutil.Encode(arr[:])
	}
	return ret
}

func _encodeAddressArraySlice(slice []ethcommon.Address) []string {
	ret := make([]string, len(slice))
	for i, arr := range slice {
		ret[i] = hexutil.Encode(arr[:])
	}
	return ret
}

func (vp *ValidatorProxyImpl) SendTransaction(ctx context.Context, tx *types.Transaction) (common.Hash, error) {
	txData, err := rlp.EncodeToBytes(tx)
	if err != nil {
		return common.Hash{}, err
	}
	request := &evm.SendTransactionArgs{SignedTransaction: hexutil.Encode(txData)}
	var response evm.SendTransactionReply
	if err := vp.doCall(ctx, "SendTransaction", request, &response); err != nil {
		log.Println("ValProxy.SendMessage: error returned from doCall:", err)
		return common.Hash{}, err
	}
	return common.HexToHash(response.TransactionHash), nil
}

func (vp *ValidatorProxyImpl) doCall(ctx context.Context, methodName string, request interface{}, response interface{}) error {
	msg, err := json.EncodeClientRequest(Namespace+"."+methodName, request)
	if err != nil {
		log.Println("ValProxy.doCall: error in json.Enc:", err)
		return err
	}
	req, err := http.NewRequest("POST", vp.url, bytes.NewBuffer(msg))
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Println("doCall error:", err)
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	ret := json.DecodeClientResponse(resp.Body, response)
	if ret != nil {
		log.Println("ValProxy.doCall: error in json.Dec from", methodName, ":", ret)
	}
	return ret
}

func (vp *ValidatorProxyImpl) BlockInfo(ctx context.Context, height uint64) (machine.BlockInfo, error) {
	request := &evm.BlockInfoArgs{
		Height: height,
	}
	var response evm.BlockInfoReply
	if err := vp.doCall(ctx, "BlockInfo", request, &response); err != nil {
		return machine.BlockInfo{}, err
	}
	bloomBytes, err := hexutil.Decode(response.Bloom)
	if err != nil {
		return machine.BlockInfo{}, err
	}
	return machine.BlockInfo{
		Hash:         common.NewHashFromEth(ethcommon.HexToHash(response.Hash)),
		StartLog:     response.StartLog,
		LogCount:     response.LogCount,
		StartMessage: response.StartMessage,
		MessageCount: response.MessageCount,
		Bloom:        types.BytesToBloom(bloomBytes),
	}, nil
}

func (vp *ValidatorProxyImpl) GetRequestResult(ctx context.Context, txHash common.Hash) (value.Value, error) {
	request := &evm.GetRequestResultArgs{
		TxHash: hexutil.Encode(txHash[:]),
	}
	var response evm.GetRequestResultReply
	if err := vp.doCall(ctx, "GetRequestResult", request, &response); err != nil {
		log.Println("ValProxy.GetRequestResult: doCall returned error:", err)
		return nil, err
	}

	if len(response.RawVal) == 0 {
		return nil, nil
	}

	data, err := hexutil.Decode(response.RawVal)
	if err != nil {
		return nil, err
	}
	val, err := value.UnmarshalValue(bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (vp *ValidatorProxyImpl) GetChainAddress(ctx context.Context) (ethcommon.Address, error) {
	request := &evm.GetChainAddressArgs{}
	var response evm.GetChainAddressReply
	if err := vp.doCall(ctx, "GetChainAddress", request, &response); err != nil {
		return ethcommon.Address{}, err
	}
	return ethcommon.HexToAddress(response.ChainAddress), nil
}

func (vp *ValidatorProxyImpl) GetBlockCount(ctx context.Context) (uint64, error) {
	request := &evm.BlockCountArgs{}
	var response evm.BlockCountReply
	if err := vp.doCall(ctx, "GetBlockCount", request, &response); err != nil {
		return 0, err
	}
	return response.Height, nil
}

func (vp *ValidatorProxyImpl) FindLogs(ctx context.Context, fromHeight, toHeight *uint64, addresses []ethcommon.Address, topicGroups [][]ethcommon.Hash) ([]evm.FullLog, error) {
	tgs := make([]*evm.TopicGroup, 0, len(topicGroups))
	for _, topicGroup := range topicGroups {
		tgs = append(tgs, &evm.TopicGroup{Topics: _encodeByteArraySlice(topicGroup)})
	}
	request := &evm.FindLogsArgs{
		FromHeight:  _encodeInt(fromHeight),
		ToHeight:    _encodeInt(toHeight),
		Addresses:   _encodeAddressArraySlice(addresses),
		TopicGroups: tgs,
	}
	var response evm.FindLogsReply
	if err := vp.doCall(ctx, "FindLogs", request, &response); err != nil {
		log.Println("Error finding logs", err)
		return nil, err
	}

	logs := make([]evm.FullLog, 0, len(response.Logs))
	for _, l := range response.Logs {
		parsedLog, err := l.Unmarshal()
		if err != nil {
			return nil, err
		}
		logs = append(logs, parsedLog)
	}
	return logs, nil
}

func hexToValue(rawVal string) (value.Value, error) {
	retBuf, err := hexutil.Decode(rawVal)
	if err != nil {
		return nil, err
	}
	return value.UnmarshalValue(bytes.NewReader(retBuf))
}

func (vp *ValidatorProxyImpl) Call(ctx context.Context, msg message.ContractTransaction, sender ethcommon.Address) (value.Value, error) {
	request := &evm.CallMessageArgs{
		Data:   hexutil.Encode(msg.AsData()),
		Sender: hexutil.Encode(sender[:]),
	}
	var response evm.CallMessageReply
	if err := vp.doCall(ctx, "Call", request, &response); err != nil {
		return nil, err
	}
	return hexToValue(response.RawVal)
}

func (vp *ValidatorProxyImpl) PendingCall(ctx context.Context, msg message.ContractTransaction, sender ethcommon.Address) (value.Value, error) {
	request := &evm.CallMessageArgs{
		Data:   hexutil.Encode(msg.AsData()),
		Sender: hexutil.Encode(sender[:]),
	}
	var response evm.CallMessageReply
	if err := vp.doCall(ctx, "PendingCall", request, &response); err != nil {
		return nil, err
	}
	return hexToValue(response.RawVal)
}
