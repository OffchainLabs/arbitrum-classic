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

// RollupAdminFacetABI is the input ABI used to generate the binding from.
const RollupAdminFacetABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterSendAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterSendCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterLogAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterLogCount\",\"type\":\"uint256\"}],\"name\":\"NodeConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"parentNodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"executionHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"inboxMaxCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"afterInboxBatchEndCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"afterInboxBatchAcc\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32[3][2]\",\"name\":\"assertionBytes32Fields\",\"type\":\"bytes32[3][2]\"},{\"indexed\":false,\"internalType\":\"uint256[4][2]\",\"name\":\"assertionIntFields\",\"type\":\"uint256[4][2]\"}],\"name\":\"NodeCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"NodeRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startNode\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endNode\",\"type\":\"uint256\"}],\"name\":\"NodesDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"OwnerFunctionCalled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"challengeContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asserter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"challenger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"challengedNode\",\"type\":\"uint256\"}],\"name\":\"RollupChallengeStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"machineHash\",\"type\":\"bytes32\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNode\",\"type\":\"uint256\"}],\"name\":\"StakerReassigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_stakerMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestStakedNode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountStaked\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"currentChallenge\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isStaked\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"amountStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"arbGasSpeedLimitPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseStake\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeExecutionBisectionDegree\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"contractIChallengeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"confirmPeriodBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"currentChallenge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayedBridge\",\"outputs\":[{\"internalType\":\"contractIBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"extraChallengeTimeBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"firstUnresolvedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"getNode\",\"outputs\":[{\"internalType\":\"contractINode\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getNodeHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakerNum\",\"type\":\"uint256\"}],\"name\":\"getStakerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMaster\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"isZombie\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastStakeBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestConfirmed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestNodeCreated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"latestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumAssertionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"contractINodeFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"outbox\",\"outputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"removeOldOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupEventBridge\",\"outputs\":[{\"internalType\":\"contractRollupEventBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerBridge\",\"outputs\":[{\"internalType\":\"contractISequencerInbox\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInboxMaxDelayBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sequencerInboxMaxDelaySeconds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newArbGasSpeedLimitPerBlock\",\"type\":\"uint256\"}],\"name\":\"setArbGasSpeedLimitPerBlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newBaseStake\",\"type\":\"uint256\"}],\"name\":\"setBaseStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newChallengeExecutionBisectionDegree\",\"type\":\"uint256\"}],\"name\":\"setChallengeExecutionBisectionDegree\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newConfirmPeriod\",\"type\":\"uint256\"}],\"name\":\"setConfirmPeriodBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newExtraTimeBlocks\",\"type\":\"uint256\"}],\"name\":\"setExtraChallengeTimeBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdminFacet\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newUserFacet\",\"type\":\"address\"}],\"name\":\"setFacets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_inbox\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setInbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"setMinimumAssertionPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIOutbox\",\"name\":\"_outbox\",\"type\":\"address\"}],\"name\":\"setOutbox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSequencerInboxMaxDelayBlocks\",\"type\":\"uint256\"}],\"name\":\"setSequencerInboxMaxDelayBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSequencerInboxMaxDelaySeconds\",\"type\":\"uint256\"}],\"name\":\"setSequencerInboxMaxDelaySeconds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newStakeToken\",\"type\":\"address\"}],\"name\":\"setStakeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_validator\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"_val\",\"type\":\"bool[]\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"withdrawableFunds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"zombieCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"zombieNum\",\"type\":\"uint256\"}],\"name\":\"zombieLatestStakedNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// RollupAdminFacetBin is the compiled bytecode used for deploying new contracts.
var RollupAdminFacetBin = "0x608060405234801561001057600080fd5b506000805460ff19908116600117909155600b80549091169055611227806100396000396000f3fe608060405234801561001057600080fd5b50600436106102d75760003560e01c806376e7e23b11610188578063ce11e6ab116100e4578063e8bd492211610092578063e8bd4922146107b9578063ef40a67014610815578063f322c0bb1461083b578063f33e1fac14610858578063f51de41b14610875578063f8d1f1941461087d578063ff204f3b1461089a576102d7565b8063ce11e6ab14610746578063d01e66021461074e578063d735e21d1461076b578063d93fe9c414610773578063dc72a33b1461077b578063dff6978714610783578063e45b7ce61461078b576102d7565b80639161d535116101415780639161d5351461059657806391c657e8146105b3578063948d6588146105d95780639e8a713f146105f6578063a3ffb772146105fe578063addd678414610721578063cd6bf14d14610729576102d7565b806376e7e23b14610566578063771b2f971461056e5780637ba9534a146105765780638456cb591461057e5780638640ce5f146105865780638da5cb5b1461058e576102d7565b80634f0f4aa9116102375780636177fd18116101f05780636177fd18146104c857806362a82d7d146104ee57806363721d6b1461050b57806365f7f80d1461051357806369fd251c1461051b5780636aef131a146105415780636f791d291461055e576102d7565b80634f0f4aa91461045157806351ed6a301461046e578063567ca41b146104765780635c975abb1461049c5780635dbaf68b146104b85780635e8ef106146104c0576102d7565b80632e7acfa6116102945780632e7acfa6146103975780632f30cabd1461039f5780633e55c0c7146103c55780633e96576e146103e95780633ea410981461040f57806345e38b641461042c5780634a56bab614610434576102d7565b80630397d458146102dc578063046f7da21461030457806306ae58511461030c57806313af40351461032957806314828f921461034f5780632703585914610369575b600080fd5b610302600480360360208110156102f257600080fd5b50356001600160a01b03166108c0565b005b6103026108f8565b6103026004803603602081101561032257600080fd5b503561091c565b6103026004803603602081101561033f57600080fd5b50356001600160a01b031661093e565b610357610976565b60408051918252519081900360200190f35b6103026004803603604081101561037f57600080fd5b506001600160a01b038135811691602001351661097c565b610357610a0c565b610357600480360360208110156103b557600080fd5b50356001600160a01b0316610a12565b6103cd610a31565b604080516001600160a01b039092168252519081900360200190f35b610357600480360360208110156103ff57600080fd5b50356001600160a01b0316610a40565b6103026004803603602081101561042557600080fd5b5035610a5e565b610357610a80565b6103026004803603602081101561044a57600080fd5b5035610a86565b6103cd6004803603602081101561046757600080fd5b5035610aa8565b6103cd610ac3565b6103026004803603602081101561048c57600080fd5b50356001600160a01b0316610ad2565b6104a4610baa565b604080519115158252519081900360200190f35b6103cd610bb3565b610357610bc2565b6104a4600480360360208110156104de57600080fd5b50356001600160a01b0316610bc8565b6103cd6004803603602081101561050457600080fd5b5035610bf0565b610357610c1a565b610357610c20565b6103cd6004803603602081101561053157600080fd5b50356001600160a01b0316610c26565b6103026004803603602081101561055757600080fd5b5035610c47565b6104a4610c69565b610357610c72565b610357610c78565b610357610c7e565b610302610c84565b610357610ca8565b6103cd610cae565b610302600480360360208110156105ac57600080fd5b5035610cbd565b6104a4600480360360208110156105c957600080fd5b50356001600160a01b0316610cdf565b610302600480360360208110156105ef57600080fd5b5035610d39565b6103cd610d5b565b6103026004803603604081101561061457600080fd5b810190602081018135600160201b81111561062e57600080fd5b82018360208201111561064057600080fd5b803590602001918460208302840111600160201b8311171561066157600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295949360208101935035915050600160201b8111156106b057600080fd5b8201836020820111156106c257600080fd5b803590602001918460208302840111600160201b831117156106e357600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929550610d6a945050505050565b610357610e38565b6103026004803603602081101561073f57600080fd5b5035610e3e565b6103cd610e60565b6103cd6004803603602081101561076457600080fd5b5035610e6f565b610357610e9e565b6103cd610ea4565b610357610eb3565b610357610eb9565b610302600480360360408110156107a157600080fd5b506001600160a01b0381351690602001351515610ebf565b6107df600480360360208110156107cf57600080fd5b50356001600160a01b0316610f4b565b604080519586526020860194909452848401929092526001600160a01b0316606084015215156080830152519081900360a00190f35b6103576004803603602081101561082b57600080fd5b50356001600160a01b0316610f87565b6103026004803603602081101561085157600080fd5b5035610fa5565b6103576004803603602081101561086e57600080fd5b5035610fc7565b6103cd610fef565b6103576004803603602081101561089357600080fd5b5035610ffe565b610302600480360360208110156108b057600080fd5b50356001600160a01b0316611010565b601780546001600160a01b0319166001600160a01b038316179055604051600d906000805160206111d283398151915290600090a250565b6109006110ae565b6040516004906000805160206111d283398151915290600090a2565b600f819055604051600c906000805160206111d283398151915290600090a250565b601680546001600160a01b0319166001600160a01b0383161790556040516007906000805160206111d283398151915290600090a250565b60195481565b81601c60008154811061098b57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080601c6001815481106109c857fe5b6000918252602082200180546001600160a01b0319166001600160a01b0393909316929092179091556040516005916000805160206111d283398151915291a25050565b600c5481565b6001600160a01b0381166000908152600a60205260409020545b919050565b6011546001600160a01b031681565b6001600160a01b031660009081526008602052604090206001015490565b600c8190556040516009906000805160206111d283398151915290600090a250565b60185481565b6019819055604051600e906000805160206111d283398151915290600090a250565b6000908152600560205260409020546001600160a01b031690565b6017546001600160a01b031681565b6012546001600160a01b0382811691161415610b22576040805162461bcd60e51b815260206004820152600a602482015269086aaa4be9eaaa8849eb60b31b604482015290519081900360640190fd5b601054604080516319dc7ae560e31b81526001600160a01b038481166004830152600060248301819052925193169263cee3d7289260448084019391929182900301818387803b158015610b7557600080fd5b505af1158015610b89573d6000803e3d6000fd5b5050604051600192506000805160206111d28339815191529150600090a250565b600b5460ff1690565b6014546001600160a01b031681565b600e5481565b6001600160a01b0316600090815260086020526040902060030154600160a01b900460ff1690565b600060078281548110610bff57fe5b6000918252602090912001546001600160a01b031692915050565b60095490565b60015490565b6001600160a01b039081166000908152600860205260409020600301541690565b601b8190556040516010906000805160206111d283398151915290600090a250565b60005460ff1690565b600f5481565b600d5481565b60035490565b610c8c61114e565b6040516003906000805160206111d283398151915290600090a2565b60045490565b6016546001600160a01b031681565b600d819055604051600a906000805160206111d283398151915290600090a250565b6000805b600954811015610d305760098181548110610cfa57fe5b60009182526020909120600290910201546001600160a01b0384811691161415610d28576001915050610a2c565b600101610ce3565b50600092915050565b60188190556040516008906000805160206111d283398151915290600090a250565b6013546001600160a01b031681565b8051825114610daf576040805162461bcd60e51b815260206004820152600c60248201526b0aea49e9c8ebe988a9c8ea8960a31b604482015290519081900360640190fd5b60005b8251811015610e1957818181518110610dc757fe5b6020026020010151601d6000858481518110610ddf57fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916911515919091179055600101610db2565b506040516006906000805160206111d283398151915290600090a25050565b601a5481565b600e819055604051600b906000805160206111d283398151915290600090a250565b6012546001600160a01b031681565b600060098281548110610e7e57fe5b60009182526020909120600290910201546001600160a01b031692915050565b60025490565b6015546001600160a01b031681565b601b5481565b60075490565b6010546040805163722dbe7360e11b81526001600160a01b03858116600483015284151560248301529151919092169163e45b7ce691604480830192600092919082900301818387803b158015610f1557600080fd5b505af1158015610f29573d6000803e3d6000fd5b5050604051600292506000805160206111d28339815191529150600090a25050565b6008602052600090815260409020805460018201546002830154600390930154919290916001600160a01b03811690600160a01b900460ff1685565b6001600160a01b031660009081526008602052604090206002015490565b601a819055604051600f906000805160206111d283398151915290600090a250565b600060098281548110610fd657fe5b9060005260206000209060020201600101549050919050565b6010546001600160a01b031681565b60009081526006602052604090205490565b601280546001600160a01b0319166001600160a01b03838116918217909255601054604080516319dc7ae560e31b81526004810193909352600160248401525192169163cee3d7289160448082019260009290919082900301818387803b15801561107a57600080fd5b505af115801561108e573d6000803e3d6000fd5b5050604051600092506000805160206111d283398151915291508290a250565b6110b6610baa565b6110fe576040805162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b604482015290519081900360640190fd5b600b805460ff191690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6111316111cd565b604080516001600160a01b039092168252519081900360200190a1565b611156610baa565b1561119b576040805162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b604482015290519081900360640190fd5b600b805460ff191660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2586111315b339056feea8787f128d10b2cc0317b0c3960f9ad447f7f6c1ed189db1083ccffd20f456ea264697066735822122027492c3eaf54f39c148898ccecdeb6f846448206f00d0173c047e92198cf367564736f6c634300060b0033"

