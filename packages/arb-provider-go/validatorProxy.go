package goarbitrum

import (
	"bytes"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/evm"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/rpc/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator-core/validatorserver"
)

type ValidatorProxy interface {
	//SendMessage(val value.Value, hexPubkey string, signature []byte) ([]byte, error)
	GetMessageResult(txHash []byte) (evm.TxInfo, error)
	GetAssertionCount() (int, error)
	GetVMInfo() (string, error)
	FindLogs(fromHeight, toHeight *uint64, addresses []common.Address, topics [][]common.Hash) ([]evm.FullLog, error)
	CallMessage(contract common.Address, sender common.Address, data []byte) (value.Value, error)
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

func _encodeByteArraySlice(slice []common.Hash) []string {
	ret := make([]string, len(slice))
	for i, arr := range slice {
		ret[i] = hexutil.Encode(arr[:])
	}
	return ret
}

func _encodeAddressArraySlice(slice []common.Address) []string {
	ret := make([]string, len(slice))
	for i, arr := range slice {
		ret[i] = hexutil.Encode(arr[:])
	}
	return ret
}

func (vp *ValidatorProxyImpl) doCall(methodName string, request interface{}, response interface{}) error {
	message, err := json.EncodeClientRequest("Validator."+methodName, request)
	if err != nil {
		log.Println("ValProxy.doCall: error in json.Enc:", err)
		return err
	}
	req, err := http.NewRequest("POST", vp.url, bytes.NewBuffer(message))
	if err != nil {
		return err
	}
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

//
//func (vp *ValidatorProxyImpl) SendMessage(val value.Value, hexPubkey string, signature []byte) ([]byte, error) {
//	var buf bytes.Buffer
//	if err := value.MarshalValue(val, &buf); err != nil {
//		log.Println("ValProxy.SendMessage: marshaling error:", err)
//		return nil, err
//	}
//	request := &rollupvalidator.SendMessageArgs{
//		Data:      hexutil.Encode(buf.Bytes()),
//		Pubkey:    hexPubkey,
//		Signature: hexutil.Encode(signature),
//	}
//	var response rollupvalidator.SendMessageReply
//	if err := vp.doCall("SendMessage", request, &response); err != nil {
//		log.Println("ValProxy.SendMessage: error returned from doCall:", err)
//		return nil, err
//	}
//	bs, err := hexutil.Decode(response.TxHash)
//	if err != nil {
//		log.Println("ValProxy.SendMessage error:", err)
//	}
//	return bs, err
//}

func (vp *ValidatorProxyImpl) GetMessageResult(txHash []byte) (evm.TxInfo, error) {
	request := &validatorserver.GetMessageResultArgs{
		TxHash: hexutil.Encode(txHash),
	}
	var response validatorserver.GetMessageResultReply
	if err := vp.doCall("GetMessageResult", request, &response); err != nil {
		log.Println("ValProxy.GetMessageResult: doCall returned error:", err)
		return evm.TxInfo{}, err
	}
	return response.Tx.Unmarshal()
}

func (vp *ValidatorProxyImpl) GetAssertionCount() (int, error) {
	request := &struct{}{}
	var response validatorserver.GetAssertionCountReply
	if err := vp.doCall("GetAssertionCount", request, &response); err != nil {
		return 0, err
	}
	return int(response.AssertionCount), nil
}

func (vp *ValidatorProxyImpl) GetVMInfo() (string, error) {
	request := &struct{}{}
	var response validatorserver.GetVMInfoReply
	if err := vp.doCall("GetVMInfo", request, &response); err != nil {
		return "", err
	}
	return response.VmID, nil
}

func (vp *ValidatorProxyImpl) FindLogs(fromHeight, toHeight *uint64, addresses []common.Address, topicGroups [][]common.Hash) ([]evm.FullLog, error) {
	tgs := make([]*validatorserver.TopicGroup, 0, len(topicGroups))
	for _, topicGroup := range topicGroups {
		tgs = append(tgs, &validatorserver.TopicGroup{Topics: _encodeByteArraySlice(topicGroup)})
	}
	request := &validatorserver.FindLogsArgs{
		FromHeight:  _encodeInt(fromHeight),
		ToHeight:    _encodeInt(toHeight),
		Addresses:   _encodeAddressArraySlice(addresses),
		TopicGroups: tgs,
	}
	var response validatorserver.FindLogsReply
	if err := vp.doCall("FindLogs", request, &response); err != nil {
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

func (vp *ValidatorProxyImpl) CallMessage(contract common.Address, sender common.Address, data []byte) (value.Value, error) {
	request := &validatorserver.CallMessageArgs{
		ContractAddress: hexutil.Encode(contract[:]),
		Sender:          hexutil.Encode(sender[:]),
		Data:            hexutil.Encode(data),
	}
	var response validatorserver.CallMessageReply
	if err := vp.doCall("CallMessage", request, &response); err != nil {
		return nil, err
	}
	retBuf, err := hexutil.Decode(response.RawVal)
	if err != nil {
		log.Println("CallMessage error:", err)
		return nil, err
	}
	retVal, err := value.UnmarshalValue(bytes.NewReader(retBuf))
	if err != nil {
		log.Println("ValProxy.CallMessage: UnmarshalValue returned error:", err)
	}
	return retVal, err
}
