package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "igmf/testutil/keeper"
	"igmf/testutil/nullify"
	"igmf/x/otc/types"
)

func TestModuleInfoQuery(t *testing.T) {
	keeper, ctx := keepertest.OtcKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestModuleInfo(keeper, ctx)
	for _, tc := range []struct {
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
	} {
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