// DeployRollupAdminFacet deploys a new Ethereum contract, binding an instance of RollupAdminFacet to it.
func DeployRollupAdminFacet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupAdminFacet, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupAdminFacetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupAdminFacetBin), backend)
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

// SequencerInboxMaxDelayBlocks is a free data retrieval call binding the contract method 0x14828f92.
//
// Solidity: function sequencerInboxMaxDelayBlocks() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) SequencerInboxMaxDelayBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "sequencerInboxMaxDelayBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerInboxMaxDelayBlocks is a free data retrieval call binding the contract method 0x14828f92.
//
// Solidity: function sequencerInboxMaxDelayBlocks() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) SequencerInboxMaxDelayBlocks() (*big.Int, error) {
	return _RollupAdminFacet.Contract.SequencerInboxMaxDelayBlocks(&_RollupAdminFacet.CallOpts)
}

// SequencerInboxMaxDelayBlocks is a free data retrieval call binding the contract method 0x14828f92.
//
// Solidity: function sequencerInboxMaxDelayBlocks() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) SequencerInboxMaxDelayBlocks() (*big.Int, error) {
	return _RollupAdminFacet.Contract.SequencerInboxMaxDelayBlocks(&_RollupAdminFacet.CallOpts)
}

// SequencerInboxMaxDelaySeconds is a free data retrieval call binding the contract method 0xaddd6784.
//
// Solidity: function sequencerInboxMaxDelaySeconds() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCaller) SequencerInboxMaxDelaySeconds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RollupAdminFacet.contract.Call(opts, &out, "sequencerInboxMaxDelaySeconds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SequencerInboxMaxDelaySeconds is a free data retrieval call binding the contract method 0xaddd6784.
//
// Solidity: function sequencerInboxMaxDelaySeconds() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetSession) SequencerInboxMaxDelaySeconds() (*big.Int, error) {
	return _RollupAdminFacet.Contract.SequencerInboxMaxDelaySeconds(&_RollupAdminFacet.CallOpts)
}

