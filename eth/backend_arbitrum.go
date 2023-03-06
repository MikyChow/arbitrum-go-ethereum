package eth

import (
	"github.com/MikyChow/arbitrum-go-ethereum/core"
	"github.com/MikyChow/arbitrum-go-ethereum/core/state"
	"github.com/MikyChow/arbitrum-go-ethereum/core/types"
	"github.com/MikyChow/arbitrum-go-ethereum/core/vm"
	"github.com/MikyChow/arbitrum-go-ethereum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, error) {
	return eth.stateAtTransaction(block, txIndex, reexec)
}
