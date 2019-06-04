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

package ethvalidator

import (
	"crypto/ecdsa"
	"crypto/tls"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-validator/valmessage"
	"log"
	"math"
	"time"

	"github.com/offchainlabs/arb-avm/vm"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 8192
)

type UnanimousAssertionRequest struct {
	requestData valmessage.UnanimousRequestData
	newMessages []protocol.Message
}

type ValidatorFollower struct {
	*EthValidator

	coordinatorConn *websocket.Conn
	FromCoordinator chan *ValidatorRequest
	ToCoordinator   chan *FollowerResponse

	unanimousRequests map[[32]byte]UnanimousAssertionRequest
}

func NewValidatorFollower(
	name string,
	machine *vm.Machine,
	key *ecdsa.PrivateKey,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
	connectionInfo ArbAddresses,
	ethURL string,
	coordinatorURL string,
) (*ValidatorFollower, error) {

	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	coordinatorConn, resp, err := dialer.Dial(coordinatorURL, nil)
	if err != nil {
		if resp != nil {
			return nil, fmt.Errorf("Coordinator handshake failed with error %v and status %d", err, resp.StatusCode)
		} else {
			return nil, fmt.Errorf("Coordinator handshake failed with error %v and response <nil>", err)
		}
	}
	tlsCon, ok := coordinatorConn.UnderlyingConn().(*tls.Conn)
	if !ok {
		return nil, errors.New("Must connect to coordinator with TLS")
	}
	uniqueVal := tlsCon.ConnectionState().TLSUnique
	hashVal := crypto.Keccak256(uniqueVal)
	sigData, err := crypto.Sign(hashVal, key)
	if err != nil {
		return nil, err
	}
	wr, err := coordinatorConn.NextWriter(websocket.BinaryMessage)
	wr.Write(sigData)
	if err := wr.Close(); err != nil {
		return nil, err
	}
	_, vmData, err := coordinatorConn.ReadMessage()
	if err != nil {
		return nil, err
	}
	var vmId [32]byte
	copy(vmId[:], vmData)
	pubkey, err := crypto.SigToPub(hashVal, vmData[32:])
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(*pubkey)

	c, err := NewEthValidator(name, vmId, machine, key, config, challengeEverything, connectionInfo, ethURL)
	if err != nil {
		return nil, err
	}

	if _, ok := c.Validators[address] ; !ok {
		return nil, errors.New("Coordinator had bad pubkey")
	}

	log.Println("Validator formed connected with coordinator")
	fromCoordinator := make(chan *ValidatorRequest, 128)
	toCoordinator := make(chan *FollowerResponse, 128)
	unanimousRequests := make(map[[32]byte]UnanimousAssertionRequest)
	return &ValidatorFollower{c, coordinatorConn, fromCoordinator, toCoordinator, unanimousRequests}, nil
}

