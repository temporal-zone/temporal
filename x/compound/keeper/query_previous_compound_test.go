package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/temporal-zone/temporal/testutil/keeper"
	"github.com/temporal-zone/temporal/testutil/nullify"
	"github.com/temporal-zone/temporal/x/compound/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestPreviousCompoundQuerySingle(t *testing.T) {
	compKeeper, ctx := keepertest.CompoundKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPreviousCompound(compKeeper, ctx, 2)
	tests := []struct {
		desc     string
		request  *types.QueryGetPreviousCompoundRequest
		response *types.QueryGetPreviousCompoundResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetPreviousCompoundRequest{
				Delegator: msgs[0].Delegator,
			},
			response: &types.QueryGetPreviousCompoundResponse{PreviousCompound: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetPreviousCompoundRequest{
				Delegator: msgs[1].Delegator,
			},
			response: &types.QueryGetPreviousCompoundResponse{PreviousCompound: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetPreviousCompoundRequest{
				Delegator: strconv.Itoa(100000),
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
			response, err := compKeeper.PreviousCompound(wctx, tc.request)
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

func TestPreviousCompoundQueryPaginated(t *testing.T) {
	compKeeper, ctx := keepertest.CompoundKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNPreviousCompound(compKeeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllPreviousCompoundRequest {
		return &types.QueryAllPreviousCompoundRequest{
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
			resp, err := compKeeper.PreviousCompoundAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PreviousCompound), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PreviousCompound),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := compKeeper.PreviousCompoundAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.PreviousCompound), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.PreviousCompound),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := compKeeper.PreviousCompoundAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.PreviousCompound),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := compKeeper.PreviousCompoundAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