// SequencerInboxMaxDelaySeconds is a free data retrieval call binding the contract method 0xaddd6784.
//
// Solidity: function sequencerInboxMaxDelaySeconds() view returns(uint256)
func (_RollupAdminFacet *RollupAdminFacetCallerSession) SequencerInboxMaxDelaySeconds() (*big.Int, error) {
	return _RollupAdminFacet.Contract.SequencerInboxMaxDelaySeconds(&_RollupAdminFacet.CallOpts)
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

// SetArbGasSpeedLimitPerBlock is a paid mutator transaction binding the contract method 0xcd6bf14d.
//
// Solidity: function setArbGasSpeedLimitPerBlock(uint256 newArbGasSpeedLimitPerBlock) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetArbGasSpeedLimitPerBlock(opts *bind.TransactOpts, newArbGasSpeedLimitPerBlock *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setArbGasSpeedLimitPerBlock", newArbGasSpeedLimitPerBlock)
}

// SetArbGasSpeedLimitPerBlock is a paid mutator transaction binding the contract method 0xcd6bf14d.
//
// Solidity: function setArbGasSpeedLimitPerBlock(uint256 newArbGasSpeedLimitPerBlock) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetArbGasSpeedLimitPerBlock(newArbGasSpeedLimitPerBlock *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetArbGasSpeedLimitPerBlock(&_RollupAdminFacet.TransactOpts, newArbGasSpeedLimitPerBlock)
}

