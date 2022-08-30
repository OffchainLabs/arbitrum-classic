// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgecontracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// RollupAdminFacetMetaData contains all meta data concerning the RollupAdminFacet contract.
var RollupAdminFacetMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenge\",\"type\":\"address\"}],\"name\":\"ChallengeDestroyedInMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterSendAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"NodeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentNodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterInboxBatchEndCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterInboxBatchAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"indexed\":false,\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"NodeDestroyedInMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"NodeRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"StakerWithdrawnInMigration\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalBalance\",\"type\":\"uint256\"}],\"name\":\"UserStakeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"initialBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"finalBalance\",\"type\":\"uint256\"}],\"name\":\"UserWithdrawableFundsUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"STORAGE_GAP_1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STORAGE_GAP_2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"avmGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeExecutionBisectionDegree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedBridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraChallengeTimeBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"beforeSendAcc\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"sendsData\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"sendLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"forceConfirmNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expectedNodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"},{\"internalType\":\"bytes\",\"name\":\"sequencerBatchProof\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"beforeProposedBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beforeInboxMaxCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevNode\",\"type\":\"uint256\"}],\"name\":\"forceCreateNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"staker\",\"type\":\"address[]\"}],\"name\":\"forceRefundStaker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"stakerA\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"stakerB\",\"type\":\"address[]\"}],\"name\":\"forceResolveChallenge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isNitroReady\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isZombie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEventBridge\",\"outputs\":[{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerBridge\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newAvmGasSpeedLimitPerBlock\",\"type\":\"uint256\"}],\"name\":\"setAvmGasSpeedLimitPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBaseStake\",\"type\":\"uint256\"}],\"name\":\"setBaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChallengeExecutionBisectionDegree\",\"type\":\"uint256\"}],\"name\":\"setChallengeExecutionBisectionDegree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newConfirmPeriod\",\"type\":\"uint256\"}],\"name\":\"setConfirmPeriodBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newExtraTimeBlocks\",\"type\":\"uint256\"}],\"name\":\"setExtraChallengeTimeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdminFacet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newUserFacet\",\"type\":\"address\"}],\"name\":\"setFacets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSequencer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isSequencer\",\"type\":\"bool\"}],\"name\":\"setIsSequencer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"setMinimumAssertionPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSequencerInboxMaxDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newSequencerInboxMaxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"setSequencerInboxMaxDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newStakeToken\",\"type\":\"address\"}],\"name\":\"setStakeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validator\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_val\",\"type\":\"bool[]\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whitelist\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"user\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"val\",\"type\":\"bool[]\"}],\"name\":\"setWhitelistEntries\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"finalNodeNum\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"destroyAlternatives\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"destroyChallenges\",\"type\":\"bool\"}],\"name\":\"shutdownForNitro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdownForNitroBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shutdownForNitroMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractOwnable\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"undoShutdownForNitro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"whitelist\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newWhitelist\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"}],\"name\":\"updateWhitelistConsumers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeBeacon\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506000805460ff19908116600117909155600b80549091169055613fb6806100396000396000f3fe608060405234801561001057600080fd5b50600436106103715760003560e01c80637c75c298116101d5578063cf47bb8411610105578063e4781e10116100a8578063e4781e1014610c4c578063e8bd492214610c54578063ef40a67014610cb0578063f33e1fac14610cd6578063f38c937914610cf3578063f51de41b14610e16578063f53f5afa14610e1e578063f8d1f19414610ef4578063ff204f3b14610f1157610371565b8063cf47bb8414610aa6578063d01e660214610bd9578063d735e21d14610bf6578063d7445bc814610bfe578063d93fe9c414610c06578063dc72a33b14610c0e578063dff6978714610c16578063e45b7ce614610c1e57610371565b80639161d535116101785780639161d5351461086057806391c657e81461087d578063948d6588146108a35780639e8a713f146108c05780639ea28e65146108c8578063a3ffb77214610956578063a5cc82f814610a79578063a8929e0b14610a96578063ce11e6ab14610a9e57610371565b80637c75c298146107615780637e6c255f146108025780637f4320ce1461080a5780637f60abbb146108125780638456cb591461081a578063848bf918146108225780638640ce5f146108505780638da5cb5b1461085857610371565b806351ed6a30116102b057806365f7f80d1161025357806365f7f80d1461060e578063661d27221461061657806369fd251c146106d05780636aef131a146106f65780636d435421146107135780636f791d291461074157806376e7e23b14610749578063771b2f97146107515780637ba9534a1461075957610371565b806351ed6a301461057d578063567ca41b146105855780635c975abb146105ab5780635dbaf68b146105b35780635e8ef106146105bb5780636177fd18146105c357806362a82d7d146105e957806363721d6b1461060657610371565b80632f30cabd116103185780632f30cabd1461048c578063313a04fa146104b25780633e55c0c7146104ce5780633e96576e146104f25780633ea410981461051857806340b570f41461053557806345e38b64146105585780634f0f4aa91461056057610371565b80630397d45814610376578063046f7da21461039e57806306ae5851146103a657806313af4035146103c35780631d0ada65146103e95780631f9566321461041657806327035859146104445780632e7acfa614610472575b600080fd5b61039c6004803603602081101561038c57600080fd5b50356001600160a01b0316610f37565b005b61039c610f6f565b61039c600480360360208110156103bc57600080fd5b5035610f93565b61039c600480360360208110156103d957600080fd5b50356001600160a01b0316610fb5565b61039c600480360360608110156103ff57600080fd5b508035906020810135151590604001351515610fed565b61039c6004803603604081101561042c57600080fd5b506001600160a01b0381351690602001351515611538565b61039c6004803603604081101561045a57600080fd5b506001600160a01b03813581169160200135166115c4565b61047a611654565b60408051918252519081900360200190f35b61047a600480360360208110156104a257600080fd5b50356001600160a01b031661165a565b6104ba611679565b604080519115158252519081900360200190f35b6104d6611682565b604080516001600160a01b039092168252519081900360200190f35b61047a6004803603602081101561050857600080fd5b50356001600160a01b0316611691565b61039c6004803603602081101561052e57600080fd5b50356116af565b61039c6004803603604081101561054b57600080fd5b50803590602001356116d1565b61047a61175b565b6104d66004803603602081101561057657600080fd5b5035611761565b6104d661177c565b61039c6004803603602081101561059b57600080fd5b50356001600160a01b031661178b565b6104ba611863565b6104d661186c565b61047a61187b565b6104ba600480360360208110156105d957600080fd5b50356001600160a01b0316611881565b6104d6600480360360208110156105ff57600080fd5b50356118a9565b61047a6118d3565b61047a6118d9565b61039c6004803603606081101561062c57600080fd5b6001600160a01b038235811692602081013590911691810190606081016040820135600160201b81111561065f57600080fd5b82018360208201111561067157600080fd5b803590602001918460208302840111600160201b8311171561069257600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506118df945050505050565b6104d6600480360360208110156106e657600080fd5b50356001600160a01b03166119a5565b61039c6004803603602081101561070c57600080fd5b50356119c6565b61039c6004803603604081101561072957600080fd5b506001600160a01b03813581169160200135166119e8565b6104ba611a76565b61047a611a7f565b61047a611a85565b61047a611a8b565b61039c6004803603602081101561077757600080fd5b810190602081018135600160201b81111561079157600080fd5b8201836020820111156107a357600080fd5b803590602001918460208302840111600160201b831117156107c457600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611a91945050505050565b61039c611b45565b61047a611c03565b61047a611c09565b61039c611c0f565b61039c6004803603604081101561083857600080fd5b506001600160a01b0381358116916020013516611c33565b61047a611cc1565b6104d6611cc7565b61039c6004803603602081101561087657600080fd5b5035611cd6565b6104ba6004803603602081101561089357600080fd5b50356001600160a01b0316611cf8565b61039c600480360360208110156108b957600080fd5b5035611d52565b6104d6611d74565b61039c60048036036102608110156108df57600080fd5b813591602081019160e08201919081019061020081016101e0820135600160201b81111561090c57600080fd5b82018360208201111561091e57600080fd5b803590602001918460018302840111600160201b8311171561093f57600080fd5b919350915080359060208101359060400135611d83565b61039c6004803603604081101561096c57600080fd5b810190602081018135600160201b81111561098657600080fd5b82018360208201111561099857600080fd5b803590602001918460208302840111600160201b831117156109b957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610a0857600080fd5b820183602082011115610a1a57600080fd5b803590602001918460208302840111600160201b83111715610a3b57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550611fda945050505050565b61039c60048036036020811015610a8f57600080fd5b50356120a8565b61047a6120ca565b6104d66120d0565b61039c60048036036060811015610abc57600080fd5b6001600160a01b038235169190810190604081016020820135600160201b811115610ae657600080fd5b820183602082011115610af857600080fd5b803590602001918460208302840111600160201b83111715610b1957600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610b6857600080fd5b820183602082011115610b7a57600080fd5b803590602001918460208302840111600160201b83111715610b9b57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506120df945050505050565b6104d660048036036020811015610bef57600080fd5b5035612226565b61047a612255565b61047a61225b565b6104d6612261565b61047a612270565b61047a612276565b61039c60048036036040811015610c3457600080fd5b506001600160a01b038135169060200135151561227c565b61047a612308565b610c7a60048036036020811015610c6a57600080fd5b50356001600160a01b031661230e565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b61047a60048036036020811015610cc657600080fd5b50356001600160a01b031661234a565b61047a60048036036020811015610cec57600080fd5b5035612368565b61039c60048036036040811015610d0957600080fd5b810190602081018135600160201b811115610d2357600080fd5b820183602082011115610d3557600080fd5b803590602001918460208302840111600160201b83111715610d5657600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b811115610da557600080fd5b820183602082011115610db757600080fd5b803590602001918460208302840111600160201b83111715610dd857600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550612390945050505050565b6104d661254f565b61039c600480360360e0811015610e3457600080fd5b813591602081013591810190606081016040820135600160201b811115610e5a57600080fd5b820183602082011115610e6c57600080fd5b803590602001918460018302840111600160201b83111715610e8d57600080fd5b919390929091602081019035600160201b811115610eaa57600080fd5b820183602082011115610ebc57600080fd5b803590602001918460208302840111600160201b83111715610edd57600080fd5b91935091508035906020810135906040013561255e565b61047a60048036036020811015610f0a57600080fd5b50356125f8565b61039c60048036036020811015610f2757600080fd5b50356001600160a01b031661260a565b601780546001600160a01b0319166001600160a01b038316179055604051600d90600080516020613f6183398151915290600090a250565b610f776126a8565b604051600490600080516020613f6183398151915290600090a2565b600f819055604051600c90600080516020613f6183398151915290600090a250565b601680546001600160a01b0319166001600160a01b038316179055604051600790600080516020613f6183398151915290600090a250565b610ff5611863565b1561103a576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b611042611679565b1561108c576040805162461bcd60e51b8152602060048201526015602482015274414c52454144595f53485554444f574e5f4d4f444560581b604482015290519081900360640190fd5b60006110966118d9565b905060006110a2611a8b565b9050845b8282146111c757808214156111305760006110c083611761565b9050806001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156110fb57600080fd5b505afa15801561110f573d6000803e3d6000fd5b505050506040513d602081101561112557600080fd5b505191506111bb9050565b8461117e576040805162461bcd60e51b8152602060048201526019602482015278105315115493905512559154d7d393d517d156141150d51151603a1b604482015290519081900360640190fd5b6111878261273f565b6040805183815290517fc48f1661fe65917dbe9d175ac4cb62063ef44afe989dcd3dbf470ac5a1c77bcb9181900360200190a15b600019909101906110a6565b60006111d1612276565b905060608167ffffffffffffffff811180156111ec57600080fd5b50604051908082528060200260200182016040528015611216578160200160208202803683370190505b50905060005b828167ffffffffffffffff161015611277576112418167ffffffffffffffff166118a9565b828267ffffffffffffffff168151811061125757fe5b6001600160a01b039092166020928302919091019091015260010161121c565b5060005b828167ffffffffffffffff161015611507576000828267ffffffffffffffff16815181106112a557fe5b6020026020010151905060006112ba826119a5565b90506001600160a01b038116156114955788611316576040805162461bcd60e51b815260206004820152601660248201527510d2105313115391d157d393d517d156141150d5115160521b604482015290519081900360640190fd5b6000816001600160a01b031663bb4af0b16040518163ffffffff1660e01b815260040160206040518083038186803b15801561135157600080fd5b505afa158015611365573d6000803e3d6000fd5b505050506040513d602081101561137b57600080fd5b5051604080516329a6d87160e11b815290519192506000916001600160a01b0385169163534db0e2916004808301926020929190829003018186803b1580156113c357600080fd5b505afa1580156113d7573d6000803e3d6000fd5b505050506040513d60208110156113ed57600080fd5b505190506113fa826127c1565b611403816127c1565b826001600160a01b03166214ebe76040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561143d57600080fd5b505af1158015611451573d6000803e3d6000fd5b5050604080516001600160a01b038716815290517fec0b2d27905d678228bae2ed41ea32ea8bdbdd62e81edec23035cbde03f122a29350908190036020019150a150505b60006114a86114a384611691565b611761565b6001600160a01b031614156114fd576114c0826127eb565b604080516001600160a01b038416815290517fe695a52cb984a997f48f43d18e30b3ea892d024ca5410f74e5fded60b4a033ef9181900360200190a15b505060010161127b565b5043601e55611514612863565b604051601990600080516020613f6183398151915290600090a25050505050505050565b60115460408051630fcab31960e11b81526001600160a01b038581166004830152841515602483015291519190921691631f95663291604480830192600092919082900301818387803b15801561158e57600080fd5b505af11580156115a2573d6000803e3d6000fd5b505060405160139250600080516020613f618339815191529150600090a25050565b81601c6000815481106115d357fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080601c60018154811061161057fe5b6000918252602082200180546001600160a01b0319166001600160a01b039390931692909217909155604051600591600080516020613f6183398151915291a25050565b600c5481565b6001600160a01b0381166000908152600a60205260409020545b919050565b601e5443101590565b6011546001600160a01b031681565b6001600160a01b031660009081526008602052604090206001015490565b600c819055604051600990600080516020613f6183398151915290600090a250565b601154604080516326a407d560e11b8152600481018590526024810184905290516001600160a01b0390921691634d480faa9160448082019260009290919082900301818387803b15801561172557600080fd5b505af1158015611739573d6000803e3d6000fd5b5050604051600e9250600080516020613f618339815191529150600090a25050565b60185481565b6000908152600560205260409020546001600160a01b031690565b6017546001600160a01b031681565b6012546001600160a01b03828116911614156117db576040805162461bcd60e51b815260206004820152600a602482015269086aaa4be9eaaa8849eb60b31b604482015290519081900360640190fd5b601054604080516319dc7ae560e31b81526001600160a01b038481166004830152600060248301819052925193169263cee3d7289260448084019391929182900301818387803b15801561182e57600080fd5b505af1158015611842573d6000803e3d6000fd5b505060405160019250600080516020613f618339815191529150600090a250565b600b5460ff1690565b6014546001600160a01b031681565b600e5490565b6001600160a01b0316600090815260086020526040902060030154600160a01b900460ff1690565b6000600782815481106118b857fe5b6000918252602090912001546001600160a01b031692915050565b60095490565b60015490565b604080516337ca261760e01b81526001600160a01b038481166004830190815260248301938452845160448401528451918716936337ca261793879387939291606401906020808601910280838360005b83811015611948578181015183820152602001611930565b505050509050019350505050600060405180830381600087803b15801561196e57600080fd5b505af1158015611982573d6000803e3d6000fd5b505060405160119250600080516020613f618339815191529150600090a2505050565b6001600160a01b039081166000908152600860205260409020600301541690565b601b819055604051601090600080516020613f6183398151915290600090a250565b816001600160a01b031663f2fde38b826040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050600060405180830381600087803b158015611a4057600080fd5b505af1158015611a54573d6000803e3d6000fd5b5050604051601b9250600080516020613f618339815191529150600090a25050565b60005460ff1690565b600f5481565b600d5481565b60035490565b611a99611863565b611ad8576040805162461bcd60e51b81526020600482015260146024820152600080516020613f41833981519152604482015290519081900360640190fd5b60005b8151811015611b2757611b02828281518110611af357fe5b602002602001015160006128e6565b50611b1f828281518110611b1257fe5b60200260200101516129bc565b600101611adb565b50604051601690600080516020613f6183398151915290600090a250565b611b4d611863565b611b8c576040805162461bcd60e51b81526020600482015260146024820152600080516020613f41833981519152604482015290519081900360640190fd5b611b94611679565b611bd9576040805162461bcd60e51b81526020600482015260116024820152704e4f545f53485554444f574e5f4d4f444560781b604482015290519081900360640190fd5b600019601e55611be76126a8565b604051601a90600080516020613f6183398151915290600090a2565b601a5481565b601e5481565b611c17612863565b604051600390600080516020613f6183398151915290600090a2565b816001600160a01b0316633659cfe6826040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050600060405180830381600087803b158015611c8b57600080fd5b505af1158015611c9f573d6000803e3d6000fd5b505060405160149250600080516020613f618339815191529150600090a25050565b60045490565b6016546001600160a01b031681565b600d819055604051600a90600080516020613f6183398151915290600090a250565b6000805b600954811015611d495760098181548110611d1357fe5b60009182526020909120600290910201546001600160a01b0384811691161415611d41576001915050611674565b600101611cfc565b50600092915050565b6018819055604051600890600080516020613f6183398151915290600090a250565b6013546001600160a01b031681565b611d8b611863565b611dca576040805162461bcd60e51b81526020600482015260146024820152600080516020613f41833981519152604482015290519081900360640190fd5b611dd26118d9565b8114611e1d576040805162461bcd60e51b815260206004820152601560248201527413d3931657d310551154d517d0d3d3919254935151605a1b604482015290519081900360640190fd5b611e25613e72565b60408051808201909152611f58908960026000835b82821015611e7b5760408051606081810190925290808402860190600390839083908082843760009201919091525050508152600190910190602001611e3a565b50506040805180820190915291508a905060026000835b82821015611ed35760408051608081810190925290808402860190600490839083908082843760009201919091525050508152600190910190602001611e92565b505050508686601160009054906101000a90046001600160a01b03166001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b158015611f2757600080fd5b505afa158015611f3b573d6000803e3d6000fd5b505050506040513d6020811015611f5157600080fd5b5051612a70565b6040805160c081018252848152600c546020820152600e54918101919091526011546001600160a01b039081166060830152601354811660808301526015541660a0820152909150611fb49082908a908a908a908a908f612abe565b50604051601790600080516020613f6183398151915290600090a2505050505050505050565b805182511461201f576040805162461bcd60e51b815260206004820152600c60248201526b0aea49e9c8ebe988a9c8ea8960a31b604482015290519081900360640190fd5b60005b82518110156120895781818151811061203757fe5b6020026020010151601d600085848151811061204f57fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055600101612022565b50604051600690600080516020613f6183398151915290600090a25050565b600e819055604051600b90600080516020613f6183398151915290600090a250565b61a4b290565b6012546001600160a01b031681565b8051825114612125576040805162461bcd60e51b815260206004820152600d60248201526c1253959053125117d253941555609a1b604482015290519081900360640190fd5b60408051633b99adf760e01b8152600481019182528351604482015283516001600160a01b03861692633b99adf792869286929182916024820191606401906020808801910280838360005b83811015612189578181015183820152602001612171565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156121c85781810151838201526020016121b0565b50505050905001945050505050600060405180830381600087803b1580156121ef57600080fd5b505af1158015612203573d6000803e3d6000fd5b505060405160129250600080516020613f618339815191529150600090a2505050565b60006009828154811061223557fe5b60009182526020909120600290910201546001600160a01b031692915050565b60025490565b600e5481565b6015546001600160a01b031681565b601b5481565b60075490565b6010546040805163722dbe7360e11b81526001600160a01b03858116600483015284151560248301529151919092169163e45b7ce691604480830192600092919082900301818387803b1580156122d257600080fd5b505af11580156122e6573d6000803e3d6000fd5b505060405160029250600080516020613f618339815191529150600090a25050565b60195481565b6008602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6001600160a01b031660009081526008602052604090206002015490565b60006009828154811061237757fe5b9060005260206000209060020201600101549050919050565b612398611863565b6123d7576040805162461bcd60e51b81526020600482015260146024820152600080516020613f41833981519152604482015290519081900360640190fd5b805182511461241c576040805162461bcd60e51b815260206004820152600c60248201526b0aea49e9c8ebe988a9c8ea8960a31b604482015290519081900360640190fd5b60005b825181101561253057600061245a84838151811061243957fe5b602002602001015184848151811061244d57fe5b6020026020010151613178565b90506001600160a01b0381166124a6576040805162461bcd60e51b815260206004820152600c60248201526b1393d517d25397d0d210531360a21b604482015290519081900360640190fd5b6124c28483815181106124b557fe5b60200260200101516127c1565b6124d18383815181106124b557fe5b806001600160a01b03166214ebe76040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561250b57600080fd5b505af115801561251f573d6000803e3d6000fd5b50506001909301925061241f915050565b50604051601590600080516020613f6183398151915290600090a25050565b6010546001600160a01b031681565b612566611863565b6125a5576040805162461bcd60e51b81526020600482015260146024820152600080516020613f41833981519152604482015290519081900360640190fd5b6012546013546125d3918b918b918b918b918b918b918b918b918b916001600160a01b03918216911661323e565b604051601890600080516020613f6183398151915290600090a2505050505050505050565b60009081526006602052604090205490565b601280546001600160a01b0319166001600160a01b03838116918217909255601054604080516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b15801561267457600080fd5b505af1158015612688573d6000803e3d6000fd5b505060405160009250600080516020613f6183398151915291508290a250565b6126b0611863565b6126ef576040805162461bcd60e51b81526020600482015260146024820152600080516020613f41833981519152604482015290519081900360640190fd5b600b805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa61272261350f565b604080516001600160a01b039092168252519081900360200190a1565b60008181526005602052604080822054815163083197ef60e41b815291516001600160a01b03909116926383197ef0926004808201939182900301818387803b15801561278b57600080fd5b505af115801561279f573d6000803e3d6000fd5b50505060009182525060056020526040902080546001600160a01b0319169055565b6001600160a01b0316600090815260086020526040902060030180546001600160a01b0319169055565b6001600160a01b038116600090815260086020526040902060028101546128128382613513565b61281b8361359e565b604080518281526000602082015281516001600160a01b038616927febd093d389ab57f3566918d2c379a2b4d9539e8eb95efad9d5e465457833fde6928290030190a2505050565b61286b611863565b156128b0576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b600b805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861272261350f565b6001600160a01b038216600090815260086020526040812060028101548084111561294b576040805162461bcd60e51b815260206004820152601060248201526f544f4f5f4c4954544c455f5354414b4560801b604482015290519081900360640190fd5b600061295d828663ffffffff6136c416565b6002840186905590506129708682613513565b604080518381526020810187905281516001600160a01b038916927febd093d389ab57f3566918d2c379a2b4d9539e8eb95efad9d5e465457833fde6928290030190a295945050505050565b6001600160a01b0381811660008181526008602090815260408083208151808301909252938152600180850154928201928352600980549182018155909352517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af600290930292830180546001600160a01b031916919095161790935591517f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7b090920191909155612a6c8261359e565b5050565b612a78613e72565b60408051808201909152865186518291612a93918888613721565b8152602001612ab2886001602002015188600160200201514387613721565b90529695505050505050565b6000612ac8613e97565b612ad1896137bf565b60e08201528351612ae190611761565b81606001906001600160a01b031690816001600160a01b03168152505083606001516001600160a01b0316633dbcc8d16040518163ffffffff1660e01b815260040160206040518083038186803b158015612b3b57600080fd5b505afa158015612b4f573d6000803e3d6000fd5b505050506040513d6020811015612b6557600080fd5b5051815260608101516040805163380ed4c760e11b815290516001600160a01b039092169163701da98e91600480820192602092909190829003018186803b158015612bb057600080fd5b505afa158015612bc4573d6000803e3d6000fd5b505050506040513d6020811015612bda57600080fd5b50518951612be7906137e1565b14612c2b576040805162461bcd60e51b815260206004820152600f60248201526e0a0a48aacbea6a882a88abe9082a69608b1b604482015290519081900360640190fd5b805160208a0151604001511115612c7a576040805162461bcd60e51b815260206004820152600e60248201526d12539093d617d41054d517d1539160921b604482015290519081900360640190fd5b83606001516001600160a01b031663dc1b7b1f87878c60200151604001516040518463ffffffff1660e01b815260040180806020018381526020018281038252858582818152602001925080828437600081840152601f19601f820116905080830192505050945050505050604080518083038186803b158015612cfd57600080fd5b505afa158015612d11573d6000803e3d6000fd5b505050506040513d6040811015612d2757600080fd5b508051602090910151610120830152610100820152612d4589613876565b816040018181525050612d6a84604001518260e00151866020015184606001516138a7565b8160c0018181525050600081606001516001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b158015612db257600080fd5b505afa158015612dc6573d6000803e3d6000fd5b505050506040513d6020811015612ddc57600080fd5b50511160a0820181905215612e6657612e5c81606001516001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b158015612e2b57600080fd5b505afa158015612e3f573d6000803e3d6000fd5b505050506040513d6020811015612e5557600080fd5b50516125f8565b6080820152612e77565b8351612e71906125f8565b60808201525b8360a001516001600160a01b031663d45ab2b5612e978b602001516137e1565b612ea68c856040015143613a2d565b612eaf8d613a4a565b88600001518660c001516040518663ffffffff1660e01b81526004018086815260200185815260200184815260200183815260200182815260200195505050505050602060405180830381600087803b158015612f0b57600080fd5b505af1158015612f1f573d6000803e3d6000fd5b505050506040513d6020811015612f3557600080fd5b50516001600160a01b031660208201526000612f4f611a8b565b600101905081606001516001600160a01b0316631bc09d0a826040518263ffffffff1660e01b815260040180828152602001915050600060405180830381600087803b158015612f9e57600080fd5b505af1158015612fb2573d6000803e3d6000fd5b50505050612fd38260a0015183608001518460400151856101200151613a7a565b9250838314613020576040805162461bcd60e51b81526020600482015260146024820152730aa9c8ab0a08a86a88a88be9c9e888abe9082a6960631b604482015290519081900360640190fd5b61302e826020015184613ae1565b6080850151855160c084015160408051638b8ca19960e01b81526004810186905260248101939093526044830191909152336064830152516001600160a01b0390921691638b8ca1999160848082019260009290919082900301818387803b15801561309957600080fd5b505af11580156130ad573d6000803e3d6000fd5b50505050506130bf84600001516125f8565b6130c7611a8b565b7f8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a5484846040015185600001518661010001518761012001518f8f6040518088815260200187815260200186815260200185815260200184815260200183600260600280828437600083820152601f01601f191690910190508261010080828437600083820152604051601f909101601f1916909201829003995090975050505050505050a350979650505050505050565b6001600160a01b038083166000908152600860205260408082208484168352908220600382015492939192909116806131e2576040805162461bcd60e51b81526020600482015260076024820152661393d7d0d2105360ca1b604482015290519081900360640190fd5b60038201546001600160a01b03828116911614613235576040805162461bcd60e51b815260206004820152600c60248201526b1112519197d25397d0d2105360a21b604482015290519081900360640190fd5b95945050505050565b60006132bf8a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020808e0282810182019093528d82529093508d92508c918291850190849080828437600081840152601f19601f820116905080830192505050505050508d613b2b565b905060006132cc8d611761565b90506132db8c83888a89613c2c565b816001600160a01b03166397bdc5106040518163ffffffff1660e01b815260040160206040518083038186803b15801561331457600080fd5b505afa158015613328573d6000803e3d6000fd5b505050506040513d602081101561333e57600080fd5b505114613381576040805162461bcd60e51b815260206004820152600c60248201526b434f4e4649524d5f4441544160a01b604482015290519081900360640190fd5b836001600160a01b0316630c7268478c8c8c8c6040518563ffffffff1660e01b81526004018080602001806020018381038352878782818152602001925080828437600083820152601f01601f19169091018481038352858152602090810191508690860280828437600081840152601f19601f8201169050808301925050509650505050505050600060405180830381600087803b15801561342357600080fd5b505af1158015613437573d6000803e3d6000fd5b5050505061344660015461273f565b60018d81558d01600255604080516316b9109b60e01b8152600481018f905290516001600160a01b038516916316b9109b91602480830192600092919082900301818387803b15801561349857600080fd5b505af11580156134ac573d6000803e3d6000fd5b505050508c7f2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2838989896040518085815260200184815260200183815260200182815260200194505050505060405180910390a250505050505050505050505050565b3390565b6001600160a01b0382166000908152600a60205260408120549061353d828463ffffffff613c7316565b6001600160a01b0385166000818152600a60209081526040918290208490558151868152908101849052815193945091927fa740af14c56e4e04a617b1de1eb20de73270decbaaead14f142aabf3038e5ae29281900390910190a250505050565b6001600160a01b038116600090815260086020526040902080546007805460001981019081106135ca57fe5b600091825260209091200154600780546001600160a01b0390921691839081106135f057fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080600860006007848154811061363057fe5b60009182526020808320909101546001600160a01b03168352820192909252604001902055600780548061366057fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b03949094168152600890935250506040812081815560018101829055600281019190915560030180546001600160a81b0319169055565b60008282111561371b576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b613729613eeb565b60408051610120810182528551815286516020820152908101856001602002015181526020018560026004811061375c57fe5b602002015181526020018560036004811061377357fe5b602002015181526020018660016003811061378a57fe5b60200201518152602001866002600381106137a157fe5b60200201518152602001848152602001838152509050949350505050565b8051516020820151516000916137db919063ffffffff6136c416565b92915050565b6000816000015182602001518360400151846060015185608001518660a001518760c001518860e00151896101000151604051602001808a81526020018981526020018881526020018781526020018681526020018581526020018481526020018381526020018281526020019950505050505050505050604051602081830303815290604052805190602001209050919050565b805180516020830151516000926137db92918290039061389590613cd4565b6138a28660200151613cd4565b613d09565b6000806138db866138cf6138c282600163ffffffff6136c416565b889063ffffffff613c7316565b9063ffffffff613d4716565b905061396a8161395e6138f4438863ffffffff613c7316565b866001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561392d57600080fd5b505afa158015613941573d6000803e3d6000fd5b505050506040513d602081101561395757600080fd5b5051613dae565b9063ffffffff613c7316565b91506000836001600160a01b031663f0dd77ff6040518163ffffffff1660e01b815260040160206040518083038186803b1580156139a757600080fd5b505afa1580156139bb573d6000803e3d6000fd5b505050506040513d60208110156139d157600080fd5b505190508015613a2357613a20836139e883611761565b6001600160a01b0316632edfb42a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561392d57600080fd5b92505b5050949350505050565b6000613a428383866020015160400151613dc4565b949350505050565b805160a09081015160208301519182015160c083015160608401516080909401516000946137db94939291613c2c565b60008085613a89576000613a8c565b60015b905080858585604051602001808560ff1660ff1660f81b815260010184815260200183815260200182815260200194505050505060405160208183030381529060405280519060200120915050949350505050565b60038054600101808255600090815260056020908152604080832080546001600160a01b0319166001600160a01b0397909716969096179095559154815260069091529190912055565b81518351600091829184835b83811015613bde576000888281518110613b4d57fe5b60200260200101519050838187011115613b9d576040805162461bcd60e51b815260206004820152600c60248201526b2220aa20afa7ab22a9292aa760a11b604482015290519081900360640190fd5b6020868b0181018290206040805180840196909652858101919091528051808603820181526060909501905283519301929092209190940193600101613b37565b50818414613c21576040805162461bcd60e51b815260206004820152600b60248201526a08882a882be988a9c8ea8960ab1b604482015290519081900360640190fd5b979650505050505050565b60408051602080820197909752808201959095526060850192909252608084019290925260a0808401929092528051808403909201825260c0909201909152805191012090565b600082820183811015613ccd576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b60006137db8260000151613d04846040015185602001518660a0015187606001518860c001518960800151613dfb565b613e46565b604080516020808201969096528082019490945260608401929092526080808401919091528151808403909101815260a09092019052805191012090565b6000808211613d9d576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b818381613da657fe5b049392505050565b6000818311613dbd5781613ccd565b5090919050565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b60408051602080820198909852808201969096526060860194909452608085019290925260a084015260c0808401919091528151808403909101815260e09092019052805191012090565b604080516020808201949094528082019290925280518083038201815260609092019052805191012090565b6040518060400160405280613e85613eeb565b8152602001613e92613eeb565b905290565b6040805161014081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081019190915290565b604051806101200160405280600081526020016000801916815260200160008152602001600081526020016000815260200160008019168152602001600080191681526020016000815260200160008152509056fe5061757361626c653a206e6f7420706175736564000000000000000000000000ea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456ea2646970667358221220e791981d28606560ce5b1b12a67e3d55219868cebbf9fbfd786f0fc78dc6131264736f6c634300060b0033",
}

// RollupAdminFacetABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupAdminFacetMetaData.ABI instead.
var RollupAdminFacetABI = RollupAdminFacetMetaData.ABI

// RollupAdminFacetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupAdminFacetMetaData.Bin instead.
var RollupAdminFacetBin = RollupAdminFacetMetaData.Bin

// DeployRollupAdminFacet deploys a new Ethereum contract, binding an instance of RollupAdminFacet to it.
func DeployRollupAdminFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupAdminFacet, error) {
	parsed, err := RollupAdminFacetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupAdminFacetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupAdminFacet{RollupAdminFacetCaller: RollupAdminFacetCaller{contract: contract}, RollupAdminFacetTransactor: RollupAdminFacetTransactor{contract: contract}, RollupAdminFacetFilterer: RollupAdminFacetFilterer{contract: contract}}, nil
}

// RollupAdminFacet is an auto generated Go binding around an Ethereum contract.
type RollupAdminFacet struct {
	RollupAdminFacetCaller     // Read-only binding to the contract
	RollupAdminFacetTransactor // Write-only binding to the contract
	RollupAdminFacetFilterer   // Log filterer for contract events
}

// RollupAdminFacetCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupAdminFacetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupAdminFacetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupAdminFacetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupAdminFacetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupAdminFacetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupAdminFacetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupAdminFacetSession struct {
	Contract     *RollupAdminFacet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupAdminFacetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupAdminFacetCallerSession struct {
	Contract *RollupAdminFacetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// RollupAdminFacetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupAdminFacetTransactorSession struct {
	Contract     *RollupAdminFacetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// RollupAdminFacetRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupAdminFacetRaw struct {
	Contract *RollupAdminFacet // Generic contract binding to access the raw methods on
}

// RollupAdminFacetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupAdminFacetCallerRaw struct {
	Contract *RollupAdminFacetCaller // Generic read-only contract binding to access the raw methods on
}

// RollupAdminFacetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupAdminFacetTransactorRaw struct {
	Contract *RollupAdminFacetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupAdminFacet creates a new instance of RollupAdminFacet, bound to a specific deployed contract.
func NewRollupAdminFacet(address common.Address, backend bind.ContractBackend) (*RollupAdminFacet, error) {
	contract, err := bindRollupAdminFacet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacet{RollupAdminFacetCaller: RollupAdminFacetCaller{contract: contract}, RollupAdminFacetTransactor: RollupAdminFacetTransactor{contract: contract}, RollupAdminFacetFilterer: RollupAdminFacetFilterer{contract: contract}}, nil
}

// NewRollupAdminFacetCaller creates a new read-only instance of RollupAdminFacet, bound to a specific deployed contract.
func NewRollupAdminFacetCaller(address common.Address, caller bind.ContractCaller) (*RollupAdminFacetCaller, error) {
	contract, err := bindRollupAdminFacet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetCaller{contract: contract}, nil
}

// NewRollupAdminFacetTransactor creates a new write-only instance of RollupAdminFacet, bound to a specific deployed contract.
func NewRollupAdminFacetTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupAdminFacetTransactor, error) {
	contract, err := bindRollupAdminFacet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetTransactor{contract: contract}, nil
}

// NewRollupAdminFacetFilterer creates a new log filterer instance of RollupAdminFacet, bound to a specific deployed contract.
func NewRollupAdminFacetFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupAdminFacetFilterer, error) {
	contract, err := bindRollupAdminFacet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetFilterer{contract: contract}, nil
}

// bindRollupAdminFacet binds a generic wrapper to an already deployed contract.
func bindRollupAdminFacet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupAdminFacetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupAdminFacet *RollupAdminFacetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupAdminFacet.Contract.RollupAdminFacetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupAdminFacet *RollupAdminFacetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.RollupAdminFacetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupAdminFacet *RollupAdminFacetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.RollupAdminFacetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupAdminFacet *RollupAdminFacetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupAdminFacet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupAdminFacet *RollupAdminFacetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupAdminFacet *RollupAdminFacetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.contract.Transact(opts, method, params...)
}

// STORAGEGAP1 is a free data retrieval call binding the contract method 0xe4781e10.
//
// Solidity: function STORAGE_GAP_1() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) STORAGEGAP1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "STORAGE_GAP_1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STORAGEGAP1 is a free data retrieval call binding the contract method 0xe4781e10.
//
// Solidity: function STORAGE_GAP_1() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) STORAGEGAP1() (*big.Int, error) {
	return _RollupAdminFacet.Contract.STORAGEGAP1(&_RollupAdminFacet.CallOpts)
}

// STORAGEGAP1 is a free data retrieval call binding the contract method 0xe4781e10.
//
// Solidity: function STORAGE_GAP_1() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) STORAGEGAP1() (*big.Int, error) {
	return _RollupAdminFacet.Contract.STORAGEGAP1(&_RollupAdminFacet.CallOpts)
}

// STORAGEGAP2 is a free data retrieval call binding the contract method 0x7f4320ce.
//
// Solidity: function STORAGE_GAP_2() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) STORAGEGAP2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "STORAGE_GAP_2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STORAGEGAP2 is a free data retrieval call binding the contract method 0x7f4320ce.
//
// Solidity: function STORAGE_GAP_2() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) STORAGEGAP2() (*big.Int, error) {
	return _RollupAdminFacet.Contract.STORAGEGAP2(&_RollupAdminFacet.CallOpts)
}

// STORAGEGAP2 is a free data retrieval call binding the contract method 0x7f4320ce.
//
// Solidity: function STORAGE_GAP_2() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) STORAGEGAP2() (*big.Int, error) {
	return _RollupAdminFacet.Contract.STORAGEGAP2(&_RollupAdminFacet.CallOpts)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupAdminFacet *RollupAdminFacetCaller) StakerMap(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "_stakerMap", arg0)

	outstruct := new(struct {
		Index            *big.Int
		LatestStakedNode *big.Int
		AmountStaked     *big.Int
		CurrentChallenge common.Address
		IsStaked         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Index = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LatestStakedNode = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentChallenge = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.IsStaked = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupAdminFacet *RollupAdminFacetSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _RollupAdminFacet.Contract.StakerMap(&_RollupAdminFacet.CallOpts, arg0)
}

// StakerMap is a free data retrieval call binding the contract method 0xe8bd4922.
//
// Solidity: function _stakerMap(address ) view returns(uint256 index, uint256 latestStakedNode, uint256 amountStaked, address currentChallenge, bool isStaked)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) StakerMap(arg0 common.Address) (struct {
	Index            *big.Int
	LatestStakedNode *big.Int
	AmountStaked     *big.Int
	CurrentChallenge common.Address
	IsStaked         bool
}, error) {
	return _RollupAdminFacet.Contract.StakerMap(&_RollupAdminFacet.CallOpts, arg0)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) AmountStaked(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "amountStaked", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _RollupAdminFacet.Contract.AmountStaked(&_RollupAdminFacet.CallOpts, staker)
}

