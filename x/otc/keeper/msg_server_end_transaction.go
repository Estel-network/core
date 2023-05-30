package keeper

import (
	"context"

	"igmf/x/otc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) EndTransaction(goCtx context.Context, msg *types.MsgEndTransaction) (*types.MsgEndTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	transaction, found := k.Keeper.GetTransactions(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "Position %d doesn't exist", msg.Id)
	}

	if transaction.Completed {
		return nil, nil // TODO: Transaction is closed
	}

	if transaction.Buyer != "Any" && transaction.Buyer != msg.Creator {
		return nil, nil // TODO: Create a not allowed error
	}

	buyer, _ := sdk.AccAddressFromBech32(msg.Creator)
	seller, _ := sdk.AccAddressFromBech32(transaction.Seller)

	denomBuyAmountunt, err := sdk.ParseCoinsNormalized(transaction.DenomBuyAmount)
	if err != nil {
		return nil, err
	}

	denomSellAmount, err := sdk.ParseCoinsNormalized(transaction.DenomSellAmount)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoins(ctx, buyer, seller, denomBuyAmountunt)
	if err != nil {
		return nil, err
	}

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, buyer, denomSellAmount)
	if err != nil {
		return nil, err // TODO: Need better checks so never gives errors
	}

	transaction.Completed = true
	k.Keeper.SetTransactions(ctx, transaction)

	return &types.MsgEndTransactionResponse{}, nil
}
