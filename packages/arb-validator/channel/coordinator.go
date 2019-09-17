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

package channel

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"math"
	"net/http"
	"time"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethvalidator"

	"github.com/offchainlabs/arbitrum/packages/arb-validator/validator"

	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	errors2 "github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ValidatorLeaderRequest interface {
}

type LabeledFollowerResponse struct {
	address  common.Address
	response *valmessage.FollowerResponse
}

type ClientManager struct {
	clients         map[*Client]bool
	broadcast       chan *valmessage.ValidatorRequest
	register        chan *Client
	unregister      chan *Client
	waitRequestChan chan chan bool
	sigRequestChan  chan GatherSignatureRequest
	waitingChans    map[chan bool]bool
	responses       map[[32]byte]chan LabeledFollowerResponse

	val        *ethvalidator.Validator
	vmID       common.Address
	validators map[common.Address]validatorInfo
}

func NewClientManager(val *ethvalidator.Validator, vmID common.Address, validators map[common.Address]validatorInfo) *ClientManager {
	return &ClientManager{
		clients:         make(map[*Client]bool),
		broadcast:       make(chan *valmessage.ValidatorRequest, 10),
		register:        make(chan *Client, 10),
		unregister:      make(chan *Client, 10),
		waitRequestChan: make(chan chan bool, 128),
		sigRequestChan:  make(chan GatherSignatureRequest, 10),
		waitingChans:    make(map[chan bool]bool),
		responses:       make(map[[32]byte]chan LabeledFollowerResponse),
		val:             val,
		vmID:            vmID,
		validators:      validators,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (m *ClientManager) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := func() (*Client, error) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return nil, err
		}
		tlsCon, ok := conn.UnderlyingConn().(*tls.Conn)
		if !ok {
			return nil, errors.New("made non tls connection")
		}

		_, signedUnique, err := conn.ReadMessage()
		if err != nil {
			return nil, errors2.Wrap(err, "failed to get message from follower")
		}
		uniqueVal := tlsCon.ConnectionState().TLSUnique
		hashVal := crypto.Keccak256(uniqueVal)
		pubkey, err := ethvalidator.EthSigToPub(hashVal, signedUnique)
		if err != nil {
			return nil, err
		}
		address := crypto.PubkeyToAddress(*pubkey)
		if _, ok := m.validators[address]; !ok {
			return nil, errors.New("follower tried to connect with bad pubkey")
		}
		sigData, err := m.val.Sign(hashVal)
		if err != nil {
			return nil, err
		}
		wr, err := conn.NextWriter(websocket.BinaryMessage)
		if err != nil {
			return nil, err
		}
		if _, err := wr.Write(m.vmID[:]); err != nil {
			return nil, err
		}

		if _, err := wr.Write(sigData); err != nil {
			return nil, err
		}

		if err := wr.Close(); err != nil {
			return nil, err
		}
		return NewClient(conn, address), nil
	}()
	if err != nil {
		log.Printf("Coordinator failed to connet with follower: %v\n", err)
		return
	}

	log.Println("Coordinator connected with follower", hexutil.Encode(c.Address[:]))
	m.register <- c

	go func() {
		if err := c.run(); err != nil {
			log.Printf("Coordinator lost connection to client with error: %v\n", err)
		}
		m.unregister <- c
	}()
}

func (m *ClientManager) RunServer(ctx context.Context) error {
	srv := &http.Server{
		Addr:    ":1236",
		Handler: m,
	}
	errChan := make(chan error)
	go func() {
		errChan <- srv.ListenAndServeTLS("server.crt", "server.key")
	}()
	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Second*2)
		defer cancel()
		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			return err
		}
		return errors.New("Stop requested")
	case err := <-errChan:
		return err
	}
}

type GatherSignatureRequest struct {
	request      *valmessage.ValidatorRequest
	responseChan chan LabeledFollowerResponse
	requestID    [32]byte
}

