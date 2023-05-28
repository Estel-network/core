package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"igmf/x/credit/types"
)

func (k Keeper) PositionsAll(goCtx context.Context, req *types.QueryAllPositionsRequest) (*types.QueryAllPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var positionss []types.Positions
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	positionsStore := prefix.NewStore(store, types.KeyPrefix(types.PositionsKeyPrefix))

	pageRes, err := query.Paginate(positionsStore, req.Pagination, func(key []byte, value []byte) error {
		var positions types.Positions
		if err := k.cdc.Unmarshal(value, &positions); err != nil {
			return err
		}

		positionss = append(positionss, positions)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllPositionsResponse{Positions: positionss, Pagination: pageRes}, nil
}

func (k Keeper) Positions(goCtx context.Context, req *types.QueryGetPositionsRequest) (*types.QueryGetPositionsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPositions(
		ctx,
		req.Address,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetPositionsResponse{Positions: val}, nil
}
