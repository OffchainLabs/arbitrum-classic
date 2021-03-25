/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Signer } from 'ethers'
import { Provider, TransactionRequest } from '@ethersproject/providers'
import { Contract, ContractFactory, Overrides } from '@ethersproject/contracts'

import type { ArbTokenBridge } from '../ArbTokenBridge'

export class ArbTokenBridge__factory extends ContractFactory {
  constructor(signer?: Signer) {
    super(_abi, _bytecode, signer)
  }

  deploy(
    _l1Pair: string,
    _templateERC777: string,
    _templateERC20: string,
    overrides?: Overrides
  ): Promise<ArbTokenBridge> {
    return super.deploy(
      _l1Pair,
      _templateERC777,
      _templateERC20,
      overrides || {}
    ) as Promise<ArbTokenBridge>
  }
  getDeployTransaction(
    _l1Pair: string,
    _templateERC777: string,
    _templateERC20: string,
    overrides?: Overrides
  ): TransactionRequest {
    return super.getDeployTransaction(
      _l1Pair,
      _templateERC777,
      _templateERC20,
      overrides || {}
    )
  }
  attach(address: string): ArbTokenBridge {
    return super.attach(address) as ArbTokenBridge
  }
  connect(signer: Signer): ArbTokenBridge__factory {
    return super.connect(signer) as ArbTokenBridge__factory
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): ArbTokenBridge {
    return new Contract(address, _abi, signerOrProvider) as ArbTokenBridge
  }
}

const _abi = [
  {
    inputs: [
      {
        internalType: 'address',
        name: '_l1Pair',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_templateERC777',
        type: 'address',
      },
      {
        internalType: 'address',
        name: '_templateERC20',
        type: 'address',
      },
    ],
    stateMutability: 'nonpayable',
    type: 'constructor',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'bool',
        name: 'success',
        type: 'bool',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'sender',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'dest',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'MintAndCallTriggered',
    type: 'event',
  },
  {
    anonymous: false,
    inputs: [
      {
        indexed: false,
        internalType: 'uint256',
        name: 'id',
        type: 'uint256',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'l1Address',
        type: 'address',
      },
      {
        indexed: true,
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
      {
        indexed: true,
        internalType: 'address',
        name: 'destination',
        type: 'address',
      },
      {
        indexed: false,
        internalType: 'uint256',
        name: 'exitNum',
        type: 'uint256',
      },
    ],
    name: 'WithdrawToken',
    type: 'event',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
    ],
    name: 'calculateBridgedERC20Address',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
    ],
    name: 'calculateBridgedERC777Address',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    name: 'customToken',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1Address',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'l2Address',
        type: 'address',
      },
    ],
    name: 'customTokenRegistered',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'l1Pair',
    outputs: [
      {
        internalType: 'address',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'target',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'account',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: 'data',
        type: 'bytes',
      },
    ],
    name: 'migrate',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'contract IArbToken',
        name: 'token',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
      {
        internalType: 'address',
        name: 'sender',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'dest',
        type: 'address',
      },
      {
        internalType: 'bytes',
        name: 'data',
        type: 'bytes',
      },
    ],
    name: 'mintAndCall',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'sender',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'dest',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
      {
        internalType: 'bytes',
        name: 'callHookData',
        type: 'bytes',
      },
    ],
    name: 'mintCustomTokenFromL1',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'sender',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'dest',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
      {
        internalType: 'uint8',
        name: 'decimals',
        type: 'uint8',
      },
      {
        internalType: 'bytes',
        name: 'callHookData',
        type: 'bytes',
      },
    ],
    name: 'mintERC20FromL1',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'sender',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'dest',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
      {
        internalType: 'uint8',
        name: 'decimals',
        type: 'uint8',
      },
      {
        internalType: 'bytes',
        name: 'callHookData',
        type: 'bytes',
      },
    ],
    name: 'mintERC777FromL1',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [],
    name: 'templateERC20',
    outputs: [
      {
        internalType: 'contract ICloneable',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [],
    name: 'templateERC777',
    outputs: [
      {
        internalType: 'contract ICloneable',
        name: '',
        type: 'address',
      },
    ],
    stateMutability: 'view',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
      {
        internalType: 'bytes',
        name: '_name',
        type: 'bytes',
      },
      {
        internalType: 'bytes',
        name: '_symbol',
        type: 'bytes',
      },
      {
        internalType: 'bytes',
        name: '_decimals',
        type: 'bytes',
      },
    ],
    name: 'updateERC20TokenInfo',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
      {
        internalType: 'bytes',
        name: '_name',
        type: 'bytes',
      },
      {
        internalType: 'bytes',
        name: '_symbol',
        type: 'bytes',
      },
      {
        internalType: 'bytes',
        name: '_decimals',
        type: 'bytes',
      },
    ],
    name: 'updateERC777TokenInfo',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
  {
    inputs: [
      {
        internalType: 'address',
        name: 'l1ERC20',
        type: 'address',
      },
      {
        internalType: 'address',
        name: 'destination',
        type: 'address',
      },
      {
        internalType: 'uint256',
        name: 'amount',
        type: 'uint256',
      },
    ],
    name: 'withdraw',
    outputs: [],
    stateMutability: 'nonpayable',
    type: 'function',
  },
]