// AmountStaked is a free data retrieval call binding the contract method 0xef40a670.
//
// Solidity: function amountStaked(address staker) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) AmountStaked(staker common.Address) (*big.Int, error) {
	return _RollupAdminFacet.Contract.AmountStaked(&_RollupAdminFacet.CallOpts, staker)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) ArbGasSpeedLimitPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "arbGasSpeedLimitPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ArbGasSpeedLimitPerBlock(&_RollupAdminFacet.CallOpts)
}

// ArbGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0x5e8ef106.
//
// Solidity: function arbGasSpeedLimitPerBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ArbGasSpeedLimitPerBlock() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ArbGasSpeedLimitPerBlock(&_RollupAdminFacet.CallOpts)
}

// AvmGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0xd7445bc8.
//
// Solidity: function avmGasSpeedLimitPerBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) AvmGasSpeedLimitPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "avmGasSpeedLimitPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvmGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0xd7445bc8.
//
// Solidity: function avmGasSpeedLimitPerBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) AvmGasSpeedLimitPerBlock() (*big.Int, error) {
	return _RollupAdminFacet.Contract.AvmGasSpeedLimitPerBlock(&_RollupAdminFacet.CallOpts)
}

// AvmGasSpeedLimitPerBlock is a free data retrieval call binding the contract method 0xd7445bc8.
//
// Solidity: function avmGasSpeedLimitPerBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) AvmGasSpeedLimitPerBlock() (*big.Int, error) {
	return _RollupAdminFacet.Contract.AvmGasSpeedLimitPerBlock(&_RollupAdminFacet.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) BaseStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "baseStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) BaseStake() (*big.Int, error) {
	return _RollupAdminFacet.Contract.BaseStake(&_RollupAdminFacet.CallOpts)
}

// BaseStake is a free data retrieval call binding the contract method 0x76e7e23b.
//
// Solidity: function baseStake() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) BaseStake() (*big.Int, error) {
	return _RollupAdminFacet.Contract.BaseStake(&_RollupAdminFacet.CallOpts)
}

// ChallengeExecutionBisectionDegree is a free data retrieval call binding the contract method 0xdc72a33b.
//
// Solidity: function challengeExecutionBisectionDegree() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) ChallengeExecutionBisectionDegree(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "challengeExecutionBisectionDegree")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChallengeExecutionBisectionDegree is a free data retrieval call binding the contract method 0xdc72a33b.
//
// Solidity: function challengeExecutionBisectionDegree() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) ChallengeExecutionBisectionDegree() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ChallengeExecutionBisectionDegree(&_RollupAdminFacet.CallOpts)
}

// ChallengeExecutionBisectionDegree is a free data retrieval call binding the contract method 0xdc72a33b.
//
// Solidity: function challengeExecutionBisectionDegree() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ChallengeExecutionBisectionDegree() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ChallengeExecutionBisectionDegree(&_RollupAdminFacet.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) ChallengeFactory() (common.Address, error) {
	return _RollupAdminFacet.Contract.ChallengeFactory(&_RollupAdminFacet.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ChallengeFactory() (common.Address, error) {
	return _RollupAdminFacet.Contract.ChallengeFactory(&_RollupAdminFacet.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) ConfirmPeriodBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "confirmPeriodBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ConfirmPeriodBlocks(&_RollupAdminFacet.CallOpts)
}

// ConfirmPeriodBlocks is a free data retrieval call binding the contract method 0x2e7acfa6.
//
// Solidity: function confirmPeriodBlocks() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ConfirmPeriodBlocks() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ConfirmPeriodBlocks(&_RollupAdminFacet.CallOpts)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) CurrentChallenge(opts *bind.CallOpts, staker common.Address) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "currentChallenge", staker)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _RollupAdminFacet.Contract.CurrentChallenge(&_RollupAdminFacet.CallOpts, staker)
}

// CurrentChallenge is a free data retrieval call binding the contract method 0x69fd251c.
//
// Solidity: function currentChallenge(address staker) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) CurrentChallenge(staker common.Address) (common.Address, error) {
	return _RollupAdminFacet.Contract.CurrentChallenge(&_RollupAdminFacet.CallOpts, staker)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) DelayedBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "delayedBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) DelayedBridge() (common.Address, error) {
	return _RollupAdminFacet.Contract.DelayedBridge(&_RollupAdminFacet.CallOpts)
}

// DelayedBridge is a free data retrieval call binding the contract method 0xf51de41b.
//
// Solidity: function delayedBridge() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) DelayedBridge() (common.Address, error) {
	return _RollupAdminFacet.Contract.DelayedBridge(&_RollupAdminFacet.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) ExtraChallengeTimeBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "extraChallengeTimeBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ExtraChallengeTimeBlocks(&_RollupAdminFacet.CallOpts)
}

// ExtraChallengeTimeBlocks is a free data retrieval call binding the contract method 0x771b2f97.
//
// Solidity: function extraChallengeTimeBlocks() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ExtraChallengeTimeBlocks() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ExtraChallengeTimeBlocks(&_RollupAdminFacet.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) FirstUnresolvedNode(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "firstUnresolvedNode")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) FirstUnresolvedNode() (*big.Int, error) {
	return _RollupAdminFacet.Contract.FirstUnresolvedNode(&_RollupAdminFacet.CallOpts)
}

// FirstUnresolvedNode is a free data retrieval call binding the contract method 0xd735e21d.
//
// Solidity: function firstUnresolvedNode() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) FirstUnresolvedNode() (*big.Int, error) {
	return _RollupAdminFacet.Contract.FirstUnresolvedNode(&_RollupAdminFacet.CallOpts)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) GetNode(opts *bind.CallOpts, nodeNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "getNode", nodeNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _RollupAdminFacet.Contract.GetNode(&_RollupAdminFacet.CallOpts, nodeNum)
}

// GetNode is a free data retrieval call binding the contract method 0x4f0f4aa9.
//
// Solidity: function getNode(uint256 nodeNum) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) GetNode(nodeNum *big.Int) (common.Address, error) {
	return _RollupAdminFacet.Contract.GetNode(&_RollupAdminFacet.CallOpts, nodeNum)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupAdminFacet *RollupAdminFacetCaller) GetNodeHash(opts *bind.CallOpts, index *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "getNodeHash", index)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupAdminFacet *RollupAdminFacetSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _RollupAdminFacet.Contract.GetNodeHash(&_RollupAdminFacet.CallOpts, index)
}

// GetNodeHash is a free data retrieval call binding the contract method 0xf8d1f194.
//
// Solidity: function getNodeHash(uint256 index) view returns(bytes32)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) GetNodeHash(index *big.Int) ([32]byte, error) {
	return _RollupAdminFacet.Contract.GetNodeHash(&_RollupAdminFacet.CallOpts, index)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) GetStakerAddress(opts *bind.CallOpts, stakerNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "getStakerAddress", stakerNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _RollupAdminFacet.Contract.GetStakerAddress(&_RollupAdminFacet.CallOpts, stakerNum)
}

// GetStakerAddress is a free data retrieval call binding the contract method 0x62a82d7d.
//
// Solidity: function getStakerAddress(uint256 stakerNum) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) GetStakerAddress(stakerNum *big.Int) (common.Address, error) {
	return _RollupAdminFacet.Contract.GetStakerAddress(&_RollupAdminFacet.CallOpts, stakerNum)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCaller) IsMaster(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "isMaster")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetSession) IsMaster() (bool, error) {
	return _RollupAdminFacet.Contract.IsMaster(&_RollupAdminFacet.CallOpts)
}

// IsMaster is a free data retrieval call binding the contract method 0x6f791d29.
//
// Solidity: function isMaster() view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) IsMaster() (bool, error) {
	return _RollupAdminFacet.Contract.IsMaster(&_RollupAdminFacet.CallOpts)
}

// IsNitroReady is a free data retrieval call binding the contract method 0xa8929e0b.
//
// Solidity: function isNitroReady() pure returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) IsNitroReady(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "isNitroReady")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsNitroReady is a free data retrieval call binding the contract method 0xa8929e0b.
//
// Solidity: function isNitroReady() pure returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) IsNitroReady() (*big.Int, error) {
	return _RollupAdminFacet.Contract.IsNitroReady(&_RollupAdminFacet.CallOpts)
}

// IsNitroReady is a free data retrieval call binding the contract method 0xa8929e0b.
//
// Solidity: function isNitroReady() pure returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) IsNitroReady() (*big.Int, error) {
	return _RollupAdminFacet.Contract.IsNitroReady(&_RollupAdminFacet.CallOpts)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCaller) IsStaked(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "isStaked", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetSession) IsStaked(staker common.Address) (bool, error) {
	return _RollupAdminFacet.Contract.IsStaked(&_RollupAdminFacet.CallOpts, staker)
}

// IsStaked is a free data retrieval call binding the contract method 0x6177fd18.
//
// Solidity: function isStaked(address staker) view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) IsStaked(staker common.Address) (bool, error) {
	return _RollupAdminFacet.Contract.IsStaked(&_RollupAdminFacet.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCaller) IsZombie(opts *bind.CallOpts, staker common.Address) (bool, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "isZombie", staker)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetSession) IsZombie(staker common.Address) (bool, error) {
	return _RollupAdminFacet.Contract.IsZombie(&_RollupAdminFacet.CallOpts, staker)
}

// IsZombie is a free data retrieval call binding the contract method 0x91c657e8.
//
// Solidity: function isZombie(address staker) view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) IsZombie(staker common.Address) (bool, error) {
	return _RollupAdminFacet.Contract.IsZombie(&_RollupAdminFacet.CallOpts, staker)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) LastStakeBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "lastStakeBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) LastStakeBlock() (*big.Int, error) {
	return _RollupAdminFacet.Contract.LastStakeBlock(&_RollupAdminFacet.CallOpts)
}

// LastStakeBlock is a free data retrieval call binding the contract method 0x8640ce5f.
//
// Solidity: function lastStakeBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) LastStakeBlock() (*big.Int, error) {
	return _RollupAdminFacet.Contract.LastStakeBlock(&_RollupAdminFacet.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) LatestConfirmed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "latestConfirmed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) LatestConfirmed() (*big.Int, error) {
	return _RollupAdminFacet.Contract.LatestConfirmed(&_RollupAdminFacet.CallOpts)
}

// LatestConfirmed is a free data retrieval call binding the contract method 0x65f7f80d.
//
// Solidity: function latestConfirmed() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) LatestConfirmed() (*big.Int, error) {
	return _RollupAdminFacet.Contract.LatestConfirmed(&_RollupAdminFacet.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) LatestNodeCreated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "latestNodeCreated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) LatestNodeCreated() (*big.Int, error) {
	return _RollupAdminFacet.Contract.LatestNodeCreated(&_RollupAdminFacet.CallOpts)
}

// LatestNodeCreated is a free data retrieval call binding the contract method 0x7ba9534a.
//
// Solidity: function latestNodeCreated() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) LatestNodeCreated() (*big.Int, error) {
	return _RollupAdminFacet.Contract.LatestNodeCreated(&_RollupAdminFacet.CallOpts)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) LatestStakedNode(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "latestStakedNode", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _RollupAdminFacet.Contract.LatestStakedNode(&_RollupAdminFacet.CallOpts, staker)
}

// LatestStakedNode is a free data retrieval call binding the contract method 0x3e96576e.
//
// Solidity: function latestStakedNode(address staker) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) LatestStakedNode(staker common.Address) (*big.Int, error) {
	return _RollupAdminFacet.Contract.LatestStakedNode(&_RollupAdminFacet.CallOpts, staker)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) MinimumAssertionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "minimumAssertionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _RollupAdminFacet.Contract.MinimumAssertionPeriod(&_RollupAdminFacet.CallOpts)
}

// MinimumAssertionPeriod is a free data retrieval call binding the contract method 0x45e38b64.
//
// Solidity: function minimumAssertionPeriod() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) MinimumAssertionPeriod() (*big.Int, error) {
	return _RollupAdminFacet.Contract.MinimumAssertionPeriod(&_RollupAdminFacet.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) NodeFactory() (common.Address, error) {
	return _RollupAdminFacet.Contract.NodeFactory(&_RollupAdminFacet.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) NodeFactory() (common.Address, error) {
	return _RollupAdminFacet.Contract.NodeFactory(&_RollupAdminFacet.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) Outbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "outbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) Outbox() (common.Address, error) {
	return _RollupAdminFacet.Contract.Outbox(&_RollupAdminFacet.CallOpts)
}

// Outbox is a free data retrieval call binding the contract method 0xce11e6ab.
//
// Solidity: function outbox() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) Outbox() (common.Address, error) {
	return _RollupAdminFacet.Contract.Outbox(&_RollupAdminFacet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) Owner() (common.Address, error) {
	return _RollupAdminFacet.Contract.Owner(&_RollupAdminFacet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) Owner() (common.Address, error) {
	return _RollupAdminFacet.Contract.Owner(&_RollupAdminFacet.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetSession) Paused() (bool, error) {
	return _RollupAdminFacet.Contract.Paused(&_RollupAdminFacet.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) Paused() (bool, error) {
	return _RollupAdminFacet.Contract.Paused(&_RollupAdminFacet.CallOpts)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) RollupEventBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "rollupEventBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) RollupEventBridge() (common.Address, error) {
	return _RollupAdminFacet.Contract.RollupEventBridge(&_RollupAdminFacet.CallOpts)
}

// RollupEventBridge is a free data retrieval call binding the contract method 0x9e8a713f.
//
// Solidity: function rollupEventBridge() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) RollupEventBridge() (common.Address, error) {
	return _RollupAdminFacet.Contract.RollupEventBridge(&_RollupAdminFacet.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) SequencerBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "sequencerBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) SequencerBridge() (common.Address, error) {
	return _RollupAdminFacet.Contract.SequencerBridge(&_RollupAdminFacet.CallOpts)
}

// SequencerBridge is a free data retrieval call binding the contract method 0x3e55c0c7.
//
// Solidity: function sequencerBridge() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) SequencerBridge() (common.Address, error) {
	return _RollupAdminFacet.Contract.SequencerBridge(&_RollupAdminFacet.CallOpts)
}

// ShutdownForNitroBlock is a free data retrieval call binding the contract method 0x7f60abbb.
//
// Solidity: function shutdownForNitroBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) ShutdownForNitroBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "shutdownForNitroBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShutdownForNitroBlock is a free data retrieval call binding the contract method 0x7f60abbb.
//
// Solidity: function shutdownForNitroBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) ShutdownForNitroBlock() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ShutdownForNitroBlock(&_RollupAdminFacet.CallOpts)
}

