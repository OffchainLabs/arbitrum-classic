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

// ValidatorUtilsABI is the input ABI used to generate the binding from.
const ValidatorUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkConfirmableNextNode\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkDecidableNextNode\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.ConfirmType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkRejectableNextNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkRejectableOutOfOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"refundStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"refundableStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakedNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"successorNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorUtilsFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorUtilsFuncSigs = map[string]string{
	"770db480": "checkConfirmableNextNode(address)",
	"2b1062cf": "checkDecidableNextNode(address,uint256,uint256,uint256,uint256)",
	"422e3550": "checkRejectableNextNode(address,uint256,uint256,uint256,uint256)",
	"ea3ca9b2": "checkRejectableOutOfOrder(address)",
	"e48a5f7b": "getConfig(address)",
	"d08272d2": "refundStakers(address,address[])",
	"7464ae06": "refundableStakers(address)",
	"c308eaaf": "stakedNodes(address,address)",
	"8730825e": "successorNodes(address,uint256)",
}

// ValidatorUtilsBin is the compiled bytecode used for deploying new contracts.
var ValidatorUtilsBin = "0x608060405234801561001057600080fd5b50611ea2806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638730825e116100665780638730825e1461020b578063c308eaaf14610237578063d08272d214610265578063e48a5f7b146102e5578063ea3ca9b21461033a57610093565b80632b1062cf14610098578063422e35501461010e5780637464ae061461016d578063770db480146101e3575b600080fd5b6100d6600480360360a08110156100ae57600080fd5b506001600160a01b038135169060208101359060408101359060608101359060800135610360565b604051808460038111156100e657fe5b8152602001838152602001826001600160a01b03168152602001935050505060405180910390f35b61014c600480360360a081101561012457600080fd5b506001600160a01b03813516906020810135906040810135906060810135906080013561089b565b604080519283526001600160a01b0390911660208301528051918290030190f35b6101936004803603602081101561018357600080fd5b50356001600160a01b0316610b5d565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156101cf5781810151838201526020016101b7565b505050509050019250505060405180910390f35b610209600480360360208110156101f957600080fd5b50356001600160a01b0316610dbe565b005b6101936004803603604081101561022157600080fd5b506001600160a01b038135169060200135611112565b6101936004803603604081101561024d57600080fd5b506001600160a01b03813581169160200135166112cb565b6102096004803603604081101561027b57600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156102a657600080fd5b8201836020820111156102b857600080fd5b803590602001918460208302840111640100000000831117156102da57600080fd5b5090925090506114e5565b61030b600480360360208110156102fb57600080fd5b50356001600160a01b031661157c565b604080519485526020850193909352838301919091526001600160a01b03166060830152519081900360800190f35b6102096004803603602081101561035057600080fd5b50356001600160a01b0316611743565b6000806000876001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b15801561039e57600080fd5b505afa9250505080156103af575060015b6103c157506000915081905080610890565b6000886001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156103fc57600080fd5b505afa158015610410573d6000803e3d6000fd5b505050506040513d602081101561042657600080fd5b50516040805163d735e21d60e01b815290519192506000916001600160a01b038c169163d735e21d916004808301926020929190829003018186803b15801561046e57600080fd5b505afa158015610482573d6000803e3d6000fd5b505050506040513d602081101561049857600080fd5b5051604080516238a78560e71b81526004810183905290519192506000916001600160a01b038d1691631c53c280916024808301926020929190829003018186803b1580156104e657600080fd5b505afa1580156104fa573d6000803e3d6000fd5b505050506040513d602081101561051057600080fd5b505160408051631422135960e11b81526004810186905290519192506001600160a01b0383169163284426b291602480820192600092909190829003018186803b15801561055d57600080fd5b505afa92505050801561056e575060015b6105775761058a565b6002600080955095509550505050610890565b8a6001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b1580156105c357600080fd5b505afa9250505080156105d4575060015b6105eb576000806000955095509550505050610890565b60008b6001600160a01b03166304a28064836040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b15801561063a57600080fd5b505afa15801561064e573d6000803e3d6000fd5b505050506040513d602081101561066457600080fd5b81019080805190602001909291905050509050816001600160a01b0316636cf00e7e828e6001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b1580156106c057600080fd5b505afa1580156106d4573d6000803e3d6000fd5b505050506040513d60208110156106ea57600080fd5b5051604080516001600160e01b031960e086901b16815292909101600483015260248201889052516044808301926000929190829003018186803b15801561073157600080fd5b505afa925050508015610742575060015b61074b5761075f565b600160008096509650965050505050610890565b816001600160a01b0316631a8a092b8d6001600160a01b03166304a28064856040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156107bb57600080fd5b505afa1580156107cf573d6000803e3d6000fd5b505050506040513d60208110156107e557600080fd5b5051604080516001600160e01b031960e085901b1681526004810192909252516024808301926000929190829003018186803b15801561082457600080fd5b505afa925050508015610835575060015b61084d57600080600096509650965050505050610890565b600080600061085f8f8f8f8f8f61194e565b9250925092508261088157600080600099509950995050505050505050610890565b60039950909750955050505050505b955095509592505050565b600080866001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b1580156108d757600080fd5b505afa1580156108eb573d6000803e3d6000fd5b50505050866001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b15801561092857600080fd5b505afa15801561093c573d6000803e3d6000fd5b505050506000876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561097b57600080fd5b505afa15801561098f573d6000803e3d6000fd5b505050506040513d60208110156109a557600080fd5b5051604080516238a78560e71b81526004810183905290519192506000916001600160a01b038b1691631c53c280916024808301926020929190829003018186803b1580156109f357600080fd5b505afa158015610a07573d6000803e3d6000fd5b505050506040513d6020811015610a1d57600080fd5b505160408051630128a01960e21b81526001600160a01b038084166004830181905292519394509192631a8a092b928d16916304a28064916024808301926020929190829003018186803b158015610a7457600080fd5b505afa158015610a88573d6000803e3d6000fd5b505050506040513d6020811015610a9e57600080fd5b5051604080516001600160e01b031960e085901b1681526004810192909252516024808301926000929190829003018186803b158015610add57600080fd5b505afa158015610af1573d6000803e3d6000fd5b505050506000806000610b078c8c8c8c8c61194e565b92509250925082610b4c576040805162461bcd60e51b815260206004820152600a6024820152694e4f5f4558414d504c4560b01b604482015290519081900360640190fd5b909b909a5098505050505050505050565b60606000826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015610b9a57600080fd5b505afa158015610bae573d6000803e3d6000fd5b505050506040513d6020811015610bc457600080fd5b5051905060608167ffffffffffffffff81118015610be157600080fd5b50604051908082528060200260200182016040528015610c0b578160200160208202803683370190505b5090506000846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c4957600080fd5b505afa158015610c5d573d6000803e3d6000fd5b505050506040513d6020811015610c7357600080fd5b505190506000805b84811015610db3576000876001600160a01b031663348e50c6836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610cc957600080fd5b505afa158015610cdd573d6000803e3d6000fd5b505050506040513d6020811015610cf357600080fd5b50516040805163729cfe3b60e01b81526001600160a01b0380841660048301529151929350600092918b169163729cfe3b9160248082019260a092909190829003018186803b158015610d4557600080fd5b505afa158015610d59573d6000803e3d6000fd5b505050506040513d60a0811015610d6f57600080fd5b50602001519050848111610da95781868581518110610d8a57fe5b6001600160a01b03909216602092830291909101909101526001909301925b5050600101610c7b565b508252509392505050565b806001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b158015610df757600080fd5b505afa158015610e0b573d6000803e3d6000fd5b50505050806001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b158015610e4857600080fd5b505afa158015610e5c573d6000803e3d6000fd5b505050506000816001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610e9b57600080fd5b505afa158015610eaf573d6000803e3d6000fd5b505050506040513d6020811015610ec557600080fd5b5051604080516365f7f80d60e01b815290519192506000916001600160a01b038516916365f7f80d916004808301926020929190829003018186803b158015610f0d57600080fd5b505afa158015610f21573d6000803e3d6000fd5b505050506040513d6020811015610f3757600080fd5b50516040805163dff6978760e01b815290519192506000916001600160a01b0386169163dff69787916004808301926020929190829003018186803b158015610f7f57600080fd5b505afa158015610f93573d6000803e3d6000fd5b505050506040513d6020811015610fa957600080fd5b5051604080516238a78560e71b81526004810186905290519192506000916001600160a01b03871691631c53c280916024808301926020929190829003018186803b158015610ff757600080fd5b505afa15801561100b573d6000803e3d6000fd5b505050506040513d602081101561102157600080fd5b505160408051630128a01960e21b81526001600160a01b0380841660048301529151929350600092918816916304a2806491602480820192602092909190829003018186803b15801561107357600080fd5b505afa158015611087573d6000803e3d6000fd5b505050506040513d602081101561109d57600080fd5b505160408051633678073f60e11b815285830160048201526024810187905290519192506001600160a01b03841691636cf00e7e91604480820192600092909190829003018186803b1580156110f257600080fd5b505afa158015611106573d6000803e3d6000fd5b50505050505050505050565b60408051620186a08082526230d4208201909252606091829190602082016230d400803683370190505090506000600184015b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561117e57600080fd5b505afa158015611192573d6000803e3d6000fd5b505050506040513d60208110156111a857600080fd5b505181116112c1576000866001600160a01b0316631c53c280836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156111f657600080fd5b505afa15801561120a573d6000803e3d6000fd5b505050506040513d602081101561122057600080fd5b5051604080516311e7249560e21b8152905191925087916001600160a01b0384169163479c9254916004808301926020929190829003018186803b15801561126757600080fd5b505afa15801561127b573d6000803e3d6000fd5b505050506040513d602081101561129157600080fd5b505114156112b857818484815181106112a657fe5b60209081029190910101526001909201915b50600101611145565b5081529392505050565b60408051620186a08082526230d4208201909252606091829190602082016230d40080368337019050509050600080856001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561133357600080fd5b505afa158015611347573d6000803e3d6000fd5b505050506040513d602081101561135d57600080fd5b505190505b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561139b57600080fd5b505afa1580156113af573d6000803e3d6000fd5b505050506040513d60208110156113c557600080fd5b505181116112c1576000866001600160a01b0316631c53c280836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b15801561141357600080fd5b505afa158015611427573d6000803e3d6000fd5b505050506040513d602081101561143d57600080fd5b5051604080516348b4573960e11b81526001600160a01b038981166004830152915192935090831691639168ae7291602480820192602092909190829003018186803b15801561148c57600080fd5b505afa1580156114a0573d6000803e3d6000fd5b505050506040513d60208110156114b657600080fd5b5051156114dc57818484815181106114ca57fe5b60209081029190910101526001909201915b50600101611362565b8060005b8181101561157557846001600160a01b0316637427be5185858481811061150c57fe5b905060200201356001600160a01b03166040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b15801561155b57600080fd5b505af192505050801561156c575060015b506001016114e9565b5050505050565b600080600080846001600160a01b03166346c2781a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156115bb57600080fd5b505afa1580156115cf573d6000803e3d6000fd5b505050506040513d60208110156115e557600080fd5b505160408051632f47788360e11b815290519195506001600160a01b03871691635e8ef10691600480820192602092909190829003018186803b15801561162b57600080fd5b505afa15801561163f573d6000803e3d6000fd5b505050506040513d602081101561165557600080fd5b5051604080516376e7e23b60e01b815290519194506001600160a01b038716916376e7e23b91600480820192602092909190829003018186803b15801561169b57600080fd5b505afa1580156116af573d6000803e3d6000fd5b505050506040513d60208110156116c557600080fd5b50516040805163051ed6a360e41b815290519193506001600160a01b038716916351ed6a3091600480820192602092909190829003018186803b15801561170b57600080fd5b505afa15801561171f573d6000803e3d6000fd5b505050506040513d602081101561173557600080fd5b505193959294509092919050565b806001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b15801561177c57600080fd5b505afa158015611790573d6000803e3d6000fd5b505050506000816001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117cf57600080fd5b505afa1580156117e3573d6000803e3d6000fd5b505050506040513d60208110156117f957600080fd5b50516040805163d735e21d60e01b815290519192506000916001600160a01b0385169163d735e21d916004808301926020929190829003018186803b15801561184157600080fd5b505afa158015611855573d6000803e3d6000fd5b505050506040513d602081101561186b57600080fd5b5051604080516238a78560e71b81526004810183905290519192506000916001600160a01b03861691631c53c280916024808301926020929190829003018186803b1580156118b957600080fd5b505afa1580156118cd573d6000803e3d6000fd5b505050506040513d60208110156118e357600080fd5b505160408051631422135960e11b81526004810186905290519192506001600160a01b0383169163284426b291602480820192600092909190829003018186803b15801561193057600080fd5b505afa158015611944573d6000803e3d6000fd5b5050505050505050565b600080600080886001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561198d57600080fd5b505afa1580156119a1573d6000803e3d6000fd5b505050506040513d60208110156119b757600080fd5b5051604080516356b8de9b60e11b8152600481018990526024810188905290519192506060916001600160a01b038c169163ad71bd36916044808301926000929190829003018186803b158015611a0d57600080fd5b505afa158015611a21573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015611a4a57600080fd5b8101908080516040519392919084640100000000821115611a6a57600080fd5b908301906020820185811115611a7f57600080fd5b8251866020820283011164010000000082111715611a9c57600080fd5b82525081516020918201928201910280838360005b83811015611ac9578181015183820152602001611ab1565b50505050905001604052505050905060008a6001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611b1357600080fd5b505afa158015611b27573d6000803e3d6000fd5b505050506040513d6020811015611b3d57600080fd5b50519050828a01600101811015611b61576000806000955095509550505050610890565b828a01810389811115611b715750885b611b7d8c828d86611b92565b96509650965050505050955095509592505050565b600080600080876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611bd157600080fd5b505afa158015611be5573d6000803e3d6000fd5b505050506040513d6020811015611bfb57600080fd5b5051604080516365f7f80d60e01b815290519192506000916001600160a01b038b16916365f7f80d916004808301926020929190829003018186803b158015611c4357600080fd5b505afa158015611c57573d6000803e3d6000fd5b505050506040513d6020811015611c6d57600080fd5b5051865190915060005b898111611e52576000818a866001010101905060008c6001600160a01b0316631c53c280836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015611cd057600080fd5b505afa158015611ce4573d6000803e3d6000fd5b505050506040513d6020811015611cfa57600080fd5b5051604080516311e7249560e21b8152905191925086916001600160a01b0384169163479c9254916004808301926020929190829003018186803b158015611d4157600080fd5b505afa158015611d55573d6000803e3d6000fd5b505050506040513d6020811015611d6b57600080fd5b505114611d79575050611e4a565b60005b84811015611e4657816001600160a01b0316639168ae728c8381518110611d9f57fe5b60200260200101516040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b158015611de457600080fd5b505afa158015611df8573d6000803e3d6000fd5b505050506040513d6020811015611e0e57600080fd5b505115611e3e576001838c8381518110611e2457fe5b602002602001015199509950995050505050505050611e62565b600101611d7c565b5050505b600101611c77565b5060008060009550955095505050505b945094509491505056fea2646970667358221220b7d94324dd0e59dc2e9f83e639b02adc9a06b22ca1b4657e4a537e732751178364736f6c634300060c0033"