const _bytecode =
  '0x60e060405234801561001057600080fd5b506040516120a73803806120a78339818101604052606081101561003357600080fd5b50805160208201516040909201519091906001600160a01b03831661009f576040805162461bcd60e51b815260206004820152601a60248201527f4c3120706169722063616e277420626520616464726573732030000000000000604482015290519081900360640190fd5b6001600160601b0319606082811b821660805283811b821660a05284901b1660c0526001600160a01b03908116929181169116611f7b61012c60003980610794528061091a5280610ba65280610bd35280610d2c5280610fe7528061119052806114e35250806108d9528061163b5280611a52525080610cf35280610fba5280611a785250611f7b6000f3fe608060405234801561001057600080fd5b50600436106100d45760003560e01c80635f292aaa116100875780635f292aaa146103e957806386038bd11461040f578063998ea6261461052d5780639de96c5c14610535578063d9caed1214610563578063db70cf6e14610599578063e0aaf02f14610663578063fa09eb4514610781576100d4565b806203d885146100d9578063123170aa1461017957806313c5cd0f146101bb57806329fd79e914610259578063429866e61461027f5780634b805f481461034b57806353ab535214610353575b600080fd5b610177600480360360c08110156100ef57600080fd5b6001600160a01b038235811692602081013582169260408201359092169160608201359160ff6080820135169181019060c0810160a0820135600160201b81111561013957600080fd5b82018360208201111561014b57600080fd5b803590602001918460018302840111600160201b8311171561016c57600080fd5b509092509050610789565b005b61019f6004803603602081101561018f57600080fd5b50356001600160a01b03166108d2565b604080516001600160a01b039092168252519081900360200190f35b610177600480360360c08110156101d157600080fd5b6001600160a01b038235811692602081013582169260408201359092169160608201359160ff6080820135169181019060c0810160a0820135600160201b81111561021b57600080fd5b82018360208201111561022d57600080fd5b803590602001918460018302840111600160201b8311171561024e57600080fd5b50909250905061090f565b61019f6004803603602081101561026f57600080fd5b50356001600160a01b0316610988565b610177600480360360a081101561029557600080fd5b6001600160a01b03823581169260208101358216926040820135909216916060820135919081019060a081016080820135600160201b8111156102d757600080fd5b8201836020820111156102e957600080fd5b803590602001918460018302840111600160201b8311171561030a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506109a3945050505050565b61019f610ba4565b610177600480360360a081101561036957600080fd5b6001600160a01b03823581169260208101358216926040820135909216916060820135919081019060a081016080820135600160201b8111156103ab57600080fd5b8201836020820111156103bd57600080fd5b803590602001918460018302840111600160201b831117156103de57600080fd5b509092509050610bc8565b61019f600480360360208110156103ff57600080fd5b50356001600160a01b0316610cec565b6101776004803603608081101561042557600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561044f57600080fd5b82018360208201111561046157600080fd5b803590602001918460018302840111600160201b8311171561048257600080fd5b919390929091602081019035600160201b81111561049f57600080fd5b8201836020820111156104b157600080fd5b803590602001918460018302840111600160201b831117156104d257600080fd5b919390929091602081019035600160201b8111156104ef57600080fd5b82018360208201111561050157600080fd5b803590602001918460018302840111600160201b8311171561052257600080fd5b509092509050610d21565b61019f610fb8565b6101776004803603604081101561054b57600080fd5b506001600160a01b0381358116916020013516610fdc565b6101776004803603606081101561057957600080fd5b506001600160a01b03813581169160208101359091169060400135611077565b610177600480360360a08110156105af57600080fd5b6001600160a01b0382358116926020810135926040820135831692606083013516919081019060a081016080820135600160201b8111156105ef57600080fd5b82018360208201111561060157600080fd5b803590602001918460018302840111600160201b8311171561062257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506112d1945050505050565b6101776004803603608081101561067957600080fd5b6001600160a01b038235169190810190604081016020820135600160201b8111156106a357600080fd5b8201836020820111156106b557600080fd5b803590602001918460018302840111600160201b831117156106d657600080fd5b919390929091602081019035600160201b8111156106f357600080fd5b82018360208201111561070557600080fd5b803590602001918460018302840111600160201b8311171561072657600080fd5b919390929091602081019035600160201b81111561074357600080fd5b82018360208201111561075557600080fd5b803590602001918460018302840111600160201b8311171561077657600080fd5b5090925090506114d8565b61019f611639565b326001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146107f6576040805162461bcd60e51b815260206004820152600d60248201526c27a7262cafa2aa242fa820a4a960991b604482015290519081900360640190fd5b6000610802888561165d565b905081156108525761084d8186898987878080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061167492505050565b6108c8565b604080516332dbacfd60e21b81526001600160a01b038881166004830152602482018890526060604483015260006064830181905292519084169263cb6eb3f49260a4808201939182900301818387803b1580156108af57600080fd5b505af11580156108c3573d6000803e3d6000fd5b505050505b5050505050505050565b60006109077f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038416611877565b90505b919050565b326001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461097c576040805162461bcd60e51b815260206004820152600d60248201526c27a7262cafa2aa242fa820a4a960991b604482015290519081900360640190fd5b60006108028885611893565b6000602081905290815260409020546001600160a01b031681565b846109ad816108d2565b6001600160a01b0316336001600160a01b031614806109e557506109d081610cec565b6001600160a01b0316336001600160a01b0316145b610a30576040805162461bcd60e51b81526020600482015260176024820152762727aa2fa32927a6afa9aa20a72220a9222faa27a5a2a760491b604482015290519081900360640190fd5b8585610a3b826108d2565b6001600160a01b0316816001600160a01b03161480610a735750610a5e82610cec565b6001600160a01b0316816001600160a01b0316145b80610a9a57506001600160a01b038083166000908152602081905260409020548282169116145b610ada576040805162461bcd60e51b815260206004820152600c60248201526b2727aa2faa27afaa27a5a2a760a11b604482015290519081900360640190fd5b866001600160a01b031663cb6eb3f48787876040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610b56578181015183820152602001610b3e565b50505050905090810190601f168015610b835780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b1580156108af57600080fd5b7f000000000000000000000000000000000000000000000000000000000000000081565b326001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610c35576040805162461bcd60e51b815260206004820152600d60248201526c27a7262cafa2aa242fa820a4a960991b604482015290519081900360640190fd5b6001600160a01b038087166000908152602081905260409020541680610ca2576040805162461bcd60e51b815260206004820152601a60248201527f437573746f6d20546f6b656e20646f65736e2774206578697374000000000000604482015290519081900360640190fd5b8082156108525761084d8186898988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061167492505050565b60006109077f00000000000000000000000000000000000000000000000000000000000000006001600160a01b038416611877565b326001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610d8e576040805162461bcd60e51b815260206004820152600d60248201526c27a7262cafa2aa242fa820a4a960991b604482015290519081900360640190fd5b6060610ddd87878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250604080516020810190915290815292506118a1915050565b90506060610e2e86868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250604080516020810190915290815292506118a1915050565b90506000610e7485858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250601292506119a8915050565b90506000610e828b83611893565b9050806001600160a01b03166347d5a09185856040518363ffffffff1660e01b8152600401808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610ee5578181015183820152602001610ecd565b50505050905090810190601f168015610f125780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610f45578181015183820152602001610f2d565b50505050905090810190601f168015610f725780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610f9357600080fd5b505af1158015610fa7573d6000803e3d6000fd5b505050505050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000081565b326001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611049576040805162461bcd60e51b815260206004820152600d60248201526c27a7262cafa2aa242fa820a4a960991b604482015290519081900360640190fd5b6001600160a01b03918216600090815260208190526040902080546001600160a01b03191691909216179055565b82611081816108d2565b6001600160a01b0316336001600160a01b031614806110b957506110a481610cec565b6001600160a01b0316336001600160a01b0316145b806110dd57506001600160a01b038181166000908152602081905260409020541633145b61111f576040805162461bcd60e51b815260206004820152600e60248201526d2727aa2fa32927a6afaa27a5a2a760911b604482015290519081900360640190fd5b600154604080516024808201939093526001600160a01b0387811660448084019190915287821660648481019190915260848085018990528551808603909101815260a490940185526020840180516001600160e01b031663e0a345fd60e01b17815285516349460b4d60e11b81527f0000000000000000000000000000000000000000000000000000000000000000948516600482019081529781019687528551938101939093528451600097929663928c169a96949092908801918083838c5b838110156111f95781810151838201526020016111e1565b50505050905090810190601f1680156112265780820380516001836020036101000a031916815260200191505b509350505050602060405180830381600087803b15801561124657600080fd5b505af115801561125a573d6000803e3d6000fd5b505050506040513d602081101561127057600080fd5b50516001805481019081905560408051838152602081019290925280519293506001600160a01b03808816938793918a16927f5f46021d14caa00be09900ec076911cd350172076472afa1697c672b8035c17c928290030190a45050505050565b333014611325576040805162461bcd60e51b815260206004820152601f60248201527f4d696e742063616e206f6e6c792062652063616c6c65642062792073656c6600604482015290519081900360640190fd5b604080516332dbacfd60e21b81526001600160a01b038481166004830152602482018790526060604483015260006064830181905292519088169263cb6eb3f49260a4808201939182900301818387803b15801561138257600080fd5b505af1158015611396573d6000803e3d6000fd5b505050506000826001600160a01b031663a4c0ed368587856040518463ffffffff1660e01b815260040180846001600160a01b03166001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015611418578181015183820152602001611400565b50505050905090810190601f1680156114455780820380516001836020036101000a031916815260200191505b50945050505050602060405180830381600087803b15801561146657600080fd5b505af115801561147a573d6000803e3d6000fd5b505050506040513d602081101561149057600080fd5b50519050806114d05760405162461bcd60e51b8152600401808060200182810382526021815260200180611f256021913960400191505060405180910390fd5b505050505050565b326001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614611545576040805162461bcd60e51b815260206004820152600d60248201526c27a7262cafa2aa242fa820a4a960991b604482015290519081900360640190fd5b606061159487878080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250604080516020810190915290815292506118a1915050565b905060606115e586868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250604080516020810190915290815292506118a1915050565b9050600061162b85858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250601292506119a8915050565b90506000610e828b8361165d565b7f000000000000000000000000000000000000000000000000000000000000000081565b600061166b838360006119cc565b90505b92915050565b604051636db867b760e11b81526001600160a01b0386811660048301908152602483018790528582166044840152908416606483015260a060848301908152835160a48401528351309363db70cf6e938a938a938a938a938a939092909160c40190602085019080838360005b838110156116f95781810151838201526020016116e1565b50505050905090810190601f1680156117265780820380516001836020036101000a031916815260200191505b509650505050505050600060405180830381600087803b15801561174957600080fd5b505af192505050801561175a575060015b61182357604080516332dbacfd60e21b81526001600160a01b038581166004830152602482018790526060604483015260006064830181905292519088169263cb6eb3f49260a4808201939182900301818387803b1580156117bb57600080fd5b505af11580156117cf573d6000803e3d6000fd5b505060408051600081526020810188905281516001600160a01b038088169550881693507f1ae40885bec8e271ecdd13e172c58836ad9dd6f4098020ebeea97d6a22e1f26c929181900390910190a3611870565b60408051600181526020810186905281516001600160a01b0380861693908716927f1ae40885bec8e271ecdd13e172c58836ad9dd6f4098020ebeea97d6a22e1f26c929081900390910190a35b5050505050565b600061166b6001600160a01b038416833063ffffffff611b4216565b600061166b838360016119cc565b60608251600014156118b457508061166e565b8251602014156118de576118d76118d284600063ffffffff611ba016565b611bf9565b905061166e565b8280602001905160208110156118f357600080fd5b8101908080516040519392919084600160201b82111561191257600080fd5b90830190602082018581111561192757600080fd5b8251600160201b81118282018810171561194057600080fd5b82525081516020918201929091019080838360005b8381101561196d578181015183820152602001611955565b50505050905090810190601f16801561199a5780820380516001836020036101000a031916815260200191505b50604052505050905061166e565b60008251600014156119bb57508061166e565b6118d783600063ffffffff611cc816565b6001600160a01b0380841660009081526020819052604081205490911680156119f6579050611b3b565b600080846001811115611a0557fe5b14611a1857611a13866108d2565b611a21565b611a2186610cec565b9050611a35816001600160a01b0316611d21565b611b37576000611aa781866001811115611a4b57fe5b14611a76577f0000000000000000000000000000000000000000000000000000000000000000611a98565b7f00000000000000000000000000000000000000000000000000000000000000005b6001600160a01b038916611d27565b9050816001600160a01b0316816001600160a01b031614611ac457fe5b604080516244919560e91b81523060048201526001600160a01b03898116602483015260ff891660448301529151918416916389232a009160648082019260009290919082900301818387803b158015611b1d57600080fd5b505af1158015611b31573d6000803e3d6000fd5b50505050505b9150505b9392505050565b604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b8152606093841b60148201526f5af43d82803e903d91602b57fd5bf3ff60801b6028820152921b6038830152604c8201526037808220606c830152605591012090565b60008160200183511015611bf0576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016020015190565b60408051818152606081810183529182919060208201818036833701905050905060005b6020811015611cc1576000848260208110611c3457fe5b1a60f881811b9250601080830480831b9360ff9091169091029003901b611c5a82611e54565b858560020281518110611c6957fe5b60200101906001600160f81b031916908160001a905350611c8981611e54565b858560020260010181518110611c9b57fe5b60200101906001600160f81b031916908160001a9053505060019092019150611c1d9050565b5092915050565b60008160010183511015611d18576040805162461bcd60e51b815260206004820152601260248201527152656164206f7574206f6620626f756e647360701b604482015290519081900360640190fd5b50016001015190565b3b151590565b6000826001600160a01b0316636f791d296040518163ffffffff1660e01b815260040160206040518083038186803b158015611d6257600080fd5b505afa158015611d76573d6000803e3d6000fd5b505050506040513d6020811015611d8c57600080fd5b505160408051808201909152600c81526b21a627a722afa6a0a9aa22a960a11b602082015290611e3a5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611dff578181015183820152602001611de7565b50505050905090810190601f168015611e2c5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5061166b6001600160a01b0384168363ffffffff611e8516565b6000600a60f883901c1015611e74578160f81c60300160f81b905061090a565b8160f81c60570160f81b905061090a565b6000604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b81528360601b60148201526e5af43d82803e903d91602b57fd5bf360881b6028820152826037826000f59150506001600160a01b03811661166e576040805162461bcd60e51b8152602060048201526017602482015276115490cc4c4d8dce8818dc99585d194c8819985a5b1959604a1b604482015290519081900360640190fdfe45787465726e616c206f6e546f6b656e5472616e73666572207265766572746564a26469706673582212204c4ecaf027c2ef03647fe894a6630e32b59aafc2b0e9b095c898b95e855dfa4a64736f6c634300060b0033'
