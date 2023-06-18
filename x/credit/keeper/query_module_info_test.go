package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "core/testutil/keeper"
	"core/testutil/nullify"
	"core/x/credit/types"
)

func TestModuleInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestModuleInfo(keeper, ctx)
	tests := []struct {
		desc     string
		request  *types.QueryGetModuleInfoRequest
		response *types.QueryGetModuleInfoResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetModuleInfoRequest{},
			response: &types.QueryGetModuleInfoResponse{ModuleInfo: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.ModuleInfo(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