// DeployValidatorUtils deploys a new Ethereum contract, binding an instance of ValidatorUtils to it.
func DeployValidatorUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ValidatorUtils, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorUtilsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ValidatorUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ValidatorUtils{ValidatorUtilsCaller: ValidatorUtilsCaller{contract: contract}, ValidatorUtilsTransactor: ValidatorUtilsTransactor{contract: contract}, ValidatorUtilsFilterer: ValidatorUtilsFilterer{contract: contract}}, nil
}

// ValidatorUtils is an auto generated Go binding around an Ethereum contract.
type ValidatorUtils struct {
	ValidatorUtilsCaller     // Read-only binding to the contract
	ValidatorUtilsTransactor // Write-only binding to the contract
	ValidatorUtilsFilterer   // Log filterer for contract events
}

// ValidatorUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorUtilsSession struct {
	Contract     *ValidatorUtils   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorUtilsCallerSession struct {
	Contract *ValidatorUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ValidatorUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorUtilsTransactorSession struct {
	Contract     *ValidatorUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ValidatorUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorUtilsRaw struct {
	Contract *ValidatorUtils // Generic contract binding to access the raw methods on
}

// ValidatorUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorUtilsCallerRaw struct {
	Contract *ValidatorUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorUtilsTransactorRaw struct {
	Contract *ValidatorUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidatorUtils creates a new instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtils(address common.Address, backend bind.ContractBackend) (*ValidatorUtils, error) {
	contract, err := bindValidatorUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtils{ValidatorUtilsCaller: ValidatorUtilsCaller{contract: contract}, ValidatorUtilsTransactor: ValidatorUtilsTransactor{contract: contract}, ValidatorUtilsFilterer: ValidatorUtilsFilterer{contract: contract}}, nil
}

// NewValidatorUtilsCaller creates a new read-only instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtilsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorUtilsCaller, error) {
	contract, err := bindValidatorUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtilsCaller{contract: contract}, nil
}

// NewValidatorUtilsTransactor creates a new write-only instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorUtilsTransactor, error) {
	contract, err := bindValidatorUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtilsTransactor{contract: contract}, nil
}

// NewValidatorUtilsFilterer creates a new log filterer instance of ValidatorUtils, bound to a specific deployed contract.
func NewValidatorUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorUtilsFilterer, error) {
	contract, err := bindValidatorUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorUtilsFilterer{contract: contract}, nil
}

