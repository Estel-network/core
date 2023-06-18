package keeper

import (
	"context"

	"core/x/credit/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) CollateralAll(goCtx context.Context, req *types.QueryAllCollateralRequest) (*types.QueryAllCollateralResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var collaterals []types.Collateral
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	collateralStore := prefix.NewStore(store, types.KeyPrefix(types.CollateralKeyPrefix))

	pageRes, err := query.Paginate(collateralStore, req.Pagination, func(key []byte, value []byte) error {
		var collateral types.Collateral
		if err := k.cdc.Unmarshal(value, &collateral); err != nil {
			return err
		}

		collaterals = append(collaterals, collateral)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCollateralResponse{Collateral: collaterals, Pagination: pageRes}, nil
}

func (k Keeper) Collateral(goCtx context.Context, req *types.QueryGetCollateralRequest) (*types.QueryGetCollateralResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetCollateral(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetCollateralResponse{Collateral: val}, nil
}
