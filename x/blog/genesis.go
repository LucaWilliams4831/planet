package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"planet/x/blog/keeper"
	"planet/x/blog/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the post
	for _, elem := range genState.PostList {
		k.SetPost(ctx, elem)
	}

	// Set post count
	k.SetPostCount(ctx, genState.PostCount)
	// Set all the sendPost
	for _, elem := range genState.SendPostList {
		k.SetSendPost(ctx, elem)
	}

	// Set sendPost count
	k.SetSendPostCount(ctx, genState.SendPostCount)
	// Set all the timedoutPost
	for _, elem := range genState.TimedoutPostList {
		k.SetTimedoutPost(ctx, elem)
	}

	// Set timedoutPost count
	k.SetTimedoutPostCount(ctx, genState.TimedoutPostCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.PostList = k.GetAllPost(ctx)
	genesis.PostCount = k.GetPostCount(ctx)
	genesis.SendPostList = k.GetAllSendPost(ctx)
	genesis.SendPostCount = k.GetSendPostCount(ctx)
	genesis.TimedoutPostList = k.GetAllTimedoutPost(ctx)
	genesis.TimedoutPostCount = k.GetTimedoutPostCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