func (m *ValidatorFollower) readPump() {
	defer func() {
		m.coordinatorConn.Close()
	}()
	m.coordinatorConn.SetReadLimit(maxMessageSize)
	m.coordinatorConn.SetReadDeadline(time.Now().Add(pongWait))
	m.coordinatorConn.SetPongHandler(func(string) error { m.coordinatorConn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := m.coordinatorConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		req := &ValidatorRequest{}
		err = proto.Unmarshal(message, req)
		if err != nil {
			log.Printf("Validator recieved malformed message")
			continue
		}
		m.FromCoordinator <- req
	}
}

func (m *ValidatorFollower) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		m.coordinatorConn.Close()
	}()
	for {
		select {
		case message, ok := <-m.ToCoordinator:
			m.coordinatorConn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				m.coordinatorConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := m.coordinatorConn.NextWriter(websocket.BinaryMessage)
			if err != nil {
				return
			}

			raw, err := proto.Marshal(message)
			if err != nil {
				log.Fatalln("Follower failed to marshal response")
			}
			w.Write(raw)
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			m.coordinatorConn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := m.coordinatorConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (m *ValidatorFollower) HandleUnanimousRequest(
	request *UnanimousAssertionValidatorRequest,
	requestId [32]byte,
) {
	unanRequest := valmessage.UnanimousRequestData{
		BeforeHash:  value.NewHashFromBuf(request.BeforeHash),
		BeforeInbox: value.NewHashFromBuf(request.BeforeInbox),
		SequenceNum: request.SequenceNum,
		TimeBounds:  protocol.NewTimeBoundsFromBuf(request.TimeBounds),
	}

	notifyCoordinator := func(msg *UnanimousAssertionFollowerResponse) {
		m.ToCoordinator <- &FollowerResponse{
			RequestId: value.NewHashBuf(unanRequest.Hash()),
			Response:  &FollowerResponse_Unanimous{msg},
		}
	}

	notifyFailed := func(err error) {
		log.Println(err)
		notifyCoordinator(&UnanimousAssertionFollowerResponse{
			Accepted: false,
		})
	}

	messages := make([]protocol.Message, 0, len(request.SignedMessages))
	for _, signedMsg := range request.SignedMessages {
		msg, err := protocol.NewMessageFromBuf(signedMsg.Message)
		tup, ok := msg.Data.(value.TupleValue)
		if !ok || tup.Len() != 4 {
			notifyFailed(fmt.Errorf("Follower recieved message in bad format"))
			return
		}
		signedVal, _ := tup.GetByInt64(0)
		messageHash := solsha3.SoliditySHA3(
			solsha3.Bytes32(m.VmId),
			solsha3.Bytes32(signedVal.Hash()),
			solsha3.Uint256(msg.Currency),
			msg.TokenType[:],
		)

		signedMsgHash := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(messageHash))
		pubkey, err := crypto.SigToPub(signedMsgHash, signedMsg.Signature)
		sender := crypto.PubkeyToAddress(*pubkey)
		senderArr := [32]byte{}
		copy(senderArr[12:], sender.Bytes())
		if senderArr != msg.Destination {
			notifyFailed(fmt.Errorf("Follower recieved message with incorrect signature"))
			return
		}
		if err != nil {
			notifyFailed(fmt.Errorf("Follower failed to unmarshal valmessage: %v", err))
			return
		}
		messages = append(messages, msg)
	}

	resultsChan, unanErrChan := m.Bot.RequestFollowUnanimous(unanRequest, messages)
	var unanUpdate valmessage.UnanimousUpdateResults
	select {
	case unanUpdate = <-resultsChan:
		break
	case err := <-unanErrChan:
		notifyFailed(fmt.Errorf("Follower failed to follow assertion error: %v", err))
		return
	}

	// Force onchain assertion if there are outgoing messages
	if len(unanUpdate.Assertion.OutMsgs) > 0 {
		unanUpdate.SequenceNum = math.MaxUint64
	}

	unanHash, err := m.UnanimousAssertHash(
		unanUpdate.SequenceNum,
		unanUpdate.BeforeHash,
		unanUpdate.TimeBounds,
		unanUpdate.NewInboxHash,
		unanUpdate.OriginalInboxHash,
		unanUpdate.Assertion,
	)
	if err != nil {
		notifyFailed(fmt.Errorf("Follower failed to generate hash: %v", err))
		return
	}
	sig, err := m.Sign(unanHash)
	if err != nil {
		notifyFailed(fmt.Errorf("Follower failed to sign: %v", err))
		return
	}

	notifyCoordinator(&UnanimousAssertionFollowerResponse{
		Accepted: true,
		Signature: &Signature{
			R: value.NewHashBuf(sig.R),
			S: value.NewHashBuf(sig.S),
			V: uint32(sig.V),
		},
		AssertionHash: value.NewHashBuf(unanHash),
	})
	m.unanimousRequests[requestId] = UnanimousAssertionRequest{
		unanRequest,
		messages,
	}
}

func (m *ValidatorFollower) HandleCreateVM(request *CreateVMValidatorRequest) {
	createHash := CreateVMHash(request)
	sig, err := m.Sign(createHash)
	if err != nil {
		log.Printf("Follower failed to sign1: %v", err)
		m.ToCoordinator <- &FollowerResponse{
			Response: &FollowerResponse_Create{
				&CreateVMFollowerResponse{
					Accepted: false,
				},
			},
			RequestId: value.NewHashBuf(createHash),
		}
	}

	m.ToCoordinator <- &FollowerResponse{
		Response: &FollowerResponse_Create{
			&CreateVMFollowerResponse{
				Accepted: true,
				Signature: &Signature{
					R: value.NewHashBuf(sig.R),
					S: value.NewHashBuf(sig.S),
					V: uint32(sig.V),
				},
			},
		},
		RequestId: value.NewHashBuf(createHash),
	}
}

func (m *ValidatorFollower) Run() error {
	go m.readPump()
	go m.writePump()

	go func() {
		for {
			select {
			case req := <-m.FromCoordinator:
				switch request := req.Request.(type) {
				case *ValidatorRequest_Unanimous:
					m.HandleUnanimousRequest(request.Unanimous, value.NewHashFromBuf(req.RequestId))
				case *ValidatorRequest_UnanimousNotification:
					requestInfo := m.unanimousRequests[value.NewHashFromBuf(req.RequestId)]
					if request.UnanimousNotification.Accepted {
						sigs := make([]valmessage.Signature, len(request.UnanimousNotification.Signatures))
						for _, sig := range request.UnanimousNotification.Signatures {
							sigs = append(sigs, valmessage.Signature{
								value.NewHashFromBuf(sig.R),
								value.NewHashFromBuf(sig.S),
								uint8(sig.V),
							})
						}
						_, _ = m.Bot.ConfirmOffchainUnanimousAssertion(
							requestInfo.requestData,
							sigs,
						)
					}
				case *ValidatorRequest_Create:
					m.HandleCreateVM(request.Create)
				case *ValidatorRequest_CreateNotification:
				}
			}
		}
	}()
	return m.StartListening()
}
