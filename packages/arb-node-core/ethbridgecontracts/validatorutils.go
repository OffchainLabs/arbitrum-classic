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
const ValidatorUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkConfirmableNextNode\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkDecidableNextNode\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.ConfirmType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkRejectableNextNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkRejectableOutOfOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"node1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"node2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findNodeConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findStakerConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"refundStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"refundableStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakedNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"successorNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorUtilsFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorUtilsFuncSigs = map[string]string{
	"770db480": "checkConfirmableNextNode(address)",
	"2b1062cf": "checkDecidableNextNode(address,uint256,uint256,uint256,uint256)",
	"422e3550": "checkRejectableNextNode(address,uint256,uint256,uint256,uint256)",
	"ea3ca9b2": "checkRejectableOutOfOrder(address)",
	"3082d029": "findNodeConflict(address,uint256,uint256,uint256)",
	"7988ad37": "findStakerConflict(address,address,address,uint256)",
	"e48a5f7b": "getConfig(address)",
	"d08272d2": "refundStakers(address,address[])",
	"7464ae06": "refundableStakers(address)",
	"c308eaaf": "stakedNodes(address,address)",
	"8730825e": "successorNodes(address,uint256)",
}

// ValidatorUtilsBin is the compiled bytecode used for deploying new contracts.
var ValidatorUtilsBin = "0x608060405234801561001057600080fd5b506123ac806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80637988ad37116100715780637988ad37146102885780638730825e146102c4578063c308eaaf146102f0578063d08272d21461031e578063e48a5f7b1461039e578063ea3ca9b2146103f3576100a9565b80632b1062cf146100ae5780633082d02914610124578063422e35501461018b5780637464ae06146101ea578063770db48014610260575b600080fd5b6100ec600480360360a08110156100c457600080fd5b506001600160a01b038135169060208101359060408101359060608101359060800135610419565b604051808460038111156100fc57fe5b8152602001838152602001826001600160a01b03168152602001935050505060405180910390f35b61015c6004803603608081101561013a57600080fd5b506001600160a01b0381351690602081013590604081013590606001356108d8565b6040518084600381111561016c57fe5b8152602001838152602001828152602001935050505060405180910390f35b6101c9600480360360a08110156101a157600080fd5b506001600160a01b038135169060208101359060408101359060608101359060800135610d6d565b604080519283526001600160a01b0390911660208301528051918290030190f35b6102106004803603602081101561020057600080fd5b50356001600160a01b0316611034565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561024c578181015183820152602001610234565b505050509050019250505060405180910390f35b6102866004803603602081101561027657600080fd5b50356001600160a01b0316611295565b005b61015c6004803603608081101561029e57600080fd5b506001600160a01b038135811691602081013582169160408201351690606001356115e9565b610210600480360360408110156102da57600080fd5b506001600160a01b03813516906020013561170a565b6102106004803603604081101561030657600080fd5b506001600160a01b03813581169160200135166118c3565b6102866004803603604081101561033457600080fd5b6001600160a01b03823516919081019060408101602082013564010000000081111561035f57600080fd5b82018360208201111561037157600080fd5b8035906020019184602083028401116401000000008311171561039357600080fd5b509092509050611add565b6103c4600480360360208110156103b457600080fd5b50356001600160a01b0316611b74565b604080519485526020850193909352838301919091526001600160a01b03166060830152519081900360800190f35b6102866004803603602081101561040957600080fd5b50356001600160a01b0316611d3b565b6000806000876001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b15801561045757600080fd5b505afa925050508015610468575060015b61047a575060009150819050806108cd565b6000886001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156104b557600080fd5b505afa1580156104c9573d6000803e3d6000fd5b505050506040513d60208110156104df57600080fd5b50516040805163d735e21d60e01b815290519192506000916001600160a01b038c169163d735e21d916004808301926020929190829003018186803b15801561052757600080fd5b505afa15801561053b573d6000803e3d6000fd5b505050506040513d602081101561055157600080fd5b5051604080516238a78560e71b81526004810183905290519192506000916001600160a01b038d1691631c53c280916024808301926020929190829003018186803b15801561059f57600080fd5b505afa1580156105b3573d6000803e3d6000fd5b505050506040513d60208110156105c957600080fd5b505160408051631422135960e11b81526004810186905290519192506001600160a01b0383169163284426b291602480820192600092909190829003018186803b15801561061657600080fd5b505afa925050508015610627575060015b61063057610643565b60026000809550955095505050506108cd565b8a6001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b15801561067c57600080fd5b505afa92505050801561068d575060015b6106a45760008060009550955095505050506108cd565b60008b6001600160a01b03166304a28064836040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156106f357600080fd5b505afa158015610707573d6000803e3d6000fd5b505050506040513d602081101561071d57600080fd5b81019080805190602001909291905050509050816001600160a01b0316636cf00e7e828e6001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561077957600080fd5b505afa15801561078d573d6000803e3d6000fd5b505050506040513d60208110156107a357600080fd5b5051604080516001600160e01b031960e086901b16815292909101600483015260248201889052516044808301926000929190829003018186803b1580156107ea57600080fd5b505afa9250505080156107fb575060015b61080457610818565b6001600080965096509650505050506108cd565b816001600160a01b0316631a8a092b826040518263ffffffff1660e01b81526004018082815260200191505060006040518083038186803b15801561085c57600080fd5b505afa92505050801561086d575060015b610885576000806000965096509650505050506108cd565b600080600061089c8f8f88600101018f8f8f611f46565b925092509250826108be576000806000995099509950505050505050506108cd565b60039950909750955050505050505b955095509592505050565b600080600080876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561091757600080fd5b505afa15801561092b573d6000803e3d6000fd5b505050506040513d602081101561094157600080fd5b5051604080516238a78560e71b8152600481018a905290519192506000916001600160a01b038b1691631c53c280916024808301926020929190829003018186803b15801561098f57600080fd5b505afa1580156109a3573d6000803e3d6000fd5b505050506040513d60208110156109b957600080fd5b5051604080516311e7249560e21b815290516001600160a01b039092169163479c925491600480820192602092909190829003018186803b1580156109fd57600080fd5b505afa158015610a11573d6000803e3d6000fd5b505050506040513d6020811015610a2757600080fd5b5051604080516238a78560e71b8152600481018a905290519192506000916001600160a01b038c1691631c53c280916024808301926020929190829003018186803b158015610a7557600080fd5b505afa158015610a89573d6000803e3d6000fd5b505050506040513d6020811015610a9f57600080fd5b5051604080516311e7249560e21b815290516001600160a01b039092169163479c925491600480820192602092909190829003018186803b158015610ae357600080fd5b505afa158015610af7573d6000803e3d6000fd5b505050506040513d6020811015610b0d57600080fd5b5051905060005b87811015610d5457888a1415610b375760008a8a96509650965050505050610d63565b81831415610b525760018a8a96509650965050505050610d63565b83831080610b5f57508382105b15610b7857600260008096509650965050505050610d63565b888a1015610c68578198508a6001600160a01b0316631c53c2808a6040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610bc757600080fd5b505afa158015610bdb573d6000803e3d6000fd5b505050506040513d6020811015610bf157600080fd5b5051604080516311e7249560e21b815290516001600160a01b039092169163479c925491600480820192602092909190829003018186803b158015610c3557600080fd5b505afa158015610c49573d6000803e3d6000fd5b505050506040513d6020811015610c5f57600080fd5b50519150610d4c565b8299508a6001600160a01b0316631c53c2808b6040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015610caf57600080fd5b505afa158015610cc3573d6000803e3d6000fd5b505050506040513d6020811015610cd957600080fd5b5051604080516311e7249560e21b815290516001600160a01b039092169163479c925491600480820192602092909190829003018186803b158015610d1d57600080fd5b505afa158015610d31573d6000803e3d6000fd5b505050506040513d6020811015610d4757600080fd5b505192505b600101610b14565b50600389899550955095505050505b9450945094915050565b600080866001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b158015610da957600080fd5b505afa158015610dbd573d6000803e3d6000fd5b50505050866001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b158015610dfa57600080fd5b505afa158015610e0e573d6000803e3d6000fd5b505050506000876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610e4d57600080fd5b505afa158015610e61573d6000803e3d6000fd5b505050506040513d6020811015610e7757600080fd5b5051604080516238a78560e71b81526004810183905290519192506000916001600160a01b038b1691631c53c280916024808301926020929190829003018186803b158015610ec557600080fd5b505afa158015610ed9573d6000803e3d6000fd5b505050506040513d6020811015610eef57600080fd5b505160408051630128a01960e21b81526001600160a01b038084166004830181905292519394509192631a8a092b928d16916304a28064916024808301926020929190829003018186803b158015610f4657600080fd5b505afa158015610f5a573d6000803e3d6000fd5b505050506040513d6020811015610f7057600080fd5b5051604080516001600160e01b031960e085901b1681526004810192909252516024808301926000929190829003018186803b158015610faf57600080fd5b505afa158015610fc3573d6000803e3d6000fd5b505050506000806000610fde8c8c87600101018c8c8c611f46565b92509250925082611023576040805162461bcd60e51b815260206004820152600a6024820152694e4f5f4558414d504c4560b01b604482015290519081900360640190fd5b909b909a5098505050505050505050565b60606000826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561107157600080fd5b505afa158015611085573d6000803e3d6000fd5b505050506040513d602081101561109b57600080fd5b5051905060608167ffffffffffffffff811180156110b857600080fd5b506040519080825280602002602001820160405280156110e2578160200160208202803683370190505b5090506000846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561112057600080fd5b505afa158015611134573d6000803e3d6000fd5b505050506040513d602081101561114a57600080fd5b505190506000805b8481101561128a576000876001600160a01b031663348e50c6836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156111a057600080fd5b505afa1580156111b4573d6000803e3d6000fd5b505050506040513d60208110156111ca57600080fd5b50516040805163729cfe3b60e01b81526001600160a01b0380841660048301529151929350600092918b169163729cfe3b9160248082019260a092909190829003018186803b15801561121c57600080fd5b505afa158015611230573d6000803e3d6000fd5b505050506040513d60a081101561124657600080fd5b50602001519050848111611280578186858151811061126157fe5b6001600160a01b03909216602092830291909101909101526001909301925b5050600101611152565b508252509392505050565b806001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b1580156112ce57600080fd5b505afa1580156112e2573d6000803e3d6000fd5b50505050806001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b15801561131f57600080fd5b505afa158015611333573d6000803e3d6000fd5b505050506000816001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561137257600080fd5b505afa158015611386573d6000803e3d6000fd5b505050506040513d602081101561139c57600080fd5b5051604080516365f7f80d60e01b815290519192506000916001600160a01b038516916365f7f80d916004808301926020929190829003018186803b1580156113e457600080fd5b505afa1580156113f8573d6000803e3d6000fd5b505050506040513d602081101561140e57600080fd5b50516040805163dff6978760e01b815290519192506000916001600160a01b0386169163dff69787916004808301926020929190829003018186803b15801561145657600080fd5b505afa15801561146a573d6000803e3d6000fd5b505050506040513d602081101561148057600080fd5b5051604080516238a78560e71b81526004810186905290519192506000916001600160a01b03871691631c53c280916024808301926020929190829003018186803b1580156114ce57600080fd5b505afa1580156114e2573d6000803e3d6000fd5b505050506040513d60208110156114f857600080fd5b505160408051630128a01960e21b81526001600160a01b0380841660048301529151929350600092918816916304a2806491602480820192602092909190829003018186803b15801561154a57600080fd5b505afa15801561155e573d6000803e3d6000fd5b505050506040513d602081101561157457600080fd5b505160408051633678073f60e11b815285830160048201526024810187905290519192506001600160a01b03841691636cf00e7e91604480820192600092909190829003018186803b1580156115c957600080fd5b505afa1580156115dd573d6000803e3d6000fd5b50505050505050505050565b600080600080876001600160a01b031663729cfe3b886040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060a06040518083038186803b15801561163c57600080fd5b505afa158015611650573d6000803e3d6000fd5b505050506040513d60a081101561166657600080fd5b50602001516040805163729cfe3b60e01b81526001600160a01b0389811660048301529151929350600092918b169163729cfe3b9160248082019260a092909190829003018186803b1580156116bb57600080fd5b505afa1580156116cf573d6000803e3d6000fd5b505050506040513d60a08110156116e557600080fd5b506020015190506116f8898383896108d8565b94509450945050509450945094915050565b60408051620186a08082526230d4208201909252606091829190602082016230d400803683370190505090506000600184015b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561177657600080fd5b505afa15801561178a573d6000803e3d6000fd5b505050506040513d60208110156117a057600080fd5b505181116118b9576000866001600160a01b0316631c53c280836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156117ee57600080fd5b505afa158015611802573d6000803e3d6000fd5b505050506040513d602081101561181857600080fd5b5051604080516311e7249560e21b8152905191925087916001600160a01b0384169163479c9254916004808301926020929190829003018186803b15801561185f57600080fd5b505afa158015611873573d6000803e3d6000fd5b505050506040513d602081101561188957600080fd5b505114156118b0578184848151811061189e57fe5b60209081029190910101526001909201915b5060010161173d565b5081529392505050565b60408051620186a08082526230d4208201909252606091829190602082016230d40080368337019050509050600080856001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561192b57600080fd5b505afa15801561193f573d6000803e3d6000fd5b505050506040513d602081101561195557600080fd5b505190505b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561199357600080fd5b505afa1580156119a7573d6000803e3d6000fd5b505050506040513d60208110156119bd57600080fd5b505181116118b9576000866001600160a01b0316631c53c280836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b158015611a0b57600080fd5b505afa158015611a1f573d6000803e3d6000fd5b505050506040513d6020811015611a3557600080fd5b5051604080516348b4573960e11b81526001600160a01b038981166004830152915192935090831691639168ae7291602480820192602092909190829003018186803b158015611a8457600080fd5b505afa158015611a98573d6000803e3d6000fd5b505050506040513d6020811015611aae57600080fd5b505115611ad45781848481518110611ac257fe5b60209081029190910101526001909201915b5060010161195a565b8060005b81811015611b6d57846001600160a01b0316637427be51858584818110611b0457fe5b905060200201356001600160a01b03166040518263ffffffff1660e01b815260040180826001600160a01b03168152602001915050600060405180830381600087803b158015611b5357600080fd5b505af1925050508015611b64575060015b50600101611ae1565b5050505050565b600080600080846001600160a01b03166346c2781a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611bb357600080fd5b505afa158015611bc7573d6000803e3d6000fd5b505050506040513d6020811015611bdd57600080fd5b505160408051632f47788360e11b815290519195506001600160a01b03871691635e8ef10691600480820192602092909190829003018186803b158015611c2357600080fd5b505afa158015611c37573d6000803e3d6000fd5b505050506040513d6020811015611c4d57600080fd5b5051604080516376e7e23b60e01b815290519194506001600160a01b038716916376e7e23b91600480820192602092909190829003018186803b158015611c9357600080fd5b505afa158015611ca7573d6000803e3d6000fd5b505050506040513d6020811015611cbd57600080fd5b50516040805163051ed6a360e41b815290519193506001600160a01b038716916351ed6a3091600480820192602092909190829003018186803b158015611d0357600080fd5b505afa158015611d17573d6000803e3d6000fd5b505050506040513d6020811015611d2d57600080fd5b505193959294509092919050565b806001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b158015611d7457600080fd5b505afa158015611d88573d6000803e3d6000fd5b505050506000816001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611dc757600080fd5b505afa158015611ddb573d6000803e3d6000fd5b505050506040513d6020811015611df157600080fd5b50516040805163d735e21d60e01b815290519192506000916001600160a01b0385169163d735e21d916004808301926020929190829003018186803b158015611e3957600080fd5b505afa158015611e4d573d6000803e3d6000fd5b505050506040513d6020811015611e6357600080fd5b5051604080516238a78560e71b81526004810183905290519192506000916001600160a01b03861691631c53c280916024808301926020929190829003018186803b158015611eb157600080fd5b505afa158015611ec5573d6000803e3d6000fd5b505050506040513d6020811015611edb57600080fd5b505160408051631422135960e11b81526004810186905290519192506001600160a01b0383169163284426b291602480820192600092909190829003018186803b158015611f2857600080fd5b505afa158015611f3c573d6000803e3d6000fd5b5050505050505050565b600080600080886001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611f8557600080fd5b505afa158015611f99573d6000803e3d6000fd5b505050506040513d6020811015611faf57600080fd5b5051905080881115611fcc576000806000935093509350506108cd565b87810387811115611fda5750865b61216e8a8a8c6001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561201857600080fd5b505afa15801561202c573d6000803e3d6000fd5b505050506040513d602081101561204257600080fd5b8101908080519060200190929190505050848e6001600160a01b031663ad71bd368d8d6040518363ffffffff1660e01b8152600401808381526020018281526020019250505060006040518083038186803b1580156120a057600080fd5b505afa1580156120b4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405260208110156120dd57600080fd5b81019080805160405193929190846401000000008211156120fd57600080fd5b90830190602082018581111561211257600080fd5b825186602082028301116401000000008211171561212f57600080fd5b82525081516020918201928201910280838360005b8381101561215c578181015183820152602001612144565b50505050905001604052505050612181565b9450945094505050955095509592505050565b6000806000808451905060005b868111612362576000818a01905060008b6001600160a01b0316631c53c280836040518263ffffffff1660e01b81526004018082815260200191505060206040518083038186803b1580156121e257600080fd5b505afa1580156121f6573d6000803e3d6000fd5b505050506040513d602081101561220c57600080fd5b5051604080516311e7249560e21b815290519192508b916001600160a01b0384169163479c9254916004808301926020929190829003018186803b15801561225357600080fd5b505afa158015612267573d6000803e3d6000fd5b505050506040513d602081101561227d57600080fd5b50511461228b57505061235a565b60005b8481101561235657816001600160a01b0316639168ae728a83815181106122b157fe5b60200260200101516040518263ffffffff1660e01b815260040180826001600160a01b0316815260200191505060206040518083038186803b1580156122f657600080fd5b505afa15801561230a573d6000803e3d6000fd5b505050506040513d602081101561232057600080fd5b50511561234e576001838a838151811061233657fe5b602002602001015197509750975050505050506108cd565b60010161228e565b5050505b60010161218e565b506000998a9950899850965050505050505056fea2646970667358221220413afdb24563f762b05c7132248ce2cdd666df9bf31296558f57b432f09e3bd864736f6c634300060c0033"

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

