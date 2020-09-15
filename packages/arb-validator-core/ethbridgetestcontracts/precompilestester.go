// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethbridgetestcontracts

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

// PrecompilesTesterABI is the input ABI used to generate the binding from.
const PrecompilesTesterABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"input\",\"type\":\"uint256[25]\"}],\"name\":\"keccakF\",\"outputs\":[{\"internalType\":\"uint256[25]\",\"name\":\"\",\"type\":\"uint256[25]\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"inputChunk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"hashState\",\"type\":\"uint256\"}],\"name\":\"sha256Block\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// PrecompilesTesterFuncSigs maps the 4-byte function signature to its string representation.
var PrecompilesTesterFuncSigs = map[string]string{
	"ac90ed46": "keccakF(uint256[25])",
	"7757783d": "sha256Block(uint256[2],uint256)",
}

// PrecompilesTesterBin is the compiled bytecode used for deploying new contracts.
var PrecompilesTesterBin = "0x61155f610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c80637757783d14610045578063ac90ed46146100a4575b600080fd5b6100926004803603606081101561005b57600080fd5b6040805180820182529183019291818301918390600290839083908082843760009201919091525091945050903591506101319050565b60408051918252519081900360200190f35b6100f860048036036103208110156100bb57600080fd5b8101908080610320019060198060200260405190810160405280929190826019602002808284376000920191909152509194506101449350505050565b604051808261032080838360005b8381101561011e578181015183820152602001610106565b5050505090500191505060405180910390f35b600061013d838361015b565b9392505050565b61014c611490565b61015582610878565b92915050565b60006101656114af565b50604080516108008101825263428a2f9881526371374491602082015263b5c0fbcf9181019190915263e9b5dba56060820152633956c25b60808201526359f111f160a082015263923f82a460c082015263ab1c5ed560e082015263d807aa986101008201526312835b0161012082015263243185be61014082015263550c7dc36101608201526372be5d746101808201526380deb1fe6101a0820152639bdc06a76101c082015263c19bf1746101e082015263e49b69c161020082015263efbe4786610220820152630fc19dc661024082015263240ca1cc610260820152632de92c6f610280820152634a7484aa6102a0820152635cb0a9dc6102c08201526376f988da6102e082015263983e515261030082015263a831c66d61032082015263b00327c861034082015263bf597fc761036082015263c6e00bf361038082015263d5a791476103a08201526306ca63516103c082015263142929676103e08201526327b70a85610400820152632e1b2138610420820152634d2c6dfc6104408201526353380d1361046082015263650a735461048082015263766a0abb6104a08201526381c2c92e6104c08201526392722c856104e082015263a2bfe8a161050082015263a81a664b61052082015263c24b8b7061054082015263c76c51a361056082015263d192e81961058082015263d69906246105a082015263f40e35856105c082015263106aa0706105e08201526319a4c116610600820152631e376c08610620820152632748774c6106408201526334b0bcb561066082015263391c0cb3610680820152634ed8aa4a6106a0820152635b9cca4f6106c082015263682e6ff36106e082015263748f82ee6107008201526378a5636f6107208201526384c87814610740820152638cc702086107608201526390befffa61078082015263a4506ceb6107a082015263bef9a3f76107c082015263c67178f26107e08201526104306114af565b60005b60088163ffffffff1610156104bd5763ffffffff6020820260e003168660006020020151901c828263ffffffff166040811061046b57fe5b63ffffffff92831660209182029290920191909152820260e003168660016020020151901c828260080163ffffffff16604081106104a557fe5b63ffffffff9092166020929092020152600101610433565b5060106000805b60408363ffffffff16101561061957600384600f850363ffffffff16604081106104ea57fe5b602002015163ffffffff16901c61051b85600f860363ffffffff166040811061050f57fe5b60200201516012611469565b61053f86600f870363ffffffff166040811061053357fe5b60200201516007611469565b18189150600a846002850363ffffffff166040811061055a57fe5b602002015163ffffffff16901c61058b856002860363ffffffff166040811061057f57fe5b60200201516013611469565b6105af866002870363ffffffff16604081106105a357fe5b60200201516011611469565b1818905080846007850363ffffffff16604081106105c957fe5b602002015183866010870363ffffffff16604081106105e457fe5b6020020151010101848463ffffffff16604081106105fe57fe5b63ffffffff90921660209290920201526001909201916104c4565b6106216114ce565b600093505b60088463ffffffff161015610672578360200260e00363ffffffff1688901c818563ffffffff166008811061065757fe5b63ffffffff9092166020929092020152600190930192610626565b60008060008096505b60408763ffffffff1610156107c757608084015161069a906019611469565b60808501516106aa90600b611469565b60808601516106ba906006611469565b18189450878763ffffffff16604081106106d057fe5b6020020151898863ffffffff16604081106106e757fe5b6020020151608086015160a087015160c0880151610706929190611487565b87876007602002015101010101925061072784600060200201516016611469565b845161073490600d611469565b8551610741906002611469565b6040870180516020890180518a5160c08c01805163ffffffff90811660e08f015260a08e018051821690925260808e018051821690925260608e0180518e0182169092528086169091528083169095528481169092528083189190911691161892909118929092188181018681019093168752600199909901989750909250905061067b565b600096505b60088763ffffffff16101561081b578660200260e00363ffffffff168b901c848863ffffffff16600881106107fd57fe5b60200201805163ffffffff92019190911690526001909601956107cc565b60008097505b60088863ffffffff161015610868578760200260e00363ffffffff16858963ffffffff166008811061084f57fe5b602002015160019099019863ffffffff16901b17610821565b9c9b505050505050505050505050565b610880611490565b6108886114ed565b6108906114ed565b610898611490565b6108a061150b565b60405180610300016040528060018152602001618082815260200167800000000000808a8152602001678000000080008000815260200161808b81526020016380000001815260200167800000008000808181526020016780000000000080098152602001608a81526020016088815260200163800080098152602001638000000a8152602001638000808b815260200167800000000000008b8152602001678000000000008089815260200167800000000000800381526020016780000000000080028152602001678000000000000080815260200161800a815260200167800000008000000a81526020016780000000800080818152602001678000000000008080815260200163800000018152602001678000000080008008815250905060008090505b601881101561145e576080878101516060808a01516040808c01516020808e01518e511890911890921890931889526101208b01516101008c015160e08d015160c08e015160a08f0151181818189089018190526101c08b01516101a08c01516101808d01516101608e01516101408f0151181818189289019283526102608b01516102408c01516102208d01516102008e01516101e08f015118181818918901919091526103008a01516102e08b01516102c08c01516102a08d01516102808e0151181818189288018390526001600160401b0360028202166001603f1b91829004179092188652510485600260200201516002026001600160401b03161785600060200201511884600160200201526001603f1b856003602002015181610aec57fe5b0485600360200201516002026001600160401b03161785600160200201511884600260200201526001603f1b856004602002015181610b2757fe5b0485600460200201516002026001600160401b03161785600260058110610b4a57fe5b602002015118606085015284516001603f1b9086516060808901519390920460029091026001600160401b031617909118608086810191825286518a5118808b5287516020808d018051909218825289516040808f0180519092189091528a518e8801805190911890528a51948e0180519095189094528901805160a08e0180519091189052805160c08e0180519091189052805160e08e018051909118905280516101008e0180519091189052516101208d018051909118905291880180516101408d018051909118905280516101608d018051909118905280516101808d018051909118905280516101a08d0180519091189052516101c08c018051909118905292870180516101e08c018051909118905280516102008c018051909118905280516102208c018051909118905280516102408c0180519091189052516102608b018051909118905281516102808b018051909118905281516102a08b018051909118905281516102c08b018051909118905281516102e08b018051909118905290516103008a01805190911890529084525163100000009060208901516001600160401b03641000000000909102169190041761010084015260408701516001603d1b9060408901516001600160401b03600890910216919004176101608401526060870151628000009060608901516001600160401b036502000000000090910216919004176102608401526080870151654000000000009060808901516001600160401b036204000090910216919004176102c084015260a08701516001603f1b900487600560200201516002026001600160401b03161783600260198110610db457fe5b602002015260c0870151621000008104651000000000009091026001600160401b039081169190911760a085015260e0880151664000000000000081046104009091028216176101a08501526101008801516208000081046520000000000090910282161761020085015261012088015160048082029092166001603e1b909104176103008501526101408801516101408901516001600160401b036001603e1b90910216919004176080840152610160870151670400000000000000906101608901516001600160401b036040909102169190041760e084015261018087015162200000906101808901516001600160401b036508000000000090910216919004176101408401526101a08701516602000000000000906101a08901516001600160401b0361800090910216919004176102408401526101c08701516008906101c08901516001600160401b036001603d1b90910216919004176102a08401526101e0870151641000000000906101e08901516001600160401b03631000000090910216919004176020840152610200808801516102008901516001600160401b0366800000000000009091021691900417610120840152610220870151648000000000906102208901516001600160401b036302000000909102169190041761018084015261024087015165080000000000906102408901516001600160401b036220000090910216919004176101e0840152610260870151610100906102608901516001600160401b03600160381b90910216919004176102e0840152610280870151642000000000906102808901516001600160401b036308000000909102169190041760608401526102a087015165100000000000906102a08901516001600160401b0362100000909102169190041760c08401526102c08701516302000000906102c08901516001600160401b0364800000000090910216919004176101c08401526102e0870151600160381b906102e08901516001600160401b036101009091021691900417610220840152610300870151660400000000000090048760186020020151614000026001600160401b031617836014602002015282600a602002015183600560200201511916836000602002015118876000602002015282600b602002015183600660200201511916836001602002015118876001602002015282600c602002015183600760200201511916836002602002015118876002602002015282600d602002015183600860200201511916836003602002015118876003602002015282600e602002015183600960200201511916836004602002015118876004602002015282600f602002015183600a602002015119168360056020020151188760056020020152826010602002015183600b602002015119168360066020020151188760066020020152826011602002015183600c602002015119168360076020020151188760076020020152826012602002015183600d602002015119168360086020020151188760086020020152826013602002015183600e602002015119168360096020020151188760096020020152826014602002015183600f6020020151191683600a60200201511887600a602002015282601560200201518360106020020151191683600b60200201511887600b602002015282601660200201518360116020020151191683600c60200201511887600c602002015282601760200201518360126020020151191683600d60200201511887600d602002015282601860200201518360136020020151191683600e60200201511887600e602002015282600060200201518360146020020151191683600f60200201511887600f602002015282600160200201518360156020020151191683601060200201511887601060200201528260026020020151836016602002015119168360116020020151188760116020020152826003602002015183601760200201511916836012602002015118876012602002015282600460200201518360186020020151191683601360200201511887601360200201528260056020020151836000602002015119168360146020020151188760146020020152826006602002015183600160200201511916836015602002015118876015602002015282600760200201518360026020020151191683601660200201511887601660200201528260086020020151836003602002015119168360176020020151188760176020020152826009602002015183600460200201511916836018602002015118876018602002015281816018811061144c57fe5b602002015187511887526001016109c7565b509495945050505050565b63ffffffff9182166020829003831681901b919092169190911c1790565b82191691161890565b6040518061032001604052806019906020820280388339509192915050565b6040518061080001604052806040906020820280388339509192915050565b6040518061010001604052806008906020820280388339509192915050565b6040518060a001604052806005906020820280388339509192915050565b604051806103000160405280601890602082028038833950919291505056fea265627a7a723158208bbda1be8992125a3bbeef1204a9878f79ef921b2402cb6859493864bac4cc0f64736f6c63430005110032"