// bindValidatorUtils binds a generic wrapper to an already deployed contract.
func bindValidatorUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorUtils *ValidatorUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorUtils.Contract.ValidatorUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorUtils *ValidatorUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.ValidatorUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorUtils *ValidatorUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.ValidatorUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ValidatorUtils *ValidatorUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ValidatorUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ValidatorUtils *ValidatorUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ValidatorUtils *ValidatorUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.contract.Transact(opts, method, params...)
}

// CheckConfirmableNextNode is a free data retrieval call binding the contract method 0x770db480.
//
// Solidity: function checkConfirmableNextNode(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCaller) CheckConfirmableNextNode(opts *bind.CallOpts, rollup common.Address) error {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "checkConfirmableNextNode", rollup)

	if err != nil {
		return err
	}

	return err

}

// CheckConfirmableNextNode is a free data retrieval call binding the contract method 0x770db480.
//
// Solidity: function checkConfirmableNextNode(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsSession) CheckConfirmableNextNode(rollup common.Address) error {
	return _ValidatorUtils.Contract.CheckConfirmableNextNode(&_ValidatorUtils.CallOpts, rollup)
}

// CheckConfirmableNextNode is a free data retrieval call binding the contract method 0x770db480.
//
// Solidity: function checkConfirmableNextNode(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCallerSession) CheckConfirmableNextNode(rollup common.Address) error {
	return _ValidatorUtils.Contract.CheckConfirmableNextNode(&_ValidatorUtils.CallOpts, rollup)
}

