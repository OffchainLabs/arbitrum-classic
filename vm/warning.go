package vm

import (
	"fmt"
)

type WarningHandler interface {
	AnyWarnings() bool
	Warn(string)
	Clone() WarningHandler
	SwitchMachinePC(stack *MachinePC)
}

type VerboseWarningHandler struct {
	pc          *MachinePC
	locations   []string
	anyWarnings bool
	num         int
}

func NewVerboseWarningHandler(m *MachinePC, locations []string) *VerboseWarningHandler {
	return &VerboseWarningHandler{m, locations, false, 0}
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
		fmt.Println(hand.locations[hand.pc.pc], ":", wstr)
	}
}

func (hand *VerboseWarningHandler) Clone() WarningHandler {
	return &VerboseWarningHandler{hand.pc, hand.locations, hand.anyWarnings, hand.num}
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
