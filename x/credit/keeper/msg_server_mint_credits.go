package keeper

import (
	"context"
	"core/x/credit/types"
	"strconv"

	sdkmath "cosmossdk.io/math"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) MintCredits(goCtx context.Context, msg *types.MsgMintCredits) (*types.MsgMintCreditsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleInfo, found := k.Keeper.GetModuleInfo(ctx)
	if !found {
		panic("moduleInfo not found")
	}

	if moduleInfo.Enabled == false {
		return &types.MsgMintCreditsResponse{Response: "Module is disabled"}, nil
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	credit, found := k.Keeper.GetCredit(ctx, msg.Creator)
	if !found {
		return &types.MsgMintCreditsResponse{Response: "Inexistent position"}, nil
	}

	remainingValue := uint64(0)

	for i := 0; i < len(credit.Collateral); i++ {
		colateral, found := k.Keeper.GetCollateral(ctx, strconv.FormatUint(credit.Collateral[i], 10))
		if !found {
			panic("Unexistent position")
		}

		if colateral.Owner != msg.Creator {
			return &types.MsgMintCreditsResponse{Response: "Error"}, nil
		}

		denomValue, err := k.oracleKeeper.GetCoinPrice(colateral.Denom)
		if err != nil {
			return &types.MsgMintCreditsResponse{}, err
		}

		remainingValue += denomValue * colateral.Amount
	}

	remainingValue = (remainingValue * moduleInfo.LiquidationThreshold) / 100

	fee := (msg.Amount * moduleInfo.CreditFee) / 100

	remainingValue -= credit.Credited * moduleInfo.TargetPrice

	if (moduleInfo.TargetPrice * (msg.Amount + fee)) > remainingValue {
		return &types.MsgMintCreditsResponse{Response: "Not enought colateral"}, nil
	}

	coins := sdk.NewCoin("ccrd", sdkmath.NewInt(int64(msg.Amount)))

	k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coins))

	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, creator, sdk.NewCoins(coins))

	moduleInfo.TotalCredited += msg.Amount
	credit.Credited += msg.Amount + fee
	k.Keeper.SetModuleInfo(ctx, moduleInfo)
	k.Keeper.SetCredit(ctx, credit)

	return &types.MsgMintCreditsResponse{}, nil
}
