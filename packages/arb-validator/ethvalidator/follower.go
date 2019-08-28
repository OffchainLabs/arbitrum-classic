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
	"context"
	"crypto/ecdsa"
	"crypto/tls"
	"errors"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	errors2 "github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/machine"
	"github.com/offchainlabs/arbitrum/packages/arb-util/protocol"
	"github.com/offchainlabs/arbitrum/packages/arb-util/value"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/ethbridge"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/hashing"
	"github.com/offchainlabs/arbitrum/packages/arb-validator/valmessage"
)

type ValidatorFollower struct {
	*EthValidator

	client *Client

	unanimousRequests map[[32]byte]valmessage.UnanimousRequestData
	maxStepsUnanSteps int32
}

func NewValidatorFollower(
	name string,
	machine machine.Machine,
	key *ecdsa.PrivateKey,
	config *valmessage.VMConfiguration,
	challengeEverything bool,
	maxCallSteps int32,
	connectionInfo ethbridge.ArbAddresses,
	ethURL string,
	coordinatorURL string,
	maxStepsUnanSteps int32,
) (*ValidatorFollower, error) {
	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = &tls.Config{
		InsecureSkipVerify: true,
	}

	coordinatorConn, resp, err := dialer.Dial(coordinatorURL, nil)
	if err != nil {
		if resp != nil {
			return nil, errors2.Wrapf(err, "coordinator handshake failed with status %d", resp.StatusCode)
		}
		return nil, errors2.Wrap(err, "coordinator handshake failed with empty response")
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
	var vmID [32]byte
	copy(vmID[:], vmData)
	pubkey, err := crypto.SigToPub(hashVal, vmData[32:])
	if err != nil {
		return nil, err
	}
	address := crypto.PubkeyToAddress(*pubkey)

	c, err := NewEthValidator(name, vmID, machine, key, config, challengeEverything, maxCallSteps, connectionInfo, ethURL)
	if err != nil {
		return nil, err
	}

	if _, ok := c.Validators[address]; !ok {
		return nil, errors.New("coordinator had bad pubkey")
	}

	log.Println("Validator formed connection with coordinator")
	unanimousRequests := make(map[[32]byte]valmessage.UnanimousRequestData)
	client := NewClient(coordinatorConn, address)
	return &ValidatorFollower{
		EthValidator:      c,
		client:            client,
		unanimousRequests: unanimousRequests,
		maxStepsUnanSteps: maxStepsUnanSteps,
	}, nil
}

func (m *ValidatorFollower) HandleUnanimousRequest(
	request *valmessage.UnanimousAssertionValidatorRequest,
	requestID [32]byte,
) error {
	sig, unanHash, err := func() ([]byte, [32]byte, error) {
		messages := make([]protocol.Message, 0, len(request.SignedMessages))
		for _, signedMsg := range request.SignedMessages {
			msg, err := protocol.NewMessageFromBuf(signedMsg.Message)
			if err != nil {
				return nil, [32]byte{}, errors2.Wrap(err, "Follower recieved message in bad format")
			}
			tup, ok := msg.Data.(value.TupleValue)
			if !ok || tup.Len() != 4 {
				return nil, [32]byte{}, errors2.Wrap(err, "Follower recieved message in bad format")
			}
			// Access is safe since we already did a length check
			signedVal, _ := tup.GetByInt64(0)
			messageHash := solsha3.SoliditySHA3(
				solsha3.Bytes32(m.VMID),
				solsha3.Bytes32(signedVal.Hash()),
				solsha3.Uint256(msg.Currency),
				msg.TokenType[:],
			)

			signedMsgHash := solsha3.SoliditySHA3WithPrefix(solsha3.Bytes32(messageHash))
			pubkey, err := crypto.SigToPub(signedMsgHash, signedMsg.Signature)
			if err != nil {
				return nil, [32]byte{}, errors2.Wrap(err, "Follower recieved message with bad signature")
			}
			sender := crypto.PubkeyToAddress(*pubkey)
			senderArr := [32]byte{}
			copy(senderArr[12:], sender.Bytes())
			if senderArr != msg.Destination {
				return nil, [32]byte{}, errors2.Wrap(err, "Follower recieved message with incorrect signature")
			}
			messages = append(messages, msg)
		}

		// Force onchain assertion if there are outgoing messages
		shouldFinalize := func(a *protocol.Assertion) bool {
			return len(a.OutMsgs) > 0
		}

		resultsChan, unanErrChan := m.Bot.RequestFollowUnanimous(
			valmessage.UnanimousRequestData{
				BeforeHash:  value.NewHashFromBuf(request.BeforeHash),
				BeforeInbox: value.NewHashFromBuf(request.BeforeInbox),
				SequenceNum: request.SequenceNum,
				TimeBounds:  protocol.NewTimeBoundsFromBuf(request.TimeBounds),
			},
			messages,
			m.maxStepsUnanSteps,
			shouldFinalize,
		)
		var unanUpdate valmessage.UnanimousUpdateResults
		select {
		case unanUpdate = <-resultsChan:
			break
		case err := <-unanErrChan:
			return nil, [32]byte{}, errors2.Wrap(err, "Follower failed to follow assertion")
		}

		unanHash, err := m.UnanimousAssertHash(
			unanUpdate.SequenceNum,
			unanUpdate.BeforeHash,
			unanUpdate.NewInboxHash,
			unanUpdate.BeforeInbox,
			unanUpdate.Assertion,
		)
		if err != nil {
			return nil, [32]byte{}, errors2.Wrap(err, "Follower failed to generate hash")
		}
		sig, err := m.Sign(unanHash)
		if err != nil {
			return nil, [32]byte{}, errors2.Wrap(err, "Follower failed to sign")
		}

		m.unanimousRequests[requestID] = unanUpdate.UnanimousRequestData
		return sig, unanHash, nil
	}()

	var msg *valmessage.UnanimousAssertionFollowerResponse
	if err != nil {
		log.Println(err)
		msg = &valmessage.UnanimousAssertionFollowerResponse{
			Accepted: false,
		}
	} else {
		msg = &valmessage.UnanimousAssertionFollowerResponse{
			Accepted:      true,
			Signature:     sig,
			AssertionHash: value.NewHashBuf(unanHash),
		}
	}
	message := &valmessage.FollowerResponse{
		RequestId: value.NewHashBuf(requestID),
		Response:  &valmessage.FollowerResponse_Unanimous{Unanimous: msg},
	}
	raw, err := proto.Marshal(message)
	if err != nil {
		return err
	}
	m.client.ToClient <- raw
	return nil
}

func (m *ValidatorFollower) HandleCreateVM(ctx context.Context, request *valmessage.CreateVMValidatorRequest) *valmessage.FollowerResponse {
	createHash := hashing.CreateVMHash(request)
	failedReply := &valmessage.FollowerResponse{
		Response: &valmessage.FollowerResponse_Create{
			Create: &valmessage.CreateVMFollowerResponse{
				Accepted: false,
			},
		},
		RequestId: value.NewHashBuf(createHash),
	}
	var escrowCurrency common.Address
	copy(escrowCurrency[:], request.Config.EscrowCurrency.Value)
	escrowRequired := value.NewBigIntFromBuf(request.Config.EscrowRequired)
	address := m.Address()
	var user [32]byte
	copy(user[:], address[:])
	if err := m.WaitForTokenBalance(ctx, user, escrowCurrency, escrowRequired); err != nil {
		log.Printf("Follower failed to meet balance requirement: %v", err)
		return failedReply
	}
	sig, err := m.Sign(createHash)
	if err != nil {
		log.Printf("Follower failed to sign1: %v", err)
		return failedReply
	}

	return &valmessage.FollowerResponse{
		Response: &valmessage.FollowerResponse_Create{
			Create: &valmessage.CreateVMFollowerResponse{
				Accepted:  true,
				Signature: sig,
			},
		},
		RequestId: value.NewHashBuf(createHash),
	}
}

func (m *ValidatorFollower) Run() error {
	go func() {
		if err := m.client.run(); err != nil {
			log.Printf("Follower connection to coordinator ended with error %v\n", err)
		}
	}()

	go func() {
		for {
			message, more := <-m.client.FromClient
			if !more {
				break
			}
			req := new(valmessage.ValidatorRequest)
			err := proto.Unmarshal(message, req)
			if err != nil {
				log.Printf("Validator recieved malformed message")
				continue
			}
			switch request := req.Request.(type) {
			case *valmessage.ValidatorRequest_Unanimous:
				err := m.HandleUnanimousRequest(request.Unanimous, value.NewHashFromBuf(req.RequestId))
				if err != nil {
					log.Printf("Follower error while trying to handle unanimous assertion request from coordinator")
				}
			case *valmessage.ValidatorRequest_UnanimousNotification:
				requestInfo := m.unanimousRequests[value.NewHashFromBuf(req.RequestId)]
				if request.UnanimousNotification.Accepted {
					resultChan, errChan := m.Bot.ConfirmOffchainUnanimousAssertion(
						requestInfo,
						request.UnanimousNotification.Signatures,
						false,
					)
					select {
					case _ = <-resultChan:
					case err := <-errChan:
						log.Fatalln("Follower failed to confirm unanimous assertion", err)
					}
				}
			case *valmessage.ValidatorRequest_Create:
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
				response := m.HandleCreateVM(ctx, request.Create)
				cancel()
				raw, err := proto.Marshal(response)
				if err != nil {
					log.Fatalln("Follower failed to marshal response")
				}
				m.client.ToClient <- raw
			case *valmessage.ValidatorRequest_CreateNotification:
			}
		}
	}()
	return m.StartListening()
}
