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
	"github.com/rs/zerolog"
	"io"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"github.com/offchainlabs/arbitrum/packages/arb-util/configuration"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks/accounttype"
	"github.com/offchainlabs/arbitrum/packages/arb-util/fireblocks/operationtype"
)

var logger = log.With().Caller().Stack().Str("component", "fireblocks").Logger()

// Transaction status values
const (
	Submitted                     = "SUBMITTED"
	Queued                        = "QEUEUED"
	PendingAuthorization          = "PENDING_AUTHORIZATION"
	PendingSignature              = "PENDING_SIGNATURE"
	Broadcasting                  = "BROADCASTING"
	Pending3rdPartyManualApproval = "PENDING_3RD_PARTY_MANUAL_APPROVAL"
	Pending3rdParty               = "PENDING_3RD_PARTY"
	Confirming                    = "CONFIRMING"
	PartiallyCompleted            = "PARTIALLY_COMPLETED"
	PendingAMLScreening           = "PENDING_AML_SCREENING"
	Cancelled                     = "CANCELLED"
	Rejected                      = "REJECTED"
	Blocked                       = "BLOCKED"
	Failed                        = "FAILED"
)

type Fireblocks struct {
	apiKey            string
	assetId           string
	baseUrl           string
	signKey           *rsa.PrivateKey
	sourceId          string
	sourceType        accounttype.AccountType
	internalWalletIds *map[string]string
	externalWalletIds *map[string]string
}

type StatusBody struct {
	Success bool `json:"success"`
}

