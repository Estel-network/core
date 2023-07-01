package types

import (
	"fmt"
)

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		ModuleInfo: ModuleInfo{
			Enabled:              true,
			TargetPrice:          1 * 1000000000,
			TotalPositions:       0,
			TotalCredited:        0,
			CreditFee:            0,
			RewardAmount:         0,
			RewardTime:           1400,
			LiquidationThreshold: 90,
			Liquidable:           false,
		},
		CreditList:     []Credit{},
		CollateralList: []Collateral{},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in credit
	creditIndexMap := make(map[string]struct{})

	for _, elem := range gs.CreditList {
		index := string(CreditKey(elem.Owner))
		if _, ok := creditIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for credit")
		}
		creditIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in collateral
	collateralIndexMap := make(map[string]struct{})

	for _, elem := range gs.CollateralList {
		index := string(CollateralKey(elem.Index))
		if _, ok := collateralIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for collateral")
		}
		collateralIndexMap[index] = struct{}{}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}