func (m *ClientManager) Run() {
	aggResponseChan := make(chan LabeledFollowerResponse, 32)
	for {
		select {
		case waitRequest := <-m.waitRequestChan:
			if len(m.clients) == len(m.validators)-1 {
				waitRequest <- true
			} else {
				m.waitingChans[waitRequest] = true
			}
		case response := <-aggResponseChan:
			m.responses[value.NewHashFromBuf(response.response.RequestId)] <- response
		case request := <-m.sigRequestChan:
			m.broadcast <- request.request
			m.responses[request.requestID] = request.responseChan
		case client := <-m.register:
			m.clients[client] = true
			go func() {
				for message := range client.FromClient {
					response := new(valmessage.FollowerResponse)
					err := proto.Unmarshal(message, response)
					if err != nil {
						log.Println("Recieved bad message from follower")
						continue
					}
					aggResponseChan <- LabeledFollowerResponse{client.Address, response}
				}
			}()
			if len(m.clients) == len(m.validators)-1 {
				for waitChan := range m.waitingChans {
					waitChan <- true
				}
				m.waitingChans = make(map[chan bool]bool)
			}
		case client := <-m.unregister:
			if _, ok := m.clients[client]; ok {
				delete(m.clients, client)
				close(client.ToClient)
			}
		case message := <-m.broadcast:
			raw, err := proto.Marshal(message)
			if err != nil {
				continue
			}
			for client := range m.clients {
				select {
				case client.ToClient <- raw:
				default:
					close(client.ToClient)
					delete(m.clients, client)
				}
			}
		}
	}
}

func (m *ClientManager) gatherSignatures(
	ctx context.Context,
	request *valmessage.ValidatorRequest,
	requestID [32]byte,
) []LabeledFollowerResponse {
	responseChan := make(chan LabeledFollowerResponse, len(m.validators)-1)
	m.sigRequestChan <- GatherSignatureRequest{
		request,
		responseChan,
		requestID,
	}
	responseList := make([]LabeledFollowerResponse, 0, len(m.validators)-1)
Loop:
	for {
		select {
		case response := <-responseChan:
			responseList = append(responseList, response)
		case <-ctx.Done():
			log.Println("Coordinator cancelled gathering signatures")
			break Loop
		}
		if len(responseList) == len(m.validators)-1 {
			break
		}
	}
	return responseList
}

func (m *ClientManager) WaitForFollowers(ctx context.Context) bool {
	waitChan := make(chan bool, 1)
	m.waitRequestChan <- waitChan
	select {
	case <-waitChan:
		return true
	case <-ctx.Done():
		return false
	}
}

type OffchainMessage struct {
	Message   protocol.Message
	Hash      []byte
	Signature []byte
}

type MessageProcessingQueue struct {
	queuedMessages []OffchainMessage
	actions        chan func(*MessageProcessingQueue)
}

func NewMessageProcessingQueue() *MessageProcessingQueue {
	return &MessageProcessingQueue{
		queuedMessages: make([]OffchainMessage, 0),
		actions:        make(chan func(*MessageProcessingQueue), 10),
	}
}

func (m *MessageProcessingQueue) Fetch() chan []OffchainMessage {
	retChan := make(chan []OffchainMessage, 1)
	m.actions <- func(m *MessageProcessingQueue) {
		retChan <- m.queuedMessages
		m.queuedMessages = nil
	}
	return retChan
}

func (m *MessageProcessingQueue) HasMessages() chan bool {
	retChan := make(chan bool, 1)
	m.actions <- func(m *MessageProcessingQueue) {
		retChan <- len(m.queuedMessages) > 0
	}
	return retChan
}

func (m *MessageProcessingQueue) Return(messages []OffchainMessage) {
	m.actions <- func(m *MessageProcessingQueue) {
		m.queuedMessages = append(messages, m.queuedMessages...)
	}
}

