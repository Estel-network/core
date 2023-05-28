package keeper

import (
	"context"
	"strconv"

	"igmf/x/credit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleInfo, found := k.Keeper.GetModuleInfo(ctx)
	if !found {
		panic("moduleInfo not found")
	}

	amount, err := strconv.ParseUint(msg.Amount, 10, 64)
	if err != nil {
		panic("Invalid amount")
	}

	var credit = types.Credit{
		Index:             strconv.FormatUint(moduleInfo.TotalPositions+1, 10),
		Owner:             msg.Creator,
		Denom:             msg.Denom,
		DenomAmount:       amount,
		Credited:          0,
		Rewarding:         0,
		LastRewardedBlock: 0,
		Bond:              0,
	}
	moduleInfo.TotalPositions++
	k.Keeper.SetModuleInfo(ctx, moduleInfo)
	k.Keeper.SetCredit(ctx, credit)

	position, found := k.Keeper.GetPositions(ctx, msg.Creator)
	if !found {
		var newPosition = types.Positions{
			Address:   msg.Creator,
			CreditIDS: strconv.FormatUint(moduleInfo.TotalPositions, 10),
		}

		k.Keeper.SetPositions(ctx, newPosition)
	} else {
		position.CreditIDS += " " + strconv.FormatUint(moduleInfo.TotalPositions, 10)
		k.Keeper.SetPositions(ctx, position)
	}

	return &types.MsgDepositResponse{}, nil
}
