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

// RollupCreatorABI is the input ABI used to generate the binding from.
const RollupCreatorABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rollupAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"inboxAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"adminProxy\",\"type\":\"address\"}],\"name\":\"RollupCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"TemplatesUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeCreator\",\"outputs\":[{\"internalType\":\"contractBridgeCreator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"challengeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_machineHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"_confirmPeriodBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_extraChallengeTimeBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_arbGasSpeedLimitPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_baseStake\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_stakeToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_sequencer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_sequencerDelayBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sequencerDelaySeconds\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_extraConfig\",\"type\":\"bytes\"}],\"name\":\"createRollup\",\"outputs\":[{\"internalType\":\"contractIRollup\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodeFactory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rollupTemplate\",\"outputs\":[{\"internalType\":\"contractICloneable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBridgeCreator\",\"name\":\"_bridgeCreator\",\"type\":\"address\"},{\"internalType\":\"contractICloneable\",\"name\":\"_rollupTemplate\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_challengeFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nodeFactory\",\"type\":\"address\"}],\"name\":\"setTemplates\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RollupCreatorBin is the compiled bytecode used for deploying new contracts.
var RollupCreatorBin = "0x608060405234801561001057600080fd5b5060006100246001600160e01b0361007316565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350610077565b3390565b6121a4806100866000396000f3fe60806040523480156200001157600080fd5b50600436106200008e5760003560e01c80634b1ef03014620000935780635dbaf68b1462000176578063715018a614620001805780638689d996146200018c5780638da5cb5b1462000196578063c8a7cb2114620001a0578063d93fe9c414620001e1578063f2fde38b14620001eb578063f860cefa1462000214575b600080fd5b6200015a6004803603610160811015620000ac57600080fd5b8135916020810135916040820135916060810135916080820135916001600160a01b0360a082013581169260c083013582169260e08101359092169161010081013591610120820135919081019061016081016101408201356401000000008111156200011857600080fd5b8201836020820111156200012b57600080fd5b803590602001918460018302840111640100000000831117156200014e57600080fd5b5090925090506200021e565b604080516001600160a01b039092168252519081900360200190f35b6200015a620002d6565b6200018a620002e5565b005b6200015a62000397565b6200015a620003a6565b6200018a60048036036080811015620001b857600080fd5b506001600160a01b038135811691602081013582169160408201358116916060013516620003b5565b6200015a62000496565b6200018a600480360360208110156200020357600080fd5b50356001600160a01b0316620004a5565b6200015a620005af565b6000620002c56040518061016001604052808f81526020018e81526020018d81526020018c81526020018b81526020018a6001600160a01b03168152602001896001600160a01b03168152602001886001600160a01b0316815260200187815260200186815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152505050915250620005be565b9d9c50505050505050505050505050565b6003546001600160a01b031681565b620002ef62000a60565b6001600160a01b031662000302620003a6565b6001600160a01b0316146200034d576040805162461bcd60e51b815260206004820181905260248201526000805160206200214f833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6002546001600160a01b031681565b6000546001600160a01b031690565b620003bf62000a60565b6001600160a01b0316620003d2620003a6565b6001600160a01b0316146200041d576040805162461bcd60e51b815260206004820181905260248201526000805160206200214f833981519152604482015290519081900360640190fd5b600180546001600160a01b038087166001600160a01b031992831617909255600280548684169083161790556003805485841690831617905560048054928416929091169190911790556040517fc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b90600090a150505050565b6004546001600160a01b031681565b620004af62000a60565b6001600160a01b0316620004c2620003a6565b6001600160a01b0316146200050d576040805162461bcd60e51b815260206004820181905260248201526000805160206200214f833981519152604482015290519081900360640190fd5b6001600160a01b038116620005545760405162461bcd60e51b8152600401808060200182810382526026815260200180620021296026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6001546001600160a01b031681565b6000620005ca62000a64565b604051620005d89062000aa0565b604051809103906000f080158015620005f5573d6000803e3d6000fd5b506001600160a01b03908116808352600254604051921691620006189062000aae565b6001600160a01b03928316815291166020820152606060408083018290526000918301829052519182900360a0019190f0801580156200065c573d6000803e3d6000fd5b506001600160a01b0390811660c08301819052600154835160e08701516101008801516101208901516040805163ed33557960e01b815294881660048601526024850196909652918616604484015260648301526084820152915192169163ed3355799160a48082019260a0929091908290030181600087803b158015620006e357600080fd5b505af1158015620006f8573d6000803e3d6000fd5b505050506040513d60a08110156200070f57600080fd5b5080516020808301516040808501516060808701516080978801516001600160a01b0390811660a08b015290811697890197909752908616908701529084168582015291831690840152825160c0860151825163f2fde38b60e01b81529084166004820152915192169163f2fde38b9160248082019260009290919082900301818387803b158015620007a157600080fd5b505af1158015620007b6573d6000803e3d6000fd5b505050508060c001516001600160a01b031663fdaf5797846000015185602001518660400151876060015188608001518960a001518a60c001518b61014001516040518060c001604052808c602001516001600160a01b03166001600160a01b031681526020018c604001516001600160a01b03166001600160a01b031681526020018c60a001516001600160a01b03166001600160a01b031681526020018c608001516001600160a01b03166001600160a01b03168152602001600360009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152602001600460009054906101000a90046001600160a01b03166001600160a01b03166001600160a01b03168152506040518a63ffffffff1660e01b8152600401808a8152602001898152602001888152602001878152602001868152602001856001600160a01b03166001600160a01b03168152602001846001600160a01b03166001600160a01b031681526020018060200183600660200280838360005b83811015620009545781810151838201526020016200093a565b50505050905001828103825284818151815260200191508051906020019080838360005b838110156200099257818101518382015260200162000978565b50505050905090810190601f168015620009c05780820380516001836020036101000a031916815260200191505b509a5050505050505050505050600060405180830381600087803b158015620009e857600080fd5b505af1158015620009fd573d6000803e3d6000fd5b50505060c082015160608301518351604080516001600160a01b039384168152918316602083015280519290931693507fd508a734b33000eb18068aa34f20f8014fa578d682a9d355017efcd93e1b4f1092908290030190a260c0015192915050565b3390565b6040805160e081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c081019190915290565b61090c8062000abd83390190565b610d6080620013c98339019056fe608060405234801561001057600080fd5b5060006100246001600160e01b0361007316565b600080546001600160a01b0319166001600160a01b0383169081178255604051929350917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350610077565b3390565b610886806100866000396000f3fe60806040526004361061006b5760003560e01c8063204e1c7a14610070578063715018a6146100bf5780637eff275e146100d65780638da5cb5b146101115780639623609d1461012657806399a88ec4146101e5578063f2fde38b14610220578063f3b7dead14610253575b600080fd5b34801561007c57600080fd5b506100a36004803603602081101561009357600080fd5b50356001600160a01b0316610286565b604080516001600160a01b039092168252519081900360200190f35b3480156100cb57600080fd5b506100d4610318565b005b3480156100e257600080fd5b506100d4600480360360408110156100f957600080fd5b506001600160a01b03813581169160200135166103c4565b34801561011d57600080fd5b506100a361049a565b6100d46004803603606081101561013c57600080fd5b6001600160a01b03823581169260208101359091169181019060608101604082013564010000000081111561017057600080fd5b82018360208201111561018257600080fd5b803590602001918460018302840111640100000000831117156101a457600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506104a9945050505050565b3480156101f157600080fd5b506100d46004803603604081101561020857600080fd5b506001600160a01b03813581169160200135166105eb565b34801561022c57600080fd5b506100d46004803603602081101561024357600080fd5b50356001600160a01b03166106a5565b34801561025f57600080fd5b506100a36004803603602081101561027657600080fd5b50356001600160a01b03166107a7565b6000806060836001600160a01b03166040518080635c60da1b60e01b8152506004019050600060405180830381855afa9150503d80600081146102e5576040519150601f19603f3d011682016040523d82523d6000602084013e6102ea565b606091505b5091509150816102f957600080fd5b80806020019051602081101561030e57600080fd5b5051949350505050565b610320610806565b6001600160a01b031661033161049a565b6001600160a01b03161461037a576040805162461bcd60e51b81526020600482018190526024820152600080516020610831833981519152604482015290519081900360640190fd5b600080546040516001600160a01b03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3600080546001600160a01b0319169055565b6103cc610806565b6001600160a01b03166103dd61049a565b6001600160a01b031614610426576040805162461bcd60e51b81526020600482018190526024820152600080516020610831833981519152604482015290519081900360640190fd5b816001600160a01b0316638f283970826040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050600060405180830381600087803b15801561047e57600080fd5b505af1158015610492573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031690565b6104b1610806565b6001600160a01b03166104c261049a565b6001600160a01b03161461050b576040805162461bcd60e51b81526020600482018190526024820152600080516020610831833981519152604482015290519081900360640190fd5b826001600160a01b0316634f1ef2863484846040518463ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610581578181015183820152602001610569565b50505050905090810190601f1680156105ae5780820380516001836020036101000a031916815260200191505b5093505050506000604051808303818588803b1580156105cd57600080fd5b505af11580156105e1573d6000803e3d6000fd5b5050505050505050565b6105f3610806565b6001600160a01b031661060461049a565b6001600160a01b03161461064d576040805162461bcd60e51b81526020600482018190526024820152600080516020610831833981519152604482015290519081900360640190fd5b816001600160a01b0316633659cfe6826040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b03168152602001915050600060405180830381600087803b15801561047e57600080fd5b6106ad610806565b6001600160a01b03166106be61049a565b6001600160a01b031614610707576040805162461bcd60e51b81526020600482018190526024820152600080516020610831833981519152604482015290519081900360640190fd5b6001600160a01b03811661074c5760405162461bcd60e51b815260040180806020018281038252602681526020018061080b6026913960400191505060405180910390fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806060836001600160a01b031660405180806303e1469160e61b8152506004019050600060405180830381855afa9150503d80600081146102e5576040519150601f19603f3d011682016040523d82523d6000602084013e6102ea565b339056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a264697066735822122005019646333ea388b0c67bda202f0b0f9e694623274a3b0b0cd706bdc2821b5464736f6c634300060b0033608060405260405162000d6038038062000d60833981810160405260608110156200002957600080fd5b815160208301516040808501805191519395929483019291846401000000008211156200005557600080fd5b9083019060208201858111156200006b57600080fd5b82516401000000008111828201881017156200008657600080fd5b82525081516020918201929091019080838360005b83811015620000b55781810151838201526020016200009b565b50505050905090810190601f168015620000e35780820380516001836020036101000a031916815260200191505b5060408181527f656970313936372e70726f78792e696d706c656d656e746174696f6e0000000082525190819003601c01902086935084925060008051602062000cbd8339815191526000199091011490506200013c57fe5b62000150826001600160e01b03620001e016565b80511562000171576200016f82826200024660201b620003841760201c565b505b5050604080517f656970313936372e70726f78792e61646d696e000000000000000000000000008152905190819003601301902060008051602062000c9d83398151915260001990910114620001c357fe5b620001d7826001600160e01b036200027e16565b50505062000461565b620001f6816200029160201b620003b01760201c565b620002335760405162461bcd60e51b815260040180806020018281038252603681526020018062000d046036913960400191505060405180910390fd5b60008051602062000cbd83398151915255565b606062000277838360405180606001604052806027815260200162000cdd602791396001600160e01b036200029716565b9392505050565b60008051602062000c9d83398151915255565b3b151590565b6060620002ad846001600160e01b036200029116565b620002ea5760405162461bcd60e51b815260040180806020018281038252602681526020018062000d3a6026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b602083106200032a5780518252601f19909201916020918201910162000309565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146200038c576040519150601f19603f3d011682016040523d82523d6000602084013e62000391565b606091505b509092509050620003ad8282866001600160e01b03620003b716565b9695505050505050565b60608315620003c857508162000277565b825115620003d95782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015620004255781810151838201526020016200040b565b50505050905090810190601f168015620004535780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b61082c80620004716000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100985780635c60da1b146101185780638f28397014610149578063f851a4401461017c5761005d565b3661005d5761005b610191565b005b61005b610191565b34801561007157600080fd5b5061005b6004803603602081101561008857600080fd5b50356001600160a01b03166101ab565b61005b600480360360408110156100ae57600080fd5b6001600160a01b0382351691908101906040810160208201356401000000008111156100d957600080fd5b8201836020820111156100eb57600080fd5b8035906020019184600183028401116401000000008311171561010d57600080fd5b5090925090506101e5565b34801561012457600080fd5b5061012d610262565b604080516001600160a01b039092168252519081900360200190f35b34801561015557600080fd5b5061005b6004803603602081101561016c57600080fd5b50356001600160a01b031661029f565b34801561018857600080fd5b5061012d610359565b6101996103b6565b6101a96101a4610416565b61043b565b565b6101b361045f565b6001600160a01b0316336001600160a01b031614156101da576101d581610484565b6101e2565b6101e2610191565b50565b6101ed61045f565b6001600160a01b0316336001600160a01b031614156102555761020f83610484565b61024f8383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061038492505050565b5061025d565b61025d610191565b505050565b600061026c61045f565b6001600160a01b0316336001600160a01b031614156102945761028d610416565b905061029c565b61029c610191565b90565b6102a761045f565b6001600160a01b0316336001600160a01b031614156101da576001600160a01b0381166103055760405162461bcd60e51b815260040180806020018281038252603a8152602001806106f8603a913960400191505060405180910390fd5b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61032e61045f565b604080516001600160a01b03928316815291841660208301528051918290030190a16101d5816104c4565b600061036361045f565b6001600160a01b0316336001600160a01b031614156102945761028d61045f565b60606103a98383604051806060016040528060278152602001610732602791396104e8565b9392505050565b3b151590565b6103be61045f565b6001600160a01b0316336001600160a01b0316141561040e5760405162461bcd60e51b81526004018080602001828103825260428152602001806107b56042913960600191505060405180910390fd5b6101a96101a9565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc5490565b3660008037600080366000845af43d6000803e80801561045a573d6000f35b3d6000fd5b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d61035490565b61048d816105eb565b6040516001600160a01b038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b7fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d610355565b60606104f3846103b0565b61052e5760405162461bcd60e51b815260040180806020018281038252602681526020018061078f6026913960400191505060405180910390fd5b60006060856001600160a01b0316856040518082805190602001908083835b6020831061056c5780518252601f19909201916020918201910161054d565b6001836020036101000a038019825116818451168082178552505050505050905001915050600060405180830381855af49150503d80600081146105cc576040519150601f19603f3d011682016040523d82523d6000602084013e6105d1565b606091505b50915091506105e1828286610653565b9695505050505050565b6105f4816103b0565b61062f5760405162461bcd60e51b81526004018080602001828103825260368152602001806107596036913960400191505060405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc55565b606083156106625750816103a9565b8251156106725782518084602001fd5b8160405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b838110156106bc5781810151838201526020016106a4565b50505050905090810190601f1680156106e95780820380516001836020036101000a031916815260200191505b509250505060405180910390fdfe5472616e73706172656e745570677261646561626c6550726f78793a206e65772061646d696e20697320746865207a65726f2061646472657373416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163745472616e73706172656e745570677261646561626c6550726f78793a2061646d696e2063616e6e6f742066616c6c6261636b20746f2070726f787920746172676574a2646970667358221220175110956fa0a7ff1615f55e1422acff6edcec0099d7ea0bae101f4f6228c8bd64736f6c634300060b0033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c65645570677261646561626c6550726f78793a206e657720696d706c656d656e746174696f6e206973206e6f74206120636f6e7472616374416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6e74726163744f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572a264697066735822122031d6570d0556206091943405ebac4088cbe1f6e42de0f5b185bdc3a04b1fc0ac64736f6c634300060b0033"