// FindNodeConflict is a free data retrieval call binding the contract method 0x3082d029.
//
// Solidity: function findNodeConflict(address rollup, uint256 node1, uint256 node2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCaller) FindNodeConflict(opts *bind.CallOpts, rollup common.Address, node1 *big.Int, node2 *big.Int, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "findNodeConflict", rollup, node1, node2, maxDepth)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// FindNodeConflict is a free data retrieval call binding the contract method 0x3082d029.
//
// Solidity: function findNodeConflict(address rollup, uint256 node1, uint256 node2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsSession) FindNodeConflict(rollup common.Address, node1 *big.Int, node2 *big.Int, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindNodeConflict(&_ValidatorUtils.CallOpts, rollup, node1, node2, maxDepth)
}

// FindNodeConflict is a free data retrieval call binding the contract method 0x3082d029.
//
// Solidity: function findNodeConflict(address rollup, uint256 node1, uint256 node2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCallerSession) FindNodeConflict(rollup common.Address, node1 *big.Int, node2 *big.Int, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindNodeConflict(&_ValidatorUtils.CallOpts, rollup, node1, node2, maxDepth)
}

// FindStakerConflict is a free data retrieval call binding the contract method 0x7988ad37.
//
// Solidity: function findStakerConflict(address rollup, address staker1, address staker2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCaller) FindStakerConflict(opts *bind.CallOpts, rollup common.Address, staker1 common.Address, staker2 common.Address, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ValidatorUtils.contract.Call(opts, &out, "findStakerConflict", rollup, staker1, staker2, maxDepth)

	if err != nil {
		return *new(uint8), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// FindStakerConflict is a free data retrieval call binding the contract method 0x7988ad37.
//
// Solidity: function findStakerConflict(address rollup, address staker1, address staker2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsSession) FindStakerConflict(rollup common.Address, staker1 common.Address, staker2 common.Address, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindStakerConflict(&_ValidatorUtils.CallOpts, rollup, staker1, staker2, maxDepth)
}

// FindStakerConflict is a free data retrieval call binding the contract method 0x7988ad37.
//
// Solidity: function findStakerConflict(address rollup, address staker1, address staker2, uint256 maxDepth) view returns(uint8, uint256, uint256)
func (_ValidatorUtils *ValidatorUtilsCallerSession) FindStakerConflict(rollup common.Address, staker1 common.Address, staker2 common.Address, maxDepth *big.Int) (uint8, *big.Int, *big.Int, error) {
	return _ValidatorUtils.Contract.FindStakerConflict(&_ValidatorUtils.CallOpts, rollup, staker1, staker2, maxDepth)
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
