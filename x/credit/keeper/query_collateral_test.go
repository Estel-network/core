package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "core/testutil/keeper"
	"core/testutil/nullify"
	"core/x/credit/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestCollateralQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCollateral(keeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetCollateralRequest
		response *types.QueryGetCollateralResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetCollateralRequest{
				Index: msgs[0].Index,
			},
			response: &types.QueryGetCollateralResponse{Collateral: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetCollateralRequest{
				Index: msgs[1].Index,
			},
			response: &types.QueryGetCollateralResponse{Collateral: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetCollateralRequest{
				Index: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	}
	for _, tc := range tests {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Collateral(wctx, tc.request)
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

func TestCollateralQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.CreditKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNCollateral(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllCollateralRequest {
		return &types.QueryAllCollateralRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CollateralAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Collateral), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Collateral),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.CollateralAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Collateral), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Collateral),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.CollateralAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Collateral),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.CollateralAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