// DeployRollupCreator deploys a new Ethereum contract, binding an instance of RollupCreator to it.
func DeployRollupCreator(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RollupCreator, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RollupCreatorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RollupCreator{RollupCreatorCaller: RollupCreatorCaller{contract: contract}, RollupCreatorTransactor: RollupCreatorTransactor{contract: contract}, RollupCreatorFilterer: RollupCreatorFilterer{contract: contract}}, nil
}

// RollupCreator is an auto generated Go binding around an Ethereum contract.
type RollupCreator struct {
	RollupCreatorCaller     // Read-only binding to the contract
	RollupCreatorTransactor // Write-only binding to the contract
	RollupCreatorFilterer   // Log filterer for contract events
}

// RollupCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupCreatorSession struct {
	Contract     *RollupCreator    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCreatorCallerSession struct {
	Contract *RollupCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RollupCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupCreatorTransactorSession struct {
	Contract     *RollupCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RollupCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupCreatorRaw struct {
	Contract *RollupCreator // Generic contract binding to access the raw methods on
}

// RollupCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCreatorCallerRaw struct {
	Contract *RollupCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// RollupCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupCreatorTransactorRaw struct {
	Contract *RollupCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollupCreator creates a new instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreator(address common.Address, backend bind.ContractBackend) (*RollupCreator, error) {
	contract, err := bindRollupCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RollupCreator{RollupCreatorCaller: RollupCreatorCaller{contract: contract}, RollupCreatorTransactor: RollupCreatorTransactor{contract: contract}, RollupCreatorFilterer: RollupCreatorFilterer{contract: contract}}, nil
}

// NewRollupCreatorCaller creates a new read-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorCaller(address common.Address, caller bind.ContractCaller) (*RollupCreatorCaller, error) {
	contract, err := bindRollupCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorCaller{contract: contract}, nil
}

// NewRollupCreatorTransactor creates a new write-only instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupCreatorTransactor, error) {
	contract, err := bindRollupCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorTransactor{contract: contract}, nil
}

