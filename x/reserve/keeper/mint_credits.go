package keeper

import (
	"igmf/x/reserve/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) MintCredits(ctx sdk.Context, receiver sdk.AccAddress, amount int) error {
	// Mint the tokens
	credit := sdk.NewCoin("CCRD", sdkmath.NewInt(int64(amount)))
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(credit))
	if err != nil {
		return err
	}
	// Send the tokens
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, receiver, sdk.NewCoins(credit))

	return nil
}
