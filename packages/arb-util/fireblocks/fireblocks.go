/*
 * Copyright 2021, Offchain Labs, Inc.
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

package fireblocks

import (
	"bytes"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks/accounttype"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks/operationtype"
)

var logger = log.With().Caller().Stack().Str("component", "fireblocks").Logger()

type Fireblocks struct {
	apiKey     string
	assetId    string
	baseUrl    string
	signKey    *rsa.PrivateKey
	sourceId   string
	sourceType accounttype.AccountType
}

type CreateNewTransactionBody struct {
	AssetId         string                          `json:"assetId"`
	Source          TransferPeerPath                `json:"source"`
	Destination     DestinationTransferPeerPath     `json:"destination"`
	Amount          string                          `json:"amount"`
	Fee             string                          `json:"fee,omitempty"`
	GasPrice        string                          `json:"gasPrice,omitempty"`
	GasLimit        string                          `json:"gasLimit,omitempty"`
	NetworkFee      string                          `json:"networkFee,omitempty"`
	FeeLevel        string                          `json:"feeLevel,omitempty"`
	MaxFee          string                          `json:"maxFee,omitempty"`
	FailOnLowFee    bool                            `json:"failOnLowFee,omitempty"`
	Note            string                          `json:"note,omitempty"`
	Operation       operationtype.OperationType     `json:"operation,omitempty"`
	CustomerRefId   string                          `json:"customerRefId,omitempty"`
	Destinations    []TransactionRequestDestination `json:"destinations,omitempty"`
	ExtraParameters TransactionExtraParameters      `json:"extraParameters"`
}

type EstimateTransactionFeeRequestBody struct {
	AssetId     string                      `json:"assetId"`
	Amount      string                      `json:"amount"`
	Source      TransferPeerPath            `json:"source"`
	Destination DestinationTransferPeerPath `json:"destination"`
	Operation   operationtype.OperationType `json:"operation"`
}

type EstimatedTransactionFeeResponse struct {
	Low    TransactionFee `json:"low"`
	Medium TransactionFee `json:"medium"`
	High   TransactionFee `json:"high"`
}

type TransactionFee struct {
	FeePerByte string `json:"feePerByte,omitempty"`
	GasPrice   string `json:"gasPrice,omitempty"`
	GasLimit   string `json:"gasLimit,omitempty"`
	NetworkFee string `json:"networkFee,omitempty"`
}

type TransactionExtraParameters struct {
	ContractCallData string `json:"contractCallData"`
}

type TransferPeerPath struct {
	Type accounttype.AccountType `json:"type"`
	Id   string                  `json:"id"`
}

type DestinationTransferPeerPath struct {
	Type           accounttype.AccountType `json:"type"`
	Id             string                  `json:"id,omitempty"`
	OneTimeAddress OneTimeAddress          `json:"oneTimeAddress,omitempty"`
}

func NewDestinationTransferPeerPath(destinationType accounttype.AccountType, destinationId string, destinationTag string) *DestinationTransferPeerPath {
	destination := DestinationTransferPeerPath{Type: destinationType}
	if destination.Type == accounttype.OneTimeAddress {
		destination.OneTimeAddress = OneTimeAddress{
			Address: destinationId,
		}
		if len(destinationTag) > 0 {
			destination.OneTimeAddress.Tag = destinationTag
		}
	} else {
		destination.Id = destinationId
	}

	return &destination
}

type OneTimeAddress struct {
	Address string `json:"address"`
	Tag     string `json:"tag"`
}

type CreateTransactionResponse struct {
	Id     string `json:"id"`
	Status string `json:"status"`
}

type TransactionDetails struct {
	Id                            string                   `json:"id"`
	AssetId                       string                   `json:"assetId"`
	Source                        TransferPeerPathResponse `json:"source"`
	Destination                   TransferPeerPathResponse `json:"destination"`
	RequestedAmount               int64                    `json:"requestedAmount"`
	AmountInfo                    AmountInfo               `json:"amountInfo"`
	FeeInfo                       FeeInfo                  `json:"feeInfo"`
	Amount                        int64                    `json:"amount"`
	NetAmount                     float32                  `json:"NetAmount"`
	AmountUSD                     float32                  `json:"amountUSD"`
	ServiceFee                    float32                  `json:"serviceFee"`
	NetworkFee                    float32                  `json:"networkFee"`
	CreatedAt                     int64                    `json:"createdAt"`
	LastUpdate                    int64                    `json:"lastUpdate"`
	Status                        string                   `json:"transactionStatus"`
	TxHash                        string                   `json:"txHash"`
	SubStatus                     string                   `json:"transactionSubStatus"`
	DestinationAddress            string                   `json:"destinationAddress"`
	DestinationAddressDescription string                   `json:"destinationAddressDescription"`
	DestinationTag                string                   `json:"destinationTag"`
	SignedBy                      []string                 `json:"signedBy"`
	CreatedBy                     string                   `json:"createdBy"`
	RejectedBy                    string                   `json:"rejectedBy"`
	AddressType                   string                   `json:"addressType"`
	Note                          string                   `json:"note"`
	ExchangeTxId                  string                   `json:"exchangeTxId"`
	FeeCurrency                   string                   `json:"feeCurrency"`
	Operation                     string                   `json:"operation"`
	AmlScreeningResult            AmlScreeningResult       `json:"amlScreeningResult"`
	CustomerRefId                 string                   `json:"customerRefId"`
	NumOfConfirmations            int                      `json:"numOfConfirmations"`
	NetworkRecords                []string                 `json:"networkRecords"`
	ReplacedTxHash                string                   `json:"replacedTxHash"`
	Destinations                  []DestinationsResponse   `json:"destinations"`
	SignedMessages                []SignedMessage          `json:"signedMessages"`
	ExtraParameters               string                   `json:"extraParameters"`
}

type TransferPeerPathResponse struct {
	Type    string `json:"type"`
	Id      string `json:"id"`
	Name    string `json:"name"`
	SubType string `json:"subType"`
}

type AmountInfo struct {
	Amount          string `json:"amount"`
	RequestedAmount string `json:"requestedAmount"`
	NetAmount       string `json:"netAmount"`
	AmountUSD       string `json:"amountUSD"`
}

type FeeInfo struct {
	NetworkFee string `json:"networkFee"`
	ServiceFee string `json:"serviceFee"`
}

type NetworkRecord struct {
	Source             TransferPeerPathResponse `json:"source"`
	Destination        TransferPeerPathResponse `json:"destination"`
	TxHash             string                   `json:"txHash"`
	NetworkFee         int64                    `json:"networkFee"`
	AssetId            string                   `json:"assetId"`
	NetAmount          int64                    `json:"netAmount"`
	Status             string                   `json:"status"`
	Type               string                   `json:"type"`
	DestinationAddress string                   `json:"destinationAddress"`
}

type AmlScreeningResult struct {
	Operation string `json:"amlScreeningResult"`
}

type SignedMessage struct {
	Content        string     `json:"content"`
	Algorithm      string     `json:"algorithm"`
	DerivationPath []int      `json:"derivationPath"`
	Signature      Dictionary `json:"signature"`
	PublicKey      string     `json:"publicKey"`
}

type Dictionary map[string]interface{}

type DestinationsResponse struct {
	Amount                        string             `json:"amount"`
	Destination                   string             `json:"destination"`
	AmountUSD                     string             `json:"amountUSD"`
	DestinationAddress            string             `json:"destinationAddress"`
	DestinationAddressDescription string             `json:"destinationAddressDescription"`
	AmlScreeningResult            AmlScreeningResult `json:"amlScreeningResult"`
	CustomerRefId                 string             `json:"customerRefId"`
}

type VaultAccount struct {
	Id            string       `json:"id"`
	Name          string       `json:"name"`
	HiddenOnUI    bool         `json:"hiddenOnUI"`
	CustomerRefId string       `json:"customerRefId,omitempty"`
	AutoFuel      bool         `json:"autoFuel"`
	Assets        []VaultAsset `json:"assets"`
}

type VaultAsset struct {
	Id                   string `json:"id"`
	Total                string `json:"total"`
	Balance              string `json:"balance"`
	Available            string `json:"available"`
	Pending              string `json:"pending"`
	LockedAmount         string `json:"lockedAmount"`
	TotalStakedCPU       string `json:"totalStakedCPU"`
	SelfStakedCPU        string `json:"selfStakedCPU"`
	SelfStakedNetwork    string `json:"selfStakedNetwork"`
	PendingRefundCPU     string `json:"pendingRefundCPU"`
	PendingRefundNetwork string `json:"pendingRefundNetwork"`
}

type TransactionRequestDestination struct {
	Amount      string `json:"amount"`
	Destination string `json:"destination"`
}

type fireblocksClaims struct {
	Uri      string `json:"uri"`
	Nonce    int64  `json:"nonce"`
	Iat      int64  `json:"iat"`
	Exp      int64  `json:"exp"`
	Sub      string `json:"sub"`
	BodyHash string `json:"bodyHash"`
	jwt.StandardClaims
}

func New(assetId string, baseUrl string, sourceType accounttype.AccountType, sourceId string, apiKey string, signKey *rsa.PrivateKey) *Fireblocks {
	rand.Seed(time.Now().UnixNano())
	return &Fireblocks{
		apiKey:     apiKey,
		assetId:    assetId,
		baseUrl:    baseUrl,
		signKey:    signKey,
		sourceId:   sourceId,
		sourceType: sourceType,
	}
}

func (fb *Fireblocks) ListTransactions() (*[]TransactionDetails, error) {
	resp, err := fb.getRequest("/v1/transactions", url.Values{})
	if err != nil {
		return nil, err
	}

	var result []TransactionDetails
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (fb *Fireblocks) ListVaultAccounts() (*[]VaultAccount, error) {
	resp, err := fb.getRequest("/v1/vault/accounts", url.Values{})
	if err != nil {
		return nil, err
	}

	var result []VaultAccount
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (fb *Fireblocks) CreateNewContractCall(destinationType accounttype.AccountType, destinationId string, destinationTag string, callData string) (*CreateTransactionResponse, error) {
	return fb.CreateNewTransaction(destinationType, destinationId, destinationTag, "0", operationtype.ContractCall, callData)
}

func (fb *Fireblocks) CreateNewTransaction(destinationType accounttype.AccountType, destinationId string, destinationTag string, amount string, operation operationtype.OperationType, callData string) (*CreateTransactionResponse, error) {

	body := &CreateNewTransactionBody{
		AssetId:         fb.assetId,
		Source:          TransferPeerPath{Type: fb.sourceType, Id: fb.sourceId},
		Destination:     *NewDestinationTransferPeerPath(destinationType, destinationId, destinationTag),
		Amount:          amount,
		Operation:       operation,
		ExtraParameters: TransactionExtraParameters{ContractCallData: callData},
	}

	resp, err := fb.postRequest("/v1/transactions", url.Values{}, body)
	if err != nil {
		return nil, err
	}

	var result CreateTransactionResponse
	response, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Response: ", string(response))
	if err != nil {
		return nil, errors.Wrapf(err, "error reading fireblocks create transaction response")
	}
	err = json.NewDecoder(strings.NewReader(string(response))).Decode(&result)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding fireblocks create transaction response")
	}

	return &result, err
}

func (fb *Fireblocks) CancelTransaction(txid string) (string, error) {

	resp, err := fb.postRequest("/v1/transactions/"+txid+"/cancel", url.Values{}, nil)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), err
}

type DropTransactionRequestBody struct {
	FeeLevel     string `json:"feeLevel,omitempty"`
	RequestedFee string `json:"requestedFee,omitempty"`
}

func (fb *Fireblocks) DropTransaction(txid string, feeLevel string, requestedFee string) (string, error) {
	var body DropTransactionRequestBody
	if len(feeLevel) > 0 {
		body.FeeLevel = feeLevel
	}
	if len(requestedFee) > 0 {
		body.RequestedFee = requestedFee
	}

	resp, err := fb.postRequest("/v1/transactions/"+txid+"/cancel", url.Values{}, body)
	if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(data), err
}

func (fb *Fireblocks) getRequest(path string, params url.Values) (*http.Response, error) {
	return fb.sendRequest(http.MethodGet, path, params, "")
}

func (fb *Fireblocks) postRequest(path string, params url.Values, body interface{}) (*http.Response, error) {
	return fb.sendRequest(http.MethodPost, path, params, body)
}

func (fb *Fireblocks) sendRequest(method string, path string, params url.Values, body interface{}) (*http.Response, error) {
	var jsonData []byte
	if body != nil {
		var err error
		jsonData, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	logger.Debug().RawJSON("request", jsonData).Msg("creating new fireblocks transaction")
	for i := 0; i < 3; i++ {
		resp, err := fb.sendRequestImpl(method, path, params, jsonData)
		if err != nil && strings.Contains(err.Error(), "nonce was already used") {
			// Duplicate nonce used, try again
			logger.Warn().Msg("duplicate nonce sent to fireblocks")
			continue
		}

		return resp, err
	}

	return nil, errors.New("too many fireblocks duplicate nonce errors")
}
func (fb *Fireblocks) sendRequestImpl(method string, path string, params url.Values, body []byte) (*http.Response, error) {
	token, err := fb.signJWT(path, body)
	if err != nil {
		return nil, err
	}

	uri, err := url.ParseRequestURI(fb.baseUrl)
	if err != nil {
		return nil, err
	}

	uri.Path = path
	uri.RawQuery = params.Encode()

	client := &http.Client{}
	var req *http.Request
	req, err = http.NewRequest(method, uri.String(), bytes.NewBuffer(body))
	if err != nil {
		logger.
			Error().
			Err(err).
			Str("url", uri.String()).
			Str("method", method).
			Msg("error creating new fireblocks request")
		return nil, err
	}
	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("X-API-Key", fb.apiKey)
	req.Header.Set("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		logger.
			Error().
			Err(err).
			Str("url", fb.baseUrl).
			Msg("error doing fireblocks request")
		return nil, err
	}

	if resp.StatusCode >= 300 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.
				Error().
				Err(err).
				Str("url", uri.String()).
				Str("status", resp.Status).
				Msg("error reading body after sending fireblocks request")
			return nil, fmt.Errorf("status '%s' receiving fireblocks response", resp.Status)
		}

		bodyStr := string(body)
		if strings.Contains(bodyStr, "nonce was already used") {
			logger.
				Warn().
				Str("url", uri.String()).
				Str("status", resp.Status).
				Str("body", bodyStr).
				Msg("error returned when posting fireblocks request")
			return nil, fmt.Errorf("nonce was already used")
		}

		logger.
			Error().
			Str("url", uri.String()).
			Str("status", resp.Status).
			Str("body", bodyStr).
			Msg("error returned when posting fireblocks request")
		return nil, fmt.Errorf("status '%s' fireblocks response", resp.Status)
	}

	return resp, nil
}

func (fb *Fireblocks) signJWT(path string, body []byte) (string, error) {
	newPath := strings.Replace(path, "[", "%5B", -1)
	newPath = strings.Replace(newPath, "]", "%5D", -1)
	now := time.Now().Unix()
	if body == nil {
		body = []byte("null")
	}

	bodyHash := sha256.Sum256([]byte(body))
	fmt.Println("body: (", string(body), ")[", len(body), "], Hash ", hex.EncodeToString(bodyHash[:]))
	claims := fireblocksClaims{
		Uri:      newPath,
		Nonce:    rand.Int63(),
		Iat:      now,
		Exp:      now + 55,
		Sub:      fb.apiKey,
		BodyHash: hex.EncodeToString(bodyHash[:]),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(fb.signKey)
}