// CheckDecidableNextNode is a free data retrieval call binding the contract method 0x2b1062cf.
//
// Solidity: function checkDecidableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint8, uint256, address)
func (_ValidatorUtils *ValidatorUtilsCaller) CheckDecidableNextNode(opts *bind.CallOpts, rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (uint8, *big.Int, common.Address, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "checkDecidableNextNode", rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// CheckDecidableNextNode is a free data retrieval call binding the contract method 0x2b1062cf.
//
// Solidity: function checkDecidableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint8, uint256, address)
func (_ValidatorUtils *ValidatorUtilsSession) CheckDecidableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (uint8, *big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.CheckDecidableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
}

// CheckDecidableNextNode is a free data retrieval call binding the contract method 0x2b1062cf.
//
// Solidity: function checkDecidableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint8, uint256, address)
func (_ValidatorUtils *ValidatorUtilsCallerSession) CheckDecidableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (uint8, *big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.CheckDecidableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
}

// CheckRejectableNextNode is a free data retrieval call binding the contract method 0x422e3550.
//
// Solidity: function checkRejectableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint256, address)
func (_ValidatorUtils *ValidatorUtilsCaller) CheckRejectableNextNode(opts *bind.CallOpts, rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "checkRejectableNextNode", rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// CheckRejectableNextNode is a free data retrieval call binding the contract method 0x422e3550.
//
// Solidity: function checkRejectableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint256, address)
func (_ValidatorUtils *ValidatorUtilsSession) CheckRejectableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (*big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.CheckRejectableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
}

// CheckRejectableNextNode is a free data retrieval call binding the contract method 0x422e3550.
//
// Solidity: function checkRejectableNextNode(address rollup, uint256 startNodeOffset, uint256 maxNodeCount, uint256 startStakerIndex, uint256 maxStakerCount) view returns(uint256, address)
func (_ValidatorUtils *ValidatorUtilsCallerSession) CheckRejectableNextNode(rollup common.Address, startNodeOffset *big.Int, maxNodeCount *big.Int, startStakerIndex *big.Int, maxStakerCount *big.Int) (*big.Int, common.Address, error) {
	return _ValidatorUtils.Contract.CheckRejectableNextNode(&_ValidatorUtils.CallOpts, rollup, startNodeOffset, maxNodeCount, startStakerIndex, maxStakerCount)
}

// CheckRejectableOutOfOrder is a free data retrieval call binding the contract method 0xea3ca9b2.
//
// Solidity: function checkRejectableOutOfOrder(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCaller) CheckRejectableOutOfOrder(opts *bind.CallOpts, rollup common.Address) error {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "checkRejectableOutOfOrder", rollup)

	if err != nil {
		return err
	}

	return err

}

// CheckRejectableOutOfOrder is a free data retrieval call binding the contract method 0xea3ca9b2.
//
// Solidity: function checkRejectableOutOfOrder(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsSession) CheckRejectableOutOfOrder(rollup common.Address) error {
	return _ValidatorUtils.Contract.CheckRejectableOutOfOrder(&_ValidatorUtils.CallOpts, rollup)
}

// CheckRejectableOutOfOrder is a free data retrieval call binding the contract method 0xea3ca9b2.
//
// Solidity: function checkRejectableOutOfOrder(address rollup) view returns()
func (_ValidatorUtils *ValidatorUtilsCallerSession) CheckRejectableOutOfOrder(rollup common.Address) error {
	return _ValidatorUtils.Contract.CheckRejectableOutOfOrder(&_ValidatorUtils.CallOpts, rollup)
}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 challengePeriodBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken)
func (_ValidatorUtils *ValidatorUtilsCaller) GetConfig(opts *bind.CallOpts, rollup common.Address) (struct {
	ChallengePeriodBlocks    *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "getConfig", rollup)

	outstruct := new(struct {
		ChallengePeriodBlocks    *big.Int
		ArbGasSpeedLimitPerBlock *big.Int
		BaseStake                *big.Int
		StakeToken               common.Address
	})

	outstruct.ChallengePeriodBlocks = out[0].(*big.Int)
	outstruct.ArbGasSpeedLimitPerBlock = out[1].(*big.Int)
	outstruct.BaseStake = out[2].(*big.Int)
	outstruct.StakeToken = out[3].(common.Address)

	return *outstruct, err

}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 challengePeriodBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken)
func (_ValidatorUtils *ValidatorUtilsSession) GetConfig(rollup common.Address) (struct {
	ChallengePeriodBlocks    *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}, error) {
	return _ValidatorUtils.Contract.GetConfig(&_ValidatorUtils.CallOpts, rollup)
}

