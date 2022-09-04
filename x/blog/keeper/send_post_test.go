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

func createNSendPost(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SendPost {
	items := make([]types.SendPost, n)
	for i := range items {
		items[i].Id = keeper.AppendSendPost(ctx, items[i])
	}
	return items
}

func TestSendPostGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSendPost(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSendPost(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSendPostRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSendPost(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSendPost(ctx, item.Id)
		_, found := keeper.GetSendPost(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSendPostGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSendPost(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSendPost(ctx)),
	)
}

func TestSendPostCount(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNSendPost(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSendPostCount(ctx))
}
