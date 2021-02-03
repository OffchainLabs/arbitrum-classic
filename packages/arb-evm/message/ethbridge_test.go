/*
 * Copyright 2019-2020, Offchain Labs, Inc.
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

package message

//
//func TestUnmarshalOutgoing(t *testing.T) {
//	msg := NewRandomOutMessage(NewRandomEth())
//	var valData bytes.Buffer
//	if err := value.MarshalValue(msg.AsValue(), &valData); err != nil {
//		t.Fatal(err)
//	}
//	valid, offset, kind, sender, data, err := tester.UnmarshalOutgoingMessage(nil, valData.Bytes(), big.NewInt(0))
//	if err != nil {
//		t.Fatal(err)
//	}
//	if !valid {
//		t.Fatal("invalid l2message")
//	}
//	if offset.Uint64() != uint64(len(valData.Bytes())) {
//		t.Error("incorrect offset")
//	}
//	if inbox.Type(kind) != msg.Kind {
//		t.Error("incorrect l2message type")
//	}
//	if sender != msg.Sender.ToEthAddress() {
//		t.Error("incorrect sender")
//	}
//	if !bytes.Equal(data, msg.Data) {
//		t.Error("incorrect data")
//	}
//}
//
//func TestParseEthMessage(t *testing.T) {
//	msg := NewRandomEth()
//	ret, err := tester.ParseEthMessage(nil, msg.AsData())
//	if err != nil {
//		t.Fatal(err)
//	}
//	if !ret.Valid {
//		t.Error("invalid l2message")
//	}
//	if ret.Value.Cmp(msg.Value) != 0 {
//		t.Error("incorrect value")
//	}
//	if ret.Dest != msg.Dest.ToEthAddress() {
//		t.Error("incorrect address")
//	}
//}
//
//func TestParseERC20Message(t *testing.T) {
//	msg := NewRandomERC20()
//	ret, err := tester.ParseERC20Message(nil, msg.AsData())
//	if err != nil {
//		t.Fatal(err)
//	}
//	if !ret.Valid {
//		t.Error("invalid l2message")
//	}
//	if ret.Token != msg.Token.ToEthAddress() {
//		t.Error("incorrect token")
//	}
//	if ret.Value.Cmp(msg.Value) != 0 {
//		t.Error("incorrect value")
//	}
//	if ret.Dest != msg.Dest.ToEthAddress() {
//		t.Error("incorrect address")
//	}
//}
//
//func TestParseERC721Message(t *testing.T) {
//	msg := NewRandomERC721()
//	ret, err := tester.ParseERC721Message(nil, msg.AsData())
//	if err != nil {
//		t.Error(err)
//	}
//	if !ret.Valid {
//		t.Error("invalid l2message")
//	}
//	if ret.Token != msg.Token.ToEthAddress() {
//		t.Error("incorrect token")
//	}
//	if ret.Id.Cmp(msg.ID) != 0 {
//		t.Error("incorrect value")
//	}
//	if ret.Dest != msg.Dest.ToEthAddress() {
//		t.Error("incorrect address")
//	}
//}
