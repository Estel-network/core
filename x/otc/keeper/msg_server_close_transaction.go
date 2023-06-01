package keeper

import (
	"context"

	"igmf/x/otc/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CloseTransaction(goCtx context.Context, msg *types.MsgCloseTransaction) (*types.MsgCloseTransactionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	transaction, found := k.Keeper.GetTransactions(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrKeyNotFound, "Position %d doesn't exist", msg.Id)
	}

	if transaction.Seller != msg.Creator {
		return nil, nil // TODO: Not the owner
	}

	if transaction.Completed {
		return nil, nil // TODO: Transaction is closed
	}

	denomSellAmount, err := sdk.ParseCoinsNormalized(transaction.DenomSellAmount)
	if err != nil {
		return nil, err
	}

	seller, _ := sdk.AccAddressFromBech32(transaction.Seller)

	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, seller, denomSellAmount)
	if err != nil {
		return nil, err // TODO: Need better checks so never gives errors
	}

	transaction.Completed = true
	k.Keeper.SetTransactions(ctx, transaction)

	return &types.MsgCloseTransactionResponse{}, nil
}
