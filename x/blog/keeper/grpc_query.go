package keeper

import (
	"planet/x/blog/types"
)

var _ types.QueryServer = Keeper{}
