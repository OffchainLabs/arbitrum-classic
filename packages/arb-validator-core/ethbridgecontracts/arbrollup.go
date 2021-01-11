// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ArbRollupABI is the input ABI used to generate the binding from.
const ArbRollupABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"logsAccHash\",\"type\":\"bytes32[]\"}],\"name\":\"ConfirmedAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"ConfirmedValidAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[7]\",\"name\":\"fields\",\"type\":\"bytes32[7]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"importedMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numArbGas\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numSteps\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"messageCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeLogCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"logCount\",\"type\":\"uint64\"}],\"name\":\"RollupAsserted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"RollupChallengeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"RollupConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"initVMHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"gracePeriodTicks\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"arbGasSpeedLimitPerTick\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxExecutionSteps\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"stakeRequirement\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraConfig\",\"type\":\"bytes\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"RollupPruned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"RollupStakeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"toNodeHash\",\"type\":\"bytes32\"}],\"name\":\"RollupStakeMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"RollupStakeRefunded\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedStakers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"initalProtoStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"beforeSendCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"branches\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"deadlineTicks\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"challengeNodeData\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"logsAcc\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"vmProtoStateHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"messageCounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"address[]\",\"name\":\"stakerAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"stakerProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakerProofOffsets\",\"type\":\"uint256[]\"}],\"name\":\"confirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"enableStakerAllowList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeRequired\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getWithdrawnStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"globalInbox\",\"outputs\":[{\"internalType\":\"contractIGlobalInbox\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_vmState\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"_gracePeriodTicks\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_arbGasSpeedLimitPerTick\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"_maxExecutionSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_stakeRequirement\",\"type\":\"uint128\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactoryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_globalInboxAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"}],\"name\":\"init\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakerAddress\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"isValidLeaf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestConfirmedPriv\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"leaves\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[8]\",\"name\":\"fields\",\"type\":\"bytes32[8]\"},{\"internalType\":\"uint256[5]\",\"name\":\"fields2\",\"type\":\"uint256[5]\"},{\"internalType\":\"bytes32\",\"name\":\"validBlockHashPrecondition\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"validBlockHeightPrecondition\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"messageCount\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"logCount\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"prevChildType\",\"type\":\"uint32\"},{\"internalType\":\"uint64\",\"name\":\"numSteps\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"numArbGas\",\"type\":\"uint64\"},{\"internalType\":\"bytes32[]\",\"name\":\"stakerProof\",\"type\":\"bytes32[]\"}],\"name\":\"makeAssertion\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof1\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof2\",\"type\":\"bytes32[]\"}],\"name\":\"moveStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"ownerAddAllowedStaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"logsAcc\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"validNodeHashes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"finalNodeHash\",\"type\":\"bytes32\"}],\"name\":\"ownerConfirm\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"ownerRefundStaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"ownerRemoveAllowedStaker\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"messages\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"initialMaxSendCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"finalMaxSendCount\",\"type\":\"uint256\"}],\"name\":\"ownerSendMessages\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"ownerShutdown\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"shouldRequire\",\"type\":\"bool\"}],\"name\":\"ownerToggleStakerAllowListed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof1\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof2\",\"type\":\"bytes32[]\"}],\"name\":\"placeStake\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"fromNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"leafProofLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"latestConfProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"latestConfirmedProofLengths\",\"type\":\"uint256[]\"}],\"name\":\"pruneLeaves\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"recoverStakeConfirmed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"latestConfirmedProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"stakerProof\",\"type\":\"bytes32[]\"}],\"name\":\"recoverStakeMooted\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"recoverStakeOld\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"disputableNodeHashVal\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"childType\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"vmProtoStateHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"recoverStakePassedDeadline\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeRequirement\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"location\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"creationTimeBlocks\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"inChallenge\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"asserterAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"challengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"prevNode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"stakerNodeTypes\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"vmProtoHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[]\",\"name\":\"asserterProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"challengerProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"asserterNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengerDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"challengerPeriodTicks\",\"type\":\"uint128\"}],\"name\":\"startChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vmParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gracePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerTick\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"maxExecutionSteps\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ArbRollupFuncSigs maps the 4-byte function signature to its string representation.
var ArbRollupFuncSigs = map[string]string{
	"ffa1ad74": "VERSION()",
	"aa0ed4cb": "allowedStakers(address)",
	"5dbaf68b": "challengeFactory()",
	"08b43a19": "challenges(address)",
	"6e64beaa": "confirm(bytes32,uint256,uint256[],uint256[],bytes32[],bytes32[],bytes32[],uint256[],bytes,address[],bytes32[],uint256[])",
	"146b459d": "enableStakerAllowList()",
	"d16c305d": "getStakeRequired()",
	"500a1564": "getStakeToken()",
	"3fcc045a": "getWithdrawnStake(address)",
	"d489113a": "globalInbox()",
	"1dceffe7": "init(bytes32,uint128,uint128,uint64,uint128,address,address,address,address,bytes)",
	"6f791d29": "isMaster()",
	"6177fd18": "isStaked(address)",
	"57ca6d1b": "isValidLeaf(bytes32)",
	"65f7f80d": "latestConfirmed()",
	"3a218e98": "latestConfirmedPriv()",
	"151bcd2c": "leaves(bytes32)",
	"a0e9f382": "makeAssertion(bytes32[8],uint256[5],bytes32,uint256,uint64,uint64,uint32,uint64,uint64,bytes32[])",
	"dbad0a39": "moveStake(bytes32[],bytes32[])",
	"8da5cb5b": "owner()",
	"3d3f66d1": "ownerAddAllowedStaker(address)",
	"10e7e692": "ownerConfirm(bytes32[],bytes32[],bytes32)",
	"9b6c85fd": "ownerRefundStaker(address)",
	"eb908433": "ownerRemoveAllowedStaker(address)",
	"56373597": "ownerSendMessages(bytes,uint256,uint256)",
	"cfa80707": "ownerShutdown()",
	"27258fb7": "ownerToggleStakerAllowListed(bool)",
	"e0620d64": "placeStake(bytes32[],bytes32[])",
	"fcfd8d3f": "pruneLeaves(bytes32[],bytes32[],uint256[],bytes32[],uint256[])",
	"7cfaaf67": "recoverStakeConfirmed(bytes32[])",
	"33554032": "recoverStakeMooted(address,bytes32,bytes32[],bytes32[])",
	"113ec9d8": "recoverStakeOld(address,bytes32[])",
	"badb3f14": "recoverStakePassedDeadline(address,uint256,bytes32,uint256,bytes32,bytes32[])",
	"396f51cf": "resolveChallenge(address,address)",
	"b6f9bbb9": "stakeRequirement()",
	"51ed6a30": "stakeToken()",
	"dff69787": "stakerCount()",
	"9168ae72": "stakers(address)",
	"bac5963f": "startChallenge(address,address,bytes32,uint256,uint256[2],bytes32[2],bytes32[],bytes32[],bytes32,bytes32,uint128)",
	"bbc2cc00": "vmParams()",
	"eb2e74cb": "withdrawnStakes(address)",
}

