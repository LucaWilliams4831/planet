package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "planet/testutil/keeper"
	"planet/testutil/nullify"
	"planet/x/blog/keeper"
	"planet/x/blog/types"
)

func createNTimedoutPost(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TimedoutPost {
	items := make([]types.TimedoutPost, n)
	for i := range items {
		items[i].Id = keeper.AppendTimedoutPost(ctx, items[i])
	}
	return items
}

func TestTimedoutPostGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNTimedoutPost(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTimedoutPost(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTimedoutPostRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNTimedoutPost(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTimedoutPost(ctx, item.Id)
		_, found := keeper.GetTimedoutPost(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTimedoutPostGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNTimedoutPost(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTimedoutPost(ctx)),
	)
}

func TestTimedoutPostCount(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNTimedoutPost(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTimedoutPostCount(ctx))
}
