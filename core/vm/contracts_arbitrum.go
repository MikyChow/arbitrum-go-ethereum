package vm

import "github.com/MikyChow/arbitrum-go-ethereum/common"

var (
	PrecompiledContractsArbitrum = make(map[common.Address]PrecompiledContract)
	PrecompiledAddressesArbitrum []common.Address
)
