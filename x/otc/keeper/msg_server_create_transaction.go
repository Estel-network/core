package keeper

import (
	"context"
	"strconv"

	"igmf/x/otc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateTransaction(goCtx context.Context, msg *types.MsgCreateTransaction) (*types.MsgCreateTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	moduleInfo, found := k.Keeper.GetModuleInfo(ctx)
	if !found {
		panic("moduleInfo not found")
	}

	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}

	buyer := msg.Buyer

	x, err := sdk.AccAddressFromBech32(buyer)
	_ = x // Why? Why? Why?!
	if err != nil {
		buyer = "Any" // if no buyer is set or not valid anyone can buy
	}

	denomSellAmount, err := sdk.ParseCoinsNormalized(msg.DenomSellAmount)
	if err != nil {
		panic(err)
	}

	denomBuyAmount, err := sdk.ParseCoinsNormalized(msg.DenomBuyAmount)
	if err != nil {
		panic(err)
	}

	_ = denomBuyAmount

	var transaction = types.Transactions{
		Index:           strconv.FormatUint(moduleInfo.ModuleIndex, 10),
		Seller:          msg.Creator,
		Buyer:           buyer,
		DenomSellAmount: msg.DenomSellAmount,
		DenomBuyAmount:  msg.DenomBuyAmount,
		Completed:       false,
	}

	moduleInfo.ModuleIndex++
	k.Keeper.SetModuleInfo(ctx, moduleInfo)

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creator, types.ModuleName, denomSellAmount)
	if sdkError != nil {
		return nil, sdkError
	}

	k.Keeper.SetTransactions(ctx, transaction)

	return &types.MsgCreateTransactionResponse{}, nil
}