// NewRollupCreatorFilterer creates a new log filterer instance of RollupCreator, bound to a specific deployed contract.
func NewRollupCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupCreatorFilterer, error) {
	contract, err := bindRollupCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorFilterer{contract: contract}, nil
}

// bindRollupCreator binds a generic wrapper to an already deployed contract.
func bindRollupCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.RollupCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.RollupCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RollupCreator *RollupCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RollupCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RollupCreator *RollupCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RollupCreator *RollupCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RollupCreator.Contract.contract.Transact(opts, method, params...)
}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreator *RollupCreatorCaller) BridgeCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "bridgeCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreator *RollupCreatorSession) BridgeCreator() (common.Address, error) {
	return _RollupCreator.Contract.BridgeCreator(&_RollupCreator.CallOpts)
}

// BridgeCreator is a free data retrieval call binding the contract method 0xf860cefa.
//
// Solidity: function bridgeCreator() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) BridgeCreator() (common.Address, error) {
	return _RollupCreator.Contract.BridgeCreator(&_RollupCreator.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupCreator *RollupCreatorCaller) ChallengeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "challengeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupCreator *RollupCreatorSession) ChallengeFactory() (common.Address, error) {
	return _RollupCreator.Contract.ChallengeFactory(&_RollupCreator.CallOpts)
}

