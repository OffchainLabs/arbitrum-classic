package goarbitrum

import (
	"bytes"
	"log"
	"net/http"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/gorilla/rpc/json"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/coordinator"
)

type ValidatorProxy interface {
	SendMessage(val value.Value, hexPubkey string, signature []byte) ([]byte, error)
	GetMessageResult(txHash []byte) (value.Value, bool, error)
	GetAssertionCount() (int, error)
	GetVMInfo() (string, error)
	FindLogs(fromHeight, toHeight int64, address []byte, topics [][32]byte) ([]*coordinator.LogInfo, error)
	CallMessage(val value.Value, sender common.Address) (value.Value, error)
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

func _encodeInt(i int64) string {
	return "0x" + strconv.FormatInt(i, 16)
}

func _encodeByteArraySlice(slice [][32]byte) []string {
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

func (vp *ValidatorProxyImpl) SendMessage(val value.Value, hexPubkey string, signature []byte) ([]byte, error) {
	var buf bytes.Buffer
	if err := value.MarshalValue(val, &buf); err != nil {
		log.Println("ValProxy.SendMessage: marshaling error:", err)
		return nil, err
	}
	request := &coordinator.SendMessageArgs{
		Data:      hexutil.Encode(buf.Bytes()),
		Pubkey:    hexPubkey,
		Signature: hexutil.Encode(signature),
	}
	var response coordinator.SendMessageReply
	if err := vp.doCall("SendMessage", request, &response); err != nil {
		log.Println("ValProxy.SendMessage: error returned from doCall:", err)
		return nil, err
	}
	bs, err := hexutil.Decode(response.TxHash)
	if err != nil {
		log.Println("ValProxy.SendMessage error:", err)
	}
	return bs, err
}

func (vp *ValidatorProxyImpl) GetMessageResult(txHash []byte) (value.Value, bool, error) {
	request := &coordinator.GetMessageResultArgs{
		TxHash: hexutil.Encode(txHash),
	}
	var response coordinator.GetMessageResultReply
	if err := vp.doCall("GetMessageResult", request, &response); err != nil {
		log.Println("ValProxy.GetMessageResult: doCall returned error:", err)
		return nil, false, err
	}
	if response.Found {
		buf, err := hexutil.Decode(response.RawVal)
		if err != nil {
			log.Println("GetMessageResult error:", err)
			return nil, false, err
		}
		val, err := value.UnmarshalValue(bytes.NewReader(buf))
		if err != nil {
			log.Println("ValProxy.GetMessageResult: UnmarshalValue returned error:", err)
		}
		return val, true, err
	} else {
		return nil, false, nil
	}
}

func (vp *ValidatorProxyImpl) GetAssertionCount() (int, error) {
	request := &struct{}{}
	var response coordinator.GetAssertionCountReply
	if err := vp.doCall("GetAssertionCount", request, &response); err != nil {
		return 0, err
	}
	return int(response.AssertionCount), nil
}

func (vp *ValidatorProxyImpl) GetVMInfo() (string, error) {
	request := &struct{}{}
	var response coordinator.GetVMInfoReply
	if err := vp.doCall("GetVMInfo", request, &response); err != nil {
		return "", err
	}
	return response.VmID, nil
}

func (vp *ValidatorProxyImpl) FindLogs(fromHeight, toHeight int64, address []byte, topics [][32]byte) ([]*coordinator.LogInfo, error) {
	request := &coordinator.FindLogsArgs{
		FromHeight: _encodeInt(fromHeight),
		ToHeight:   _encodeInt(toHeight),
		Address:    hexutil.Encode(address),
		Topics:     _encodeByteArraySlice(topics),
	}
	var response coordinator.FindLogsReply
	if err := vp.doCall("FindLogs", request, &response); err != nil {
		return nil, err
	}
	return response.Logs, nil
}

func (vp *ValidatorProxyImpl) CallMessage(val value.Value, sender common.Address) (value.Value, error) {
	var buf bytes.Buffer
	if err := value.MarshalValue(val, &buf); err != nil {
		return nil, err
	}
	request := &coordinator.CallMessageArgs{
		Data:   hexutil.Encode(buf.Bytes()),
		Sender: hexutil.Encode(sender[:]),
	}
	var response coordinator.CallMessageReply
	if err := vp.doCall("CallMessage", request, &response); err != nil {
		return nil, err
	}
	retBuf, err := hexutil.Decode(response.RawVal)
	if err != nil {
		log.Println("GetMessageResult error:", err)
		return nil, err
	}
	retVal, err := value.UnmarshalValue(bytes.NewReader(retBuf))
	if err != nil {
		log.Println("ValProxy.GetMessageResult: UnmarshalValue returned error:", err)
	}
	return retVal, err
}