type CreateTransactionBody struct {
	AssetId          string                          `json:"assetId"`
	Source           TransferPeerPath                `json:"source"`
	Destination      DestinationTransferPeerPath     `json:"destination"`
	Amount           string                          `json:"amount"`
	Fee              string                          `json:"fee,omitempty"`
	GasPrice         string                          `json:"gasPrice,omitempty"`
	GasLimit         string                          `json:"gasLimit,omitempty"`
	MaxPriorityFee   string                          `json:"maxPriorityFee"`
	MaxTotalGasPrice string                          `json:"maxTotalGasPrice"`
	NetworkFee       string                          `json:"networkFee,omitempty"`
	FeeLevel         string                          `json:"feeLevel,omitempty"`
	MaxFee           string                          `json:"maxFee,omitempty"`
	FailOnLowFee     bool                            `json:"failOnLowFee,omitempty"`
	Note             string                          `json:"note,omitempty"`
	Operation        operationtype.OperationType     `json:"operation,omitempty"`
	CustomerRefId    string                          `json:"customerRefId,omitempty"`
	ReplaceTxByHash  string                          `json:"replaceTxByHash,omitempty"`
	Destinations     []TransactionRequestDestination `json:"destinations,omitempty"`
	ExtraParameters  TransactionExtraParameters      `json:"extraParameters"`
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

func (fb *Fireblocks) NewDestinationTransferUsingAddress(addr string, tag string) *DestinationTransferPeerPath {
	if id, found := (*fb.internalWalletIds)[addr]; found {
		return NewDestinationTransferPeerPath(accounttype.InternalWallet, id, tag)
	}

	if id, found := (*fb.externalWalletIds)[addr]; found {
		return NewDestinationTransferPeerPath(accounttype.ExternalWallet, id, tag)
	}

	return NewDestinationTransferPeerPath(accounttype.OneTimeAddress, addr, tag)
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
	AddressType                   string                     `json:"addressType"`
	AmlScreeningResult            AmlScreeningResult         `json:"amlScreeningResult"`
	Amount                        int64                      `json:"amount"`
	AmountInfo                    AmountInfo                 `json:"amountInfo"`
	AmountUSD                     float32                    `json:"amountUSD"`
	AssetId                       string                     `json:"assetId"`
	CreatedAt                     int64                      `json:"createdAt"`
	CreatedBy                     string                     `json:"createdBy"`
	CustomerRefId                 string                     `json:"customerRefId"`
	Destination                   TransferPeerPathResponse   `json:"destination"`
	DestinationAddress            string                     `json:"destinationAddress"`
	DestinationAddressDescription string                     `json:"destinationAddressDescription"`
	DestinationTag                string                     `json:"destinationTag"`
	Destinations                  []DestinationsResponse     `json:"destinations"`
	ExchangeTxId                  string                     `json:"exchangeTxId"`
	ExtraParameters               TransactionExtraParameters `json:"extraParameters"`
	Fee                           float32                    `json:"fee"`
	FeeCurrency                   string                     `json:"feeCurrency"`
	FeeInfo                       FeeInfo                    `json:"feeInfo"`
	Id                            string                     `json:"id"`
	LastUpdated                   int64                      `json:"lastUpdated"`
	NetAmount                     float32                    `json:"NetAmount"`
	NetworkFee                    float32                    `json:"networkFee"`
	NetworkRecords                []NetworkRecord            `json:"networkRecords"`
	Note                          string                     `json:"note"`
	NumOfConfirmations            int                        `json:"numOfConfirmations"`
	Operation                     string                     `json:"operation"`
	RejectedBy                    string                     `json:"rejectedBy"`
	ReplacedTxHash                string                     `json:"replacedTxHash"`
	RequestedAmount               int64                      `json:"requestedAmount"`
	ServiceFee                    float32                    `json:"serviceFee"`
	SignedBy                      []string                   `json:"signedBy"`
	SignedMessages                []SignedMessage            `json:"signedMessages"`
	Source                        TransferPeerPathResponse   `json:"source"`
	SourceAddress                 string                     `json:"sourceAddress"`
	Status                        string                     `json:"status"`
	SubStatus                     string                     `json:"subStatus"`
	TxHash                        string                     `json:"txHash"`
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
	NetworkFee         string                   `json:"networkFee"`
	AssetId            string                   `json:"assetId"`
	NetAmount          string                   `json:"netAmount"`
	IsDropped          bool                     `json:"isDropped"`
	Status             string                   `json:"status"`
	Type               string                   `json:"type"`
	DestinationAddress string                   `json:"destinationAddress"`
	AmountUSD          string                   `json:"amountUSD"`
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

type NetworkFees struct {
	Low    NetworkFee `json:"low"`
	Medium NetworkFee `json:"medium"`
	High   NetworkFee `json:"high"`
}

type NetworkFee struct {
	BaseFee     string `json:"baseFee"`
	FeePerByte  string `json:"feePerByte"`
	GasPrice    string `json:"gasPrice"`
	NetworkFee  string `json:"networkFee"`
	PriorityFee string `json:"priorityFee"`
}

func New(fireblocksConfig configuration.WalletFireblocks) (*Fireblocks, error) {
	rand.Seed(time.Now().UnixNano())

	var signKey *rsa.PrivateKey
	var err error
	if len(fireblocksConfig.SSLKeyPassword) != 0 {
		signKey, err = jwt.ParseRSAPrivateKeyFromPEMWithPassword([]byte(fireblocksConfig.SSLKey), fireblocksConfig.SSLKeyPassword)
		if err != nil {
			return nil, errors.Wrap(err, "problem with fireblocks privatekey with password")
		}
	} else {
		signKey, err = jwt.ParseRSAPrivateKeyFromPEM([]byte(fireblocksConfig.SSLKey))
		if err != nil {
			return nil, errors.Wrap(err, "problem with fireblocks privatekey")
		}
	}

	sourceType, err := accounttype.New(fireblocksConfig.SourceType)
	if err != nil {
		return nil, errors.Wrap(err, "problem with fireblocks source-type")
	}

	return &Fireblocks{
		apiKey:            fireblocksConfig.APIKey,
		assetId:           fireblocksConfig.AssetId,
		baseUrl:           fireblocksConfig.BaseURL,
		signKey:           signKey,
		sourceId:          fireblocksConfig.SourceId,
		sourceType:        *sourceType,
		internalWalletIds: configuration.UnmarshalMap(fireblocksConfig.InternalWallets),
		externalWalletIds: configuration.UnmarshalMap(fireblocksConfig.ExternalWallets),
	}, nil
}

func (fb *Fireblocks) ListPendingTransactions() (*[]TransactionDetails, error) {
	statusList := []string{
		Submitted,
		Queued,
		PendingAuthorization,
		PendingSignature,
		Broadcasting,
		Pending3rdPartyManualApproval,
		Pending3rdParty,
		Confirming,
		PartiallyCompleted,
		PendingAMLScreening,
	}
	return fb.ListTransactions(statusList)
}

func (fb *Fireblocks) ListTransactions(statusList []string) (*[]TransactionDetails, error) {
	values := url.Values{}
	values.Set("sourceType", fb.sourceType.String())
	values.Set("sourceId", fb.sourceId)
	if len(statusList) > 0 {
		values.Set("status", strings.Join(statusList, ","))
	}
	resp, err := fb.getRequest("/v1/transactions", url.Values{})
	if err != nil {
		return nil, err
	}

	var result []TransactionDetails
	err = fb.parseBody(resp.Body, &result)
	if err != nil {
		return nil, errors.Wrap(err, "list transactions")
	}

	return &result, nil
}

func (fb *Fireblocks) ListVaultAccounts() (*[]VaultAccount, error) {
	resp, err := fb.getRequest("/v1/vault/accounts", url.Values{})
	if err != nil {
		return nil, err
	}

	var result []VaultAccount
	err = fb.parseBody(resp.Body, &result)
	if err != nil {
		return nil, errors.Wrap(err, "list vault accounts")
	}

	return &result, nil
}

func (fb *Fireblocks) EstimateNetworkFees() (*NetworkFees, error) {
	values := url.Values{}
	values.Set("assetId", fb.assetId)
	resp, err := fb.getRequest("/v1/estimate_network_fee", values)
	if err != nil {
		return nil, err
	}

	var result NetworkFees
	err = fb.parseBody(resp.Body, &result)
	if err != nil {
		return nil, errors.Wrap(err, "list vault accounts")
	}

	return &result, nil
}

func (fb *Fireblocks) CreateContractCall(
	destinationType accounttype.AccountType,
	destinationId string,
	destinationTag string,
	amount *big.Int,
	gasLimitWei *big.Int,
	gasPriceWei *big.Int,
	maxPriorityFeeWei *big.Int,
	maxTotalGasPriceWei *big.Int,
	replaceTxByHash string,
	callData string,
) (*CreateTransactionResponse, error) {
	return fb.CreateTransaction(
		destinationType,
		destinationId,
		destinationTag,
		amount,
		operationtype.ContractCall,
		gasLimitWei,
		gasPriceWei,
		maxPriorityFeeWei,
		maxTotalGasPriceWei,
		replaceTxByHash,
		callData,
	)
}

func (fb *Fireblocks) CreateTransaction(
	destinationType accounttype.AccountType,
	destinationId string,
	destinationTag string,
	amountWei *big.Int,
	operation operationtype.OperationType,
	gasLimitWei *big.Int,
	gasPriceWei *big.Int,
	maxPriorityFeeWei *big.Int,
	maxTotalGasPriceWei *big.Int,
	replaceTxByHash string,
	callData string,
) (*CreateTransactionResponse, error) {
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil)
	amountEth := new(big.Rat).SetFrac(amountWei, divisor)
	maxPriorityFee := new(big.Rat).SetFrac(maxPriorityFeeWei, divisor)
	maxTotalGasPrice := new(big.Rat).SetFrac(maxTotalGasPriceWei, divisor)
	gasLimit := new(big.Rat).SetFrac(gasLimitWei, divisor)
	gasPrice := new(big.Rat).SetFrac(gasPriceWei, divisor)

	body := &CreateTransactionBody{
		AssetId:          fb.assetId,
		Source:           TransferPeerPath{Type: fb.sourceType, Id: fb.sourceId},
		Destination:      *NewDestinationTransferPeerPath(destinationType, destinationId, destinationTag),
		Amount:           amountEth.FloatString(18),
		Operation:        operation,
		GasLimit:         gasLimit.FloatString(18),
		GasPrice:         gasPrice.FloatString(18),
		MaxPriorityFee:   maxPriorityFee.FloatString(18),
		MaxTotalGasPrice: maxTotalGasPrice.FloatString(18),
		ReplaceTxByHash:  replaceTxByHash,
		ExtraParameters:  TransactionExtraParameters{ContractCallData: callData},
	}

	resp, err := fb.postRequest("/v1/transactions", url.Values{}, body)
	if err != nil {
		return nil, err
	}

	var result CreateTransactionResponse
	err = fb.parseBody(resp.Body, &result)
	if err != nil {
		return nil, errors.Wrap(err, "create transaction")
	}

	return &result, nil
}

func (fb *Fireblocks) GetTransaction(id string) (*TransactionDetails, error) {
	resp, err := fb.getRequest("/v1/transactions/"+id, url.Values{})
	if err != nil {
		return nil, err
	}

	var result TransactionDetails
	err = fb.parseBody(resp.Body, &result)
	if err != nil {
		return nil, errors.Wrap(err, "get transaction by external id")
	}

	return &result, nil
}

func (fb *Fireblocks) GetTransactionByExternalId(externalId string) (*TransactionDetails, error) {
	resp, err := fb.postRequest("/v1/transactions/external_tx_id/"+externalId, url.Values{}, nil)
	if err != nil {
		return nil, err
	}

	var result TransactionDetails
	err = fb.parseBody(resp.Body, &result)
	if err != nil {
		return nil, errors.Wrap(err, "get transaction by external id")
	}

	return &result, nil
}

func (fb *Fireblocks) IsTransactionStatusFailed(status string) bool {
	if len(status) == 0 || status == Cancelled || status == Rejected || status == Blocked || status == Failed {
		return true
	}

	return false
}

func (fb *Fireblocks) CancelTransaction(txid string) error {

	resp, err := fb.postRequest("/v1/transactions/"+txid+"/cancel", url.Values{}, nil)
	if err != nil {
		return err
	}

	var result StatusBody
	err = fb.parseBody(resp.Body, &result)
	if err != nil {
		return errors.Wrap(err, "cancel transaction")
	}

	if !result.Success {
		return fmt.Errorf("transaction %s not cancelled", txid)
	}

	return nil
}

type DropTransactionRequestBody struct {
	FeeLevel     string `json:"feeLevel,omitempty"`
	RequestedFee string `json:"requestedFee,omitempty"`
}

func (fb *Fireblocks) DropTransaction(txid string, feeLevel string, requestedFee string) error {
	var body DropTransactionRequestBody
	if len(feeLevel) > 0 {
		body.FeeLevel = feeLevel
	}
	if len(requestedFee) > 0 {
		body.RequestedFee = requestedFee
	}

	resp, err := fb.postRequest("/v1/transactions/"+txid+"/cancel", url.Values{}, body)
	if err != nil {
		return err
	}

	var result StatusBody
	err = fb.parseBody(resp.Body, &result)
	if err != nil {
		return errors.Wrap(err, "drop transaction")
	}

	if !result.Success {
		return fmt.Errorf("transaction %s not dropped", txid)
	}

	return nil
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

func (fb *Fireblocks) sendRequestImpl(method string, path string, params url.Values, requestBody []byte) (*http.Response, error) {
	token, err := fb.signJWT(path, requestBody)
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
	req, err = http.NewRequest(method, uri.String(), bytes.NewBuffer(requestBody))
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
			Str("method", method).
			Str("url", fb.baseUrl).
			Msg("error doing fireblocks request")
		return nil, err
	}

	if resp.StatusCode >= 300 {
		responseBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logger.
				Error().
				Err(err).
				Str("method", method).
				Str("url", uri.String()).
				Str("status", resp.Status).
				Msg("error reading body after sending fireblocks request")
			return nil, fmt.Errorf("status '%s' receiving fireblocks response", resp.Status)
		}

		responseBodyStr := string(responseBody)
		if strings.Contains(responseBodyStr, "nonce was already used") {
			logger.
				Warn().
				Str("method", method).
				Str("url", uri.String()).
				Str("status", resp.Status).
				Str("body", responseBodyStr).
				Msg("error returned when posting fireblocks request")
			return nil, fmt.Errorf("nonce was already used")
		}

		if resp.StatusCode == 404 {
			logger.
				Warn().
				Str("method", method).
				Str("url", uri.String()).
				Str("status", resp.Status).
				Str("body", responseBodyStr).
				Msg("fireblocks requested object not found")
			return nil, fmt.Errorf("status '%s' fireblocks requested object not found", resp.Status)
		}

		logger.
			Error().
			Str("method", method).
			Str("url", uri.String()).
			Str("status", resp.Status).
			Str("body", responseBodyStr).
			Msg("error returned when posting fireblocks request")
		return nil, fmt.Errorf("status '%s' fireblocks response", resp.Status)
	}

	logger.
		Debug().
		Str("method", method).
		Str("url", uri.String()).
		RawJSON("body", requestBody).
		Str("status", resp.Status).
		Msg("sent fireblocks request")

	return resp, nil
}

func (fb *Fireblocks) signJWT(path string, body []byte) (string, error) {
	newPath := strings.Replace(path, "[", "%5B", -1)
	newPath = strings.Replace(newPath, "]", "%5D", -1)
	now := time.Now().Unix()
	if body == nil {
		body = []byte("null")
	}

	bodyHash := sha256.Sum256(body)
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

func (fb *Fireblocks) parseBody(body io.Reader, result interface{}) error {
	if zerolog.GlobalLevel() <= zerolog.DebugLevel {
		response, err := ioutil.ReadAll(body)
		if err != nil {
			return errors.Wrapf(err, "error reading fireblocks response")
		}
		logger.Debug().RawJSON("body", response).Msgf("received fireblocks response")
		err = json.NewDecoder(strings.NewReader(string(response))).Decode(&result)
		if err != nil {
			return errors.Wrapf(err, "error decoding fireblocks response")
		}
	} else {
		err := json.NewDecoder(body).Decode(&result)
		if err != nil {
			return errors.Wrapf(err, "error decoding fireblocks response")
		}
	}

	return nil
}
