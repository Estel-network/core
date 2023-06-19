package keeper

import (
	"context"
	"strconv"

	"core/x/credit/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Deposit(goCtx context.Context, msg *types.MsgDeposit) (*types.MsgDepositResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleInfo, found := k.Keeper.GetModuleInfo(ctx)
	if !found {
		panic("moduleInfo not found")
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	amount, err := strconv.ParseUint(msg.Amount, 10, 64)
	if err != nil {
		panic("Invalid amount")
	}

	denomAmount, err := sdk.ParseCoinsNormalized(msg.Amount + msg.Denom)
	if err != nil {
		return nil, err
	}

	credit, found := k.Keeper.GetCredit(ctx, msg.Creator)
	if !found {
		var newCredit = types.Credit{
			Owner:      msg.Creator,
			Credited:   0,
			Rewarding:  0,
			Collateral: []uint64{moduleInfo.TotalPositions + 1},
			Bond:       0,
		}

		var newColateral = types.Collateral{
			Index:  strconv.FormatUint(moduleInfo.TotalPositions+1, 10),
			Owner:  msg.Creator,
			Denom:  msg.Denom,
			Amount: amount,
		}

		sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, denomAmount)
		if sdkError != nil {
			return nil, sdkError
		}

		k.Keeper.SetCredit(ctx, newCredit)
		k.Keeper.SetCollateral(ctx, newColateral)
		moduleInfo.TotalPositions++
		k.Keeper.SetModuleInfo(ctx, moduleInfo)
	} else {
		indexed := false
		for i := 0; i < len(credit.Collateral); i++ {
			colateral, found := k.Keeper.GetCollateral(ctx, strconv.FormatUint(credit.Collateral[i], 10))
			if !found {
				panic("Unexistent position")
			}

			if colateral.Denom == msg.Denom {
				colateral.Amount += amount
				indexed = true

				sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, denomAmount)
				if sdkError != nil {
					return nil, sdkError
				}
				k.Keeper.SetCollateral(ctx, colateral)
			}
		}

		if !indexed {
			var newColateral = types.Collateral{
				Index:  strconv.FormatUint(moduleInfo.TotalPositions+1, 10),
				Owner:  msg.Creator,
				Denom:  msg.Denom,
				Amount: amount,
			}

			credit.Collateral = append(credit.Collateral, moduleInfo.TotalPositions+1)

			sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, denomAmount)
			if sdkError != nil {
				return nil, sdkError
			}

			k.Keeper.SetCredit(ctx, credit)
			k.Keeper.SetCollateral(ctx, newColateral)
			moduleInfo.TotalPositions++
			k.Keeper.SetModuleInfo(ctx, moduleInfo)
		}

	}

	return &types.MsgDepositResponse{}, nil
}