// ChallengeFactory is a free data retrieval call binding the contract method 0x5dbaf68b.
//
// Solidity: function challengeFactory() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) ChallengeFactory() (common.Address, error) {
	return _RollupCreator.Contract.ChallengeFactory(&_RollupCreator.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupCreator *RollupCreatorCaller) NodeFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "nodeFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupCreator *RollupCreatorSession) NodeFactory() (common.Address, error) {
	return _RollupCreator.Contract.NodeFactory(&_RollupCreator.CallOpts)
}

// NodeFactory is a free data retrieval call binding the contract method 0xd93fe9c4.
//
// Solidity: function nodeFactory() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) NodeFactory() (common.Address, error) {
	return _RollupCreator.Contract.NodeFactory(&_RollupCreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreator *RollupCreatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreator *RollupCreatorSession) Owner() (common.Address, error) {
	return _RollupCreator.Contract.Owner(&_RollupCreator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) Owner() (common.Address, error) {
	return _RollupCreator.Contract.Owner(&_RollupCreator.CallOpts)
}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_RollupCreator *RollupCreatorCaller) RollupTemplate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RollupCreator.contract.Call(opts, &out, "rollupTemplate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_RollupCreator *RollupCreatorSession) RollupTemplate() (common.Address, error) {
	return _RollupCreator.Contract.RollupTemplate(&_RollupCreator.CallOpts)
}

