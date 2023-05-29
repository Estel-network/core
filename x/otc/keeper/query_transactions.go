package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"igmf/x/otc/types"
)

func (k Keeper) TransactionsAll(goCtx context.Context, req *types.QueryAllTransactionsRequest) (*types.QueryAllTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var transactionss []types.Transactions
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	transactionsStore := prefix.NewStore(store, types.KeyPrefix(types.TransactionsKeyPrefix))

	pageRes, err := query.Paginate(transactionsStore, req.Pagination, func(key []byte, value []byte) error {
		var transactions types.Transactions
		if err := k.cdc.Unmarshal(value, &transactions); err != nil {
			return err
		}

		transactionss = append(transactionss, transactions)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTransactionsResponse{Transactions: transactionss, Pagination: pageRes}, nil
}

func (k Keeper) Transactions(goCtx context.Context, req *types.QueryGetTransactionsRequest) (*types.QueryGetTransactionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetTransactions(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetTransactionsResponse{Transactions: val}, nil
}