// ArbRollupBin is the compiled bytecode used for deploying new contracts.
var ArbRollupBin = "0x60806040526000805460ff19166001179055616096806100206000396000f3fe6080604052600436106102465760003560e01c80636f791d2911610139578063bbc2cc00116100b6578063dff697871161007a578063dff6978714611382578063e0620d6414611397578063eb2e74cb14611455578063eb90843314611488578063fcfd8d3f146114bb578063ffa1ad741461167657610246565b8063bbc2cc001461123c578063cfa8070714611278578063d16c305d1461128d578063d489113a146112a2578063dbad0a39146112b757610246565b8063a0e9f382116100fd578063a0e9f38214610ea9578063aa0ed4cb14610f76578063b6f9bbb914610fa9578063bac5963f14610fda578063badb3f141461119957610246565b80636f791d2914610d765780637cfaaf6714610d8b5780638da5cb5b14610e065780639168ae7214610e1b5780639b6c85fd14610e7657610246565b80633d3f66d1116101c757806357ca6d1b1161018b57806357ca6d1b1461079f5780635dbaf68b146107c95780636177fd18146107de57806365f7f80d146108115780636e64beaa1461082657610246565b80633d3f66d1146106725780633fcc045a146106a5578063500a1564146106d857806351ed6a3014610709578063563735971461071e57610246565b80631dceffe71161020e5780631dceffe71461042957806327258fb7146105045780633355403214610530578063396f51cf146106105780633a218e981461064b57610246565b806308b43a191461024b57806310e7e69214610292578063113ec9d81461035f578063146b459d146103ea578063151bcd2c146103ff575b600080fd5b34801561025757600080fd5b5061027e6004803603602081101561026e57600080fd5b50356001600160a01b0316611700565b604080519115158252519081900360200190f35b34801561029e57600080fd5b5061035d600480360360608110156102b557600080fd5b810190602081018135600160201b8111156102cf57600080fd5b8201836020820111156102e157600080fd5b803590602001918460208302840111600160201b8311171561030257600080fd5b919390929091602081019035600160201b81111561031f57600080fd5b82018360208201111561033157600080fd5b803590602001918460208302840111600160201b8311171561035257600080fd5b919350915035611715565b005b34801561036b57600080fd5b5061035d6004803603604081101561038257600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156103ac57600080fd5b8201836020820111156103be57600080fd5b803590602001918460208302840111600160201b831117156103df57600080fd5b50909250905061189d565b3480156103f657600080fd5b5061027e611951565b34801561040b57600080fd5b5061027e6004803603602081101561042257600080fd5b5035611961565b34801561043557600080fd5b5061035d600480360361014081101561044d57600080fd5b8135916001600160801b03602082013581169260408301358216926001600160401b0360608201351692608082013516916001600160a01b0360a083013581169260c081013582169260e08201358316926101008301351691908101906101408101610120820135600160201b8111156104c657600080fd5b8201836020820111156104d857600080fd5b803590602001918460018302840111600160201b831117156104f957600080fd5b509092509050611976565b34801561051057600080fd5b5061035d6004803603602081101561052757600080fd5b50351515611c05565b34801561053c57600080fd5b5061035d6004803603608081101561055357600080fd5b6001600160a01b0382351691602081013591810190606081016040820135600160201b81111561058257600080fd5b82018360208201111561059457600080fd5b803590602001918460208302840111600160201b831117156105b557600080fd5b919390929091602081019035600160201b8111156105d257600080fd5b8201836020820111156105e457600080fd5b803590602001918460208302840111600160201b8311171561060557600080fd5b509092509050611c9f565b34801561061c57600080fd5b5061035d6004803603604081101561063357600080fd5b506001600160a01b0381358116916020013516611df7565b34801561065757600080fd5b50610660611f3b565b60408051918252519081900360200190f35b34801561067e57600080fd5b5061035d6004803603602081101561069557600080fd5b50356001600160a01b0316611f41565b3480156106b157600080fd5b5061035d600480360360208110156106c857600080fd5b50356001600160a01b0316611fe1565b3480156106e457600080fd5b506106ed61214a565b604080516001600160a01b039092168252519081900360200190f35b34801561071557600080fd5b506106ed612159565b34801561072a57600080fd5b5061035d6004803603606081101561074157600080fd5b810190602081018135600160201b81111561075b57600080fd5b82018360208201111561076d57600080fd5b803590602001918460018302840111600160201b8311171561078e57600080fd5b919350915080359060200135612168565b3480156107ab57600080fd5b5061027e600480360360208110156107c257600080fd5b503561228d565b3480156107d557600080fd5b506106ed6122a5565b3480156107ea57600080fd5b5061027e6004803603602081101561080157600080fd5b50356001600160a01b03166122b4565b34801561081d57600080fd5b506106606122d1565b34801561083257600080fd5b5061035d600480360361018081101561084a57600080fd5b813591602081013591810190606081016040820135600160201b81111561087057600080fd5b82018360208201111561088257600080fd5b803590602001918460208302840111600160201b831117156108a357600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156108f257600080fd5b82018360208201111561090457600080fd5b803590602001918460208302840111600160201b8311171561092557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561097457600080fd5b82018360208201111561098657600080fd5b803590602001918460208302840111600160201b831117156109a757600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156109f657600080fd5b820183602082011115610a0857600080fd5b803590602001918460208302840111600160201b83111715610a2957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a7857600080fd5b820183602082011115610a8a57600080fd5b803590602001918460208302840111600160201b83111715610aab57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610afa57600080fd5b820183602082011115610b0c57600080fd5b803590602001918460208302840111600160201b83111715610b2d57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b7c57600080fd5b820183602082011115610b8e57600080fd5b803590602001918460018302840111600160201b83111715610baf57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610c0157600080fd5b820183602082011115610c1357600080fd5b803590602001918460208302840111600160201b83111715610c3457600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610c8357600080fd5b820183602082011115610c9557600080fd5b803590602001918460208302840111600160201b83111715610cb657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610d0557600080fd5b820183602082011115610d1757600080fd5b803590602001918460208302840111600160201b83111715610d3857600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506122d7945050505050565b348015610d8257600080fd5b5061027e612330565b348015610d9757600080fd5b5061035d60048036036020811015610dae57600080fd5b810190602081018135600160201b811115610dc857600080fd5b820183602082011115610dda57600080fd5b803590602001918460208302840111600160201b83111715610dfb57600080fd5b509092509050612339565b348015610e1257600080fd5b506106ed612376565b348015610e2757600080fd5b50610e4e60048036036020811015610e3e57600080fd5b50356001600160a01b0316612385565b604080519384526001600160801b039092166020840152151582820152519081900360600190f35b348015610e8257600080fd5b5061035d60048036036020811015610e9957600080fd5b50356001600160a01b03166123b3565b348015610eb557600080fd5b5061035d60048036036102a0811015610ecd57600080fd5b6101008201906101a0830135906101c0840135906001600160401b036101e086013581169161020087013582169163ffffffff61022089013516916102408901358216916102608a013516908901896102a08101610280820135600160201b811115610f3857600080fd5b820183602082011115610f4a57600080fd5b803590602001918460208302840111600160201b83111715610f6b57600080fd5b509092509050612438565b348015610f8257600080fd5b5061027e60048036036020811015610f9957600080fd5b50356001600160a01b031661265b565b348015610fb557600080fd5b50610fbe612670565b604080516001600160801b039092168252519081900360200190f35b348015610fe657600080fd5b5061035d60048036036101a0811015610ffe57600080fd5b6040805180820182526001600160a01b0384358116946020810135909116938382013593606083013593918301929160c083019160808401906002908390839080828437600092019190915250506040805180820182529295949381810193925090600290839083908082843760009201919091525091949392602081019250359050600160201b81111561109257600080fd5b8201836020820111156110a457600080fd5b803590602001918460208302840111600160201b831117156110c557600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b81111561111457600080fd5b82018360208201111561112657600080fd5b803590602001918460208302840111600160201b8311171561114757600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050823593505050602081013590604001356001600160801b031661267f565b3480156111a557600080fd5b5061035d600480360360c08110156111bc57600080fd5b6001600160a01b03823516916020810135916040820135916060810135916080820135919081019060c0810160a0820135600160201b8111156111fe57600080fd5b82018360208201111561121057600080fd5b803590602001918460208302840111600160201b8311171561123157600080fd5b509092509050612ab1565b34801561124857600080fd5b50611251612c35565b6040805193845260208401929092526001600160401b031682820152519081900360600190f35b34801561128457600080fd5b5061035d612c4a565b34801561129957600080fd5b50610fbe612cd1565b3480156112ae57600080fd5b506106ed612ce0565b3480156112c357600080fd5b5061035d600480360360408110156112da57600080fd5b810190602081018135600160201b8111156112f457600080fd5b82018360208201111561130657600080fd5b803590602001918460208302840111600160201b8311171561132757600080fd5b919390929091602081019035600160201b81111561134457600080fd5b82018360208201111561135657600080fd5b803590602001918460208302840111600160201b8311171561137757600080fd5b509092509050612cef565b34801561138e57600080fd5b50610660612e00565b61035d600480360360408110156113ad57600080fd5b810190602081018135600160201b8111156113c757600080fd5b8201836020820111156113d957600080fd5b803590602001918460208302840111600160201b831117156113fa57600080fd5b919390929091602081019035600160201b81111561141757600080fd5b82018360208201111561142957600080fd5b803590602001918460208302840111600160201b8311171561144a57600080fd5b509092509050612e06565b34801561146157600080fd5b506106606004803603602081101561147857600080fd5b50356001600160a01b0316612f81565b34801561149457600080fd5b5061035d600480360360208110156114ab57600080fd5b50356001600160a01b0316612f93565b3480156114c757600080fd5b5061035d600480360360a08110156114de57600080fd5b810190602081018135600160201b8111156114f857600080fd5b82018360208201111561150a57600080fd5b803590602001918460208302840111600160201b8311171561152b57600080fd5b919390929091602081019035600160201b81111561154857600080fd5b82018360208201111561155a57600080fd5b803590602001918460208302840111600160201b8311171561157b57600080fd5b919390929091602081019035600160201b81111561159857600080fd5b8201836020820111156115aa57600080fd5b803590602001918460208302840111600160201b831117156115cb57600080fd5b919390929091602081019035600160201b8111156115e857600080fd5b8201836020820111156115fa57600080fd5b803590602001918460208302840111600160201b8311171561161b57600080fd5b919390929091602081019035600160201b81111561163857600080fd5b82018360208201111561164a57600080fd5b803590602001918460208302840111600160201b8311171561166b57600080fd5b509092509050613030565b34801561168257600080fd5b5061168b61316e565b6040805160208082528351818301528351919283929083019185019080838360005b838110156116c55781810151838201526020016116ad565b50505050905090810190601f1680156116f25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600b6020526000908152604090205460ff1681565b600d5460408051808201909152600a81526927a7262cafa7aba722a960b11b6020820152906001600160a01b031633146117cd5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561179257818101518382015260200161177a565b50505050905090810190601f1680156117bf5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b508160005b81811015611821578484828181106117e657fe5b905060200201357f89cc5e236414c34f1206c0c14d8ac5b0e5444b669b309aaca16fe3d27749f50e60405160405180910390a26001016117d2565b5061182b8261318f565b8015611895577fded5fa103431438087188a5f8c6a4c3ea90996bbd63be7b1b3fa0a425b37fdd5868660405180806020018281038252848482818152602001925060200280828437600083820152604051601f909101601f19169092018290039550909350505050a15b505050505050565b60408051808201909152600e81526d0a48a86ac9e9888be988a9c8ea8960931b60208201528161190e5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5061194c838383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506131ca92505050565b505050565b600e54600160a01b900460ff1681565b60046020526000908152604090205460ff1681565b7f3f3efae8ec7ea5f2d06aa37b37bb676f94c915ba05679d32ccdd0dc570dd58648b8b8b8b8b8a888860405180898152602001886001600160801b03166001600160801b03168152602001876001600160801b03166001600160801b03168152602001866001600160401b03166001600160401b03168152602001856001600160801b03166001600160801b03168152602001846001600160a01b03166001600160a01b03168152602001806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039b50909950505050505050505050a1611a6c8b8b8b8b613269565b611a778787866132e2565b600e80546001600160a01b038086166001600160a01b03199283161792839055600d8054898316931692909217909155604080516001600160801b03808f16602083018181528f83169484018590526001600160401b038f166060858101829052938f16608086018190526bffffffffffffffffffffffff198f861b811660a08801819052958f901b1660c087018190529890971697635cc96efa9793969591949290918b918b9160e00183838082843780830192505050985050505050505050506040516020818303038152906040526040518263ffffffff1660e01b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611b94578181015183820152602001611b7c565b50505050905090810190601f168015611bc15780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015611be057600080fd5b505af1158015611bf4573d6000803e3d6000fd5b505050505050505050505050505050565b600d5460408051808201909152600a81526927a7262cafa7aba722a960b11b6020820152906001600160a01b03163314611c805760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50600e8054911515600160a01b0260ff60a01b19909216919091179055565b6000611caa87613424565b905082826000818110611cb957fe5b9050602002013585856000818110611ccd57fe5b9050602002013514158015611d245750611ce56122d1565b611d22878787808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134b492505050565b145b8015611d6b575080611d69878585808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134b492505050565b145b604051806040016040528060148152602001732922a1a7ab2fa1a7a7232624a1aa2fa82927a7a360611b81525090611de45760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50611dee876134cb565b50505050505050565b336000908152600b6020908152604091829020548251808401909352600f83526e2922a9afa1a420a62fa9a2a72222a960891b9183019190915260ff16611e7f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50336000908152600b60205260408120805460ff19169055611ea08361353a565b6007546001600160a01b0385166000908152600c60205260409020805460026001600160801b0393841604909216909101905560018101805460ff60801b191690559050611eed826135c4565b604080513381526001600160a01b03808616602083015284168183015290517f468aa7d460319b17466ca163bca353a0c62fff0d7d0fa287f634ef305d946f299181900360600190a1505050565b60055481565b600d5460408051808201909152600a81526927a7262cafa7aba722a960b11b6020820152906001600160a01b03163314611fbc5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b506001600160a01b03166000908152600f60205260409020805460ff19166001179055565b6001600160a01b0381166000908152600c6020526040902054806120055750612147565b6008546001600160a01b0316612051576040516001600160a01b0383169082156108fc029083906000818181858888f1935050505015801561204b573d6000803e3d6000fd5b50612145565b6008546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b1580156120a757600080fd5b505af11580156120bb573d6000803e3d6000fd5b505050506040513d60208110156120d157600080fd5b505160408051808201909152600f81526e1514905394d1915497d19052531151608a1b60208201529061194c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b505b50565b6008546001600160a01b031690565b6008546001600160a01b031681565b600d5460408051808201909152600a81526927a7262cafa7aba722a960b11b6020820152906001600160a01b031633146121e35760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50600e54604051630caba3af60e41b81526024810184905260448101839052606060048201908152606482018690526001600160a01b039092169163caba3af09187918791879187918190608401868680828437600081840152601f19601f82011690508083019250505095505050505050600060405180830381600087803b15801561226f57600080fd5b505af1158015612283573d6000803e3d6000fd5b5050505050505050565b60008181526004602052604090205460ff165b919050565b6006546001600160a01b031681565b6001600160a01b0316600090815260096020526040902054151590565b60055490565b6123226040518061012001604052808e81526020018d81526020018c81526020018b81526020018a815260200189815260200188815260200187815260200186815250848484613605565b505050505050505050505050565b60005460ff1690565b612145338383808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506131ca92505050565b600d546001600160a01b031681565b600960205260009081526040902080546001909101546001600160801b03811690600160801b900460ff1683565b600d5460408051808201909152600a81526927a7262cafa7aba722a960b11b6020820152906001600160a01b0316331461242e5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50612147816134cb565b88884014612483576040805162461bcd60e51b8152602060048201526013602482015272696e76616c6964206b6e6f776e20626c6f636b60681b604482015290519081900360640190fd5b61248b615f11565b604080516101008181019092526124e5918e906008908390839080828437600092019190915250506040805160a081810190925291508e9060059083908390808284376000920191909152508a9150899050888d8d61393d565b600e5460408051630220168160e01b8152306004820152815193945060009384936001600160a01b03169263022016819260248082019391829003018186803b15801561253157600080fd5b505afa158015612545573d6000803e3d6000fd5b505050506040513d604081101561255b57600080fd5b5080516020909101519092509050600080612577858585613ab5565b91509150600061258633613424565b9050826125c6828a8a808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134b492505050565b146040518060400160405280601181526020017026a0a5a2afa9aa20a5a2a92fa82927a7a360791b8152509061263d5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b506126483383613c1b565b5050505050505050505050505050505050565b600f6020526000908152604090205460ff1681565b6007546001600160801b031681565b600061268a8c61353a565b905060006126978c61353a565b60018301549091508a906126b3906001600160801b0316613c74565b106040518060400160405280600d81526020016c53544b315f444541444c494e4560981b815250906127265760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5060018101548a90612740906001600160801b0316613c74565b106040518060400160405280600d81526020016c53544b325f444541444c494e4560981b815250906127b35760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50600182015460408051808201909152600c81526b14d512cc57d25397d0d2105360a21b602082015290600160801b900460ff16156128335760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50600181015460408051808201909152600c81526b14d512cc97d25397d0d2105360a21b602082015290600160801b900460ff16156128b35760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b506020808a01518a5160408051808201909152600a8152692a2ca822afa7a92222a960b11b938101939093521161292b5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50815461295361294d8d8d898e600060200201518e60005b6020020151613c7b565b896134b4565b146040518060400160405280600c81526020016b20a9a9a2a92a2fa82927a7a360a11b815250906129c55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5080546129f66129f08d8d6129e3896001600160801b038a16613ce3565b60208f01518e6001612943565b886134b4565b146040518060400160405280600a81526020016921a420a62fa82927a7a360b11b81525090612a665760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5060018281018054600160801b60ff60801b19918216811790925591830180549092161790556020890151612aa2908e908e9086908890613d0f565b50505050505050505050505050565b6000612abc88613424565b90506000612acd8289898989613c7b565b90506000612b0e828686808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134b492505050565b9050612b198161228d565b604051806040016040528060138152602001722922a1a7ab2fa222a0a22624a722afa622a0a360691b81525090612b915760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50612b9b89613e2d565b6001600160801b0316431015604051806040016040528060138152602001725245434f565f444541444c494e455f54494d4560681b81525090612c1f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50612c298a6134cb565b50505050505050505050565b6001546002546003546001600160401b031683565b600d5460408051808201909152600a81526927a7262cafa7aba722a960b11b6020820152906001600160a01b03163314612cc55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50612ccf33613e35565b565b6007546001600160801b031690565b600e546001600160a01b031681565b6000612cfa33613424565b90506000612d3b828787808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134b492505050565b90506000612d7c828686808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134b492505050565b9050612d878161228d565b6040518060400160405280600981526020016826a7ab22afa622a0a360b91b81525090612df55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50611dee3383613c1b565b600a5481565b600e54600160a01b900460ff161580612e2e5750336000908152600f602052604090205460ff165b612e76576040805162461bcd60e51b81526020600482015260146024820152734e4f545f414c4c4f5745445f544f5f5354414b4560601b604482015290519081900360640190fd5b6000612ebc612e836122d1565b8686808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134b492505050565b90506000612efd828585808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152506134b492505050565b9050612f088161228d565b6040518060400160405280600a815260200169282620a1a2afa622a0a360b11b81525090612f775760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5061189582613eb5565b600c6020526000908152604090205481565b600d5460408051808201909152600a81526927a7262cafa7aba722a960b11b6020820152906001600160a01b0316331461300e5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b506001600160a01b03166000908152600f60205260409020805460ff19169055565b88858114801561303f57508181145b613089576040805162461bcd60e51b81526020600482015260166024820152750d2dce0eae840d8cadccee8d040dad2e6e8dac2e8c6d60531b604482015290519081900360640190fd5b600080805b8381101561315e576131518e8e838181106130a557fe5b905060200201358787848181106130b857fe5b905060200201358c8c858181106130cb57fe5b905060200201358f8f80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508b92508a91506141e99050565b909350915060010161308e565b5050505050505050505050505050565b6040518060400160405280600581526020016418171b971960d91b81525081565b60058190556040805182815290517f9d13d0ad532ca8e545a3b66828cb99a18c3bc98e2a50b4db1990a033fdba6f639181900360200190a150565b60006131d533613424565b90506131df6122d1565b6131e982846134b4565b146040518060400160405280601081526020016f2922a1a7ab2fa820aa242fa82927a7a360811b8152509061325f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5061194c836134cb565b600061327885828080806143b2565b905060006132898180808086613c7b565b600581905560009081526004602052604090208054600160ff1990911681179091556001600160801b039586169055505091166002556003805467ffffffffffffffff19166001600160401b0390921691909117905550565b60065460408051808201909152600a815269494e49545f545749434560b01b6020820152906001600160a01b03161561335c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5060408051808201909152600c81526b494e49545f4e4f4e5a45524f60a01b60208201526001600160a01b0382166133d55760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50600680546001600160a01b03199081166001600160a01b0393841617909155600780546001600160801b0319166001600160801b039590951694909417909355600880549093169116179055565b6001600160a01b0381166000908152600960209081526040808320548151808301909252600a82526924a72b2fa9aa20a5a2a960b11b92820192909252816134ad5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5092915050565b60006134c48383600085516143f8565b9392505050565b6134d4816135c4565b6007546001600160a01b0382166000818152600c602090815260409182902080546001600160801b039095169094019093558051918252517f953ab9eece73c907353307064109cf873462177a0e358e463fd89f5b206daa6c929181900390910190a150565b6001600160a01b038116600090815260096020908152604080832080548251808401909352600a83526924a72b2fa9aa20a5a2a960b11b93830193909352916134ad5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b6001600160a01b03166000908152600960205260408120908155600101805470ffffffffffffffffffffffffffffffffff19169055600a8054600019019055565b60408401515160608501518051600019830190811061362057fe5b602002602001015161363143613c74565b101560405180604001604052806009815260200168434f4e465f54494d4560b81b815250906136a15760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5060606136ac615f6e565b6136bd876136b86122d1565b614460565b8151919350915060005b81811015613717578381815181106136db57fe5b60200260200101517f89cc5e236414c34f1206c0c14d8ac5b0e5444b669b309aaca16fe3d27749f50e60405160405180910390a26001016136c7565b5060006137458360a001518a60600151600188038151811061373557fe5b60200260200101518a8a8a614524565b9050600081116040518060400160405280600f81526020016e21a7a7232fa420a9afa9aa20a5a2a960891b815250906137bf5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b506137cd8360a0015161318f565b600e546101008a01516020808c01516080870151604051630caba3af60e41b815260248101839052604481018290526060600482019081528551606483015285516001600160a01b039097169663caba3af09695919283926084909101919087019080838360005b8381101561384d578181015183820152602001613835565b50505050905090810190601f16801561387a5780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561389b57600080fd5b505af11580156138af573d6000803e3d6000fd5b5050505060008211156139325760a089015160408051602080825283518183015283517fded5fa103431438087188a5f8c6a4c3ea90996bbd63be7b1b3fa0a425b37fdd59493839290830191818601910280838360005b8381101561391e578181015183820152602001613906565b505050509050019250505060405180910390a15b505050505050505050565b613945615f11565b61394d615faa565b60408051610180810182526001600160401b03888116825287166020808301919091528b51828401528b0151606080830191909152918b01516080820152908a015160a0820152600060c082015260e081018a600460200201518152602001856001600160401b031681526020016000801b81526020018a6005600881106139d157fe5b60200201518152602001846001600160401b0316815250905060405180610120016040528089600060058110613a0357fe5b602002015181526020018a600660088110613a1a57fe5b6020020151815260200189600160058110613a3157fe5b602002015181526020018a600760088110613a4857fe5b602002015181526020018863ffffffff16815260200189600260058110613a6b57fe5b6020020151815260200189600360058110613a8257fe5b6020020151815260200189600460058110613a9957fe5b6020020151815260200182815250915050979650505050505050565b600080600080613ac4876146a2565b91509150613ad18261228d565b6040518060400160405280600981526020016826a0a5a2afa622a0a360b91b81525090613b3f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50613b49876146fc565b8651613b5c90869063ffffffff61481d16565b8760a0015111156040518060400160405280601081526020016f135052d157d35154d4d051d157d0d39560821b81525090613bd85760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b506000613be88884848a8a614856565b6000848152600460205260409020805460ff191690559050613c0d8884838a8a61491b565b919791965090945050505050565b6001600160a01b0382166000818152600960209081526040918290208490558151928352820183905280517fbe690ac5fe353c094bcc6f187eeb841c0ca61b6edf32c142eadad655b7d173f49281900390910190a15050565b6103e80290565b6040805160208082018490528183018790526060820186905260808083018690528351808403909101815260a08301845280519082012060c0830189905260e08084019190915283518084039091018152610100909201909252805191012095945050505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6006546040805163432ed0e160e11b81526001600160a01b03888116600483015287811660248301526001600160801b038716604483015260648201869052608482018590529151600093929092169163865da1c29160a48082019260209290919082900301818787803b158015613d8657600080fd5b505af1158015613d9a573d6000803e3d6000fd5b505050506040513d6020811015613db057600080fd5b50516001600160a01b038082166000818152600b6020908152604091829020805460ff1916600117905581518b85168152938a16908401528281018690526060830191909152519192507f6c69257ddf620994c6fb9e5304db0e5563db3765bee033ddd61b6a1caa7d043f919081900360800190a1505050505050565b6103e8900490565b6000546040805180820190915260098152684e4f545f434c4f4e4560b81b60208201529060ff1615613ea85760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50806001600160a01b0316ff5b6008546001600160a01b0316613f425760078054604080518082019091529182526614d512d7d0535560ca1b60208301526001600160801b03163414613f3c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b506140b3565b60408051808201909152600781526614d512d7d0535560ca1b60208201523415613fad5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50600854600754604080516323b872dd60e01b81523360048201523060248201526001600160801b039092166044830152516001600160a01b03909216916323b872dd916064808201926020929091908290030181600087803b15801561401357600080fd5b505af1158015614027573d6000803e3d6000fd5b505050506040513d602081101561403d57600080fd5b505160408051808201909152600f81526e1514905394d1915497d19052531151608a1b6020820152906140b15760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b505b33600090815260096020908152604091829020548251808401909352600c83526b105314911657d4d51052d15160a21b91830191909152156141365760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50604080516060810182528281526001600160801b03438116602080840191825260008486018181523380835260098452918790209551865592516001958601805494511515600160801b0260ff60801b19929096166001600160801b0319909516949094171693909317909155600a8054909301909255825190815290810183905281517fcbafbb223ed21c82af9e2ad20cdfdf55d3263d06f9a65b3f70da613f32d81f88929181900390910190a150565b6000806000871180156141fc5750600088115b6040518060400160405280600e81526020016d28292aa722afa82927a7a32622a760911b8152509061426f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5083870183890160006142806122d1565b61428c8d8a89866143f8565b1490508080156142c257508786815181106142a357fe5b60200260200101518988815181106142b757fe5b602002602001015114155b6040518060400160405280600e81526020016d141495539157d0d3d391931250d560921b815250906143355760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5060006143448d8b8a876143f8565b905061434f8161228d565b156143a057600081815260046020908152604091829020805460ff19169055815183815291517f3d3e2ada9638548d1bb115fd766ef675213d953efe8d433bbd8d6718f44909509281900390910190a15b50919b909a5098505050505050505050565b60408051602080820197909752808201959095526060850193909352608084019190915260a0808401919091528151808403909101815260c09092019052805191012090565b600084835b83811015614456578186828151811061441257fe5b6020026020010151604051602001808381526020018281526020019250505060405160208183030381529060405280519060200120915080806001019150506143fd565b5095945050505050565b606061446a615f6e565b61447384614a55565b6040808501515160e0860151518251818152602080830282010190935290918180156144a9578160200160208202803883390190505b5093506144b4615f6e565b6144c78760000151886020015188614c40565b905060005b838110156145175760006144e1898484614c81565b9050801561450e578260a001518760018560000151038151811061450157fe5b6020026020010181815250505b506001016144cc565b50925050505b9250929050565b60008084519050600a5481146040518060400160405280600a81526020016910d210d2d7d0d3d5539560b21b8152509061459f5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50825181600101146040518060400160405280600c81526020016b4348434b5f4f46465345545360a01b815250906146185760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5060008080805b8481101561469357600089828151811061463557fe5b602002602001015190506146778c8a8d84898d888151811061465357fe5b60200260200101518e896001018151811061466a57fe5b6020026020010151614d5a565b92508215614686576001909301925b60601b935060010161461f565b50909998505050505050505050565b6000806146d0836101000151604001518461010001516080015185600001518660c001518760e001516143b2565b90506146f5836020015184604001518560600151866080015163ffffffff1685613c7b565b9150915091565b61470e81610100015160400151614eb7565b15801561472a575061472881610100015160400151614ebd565b155b6040518060400160405280600881526020016726a0a5a2afa92aa760c11b815250906147975760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b506003546101008201515160408051808201909152600981526804d414b455f535445560bc1b6020820152916001600160401b03908116911611156121455760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b60006134c48383604051806040016040528060148152602001737375627472616374696f6e206f766572666c6f7760601b815250614ec1565b60408051606081018252600154815260025460208201526003546001600160401b03169181019190915260009081908190614892908943614f1b565b9150915060006148ac89898489898c600160000154614f76565b905060006148c38a8a858b60016000015489614fc3565b905060006148d28b8b86614ff6565b6000938452600460205260408085208054600160ff1991821681179092559486528186208054861682179055828652942080549093169093179091555098975050505050505050565b7f3112f8dec1eebe04bc7f92bf1031a1c749e09e57836222fef69df63d591bf6036040518060e0016040528086815260200184815260200187610100015160600151815260200187610100015160a00151815260200187610100015160e0015181526020018761010001516101400151815260200185815250828760a0015188610100015160200151896101000151600001518a60c001518b610100015161010001518c60e001518d61010001516101600151604051808a600760200280838360005b838110156149f65781810151838201526020016149de565b5050505091909101998a52505060208801969096526001600160401b0394851660408089019190915293851660608801526080870192909252831660a086015260c0850152911660e08301525190819003610100019150a15050505050565b6040808201515160e08301515160c084015151835180850190945260088452670434f4e465f494e560c41b6020850152919290918214614ad65760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50808360a001515114604051806040016040528060088152602001670434f4e465f494e560c41b81525090614b4c5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b508183606001515114604051806040016040528060088152602001670434f4e465f494e560c41b81525090614bc25760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b5080820383608001515114604051806040016040528060088152602001670434f4e465f494e560c41b81525090614c3a5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b50505050565b614c48615f6e565b6040518060c001604052806000815260200160008152602001600081526020018581526020018481526020018381525090509392505050565b60008084604001518381518110614c9457fe5b602002602001015190506000614ca861508f565b8214905060008115614cef57614ccc87876000015188608001518960400151615094565b60608a015260408901919091526080880191909152865160010187529050614d1e565b8660800151866020015181518110614d0357fe5b60200260200101519050856020018051809190600101815250505b614d4a8660a0015188606001518781518110614d3657fe5b602002602001015183868a60600151613c7b565b60a0870152509150509392505050565b6000836bffffffffffffffffffffffff19168560601b6bffffffffffffffffffffffff1916116040518060400160405280600a81526020016921a421a5afa7a92222a960b11b81525090614def5760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b506000614dfb8661353a565b60018101549091506000908890614e1a906001600160801b0316613c74565b1090508015614eaa578154614e318b8b88886143f8565b146040518060400160405280601181526020017021a421a5afa9aa20a5a2a92fa82927a7a360791b81525090614ea85760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b505b9998505050505050505050565b60011490565b1590565b60008184841115614f135760405162461bcd60e51b815260206004820181815283516024840152835190928392604490910191908501908083836000831561179257818101518382015260200161177a565b505050900390565b60008060008560200151856101000151602001516001600160401b031681614f3f57fe5b04905060008660000151614f5286613c74565b0190508560400151811015614f68575060408501515b909690870195509350505050565b600080614f9989610100015160a00151878b60a001518c60000151018803615124565b9050614eaa8888614fb584614fae6001613c74565b8801613ce3565b614fbd61515b565b88613c7b565b6000614feb8686614fe3614fdb8b6101000151615160565b868801613ce3565b614fbd615270565b979650505050505050565b6000615087838361501f8760c0015188610100015160e001518961010001516101400151615124565b61502761508f565b615082896101000151606001518a610100015160a001518b60a001518c60000151018c610100015161010001516001600160401b03168d60c00151018d610100015161016001516001600160401b03168e60e00151016143b2565b613c7b565b949350505050565b600290565b60008060008060008860e0015188815181106150ac57fe5b602002602001015190506000806150c98b61010001518985615275565b9150915060006150f18a848e60a001518e815181106150e457fe5b6020026020010151615124565b905060008c60c001518c8151811061510557fe5b60209081029190910101519a9094019c919b5099509650505050505050565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b600090565b6000816000015182602001518360400151846060015185608001518660a001518760c001518860e001518961010001518a61012001518b61014001518c6101600151604051602001808d6001600160401b03166001600160401b031660c01b81526008018c6001600160401b03166001600160401b031660c01b81526008018b81526020018a8152602001898152602001888152602001878152602001868152602001856001600160401b03166001600160401b031660c01b8152600801848152602001838152602001826001600160401b03166001600160401b031660c01b81526008019c50505050505050505050505050604051602081830303815290604052805190602001209050919050565b600190565b6000808061528161600e565b8560005b868110156152d85761529789836152e7565b93509150836152a584615472565b60408051602080820194909452808201929092528051808303820181526060909201905280519101209350600101615285565b50919791965090945050505050565b60006152f161600e565b83518310615337576040805162461bcd60e51b815260206004820152600e60248201526d1a5b9d985b1a59081bd9999cd95d60921b604482015290519081900360640190fd5b6000806153448686615592565b9150915061535061515b565b60ff168160ff16141561538457600061536987846155b9565b9093509050826153788261562d565b9450945050505061451d565b61538c615270565b60ff168160ff1614156153ae576153a386836156df565b93509350505061451d565b6153b661508f565b60ff168160ff1614156153cd576153a38683615781565b6153d5615815565b60ff168160ff16101580156153f657506153ed61581a565b60ff168160ff16105b15615432576000615405615815565b82039050606061541682898661581f565b909450905083615425826158b8565b955095505050505061451d565b6040805162461bcd60e51b815260206004820152601060248201526f696e76616c69642074797065636f646560801b604482015290519081900360640190fd5b600061547c61515b565b60ff16826060015160ff16141561549f578151615498906159ca565b90506122a0565b6154a7615270565b60ff16826060015160ff1614156154c55761549882602001516159ee565b6154cd61508f565b60ff16826060015160ff1614156154ef57815160808301516154989190615aeb565b6154f7615815565b60ff16826060015160ff1614156155305761551061600e565b61551d8360400151615b3c565b905061552881615472565b9150506122a0565b615538615c9e565b60ff16826060015160ff161415615551575080516122a0565b6040805162461bcd60e51b8152602060048201526011602482015270496e76616c6964207479706520636f646560781b604482015290519081900360640190fd5b600080826001018484815181106155a557fe5b016020015190925060f81c90509250929050565b600080828451101580156155d1575060208385510310155b61560e576040805162461bcd60e51b81526020600482015260096024820152681d1bdbc81cda1bdc9d60ba1b604482015290519081900360640190fd5b60208301615622858563ffffffff615ca316565b915091509250929050565b61563561600e565b6040805160a081018252838152815160608101835260008082526020828101829052845182815280820186529394908501939083019161568b565b61567861600e565b8152602001906001900390816156705790505b509052815260408051600080825260208281019093529190920191906156c7565b6156b461600e565b8152602001906001900390816156ac5790505b50815260006020820152600160409091015292915050565b60006156e961600e565b826000806156f561600e565b60006157018986615592565b90955093506157108986615592565b9095509250600160ff851614156157315761572b89866152e7565b90955091505b61573b8986615cfc565b9095509050600160ff851614156157665784615758848385615d13565b96509650505050505061451d565b846157718483615d97565b9650965050505050509250929050565b600061578b61600e565b828451101580156157a0575060408385510310155b6157dc576040805162461bcd60e51b81526020600482015260086024820152671d1bc81cda1bdc9d60c21b604482015290519081900360640190fd5b6000806157e98686615cfc565b90945091506157f886856155b9565b9094509050836158088383615df2565b9350935050509250929050565b600390565b600c90565b60006060600083905060608660ff1660405190808252806020026020018201604052801561586757816020015b61585461600e565b81526020019060019003908161584c5790505b50905060005b8760ff168160ff1610156158ab5761588587846152e7565b8351849060ff851690811061589657fe5b6020908102919091010152925060010161586d565b5090969095509350505050565b6158c061600e565b6158ca8251615ea3565b61591b576040805162461bcd60e51b815260206004820152601a60248201527f5475706c65206d75737420686176652076616c69642073697a65000000000000604482015290519081900360640190fd5b600160005b83518110156159525783818151811061593557fe5b602002602001015160800151820191508080600101915050615920565b506040805160a08101825260008082528251606081018452818152602081810183905284518381528082018652939490850193919290830191906159ac565b61599961600e565b8152602001906001900390816159915790505b50905281526020810194909452600360408501526060909301525090565b60408051602080820193909352815180820384018152908201909152805191012090565b60006002826040015151106159ff57fe5b604082015151615a6457615a11615270565b8251602080850151604080516001600160f81b031960f896871b8116828601529490951b9093166021850152602280850191909152825180850390910181526042909301909152815191012090506122a0565b615a6c615270565b8260000151615a928460400151600081518110615a8557fe5b6020026020010151615472565b8460200151604051602001808560ff1660ff1660f81b81526001018460ff1660ff1660f81b8152600101838152602001828152602001945050505050604051602081830303815290604052805190602001209050919050565b6000615af5615815565b8383604051602001808460ff1660ff1660f81b8152600101838152602001828152602001935050505060405160208183030381529060405280519060200120905092915050565b615b4461600e565b600882511115615b92576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840e8eae0d8ca40d8cadccee8d60631b604482015290519081900360640190fd5b60608251604051908082528060200260200182016040528015615bbf578160200160208202803883390190505b508051909150600160005b82811015615c2257615be1868281518110615a8557fe5b848281518110615bed57fe5b602002602001018181525050858181518110615c0557fe5b602002602001015160800151820191508080600101915050615bca565b506000835184604051602001808360ff1660ff1660f81b8152600101828051906020019060200280838360005b83811015615c67578181015183820152602001615c4f565b5050505090500192505050604051602081830303815290604052805190602001209050615c948183615df2565b9695505050505050565b606490565b60008160200183511015615cf3576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b60008060208301615622858563ffffffff615ca316565b615d1b61600e565b604080516001808252818301909252606091816020015b615d3a61600e565b815260200190600190039081615d325790505090508281600081518110615d5d57fe5b6020026020010181905250615d8e60405180606001604052808760ff16815260200186815260200183815250615eaa565b95945050505050565b615d9f61600e565b6040805160608101825260ff8516815260208082018590528251600080825291810184526134c493830191615dea565b615dd761600e565b815260200190600190039081615dcf5790505b509052615eaa565b615dfa61600e565b6040805160a0810182528481528151606081018352600080825260208281018290528451828152808201865293949085019390830191615e50565b615e3d61600e565b815260200190600190039081615e355790505b50905281526040805160008082526020828101909352919092019190615e8c565b615e7961600e565b815260200190600190039081615e715790505b508152600260208201526040019290925250919050565b6008101590565b615eb261600e565b6040805160a0810182526000808252602080830186905283518281529081018452919283019190615ef9565b615ee661600e565b815260200190600190039081615ede5790505b50815260016020820181905260409091015292915050565b60405180610120016040528060008152602001600080191681526020016000815260200160008019168152602001600063ffffffff168152602001600081526020016000815260200160008152602001615f69615faa565b905290565b6040518060c001604052806000815260200160008152602001600081526020016000801916815260200160008152602001600080191681525090565b6040805161018081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810182905261010081018290526101208101829052610140810182905261016081019190915290565b6040518060a0016040528060008152602001616028616042565b815260606020820181905260006040830181905291015290565b604080516060808201835260008083526020830152918101919091529056fea265627a7a72315820f5a368acabeb33f1af8537092b65e25780901489798330c0e71a3df9f49bdcca64736f6c63430005110032"