// SetArbGasSpeedLimitPerBlock is a paid mutator transaction binding the contract method 0xcd6bf14d.
//
// Solidity: function setArbGasSpeedLimitPerBlock(uint256 newArbGasSpeedLimitPerBlock) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetArbGasSpeedLimitPerBlock(newArbGasSpeedLimitPerBlock *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetArbGasSpeedLimitPerBlock(&_RollupAdminFacet.TransactOpts, newArbGasSpeedLimitPerBlock)
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

// SetSequencerInboxMaxDelayBlocks is a paid mutator transaction binding the contract method 0x4a56bab6.
//
// Solidity: function setSequencerInboxMaxDelayBlocks(uint256 newSequencerInboxMaxDelayBlocks) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetSequencerInboxMaxDelayBlocks(opts *bind.TransactOpts, newSequencerInboxMaxDelayBlocks *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setSequencerInboxMaxDelayBlocks", newSequencerInboxMaxDelayBlocks)
}

// SetSequencerInboxMaxDelayBlocks is a paid mutator transaction binding the contract method 0x4a56bab6.
//
// Solidity: function setSequencerInboxMaxDelayBlocks(uint256 newSequencerInboxMaxDelayBlocks) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetSequencerInboxMaxDelayBlocks(newSequencerInboxMaxDelayBlocks *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetSequencerInboxMaxDelayBlocks(&_RollupAdminFacet.TransactOpts, newSequencerInboxMaxDelayBlocks)
}