// ShutdownForNitroBlock is a free data retrieval call binding the contract method 0x7f60abbb.
//
// Solidity: function shutdownForNitroBlock() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ShutdownForNitroBlock() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ShutdownForNitroBlock(&_RollupAdminFacet.CallOpts)
}

// ShutdownForNitroMode is a free data retrieval call binding the contract method 0x313a04fa.
//
// Solidity: function shutdownForNitroMode() view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCaller) ShutdownForNitroMode(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "shutdownForNitroMode")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ShutdownForNitroMode is a free data retrieval call binding the contract method 0x313a04fa.
//
// Solidity: function shutdownForNitroMode() view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetSession) ShutdownForNitroMode() (bool, error) {
	return _RollupAdminFacet.Contract.ShutdownForNitroMode(&_RollupAdminFacet.CallOpts)
}

// ShutdownForNitroMode is a free data retrieval call binding the contract method 0x313a04fa.
//
// Solidity: function shutdownForNitroMode() view returns(bool)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ShutdownForNitroMode() (bool, error) {
	return _RollupAdminFacet.Contract.ShutdownForNitroMode(&_RollupAdminFacet.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) StakeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "stakeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) StakeToken() (common.Address, error) {
	return _RollupAdminFacet.Contract.StakeToken(&_RollupAdminFacet.CallOpts)
}

// StakeToken is a free data retrieval call binding the contract method 0x51ed6a30.
//
// Solidity: function stakeToken() view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) StakeToken() (common.Address, error) {
	return _RollupAdminFacet.Contract.StakeToken(&_RollupAdminFacet.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) StakerCount() (*big.Int, error) {
	return _RollupAdminFacet.Contract.StakerCount(&_RollupAdminFacet.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) StakerCount() (*big.Int, error) {
	return _RollupAdminFacet.Contract.StakerCount(&_RollupAdminFacet.CallOpts)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) WithdrawableFunds(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "withdrawableFunds", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _RollupAdminFacet.Contract.WithdrawableFunds(&_RollupAdminFacet.CallOpts, owner)
}

// WithdrawableFunds is a free data retrieval call binding the contract method 0x2f30cabd.
//
// Solidity: function withdrawableFunds(address owner) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) WithdrawableFunds(owner common.Address) (*big.Int, error) {
	return _RollupAdminFacet.Contract.WithdrawableFunds(&_RollupAdminFacet.CallOpts, owner)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCaller) ZombieAddress(opts *bind.CallOpts, zombieNum *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "zombieAddress", zombieNum)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _RollupAdminFacet.Contract.ZombieAddress(&_RollupAdminFacet.CallOpts, zombieNum)
}

// ZombieAddress is a free data retrieval call binding the contract method 0xd01e6602.
//
// Solidity: function zombieAddress(uint256 zombieNum) view returns(address)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ZombieAddress(zombieNum *big.Int) (common.Address, error) {
	return _RollupAdminFacet.Contract.ZombieAddress(&_RollupAdminFacet.CallOpts, zombieNum)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) ZombieCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "zombieCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) ZombieCount() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ZombieCount(&_RollupAdminFacet.CallOpts)
}

// ZombieCount is a free data retrieval call binding the contract method 0x63721d6b.
//
// Solidity: function zombieCount() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ZombieCount() (*big.Int, error) {
	return _RollupAdminFacet.Contract.ZombieCount(&_RollupAdminFacet.CallOpts)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) ZombieLatestStakedNode(opts *bind.CallOpts, zombieNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "zombieLatestStakedNode", zombieNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _RollupAdminFacet.Contract.ZombieLatestStakedNode(&_RollupAdminFacet.CallOpts, zombieNum)
}

// ZombieLatestStakedNode is a free data retrieval call binding the contract method 0xf33e1fac.
//
// Solidity: function zombieLatestStakedNode(uint256 zombieNum) view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) ZombieLatestStakedNode(zombieNum *big.Int) (*big.Int, error) {
	return _RollupAdminFacet.Contract.ZombieLatestStakedNode(&_RollupAdminFacet.CallOpts, zombieNum)
}

// ForceConfirmNode is a paid mutator transaction binding the contract method 0xf53f5afa.
//
// Solidity: function forceConfirmNode(uint256 nodeNum, bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) ForceConfirmNode(opts *bind.TransactOpts, nodeNum *big.Int, beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "forceConfirmNode", nodeNum, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ForceConfirmNode is a paid mutator transaction binding the contract method 0xf53f5afa.
//
// Solidity: function forceConfirmNode(uint256 nodeNum, bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) ForceConfirmNode(nodeNum *big.Int, beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ForceConfirmNode(&_RollupAdminFacet.TransactOpts, nodeNum, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ForceConfirmNode is a paid mutator transaction binding the contract method 0xf53f5afa.
//
// Solidity: function forceConfirmNode(uint256 nodeNum, bytes32 beforeSendAcc, bytes sendsData, uint256[] sendLengths, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) ForceConfirmNode(nodeNum *big.Int, beforeSendAcc [32]byte, sendsData []byte, sendLengths []*big.Int, afterSendCount *big.Int, afterLogAcc [32]byte, afterLogCount *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ForceConfirmNode(&_RollupAdminFacet.TransactOpts, nodeNum, beforeSendAcc, sendsData, sendLengths, afterSendCount, afterLogAcc, afterLogCount)
}

// ForceCreateNode is a paid mutator transaction binding the contract method 0x9ea28e65.
//
// Solidity: function forceCreateNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, bytes sequencerBatchProof, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, uint256 prevNode) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) ForceCreateNode(opts *bind.TransactOpts, expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, sequencerBatchProof []byte, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, prevNode *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "forceCreateNode", expectedNodeHash, assertionBytes32Fields, assertionIntFields, sequencerBatchProof, beforeProposedBlock, beforeInboxMaxCount, prevNode)
}

// ForceCreateNode is a paid mutator transaction binding the contract method 0x9ea28e65.
//
// Solidity: function forceCreateNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, bytes sequencerBatchProof, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, uint256 prevNode) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) ForceCreateNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, sequencerBatchProof []byte, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, prevNode *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ForceCreateNode(&_RollupAdminFacet.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, sequencerBatchProof, beforeProposedBlock, beforeInboxMaxCount, prevNode)
}

// ForceCreateNode is a paid mutator transaction binding the contract method 0x9ea28e65.
//
// Solidity: function forceCreateNode(bytes32 expectedNodeHash, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields, bytes sequencerBatchProof, uint256 beforeProposedBlock, uint256 beforeInboxMaxCount, uint256 prevNode) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) ForceCreateNode(expectedNodeHash [32]byte, assertionBytes32Fields [2][3][32]byte, assertionIntFields [2][4]*big.Int, sequencerBatchProof []byte, beforeProposedBlock *big.Int, beforeInboxMaxCount *big.Int, prevNode *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ForceCreateNode(&_RollupAdminFacet.TransactOpts, expectedNodeHash, assertionBytes32Fields, assertionIntFields, sequencerBatchProof, beforeProposedBlock, beforeInboxMaxCount, prevNode)
}

// ForceRefundStaker is a paid mutator transaction binding the contract method 0x7c75c298.
//
// Solidity: function forceRefundStaker(address[] staker) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) ForceRefundStaker(opts *bind.TransactOpts, staker []common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "forceRefundStaker", staker)
}

// ForceRefundStaker is a paid mutator transaction binding the contract method 0x7c75c298.
//
// Solidity: function forceRefundStaker(address[] staker) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) ForceRefundStaker(staker []common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ForceRefundStaker(&_RollupAdminFacet.TransactOpts, staker)
}

// ForceRefundStaker is a paid mutator transaction binding the contract method 0x7c75c298.
//
// Solidity: function forceRefundStaker(address[] staker) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) ForceRefundStaker(staker []common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ForceRefundStaker(&_RollupAdminFacet.TransactOpts, staker)
}

// ForceResolveChallenge is a paid mutator transaction binding the contract method 0xf38c9379.
//
// Solidity: function forceResolveChallenge(address[] stakerA, address[] stakerB) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) ForceResolveChallenge(opts *bind.TransactOpts, stakerA []common.Address, stakerB []common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "forceResolveChallenge", stakerA, stakerB)
}

// ForceResolveChallenge is a paid mutator transaction binding the contract method 0xf38c9379.
//
// Solidity: function forceResolveChallenge(address[] stakerA, address[] stakerB) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) ForceResolveChallenge(stakerA []common.Address, stakerB []common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ForceResolveChallenge(&_RollupAdminFacet.TransactOpts, stakerA, stakerB)
}

// ForceResolveChallenge is a paid mutator transaction binding the contract method 0xf38c9379.
//
// Solidity: function forceResolveChallenge(address[] stakerA, address[] stakerB) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) ForceResolveChallenge(stakerA []common.Address, stakerB []common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ForceResolveChallenge(&_RollupAdminFacet.TransactOpts, stakerA, stakerB)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RollupAdminFacet *RollupAdminFacetSession) Pause() (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.Pause(&_RollupAdminFacet.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) Pause() (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.Pause(&_RollupAdminFacet.TransactOpts)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) RemoveOldOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "removeOldOutbox", _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.RemoveOldOutbox(&_RollupAdminFacet.TransactOpts, _outbox)
}

// RemoveOldOutbox is a paid mutator transaction binding the contract method 0x567ca41b.
//
// Solidity: function removeOldOutbox(address _outbox) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) RemoveOldOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.RemoveOldOutbox(&_RollupAdminFacet.TransactOpts, _outbox)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_RollupAdminFacet *RollupAdminFacetSession) Resume() (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.Resume(&_RollupAdminFacet.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) Resume() (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.Resume(&_RollupAdminFacet.TransactOpts)
}

// SetAvmGasSpeedLimitPerBlock is a paid mutator transaction binding the contract method 0xa5cc82f8.
//
// Solidity: function setAvmGasSpeedLimitPerBlock(uint256 newAvmGasSpeedLimitPerBlock) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetAvmGasSpeedLimitPerBlock(opts *bind.TransactOpts, newAvmGasSpeedLimitPerBlock *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setAvmGasSpeedLimitPerBlock", newAvmGasSpeedLimitPerBlock)
}

// SetAvmGasSpeedLimitPerBlock is a paid mutator transaction binding the contract method 0xa5cc82f8.
//
// Solidity: function setAvmGasSpeedLimitPerBlock(uint256 newAvmGasSpeedLimitPerBlock) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetAvmGasSpeedLimitPerBlock(newAvmGasSpeedLimitPerBlock *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetAvmGasSpeedLimitPerBlock(&_RollupAdminFacet.TransactOpts, newAvmGasSpeedLimitPerBlock)
}

// SetAvmGasSpeedLimitPerBlock is a paid mutator transaction binding the contract method 0xa5cc82f8.
//
// Solidity: function setAvmGasSpeedLimitPerBlock(uint256 newAvmGasSpeedLimitPerBlock) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetAvmGasSpeedLimitPerBlock(newAvmGasSpeedLimitPerBlock *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetAvmGasSpeedLimitPerBlock(&_RollupAdminFacet.TransactOpts, newAvmGasSpeedLimitPerBlock)
}

// SetBaseStake is a paid mutator transaction binding the contract method 0x06ae5851.
//
// Solidity: function setBaseStake(uint256 newBaseStake) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetBaseStake(opts *bind.TransactOpts, newBaseStake *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setBaseStake", newBaseStake)
}

// SetBaseStake is a paid mutator transaction binding the contract method 0x06ae5851.
//
// Solidity: function setBaseStake(uint256 newBaseStake) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetBaseStake(newBaseStake *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetBaseStake(&_RollupAdminFacet.TransactOpts, newBaseStake)
}

// SetBaseStake is a paid mutator transaction binding the contract method 0x06ae5851.
//
// Solidity: function setBaseStake(uint256 newBaseStake) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetBaseStake(newBaseStake *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetBaseStake(&_RollupAdminFacet.TransactOpts, newBaseStake)
}

// SetChallengeExecutionBisectionDegree is a paid mutator transaction binding the contract method 0x6aef131a.
//
// Solidity: function setChallengeExecutionBisectionDegree(uint256 newChallengeExecutionBisectionDegree) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetChallengeExecutionBisectionDegree(opts *bind.TransactOpts, newChallengeExecutionBisectionDegree *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setChallengeExecutionBisectionDegree", newChallengeExecutionBisectionDegree)
}