// DeployArbRollup deploys a new Ethereum contract, binding an instance of ArbRollup to it.
func DeployArbRollup(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ArbRollup, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbRollupABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ArbRollupBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ArbRollup{ArbRollupCaller: ArbRollupCaller{contract: contract}, ArbRollupTransactor: ArbRollupTransactor{contract: contract}, ArbRollupFilterer: ArbRollupFilterer{contract: contract}}, nil
}

// ArbRollup is an auto generated Go binding around an Ethereum contract.
type ArbRollup struct {
	ArbRollupCaller     // Read-only binding to the contract
	ArbRollupTransactor // Write-only binding to the contract
	ArbRollupFilterer   // Log filterer for contract events
}

// ArbRollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbRollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbRollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbRollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbRollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbRollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbRollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbRollupSession struct {
	Contract     *ArbRollup        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbRollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbRollupCallerSession struct {
	Contract *ArbRollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ArbRollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbRollupTransactorSession struct {
	Contract     *ArbRollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ArbRollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbRollupRaw struct {
	Contract *ArbRollup // Generic contract binding to access the raw methods on
}

// ArbRollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbRollupCallerRaw struct {
	Contract *ArbRollupCaller // Generic read-only contract binding to access the raw methods on
}

// ArbRollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbRollupTransactorRaw struct {
	Contract *ArbRollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbRollup creates a new instance of ArbRollup, bound to a specific deployed contract.
func NewArbRollup(address common.Address, backend bind.ContractBackend) (*ArbRollup, error) {
	contract, err := bindArbRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbRollup{ArbRollupCaller: ArbRollupCaller{contract: contract}, ArbRollupTransactor: ArbRollupTransactor{contract: contract}, ArbRollupFilterer: ArbRollupFilterer{contract: contract}}, nil
}

// NewArbRollupCaller creates a new read-only instance of ArbRollup, bound to a specific deployed contract.
func NewArbRollupCaller(address common.Address, caller bind.ContractCaller) (*ArbRollupCaller, error) {
	contract, err := bindArbRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbRollupCaller{contract: contract}, nil
}

// NewArbRollupTransactor creates a new write-only instance of ArbRollup, bound to a specific deployed contract.
func NewArbRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbRollupTransactor, error) {
	contract, err := bindArbRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbRollupTransactor{contract: contract}, nil
}

// NewArbRollupFilterer creates a new log filterer instance of ArbRollup, bound to a specific deployed contract.
func NewArbRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbRollupFilterer, error) {
	contract, err := bindArbRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbRollupFilterer{contract: contract}, nil
}