// SetSequencerInboxMaxDelayBlocks is a paid mutator transaction binding the contract method 0x4a56bab6.
//
// Solidity: function setSequencerInboxMaxDelayBlocks(uint256 newSequencerInboxMaxDelayBlocks) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetSequencerInboxMaxDelayBlocks(newSequencerInboxMaxDelayBlocks *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetSequencerInboxMaxDelayBlocks(&_RollupAdminFacet.TransactOpts, newSequencerInboxMaxDelayBlocks)
}

// SetSequencerInboxMaxDelaySeconds is a paid mutator transaction binding the contract method 0xf322c0bb.
//
// Solidity: function setSequencerInboxMaxDelaySeconds(uint256 newSequencerInboxMaxDelaySeconds) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactor) SetSequencerInboxMaxDelaySeconds(opts *bind.TransactOpts, newSequencerInboxMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.contract.Transact(opts, "setSequencerInboxMaxDelaySeconds", newSequencerInboxMaxDelaySeconds)
}

// SetSequencerInboxMaxDelaySeconds is a paid mutator transaction binding the contract method 0xf322c0bb.
//
// Solidity: function setSequencerInboxMaxDelaySeconds(uint256 newSequencerInboxMaxDelaySeconds) returns()
func (_RollupAdminFacet *RollupAdminFacetSession) SetSequencerInboxMaxDelaySeconds(newSequencerInboxMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetSequencerInboxMaxDelaySeconds(&_RollupAdminFacet.TransactOpts, newSequencerInboxMaxDelaySeconds)
}

