package keeper

import (
	"context"

	"core/x/credit/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ModuleInfo(goCtx context.Context, req *types.QueryGetModuleInfoRequest) (*types.QueryGetModuleInfoResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetModuleInfo(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryGetModuleInfoResponse{ModuleInfo: val}, nil
}