func (m *MessageProcessingQueue) Send(message OffchainMessage) {
	m.actions <- func(m *MessageProcessingQueue) {
		m.queuedMessages = append(m.queuedMessages, message)
	}
}

func (m *MessageProcessingQueue) run() {
	go func() {
		for action := range m.actions {
			action(m)
		}
	}()
}

type ValidatorCoordinator struct {
	Val        *Validator
	ChannelVal *validator.ChannelValidator
	cm         *ClientManager

	actions chan func(*ValidatorCoordinator)

	mpq               *MessageProcessingQueue
	maxStepsUnanSteps int32
}

func NewCoordinator(
	name string,
	val *ethvalidator.Validator,
	vmID common.Address,
	machine machine.Machine,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
	maxCallSteps int32,
	maxStepsUnanSteps int32,
) (*ValidatorCoordinator, error) {
	header, err := val.LatestHeader(context.Background())
	if err != nil {
		return nil, errors2.Wrap(err, "Validator couldn't get latest error")
	}

	c, err := NewValidator(val, vmID, machine, config)
	if err != nil {
		return nil, errors2.Wrap(err, "Error initializing Validator in coordinator")
	}

	channelVal := validator.NewChannelValidator(
		name,
		c,
		vmID,
		header,
		protocol.NewBalanceTracker(),
		config,
		machine,
		challengeEverything,
		maxCallSteps,
	)

	return &ValidatorCoordinator{
		Val:               c,
		ChannelVal:        channelVal,
		cm:                NewClientManager(val, vmID, c.Validators),
		actions:           make(chan func(*ValidatorCoordinator), 10),
		mpq:               NewMessageProcessingQueue(),
		maxStepsUnanSteps: maxStepsUnanSteps,
	}, nil
}

func (m *ValidatorCoordinator) SendMessage(msg OffchainMessage) {
	m.mpq.Send(msg)
}

func (m *ValidatorCoordinator) StartServer(ctx context.Context) {
	go func() {
		err := m.cm.RunServer(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}()
	go m.mpq.run()
	go m.cm.Run()
}

func (m *ValidatorCoordinator) Run(ctx context.Context) error {
	parsedChan, err := m.Val.StartListening(ctx)
	if err != nil {
		return err
	}

	go func() {
		m.ChannelVal.Run(ctx, parsedChan)
	}()

	go func() {
		for {
			select {
			case action := <-m.actions:
				action(m)
			case <-time.After(time.Second):
				if !<-m.ChannelVal.CanRun() {
					break
				}
				shouldUnan := false
				forceFinal := false
				pendingCount := <-m.ChannelVal.PendingMessageCount()
				if pendingCount > 0 {
					// Force onchain assertion if there are pending on chain messages, then force an offchain assertion
					log.Println("Unanimous asserting because of pending on chain messages")
					shouldUnan = true
					forceFinal = true
				} else if <-m.mpq.HasMessages() {
					log.Println("Unanimous asserting because of pending off chain messages")
					shouldUnan = true
				} else if <-m.ChannelVal.CanContinueRunning() {
					log.Println("Unanimous asserting because machine is not blocked")
					shouldUnan = true
				}
				if !shouldUnan {
					break
				}

				ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
				err := m.initiateUnanimousAssertionImpl(ctx, forceFinal, m.maxStepsUnanSteps)
				cancel()
				if err == nil {
					// Assertion was successful so we are done
					break
				}
				log.Println("Coordinator hit problem unanimously asserting", err)
				if <-m.ChannelVal.HasOpenAssertion() {
					log.Println("Coordinator is closing channel")
					closedChan, errChan := m.ChannelVal.CloseUnanimousAssertionRequest()
					select {
					case _ = <-closedChan:
						log.Println("Coordinator successfully closed channel")
					case err := <-errChan:
						log.Println("Coordinator failed to close channel", err)
					}
				} else {
					log.Println("Coordinator is creating a disputable assertion")
					// Get the message on-chain (in the inbox)
					// Do the disputable assertion
					messages := <-m.mpq.Fetch()
					for _, msg := range messages {
						receiptChan, errChan := m.Val.ForwardMessage(context.Background(), msg.Message.Data, msg.Message.TokenType, msg.Message.Currency, msg.Signature)
						select {
						case _ = <-receiptChan:
						case err := <-errChan:
							log.Fatalln("ForwardMessage err", err)
						}
					}
					m.initiateDisputableAssertionImpl()
				}
			}
		}
	}()
	return nil
}

type CoordinatorUnanimousRequest struct {
	final   bool
	retChan chan bool
	errChan chan error
}

func (m *ValidatorCoordinator) InitiateDisputableAssertion() chan bool {
	retChan := make(chan bool, 1)
	m.actions <- func(m *ValidatorCoordinator) {
		retChan <- m.initiateDisputableAssertionImpl()
	}

	return retChan
}

func (m *ValidatorCoordinator) InitiateUnanimousAssertion(final bool) (chan bool, chan error) {
	retChan := make(chan bool, 1)
	errChan := make(chan error, 1)
	m.actions <- func(m *ValidatorCoordinator) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
		err := m.initiateUnanimousAssertionImpl(ctx, final, m.maxStepsUnanSteps)
		cancel()
		if err != nil {
			errChan <- err
		} else {
			retChan <- true
		}
	}
	return retChan, errChan
}