// GetConfig is a free data retrieval call binding the contract method 0xe48a5f7b.
//
// Solidity: function getConfig(address rollup) view returns(uint256 challengePeriodBlocks, uint256 arbGasSpeedLimitPerBlock, uint256 baseStake, address stakeToken)
func (_ValidatorUtils *ValidatorUtilsCallerSession) GetConfig(rollup common.Address) (struct {
	ChallengePeriodBlocks    *big.Int
	ArbGasSpeedLimitPerBlock *big.Int
	BaseStake                *big.Int
	StakeToken               common.Address
}, error) {
	return _ValidatorUtils.Contract.GetConfig(&_ValidatorUtils.CallOpts, rollup)
}

// RefundableStakers is a free data retrieval call binding the contract method 0x7464ae06.
//
// Solidity: function refundableStakers(address rollup) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsCaller) RefundableStakers(opts *bind.CallOpts, rollup common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "refundableStakers", rollup)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// RefundableStakers is a free data retrieval call binding the contract method 0x7464ae06.
//
// Solidity: function refundableStakers(address rollup) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsSession) RefundableStakers(rollup common.Address) ([]common.Address, error) {
	return _ValidatorUtils.Contract.RefundableStakers(&_ValidatorUtils.CallOpts, rollup)
}

// RefundableStakers is a free data retrieval call binding the contract method 0x7464ae06.
//
// Solidity: function refundableStakers(address rollup) view returns(address[])
func (_ValidatorUtils *ValidatorUtilsCallerSession) RefundableStakers(rollup common.Address) ([]common.Address, error) {
	return _ValidatorUtils.Contract.RefundableStakers(&_ValidatorUtils.CallOpts, rollup)
}

