package core

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var pauserABI string = "[{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

type Pauser struct {
	ABI     abi.ABI
	backend bind.ContractBackend
}

func NewPauser(backend bind.ContractBackend) (*Pauser, error) {
	return &Pauser{
		ABI:     *GetAbi("Pauser"),
		backend: backend,
	}, nil
}

func (p *Pauser) IsPaused(opts *bind.CallOpts, address common.Address) (bool, error) {
	var out []interface{}
	err := bind.NewBoundContract(address, p.ABI, p.backend, nil, nil).Call(opts, &out, "paused")
	if err != nil {
		return *new(bool), err
	}
	return out[0].(bool), err
}

func (p *Pauser) Pause(opts *bind.TransactOpts, address common.Address) (*types.Transaction, error) {
	return bind.NewBoundContract(address, p.ABI, p.backend, p.backend, nil).Transact(opts, "pause")
}

func (p *Pauser) Unpause(opts *bind.TransactOpts, address common.Address) (*types.Transaction, error) {
	return bind.NewBoundContract(address, p.ABI, p.backend, p.backend, nil).Transact(opts, "unpause")
}