// DeployPrecompilesTester deploys a new Ethereum contract, binding an instance of PrecompilesTester to it.
func DeployPrecompilesTester(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrecompilesTester, error) {
	parsed, err := abi.JSON(strings.NewReader(PrecompilesTesterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrecompilesTesterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrecompilesTester{PrecompilesTesterCaller: PrecompilesTesterCaller{contract: contract}, PrecompilesTesterTransactor: PrecompilesTesterTransactor{contract: contract}, PrecompilesTesterFilterer: PrecompilesTesterFilterer{contract: contract}}, nil
}

// PrecompilesTester is an auto generated Go binding around an Ethereum contract.
type PrecompilesTester struct {
	PrecompilesTesterCaller     // Read-only binding to the contract
	PrecompilesTesterTransactor // Write-only binding to the contract
	PrecompilesTesterFilterer   // Log filterer for contract events
}

// PrecompilesTesterCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrecompilesTesterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompilesTesterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrecompilesTesterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompilesTesterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrecompilesTesterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrecompilesTesterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrecompilesTesterSession struct {
	Contract     *PrecompilesTester // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PrecompilesTesterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrecompilesTesterCallerSession struct {
	Contract *PrecompilesTesterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PrecompilesTesterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrecompilesTesterTransactorSession struct {
	Contract     *PrecompilesTesterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PrecompilesTesterRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrecompilesTesterRaw struct {
	Contract *PrecompilesTester // Generic contract binding to access the raw methods on
}

// PrecompilesTesterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrecompilesTesterCallerRaw struct {
	Contract *PrecompilesTesterCaller // Generic read-only contract binding to access the raw methods on
}

// PrecompilesTesterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrecompilesTesterTransactorRaw struct {
	Contract *PrecompilesTesterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrecompilesTester creates a new instance of PrecompilesTester, bound to a specific deployed contract.
func NewPrecompilesTester(address common.Address, backend bind.ContractBackend) (*PrecompilesTester, error) {
	contract, err := bindPrecompilesTester(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrecompilesTester{PrecompilesTesterCaller: PrecompilesTesterCaller{contract: contract}, PrecompilesTesterTransactor: PrecompilesTesterTransactor{contract: contract}, PrecompilesTesterFilterer: PrecompilesTesterFilterer{contract: contract}}, nil
}

// NewPrecompilesTesterCaller creates a new read-only instance of PrecompilesTester, bound to a specific deployed contract.
func NewPrecompilesTesterCaller(address common.Address, caller bind.ContractCaller) (*PrecompilesTesterCaller, error) {
	contract, err := bindPrecompilesTester(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompilesTesterCaller{contract: contract}, nil
}

// NewPrecompilesTesterTransactor creates a new write-only instance of PrecompilesTester, bound to a specific deployed contract.
func NewPrecompilesTesterTransactor(address common.Address, transactor bind.ContractTransactor) (*PrecompilesTesterTransactor, error) {
	contract, err := bindPrecompilesTester(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrecompilesTesterTransactor{contract: contract}, nil
}

// NewPrecompilesTesterFilterer creates a new log filterer instance of PrecompilesTester, bound to a specific deployed contract.
func NewPrecompilesTesterFilterer(address common.Address, filterer bind.ContractFilterer) (*PrecompilesTesterFilterer, error) {
	contract, err := bindPrecompilesTester(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrecompilesTesterFilterer{contract: contract}, nil
}

// bindPrecompilesTester binds a generic wrapper to an already deployed contract.
func bindPrecompilesTester(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrecompilesTesterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompilesTester *PrecompilesTesterRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrecompilesTester.Contract.PrecompilesTesterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompilesTester *PrecompilesTesterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompilesTester.Contract.PrecompilesTesterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompilesTester *PrecompilesTesterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompilesTester.Contract.PrecompilesTesterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrecompilesTester *PrecompilesTesterCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PrecompilesTester.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrecompilesTester *PrecompilesTesterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrecompilesTester.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrecompilesTester *PrecompilesTesterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrecompilesTester.Contract.contract.Transact(opts, method, params...)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_PrecompilesTester *PrecompilesTesterCaller) KeccakF(opts *bind.CallOpts, input [25]*big.Int) ([25]*big.Int, error) {
	var (
		ret0 = new([25]*big.Int)
	)
	out := ret0
	err := _PrecompilesTester.contract.Call(opts, out, "keccakF", input)
	return *ret0, err
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_PrecompilesTester *PrecompilesTesterSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _PrecompilesTester.Contract.KeccakF(&_PrecompilesTester.CallOpts, input)
}

// KeccakF is a free data retrieval call binding the contract method 0xac90ed46.
//
// Solidity: function keccakF(uint256[25] input) pure returns(uint256[25])
func (_PrecompilesTester *PrecompilesTesterCallerSession) KeccakF(input [25]*big.Int) ([25]*big.Int, error) {
	return _PrecompilesTester.Contract.KeccakF(&_PrecompilesTester.CallOpts, input)
}

// Sha256Block is a free data retrieval call binding the contract method 0x7757783d.
//
// Solidity: function sha256Block(uint256[2] inputChunk, uint256 hashState) pure returns(uint256)
func (_PrecompilesTester *PrecompilesTesterCaller) Sha256Block(opts *bind.CallOpts, inputChunk [2]*big.Int, hashState *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PrecompilesTester.contract.Call(opts, out, "sha256Block", inputChunk, hashState)
	return *ret0, err
}

// Sha256Block is a free data retrieval call binding the contract method 0x7757783d.
//
// Solidity: function sha256Block(uint256[2] inputChunk, uint256 hashState) pure returns(uint256)
func (_PrecompilesTester *PrecompilesTesterSession) Sha256Block(inputChunk [2]*big.Int, hashState *big.Int) (*big.Int, error) {
	return _PrecompilesTester.Contract.Sha256Block(&_PrecompilesTester.CallOpts, inputChunk, hashState)
}

// Sha256Block is a free data retrieval call binding the contract method 0x7757783d.
//
// Solidity: function sha256Block(uint256[2] inputChunk, uint256 hashState) pure returns(uint256)
func (_PrecompilesTester *PrecompilesTesterCallerSession) Sha256Block(inputChunk [2]*big.Int, hashState *big.Int) (*big.Int, error) {
	return _PrecompilesTester.Contract.Sha256Block(&_PrecompilesTester.CallOpts, inputChunk, hashState)
}
