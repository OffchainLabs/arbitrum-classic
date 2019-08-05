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

package vm

import "fmt"

type WarningHandler interface {
	AnyWarnings() bool
	Warn(string)
	Clone() WarningHandler
	SwitchMachinePC(stack *MachinePC)
}

type VerboseWarningHandler struct {
	pc          *MachinePC
	anyWarnings bool
	num         int
}

func NewVerboseWarningHandler(m *MachinePC) *VerboseWarningHandler {
	return &VerboseWarningHandler{m, false, 0}
}

func (hand *VerboseWarningHandler) AnyWarnings() bool {
	return hand.anyWarnings
}

func (hand *VerboseWarningHandler) Warn(wstr string) {
	hand.num = hand.num + 1
	if hand.num >= 10 {
		panic("Too many warnings")
	}
	if hand.pc != nil {
		fmt.Println(hand.pc, ":", wstr)
		// fmt.Println(hand.locations[hand.pc.pc], ":", wstr)
	}
}

func (hand *VerboseWarningHandler) Clone() WarningHandler {
	return &VerboseWarningHandler{hand.pc, hand.anyWarnings, hand.num}
}

func (hand *VerboseWarningHandler) SwitchMachinePC(m *MachinePC) {
	hand.pc = m
}

type SilentWarningHandler struct {
	anyWarnings bool
}

func NewSilentWarningHandler() *SilentWarningHandler {
	return &SilentWarningHandler{false}
}

func (hand *SilentWarningHandler) AnyWarnings() bool {
	return hand.anyWarnings
}

func (hand *SilentWarningHandler) Warn(wstr string) {
	hand.anyWarnings = true
}

func (hand *SilentWarningHandler) Clone() WarningHandler {
	return &SilentWarningHandler{hand.anyWarnings}
}

func (hand *SilentWarningHandler) SwitchMachinePC(m *MachinePC) {
	// do nothing
}