func (m *ValidatorCoordinator) initiateDisputableAssertionImpl() bool {
	start := time.Now()
	resultChan, errChan := m.ChannelVal.RequestDisputableAssertion(10000)

	select {
	case <-resultChan:
		log.Printf("Coordinator made disputable assertion in %s seconds", time.Since(start))
		return true
	case err := <-errChan:
		log.Println("Disputable assertion failed", err)
		return false
	}
}

func (m *ValidatorCoordinator) initiateUnanimousAssertionImpl(ctx context.Context, forceFinal bool, maxSteps int32) error {
	if !m.cm.WaitForFollowers(ctx) {
		return errors.New("failed to connect to other validators when unanimous asserting")
	}

	queuedMessages := <-m.mpq.Fetch()

	err := func() error {
		log.Println("Coordinator making unanimous assertion with", len(queuedMessages), "messages")
		newMessages := make([]protocol.Message, 0, len(queuedMessages))
		messageHashes := make([][]byte, 0, len(newMessages))
		for _, msg := range queuedMessages {
			newMessages = append(newMessages, msg.Message)
			messageHashes = append(messageHashes, msg.Hash)
		}

		// Force onchain assertion if there are outgoing messages
		shouldFinalize := func(a *protocol.Assertion) bool {
			return len(a.OutMsgs) > 0
		}

		start := time.Now()
		requestChan, resultsChan, unanErrChan := m.ChannelVal.InitiateUnanimousRequest(10000, newMessages, messageHashes, forceFinal, maxSteps, shouldFinalize)
		responsesChan := make(chan []LabeledFollowerResponse, 1)

		var unanRequest valmessage.UnanimousRequest
		select {
		case unanRequest = <-requestChan:
			break
		case err := <-unanErrChan:
			return err
		}

		requestMessages := make([]*valmessage.SignedMessage, 0, len(unanRequest.NewMessages))
		for i, msg := range unanRequest.NewMessages {
			requestMessages = append(requestMessages, &valmessage.SignedMessage{
				Message:   protocol.NewMessageBuf(msg),
				Signature: queuedMessages[i].Signature,
			})
		}
		hashID := unanRequest.Hash()

		notifyFollowers := func(msg *valmessage.UnanimousAssertionValidatorNotification) {
			m.cm.broadcast <- &valmessage.ValidatorRequest{
				RequestId: value.NewHashBuf(hashID),
				Request:   &valmessage.ValidatorRequest_UnanimousNotification{UnanimousNotification: msg},
			}
		}

		go func() {
			request := &valmessage.UnanimousAssertionValidatorRequest{
				BeforeHash:     value.NewHashBuf(unanRequest.BeforeHash),
				BeforeInbox:    value.NewHashBuf(unanRequest.BeforeInbox),
				SequenceNum:    unanRequest.SequenceNum,
				TimeBounds:     protocol.NewTimeBoundsBuf(unanRequest.TimeBounds),
				SignedMessages: requestMessages,
			}
			responsesChan <- m.cm.gatherSignatures(
				ctx,
				&valmessage.ValidatorRequest{
					RequestId: value.NewHashBuf(hashID),
					Request: &valmessage.ValidatorRequest_Unanimous{
						Unanimous: request,
					},
				},
				hashID,
			)
		}()

		var unanUpdate valmessage.UnanimousUpdateResults
		select {
		case unanUpdate = <-resultsChan:
			break
		case err := <-unanErrChan:
			notifyFollowers(&valmessage.UnanimousAssertionValidatorNotification{
				Accepted: false,
			})
			return err
		}

		unanHash, err := m.Val.UnanimousAssertHash(
			unanUpdate.SequenceNum,
			unanUpdate.BeforeHash,
			unanUpdate.NewInboxHash,
			unanUpdate.BeforeInbox,
			unanUpdate.Assertion,
		)
		if err != nil {
			log.Println("Coordinator failed to hash unanimous assertion")
			notifyFollowers(&valmessage.UnanimousAssertionValidatorNotification{
				Accepted: false,
			})
			return err
		}
		sig, err := m.Val.Sign(unanHash)
		if err != nil {
			log.Println("Coordinator failed to sign unanimous assertion")
			notifyFollowers(&valmessage.UnanimousAssertionValidatorNotification{
				Accepted: false,
			})
			return err
		}

		responses := <-responsesChan
		if len(responses) != m.Val.ValidatorCount()-1 {
			log.Println("Coordinator failed to collect unanimous assertion sigs")
			notifyFollowers(&valmessage.UnanimousAssertionValidatorNotification{
				Accepted: false,
			})
			return errors.New("some Validators didn't respond")
		}

		signatures := make([][]byte, m.Val.ValidatorCount())
		signatures[m.Val.Validators[m.Val.Address()].indexNum] = sig
		for _, response := range responses {
			r := response.response.Unanimous
			if !r.Accepted {
				notifyFollowers(&valmessage.UnanimousAssertionValidatorNotification{
					Accepted: false,
				})
				return errors.New("some Validators refused to sign")
			}
			if value.NewHashFromBuf(r.AssertionHash) != unanHash {
				notifyFollowers(&valmessage.UnanimousAssertionValidatorNotification{
					Accepted: false,
				})
				return errors.New("some Validators signed the wrong assertion")
			}
			signatures[m.Val.Validators[response.address].indexNum] = r.Signature
		}

		elapsed := time.Since(start)
		log.Printf("Coordinator succeeded signing unanimous assertion in %s\n", elapsed)
		notifyFollowers(&valmessage.UnanimousAssertionValidatorNotification{
			Accepted:   true,
			Signatures: signatures,
		})

		confRetChan, confErrChan := m.ChannelVal.ConfirmOffchainUnanimousAssertion(
			unanUpdate.UnanimousRequestData,
			signatures,
			true,
		)

		wasFinal := unanUpdate.SequenceNum == math.MaxUint64

		if wasFinal {
			log.Println("Coordinator is closing unanimous assertion")
		}

		select {
		case <-confRetChan:
			if wasFinal {
				log.Println("Coordinator successfully closed channel")
			} else {
				log.Println("Coordinator is keeping unanimous assertion chain open")
			}
		case err := <-confErrChan:
			log.Println("Coordinator failed to complete assertion", err)
			return err
		}
		return nil
	}()

	if err != nil {
		m.mpq.Return(queuedMessages)
		return err
	}
	return nil
}
