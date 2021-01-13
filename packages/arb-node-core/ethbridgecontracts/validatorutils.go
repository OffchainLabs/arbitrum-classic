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
const ValidatorUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkConfirmableNextNode\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkDecidableNextNode\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.ConfirmType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkRejectableNextNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkRejectableOutOfOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"node1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"node2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findNodeConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findStakerConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"addresspayable[]\",\"name\":\"stakers\",\"type\":\"address[]\"}],\"name\":\"refundStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"refundableStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakedNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"successorNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

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
var ValidatorUtilsBin = "0x608060405234801561001057600080fd5b506126b5806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80637988ad37116100715780637988ad37146101515780638730825e14610164578063c308eaaf14610184578063d08272d214610197578063e48a5f7b146101aa578063ea3ca9b2146101cd576100a9565b80632b1062cf146100ae5780633082d029146100d9578063422e3550146100fb5780637464ae061461011c578063770db4801461013c575b600080fd5b6100c16100bc36600461247b565b6101e0565b6040516100d09392919061256f565b60405180910390f35b6100ec6100e7366004612441565b6106a5565b6040516100d09392919061259a565b61010e61010936600461247b565b610b76565b6040516100d09291906125e9565b61012f61012a3660046121de565b610e3b565b6040516100d091906124ea565b61014f61014a3660046121de565b6110c5565b005b6100ec61015f366004612344565b61142f565b610177610172366004612416565b611557565b6040516100d09190612537565b61017761019236600461230c565b61172d565b61014f6101a5366004612394565b611971565b6101bd6101b83660046121de565b611a04565b6040516100d0949392919061260e565b61014f6101db3660046121de565b611bdd565b6000806000876001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b15801561021e57600080fd5b505afa92505050801561022f575060015b6102415750600091508190508061069a565b6000886001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561027c57600080fd5b505afa158015610290573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102b491906124be565b90506000896001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156102f157600080fd5b505afa158015610305573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061032991906124be565b905060008a6001600160a01b0316631c53c280836040518263ffffffff1660e01b815260040161035991906125e0565b60206040518083038186803b15801561037157600080fd5b505afa158015610385573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103a991906121bb565b604051631422135960e11b81529091506001600160a01b0382169063284426b2906103d89086906004016125e0565b60006040518083038186803b1580156103f057600080fd5b505afa925050508015610401575060015b61040a5761041d565b600260008095509550955050505061069a565b8a6001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b15801561045657600080fd5b505afa925050508015610467575060015b61047e57600080600095509550955050505061069a565b604051630128a01960e21b81526000906001600160a01b038d16906304a28064906104ad9085906004016124d6565b60206040518083038186803b1580156104c557600080fd5b505afa1580156104d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104fd91906124be565b9050816001600160a01b0316636cf00e7e828e6001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561054857600080fd5b505afa15801561055c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061058091906124be565b01866040518363ffffffff1660e01b815260040161059f929190612600565b60006040518083038186803b1580156105b757600080fd5b505afa9250505080156105c8575060015b6105d1576105e5565b60016000809650965096505050505061069a565b604051631a8a092b60e01b81526001600160a01b03831690631a8a092b906106119084906004016125e0565b60006040518083038186803b15801561062957600080fd5b505afa92505050801561063a575060015b6106525760008060009650965096505050505061069a565b60008060006106698f8f88600101018f8f8f611dfb565b9250925092508261068b5760008060009950995099505050505050505061069a565b60029950909750955050505050505b955095509592505050565b600080600080876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156106e457600080fd5b505afa1580156106f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071c91906124be565b90506000886001600160a01b0316631c53c280896040518263ffffffff1660e01b815260040161074c91906125e0565b60206040518083038186803b15801561076457600080fd5b505afa158015610778573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079c91906121bb565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156107d457600080fd5b505afa1580156107e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061080c91906124be565b90506000896001600160a01b0316631c53c280896040518263ffffffff1660e01b815260040161083c91906125e0565b60206040518083038186803b15801561085457600080fd5b505afa158015610868573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061088c91906121bb565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156108c457600080fd5b505afa1580156108d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108fc91906124be565b905060005b87811015610b5d57888a14156109245760008a8a96509650965050505050610b6c565b8183141561093f5760018a8a96509650965050505050610b6c565b8383108061094c57508382105b1561096557600260008096509650965050505050610b6c565b888a1015610a63578198508a6001600160a01b0316631c53c2808a6040518263ffffffff1660e01b815260040161099c91906125e0565b60206040518083038186803b1580156109b457600080fd5b505afa1580156109c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ec91906121bb565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610a2457600080fd5b505afa158015610a38573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a5c91906124be565b9150610b55565b8299508a6001600160a01b0316631c53c2808b6040518263ffffffff1660e01b8152600401610a9291906125e0565b60206040518083038186803b158015610aaa57600080fd5b505afa158015610abe573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae291906121bb565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610b1a57600080fd5b505afa158015610b2e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b5291906124be565b92505b600101610901565b50600389899550955095505050505b9450945094915050565b600080866001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b158015610bb257600080fd5b505afa158015610bc6573d6000803e3d6000fd5b50505050866001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b158015610c0357600080fd5b505afa158015610c17573d6000803e3d6000fd5b505050506000876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c5657600080fd5b505afa158015610c6a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c8e91906124be565b90506000886001600160a01b0316631c53c280836040518263ffffffff1660e01b8152600401610cbe91906125e0565b60206040518083038186803b158015610cd657600080fd5b505afa158015610cea573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0e91906121bb565b9050806001600160a01b0316631a8a092b8a6001600160a01b03166304a28064846040518263ffffffff1660e01b8152600401610d4b91906124d6565b60206040518083038186803b158015610d6357600080fd5b505afa158015610d77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d9b91906124be565b6040518263ffffffff1660e01b8152600401610db791906125e0565b60006040518083038186803b158015610dcf57600080fd5b505afa158015610de3573d6000803e3d6000fd5b505050506000806000610dfe8c8c87600101018c8c8c611dfb565b92509250925082610e2a5760405162461bcd60e51b8152600401610e21906125bc565b60405180910390fd5b909b909a5098505050505050505050565b60606000826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015610e7857600080fd5b505afa158015610e8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610eb091906124be565b905060608167ffffffffffffffff81118015610ecb57600080fd5b50604051908082528060200260200182016040528015610ef5578160200160208202803683370190505b5090506000846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f3357600080fd5b505afa158015610f47573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f6b91906124be565b90506000805b848110156110ba57604051631a47286360e11b81526000906001600160a01b0389169063348e50c690610fa89085906004016125e0565b60206040518083038186803b158015610fc057600080fd5b505afa158015610fd4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff891906121bb565b90506000886001600160a01b0316634e745f1f836040518263ffffffff1660e01b815260040161102891906124d6565b60806040518083038186803b15801561104057600080fd5b505afa158015611054573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061107891906122c1565b50509150508481116110b0578186858151811061109157fe5b6001600160a01b03909216602092830291909101909101526001909301925b5050600101610f71565b508252509392505050565b806001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b1580156110fe57600080fd5b505afa158015611112573d6000803e3d6000fd5b50505050806001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b15801561114f57600080fd5b505afa158015611163573d6000803e3d6000fd5b505050506000816001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156111a257600080fd5b505afa1580156111b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111da91906124be565b90506000826001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561121757600080fd5b505afa15801561122b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061124f91906124be565b90506000836001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561128c57600080fd5b505afa1580156112a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112c491906124be565b90506000846001600160a01b0316631c53c280856040518263ffffffff1660e01b81526004016112f491906125e0565b60206040518083038186803b15801561130c57600080fd5b505afa158015611320573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134491906121bb565b90506000856001600160a01b03166304a28064836040518263ffffffff1660e01b815260040161137491906124d6565b60206040518083038186803b15801561138c57600080fd5b505afa1580156113a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113c491906124be565b604051633678073f60e11b81529091506001600160a01b03831690636cf00e7e906113f790868501908890600401612600565b60006040518083038186803b15801561140f57600080fd5b505afa158015611423573d6000803e3d6000fd5b50505050505050505050565b600080600080876001600160a01b0316634e745f1f886040518263ffffffff1660e01b815260040161146191906124d6565b60806040518083038186803b15801561147957600080fd5b505afa15801561148d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b191906122c1565b50509150506000886001600160a01b0316634e745f1f886040518263ffffffff1660e01b81526004016114e491906124d6565b60806040518083038186803b1580156114fc57600080fd5b505afa158015611510573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061153491906122c1565b5050915050611545898383896106a5565b94509450945050509450945094915050565b60408051620186a08082526230d4208201909252606091829190602082016230d400803683370190505090506000600184015b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156115c357600080fd5b505afa1580156115d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115fb91906124be565b8111611721576040516238a78560e71b81526000906001600160a01b03881690631c53c2809061162f9085906004016125e0565b60206040518083038186803b15801561164757600080fd5b505afa15801561165b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061167f91906121bb565b905085816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156116bb57600080fd5b505afa1580156116cf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116f391906124be565b1415611718578184848151811061170657fe5b60209081029190910101526001909201915b5060010161158a565b50815290505b92915050565b60408051620186a08082526230d4208201909252606091829190602082016230d40080368337019050509050600080856001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561179557600080fd5b505afa1580156117a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117cd91906124be565b90505b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561180957600080fd5b505afa15801561181d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061184191906124be565b8111611721576040516238a78560e71b81526000906001600160a01b03881690631c53c280906118759085906004016125e0565b60206040518083038186803b15801561188d57600080fd5b505afa1580156118a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118c591906121bb565b6040516348b4573960e11b81529091506001600160a01b03821690639168ae72906118f49089906004016124d6565b60206040518083038186803b15801561190c57600080fd5b505afa158015611920573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061194491906122a5565b15611968578184848151811061195657fe5b60209081029190910101526001909201915b506001016117d0565b8060005b818110156119fd57846001600160a01b0316637427be5185858481811061199857fe5b90506020020160208101906119ad91906121de565b6040518263ffffffff1660e01b81526004016119c991906124d6565b600060405180830381600087803b1580156119e357600080fd5b505af19250505080156119f4575060015b50600101611975565b5050505050565b600080600080846001600160a01b03166346c2781a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611a4357600080fd5b505afa158015611a57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7b91906124be565b9350846001600160a01b0316635e8ef1066040518163ffffffff1660e01b815260040160206040518083038186803b158015611ab657600080fd5b505afa158015611aca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611aee91906124be565b9250846001600160a01b03166376e7e23b6040518163ffffffff1660e01b815260040160206040518083038186803b158015611b2957600080fd5b505afa158015611b3d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b6191906124be565b9150846001600160a01b03166351ed6a306040518163ffffffff1660e01b815260040160206040518083038186803b158015611b9c57600080fd5b505afa158015611bb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bd491906121bb565b90509193509193565b806001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b158015611c1657600080fd5b505afa158015611c2a573d6000803e3d6000fd5b505050506000816001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c6957600080fd5b505afa158015611c7d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ca191906124be565b90506000826001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611cde57600080fd5b505afa158015611cf2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d1691906124be565b90506000836001600160a01b0316631c53c280836040518263ffffffff1660e01b8152600401611d4691906125e0565b60206040518083038186803b158015611d5e57600080fd5b505afa158015611d72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d9691906121bb565b604051631422135960e11b81529091506001600160a01b0382169063284426b290611dc59086906004016125e0565b60006040518083038186803b158015611ddd57600080fd5b505afa158015611df1573d6000803e3d6000fd5b5050505050505050565b600080600080886001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611e3a57600080fd5b505afa158015611e4e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e7291906124be565b905080881115611e8d5760008060009350935093505061069a565b87810387811115611e9b5750865b611f998a8a8c6001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611ed957600080fd5b505afa158015611eed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f1191906124be565b848e6001600160a01b031663ad71bd368d8d6040518363ffffffff1660e01b8152600401611f40929190612600565b60006040518083038186803b158015611f5857600080fd5b505afa158015611f6c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611f9491908101906121fa565b611fac565b9450945094505050955095509592505050565b6000806000808451905060005b86811161219c576040516238a78560e71b8152898201906000906001600160a01b038d1690631c53c28090611ff29085906004016125e0565b60206040518083038186803b15801561200a57600080fd5b505afa15801561201e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061204291906121bb565b905089816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561207e57600080fd5b505afa158015612092573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120b691906124be565b146120c2575050612194565b60005b8481101561219057816001600160a01b0316639168ae728a83815181106120e857fe5b60200260200101516040518263ffffffff1660e01b815260040161210c91906124d6565b60206040518083038186803b15801561212457600080fd5b505afa158015612138573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061215c91906122a5565b15612188576001838a838151811061217057fe5b6020026020010151975097509750505050505061069a565b6001016120c5565b5050505b600101611fb9565b506000998a99508998509650505050505050565b805161172781612659565b6000602082840312156121cc578081fd5b81516121d781612659565b9392505050565b6000602082840312156121ef578081fd5b81356121d781612659565b6000602080838503121561220c578182fd5b825167ffffffffffffffff80821115612223578384fd5b818501915085601f830112612236578384fd5b815181811115612244578485fd5b8381029150612254848301612632565b8181528481019084860184860187018a101561226e578788fd5b8795505b83861015612298576122848a826121b0565b835260019590950194918601918601612272565b5098975050505050505050565b6000602082840312156122b6578081fd5b81516121d781612671565b600080600080608085870312156122d6578283fd5b84516122e181612671565b809450506020850151925060408501519150606085015161230181612659565b939692955090935050565b6000806040838503121561231e578182fd5b823561232981612659565b9150602083013561233981612659565b809150509250929050565b60008060008060808587031215612359578384fd5b843561236481612659565b9350602085013561237481612659565b9250604085013561238481612659565b9396929550929360600135925050565b6000806000604084860312156123a8578283fd5b83356123b381612659565b9250602084013567ffffffffffffffff808211156123cf578384fd5b818601915086601f8301126123e2578384fd5b8135818111156123f0578485fd5b8760208083028501011115612403578485fd5b6020830194508093505050509250925092565b60008060408385031215612428578182fd5b823561243381612659565b946020939093013593505050565b60008060008060808587031215612456578384fd5b843561246181612659565b966020860135965060408601359560600135945092505050565b600080600080600060a08688031215612492578283fd5b853561249d81612659565b97602087013597506040870135966060810135965060800135945092505050565b6000602082840312156124cf578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b8181101561252b5783516001600160a01b031683529284019291840191600101612506565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b8181101561252b57835183529284019291840191600101612553565b606081016003851061257d57fe5b93815260208101929092526001600160a01b031660409091015290565b60608101600485106125a857fe5b938152602081019290925260409091015290565b6020808252600a90820152694e4f5f4558414d504c4560b01b604082015260600190565b90815260200190565b9182526001600160a01b0316602082015260400190565b918252602082015260400190565b938452602084019290925260408301526001600160a01b0316606082015260800190565b60405181810167ffffffffffffffff8111828210171561265157600080fd5b604052919050565b6001600160a01b038116811461266e57600080fd5b50565b801515811461266e57600080fdfea2646970667358221220945dba8e01e2ec3a3d29ef6b1240f58b3c25cc314e45141921b6eae9d21e961464736f6c634300060c0033"

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
