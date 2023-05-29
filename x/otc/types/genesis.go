package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		TransactionsList: []Transactions{},
		ModuleInfo: ModuleInfo{
			ModuleIndex: uint64(DefaultIndex),
			ServiceFee:  int32(0),
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in transactions
	transactionsIndexMap := make(map[string]struct{})

	for _, elem := range gs.TransactionsList {
		index := string(TransactionsKey(elem.Index))
		if _, ok := transactionsIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for transactions")
		}
		transactionsIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