// SetChallengeExecutionBisectionDegree is a paid mutator transaction binding the contract method 0x6aef131a.
//
// Solidity: function setChallengeExecutionBisectionDegree(uint256 newChallengeExecutionBisectionDegree) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetChallengeExecutionBisectionDegree(newChallengeExecutionBisectionDegree *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetChallengeExecutionBisectionDegree(&_RollupAdminFacet.TransactOpts, newChallengeExecutionBisectionDegree)
}

// SetChallengeExecutionBisectionDegree is a paid mutator transaction binding the contract method 0x6aef131a.
//
// Solidity: function setChallengeExecutionBisectionDegree(uint256 newChallengeExecutionBisectionDegree) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetChallengeExecutionBisectionDegree(newChallengeExecutionBisectionDegree *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetChallengeExecutionBisectionDegree(&_RollupAdminFacet.TransactOpts, newChallengeExecutionBisectionDegree)
}

// SetConfirmPeriodBlocks is a paid mutator transaction binding the contract method 0x3ea41098.
//
// Solidity: function setConfirmPeriodBlocks(uint256 newConfirmPeriod) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetConfirmPeriodBlocks(opts *bind.TransactOpts, newConfirmPeriod *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setConfirmPeriodBlocks", newConfirmPeriod)
}

// SetConfirmPeriodBlocks is a paid mutator transaction binding the contract method 0x3ea41098.
//
// Solidity: function setConfirmPeriodBlocks(uint256 newConfirmPeriod) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetConfirmPeriodBlocks(newConfirmPeriod *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetConfirmPeriodBlocks(&_RollupAdminFacet.TransactOpts, newConfirmPeriod)
}

// SetConfirmPeriodBlocks is a paid mutator transaction binding the contract method 0x3ea41098.
//
// Solidity: function setConfirmPeriodBlocks(uint256 newConfirmPeriod) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetConfirmPeriodBlocks(newConfirmPeriod *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetConfirmPeriodBlocks(&_RollupAdminFacet.TransactOpts, newConfirmPeriod)
}

// SetExtraChallengeTimeBlocks is a paid mutator transaction binding the contract method 0x9161d535.
//
// Solidity: function setExtraChallengeTimeBlocks(uint256 newExtraTimeBlocks) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetExtraChallengeTimeBlocks(opts *bind.TransactOpts, newExtraTimeBlocks *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setExtraChallengeTimeBlocks", newExtraTimeBlocks)
}

// SetExtraChallengeTimeBlocks is a paid mutator transaction binding the contract method 0x9161d535.
//
// Solidity: function setExtraChallengeTimeBlocks(uint256 newExtraTimeBlocks) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetExtraChallengeTimeBlocks(newExtraTimeBlocks *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetExtraChallengeTimeBlocks(&_RollupAdminFacet.TransactOpts, newExtraTimeBlocks)
}

// SetExtraChallengeTimeBlocks is a paid mutator transaction binding the contract method 0x9161d535.
//
// Solidity: function setExtraChallengeTimeBlocks(uint256 newExtraTimeBlocks) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetExtraChallengeTimeBlocks(newExtraTimeBlocks *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetExtraChallengeTimeBlocks(&_RollupAdminFacet.TransactOpts, newExtraTimeBlocks)
}

// SetFacets is a paid mutator transaction binding the contract method 0x27035859.
//
// Solidity: function setFacets(address newAdminFacet, address newUserFacet) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetFacets(opts *bind.TransactOpts, newAdminFacet common.Address, newUserFacet common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setFacets", newAdminFacet, newUserFacet)
}

// SetFacets is a paid mutator transaction binding the contract method 0x27035859.
//
// Solidity: function setFacets(address newAdminFacet, address newUserFacet) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetFacets(newAdminFacet common.Address, newUserFacet common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetFacets(&_RollupAdminFacet.TransactOpts, newAdminFacet, newUserFacet)
}

// SetFacets is a paid mutator transaction binding the contract method 0x27035859.
//
// Solidity: function setFacets(address newAdminFacet, address newUserFacet) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetFacets(newAdminFacet common.Address, newUserFacet common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetFacets(&_RollupAdminFacet.TransactOpts, newAdminFacet, newUserFacet)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetInbox(opts *bind.TransactOpts, _inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setInbox", _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetInbox(&_RollupAdminFacet.TransactOpts, _inbox, _enabled)
}

// SetInbox is a paid mutator transaction binding the contract method 0xe45b7ce6.
//
// Solidity: function setInbox(address _inbox, bool _enabled) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetInbox(_inbox common.Address, _enabled bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetInbox(&_RollupAdminFacet.TransactOpts, _inbox, _enabled)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address newSequencer, bool isSequencer) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetIsSequencer(opts *bind.TransactOpts, newSequencer common.Address, isSequencer bool) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setIsSequencer", newSequencer, isSequencer)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address newSequencer, bool isSequencer) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetIsSequencer(newSequencer common.Address, isSequencer bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetIsSequencer(&_RollupAdminFacet.TransactOpts, newSequencer, isSequencer)
}

// SetIsSequencer is a paid mutator transaction binding the contract method 0x1f956632.
//
// Solidity: function setIsSequencer(address newSequencer, bool isSequencer) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetIsSequencer(newSequencer common.Address, isSequencer bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetIsSequencer(&_RollupAdminFacet.TransactOpts, newSequencer, isSequencer)
}

// SetMinimumAssertionPeriod is a paid mutator transaction binding the contract method 0x948d6588.
//
// Solidity: function setMinimumAssertionPeriod(uint256 newPeriod) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetMinimumAssertionPeriod(opts *bind.TransactOpts, newPeriod *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setMinimumAssertionPeriod", newPeriod)
}

// SetMinimumAssertionPeriod is a paid mutator transaction binding the contract method 0x948d6588.
//
// Solidity: function setMinimumAssertionPeriod(uint256 newPeriod) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetMinimumAssertionPeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetMinimumAssertionPeriod(&_RollupAdminFacet.TransactOpts, newPeriod)
}

// SetMinimumAssertionPeriod is a paid mutator transaction binding the contract method 0x948d6588.
//
// Solidity: function setMinimumAssertionPeriod(uint256 newPeriod) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetMinimumAssertionPeriod(newPeriod *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetMinimumAssertionPeriod(&_RollupAdminFacet.TransactOpts, newPeriod)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetOutbox(opts *bind.TransactOpts, _outbox common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setOutbox", _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetOutbox(&_RollupAdminFacet.TransactOpts, _outbox)
}

// SetOutbox is a paid mutator transaction binding the contract method 0xff204f3b.
//
// Solidity: function setOutbox(address _outbox) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetOutbox(_outbox common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetOutbox(&_RollupAdminFacet.TransactOpts, _outbox)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setOwner", newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetOwner(&_RollupAdminFacet.TransactOpts, newOwner)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(address newOwner) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetOwner(newOwner common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetOwner(&_RollupAdminFacet.TransactOpts, newOwner)
}

// SetSequencerInboxMaxDelay is a paid mutator transaction binding the contract method 0x40b570f4.
//
// Solidity: function setSequencerInboxMaxDelay(uint256 newSequencerInboxMaxDelayBlocks, uint256 newSequencerInboxMaxDelaySeconds) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetSequencerInboxMaxDelay(opts *bind.TransactOpts, newSequencerInboxMaxDelayBlocks *big.Int, newSequencerInboxMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setSequencerInboxMaxDelay", newSequencerInboxMaxDelayBlocks, newSequencerInboxMaxDelaySeconds)
}

// SetSequencerInboxMaxDelay is a paid mutator transaction binding the contract method 0x40b570f4.
//
// Solidity: function setSequencerInboxMaxDelay(uint256 newSequencerInboxMaxDelayBlocks, uint256 newSequencerInboxMaxDelaySeconds) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetSequencerInboxMaxDelay(newSequencerInboxMaxDelayBlocks *big.Int, newSequencerInboxMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetSequencerInboxMaxDelay(&_RollupAdminFacet.TransactOpts, newSequencerInboxMaxDelayBlocks, newSequencerInboxMaxDelaySeconds)
}

// SetSequencerInboxMaxDelay is a paid mutator transaction binding the contract method 0x40b570f4.
//
// Solidity: function setSequencerInboxMaxDelay(uint256 newSequencerInboxMaxDelayBlocks, uint256 newSequencerInboxMaxDelaySeconds) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetSequencerInboxMaxDelay(newSequencerInboxMaxDelayBlocks *big.Int, newSequencerInboxMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetSequencerInboxMaxDelay(&_RollupAdminFacet.TransactOpts, newSequencerInboxMaxDelayBlocks, newSequencerInboxMaxDelaySeconds)
}

// SetStakeToken is a paid mutator transaction binding the contract method 0x0397d458.
//
// Solidity: function setStakeToken(address newStakeToken) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetStakeToken(opts *bind.TransactOpts, newStakeToken common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setStakeToken", newStakeToken)
}

// SetStakeToken is a paid mutator transaction binding the contract method 0x0397d458.
//
// Solidity: function setStakeToken(address newStakeToken) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetStakeToken(newStakeToken common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetStakeToken(&_RollupAdminFacet.TransactOpts, newStakeToken)
}

// SetStakeToken is a paid mutator transaction binding the contract method 0x0397d458.
//
// Solidity: function setStakeToken(address newStakeToken) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetStakeToken(newStakeToken common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetStakeToken(&_RollupAdminFacet.TransactOpts, newStakeToken)
}

// SetValidator is a paid mutator transaction binding the contract method 0xa3ffb772.
//
// Solidity: function setValidator(address[] _validator, bool[] _val) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetValidator(opts *bind.TransactOpts, _validator []common.Address, _val []bool) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setValidator", _validator, _val)
}

// SetValidator is a paid mutator transaction binding the contract method 0xa3ffb772.
//
// Solidity: function setValidator(address[] _validator, bool[] _val) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetValidator(_validator []common.Address, _val []bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetValidator(&_RollupAdminFacet.TransactOpts, _validator, _val)
}

// SetValidator is a paid mutator transaction binding the contract method 0xa3ffb772.
//
// Solidity: function setValidator(address[] _validator, bool[] _val) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetValidator(_validator []common.Address, _val []bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetValidator(&_RollupAdminFacet.TransactOpts, _validator, _val)
}

// SetWhitelistEntries is a paid mutator transaction binding the contract method 0xcf47bb84.
//
// Solidity: function setWhitelistEntries(address whitelist, address[] user, bool[] val) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetWhitelistEntries(opts *bind.TransactOpts, whitelist common.Address, user []common.Address, val []bool) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setWhitelistEntries", whitelist, user, val)
}

// SetWhitelistEntries is a paid mutator transaction binding the contract method 0xcf47bb84.
//
// Solidity: function setWhitelistEntries(address whitelist, address[] user, bool[] val) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetWhitelistEntries(whitelist common.Address, user []common.Address, val []bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetWhitelistEntries(&_RollupAdminFacet.TransactOpts, whitelist, user, val)
}

// SetWhitelistEntries is a paid mutator transaction binding the contract method 0xcf47bb84.
//
// Solidity: function setWhitelistEntries(address whitelist, address[] user, bool[] val) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetWhitelistEntries(whitelist common.Address, user []common.Address, val []bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetWhitelistEntries(&_RollupAdminFacet.TransactOpts, whitelist, user, val)
}

// ShutdownForNitro is a paid mutator transaction binding the contract method 0x1d0ada65.
//
// Solidity: function shutdownForNitro(uint256 finalNodeNum, bool destroyAlternatives, bool destroyChallenges) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) ShutdownForNitro(opts *bind.TransactOpts, finalNodeNum *big.Int, destroyAlternatives bool, destroyChallenges bool) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "shutdownForNitro", finalNodeNum, destroyAlternatives, destroyChallenges)
}

// ShutdownForNitro is a paid mutator transaction binding the contract method 0x1d0ada65.
//
// Solidity: function shutdownForNitro(uint256 finalNodeNum, bool destroyAlternatives, bool destroyChallenges) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) ShutdownForNitro(finalNodeNum *big.Int, destroyAlternatives bool, destroyChallenges bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ShutdownForNitro(&_RollupAdminFacet.TransactOpts, finalNodeNum, destroyAlternatives, destroyChallenges)
}

// ShutdownForNitro is a paid mutator transaction binding the contract method 0x1d0ada65.
//
// Solidity: function shutdownForNitro(uint256 finalNodeNum, bool destroyAlternatives, bool destroyChallenges) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) ShutdownForNitro(finalNodeNum *big.Int, destroyAlternatives bool, destroyChallenges bool) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.ShutdownForNitro(&_RollupAdminFacet.TransactOpts, finalNodeNum, destroyAlternatives, destroyChallenges)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) TransferOwnership(opts *bind.TransactOpts, target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "transferOwnership", target, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) TransferOwnership(target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.TransferOwnership(&_RollupAdminFacet.TransactOpts, target, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x6d435421.
//
// Solidity: function transferOwnership(address target, address newOwner) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) TransferOwnership(target common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.TransferOwnership(&_RollupAdminFacet.TransactOpts, target, newOwner)
}

// UndoShutdownForNitro is a paid mutator transaction binding the contract method 0x7e6c255f.
//
// Solidity: function undoShutdownForNitro() returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) UndoShutdownForNitro(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "undoShutdownForNitro")
}

// UndoShutdownForNitro is a paid mutator transaction binding the contract method 0x7e6c255f.
//
// Solidity: function undoShutdownForNitro() returns()
func (_RollupAdminFacet *RollupAdminFacetSession) UndoShutdownForNitro() (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.UndoShutdownForNitro(&_RollupAdminFacet.TransactOpts)
}

