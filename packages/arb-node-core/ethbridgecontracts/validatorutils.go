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
const ValidatorUtilsABI = "[{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkConfirmableNextNode\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkDecidableNextNode\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.ConfirmType\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startNodeOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxNodeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startStakerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxStakerCount\",\"type\":\"uint256\"}],\"name\":\"checkRejectableNextNode\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"checkRejectableOutOfOrder\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"node1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"node2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findNodeConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker2\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxDepth\",\"type\":\"uint256\"}],\"name\":\"findStakerConflict\",\"outputs\":[{\"internalType\":\"enumValidatorUtils.NodeConflict\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"challengePeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"stakeToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"}],\"name\":\"refundableStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"stakedNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"rollup\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeNum\",\"type\":\"uint256\"}],\"name\":\"successorNodes\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ValidatorUtilsFuncSigs maps the 4-byte function signature to its string representation.
var ValidatorUtilsFuncSigs = map[string]string{
	"770db480": "checkConfirmableNextNode(address)",
	"2b1062cf": "checkDecidableNextNode(address,uint256,uint256,uint256,uint256)",
	"422e3550": "checkRejectableNextNode(address,uint256,uint256,uint256,uint256)",
	"ea3ca9b2": "checkRejectableOutOfOrder(address)",
	"3082d029": "findNodeConflict(address,uint256,uint256,uint256)",
	"7988ad37": "findStakerConflict(address,address,address,uint256)",
	"e48a5f7b": "getConfig(address)",
	"7464ae06": "refundableStakers(address)",
	"c308eaaf": "stakedNodes(address,address)",
	"8730825e": "successorNodes(address,uint256)",
}

