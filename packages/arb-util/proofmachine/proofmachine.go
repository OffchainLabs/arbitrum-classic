package proofmachine

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"

	"github.com/offchainlabs/arbitrum/packages/arb-util/common"
	"github.com/offchainlabs/arbitrum/packages/arb-util/core"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgecontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethbridgetestcontracts"
	"github.com/offchainlabs/arbitrum/packages/arb-util/ethutils"
)

type ProofData struct {
	Assertion   *core.Assertion
	Proof       hexutil.Bytes
	BufferProof hexutil.Bytes
}

type ProofChecker struct {
	osps            []*ethbridgetestcontracts.IOneStepProof
	sequencerBridge ethcommon.Address
	delayedBridge   ethcommon.Address
}

func NewProofChecker(auth *bind.TransactOpts, client *ethutils.SimulatedEthClient) (*ProofChecker, error) {
	sequencer := common.RandAddress().ToEthAddress()
	maxDelayBlocks := big.NewInt(60)
	maxDelaySeconds := big.NewInt(900)

	osp1Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof(auth, client)
	if err != nil {
		return nil, err
	}
	osp2Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProof2(auth, client)
	if err != nil {
		return nil, err
	}
	osp3Addr, _, _, err := ethbridgetestcontracts.DeployOneStepProofHash(auth, client)
	if err != nil {
		return nil, err
	}
	delayedBridgeAddr, _, _, err := ethbridgecontracts.DeployBridge(auth, client)
	if err != nil {
		return nil, err
	}
	sequencerAddr, _, sequencerCon, err := ethbridgecontracts.DeploySequencerInbox(auth, client)
	if err != nil {
		return nil, err
	}
	rollupAddr, _, rollup, err := ethbridgetestcontracts.DeployRollupMock(auth, client)
	if err != nil {
		return nil, err
	}
	client.Commit()

	_, err = rollup.SetMock(auth, maxDelayBlocks, maxDelaySeconds)
	if err != nil {
		return nil, err
	}
	_, err = sequencerCon.Initialize(auth, delayedBridgeAddr, sequencer, rollupAddr)
	if err != nil {
		return nil, err
	}
	client.Commit()

	osp1, err := ethbridgetestcontracts.NewIOneStepProof(osp1Addr, client)
	if err != nil {
		return nil, err
	}
	osp2, err := ethbridgetestcontracts.NewIOneStepProof(osp2Addr, client)
	if err != nil {
		return nil, err
	}
	osp3, err := ethbridgetestcontracts.NewIOneStepProof(osp3Addr, client)
	if err != nil {
		return nil, err
	}
	return &ProofChecker{
		osps:            []*ethbridgetestcontracts.IOneStepProof{osp1, osp2, osp3},
		sequencerBridge: sequencerAddr,
		delayedBridge:   delayedBridgeAddr,
	}, nil
}

func getProverNum(op uint8) uint8 {
	if (op >= 0xa1 && op <= 0xa6) || op == 0x70 {
		return 1
	} else if op >= 0x20 && op <= 0x24 {
		return 2
	} else {
		return 0
	}
}

func (p *ProofChecker) CheckProof(proof *ProofData) []error {
	op := proof.Proof[0]
	prover := getProverNum(op)
	machineData, err := p.osps[prover].ExecuteStep(
		&bind.CallOpts{},
		[2]ethcommon.Address{p.sequencerBridge, p.delayedBridge},
		proof.Assertion.Before.TotalMessagesRead,
		[2][32]byte{
			proof.Assertion.Before.SendAcc,
			proof.Assertion.Before.LogAcc,
		},
		proof.Proof,
		proof.BufferProof,
	)
	if err != nil {
		return []error{errors.Wrapf(err, "Solidity OSP execution failed for opcode 0x%x with prover %v", op, prover)}
	}

	correctGasUsed := proof.Assertion.GasUsed()
	var errorList []error
	if new(big.Int).SetUint64(machineData.Gas).Cmp(correctGasUsed) != 0 {
		err = errors.Errorf("wrong gas %v instead of %v", machineData.Gas, correctGasUsed)
		errorList = append(errorList, err)
	}
	if machineData.AfterMessagesRead.Cmp(proof.Assertion.After.TotalMessagesRead) != 0 {
		err = errors.Errorf("wrong total messages read %v %v", machineData.AfterMessagesRead, proof.Assertion.After.TotalMessagesRead)
		errorList = append(errorList, err)
	}
	if machineData.Fields[0] != proof.Assertion.Before.MachineHash {
		err = errors.Errorf("wrong before machine 0x%x 0x%x", machineData.Fields[0][:], proof.Assertion.After.MachineHash[:])
		errorList = append(errorList, err)
	}
	if machineData.Fields[2] != proof.Assertion.After.SendAcc {
		err = errors.Errorf("wrong send acc 0x%x 0x%x", machineData.Fields[2][:], proof.Assertion.After.SendAcc[:])
		errorList = append(errorList, err)
	}
	if machineData.Fields[3] != proof.Assertion.After.LogAcc {
		err = errors.Errorf("wrong log acc 0x%x 0x%x", machineData.Fields[3][:], proof.Assertion.After.LogAcc[:])
		errorList = append(errorList, err)
	}
	if machineData.Fields[1] != proof.Assertion.After.MachineHash {
		err = errors.Errorf("wrong after machine 0x%x 0x%x", machineData.Fields[1][:], proof.Assertion.After.MachineHash[:])
		errorList = append(errorList, err)
	}
	return errorList
}