// SetSequencerInboxMaxDelaySeconds is a paid mutator transaction binding the contract method 0xf322c0bb.
//
// Solidity: function setSequencerInboxMaxDelaySeconds(uint256 newSequencerInboxMaxDelaySeconds) returns()
func (_RollupAdminFacet *RollupAdminFacetTransactorSession) SetSequencerInboxMaxDelaySeconds(newSequencerInboxMaxDelaySeconds *big.Int) (*types.Transaction, error) {
	return _RollupAdminFacet.Contract.SetSequencerInboxMaxDelaySeconds(&_RollupAdminFacet.TransactOpts, newSequencerInboxMaxDelaySeconds)
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

// RollupAdminFacetNodesDestroyedIterator is returned from FilterNodesDestroyed and is used to iterate over the raw logs and unpacked data for NodesDestroyed events raised by the RollupAdminFacet contract.
type RollupAdminFacetNodesDestroyedIterator struct {
	Event *RollupAdminFacetNodesDestroyed // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetNodesDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetNodesDestroyed)
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
		it.Event = new(RollupAdminFacetNodesDestroyed)
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
func (it *RollupAdminFacetNodesDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetNodesDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetNodesDestroyed represents a NodesDestroyed event raised by the RollupAdminFacet contract.
type RollupAdminFacetNodesDestroyed struct {
	StartNode *big.Int
	EndNode   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNodesDestroyed is a free log retrieval operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterNodesDestroyed(opts *bind.FilterOpts, startNode []*big.Int, endNode []*big.Int) (*RollupAdminFacetNodesDestroyedIterator, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetNodesDestroyedIterator{contract: _RollupAdminFacet.contract, event: "NodesDestroyed", logs: logs, sub: sub}, nil
}

// WatchNodesDestroyed is a free log subscription operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchNodesDestroyed(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetNodesDestroyed, startNode []*big.Int, endNode []*big.Int) (event.Subscription, error) {

	var startNodeRule []interface{}
	for _, startNodeItem := range startNode {
		startNodeRule = append(startNodeRule, startNodeItem)
	}
	var endNodeRule []interface{}
	for _, endNodeItem := range endNode {
		endNodeRule = append(endNodeRule, endNodeItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "NodesDestroyed", startNodeRule, endNodeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetNodesDestroyed)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
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

// ParseNodesDestroyed is a log parse operation binding the contract event 0x9455d3b30b954764ff9f6ebe9120d1d8bb842ba3923bb5e0f71317b04d8a272d.
//
// Solidity: event NodesDestroyed(uint256 indexed startNode, uint256 indexed endNode)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseNodesDestroyed(log types.Log) (*RollupAdminFacetNodesDestroyed, error) {
	event := new(RollupAdminFacetNodesDestroyed)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "NodesDestroyed", log); err != nil {
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

// RollupAdminFacetStakerReassignedIterator is returned from FilterStakerReassigned and is used to iterate over the raw logs and unpacked data for StakerReassigned events raised by the RollupAdminFacet contract.
type RollupAdminFacetStakerReassignedIterator struct {
	Event *RollupAdminFacetStakerReassigned // Event containing the contract specifics and raw log

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
func (it *RollupAdminFacetStakerReassignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupAdminFacetStakerReassigned)
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
		it.Event = new(RollupAdminFacetStakerReassigned)
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
func (it *RollupAdminFacetStakerReassignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupAdminFacetStakerReassignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupAdminFacetStakerReassigned represents a StakerReassigned event raised by the RollupAdminFacet contract.
type RollupAdminFacetStakerReassigned struct {
	Staker  common.Address
	NewNode *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStakerReassigned is a free log retrieval operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_RollupAdminFacet *RollupAdminFacetFilterer) FilterStakerReassigned(opts *bind.FilterOpts, staker []common.Address) (*RollupAdminFacetStakerReassignedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.FilterLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return &RollupAdminFacetStakerReassignedIterator{contract: _RollupAdminFacet.contract, event: "StakerReassigned", logs: logs, sub: sub}, nil
}

// WatchStakerReassigned is a free log subscription operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_RollupAdminFacet *RollupAdminFacetFilterer) WatchStakerReassigned(opts *bind.WatchOpts, sink chan<- *RollupAdminFacetStakerReassigned, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _RollupAdminFacet.contract.WatchLogs(opts, "StakerReassigned", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupAdminFacetStakerReassigned)
				if err := _RollupAdminFacet.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
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

// ParseStakerReassigned is a log parse operation binding the contract event 0x8d475b2086edfd0e8badb5d935b5e14f0e09686368da62192932aaf86c137870.
//
// Solidity: event StakerReassigned(address indexed staker, uint256 newNode)
func (_RollupAdminFacet *RollupAdminFacetFilterer) ParseStakerReassigned(log types.Log) (*RollupAdminFacetStakerReassigned, error) {
	event := new(RollupAdminFacetStakerReassigned)
	if err := _RollupAdminFacet.contract.UnpackLog(event, "StakerReassigned", log); err != nil {
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
