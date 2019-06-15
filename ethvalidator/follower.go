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
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"github.com/offchainlabs/arb-avm/protocol"
	"github.com/offchainlabs/arb-avm/value"
	"github.com/offchainlabs/arb-validator/valmessage"
	errors2 "github.com/pkg/errors"
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
			return nil, errors2.Wrapf(err, "coordinator handshake failed with status %d", resp.StatusCode)
		} else {
			return nil, errors2.Wrap(err, "coordinator handshake failed with empty response")
		}
	}
	tlsCon, ok := coordinatorConn.UnderlyingConn().(*tls.Conn)
	if !ok {
		return nil, errors.New("must connect to coordinator with TLS")
	}
	uniqueVal := tlsCon.ConnectionState().TLSUnique
	hashVal := crypto.Keccak256(uniqueVal)
	sigData, err := crypto.Sign(hashVal, key)
	if err != nil {
		return nil, err
	}
	wr, err := coordinatorConn.NextWriter(websocket.BinaryMessage)
	if err != nil {
		return nil, err
	}
	if _, err := wr.Write(sigData); err != nil {
		return nil, err
	}
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

	if _, ok := c.Validators[address]; !ok {
		return nil, errors.New("coordinator had bad pubkey")
	}

	log.Println("Validator formed connected with coordinator")
	fromCoordinator := make(chan *ValidatorRequest, 128)
	toCoordinator := make(chan *FollowerResponse, 128)
	unanimousRequests := make(map[[32]byte]UnanimousAssertionRequest)
	return &ValidatorFollower{c, coordinatorConn, fromCoordinator, toCoordinator, unanimousRequests}, nil
}

func (m *ValidatorFollower) readPump() {
	defer func() error {
		return m.coordinatorConn.Close()
	}()
	m.coordinatorConn.SetReadLimit(maxMessageSize)
	if err := m.coordinatorConn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		return
	}
	m.coordinatorConn.SetPongHandler(func(string) error {
		return m.coordinatorConn.SetReadDeadline(time.Now().Add(pongWait))
	})
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
	defer func() error {
		ticker.Stop()
		return m.coordinatorConn.Close()
	}()
	for {
		select {
		case message, ok := <-m.ToCoordinator:
			if err := m.coordinatorConn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return
			}
			if !ok {
				// The hub closed the channel.
				// We're already bailing from the channel so failure of CloseMessage can be ignored
				_ = m.coordinatorConn.WriteMessage(websocket.CloseMessage, []byte{})
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
			if _, err := w.Write(raw); err != nil {
				return
			}
			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			if err := m.coordinatorConn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				return
			}
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

	sig, unanHash, err := func() (valmessage.Signature, [32]byte, error) {
		messages := make([]protocol.Message, 0, len(request.SignedMessages))
		for _, signedMsg := range request.SignedMessages {
			msg, err := protocol.NewMessageFromBuf(signedMsg.Message)
			if err != nil {
				return valmessage.Signature{}, [32]byte{}, errors2.Wrap(err, "Follower recieved message in bad format")
			}
			tup, ok := msg.Data.(value.TupleValue)
			if !ok || tup.Len() != 4 {
				return valmessage.Signature{}, [32]byte{}, errors2.Wrap(err, "Follower recieved message in bad format")
			}
			// Access is safe since we already did a length check
			signedVal, _ := tup.GetByInt64(0)
			messageHash := solsha3.SoliditySHA3(
				solsha3.Bytes32(m.VmId),
				solsha3.Bytes32(signedVal.Hash()),
				solsha3.Uint256(msg.Currency),
				msg.TokenType[:],
			)

			signedMsgHash := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(messageHash))
			pubkey, err := crypto.SigToPub(signedMsgHash, signedMsg.Signature)
			if err != nil {
				return valmessage.Signature{}, [32]byte{}, errors2.Wrap(err, "Follower recieved message with bad signature")
			}
			sender := crypto.PubkeyToAddress(*pubkey)
			senderArr := [32]byte{}
			copy(senderArr[12:], sender.Bytes())
			if senderArr != msg.Destination {
				return valmessage.Signature{}, [32]byte{}, errors2.Wrap(err, "Follower recieved message with incorrect signature")
			}
			messages = append(messages, msg)
		}

		resultsChan, unanErrChan := m.Bot.RequestFollowUnanimous(unanRequest, messages)
		var unanUpdate valmessage.UnanimousUpdateResults
		select {
		case unanUpdate = <-resultsChan:
			break
		case err := <-unanErrChan:
			return valmessage.Signature{}, [32]byte{}, errors2.Wrap(err, "Follower failed to follow assertion")
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
			return valmessage.Signature{}, [32]byte{}, errors2.Wrap(err, "Follower failed to generate hash")
		}
		sig, err := m.Sign(unanHash)
		if err != nil {
			return valmessage.Signature{}, [32]byte{}, errors2.Wrap(err, "Follower failed to sign")
		}

		m.unanimousRequests[requestId] = UnanimousAssertionRequest{
			unanRequest,
			messages,
		}
		return sig, unanHash, nil
	}()

	notifyCoordinator := func(msg *UnanimousAssertionFollowerResponse) {
		m.ToCoordinator <- &FollowerResponse{
			RequestId: value.NewHashBuf(unanRequest.Hash()),
			Response:  &FollowerResponse_Unanimous{msg},
		}
	}

	if err != nil {
		log.Println(err)
		notifyCoordinator(&UnanimousAssertionFollowerResponse{
			Accepted: false,
		})
	} else {
		notifyCoordinator(&UnanimousAssertionFollowerResponse{
			Accepted: true,
			Signature: &Signature{
				R: value.NewHashBuf(sig.R),
				S: value.NewHashBuf(sig.S),
				V: uint32(sig.V),
			},
			AssertionHash: value.NewHashBuf(unanHash),
		})
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
								R: value.NewHashFromBuf(sig.R),
								S: value.NewHashFromBuf(sig.S),
								V: uint8(sig.V),
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