// StakedNodes is a free data retrieval call binding the contract method 0xc308eaaf.
//
// Solidity: function stakedNodes(address rollup, address staker) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCaller) StakedNodes(opts *bind.CallOpts, rollup common.Address, staker common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "stakedNodes", rollup, staker)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// StakedNodes is a free data retrieval call binding the contract method 0xc308eaaf.
//
// Solidity: function stakedNodes(address rollup, address staker) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsSession) StakedNodes(rollup common.Address, staker common.Address) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.StakedNodes(&_ValidatorUtils.CallOpts, rollup, staker)
}

// StakedNodes is a free data retrieval call binding the contract method 0xc308eaaf.
//
// Solidity: function stakedNodes(address rollup, address staker) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCallerSession) StakedNodes(rollup common.Address, staker common.Address) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.StakedNodes(&_ValidatorUtils.CallOpts, rollup, staker)
}

// SuccessorNodes is a free data retrieval call binding the contract method 0x8730825e.
//
// Solidity: function successorNodes(address rollup, uint256 nodeNum) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCaller) SuccessorNodes(opts *bind.CallOpts, rollup common.Address, nodeNum *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "successorNodes", rollup, nodeNum)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// SuccessorNodes is a free data retrieval call binding the contract method 0x8730825e.
//
// Solidity: function successorNodes(address rollup, uint256 nodeNum) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsSession) SuccessorNodes(rollup common.Address, nodeNum *big.Int) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.SuccessorNodes(&_ValidatorUtils.CallOpts, rollup, nodeNum)
}