// UndoShutdownForNitro is a paid mutator transaction binding the contract method 0x7e6c255f.
//
// Solidity: function undoShutdownForNitro() returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) UndoShutdownForNitro() (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.UndoShutdownForNitro(&_RollupAdminFacet.TransactOpts)
}

// UpdateWhitelistConsumers is a paid mutator transaction binding the contract method 0x661d2722.
//
// Solidity: function updateWhitelistConsumers(address whitelist, address newWhitelist, address[] targets) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) UpdateWhitelistConsumers(opts *bind.TransactOpts, whitelist common.Address, newWhitelist common.Address, targets []common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "updateWhitelistConsumers", whitelist, newWhitelist, targets)
}

// UpdateWhitelistConsumers is a paid mutator transaction binding the contract method 0x661d2722.
//
// Solidity: function updateWhitelistConsumers(address whitelist, address newWhitelist, address[] targets) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) UpdateWhitelistConsumers(whitelist common.Address, newWhitelist common.Address, targets []common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.UpdateWhitelistConsumers(&_RollupAdminFacet.TransactOpts, whitelist, newWhitelist, targets)
}

// UpdateWhitelistConsumers is a paid mutator transaction binding the contract method 0x661d2722.
//
// Solidity: function updateWhitelistConsumers(address whitelist, address newWhitelist, address[] targets) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) UpdateWhitelistConsumers(whitelist common.Address, newWhitelist common.Address, targets []common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.UpdateWhitelistConsumers(&_RollupAdminFacet.TransactOpts, whitelist, newWhitelist, targets)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x848bf918.
//
// Solidity: function upgradeBeacon(address beacon, address newImplementation) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) UpgradeBeacon(opts *bind.TransactOpts, beacon common.Address, newImplementation common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "upgradeBeacon", beacon, newImplementation)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x848bf918.
//
// Solidity: function upgradeBeacon(address beacon, address newImplementation) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) UpgradeBeacon(beacon common.Address, newImplementation common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.UpgradeBeacon(&_RollupAdminFacet.TransactOpts, beacon, newImplementation)
}

// UpgradeBeacon is a paid mutator transaction binding the contract method 0x848bf918.
//
// Solidity: function upgradeBeacon(address beacon, address newImplementation) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) UpgradeBeacon(beacon common.Address, newImplementation common.Address) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.UpgradeBeacon(&_RollupAdminFacet.TransactOpts, beacon, newImplementation)
}

