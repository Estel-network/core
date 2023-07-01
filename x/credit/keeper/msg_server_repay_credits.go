package keeper

import (
	"context"

	"core/x/credit/types"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) RepayCredits(goCtx context.Context, msg *types.MsgRepayCredits) (*types.MsgRepayCreditsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleInfo, found := k.Keeper.GetModuleInfo(ctx)
	if !found {
		panic("moduleInfo not found")
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	credit, found := k.Keeper.GetCredit(ctx, msg.Creator)
	if !found {
		return &types.MsgRepayCreditsResponse{Response: "Inexistent position"}, nil
	}

	if msg.Amount > credit.Credited {
		msg.Amount = credit.Credited
	}

	coins := sdk.NewCoin("ccrd", sdkmath.NewInt(int64(msg.Amount)))

	k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, sdk.NewCoins(coins))
	k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.NewCoins(coins))

	credit.Credited -= msg.Amount
	moduleInfo.TotalCredited -= msg.Amount
	k.Keeper.SetModuleInfo(ctx, moduleInfo)
	k.Keeper.SetCredit(ctx, credit)

	return &types.MsgRepayCreditsResponse{}, nil
}
