package goarbitrum

import (
	"bytes"
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/rpc/json"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/rollupvalidator"
)

type ValidatorProxy interface {
	//SendMessage(val value.Value, hexPubkey string, signature []byte) ([]byte, error)
	GetMessageResult(ctx context.Context, txHash []byte) (value.Value, bool, error)
	GetAssertionCount(ctx context.Context) (int, error)
	GetVMInfo(ctx context.Context) (string, error)
	FindLogs(
		ctx context.Context,
		fromHeight, toHeight int64,
		address []byte,
		topics [][32]byte,
	) ([]*rollupvalidator.LogInfo, error)
	CallMessage(ctx context.Context, contract common.Address, sender common.Address, data []byte) (value.Value, error)
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

func (vp *ValidatorProxyImpl) doCall(
	ctx context.Context,
	methodName string,
	request interface{},
	response interface{},
) error {
	message, err := json.EncodeClientRequest("Validator."+methodName, request)
	if err != nil {
		log.Println("ValProxy.doCall: error in json.Enc:", err)
		return err
	}
	req, err := http.NewRequest("POST", vp.url, bytes.NewBuffer(message))
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

func (vp *ValidatorProxyImpl) GetMessageResult(ctx context.Context, txHash []byte) (value.Value, bool, error) {
	request := &rollupvalidator.GetMessageResultArgs{
		TxHash: hexutil.Encode(txHash),
	}
	var response rollupvalidator.GetMessageResultReply
	if err := vp.doCall(ctx, "GetMessageResult", request, &response); err != nil {
		log.Println("ValProxy.GetMessageResult: doCall returned error:", err)
		return nil, false, err
	}
	if !response.Found {
		return nil, false, nil
	}
	buf, err := hexutil.Decode(response.RawVal)
	if err != nil {
		log.Println("GetMessageResult error:", err)
		return nil, false, err
	}
	val, err := value.UnmarshalValue(bytes.NewReader(buf))
	if err != nil {
		log.Println("ValProxy.GetMessageResult: UnmarshalValue returned error:", err)
		return nil, false, err
	}
	return val, true, err
}

func (vp *ValidatorProxyImpl) GetAssertionCount(ctx context.Context) (int, error) {
	request := &struct{}{}
	var response rollupvalidator.GetAssertionCountReply
	if err := vp.doCall(ctx, "GetAssertionCount", request, &response); err != nil {
		return 0, err
	}
	return int(response.AssertionCount), nil
}

func (vp *ValidatorProxyImpl) GetVMInfo(ctx context.Context) (string, error) {
	request := &struct{}{}
	var response rollupvalidator.GetVMInfoReply
	if err := vp.doCall(ctx, "GetVMInfo", request, &response); err != nil {
		return "", err
	}
	return response.VmID, nil
}

func (vp *ValidatorProxyImpl) FindLogs(
	ctx context.Context,
	fromHeight, toHeight int64,
	address []byte,
	topics [][32]byte,
) ([]*rollupvalidator.LogInfo, error) {
	request := &rollupvalidator.FindLogsArgs{
		FromHeight: _encodeInt(fromHeight),
		ToHeight:   _encodeInt(toHeight),
		Address:    hexutil.Encode(address),
		Topics:     _encodeByteArraySlice(topics),
	}
	var response rollupvalidator.FindLogsReply
	if err := vp.doCall(ctx, "FindLogs", request, &response); err != nil {
		return nil, err
	}
	return response.Logs, nil
}

func (vp *ValidatorProxyImpl) CallMessage(
	ctx context.Context,
	contract common.Address,
	sender common.Address,
	data []byte,
) (value.Value, error) {
	request := &rollupvalidator.CallMessageArgs{
		ContractAddress: hexutil.Encode(contract[:]),
		Sender:          hexutil.Encode(sender[:]),
		Data:            hexutil.Encode(data),
	}
	var response rollupvalidator.CallMessageReply
	if err := vp.doCall(ctx, "CallMessage", request, &response); err != nil {
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