// RollupTemplate is a free data retrieval call binding the contract method 0x8689d996.
//
// Solidity: function rollupTemplate() view returns(address)
func (_RollupCreator *RollupCreatorCallerSession) RollupTemplate() (common.Address, error) {
	return _RollupCreator.Contract.RollupTemplate(&_RollupCreator.CallOpts)
}

// CreateRollup is a paid mutator transaction binding the contract method 0x4b1ef030.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorTransactor) CreateRollup(opts *bind.TransactOpts, _machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _sequencer common.Address, _sequencerDelayBlocks *big.Int, _sequencerDelaySeconds *big.Int, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "createRollup", _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _sequencer, _sequencerDelayBlocks, _sequencerDelaySeconds, _extraConfig)
}

// CreateRollup is a paid mutator transaction binding the contract method 0x4b1ef030.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorSession) CreateRollup(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _sequencer common.Address, _sequencerDelayBlocks *big.Int, _sequencerDelaySeconds *big.Int, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _sequencer, _sequencerDelayBlocks, _sequencerDelaySeconds, _extraConfig)
}

// CreateRollup is a paid mutator transaction binding the contract method 0x4b1ef030.
//
// Solidity: function createRollup(bytes32 _machineHash, uint256 _confirmPeriodBlocks, uint256 _extraChallengeTimeBlocks, uint256 _arbGasSpeedLimitPerBlock, uint256 _baseStake, address _stakeToken, address _owner, address _sequencer, uint256 _sequencerDelayBlocks, uint256 _sequencerDelaySeconds, bytes _extraConfig) returns(address)
func (_RollupCreator *RollupCreatorTransactorSession) CreateRollup(_machineHash [32]byte, _confirmPeriodBlocks *big.Int, _extraChallengeTimeBlocks *big.Int, _arbGasSpeedLimitPerBlock *big.Int, _baseStake *big.Int, _stakeToken common.Address, _owner common.Address, _sequencer common.Address, _sequencerDelayBlocks *big.Int, _sequencerDelaySeconds *big.Int, _extraConfig []byte) (*types.Transaction, error) {
	return _RollupCreator.Contract.CreateRollup(&_RollupCreator.TransactOpts, _machineHash, _confirmPeriodBlocks, _extraChallengeTimeBlocks, _arbGasSpeedLimitPerBlock, _baseStake, _stakeToken, _owner, _sequencer, _sequencerDelayBlocks, _sequencerDelaySeconds, _extraConfig)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreator *RollupCreatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreator *RollupCreatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreator.Contract.RenounceOwnership(&_RollupCreator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RollupCreator *RollupCreatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RollupCreator.Contract.RenounceOwnership(&_RollupCreator.TransactOpts)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xc8a7cb21.
//
// Solidity: function setTemplates(address _bridgeCreator, address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreator *RollupCreatorTransactor) SetTemplates(opts *bind.TransactOpts, _bridgeCreator common.Address, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "setTemplates", _bridgeCreator, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xc8a7cb21.
//
// Solidity: function setTemplates(address _bridgeCreator, address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreator *RollupCreatorSession) SetTemplates(_bridgeCreator common.Address, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.SetTemplates(&_RollupCreator.TransactOpts, _bridgeCreator, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// SetTemplates is a paid mutator transaction binding the contract method 0xc8a7cb21.
//
// Solidity: function setTemplates(address _bridgeCreator, address _rollupTemplate, address _challengeFactory, address _nodeFactory) returns()
func (_RollupCreator *RollupCreatorTransactorSession) SetTemplates(_bridgeCreator common.Address, _rollupTemplate common.Address, _challengeFactory common.Address, _nodeFactory common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.SetTemplates(&_RollupCreator.TransactOpts, _bridgeCreator, _rollupTemplate, _challengeFactory, _nodeFactory)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreator *RollupCreatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreator *RollupCreatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.TransferOwnership(&_RollupCreator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RollupCreator *RollupCreatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RollupCreator.Contract.TransferOwnership(&_RollupCreator.TransactOpts, newOwner)
}

// RollupCreatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RollupCreator contract.
type RollupCreatorOwnershipTransferredIterator struct {
	Event *RollupCreatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RollupCreatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorOwnershipTransferred)
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
		it.Event = new(RollupCreatorOwnershipTransferred)
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
func (it *RollupCreatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorOwnershipTransferred represents a OwnershipTransferred event raised by the RollupCreator contract.
type RollupCreatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreator *RollupCreatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RollupCreatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorOwnershipTransferredIterator{contract: _RollupCreator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreator *RollupCreatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RollupCreatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorOwnershipTransferred)
				if err := _RollupCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RollupCreator *RollupCreatorFilterer) ParseOwnershipTransferred(log types.Log) (*RollupCreatorOwnershipTransferred, error) {
	event := new(RollupCreatorOwnershipTransferred)
	if err := _RollupCreator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCreatorRollupCreatedIterator is returned from FilterRollupCreated and is used to iterate over the raw logs and unpacked data for RollupCreated events raised by the RollupCreator contract.
type RollupCreatorRollupCreatedIterator struct {
	Event *RollupCreatorRollupCreated // Event containing the contract specifics and raw log

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
func (it *RollupCreatorRollupCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorRollupCreated)
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
		it.Event = new(RollupCreatorRollupCreated)
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
func (it *RollupCreatorRollupCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorRollupCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorRollupCreated represents a RollupCreated event raised by the RollupCreator contract.
type RollupCreatorRollupCreated struct {
	RollupAddress common.Address
	InboxAddress  common.Address
	AdminProxy    common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRollupCreated is a free log retrieval operation binding the contract event 0xd508a734b33000eb18068aa34f20f8014fa578d682a9d355017efcd93e1b4f10.
//
// Solidity: event RollupCreated(address indexed rollupAddress, address inboxAddress, address adminProxy)
func (_RollupCreator *RollupCreatorFilterer) FilterRollupCreated(opts *bind.FilterOpts, rollupAddress []common.Address) (*RollupCreatorRollupCreatedIterator, error) {

	var rollupAddressRule []interface{}
	for _, rollupAddressItem := range rollupAddress {
		rollupAddressRule = append(rollupAddressRule, rollupAddressItem)
	}

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "RollupCreated", rollupAddressRule)
	if err != nil {
		return nil, err
	}
	return &RollupCreatorRollupCreatedIterator{contract: _RollupCreator.contract, event: "RollupCreated", logs: logs, sub: sub}, nil
}

// WatchRollupCreated is a free log subscription operation binding the contract event 0xd508a734b33000eb18068aa34f20f8014fa578d682a9d355017efcd93e1b4f10.
//
// Solidity: event RollupCreated(address indexed rollupAddress, address inboxAddress, address adminProxy)
func (_RollupCreator *RollupCreatorFilterer) WatchRollupCreated(opts *bind.WatchOpts, sink chan<- *RollupCreatorRollupCreated, rollupAddress []common.Address) (event.Subscription, error) {

	var rollupAddressRule []interface{}
	for _, rollupAddressItem := range rollupAddress {
		rollupAddressRule = append(rollupAddressRule, rollupAddressItem)
	}

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "RollupCreated", rollupAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorRollupCreated)
				if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
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

// ParseRollupCreated is a log parse operation binding the contract event 0xd508a734b33000eb18068aa34f20f8014fa578d682a9d355017efcd93e1b4f10.
//
// Solidity: event RollupCreated(address indexed rollupAddress, address inboxAddress, address adminProxy)
func (_RollupCreator *RollupCreatorFilterer) ParseRollupCreated(log types.Log) (*RollupCreatorRollupCreated, error) {
	event := new(RollupCreatorRollupCreated)
	if err := _RollupCreator.contract.UnpackLog(event, "RollupCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupCreatorTemplatesUpdatedIterator is returned from FilterTemplatesUpdated and is used to iterate over the raw logs and unpacked data for TemplatesUpdated events raised by the RollupCreator contract.
type RollupCreatorTemplatesUpdatedIterator struct {
	Event *RollupCreatorTemplatesUpdated // Event containing the contract specifics and raw log

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
func (it *RollupCreatorTemplatesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupCreatorTemplatesUpdated)
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
		it.Event = new(RollupCreatorTemplatesUpdated)
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
func (it *RollupCreatorTemplatesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupCreatorTemplatesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupCreatorTemplatesUpdated represents a TemplatesUpdated event raised by the RollupCreator contract.
type RollupCreatorTemplatesUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTemplatesUpdated is a free log retrieval operation binding the contract event 0xc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b.
//
// Solidity: event TemplatesUpdated()
func (_RollupCreator *RollupCreatorFilterer) FilterTemplatesUpdated(opts *bind.FilterOpts) (*RollupCreatorTemplatesUpdatedIterator, error) {

	logs, sub, err := _RollupCreator.contract.FilterLogs(opts, "TemplatesUpdated")
	if err != nil {
		return nil, err
	}
	return &RollupCreatorTemplatesUpdatedIterator{contract: _RollupCreator.contract, event: "TemplatesUpdated", logs: logs, sub: sub}, nil
}

// WatchTemplatesUpdated is a free log subscription operation binding the contract event 0xc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b.
//
// Solidity: event TemplatesUpdated()
func (_RollupCreator *RollupCreatorFilterer) WatchTemplatesUpdated(opts *bind.WatchOpts, sink chan<- *RollupCreatorTemplatesUpdated) (event.Subscription, error) {

	logs, sub, err := _RollupCreator.contract.WatchLogs(opts, "TemplatesUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupCreatorTemplatesUpdated)
				if err := _RollupCreator.contract.UnpackLog(event, "TemplatesUpdated", log); err != nil {
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

// ParseTemplatesUpdated is a log parse operation binding the contract event 0xc9d3947d22fa124aaec4c7e8c919f79016e2d7b48eee10568375d98b86460d1b.
//
// Solidity: event TemplatesUpdated()
func (_RollupCreator *RollupCreatorFilterer) ParseTemplatesUpdated(log types.Log) (*RollupCreatorTemplatesUpdated, error) {
	event := new(RollupCreatorTemplatesUpdated)
	if err := _RollupCreator.contract.UnpackLog(event, "TemplatesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