// bindArbRollup binds a generic wrapper to an already deployed contract.
func bindArbRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ArbRollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbRollup *ArbRollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbRollup.Contract.ArbRollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbRollup *ArbRollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbRollup.Contract.ArbRollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbRollup *ArbRollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbRollup.Contract.ArbRollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbRollup *ArbRollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbRollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbRollup *ArbRollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbRollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbRollup *ArbRollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbRollup.Contract.contract.Transact(opts, method, params...)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ArbRollup *ArbRollupCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ArbRollup *ArbRollupSession) VERSION() (string, error) {
	return _ArbRollup.Contract.VERSION(&_ArbRollup.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_ArbRollup *ArbRollupCallerSession) VERSION() (string, error) {
	return _ArbRollup.Contract.VERSION(&_ArbRollup.CallOpts)
}

// AllowedStakers is a free data retrieval call binding the contract method 0xaa0ed4cb.
//
// Solidity: function allowedStakers(address ) view returns(bool)
func (_ArbRollup *ArbRollupCaller) AllowedStakers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "allowedStakers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedStakers is a free data retrieval call binding the contract method 0xaa0ed4cb.
//
// Solidity: function allowedStakers(address ) view returns(bool)
func (_ArbRollup *ArbRollupSession) AllowedStakers(arg0 common.Address) (bool, error) {
	return _ArbRollup.Contract.AllowedStakers(&_ArbRollup.CallOpts, arg0)
}

// AllowedStakers is a free data retrieval call binding the contract method 0xaa0ed4cb.
//
// Solidity: function allowedStakers(address ) view returns(bool)
func (_ArbRollup *ArbRollupCallerSession) AllowedStakers(arg0 common.Address) (bool, error) {
	return _ArbRollup.Contract.AllowedStakers(&_ArbRollup.CallOpts, arg0)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_ArbRollup *ArbRollupCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_ArbRollup *ArbRollupSession) ChallengeFactory() (common.Address, error) {
	return _ArbRollup.Contract.ChallengeFactory(&_ArbRollup.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_ArbRollup *ArbRollupCallerSession) ChallengeFactory() (common.Address, error) {
	return _ArbRollup.Contract.ChallengeFactory(&_ArbRollup.CallOpts)
}

// Challenges is a free data retrieval call binding the contract method 0x08b43a19.
//
// Solidity: function challenges(address ) view returns(bool)
func (_ArbRollup *ArbRollupCaller) Challenges(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "challenges", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Challenges is a free data retrieval call binding the contract method 0x08b43a19.
//
// Solidity: function challenges(address ) view returns(bool)
func (_ArbRollup *ArbRollupSession) Challenges(arg0 common.Address) (bool, error) {
	return _ArbRollup.Contract.Challenges(&_ArbRollup.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x08b43a19.
//
// Solidity: function challenges(address ) view returns(bool)
func (_ArbRollup *ArbRollupCallerSession) Challenges(arg0 common.Address) (bool, error) {
	return _ArbRollup.Contract.Challenges(&_ArbRollup.CallOpts, arg0)
}

// EnableStakerAllowList is a free data retrieval call binding the contract method 0x146b459d.
//
// Solidity: function enableStakerAllowList() view returns(bool)
func (_ArbRollup *ArbRollupCaller) EnableStakerAllowList(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "enableStakerAllowList")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EnableStakerAllowList is a free data retrieval call binding the contract method 0x146b459d.
//
// Solidity: function enableStakerAllowList() view returns(bool)
func (_ArbRollup *ArbRollupSession) EnableStakerAllowList() (bool, error) {
	return _ArbRollup.Contract.EnableStakerAllowList(&_ArbRollup.CallOpts)
}

// EnableStakerAllowList is a free data retrieval call binding the contract method 0x146b459d.
//
// Solidity: function enableStakerAllowList() view returns(bool)
func (_ArbRollup *ArbRollupCallerSession) EnableStakerAllowList() (bool, error) {
	return _ArbRollup.Contract.EnableStakerAllowList(&_ArbRollup.CallOpts)
}

// GetStakeRequired is a free data retrieval call binding the contract method 0xd16c305d.
//
// Solidity: function getStakeRequired() view returns(uint128)
func (_ArbRollup *ArbRollupCaller) GetStakeRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "getStakeRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeRequired is a free data retrieval call binding the contract method 0xd16c305d.
//
// Solidity: function getStakeRequired() view returns(uint128)
func (_ArbRollup *ArbRollupSession) GetStakeRequired() (*big.Int, error) {
	return _ArbRollup.Contract.GetStakeRequired(&_ArbRollup.CallOpts)
}

// GetStakeRequired is a free data retrieval call binding the contract method 0xd16c305d.
//
// Solidity: function getStakeRequired() view returns(uint128)
func (_ArbRollup *ArbRollupCallerSession) GetStakeRequired() (*big.Int, error) {
	return _ArbRollup.Contract.GetStakeRequired(&_ArbRollup.CallOpts)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x500a1564.
//
// Solidity: function getStakeToken() view returns(address)
func (_ArbRollup *ArbRollupCaller) GetStakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "getStakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakeToken is a free data retrieval call binding the contract method 0x500a1564.
//
// Solidity: function getStakeToken() view returns(address)
func (_ArbRollup *ArbRollupSession) GetStakeToken() (common.Address, error) {
	return _ArbRollup.Contract.GetStakeToken(&_ArbRollup.CallOpts)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x500a1564.
//
// Solidity: function getStakeToken() view returns(address)
func (_ArbRollup *ArbRollupCallerSession) GetStakeToken() (common.Address, error) {
	return _ArbRollup.Contract.GetStakeToken(&_ArbRollup.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() view returns(address)
func (_ArbRollup *ArbRollupCaller) GlobalInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "globalInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() view returns(address)
func (_ArbRollup *ArbRollupSession) GlobalInbox() (common.Address, error) {
	return _ArbRollup.Contract.GlobalInbox(&_ArbRollup.CallOpts)
}

// GlobalInbox is a free data retrieval call binding the contract method 0xd489113a.
//
// Solidity: function globalInbox() view returns(address)
func (_ArbRollup *ArbRollupCallerSession) GlobalInbox() (common.Address, error) {
	return _ArbRollup.Contract.GlobalInbox(&_ArbRollup.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ArbRollup *ArbRollupCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ArbRollup *ArbRollupSession) IsMaster() (bool, error) {
	return _ArbRollup.Contract.IsMaster(&_ArbRollup.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_ArbRollup *ArbRollupCallerSession) IsMaster() (bool, error) {
	return _ArbRollup.Contract.IsMaster(&_ArbRollup.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address _stakerAddress) view returns(bool)
func (_ArbRollup *ArbRollupCaller) IsStaked(opts *bind.CallOpts, _stakerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "isStaked", _stakerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address _stakerAddress) view returns(bool)
func (_ArbRollup *ArbRollupSession) IsStaked(_stakerAddress common.Address) (bool, error) {
	return _ArbRollup.Contract.IsStaked(&_ArbRollup.CallOpts, _stakerAddress)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address _stakerAddress) view returns(bool)
func (_ArbRollup *ArbRollupCallerSession) IsStaked(_stakerAddress common.Address) (bool, error) {
	return _ArbRollup.Contract.IsStaked(&_ArbRollup.CallOpts, _stakerAddress)
}

// IsValidLeaf is a free data retrieval call binding the contract method 0x57ca6d1b.
//
// Solidity: function isValidLeaf(bytes32 leaf) view returns(bool)
func (_ArbRollup *ArbRollupCaller) IsValidLeaf(opts *bind.CallOpts, leaf [32]byte) (bool, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "isValidLeaf", leaf)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidLeaf is a free data retrieval call binding the contract method 0x57ca6d1b.
//
// Solidity: function isValidLeaf(bytes32 leaf) view returns(bool)
func (_ArbRollup *ArbRollupSession) IsValidLeaf(leaf [32]byte) (bool, error) {
	return _ArbRollup.Contract.IsValidLeaf(&_ArbRollup.CallOpts, leaf)
}

// IsValidLeaf is a free data retrieval call binding the contract method 0x57ca6d1b.
//
// Solidity: function isValidLeaf(bytes32 leaf) view returns(bool)
func (_ArbRollup *ArbRollupCallerSession) IsValidLeaf(leaf [32]byte) (bool, error) {
	return _ArbRollup.Contract.IsValidLeaf(&_ArbRollup.CallOpts, leaf)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(bytes32)
func (_ArbRollup *ArbRollupCaller) LatestConfirmed(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(bytes32)
func (_ArbRollup *ArbRollupSession) LatestConfirmed() ([32]byte, error) {
	return _ArbRollup.Contract.LatestConfirmed(&_ArbRollup.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(bytes32)
func (_ArbRollup *ArbRollupCallerSession) LatestConfirmed() ([32]byte, error) {
	return _ArbRollup.Contract.LatestConfirmed(&_ArbRollup.CallOpts)
}

// LatestConfirmedPriv is a free data retrieval call binding the contract method 0x3a218e98.
//
// Solidity: function latestConfirmedPriv() view returns(bytes32)
func (_ArbRollup *ArbRollupCaller) LatestConfirmedPriv(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "latestConfirmedPriv")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestConfirmedPriv is a free data retrieval call binding the contract method 0x3a218e98.
//
// Solidity: function latestConfirmedPriv() view returns(bytes32)
func (_ArbRollup *ArbRollupSession) LatestConfirmedPriv() ([32]byte, error) {
	return _ArbRollup.Contract.LatestConfirmedPriv(&_ArbRollup.CallOpts)
}

// LatestConfirmedPriv is a free data retrieval call binding the contract method 0x3a218e98.
//
// Solidity: function latestConfirmedPriv() view returns(bytes32)
func (_ArbRollup *ArbRollupCallerSession) LatestConfirmedPriv() ([32]byte, error) {
	return _ArbRollup.Contract.LatestConfirmedPriv(&_ArbRollup.CallOpts)
}

// Leaves is a free data retrieval call binding the contract method 0x151bcd2c.
//
// Solidity: function leaves(bytes32 ) view returns(bool)
func (_ArbRollup *ArbRollupCaller) Leaves(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "leaves", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Leaves is a free data retrieval call binding the contract method 0x151bcd2c.
//
// Solidity: function leaves(bytes32 ) view returns(bool)
func (_ArbRollup *ArbRollupSession) Leaves(arg0 [32]byte) (bool, error) {
	return _ArbRollup.Contract.Leaves(&_ArbRollup.CallOpts, arg0)
}

// Leaves is a free data retrieval call binding the contract method 0x151bcd2c.
//
// Solidity: function leaves(bytes32 ) view returns(bool)
func (_ArbRollup *ArbRollupCallerSession) Leaves(arg0 [32]byte) (bool, error) {
	return _ArbRollup.Contract.Leaves(&_ArbRollup.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArbRollup *ArbRollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArbRollup *ArbRollupSession) Owner() (common.Address, error) {
	return _ArbRollup.Contract.Owner(&_ArbRollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArbRollup *ArbRollupCallerSession) Owner() (common.Address, error) {
	return _ArbRollup.Contract.Owner(&_ArbRollup.CallOpts)
}

// StakeRequirement is a free data retrieval call binding the contract method 0xb6f9bbb9.
//
// Solidity: function stakeRequirement() view returns(uint128)
func (_ArbRollup *ArbRollupCaller) StakeRequirement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "stakeRequirement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeRequirement is a free data retrieval call binding the contract method 0xb6f9bbb9.
//
// Solidity: function stakeRequirement() view returns(uint128)
func (_ArbRollup *ArbRollupSession) StakeRequirement() (*big.Int, error) {
	return _ArbRollup.Contract.StakeRequirement(&_ArbRollup.CallOpts)
}

// StakeRequirement is a free data retrieval call binding the contract method 0xb6f9bbb9.
//
// Solidity: function stakeRequirement() view returns(uint128)
func (_ArbRollup *ArbRollupCallerSession) StakeRequirement() (*big.Int, error) {
	return _ArbRollup.Contract.StakeRequirement(&_ArbRollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ArbRollup *ArbRollupCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ArbRollup *ArbRollupSession) StakeToken() (common.Address, error) {
	return _ArbRollup.Contract.StakeToken(&_ArbRollup.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_ArbRollup *ArbRollupCallerSession) StakeToken() (common.Address, error) {
	return _ArbRollup.Contract.StakeToken(&_ArbRollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_ArbRollup *ArbRollupCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_ArbRollup *ArbRollupSession) StakerCount() (*big.Int, error) {
	return _ArbRollup.Contract.StakerCount(&_ArbRollup.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_ArbRollup *ArbRollupCallerSession) StakerCount() (*big.Int, error) {
	return _ArbRollup.Contract.StakerCount(&_ArbRollup.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bytes32 location, uint128 creationTimeBlocks, bool inChallenge)
func (_ArbRollup *ArbRollupCaller) Stakers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Location           [32]byte
	CreationTimeBlocks *big.Int
	InChallenge        bool
}, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		Location           [32]byte
		CreationTimeBlocks *big.Int
		InChallenge        bool
	})

	outstruct.Location = out[0].([32]byte)
	outstruct.CreationTimeBlocks = out[1].(*big.Int)
	outstruct.InChallenge = out[2].(bool)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bytes32 location, uint128 creationTimeBlocks, bool inChallenge)
func (_ArbRollup *ArbRollupSession) Stakers(arg0 common.Address) (struct {
	Location           [32]byte
	CreationTimeBlocks *big.Int
	InChallenge        bool
}, error) {
	return _ArbRollup.Contract.Stakers(&_ArbRollup.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bytes32 location, uint128 creationTimeBlocks, bool inChallenge)
func (_ArbRollup *ArbRollupCallerSession) Stakers(arg0 common.Address) (struct {
	Location           [32]byte
	CreationTimeBlocks *big.Int
	InChallenge        bool
}, error) {
	return _ArbRollup.Contract.Stakers(&_ArbRollup.CallOpts, arg0)
}

// VmParams is a free data retrieval call binding the contract method 0xbbc2cc00.
//
// Solidity: function vmParams() view returns(uint256 gracePeriodTicks, uint256 arbGasSpeedLimitPerTick, uint64 maxExecutionSteps)
func (_ArbRollup *ArbRollupCaller) VmParams(opts *bind.CallOpts) (struct {
	GracePeriodTicks        *big.Int
	ArbGasSpeedLimitPerTick *big.Int
	MaxExecutionSteps       uint64
}, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "vmParams")

	outstruct := new(struct {
		GracePeriodTicks        *big.Int
		ArbGasSpeedLimitPerTick *big.Int
		MaxExecutionSteps       uint64
	})

	outstruct.GracePeriodTicks = out[0].(*big.Int)
	outstruct.ArbGasSpeedLimitPerTick = out[1].(*big.Int)
	outstruct.MaxExecutionSteps = out[2].(uint64)

	return *outstruct, err

}

// VmParams is a free data retrieval call binding the contract method 0xbbc2cc00.
//
// Solidity: function vmParams() view returns(uint256 gracePeriodTicks, uint256 arbGasSpeedLimitPerTick, uint64 maxExecutionSteps)
func (_ArbRollup *ArbRollupSession) VmParams() (struct {
	GracePeriodTicks        *big.Int
	ArbGasSpeedLimitPerTick *big.Int
	MaxExecutionSteps       uint64
}, error) {
	return _ArbRollup.Contract.VmParams(&_ArbRollup.CallOpts)
}

// VmParams is a free data retrieval call binding the contract method 0xbbc2cc00.
//
// Solidity: function vmParams() view returns(uint256 gracePeriodTicks, uint256 arbGasSpeedLimitPerTick, uint64 maxExecutionSteps)
func (_ArbRollup *ArbRollupCallerSession) VmParams() (struct {
	GracePeriodTicks        *big.Int
	ArbGasSpeedLimitPerTick *big.Int
	MaxExecutionSteps       uint64
}, error) {
	return _ArbRollup.Contract.VmParams(&_ArbRollup.CallOpts)
}

// WithdrawnStakes is a free data retrieval call binding the contract method 0xeb2e74cb.
//
// Solidity: function withdrawnStakes(address ) view returns(uint256)
func (_ArbRollup *ArbRollupCaller) WithdrawnStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArbRollup.contract.Call(opts, &out, "withdrawnStakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnStakes is a free data retrieval call binding the contract method 0xeb2e74cb.
//
// Solidity: function withdrawnStakes(address ) view returns(uint256)
func (_ArbRollup *ArbRollupSession) WithdrawnStakes(arg0 common.Address) (*big.Int, error) {
	return _ArbRollup.Contract.WithdrawnStakes(&_ArbRollup.CallOpts, arg0)
}

// WithdrawnStakes is a free data retrieval call binding the contract method 0xeb2e74cb.
//
// Solidity: function withdrawnStakes(address ) view returns(uint256)
func (_ArbRollup *ArbRollupCallerSession) WithdrawnStakes(arg0 common.Address) (*big.Int, error) {
	return _ArbRollup.Contract.WithdrawnStakes(&_ArbRollup.CallOpts, arg0)
}

// Confirm is a paid mutator transaction binding the contract method 0x6e64beaa.
//
// Solidity: function confirm(bytes32 initalProtoStateHash, uint256 beforeSendCount, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages, address[] stakerAddresses, bytes32[] stakerProofs, uint256[] stakerProofOffsets) returns()
func (_ArbRollup *ArbRollupTransactor) Confirm(opts *bind.TransactOpts, initalProtoStateHash [32]byte, beforeSendCount *big.Int, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte, stakerAddresses []common.Address, stakerProofs [][32]byte, stakerProofOffsets []*big.Int) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "confirm", initalProtoStateHash, beforeSendCount, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages, stakerAddresses, stakerProofs, stakerProofOffsets)
}

// Confirm is a paid mutator transaction binding the contract method 0x6e64beaa.
//
// Solidity: function confirm(bytes32 initalProtoStateHash, uint256 beforeSendCount, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages, address[] stakerAddresses, bytes32[] stakerProofs, uint256[] stakerProofOffsets) returns()
func (_ArbRollup *ArbRollupSession) Confirm(initalProtoStateHash [32]byte, beforeSendCount *big.Int, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte, stakerAddresses []common.Address, stakerProofs [][32]byte, stakerProofOffsets []*big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.Confirm(&_ArbRollup.TransactOpts, initalProtoStateHash, beforeSendCount, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages, stakerAddresses, stakerProofs, stakerProofOffsets)
}

// Confirm is a paid mutator transaction binding the contract method 0x6e64beaa.
//
// Solidity: function confirm(bytes32 initalProtoStateHash, uint256 beforeSendCount, uint256[] branches, uint256[] deadlineTicks, bytes32[] challengeNodeData, bytes32[] logsAcc, bytes32[] vmProtoStateHashes, uint256[] messageCounts, bytes messages, address[] stakerAddresses, bytes32[] stakerProofs, uint256[] stakerProofOffsets) returns()
func (_ArbRollup *ArbRollupTransactorSession) Confirm(initalProtoStateHash [32]byte, beforeSendCount *big.Int, branches []*big.Int, deadlineTicks []*big.Int, challengeNodeData [][32]byte, logsAcc [][32]byte, vmProtoStateHashes [][32]byte, messageCounts []*big.Int, messages []byte, stakerAddresses []common.Address, stakerProofs [][32]byte, stakerProofOffsets []*big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.Confirm(&_ArbRollup.TransactOpts, initalProtoStateHash, beforeSendCount, branches, deadlineTicks, challengeNodeData, logsAcc, vmProtoStateHashes, messageCounts, messages, stakerAddresses, stakerProofs, stakerProofOffsets)
}

// GetWithdrawnStake is a paid mutator transaction binding the contract method 0x3fcc045a.
//
// Solidity: function getWithdrawnStake(address _staker) returns()
func (_ArbRollup *ArbRollupTransactor) GetWithdrawnStake(opts *bind.TransactOpts, _staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "getWithdrawnStake", _staker)
}

// GetWithdrawnStake is a paid mutator transaction binding the contract method 0x3fcc045a.
//
// Solidity: function getWithdrawnStake(address _staker) returns()
func (_ArbRollup *ArbRollupSession) GetWithdrawnStake(_staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.GetWithdrawnStake(&_ArbRollup.TransactOpts, _staker)
}

// GetWithdrawnStake is a paid mutator transaction binding the contract method 0x3fcc045a.
//
// Solidity: function getWithdrawnStake(address _staker) returns()
func (_ArbRollup *ArbRollupTransactorSession) GetWithdrawnStake(_staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.GetWithdrawnStake(&_ArbRollup.TransactOpts, _staker)
}

// Init is a paid mutator transaction binding the contract method 0x1dceffe7.
//
// Solidity: function init(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint128 _stakeRequirement, address _stakeToken, address _owner, address _challengeFactoryAddress, address _globalInboxAddress, bytes _extraConfig) returns()
func (_ArbRollup *ArbRollupTransactor) Init(opts *bind.TransactOpts, _vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _stakeRequirement *big.Int, _stakeToken common.Address, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "init", _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _stakeRequirement, _stakeToken, _owner, _challengeFactoryAddress, _globalInboxAddress, _extraConfig)
}

// Init is a paid mutator transaction binding the contract method 0x1dceffe7.
//
// Solidity: function init(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint128 _stakeRequirement, address _stakeToken, address _owner, address _challengeFactoryAddress, address _globalInboxAddress, bytes _extraConfig) returns()
func (_ArbRollup *ArbRollupSession) Init(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _stakeRequirement *big.Int, _stakeToken common.Address, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.Init(&_ArbRollup.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _stakeRequirement, _stakeToken, _owner, _challengeFactoryAddress, _globalInboxAddress, _extraConfig)
}

// Init is a paid mutator transaction binding the contract method 0x1dceffe7.
//
// Solidity: function init(bytes32 _vmState, uint128 _gracePeriodTicks, uint128 _arbGasSpeedLimitPerTick, uint64 _maxExecutionSteps, uint128 _stakeRequirement, address _stakeToken, address _owner, address _challengeFactoryAddress, address _globalInboxAddress, bytes _extraConfig) returns()
func (_ArbRollup *ArbRollupTransactorSession) Init(_vmState [32]byte, _gracePeriodTicks *big.Int, _arbGasSpeedLimitPerTick *big.Int, _maxExecutionSteps uint64, _stakeRequirement *big.Int, _stakeToken common.Address, _owner common.Address, _challengeFactoryAddress common.Address, _globalInboxAddress common.Address, _extraConfig []byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.Init(&_ArbRollup.TransactOpts, _vmState, _gracePeriodTicks, _arbGasSpeedLimitPerTick, _maxExecutionSteps, _stakeRequirement, _stakeToken, _owner, _challengeFactoryAddress, _globalInboxAddress, _extraConfig)
}

// MakeAssertion is a paid mutator transaction binding the contract method 0xa0e9f382.
//
// Solidity: function makeAssertion(bytes32[8] fields, uint256[5] fields2, bytes32 validBlockHashPrecondition, uint256 validBlockHeightPrecondition, uint64 messageCount, uint64 logCount, uint32 prevChildType, uint64 numSteps, uint64 numArbGas, bytes32[] stakerProof) returns()
func (_ArbRollup *ArbRollupTransactor) MakeAssertion(opts *bind.TransactOpts, fields [8][32]byte, fields2 [5]*big.Int, validBlockHashPrecondition [32]byte, validBlockHeightPrecondition *big.Int, messageCount uint64, logCount uint64, prevChildType uint32, numSteps uint64, numArbGas uint64, stakerProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "makeAssertion", fields, fields2, validBlockHashPrecondition, validBlockHeightPrecondition, messageCount, logCount, prevChildType, numSteps, numArbGas, stakerProof)
}

// MakeAssertion is a paid mutator transaction binding the contract method 0xa0e9f382.
//
// Solidity: function makeAssertion(bytes32[8] fields, uint256[5] fields2, bytes32 validBlockHashPrecondition, uint256 validBlockHeightPrecondition, uint64 messageCount, uint64 logCount, uint32 prevChildType, uint64 numSteps, uint64 numArbGas, bytes32[] stakerProof) returns()
func (_ArbRollup *ArbRollupSession) MakeAssertion(fields [8][32]byte, fields2 [5]*big.Int, validBlockHashPrecondition [32]byte, validBlockHeightPrecondition *big.Int, messageCount uint64, logCount uint64, prevChildType uint32, numSteps uint64, numArbGas uint64, stakerProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.MakeAssertion(&_ArbRollup.TransactOpts, fields, fields2, validBlockHashPrecondition, validBlockHeightPrecondition, messageCount, logCount, prevChildType, numSteps, numArbGas, stakerProof)
}

// MakeAssertion is a paid mutator transaction binding the contract method 0xa0e9f382.
//
// Solidity: function makeAssertion(bytes32[8] fields, uint256[5] fields2, bytes32 validBlockHashPrecondition, uint256 validBlockHeightPrecondition, uint64 messageCount, uint64 logCount, uint32 prevChildType, uint64 numSteps, uint64 numArbGas, bytes32[] stakerProof) returns()
func (_ArbRollup *ArbRollupTransactorSession) MakeAssertion(fields [8][32]byte, fields2 [5]*big.Int, validBlockHashPrecondition [32]byte, validBlockHeightPrecondition *big.Int, messageCount uint64, logCount uint64, prevChildType uint32, numSteps uint64, numArbGas uint64, stakerProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.MakeAssertion(&_ArbRollup.TransactOpts, fields, fields2, validBlockHashPrecondition, validBlockHeightPrecondition, messageCount, logCount, prevChildType, numSteps, numArbGas, stakerProof)
}

// MoveStake is a paid mutator transaction binding the contract method 0xdbad0a39.
//
// Solidity: function moveStake(bytes32[] proof1, bytes32[] proof2) returns()
func (_ArbRollup *ArbRollupTransactor) MoveStake(opts *bind.TransactOpts, proof1 [][32]byte, proof2 [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "moveStake", proof1, proof2)
}

// MoveStake is a paid mutator transaction binding the contract method 0xdbad0a39.
//
// Solidity: function moveStake(bytes32[] proof1, bytes32[] proof2) returns()
func (_ArbRollup *ArbRollupSession) MoveStake(proof1 [][32]byte, proof2 [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.MoveStake(&_ArbRollup.TransactOpts, proof1, proof2)
}

// MoveStake is a paid mutator transaction binding the contract method 0xdbad0a39.
//
// Solidity: function moveStake(bytes32[] proof1, bytes32[] proof2) returns()
func (_ArbRollup *ArbRollupTransactorSession) MoveStake(proof1 [][32]byte, proof2 [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.MoveStake(&_ArbRollup.TransactOpts, proof1, proof2)
}

// OwnerAddAllowedStaker is a paid mutator transaction binding the contract method 0x3d3f66d1.
//
// Solidity: function ownerAddAllowedStaker(address staker) returns()
func (_ArbRollup *ArbRollupTransactor) OwnerAddAllowedStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "ownerAddAllowedStaker", staker)
}

// OwnerAddAllowedStaker is a paid mutator transaction binding the contract method 0x3d3f66d1.
//
// Solidity: function ownerAddAllowedStaker(address staker) returns()
func (_ArbRollup *ArbRollupSession) OwnerAddAllowedStaker(staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerAddAllowedStaker(&_ArbRollup.TransactOpts, staker)
}

// OwnerAddAllowedStaker is a paid mutator transaction binding the contract method 0x3d3f66d1.
//
// Solidity: function ownerAddAllowedStaker(address staker) returns()
func (_ArbRollup *ArbRollupTransactorSession) OwnerAddAllowedStaker(staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerAddAllowedStaker(&_ArbRollup.TransactOpts, staker)
}

// OwnerConfirm is a paid mutator transaction binding the contract method 0x10e7e692.
//
// Solidity: function ownerConfirm(bytes32[] logsAcc, bytes32[] validNodeHashes, bytes32 finalNodeHash) returns()
func (_ArbRollup *ArbRollupTransactor) OwnerConfirm(opts *bind.TransactOpts, logsAcc [][32]byte, validNodeHashes [][32]byte, finalNodeHash [32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "ownerConfirm", logsAcc, validNodeHashes, finalNodeHash)
}

// OwnerConfirm is a paid mutator transaction binding the contract method 0x10e7e692.
//
// Solidity: function ownerConfirm(bytes32[] logsAcc, bytes32[] validNodeHashes, bytes32 finalNodeHash) returns()
func (_ArbRollup *ArbRollupSession) OwnerConfirm(logsAcc [][32]byte, validNodeHashes [][32]byte, finalNodeHash [32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerConfirm(&_ArbRollup.TransactOpts, logsAcc, validNodeHashes, finalNodeHash)
}

// OwnerConfirm is a paid mutator transaction binding the contract method 0x10e7e692.
//
// Solidity: function ownerConfirm(bytes32[] logsAcc, bytes32[] validNodeHashes, bytes32 finalNodeHash) returns()
func (_ArbRollup *ArbRollupTransactorSession) OwnerConfirm(logsAcc [][32]byte, validNodeHashes [][32]byte, finalNodeHash [32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerConfirm(&_ArbRollup.TransactOpts, logsAcc, validNodeHashes, finalNodeHash)
}

// OwnerRefundStaker is a paid mutator transaction binding the contract method 0x9b6c85fd.
//
// Solidity: function ownerRefundStaker(address staker) returns()
func (_ArbRollup *ArbRollupTransactor) OwnerRefundStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "ownerRefundStaker", staker)
}

// OwnerRefundStaker is a paid mutator transaction binding the contract method 0x9b6c85fd.
//
// Solidity: function ownerRefundStaker(address staker) returns()
func (_ArbRollup *ArbRollupSession) OwnerRefundStaker(staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerRefundStaker(&_ArbRollup.TransactOpts, staker)
}

// OwnerRefundStaker is a paid mutator transaction binding the contract method 0x9b6c85fd.
//
// Solidity: function ownerRefundStaker(address staker) returns()
func (_ArbRollup *ArbRollupTransactorSession) OwnerRefundStaker(staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerRefundStaker(&_ArbRollup.TransactOpts, staker)
}

// OwnerRemoveAllowedStaker is a paid mutator transaction binding the contract method 0xeb908433.
//
// Solidity: function ownerRemoveAllowedStaker(address staker) returns()
func (_ArbRollup *ArbRollupTransactor) OwnerRemoveAllowedStaker(opts *bind.TransactOpts, staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "ownerRemoveAllowedStaker", staker)
}

// OwnerRemoveAllowedStaker is a paid mutator transaction binding the contract method 0xeb908433.
//
// Solidity: function ownerRemoveAllowedStaker(address staker) returns()
func (_ArbRollup *ArbRollupSession) OwnerRemoveAllowedStaker(staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerRemoveAllowedStaker(&_ArbRollup.TransactOpts, staker)
}

// OwnerRemoveAllowedStaker is a paid mutator transaction binding the contract method 0xeb908433.
//
// Solidity: function ownerRemoveAllowedStaker(address staker) returns()
func (_ArbRollup *ArbRollupTransactorSession) OwnerRemoveAllowedStaker(staker common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerRemoveAllowedStaker(&_ArbRollup.TransactOpts, staker)
}

// OwnerSendMessages is a paid mutator transaction binding the contract method 0x56373597.
//
// Solidity: function ownerSendMessages(bytes messages, uint256 initialMaxSendCount, uint256 finalMaxSendCount) returns()
func (_ArbRollup *ArbRollupTransactor) OwnerSendMessages(opts *bind.TransactOpts, messages []byte, initialMaxSendCount *big.Int, finalMaxSendCount *big.Int) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "ownerSendMessages", messages, initialMaxSendCount, finalMaxSendCount)
}

// OwnerSendMessages is a paid mutator transaction binding the contract method 0x56373597.
//
// Solidity: function ownerSendMessages(bytes messages, uint256 initialMaxSendCount, uint256 finalMaxSendCount) returns()
func (_ArbRollup *ArbRollupSession) OwnerSendMessages(messages []byte, initialMaxSendCount *big.Int, finalMaxSendCount *big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerSendMessages(&_ArbRollup.TransactOpts, messages, initialMaxSendCount, finalMaxSendCount)
}

// OwnerSendMessages is a paid mutator transaction binding the contract method 0x56373597.
//
// Solidity: function ownerSendMessages(bytes messages, uint256 initialMaxSendCount, uint256 finalMaxSendCount) returns()
func (_ArbRollup *ArbRollupTransactorSession) OwnerSendMessages(messages []byte, initialMaxSendCount *big.Int, finalMaxSendCount *big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerSendMessages(&_ArbRollup.TransactOpts, messages, initialMaxSendCount, finalMaxSendCount)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbRollup *ArbRollupTransactor) OwnerShutdown(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "ownerShutdown")
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbRollup *ArbRollupSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerShutdown(&_ArbRollup.TransactOpts)
}

// OwnerShutdown is a paid mutator transaction binding the contract method 0xcfa80707.
//
// Solidity: function ownerShutdown() returns()
func (_ArbRollup *ArbRollupTransactorSession) OwnerShutdown() (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerShutdown(&_ArbRollup.TransactOpts)
}

// OwnerToggleStakerAllowListed is a paid mutator transaction binding the contract method 0x27258fb7.
//
// Solidity: function ownerToggleStakerAllowListed(bool shouldRequire) returns()
func (_ArbRollup *ArbRollupTransactor) OwnerToggleStakerAllowListed(opts *bind.TransactOpts, shouldRequire bool) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "ownerToggleStakerAllowListed", shouldRequire)
}

// OwnerToggleStakerAllowListed is a paid mutator transaction binding the contract method 0x27258fb7.
//
// Solidity: function ownerToggleStakerAllowListed(bool shouldRequire) returns()
func (_ArbRollup *ArbRollupSession) OwnerToggleStakerAllowListed(shouldRequire bool) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerToggleStakerAllowListed(&_ArbRollup.TransactOpts, shouldRequire)
}

// OwnerToggleStakerAllowListed is a paid mutator transaction binding the contract method 0x27258fb7.
//
// Solidity: function ownerToggleStakerAllowListed(bool shouldRequire) returns()
func (_ArbRollup *ArbRollupTransactorSession) OwnerToggleStakerAllowListed(shouldRequire bool) (*types.Transaction, error) {
	return _ArbRollup.Contract.OwnerToggleStakerAllowListed(&_ArbRollup.TransactOpts, shouldRequire)
}

// PlaceStake is a paid mutator transaction binding the contract method 0xe0620d64.
//
// Solidity: function placeStake(bytes32[] proof1, bytes32[] proof2) payable returns()
func (_ArbRollup *ArbRollupTransactor) PlaceStake(opts *bind.TransactOpts, proof1 [][32]byte, proof2 [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "placeStake", proof1, proof2)
}

// PlaceStake is a paid mutator transaction binding the contract method 0xe0620d64.
//
// Solidity: function placeStake(bytes32[] proof1, bytes32[] proof2) payable returns()
func (_ArbRollup *ArbRollupSession) PlaceStake(proof1 [][32]byte, proof2 [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.PlaceStake(&_ArbRollup.TransactOpts, proof1, proof2)
}

// PlaceStake is a paid mutator transaction binding the contract method 0xe0620d64.
//
// Solidity: function placeStake(bytes32[] proof1, bytes32[] proof2) payable returns()
func (_ArbRollup *ArbRollupTransactorSession) PlaceStake(proof1 [][32]byte, proof2 [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.PlaceStake(&_ArbRollup.TransactOpts, proof1, proof2)
}

// PruneLeaves is a paid mutator transaction binding the contract method 0xfcfd8d3f.
//
// Solidity: function pruneLeaves(bytes32[] fromNodes, bytes32[] leafProofs, uint256[] leafProofLengths, bytes32[] latestConfProofs, uint256[] latestConfirmedProofLengths) returns()
func (_ArbRollup *ArbRollupTransactor) PruneLeaves(opts *bind.TransactOpts, fromNodes [][32]byte, leafProofs [][32]byte, leafProofLengths []*big.Int, latestConfProofs [][32]byte, latestConfirmedProofLengths []*big.Int) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "pruneLeaves", fromNodes, leafProofs, leafProofLengths, latestConfProofs, latestConfirmedProofLengths)
}

// PruneLeaves is a paid mutator transaction binding the contract method 0xfcfd8d3f.
//
// Solidity: function pruneLeaves(bytes32[] fromNodes, bytes32[] leafProofs, uint256[] leafProofLengths, bytes32[] latestConfProofs, uint256[] latestConfirmedProofLengths) returns()
func (_ArbRollup *ArbRollupSession) PruneLeaves(fromNodes [][32]byte, leafProofs [][32]byte, leafProofLengths []*big.Int, latestConfProofs [][32]byte, latestConfirmedProofLengths []*big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.PruneLeaves(&_ArbRollup.TransactOpts, fromNodes, leafProofs, leafProofLengths, latestConfProofs, latestConfirmedProofLengths)
}

// PruneLeaves is a paid mutator transaction binding the contract method 0xfcfd8d3f.
//
// Solidity: function pruneLeaves(bytes32[] fromNodes, bytes32[] leafProofs, uint256[] leafProofLengths, bytes32[] latestConfProofs, uint256[] latestConfirmedProofLengths) returns()
func (_ArbRollup *ArbRollupTransactorSession) PruneLeaves(fromNodes [][32]byte, leafProofs [][32]byte, leafProofLengths []*big.Int, latestConfProofs [][32]byte, latestConfirmedProofLengths []*big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.PruneLeaves(&_ArbRollup.TransactOpts, fromNodes, leafProofs, leafProofLengths, latestConfProofs, latestConfirmedProofLengths)
}

// RecoverStakeConfirmed is a paid mutator transaction binding the contract method 0x7cfaaf67.
//
// Solidity: function recoverStakeConfirmed(bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactor) RecoverStakeConfirmed(opts *bind.TransactOpts, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "recoverStakeConfirmed", proof)
}

// RecoverStakeConfirmed is a paid mutator transaction binding the contract method 0x7cfaaf67.
//
// Solidity: function recoverStakeConfirmed(bytes32[] proof) returns()
func (_ArbRollup *ArbRollupSession) RecoverStakeConfirmed(proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeConfirmed(&_ArbRollup.TransactOpts, proof)
}

// RecoverStakeConfirmed is a paid mutator transaction binding the contract method 0x7cfaaf67.
//
// Solidity: function recoverStakeConfirmed(bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactorSession) RecoverStakeConfirmed(proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeConfirmed(&_ArbRollup.TransactOpts, proof)
}

// RecoverStakeMooted is a paid mutator transaction binding the contract method 0x33554032.
//
// Solidity: function recoverStakeMooted(address stakerAddress, bytes32 node, bytes32[] latestConfirmedProof, bytes32[] stakerProof) returns()
func (_ArbRollup *ArbRollupTransactor) RecoverStakeMooted(opts *bind.TransactOpts, stakerAddress common.Address, node [32]byte, latestConfirmedProof [][32]byte, stakerProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "recoverStakeMooted", stakerAddress, node, latestConfirmedProof, stakerProof)
}

// RecoverStakeMooted is a paid mutator transaction binding the contract method 0x33554032.
//
// Solidity: function recoverStakeMooted(address stakerAddress, bytes32 node, bytes32[] latestConfirmedProof, bytes32[] stakerProof) returns()
func (_ArbRollup *ArbRollupSession) RecoverStakeMooted(stakerAddress common.Address, node [32]byte, latestConfirmedProof [][32]byte, stakerProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeMooted(&_ArbRollup.TransactOpts, stakerAddress, node, latestConfirmedProof, stakerProof)
}

// RecoverStakeMooted is a paid mutator transaction binding the contract method 0x33554032.
//
// Solidity: function recoverStakeMooted(address stakerAddress, bytes32 node, bytes32[] latestConfirmedProof, bytes32[] stakerProof) returns()
func (_ArbRollup *ArbRollupTransactorSession) RecoverStakeMooted(stakerAddress common.Address, node [32]byte, latestConfirmedProof [][32]byte, stakerProof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeMooted(&_ArbRollup.TransactOpts, stakerAddress, node, latestConfirmedProof, stakerProof)
}

// RecoverStakeOld is a paid mutator transaction binding the contract method 0x113ec9d8.
//
// Solidity: function recoverStakeOld(address stakerAddress, bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactor) RecoverStakeOld(opts *bind.TransactOpts, stakerAddress common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "recoverStakeOld", stakerAddress, proof)
}

// RecoverStakeOld is a paid mutator transaction binding the contract method 0x113ec9d8.
//
// Solidity: function recoverStakeOld(address stakerAddress, bytes32[] proof) returns()
func (_ArbRollup *ArbRollupSession) RecoverStakeOld(stakerAddress common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeOld(&_ArbRollup.TransactOpts, stakerAddress, proof)
}

// RecoverStakeOld is a paid mutator transaction binding the contract method 0x113ec9d8.
//
// Solidity: function recoverStakeOld(address stakerAddress, bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactorSession) RecoverStakeOld(stakerAddress common.Address, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakeOld(&_ArbRollup.TransactOpts, stakerAddress, proof)
}

// RecoverStakePassedDeadline is a paid mutator transaction binding the contract method 0xbadb3f14.
//
// Solidity: function recoverStakePassedDeadline(address stakerAddress, uint256 deadlineTicks, bytes32 disputableNodeHashVal, uint256 childType, bytes32 vmProtoStateHash, bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactor) RecoverStakePassedDeadline(opts *bind.TransactOpts, stakerAddress common.Address, deadlineTicks *big.Int, disputableNodeHashVal [32]byte, childType *big.Int, vmProtoStateHash [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "recoverStakePassedDeadline", stakerAddress, deadlineTicks, disputableNodeHashVal, childType, vmProtoStateHash, proof)
}

// RecoverStakePassedDeadline is a paid mutator transaction binding the contract method 0xbadb3f14.
//
// Solidity: function recoverStakePassedDeadline(address stakerAddress, uint256 deadlineTicks, bytes32 disputableNodeHashVal, uint256 childType, bytes32 vmProtoStateHash, bytes32[] proof) returns()
func (_ArbRollup *ArbRollupSession) RecoverStakePassedDeadline(stakerAddress common.Address, deadlineTicks *big.Int, disputableNodeHashVal [32]byte, childType *big.Int, vmProtoStateHash [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakePassedDeadline(&_ArbRollup.TransactOpts, stakerAddress, deadlineTicks, disputableNodeHashVal, childType, vmProtoStateHash, proof)
}

// RecoverStakePassedDeadline is a paid mutator transaction binding the contract method 0xbadb3f14.
//
// Solidity: function recoverStakePassedDeadline(address stakerAddress, uint256 deadlineTicks, bytes32 disputableNodeHashVal, uint256 childType, bytes32 vmProtoStateHash, bytes32[] proof) returns()
func (_ArbRollup *ArbRollupTransactorSession) RecoverStakePassedDeadline(stakerAddress common.Address, deadlineTicks *big.Int, disputableNodeHashVal [32]byte, childType *big.Int, vmProtoStateHash [32]byte, proof [][32]byte) (*types.Transaction, error) {
	return _ArbRollup.Contract.RecoverStakePassedDeadline(&_ArbRollup.TransactOpts, stakerAddress, deadlineTicks, disputableNodeHashVal, childType, vmProtoStateHash, proof)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_ArbRollup *ArbRollupTransactor) ResolveChallenge(opts *bind.TransactOpts, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "resolveChallenge", winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_ArbRollup *ArbRollupSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.ResolveChallenge(&_ArbRollup.TransactOpts, winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_ArbRollup *ArbRollupTransactorSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _ArbRollup.Contract.ResolveChallenge(&_ArbRollup.TransactOpts, winner, loser)
}

// StartChallenge is a paid mutator transaction binding the contract method 0xbac5963f.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, bytes32 prevNode, uint256 deadlineTicks, uint256[2] stakerNodeTypes, bytes32[2] vmProtoHashes, bytes32[] asserterProof, bytes32[] challengerProof, bytes32 asserterNodeHash, bytes32 challengerDataHash, uint128 challengerPeriodTicks) returns()
func (_ArbRollup *ArbRollupTransactor) StartChallenge(opts *bind.TransactOpts, asserterAddress common.Address, challengerAddress common.Address, prevNode [32]byte, deadlineTicks *big.Int, stakerNodeTypes [2]*big.Int, vmProtoHashes [2][32]byte, asserterProof [][32]byte, challengerProof [][32]byte, asserterNodeHash [32]byte, challengerDataHash [32]byte, challengerPeriodTicks *big.Int) (*types.Transaction, error) {
	return _ArbRollup.contract.Transact(opts, "startChallenge", asserterAddress, challengerAddress, prevNode, deadlineTicks, stakerNodeTypes, vmProtoHashes, asserterProof, challengerProof, asserterNodeHash, challengerDataHash, challengerPeriodTicks)
}

// StartChallenge is a paid mutator transaction binding the contract method 0xbac5963f.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, bytes32 prevNode, uint256 deadlineTicks, uint256[2] stakerNodeTypes, bytes32[2] vmProtoHashes, bytes32[] asserterProof, bytes32[] challengerProof, bytes32 asserterNodeHash, bytes32 challengerDataHash, uint128 challengerPeriodTicks) returns()
func (_ArbRollup *ArbRollupSession) StartChallenge(asserterAddress common.Address, challengerAddress common.Address, prevNode [32]byte, deadlineTicks *big.Int, stakerNodeTypes [2]*big.Int, vmProtoHashes [2][32]byte, asserterProof [][32]byte, challengerProof [][32]byte, asserterNodeHash [32]byte, challengerDataHash [32]byte, challengerPeriodTicks *big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.StartChallenge(&_ArbRollup.TransactOpts, asserterAddress, challengerAddress, prevNode, deadlineTicks, stakerNodeTypes, vmProtoHashes, asserterProof, challengerProof, asserterNodeHash, challengerDataHash, challengerPeriodTicks)
}

// StartChallenge is a paid mutator transaction binding the contract method 0xbac5963f.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, bytes32 prevNode, uint256 deadlineTicks, uint256[2] stakerNodeTypes, bytes32[2] vmProtoHashes, bytes32[] asserterProof, bytes32[] challengerProof, bytes32 asserterNodeHash, bytes32 challengerDataHash, uint128 challengerPeriodTicks) returns()
func (_ArbRollup *ArbRollupTransactorSession) StartChallenge(asserterAddress common.Address, challengerAddress common.Address, prevNode [32]byte, deadlineTicks *big.Int, stakerNodeTypes [2]*big.Int, vmProtoHashes [2][32]byte, asserterProof [][32]byte, challengerProof [][32]byte, asserterNodeHash [32]byte, challengerDataHash [32]byte, challengerPeriodTicks *big.Int) (*types.Transaction, error) {
	return _ArbRollup.Contract.StartChallenge(&_ArbRollup.TransactOpts, asserterAddress, challengerAddress, prevNode, deadlineTicks, stakerNodeTypes, vmProtoHashes, asserterProof, challengerProof, asserterNodeHash, challengerDataHash, challengerPeriodTicks)
}

// ArbRollupConfirmedAssertionIterator is returned from FilterConfirmedAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedAssertion events raised by the ArbRollup contract.
type ArbRollupConfirmedAssertionIterator struct {
	Event *ArbRollupConfirmedAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupConfirmedAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupConfirmedAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupConfirmedAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupConfirmedAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupConfirmedAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupConfirmedAssertion represents a ConfirmedAssertion event raised by the ArbRollup contract.
type ArbRollupConfirmedAssertion struct {
	LogsAccHash [][32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConfirmedAssertion is a free log retrieval operation binding the contract event 0xded5fa103431438087188a5f8c6a4c3ea90996bbd63be7b1b3fa0a425b37fdd5.
//
// Solidity: event ConfirmedAssertion(bytes32[] logsAccHash)
func (_ArbRollup *ArbRollupFilterer) FilterConfirmedAssertion(opts *bind.FilterOpts) (*ArbRollupConfirmedAssertionIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "ConfirmedAssertion")
	if err != nil {
		return nil, err
	}
	return &ArbRollupConfirmedAssertionIterator{contract: _ArbRollup.contract, event: "ConfirmedAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedAssertion is a free log subscription operation binding the contract event 0xded5fa103431438087188a5f8c6a4c3ea90996bbd63be7b1b3fa0a425b37fdd5.
//
// Solidity: event ConfirmedAssertion(bytes32[] logsAccHash)
func (_ArbRollup *ArbRollupFilterer) WatchConfirmedAssertion(opts *bind.WatchOpts, sink chan<- *ArbRollupConfirmedAssertion) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "ConfirmedAssertion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupConfirmedAssertion)
				if err := _ArbRollup.contract.UnpackLog(event, "ConfirmedAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmedAssertion is a log parse operation binding the contract event 0xded5fa103431438087188a5f8c6a4c3ea90996bbd63be7b1b3fa0a425b37fdd5.
//
// Solidity: event ConfirmedAssertion(bytes32[] logsAccHash)
func (_ArbRollup *ArbRollupFilterer) ParseConfirmedAssertion(log types.Log) (*ArbRollupConfirmedAssertion, error) {
	event := new(ArbRollupConfirmedAssertion)
	if err := _ArbRollup.contract.UnpackLog(event, "ConfirmedAssertion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupConfirmedValidAssertionIterator is returned from FilterConfirmedValidAssertion and is used to iterate over the raw logs and unpacked data for ConfirmedValidAssertion events raised by the ArbRollup contract.
type ArbRollupConfirmedValidAssertionIterator struct {
	Event *ArbRollupConfirmedValidAssertion // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupConfirmedValidAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupConfirmedValidAssertion)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupConfirmedValidAssertion)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupConfirmedValidAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupConfirmedValidAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupConfirmedValidAssertion represents a ConfirmedValidAssertion event raised by the ArbRollup contract.
type ArbRollupConfirmedValidAssertion struct {
	NodeHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConfirmedValidAssertion is a free log retrieval operation binding the contract event 0x89cc5e236414c34f1206c0c14d8ac5b0e5444b669b309aaca16fe3d27749f50e.
//
// Solidity: event ConfirmedValidAssertion(bytes32 indexed nodeHash)
func (_ArbRollup *ArbRollupFilterer) FilterConfirmedValidAssertion(opts *bind.FilterOpts, nodeHash [][32]byte) (*ArbRollupConfirmedValidAssertionIterator, error) {

	var nodeHashRule []interface{}
	for _, nodeHashItem := range nodeHash {
		nodeHashRule = append(nodeHashRule, nodeHashItem)
	}

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "ConfirmedValidAssertion", nodeHashRule)
	if err != nil {
		return nil, err
	}
	return &ArbRollupConfirmedValidAssertionIterator{contract: _ArbRollup.contract, event: "ConfirmedValidAssertion", logs: logs, sub: sub}, nil
}

// WatchConfirmedValidAssertion is a free log subscription operation binding the contract event 0x89cc5e236414c34f1206c0c14d8ac5b0e5444b669b309aaca16fe3d27749f50e.
//
// Solidity: event ConfirmedValidAssertion(bytes32 indexed nodeHash)
func (_ArbRollup *ArbRollupFilterer) WatchConfirmedValidAssertion(opts *bind.WatchOpts, sink chan<- *ArbRollupConfirmedValidAssertion, nodeHash [][32]byte) (event.Subscription, error) {

	var nodeHashRule []interface{}
	for _, nodeHashItem := range nodeHash {
		nodeHashRule = append(nodeHashRule, nodeHashItem)
	}

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "ConfirmedValidAssertion", nodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupConfirmedValidAssertion)
				if err := _ArbRollup.contract.UnpackLog(event, "ConfirmedValidAssertion", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfirmedValidAssertion is a log parse operation binding the contract event 0x89cc5e236414c34f1206c0c14d8ac5b0e5444b669b309aaca16fe3d27749f50e.
//
// Solidity: event ConfirmedValidAssertion(bytes32 indexed nodeHash)
func (_ArbRollup *ArbRollupFilterer) ParseConfirmedValidAssertion(log types.Log) (*ArbRollupConfirmedValidAssertion, error) {
	event := new(ArbRollupConfirmedValidAssertion)
	if err := _ArbRollup.contract.UnpackLog(event, "ConfirmedValidAssertion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupRollupAssertedIterator is returned from FilterRollupAsserted and is used to iterate over the raw logs and unpacked data for RollupAsserted events raised by the ArbRollup contract.
type ArbRollupRollupAssertedIterator struct {
	Event *ArbRollupRollupAsserted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupAssertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupAsserted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupAsserted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupAssertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupAssertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupAsserted represents a RollupAsserted event raised by the ArbRollup contract.
type ArbRollupRollupAsserted struct {
	Fields               [7][32]byte
	InboxCount           *big.Int
	ImportedMessageCount *big.Int
	NumArbGas            uint64
	NumSteps             uint64
	BeforeMessageCount   *big.Int
	MessageCount         uint64
	BeforeLogCount       *big.Int
	LogCount             uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRollupAsserted is a free log retrieval operation binding the contract event 0x3112f8dec1eebe04bc7f92bf1031a1c749e09e57836222fef69df63d591bf603.
//
// Solidity: event RollupAsserted(bytes32[7] fields, uint256 inboxCount, uint256 importedMessageCount, uint64 numArbGas, uint64 numSteps, uint256 beforeMessageCount, uint64 messageCount, uint256 beforeLogCount, uint64 logCount)
func (_ArbRollup *ArbRollupFilterer) FilterRollupAsserted(opts *bind.FilterOpts) (*ArbRollupRollupAssertedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "RollupAsserted")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupAssertedIterator{contract: _ArbRollup.contract, event: "RollupAsserted", logs: logs, sub: sub}, nil
}

// WatchRollupAsserted is a free log subscription operation binding the contract event 0x3112f8dec1eebe04bc7f92bf1031a1c749e09e57836222fef69df63d591bf603.
//
// Solidity: event RollupAsserted(bytes32[7] fields, uint256 inboxCount, uint256 importedMessageCount, uint64 numArbGas, uint64 numSteps, uint256 beforeMessageCount, uint64 messageCount, uint256 beforeLogCount, uint64 logCount)
func (_ArbRollup *ArbRollupFilterer) WatchRollupAsserted(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupAsserted) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "RollupAsserted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupAsserted)
				if err := _ArbRollup.contract.UnpackLog(event, "RollupAsserted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupAsserted is a log parse operation binding the contract event 0x3112f8dec1eebe04bc7f92bf1031a1c749e09e57836222fef69df63d591bf603.
//
// Solidity: event RollupAsserted(bytes32[7] fields, uint256 inboxCount, uint256 importedMessageCount, uint64 numArbGas, uint64 numSteps, uint256 beforeMessageCount, uint64 messageCount, uint256 beforeLogCount, uint64 logCount)
func (_ArbRollup *ArbRollupFilterer) ParseRollupAsserted(log types.Log) (*ArbRollupRollupAsserted, error) {
	event := new(ArbRollupRollupAsserted)
	if err := _ArbRollup.contract.UnpackLog(event, "RollupAsserted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupRollupChallengeCompletedIterator is returned from FilterRollupChallengeCompleted and is used to iterate over the raw logs and unpacked data for RollupChallengeCompleted events raised by the ArbRollup contract.
type ArbRollupRollupChallengeCompletedIterator struct {
	Event *ArbRollupRollupChallengeCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupChallengeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupChallengeCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupChallengeCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupChallengeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupChallengeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupChallengeCompleted represents a RollupChallengeCompleted event raised by the ArbRollup contract.
type ArbRollupRollupChallengeCompleted struct {
	ChallengeContract common.Address
	Winner            common.Address
	Loser             common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeCompleted is a free log retrieval operation binding the contract event 0x468aa7d460319b17466ca163bca353a0c62fff0d7d0fa287f634ef305d946f29.
//
// Solidity: event RollupChallengeCompleted(address challengeContract, address winner, address loser)
func (_ArbRollup *ArbRollupFilterer) FilterRollupChallengeCompleted(opts *bind.FilterOpts) (*ArbRollupRollupChallengeCompletedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "RollupChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupChallengeCompletedIterator{contract: _ArbRollup.contract, event: "RollupChallengeCompleted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeCompleted is a free log subscription operation binding the contract event 0x468aa7d460319b17466ca163bca353a0c62fff0d7d0fa287f634ef305d946f29.
//
// Solidity: event RollupChallengeCompleted(address challengeContract, address winner, address loser)
func (_ArbRollup *ArbRollupFilterer) WatchRollupChallengeCompleted(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupChallengeCompleted) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "RollupChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupChallengeCompleted)
				if err := _ArbRollup.contract.UnpackLog(event, "RollupChallengeCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupChallengeCompleted is a log parse operation binding the contract event 0x468aa7d460319b17466ca163bca353a0c62fff0d7d0fa287f634ef305d946f29.
//
// Solidity: event RollupChallengeCompleted(address challengeContract, address winner, address loser)
func (_ArbRollup *ArbRollupFilterer) ParseRollupChallengeCompleted(log types.Log) (*ArbRollupRollupChallengeCompleted, error) {
	event := new(ArbRollupRollupChallengeCompleted)
	if err := _ArbRollup.contract.UnpackLog(event, "RollupChallengeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the ArbRollup contract.
type ArbRollupRollupChallengeStartedIterator struct {
	Event *ArbRollupRollupChallengeStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupChallengeStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupChallengeStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupChallengeStarted represents a RollupChallengeStarted event raised by the ArbRollup contract.
type ArbRollupRollupChallengeStarted struct {
	Asserter          common.Address
	Challenger        common.Address
	ChallengeType     *big.Int
	ChallengeContract common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0x6c69257ddf620994c6fb9e5304db0e5563db3765bee033ddd61b6a1caa7d043f.
//
// Solidity: event RollupChallengeStarted(address asserter, address challenger, uint256 challengeType, address challengeContract)
func (_ArbRollup *ArbRollupFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts) (*ArbRollupRollupChallengeStartedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "RollupChallengeStarted")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupChallengeStartedIterator{contract: _ArbRollup.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0x6c69257ddf620994c6fb9e5304db0e5563db3765bee033ddd61b6a1caa7d043f.
//
// Solidity: event RollupChallengeStarted(address asserter, address challenger, uint256 challengeType, address challengeContract)
func (_ArbRollup *ArbRollupFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupChallengeStarted) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "RollupChallengeStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupChallengeStarted)
				if err := _ArbRollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0x6c69257ddf620994c6fb9e5304db0e5563db3765bee033ddd61b6a1caa7d043f.
//
// Solidity: event RollupChallengeStarted(address asserter, address challenger, uint256 challengeType, address challengeContract)
func (_ArbRollup *ArbRollupFilterer) ParseRollupChallengeStarted(log types.Log) (*ArbRollupRollupChallengeStarted, error) {
	event := new(ArbRollupRollupChallengeStarted)
	if err := _ArbRollup.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupRollupConfirmedIterator is returned from FilterRollupConfirmed and is used to iterate over the raw logs and unpacked data for RollupConfirmed events raised by the ArbRollup contract.
type ArbRollupRollupConfirmedIterator struct {
	Event *ArbRollupRollupConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupConfirmed represents a RollupConfirmed event raised by the ArbRollup contract.
type ArbRollupRollupConfirmed struct {
	NodeHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRollupConfirmed is a free log retrieval operation binding the contract event 0x9d13d0ad532ca8e545a3b66828cb99a18c3bc98e2a50b4db1990a033fdba6f63.
//
// Solidity: event RollupConfirmed(bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) FilterRollupConfirmed(opts *bind.FilterOpts) (*ArbRollupRollupConfirmedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "RollupConfirmed")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupConfirmedIterator{contract: _ArbRollup.contract, event: "RollupConfirmed", logs: logs, sub: sub}, nil
}

// WatchRollupConfirmed is a free log subscription operation binding the contract event 0x9d13d0ad532ca8e545a3b66828cb99a18c3bc98e2a50b4db1990a033fdba6f63.
//
// Solidity: event RollupConfirmed(bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) WatchRollupConfirmed(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupConfirmed) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "RollupConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupConfirmed)
				if err := _ArbRollup.contract.UnpackLog(event, "RollupConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupConfirmed is a log parse operation binding the contract event 0x9d13d0ad532ca8e545a3b66828cb99a18c3bc98e2a50b4db1990a033fdba6f63.
//
// Solidity: event RollupConfirmed(bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) ParseRollupConfirmed(log types.Log) (*ArbRollupRollupConfirmed, error) {
	event := new(ArbRollupRollupConfirmed)
	if err := _ArbRollup.contract.UnpackLog(event, "RollupConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the ArbRollup contract.
type ArbRollupRollupCreatedIterator struct {
	Event *ArbRollupRollupCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupCreated represents a RollupCreated event raised by the ArbRollup contract.
type ArbRollupRollupCreated struct {
	InitVMHash              [32]byte
	GracePeriodTicks        *big.Int
	ArbGasSpeedLimitPerTick *big.Int
	MaxExecutionSteps       uint64
	StakeRequirement        *big.Int
	Owner                   common.Address
	ExtraConfig             []byte
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x3f3efae8ec7ea5f2d06aa37b37bb676f94c915ba05679d32ccdd0dc570dd5864.
//
// Solidity: event RollupCreated(bytes32 initVMHash, uint128 gracePeriodTicks, uint128 arbGasSpeedLimitPerTick, uint64 maxExecutionSteps, uint128 stakeRequirement, address owner, bytes extraConfig)
func (_ArbRollup *ArbRollupFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*ArbRollupRollupCreatedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupCreatedIterator{contract: _ArbRollup.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x3f3efae8ec7ea5f2d06aa37b37bb676f94c915ba05679d32ccdd0dc570dd5864.
//
// Solidity: event RollupCreated(bytes32 initVMHash, uint128 gracePeriodTicks, uint128 arbGasSpeedLimitPerTick, uint64 maxExecutionSteps, uint128 stakeRequirement, address owner, bytes extraConfig)
func (_ArbRollup *ArbRollupFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupCreated) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupCreated)
				if err := _ArbRollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupCreated is a log parse operation binding the contract event 0x3f3efae8ec7ea5f2d06aa37b37bb676f94c915ba05679d32ccdd0dc570dd5864.
//
// Solidity: event RollupCreated(bytes32 initVMHash, uint128 gracePeriodTicks, uint128 arbGasSpeedLimitPerTick, uint64 maxExecutionSteps, uint128 stakeRequirement, address owner, bytes extraConfig)
func (_ArbRollup *ArbRollupFilterer) ParseRollupCreated(log types.Log) (*ArbRollupRollupCreated, error) {
	event := new(ArbRollupRollupCreated)
	if err := _ArbRollup.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupRollupPrunedIterator is returned from FilterRollupPruned and is used to iterate over the raw logs and unpacked data for RollupPruned events raised by the ArbRollup contract.
type ArbRollupRollupPrunedIterator struct {
	Event *ArbRollupRollupPruned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupPrunedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupPruned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupPruned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupPrunedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupPrunedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupPruned represents a RollupPruned event raised by the ArbRollup contract.
type ArbRollupRollupPruned struct {
	Leaf [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRollupPruned is a free log retrieval operation binding the contract event 0x3d3e2ada9638548d1bb115fd766ef675213d953efe8d433bbd8d6718f4490950.
//
// Solidity: event RollupPruned(bytes32 leaf)
func (_ArbRollup *ArbRollupFilterer) FilterRollupPruned(opts *bind.FilterOpts) (*ArbRollupRollupPrunedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "RollupPruned")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupPrunedIterator{contract: _ArbRollup.contract, event: "RollupPruned", logs: logs, sub: sub}, nil
}

// WatchRollupPruned is a free log subscription operation binding the contract event 0x3d3e2ada9638548d1bb115fd766ef675213d953efe8d433bbd8d6718f4490950.
//
// Solidity: event RollupPruned(bytes32 leaf)
func (_ArbRollup *ArbRollupFilterer) WatchRollupPruned(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupPruned) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "RollupPruned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupPruned)
				if err := _ArbRollup.contract.UnpackLog(event, "RollupPruned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupPruned is a log parse operation binding the contract event 0x3d3e2ada9638548d1bb115fd766ef675213d953efe8d433bbd8d6718f4490950.
//
// Solidity: event RollupPruned(bytes32 leaf)
func (_ArbRollup *ArbRollupFilterer) ParseRollupPruned(log types.Log) (*ArbRollupRollupPruned, error) {
	event := new(ArbRollupRollupPruned)
	if err := _ArbRollup.contract.UnpackLog(event, "RollupPruned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupRollupStakeCreatedIterator is returned from FilterRollupStakeCreated and is used to iterate over the raw logs and unpacked data for RollupStakeCreated events raised by the ArbRollup contract.
type ArbRollupRollupStakeCreatedIterator struct {
	Event *ArbRollupRollupStakeCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupStakeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupStakeCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupStakeCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupStakeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupStakeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupStakeCreated represents a RollupStakeCreated event raised by the ArbRollup contract.
type ArbRollupRollupStakeCreated struct {
	Staker   common.Address
	NodeHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRollupStakeCreated is a free log retrieval operation binding the contract event 0xcbafbb223ed21c82af9e2ad20cdfdf55d3263d06f9a65b3f70da613f32d81f88.
//
// Solidity: event RollupStakeCreated(address staker, bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) FilterRollupStakeCreated(opts *bind.FilterOpts) (*ArbRollupRollupStakeCreatedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "RollupStakeCreated")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupStakeCreatedIterator{contract: _ArbRollup.contract, event: "RollupStakeCreated", logs: logs, sub: sub}, nil
}

// WatchRollupStakeCreated is a free log subscription operation binding the contract event 0xcbafbb223ed21c82af9e2ad20cdfdf55d3263d06f9a65b3f70da613f32d81f88.
//
// Solidity: event RollupStakeCreated(address staker, bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) WatchRollupStakeCreated(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupStakeCreated) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "RollupStakeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupStakeCreated)
				if err := _ArbRollup.contract.UnpackLog(event, "RollupStakeCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStakeCreated is a log parse operation binding the contract event 0xcbafbb223ed21c82af9e2ad20cdfdf55d3263d06f9a65b3f70da613f32d81f88.
//
// Solidity: event RollupStakeCreated(address staker, bytes32 nodeHash)
func (_ArbRollup *ArbRollupFilterer) ParseRollupStakeCreated(log types.Log) (*ArbRollupRollupStakeCreated, error) {
	event := new(ArbRollupRollupStakeCreated)
	if err := _ArbRollup.contract.UnpackLog(event, "RollupStakeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupRollupStakeMovedIterator is returned from FilterRollupStakeMoved and is used to iterate over the raw logs and unpacked data for RollupStakeMoved events raised by the ArbRollup contract.
type ArbRollupRollupStakeMovedIterator struct {
	Event *ArbRollupRollupStakeMoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupStakeMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupStakeMoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupStakeMoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupStakeMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupStakeMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupStakeMoved represents a RollupStakeMoved event raised by the ArbRollup contract.
type ArbRollupRollupStakeMoved struct {
	Staker     common.Address
	ToNodeHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRollupStakeMoved is a free log retrieval operation binding the contract event 0xbe690ac5fe353c094bcc6f187eeb841c0ca61b6edf32c142eadad655b7d173f4.
//
// Solidity: event RollupStakeMoved(address staker, bytes32 toNodeHash)
func (_ArbRollup *ArbRollupFilterer) FilterRollupStakeMoved(opts *bind.FilterOpts) (*ArbRollupRollupStakeMovedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "RollupStakeMoved")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupStakeMovedIterator{contract: _ArbRollup.contract, event: "RollupStakeMoved", logs: logs, sub: sub}, nil
}

// WatchRollupStakeMoved is a free log subscription operation binding the contract event 0xbe690ac5fe353c094bcc6f187eeb841c0ca61b6edf32c142eadad655b7d173f4.
//
// Solidity: event RollupStakeMoved(address staker, bytes32 toNodeHash)
func (_ArbRollup *ArbRollupFilterer) WatchRollupStakeMoved(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupStakeMoved) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "RollupStakeMoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupStakeMoved)
				if err := _ArbRollup.contract.UnpackLog(event, "RollupStakeMoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStakeMoved is a log parse operation binding the contract event 0xbe690ac5fe353c094bcc6f187eeb841c0ca61b6edf32c142eadad655b7d173f4.
//
// Solidity: event RollupStakeMoved(address staker, bytes32 toNodeHash)
func (_ArbRollup *ArbRollupFilterer) ParseRollupStakeMoved(log types.Log) (*ArbRollupRollupStakeMoved, error) {
	event := new(ArbRollupRollupStakeMoved)
	if err := _ArbRollup.contract.UnpackLog(event, "RollupStakeMoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbRollupRollupStakeRefundedIterator is returned from FilterRollupStakeRefunded and is used to iterate over the raw logs and unpacked data for RollupStakeRefunded events raised by the ArbRollup contract.
type ArbRollupRollupStakeRefundedIterator struct {
	Event *ArbRollupRollupStakeRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ArbRollupRollupStakeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbRollupRollupStakeRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ArbRollupRollupStakeRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ArbRollupRollupStakeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbRollupRollupStakeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbRollupRollupStakeRefunded represents a RollupStakeRefunded event raised by the ArbRollup contract.
type ArbRollupRollupStakeRefunded struct {
	Staker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRollupStakeRefunded is a free log retrieval operation binding the contract event 0x953ab9eece73c907353307064109cf873462177a0e358e463fd89f5b206daa6c.
//
// Solidity: event RollupStakeRefunded(address staker)
func (_ArbRollup *ArbRollupFilterer) FilterRollupStakeRefunded(opts *bind.FilterOpts) (*ArbRollupRollupStakeRefundedIterator, error) {

	logs, sub, err := _ArbRollup.contract.FilterLogs(opts, "RollupStakeRefunded")
	if err != nil {
		return nil, err
	}
	return &ArbRollupRollupStakeRefundedIterator{contract: _ArbRollup.contract, event: "RollupStakeRefunded", logs: logs, sub: sub}, nil
}

// WatchRollupStakeRefunded is a free log subscription operation binding the contract event 0x953ab9eece73c907353307064109cf873462177a0e358e463fd89f5b206daa6c.
//
// Solidity: event RollupStakeRefunded(address staker)
func (_ArbRollup *ArbRollupFilterer) WatchRollupStakeRefunded(opts *bind.WatchOpts, sink chan<- *ArbRollupRollupStakeRefunded) (event.Subscription, error) {

	logs, sub, err := _ArbRollup.contract.WatchLogs(opts, "RollupStakeRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbRollupRollupStakeRefunded)
				if err := _ArbRollup.contract.UnpackLog(event, "RollupStakeRefunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStakeRefunded is a log parse operation binding the contract event 0x953ab9eece73c907353307064109cf873462177a0e358e463fd89f5b206daa6c.
//
// Solidity: event RollupStakeRefunded(address staker)
func (_ArbRollup *ArbRollupFilterer) ParseRollupStakeRefunded(log types.Log) (*ArbRollupRollupStakeRefunded, error) {
	event := new(ArbRollupRollupStakeRefunded)
	if err := _ArbRollup.contract.UnpackLog(event, "RollupStakeRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeGraphABI is the input ABI used to generate the binding from.
const NodeGraphABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[7]\",\"name\":\"fields\",\"type\":\"bytes32[7]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"importedMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numArbGas\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"numSteps\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeMessageCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"messageCount\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beforeLogCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"logCount\",\"type\":\"uint64\"}],\"name\":\"RollupAsserted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"RollupConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"RollupPruned\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"isValidLeaf\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"latestConfirmedPriv\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"leaves\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"fromNodes\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"leafProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"leafProofLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"latestConfProofs\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"latestConfirmedProofLengths\",\"type\":\"uint256[]\"}],\"name\":\"pruneLeaves\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vmParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gracePeriodTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerTick\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"maxExecutionSteps\",\"type\":\"uint64\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NodeGraphFuncSigs maps the 4-byte function signature to its string representation.
var NodeGraphFuncSigs = map[string]string{
	"57ca6d1b": "isValidLeaf(bytes32)",
	"65f7f80d": "latestConfirmed()",
	"3a218e98": "latestConfirmedPriv()",
	"151bcd2c": "leaves(bytes32)",
	"fcfd8d3f": "pruneLeaves(bytes32[],bytes32[],uint256[],bytes32[],uint256[])",
	"bbc2cc00": "vmParams()",
}

// NodeGraphBin is the compiled bytecode used for deploying new contracts.
var NodeGraphBin = "0x608060405234801561001057600080fd5b506106e4806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063151bcd2c146100675780633a218e981461009857806357ca6d1b146100b257806365f7f80d146100cf578063bbc2cc00146100d7578063fcfd8d3f14610107575b600080fd5b6100846004803603602081101561007d57600080fd5b50356102b7565b604080519115158252519081900360200190f35b6100a06102cc565b60408051918252519081900360200190f35b610084600480360360208110156100c857600080fd5b50356102d2565b6100a06102e7565b6100df6102ed565b60408051938452602084019290925267ffffffffffffffff1682820152519081900360600190f35b6102b5600480360360a081101561011d57600080fd5b810190602081018135600160201b81111561013757600080fd5b82018360208201111561014957600080fd5b803590602001918460208302840111600160201b8311171561016a57600080fd5b919390929091602081019035600160201b81111561018757600080fd5b82018360208201111561019957600080fd5b803590602001918460208302840111600160201b831117156101ba57600080fd5b919390929091602081019035600160201b8111156101d757600080fd5b8201836020820111156101e957600080fd5b803590602001918460208302840111600160201b8311171561020a57600080fd5b919390929091602081019035600160201b81111561022757600080fd5b82018360208201111561023957600080fd5b803590602001918460208302840111600160201b8311171561025a57600080fd5b919390929091602081019035600160201b81111561027757600080fd5b82018360208201111561028957600080fd5b803590602001918460208302840111600160201b831117156102aa57600080fd5b509092509050610303565b005b60036020526000908152604090205460ff1681565b60045481565b60009081526003602052604090205460ff1690565b60045490565b60005460015460025467ffffffffffffffff1683565b88858114801561031257508181145b61035c576040805162461bcd60e51b81526020600482015260166024820152750d2dce0eae840d8cadccee8d040dad2e6e8dac2e8c6d60531b604482015290519081900360640190fd5b600080805b83811015610431576104248e8e8381811061037857fe5b9050602002013587878481811061038b57fe5b905060200201358c8c8581811061039e57fe5b905060200201358f8f80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050508c8c808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152508b92508a91506104419050565b9093509150600101610361565b5050505050505050505050505050565b6000806000871180156104545750600088115b6040518060400160405280600e81526020016d28292aa722afa82927a7a32622a760911b815250906105045760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156104c95781810151838201526020016104b1565b50505050905090810190601f1680156104f65780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5083870183890160006105156102e7565b6105218d8a8986610647565b149050808015610557575087868151811061053857fe5b602002602001015189888151811061054c57fe5b602002602001015114155b6040518060400160405280600e81526020016d141495539157d0d3d391931250d560921b815250906105ca5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104c95781810151838201526020016104b1565b5060006105d98d8b8a87610647565b90506105e4816102d2565b1561063557600081815260036020908152604091829020805460ff19169055815183815291517f3d3e2ada9638548d1bb115fd766ef675213d953efe8d433bbd8d6718f44909509281900390910190a15b50919b909a5098505050505050505050565b600084835b838110156106a5578186828151811061066157fe5b60200260200101516040516020018083815260200182815260200192505050604051602081830303815290604052805190602001209150808060010191505061064c565b509594505050505056fea265627a7a723158209e89fa471b6c307c5ce51410302cf9d71ab19094c4f291fa28568ab72205cedc64736f6c63430005110032"

// DeployNodeGraph deploys a new Ethereum contract, binding an instance of NodeGraph to it.
func DeployNodeGraph(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeGraph, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeGraphABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeGraphBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeGraph{NodeGraphCaller: NodeGraphCaller{contract: contract}, NodeGraphTransactor: NodeGraphTransactor{contract: contract}, NodeGraphFilterer: NodeGraphFilterer{contract: contract}}, nil
}

// NodeGraph is an auto generated Go binding around an Ethereum contract.
type NodeGraph struct {
	NodeGraphCaller     // Read-only binding to the contract
	NodeGraphTransactor // Write-only binding to the contract
	NodeGraphFilterer   // Log filterer for contract events
}

// NodeGraphCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeGraphCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeGraphTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeGraphTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeGraphFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeGraphFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeGraphSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeGraphSession struct {
	Contract     *NodeGraph        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeGraphCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeGraphCallerSession struct {
	Contract *NodeGraphCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// NodeGraphTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeGraphTransactorSession struct {
	Contract     *NodeGraphTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// NodeGraphRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeGraphRaw struct {
	Contract *NodeGraph // Generic contract binding to access the raw methods on
}

// NodeGraphCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeGraphCallerRaw struct {
	Contract *NodeGraphCaller // Generic read-only contract binding to access the raw methods on
}

// NodeGraphTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeGraphTransactorRaw struct {
	Contract *NodeGraphTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeGraph creates a new instance of NodeGraph, bound to a specific deployed contract.
func NewNodeGraph(address common.Address, backend bind.ContractBackend) (*NodeGraph, error) {
	contract, err := bindNodeGraph(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeGraph{NodeGraphCaller: NodeGraphCaller{contract: contract}, NodeGraphTransactor: NodeGraphTransactor{contract: contract}, NodeGraphFilterer: NodeGraphFilterer{contract: contract}}, nil
}

// NewNodeGraphCaller creates a new read-only instance of NodeGraph, bound to a specific deployed contract.
func NewNodeGraphCaller(address common.Address, caller bind.ContractCaller) (*NodeGraphCaller, error) {
	contract, err := bindNodeGraph(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeGraphCaller{contract: contract}, nil
}

// NewNodeGraphTransactor creates a new write-only instance of NodeGraph, bound to a specific deployed contract.
func NewNodeGraphTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeGraphTransactor, error) {
	contract, err := bindNodeGraph(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeGraphTransactor{contract: contract}, nil
}

// NewNodeGraphFilterer creates a new log filterer instance of NodeGraph, bound to a specific deployed contract.
func NewNodeGraphFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeGraphFilterer, error) {
	contract, err := bindNodeGraph(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeGraphFilterer{contract: contract}, nil
}

// bindNodeGraph binds a generic wrapper to an already deployed contract.
func bindNodeGraph(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeGraphABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeGraph *NodeGraphRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeGraph.Contract.NodeGraphCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeGraph *NodeGraphRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeGraph.Contract.NodeGraphTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeGraph *NodeGraphRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeGraph.Contract.NodeGraphTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeGraph *NodeGraphCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeGraph.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeGraph *NodeGraphTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeGraph.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeGraph *NodeGraphTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeGraph.Contract.contract.Transact(opts, method, params...)
}

// IsValidLeaf is a free data retrieval call binding the contract method 0x57ca6d1b.
//
// Solidity: function isValidLeaf(bytes32 leaf) view returns(bool)
func (_NodeGraph *NodeGraphCaller) IsValidLeaf(opts *bind.CallOpts, leaf [32]byte) (bool, error) {
	var out []interface{}
	err := _NodeGraph.contract.Call(opts, &out, "isValidLeaf", leaf)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidLeaf is a free data retrieval call binding the contract method 0x57ca6d1b.
//
// Solidity: function isValidLeaf(bytes32 leaf) view returns(bool)
func (_NodeGraph *NodeGraphSession) IsValidLeaf(leaf [32]byte) (bool, error) {
	return _NodeGraph.Contract.IsValidLeaf(&_NodeGraph.CallOpts, leaf)
}

// IsValidLeaf is a free data retrieval call binding the contract method 0x57ca6d1b.
//
// Solidity: function isValidLeaf(bytes32 leaf) view returns(bool)
func (_NodeGraph *NodeGraphCallerSession) IsValidLeaf(leaf [32]byte) (bool, error) {
	return _NodeGraph.Contract.IsValidLeaf(&_NodeGraph.CallOpts, leaf)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(bytes32)
func (_NodeGraph *NodeGraphCaller) LatestConfirmed(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeGraph.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(bytes32)
func (_NodeGraph *NodeGraphSession) LatestConfirmed() ([32]byte, error) {
	return _NodeGraph.Contract.LatestConfirmed(&_NodeGraph.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(bytes32)
func (_NodeGraph *NodeGraphCallerSession) LatestConfirmed() ([32]byte, error) {
	return _NodeGraph.Contract.LatestConfirmed(&_NodeGraph.CallOpts)
}

// LatestConfirmedPriv is a free data retrieval call binding the contract method 0x3a218e98.
//
// Solidity: function latestConfirmedPriv() view returns(bytes32)
func (_NodeGraph *NodeGraphCaller) LatestConfirmedPriv(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NodeGraph.contract.Call(opts, &out, "latestConfirmedPriv")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// LatestConfirmedPriv is a free data retrieval call binding the contract method 0x3a218e98.
//
// Solidity: function latestConfirmedPriv() view returns(bytes32)
func (_NodeGraph *NodeGraphSession) LatestConfirmedPriv() ([32]byte, error) {
	return _NodeGraph.Contract.LatestConfirmedPriv(&_NodeGraph.CallOpts)
}

// LatestConfirmedPriv is a free data retrieval call binding the contract method 0x3a218e98.
//
// Solidity: function latestConfirmedPriv() view returns(bytes32)
func (_NodeGraph *NodeGraphCallerSession) LatestConfirmedPriv() ([32]byte, error) {
	return _NodeGraph.Contract.LatestConfirmedPriv(&_NodeGraph.CallOpts)
}

// Leaves is a free data retrieval call binding the contract method 0x151bcd2c.
//
// Solidity: function leaves(bytes32 ) view returns(bool)
func (_NodeGraph *NodeGraphCaller) Leaves(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _NodeGraph.contract.Call(opts, &out, "leaves", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Leaves is a free data retrieval call binding the contract method 0x151bcd2c.
//
// Solidity: function leaves(bytes32 ) view returns(bool)
func (_NodeGraph *NodeGraphSession) Leaves(arg0 [32]byte) (bool, error) {
	return _NodeGraph.Contract.Leaves(&_NodeGraph.CallOpts, arg0)
}

// Leaves is a free data retrieval call binding the contract method 0x151bcd2c.
//
// Solidity: function leaves(bytes32 ) view returns(bool)
func (_NodeGraph *NodeGraphCallerSession) Leaves(arg0 [32]byte) (bool, error) {
	return _NodeGraph.Contract.Leaves(&_NodeGraph.CallOpts, arg0)
}

// VmParams is a free data retrieval call binding the contract method 0xbbc2cc00.
//
// Solidity: function vmParams() view returns(uint256 gracePeriodTicks, uint256 arbGasSpeedLimitPerTick, uint64 maxExecutionSteps)
func (_NodeGraph *NodeGraphCaller) VmParams(opts *bind.CallOpts) (struct {
	GracePeriodTicks        *big.Int
	ArbGasSpeedLimitPerTick *big.Int
	MaxExecutionSteps       uint64
}, error) {
	var out []interface{}
	err := _NodeGraph.contract.Call(opts, &out, "vmParams")

	outstruct := new(struct {
		GracePeriodTicks        *big.Int
		ArbGasSpeedLimitPerTick *big.Int
		MaxExecutionSteps       uint64
	})

	outstruct.GracePeriodTicks = out[0].(*big.Int)
	outstruct.ArbGasSpeedLimitPerTick = out[1].(*big.Int)
	outstruct.MaxExecutionSteps = out[2].(uint64)

	return *outstruct, err

}

// VmParams is a free data retrieval call binding the contract method 0xbbc2cc00.
//
// Solidity: function vmParams() view returns(uint256 gracePeriodTicks, uint256 arbGasSpeedLimitPerTick, uint64 maxExecutionSteps)
func (_NodeGraph *NodeGraphSession) VmParams() (struct {
	GracePeriodTicks        *big.Int
	ArbGasSpeedLimitPerTick *big.Int
	MaxExecutionSteps       uint64
}, error) {
	return _NodeGraph.Contract.VmParams(&_NodeGraph.CallOpts)
}

// VmParams is a free data retrieval call binding the contract method 0xbbc2cc00.
//
// Solidity: function vmParams() view returns(uint256 gracePeriodTicks, uint256 arbGasSpeedLimitPerTick, uint64 maxExecutionSteps)
func (_NodeGraph *NodeGraphCallerSession) VmParams() (struct {
	GracePeriodTicks        *big.Int
	ArbGasSpeedLimitPerTick *big.Int
	MaxExecutionSteps       uint64
}, error) {
	return _NodeGraph.Contract.VmParams(&_NodeGraph.CallOpts)
}

// PruneLeaves is a paid mutator transaction binding the contract method 0xfcfd8d3f.
//
// Solidity: function pruneLeaves(bytes32[] fromNodes, bytes32[] leafProofs, uint256[] leafProofLengths, bytes32[] latestConfProofs, uint256[] latestConfirmedProofLengths) returns()
func (_NodeGraph *NodeGraphTransactor) PruneLeaves(opts *bind.TransactOpts, fromNodes [][32]byte, leafProofs [][32]byte, leafProofLengths []*big.Int, latestConfProofs [][32]byte, latestConfirmedProofLengths []*big.Int) (*types.Transaction, error) {
	return _NodeGraph.contract.Transact(opts, "pruneLeaves", fromNodes, leafProofs, leafProofLengths, latestConfProofs, latestConfirmedProofLengths)
}

// PruneLeaves is a paid mutator transaction binding the contract method 0xfcfd8d3f.
//
// Solidity: function pruneLeaves(bytes32[] fromNodes, bytes32[] leafProofs, uint256[] leafProofLengths, bytes32[] latestConfProofs, uint256[] latestConfirmedProofLengths) returns()
func (_NodeGraph *NodeGraphSession) PruneLeaves(fromNodes [][32]byte, leafProofs [][32]byte, leafProofLengths []*big.Int, latestConfProofs [][32]byte, latestConfirmedProofLengths []*big.Int) (*types.Transaction, error) {
	return _NodeGraph.Contract.PruneLeaves(&_NodeGraph.TransactOpts, fromNodes, leafProofs, leafProofLengths, latestConfProofs, latestConfirmedProofLengths)
}

// PruneLeaves is a paid mutator transaction binding the contract method 0xfcfd8d3f.
//
// Solidity: function pruneLeaves(bytes32[] fromNodes, bytes32[] leafProofs, uint256[] leafProofLengths, bytes32[] latestConfProofs, uint256[] latestConfirmedProofLengths) returns()
func (_NodeGraph *NodeGraphTransactorSession) PruneLeaves(fromNodes [][32]byte, leafProofs [][32]byte, leafProofLengths []*big.Int, latestConfProofs [][32]byte, latestConfirmedProofLengths []*big.Int) (*types.Transaction, error) {
	return _NodeGraph.Contract.PruneLeaves(&_NodeGraph.TransactOpts, fromNodes, leafProofs, leafProofLengths, latestConfProofs, latestConfirmedProofLengths)
}

// NodeGraphRollupAssertedIterator is returned from FilterRollupAsserted and is used to iterate over the raw logs and unpacked data for RollupAsserted events raised by the NodeGraph contract.
type NodeGraphRollupAssertedIterator struct {
	Event *NodeGraphRollupAsserted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeGraphRollupAssertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeGraphRollupAsserted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeGraphRollupAsserted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeGraphRollupAssertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeGraphRollupAssertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeGraphRollupAsserted represents a RollupAsserted event raised by the NodeGraph contract.
type NodeGraphRollupAsserted struct {
	Fields               [7][32]byte
	InboxCount           *big.Int
	ImportedMessageCount *big.Int
	NumArbGas            uint64
	NumSteps             uint64
	BeforeMessageCount   *big.Int
	MessageCount         uint64
	BeforeLogCount       *big.Int
	LogCount             uint64
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterRollupAsserted is a free log retrieval operation binding the contract event 0x3112f8dec1eebe04bc7f92bf1031a1c749e09e57836222fef69df63d591bf603.
//
// Solidity: event RollupAsserted(bytes32[7] fields, uint256 inboxCount, uint256 importedMessageCount, uint64 numArbGas, uint64 numSteps, uint256 beforeMessageCount, uint64 messageCount, uint256 beforeLogCount, uint64 logCount)
func (_NodeGraph *NodeGraphFilterer) FilterRollupAsserted(opts *bind.FilterOpts) (*NodeGraphRollupAssertedIterator, error) {

	logs, sub, err := _NodeGraph.contract.FilterLogs(opts, "RollupAsserted")
	if err != nil {
		return nil, err
	}
	return &NodeGraphRollupAssertedIterator{contract: _NodeGraph.contract, event: "RollupAsserted", logs: logs, sub: sub}, nil
}

// WatchRollupAsserted is a free log subscription operation binding the contract event 0x3112f8dec1eebe04bc7f92bf1031a1c749e09e57836222fef69df63d591bf603.
//
// Solidity: event RollupAsserted(bytes32[7] fields, uint256 inboxCount, uint256 importedMessageCount, uint64 numArbGas, uint64 numSteps, uint256 beforeMessageCount, uint64 messageCount, uint256 beforeLogCount, uint64 logCount)
func (_NodeGraph *NodeGraphFilterer) WatchRollupAsserted(opts *bind.WatchOpts, sink chan<- *NodeGraphRollupAsserted) (event.Subscription, error) {

	logs, sub, err := _NodeGraph.contract.WatchLogs(opts, "RollupAsserted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeGraphRollupAsserted)
				if err := _NodeGraph.contract.UnpackLog(event, "RollupAsserted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupAsserted is a log parse operation binding the contract event 0x3112f8dec1eebe04bc7f92bf1031a1c749e09e57836222fef69df63d591bf603.
//
// Solidity: event RollupAsserted(bytes32[7] fields, uint256 inboxCount, uint256 importedMessageCount, uint64 numArbGas, uint64 numSteps, uint256 beforeMessageCount, uint64 messageCount, uint256 beforeLogCount, uint64 logCount)
func (_NodeGraph *NodeGraphFilterer) ParseRollupAsserted(log types.Log) (*NodeGraphRollupAsserted, error) {
	event := new(NodeGraphRollupAsserted)
	if err := _NodeGraph.contract.UnpackLog(event, "RollupAsserted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeGraphRollupConfirmedIterator is returned from FilterRollupConfirmed and is used to iterate over the raw logs and unpacked data for RollupConfirmed events raised by the NodeGraph contract.
type NodeGraphRollupConfirmedIterator struct {
	Event *NodeGraphRollupConfirmed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeGraphRollupConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeGraphRollupConfirmed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeGraphRollupConfirmed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeGraphRollupConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeGraphRollupConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeGraphRollupConfirmed represents a RollupConfirmed event raised by the NodeGraph contract.
type NodeGraphRollupConfirmed struct {
	NodeHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRollupConfirmed is a free log retrieval operation binding the contract event 0x9d13d0ad532ca8e545a3b66828cb99a18c3bc98e2a50b4db1990a033fdba6f63.
//
// Solidity: event RollupConfirmed(bytes32 nodeHash)
func (_NodeGraph *NodeGraphFilterer) FilterRollupConfirmed(opts *bind.FilterOpts) (*NodeGraphRollupConfirmedIterator, error) {

	logs, sub, err := _NodeGraph.contract.FilterLogs(opts, "RollupConfirmed")
	if err != nil {
		return nil, err
	}
	return &NodeGraphRollupConfirmedIterator{contract: _NodeGraph.contract, event: "RollupConfirmed", logs: logs, sub: sub}, nil
}

// WatchRollupConfirmed is a free log subscription operation binding the contract event 0x9d13d0ad532ca8e545a3b66828cb99a18c3bc98e2a50b4db1990a033fdba6f63.
//
// Solidity: event RollupConfirmed(bytes32 nodeHash)
func (_NodeGraph *NodeGraphFilterer) WatchRollupConfirmed(opts *bind.WatchOpts, sink chan<- *NodeGraphRollupConfirmed) (event.Subscription, error) {

	logs, sub, err := _NodeGraph.contract.WatchLogs(opts, "RollupConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeGraphRollupConfirmed)
				if err := _NodeGraph.contract.UnpackLog(event, "RollupConfirmed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupConfirmed is a log parse operation binding the contract event 0x9d13d0ad532ca8e545a3b66828cb99a18c3bc98e2a50b4db1990a033fdba6f63.
//
// Solidity: event RollupConfirmed(bytes32 nodeHash)
func (_NodeGraph *NodeGraphFilterer) ParseRollupConfirmed(log types.Log) (*NodeGraphRollupConfirmed, error) {
	event := new(NodeGraphRollupConfirmed)
	if err := _NodeGraph.contract.UnpackLog(event, "RollupConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeGraphRollupPrunedIterator is returned from FilterRollupPruned and is used to iterate over the raw logs and unpacked data for RollupPruned events raised by the NodeGraph contract.
type NodeGraphRollupPrunedIterator struct {
	Event *NodeGraphRollupPruned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NodeGraphRollupPrunedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NodeGraphRollupPruned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NodeGraphRollupPruned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NodeGraphRollupPrunedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NodeGraphRollupPrunedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NodeGraphRollupPruned represents a RollupPruned event raised by the NodeGraph contract.
type NodeGraphRollupPruned struct {
	Leaf [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRollupPruned is a free log retrieval operation binding the contract event 0x3d3e2ada9638548d1bb115fd766ef675213d953efe8d433bbd8d6718f4490950.
//
// Solidity: event RollupPruned(bytes32 leaf)
func (_NodeGraph *NodeGraphFilterer) FilterRollupPruned(opts *bind.FilterOpts) (*NodeGraphRollupPrunedIterator, error) {

	logs, sub, err := _NodeGraph.contract.FilterLogs(opts, "RollupPruned")
	if err != nil {
		return nil, err
	}
	return &NodeGraphRollupPrunedIterator{contract: _NodeGraph.contract, event: "RollupPruned", logs: logs, sub: sub}, nil
}

// WatchRollupPruned is a free log subscription operation binding the contract event 0x3d3e2ada9638548d1bb115fd766ef675213d953efe8d433bbd8d6718f4490950.
//
// Solidity: event RollupPruned(bytes32 leaf)
func (_NodeGraph *NodeGraphFilterer) WatchRollupPruned(opts *bind.WatchOpts, sink chan<- *NodeGraphRollupPruned) (event.Subscription, error) {

	logs, sub, err := _NodeGraph.contract.WatchLogs(opts, "RollupPruned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NodeGraphRollupPruned)
				if err := _NodeGraph.contract.UnpackLog(event, "RollupPruned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupPruned is a log parse operation binding the contract event 0x3d3e2ada9638548d1bb115fd766ef675213d953efe8d433bbd8d6718f4490950.
//
// Solidity: event RollupPruned(bytes32 leaf)
func (_NodeGraph *NodeGraphFilterer) ParseRollupPruned(log types.Log) (*NodeGraphRollupPruned, error) {
	event := new(NodeGraphRollupPruned)
	if err := _NodeGraph.contract.UnpackLog(event, "RollupPruned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NodeGraphUtilsABI is the input ABI used to generate the binding from.
const NodeGraphUtilsABI = "[]"

// NodeGraphUtilsBin is the compiled bytecode used for deploying new contracts.
var NodeGraphUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582086807d620e5b044209761928521ba376ee192b256394f44013a98801b9991bdd64736f6c63430005110032"

// DeployNodeGraphUtils deploys a new Ethereum contract, binding an instance of NodeGraphUtils to it.
func DeployNodeGraphUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NodeGraphUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeGraphUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NodeGraphUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NodeGraphUtils{NodeGraphUtilsCaller: NodeGraphUtilsCaller{contract: contract}, NodeGraphUtilsTransactor: NodeGraphUtilsTransactor{contract: contract}, NodeGraphUtilsFilterer: NodeGraphUtilsFilterer{contract: contract}}, nil
}

// NodeGraphUtils is an auto generated Go binding around an Ethereum contract.
type NodeGraphUtils struct {
	NodeGraphUtilsCaller     // Read-only binding to the contract
	NodeGraphUtilsTransactor // Write-only binding to the contract
	NodeGraphUtilsFilterer   // Log filterer for contract events
}

// NodeGraphUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type NodeGraphUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeGraphUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NodeGraphUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeGraphUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NodeGraphUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NodeGraphUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NodeGraphUtilsSession struct {
	Contract     *NodeGraphUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NodeGraphUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NodeGraphUtilsCallerSession struct {
	Contract *NodeGraphUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NodeGraphUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NodeGraphUtilsTransactorSession struct {
	Contract     *NodeGraphUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NodeGraphUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type NodeGraphUtilsRaw struct {
	Contract *NodeGraphUtils // Generic contract binding to access the raw methods on
}

// NodeGraphUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NodeGraphUtilsCallerRaw struct {
	Contract *NodeGraphUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// NodeGraphUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NodeGraphUtilsTransactorRaw struct {
	Contract *NodeGraphUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNodeGraphUtils creates a new instance of NodeGraphUtils, bound to a specific deployed contract.
func NewNodeGraphUtils(address common.Address, backend bind.ContractBackend) (*NodeGraphUtils, error) {
	contract, err := bindNodeGraphUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NodeGraphUtils{NodeGraphUtilsCaller: NodeGraphUtilsCaller{contract: contract}, NodeGraphUtilsTransactor: NodeGraphUtilsTransactor{contract: contract}, NodeGraphUtilsFilterer: NodeGraphUtilsFilterer{contract: contract}}, nil
}

// NewNodeGraphUtilsCaller creates a new read-only instance of NodeGraphUtils, bound to a specific deployed contract.
func NewNodeGraphUtilsCaller(address common.Address, caller bind.ContractCaller) (*NodeGraphUtilsCaller, error) {
	contract, err := bindNodeGraphUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NodeGraphUtilsCaller{contract: contract}, nil
}

// NewNodeGraphUtilsTransactor creates a new write-only instance of NodeGraphUtils, bound to a specific deployed contract.
func NewNodeGraphUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*NodeGraphUtilsTransactor, error) {
	contract, err := bindNodeGraphUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NodeGraphUtilsTransactor{contract: contract}, nil
}

// NewNodeGraphUtilsFilterer creates a new log filterer instance of NodeGraphUtils, bound to a specific deployed contract.
func NewNodeGraphUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*NodeGraphUtilsFilterer, error) {
	contract, err := bindNodeGraphUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NodeGraphUtilsFilterer{contract: contract}, nil
}

// bindNodeGraphUtils binds a generic wrapper to an already deployed contract.
func bindNodeGraphUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NodeGraphUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeGraphUtils *NodeGraphUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeGraphUtils.Contract.NodeGraphUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeGraphUtils *NodeGraphUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeGraphUtils.Contract.NodeGraphUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeGraphUtils *NodeGraphUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeGraphUtils.Contract.NodeGraphUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NodeGraphUtils *NodeGraphUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NodeGraphUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NodeGraphUtils *NodeGraphUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NodeGraphUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NodeGraphUtils *NodeGraphUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NodeGraphUtils.Contract.contract.Transact(opts, method, params...)
}

// RollupUtilsABI is the input ABI used to generate the binding from.
const RollupUtilsABI = "[]"

// RollupUtilsBin is the compiled bytecode used for deploying new contracts.
var RollupUtilsBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a7231582093fcb3c9dda65c3c67a671d93b23253577198397382f277b4e69a39194ea4b2c64736f6c63430005110032"

// DeployRollupUtils deploys a new Ethereum contract, binding an instance of RollupUtils to it.
func DeployRollupUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupUtils{RollupUtilsCaller: RollupUtilsCaller{contract: contract}, RollupUtilsTransactor: RollupUtilsTransactor{contract: contract}, RollupUtilsFilterer: RollupUtilsFilterer{contract: contract}}, nil
}

// RollupUtils is an auto generated Go binding around an Ethereum contract.
type RollupUtils struct {
	RollupUtilsCaller     // Read-only binding to the contract
	RollupUtilsTransactor // Write-only binding to the contract
	RollupUtilsFilterer   // Log filterer for contract events
}

// RollupUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupUtilsSession struct {
	Contract     *RollupUtils      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupUtilsCallerSession struct {
	Contract *RollupUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RollupUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupUtilsTransactorSession struct {
	Contract     *RollupUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RollupUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupUtilsRaw struct {
	Contract *RollupUtils // Generic contract binding to access the raw methods on
}

// RollupUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupUtilsCallerRaw struct {
	Contract *RollupUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// RollupUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupUtilsTransactorRaw struct {
	Contract *RollupUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupUtils creates a new instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtils(address common.Address, backend bind.ContractBackend) (*RollupUtils, error) {
	contract, err := bindRollupUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupUtils{RollupUtilsCaller: RollupUtilsCaller{contract: contract}, RollupUtilsTransactor: RollupUtilsTransactor{contract: contract}, RollupUtilsFilterer: RollupUtilsFilterer{contract: contract}}, nil
}

// NewRollupUtilsCaller creates a new read-only instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtilsCaller(address common.Address, caller bind.ContractCaller) (*RollupUtilsCaller, error) {
	contract, err := bindRollupUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsCaller{contract: contract}, nil
}

// NewRollupUtilsTransactor creates a new write-only instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupUtilsTransactor, error) {
	contract, err := bindRollupUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsTransactor{contract: contract}, nil
}

// NewRollupUtilsFilterer creates a new log filterer instance of RollupUtils, bound to a specific deployed contract.
func NewRollupUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupUtilsFilterer, error) {
	contract, err := bindRollupUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupUtilsFilterer{contract: contract}, nil
}

// bindRollupUtils binds a generic wrapper to an already deployed contract.
func bindRollupUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupUtils *RollupUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupUtils.Contract.RollupUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupUtils *RollupUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupUtils.Contract.RollupUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupUtils *RollupUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupUtils.Contract.RollupUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupUtils *RollupUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupUtils *RollupUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupUtils *RollupUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupUtils.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158208779b99760723161ba40a9da7d8b060f33ca5d9edad455d784d6ef6d6e9e69dc64736f6c63430005110032"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StakingABI is the input ABI used to generate the binding from.
const StakingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"winner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"RollupChallengeCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengeType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"}],\"name\":\"RollupStakeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"toNodeHash\",\"type\":\"bytes32\"}],\"name\":\"RollupStakeMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"RollupStakeRefunded\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"challenges\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeRequired\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getWithdrawnStake\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakerAddress\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"winner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"loser\",\"type\":\"address\"}],\"name\":\"resolveChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeRequirement\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"location\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"creationTimeBlocks\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"inChallenge\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"asserterAddress\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"challengerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"prevNode\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"deadlineTicks\",\"type\":\"uint256\"},{\"internalType\":\"uint256[2]\",\"name\":\"stakerNodeTypes\",\"type\":\"uint256[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"vmProtoHashes\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[]\",\"name\":\"asserterProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"challengerProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"asserterNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"challengerDataHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint128\",\"name\":\"challengerPeriodTicks\",\"type\":\"uint128\"}],\"name\":\"startChallenge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawnStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// StakingFuncSigs maps the 4-byte function signature to its string representation.
var StakingFuncSigs = map[string]string{
	"5dbaf68b": "challengeFactory()",
	"08b43a19": "challenges(address)",
	"d16c305d": "getStakeRequired()",
	"500a1564": "getStakeToken()",
	"3fcc045a": "getWithdrawnStake(address)",
	"6177fd18": "isStaked(address)",
	"396f51cf": "resolveChallenge(address,address)",
	"b6f9bbb9": "stakeRequirement()",
	"51ed6a30": "stakeToken()",
	"dff69787": "stakerCount()",
	"9168ae72": "stakers(address)",
	"bac5963f": "startChallenge(address,address,bytes32,uint256,uint256[2],bytes32[2],bytes32[],bytes32[],bytes32,bytes32,uint128)",
	"eb2e74cb": "withdrawnStakes(address)",
}

// StakingBin is the compiled bytecode used for deploying new contracts.
var StakingBin = "0x608060405234801561001057600080fd5b50610f4a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636177fd181161008c578063bac5963f11610066578063bac5963f14610230578063d16c305d146103e6578063dff69787146103ee578063eb2e74cb14610408576100cf565b80636177fd18146101985780639168ae72146101be578063b6f9bbb91461020c576100cf565b806308b43a19146100d4578063396f51cf1461010e5780633fcc045a1461013e578063500a15641461016457806351ed6a30146101885780635dbaf68b14610190575b600080fd5b6100fa600480360360208110156100ea57600080fd5b50356001600160a01b031661042e565b604080519115158252519081900360200190f35b61013c6004803603604081101561012457600080fd5b506001600160a01b0381358116916020013516610443565b005b61013c6004803603602081101561015457600080fd5b50356001600160a01b03166105c3565b61016c61072e565b604080516001600160a01b039092168252519081900360200190f35b61016c61073d565b61016c61074c565b6100fa600480360360208110156101ae57600080fd5b50356001600160a01b031661075b565b6101e4600480360360208110156101d457600080fd5b50356001600160a01b0316610778565b604080519384526001600160801b039092166020840152151582820152519081900360600190f35b6102146107a6565b604080516001600160801b039092168252519081900360200190f35b61013c60048036036101a081101561024757600080fd5b6040805180820182526001600160a01b0384358116946020810135909116938382013593606083013593918301929160c0830191608084019060029083908390808284376000920191909152505060408051808201825292959493818101939250906002908390839080828437600092019190915250919493926020810192503590506401000000008111156102dc57600080fd5b8201836020820111156102ee57600080fd5b8035906020019184602083028401116401000000008311171561031057600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561036057600080fd5b82018360208201111561037257600080fd5b8035906020019184602083028401116401000000008311171561039457600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050823593505050602081013590604001356001600160801b03166107b5565b610214610be7565b6103f6610bf6565b60408051918252519081900360200190f35b6103f66004803603602081101561041e57600080fd5b50356001600160a01b0316610bfc565b60056020526000908152604090205460ff1681565b33600090815260056020908152604091829020548251808401909352600f83526e2922a9afa1a420a62fa9a2a72222a960891b9183019190915260ff166105085760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156104cd5781810151838201526020016104b5565b50505050905090810190601f1680156104fa5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50336000908152600560205260408120805460ff1916905561052983610c0e565b600180546001600160a01b0386166000908152600660205260409020805460026001600160801b039384160490921690910190558101805460ff60801b19169055905061057582610c9f565b604080513381526001600160a01b03808616602083015284168183015290517f468aa7d460319b17466ca163bca353a0c62fff0d7d0fa287f634ef305d946f299181900360600190a1505050565b6001600160a01b038116600090815260066020526040902054806105e7575061072b565b6002546001600160a01b0316610633576040516001600160a01b0383169082156108fc029083906000818181858888f1935050505015801561062d573d6000803e3d6000fd5b50610729565b6002546040805163a9059cbb60e01b81526001600160a01b038581166004830152602482018590529151919092169163a9059cbb9160448083019260209291908290030181600087803b15801561068957600080fd5b505af115801561069d573d6000803e3d6000fd5b505050506040513d60208110156106b357600080fd5b505160408051808201909152600f81526e1514905394d1915497d19052531151608a1b6020820152906107275760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104cd5781810151838201526020016104b5565b505b505b50565b6002546001600160a01b031690565b6002546001600160a01b031681565b6000546001600160a01b031681565b6001600160a01b0316600090815260036020526040902054151590565b600360205260009081526040902080546001909101546001600160801b03811690600160801b900460ff1683565b6001546001600160801b031681565b60006107c08c610c0e565b905060006107cd8c610c0e565b60018301549091508a906107e9906001600160801b0316610ce0565b106040518060400160405280600d81526020016c53544b315f444541444c494e4560981b8152509061085c5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104cd5781810151838201526020016104b5565b5060018101548a90610876906001600160801b0316610ce0565b106040518060400160405280600d81526020016c53544b325f444541444c494e4560981b815250906108e95760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104cd5781810151838201526020016104b5565b50600182015460408051808201909152600c81526b14d512cc57d25397d0d2105360a21b602082015290600160801b900460ff16156109695760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104cd5781810151838201526020016104b5565b50600181015460408051808201909152600c81526b14d512cc97d25397d0d2105360a21b602082015290600160801b900460ff16156109e95760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104cd5781810151838201526020016104b5565b506020808a01518a5160408051808201909152600a8152692a2ca822afa7a92222a960b11b9381019390935211610a615760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104cd5781810151838201526020016104b5565b508154610a89610a838d8d898e600060200201518e60005b6020020151610ce7565b89610d4f565b146040518060400160405280600c81526020016b20a9a9a2a92a2fa82927a7a360a11b81525090610afb5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104cd5781810151838201526020016104b5565b508054610b2c610b268d8d610b19896001600160801b038a16610d66565b60208f01518e6001610a79565b88610d4f565b146040518060400160405280600a81526020016921a420a62fa82927a7a360b11b81525090610b9c5760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104cd5781810151838201526020016104b5565b5060018281018054600160801b60ff60801b19918216811790925591830180549092161790556020890151610bd8908e908e9086908890610d92565b50505050505050505050505050565b6001546001600160801b031690565b60045481565b60066020526000908152604090205481565b6001600160a01b038116600090815260036020908152604080832080548251808401909352600a83526924a72b2fa9aa20a5a2a960b11b9383019390935291610c985760405162461bcd60e51b81526020600482018181528351602484015283519092839260449091019190850190808383600083156104cd5781810151838201526020016104b5565b5092915050565b6001600160a01b03166000908152600360205260408120908155600101805470ffffffffffffffffffffffffffffffffff1916905560048054600019019055565b6103e80290565b6040805160208082018490528183018790526060820186905260808083018690528351808403909101815260a08301845280519082012060c0830189905260e08084019190915283518084039091018152610100909201909252805191012095945050505050565b6000610d5f838360008551610ead565b9392505050565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b600080546040805163432ed0e160e11b81526001600160a01b03898116600483015288811660248301526001600160801b038816604483015260648201879052608482018690529151919092169163865da1c29160a480830192602092919082900301818787803b158015610e0657600080fd5b505af1158015610e1a573d6000803e3d6000fd5b505050506040513d6020811015610e3057600080fd5b50516001600160a01b03808216600081815260056020908152604091829020805460ff1916600117905581518b85168152938a16908401528281018690526060830191909152519192507f6c69257ddf620994c6fb9e5304db0e5563db3765bee033ddd61b6a1caa7d043f919081900360800190a1505050505050565b600084835b83811015610f0b5781868281518110610ec757fe5b602002602001015160405160200180838152602001828152602001925050506040516020818303038152906040528051906020012091508080600101915050610eb2565b509594505050505056fea265627a7a7231582070a79b8a7075439996309960bb65627aeb2fe928fa631dc4d6ea8f0c96d296b664736f6c63430005110032"

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StakingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Staking *StakingCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Staking *StakingSession) ChallengeFactory() (common.Address, error) {
	return _Staking.Contract.ChallengeFactory(&_Staking.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_Staking *StakingCallerSession) ChallengeFactory() (common.Address, error) {
	return _Staking.Contract.ChallengeFactory(&_Staking.CallOpts)
}

// Challenges is a free data retrieval call binding the contract method 0x08b43a19.
//
// Solidity: function challenges(address ) view returns(bool)
func (_Staking *StakingCaller) Challenges(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "challenges", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Challenges is a free data retrieval call binding the contract method 0x08b43a19.
//
// Solidity: function challenges(address ) view returns(bool)
func (_Staking *StakingSession) Challenges(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Challenges(&_Staking.CallOpts, arg0)
}

// Challenges is a free data retrieval call binding the contract method 0x08b43a19.
//
// Solidity: function challenges(address ) view returns(bool)
func (_Staking *StakingCallerSession) Challenges(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Challenges(&_Staking.CallOpts, arg0)
}

// GetStakeRequired is a free data retrieval call binding the contract method 0xd16c305d.
//
// Solidity: function getStakeRequired() view returns(uint128)
func (_Staking *StakingCaller) GetStakeRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStakeRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakeRequired is a free data retrieval call binding the contract method 0xd16c305d.
//
// Solidity: function getStakeRequired() view returns(uint128)
func (_Staking *StakingSession) GetStakeRequired() (*big.Int, error) {
	return _Staking.Contract.GetStakeRequired(&_Staking.CallOpts)
}

// GetStakeRequired is a free data retrieval call binding the contract method 0xd16c305d.
//
// Solidity: function getStakeRequired() view returns(uint128)
func (_Staking *StakingCallerSession) GetStakeRequired() (*big.Int, error) {
	return _Staking.Contract.GetStakeRequired(&_Staking.CallOpts)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x500a1564.
//
// Solidity: function getStakeToken() view returns(address)
func (_Staking *StakingCaller) GetStakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakeToken is a free data retrieval call binding the contract method 0x500a1564.
//
// Solidity: function getStakeToken() view returns(address)
func (_Staking *StakingSession) GetStakeToken() (common.Address, error) {
	return _Staking.Contract.GetStakeToken(&_Staking.CallOpts)
}

// GetStakeToken is a free data retrieval call binding the contract method 0x500a1564.
//
// Solidity: function getStakeToken() view returns(address)
func (_Staking *StakingCallerSession) GetStakeToken() (common.Address, error) {
	return _Staking.Contract.GetStakeToken(&_Staking.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address _stakerAddress) view returns(bool)
func (_Staking *StakingCaller) IsStaked(opts *bind.CallOpts, _stakerAddress common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isStaked", _stakerAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address _stakerAddress) view returns(bool)
func (_Staking *StakingSession) IsStaked(_stakerAddress common.Address) (bool, error) {
	return _Staking.Contract.IsStaked(&_Staking.CallOpts, _stakerAddress)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address _stakerAddress) view returns(bool)
func (_Staking *StakingCallerSession) IsStaked(_stakerAddress common.Address) (bool, error) {
	return _Staking.Contract.IsStaked(&_Staking.CallOpts, _stakerAddress)
}

// StakeRequirement is a free data retrieval call binding the contract method 0xb6f9bbb9.
//
// Solidity: function stakeRequirement() view returns(uint128)
func (_Staking *StakingCaller) StakeRequirement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakeRequirement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeRequirement is a free data retrieval call binding the contract method 0xb6f9bbb9.
//
// Solidity: function stakeRequirement() view returns(uint128)
func (_Staking *StakingSession) StakeRequirement() (*big.Int, error) {
	return _Staking.Contract.StakeRequirement(&_Staking.CallOpts)
}

// StakeRequirement is a free data retrieval call binding the contract method 0xb6f9bbb9.
//
// Solidity: function stakeRequirement() view returns(uint128)
func (_Staking *StakingCallerSession) StakeRequirement() (*big.Int, error) {
	return _Staking.Contract.StakeRequirement(&_Staking.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Staking *StakingCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Staking *StakingSession) StakeToken() (common.Address, error) {
	return _Staking.Contract.StakeToken(&_Staking.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_Staking *StakingCallerSession) StakeToken() (common.Address, error) {
	return _Staking.Contract.StakeToken(&_Staking.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Staking *StakingCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Staking *StakingSession) StakerCount() (*big.Int, error) {
	return _Staking.Contract.StakerCount(&_Staking.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_Staking *StakingCallerSession) StakerCount() (*big.Int, error) {
	return _Staking.Contract.StakerCount(&_Staking.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bytes32 location, uint128 creationTimeBlocks, bool inChallenge)
func (_Staking *StakingCaller) Stakers(opts *bind.CallOpts, arg0 common.Address) (struct {
	Location           [32]byte
	CreationTimeBlocks *big.Int
	InChallenge        bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		Location           [32]byte
		CreationTimeBlocks *big.Int
		InChallenge        bool
	})

	outstruct.Location = out[0].([32]byte)
	outstruct.CreationTimeBlocks = out[1].(*big.Int)
	outstruct.InChallenge = out[2].(bool)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bytes32 location, uint128 creationTimeBlocks, bool inChallenge)
func (_Staking *StakingSession) Stakers(arg0 common.Address) (struct {
	Location           [32]byte
	CreationTimeBlocks *big.Int
	InChallenge        bool
}, error) {
	return _Staking.Contract.Stakers(&_Staking.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(bytes32 location, uint128 creationTimeBlocks, bool inChallenge)
func (_Staking *StakingCallerSession) Stakers(arg0 common.Address) (struct {
	Location           [32]byte
	CreationTimeBlocks *big.Int
	InChallenge        bool
}, error) {
	return _Staking.Contract.Stakers(&_Staking.CallOpts, arg0)
}

// WithdrawnStakes is a free data retrieval call binding the contract method 0xeb2e74cb.
//
// Solidity: function withdrawnStakes(address ) view returns(uint256)
func (_Staking *StakingCaller) WithdrawnStakes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "withdrawnStakes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawnStakes is a free data retrieval call binding the contract method 0xeb2e74cb.
//
// Solidity: function withdrawnStakes(address ) view returns(uint256)
func (_Staking *StakingSession) WithdrawnStakes(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawnStakes(&_Staking.CallOpts, arg0)
}

// WithdrawnStakes is a free data retrieval call binding the contract method 0xeb2e74cb.
//
// Solidity: function withdrawnStakes(address ) view returns(uint256)
func (_Staking *StakingCallerSession) WithdrawnStakes(arg0 common.Address) (*big.Int, error) {
	return _Staking.Contract.WithdrawnStakes(&_Staking.CallOpts, arg0)
}

// GetWithdrawnStake is a paid mutator transaction binding the contract method 0x3fcc045a.
//
// Solidity: function getWithdrawnStake(address _staker) returns()
func (_Staking *StakingTransactor) GetWithdrawnStake(opts *bind.TransactOpts, _staker common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "getWithdrawnStake", _staker)
}

// GetWithdrawnStake is a paid mutator transaction binding the contract method 0x3fcc045a.
//
// Solidity: function getWithdrawnStake(address _staker) returns()
func (_Staking *StakingSession) GetWithdrawnStake(_staker common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GetWithdrawnStake(&_Staking.TransactOpts, _staker)
}

// GetWithdrawnStake is a paid mutator transaction binding the contract method 0x3fcc045a.
//
// Solidity: function getWithdrawnStake(address _staker) returns()
func (_Staking *StakingTransactorSession) GetWithdrawnStake(_staker common.Address) (*types.Transaction, error) {
	return _Staking.Contract.GetWithdrawnStake(&_Staking.TransactOpts, _staker)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_Staking *StakingTransactor) ResolveChallenge(opts *bind.TransactOpts, winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "resolveChallenge", winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_Staking *StakingSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ResolveChallenge(&_Staking.TransactOpts, winner, loser)
}

// ResolveChallenge is a paid mutator transaction binding the contract method 0x396f51cf.
//
// Solidity: function resolveChallenge(address winner, address loser) returns()
func (_Staking *StakingTransactorSession) ResolveChallenge(winner common.Address, loser common.Address) (*types.Transaction, error) {
	return _Staking.Contract.ResolveChallenge(&_Staking.TransactOpts, winner, loser)
}

// StartChallenge is a paid mutator transaction binding the contract method 0xbac5963f.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, bytes32 prevNode, uint256 deadlineTicks, uint256[2] stakerNodeTypes, bytes32[2] vmProtoHashes, bytes32[] asserterProof, bytes32[] challengerProof, bytes32 asserterNodeHash, bytes32 challengerDataHash, uint128 challengerPeriodTicks) returns()
func (_Staking *StakingTransactor) StartChallenge(opts *bind.TransactOpts, asserterAddress common.Address, challengerAddress common.Address, prevNode [32]byte, deadlineTicks *big.Int, stakerNodeTypes [2]*big.Int, vmProtoHashes [2][32]byte, asserterProof [][32]byte, challengerProof [][32]byte, asserterNodeHash [32]byte, challengerDataHash [32]byte, challengerPeriodTicks *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "startChallenge", asserterAddress, challengerAddress, prevNode, deadlineTicks, stakerNodeTypes, vmProtoHashes, asserterProof, challengerProof, asserterNodeHash, challengerDataHash, challengerPeriodTicks)
}

// StartChallenge is a paid mutator transaction binding the contract method 0xbac5963f.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, bytes32 prevNode, uint256 deadlineTicks, uint256[2] stakerNodeTypes, bytes32[2] vmProtoHashes, bytes32[] asserterProof, bytes32[] challengerProof, bytes32 asserterNodeHash, bytes32 challengerDataHash, uint128 challengerPeriodTicks) returns()
func (_Staking *StakingSession) StartChallenge(asserterAddress common.Address, challengerAddress common.Address, prevNode [32]byte, deadlineTicks *big.Int, stakerNodeTypes [2]*big.Int, vmProtoHashes [2][32]byte, asserterProof [][32]byte, challengerProof [][32]byte, asserterNodeHash [32]byte, challengerDataHash [32]byte, challengerPeriodTicks *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StartChallenge(&_Staking.TransactOpts, asserterAddress, challengerAddress, prevNode, deadlineTicks, stakerNodeTypes, vmProtoHashes, asserterProof, challengerProof, asserterNodeHash, challengerDataHash, challengerPeriodTicks)
}

// StartChallenge is a paid mutator transaction binding the contract method 0xbac5963f.
//
// Solidity: function startChallenge(address asserterAddress, address challengerAddress, bytes32 prevNode, uint256 deadlineTicks, uint256[2] stakerNodeTypes, bytes32[2] vmProtoHashes, bytes32[] asserterProof, bytes32[] challengerProof, bytes32 asserterNodeHash, bytes32 challengerDataHash, uint128 challengerPeriodTicks) returns()
func (_Staking *StakingTransactorSession) StartChallenge(asserterAddress common.Address, challengerAddress common.Address, prevNode [32]byte, deadlineTicks *big.Int, stakerNodeTypes [2]*big.Int, vmProtoHashes [2][32]byte, asserterProof [][32]byte, challengerProof [][32]byte, asserterNodeHash [32]byte, challengerDataHash [32]byte, challengerPeriodTicks *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.StartChallenge(&_Staking.TransactOpts, asserterAddress, challengerAddress, prevNode, deadlineTicks, stakerNodeTypes, vmProtoHashes, asserterProof, challengerProof, asserterNodeHash, challengerDataHash, challengerPeriodTicks)
}

// StakingRollupChallengeCompletedIterator is returned from FilterRollupChallengeCompleted and is used to iterate over the raw logs and unpacked data for RollupChallengeCompleted events raised by the Staking contract.
type StakingRollupChallengeCompletedIterator struct {
	Event *StakingRollupChallengeCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingRollupChallengeCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRollupChallengeCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingRollupChallengeCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingRollupChallengeCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRollupChallengeCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRollupChallengeCompleted represents a RollupChallengeCompleted event raised by the Staking contract.
type StakingRollupChallengeCompleted struct {
	ChallengeContract common.Address
	Winner            common.Address
	Loser             common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeCompleted is a free log retrieval operation binding the contract event 0x468aa7d460319b17466ca163bca353a0c62fff0d7d0fa287f634ef305d946f29.
//
// Solidity: event RollupChallengeCompleted(address challengeContract, address winner, address loser)
func (_Staking *StakingFilterer) FilterRollupChallengeCompleted(opts *bind.FilterOpts) (*StakingRollupChallengeCompletedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RollupChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return &StakingRollupChallengeCompletedIterator{contract: _Staking.contract, event: "RollupChallengeCompleted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeCompleted is a free log subscription operation binding the contract event 0x468aa7d460319b17466ca163bca353a0c62fff0d7d0fa287f634ef305d946f29.
//
// Solidity: event RollupChallengeCompleted(address challengeContract, address winner, address loser)
func (_Staking *StakingFilterer) WatchRollupChallengeCompleted(opts *bind.WatchOpts, sink chan<- *StakingRollupChallengeCompleted) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RollupChallengeCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRollupChallengeCompleted)
				if err := _Staking.contract.UnpackLog(event, "RollupChallengeCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupChallengeCompleted is a log parse operation binding the contract event 0x468aa7d460319b17466ca163bca353a0c62fff0d7d0fa287f634ef305d946f29.
//
// Solidity: event RollupChallengeCompleted(address challengeContract, address winner, address loser)
func (_Staking *StakingFilterer) ParseRollupChallengeCompleted(log types.Log) (*StakingRollupChallengeCompleted, error) {
	event := new(StakingRollupChallengeCompleted)
	if err := _Staking.contract.UnpackLog(event, "RollupChallengeCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the Staking contract.
type StakingRollupChallengeStartedIterator struct {
	Event *StakingRollupChallengeStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRollupChallengeStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingRollupChallengeStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRollupChallengeStarted represents a RollupChallengeStarted event raised by the Staking contract.
type StakingRollupChallengeStarted struct {
	Asserter          common.Address
	Challenger        common.Address
	ChallengeType     *big.Int
	ChallengeContract common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0x6c69257ddf620994c6fb9e5304db0e5563db3765bee033ddd61b6a1caa7d043f.
//
// Solidity: event RollupChallengeStarted(address asserter, address challenger, uint256 challengeType, address challengeContract)
func (_Staking *StakingFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts) (*StakingRollupChallengeStartedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RollupChallengeStarted")
	if err != nil {
		return nil, err
	}
	return &StakingRollupChallengeStartedIterator{contract: _Staking.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0x6c69257ddf620994c6fb9e5304db0e5563db3765bee033ddd61b6a1caa7d043f.
//
// Solidity: event RollupChallengeStarted(address asserter, address challenger, uint256 challengeType, address challengeContract)
func (_Staking *StakingFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *StakingRollupChallengeStarted) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RollupChallengeStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRollupChallengeStarted)
				if err := _Staking.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0x6c69257ddf620994c6fb9e5304db0e5563db3765bee033ddd61b6a1caa7d043f.
//
// Solidity: event RollupChallengeStarted(address asserter, address challenger, uint256 challengeType, address challengeContract)
func (_Staking *StakingFilterer) ParseRollupChallengeStarted(log types.Log) (*StakingRollupChallengeStarted, error) {
	event := new(StakingRollupChallengeStarted)
	if err := _Staking.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRollupStakeCreatedIterator is returned from FilterRollupStakeCreated and is used to iterate over the raw logs and unpacked data for RollupStakeCreated events raised by the Staking contract.
type StakingRollupStakeCreatedIterator struct {
	Event *StakingRollupStakeCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingRollupStakeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRollupStakeCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingRollupStakeCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingRollupStakeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRollupStakeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRollupStakeCreated represents a RollupStakeCreated event raised by the Staking contract.
type StakingRollupStakeCreated struct {
	Staker   common.Address
	NodeHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRollupStakeCreated is a free log retrieval operation binding the contract event 0xcbafbb223ed21c82af9e2ad20cdfdf55d3263d06f9a65b3f70da613f32d81f88.
//
// Solidity: event RollupStakeCreated(address staker, bytes32 nodeHash)
func (_Staking *StakingFilterer) FilterRollupStakeCreated(opts *bind.FilterOpts) (*StakingRollupStakeCreatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RollupStakeCreated")
	if err != nil {
		return nil, err
	}
	return &StakingRollupStakeCreatedIterator{contract: _Staking.contract, event: "RollupStakeCreated", logs: logs, sub: sub}, nil
}

// WatchRollupStakeCreated is a free log subscription operation binding the contract event 0xcbafbb223ed21c82af9e2ad20cdfdf55d3263d06f9a65b3f70da613f32d81f88.
//
// Solidity: event RollupStakeCreated(address staker, bytes32 nodeHash)
func (_Staking *StakingFilterer) WatchRollupStakeCreated(opts *bind.WatchOpts, sink chan<- *StakingRollupStakeCreated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RollupStakeCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRollupStakeCreated)
				if err := _Staking.contract.UnpackLog(event, "RollupStakeCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStakeCreated is a log parse operation binding the contract event 0xcbafbb223ed21c82af9e2ad20cdfdf55d3263d06f9a65b3f70da613f32d81f88.
//
// Solidity: event RollupStakeCreated(address staker, bytes32 nodeHash)
func (_Staking *StakingFilterer) ParseRollupStakeCreated(log types.Log) (*StakingRollupStakeCreated, error) {
	event := new(StakingRollupStakeCreated)
	if err := _Staking.contract.UnpackLog(event, "RollupStakeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRollupStakeMovedIterator is returned from FilterRollupStakeMoved and is used to iterate over the raw logs and unpacked data for RollupStakeMoved events raised by the Staking contract.
type StakingRollupStakeMovedIterator struct {
	Event *StakingRollupStakeMoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingRollupStakeMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRollupStakeMoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingRollupStakeMoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingRollupStakeMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRollupStakeMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRollupStakeMoved represents a RollupStakeMoved event raised by the Staking contract.
type StakingRollupStakeMoved struct {
	Staker     common.Address
	ToNodeHash [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRollupStakeMoved is a free log retrieval operation binding the contract event 0xbe690ac5fe353c094bcc6f187eeb841c0ca61b6edf32c142eadad655b7d173f4.
//
// Solidity: event RollupStakeMoved(address staker, bytes32 toNodeHash)
func (_Staking *StakingFilterer) FilterRollupStakeMoved(opts *bind.FilterOpts) (*StakingRollupStakeMovedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RollupStakeMoved")
	if err != nil {
		return nil, err
	}
	return &StakingRollupStakeMovedIterator{contract: _Staking.contract, event: "RollupStakeMoved", logs: logs, sub: sub}, nil
}

// WatchRollupStakeMoved is a free log subscription operation binding the contract event 0xbe690ac5fe353c094bcc6f187eeb841c0ca61b6edf32c142eadad655b7d173f4.
//
// Solidity: event RollupStakeMoved(address staker, bytes32 toNodeHash)
func (_Staking *StakingFilterer) WatchRollupStakeMoved(opts *bind.WatchOpts, sink chan<- *StakingRollupStakeMoved) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RollupStakeMoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRollupStakeMoved)
				if err := _Staking.contract.UnpackLog(event, "RollupStakeMoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStakeMoved is a log parse operation binding the contract event 0xbe690ac5fe353c094bcc6f187eeb841c0ca61b6edf32c142eadad655b7d173f4.
//
// Solidity: event RollupStakeMoved(address staker, bytes32 toNodeHash)
func (_Staking *StakingFilterer) ParseRollupStakeMoved(log types.Log) (*StakingRollupStakeMoved, error) {
	event := new(StakingRollupStakeMoved)
	if err := _Staking.contract.UnpackLog(event, "RollupStakeMoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingRollupStakeRefundedIterator is returned from FilterRollupStakeRefunded and is used to iterate over the raw logs and unpacked data for RollupStakeRefunded events raised by the Staking contract.
type StakingRollupStakeRefundedIterator struct {
	Event *StakingRollupStakeRefunded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakingRollupStakeRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingRollupStakeRefunded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakingRollupStakeRefunded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakingRollupStakeRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRollupStakeRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingRollupStakeRefunded represents a RollupStakeRefunded event raised by the Staking contract.
type StakingRollupStakeRefunded struct {
	Staker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRollupStakeRefunded is a free log retrieval operation binding the contract event 0x953ab9eece73c907353307064109cf873462177a0e358e463fd89f5b206daa6c.
//
// Solidity: event RollupStakeRefunded(address staker)
func (_Staking *StakingFilterer) FilterRollupStakeRefunded(opts *bind.FilterOpts) (*StakingRollupStakeRefundedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "RollupStakeRefunded")
	if err != nil {
		return nil, err
	}
	return &StakingRollupStakeRefundedIterator{contract: _Staking.contract, event: "RollupStakeRefunded", logs: logs, sub: sub}, nil
}

// WatchRollupStakeRefunded is a free log subscription operation binding the contract event 0x953ab9eece73c907353307064109cf873462177a0e358e463fd89f5b206daa6c.
//
// Solidity: event RollupStakeRefunded(address staker)
func (_Staking *StakingFilterer) WatchRollupStakeRefunded(opts *bind.WatchOpts, sink chan<- *StakingRollupStakeRefunded) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "RollupStakeRefunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingRollupStakeRefunded)
				if err := _Staking.contract.UnpackLog(event, "RollupStakeRefunded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRollupStakeRefunded is a log parse operation binding the contract event 0x953ab9eece73c907353307064109cf873462177a0e358e463fd89f5b206daa6c.
//
// Solidity: event RollupStakeRefunded(address staker)
func (_Staking *StakingFilterer) ParseRollupStakeRefunded(log types.Log) (*StakingRollupStakeRefunded, error) {
	event := new(StakingRollupStakeRefunded)
	if err := _Staking.contract.UnpackLog(event, "RollupStakeRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VMABI is the input ABI used to generate the binding from.
const VMABI = "[]"

// VMBin is the compiled bytecode used for deploying new contracts.
var VMBin = "0x60556023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea265627a7a723158207a2f4d76c7ac5af065616b74405a4a9f6a6aafe38708dfab7489b4acf3131b9864736f6c63430005110032"

// DeployVM deploys a new Ethereum contract, binding an instance of VM to it.
func DeployVM(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VM, error) {
	parsed, err := abi.JSON(strings.NewReader(VMABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VMBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VM{VMCaller: VMCaller{contract: contract}, VMTransactor: VMTransactor{contract: contract}, VMFilterer: VMFilterer{contract: contract}}, nil
}

// VM is an auto generated Go binding around an Ethereum contract.
type VM struct {
	VMCaller     // Read-only binding to the contract
	VMTransactor // Write-only binding to the contract
	VMFilterer   // Log filterer for contract events
}

// VMCaller is an auto generated read-only Go binding around an Ethereum contract.
type VMCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VMTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VMFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VMSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VMSession struct {
	Contract     *VM               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VMCallerSession struct {
	Contract *VMCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VMTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VMTransactorSession struct {
	Contract     *VMTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VMRaw is an auto generated low-level Go binding around an Ethereum contract.
type VMRaw struct {
	Contract *VM // Generic contract binding to access the raw methods on
}

// VMCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VMCallerRaw struct {
	Contract *VMCaller // Generic read-only contract binding to access the raw methods on
}

// VMTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VMTransactorRaw struct {
	Contract *VMTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVM creates a new instance of VM, bound to a specific deployed contract.
func NewVM(address common.Address, backend bind.ContractBackend) (*VM, error) {
	contract, err := bindVM(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VM{VMCaller: VMCaller{contract: contract}, VMTransactor: VMTransactor{contract: contract}, VMFilterer: VMFilterer{contract: contract}}, nil
}

// NewVMCaller creates a new read-only instance of VM, bound to a specific deployed contract.
func NewVMCaller(address common.Address, caller bind.ContractCaller) (*VMCaller, error) {
	contract, err := bindVM(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VMCaller{contract: contract}, nil
}

// NewVMTransactor creates a new write-only instance of VM, bound to a specific deployed contract.
func NewVMTransactor(address common.Address, transactor bind.ContractTransactor) (*VMTransactor, error) {
	contract, err := bindVM(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VMTransactor{contract: contract}, nil
}

// NewVMFilterer creates a new log filterer instance of VM, bound to a specific deployed contract.
func NewVMFilterer(address common.Address, filterer bind.ContractFilterer) (*VMFilterer, error) {
	contract, err := bindVM(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VMFilterer{contract: contract}, nil
}

// bindVM binds a generic wrapper to an already deployed contract.
func bindVM(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VMABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VM *VMRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VM.Contract.VMCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VM *VMRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VM.Contract.VMTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VM *VMRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VM.Contract.VMTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VM *VMCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VM.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VM *VMTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VM.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VM *VMTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VM.Contract.contract.Transact(opts, method, params...)
}