// ValidatorUtilsBin is the compiled bytecode used for deploying new contracts.
var ValidatorUtilsBin = "0x608060405234801561001057600080fd5b50612582806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80637988ad37116100665780637988ad37146101465780638730825e14610159578063c308eaaf14610179578063e48a5f7b1461018c578063ea3ca9b2146101af5761009e565b80632b1062cf146100a35780633082d029146100ce578063422e3550146100f05780637464ae0614610111578063770db48014610131575b600080fd5b6100b66100b1366004612348565b6101c2565b6040516100c59392919061243c565b60405180910390f35b6100e16100dc36600461230e565b610687565b6040516100c593929190612467565b6101036100fe366004612348565b610b58565b6040516100c59291906124b6565b61012461011f36600461223f565b610e1d565b6040516100c591906123b7565b61014461013f36600461223f565b6110a7565b005b6100e1610154366004612293565b611411565b61016c6101673660046122e3565b611539565b6040516100c59190612404565b61016c61018736600461225b565b61170f565b61019f61019a36600461223f565b611953565b6040516100c594939291906124db565b6101446101bd36600461223f565b611b2c565b6000806000876001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b15801561020057600080fd5b505afa925050508015610211575060015b6102235750600091508190508061067c565b6000886001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561025e57600080fd5b505afa158015610272573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610296919061238b565b90506000896001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156102d357600080fd5b505afa1580156102e7573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061030b919061238b565b905060008a6001600160a01b0316631c53c280836040518263ffffffff1660e01b815260040161033b91906124ad565b60206040518083038186803b15801561035357600080fd5b505afa158015610367573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061038b919061210a565b604051631422135960e11b81529091506001600160a01b0382169063284426b2906103ba9086906004016124ad565b60006040518083038186803b1580156103d257600080fd5b505afa9250505080156103e3575060015b6103ec576103ff565b600260008095509550955050505061067c565b8a6001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b15801561043857600080fd5b505afa925050508015610449575060015b61046057600080600095509550955050505061067c565b604051630128a01960e21b81526000906001600160a01b038d16906304a280649061048f9085906004016123a3565b60206040518083038186803b1580156104a757600080fd5b505afa1580156104bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104df919061238b565b9050816001600160a01b0316636cf00e7e828e6001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561052a57600080fd5b505afa15801561053e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610562919061238b565b01866040518363ffffffff1660e01b81526004016105819291906124cd565b60006040518083038186803b15801561059957600080fd5b505afa9250505080156105aa575060015b6105b3576105c7565b60016000809650965096505050505061067c565b604051631a8a092b60e01b81526001600160a01b03831690631a8a092b906105f39084906004016124ad565b60006040518083038186803b15801561060b57600080fd5b505afa92505050801561061c575060015b6106345760008060009650965096505050505061067c565b600080600061064b8f8f88600101018f8f8f611d4a565b9250925092508261066d5760008060009950995099505050505050505061067c565b60029950909750955050505050505b955095509592505050565b600080600080876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156106c657600080fd5b505afa1580156106da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106fe919061238b565b90506000886001600160a01b0316631c53c280896040518263ffffffff1660e01b815260040161072e91906124ad565b60206040518083038186803b15801561074657600080fd5b505afa15801561075a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061077e919061210a565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156107b657600080fd5b505afa1580156107ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107ee919061238b565b90506000896001600160a01b0316631c53c280896040518263ffffffff1660e01b815260040161081e91906124ad565b60206040518083038186803b15801561083657600080fd5b505afa15801561084a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061086e919061210a565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b1580156108a657600080fd5b505afa1580156108ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108de919061238b565b905060005b87811015610b3f57888a14156109065760008a8a96509650965050505050610b4e565b818314156109215760018a8a96509650965050505050610b4e565b8383108061092e57508382105b1561094757600260008096509650965050505050610b4e565b888a1015610a45578198508a6001600160a01b0316631c53c2808a6040518263ffffffff1660e01b815260040161097e91906124ad565b60206040518083038186803b15801561099657600080fd5b505afa1580156109aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ce919061210a565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610a0657600080fd5b505afa158015610a1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a3e919061238b565b9150610b37565b8299508a6001600160a01b0316631c53c2808b6040518263ffffffff1660e01b8152600401610a7491906124ad565b60206040518083038186803b158015610a8c57600080fd5b505afa158015610aa0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ac4919061210a565b6001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015610afc57600080fd5b505afa158015610b10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b34919061238b565b92505b6001016108e3565b50600389899550955095505050505b9450945094915050565b600080866001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b158015610b9457600080fd5b505afa158015610ba8573d6000803e3d6000fd5b50505050866001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b158015610be557600080fd5b505afa158015610bf9573d6000803e3d6000fd5b505050506000876001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c3857600080fd5b505afa158015610c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c70919061238b565b90506000886001600160a01b0316631c53c280836040518263ffffffff1660e01b8152600401610ca091906124ad565b60206040518083038186803b158015610cb857600080fd5b505afa158015610ccc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cf0919061210a565b9050806001600160a01b0316631a8a092b8a6001600160a01b03166304a28064846040518263ffffffff1660e01b8152600401610d2d91906123a3565b60206040518083038186803b158015610d4557600080fd5b505afa158015610d59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7d919061238b565b6040518263ffffffff1660e01b8152600401610d9991906124ad565b60006040518083038186803b158015610db157600080fd5b505afa158015610dc5573d6000803e3d6000fd5b505050506000806000610de08c8c87600101018c8c8c611d4a565b92509250925082610e0c5760405162461bcd60e51b8152600401610e0390612489565b60405180910390fd5b909b909a5098505050505050505050565b60606000826001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b158015610e5a57600080fd5b505afa158015610e6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e92919061238b565b905060608167ffffffffffffffff81118015610ead57600080fd5b50604051908082528060200260200182016040528015610ed7578160200160208202803683370190505b5090506000846001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015610f1557600080fd5b505afa158015610f29573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f4d919061238b565b90506000805b8481101561109c57604051631a47286360e11b81526000906001600160a01b0389169063348e50c690610f8a9085906004016124ad565b60206040518083038186803b158015610fa257600080fd5b505afa158015610fb6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fda919061210a565b90506000886001600160a01b0316634e745f1f836040518263ffffffff1660e01b815260040161100a91906123a3565b60806040518083038186803b15801561102257600080fd5b505afa158015611036573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061105a91906121f4565b5050915050848111611092578186858151811061107357fe5b6001600160a01b03909216602092830291909101909101526001909301925b5050600101610f53565b508252509392505050565b806001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b1580156110e057600080fd5b505afa1580156110f4573d6000803e3d6000fd5b50505050806001600160a01b031663be211c9a6040518163ffffffff1660e01b815260040160006040518083038186803b15801561113157600080fd5b505afa158015611145573d6000803e3d6000fd5b505050506000816001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561118457600080fd5b505afa158015611198573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111bc919061238b565b90506000826001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b1580156111f957600080fd5b505afa15801561120d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611231919061238b565b90506000836001600160a01b031663dff697876040518163ffffffff1660e01b815260040160206040518083038186803b15801561126e57600080fd5b505afa158015611282573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112a6919061238b565b90506000846001600160a01b0316631c53c280856040518263ffffffff1660e01b81526004016112d691906124ad565b60206040518083038186803b1580156112ee57600080fd5b505afa158015611302573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611326919061210a565b90506000856001600160a01b03166304a28064836040518263ffffffff1660e01b815260040161135691906123a3565b60206040518083038186803b15801561136e57600080fd5b505afa158015611382573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113a6919061238b565b604051633678073f60e11b81529091506001600160a01b03831690636cf00e7e906113d9908685019088906004016124cd565b60006040518083038186803b1580156113f157600080fd5b505afa158015611405573d6000803e3d6000fd5b50505050505050505050565b600080600080876001600160a01b0316634e745f1f886040518263ffffffff1660e01b815260040161144391906123a3565b60806040518083038186803b15801561145b57600080fd5b505afa15801561146f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061149391906121f4565b50509150506000886001600160a01b0316634e745f1f886040518263ffffffff1660e01b81526004016114c691906123a3565b60806040518083038186803b1580156114de57600080fd5b505afa1580156114f2573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061151691906121f4565b505091505061152789838389610687565b94509450945050509450945094915050565b60408051620186a08082526230d4208201909252606091829190602082016230d400803683370190505090506000600184015b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156115a557600080fd5b505afa1580156115b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115dd919061238b565b8111611703576040516238a78560e71b81526000906001600160a01b03881690631c53c280906116119085906004016124ad565b60206040518083038186803b15801561162957600080fd5b505afa15801561163d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611661919061210a565b905085816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b15801561169d57600080fd5b505afa1580156116b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116d5919061238b565b14156116fa57818484815181106116e857fe5b60209081029190910101526001909201915b5060010161156c565b50815290505b92915050565b60408051620186a08082526230d4208201909252606091829190602082016230d40080368337019050509050600080856001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561177757600080fd5b505afa15801561178b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117af919061238b565b90505b856001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b1580156117eb57600080fd5b505afa1580156117ff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611823919061238b565b8111611703576040516238a78560e71b81526000906001600160a01b03881690631c53c280906118579085906004016124ad565b60206040518083038186803b15801561186f57600080fd5b505afa158015611883573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118a7919061210a565b6040516348b4573960e11b81529091506001600160a01b03821690639168ae72906118d69089906004016123a3565b60206040518083038186803b1580156118ee57600080fd5b505afa158015611902573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061192691906121d8565b1561194a578184848151811061193857fe5b60209081029190910101526001909201915b506001016117b2565b600080600080846001600160a01b03166346c2781a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561199257600080fd5b505afa1580156119a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119ca919061238b565b9350846001600160a01b0316635e8ef1066040518163ffffffff1660e01b815260040160206040518083038186803b158015611a0557600080fd5b505afa158015611a19573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a3d919061238b565b9250846001600160a01b03166376e7e23b6040518163ffffffff1660e01b815260040160206040518083038186803b158015611a7857600080fd5b505afa158015611a8c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ab0919061238b565b9150846001600160a01b03166351ed6a306040518163ffffffff1660e01b815260040160206040518083038186803b158015611aeb57600080fd5b505afa158015611aff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b23919061210a565b90509193509193565b806001600160a01b03166373f33b066040518163ffffffff1660e01b815260040160006040518083038186803b158015611b6557600080fd5b505afa158015611b79573d6000803e3d6000fd5b505050506000816001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611bb857600080fd5b505afa158015611bcc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bf0919061238b565b90506000826001600160a01b031663d735e21d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611c2d57600080fd5b505afa158015611c41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c65919061238b565b90506000836001600160a01b0316631c53c280836040518263ffffffff1660e01b8152600401611c9591906124ad565b60206040518083038186803b158015611cad57600080fd5b505afa158015611cc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611ce5919061210a565b604051631422135960e11b81529091506001600160a01b0382169063284426b290611d149086906004016124ad565b60006040518083038186803b158015611d2c57600080fd5b505afa158015611d40573d6000803e3d6000fd5b5050505050505050565b600080600080886001600160a01b0316637ba9534a6040518163ffffffff1660e01b815260040160206040518083038186803b158015611d8957600080fd5b505afa158015611d9d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dc1919061238b565b905080881115611ddc5760008060009350935093505061067c565b87810387811115611dea5750865b611ee88a8a8c6001600160a01b03166365f7f80d6040518163ffffffff1660e01b815260040160206040518083038186803b158015611e2857600080fd5b505afa158015611e3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e60919061238b565b848e6001600160a01b031663ad71bd368d8d6040518363ffffffff1660e01b8152600401611e8f9291906124cd565b60006040518083038186803b158015611ea757600080fd5b505afa158015611ebb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611ee3919081019061212d565b611efb565b9450945094505050955095509592505050565b6000806000808451905060005b8681116120eb576040516238a78560e71b8152898201906000906001600160a01b038d1690631c53c28090611f419085906004016124ad565b60206040518083038186803b158015611f5957600080fd5b505afa158015611f6d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f91919061210a565b905089816001600160a01b031663479c92546040518163ffffffff1660e01b815260040160206040518083038186803b158015611fcd57600080fd5b505afa158015611fe1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612005919061238b565b146120115750506120e3565b60005b848110156120df57816001600160a01b0316639168ae728a838151811061203757fe5b60200260200101516040518263ffffffff1660e01b815260040161205b91906123a3565b60206040518083038186803b15801561207357600080fd5b505afa158015612087573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120ab91906121d8565b156120d7576001838a83815181106120bf57fe5b6020026020010151975097509750505050505061067c565b600101612014565b5050505b600101611f08565b506000998a99508998509650505050505050565b805161170981612526565b60006020828403121561211b578081fd5b815161212681612526565b9392505050565b6000602080838503121561213f578182fd5b825167ffffffffffffffff80821115612156578384fd5b818501915085601f830112612169578384fd5b815181811115612177578485fd5b83810291506121878483016124ff565b8181528481019084860184860187018a10156121a1578788fd5b8795505b838610156121cb576121b78a826120ff565b8352600195909501949186019186016121a5565b5098975050505050505050565b6000602082840312156121e9578081fd5b81516121268161253e565b60008060008060808587031215612209578283fd5b84516122148161253e565b809450506020850151925060408501519150606085015161223481612526565b939692955090935050565b600060208284031215612250578081fd5b813561212681612526565b6000806040838503121561226d578182fd5b823561227881612526565b9150602083013561228881612526565b809150509250929050565b600080600080608085870312156122a8578384fd5b84356122b381612526565b935060208501356122c381612526565b925060408501356122d381612526565b9396929550929360600135925050565b600080604083850312156122f5578182fd5b823561230081612526565b946020939093013593505050565b60008060008060808587031215612323578384fd5b843561232e81612526565b966020860135965060408601359560600135945092505050565b600080600080600060a0868803121561235f578081fd5b853561236a81612526565b97602087013597506040870135966060810135965060800135945092505050565b60006020828403121561239c578081fd5b5051919050565b6001600160a01b0391909116815260200190565b6020808252825182820181905260009190848201906040850190845b818110156123f85783516001600160a01b0316835292840192918401916001016123d3565b50909695505050505050565b6020808252825182820181905260009190848201906040850190845b818110156123f857835183529284019291840191600101612420565b606081016003851061244a57fe5b93815260208101929092526001600160a01b031660409091015290565b606081016004851061247557fe5b938152602081019290925260409091015290565b6020808252600a90820152694e4f5f4558414d504c4560b01b604082015260600190565b90815260200190565b9182526001600160a01b0316602082015260400190565b918252602082015260400190565b938452602084019290925260408301526001600160a01b0316606082015260800190565b60405181810167ffffffffffffffff8111828210171561251e57600080fd5b604052919050565b6001600160a01b038116811461253b57600080fd5b50565b801515811461253b57600080fdfea26469706673582212205c16b95a1471228a3981b6b007ee9be409e561c9e2f212fdae4012708d01ba6064736f6c634300060c0033"

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