// RollupAdminFacetChallengeDestroyedInMigrationIterator is returned from FilterChallengeDestroyedInMigration and is used to iterate over the raw logs and unpacked data for ChallengeDestroyedInMigration events raised by the RollupAdminFacet contract.
type RollupAdminFacetChallengeDestroyedInMigrationIterator struct {
	Event *RollupAdminFacetChallengeDestroyedInMigration // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetChallengeDestroyedInMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetChallengeDestroyedInMigration)
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
		it.Event = new(RollupAdminFacetChallengeDestroyedInMigration)
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
func (it *RollupAdminFacetChallengeDestroyedInMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetChallengeDestroyedInMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetChallengeDestroyedInMigration represents a ChallengeDestroyedInMigration event raised by the RollupAdminFacet contract.
type RollupAdminFacetChallengeDestroyedInMigration struct {
	Challenge common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChallengeDestroyedInMigration is a free log retrieval operation binding the contract event 0xec0b2d27905d678228bae2ed41ea32ea8bdbdd62e81edec23035cbde03f122a2.
//
// Solidity: event ChallengeDestroyedInMigration(address challenge)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterChallengeDestroyedInMigration(opts *bind.FilterOpts) (*RollupAdminFacetChallengeDestroyedInMigrationIterator, error) {

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "ChallengeDestroyedInMigration")
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetChallengeDestroyedInMigrationIterator{contract: _RollupAdminFacet.contract, event: "ChallengeDestroyedInMigration", logs: logs, sub: sub}, nil
}

// WatchChallengeDestroyedInMigration is a free log subscription operation binding the contract event 0xec0b2d27905d678228bae2ed41ea32ea8bdbdd62e81edec23035cbde03f122a2.
//
// Solidity: event ChallengeDestroyedInMigration(address challenge)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchChallengeDestroyedInMigration(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetChallengeDestroyedInMigration) (event.Subscription, error) {

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "ChallengeDestroyedInMigration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetChallengeDestroyedInMigration)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "ChallengeDestroyedInMigration", log); err != nil {
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

// ParseChallengeDestroyedInMigration is a log parse operation binding the contract event 0xec0b2d27905d678228bae2ed41ea32ea8bdbdd62e81edec23035cbde03f122a2.
//
// Solidity: event ChallengeDestroyedInMigration(address challenge)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseChallengeDestroyedInMigration(log types.Log) (*RollupAdminFacetChallengeDestroyedInMigration, error) {
	event := new(RollupAdminFacetChallengeDestroyedInMigration)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "ChallengeDestroyedInMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetNodeConfirmedIterator is returned from FilterNodeConfirmed and is used to iterate over the raw logs and unpacked data for NodeConfirmed events raised by the RollupAdminFacet contract.
type RollupAdminFacetNodeConfirmedIterator struct {
	Event *RollupAdminFacetNodeConfirmed // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetNodeConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetNodeConfirmed)
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
		it.Event = new(RollupAdminFacetNodeConfirmed)
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
func (it *RollupAdminFacetNodeConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetNodeConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetNodeConfirmed represents a NodeConfirmed event raised by the RollupAdminFacet contract.
type RollupAdminFacetNodeConfirmed struct {
	NodeNum        *big.Int
	AfterSendAcc   [32]byte
	AfterSendCount *big.Int
	AfterLogAcc    [32]byte
	AfterLogCount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeConfirmed is a free log retrieval operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterNodeConfirmed(opts *bind.FilterOpts, nodeNum []*big.Int) (*RollupAdminFacetNodeConfirmedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetNodeConfirmedIterator{contract: _RollupAdminFacet.contract, event: "NodeConfirmed", logs: logs, sub: sub}, nil
}

// WatchNodeConfirmed is a free log subscription operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchNodeConfirmed(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetNodeConfirmed, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "NodeConfirmed", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetNodeConfirmed)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
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

// ParseNodeConfirmed is a log parse operation binding the contract event 0x2400bd6e429cfcd98fe43a75bbbe4702c59c99d636100690130cc1ebb611c5a2.
//
// Solidity: event NodeConfirmed(uint256 indexed nodeNum, bytes32 afterSendAcc, uint256 afterSendCount, bytes32 afterLogAcc, uint256 afterLogCount)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseNodeConfirmed(log types.Log) (*RollupAdminFacetNodeConfirmed, error) {
	event := new(RollupAdminFacetNodeConfirmed)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "NodeConfirmed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetNodeCreatedIterator is returned from FilterNodeCreated and is used to iterate over the raw logs and unpacked data for NodeCreated events raised by the RollupAdminFacet contract.
type RollupAdminFacetNodeCreatedIterator struct {
	Event *RollupAdminFacetNodeCreated // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetNodeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetNodeCreated)
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
		it.Event = new(RollupAdminFacetNodeCreated)
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
func (it *RollupAdminFacetNodeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetNodeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetNodeCreated represents a NodeCreated event raised by the RollupAdminFacet contract.
type RollupAdminFacetNodeCreated struct {
	NodeNum                 *big.Int
	ParentNodeHash          [32]byte
	NodeHash                [32]byte
	ExecutionHash           [32]byte
	InboxMaxCount           *big.Int
	AfterInboxBatchEndCount *big.Int
	AfterInboxBatchAcc      [32]byte
	AssertionBytes32Fields  [2][3][32]byte
	AssertionIntFields      [2][4]*big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterNodeCreated is a free log retrieval operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterNodeCreated(opts *bind.FilterOpts, nodeNum []*big.Int, parentNodeHash [][32]byte) (*RollupAdminFacetNodeCreatedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetNodeCreatedIterator{contract: _RollupAdminFacet.contract, event: "NodeCreated", logs: logs, sub: sub}, nil
}

// WatchNodeCreated is a free log subscription operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchNodeCreated(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetNodeCreated, nodeNum []*big.Int, parentNodeHash [][32]byte) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}
	var parentNodeHashRule []interface{}
	for _, parentNodeHashItem := range parentNodeHash {
		parentNodeHashRule = append(parentNodeHashRule, parentNodeHashItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "NodeCreated", nodeNumRule, parentNodeHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetNodeCreated)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "NodeCreated", log); err != nil {
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

// ParseNodeCreated is a log parse operation binding the contract event 0x8016306209aff73e79f274cf38a41928996f746e2953111902e1f55be1713a54.
//
// Solidity: event NodeCreated(uint256 indexed nodeNum, bytes32 indexed parentNodeHash, bytes32 nodeHash, bytes32 executionHash, uint256 inboxMaxCount, uint256 afterInboxBatchEndCount, bytes32 afterInboxBatchAcc, bytes32[3][2] assertionBytes32Fields, uint256[4][2] assertionIntFields)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseNodeCreated(log types.Log) (*RollupAdminFacetNodeCreated, error) {
	event := new(RollupAdminFacetNodeCreated)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "NodeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetNodeDestroyedInMigrationIterator is returned from FilterNodeDestroyedInMigration and is used to iterate over the raw logs and unpacked data for NodeDestroyedInMigration events raised by the RollupAdminFacet contract.
type RollupAdminFacetNodeDestroyedInMigrationIterator struct {
	Event *RollupAdminFacetNodeDestroyedInMigration // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetNodeDestroyedInMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetNodeDestroyedInMigration)
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
		it.Event = new(RollupAdminFacetNodeDestroyedInMigration)
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
func (it *RollupAdminFacetNodeDestroyedInMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetNodeDestroyedInMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetNodeDestroyedInMigration represents a NodeDestroyedInMigration event raised by the RollupAdminFacet contract.
type RollupAdminFacetNodeDestroyedInMigration struct {
	NodeNum *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeDestroyedInMigration is a free log retrieval operation binding the contract event 0xc48f1661fe65917dbe9d175ac4cb62063ef44afe989dcd3dbf470ac5a1c77bcb.
//
// Solidity: event NodeDestroyedInMigration(uint256 nodeNum)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterNodeDestroyedInMigration(opts *bind.FilterOpts) (*RollupAdminFacetNodeDestroyedInMigrationIterator, error) {

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "NodeDestroyedInMigration")
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetNodeDestroyedInMigrationIterator{contract: _RollupAdminFacet.contract, event: "NodeDestroyedInMigration", logs: logs, sub: sub}, nil
}

// WatchNodeDestroyedInMigration is a free log subscription operation binding the contract event 0xc48f1661fe65917dbe9d175ac4cb62063ef44afe989dcd3dbf470ac5a1c77bcb.
//
// Solidity: event NodeDestroyedInMigration(uint256 nodeNum)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchNodeDestroyedInMigration(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetNodeDestroyedInMigration) (event.Subscription, error) {

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "NodeDestroyedInMigration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetNodeDestroyedInMigration)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "NodeDestroyedInMigration", log); err != nil {
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

// ParseNodeDestroyedInMigration is a log parse operation binding the contract event 0xc48f1661fe65917dbe9d175ac4cb62063ef44afe989dcd3dbf470ac5a1c77bcb.
//
// Solidity: event NodeDestroyedInMigration(uint256 nodeNum)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseNodeDestroyedInMigration(log types.Log) (*RollupAdminFacetNodeDestroyedInMigration, error) {
	event := new(RollupAdminFacetNodeDestroyedInMigration)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "NodeDestroyedInMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetNodeRejectedIterator is returned from FilterNodeRejected and is used to iterate over the raw logs and unpacked data for NodeRejected events raised by the RollupAdminFacet contract.
type RollupAdminFacetNodeRejectedIterator struct {
	Event *RollupAdminFacetNodeRejected // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetNodeRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetNodeRejected)
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
		it.Event = new(RollupAdminFacetNodeRejected)
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
func (it *RollupAdminFacetNodeRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetNodeRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetNodeRejected represents a NodeRejected event raised by the RollupAdminFacet contract.
type RollupAdminFacetNodeRejected struct {
	NodeNum *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNodeRejected is a free log retrieval operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterNodeRejected(opts *bind.FilterOpts, nodeNum []*big.Int) (*RollupAdminFacetNodeRejectedIterator, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetNodeRejectedIterator{contract: _RollupAdminFacet.contract, event: "NodeRejected", logs: logs, sub: sub}, nil
}

// WatchNodeRejected is a free log subscription operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchNodeRejected(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetNodeRejected, nodeNum []*big.Int) (event.Subscription, error) {

	var nodeNumRule []interface{}
	for _, nodeNumItem := range nodeNum {
		nodeNumRule = append(nodeNumRule, nodeNumItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "NodeRejected", nodeNumRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetNodeRejected)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "NodeRejected", log); err != nil {
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

// ParseNodeRejected is a log parse operation binding the contract event 0x9f7eee12f08e41a1d1a617e76576aa2d6a1e06dbdd72d817e62b6e8dfdebe2a3.
//
// Solidity: event NodeRejected(uint256 indexed nodeNum)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseNodeRejected(log types.Log) (*RollupAdminFacetNodeRejected, error) {
	event := new(RollupAdminFacetNodeRejected)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "NodeRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetOwnerFunctionCalledIterator is returned from FilterOwnerFunctionCalled and is used to iterate over the raw logs and unpacked data for OwnerFunctionCalled events raised by the RollupAdminFacet contract.
type RollupAdminFacetOwnerFunctionCalledIterator struct {
	Event *RollupAdminFacetOwnerFunctionCalled // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetOwnerFunctionCalledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetOwnerFunctionCalled)
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
		it.Event = new(RollupAdminFacetOwnerFunctionCalled)
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
func (it *RollupAdminFacetOwnerFunctionCalledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetOwnerFunctionCalledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetOwnerFunctionCalled represents a OwnerFunctionCalled event raised by the RollupAdminFacet contract.
type RollupAdminFacetOwnerFunctionCalled struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOwnerFunctionCalled is a free log retrieval operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterOwnerFunctionCalled(opts *bind.FilterOpts, id []*big.Int) (*RollupAdminFacetOwnerFunctionCalledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetOwnerFunctionCalledIterator{contract: _RollupAdminFacet.contract, event: "OwnerFunctionCalled", logs: logs, sub: sub}, nil
}

// WatchOwnerFunctionCalled is a free log subscription operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchOwnerFunctionCalled(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetOwnerFunctionCalled, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "OwnerFunctionCalled", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetOwnerFunctionCalled)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
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

// ParseOwnerFunctionCalled is a log parse operation binding the contract event 0xea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456e.
//
// Solidity: event OwnerFunctionCalled(uint256 indexed id)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseOwnerFunctionCalled(log types.Log) (*RollupAdminFacetOwnerFunctionCalled, error) {
	event := new(RollupAdminFacetOwnerFunctionCalled)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "OwnerFunctionCalled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the RollupAdminFacet contract.
type RollupAdminFacetPausedIterator struct {
	Event *RollupAdminFacetPaused // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetPaused)
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
		it.Event = new(RollupAdminFacetPaused)
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
func (it *RollupAdminFacetPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetPaused represents a Paused event raised by the RollupAdminFacet contract.
type RollupAdminFacetPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterPaused(opts *bind.FilterOpts) (*RollupAdminFacetPausedIterator, error) {

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetPausedIterator{contract: _RollupAdminFacet.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetPaused) (event.Subscription, error) {

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetPaused)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParsePaused(log types.Log) (*RollupAdminFacetPaused, error) {
	event := new(RollupAdminFacetPaused)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetRollupChallengeStartedIterator is returned from FilterRollupChallengeStarted and is used to iterate over the raw logs and unpacked data for RollupChallengeStarted events raised by the RollupAdminFacet contract.
type RollupAdminFacetRollupChallengeStartedIterator struct {
	Event *RollupAdminFacetRollupChallengeStarted // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetRollupChallengeStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetRollupChallengeStarted)
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
		it.Event = new(RollupAdminFacetRollupChallengeStarted)
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
func (it *RollupAdminFacetRollupChallengeStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetRollupChallengeStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetRollupChallengeStarted represents a RollupChallengeStarted event raised by the RollupAdminFacet contract.
type RollupAdminFacetRollupChallengeStarted struct {
	ChallengeContract common.Address
	Asserter          common.Address
	Challenger        common.Address
	ChallengedNode    *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRollupChallengeStarted is a free log retrieval operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterRollupChallengeStarted(opts *bind.FilterOpts, challengeContract []common.Address) (*RollupAdminFacetRollupChallengeStartedIterator, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetRollupChallengeStartedIterator{contract: _RollupAdminFacet.contract, event: "RollupChallengeStarted", logs: logs, sub: sub}, nil
}

// WatchRollupChallengeStarted is a free log subscription operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchRollupChallengeStarted(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetRollupChallengeStarted, challengeContract []common.Address) (event.Subscription, error) {

	var challengeContractRule []interface{}
	for _, challengeContractItem := range challengeContract {
		challengeContractRule = append(challengeContractRule, challengeContractItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "RollupChallengeStarted", challengeContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetRollupChallengeStarted)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
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

// ParseRollupChallengeStarted is a log parse operation binding the contract event 0xa5256d19d4ddaf646f4b5c1861b8d4c08238e6356b8ae36dcc49ac67fda75879.
//
// Solidity: event RollupChallengeStarted(address indexed challengeContract, address asserter, address challenger, uint256 challengedNode)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseRollupChallengeStarted(log types.Log) (*RollupAdminFacetRollupChallengeStarted, error) {
	event := new(RollupAdminFacetRollupChallengeStarted)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "RollupChallengeStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the RollupAdminFacet contract.
type RollupAdminFacetRollupCreatedIterator struct {
	Event *RollupAdminFacetRollupCreated // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetRollupCreated)
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
		it.Event = new(RollupAdminFacetRollupCreated)
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
func (it *RollupAdminFacetRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetRollupCreated represents a RollupCreated event raised by the RollupAdminFacet contract.
type RollupAdminFacetRollupCreated struct {
	MachineHash [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterRollupCreated(opts *bind.FilterOpts) (*RollupAdminFacetRollupCreatedIterator, error) {

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetRollupCreatedIterator{contract: _RollupAdminFacet.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetRollupCreated) (event.Subscription, error) {

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "RollupCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetRollupCreated)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0x4ac0014773275a3dfb58c58539631006301de41998cce7c4f8698d297c88bb2d.
//
// Solidity: event RollupCreated(bytes32 machineHash)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseRollupCreated(log types.Log) (*RollupAdminFacetRollupCreated, error) {
	event := new(RollupAdminFacetRollupCreated)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetStakerWithdrawnInMigrationIterator is returned from FilterStakerWithdrawnInMigration and is used to iterate over the raw logs and unpacked data for StakerWithdrawnInMigration events raised by the RollupAdminFacet contract.
type RollupAdminFacetStakerWithdrawnInMigrationIterator struct {
	Event *RollupAdminFacetStakerWithdrawnInMigration // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetStakerWithdrawnInMigrationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetStakerWithdrawnInMigration)
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
		it.Event = new(RollupAdminFacetStakerWithdrawnInMigration)
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
func (it *RollupAdminFacetStakerWithdrawnInMigrationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetStakerWithdrawnInMigrationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetStakerWithdrawnInMigration represents a StakerWithdrawnInMigration event raised by the RollupAdminFacet contract.
type RollupAdminFacetStakerWithdrawnInMigration struct {
	Staker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakerWithdrawnInMigration is a free log retrieval operation binding the contract event 0xe695a52cb984a997f48f43d18e30b3ea892d024ca5410f74e5fded60b4a033ef.
//
// Solidity: event StakerWithdrawnInMigration(address staker)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterStakerWithdrawnInMigration(opts *bind.FilterOpts) (*RollupAdminFacetStakerWithdrawnInMigrationIterator, error) {

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "StakerWithdrawnInMigration")
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetStakerWithdrawnInMigrationIterator{contract: _RollupAdminFacet.contract, event: "StakerWithdrawnInMigration", logs: logs, sub: sub}, nil
}

// WatchStakerWithdrawnInMigration is a free log subscription operation binding the contract event 0xe695a52cb984a997f48f43d18e30b3ea892d024ca5410f74e5fded60b4a033ef.
//
// Solidity: event StakerWithdrawnInMigration(address staker)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchStakerWithdrawnInMigration(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetStakerWithdrawnInMigration) (event.Subscription, error) {

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "StakerWithdrawnInMigration")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetStakerWithdrawnInMigration)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "StakerWithdrawnInMigration", log); err != nil {
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

// ParseStakerWithdrawnInMigration is a log parse operation binding the contract event 0xe695a52cb984a997f48f43d18e30b3ea892d024ca5410f74e5fded60b4a033ef.
//
// Solidity: event StakerWithdrawnInMigration(address staker)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseStakerWithdrawnInMigration(log types.Log) (*RollupAdminFacetStakerWithdrawnInMigration, error) {
	event := new(RollupAdminFacetStakerWithdrawnInMigration)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "StakerWithdrawnInMigration", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the RollupAdminFacet contract.
type RollupAdminFacetUnpausedIterator struct {
	Event *RollupAdminFacetUnpaused // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetUnpaused)
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
		it.Event = new(RollupAdminFacetUnpaused)
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
func (it *RollupAdminFacetUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetUnpaused represents a Unpaused event raised by the RollupAdminFacet contract.
type RollupAdminFacetUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterUnpaused(opts *bind.FilterOpts) (*RollupAdminFacetUnpausedIterator, error) {

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetUnpausedIterator{contract: _RollupAdminFacet.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetUnpaused) (event.Subscription, error) {

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetUnpaused)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseUnpaused(log types.Log) (*RollupAdminFacetUnpaused, error) {
	event := new(RollupAdminFacetUnpaused)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetUserStakeUpdatedIterator is returned from FilterUserStakeUpdated and is used to iterate over the raw logs and unpacked data for UserStakeUpdated events raised by the RollupAdminFacet contract.
type RollupAdminFacetUserStakeUpdatedIterator struct {
	Event *RollupAdminFacetUserStakeUpdated // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetUserStakeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetUserStakeUpdated)
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
		it.Event = new(RollupAdminFacetUserStakeUpdated)
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
func (it *RollupAdminFacetUserStakeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetUserStakeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetUserStakeUpdated represents a UserStakeUpdated event raised by the RollupAdminFacet contract.
type RollupAdminFacetUserStakeUpdated struct {
	User           common.Address
	InitialBalance *big.Int
	FinalBalance   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUserStakeUpdated is a free log retrieval operation binding the contract event 0xebd093d389ab57f3566918d2c379a2b4d9539e8eb95efad9d5e465457833fde6.
//
// Solidity: event UserStakeUpdated(address indexed user, uint256 initialBalance, uint256 finalBalance)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterUserStakeUpdated(opts *bind.FilterOpts, user []common.Address) (*RollupAdminFacetUserStakeUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "UserStakeUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetUserStakeUpdatedIterator{contract: _RollupAdminFacet.contract, event: "UserStakeUpdated", logs: logs, sub: sub}, nil
}

// WatchUserStakeUpdated is a free log subscription operation binding the contract event 0xebd093d389ab57f3566918d2c379a2b4d9539e8eb95efad9d5e465457833fde6.
//
// Solidity: event UserStakeUpdated(address indexed user, uint256 initialBalance, uint256 finalBalance)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchUserStakeUpdated(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetUserStakeUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "UserStakeUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetUserStakeUpdated)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "UserStakeUpdated", log); err != nil {
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

// ParseUserStakeUpdated is a log parse operation binding the contract event 0xebd093d389ab57f3566918d2c379a2b4d9539e8eb95efad9d5e465457833fde6.
//
// Solidity: event UserStakeUpdated(address indexed user, uint256 initialBalance, uint256 finalBalance)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseUserStakeUpdated(log types.Log) (*RollupAdminFacetUserStakeUpdated, error) {
	event := new(RollupAdminFacetUserStakeUpdated)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "UserStakeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupAdminFacetUserWithdrawableFundsUpdatedIterator is returned from FilterUserWithdrawableFundsUpdated and is used to iterate over the raw logs and unpacked data for UserWithdrawableFundsUpdated events raised by the RollupAdminFacet contract.
type RollupAdminFacetUserWithdrawableFundsUpdatedIterator struct {
	Event *RollupAdminFacetUserWithdrawableFundsUpdated // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetUserWithdrawableFundsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetUserWithdrawableFundsUpdated)
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
		it.Event = new(RollupAdminFacetUserWithdrawableFundsUpdated)
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
func (it *RollupAdminFacetUserWithdrawableFundsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetUserWithdrawableFundsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetUserWithdrawableFundsUpdated represents a UserWithdrawableFundsUpdated event raised by the RollupAdminFacet contract.
type RollupAdminFacetUserWithdrawableFundsUpdated struct {
	User           common.Address
	InitialBalance *big.Int
	FinalBalance   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUserWithdrawableFundsUpdated is a free log retrieval operation binding the contract event 0xa740af14c56e4e04a617b1de1eb20de73270decbaaead14f142aabf3038e5ae2.
//
// Solidity: event UserWithdrawableFundsUpdated(address indexed user, uint256 initialBalance, uint256 finalBalance)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterUserWithdrawableFundsUpdated(opts *bind.FilterOpts, user []common.Address) (*RollupAdminFacetUserWithdrawableFundsUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "UserWithdrawableFundsUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetUserWithdrawableFundsUpdatedIterator{contract: _RollupAdminFacet.contract, event: "UserWithdrawableFundsUpdated", logs: logs, sub: sub}, nil
}

// WatchUserWithdrawableFundsUpdated is a free log subscription operation binding the contract event 0xa740af14c56e4e04a617b1de1eb20de73270decbaaead14f142aabf3038e5ae2.
//
// Solidity: event UserWithdrawableFundsUpdated(address indexed user, uint256 initialBalance, uint256 finalBalance)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchUserWithdrawableFundsUpdated(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetUserWithdrawableFundsUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "UserWithdrawableFundsUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetUserWithdrawableFundsUpdated)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "UserWithdrawableFundsUpdated", log); err != nil {
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

// ParseUserWithdrawableFundsUpdated is a log parse operation binding the contract event 0xa740af14c56e4e04a617b1de1eb20de73270decbaaead14f142aabf3038e5ae2.
//
// Solidity: event UserWithdrawableFundsUpdated(address indexed user, uint256 initialBalance, uint256 finalBalance)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseUserWithdrawableFundsUpdated(log types.Log) (*RollupAdminFacetUserWithdrawableFundsUpdated, error) {
	event := new(RollupAdminFacetUserWithdrawableFundsUpdated)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "UserWithdrawableFundsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