// SuccessorNodes is a free data retrieval call binding the contract method 0x8730825e.
//
// Solidity: function successorNodes(address rollup, uint256 nodeNum) view returns(uint256[])
func (_ValidatorUtils *ValidatorUtilsCallerSession) SuccessorNodes(rollup common.Address, nodeNum *big.Int) ([]*big.Int, error) {
	return _ValidatorUtils.Contract.SuccessorNodes(&_ValidatorUtils.CallOpts, rollup, nodeNum)
}

// RefundStakers is a paid mutator transaction binding the contract method 0xd08272d2.
//
// Solidity: function refundStakers(address rollup, address[] stakers) returns()
func (_ValidatorUtils *ValidatorUtilsTransactor) RefundStakers(opts *bind.TransactOpts, rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _ValidatorUtils.contract.Transact(opts, "refundStakers", rollup, stakers)
}

// RefundStakers is a paid mutator transaction binding the contract method 0xd08272d2.
//
// Solidity: function refundStakers(address rollup, address[] stakers) returns()
func (_ValidatorUtils *ValidatorUtilsSession) RefundStakers(rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.RefundStakers(&_ValidatorUtils.TransactOpts, rollup, stakers)
}

// RefundStakers is a paid mutator transaction binding the contract method 0xd08272d2.
//
// Solidity: function refundStakers(address rollup, address[] stakers) returns()
func (_ValidatorUtils *ValidatorUtilsTransactorSession) RefundStakers(rollup common.Address, stakers []common.Address) (*types.Transaction, error) {
	return _ValidatorUtils.Contract.RefundStakers(&_ValidatorUtils.TransactOpts, rollup, stakers)
}
