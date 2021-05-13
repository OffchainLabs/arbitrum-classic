package web3

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

type Accounts struct {
	srv         *Server
	addresses   []common.Address
	privateKeys map[common.Address]*ecdsa.PrivateKey
	signer      types.Signer
}

func NewAccounts(ethServer *Server, privateKeys []*ecdsa.PrivateKey) *Accounts {
	keys := make(map[common.Address]*ecdsa.PrivateKey)
	addresses := make([]common.Address, 0, len(privateKeys))
	for _, privKey := range privateKeys {
		addr := crypto.PubkeyToAddress(privKey.PublicKey)
		keys[addr] = privKey
		addresses = append(addresses, addr)
	}
	return &Accounts{
		srv:         ethServer,
		addresses:   addresses,
		privateKeys: keys,
		signer:      types.NewEIP155Signer(new(big.Int).SetUint64(uint64(ethServer.ChainId()))),
	}
}

func (s *Accounts) Accounts() []common.Address {
	return s.addresses
}

type SendTransactionArgs struct {
	From     *common.Address `json:"from"`
	To       *common.Address `json:"to"`
	Gas      *hexutil.Uint64 `json:"gas"`
	GasPrice *hexutil.Big    `json:"gasPrice"`
	Value    *hexutil.Big    `json:"value"`
	Nonce    *hexutil.Uint64 `json:"nonce"`
	Data     *hexutil.Bytes  `json:"data"`
}

func (s *Accounts) SendTransaction(ctx context.Context, args *SendTransactionArgs) (common.Hash, error) {
	sender := s.addresses[0]
	if args.From != nil {
		sender = *args.From
	}
	privKey, ok := s.privateKeys[sender]
	if !ok {
		return common.Hash{}, errors.New("sender does not have unlocked wallet")
	}

	var nonce uint64
	if args.Nonce != nil {
		nonce = uint64(*args.Nonce)
	} else {
		pendingNum := rpc.PendingBlockNumber
		rawNonce, err := s.srv.GetTransactionCount(ctx, &sender, &pendingNum)
		if err != nil {
			return common.Hash{}, err
		}
		nonce = uint64(rawNonce)
	}
	gas := uint64(2000000)
	if args.Gas != nil {
		gas = uint64(*args.Gas)
	}
	val := (*big.Int)(args.Value)
	if val == nil {
		val = big.NewInt(0)
	}
	var data []byte
	if args.Data != nil {
		data = *args.Data
	}
	gasPrice := (*big.Int)(args.GasPrice)
	if gasPrice == nil {
		gasPriceRaw, err := s.srv.GasPrice()
		if err != nil {
			return [32]byte{}, err
		}
		gasPrice = (*big.Int)(gasPriceRaw)
	}
	var tx *types.Transaction
	if args.To != nil {
		tx = types.NewTransaction(
			nonce,
			*args.To,
			val,
			gas,
			gasPrice,
			data,
		)
	} else {
		tx = types.NewContractCreation(
			nonce,
			val,
			gas,
			gasPrice,
			data,
		)
	}
	signedTx, err := types.SignTx(tx, s.signer, privKey)
	if err != nil {
		return [32]byte{}, err
	}

	if err := s.srv.srv.SendTransaction(ctx, signedTx); err != nil {
		return [32]byte{}, err
	}
	return signedTx.Hash(), nil
}

func (s *Accounts) Sign(account common.Address, data hexutil.Bytes) (hexutil.Bytes, error) {
	privKey, ok := s.privateKeys[account]
	if !ok {
		return nil, errors.New("signer does not have unlocked wallet")
	}
	sig, err := crypto.Sign(accounts.TextHash(data), privKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27
	return sig, nil
}

type PersonalAccounts struct {
	privateKeys map[common.Address]*ecdsa.PrivateKey
}

func NewPersonalAccounts(privateKeys []*ecdsa.PrivateKey) *PersonalAccounts {
	keys := make(map[common.Address]*ecdsa.PrivateKey)
	for _, privKey := range privateKeys {
		addr := crypto.PubkeyToAddress(privKey.PublicKey)
		keys[addr] = privKey
	}
	return &PersonalAccounts{
		privateKeys: keys,
	}
}

func (s *PersonalAccounts) Sign(data hexutil.Bytes, account common.Address, _ *hexutil.Bytes) (hexutil.Bytes, error) {
	// Password ignored
	privKey, ok := s.privateKeys[account]
	if !ok {
		return nil, errors.New("signer does not have unlocked wallet")
	}
	sig, err := crypto.Sign(accounts.TextHash(data), privKey)
	if err != nil {
		return nil, err
	}
	sig[64] += 27
	return sig, nil
}
